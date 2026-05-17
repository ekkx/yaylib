package yaylib

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"sync/atomic"
	"testing"
	"time"

	"github.com/coder/websocket"
)

// fakeServer stands up an HTTP+WebSocket endpoint that mimics the
// minimum ActionCable surface the SDK needs: a /v1/users/ws_token JSON
// handler and a WebSocket upgrade on /. Tests inject behavior through
// `onConnect`, which gets a fresh wsSession per accepted socket.
type fakeServer struct {
	srv       *httptest.Server
	onConnect func(s *wsSession)
	conns     atomic.Int32
}

type wsSession struct {
	t        *testing.T
	ws       *websocket.Conn
	ctx      context.Context
	received chan map[string]string
}

func newFakeServer(t *testing.T, onConnect func(*wsSession)) *fakeServer {
	t.Helper()
	fs := &fakeServer{onConnect: onConnect}
	mux := http.NewServeMux()
	mux.HandleFunc("/v1/users/ws_token", func(w http.ResponseWriter, _ *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"token":"test-token"}`))
	})
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		ws, err := websocket.Accept(w, r, &websocket.AcceptOptions{
			InsecureSkipVerify: true,
		})
		if err != nil {
			t.Logf("accept: %v", err)
			return
		}
		fs.conns.Add(1)
		ctx, cancel := context.WithTimeout(r.Context(), 10*time.Second)
		defer cancel()
		sess := &wsSession{
			t: t, ws: ws, ctx: ctx,
			received: make(chan map[string]string, 8),
		}
		// Read pump for incoming subscribe/unsubscribe commands.
		go func() {
			for {
				_, data, err := ws.Read(ctx)
				if err != nil {
					return
				}
				var msg map[string]string
				if json.Unmarshal(data, &msg) == nil {
					sess.received <- msg
				}
			}
		}()
		fs.onConnect(sess)
	})
	fs.srv = httptest.NewServer(mux)
	t.Cleanup(fs.srv.Close)
	return fs
}

func (s *wsSession) sendWelcome() {
	_ = s.ws.Write(s.ctx, websocket.MessageText, []byte(`{"type":"welcome","sid":"test-sid"}`))
}

func (s *wsSession) confirm(identifier string) {
	body, _ := json.Marshal(map[string]string{"identifier": identifier, "type": "confirm_subscription"})
	_ = s.ws.Write(s.ctx, websocket.MessageText, body)
}

// pushEvent wraps an event payload in the channel envelope the server
// sends for actual events.
func (s *wsSession) pushEvent(identifier, event string, data any) {
	dataJSON, _ := json.Marshal(data)
	innerJSON, _ := json.Marshal(map[string]any{
		"event": event,
		"data":  json.RawMessage(dataJSON),
	})
	frame, _ := json.Marshal(map[string]any{
		"identifier": identifier,
		"message":    json.RawMessage(innerJSON),
	})
	_ = s.ws.Write(s.ctx, websocket.MessageText, frame)
}

func (s *wsSession) close() {
	_ = s.ws.Close(websocket.StatusNormalClosure, "")
}

func newStreamTestClient(httpURL string) *Client {
	return NewClient(
		WithBaseURL(httpURL),
		WithEventStreamURL(strings.Replace(httpURL, "http://", "ws://", 1)),
		WithClientIP("127.0.0.1"),
	)
}

func TestEventStream_Unsubscribe(t *testing.T) {
	gotUnsub := make(chan struct{})
	fs := newFakeServer(t, func(s *wsSession) {
		s.sendWelcome()
		msg := <-s.received
		s.confirm(msg["identifier"])
		// Wait for unsubscribe.
		for {
			next := <-s.received
			if next["command"] == "unsubscribe" {
				close(gotUnsub)
				return
			}
		}
	})

	c := newStreamTestClient(fs.srv.URL)
	c.SetTokens("stub", "")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	conn, err := c.OpenEventStream(ctx, EventStreamOptions{Reconnect: ReconnectPolicy{Disabled: true}})
	if err != nil {
		t.Fatalf("OpenEventStream: %v", err)
	}
	defer conn.Close()

	sub, err := conn.Subscribe(ctx, ChatRoomChannel())
	if err != nil {
		t.Fatalf("Subscribe: %v", err)
	}
	if err := sub.Unsubscribe(ctx); err != nil {
		t.Fatalf("Unsubscribe: %v", err)
	}
	select {
	case <-gotUnsub:
	case <-time.After(2 * time.Second):
		t.Fatal("server did not see unsubscribe command")
	}
	select {
	case _, ok := <-sub.Events():
		if ok {
			t.Error("Events should be closed after Unsubscribe")
		}
	case <-time.After(time.Second):
		t.Error("Events channel did not close after Unsubscribe")
	}
}

func TestEventStream_ErrAfterReconnectExhausted(t *testing.T) {
	fs := newFakeServer(t, func(s *wsSession) {
		s.sendWelcome()
		// Drop immediately so the client tries to reconnect.
		s.close()
	})

	c := newStreamTestClient(fs.srv.URL)
	c.SetTokens("stub", "")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	conn, err := c.OpenEventStream(ctx, EventStreamOptions{
		Reconnect: ReconnectPolicy{
			MaxAttempts:  1,
			InitialDelay: 10 * time.Millisecond,
			MaxDelay:     20 * time.Millisecond,
		},
	})
	if err != nil {
		t.Fatalf("OpenEventStream: %v", err)
	}
	defer conn.Close()

	select {
	case <-conn.Done():
	case <-time.After(3 * time.Second):
		t.Fatal("Done did not fire after reconnect exhaustion")
	}
	if err := conn.Err(); err == nil {
		t.Error("Err after reconnect exhausted = nil, want non-nil")
	}
}

func TestEventStream_OnDropFiresWhenBufferFull(t *testing.T) {
	const ident = `{"channel":"ChatRoomChannel"}`
	send := make(chan struct{})
	fs := newFakeServer(t, func(s *wsSession) {
		s.sendWelcome()
		msg := <-s.received
		s.confirm(msg["identifier"])
		<-send
		// Push more events than the EventBuffer holds. The first
		// fills the buffer; the rest must drop.
		for i := 0; i < 5; i++ {
			s.pushEvent(ident, "chat_deleted", map[string]int64{"room_id": int64(i)})
		}
		<-s.ctx.Done()
	})

	c := newStreamTestClient(fs.srv.URL)
	c.SetTokens("stub", "")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var dropped atomic.Int32
	conn, err := c.OpenEventStream(ctx, EventStreamOptions{
		Reconnect:   ReconnectPolicy{Disabled: true},
		EventBuffer: 1,
		OnDrop:      func(Event) { dropped.Add(1) },
	})
	if err != nil {
		t.Fatalf("OpenEventStream: %v", err)
	}
	defer conn.Close()

	if _, err := conn.Subscribe(ctx, ChatRoomChannel()); err != nil {
		t.Fatalf("Subscribe: %v", err)
	}
	close(send) // unleash the burst — we never consume sub.Events()

	deadline := time.Now().Add(2 * time.Second)
	for time.Now().Before(deadline) {
		if dropped.Load() > 0 {
			break
		}
		time.Sleep(10 * time.Millisecond)
	}
	if got := dropped.Load(); got == 0 {
		t.Errorf("OnDrop fired %d times, want >= 1", got)
	}
}

func TestEventStream_StableConnectionResetsAttemptBudget(t *testing.T) {
	// Shrink the threshold so a tightly-timed test can still cross it.
	old := reconnectStableThreshold
	reconnectStableThreshold = 50 * time.Millisecond
	t.Cleanup(func() { reconnectStableThreshold = old })

	var attempts atomic.Int32
	dropAfter := make(chan struct{})
	fs := newFakeServer(t, func(s *wsSession) {
		n := attempts.Add(1)
		s.sendWelcome()
		msg := <-s.received
		s.confirm(msg["identifier"])
		// First two connections each stay alive past
		// reconnectStableThreshold, then drop. Without the time-based
		// reset the loop would exhaust at MaxAttempts=1 after the
		// second connection — only the reset lets a 3rd dial happen.
		if n <= 2 {
			time.Sleep(80 * time.Millisecond)
			s.close()
			return
		}
		<-dropAfter
		<-s.ctx.Done()
	})

	c := newStreamTestClient(fs.srv.URL)
	c.SetTokens("stub", "")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	conn, err := c.OpenEventStream(ctx, EventStreamOptions{
		Reconnect: ReconnectPolicy{
			MaxAttempts:  1,
			InitialDelay: 10 * time.Millisecond,
			MaxDelay:     20 * time.Millisecond,
		},
	})
	if err != nil {
		t.Fatalf("OpenEventStream: %v", err)
	}
	defer conn.Close()
	if _, err := conn.Subscribe(ctx, ChatRoomChannel()); err != nil {
		t.Fatalf("Subscribe: %v", err)
	}

	deadline := time.Now().Add(2 * time.Second)
	for time.Now().Before(deadline) {
		if attempts.Load() >= 3 {
			break
		}
		time.Sleep(10 * time.Millisecond)
	}
	close(dropAfter)
	if got := attempts.Load(); got < 3 {
		t.Errorf("attempts = %d, want >= 3 (stable connection should have reset the failure budget)", got)
	}
}

func TestSubscription_DoneFiresOnUnsubscribe(t *testing.T) {
	fs := newFakeServer(t, func(s *wsSession) {
		s.sendWelcome()
		msg := <-s.received
		s.confirm(msg["identifier"])
		<-s.ctx.Done()
	})

	c := newStreamTestClient(fs.srv.URL)
	c.SetTokens("stub", "")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	conn, err := c.OpenEventStream(ctx, EventStreamOptions{Reconnect: ReconnectPolicy{Disabled: true}})
	if err != nil {
		t.Fatalf("OpenEventStream: %v", err)
	}
	defer conn.Close()

	sub, err := conn.Subscribe(ctx, ChatRoomChannel())
	if err != nil {
		t.Fatalf("Subscribe: %v", err)
	}
	_ = sub.Unsubscribe(ctx)
	select {
	case <-sub.Done():
	case <-time.After(2 * time.Second):
		t.Fatal("Subscription.Done did not fire after Unsubscribe")
	}
}
