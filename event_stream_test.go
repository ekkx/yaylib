package yaylib

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"sync"
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

func (s *wsSession) reject(identifier string) {
	body, _ := json.Marshal(map[string]string{"identifier": identifier, "type": "reject_subscription"})
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

func TestEventStream_SubscribeAndReceiveEvent(t *testing.T) {
	var sess *wsSession
	ready := make(chan struct{})
	fs := newFakeServer(t, func(s *wsSession) {
		sess = s
		s.sendWelcome()
		close(ready)
		// Wait for subscribe, confirm, push one event.
		msg := <-s.received
		if msg["command"] != "subscribe" {
			t.Errorf("first command = %q, want subscribe", msg["command"])
		}
		s.confirm(msg["identifier"])
		s.pushEvent(msg["identifier"], "chat_deleted", map[string]int64{"room_id": 123})
		<-s.ctx.Done()
	})

	c := newStreamTestClient(fs.srv.URL)
	c.Tokens.Access = "stub" // skip auth-refresh path
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	conn, err := c.OpenEventStream(ctx, EventStreamOptions{Reconnect: ReconnectPolicy{Disabled: true}})
	if err != nil {
		t.Fatalf("OpenEventStream: %v", err)
	}
	defer conn.Close()
	<-ready

	sub, err := conn.Subscribe(ctx, ChatRoomChannel())
	if err != nil {
		t.Fatalf("Subscribe: %v", err)
	}

	select {
	case ev := <-sub.Events():
		cd, ok := ev.(*ChatDeletedEvent)
		if !ok {
			t.Fatalf("event = %T, want *ChatDeletedEvent", ev)
		}
		if cd.RoomID != 123 {
			t.Errorf("RoomID = %d, want 123", cd.RoomID)
		}
	case <-time.After(2 * time.Second):
		t.Fatal("no event received")
	}
	_ = sess
}

func TestEventStream_RejectedSubscription(t *testing.T) {
	fs := newFakeServer(t, func(s *wsSession) {
		s.sendWelcome()
		msg := <-s.received
		s.reject(msg["identifier"])
		<-s.ctx.Done()
	})

	c := newStreamTestClient(fs.srv.URL)
	c.Tokens.Access = "stub"
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	conn, err := c.OpenEventStream(ctx, EventStreamOptions{Reconnect: ReconnectPolicy{Disabled: true}})
	if err != nil {
		t.Fatalf("OpenEventStream: %v", err)
	}
	defer conn.Close()

	if _, err := conn.Subscribe(ctx, ChatRoomChannel()); err == nil {
		t.Error("Subscribe: expected rejection error, got nil")
	}
}

func TestEventStream_MultipleChannels(t *testing.T) {
	fs := newFakeServer(t, func(s *wsSession) {
		s.sendWelcome()
		// Confirm whatever subscriptions arrive.
		for i := 0; i < 2; i++ {
			msg := <-s.received
			s.confirm(msg["identifier"])
		}
		// Push events on each channel.
		s.pushEvent(`{"channel":"ChatRoomChannel"}`, "total_chat_request",
			map[string]int{"total_count": 7})
		s.pushEvent(`{"channel":"GroupUpdatesChannel"}`, "new_post",
			map[string]int64{"group_id": 42})
		<-s.ctx.Done()
	})

	c := newStreamTestClient(fs.srv.URL)
	c.Tokens.Access = "stub"
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	conn, err := c.OpenEventStream(ctx, EventStreamOptions{Reconnect: ReconnectPolicy{Disabled: true}})
	if err != nil {
		t.Fatalf("OpenEventStream: %v", err)
	}
	defer conn.Close()

	sub1, err := conn.Subscribe(ctx, ChatRoomChannel())
	if err != nil {
		t.Fatalf("Subscribe ChatRoom: %v", err)
	}
	sub2, err := conn.Subscribe(ctx, GroupUpdatesChannel())
	if err != nil {
		t.Fatalf("Subscribe GroupUpdates: %v", err)
	}

	gotChat := false
	gotGroup := false
	deadline := time.After(3 * time.Second)
	for !(gotChat && gotGroup) {
		select {
		case ev := <-sub1.Events():
			if e, ok := ev.(*TotalChatRequestEvent); ok && e.TotalCount == 7 {
				gotChat = true
			}
		case ev := <-sub2.Events():
			if e, ok := ev.(*GroupUpdatedEvent); ok && e.GroupID == 42 {
				gotGroup = true
			}
		case <-deadline:
			t.Fatalf("missing events: chat=%v group=%v", gotChat, gotGroup)
		}
	}
}

func TestEventStream_ReconnectAfterServerClose(t *testing.T) {
	var attempts atomic.Int32
	var subOnSecond sync.WaitGroup
	subOnSecond.Add(1)
	fs := newFakeServer(t, func(s *wsSession) {
		n := attempts.Add(1)
		s.sendWelcome()
		msg := <-s.received
		if msg["command"] != "subscribe" {
			t.Errorf("attempt %d: first command = %q", n, msg["command"])
		}
		s.confirm(msg["identifier"])
		if n == 1 {
			// First connection: drop right after confirm so the client
			// retries.
			s.close()
			return
		}
		// Second connection: re-subscribe should land here. Push an
		// event so the client can prove it survived.
		s.pushEvent(msg["identifier"], "chat_deleted", map[string]int64{"room_id": 999})
		subOnSecond.Done()
		<-s.ctx.Done()
	})

	c := newStreamTestClient(fs.srv.URL)
	c.Tokens.Access = "stub"
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	conn, err := c.OpenEventStream(ctx, EventStreamOptions{
		Reconnect: ReconnectPolicy{
			InitialDelay: 50 * time.Millisecond,
			MaxDelay:     200 * time.Millisecond,
		},
	})
	if err != nil {
		t.Fatalf("OpenEventStream: %v", err)
	}
	defer conn.Close()

	sub, err := conn.Subscribe(ctx, ChatRoomChannel())
	if err != nil {
		t.Fatalf("Subscribe: %v", err)
	}

	subOnSecond.Wait()

	select {
	case ev := <-sub.Events():
		cd, ok := ev.(*ChatDeletedEvent)
		if !ok || cd.RoomID != 999 {
			t.Fatalf("post-reconnect event = %#v", ev)
		}
	case <-time.After(3 * time.Second):
		t.Fatal("no event after reconnect")
	}
	if got := attempts.Load(); got < 2 {
		t.Errorf("server saw %d connections, want >= 2", got)
	}
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
	c.Tokens.Access = "stub"
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

func TestEventStream_SubscribeTimeout(t *testing.T) {
	fs := newFakeServer(t, func(s *wsSession) {
		s.sendWelcome()
		// Read incoming subscribe but never confirm.
		<-s.received
		<-s.ctx.Done()
	})

	c := newStreamTestClient(fs.srv.URL)
	c.Tokens.Access = "stub"
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	conn, err := c.OpenEventStream(ctx, EventStreamOptions{
		Reconnect:        ReconnectPolicy{Disabled: true},
		SubscribeTimeout: 200 * time.Millisecond,
	})
	if err != nil {
		t.Fatalf("OpenEventStream: %v", err)
	}
	defer conn.Close()

	if _, err := conn.Subscribe(ctx, ChatRoomChannel()); err == nil {
		t.Error("expected subscribe timeout, got nil")
	}
}

func TestEventStream_DoneAndErrOnCleanClose(t *testing.T) {
	fs := newFakeServer(t, func(s *wsSession) {
		s.sendWelcome()
		<-s.ctx.Done()
	})

	c := newStreamTestClient(fs.srv.URL)
	c.Tokens.Access = "stub"
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	conn, err := c.OpenEventStream(ctx, EventStreamOptions{Reconnect: ReconnectPolicy{Disabled: true}})
	if err != nil {
		t.Fatalf("OpenEventStream: %v", err)
	}
	if err := conn.Err(); err != nil {
		t.Errorf("Err before Done = %v, want nil", err)
	}
	_ = conn.Close()
	select {
	case <-conn.Done():
	case <-time.After(2 * time.Second):
		t.Fatal("Done did not fire after Close")
	}
	if err := conn.Err(); err != nil {
		t.Errorf("Err after clean Close = %v, want nil", err)
	}
}

func TestEventStream_ErrAfterReconnectExhausted(t *testing.T) {
	fs := newFakeServer(t, func(s *wsSession) {
		s.sendWelcome()
		// Drop immediately so the client tries to reconnect.
		s.close()
	})

	c := newStreamTestClient(fs.srv.URL)
	c.Tokens.Access = "stub"
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

func TestSubscription_DoneFiresOnUnsubscribe(t *testing.T) {
	fs := newFakeServer(t, func(s *wsSession) {
		s.sendWelcome()
		msg := <-s.received
		s.confirm(msg["identifier"])
		<-s.ctx.Done()
	})

	c := newStreamTestClient(fs.srv.URL)
	c.Tokens.Access = "stub"
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
