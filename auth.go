package yaylib

import (
	"context"
	"net/http"

	"github.com/ekkx/yaylib/gen"
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
func (r LoginWithEmailRequest) Execute() (*gen.LoginUserResponse, *http.Response, error) {
	body := r.client.fillEmailLogin(r.body)
	email := body.GetEmail()
	if resp, ok := r.client.maybeRestore(email, r.skipCache); ok {
		return resp, nil, nil
	}
	resp, httpResp, err := r.client.UsersAPIService.
		LoginWithEmail(r.ctx).
		LoginEmailUserRequest(body).
		Execute()
	if err != nil {
		return resp, httpResp, err
	}
	r.client.acceptLogin(email, resp)
	return resp, httpResp, err
}

// ---- internal helpers ----

// fillEmailLogin populates APIKey and Uuid from the client if they weren't
// set by the caller, and returns the completed body.
func (c *Client) fillEmailLogin(body gen.LoginEmailUserRequest) gen.LoginEmailUserRequest {
	if body.GetApiKey() == "" {
		body.SetApiKey(c.APIKey)
	}
	if body.GetUuid() == "" {
		body.SetUuid(c.DeviceUUID)
	}
	return body
}

// maybeRestore checks the session store for a cached session and, on hit,
// applies it to the client and returns a synthesized response. Returns
// ok=false when skipped, missing, or on store error.
func (c *Client) maybeRestore(email string, skip bool) (*gen.LoginUserResponse, bool) {
	if skip || c.sessionDB == nil || email == "" {
		return nil, false
	}
	sess, err := c.LoadSession(email)
	if err != nil {
		return nil, false
	}
	resp := gen.NewLoginUserResponse()
	resp.SetUserId(sess.UserID)
	resp.SetAccessToken(sess.AccessToken)
	resp.SetRefreshToken(sess.RefreshToken)
	return resp, true
}

// acceptLogin is called after a successful fresh login. It activates the
// tokens on the client and, if a session store is configured, persists them.
func (c *Client) acceptLogin(email string, resp *gen.LoginUserResponse) {
	if c.sessionDB != nil && email != "" {
		_ = c.SaveSession(&Session{
			Email:        email,
			UserID:       resp.GetUserId(),
			AccessToken:  resp.GetAccessToken(),
			RefreshToken: resp.GetRefreshToken(),
		})
		return
	}
	c.SetTokens(resp.GetAccessToken(), resp.GetRefreshToken())
}
