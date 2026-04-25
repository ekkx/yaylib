package yaylib

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"math/rand"
	"net/http"
	"time"
)

// RetryPolicy controls how Transport retries failed requests on 5xx, 429,
// and network errors. The zero value disables retries entirely
// (MaxAttempts == 0 ⇒ a single attempt and no retry); use
// DefaultRetryPolicy() for the values NewClient applies by default.
type RetryPolicy struct {
	// MaxAttempts caps the total number of attempts (including the
	// initial one). Values < 2 disable retry — the request is sent at
	// most once.
	MaxAttempts int

	// BaseDelay is the starting delay for exponential backoff. The
	// actual sleep is BaseDelay * 2^(attempt-2) bounded by MaxDelay,
	// with full jitter (random uniform in [0, computed]).
	BaseDelay time.Duration

	// MaxDelay caps any single sleep between attempts (also applied to
	// server-suggested delays).
	MaxDelay time.Duration

	// RetryOnPOST allows POST and PATCH requests to be retried on 5xx
	// and network errors. Off by default because Yay! has no
	// idempotency-key concept and a retry could double-create resources
	// (post, message, follow). Turn this on only when you've verified
	// the endpoints you call are idempotent.
	//
	// 429 (TooManyRequests) responses always retry regardless of this
	// setting — the server has explicitly requested the retry and the
	// request was rejected before processing, so there is no
	// duplicate-creation risk.
	RetryOnPOST bool
}

// DefaultRetryPolicy returns the policy NewClient applies when callers do
// not override it: 3 attempts, 200ms base delay, 30s ceiling, no POST
// retry.
func DefaultRetryPolicy() RetryPolicy {
	return RetryPolicy{
		MaxAttempts: 3,
		BaseDelay:   200 * time.Millisecond,
		MaxDelay:    30 * time.Second,
		RetryOnPOST: false,
	}
}

// shouldRetry reports whether a (resp, err) pair from one HTTP attempt
// merits another attempt under the given policy. It does not consult the
// attempt counter — the caller is responsible for stopping at MaxAttempts.
func shouldRetry(resp *http.Response, err error, method string, policy RetryPolicy) bool {
	// 429 means the server explicitly asked us to retry. The request was
	// rejected at the rate limiter before any processing, so there is no
	// duplicate-creation risk — retry regardless of method.
	if err == nil && resp != nil && resp.StatusCode == http.StatusTooManyRequests {
		return true
	}
	// 5xx and network failures carry uncertainty about server-side state
	// (the request may or may not have been processed before the failure).
	// Gate by method to avoid double-creating resources on POST/PATCH.
	if !canRetryMethod(method, policy) {
		return false
	}
	if err != nil {
		return true
	}
	if resp == nil {
		return false
	}
	// 401 is owned by the auth-refresh path; surviving 401s don't fix
	// themselves with more retries.
	if resp.StatusCode == http.StatusUnauthorized {
		return false
	}
	if resp.StatusCode >= 500 && resp.StatusCode <= 599 {
		return true
	}
	return false
}

// canRetryMethod gates retries on HTTP method semantics. Methods that are
// idempotent per RFC 7231 are always retriable; POST/PATCH require
// RetryOnPOST.
func canRetryMethod(method string, policy RetryPolicy) bool {
	switch method {
	case http.MethodGet, http.MethodHead, http.MethodPut, http.MethodDelete, http.MethodOptions:
		return true
	case http.MethodPost, http.MethodPatch:
		return policy.RetryOnPOST
	default:
		return false
	}
}

// nextDelay computes the wait before the upcoming attempt (>= 2). It
// honors the server-suggested delay (the JSON body's `retry_in` field)
// when present, and otherwise applies exponential backoff with full
// jitter, capped by policy.MaxDelay.
//
// Reading retry_in consumes resp.Body. To stay compatible with the
// caller's drain step we replace resp.Body with a fresh bytes.Reader
// holding the same bytes.
func nextDelay(prev *http.Response, attempt int, policy RetryPolicy) time.Duration {
	if d := retryInFromBody(prev); d > 0 {
		if policy.MaxDelay > 0 && d > policy.MaxDelay {
			return policy.MaxDelay
		}
		return d
	}
	if policy.BaseDelay <= 0 {
		return 0
	}
	exp := policy.BaseDelay << uint(attempt-2) // attempt=2 → base, =3 → base*2, ...
	if exp <= 0 || exp > policy.MaxDelay {
		exp = policy.MaxDelay
	}
	if exp <= 0 {
		return 0
	}
	return time.Duration(rand.Int63n(int64(exp) + 1))
}

// retryInFromBody parses the JSON body's `retry_in` field (in seconds).
// Returns 0 when the body doesn't carry one. The Yay! server uses this
// field rather than the standard Retry-After header to communicate
// retry timing.
func retryInFromBody(resp *http.Response) time.Duration {
	if resp == nil || resp.StatusCode < 400 || resp.Body == nil {
		return 0
	}
	body, err := io.ReadAll(io.LimitReader(resp.Body, 16*1024))
	if err != nil {
		return 0
	}
	resp.Body = io.NopCloser(bytes.NewReader(body))
	var er ErrorResponse
	if err := json.Unmarshal(body, &er); err != nil {
		return 0
	}
	if er.RetryIn == nil || *er.RetryIn <= 0 {
		return 0
	}
	return time.Duration(*er.RetryIn) * time.Second
}

// sleepWithCtx blocks for d, returning false if the context fires first.
// d <= 0 returns true immediately.
func sleepWithCtx(ctx context.Context, d time.Duration) bool {
	if d <= 0 {
		return true
	}
	t := time.NewTimer(d)
	defer t.Stop()
	select {
	case <-t.C:
		return true
	case <-ctx.Done():
		return false
	}
}

// drainAndClose drains the body to enable connection reuse and closes it.
// Safe on nil response or nil body.
func drainAndClose(resp *http.Response) {
	if resp == nil || resp.Body == nil {
		return
	}
	io.Copy(io.Discard, resp.Body)
	resp.Body.Close()
}
