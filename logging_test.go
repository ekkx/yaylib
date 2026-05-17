package yaylib

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log/slog"
	"net/http"
	"net/http/httptest"
	"strings"
	"sync"
	"sync/atomic"
	"testing"
)

// captureHandler records every slog record so a test can assert which
// events were emitted and that no banned value ever reached the logger.
type captureHandler struct {
	mu      sync.Mutex
	records []capturedRecord
}

type capturedRecord struct {
	level slog.Level
	msg   string
	attrs map[string]string
}

func (h *captureHandler) Enabled(context.Context, slog.Level) bool { return true }

func (h *captureHandler) Handle(_ context.Context, r slog.Record) error {
	rec := capturedRecord{level: r.Level, msg: r.Message, attrs: map[string]string{}}
	r.Attrs(func(a slog.Attr) bool {
		rec.attrs[a.Key] = a.Value.String()
		return true
	})
	h.mu.Lock()
	h.records = append(h.records, rec)
	h.mu.Unlock()
	return nil
}

func (h *captureHandler) WithAttrs([]slog.Attr) slog.Handler { return h }
func (h *captureHandler) WithGroup(string) slog.Handler      { return h }

func (h *captureHandler) snapshot() []capturedRecord {
	h.mu.Lock()
	defer h.mu.Unlock()
	return append([]capturedRecord(nil), h.records...)
}

func (h *captureHandler) byEvent(event string) (capturedRecord, bool) {
	for _, r := range h.snapshot() {
		if r.attrs["event"] == event {
			return r, true
		}
	}
	return capturedRecord{}, false
}

// Default client is silent: the discard handler reports not-enabled at
// every level, so slog never formats a record and nothing is written.
func TestLogging_DefaultIsSilent(t *testing.T) {
	c := NewClient()
	for _, lvl := range []slog.Level{slog.LevelDebug, slog.LevelInfo, slog.LevelWarn, slog.LevelError} {
		if c.Logger.Enabled(context.Background(), lvl) {
			t.Errorf("default logger enabled at %v; library must be silent by default", lvl)
		}
	}
}

// With a logger injected, the 401 → refresh → persist-failure path
// emits the contracted events and never leaks a token / password / API
// key into any record (PORTING.md §12.2 + §15).
func TestLogging_RefreshPersistFailEventAndRedaction(t *testing.T) {
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
		if atomic.AddInt32(&dataAttempt, 1) == 1 {
			w.WriteHeader(http.StatusUnauthorized)
			_, _ = w.Write([]byte(`{}`))
			return
		}
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{}`))
	}}
	srv := httptest.NewServer(cap)
	defer srv.Close()

	h := &captureHandler{}
	c := NewClient(
		WithBaseURL(srv.URL),
		WithRetryPolicy(RetryPolicy{}),
		WithSessionStore(&errorStore{saveErr: errors.New("disk full")}),
		WithLogger(slog.New(h)),
	)
	c.SetTokens("STALE", "REF")
	c.currentEmail = "foo@example.com" // enable the persist branch

	req, _ := http.NewRequestWithContext(context.Background(), http.MethodGet, srv.URL+"/v1/anywhere", nil)
	resp, err := c.httpClient.Do(req)
	if err != nil {
		t.Fatalf("Do: %v", err)
	}
	resp.Body.Close()

	// Contracted events present.
	if rec, ok := h.byEvent("token_persist_fail"); !ok {
		t.Error("no token_persist_fail record emitted")
	} else if rec.level != slog.LevelWarn {
		t.Errorf("token_persist_fail level = %v, want WARN", rec.level)
	}
	if rec, ok := h.byEvent("token_refresh"); !ok || rec.attrs["outcome"] != "ok" {
		t.Errorf("token_refresh ok record missing; got %+v ok=%v", rec, ok)
	}
	if _, ok := h.byEvent("http_request"); !ok {
		t.Error("no http_request debug record emitted")
	}

	// Redaction: no record (message or any attr) carries a secret.
	banned := []string{"STALE", "REF", "NEW_ACC", "NEW_REF", c.APIKey}
	for _, r := range h.snapshot() {
		hay := r.msg
		for k, v := range r.attrs {
			hay += "\x00" + k + "=" + v
		}
		for _, b := range banned {
			if b != "" && strings.Contains(hay, b) {
				t.Errorf("record %q leaked banned value %q: %s",
					r.msg, b, fmt.Sprint(r.attrs))
			}
		}
	}
}
