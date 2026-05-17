package yaylib

import (
	"crypto/rand"
	"encoding/hex"
	"io"
	"net/http"
	"os"
	"strings"
	"testing"
)

// Behavior-parity harness.
//
// These tests assert the wrapper's retry / refresh / rate-limit /
// upload / event-stream behavior against a shared behavior-contract
// server (the same instance the TypeScript and Python suites use), so
// the three languages verify identical behavior instead of three
// divergent in-process fakes.
//
// The server is discovered from the environment. When the variables
// are unset the parity tests skip, so a standalone `go test ./.` still
// passes on the pure-function and typed-decode fixtures alone.
//
//	YAYLIB_MOCK_BASE_URL   http://host:port   (REST)
//	YAYLIB_MOCK_WS_URL     ws://host:port     (event stream)
//
// Behavior is selected per request with two headers the server reads:
// X-Yay-Mock-Scenario picks the failure/latency behavior and
// X-Yay-Mock-Session isolates each test's stateful counters. The
// server echoes diagnostics back on every response it owns:
// X-Mock-Branch (decision taken) and X-Mock-Attempt (the per-session
// attempt ordinal for counter-based scenarios).
const (
	envMockBaseURL = "YAYLIB_MOCK_BASE_URL"
	envMockWSURL   = "YAYLIB_MOCK_WS_URL"

	hdrScenario = "X-Yay-Mock-Scenario"
	hdrSession  = "X-Yay-Mock-Session"

	hdrDiagBranch  = "X-Mock-Branch"
	hdrDiagAttempt = "X-Mock-Attempt"

	// Real spec routes used to drive generic transport behavior. The
	// happy fall-through needs a route the server knows so it answers
	// 200; failure scenarios are path-independent. Neither is host
	// -routed, so requests stay on the primary base URL.
	parityGETPath  = "/v1/buckets/presigned_urls"
	parityPOSTPath = "/v1/calls/leave_agora_channel"
)

func mockBaseURL() string { return os.Getenv(envMockBaseURL) }
func mockWSURL() string   { return os.Getenv(envMockWSURL) }

// requireMock skips the calling test unless the shared server is
// configured, and returns its REST base URL.
func requireMock(t *testing.T) string {
	t.Helper()
	base := mockBaseURL()
	if base == "" {
		t.Skipf("%s unset; skipping behavior-parity test", envMockBaseURL)
	}
	return base
}

// newMockSession mints an opaque per-test key so the server's
// stateful scenario counters never collide across tests (or across
// the parallel three-language run).
func newMockSession() string {
	var b [16]byte
	_, _ = rand.Read(b[:])
	return hex.EncodeToString(b[:])
}

// scenarioRT stamps the scenario and session headers on every outbound
// request. It is installed as the http.Client transport, which the SDK
// wraps as its base hop — so the headers ride every real attempt,
// including token-refresh and retry replays.
type scenarioRT struct {
	next     http.RoundTripper
	scenario string
	session  string
}

func (s scenarioRT) RoundTrip(r *http.Request) (*http.Response, error) {
	r = r.Clone(r.Context())
	if s.scenario != "" {
		r.Header.Set(hdrScenario, s.scenario)
	}
	r.Header.Set(hdrSession, s.session)
	return s.next.RoundTrip(r)
}

// mockClient builds a Client pointed at the shared server with scenario
// active for this test's isolated session. Extra options are applied
// after the parity defaults so a test can override (e.g. RetryPolicy).
func mockClient(t *testing.T, scenario string, opts ...Option) *Client {
	t.Helper()
	base := requireMock(t)
	hc := &http.Client{
		Transport: scenarioRT{
			next:     http.DefaultTransport,
			scenario: scenario,
			session:  newMockSession(),
		},
	}
	defaults := []Option{
		WithBaseURL(base),
		WithHTTPClient(hc),
		WithClientIP("127.0.0.1"), // keep the lazy IP-fetch goroutine off the test path
	}
	if ws := mockWSURL(); ws != "" {
		defaults = append(defaults, WithEventStreamURL(ws))
	}
	return NewClient(append(defaults, opts...)...)
}

// mockStreamClient builds a Client whose event stream targets the
// shared server with the given per-dial WS mode: "" confirms each
// subscribe and pushes the channel's representative event; "reject"
// rejects every subscribe; "no-confirm" stays silent (drives the
// subscribe timeout); "drop-after-confirm" confirms + pushes then
// closes (drives reconnect / re-subscribe). The mode rides the
// event-stream URL query, which the SDK preserves while appending the
// token and app_version.
func mockStreamClient(t *testing.T, mode string, opts ...Option) *Client {
	t.Helper()
	requireMock(t)
	ws := mockWSURL()
	if ws == "" {
		t.Skipf("%s unset; skipping event-stream parity test", envMockWSURL)
	}
	if mode != "" {
		ws += "/?mock=" + mode
	}
	return mockClient(t, "", append([]Option{WithEventStreamURL(ws)}, opts...)...)
}

// mockGet drives a GET against the shared server's generic primary
// route through the SDK transport, draining and closing the body. The
// returned *http.Response keeps its headers (including the X-Mock-*
// diagnostics) for assertion.
func mockGet(t *testing.T, c *Client) (*http.Response, error) {
	t.Helper()
	resp, err := c.httpClient.Get(mockBaseURL() + parityGETPath)
	if resp != nil {
		_, _ = io.Copy(io.Discard, resp.Body)
		resp.Body.Close()
	}
	return resp, err
}

// mockGetRaw is mockGet but leaves the body open so the caller can
// read it (and must close it).
func mockGetRaw(t *testing.T, c *Client) (*http.Response, error) {
	t.Helper()
	return c.httpClient.Get(mockBaseURL() + parityGETPath)
}

// mockPost drives a POST against the shared server's generic primary
// route, draining and closing the body.
func mockPost(t *testing.T, c *Client) (*http.Response, error) {
	t.Helper()
	req, _ := http.NewRequest(http.MethodPost, mockBaseURL()+parityPOSTPath, strings.NewReader("payload"))
	resp, err := c.httpClient.Do(req)
	if resp != nil {
		_, _ = io.Copy(io.Discard, resp.Body)
		resp.Body.Close()
	}
	return resp, err
}

// resetMock clears every scenario counter on the shared server. Tests
// that mint a fresh session per case rarely need it; it exists for the
// few that reuse a session across phases.
func resetMock(t *testing.T) {
	t.Helper()
	base := requireMock(t)
	req, err := http.NewRequest(http.MethodPost, base+"/__mock/reset", nil)
	if err != nil {
		t.Fatalf("reset request: %v", err)
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		t.Fatalf("reset: %v", err)
	}
	resp.Body.Close()
}
