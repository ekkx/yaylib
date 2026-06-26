use std::sync::Arc;

use crate::auth::{LoginEmailUserRequestBody, LoginResponse, LoginWithEmailRequest};
use crate::config::*;
use crate::errors::ApiError;
use crate::event_stream::{self, EventStream, EventStreamOptions, DEFAULT_EVENT_STREAM_URL};
use crate::upload::{self, Upload};
use crate::retry::RetryPolicy;
use crate::session::{Session, SessionError, SessionStore};
use crate::signing;
use crate::transport::Transport;
use crate::uuid::new_uuid_v4;

// ---------------------------------------------------------------------------
// Shared internal types
// ---------------------------------------------------------------------------

#[derive(Debug)]
pub(crate) struct ClientConfig {
    pub(crate) api_key: String,
    pub(crate) api_version_key: String,
    pub(crate) api_version_name: String,
    pub(crate) app_version: String,
    pub(crate) user_agent: String,
    pub(crate) device_info: String,
    pub(crate) connection_type: String,
    pub(crate) connection_speed: String,
    pub(crate) accept_language: String,
    pub(crate) base_url: String,
    pub(crate) cassandra_base_url: String,
    pub(crate) event_stream_url: String,
    pub(crate) retry_policy: RetryPolicy,
}

#[derive(Debug)]
pub(crate) struct SharedState {
    pub(crate) access_token: String,
    pub(crate) refresh_token: String,
    pub(crate) user_id: i64,
    pub(crate) current_email: String,
    pub(crate) device_uuid: String,
    pub(crate) client_ip: String,
    /// Reserved for lazy IP fetch (not yet implemented).
    #[allow(dead_code)]
    pub(crate) client_ip_fetching: bool,
}

// ---------------------------------------------------------------------------
// ClientBuilder
// ---------------------------------------------------------------------------

/// Builder for configuring a [`Client`] instance.
#[derive(Default)]
pub struct ClientBuilder {
    api_key: Option<String>,
    api_version_key: Option<String>,
    api_version_name: Option<String>,
    app_version: Option<String>,
    user_agent: Option<String>,
    device_info: Option<String>,
    device_uuid: Option<String>,
    base_url: Option<String>,
    cassandra_base_url: Option<String>,
    event_stream_url: Option<String>,
    connection_type: Option<String>,
    connection_speed: Option<String>,
    accept_language: Option<String>,
    client_ip: Option<String>,
    retry_policy: Option<RetryPolicy>,
    http_client: Option<reqwest::Client>,
    session_store: Option<Arc<dyn SessionStore>>,
    user_id: Option<i64>,
}

impl ClientBuilder {
    /// Sets a custom API key (defaults to the built-in key).
    pub fn api_key(mut self, v: impl Into<String>) -> Self {
        self.api_key = Some(v.into());
        self
    }
    pub fn api_version_key(mut self, v: impl Into<String>) -> Self {
        self.api_version_key = Some(v.into());
        self
    }
    pub fn api_version_name(mut self, v: impl Into<String>) -> Self {
        self.api_version_name = Some(v.into());
        self
    }
    pub fn app_version(mut self, v: impl Into<String>) -> Self {
        self.app_version = Some(v.into());
        self
    }
    pub fn user_agent(mut self, v: impl Into<String>) -> Self {
        self.user_agent = Some(v.into());
        self
    }
    pub fn device_info(mut self, v: impl Into<String>) -> Self {
        self.device_info = Some(v.into());
        self
    }
    pub fn device_uuid(mut self, v: impl Into<String>) -> Self {
        self.device_uuid = Some(v.into());
        self
    }
    /// Overrides the base API URL (defaults to `https://api.yay.space/`).
    pub fn base_url(mut self, v: impl Into<String>) -> Self {
        self.base_url = Some(v.into());
        self
    }
    /// Overrides the Cassandra service base URL.
    pub fn cassandra_base_url(mut self, v: impl Into<String>) -> Self {
        self.cassandra_base_url = Some(v.into());
        self
    }
    /// Overrides the WebSocket URL for real-time events.
    pub fn event_stream_url(mut self, v: impl Into<String>) -> Self {
        self.event_stream_url = Some(v.into());
        self
    }
    pub fn connection_type(mut self, v: impl Into<String>) -> Self {
        self.connection_type = Some(v.into());
        self
    }
    pub fn connection_speed(mut self, v: impl Into<String>) -> Self {
        self.connection_speed = Some(v.into());
        self
    }
    pub fn accept_language(mut self, v: impl Into<String>) -> Self {
        self.accept_language = Some(v.into());
        self
    }
    pub fn client_ip(mut self, v: impl Into<String>) -> Self {
        self.client_ip = Some(v.into());
        self
    }
    /// Sets the retry policy for transient failures.
    pub fn retry_policy(mut self, v: RetryPolicy) -> Self {
        self.retry_policy = Some(v);
        self
    }
    /// Supplies a pre-configured `reqwest::Client` for HTTP requests.
    pub fn http_client(mut self, v: reqwest::Client) -> Self {
        self.http_client = Some(v);
        self
    }
    /// Attaches a session store for transparent login caching.
    pub fn session_store(mut self, v: impl SessionStore + 'static) -> Self {
        self.session_store = Some(Arc::new(v));
        self
    }
    /// Like [`session_store`](Self::session_store), but accepts a pre-wrapped `Arc`.
    pub fn session_store_arc(mut self, v: Arc<dyn SessionStore>) -> Self {
        self.session_store = Some(v);
        self
    }
    pub fn user_id(mut self, v: i64) -> Self {
        self.user_id = Some(v);
        self
    }

    /// Builds the [`Client`] with the configured options.
    pub fn build(self) -> Client {
        let config = Arc::new(ClientConfig {
            api_key: self
                .api_key
                .unwrap_or_else(|| DEFAULT_API_KEY.to_string()),
            api_version_key: self
                .api_version_key
                .unwrap_or_else(|| DEFAULT_API_VERSION_KEY.to_string()),
            api_version_name: self
                .api_version_name
                .unwrap_or_else(|| DEFAULT_API_VERSION_NAME.to_string()),
            app_version: self
                .app_version
                .unwrap_or_else(|| DEFAULT_APP_VERSION.to_string()),
            user_agent: self.user_agent.unwrap_or_else(default_user_agent),
            device_info: self.device_info.unwrap_or_else(default_device_info),
            connection_type: self
                .connection_type
                .unwrap_or_else(|| DEFAULT_CONNECTION_TYPE.to_string()),
            connection_speed: self
                .connection_speed
                .unwrap_or_else(|| DEFAULT_CONNECTION_SPEED.to_string()),
            accept_language: self
                .accept_language
                .unwrap_or_else(|| DEFAULT_ACCEPT_LANGUAGE.to_string()),
            base_url: self
                .base_url
                .unwrap_or_else(|| DEFAULT_BASE_URL.to_string()),
            cassandra_base_url: self
                .cassandra_base_url
                .unwrap_or_else(|| DEFAULT_CASSANDRA_BASE_URL.to_string()),
            event_stream_url: self
                .event_stream_url
                .unwrap_or_else(|| DEFAULT_EVENT_STREAM_URL.to_string()),
            retry_policy: self.retry_policy.unwrap_or_default(),
        });

        let state = Arc::new(std::sync::RwLock::new(SharedState {
            access_token: String::new(),
            refresh_token: String::new(),
            user_id: self.user_id.unwrap_or(0),
            current_email: String::new(),
            device_uuid: self.device_uuid.unwrap_or_else(new_uuid_v4),
            client_ip: self.client_ip.unwrap_or_default(),
            client_ip_fetching: false,
        }));

        let http_client = self.http_client.unwrap_or_else(|| {
            reqwest::Client::builder()
                .timeout(DEFAULT_HTTP_TIMEOUT)
                .build()
                .expect("failed to build default HTTP client")
        });

        let transport = Transport {
            http_client,
            config: config.clone(),
            state: state.clone(),
            refresh_lock: Arc::new(tokio::sync::Mutex::new(())),
            session_store: self.session_store.clone(),
        };

        let (shutdown_tx, _) = tokio::sync::watch::channel(false);

        Client {
            config,
            state,
            transport,
            session_store: self.session_store,
            shutdown_tx: Arc::new(shutdown_tx),
        }
    }
}

// ---------------------------------------------------------------------------
// Client
// ---------------------------------------------------------------------------

/// Client for the Yay! API. Thread-safe and cloneable.
#[derive(Clone)]
pub struct Client {
    pub(crate) config: Arc<ClientConfig>,
    pub(crate) state: Arc<std::sync::RwLock<SharedState>>,
    pub(crate) transport: Transport,
    session_store: Option<Arc<dyn SessionStore>>,
    shutdown_tx: Arc<tokio::sync::watch::Sender<bool>>,
}

impl std::fmt::Debug for Client {
    fn fmt(&self, f: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
        f.debug_struct("Client")
            .field("base_url", &self.config.base_url)
            .field("api_version_name", &self.config.api_version_name)
            .finish_non_exhaustive()
    }
}

impl Default for Client {
    fn default() -> Self {
        Self::new()
    }
}

impl Client {
    /// Creates a new `Client` with default settings.
    pub fn new() -> Self {
        ClientBuilder::default().build()
    }

    /// Creates a new [`ClientBuilder`] with default settings.
    pub fn builder() -> ClientBuilder {
        ClientBuilder::default()
    }

    /// Shuts down background tasks (e.g. event streams) associated with this client.
    pub fn close(&self) {
        let _ = self.shutdown_tx.send(true);
    }

    /// Returns a watch receiver that fires when [`close`](Self::close) is
    /// called. Background tasks should select on this to stop promptly.
    #[allow(dead_code)]
    pub(crate) fn shutdown_rx(&self) -> tokio::sync::watch::Receiver<bool> {
        self.shutdown_tx.subscribe()
    }

    pub fn transport(&self) -> &Transport {
        &self.transport
    }

    // -- Read-only config accessors -------------------------------------------

    pub fn api_key(&self) -> &str {
        &self.config.api_key
    }
    pub fn api_version_key(&self) -> &str {
        &self.config.api_version_key
    }
    pub fn api_version_name(&self) -> &str {
        &self.config.api_version_name
    }
    pub fn app_version(&self) -> &str {
        &self.config.app_version
    }
    pub fn base_url(&self) -> &str {
        &self.config.base_url
    }
    pub fn cassandra_base_url(&self) -> &str {
        &self.config.cassandra_base_url
    }
    pub fn event_stream_url(&self) -> &str {
        &self.config.event_stream_url
    }
    pub fn retry_policy(&self) -> &RetryPolicy {
        &self.config.retry_policy
    }

    // -- Mutable state accessors ----------------------------------------------

    pub fn tokens(&self) -> crate::session::Tokens {
        let state = self.state.read().unwrap_or_else(|e| e.into_inner());
        crate::session::Tokens {
            access: state.access_token.clone(),
            refresh: state.refresh_token.clone(),
        }
    }

    pub fn set_tokens(&self, access: &str, refresh: &str) {
        let mut state = self.state.write().unwrap_or_else(|e| e.into_inner());
        state.access_token = access.to_string();
        state.refresh_token = refresh.to_string();
    }

    pub fn user_id(&self) -> i64 {
        self.state.read().unwrap_or_else(|e| e.into_inner()).user_id
    }

    pub fn set_user_id(&self, uid: i64) {
        self.state.write().unwrap_or_else(|e| e.into_inner()).user_id = uid;
    }

    pub fn device_uuid(&self) -> String {
        self.state.read().unwrap_or_else(|e| e.into_inner()).device_uuid.clone()
    }

    pub(crate) fn set_device_uuid(&self, uuid: String) {
        self.state.write().unwrap_or_else(|e| e.into_inner()).device_uuid = uuid;
    }

    pub fn set_login_identity(&self, email: &str, user_id: i64) {
        let mut state = self.state.write().unwrap_or_else(|e| e.into_inner());
        if !email.is_empty() {
            state.current_email = email.to_string();
        }
        if user_id != 0 {
            state.user_id = user_id;
        }
    }

    #[allow(dead_code)]
    pub(crate) fn current_email(&self) -> String {
        self.state.read().unwrap_or_else(|e| e.into_inner()).current_email.clone()
    }

    pub fn session_store(&self) -> Option<&dyn SessionStore> {
        self.session_store.as_deref()
    }

    // -- Session convenience --------------------------------------------------

    pub async fn load_session(&self, email: &str) -> Result<Session, ApiError> {
        let store = self
            .session_store
            .as_ref()
            .ok_or_else(|| ApiError::Other("session store not configured".into()))?;
        let session = store.load(email).await.map_err(|e| match e {
            SessionError::NoSession => ApiError::Other("no session for that email".into()),
            other => ApiError::Other(format!("load session: {}", other)),
        })?;
        self.set_tokens(&session.access_token, &session.refresh_token);
        if !session.device_uuid.is_empty() {
            self.set_device_uuid(session.device_uuid.clone());
        }
        self.set_login_identity(email, session.user_id);
        Ok(session)
    }

    pub async fn save_session(&self, session: &mut Session) -> Result<(), ApiError> {
        let store = self
            .session_store
            .as_ref()
            .ok_or_else(|| ApiError::Other("session store not configured".into()))?;
        if session.device_uuid.is_empty() {
            session.device_uuid = self.device_uuid();
        }
        store
            .save(session)
            .await
            .map_err(|e| ApiError::Other(format!("save session: {}", e)))?;
        self.set_tokens(&session.access_token, &session.refresh_token);
        if !session.device_uuid.is_empty() {
            self.set_device_uuid(session.device_uuid.clone());
        }
        self.set_login_identity(&session.email, session.user_id);
        Ok(())
    }

    // -- Generated API configuration -----------------------------------------

    /// Builds a [`crate::gen::apis::configuration::Configuration`] that
    /// delegates HTTP through this client's [`Transport`] (which handles
    /// auth headers, retry, host routing, etc.).  The generated free
    /// functions in `crate::gen::apis::*_api` accept this type.
    pub(crate) fn gen_configuration(&self) -> crate::gen::apis::configuration::Configuration {
        crate::gen::apis::configuration::Configuration {
            base_path: self.config.base_url.clone(),
            // Transport already sets User-Agent and all other required
            // headers in its set_headers() hook, so we leave these None
            // to avoid double-setting.
            user_agent: None,
            client: self.transport.clone(),
            basic_auth: None,
            oauth_access_token: None,
            bearer_access_token: None,
            api_key: None,
        }
    }

    // -- Signing helpers ------------------------------------------------------

    /// Fetches the server timestamp and computes `signed_info`.
    pub async fn generate_signed_info(&self) -> Result<signing::SignedInfo, ApiError> {
        let config = self.gen_configuration();
        let resp = crate::gen::apis::users_api::get_user_timestamp(&config).await?;
        let ts = resp.time.flatten().unwrap_or(0);
        Ok(self.generate_signed_info_at(ts))
    }

    /// Computes `signed_info` for a known timestamp (pure, no I/O).
    pub fn generate_signed_info_at(&self, timestamp: i64) -> signing::SignedInfo {
        let uuid = self.device_uuid();
        signing::compute_signed_info_at(self.api_key(), &uuid, timestamp)
    }

    /// Computes the `signed_version` HMAC token (pure, no I/O).
    pub fn generate_signed_version(&self) -> String {
        signing::compute_signed_version(self.api_version_key(), self.api_version_name())
    }

    /// Generates a short-lived HS256 JWT for write endpoints.
    pub fn generate_x_jwt(&self) -> String {
        signing::generate_x_jwt(self.api_version_key())
    }

    // -- Auth -----------------------------------------------------------------

    /// Starts a login-with-email request builder.
    pub fn login_with_email(&self) -> LoginWithEmailRequest<'_> {
        LoginWithEmailRequest::new(self)
    }

    pub(crate) async fn maybe_restore(
        &self,
        email: &str,
    ) -> Result<Option<LoginResponse>, ApiError> {
        let store = match &self.session_store {
            Some(s) => s,
            None => return Ok(None),
        };
        match store.load(email).await {
            Ok(session) => {
                self.set_tokens(&session.access_token, &session.refresh_token);
                if !session.device_uuid.is_empty() {
                    self.set_device_uuid(session.device_uuid);
                }
                self.set_login_identity(email, session.user_id);
                Ok(Some(LoginResponse {
                    user_id: session.user_id,
                    access_token: session.access_token,
                    refresh_token: session.refresh_token,
                }))
            }
            Err(SessionError::NoSession) => Ok(None),
            Err(e) => Err(ApiError::Other(format!("load session: {}", e))),
        }
    }

    pub(crate) async fn do_login_with_email(
        &self,
        body: LoginEmailUserRequestBody,
    ) -> Result<LoginResponse, ApiError> {
        let config = self.gen_configuration();
        let req = crate::gen::models::LoginEmailUserRequest {
            email: Some(Some(body.email)),
            password: Some(Some(body.password)),
            api_key: Some(Some(body.api_key)),
            uuid: Some(Some(body.uuid)),
            two_f_a_code: body.two_fa_code.map(Some),
        };
        let resp = crate::gen::apis::users_api::login_with_email(&config, req).await?;
        Ok(LoginResponse {
            user_id: resp.user_id.flatten().unwrap_or(0),
            access_token: resp.access_token.flatten().unwrap_or_default(),
            refresh_token: resp.refresh_token.flatten().unwrap_or_default(),
        })
    }
}

// ---------------------------------------------------------------------------
// Upload methods
// ---------------------------------------------------------------------------

impl Client {
    fn require_user_id(&self) -> Result<i64, ApiError> {
        let uid = self.user_id();
        if uid == 0 {
            return Err(ApiError::Other(
                "not logged in (call login_with_email before user-bound uploads)".into(),
            ));
        }
        Ok(uid)
    }

    fn access_token_snapshot(&self) -> Option<String> {
        let state = self.state.read().unwrap_or_else(|e| e.into_inner());
        if state.access_token.is_empty() {
            None
        } else {
            Some(state.access_token.clone())
        }
    }

    async fn upload_single_image(
        &self,
        category: &dyn upload::UploadCategory,
        file: Upload,
    ) -> Result<String, ApiError> {
        let token = self.access_token_snapshot();
        let files = [file];
        let names = upload::upload_images(
            &self.transport.http_client,
            &self.config.base_url,
            token.as_deref(),
            category,
            &files,
        )
        .await?;
        Ok(names.into_iter().next().unwrap_or_default())
    }

    async fn upload_multiple_images(
        &self,
        category: &dyn upload::UploadCategory,
        files: Vec<Upload>,
    ) -> Result<Vec<String>, ApiError> {
        let token = self.access_token_snapshot();
        upload::upload_images(
            &self.transport.http_client,
            &self.config.base_url,
            token.as_deref(),
            category,
            &files,
        )
        .await
    }

    /// Uploads the authenticated user's avatar image and returns the
    /// server-canonical filename to pass to edit_user as
    /// profile_icon_filename.
    pub async fn upload_avatar_image(&self, file: Upload) -> Result<String, ApiError> {
        let uid = self.require_user_id()?;
        let cat = upload::avatar_category(uid);
        self.upload_single_image(&cat, file).await
    }

    /// Uploads the authenticated user's profile cover image and returns
    /// the server-canonical filename.
    pub async fn upload_cover_image(&self, file: Upload) -> Result<String, ApiError> {
        let uid = self.require_user_id()?;
        let cat = upload::cover_category(uid);
        self.upload_single_image(&cat, file).await
    }

    /// Uploads up to 9 images for a new post by the authenticated user
    /// and returns the server-canonical filenames.
    pub async fn upload_post_images(&self, files: Vec<Upload>) -> Result<Vec<String>, ApiError> {
        let uid = self.require_user_id()?;
        let cat = upload::post_category(uid);
        self.upload_multiple_images(&cat, files).await
    }

    /// Uploads up to 9 images for a chat message in the given room and
    /// returns the server-canonical filenames.
    pub async fn upload_chat_message_images(
        &self,
        room_id: i64,
        files: Vec<Upload>,
    ) -> Result<Vec<String>, ApiError> {
        let uid = self.require_user_id()?;
        let cat = upload::chat_message_category(room_id, uid);
        self.upload_multiple_images(&cat, files).await
    }

    /// Uploads the background image of the chat room with the given ID
    /// and returns the server-canonical filename.
    pub async fn upload_chat_background_image(
        &self,
        room_id: i64,
        file: Upload,
    ) -> Result<String, ApiError> {
        let cat = upload::chat_background_category(room_id);
        self.upload_single_image(&cat, file).await
    }

    /// Uploads the icon of the group with the given ID and returns the
    /// server-canonical filename.
    pub async fn upload_group_icon_image(
        &self,
        group_id: i64,
        file: Upload,
    ) -> Result<String, ApiError> {
        let cat = upload::group_icon_category(group_id);
        self.upload_single_image(&cat, file).await
    }

    /// Uploads the cover image of the group with the given ID and
    /// returns the server-canonical filename.
    pub async fn upload_group_cover_image(
        &self,
        group_id: i64,
        file: Upload,
    ) -> Result<String, ApiError> {
        let cat = upload::group_cover_category(group_id);
        self.upload_single_image(&cat, file).await
    }

    /// Uploads the icon image for a group that is about to be created,
    /// before the group ID exists.
    pub async fn upload_group_creation_icon_image(
        &self,
        file: Upload,
    ) -> Result<String, ApiError> {
        let cat = upload::group_creation_icon_category();
        self.upload_single_image(&cat, file).await
    }

    /// Uploads the cover image for a group that is about to be created.
    pub async fn upload_group_creation_cover_image(
        &self,
        file: Upload,
    ) -> Result<String, ApiError> {
        let cat = upload::group_creation_cover_category();
        self.upload_single_image(&cat, file).await
    }

    /// Uploads an avatar image during the signup flow, before the
    /// account exists.
    pub async fn upload_signup_avatar_image(&self, file: Upload) -> Result<String, ApiError> {
        let cat = upload::signup_avatar_category();
        self.upload_single_image(&cat, file).await
    }

    /// Uploads the icon of a thread under the given group and returns
    /// the server-canonical filename.
    pub async fn upload_thread_icon_image(
        &self,
        group_id: i64,
        file: Upload,
    ) -> Result<String, ApiError> {
        let cat = upload::thread_icon_category(group_id);
        self.upload_single_image(&cat, file).await
    }

    /// Uploads screenshots attached to a moderation report (up to 4)
    /// and returns the server-canonical filenames.
    pub async fn upload_report_images(
        &self,
        files: Vec<Upload>,
    ) -> Result<Vec<String>, ApiError> {
        let cat = upload::report_category();
        self.upload_multiple_images(&cat, files).await
    }

    /// Uploads a snapshot taken from a video call. No thumbnail is
    /// generated for this category.
    pub async fn upload_video_call_snapshot_image(
        &self,
        file: Upload,
    ) -> Result<String, ApiError> {
        let cat = upload::video_call_snapshot_category();
        self.upload_single_image(&cat, file).await
    }

    /// Uploads a single MP4 video and returns the filename to pass to
    /// create_post as video_file_name.
    pub async fn upload_video(&self, body: Vec<u8>) -> Result<String, ApiError> {
        let token = self.access_token_snapshot();
        upload::upload_video(
            &self.transport.http_client,
            &self.config.base_url,
            token.as_deref(),
            body,
        )
        .await
    }
}

// ---------------------------------------------------------------------------
// Event stream
// ---------------------------------------------------------------------------

impl Client {
    /// Opens a real-time event stream with default options.
    pub async fn open_event_stream(&self) -> Result<EventStream, ApiError> {
        self.open_event_stream_with_options(EventStreamOptions::default())
            .await
    }

    /// Opens a real-time event stream with custom options.
    pub async fn open_event_stream_with_options(
        &self,
        opts: EventStreamOptions,
    ) -> Result<EventStream, ApiError> {
        let token = self.access_token_snapshot();
        event_stream::open_event_stream(
            self.transport.http_client.clone(),
            &self.config.base_url,
            &self.config.event_stream_url,
            &self.config.api_version_name,
            token.as_deref(),
            opts,
        )
        .await
        .map_err(ApiError::Other)
    }
}
