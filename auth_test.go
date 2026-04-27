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
	if got := c.accessSnapshot(); got != "ACC" {
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
	if c.accessSnapshot() != "AT" {
		t.Errorf("client tokens not applied from store, got %q", c.accessSnapshot())
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
	if c.accessSnapshot() != "AT" {
		t.Errorf("tokens should be active even when persist failed, got %q", c.accessSnapshot())
	}
}
