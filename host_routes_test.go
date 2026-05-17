package yaylib

import (
	"context"
	"net/http"
	"net/http/httptest"
	"sync"
	"testing"
)

// A few operations are served from an auxiliary host rather than the
// primary API host. The transport must send those — and only those —
// to the auxiliary base URL, leaving every other request on the
// primary host. These tests pin that routing against two stand-in
// servers so a regression (or a host-route table change) is caught
// without touching the network.

type hitRecorder struct {
	mu   sync.Mutex
	hits []string
}

func (h *hitRecorder) handler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		h.mu.Lock()
		h.hits = append(h.hits, r.Method+" "+r.URL.Path)
		h.mu.Unlock()
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{}`))
	})
}

func (h *hitRecorder) snapshot() []string {
	h.mu.Lock()
	defer h.mu.Unlock()
	return append([]string(nil), h.hits...)
}

func contains(ss []string, want string) bool {
	for _, s := range ss {
		if s == want {
			return true
		}
	}
	return false
}

func TestTransport_RoutesActivityFeedToAuxiliaryHost(t *testing.T) {
	primary := &hitRecorder{}
	aux := &hitRecorder{}
	ps := httptest.NewServer(primary.handler())
	defer ps.Close()
	as := httptest.NewServer(aux.handler())
	defer as.Close()

	c := NewClient(WithBaseURL(ps.URL), WithCassandraBaseURL(as.URL))
	c.SetTokens("ACC", "REF")

	if _, _, err := c.GetUserActivities(context.Background()).ExecuteRaw(); err != nil {
		t.Fatalf("GetUserActivities: %v", err)
	}

	const op = "GET /api/v2/user_activities"
	if !contains(aux.snapshot(), op) {
		t.Errorf("auxiliary host did not receive %q; got %v", op, aux.snapshot())
	}
	if contains(primary.snapshot(), op) {
		t.Errorf("primary host wrongly received host-routed %q", op)
	}
}

func TestTransport_LeavesUnroutedRequestsOnPrimaryHost(t *testing.T) {
	primary := &hitRecorder{}
	aux := &hitRecorder{}
	ps := httptest.NewServer(primary.handler())
	defer ps.Close()
	as := httptest.NewServer(aux.handler())
	defer as.Close()

	c := NewClient(WithBaseURL(ps.URL), WithCassandraBaseURL(as.URL))
	c.SetTokens("ACC", "REF")

	if _, _, err := c.GetUserTimestamp(context.Background()).ExecuteRaw(); err != nil {
		t.Fatalf("GetUserTimestamp: %v", err)
	}

	if len(aux.snapshot()) != 0 {
		t.Errorf("auxiliary host received an unrouted request: %v", aux.snapshot())
	}
	if len(primary.snapshot()) == 0 {
		t.Error("primary host received nothing; unrouted request was misdirected")
	}
}

func TestHostBaseURLFor(t *testing.T) {
	c := NewClient(WithCassandraBaseURL("https://aux.example"))

	if got, ok := c.hostBaseURLFor("GET", "/api/v2/user_activities"); !ok || got != "https://aux.example" {
		t.Errorf("host-routed op: got (%q, %v), want (%q, true)", got, ok, "https://aux.example")
	}
	if got, ok := c.hostBaseURLFor("GET", "/api/user_activities"); !ok || got != "https://aux.example" {
		t.Errorf("host-routed v1 op: got (%q, %v), want (%q, true)", got, ok, "https://aux.example")
	}
	if _, ok := c.hostBaseURLFor("GET", "/v2/users/timestamp"); ok {
		t.Error("non-routed op resolved to an auxiliary host")
	}
}
