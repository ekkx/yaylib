package yaylib

import (
	"context"
	"testing"
	"time"
)

// Event-stream behavior parity, driven against the shared server's
// WebSocket contract. The per-dial mode is selected by mockStreamClient
// (see the harness). The shared server pushes one representative event
// per channel: ChatRoomChannel → chat_deleted{room_id:123},
// GroupUpdatesChannel → new_post{group_id:1}.
//
// The Go-internal reconnect-budget mechanics (exhaustion, the 30s
// stable-connection reset, the per-sub OnDrop overflow) and the
// client-side unsubscribe lifecycle stay as local fixtures in
// event_stream_test.go: the shared WS contract pushes a single event
// per subscribe and exposes no socket-fault knob, so those are not
// expressible here (§15 covers reconnect/reject/timeout, which are).

func awaitEvent(t *testing.T, sub *Subscription, d time.Duration) Event {
	t.Helper()
	select {
	case ev := <-sub.Events():
		return ev
	case <-time.After(d):
		t.Fatalf("no event within %v", d)
		return nil
	}
}

// PORTING:S18
// PORTING:S25
func TestEventStream_SubscribeAndReceiveEvent(t *testing.T) {
	c := mockStreamClient(t, "")
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
	ev := awaitEvent(t, sub, 2*time.Second)
	cd, ok := ev.(*ChatDeletedEvent)
	if !ok {
		t.Fatalf("event = %T, want *ChatDeletedEvent", ev)
	}
	if cd.RoomID != 123 {
		t.Errorf("RoomID = %d, want 123 (server's representative payload)", cd.RoomID)
	}
}

// PORTING:S21
func TestEventStream_RejectedSubscription(t *testing.T) {
	c := mockStreamClient(t, "reject")
	c.SetTokens("stub", "")
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

// PORTING:S18
// PORTING:S25
func TestEventStream_MultipleChannels(t *testing.T) {
	c := mockStreamClient(t, "")
	c.SetTokens("stub", "")
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

	gotChat, gotGroup := false, false
	deadline := time.After(3 * time.Second)
	for !(gotChat && gotGroup) {
		select {
		case ev := <-sub1.Events():
			if e, ok := ev.(*ChatDeletedEvent); ok && e.RoomID == 123 {
				gotChat = true
			}
		case ev := <-sub2.Events():
			if e, ok := ev.(*GroupUpdatedEvent); ok && e.GroupID == 1 {
				gotGroup = true
			}
		case <-deadline:
			t.Fatalf("missing events: chat=%v group=%v", gotChat, gotGroup)
		}
	}
}

// PORTING:S19
func TestEventStream_ReconnectAfterServerClose(t *testing.T) {
	c := mockStreamClient(t, "drop-after-confirm")
	c.SetTokens("stub", "")
	ctx, cancel := context.WithTimeout(context.Background(), 8*time.Second)
	defer cancel()

	// drop-after-confirm closes the socket right after pushing the
	// event, so the client must reconnect and re-subscribe to keep
	// receiving. Seeing the representative event more than once proves
	// the reconnect + re-subscribe cycle.
	conn, err := c.OpenEventStream(ctx, EventStreamOptions{
		Reconnect: ReconnectPolicy{InitialDelay: 10 * time.Millisecond, MaxDelay: 30 * time.Millisecond},
	})
	if err != nil {
		t.Fatalf("OpenEventStream: %v", err)
	}
	defer conn.Close()

	sub, err := conn.Subscribe(ctx, ChatRoomChannel())
	if err != nil {
		t.Fatalf("Subscribe: %v", err)
	}
	for i := 0; i < 2; i++ {
		ev := awaitEvent(t, sub, 3*time.Second)
		if _, ok := ev.(*ChatDeletedEvent); !ok {
			t.Fatalf("cycle %d: event = %T, want *ChatDeletedEvent", i, ev)
		}
	}
}

// PORTING:S21
func TestEventStream_SubscribeTimeout(t *testing.T) {
	c := mockStreamClient(t, "no-confirm")
	c.SetTokens("stub", "")
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
	c := mockStreamClient(t, "")
	c.SetTokens("stub", "")
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

// PORTING:S20
func TestEventStream_MultipleSubsResubscribeAfterReconnect(t *testing.T) {
	c := mockStreamClient(t, "drop-after-confirm")
	c.SetTokens("stub", "")
	ctx, cancel := context.WithTimeout(context.Background(), 8*time.Second)
	defer cancel()

	conn, err := c.OpenEventStream(ctx, EventStreamOptions{
		Reconnect: ReconnectPolicy{InitialDelay: 10 * time.Millisecond, MaxDelay: 30 * time.Millisecond},
	})
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

	// Both subs must keep receiving across reconnect cycles, proving
	// every sub is re-subscribed on the new connection.
	chat, group := 0, 0
	deadline := time.After(6 * time.Second)
	for chat < 2 || group < 2 {
		select {
		case ev := <-sub1.Events():
			if _, ok := ev.(*ChatDeletedEvent); ok {
				chat++
			}
		case ev := <-sub2.Events():
			if _, ok := ev.(*GroupUpdatedEvent); ok {
				group++
			}
		case <-deadline:
			t.Fatalf("after reconnect, counts chat=%d group=%d, want ≥2 each", chat, group)
		}
	}
}

// PORTING:S22
func TestEventStream_WSDialDoesNotLeakBearer(t *testing.T) {
	c := mockStreamClient(t, "")
	// Tokens are set, but the WS dial must authenticate via the query
	// only. The shared server refuses the upgrade (401) if the dial
	// carries Authorization, so a successful open proves no bearer.
	c.SetTokens("ACCESS_THAT_MUST_NOT_LEAK", "REF")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	conn, err := c.OpenEventStream(ctx, EventStreamOptions{Reconnect: ReconnectPolicy{Disabled: true}})
	if err != nil {
		t.Fatalf("OpenEventStream failed; a leaked bearer would 401 the upgrade: %v", err)
	}
	defer conn.Close()

	if _, err := conn.Subscribe(ctx, ChatRoomChannel()); err != nil {
		t.Fatalf("Subscribe over a no-bearer dial failed: %v", err)
	}
}
