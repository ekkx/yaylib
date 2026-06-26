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

        // Build and execute the HTTP login request.
        let body = LoginEmailUserRequestBody {
            email: email.clone(),
            password: self.password.unwrap_or_default(),
            api_key,
            uuid: device_uuid.clone(),
            two_fa_code: self.two_fa_code,
        };

        let resp = self.client.do_login_with_email(body).await?;

        // Activate tokens immediately so follow-up requests work even if
        // session persistence fails below.
        self.client.set_login_identity(&email, resp.user_id);
        self.client.set_tokens(&resp.access_token, &resp.refresh_token);

        if let Some(store) = self.client.session_store() {
            if !email.is_empty() {
                let session = Session {
                    email: email.clone(),
                    user_id: resp.user_id,
                    access_token: resp.access_token.clone(),
                    refresh_token: resp.refresh_token.clone(),
                    device_uuid,
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
