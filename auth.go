package yaylib

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"

	"github.com/ekkx/yaylib/v2/gen"
)

// LoginWithEmailRequest is the fluent builder returned by
// (*Client).LoginWithEmail. Each field of the request body can be set
// directly on the builder (Email, Password, ...); Execute assembles the
// underlying request from them. Two conveniences are built in:
//
//   - APIKey and Uuid on the request body are auto-populated from the Client
//     if the caller left them blank.
//   - When a session store is active (WithSessionStore), Execute first tries
//     to restore a cached session keyed by the email, and on a successful
//     fresh login persists the new session automatically.
//
// Chain .NoCache() before Execute to force a fresh login regardless of
// cache contents.
//
// Why only LoginWithEmail has this wrapper:
//
// LoginWithEmail is the SDK's session entry point. The wrapper exists
// so that "log in with email/password and don't think about caching"
// is a one-liner. Other auth-flow body endpoints (CreateUser,
// EditUser, LoginWithSns, OauthToken / refresh) are intentionally NOT
// wrapped — they don't sit on the cache boundary, and adding a
// chain wrapper to each would multiply API surface without changing
// behavior. Reach those through the embedded service field directly,
// e.g. client.UsersAPIService.CreateUser(ctx)... or
// client.AuthAPIService.OauthToken(ctx)...
type LoginWithEmailRequest struct {
	client    *Client
	ctx       context.Context
	body      gen.LoginEmailUserRequest
	skipCache bool
}

// LoginWithEmail starts a login call. When a session store is active the
// resulting Execute will restore or persist tokens automatically, and APIKey
// / Uuid are auto-filled from the Client if omitted.
func (c *Client) LoginWithEmail(ctx context.Context) LoginWithEmailRequest {
	return LoginWithEmailRequest{client: c, ctx: ctx}
}

// Email sets the account email. Required.
func (r LoginWithEmailRequest) Email(v string) LoginWithEmailRequest {
	r.body.SetEmail(v)
	return r
}

// Password sets the account password. Required.
func (r LoginWithEmailRequest) Password(v string) LoginWithEmailRequest {
	r.body.SetPassword(v)
	return r
}

// ApiKey overrides the API key (defaults to the Client's APIKey).
func (r LoginWithEmailRequest) ApiKey(v string) LoginWithEmailRequest {
	r.body.SetApiKey(v)
	return r
}

// Uuid overrides the device UUID (defaults to the Client's DeviceUUID).
func (r LoginWithEmailRequest) Uuid(v string) LoginWithEmailRequest {
	r.body.SetUuid(v)
	return r
}

// TwoFACode sets the 2FA code when the account has it enabled.
func (r LoginWithEmailRequest) TwoFACode(v string) LoginWithEmailRequest {
	r.body.SetTwoFACode(v)
	return r
}

// NoCache forces a fresh HTTP login even if a cached session is available.
// Use it after changing credentials or when you need a fresh response
// payload instead of the abbreviated cached one.
func (r LoginWithEmailRequest) NoCache() LoginWithEmailRequest {
	r.skipCache = true
	return r
}

// Execute issues the login, transparently using the session store when one
// is configured. On a cache hit, the returned *http.Response is nil and the
// LoginUserResponse is synthesized from the cached tokens (UserId,
// AccessToken, RefreshToken populated; other fields zero-valued).
//
// Errors fall into three buckets:
//   - Session-store read failure (corrupt JSON, permission, I/O): returned
//     without attempting a fresh login, so a transient store fault doesn't
//     burn through the server's login rate limit.
//   - Login HTTP failure: returned as-is.
//   - Login succeeded but persisting the session failed: the response is
//     still returned so the tokens are usable; the persistence error is
//     wrapped with ErrSessionSaveFailed for callers that want to detect it.
func (r LoginWithEmailRequest) Execute() (*gen.LoginUserResponse, *http.Response, error) {
	body := r.client.fillEmailLogin(r.body)
	email := body.GetEmail()
	resp, ok, err := r.client.maybeRestore(email, r.skipCache)
	if err != nil {
		return nil, nil, err
	}
	if ok {
		return resp, nil, nil
	}
	resp, httpResp, err := r.client.UsersAPIService.
		LoginWithEmail(r.ctx).
		LoginEmailUserRequest(body).
		Execute()
	if err != nil {
		// Fallback: when the primary (app) login is rejected with a plain
		// "unauthorized", retry against the browser login endpoint, which
		// authenticates with a different key (WebAPIKey). On success the
		// web response replaces the failed one; if the fallback also fails,
		// the original error (a proper *APIError) is surfaced unchanged.
		//
		// The fallback is gated on ErrCodeUnauthorized specifically, not on
		// a bare 401 status: other rejections are also 401 but must NOT be
		// retried — notably ErrCodeRequired2FA (the caller has to supply a
		// code, not switch keys). It is also skipped when the caller pinned
		// a specific api_key via .ApiKey(), since the fallback would
		// otherwise silently authenticate under a different key.
		callerPinnedKey := body.GetApiKey() != r.client.APIKey
		if r.client.WebAPIKey != "" && !callerPinnedKey && CodeOf(err) == ErrCodeUnauthorized {
			if wResp, wHTTP, wErr := r.client.loginWithEmailWeb(r.ctx, body); wErr == nil {
				r.client.Logger.Debug("web login fallback ok",
					"event", "login_fallback", "outcome", "ok")
				resp, httpResp, err = wResp, wHTTP, nil
			} else {
				r.client.Logger.Debug("web login fallback failed",
					"event", "login_fallback", "outcome", "failed")
				// A 2xx web response we couldn't decode is a genuine
				// login result, not an auth rejection: surface it (with
				// its raw body, readable on wHTTP) rather than masking it
				// behind the primary 401. Non-2xx / transport errors fall
				// through and the primary error is returned unchanged.
				if wHTTP != nil && wHTTP.StatusCode >= 200 && wHTTP.StatusCode < 300 {
					return nil, wHTTP, wErr
				}
			}
		}
		if err != nil {
			return resp, httpResp, err
		}
	}
	if perr := r.client.acceptLogin(email, resp); perr != nil {
		return resp, httpResp, perr
	}
	return resp, httpResp, nil
}

// ---- internal helpers ----

// bareClient returns an *http.Client that sends requests through the base
// RoundTripper WITHOUT the app transport's header injection, Authorization,
// 401 auto-refresh or retry. It inherits the configured client's timeout
// and, when the app transport is in place, its underlying (caller-supplied
// or default) RoundTripper — so a custom proxy / TLS config still applies.
func (c *Client) bareClient() *http.Client {
	bare := &http.Client{Timeout: c.httpClient.Timeout}
	if t, ok := c.httpClient.Transport.(*Transport); ok {
		bare.Transport = t.Base
	}
	return bare
}

// loginWithEmailWeb is the fallback login path invoked when the primary
// email login returns 401. It posts to /v2/users/login_with_email — the
// endpoint the browser client uses — authenticating with the client's
// WebAPIKey instead of APIKey.
//
// The request is sent through the UNWRAPPED HTTP layer (the base
// RoundTripper, bypassing the app transport) for two reasons: it must
// carry the browser client's minimal header set rather than the app's
// (X-App-Version / X-Connection-* / X-Device-UUID), and it must not
// pick up the transport's Authorization injection or 401 auto-refresh —
// a login must authenticate purely by its body credentials, never by a
// pre-existing session's Bearer. This mirrors the "re-issue with the
// unwrapped client" rule the retry/refresh paths follow (PORTING §6.1).
// A caller-supplied *http.Client / proxy is still honored: the base
// RoundTripper and timeout are inherited from it. Only the login body's
// api_key and the endpoint version differ from the app login.
func (c *Client) loginWithEmailWeb(ctx context.Context, body gen.LoginEmailUserRequest) (*gen.LoginUserResponse, *http.Response, error) {
	payload := map[string]string{
		"email":    body.GetEmail(),
		"password": body.GetPassword(),
		"api_key":  c.WebAPIKey,
		"uuid":     body.GetUuid(),
	}
	if code := body.GetTwoFACode(); code != "" {
		payload["two_fa_code"] = code
	}
	buf, err := json.Marshal(payload)
	if err != nil {
		return nil, nil, err
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost,
		c.BaseURL+"/v2/users/login_with_email", bytes.NewReader(buf))
	if err != nil {
		return nil, nil, err
	}
	req.Header.Set("Content-Type", "application/json;charset=UTF-8")
	// Present as the browser client. The web version string is not
	// otherwise known to the SDK, so AppVersion stands in for it.
	const webUA = "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) " +
		"AppleWebKit/537.36 (KHTML, like Gecko) Chrome/120.0.0.0 Safari/537.36"
	req.Header.Set("User-Agent", webUA)
	req.Header.Set("Agent", "YayWeb "+c.AppVersion)
	req.Header.Set("X-Device-Info", fmt.Sprintf("Yay %s Web (%s)", c.AppVersion, webUA))

	httpResp, err := c.bareClient().Do(req)
	if err != nil {
		return nil, nil, err
	}
	respBody, err := io.ReadAll(httpResp.Body)
	httpResp.Body.Close()
	if err != nil {
		return nil, httpResp, err
	}
	// Restore a readable body on the response so callers / middleware that
	// inspect the returned *http.Response see the same bytes we decoded,
	// matching the primary login path's semantics.
	httpResp.Body = io.NopCloser(bytes.NewReader(respBody))
	if httpResp.StatusCode < 200 || httpResp.StatusCode >= 300 {
		return nil, httpResp, fmt.Errorf("yaylib: web login failed: status %d", httpResp.StatusCode)
	}
	var out gen.LoginUserResponse
	if err := json.Unmarshal(respBody, &out); err != nil {
		return nil, httpResp, fmt.Errorf("yaylib: web login decode: %w", err)
	}
	return &out, httpResp, nil
}

// fillEmailLogin populates APIKey and Uuid from the client if they weren't
// set by the caller, and returns the completed body.
func (c *Client) fillEmailLogin(body gen.LoginEmailUserRequest) gen.LoginEmailUserRequest {
	if body.GetApiKey() == "" {
		body.SetApiKey(c.APIKey)
	}
	if body.GetUuid() == "" {
		body.SetUuid(c.deviceUUIDSnapshot())
	}
	return body
}

// ErrSessionSaveFailed wraps a session-store write failure that occurred
// after an otherwise successful login. The login response is still
// returned alongside this error, so callers can choose to log-and-ignore
// (errors.Is(err, ErrSessionSaveFailed)) or treat it as fatal.
var ErrSessionSaveFailed = errors.New("yaylib: session save failed")

// maybeRestore checks the session store for a cached session and, on hit,
// applies it to the client and returns a synthesized response. ok=false /
// err=nil means "no cached session, proceed with fresh login"; a non-nil
// err signals the store itself is unhealthy (corrupt JSON, permission,
// I/O), in which case the caller MUST NOT fall through to a fresh login —
// otherwise a transient store fault burns through the login rate limit.
func (c *Client) maybeRestore(email string, skip bool) (*gen.LoginUserResponse, bool, error) {
	if skip || c.sessionStore == nil || email == "" {
		return nil, false, nil
	}
	sess, err := c.LoadSession(email)
	if errors.Is(err, ErrNoSession) {
		return nil, false, nil
	}
	if err != nil {
		return nil, false, fmt.Errorf("yaylib: load session: %w", err)
	}
	resp := gen.NewLoginUserResponse()
	resp.SetUserId(sess.UserID)
	resp.SetAccessToken(sess.AccessToken)
	resp.SetRefreshToken(sess.RefreshToken)
	return resp, true, nil
}

// acceptLogin is called after a successful fresh login. It activates the
// tokens on the client, records the account email (used by the 401 auto-
// refresh flow to key session-store writes), and persists the session when
// a store is configured. Persistence errors are wrapped with
// ErrSessionSaveFailed so the caller can detect them, but the tokens are
// already active on the client either way.
func (c *Client) acceptLogin(email string, resp *gen.LoginUserResponse) error {
	c.setLoginIdentity(email, resp.GetUserId())
	// Activate tokens unconditionally so an immediate follow-up request
	// works even when a later session-store write fails.
	c.SetTokens(resp.GetAccessToken(), resp.GetRefreshToken())
	if c.sessionStore != nil && email != "" {
		if err := c.SaveSession(&Session{
			Email:        email,
			UserID:       resp.GetUserId(),
			AccessToken:  resp.GetAccessToken(),
			RefreshToken: resp.GetRefreshToken(),
		}); err != nil {
			return fmt.Errorf("%w: %v", ErrSessionSaveFailed, err)
		}
	}
	return nil
}

// refreshTokens exchanges the current refresh token for a fresh access /
// refresh pair via POST /api/v1/oauth/token (grant_type=refresh_token).
//
// It is safe to call concurrently: only one refresh runs at a time, and a
// caller that arrives with a staleToken matching an access token that has
// already been rotated is a no-op (the new token is read via Tokens()).
//
// When a session store is configured and an account email is known, the
// refreshed tokens are persisted so they survive process restarts.
func (c *Client) refreshTokens(ctx context.Context, staleToken string) error {
	c.refreshMu.Lock()
	defer c.refreshMu.Unlock()

	// Another goroutine already rotated the access token — use theirs.
	snap := c.Tokens()
	if snap.Access != staleToken {
		return nil
	}
	currentRefresh := snap.Refresh
	if currentRefresh == "" {
		c.Logger.Debug("token refresh skipped",
			"event", "token_refresh", "outcome", "no_token")
		return nil
	}

	resp, _, err := c.OauthToken(ctx).
		GrantType("refresh_token").
		RefreshToken(currentRefresh).
		Execute()
	if err != nil {
		c.Logger.Debug("token refresh failed",
			"event", "token_refresh", "outcome", "failed")
		return err
	}

	c.SetTokens(resp.GetAccessToken(), resp.GetRefreshToken())
	c.Logger.Debug("token refresh ok",
		"event", "token_refresh", "outcome", "ok")

	if email := c.currentEmailSnapshot(); c.sessionStore != nil && email != "" {
		if err := c.SaveSession(&Session{
			Email:        email,
			UserID:       resp.GetId(),
			AccessToken:  resp.GetAccessToken(),
			RefreshToken: resp.GetRefreshToken(),
		}); err != nil {
			// The transport path can't bubble this error to the
			// user — the refreshed tokens are already active in
			// memory and the retried request is what they care
			// about. Surface persist failure via the logger so it
			// stays observable; on next process start the cached
			// refresh token will be the now-stale one and a fresh
			// login will be needed.
			c.Logger.Warn("persist refreshed tokens failed",
				"event", "token_persist_fail",
				"user_id", resp.GetId(), "err", err.Error())
		}
	}
	return nil
}
