package mock

import (
	"context"
	"encoding/json"
	"net/http"
	"regexp"
	"sync"

	"github.com/coder/websocket"
)

// The event stream speaks an ActionCable-style JSON protocol:
//
//	server -> {"type":"welcome","sid":...}
//	client -> {"command":"subscribe","identifier":"<json>"}
//	server -> {"identifier":"<json>","type":"confirm_subscription"}
//	server -> {"identifier":"<json>","message":{"event":"<n>","data":{}}}
//	client -> {"command":"unsubscribe","identifier":"<json>"}
//
// The connection is controlled per-dial by the `mock` query parameter
// (the SDK lets a test set the event-stream base URL, so a query
// string is the portable control channel — WS dials carry no custom
// headers):
//
//	(default)            confirm every subscribe, then push one
//	                     representative event for that channel
//	mock=reject          reject every subscribe
//	mock=no-confirm      ignore subscribes (exercises client timeout)
//	mock=drop-after-confirm
//	                     confirm + push, then close — the SDK
//	                     reconnects to the same URL and the next
//	                     connection confirms + pushes again, so
//	                     reconnect/re-subscribe is observable without
//	                     out-of-band orchestration
//
// A WS dial MUST NOT carry Authorization (the token rides the query),
// and the handler refuses the upgrade if it does, so all three
// languages get the same contract check.

// wsHub tracks live connections so the /__mock/ws/close admin route
// can drop them (used by multi-subscription reconnect tests).
type wsHub struct {
	mu    sync.Mutex
	conns map[*websocket.Conn]struct{}
}

func newWSHub() *wsHub { return &wsHub{conns: map[*websocket.Conn]struct{}{}} }

func (h *wsHub) add(c *websocket.Conn) {
	h.mu.Lock()
	h.conns[c] = struct{}{}
	h.mu.Unlock()
}

func (h *wsHub) remove(c *websocket.Conn) {
	h.mu.Lock()
	delete(h.conns, c)
	h.mu.Unlock()
}

func (h *wsHub) closeAll() {
	h.mu.Lock()
	conns := make([]*websocket.Conn, 0, len(h.conns))
	for c := range h.conns {
		conns = append(conns, c)
	}
	h.conns = map[*websocket.Conn]struct{}{}
	h.mu.Unlock()
	for _, c := range conns {
		_ = c.Close(websocket.StatusGoingAway, "mock close")
	}
}

func isWebSocketUpgrade(r *http.Request) bool {
	return headerContainsToken(r.Header.Get("Connection"), "upgrade") &&
		headerContainsToken(r.Header.Get("Upgrade"), "websocket")
}

func headerContainsToken(v, token string) bool {
	for _, part := range regexp.MustCompile(`[,\s]+`).Split(v, -1) {
		if part != "" && equalFold(part, token) {
			return true
		}
	}
	return false
}

func equalFold(a, b string) bool {
	if len(a) != len(b) {
		return false
	}
	for i := 0; i < len(a); i++ {
		ca, cb := a[i], b[i]
		if 'A' <= ca && ca <= 'Z' {
			ca += 'a' - 'A'
		}
		if 'A' <= cb && cb <= 'Z' {
			cb += 'a' - 'A'
		}
		if ca != cb {
			return false
		}
	}
	return true
}

var channelNameRe = regexp.MustCompile(`"channel"\s*:\s*"([A-Za-z]+)"`)

// representativeEvent returns the (event, data) pair pushed after a
// successful subscribe, chosen by the channel named in the identifier
// so each channel type yields a type-appropriate event.
func representativeEvent(identifier string) (string, any) {
	channel := ""
	if m := channelNameRe.FindStringSubmatch(identifier); m != nil {
		channel = m[1]
	}
	switch channel {
	case "ChatRoomChannel":
		return "chat_deleted", map[string]int64{"room_id": 123}
	case "MessagesChannel":
		return "new_message", map[string]int64{"id": 1}
	case "GroupUpdatesChannel", "GroupPostsChannel":
		return "new_post", map[string]int64{"group_id": 1}
	default:
		return "chat_deleted", map[string]int64{"room_id": 123}
	}
}

func (s *Server) serveWS(w http.ResponseWriter, r *http.Request) {
	if r.Header.Get("Authorization") != "" {
		// Contract: the event-stream dial authenticates via the
		// token query, never a bearer header.
		http.Error(w, "websocket dial must not carry Authorization", http.StatusUnauthorized)
		return
	}
	mode := r.URL.Query().Get("mock")

	c, err := websocket.Accept(w, r, &websocket.AcceptOptions{InsecureSkipVerify: true})
	if err != nil {
		return
	}
	s.ws.add(c)
	defer func() {
		s.ws.remove(c)
		_ = c.Close(websocket.StatusNormalClosure, "")
	}()

	ctx := r.Context()
	if err := writeFrame(ctx, c, map[string]string{"type": "welcome", "sid": "mock-sid"}); err != nil {
		return
	}

	for {
		_, data, err := c.Read(ctx)
		if err != nil {
			return
		}
		var cmd struct {
			Command    string `json:"command"`
			Identifier string `json:"identifier"`
		}
		if json.Unmarshal(data, &cmd) != nil {
			continue
		}
		switch cmd.Command {
		case "subscribe":
			if !s.handleSubscribe(ctx, c, cmd.Identifier, mode) {
				return // drop-after-confirm closed the connection
			}
		case "unsubscribe":
			// Acknowledged implicitly; nothing to tear down server-side.
		}
	}
}

// handleSubscribe applies the dial mode to one subscribe command.
// It returns false when the connection should be closed afterwards.
func (s *Server) handleSubscribe(ctx context.Context, c *websocket.Conn, ident, mode string) bool {
	switch mode {
	case "reject":
		_ = writeFrame(ctx, c, map[string]string{"identifier": ident, "type": "reject_subscription"})
		return true
	case "no-confirm":
		return true // deliberately silent
	}
	if writeFrame(ctx, c, map[string]string{"identifier": ident, "type": "confirm_subscription"}) != nil {
		return false
	}
	event, payload := representativeEvent(ident)
	if writeFrame(ctx, c, map[string]any{
		"identifier": ident,
		"message":    map[string]any{"event": event, "data": payload},
	}) != nil {
		return false
	}
	return mode != "drop-after-confirm"
}

func writeFrame(ctx context.Context, c *websocket.Conn, v any) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}
	return c.Write(ctx, websocket.MessageText, b)
}
