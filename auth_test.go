package yaylib

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"sync/atomic"
	"testing"
)

// fakeLoginServer answers LoginWithEmail with the supplied tokens and
// returns the configured user_id. Other paths get a 404.
func fakeLoginServer(t *testing.T, accessToken, refreshToken string, userID int64) *httptest.Server {
	t.Helper()
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if !strings.HasSuffix(r.URL.Path, "/login_with_email") {
			http.NotFound(w, r)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(map[string]any{
			"access_token":  accessToken,
			"refresh_token": refreshToken,
			"user_id":       userID,
		})
	}))
}

// errorStore is a SessionStore whose Load and Save can be made to fail
// with configurable errors, for exercising the auth error-handling paths.
type errorStore struct {
	loadErr  error
	saveErr  error
	loadVal  *Session
	saveSeen *Session
}

func (s *errorStore) Load(string) (*Session, error) {
	if s.loadErr != nil {
		return nil, s.loadErr
	}
	if s.loadVal != nil {
		v := *s.loadVal
		return &v, nil
	}
	return nil, ErrNoSession
}

func (s *errorStore) Save(sess *Session) error {
	v := *sess
	s.saveSeen = &v
	return s.saveErr
}

func (s *errorStore) Delete(string) error { return nil }

// PORTING:S1
func TestLoginWithEmail_FreshLoginPersistsSession(t *testing.T) {
	srv := fakeLoginServer(t, "ACC", "REF", 4242)
	defer srv.Close()

	store := NewMemoryStore()
	c := NewClient(WithBaseURL(srv.URL), WithSessionStore(store))

	resp, _, err := c.LoginWithEmail(context.Background()).
		Email("foo@example.com").Password("pw").Execute()
	if err != nil {
		t.Fatalf("LoginWithEmail: %v", err)
	}
	if resp.GetUserId() != 4242 {
		t.Errorf("UserID = %d, want 4242", resp.GetUserId())
	}
	if got := c.Tokens().Access; got != "ACC" {
		t.Errorf("access token = %q, want ACC", got)
	}
	if c.UserID != 4242 {
		t.Errorf("Client.UserID = %d, want 4242", c.UserID)
	}
	if c.currentEmail != "foo@example.com" {
		t.Errorf("currentEmail = %q, want foo@example.com", c.currentEmail)
	}

	sess, err := store.Load("foo@example.com")
	if err != nil {
		t.Fatalf("session not persisted: %v", err)
	}
	if sess.AccessToken != "ACC" || sess.RefreshToken != "REF" || sess.UserID != 4242 {
		t.Errorf("persisted session = %+v", sess)
	}
}

// PORTING:S2
func TestLoginWithEmail_RestoresFromStoreWithoutHTTP(t *testing.T) {
	var hits int32
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		atomic.AddInt32(&hits, 1)
		http.Error(w, "should not be hit", http.StatusInternalServerError)
	}))
	defer srv.Close()

	store := NewMemoryStore()
	if err := store.Save(&Session{
		Email: "foo@example.com", UserID: 99,
		AccessToken: "AT", RefreshToken: "RT",
	}); err != nil {
		t.Fatal(err)
	}

	c := NewClient(WithBaseURL(srv.URL), WithSessionStore(store))
	resp, httpResp, err := c.LoginWithEmail(context.Background()).
		Email("foo@example.com").Password("pw").Execute()
	if err != nil {
		t.Fatalf("LoginWithEmail: %v", err)
	}
	if httpResp != nil {
		t.Errorf("HTTP response should be nil on cache hit, got %v", httpResp.Status)
	}
	if resp.GetUserId() != 99 || resp.GetAccessToken() != "AT" {
		t.Errorf("synthesized resp = %+v", resp)
	}
	if c.Tokens().Access != "AT" {
		t.Errorf("client tokens not applied from store, got %q", c.Tokens().Access)
	}
	if c.UserID != 99 {
		t.Errorf("Client.UserID = %d, want 99", c.UserID)
	}
	if got := atomic.LoadInt32(&hits); got != 0 {
		t.Errorf("server was hit %d times on cache hit, want 0", got)
	}
}

func TestLoginWithEmail_NoCacheBypassesStore(t *testing.T) {
	srv := fakeLoginServer(t, "FRESH", "FRESH_REF", 7)
	defer srv.Close()

	store := NewMemoryStore()
	if err := store.Save(&Session{
		Email: "foo@example.com", UserID: 1,
		AccessToken: "OLD", RefreshToken: "OLD_REF",
	}); err != nil {
		t.Fatal(err)
	}

	c := NewClient(WithBaseURL(srv.URL), WithSessionStore(store))
	resp, _, err := c.LoginWithEmail(context.Background()).
		Email("foo@example.com").Password("pw").NoCache().Execute()
	if err != nil {
		t.Fatalf("LoginWithEmail: %v", err)
	}
	if resp.GetAccessToken() != "FRESH" {
		t.Errorf("access = %q, want FRESH", resp.GetAccessToken())
	}
	sess, _ := store.Load("foo@example.com")
	if sess.AccessToken != "FRESH" {
		t.Errorf("store should be overwritten with fresh tokens, got %q", sess.AccessToken)
	}
}

// PORTING:S3
func TestLoginWithEmail_LoadErrorReturnsErrorWithoutHTTP(t *testing.T) {
	var hits int32
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		atomic.AddInt32(&hits, 1)
		w.WriteHeader(http.StatusInternalServerError)
	}))
	defer srv.Close()

	store := &errorStore{loadErr: errors.New("corrupt JSON")}
	c := NewClient(WithBaseURL(srv.URL), WithSessionStore(store))
	_, _, err := c.LoginWithEmail(context.Background()).
		Email("foo@example.com").Password("pw").Execute()
	if err == nil {
		t.Fatal("LoginWithEmail: want error from broken store, got nil")
	}
	if !strings.Contains(err.Error(), "load session") {
		t.Errorf("error = %q, want one mentioning 'load session'", err)
	}
	if got := atomic.LoadInt32(&hits); got != 0 {
		t.Errorf("server was hit %d times on store error; rate-limit protection failed", got)
	}
}

// PORTING:S4
func TestLoginWithEmail_SaveFailureWrapsErrSessionSaveFailed(t *testing.T) {
	srv := fakeLoginServer(t, "AT", "RT", 42)
	defer srv.Close()

	store := &errorStore{saveErr: errors.New("disk full")}
	c := NewClient(WithBaseURL(srv.URL), WithSessionStore(store))
	resp, _, err := c.LoginWithEmail(context.Background()).
		Email("foo@example.com").Password("pw").Execute()
	if !errors.Is(err, ErrSessionSaveFailed) {
		t.Fatalf("err = %v, want errors.Is ErrSessionSaveFailed", err)
	}
	if resp == nil || resp.GetAccessToken() != "AT" {
		t.Errorf("resp should still be populated on persist failure, got %+v", resp)
	}
	if c.Tokens().Access != "AT" {
		t.Errorf("tokens should be active even when persist failed, got %q", c.Tokens().Access)
	}
}

// PORTING:S5
func TestLoginWithEmail_TwoFARequiredSurfacesErrorCode(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if !strings.HasSuffix(r.URL.Path, "/login_with_email") {
			http.NotFound(w, r)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusUnauthorized)
		_ = json.NewEncoder(w).Encode(map[string]any{
			"error_code": int(ErrCodeRequired2FA),
			"message":    "2FA required",
		})
	}))
	defer srv.Close()

	c := NewClient(WithBaseURL(srv.URL), WithRetryPolicy(RetryPolicy{}))
	_, _, err := c.LoginWithEmail(context.Background()).
		Email("u@example.com").Password("pw").Execute()
	if err == nil {
		t.Fatal("expected error from 2FA-required server, got nil")
	}
	if got := CodeOf(err); got != ErrCodeRequired2FA {
		t.Errorf("CodeOf(err) = %v, want ErrCodeRequired2FA", got)
	}
	if c.Tokens().Access != "" {
		t.Errorf("tokens should not be set on 2FA-required, got %q", c.Tokens().Access)
	}
}

// webFirstLoginServer models the web-first flow. The browser login (v2) is
// primary: it succeeds when webOK is true and otherwise rejects with 401. The
// app login (v3) is the fallback: it succeeds when appOK is true, else 401.
// Hits on each endpoint and the api_key each request carried are recorded so
// tests can assert which path ran and under which key.
func webFirstLoginServer(t *testing.T, webOK, appOK bool, seenWebKey, seenAppKey *string, webHit, appHit *int32) *httptest.Server {
	t.Helper()
	reject := func(w http.ResponseWriter) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusUnauthorized)
		_ = json.NewEncoder(w).Encode(map[string]any{
			"result": "error", "message": "unauthorized", "error_code": 401,
		})
	}
	record := func(r *http.Request, dst *string) {
		if dst == nil {
			return
		}
		var body map[string]any
		_ = json.NewDecoder(r.Body).Decode(&body)
		k, _ := body["api_key"].(string)
		*dst = k
	}
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch {
		case strings.HasPrefix(r.URL.Path, "/v2/users/login_with_email"):
			if webHit != nil {
				atomic.AddInt32(webHit, 1)
			}
			record(r, seenWebKey)
			if !webOK {
				reject(w)
				return
			}
			w.Header().Set("Content-Type", "application/json")
			_ = json.NewEncoder(w).Encode(map[string]any{
				"access_token": "WEBACC", "refresh_token": "WEBREF", "user_id": 7777,
			})
		case strings.HasPrefix(r.URL.Path, "/v3/users/login_with_email"):
			if appHit != nil {
				atomic.AddInt32(appHit, 1)
			}
			record(r, seenAppKey)
			if !appOK {
				reject(w)
				return
			}
			w.Header().Set("Content-Type", "application/json")
			_ = json.NewEncoder(w).Encode(map[string]any{
				"access_token": "APPACC", "refresh_token": "APPREF", "user_id": 4242,
			})
		default:
			http.NotFound(w, r)
		}
	}))
}

// PORTING §6.2 — web-first: the browser login endpoint is the primary path.
// It succeeds directly and the app (v3) endpoint is never touched.
func TestLoginWithEmail_WebPrimarySucceeds(t *testing.T) {
	var seenKey string
	var appHit int32
	srv := webFirstLoginServer(t, true, true, &seenKey, nil, nil, &appHit)
	defer srv.Close()

	c := NewClient(WithBaseURL(srv.URL), WithRetryPolicy(RetryPolicy{}))
	resp, _, err := c.LoginWithEmail(context.Background()).
		Email("u@example.com").Password("pw").Execute()
	if err != nil {
		t.Fatalf("LoginWithEmail: %v", err)
	}
	if resp.GetUserId() != 7777 {
		t.Errorf("UserID = %d, want 7777 (web response)", resp.GetUserId())
	}
	if c.Tokens().Access != "WEBACC" || c.Tokens().Refresh != "WEBREF" {
		t.Errorf("tokens = %+v, want WEBACC/WEBREF from web login", c.Tokens())
	}
	if seenKey != DefaultWebAPIKey {
		t.Errorf("web api_key = %q, want DefaultWebAPIKey", seenKey)
	}
	if c.oauthKeySnapshot() != DefaultWebAPIKey {
		t.Errorf("oauthKeySnapshot = %q, want web key", c.oauthKeySnapshot())
	}
	if n := atomic.LoadInt32(&appHit); n != 0 {
		t.Errorf("app (v3) endpoint hit %d times, want 0 (web is primary)", n)
	}
}

// When the primary web login is rejected with a plain 401, the login falls
// back to the app (v3) endpoint under the app key, and the session tracks the
// app key for refresh.
func TestLoginWithEmail_FallsBackToAppOn401(t *testing.T) {
	var seenAppKey string
	var webHit, appHit int32
	srv := webFirstLoginServer(t, false, true, nil, &seenAppKey, &webHit, &appHit)
	defer srv.Close()

	c := NewClient(WithBaseURL(srv.URL), WithRetryPolicy(RetryPolicy{}))
	resp, _, err := c.LoginWithEmail(context.Background()).
		Email("u@example.com").Password("pw").Execute()
	if err != nil {
		t.Fatalf("LoginWithEmail with app fallback: %v", err)
	}
	if resp.GetUserId() != 4242 {
		t.Errorf("UserID = %d, want 4242 (app response)", resp.GetUserId())
	}
	if c.Tokens().Access != "APPACC" {
		t.Errorf("tokens = %+v, want APPACC from app fallback", c.Tokens())
	}
	if seenAppKey != DefaultAPIKey {
		t.Errorf("app api_key = %q, want DefaultAPIKey", seenAppKey)
	}
	if c.oauthKeySnapshot() != DefaultAPIKey {
		t.Errorf("oauthKeySnapshot = %q, want app key after app fallback", c.oauthKeySnapshot())
	}
	if atomic.LoadInt32(&webHit) != 1 || atomic.LoadInt32(&appHit) != 1 {
		t.Errorf("hits web=%d app=%d, want 1/1 (web tried then app)", webHit, appHit)
	}
}

// With the web path switched off (WebAPIKey==""), login goes straight to the
// app endpoint and no v2 request is attempted; an app 401 surfaces unchanged.
func TestLoginWithEmail_WebDisabledUsesAppOnly(t *testing.T) {
	var webHit int32
	srv := webFirstLoginServer(t, true, false, nil, nil, &webHit, nil)
	defer srv.Close()

	c := NewClient(WithBaseURL(srv.URL), WithRetryPolicy(RetryPolicy{}), WithWebAPIKey(""))
	_, _, err := c.LoginWithEmail(context.Background()).
		Email("u@example.com").Password("pw").Execute()
	if err == nil {
		t.Fatal("expected app 401 to surface when web is disabled")
	}
	if CodeOf(err) != ErrCodeUnauthorized {
		t.Errorf("CodeOf(err) = %v, want ErrCodeUnauthorized", CodeOf(err))
	}
	if n := atomic.LoadInt32(&webHit); n != 0 {
		t.Errorf("web (v2) endpoint hit %d times, want 0 (web disabled)", n)
	}
	if c.Tokens().Access != "" {
		t.Errorf("tokens should stay empty, got %q", c.Tokens().Access)
	}
}

// When both the primary web login and the app fallback fail, the primary
// (web) 401 — a proper *APIError — is surfaced rather than the fallback's.
func TestLoginWithEmail_BothFailSurfacesWeb401(t *testing.T) {
	srv := webFirstLoginServer(t, false, false, nil, nil, nil, nil)
	defer srv.Close()

	c := NewClient(WithBaseURL(srv.URL), WithRetryPolicy(RetryPolicy{}))
	_, _, err := c.LoginWithEmail(context.Background()).
		Email("u@example.com").Password("pw").Execute()
	if err == nil {
		t.Fatal("expected error when both logins fail")
	}
	if CodeOf(err) != ErrCodeUnauthorized {
		t.Errorf("CodeOf(err) = %v, want ErrCodeUnauthorized", CodeOf(err))
	}
	if c.Tokens().Access != "" {
		t.Errorf("tokens should stay empty, got %q", c.Tokens().Access)
	}
}

// A 2FA-required rejection is also HTTP 401 but carries ErrCodeRequired2FA;
// the app fallback must NOT fire (the caller has to supply a code, not switch
// endpoints), so the primary web login's 2FA surfaces and the app (v3)
// endpoint is never called.
func TestLoginWithEmail_NoFallbackOn2FA(t *testing.T) {
	var v3Hit int32
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		switch {
		case strings.HasPrefix(r.URL.Path, "/v2/users/login_with_email"):
			w.WriteHeader(http.StatusUnauthorized)
			_ = json.NewEncoder(w).Encode(map[string]any{
				"error_code": int(ErrCodeRequired2FA), "message": "2FA required",
			})
		case strings.HasPrefix(r.URL.Path, "/v3/users/login_with_email"):
			atomic.AddInt32(&v3Hit, 1)
			_ = json.NewEncoder(w).Encode(map[string]any{
				"access_token": "APPACC", "refresh_token": "APPREF", "user_id": 7777,
			})
		default:
			http.NotFound(w, r)
		}
	}))
	defer srv.Close()

	c := NewClient(WithBaseURL(srv.URL), WithRetryPolicy(RetryPolicy{}))
	_, _, err := c.LoginWithEmail(context.Background()).
		Email("u@example.com").Password("pw").Execute()
	if err == nil {
		t.Fatal("expected 2FA-required error")
	}
	if CodeOf(err) != ErrCodeRequired2FA {
		t.Errorf("CodeOf(err) = %v, want ErrCodeRequired2FA", CodeOf(err))
	}
	if n := atomic.LoadInt32(&v3Hit); n != 0 {
		t.Errorf("app (v3) endpoint hit %d times, want 0 (no fallback on 2FA)", n)
	}
	if c.Tokens().Access != "" {
		t.Errorf("tokens should stay empty, got %q", c.Tokens().Access)
	}
}

// When the caller pins a specific api_key via .ApiKey(), the web-first path
// is skipped (it would otherwise authenticate under WebAPIKey), so login goes
// straight to the app endpoint under the pinned key and the browser (v2)
// endpoint is never called.
func TestLoginWithEmail_NoFallbackWhenApiKeyPinned(t *testing.T) {
	var seenAppKey string
	var webHit int32
	srv := webFirstLoginServer(t, true, false, nil, &seenAppKey, &webHit, nil)
	defer srv.Close()

	c := NewClient(WithBaseURL(srv.URL), WithRetryPolicy(RetryPolicy{}))
	_, _, err := c.LoginWithEmail(context.Background()).
		Email("u@example.com").Password("pw").ApiKey("caller-pinned-key").Execute()
	if err == nil {
		t.Fatal("expected app 401 when caller pinned a custom api_key")
	}
	if CodeOf(err) != ErrCodeUnauthorized {
		t.Errorf("CodeOf(err) = %v, want ErrCodeUnauthorized", CodeOf(err))
	}
	if n := atomic.LoadInt32(&webHit); n != 0 {
		t.Errorf("web (v2) endpoint should not have been called, hit %d times", n)
	}
	if seenAppKey != "caller-pinned-key" {
		t.Errorf("app api_key = %q, want the caller-pinned key", seenAppKey)
	}
}

// A primary web response that is 2xx but undecodable is a genuine login
// result, so its decode error (carrying the raw body on the response) is
// surfaced as a typed APIError; since it is not ErrCodeUnauthorized the app
// fallback does not fire.
func TestLoginWithEmail_WebDecodeFailureSurfaced(t *testing.T) {
	var appHit int32
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch {
		case strings.HasPrefix(r.URL.Path, "/v2/users/login_with_email"):
			w.Header().Set("Content-Type", "application/json")
			_, _ = w.Write([]byte("{ this is not valid json"))
		case strings.HasPrefix(r.URL.Path, "/v3/users/login_with_email"):
			atomic.AddInt32(&appHit, 1)
			w.Header().Set("Content-Type", "application/json")
			_ = json.NewEncoder(w).Encode(map[string]any{
				"access_token": "APPACC", "refresh_token": "APPREF", "user_id": 1,
			})
		default:
			http.NotFound(w, r)
		}
	}))
	defer srv.Close()

	c := NewClient(WithBaseURL(srv.URL), WithRetryPolicy(RetryPolicy{}))
	_, httpResp, err := c.LoginWithEmail(context.Background()).
		Email("u@example.com").Password("pw").Execute()
	if err == nil {
		t.Fatal("expected a decode error to surface")
	}
	// Not masked as an auth rejection, and no app fallback.
	if CodeOf(err) == ErrCodeUnauthorized {
		t.Errorf("decode failure was masked as ErrCodeUnauthorized: %v", err)
	}
	if !strings.Contains(err.Error(), "decode") {
		t.Errorf("err = %v, want a decode error", err)
	}
	if n := atomic.LoadInt32(&appHit); n != 0 {
		t.Errorf("app (v3) endpoint hit %d times, want 0 (2xx decode fail is not 401)", n)
	}
	// The raw web body is still readable on the returned response and on the
	// typed error.
	if httpResp == nil || httpResp.StatusCode != http.StatusOK {
		t.Fatalf("want the 2xx web response surfaced, got %v", httpResp)
	}
	var apiErr *APIError
	if !errors.As(err, &apiErr) || !strings.Contains(string(apiErr.Body()), "not valid json") {
		t.Errorf("want a typed APIError carrying the raw web body, got %T %v", err, err)
	}
	if c.Tokens().Access != "" {
		t.Errorf("tokens should stay empty, got %q", c.Tokens().Access)
	}
}

// PORTING §6.2 — a session obtained via the web fallback must refresh its
// tokens with the web key (Basic), not the app key, or the refresh 401s.
func TestRefresh_WebSessionUsesWebKeyForBasic(t *testing.T) {
	var oauthAuth string
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		switch {
		case strings.HasPrefix(r.URL.Path, "/v3/users/login_with_email"):
			w.WriteHeader(http.StatusUnauthorized)
			_ = json.NewEncoder(w).Encode(map[string]any{"result": "error", "error_code": 401})
		case strings.HasPrefix(r.URL.Path, "/v2/users/login_with_email"):
			_ = json.NewEncoder(w).Encode(map[string]any{"access_token": "A1", "refresh_token": "R1", "user_id": 1})
		case strings.Contains(r.URL.Path, "/api/v1/oauth/token"):
			oauthAuth = r.Header.Get("Authorization")
			_ = json.NewEncoder(w).Encode(map[string]any{"access_token": "A2", "refresh_token": "R2", "id": 1})
		default:
			http.NotFound(w, r)
		}
	}))
	defer srv.Close()

	c := NewClient(WithBaseURL(srv.URL), WithRetryPolicy(RetryPolicy{}))
	if _, _, err := c.LoginWithEmail(context.Background()).
		Email("u@example.com").Password("pw").Execute(); err != nil {
		t.Fatalf("login: %v", err)
	}
	if got := c.oauthKeySnapshot(); got != DefaultWebAPIKey {
		t.Fatalf("oauthKeySnapshot = %q, want web key", got)
	}
	if err := c.refreshTokens(context.Background(), c.Tokens().Access); err != nil {
		t.Fatalf("refresh: %v", err)
	}
	if want := "Basic " + basicCredentials(DefaultWebAPIKey); oauthAuth != want {
		t.Errorf("oauth Basic = %q, want web-key Basic", oauthAuth)
	}
	if c.Tokens().Access != "A2" {
		t.Errorf("access after refresh = %q, want A2", c.Tokens().Access)
	}
}

// An app login (reached here by disabling the web-first path with
// WithWebAPIKey("")) refreshes with the app key, and does not persist an
// oauth_key (app sessions stay byte-identical on disk).
func TestRefresh_AppSessionUsesAppKey(t *testing.T) {
	var oauthAuth string
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		switch {
		case strings.HasSuffix(r.URL.Path, "/login_with_email"):
			_ = json.NewEncoder(w).Encode(map[string]any{"access_token": "A1", "refresh_token": "R1", "user_id": 1})
		case strings.Contains(r.URL.Path, "/api/v1/oauth/token"):
			oauthAuth = r.Header.Get("Authorization")
			_ = json.NewEncoder(w).Encode(map[string]any{"access_token": "A2", "refresh_token": "R2", "id": 1})
		default:
			http.NotFound(w, r)
		}
	}))
	defer srv.Close()

	store := NewMemoryStore()
	c := NewClient(WithBaseURL(srv.URL), WithSessionStore(store), WithRetryPolicy(RetryPolicy{}), WithWebAPIKey(""))
	if _, _, err := c.LoginWithEmail(context.Background()).
		Email("u@example.com").Password("pw").Execute(); err != nil {
		t.Fatalf("login: %v", err)
	}
	if got := c.oauthKeySnapshot(); got != DefaultAPIKey {
		t.Errorf("oauthKeySnapshot = %q, want app key", got)
	}
	sess, _ := store.Load("u@example.com")
	if sess.OAuthKey != "" {
		t.Errorf("app session persisted oauth_key=%q, want empty", sess.OAuthKey)
	}
	if err := c.refreshTokens(context.Background(), c.Tokens().Access); err != nil {
		t.Fatalf("refresh: %v", err)
	}
	if want := "Basic " + basicCredentials(DefaultAPIKey); oauthAuth != want {
		t.Errorf("oauth Basic = %q, want app-key Basic", oauthAuth)
	}
}

// A cached web session, restored on a fresh client, still refreshes with the
// web key — the oauth_key round-trips through the session store.
func TestRefresh_RestoredWebSessionUsesWebKey(t *testing.T) {
	var oauthAuth string
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		switch {
		case strings.HasPrefix(r.URL.Path, "/v3/users/login_with_email"):
			w.WriteHeader(http.StatusUnauthorized)
			_ = json.NewEncoder(w).Encode(map[string]any{"result": "error", "error_code": 401})
		case strings.HasPrefix(r.URL.Path, "/v2/users/login_with_email"):
			_ = json.NewEncoder(w).Encode(map[string]any{"access_token": "A1", "refresh_token": "R1", "user_id": 1})
		case strings.Contains(r.URL.Path, "/api/v1/oauth/token"):
			oauthAuth = r.Header.Get("Authorization")
			_ = json.NewEncoder(w).Encode(map[string]any{"access_token": "A2", "refresh_token": "R2", "id": 1})
		default:
			http.NotFound(w, r)
		}
	}))
	defer srv.Close()

	store := NewMemoryStore()
	c1 := NewClient(WithBaseURL(srv.URL), WithSessionStore(store), WithRetryPolicy(RetryPolicy{}))
	if _, _, err := c1.LoginWithEmail(context.Background()).
		Email("u@example.com").Password("pw").Execute(); err != nil {
		t.Fatalf("login: %v", err)
	}
	if sess, _ := store.Load("u@example.com"); sess.OAuthKey != DefaultWebAPIKey {
		t.Fatalf("persisted oauth_key = %q, want web key", sess.OAuthKey)
	}

	// Fresh client restores the web session from the store (cache hit).
	c2 := NewClient(WithBaseURL(srv.URL), WithSessionStore(store), WithRetryPolicy(RetryPolicy{}))
	if _, _, err := c2.LoginWithEmail(context.Background()).
		Email("u@example.com").Password("pw").Execute(); err != nil {
		t.Fatalf("restore login: %v", err)
	}
	if got := c2.oauthKeySnapshot(); got != DefaultWebAPIKey {
		t.Errorf("restored oauthKeySnapshot = %q, want web key", got)
	}
	if err := c2.refreshTokens(context.Background(), c2.Tokens().Access); err != nil {
		t.Fatalf("refresh: %v", err)
	}
	if want := "Basic " + basicCredentials(DefaultWebAPIKey); oauthAuth != want {
		t.Errorf("restored-session oauth Basic = %q, want web-key Basic", oauthAuth)
	}
}

// PORTING §6.2 — a web session restored directly with LoadSession (not through
// the cached-login path) must also refresh with the web key. This is the
// regression the oauth_key restore fixes: LoadSession used to apply tokens,
// device UUID, and identity but not the establishing key, so a web session
// loaded straight from disk refreshed with the app key and 401'd.
func TestLoadSession_RestoresWebOAuthKey(t *testing.T) {
	var oauthAuth string
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if strings.Contains(r.URL.Path, "/api/v1/oauth/token") {
			oauthAuth = r.Header.Get("Authorization")
			w.Header().Set("Content-Type", "application/json")
			_ = json.NewEncoder(w).Encode(map[string]any{"access_token": "A2", "refresh_token": "R2", "id": 1})
			return
		}
		http.NotFound(w, r)
	}))
	defer srv.Close()

	// A web-fallback session persisted to the store (OAuthKey == web key).
	store := NewMemoryStore()
	if err := store.Save(&Session{
		Email: "u@example.com", UserID: 1,
		AccessToken: "A1", RefreshToken: "R1",
		OAuthKey: DefaultWebAPIKey,
	}); err != nil {
		t.Fatal(err)
	}

	c := NewClient(WithBaseURL(srv.URL), WithSessionStore(store), WithRetryPolicy(RetryPolicy{}))
	if _, err := c.LoadSession("u@example.com"); err != nil {
		t.Fatalf("LoadSession: %v", err)
	}
	if got := c.oauthKeySnapshot(); got != DefaultWebAPIKey {
		t.Fatalf("after LoadSession oauthKeySnapshot = %q, want web key", got)
	}
	if err := c.refreshTokens(context.Background(), c.Tokens().Access); err != nil {
		t.Fatalf("refresh: %v", err)
	}
	if want := "Basic " + basicCredentials(DefaultWebAPIKey); oauthAuth != want {
		t.Errorf("refresh Basic = %q, want web-key Basic after LoadSession", oauthAuth)
	}
	if c.Tokens().Access != "A2" {
		t.Errorf("access after refresh = %q, want A2", c.Tokens().Access)
	}
}

// App sessions carry an empty OAuthKey; LoadSession must leave refresh on the
// app key (the empty → app-key default), unchanged from before.
func TestLoadSession_AppSessionKeepsAppKey(t *testing.T) {
	store := NewMemoryStore()
	if err := store.Save(&Session{
		Email: "u@example.com", UserID: 1,
		AccessToken: "A1", RefreshToken: "R1",
		// no OAuthKey — an app session
	}); err != nil {
		t.Fatal(err)
	}

	c := NewClient(WithSessionStore(store))
	if _, err := c.LoadSession("u@example.com"); err != nil {
		t.Fatalf("LoadSession: %v", err)
	}
	if got := c.oauthKeySnapshot(); got != DefaultAPIKey {
		t.Errorf("after LoadSession oauthKeySnapshot = %q, want app key", got)
	}
}

// A web session injected directly via SaveSession (not through login) also
// activates its establishing key, so its refresh uses the web key. This
// mirrors LoadSession — SaveSession is a full activation, not just a write.
func TestSaveSession_ActivatesWebOAuthKey(t *testing.T) {
	var oauthAuth string
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if strings.Contains(r.URL.Path, "/api/v1/oauth/token") {
			oauthAuth = r.Header.Get("Authorization")
			w.Header().Set("Content-Type", "application/json")
			_ = json.NewEncoder(w).Encode(map[string]any{"access_token": "A2", "refresh_token": "R2", "id": 1})
			return
		}
		http.NotFound(w, r)
	}))
	defer srv.Close()

	c := NewClient(WithBaseURL(srv.URL), WithSessionStore(NewMemoryStore()), WithRetryPolicy(RetryPolicy{}))
	if err := c.SaveSession(&Session{
		Email: "u@example.com", UserID: 1,
		AccessToken: "A1", RefreshToken: "R1",
		OAuthKey: DefaultWebAPIKey,
	}); err != nil {
		t.Fatalf("SaveSession: %v", err)
	}
	if got := c.oauthKeySnapshot(); got != DefaultWebAPIKey {
		t.Fatalf("after SaveSession oauthKeySnapshot = %q, want web key", got)
	}
	if err := c.refreshTokens(context.Background(), c.Tokens().Access); err != nil {
		t.Fatalf("refresh: %v", err)
	}
	if want := "Basic " + basicCredentials(DefaultWebAPIKey); oauthAuth != want {
		t.Errorf("refresh Basic = %q, want web-key Basic after SaveSession", oauthAuth)
	}
}

// A primary web login rejected with a NON-Unauthorized typed error (here a
// temporary ban carrying ban_until) surfaces directly as a typed *APIError —
// no app fallback fires (it is not ErrCodeUnauthorized), and ErrorResponseOf
// / errors.As read the parsed fields and raw body. This locks the
// loginWithEmailWeb typed-error upgrade now that web is the primary path.
func TestLoginWithEmail_WebPrimaryTypedErrorSurfaced(t *testing.T) {
	var appHit int32
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch {
		case strings.HasPrefix(r.URL.Path, "/v2/users/login_with_email"):
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusForbidden)
			_ = json.NewEncoder(w).Encode(map[string]any{
				"result": "error", "message": "temporarily banned",
				"error_code": int(ErrCodeUserTemporaryBanned), "ban_until": 1893456000,
			})
		case strings.HasPrefix(r.URL.Path, "/v3/users/login_with_email"):
			atomic.AddInt32(&appHit, 1)
			w.Header().Set("Content-Type", "application/json")
			_ = json.NewEncoder(w).Encode(map[string]any{"access_token": "APPACC", "user_id": 1})
		default:
			http.NotFound(w, r)
		}
	}))
	defer srv.Close()

	c := NewClient(WithBaseURL(srv.URL), WithRetryPolicy(RetryPolicy{}))
	_, _, err := c.LoginWithEmail(context.Background()).
		Email("u@example.com").Password("pw").Execute()
	if err == nil {
		t.Fatal("expected the web ban error to surface")
	}
	if CodeOf(err) != ErrCodeUserTemporaryBanned {
		t.Errorf("CodeOf(err) = %v, want ErrCodeUserTemporaryBanned", CodeOf(err))
	}
	if n := atomic.LoadInt32(&appHit); n != 0 {
		t.Errorf("app (v3) endpoint hit %d times, want 0 (non-401 does not fall back)", n)
	}
	if r := ErrorResponseOf(err); r == nil || r.BanUntil == nil || *r.BanUntil != 1893456000 {
		t.Errorf("ErrorResponseOf did not parse ban_until from the web error: %+v", r)
	}
	var apiErr *APIError
	if !errors.As(err, &apiErr) || !strings.Contains(string(apiErr.Body()), "temporarily banned") {
		t.Errorf("want a typed APIError carrying the raw web body, got %T %v", err, err)
	}
}
