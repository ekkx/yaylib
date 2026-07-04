use std::sync::Arc;
use std::time::Duration;

use base64::Engine;
use reqwest::header::{
    HeaderMap, HeaderName, HeaderValue, ACCEPT_LANGUAGE, AUTHORIZATION, CONTENT_TYPE, USER_AGENT,
};

use crate::client::{ClientConfig, SharedState};
use crate::retry::{extract_retry_in, next_delay, should_retry};

#[derive(Clone)]
pub struct Transport {
    pub(crate) http_client: reqwest::Client,
    pub(crate) config: Arc<ClientConfig>,
    pub(crate) state: Arc<std::sync::RwLock<SharedState>>,
    pub(crate) refresh_lock: Arc<tokio::sync::Mutex<()>>,
    pub(crate) session_store: Option<Arc<dyn crate::session::SessionStore>>,
}

impl std::fmt::Debug for Transport {
    fn fmt(&self, f: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
        f.debug_struct("Transport")
            .field("config", &self.config)
            .field("session_store", &self.session_store.is_some())
            .finish_non_exhaustive()
    }
}

impl Transport {
    pub fn request(&self, method: reqwest::Method, url: &str) -> reqwest::RequestBuilder {
        self.http_client.request(method, url)
    }

    pub async fn execute(
        &self,
        req: reqwest::Request,
    ) -> Result<reqwest::Response, reqwest::Error> {
        let method = req.method().clone();
        let mut url = req.url().clone();
        self.route_host(&method, &mut url);

        let policy = &self.config.retry_policy;
        let max_attempts = policy.max_attempts.max(1);
        let mut prev_retry_in: Option<Duration> = None;

        // Buffer the body for retry only if it is byte-representable.
        // Multipart/streamed bodies return None from as_bytes() and
        // cannot be cloned; those requests use the original Request on
        // the first attempt and are not retried.
        let replayable_body: Option<Vec<u8>> = req
            .body()
            .and_then(|b| b.as_bytes())
            .map(|b| b.to_vec());
        let original_headers = req.headers().clone();
        let is_replayable = replayable_body.is_some() || req.body().is_none();

        // Hold the original request for attempt 1; taken via .take() so
        // subsequent iterations know to rebuild from the buffered body.
        let mut original_req = Some(req);

        for attempt in 1..=max_attempts {
            if attempt > 1 {
                if !is_replayable {
                    break;
                }
                let wait = next_delay(prev_retry_in, attempt, policy);
                tracing::debug!(
                    event = crate::logging::LOG_EVENT_HTTP_RETRY,
                    method = %method,
                    path = url.path(),
                    attempt,
                    delay_ms = wait.as_millis() as u64,
                );
                tokio::time::sleep(wait).await;
            }

            let attempt_req = if let Some(mut orig) = original_req.take() {
                // First attempt: mutate the original request in-place
                // to preserve non-cloneable bodies (multipart, streams).
                *orig.url_mut() = url.clone();
                self.set_headers(&mut orig);
                orig
            } else {
                // Subsequent attempts: rebuild from buffered body.
                let mut r = reqwest::Request::new(method.clone(), url.clone());
                *r.headers_mut() = original_headers.clone();
                if let Some(ref body) = replayable_body {
                    *r.body_mut() = Some(reqwest::Body::from(body.clone()));
                }
                self.set_headers(&mut r);
                r
            };

            tracing::debug!(
                event = crate::logging::LOG_EVENT_HTTP_REQUEST,
                method = %method,
                path = url.path(),
                host = url.host_str().unwrap_or(""),
            );

            match self.http_client.execute(attempt_req).await {
                Ok(mut resp) => {
                    if resp.status().as_u16() == 401 && !is_oauth_token_path(url.path()) {
                        // Skip 401 refresh+replay for non-replayable bodies:
                        // we cannot re-send the multipart/stream body after
                        // refresh, so return the 401 as-is.
                        if is_replayable {
                            resp = self
                                .handle_401(
                                    resp,
                                    &method,
                                    &url,
                                    &original_headers,
                                    replayable_body.as_deref(),
                                )
                                .await?;
                        }
                    }

                    let status = resp.status().as_u16();
                    if !should_retry(Some(status), false, method.as_str(), policy)
                        || attempt == max_attempts
                        || !is_replayable
                    {
                        return Ok(resp);
                    }

                    prev_retry_in = if status >= 400 {
                        match resp.bytes().await {
                            Ok(body) => extract_retry_in(&body),
                            Err(_) => None,
                        }
                    } else {
                        None
                    };
                }
                Err(e) => {
                    if !should_retry(None, true, method.as_str(), policy)
                        || attempt == max_attempts
                        || !is_replayable
                    {
                        return Err(e);
                    }
                    prev_retry_in = None;
                }
            }
        }

        unreachable!("retry loop must return on final attempt")
    }

    fn set_headers(&self, req: &mut reqwest::Request) {
        let (access_token, device_uuid, client_ip) = {
            let state = self.state.read().unwrap_or_else(|e| e.into_inner());
            (
                state.access_token.clone(),
                state.device_uuid.clone(),
                state.client_ip.clone(),
            )
        };

        let config = &self.config;

        // Snapshot the path before taking a mutable borrow on headers.
        let is_oauth = is_oauth_token_path(req.url().path());

        let headers = req.headers_mut();

        set_if_absent(headers, USER_AGENT, &config.user_agent);
        set_if_absent(
            headers,
            HeaderName::from_static("x-app-version"),
            &config.api_version_name,
        );
        set_if_absent(
            headers,
            HeaderName::from_static("x-device-info"),
            &config.device_info,
        );
        set_if_absent(
            headers,
            HeaderName::from_static("x-device-uuid"),
            &device_uuid,
        );
        if !client_ip.is_empty() {
            set_if_absent(
                headers,
                HeaderName::from_static("x-client-ip"),
                &client_ip,
            );
        }
        set_if_absent(
            headers,
            HeaderName::from_static("x-connection-type"),
            &config.connection_type,
        );
        set_if_absent(
            headers,
            HeaderName::from_static("x-connection-speed"),
            &config.connection_speed,
        );
        set_if_absent(headers, ACCEPT_LANGUAGE, &config.accept_language);

        // X-Timestamp must be fresh per request.
        let timestamp = std::time::SystemTime::now()
            .duration_since(std::time::UNIX_EPOCH)
            .unwrap()
            .as_secs()
            .to_string();
        headers.insert(
            HeaderName::from_static("x-timestamp"),
            HeaderValue::from_str(&timestamp).unwrap(),
        );

        // Normalize JSON Content-Type to the exact casing the server expects.
        let should_normalize = headers
            .get(CONTENT_TYPE)
            .and_then(|ct| ct.to_str().ok())
            .is_some_and(|ct| ct.starts_with("application/json"));
        if should_normalize {
            headers.insert(
                CONTENT_TYPE,
                HeaderValue::from_static("application/json;charset=UTF-8"),
            );
        }

        // Authorization: Basic for oauth/token, Bearer otherwise.
        if !headers.contains_key(AUTHORIZATION) {
            if is_oauth {
                let cred = basic_credentials(&config.api_key);
                if let Ok(v) = HeaderValue::from_str(&format!("Basic {}", cred)) {
                    headers.insert(AUTHORIZATION, v);
                }
            } else if !access_token.is_empty() {
                if let Ok(v) = HeaderValue::from_str(&format!("Bearer {}", access_token)) {
                    headers.insert(AUTHORIZATION, v);
                }
            }
        }
    }

    /// Handles a 401 response by attempting a token refresh and replaying
    /// the original request.
    ///
    /// - No refresh token / refresh failed: the original 401 is returned
    ///   so the caller sees the server's auth-failure message.
    /// - Refresh succeeded but replay network-errored: the replay error is
    ///   returned (not the stale 401).
    /// - Refresh succeeded and replay succeeded: the new response replaces
    ///   the 401.
    async fn handle_401(
        &self,
        original_resp: reqwest::Response,
        method: &reqwest::Method,
        url: &url::Url,
        original_headers: &HeaderMap,
        body_bytes: Option<&[u8]>,
    ) -> Result<reqwest::Response, reqwest::Error> {
        let (refresh_token, stale_access) = {
            let state = self.state.read().unwrap_or_else(|e| e.into_inner());
            (state.refresh_token.clone(), state.access_token.clone())
        };

        if refresh_token.is_empty() {
            tracing::debug!(event = crate::logging::LOG_EVENT_TOKEN_REFRESH, outcome = "no_token");
            return Ok(original_resp);
        }

        if self
            .refresh_tokens(&refresh_token, &stale_access)
            .await
            .is_err()
        {
            return Ok(original_resp);
        }

        let mut retry_req = reqwest::Request::new(method.clone(), url.clone());
        *retry_req.headers_mut() = original_headers.clone();
        retry_req.headers_mut().remove(AUTHORIZATION);
        if let Some(body) = body_bytes {
            *retry_req.body_mut() = Some(reqwest::Body::from(body.to_vec()));
        }
        self.set_headers(&mut retry_req);

        self.http_client.execute(retry_req).await
    }

    /// Refreshes OAuth tokens via a direct HTTP call (not through the
    /// middleware) to avoid recursion. A mutex ensures only one refresh
    /// runs at a time; concurrent 401s coalesce on the same refresh.
    async fn refresh_tokens(
        &self,
        refresh_token: &str,
        stale_access: &str,
    ) -> Result<(), Box<dyn std::error::Error + Send + Sync>> {
        let _guard = self.refresh_lock.lock().await;

        // Another task may have already refreshed while we waited for the lock.
        let (current_email, device_uuid) = {
            let state = self.state.read().unwrap_or_else(|e| e.into_inner());
            if state.access_token != stale_access {
                return Ok(());
            }
            (state.current_email.clone(), state.device_uuid.clone())
        };

        let config = &self.config;

        tracing::debug!(event = crate::logging::LOG_EVENT_TOKEN_REFRESH, outcome = "start");

        let timestamp = std::time::SystemTime::now()
            .duration_since(std::time::UNIX_EPOCH)
            .unwrap()
            .as_secs()
            .to_string();

        let resp = self
            .http_client
            .post(format!("{}/api/v1/oauth/token", config.base_url))
            .header(
                AUTHORIZATION,
                format!("Basic {}", basic_credentials(&config.api_key)),
            )
            .header(USER_AGENT, config.user_agent.as_str())
            .header("x-app-version", config.api_version_name.as_str())
            .header("x-device-info", config.device_info.as_str())
            .header("x-device-uuid", &device_uuid)
            .header("x-connection-type", config.connection_type.as_str())
            .header("x-connection-speed", config.connection_speed.as_str())
            .header("accept-language", config.accept_language.as_str())
            .header("x-timestamp", &timestamp)
            .form(&[
                ("grant_type", "refresh_token"),
                ("refresh_token", refresh_token),
            ])
            .send()
            .await?;

        if !resp.status().is_success() {
            tracing::warn!(
                event = crate::logging::LOG_EVENT_TOKEN_REFRESH,
                outcome = "failed",
                status = resp.status().as_u16(),
            );
            return Err(format!("token refresh failed: {}", resp.status()).into());
        }

        #[derive(serde::Deserialize)]
        struct TokenResponse {
            access_token: String,
            refresh_token: String,
        }

        let tokens: TokenResponse = resp.json().await?;

        let (new_access, new_refresh) = {
            let mut state = self.state.write().unwrap_or_else(|e| e.into_inner());
            state.access_token = tokens.access_token.clone();
            state.refresh_token = tokens.refresh_token.clone();
            (tokens.access_token, tokens.refresh_token)
        };

        tracing::debug!(event = crate::logging::LOG_EVENT_TOKEN_REFRESH, outcome = "ok");

        // Persist refreshed tokens to session store asynchronously.
        if let Some(store) = &self.session_store {
            if !current_email.is_empty() {
                let store = store.clone();
                let user_id = self.state.read().unwrap_or_else(|e| e.into_inner()).user_id;
                let session = crate::session::Session {
                    email: current_email,
                    user_id,
                    access_token: new_access,
                    refresh_token: new_refresh,
                    device_uuid,
                    updated_at: chrono::Utc::now(),
                };
                tokio::spawn(async move {
                    if let Err(e) = store.save(&session).await {
                        tracing::warn!(
                            event = crate::logging::LOG_EVENT_TOKEN_PERSIST_FAIL,
                            error = %e,
                        );
                    }
                });
            }
        }

        Ok(())
    }

    fn route_host(&self, method: &reqwest::Method, url: &mut url::Url) {
        let key = format!("{} {}", method.as_str(), url.path());
        let base_url = match self.host_base_url_for(&key) {
            Some(u) => u,
            None => return,
        };
        if let Ok(parsed) = url::Url::parse(base_url) {
            if let Some(host) = parsed.host_str() {
                let _ = url.set_scheme(parsed.scheme());
                let _ = url.set_host(Some(host));
                let _ = url.set_port(parsed.port());
                tracing::debug!(
                    event = crate::logging::LOG_EVENT_HOST_ROUTE_REWRITE,
                    op = key.as_str(),
                    host,
                );
            }
        }
    }

    fn host_base_url_for(&self, key: &str) -> Option<&str> {
        let alias = crate::host_routes::HOST_ROUTES.get(key).copied()?;
        match alias {
            "cassandra" => Some(&self.config.cassandra_base_url),
            _ => None,
        }
    }
}

impl Default for Transport {
    fn default() -> Self {
        let config = Arc::new(crate::client::ClientConfig {
            api_key: crate::config::DEFAULT_API_KEY.into(),
            web_api_key: crate::config::DEFAULT_WEB_API_KEY.into(),
            api_version_key: crate::config::DEFAULT_API_VERSION_KEY.into(),
            api_version_name: crate::config::DEFAULT_API_VERSION_NAME.into(),
            app_version: crate::config::DEFAULT_APP_VERSION.into(),
            user_agent: crate::config::default_user_agent(),
            device_info: crate::config::default_device_info(),
            connection_type: crate::config::DEFAULT_CONNECTION_TYPE.into(),
            connection_speed: crate::config::DEFAULT_CONNECTION_SPEED.into(),
            accept_language: crate::config::DEFAULT_ACCEPT_LANGUAGE.into(),
            base_url: crate::config::DEFAULT_BASE_URL.into(),
            cassandra_base_url: crate::config::DEFAULT_CASSANDRA_BASE_URL.into(),
            event_stream_url: crate::event_stream::DEFAULT_EVENT_STREAM_URL.into(),
            retry_policy: crate::retry::RetryPolicy::default(),
        });
        let state = Arc::new(std::sync::RwLock::new(crate::client::SharedState {
            access_token: String::new(),
            refresh_token: String::new(),
            user_id: 0,
            current_email: String::new(),
            device_uuid: crate::uuid::new_uuid_v4(),
            client_ip: String::new(),
            client_ip_fetching: false,
        }));
        Self {
            http_client: reqwest::Client::new(),
            config,
            state,
            refresh_lock: Arc::new(tokio::sync::Mutex::new(())),
            session_store: None,
        }
    }
}

fn set_if_absent(headers: &mut HeaderMap, name: HeaderName, value: &str) {
    if !headers.contains_key(&name) {
        if let Ok(v) = HeaderValue::from_str(value) {
            headers.insert(name, v);
        }
    }
}

pub(crate) fn basic_credentials(api_key: &str) -> String {
    base64::engine::general_purpose::STANDARD.encode(api_key.as_bytes())
}

fn is_oauth_token_path(path: &str) -> bool {
    path.contains("/api/v1/oauth/token")
}
