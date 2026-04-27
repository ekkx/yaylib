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

func TestTransport_401TriggersRefreshAndRetries(t *testing.T) {
	var dataAttempt int32
	cap := &requestCapturer{handler: func(w http.ResponseWriter, r *http.Request) {
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
			_, _ = w.Write([]byte(`{}`))
			return
		}
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{}`))
	}}
	srv := httptest.NewServer(cap)
	defer srv.Close()

	c := NewClient(WithBaseURL(srv.URL), WithRetryPolicy(RetryPolicy{}))
	c.SetTokens("STALE", "REF")

	req, _ := http.NewRequestWithContext(context.Background(), http.MethodGet, srv.URL+"/v1/anywhere", nil)
	resp, err := c.httpClient.Do(req)
	if err != nil {
		t.Fatalf("Do: %v", err)
	}
	resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		t.Errorf("final status = %d, want 200", resp.StatusCode)
	}
	if c.accessSnapshot() != "NEW_ACC" {
		t.Errorf("access token after refresh = %q, want NEW_ACC", c.accessSnapshot())
	}

	requests := cap.snapshot()
	var dataAuths []string
	var refreshHits int
	for _, req := range requests {
		if strings.Contains(req.path, "/api/v1/oauth/token") {
			refreshHits++
			continue
		}
		if strings.Contains(req.path, "/v1/anywhere") {
			dataAuths = append(dataAuths, req.headers.Get("Authorization"))
		}
	}
	if refreshHits != 1 {
		t.Errorf("refresh hits = %d, want 1", refreshHits)
	}
	if len(dataAuths) != 2 {
		t.Fatalf("data path attempts = %d, want 2 (initial 401 + retry)", len(dataAuths))
	}
	if dataAuths[0] != "Bearer STALE" {
		t.Errorf("first attempt Authorization = %q, want Bearer STALE", dataAuths[0])
	}
	if dataAuths[1] != "Bearer NEW_ACC" {
		t.Errorf("retry Authorization = %q, want Bearer NEW_ACC", dataAuths[1])
	}
}

func TestTransport_401RefreshFailureSurfaces401(t *testing.T) {
	cap := &requestCapturer{handler: func(w http.ResponseWriter, r *http.Request) {
		if strings.Contains(r.URL.Path, "/api/v1/oauth/token") {
			http.Error(w, `{"error_code":-1,"message":"refresh denied"}`, http.StatusUnauthorized)
			return
		}
		w.WriteHeader(http.StatusUnauthorized)
		_, _ = w.Write([]byte(`{"error_code":-2,"message":"original 401"}`))
	}}
	srv := httptest.NewServer(cap)
	defer srv.Close()

	c := NewClient(WithBaseURL(srv.URL), WithRetryPolicy(RetryPolicy{}))
	c.SetTokens("STALE", "REF")

	req, _ := http.NewRequestWithContext(context.Background(), http.MethodGet, srv.URL+"/v1/anywhere", nil)
	resp, err := c.httpClient.Do(req)
	if err != nil {
		t.Fatalf("Do: %v", err)
	}
	body, _ := io.ReadAll(resp.Body)
	resp.Body.Close()

	if resp.StatusCode != http.StatusUnauthorized {
		t.Errorf("status = %d, want 401", resp.StatusCode)
	}
	if !strings.Contains(string(body), "original 401") {
		t.Errorf("body = %q, want original 401 body to be preserved", body)
	}
}

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
