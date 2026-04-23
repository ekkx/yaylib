package yaylib

import (
	"bytes"
	"context"
	"encoding/base64"
	"io"
	"net/http"
	"strconv"
	"strings"
	"time"
)

// TokenStore holds the active OAuth credentials. The Transport reads Access
// for every request; Refresh is used for automatic token refresh on 401.
type TokenStore struct {
	Access  string
	Refresh string
}

// Transport is an http.RoundTripper that injects every header the Yay!
// servers expect from the official Android app, picks the right
// Authorization scheme (Basic for /api/v1/oauth/token*, Bearer otherwise),
// and transparently refreshes the access token on a 401 response.
//
// It is not intended for direct use by callers; NewClient wires one in
// front of the client's http.Client.
type Transport struct {
	Base   http.RoundTripper
	client *Client
}

func (t *Transport) RoundTrip(r *http.Request) (*http.Response, error) {
	r = r.Clone(r.Context())

	// Buffer the request body so a 401 retry can re-send the same payload.
	// This reads the body into memory once; Yay! requests are small so the
	// cost is negligible.
	var bodyBytes []byte
	if r.Body != nil {
		b, err := io.ReadAll(r.Body)
		if err != nil {
			return nil, err
		}
		r.Body.Close()
		bodyBytes = b
		r.Body = io.NopCloser(bytes.NewReader(bodyBytes))
	}

	// Snapshot the access token actually being sent so refreshTokens can
	// detect when another goroutine has already rotated it.
	sentAccess := ""
	if t.client.Tokens != nil {
		sentAccess = t.client.Tokens.Access
	}

	t.setHeaders(r, sentAccess)

	resp, err := t.Base.RoundTrip(r)
	if err != nil {
		return resp, err
	}

	// Trigger async client-IP fetch on the first request where we don't
	// already have one — skip the fetch endpoint itself to avoid recursion.
	if !isTimestampPath(r.URL.Path) {
		t.client.maybeFetchClientIP()
	}

	// 401 auto-refresh. The oauth/token endpoint itself is excluded so a
	// failed refresh doesn't loop.
	if resp.StatusCode == http.StatusUnauthorized && !isOAuthTokenPath(r.URL.Path) {
		return t.handle401(resp, r, bodyBytes, sentAccess), nil
	}

	return resp, err
}

// setHeaders applies the Yay! required headers and the correct Authorization
// scheme. Caller-provided headers win (set-if-absent), except X-Timestamp
// and Content-Type which are always normalized.
func (t *Transport) setHeaders(r *http.Request, accessToken string) {
	c := t.client
	setIfAbsent := func(k, v string) {
		if r.Header.Get(k) == "" {
			r.Header.Set(k, v)
		}
	}
	setIfAbsent("User-Agent", c.UserAgent)
	setIfAbsent("X-App-Version", c.APIVersionName)
	setIfAbsent("X-Device-Info", c.DeviceInfo)
	setIfAbsent("X-Device-UUID", c.DeviceUUID)
	setIfAbsent("X-Client-IP", c.getClientIP())
	setIfAbsent("X-Connection-Type", c.ConnectionType)
	setIfAbsent("X-Connection-Speed", c.ConnectionSpeed)
	setIfAbsent("Accept-Language", c.AcceptLanguage)

	// X-Timestamp must be fresh per request — don't let a caller pin a
	// stale value.
	r.Header.Set("X-Timestamp", strconv.FormatInt(time.Now().Unix(), 10))

	// Normalize any JSON Content-Type to match the server's exact casing /
	// spacing expectations.
	if ct := r.Header.Get("Content-Type"); strings.HasPrefix(ct, "application/json") {
		r.Header.Set("Content-Type", "application/json;charset=UTF-8")
	}

	// Authorization: Basic base64(apiKey) for oauth/token endpoints,
	// Bearer <access> for everything else.
	if r.Header.Get("Authorization") == "" {
		if isOAuthTokenPath(r.URL.Path) {
			r.Header.Set("Authorization", "Basic "+basicCredentials(c.APIKey))
		} else if accessToken != "" {
			r.Header.Set("Authorization", "Bearer "+accessToken)
		}
	}
}

// handle401 buffers the 401 response body, attempts a token refresh, and
// retries the original request once with the new Authorization. On any
// failure path it restores and returns the original 401 response so the
// caller still sees the server's error body.
func (t *Transport) handle401(orig *http.Response, req *http.Request, reqBody []byte, staleToken string) *http.Response {
	origBody, _ := io.ReadAll(orig.Body)
	orig.Body.Close()
	restore := func() *http.Response {
		orig.Body = io.NopCloser(bytes.NewReader(origBody))
		return orig
	}

	c := t.client
	if c.Tokens == nil || c.Tokens.Refresh == "" {
		return restore()
	}
	if err := c.refreshTokens(req.Context(), staleToken); err != nil {
		return restore()
	}

	r2 := req.Clone(req.Context())
	if reqBody != nil {
		r2.Body = io.NopCloser(bytes.NewReader(reqBody))
	}
	// Force re-computation of Authorization with the refreshed token.
	r2.Header.Del("Authorization")
	newAccess := ""
	if c.Tokens != nil {
		newAccess = c.Tokens.Access
	}
	t.setHeaders(r2, newAccess)

	resp2, err := t.Base.RoundTrip(r2)
	if err != nil {
		return restore()
	}
	return resp2
}

// getClientIP returns the cached X-Client-IP value, safe to call
// concurrently with maybeFetchClientIP.
func (c *Client) getClientIP() string {
	c.clientIPMu.Lock()
	defer c.clientIPMu.Unlock()
	return c.ClientIP
}

// maybeFetchClientIP starts a background goroutine that calls
// GetUserTimestamp and caches the returned ip_address into c.ClientIP.
// If a fetch is already in flight, or a client IP is already set, it
// returns immediately. On fetch failure the field is left empty and the
// next invocation will try again.
func (c *Client) maybeFetchClientIP() {
	c.clientIPMu.Lock()
	if c.ClientIP != "" || c.clientIPFetching {
		c.clientIPMu.Unlock()
		return
	}
	c.clientIPFetching = true
	c.clientIPMu.Unlock()

	go func() {
		ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
		defer cancel()

		ip := ""
		if resp, _, err := c.GetUserTimestamp(ctx).Execute(); err == nil && resp != nil {
			ip = resp.GetIpAddress()
		}

		c.clientIPMu.Lock()
		c.clientIPFetching = false
		if ip != "" {
			c.ClientIP = ip
		}
		c.clientIPMu.Unlock()
	}()
}

func basicCredentials(apiKey string) string {
	return base64.StdEncoding.EncodeToString([]byte(apiKey))
}

// isOAuthTokenPath matches /api/v1/oauth/token and its sub-paths, which the
// Yay! server authenticates with Basic rather than Bearer and which must be
// excluded from the 401 auto-refresh retry to prevent infinite loops.
func isOAuthTokenPath(path string) bool {
	return strings.Contains(path, "/api/v1/oauth/token")
}

// isTimestampPath matches /v2/users/timestamp, the endpoint used to look up
// the caller's IP for X-Client-IP. The lazy fetch must not re-trigger
// itself when calling this endpoint.
func isTimestampPath(path string) bool {
	return strings.HasSuffix(path, "/v2/users/timestamp")
}
