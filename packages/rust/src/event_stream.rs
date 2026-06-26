use std::collections::HashMap;
use std::sync::atomic::{AtomicBool, Ordering};
use std::sync::Arc;
use std::time::Duration;

use futures_util::{SinkExt, StreamExt};
use tokio::sync::{mpsc, Mutex, Notify, RwLock};
use tokio_tungstenite::tungstenite::Message;

use crate::events::{decode_event, Event};

pub const DEFAULT_EVENT_STREAM_URL: &str = "wss://cable.yay.space";

const DEFAULT_SUBSCRIBE_TIMEOUT: Duration = Duration::from_secs(10);
const DEFAULT_EVENT_BUFFER: usize = 64;
const RECONNECT_STABLE_THRESHOLD: Duration = Duration::from_secs(30);

/// Configuration for automatic reconnection of the event stream.
#[derive(Debug, Clone)]
pub struct ReconnectPolicy {
    pub disabled: bool,
    pub max_attempts: usize,
    pub initial_delay: Duration,
    pub max_delay: Duration,
}

impl Default for ReconnectPolicy {
    fn default() -> Self {
        Self {
            disabled: false,
            max_attempts: 0,
            initial_delay: Duration::from_millis(500),
            max_delay: Duration::from_secs(30),
        }
    }
}

#[derive(Clone)]
pub struct EventStreamOptions {
    pub reconnect: ReconnectPolicy,
    pub event_buffer: usize,
    pub subscribe_timeout: Duration,
    pub on_drop: Option<Arc<dyn Fn(Event) + Send + Sync>>,
}

impl Default for EventStreamOptions {
    fn default() -> Self {
        Self {
            reconnect: ReconnectPolicy::default(),
            event_buffer: DEFAULT_EVENT_BUFFER,
            subscribe_timeout: DEFAULT_SUBSCRIBE_TIMEOUT,
            on_drop: None,
        }
    }
}

/// Event stream channel types for subscribing to real-time updates.
#[derive(Debug, Clone)]
pub enum Channel {
    ChatRoom,
    Messages { chat_room_id: i64 },
    GroupUpdates,
    GroupPosts { group_id: i64 },
}

impl Channel {
    pub fn identifier(&self) -> String {
        match self {
            Channel::ChatRoom => r#"{"channel":"ChatRoomChannel"}"#.to_string(),
            Channel::Messages { chat_room_id } => {
                format!(
                    r#"{{"channel":"MessagesChannel", "chat_room_id":"{}"}}"#,
                    chat_room_id
                )
            }
            Channel::GroupUpdates => r#"{"channel":"GroupUpdatesChannel"}"#.to_string(),
            Channel::GroupPosts { group_id } => {
                format!(
                    r#"{{"channel":"GroupPostsChannel", "group_id":"{}"}}"#,
                    group_id
                )
            }
        }
    }
}

pub fn chat_room_channel() -> Channel {
    Channel::ChatRoom
}

pub fn messages_channel(room_id: i64) -> Channel {
    Channel::Messages {
        chat_room_id: room_id,
    }
}

pub fn group_updates_channel() -> Channel {
    Channel::GroupUpdates
}

pub fn group_posts_channel(group_id: i64) -> Channel {
    Channel::GroupPosts { group_id }
}

#[allow(dead_code)]
struct SubEntry {
    ident: String,
    tx: mpsc::Sender<Event>,
    confirmed: Arc<Notify>,
    confirm_err: Arc<Mutex<Option<String>>>,
    closed: Arc<AtomicBool>,
}

/// A subscription to events on a specific channel.
#[allow(dead_code)]
pub struct Subscription {
    ident: String,
    rx: mpsc::Receiver<Event>,
    confirmed: Arc<Notify>,
    confirm_err: Arc<Mutex<Option<String>>>,
    closed: Arc<AtomicBool>,
    stream: Arc<EventStreamInner>,
}

impl Subscription {
    /// Receives the next event, or `None` when the subscription is closed.
    pub async fn recv(&mut self) -> Option<Event> {
        self.rx.recv().await
    }

    pub async fn unsubscribe(&self) {
        if self
            .closed
            .compare_exchange(false, true, Ordering::SeqCst, Ordering::SeqCst)
            .is_err()
        {
            return;
        }
        self.stream.remove_sub(&self.ident).await;
        let _ = self.stream.write_command("unsubscribe", &self.ident).await;
    }

    pub fn is_closed(&self) -> bool {
        self.closed.load(Ordering::SeqCst)
    }
}

#[derive(Debug, serde::Deserialize)]
struct WireFrame {
    #[serde(default)]
    identifier: Option<String>,
    #[serde(default, rename = "type")]
    frame_type: Option<String>,
    #[serde(default)]
    message: Option<serde_json::Value>,
}

#[derive(Debug, serde::Deserialize)]
struct EventPayload {
    #[serde(default)]
    event: Option<String>,
    #[serde(default)]
    data: Option<serde_json::Value>,
    #[serde(default)]
    message: Option<serde_json::Value>,
}

type WsSink = futures_util::stream::SplitSink<
    tokio_tungstenite::WebSocketStream<
        tokio_tungstenite::MaybeTlsStream<tokio::net::TcpStream>,
    >,
    Message,
>;

type WsStream = futures_util::stream::SplitStream<
    tokio_tungstenite::WebSocketStream<
        tokio_tungstenite::MaybeTlsStream<tokio::net::TcpStream>,
    >,
>;

struct EventStreamInner {
    base_url: String,
    api_version_name: String,
    opts: EventStreamOptions,

    ws_sink: RwLock<Option<WsSink>>,

    subs: RwLock<HashMap<String, SubEntry>>,

    closed: AtomicBool,
    done: Notify,
    final_err: Mutex<Option<String>>,

    http: reqwest::Client,
    api_base_url: String,
    access_token: RwLock<Option<String>>,
}

impl EventStreamInner {
    async fn write_command(&self, cmd: &str, identifier: &str) -> Result<(), String> {
        let msg = serde_json::json!({
            "command": cmd,
            "identifier": identifier,
        });
        let text = serde_json::to_string(&msg).map_err(|e| e.to_string())?;

        let mut guard = self.ws_sink.write().await;
        if let Some(sink) = guard.as_mut() {
            sink.send(Message::Text(text))
                .await
                .map_err(|e| e.to_string())
        } else {
            Err("ws not connected".into())
        }
    }

    async fn remove_sub(&self, ident: &str) {
        let mut subs = self.subs.write().await;
        subs.remove(ident);
    }

    async fn set_final_err(&self, err: String) {
        let mut guard = self.final_err.lock().await;
        *guard = Some(err);
    }

    async fn get_final_err(&self) -> Option<String> {
        let guard = self.final_err.lock().await;
        guard.clone()
    }

    async fn terminate_all_subs(&self, err: &str) {
        let entries: Vec<SubEntry> = {
            let mut guard = self.subs.write().await;
            guard.drain().map(|(_, v)| v).collect()
        };
        for entry in entries {
            entry.closed.store(true, Ordering::SeqCst);
            {
                let mut cerr = entry.confirm_err.lock().await;
                if cerr.is_none() {
                    *cerr = Some(err.to_string());
                }
            }
            entry.confirmed.notify_waiters();
        }
    }

    async fn resubscribe_all(&self) {
        let idents: Vec<String> = {
            let subs = self.subs.read().await;
            subs.keys().cloned().collect()
        };
        for ident in idents {
            let _ = self.write_command("subscribe", &ident).await;
        }
    }

    async fn dispatch(&self, data: &[u8]) {
        let frame: WireFrame = match serde_json::from_slice(data) {
            Ok(f) => f,
            Err(_) => return,
        };

        match frame.frame_type.as_deref() {
            Some("welcome") | Some("ping") => return,
            Some("disconnect") => {
                let mut guard = self.ws_sink.write().await;
                if let Some(sink) = guard.as_mut() {
                    let _ = sink.close().await;
                }
                return;
            }
            Some("confirm_subscription") => {
                if let Some(ident) = &frame.identifier {
                    let subs = self.subs.read().await;
                    if let Some(entry) = subs.get(ident.as_str()) {
                        entry.confirmed.notify_waiters();
                    }
                }
                return;
            }
            Some("reject_subscription") => {
                if let Some(ident) = &frame.identifier {
                    let subs = self.subs.read().await;
                    if let Some(entry) = subs.get(ident.as_str()) {
                        {
                            let mut cerr = entry.confirm_err.lock().await;
                            *cerr = Some("subscription rejected".into());
                        }
                        entry.confirmed.notify_waiters();
                        entry.closed.store(true, Ordering::SeqCst);
                    }
                }
                return;
            }
            _ => {}
        }

        let msg = match frame.message {
            Some(m) => m,
            None => return,
        };

        let payload: EventPayload = match serde_json::from_value(msg) {
            Ok(p) => p,
            Err(_) => return,
        };

        let event_name = match payload.event.as_deref() {
            Some(name) if !name.is_empty() => name,
            _ => return,
        };

        let body = payload.data.or(payload.message);
        let body = match body {
            Some(v) => v,
            None => return,
        };

        let body_bytes = serde_json::to_vec(&body).unwrap_or_default();
        if body_bytes.is_empty() {
            return;
        }

        let ev = match decode_event(event_name, &body_bytes) {
            Some(ev) => ev,
            None => return,
        };

        if let Some(ident) = &frame.identifier {
            let subs = self.subs.read().await;
            if let Some(entry) = subs.get(ident.as_str()) {
                if entry.closed.load(Ordering::SeqCst) {
                    return;
                }
                match entry.tx.try_send(ev) {
                    Ok(()) => {}
                    Err(mpsc::error::TrySendError::Full(dropped)) => {
                        if let Some(ref cb) = self.opts.on_drop {
                            cb(dropped);
                        }
                    }
                    Err(mpsc::error::TrySendError::Closed(_)) => {}
                }
            }
        }
    }

    async fn fetch_ws_token(&self) -> Result<String, String> {
        let url = format!(
            "{}/v1/users/ws_token",
            self.api_base_url.trim_end_matches('/')
        );
        let mut req = self.http.get(&url);
        {
            let token = self.access_token.read().await;
            if let Some(t) = token.as_deref() {
                req = req.bearer_auth(t);
            }
        }
        let resp = req
            .send()
            .await
            .map_err(|e| format!("fetch ws token: {}", e))?;
        if !resp.status().is_success() {
            return Err(format!("fetch ws token: HTTP {}", resp.status()));
        }
        #[derive(serde::Deserialize)]
        struct TokenResp {
            #[serde(default)]
            token: String,
        }
        let parsed: TokenResp = resp
            .json()
            .await
            .map_err(|e| format!("parse ws token: {}", e))?;
        Ok(parsed.token)
    }
}

/// Real-time event stream over WebSocket with automatic reconnection.
pub struct EventStream {
    inner: Arc<EventStreamInner>,
    _run_handle: tokio::task::JoinHandle<()>,
}

impl EventStream {
    /// Subscribes to events on the given channel.
    pub async fn subscribe(&self, ch: Channel) -> Result<Subscription, String> {
        if self.inner.closed.load(Ordering::SeqCst) {
            return Err("event stream closed".into());
        }

        let ident = ch.identifier();
        let buffer = self.inner.opts.event_buffer;
        let timeout = self.inner.opts.subscribe_timeout;

        let (tx, rx) = mpsc::channel(buffer);
        let confirmed = Arc::new(Notify::new());
        let confirm_err: Arc<Mutex<Option<String>>> = Arc::new(Mutex::new(None));
        let closed = Arc::new(AtomicBool::new(false));

        {
            let mut subs = self.inner.subs.write().await;
            subs.insert(
                ident.clone(),
                SubEntry {
                    ident: ident.clone(),
                    tx,
                    confirmed: confirmed.clone(),
                    confirm_err: confirm_err.clone(),
                    closed: closed.clone(),
                },
            );
        }

        if let Err(e) = self.inner.write_command("subscribe", &ident).await {
            self.inner.remove_sub(&ident).await;
            return Err(format!("subscribe {:?}: {}", ident, e));
        }

        tokio::select! {
            _ = confirmed.notified() => {
                let cerr = confirm_err.lock().await;
                if let Some(err) = cerr.as_ref() {
                    self.inner.remove_sub(&ident).await;
                    return Err(err.clone());
                }
            }
            _ = tokio::time::sleep(timeout) => {
                self.inner.remove_sub(&ident).await;
                return Err(format!("subscribe {:?}: timed out after {:?}", ident, timeout));
            }
        }

        Ok(Subscription {
            ident,
            rx,
            confirmed,
            confirm_err,
            closed,
            stream: self.inner.clone(),
        })
    }

    pub async fn close(&self) {
        if self
            .inner
            .closed
            .compare_exchange(false, true, Ordering::SeqCst, Ordering::SeqCst)
            .is_err()
        {
            return;
        }
        {
            let mut guard = self.inner.ws_sink.write().await;
            if let Some(sink) = guard.as_mut() {
                let _ = sink.close().await;
            }
        }
        self.inner.done.notify_waiters();
    }

    pub async fn err(&self) -> Option<String> {
        if !self.inner.closed.load(Ordering::SeqCst) {
            return None;
        }
        self.inner.get_final_err().await
    }

    pub fn is_closed(&self) -> bool {
        self.inner.closed.load(Ordering::SeqCst)
    }
}

impl Drop for EventStream {
    fn drop(&mut self) {
        self.inner.closed.store(true, Ordering::SeqCst);
        self._run_handle.abort();
    }
}

async fn do_connect(
    inner: &EventStreamInner,
) -> Result<(WsSink, WsStream), String> {
    let token = inner.fetch_ws_token().await?;

    let mut url =
        url::Url::parse(&inner.base_url).map_err(|e| format!("parse ws url: {}", e))?;
    url.query_pairs_mut()
        .append_pair("token", &token)
        .append_pair("app_version", &inner.api_version_name);

    let (ws, _resp) = tokio_tungstenite::connect_async(url.as_str())
        .await
        .map_err(|e| format!("dial: {}", e))?;

    let (sink, mut stream) = ws.split();

    loop {
        match stream.next().await {
            Some(Ok(Message::Text(text))) => {
                if let Ok(frame) = serde_json::from_str::<WireFrame>(&text) {
                    match frame.frame_type.as_deref() {
                        Some("welcome") => break,
                        Some("disconnect") => {
                            return Err("server disconnect during welcome".into());
                        }
                        _ => continue,
                    }
                }
            }
            Some(Ok(_)) => continue,
            Some(Err(e)) => return Err(format!("read welcome: {}", e)),
            None => return Err("connection closed before welcome".into()),
        }
    }

    Ok((sink, stream))
}

async fn read_until_disconnect(inner: &EventStreamInner, mut stream: WsStream) {
    while let Some(msg) = stream.next().await {
        match msg {
            Ok(Message::Text(text)) => {
                inner.dispatch(text.as_bytes()).await;
            }
            Ok(Message::Binary(data)) => {
                inner.dispatch(&data).await;
            }
            Ok(Message::Close(_)) => {
                inner.set_final_err("connection closed".into()).await;
                return;
            }
            Ok(_) => continue,
            Err(e) => {
                inner.set_final_err(e.to_string()).await;
                return;
            }
        }
    }
    inner
        .set_final_err("connection stream ended".into())
        .await;
}

fn backoff(policy: &ReconnectPolicy, attempt: usize) -> Duration {
    let shift = (attempt - 1).min(30);
    let exp_ms = policy.initial_delay.as_millis() as u64 * (1u64 << shift);
    let max_ms = policy.max_delay.as_millis() as u64;
    let exp_ms = exp_ms.min(max_ms);
    if exp_ms == 0 {
        return Duration::ZERO;
    }
    let jitter_ms = rand::random::<u64>() % exp_ms;
    let delay_ms = jitter_ms + policy.initial_delay.as_millis() as u64;
    Duration::from_millis(delay_ms.min(max_ms))
}

async fn run_loop(inner: Arc<EventStreamInner>, initial_stream: WsStream) {
    read_until_disconnect(&inner, initial_stream).await;

    if inner.opts.reconnect.disabled || inner.closed.load(Ordering::SeqCst) {
        let err = inner
            .get_final_err()
            .await
            .unwrap_or_else(|| "stream ended".into());
        inner.terminate_all_subs(&err).await;
        inner.done.notify_waiters();
        return;
    }

    let mut attempt = 0usize;
    loop {
        if inner.closed.load(Ordering::SeqCst) {
            inner.done.notify_waiters();
            return;
        }
        attempt += 1;
        let max = inner.opts.reconnect.max_attempts;
        if max > 0 && attempt > max {
            let err = format!(
                "reconnect: exhausted after {} attempts: {}",
                max,
                inner
                    .get_final_err()
                    .await
                    .unwrap_or_else(|| "unknown".into())
            );
            inner.set_final_err(err.clone()).await;
            inner.terminate_all_subs(&err).await;
            inner.done.notify_waiters();
            return;
        }

        let delay = backoff(&inner.opts.reconnect, attempt);
        tracing::debug!(
            event = "ws_reconnect",
            attempt = attempt,
            delay_ms = delay.as_millis() as u64,
            "reconnecting event stream"
        );

        tokio::time::sleep(delay).await;

        if inner.closed.load(Ordering::SeqCst) {
            inner.done.notify_waiters();
            return;
        }

        let (sink, stream) = match do_connect(&inner).await {
            Ok(pair) => pair,
            Err(e) => {
                inner.set_final_err(e).await;
                continue;
            }
        };

        {
            let mut guard = inner.ws_sink.write().await;
            *guard = Some(sink);
        }

        inner.resubscribe_all().await;
        let connected_at = tokio::time::Instant::now();
        read_until_disconnect(&inner, stream).await;

        if connected_at.elapsed() >= RECONNECT_STABLE_THRESHOLD {
            attempt = 0;
        }

        if inner.closed.load(Ordering::SeqCst) {
            inner.done.notify_waiters();
            return;
        }
    }
}

/// Opens a real-time event stream. This is the internal entry point;
/// Client wraps it with its own configuration.
pub(crate) async fn open_event_stream(
    http: reqwest::Client,
    api_base_url: &str,
    ws_base_url: &str,
    api_version_name: &str,
    access_token: Option<&str>,
    opts: EventStreamOptions,
) -> Result<EventStream, String> {
    let inner = Arc::new(EventStreamInner {
        base_url: if ws_base_url.is_empty() {
            DEFAULT_EVENT_STREAM_URL.to_string()
        } else {
            ws_base_url.to_string()
        },
        api_version_name: api_version_name.to_string(),
        opts,
        ws_sink: RwLock::new(None),
        subs: RwLock::new(HashMap::new()),
        closed: AtomicBool::new(false),
        done: Notify::new(),
        final_err: Mutex::new(None),
        http,
        api_base_url: api_base_url.to_string(),
        access_token: RwLock::new(access_token.map(|s| s.to_string())),
    });

    let (sink, stream) = do_connect(&inner).await?;
    {
        let mut guard = inner.ws_sink.write().await;
        *guard = Some(sink);
    }

    let inner_clone = inner.clone();
    let handle = tokio::spawn(async move {
        run_loop(inner_clone, stream).await;
    });

    Ok(EventStream {
        inner,
        _run_handle: handle,
    })
}
