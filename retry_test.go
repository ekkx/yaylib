package yaylib

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"sync/atomic"
	"testing"
	"time"
)

// newTestClient builds a Client pointed at testURL with retry behavior set
// to policy and X-Client-IP pre-populated so the lazy IP-fetch goroutine
// stays out of the test path.
func newTestClient(testURL string, policy RetryPolicy) *Client {
	return NewClient(
		WithBaseURL(testURL),
		WithClientIP("127.0.0.1"),
		WithRetryPolicy(policy),
	)
}

func TestTransport_RetriesOn5xx(t *testing.T) {
	var attempts int32
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		n := atomic.AddInt32(&attempts, 1)
		if n < 3 {
			w.WriteHeader(http.StatusServiceUnavailable)
			return
		}
		w.WriteHeader(http.StatusOK)
	}))
	defer server.Close()

	c := newTestClient(server.URL, RetryPolicy{
		MaxAttempts: 5,
		BaseDelay:   time.Millisecond,
		MaxDelay:    10 * time.Millisecond,
	})

	resp, err := c.httpClient.Get(server.URL + "/test")
	if err != nil {
		t.Fatalf("Get: %v", err)
	}
	defer resp.Body.Close()
	_, _ = io.Copy(io.Discard, resp.Body)

	if resp.StatusCode != http.StatusOK {
		t.Errorf("status = %d, want 200", resp.StatusCode)
	}
	if got := atomic.LoadInt32(&attempts); got != 3 {
		t.Errorf("attempts = %d, want 3", got)
	}
}

func TestTransport_RetriesPOSTOn429(t *testing.T) {
	var attempts int32
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		n := atomic.AddInt32(&attempts, 1)
		if n < 2 {
			w.WriteHeader(http.StatusTooManyRequests)
			return
		}
		w.WriteHeader(http.StatusOK)
	}))
	defer server.Close()

	// RetryOnPOST=false (default): POST 5xx is not retried, but a 429
	// is retried regardless because the server explicitly asked for it
	// and there's no duplicate-creation risk.
	c := newTestClient(server.URL, RetryPolicy{
		MaxAttempts: 3,
		BaseDelay:   time.Millisecond,
	})

	req, _ := http.NewRequest(http.MethodPost, server.URL+"/test", strings.NewReader("payload"))
	resp, err := c.httpClient.Do(req)
	if err != nil {
		t.Fatalf("Do: %v", err)
	}
	defer resp.Body.Close()
	_, _ = io.Copy(io.Discard, resp.Body)

	if resp.StatusCode != http.StatusOK {
		t.Errorf("status = %d, want 200", resp.StatusCode)
	}
	if got := atomic.LoadInt32(&attempts); got != 2 {
		t.Errorf("POST 429 attempts = %d, want 2", got)
	}
}

func TestTransport_NoRetryPOSTByDefault(t *testing.T) {
	var attempts int32
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		atomic.AddInt32(&attempts, 1)
		w.WriteHeader(http.StatusServiceUnavailable)
	}))
	defer server.Close()

	c := newTestClient(server.URL, RetryPolicy{
		MaxAttempts: 5,
		BaseDelay:   time.Millisecond,
	})

	req, _ := http.NewRequest(http.MethodPost, server.URL+"/test", strings.NewReader("payload"))
	resp, err := c.httpClient.Do(req)
	if err != nil {
		t.Fatalf("Do: %v", err)
	}
	defer resp.Body.Close()
	_, _ = io.Copy(io.Discard, resp.Body)

	if got := atomic.LoadInt32(&attempts); got != 1 {
		t.Errorf("POST attempts = %d, want 1 (POST should not retry by default)", got)
	}
}

func TestTransport_RetryPOSTWhenEnabled(t *testing.T) {
	var attempts int32
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		n := atomic.AddInt32(&attempts, 1)
		if n < 2 {
			w.WriteHeader(http.StatusServiceUnavailable)
			return
		}
		w.WriteHeader(http.StatusOK)
	}))
	defer server.Close()

	c := newTestClient(server.URL, RetryPolicy{
		MaxAttempts: 3,
		BaseDelay:   time.Millisecond,
		RetryOnPOST: true,
	})

	req, _ := http.NewRequest(http.MethodPost, server.URL+"/test", strings.NewReader("payload"))
	resp, err := c.httpClient.Do(req)
	if err != nil {
		t.Fatalf("Do: %v", err)
	}
	defer resp.Body.Close()
	_, _ = io.Copy(io.Discard, resp.Body)

	if resp.StatusCode != http.StatusOK {
		t.Errorf("status = %d, want 200", resp.StatusCode)
	}
	if got := atomic.LoadInt32(&attempts); got != 2 {
		t.Errorf("POST attempts = %d, want 2", got)
	}
}

func TestTransport_RespectsMaxAttempts(t *testing.T) {
	var attempts int32
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		atomic.AddInt32(&attempts, 1)
		w.WriteHeader(http.StatusServiceUnavailable)
	}))
	defer server.Close()

	c := newTestClient(server.URL, RetryPolicy{
		MaxAttempts: 4,
		BaseDelay:   time.Millisecond,
		MaxDelay:    5 * time.Millisecond,
	})

	resp, err := c.httpClient.Get(server.URL + "/test")
	if err != nil {
		t.Fatalf("Get: %v", err)
	}
	defer resp.Body.Close()
	_, _ = io.Copy(io.Discard, resp.Body)

	if got := atomic.LoadInt32(&attempts); got != 4 {
		t.Errorf("attempts = %d, want 4 (MaxAttempts)", got)
	}
	if resp.StatusCode != http.StatusServiceUnavailable {
		t.Errorf("status = %d, want 503 (final response surfaced)", resp.StatusCode)
	}
}

func TestTransport_DisabledByZeroPolicy(t *testing.T) {
	var attempts int32
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		atomic.AddInt32(&attempts, 1)
		w.WriteHeader(http.StatusServiceUnavailable)
	}))
	defer server.Close()

	c := newTestClient(server.URL, RetryPolicy{}) // zero ⇒ no retry

	resp, err := c.httpClient.Get(server.URL + "/test")
	if err != nil {
		t.Fatalf("Get: %v", err)
	}
	defer resp.Body.Close()
	_, _ = io.Copy(io.Discard, resp.Body)

	if got := atomic.LoadInt32(&attempts); got != 1 {
		t.Errorf("attempts = %d, want 1 (zero policy disables retry)", got)
	}
}

func TestTransport_HonorsRetryInBody(t *testing.T) {
	var firstAt, secondAt time.Time
	var attempts int32
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		n := atomic.AddInt32(&attempts, 1)
		if n == 1 {
			firstAt = time.Now()
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusTooManyRequests)
			fmt.Fprintln(w, `{"error_code":-343,"message":"quota","retry_in":1}`)
			return
		}
		secondAt = time.Now()
		w.WriteHeader(http.StatusOK)
	}))
	defer server.Close()

	c := newTestClient(server.URL, RetryPolicy{
		MaxAttempts: 2,
		BaseDelay:   time.Millisecond,
		MaxDelay:    5 * time.Second,
	})

	resp, err := c.httpClient.Get(server.URL + "/test")
	if err != nil {
		t.Fatalf("Get: %v", err)
	}
	defer resp.Body.Close()
	_, _ = io.Copy(io.Discard, resp.Body)

	elapsed := secondAt.Sub(firstAt)
	if elapsed < 900*time.Millisecond {
		t.Errorf("elapsed = %v, want ≥1s (server said retry_in: 1)", elapsed)
	}
}

func TestTransport_RetryRespectsContextCancellation(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		w.WriteHeader(http.StatusServiceUnavailable)
	}))
	defer server.Close()

	c := newTestClient(server.URL, RetryPolicy{
		MaxAttempts: 10,
		BaseDelay:   500 * time.Millisecond,
		MaxDelay:    5 * time.Second,
	})

	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Millisecond)
	defer cancel()
	req, _ := http.NewRequestWithContext(ctx, http.MethodGet, server.URL+"/test", nil)
	_, err := c.httpClient.Do(req)
	if err == nil {
		t.Fatal("expected context error, got nil")
	}
}
