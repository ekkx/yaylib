package yaylib

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"sync"
	"sync/atomic"
	"testing"
	"time"
)

// requestCapturer collects the headers and paths of every incoming
// request, while letting each test plug in a custom handler for the
// response side.
type requestCapturer struct {
	mu       sync.Mutex
	requests []capturedRequest
	handler  func(w http.ResponseWriter, r *http.Request)
}

type capturedRequest struct {
	method  string
	path    string
	headers http.Header
	body    []byte
}

func (c *requestCapturer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	body, _ := io.ReadAll(r.Body)
	r.Body.Close()
	c.mu.Lock()
	c.requests = append(c.requests, capturedRequest{
		method: r.Method, path: r.URL.Path, headers: r.Header.Clone(), body: body,
	})
	c.mu.Unlock()
	c.handler(w, r)
}

func (c *requestCapturer) snapshot() []capturedRequest {
	c.mu.Lock()
	defer c.mu.Unlock()
	out := make([]capturedRequest, len(c.requests))
	copy(out, c.requests)
	return out
}

// PORTING:S6
func TestTransport_InjectsRequiredHeaders(t *testing.T) {
	cap := &requestCapturer{handler: func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{}`))
	}}
	srv := httptest.NewServer(cap)
	defer srv.Close()

	c := NewClient(WithBaseURL(srv.URL))
	c.SetTokens("ACC", "REF")

	req, _ := http.NewRequestWithContext(context.Background(), http.MethodGet, srv.URL+"/v1/anywhere", nil)
	resp, err := c.httpClient.Do(req)
	if err != nil {
		t.Fatalf("Do: %v", err)
	}
	resp.Body.Close()

	got := cap.snapshot()[0].headers
	for _, h := range []string{
		"User-Agent", "X-App-Version", "X-Device-Info", "X-Device-UUID",
		"X-Connection-Type", "X-Connection-Speed", "Accept-Language",
		"X-Timestamp",
	} {
		if got.Get(h) == "" {
			t.Errorf("missing required header %s", h)
		}
	}
	if got.Get("Authorization") != "Bearer ACC" {
		t.Errorf("Authorization = %q, want %q", got.Get("Authorization"), "Bearer ACC")
	}
}

// PORTING:S6
func TestTransport_OAuthEndpointUsesBasicAuth(t *testing.T) {
	cap := &requestCapturer{handler: func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{}`))
	}}
	srv := httptest.NewServer(cap)
	defer srv.Close()

	c := NewClient(WithBaseURL(srv.URL))
	c.SetTokens("ACC", "REF")

	req, _ := http.NewRequestWithContext(context.Background(), http.MethodPost, srv.URL+"/api/v1/oauth/token", strings.NewReader(""))
	resp, err := c.httpClient.Do(req)
	if err != nil {
		t.Fatalf("Do: %v", err)
	}
	resp.Body.Close()

	got := cap.snapshot()[0].headers.Get("Authorization")
	if !strings.HasPrefix(got, "Basic ") {
		t.Errorf("Authorization on /api/v1/oauth/token = %q, want Basic prefix", got)
	}
}

// expired-token: protected requests 401 (error_code -3) until the
// token endpoint is hit (a simulated refresh), then the happy path.
// This is the 401 → refresh → retry chain. Behavioral assertion:
// final success and the seeded tokens rotated to the server-issued
// ones (the Bearer-header progression is a Go-internal mechanic, and
// §15 says translate the scenario, not the mechanics).
//
// PORTING:S7
func TestTransport_401TriggersRefreshAndRetries(t *testing.T) {
	c := mockClient(t, "expired-token", WithRetryPolicy(RetryPolicy{}))
	c.SetTokens("STALE", "REF")

	resp, err := mockGet(t, c)
	if err != nil {
		t.Fatalf("Get: %v", err)
	}
	if resp.StatusCode != http.StatusOK {
		t.Errorf("final status = %d, want 200 (refresh + retry succeeded)", resp.StatusCode)
	}
	tok := c.Tokens()
	if tok.Access == "STALE" || tok.Access == "" {
		t.Errorf("access token = %q, want it rotated to the server-issued value", tok.Access)
	}
	if tok.Refresh == "REF" || tok.Refresh == "" {
		t.Errorf("refresh token = %q, want it rotated to the server-issued value", tok.Refresh)
	}
}

// fail-401-times-2: the data request 401s and the refresh call (same
// session+scenario counter) also 401s, so the refresh fails and the
// original 401 surfaces with its body intact (README composition note).
//
// PORTING:S8
func TestTransport_401RefreshFailureSurfaces401(t *testing.T) {
	c := mockClient(t, "fail-401-times-2", WithRetryPolicy(RetryPolicy{}))
	c.SetTokens("STALE", "REF")

	resp, err := mockGetRaw(t, c)
	if err != nil {
		t.Fatalf("Get: %v", err)
	}
	body, _ := io.ReadAll(resp.Body)
	resp.Body.Close()

	if resp.StatusCode != http.StatusUnauthorized {
		t.Errorf("status = %d, want 401 (refresh failed, original 401 surfaced)", resp.StatusCode)
	}
	if len(body) == 0 {
		t.Error("body empty; want the original 401 error body preserved")
	}
}

// PORTING:S9
func TestTransport_401RefreshSucceedsButRetryNetworkErrorSurfacesError(t *testing.T) {
	// Server: oauth/token returns fresh tokens. Data path returns 401
	// once, then closes the connection on the retry to simulate a
	// post-refresh network failure.
	var dataAttempt int32
	srv := httptest.NewUnstartedServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if strings.Contains(r.URL.Path, "/api/v1/oauth/token") {
			w.Header().Set("Content-Type", "application/json")
			_ = json.NewEncoder(w).Encode(map[string]any{
				"access_token":  "NEW_ACC",
				"refresh_token": "NEW_REF",
			})
			return
		}
		n := atomic.AddInt32(&dataAttempt, 1)
		if n == 1 {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		// Force a connection-level failure on the retry by hijacking
		// and abruptly closing the underlying TCP socket.
		hj, ok := w.(http.Hijacker)
		if !ok {
			t.Fatal("server doesn't support hijack")
		}
		conn, _, err := hj.Hijack()
		if err != nil {
			t.Fatalf("hijack: %v", err)
		}
		_ = conn.Close()
	}))
	srv.Start()
	defer srv.Close()

	c := NewClient(WithBaseURL(srv.URL), WithRetryPolicy(RetryPolicy{}))
	c.SetTokens("STALE", "REF")

	req, _ := http.NewRequestWithContext(context.Background(), http.MethodGet, srv.URL+"/v1/anywhere", nil)
	resp, err := c.httpClient.Do(req)
	if err == nil {
		// httpClient.Do may surface as resp + 0 bytes if the connection
		// was closed mid-response, but our handle401 path returns
		// (nil, err) so we expect err non-nil here.
		resp.Body.Close()
		t.Fatal("expected error from broken retry, got nil")
	}
	if c.Tokens().Access != "NEW_ACC" {
		t.Errorf("tokens should still be rotated, got access=%q", c.Tokens().Access)
	}
}

// PORTING:S10
func TestTransport_LazyFetchClientIPPopulatesHeader(t *testing.T) {
	const fetchedIP = "203.0.113.7"
	cap := &requestCapturer{handler: func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		if strings.HasSuffix(r.URL.Path, "/v2/users/timestamp") {
			_ = json.NewEncoder(w).Encode(map[string]any{
				"time":       int64(0),
				"ip_address": fetchedIP,
			})
			return
		}
		_, _ = w.Write([]byte(`{}`))
	}}
	srv := httptest.NewServer(cap)
	defer srv.Close()

	c := NewClient(WithBaseURL(srv.URL))

	// First request kicks off the lazy IP fetch in a background
	// goroutine; the request itself goes out without X-Client-IP.
	req1, _ := http.NewRequestWithContext(context.Background(), http.MethodGet, srv.URL+"/v1/first", nil)
	resp1, err := c.httpClient.Do(req1)
	if err != nil {
		t.Fatalf("first Do: %v", err)
	}
	resp1.Body.Close()

	deadline := time.Now().Add(2 * time.Second)
	for time.Now().Before(deadline) {
		if c.getClientIP() == fetchedIP {
			break
		}
		time.Sleep(20 * time.Millisecond)
	}
	if c.getClientIP() != fetchedIP {
		t.Fatalf("ClientIP not populated after lazy fetch, got %q", c.getClientIP())
	}

	req2, _ := http.NewRequestWithContext(context.Background(), http.MethodGet, srv.URL+"/v1/second", nil)
	resp2, err := c.httpClient.Do(req2)
	if err != nil {
		t.Fatalf("second Do: %v", err)
	}
	resp2.Body.Close()

	requests := cap.snapshot()
	var firstIP, secondIP string
	for _, req := range requests {
		switch req.path {
		case "/v1/first":
			firstIP = req.headers.Get("X-Client-IP")
		case "/v1/second":
			secondIP = req.headers.Get("X-Client-IP")
		}
	}
	if firstIP != "" {
		t.Errorf("first request X-Client-IP = %q, want empty (fetch is async)", firstIP)
	}
	if secondIP != fetchedIP {
		t.Errorf("second request X-Client-IP = %q, want %q", secondIP, fetchedIP)
	}
}
