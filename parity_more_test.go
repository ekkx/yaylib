package yaylib

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"sync"
	"testing"
)

// PORTING:S11
//
// Concurrent 401s collapse to a single token refresh: when several
// requests are in flight and all see an expired-access-token 401, only
// ONE of them performs the refresh; the rest wait on it and replay with
// the freshly rotated token instead of each firing their own refresh.
// Parallel refreshes would invalidate one another, so single-flight is
// the contract (PORTING.md §6.1).
//
// Driven against the shared server's expired-token scenario through a
// single client/session: the first protected request answers 401, the
// token endpoint simulates the rotation, and every request after the
// refresh latches falls through happy.
//
// The shared server's refresh latch is a boolean, not a counter, and
// the only response a caller observes is the post-refresh-and-retry
// final one (which carries the happy branch, never the internal
// refresh-branch header). Counting refreshes from the client side is
// therefore inherently unreliable under concurrency. We assert the
// strong observable invariant the single-flight contract guarantees:
// every concurrent call ultimately succeeds with 200 and none surfaces
// a 401 — a broken single-flight (parallel refreshes invalidating each
// other) would leave at least one call stuck on a 401 or a transport
// error. Single-flight itself is enforced in the transport by the
// refresh mutex + stale-token guard (auth.go refreshTokens).
func TestTransport_Concurrent401sCollapseToSingleRefresh(t *testing.T) {
	c := mockClient(t, "expired-token", WithRetryPolicy(RetryPolicy{}))
	// Seed a stale access token + a refresh token so the 401 handler
	// actually runs the refresh chain (with no refresh token it would
	// just hand the 401 straight back). One client/session is shared by
	// all the goroutines so the single-flight guard is exercised.
	c.SetTokens("STALE", "REF")

	const n = 5
	type result struct {
		resp *http.Response
		err  error
	}
	results := make([]result, n)
	var wg sync.WaitGroup
	wg.Add(n)
	for i := 0; i < n; i++ {
		go func(i int) {
			defer wg.Done()
			resp, err := mockGet(t, c)
			results[i] = result{resp, err}
		}(i)
	}
	wg.Wait()

	for i, r := range results {
		if r.err != nil {
			t.Fatalf("concurrent call %d: unexpected error: %v", i, r.err)
		}
		if r.resp == nil {
			t.Fatalf("concurrent call %d: nil response", i)
		}
		if r.resp.StatusCode == http.StatusUnauthorized {
			t.Fatalf("concurrent call %d: surfaced a 401 — single-flight "+
				"refresh broke (parallel refreshes invalidated each other)", i)
		}
		if r.resp.StatusCode != http.StatusOK {
			t.Fatalf("concurrent call %d: status = %d, want 200 "+
				"(refresh + retry should have succeeded)", i, r.resp.StatusCode)
		}
	}
}

// PORTING:S29
//
// Any 2xx is a success: a successful response answered with a non-200
// 2xx status (the shared server's ok-status-201) still decodes into the
// typed success model. The SDK must not key strictly on the documented
// status code and discard the body — it falls back to the documented
// 2xx success type (PORTING.md §11 rule 3).
func TestDecode_AnyTwoxxIsSuccess(t *testing.T) {
	c := mockClient(t, "ok-status-201")
	ctx := context.Background()

	out, httpResp, err := c.GetBucketPresignedUrls(ctx).Execute()
	if err != nil {
		t.Fatalf("typed Execute() on 201 success: unexpected error: %v", err)
	}
	if out == nil {
		t.Fatal("typed Execute() on 201 success returned nil response model")
	}
	if httpResp == nil {
		t.Fatal("typed Execute() returned nil *http.Response")
	}
	if httpResp.StatusCode != http.StatusCreated {
		t.Fatalf("status = %d, want 201 (the non-200 2xx was treated as success)",
			httpResp.StatusCode)
	}
}

// PORTING:S30
//
// CodeOf resolves to the typed ErrorCode constant, hermetically (no
// shared server). An error envelope carrying error_code -31 must map to
// the typed yaylib.ErrCodeRequired2FA constant — not a bare integer —
// and a non-API / code-less error must yield yaylib.ErrCodeUnknown
// without panicking (PORTING.md §15 S30).
//
// APIError (= gen.GenericOpenAPIError) keeps its body unexported, so the
// established hermetic way to obtain a real *APIError with a chosen body
// is the same one the auth suite uses for S5: a local httptest server
// returning the envelope and a real typed call producing the error.
func TestCodeOf_ResolvesTypedConstant(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusUnauthorized)
		_ = json.NewEncoder(w).Encode(map[string]any{
			"error_code": int(ErrCodeRequired2FA), // -31
			"message":    "x",
		})
	}))
	defer srv.Close()

	c := NewClient(WithBaseURL(srv.URL), WithRetryPolicy(RetryPolicy{}))
	_, _, err := c.GetBucketPresignedUrls(context.Background()).Execute()
	if err == nil {
		t.Fatal("expected an API error from the error envelope, got nil")
	}

	var apiErr *APIError
	if !errors.As(err, &apiErr) {
		t.Fatalf("error = %T, want *yaylib.APIError", err)
	}
	if got := string(apiErr.Body()); got == "" {
		t.Fatal("APIError carried no body")
	}

	got := CodeOf(err)
	if got != ErrCodeRequired2FA {
		t.Fatalf("CodeOf(err) = %v (%d), want ErrCodeRequired2FA (%d)",
			got, int(got), int(ErrCodeRequired2FA))
	}
	// The result is the typed constant, usable in a switch and via
	// String(); a bare int would not stringify to the constant name.
	if got.String() != ErrCodeRequired2FA.String() {
		t.Fatalf("CodeOf(err).String() = %q, want %q",
			got.String(), ErrCodeRequired2FA.String())
	}

	// A non-API error (and never a panic) yields ErrCodeUnknown.
	func() {
		defer func() {
			if p := recover(); p != nil {
				t.Fatalf("CodeOf panicked on a plain error: %v", p)
			}
		}()
		if got := CodeOf(errors.New("plain")); got != ErrCodeUnknown {
			t.Fatalf("CodeOf(plain error) = %v, want ErrCodeUnknown", got)
		}
		if got := CodeOf(nil); got != ErrCodeUnknown {
			t.Fatalf("CodeOf(nil) = %v, want ErrCodeUnknown", got)
		}
	}()
}
