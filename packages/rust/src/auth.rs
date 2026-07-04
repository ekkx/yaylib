use chrono::Utc;

use crate::client::Client;
use crate::session::{Session, SessionError, SessionStore};

/// Wraps a session-store write failure that occurred after an otherwise
/// successful login. The response and tokens are still available: tokens
/// are active in memory on the Client, and the response is attached to
/// this error so callers can extract it via pattern matching.
#[derive(Debug, thiserror::Error)]
#[error("session save failed: {source}")]
pub struct SessionSaveFailed {
    pub response: LoginResponse,
    #[source]
    pub source: SessionError,
}

/// Builder for a `login_with_email` call. Constructed via
/// [`Client::login_with_email`].
///
/// Each field of the request body can be set directly on the builder;
/// [`execute`](Self::execute) assembles the underlying request. Two
/// conveniences are built in:
///
/// - `api_key` and `uuid` are auto-populated from the Client if the caller
///   left them blank.
/// - When a session store is active, `execute` first tries to restore a
///   cached session keyed by the email, and on a successful fresh login
///   persists the new session automatically.
///
/// Chain [`.no_cache()`](Self::no_cache) before `execute` to force a fresh
/// login regardless of cache contents.
pub struct LoginWithEmailRequest<'a> {
    client: &'a Client,
    email: Option<String>,
    password: Option<String>,
    api_key: Option<String>,
    uuid: Option<String>,
    two_fa_code: Option<String>,
    skip_cache: bool,
}

impl<'a> LoginWithEmailRequest<'a> {
    pub(crate) fn new(client: &'a Client) -> Self {
        Self {
            client,
            email: None,
            password: None,
            api_key: None,
            uuid: None,
            two_fa_code: None,
            skip_cache: false,
        }
    }

    pub fn email(mut self, v: impl Into<String>) -> Self {
        self.email = Some(v.into());
        self
    }

    pub fn password(mut self, v: impl Into<String>) -> Self {
        self.password = Some(v.into());
        self
    }

    pub fn api_key(mut self, v: impl Into<String>) -> Self {
        self.api_key = Some(v.into());
        self
    }

    pub fn uuid(mut self, v: impl Into<String>) -> Self {
        self.uuid = Some(v.into());
        self
    }

    pub fn two_fa_code(mut self, v: impl Into<String>) -> Self {
        self.two_fa_code = Some(v.into());
        self
    }

    /// Forces a fresh HTTP login even if a cached session is available.
    pub fn no_cache(mut self) -> Self {
        self.skip_cache = true;
        self
    }

    /// Issues the login, transparently using the session store when one
    /// is configured.
    ///
    /// On a cache hit the returned [`LoginResponse`] is synthesized from
    /// the cached tokens (`user_id`, `access_token`, `refresh_token`
    /// populated; other fields default).
    ///
    /// Errors fall into three buckets:
    ///
    /// - Session-store read failure (corrupt JSON, permission, I/O):
    ///   returned without attempting a fresh login, so a transient store
    ///   fault does not burn through the server's login rate limit.
    /// - Login HTTP failure: returned as-is.
    /// - Login succeeded but persisting the session failed: the response
    ///   is still returned so the tokens are usable; the persistence
    ///   error is wrapped as [`SessionSaveFailed`].
    pub async fn execute(self) -> Result<LoginResponse, crate::errors::ApiError> {
        let api_key = self
            .api_key
            .unwrap_or_else(|| self.client.api_key().to_string());
        let device_uuid = self
            .uuid
            .unwrap_or_else(|| self.client.device_uuid());
        let email = self.email.unwrap_or_default();

        // Try restoring a cached session.
        if !self.skip_cache && !email.is_empty() {
            match self.client.maybe_restore(&email).await {
                Ok(Some(resp)) => return Ok(resp),
                Ok(None) => {}
                Err(e) => return Err(e),
            }
        }

        // Whether the caller pinned a specific api_key (vs. the client
        // default). A pinned key opts out of the browser-login fallback.
        let caller_pinned = api_key.as_str() != self.client.api_key();

        // The key that establishes the session also refreshes its tokens;
        // the fallback switches it to web_api_key below.
        let mut oauth_key = api_key.clone();

        // Build and execute the HTTP login request.
        let body = LoginEmailUserRequestBody {
            email: email.clone(),
            password: self.password.unwrap_or_default(),
            api_key,
            uuid: device_uuid.clone(),
            two_fa_code: self.two_fa_code,
        };

        let resp = match self.client.do_login_with_email(body.clone()).await {
            Ok(resp) => resp,
            Err(e) => {
                // Fallback (PORTING.md §6.2): a plain "unauthorized"
                // retries against the browser login endpoint with
                // web_api_key. Skipped for other 401s — notably
                // Required2FA (the caller must supply a code, not switch
                // keys) — when the caller pinned a specific api_key, or
                // when web_api_key is empty (the off switch).
                if !self.client.web_api_key().is_empty()
                    && !caller_pinned
                    && crate::errors::code_of(&e)
                        == crate::error_codes::ErrorCode::Unauthorized
                {
                    match self.client.do_login_with_email_web(&body).await {
                        Ok(Some(resp)) => {
                            tracing::debug!(
                                event = "login_fallback",
                                outcome = "ok"
                            );
                            // Established with the web key → refresh with it.
                            oauth_key = self.client.web_api_key().to_string();
                            resp
                        }
                        // Fallback rejected → surface the primary error.
                        Ok(None) => return Err(e),
                        // 2xx decode failure → surface the web error.
                        Err(we) => return Err(we),
                    }
                } else {
                    return Err(e);
                }
            }
        };

        // Activate tokens immediately so follow-up requests work even if
        // session persistence fails below.
        self.client.set_login_identity(&email, resp.user_id);
        self.client.set_tokens(&resp.access_token, &resp.refresh_token);
        self.client.set_oauth_key(&oauth_key);

        if let Some(store) = self.client.session_store() {
            if !email.is_empty() {
                let session = Session {
                    email: email.clone(),
                    user_id: resp.user_id,
                    access_token: resp.access_token.clone(),
                    refresh_token: resp.refresh_token.clone(),
                    device_uuid,
                    oauth_key: self.client.session_oauth_key(),
                    updated_at: Utc::now(),
                };
                if let Err(e) = store.save(&session).await {
                    return Err(crate::errors::ApiError::SessionSaveFailed(
                        SessionSaveFailed {
                            response: resp,
                            source: e,
                        },
                    ));
                }
            }
        }

        Ok(resp)
    }
}

/// The request body sent to `POST /v3/users/login_with_email`.
///
/// This is the internal representation used by the auth wrapper;
/// the generated API client has its own type.
#[derive(Debug, Clone)]
pub(crate) struct LoginEmailUserRequestBody {
    pub email: String,
    pub password: String,
    pub api_key: String,
    pub uuid: String,
    pub two_fa_code: Option<String>,
}

/// Abbreviated login response. The full generated type has many more
/// fields, but the wrapper only needs these for session management.
/// When restoring from cache, only these fields are populated.
#[derive(Debug, Clone)]
pub struct LoginResponse {
    pub user_id: i64,
    pub access_token: String,
    pub refresh_token: String,
}

// Browser-login fallback (PORTING.md §6.2).
#[cfg(test)]
mod fallback_tests {
    use crate::client::Client;
    use crate::error_codes::ErrorCode;
    use crate::errors::code_of;
    use wiremock::matchers::{body_string_contains, header, method, path};
    use wiremock::{Mock, MockServer, ResponseTemplate};

    const WEB_KEY: &str =
        "e9f1ae4c4470f29a92c0168dc42b13637cc332692103f23e626bc2b016f66603";
    const APP_KEY: &str = crate::config::DEFAULT_API_KEY;

    fn reject_401(code: i32) -> ResponseTemplate {
        ResponseTemplate::new(401).set_body_string(format!(
            r#"{{"result":"error","message":"unauthorized","error_code":{code}}}"#
        ))
    }

    async fn mock_v3_reject(server: &MockServer, code: i32) {
        Mock::given(method("POST"))
            .and(path("/v3/users/login_with_email"))
            .respond_with(reject_401(code))
            .mount(server)
            .await;
    }

    #[tokio::test]
    async fn falls_back_to_web_on_401() {
        let server = MockServer::start().await;
        mock_v3_reject(&server, 401).await;
        // Matching on the web key asserts the fallback carried it.
        Mock::given(method("POST"))
            .and(path("/v2/users/login_with_email"))
            .and(body_string_contains(WEB_KEY))
            .respond_with(ResponseTemplate::new(200).set_body_string(
                r#"{"access_token":"WEBACC","refresh_token":"WEBREF","user_id":7777}"#,
            ))
            .mount(&server)
            .await;

        let client = Client::builder().base_url(server.uri()).build();
        let resp = client
            .login_with_email()
            .email("u@example.com")
            .password("pw")
            .execute()
            .await
            .expect("fallback login should succeed");
        assert_eq!(resp.user_id, 7777);
        assert_eq!(resp.access_token, "WEBACC");
        assert_eq!(client.tokens().access, "WEBACC");
    }

    #[tokio::test]
    async fn no_fallback_on_2fa() {
        let server = MockServer::start().await;
        mock_v3_reject(&server, ErrorCode::Required2FA as i32).await;
        // The v2 endpoint must never be hit.
        Mock::given(method("POST"))
            .and(path("/v2/users/login_with_email"))
            .respond_with(ResponseTemplate::new(200).set_body_string("{}"))
            .expect(0)
            .mount(&server)
            .await;

        let client = Client::builder().base_url(server.uri()).build();
        let err = client
            .login_with_email()
            .email("u@example.com")
            .password("pw")
            .execute()
            .await
            .expect_err("2FA-required must not fall back");
        assert_eq!(code_of(&err), ErrorCode::Required2FA);
        assert_eq!(client.tokens().access, "");
    }

    #[tokio::test]
    async fn no_fallback_when_api_key_pinned() {
        let server = MockServer::start().await;
        mock_v3_reject(&server, 401).await;
        Mock::given(method("POST"))
            .and(path("/v2/users/login_with_email"))
            .respond_with(ResponseTemplate::new(200).set_body_string("{}"))
            .expect(0)
            .mount(&server)
            .await;

        let client = Client::builder().base_url(server.uri()).build();
        let err = client
            .login_with_email()
            .email("u@example.com")
            .password("pw")
            .api_key("caller-pinned-key")
            .execute()
            .await
            .expect_err("pinned api_key must not fall back");
        assert_eq!(code_of(&err), ErrorCode::Unauthorized);
    }

    #[tokio::test]
    async fn fallback_disabled_surfaces_401() {
        let server = MockServer::start().await;
        mock_v3_reject(&server, 401).await;
        Mock::given(method("POST"))
            .and(path("/v2/users/login_with_email"))
            .respond_with(ResponseTemplate::new(200).set_body_string("{}"))
            .expect(0)
            .mount(&server)
            .await;

        let client = Client::builder()
            .base_url(server.uri())
            .web_api_key("")
            .build();
        let err = client
            .login_with_email()
            .email("u@example.com")
            .password("pw")
            .execute()
            .await
            .expect_err("disabled fallback must surface the primary 401");
        assert_eq!(code_of(&err), ErrorCode::Unauthorized);
    }

    #[tokio::test]
    async fn web_decode_failure_surfaced() {
        let server = MockServer::start().await;
        mock_v3_reject(&server, 401).await;
        Mock::given(method("POST"))
            .and(path("/v2/users/login_with_email"))
            .respond_with(
                ResponseTemplate::new(200).set_body_string("{ this is not valid json"),
            )
            .mount(&server)
            .await;

        let client = Client::builder().base_url(server.uri()).build();
        let err = client
            .login_with_email()
            .email("u@example.com")
            .password("pw")
            .execute()
            .await
            .expect_err("a 2xx decode failure should surface as an error");
        // Not masked as the primary auth rejection.
        assert_ne!(code_of(&err), ErrorCode::Unauthorized);
        assert_eq!(client.tokens().access, "");
    }

    // --- OAuth refresh uses the session's key (PORTING.md §6.2) --------------

    // Mount the timestamp endpoint to 401 once (triggering a refresh) then
    // 200, and the oauth endpoint to accept ONLY the given Basic key — so a
    // successful get_user_timestamp proves the refresh used that key.
    async fn mock_refresh_flow(server: &MockServer, expected_key: &str) {
        Mock::given(method("GET"))
            .and(path("/v2/users/timestamp"))
            .respond_with(ResponseTemplate::new(401).set_body_string(r#"{"error_code":-1}"#))
            .up_to_n_times(1)
            .with_priority(1)
            .mount(server)
            .await;
        Mock::given(method("GET"))
            .and(path("/v2/users/timestamp"))
            .respond_with(ResponseTemplate::new(200).set_body_string(r#"{"time":1700000000}"#))
            .with_priority(2)
            .mount(server)
            .await;
        Mock::given(method("POST"))
            .and(path("/api/v1/oauth/token"))
            .and(header(
                "authorization",
                format!("Basic {}", crate::transport::basic_credentials(expected_key)),
            ))
            .respond_with(ResponseTemplate::new(200).set_body_string(
                r#"{"access_token":"A2","refresh_token":"R2","id":1}"#,
            ))
            .mount(server)
            .await;
    }

    #[tokio::test]
    async fn web_session_refreshes_with_web_key() {
        let server = MockServer::start().await;
        mock_v3_reject(&server, 401).await;
        Mock::given(method("POST"))
            .and(path("/v2/users/login_with_email"))
            .respond_with(ResponseTemplate::new(200).set_body_string(
                r#"{"access_token":"A1","refresh_token":"R1","user_id":1}"#,
            ))
            .mount(&server)
            .await;
        mock_refresh_flow(&server, WEB_KEY).await;

        let client = Client::builder().base_url(server.uri()).build();
        client
            .login_with_email()
            .email("u@example.com")
            .password("pw")
            .execute()
            .await
            .expect("web login");
        assert_eq!(client.oauth_key_snapshot(), WEB_KEY);
        // 401 on the timestamp GET triggers a refresh; it succeeds only if
        // the oauth request carried the web key.
        let ts = client
            .get_user_timestamp()
            .await
            .expect("timestamp after web-key refresh");
        assert_eq!(ts.time.flatten().unwrap_or(0), 1700000000);
    }

    #[tokio::test]
    async fn app_session_refreshes_with_app_key() {
        let server = MockServer::start().await;
        Mock::given(method("POST"))
            .and(path("/v3/users/login_with_email"))
            .respond_with(ResponseTemplate::new(200).set_body_string(
                r#"{"access_token":"A1","refresh_token":"R1","user_id":1}"#,
            ))
            .mount(&server)
            .await;
        mock_refresh_flow(&server, APP_KEY).await;

        let client = Client::builder().base_url(server.uri()).build();
        client
            .login_with_email()
            .email("u@example.com")
            .password("pw")
            .execute()
            .await
            .expect("app login");
        assert_eq!(client.oauth_key_snapshot(), APP_KEY);
        let ts = client
            .get_user_timestamp()
            .await
            .expect("timestamp after app-key refresh");
        assert_eq!(ts.time.flatten().unwrap_or(0), 1700000000);
    }

    #[tokio::test]
    async fn restored_web_session_tracks_web_key() {
        use std::sync::Arc;
        let server = MockServer::start().await;
        mock_v3_reject(&server, 401).await;
        Mock::given(method("POST"))
            .and(path("/v2/users/login_with_email"))
            .respond_with(ResponseTemplate::new(200).set_body_string(
                r#"{"access_token":"A1","refresh_token":"R1","user_id":1}"#,
            ))
            .mount(&server)
            .await;

        let store = Arc::new(crate::session::MemorySessionStore::new());
        let c1 = Client::builder()
            .base_url(server.uri())
            .session_store_arc(store.clone())
            .build();
        c1.login_with_email()
            .email("u@example.com")
            .password("pw")
            .execute()
            .await
            .expect("web login");

        // A fresh client restores the cached web session and tracks the key.
        let c2 = Client::builder()
            .base_url(server.uri())
            .session_store_arc(store)
            .build();
        c2.login_with_email()
            .email("u@example.com")
            .password("pw")
            .execute()
            .await
            .expect("restore login");
        assert_eq!(c2.oauth_key_snapshot(), WEB_KEY);
    }
}
