package mock

import (
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func newTestServer(t *testing.T) *httptest.Server {
	t.Helper()
	s, err := New()
	if err != nil {
		t.Fatalf("New: %v", err)
	}
	return httptest.NewServer(s)
}

// do issues a request with optional scenario/session headers.
func do(t *testing.T, ts *httptest.Server, method, path, scenario, session string) *http.Response {
	t.Helper()
	req, err := http.NewRequest(method, ts.URL+path, nil)
	if err != nil {
		t.Fatalf("NewRequest: %v", err)
	}
	if scenario != "" {
		req.Header.Set(headerScenario, scenario)
	}
	if session != "" {
		req.Header.Set(headerSession, session)
	}
	resp, err := ts.Client().Do(req)
	if err != nil {
		t.Fatalf("Do: %v", err)
	}
	return resp
}

func body(t *testing.T, r *http.Response) string {
	t.Helper()
	b, _ := io.ReadAll(r.Body)
	r.Body.Close()
	return string(b)
}

func TestAdminRoutes(t *testing.T) {
	ts := newTestServer(t)
	defer ts.Close()

	r := do(t, ts, http.MethodGet, "/__mock/health", "", "")
	if r.StatusCode != 200 || !strings.Contains(body(t, r), `"status":"ok"`) {
		t.Fatalf("health = %d", r.StatusCode)
	}
	r = do(t, ts, http.MethodPost, "/__mock/reset", "", "")
	if r.StatusCode != http.StatusNoContent {
		t.Fatalf("reset = %d, want 204", r.StatusCode)
	}
}

func TestBadScenarioIs400(t *testing.T) {
	ts := newTestServer(t)
	defer ts.Close()
	r := do(t, ts, http.MethodGet, "/api/apps/x", "totally-bogus", "s1")
	if r.StatusCode != http.StatusBadRequest {
		t.Fatalf("bad scenario = %d, want 400", r.StatusCode)
	}
}

func TestNoScenarioTokenEndpointReturnsLoginToken(t *testing.T) {
	ts := newTestServer(t)
	defer ts.Close()
	r := do(t, ts, http.MethodPost, tokenPath, "", "")
	if r.StatusCode != 200 {
		t.Fatalf("token = %d", r.StatusCode)
	}
	var tok tokenResponse
	if err := json.Unmarshal([]byte(body(t, r)), &tok); err != nil {
		t.Fatalf("decode token: %v", err)
	}
	if tok.AccessToken != freshAccessToken || tok.RefreshToken == "" {
		t.Fatalf("unexpected token body: %+v", tok)
	}
}

func TestNoScenarioOtherRouteFallsThroughTo200(t *testing.T) {
	ts := newTestServer(t)
	defer ts.Close()
	r := do(t, ts, http.MethodGet, "/api/apps/anything", "", "")
	if r.StatusCode != 200 {
		t.Fatalf("fallthrough = %d, want 200", r.StatusCode)
	}
}

func TestUnknownRouteIs404(t *testing.T) {
	ts := newTestServer(t)
	defer ts.Close()
	r := do(t, ts, http.MethodGet, "/no/such/route", "", "")
	if r.StatusCode != http.StatusNotFound {
		t.Fatalf("unknown route = %d, want 404", r.StatusCode)
	}
}

func TestFailTimesThenSuccess(t *testing.T) {
	ts := newTestServer(t)
	defer ts.Close()
	const sc, sess = "fail-503-times-2", "ft1"
	for i := 1; i <= 2; i++ {
		r := do(t, ts, http.MethodGet, "/api/apps/x", sc, sess)
		if r.StatusCode != 503 {
			t.Fatalf("attempt %d = %d, want 503", i, r.StatusCode)
		}
		if !strings.Contains(body(t, r), "mock scenario fail-503-times-2") {
			t.Fatalf("attempt %d missing CommonErrorResponse body", i)
		}
	}
	r := do(t, ts, http.MethodGet, "/api/apps/x", sc, sess)
	if r.StatusCode != 200 {
		t.Fatalf("post-budget attempt = %d, want 200", r.StatusCode)
	}
}

func TestFailTimesIsSessionScoped(t *testing.T) {
	ts := newTestServer(t)
	defer ts.Close()
	const sc = "fail-500-times-1"
	if r := do(t, ts, http.MethodGet, "/api/apps/x", sc, "A"); r.StatusCode != 500 {
		t.Fatalf("session A first = %d, want 500", r.StatusCode)
	}
	// Fresh session must get its own budget, not the exhausted one.
	if r := do(t, ts, http.MethodGet, "/api/apps/x", sc, "B"); r.StatusCode != 500 {
		t.Fatalf("session B first = %d, want 500", r.StatusCode)
	}
	if r := do(t, ts, http.MethodGet, "/api/apps/x", sc, "A"); r.StatusCode != 200 {
		t.Fatalf("session A second = %d, want 200", r.StatusCode)
	}
}

func TestRetryAfterFirstThenSuccess(t *testing.T) {
	ts := newTestServer(t)
	defer ts.Close()
	const sc, sess = "retry-after-7", "ra1"
	r := do(t, ts, http.MethodGet, "/api/apps/x", sc, sess)
	if r.StatusCode != http.StatusTooManyRequests {
		t.Fatalf("first = %d, want 429", r.StatusCode)
	}
	if got := r.Header.Get("Retry-After"); got != "7" {
		t.Fatalf("Retry-After = %q, want 7", got)
	}
	if !strings.Contains(body(t, r), `"retry_in":7`) {
		t.Fatal("body missing retry_in=7")
	}
	r = do(t, ts, http.MethodGet, "/api/apps/x", sc, sess)
	if r.StatusCode != 200 {
		t.Fatalf("second = %d, want 200", r.StatusCode)
	}
}

func TestRatelimitQuotaZeroAlways429(t *testing.T) {
	ts := newTestServer(t)
	defer ts.Close()
	for i := 0; i < 3; i++ {
		r := do(t, ts, http.MethodGet, "/api/apps/x", "ratelimit-quota-zero", "rl")
		if r.StatusCode != http.StatusTooManyRequests {
			t.Fatalf("call %d = %d, want 429", i, r.StatusCode)
		}
		if !strings.Contains(body(t, r), `"remaining_quota":0`) {
			t.Fatalf("call %d missing remaining_quota=0", i)
		}
	}
}

func TestExpiredTokenRefreshChain(t *testing.T) {
	ts := newTestServer(t)
	defer ts.Close()
	const sc, sess = "expired-token", "et1"

	// Protected request before refresh -> 401 access-token-expired.
	r := do(t, ts, http.MethodGet, "/api/apps/x", sc, sess)
	if r.StatusCode != http.StatusUnauthorized {
		t.Fatalf("pre-refresh = %d, want 401", r.StatusCode)
	}
	if !strings.Contains(body(t, r), `"error_code":-3`) {
		t.Fatal("401 body missing AccessTokenExpired code")
	}

	// SDK refreshes via the token endpoint -> 200 with a new access token.
	r = do(t, ts, http.MethodPost, tokenPath, sc, sess)
	if r.StatusCode != 200 {
		t.Fatalf("refresh = %d, want 200", r.StatusCode)
	}
	if !strings.Contains(body(t, r), refreshedAccessToken) {
		t.Fatal("refresh did not issue a new access token")
	}

	// Retry of the original request now succeeds.
	r = do(t, ts, http.MethodGet, "/api/apps/x", sc, sess)
	if r.StatusCode != 200 {
		t.Fatalf("post-refresh = %d, want 200", r.StatusCode)
	}

	// A different session is still unauthenticated.
	r = do(t, ts, http.MethodGet, "/api/apps/x", sc, "other")
	if r.StatusCode != http.StatusUnauthorized {
		t.Fatalf("other session = %d, want 401", r.StatusCode)
	}
}

func TestBodyMalformedIs200NonObject(t *testing.T) {
	ts := newTestServer(t)
	defer ts.Close()
	r := do(t, ts, http.MethodGet, "/api/apps/x", "body-malformed", "bm")
	if r.StatusCode != 200 {
		t.Fatalf("status = %d, want 200", r.StatusCode)
	}
	raw := body(t, r)
	// Decodes as a JSON string, not into a struct: typed Execute
	// would fail while ExecuteRaw returns these bytes.
	if !strings.Contains(raw, `"mock-malformed-body"`) {
		t.Fatalf("body = %q, want JSON scalar", raw)
	}
	var obj map[string]any
	if json.Unmarshal([]byte(raw), &obj) == nil {
		t.Fatal("body unexpectedly decoded into an object")
	}
}

func TestBodyUnknownEnumIs200WithOutOfSpecValue(t *testing.T) {
	ts := newTestServer(t)
	defer ts.Close()
	r := do(t, ts, http.MethodGet, "/api/apps/x", "body-unknown-enum", "ue")
	if r.StatusCode != 200 {
		t.Fatalf("status = %d, want 200", r.StatusCode)
	}
	var obj map[string]string
	if err := json.Unmarshal([]byte(body(t, r)), &obj); err != nil {
		t.Fatalf("decode: %v", err)
	}
	if obj["post_type"] != unknownEnumValue {
		t.Fatalf("post_type = %q, want %q", obj["post_type"], unknownEnumValue)
	}
}

func TestDiagnosticHeaders(t *testing.T) {
	ts := newTestServer(t)
	defer ts.Close()

	// Stateful scenario: branch + attempt ordinal are reported.
	r := do(t, ts, http.MethodGet, "/api/apps/x", "fail-503-times-2", "dg")
	body(t, r)
	if got := r.Header.Get("X-Mock-Scenario"); got != "fail-503-times-2" {
		t.Fatalf("X-Mock-Scenario = %q", got)
	}
	if got := r.Header.Get("X-Mock-Branch"); got != "fail" {
		t.Fatalf("X-Mock-Branch = %q, want fail", got)
	}
	if got := r.Header.Get("X-Mock-Attempt"); got != "1" {
		t.Fatalf("X-Mock-Attempt = %q, want 1", got)
	}
	r = do(t, ts, http.MethodGet, "/api/apps/x", "fail-503-times-2", "dg")
	body(t, r)
	if got := r.Header.Get("X-Mock-Attempt"); got != "2" {
		t.Fatalf("second attempt header = %q, want 2", got)
	}
	r = do(t, ts, http.MethodGet, "/api/apps/x", "fail-503-times-2", "dg")
	body(t, r)
	if got := r.Header.Get("X-Mock-Branch"); got != "fail-exhausted-happy" {
		t.Fatalf("post-budget branch = %q, want fail-exhausted-happy", got)
	}

	// No-scenario happy path still reports its branch.
	r = do(t, ts, http.MethodGet, "/api/apps/x", "", "")
	body(t, r)
	if got := r.Header.Get("X-Mock-Branch"); got != "happy" {
		t.Fatalf("no-scenario branch = %q, want happy", got)
	}

	// Admin route is labeled too.
	r = do(t, ts, http.MethodGet, "/__mock/health", "", "")
	body(t, r)
	if got := r.Header.Get("X-Mock-Branch"); got != "admin-health" {
		t.Fatalf("admin branch = %q, want admin-health", got)
	}
}

func TestResetClearsState(t *testing.T) {
	ts := newTestServer(t)
	defer ts.Close()
	const sc, sess = "fail-500-times-1", "rs1"
	if r := do(t, ts, http.MethodGet, "/api/apps/x", sc, sess); r.StatusCode != 500 {
		t.Fatalf("first = %d, want 500", r.StatusCode)
	}
	if r := do(t, ts, http.MethodGet, "/api/apps/x", sc, sess); r.StatusCode != 200 {
		t.Fatalf("second = %d, want 200 (budget spent)", r.StatusCode)
	}
	do(t, ts, http.MethodPost, "/__mock/reset", "", "")
	if r := do(t, ts, http.MethodGet, "/api/apps/x", sc, sess); r.StatusCode != 500 {
		t.Fatalf("after reset = %d, want 500 (budget restored)", r.StatusCode)
	}
}
