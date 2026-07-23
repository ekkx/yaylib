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
        // default). A pinned key opts out of the web-first path and logs in
        // through the app endpoint with exactly that key.
        let caller_pinned = api_key.as_str() != self.client.api_key();

        // Build the login request body.
        let body = LoginEmailUserRequestBody {
            email: email.clone(),
            password: self.password.unwrap_or_default(),
            api_key,
            uuid: device_uuid.clone(),
            two_fa_code: self.two_fa_code,
        };

        // Web-first fresh login (PORTING.md §6.2). do_fresh_login runs the
        // browser endpoint as primary and the app endpoint as fallback,
        // reporting which key established the session so token refresh uses it.
        let (resp, oauth_key) = self.client.do_fresh_login(body, caller_pinned).await?;

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

// Web-first login (PORTING.md §6.2): the browser endpoint is primary, the
// app endpoint is the fallback.
#[cfg(test)]
mod web_first_tests {
    use std::sync::Arc;

    use crate::client::Client;
    use crate::error_codes::ErrorCode;
    use crate::errors::{code_of, error_response_of, ApiError};
    use crate::session::{MemorySessionStore, Session, SessionStore};
    use wiremock::matchers::{body_string_contains, header, method, path};
    use wiremock::{Mock, MockServer, ResponseTemplate};

    const WEB_KEY: &str = crate::config::DEFAULT_WEB_API_KEY;
    const APP_KEY: &str = crate::config::DEFAULT_API_KEY;

    fn reject(code: i32) -> ResponseTemplate {
        ResponseTemplate::new(401).set_body_string(format!(
            r#"{{"result":"error","message":"unauthorized","error_code":{code}}}"#
        ))
    }

    // Browser (v2) login succeeds, returning a distinct user_id/tokens.
    async fn mount_web_ok(server: &MockServer) {
        Mock::given(method("POST"))
            .and(path("/v2/users/login_with_email"))
            .and(body_string_contains(WEB_KEY))
            .respond_with(ResponseTemplate::new(200).set_body_string(
                r#"{"access_token":"WEBACC","refresh_token":"WEBREF","user_id":7777}"#,
            ))
            .mount(server)
            .await;
    }

    // The browser endpoint is primary: it succeeds directly and the app (v3)
    // endpoint is never touched.
    #[tokio::test]
    async fn web_primary_succeeds() {
        let server = MockServer::start().await;
        mount_web_ok(&server).await;
        // The v3 endpoint must never be hit.
        Mock::given(method("POST"))
            .and(path("/v3/users/login_with_email"))
            .respond_with(ResponseTemplate::new(200).set_body_string("{}"))
            .expect(0)
            .mount(&server)
            .await;

        let client = Client::builder().base_url(server.uri()).build();
        let resp = client
            .login_with_email()
            .email("u@example.com")
            .password("pw")
            .execute()
            .await
            .expect("web primary login should succeed");
        assert_eq!(resp.user_id, 7777, "want the web response");
        assert_eq!(client.tokens().access, "WEBACC");
        assert_eq!(client.tokens().refresh, "WEBREF");
        assert_eq!(
            client.oauth_key_snapshot(),
            WEB_KEY,
            "session must track the web key"
        );
    }

    // A primary web 401 falls back to the app (v3) endpoint under the app key;
    // the session then tracks the app key for refresh. Both endpoints are hit.
    #[tokio::test]
    async fn falls_back_to_app_on_401() {
        let server = MockServer::start().await;
        Mock::given(method("POST"))
            .and(path("/v2/users/login_with_email"))
            .respond_with(reject(401))
            .expect(1)
            .mount(&server)
            .await;
        Mock::given(method("POST"))
            .and(path("/v3/users/login_with_email"))
            .and(body_string_contains(APP_KEY))
            .respond_with(ResponseTemplate::new(200).set_body_string(
                r#"{"access_token":"APPACC","refresh_token":"APPREF","user_id":4242}"#,
            ))
            .expect(1)
            .mount(&server)
            .await;

        let client = Client::builder().base_url(server.uri()).build();
        let resp = client
            .login_with_email()
            .email("u@example.com")
            .password("pw")
            .execute()
            .await
            .expect("app fallback should succeed");
        assert_eq!(resp.user_id, 4242, "want the app response");
        assert_eq!(client.tokens().access, "APPACC");
        assert_eq!(
            client.oauth_key_snapshot(),
            APP_KEY,
            "session must track the app key after fallback"
        );
    }

    // With the web path switched off (web_api_key==""), login goes straight to
    // the app endpoint and no v2 request is attempted; an app 401 surfaces.
    #[tokio::test]
    async fn web_disabled_uses_app_only() {
        let server = MockServer::start().await;
        // The v2 endpoint must never be hit.
        Mock::given(method("POST"))
            .and(path("/v2/users/login_with_email"))
            .respond_with(ResponseTemplate::new(200).set_body_string("{}"))
            .expect(0)
            .mount(&server)
            .await;
        Mock::given(method("POST"))
            .and(path("/v3/users/login_with_email"))
            .respond_with(reject(401))
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
            .expect_err("app 401 must surface when web is disabled");
        assert_eq!(code_of(&err), ErrorCode::Unauthorized);
        assert_eq!(client.tokens().access, "", "tokens must stay empty");
    }

    // When both the primary web login and the app fallback fail, the primary
    // (web) 401 is surfaced.
    #[tokio::test]
    async fn both_fail_surfaces_web_401() {
        let server = MockServer::start().await;
        Mock::given(method("POST"))
            .and(path("/v2/users/login_with_email"))
            .respond_with(reject(401))
            .mount(&server)
            .await;
        Mock::given(method("POST"))
            .and(path("/v3/users/login_with_email"))
            .respond_with(reject(401))
            .mount(&server)
            .await;

        let client = Client::builder().base_url(server.uri()).build();
        let err = client
            .login_with_email()
            .email("u@example.com")
            .password("pw")
            .execute()
            .await
            .expect_err("both logins failing must surface an error");
        assert_eq!(code_of(&err), ErrorCode::Unauthorized);
        assert_eq!(client.tokens().access, "");
    }

    // A 2FA-required rejection is also HTTP 401 but carries Required2FA; the
    // app fallback must NOT fire (the caller supplies a code, not a different
    // endpoint), so the web login's 2FA surfaces and v3 is never called.
    #[tokio::test]
    async fn no_fallback_on_2fa() {
        let server = MockServer::start().await;
        Mock::given(method("POST"))
            .and(path("/v2/users/login_with_email"))
            .respond_with(reject(ErrorCode::Required2FA as i32))
            .mount(&server)
            .await;
        Mock::given(method("POST"))
            .and(path("/v3/users/login_with_email"))
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

    // When the caller pins a specific api_key via .api_key(), the web-first
    // path is skipped (it would otherwise authenticate under web_api_key), so
    // login goes straight to the app endpoint under the pinned key and v2 is
    // never called.
    #[tokio::test]
    async fn no_fallback_when_api_key_pinned() {
        let server = MockServer::start().await;
        // The v2 endpoint must never be hit.
        Mock::given(method("POST"))
            .and(path("/v2/users/login_with_email"))
            .respond_with(ResponseTemplate::new(200).set_body_string("{}"))
            .expect(0)
            .mount(&server)
            .await;
        // The app endpoint sees exactly the caller-pinned key.
        Mock::given(method("POST"))
            .and(path("/v3/users/login_with_email"))
            .and(body_string_contains("caller-pinned-key"))
            .respond_with(reject(401))
            .expect(1)
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
            .expect_err("pinned api_key must not fall back to web");
        assert_eq!(code_of(&err), ErrorCode::Unauthorized);
    }

    // A primary web response that is 2xx but undecodable is a genuine login
    // result: its raw body rides on a typed APIError. Since it is not
    // Unauthorized the app fallback does not fire.
    #[tokio::test]
    async fn web_2xx_decode_failure_surfaced() {
        let server = MockServer::start().await;
        Mock::given(method("POST"))
            .and(path("/v2/users/login_with_email"))
            .respond_with(ResponseTemplate::new(200).set_body_string("{ this is not valid json"))
            .mount(&server)
            .await;
        // No app fallback on a 2xx decode failure.
        Mock::given(method("POST"))
            .and(path("/v3/users/login_with_email"))
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
            .expect_err("a 2xx decode failure should surface as an error");
        // Not masked as an auth rejection.
        assert_ne!(code_of(&err), ErrorCode::Unauthorized);
        // Typed APIError carrying the raw web body.
        match &err {
            ApiError::Response { status, body, .. } => {
                assert_eq!(*status, 200, "want the 2xx web response surfaced");
                assert!(
                    String::from_utf8_lossy(body).contains("not valid json"),
                    "raw web body must be attached, got {:?}",
                    String::from_utf8_lossy(body)
                );
            }
            other => panic!("want a typed ApiError::Response, got {other:?}"),
        }
        assert_eq!(client.tokens().access, "");
    }

    // A primary web login rejected with a NON-Unauthorized typed error (a
    // temporary ban carrying ban_until) surfaces directly — no app fallback
    // (it is not Unauthorized), and error_response_of / matching read the
    // parsed fields and raw body.
    #[tokio::test]
    async fn web_primary_typed_error_surfaced() {
        let server = MockServer::start().await;
        Mock::given(method("POST"))
            .and(path("/v2/users/login_with_email"))
            .respond_with(ResponseTemplate::new(403).set_body_string(format!(
                r#"{{"result":"error","message":"temporarily banned","error_code":{},"ban_until":1893456000}}"#,
                ErrorCode::UserTemporaryBanned as i32
            )))
            .mount(&server)
            .await;
        Mock::given(method("POST"))
            .and(path("/v3/users/login_with_email"))
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
            .expect_err("the web ban error must surface");
        assert_eq!(code_of(&err), ErrorCode::UserTemporaryBanned);
        let parsed = error_response_of(&err).expect("ban error must parse an error response");
        assert_eq!(
            parsed.ban_until,
            Some(1893456000),
            "error_response_of must read ban_until from the web error"
        );
        match &err {
            ApiError::Response { body, .. } => assert!(
                String::from_utf8_lossy(body).contains("temporarily banned"),
                "raw web body must be accessible"
            ),
            other => panic!("want a typed ApiError::Response, got {other:?}"),
        }
    }

    // signed_info is keyed on api_key, NOT the session's OAuth key. Under
    // web-first a web session (oauth key = web_api_key) is the common case, so
    // this app-key/OAuth-key split must hold: setting the OAuth key to the web
    // key must not change the signature.
    #[test]
    fn signed_info_ignores_oauth_key() {
        let reference = Client::builder()
            .api_key("app-key")
            .device_uuid("uuid-shared")
            .build();

        let web = Client::builder()
            .api_key("app-key")
            .device_uuid("uuid-shared")
            .build();
        web.set_oauth_key("web-key"); // establish a web session

        assert_eq!(
            web.generate_signed_info_at(1700000000).value,
            reference.generate_signed_info_at(1700000000).value,
            "signed_info must stay on api_key, not the OAuth key"
        );
    }

    // --- OAuth refresh uses the session's establishing key (PORTING.md §6.2)

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

    // A web session refreshes its tokens with the web key (Basic).
    #[tokio::test]
    async fn web_session_refreshes_with_web_key() {
        let server = MockServer::start().await;
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
        // 401 on the timestamp GET triggers a refresh; it succeeds only if the
        // oauth request carried the web key.
        let ts = client
            .get_user_timestamp()
            .await
            .expect("timestamp after web-key refresh");
        assert_eq!(ts.time.flatten().unwrap_or(0), 1700000000);
    }

    // An app login (reached by disabling the web-first path with
    // web_api_key("")) refreshes with the app key, and does not persist an
    // oauth_key (app sessions stay byte-identical on disk).
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

        let store = Arc::new(MemorySessionStore::new());
        let client = Client::builder()
            .base_url(server.uri())
            .session_store_arc(store.clone())
            .web_api_key("")
            .build();
        client
            .login_with_email()
            .email("u@example.com")
            .password("pw")
            .execute()
            .await
            .expect("app login");
        assert_eq!(client.oauth_key_snapshot(), APP_KEY);
        // App sessions omit oauth_key on disk.
        let sess = store.load("u@example.com").await.expect("session persisted");
        assert_eq!(sess.oauth_key, "", "app session must persist an empty oauth_key");

        let ts = client
            .get_user_timestamp()
            .await
            .expect("timestamp after app-key refresh");
        assert_eq!(ts.time.flatten().unwrap_or(0), 1700000000);
    }

    // A cached web session, restored on a fresh client via the cached-login
    // path, still refreshes with the web key — the oauth_key round-trips
    // through the session store.
    #[tokio::test]
    async fn restored_web_session_tracks_web_key() {
        let server = MockServer::start().await;
        Mock::given(method("POST"))
            .and(path("/v2/users/login_with_email"))
            .respond_with(ResponseTemplate::new(200).set_body_string(
                r#"{"access_token":"A1","refresh_token":"R1","user_id":1}"#,
            ))
            .mount(&server)
            .await;

        let store = Arc::new(MemorySessionStore::new());
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
        let persisted = store.load("u@example.com").await.expect("persisted");
        assert_eq!(persisted.oauth_key, WEB_KEY, "web key must round-trip on disk");

        // A fresh client restores the cached web session (cache hit) and
        // tracks the key.
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

    // A web session restored directly with the PUBLIC load_session (not the
    // cached-login path) must also refresh with the web key. This is the
    // regression the oauth_key restore fixes: load_session used to apply
    // tokens, device UUID, and identity but not the establishing key, so a web
    // session loaded straight from disk refreshed with the app key and 401'd.
    #[tokio::test]
    async fn load_session_restores_web_oauth_key() {
        let server = MockServer::start().await;
        mock_refresh_flow(&server, WEB_KEY).await;

        // A web session persisted to the store (oauth_key == web key).
        let store = Arc::new(MemorySessionStore::new());
        store
            .save(&Session {
                email: "u@example.com".into(),
                user_id: 1,
                access_token: "A1".into(),
                refresh_token: "R1".into(),
                device_uuid: String::new(),
                oauth_key: WEB_KEY.into(),
                updated_at: chrono::Utc::now(),
            })
            .await
            .expect("seed session");

        let client = Client::builder()
            .base_url(server.uri())
            .session_store_arc(store)
            .build();
        client
            .load_session("u@example.com")
            .await
            .expect("load_session");
        assert_eq!(
            client.oauth_key_snapshot(),
            WEB_KEY,
            "load_session must restore the web key"
        );
        // A refresh (triggered by the 401 timestamp) must carry the web key.
        let ts = client
            .get_user_timestamp()
            .await
            .expect("timestamp after web-key refresh via load_session");
        assert_eq!(ts.time.flatten().unwrap_or(0), 1700000000);
        assert_eq!(client.tokens().access, "A2");
    }

    // App sessions carry an empty oauth_key; load_session must leave refresh on
    // the app key (empty → app-key default).
    #[tokio::test]
    async fn load_session_app_session_keeps_app_key() {
        let store = Arc::new(MemorySessionStore::new());
        store
            .save(&Session {
                email: "u@example.com".into(),
                user_id: 1,
                access_token: "A1".into(),
                refresh_token: "R1".into(),
                device_uuid: String::new(),
                oauth_key: String::new(), // an app session
                updated_at: chrono::Utc::now(),
            })
            .await
            .expect("seed session");

        let client = Client::builder().session_store_arc(store).build();
        client
            .load_session("u@example.com")
            .await
            .expect("load_session");
        assert_eq!(client.oauth_key_snapshot(), APP_KEY);
    }

    // A web session injected directly via save_session (not through login) also
    // activates its establishing key, so its refresh uses the web key. Mirrors
    // load_session — save_session is a full activation, not just a write.
    #[tokio::test]
    async fn save_session_activates_web_oauth_key() {
        let server = MockServer::start().await;
        mock_refresh_flow(&server, WEB_KEY).await;

        let client = Client::builder()
            .base_url(server.uri())
            .session_store_arc(Arc::new(MemorySessionStore::new()))
            .build();
        client
            .save_session(&mut Session {
                email: "u@example.com".into(),
                user_id: 1,
                access_token: "A1".into(),
                refresh_token: "R1".into(),
                device_uuid: String::new(),
                oauth_key: WEB_KEY.into(),
                updated_at: chrono::Utc::now(),
            })
            .await
            .expect("save_session");
        assert_eq!(
            client.oauth_key_snapshot(),
            WEB_KEY,
            "save_session must activate the web key"
        );
        let ts = client
            .get_user_timestamp()
            .await
            .expect("timestamp after web-key refresh via save_session");
        assert_eq!(ts.time.flatten().unwrap_or(0), 1700000000);
    }
}
