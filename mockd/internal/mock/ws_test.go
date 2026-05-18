package mock

import (
	"context"
	"encoding/json"
	"net/http"
	"strings"
	"testing"
	"time"

	"github.com/coder/websocket"
)

func wsURL(ts string, query string) string {
	u := strings.Replace(ts, "http://", "ws://", 1)
	if query != "" {
		u += "?" + query
	}
	return u
}

// dialWS opens a stream and consumes the welcome frame.
func dialWS(t *testing.T, ts, query string) (*websocket.Conn, context.Context) {
	t.Helper()
	ctx := context.Background()
	c, _, err := websocket.Dial(ctx, wsURL(ts, query), nil)
	if err != nil {
		t.Fatalf("dial: %v", err)
	}
	f := readFrame(t, ctx, c)
	if f["type"] != "welcome" {
		t.Fatalf("first frame = %v, want welcome", f)
	}
	return c, ctx
}

func readFrame(t *testing.T, ctx context.Context, c *websocket.Conn) map[string]any {
	t.Helper()
	rctx, cancel := context.WithTimeout(ctx, 2*time.Second)
	defer cancel()
	_, data, err := c.Read(rctx)
	if err != nil {
		t.Fatalf("read frame: %v", err)
	}
	var m map[string]any
	if err := json.Unmarshal(data, &m); err != nil {
		t.Fatalf("decode frame %q: %v", data, err)
	}
	return m
}

func subscribe(t *testing.T, ctx context.Context, c *websocket.Conn, ident string) {
	t.Helper()
	b, _ := json.Marshal(map[string]string{"command": "subscribe", "identifier": ident})
	if err := c.Write(ctx, websocket.MessageText, b); err != nil {
		t.Fatalf("write subscribe: %v", err)
	}
}

func TestWS_SubscribeConfirmAndEvent(t *testing.T) {
	ts := newTestServer(t)
	defer ts.Close()
	c, ctx := dialWS(t, ts.URL, "token=x")
	defer c.Close(websocket.StatusNormalClosure, "")

	const ident = `{"channel":"ChatRoomChannel"}`
	subscribe(t, ctx, c, ident)

	confirm := readFrame(t, ctx, c)
	if confirm["type"] != "confirm_subscription" || confirm["identifier"] != ident {
		t.Fatalf("confirm frame = %v", confirm)
	}
	event := readFrame(t, ctx, c)
	msg, ok := event["message"].(map[string]any)
	if !ok || msg["event"] != "chat_deleted" {
		t.Fatalf("event frame = %v, want chat_deleted message", event)
	}
}

func TestWS_RepresentativeEventPerChannel(t *testing.T) {
	ts := newTestServer(t)
	defer ts.Close()
	cases := map[string]string{
		`{"channel":"ChatRoomChannel"}`:                     "chat_deleted",
		`{"channel":"MessagesChannel", "chat_room_id":"1"}`: "new_message",
		`{"channel":"GroupUpdatesChannel"}`:                 "new_post",
		`{"channel":"GroupPostsChannel", "group_id":"2"}`:   "new_post",
	}
	for ident, want := range cases {
		c, ctx := dialWS(t, ts.URL, "token=x")
		subscribe(t, ctx, c, ident)
		readFrame(t, ctx, c) // confirm
		ev := readFrame(t, ctx, c)
		msg := ev["message"].(map[string]any)
		if msg["event"] != want {
			t.Errorf("channel %s event = %v, want %s", ident, msg["event"], want)
		}
		c.Close(websocket.StatusNormalClosure, "")
	}
}

func TestWS_RejectMode(t *testing.T) {
	ts := newTestServer(t)
	defer ts.Close()
	c, ctx := dialWS(t, ts.URL, "mock=reject&token=x")
	defer c.Close(websocket.StatusNormalClosure, "")
	subscribe(t, ctx, c, `{"channel":"ChatRoomChannel"}`)
	f := readFrame(t, ctx, c)
	if f["type"] != "reject_subscription" {
		t.Fatalf("frame = %v, want reject_subscription", f)
	}
}

func TestWS_NoConfirmModeStaysSilent(t *testing.T) {
	ts := newTestServer(t)
	defer ts.Close()
	c, ctx := dialWS(t, ts.URL, "mock=no-confirm&token=x")
	defer c.Close(websocket.StatusNormalClosure, "")
	subscribe(t, ctx, c, `{"channel":"ChatRoomChannel"}`)
	rctx, cancel := context.WithTimeout(ctx, 300*time.Millisecond)
	defer cancel()
	if _, _, err := c.Read(rctx); err == nil {
		t.Fatal("expected no frame (timeout), got one")
	}
}

func TestWS_DropAfterConfirmClosesConnection(t *testing.T) {
	ts := newTestServer(t)
	defer ts.Close()
	c, ctx := dialWS(t, ts.URL, "mock=drop-after-confirm&token=x")
	defer c.Close(websocket.StatusNormalClosure, "")
	subscribe(t, ctx, c, `{"channel":"ChatRoomChannel"}`)
	readFrame(t, ctx, c) // confirm
	readFrame(t, ctx, c) // event
	rctx, cancel := context.WithTimeout(ctx, 2*time.Second)
	defer cancel()
	if _, _, err := c.Read(rctx); err == nil {
		t.Fatal("expected connection close after confirm+event")
	}
}

func TestWS_AdminCloseDropsLiveConnections(t *testing.T) {
	ts := newTestServer(t)
	defer ts.Close()
	c, ctx := dialWS(t, ts.URL, "token=x")
	defer c.Close(websocket.StatusNormalClosure, "")
	subscribe(t, ctx, c, `{"channel":"ChatRoomChannel"}`)
	readFrame(t, ctx, c) // confirm
	readFrame(t, ctx, c) // event

	resp, err := ts.Client().Post(ts.URL+"/__mock/ws/close", "", nil)
	if err != nil {
		t.Fatalf("admin close: %v", err)
	}
	resp.Body.Close()

	rctx, cancel := context.WithTimeout(ctx, 2*time.Second)
	defer cancel()
	if _, _, err := c.Read(rctx); err == nil {
		t.Fatal("connection should be closed by admin route")
	}
}

func TestWS_DialMustNotCarryAuthorization(t *testing.T) {
	ts := newTestServer(t)
	defer ts.Close()
	_, _, err := websocket.Dial(context.Background(), wsURL(ts.URL, "token=x"), &websocket.DialOptions{
		HTTPHeader: http.Header{"Authorization": {"Bearer leaked"}},
	})
	if err == nil {
		t.Fatal("dial with Authorization should be refused")
	}
}
