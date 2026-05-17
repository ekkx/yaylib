package yaylib

import (
	"context"
	"net/http"
	"testing"
	"time"
)

// Retry-policy behavior parity. Driven against the shared behavior
// server: the scenario engine enforces the failure pattern and the
// test asserts the observable outcome (final status, the echoed
// X-Mock-Attempt ordinal, elapsed time, context error) rather than a
// local request counter.

func TestTransport_RetriesOn5xx(t *testing.T) {
	c := mockClient(t, "fail-503-times-2", WithRetryPolicy(RetryPolicy{
		MaxAttempts: 5,
		BaseDelay:   time.Millisecond,
		MaxDelay:    10 * time.Millisecond,
	}))

	resp, err := mockGet(t, c)
	if err != nil {
		t.Fatalf("Get: %v", err)
	}
	if resp.StatusCode != http.StatusOK {
		t.Errorf("status = %d, want 200 (retried past the two 503s)", resp.StatusCode)
	}
	if got := resp.Header.Get(hdrDiagAttempt); got != "3" {
		t.Errorf("attempt = %q, want 3 (two failures + the success)", got)
	}
	if got := resp.Header.Get(hdrDiagBranch); got != "fail-exhausted-happy" {
		t.Errorf("branch = %q, want fail-exhausted-happy", got)
	}
}

// PORTING:S14
func TestTransport_RetriesPOSTOn429(t *testing.T) {
	// RetryOnPOST=false (default): POST 5xx is not retried, but a 429
	// is retried regardless because the server explicitly asked for it
	// and there's no duplicate-creation risk.
	c := mockClient(t, "fail-429-times-1", WithRetryPolicy(RetryPolicy{
		MaxAttempts: 3,
		BaseDelay:   time.Millisecond,
	}))

	resp, err := mockPost(t, c)
	if err != nil {
		t.Fatalf("Do: %v", err)
	}
	if resp.StatusCode != http.StatusOK {
		t.Errorf("status = %d, want 200", resp.StatusCode)
	}
	if got := resp.Header.Get(hdrDiagAttempt); got != "2" {
		t.Errorf("attempt = %q, want 2 (POST 429 retried once)", got)
	}
}

// PORTING:S14
func TestTransport_NoRetryPOSTByDefault(t *testing.T) {
	// fail-503-times-1 would succeed on a retry, but POST 5xx is not
	// retried by default, so the single 503 surfaces.
	c := mockClient(t, "fail-503-times-1", WithRetryPolicy(RetryPolicy{
		MaxAttempts: 5,
		BaseDelay:   time.Millisecond,
	}))

	resp, err := mockPost(t, c)
	if err != nil {
		t.Fatalf("Do: %v", err)
	}
	if resp.StatusCode != http.StatusServiceUnavailable {
		t.Errorf("status = %d, want 503 (POST not retried)", resp.StatusCode)
	}
	if got := resp.Header.Get(hdrDiagAttempt); got != "1" {
		t.Errorf("attempt = %q, want 1 (no retry)", got)
	}
}

func TestTransport_RetryPOSTWhenEnabled(t *testing.T) {
	c := mockClient(t, "fail-503-times-1", WithRetryPolicy(RetryPolicy{
		MaxAttempts: 3,
		BaseDelay:   time.Millisecond,
		RetryOnPOST: true,
	}))

	resp, err := mockPost(t, c)
	if err != nil {
		t.Fatalf("Do: %v", err)
	}
	if resp.StatusCode != http.StatusOK {
		t.Errorf("status = %d, want 200", resp.StatusCode)
	}
	if got := resp.Header.Get(hdrDiagAttempt); got != "2" {
		t.Errorf("attempt = %q, want 2 (POST retried when enabled)", got)
	}
}

func TestTransport_RespectsMaxAttempts(t *testing.T) {
	c := mockClient(t, "fail-503-times-99", WithRetryPolicy(RetryPolicy{
		MaxAttempts: 4,
		BaseDelay:   time.Millisecond,
		MaxDelay:    5 * time.Millisecond,
	}))

	resp, err := mockGet(t, c)
	if err != nil {
		t.Fatalf("Get: %v", err)
	}
	if resp.StatusCode != http.StatusServiceUnavailable {
		t.Errorf("status = %d, want 503 (final response surfaced)", resp.StatusCode)
	}
	if got := resp.Header.Get(hdrDiagAttempt); got != "4" {
		t.Errorf("attempt = %q, want 4 (stopped at MaxAttempts)", got)
	}
}

func TestTransport_DisabledByZeroPolicy(t *testing.T) {
	c := mockClient(t, "fail-503-times-99", WithRetryPolicy(RetryPolicy{})) // zero ⇒ no retry

	resp, err := mockGet(t, c)
	if err != nil {
		t.Fatalf("Get: %v", err)
	}
	if got := resp.Header.Get(hdrDiagAttempt); got != "1" {
		t.Errorf("attempt = %q, want 1 (zero policy disables retry)", got)
	}
}

// PORTING:S12
func TestTransport_HonorsRetryInBody(t *testing.T) {
	c := mockClient(t, "retry-after-1", WithRetryPolicy(RetryPolicy{
		MaxAttempts: 2,
		BaseDelay:   time.Millisecond,
		MaxDelay:    5 * time.Second,
	}))

	start := time.Now()
	resp, err := mockGet(t, c)
	if err != nil {
		t.Fatalf("Get: %v", err)
	}
	elapsed := time.Since(start)

	if resp.StatusCode != http.StatusOK {
		t.Errorf("status = %d, want 200 (retried after the 429)", resp.StatusCode)
	}
	if elapsed < 900*time.Millisecond {
		t.Errorf("elapsed = %v, want ≥1s (server said retry_in: 1)", elapsed)
	}
}

// PORTING:S13
func TestTransport_RetryRespectsContextCancellation(t *testing.T) {
	c := mockClient(t, "fail-503-times-99", WithRetryPolicy(RetryPolicy{
		MaxAttempts: 10,
		BaseDelay:   500 * time.Millisecond,
		MaxDelay:    5 * time.Second,
	}))

	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Millisecond)
	defer cancel()
	req, _ := http.NewRequestWithContext(ctx, http.MethodGet, mockBaseURL()+parityGETPath, nil)
	resp, err := c.httpClient.Do(req)
	if resp != nil {
		resp.Body.Close()
	}
	if err == nil {
		t.Fatal("expected context error, got nil")
	}
}
