package yaylib

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
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
		return resp, httpResp, err
	}
	if perr := r.client.acceptLogin(email, resp); perr != nil {
		return resp, httpResp, perr
	}
	return resp, httpResp, nil
}

// ---- internal helpers ----

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
// already been rotated is a no-op (the new token is in c.Tokens).
//
// When a session store is configured and an account email is known, the
// refreshed tokens are persisted so they survive process restarts.
func (c *Client) refreshTokens(ctx context.Context, staleToken string) error {
	c.refreshMu.Lock()
	defer c.refreshMu.Unlock()

	// Another goroutine already rotated the access token — use theirs.
	if c.accessSnapshot() != staleToken {
		return nil
	}
	currentRefresh := c.refreshSnapshot()
	if currentRefresh == "" {
		return nil
	}

	resp, _, err := c.OauthToken(ctx).
		GrantType("refresh_token").
		RefreshToken(currentRefresh).
		Execute()
	if err != nil {
		return err
	}

	c.SetTokens(resp.GetAccessToken(), resp.GetRefreshToken())

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
			// about. Surface persist failure via slog so it stays
			// observable; on next process start the cached refresh
			// token will be the now-stale one and a fresh login
			// will be needed.
			slog.Warn("yaylib: persist refreshed tokens failed",
				"email", email, "err", err)
		}
	}
	return nil
}
