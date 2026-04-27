package yaylib

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"math/rand"
	"net/url"
	"sync"
	"sync/atomic"
	"time"

	"github.com/coder/websocket"
)

// DefaultEventStreamURL is the production event-stream endpoint.
// Override per-Client with WithEventStreamURL when targeting a
// different environment. The URL is a wss:// because the underlying
// transport is WebSocket — useful to know when reading proxy logs or
// tcpdump output.
const DefaultEventStreamURL = "wss://cable.yay.space"

// Default subscription tunables. Adjust per-Conn via EventStreamOptions.
const (
	defaultSubscribeTimeout = 10 * time.Second
	defaultEventBuffer      = 64
)

// ReconnectPolicy controls how a *EventStream behaves after the
// underlying connection drops. The zero value (DefaultReconnectPolicy())
// keeps reconnecting forever with exponential backoff; set Disabled true
// to surface the disconnect to the caller instead.
type ReconnectPolicy struct {
	// Disabled stops the conn after the first disconnect. Subscriptions
	// terminate with the connection error.
	Disabled bool

	// MaxAttempts caps the total reconnect attempts (excluding the
	// initial dial). 0 means unlimited.
	MaxAttempts int

	// InitialDelay is the first backoff sleep. Each subsequent attempt
	// doubles up to MaxDelay, with full jitter.
	InitialDelay time.Duration

	// MaxDelay caps any single sleep between attempts.
	MaxDelay time.Duration
}

// DefaultReconnectPolicy returns the policy OpenEventStream applies when
// the caller does not override it: unlimited attempts, 500ms initial
// delay, 30s ceiling.
func DefaultReconnectPolicy() ReconnectPolicy {
	return ReconnectPolicy{
		InitialDelay: 500 * time.Millisecond,
		MaxDelay:     30 * time.Second,
	}
}

// EventStreamOptions tunes a single OpenEventStream invocation.
type EventStreamOptions struct {
	// Reconnect controls auto-reconnect behavior. Defaults to
	// DefaultReconnectPolicy() if left zero (use Disabled=true to opt out).
	Reconnect ReconnectPolicy

	// EventBuffer is the per-Subscription channel buffer. Slow consumers
	// past this point cause events to be dropped silently — choose a
	// value large enough for your traffic. Defaults to 64.
	EventBuffer int

	// SubscribeTimeout caps how long Subscribe blocks waiting for the
	// server's confirm_subscription frame. Defaults to 10s.
	SubscribeTimeout time.Duration
}

// Channel identifies one subscribable event-stream channel. Use the
// constructors (ChatRoomChannel, MessagesChannel, GroupUpdatesChannel,
// GroupPostsChannel) rather than implementing this yourself — the
// identifier string format is wire-significant.
type Channel interface {
	Identifier() string
}

type chatRoomChannel struct{}

func (chatRoomChannel) Identifier() string { return `{"channel":"ChatRoomChannel"}` }

// ChatRoomChannel subscribes to global chat-room signals: new chat
// requests, chat deletions, and forced unsubscribe notices.
func ChatRoomChannel() Channel { return chatRoomChannel{} }

type messagesChannel struct{ roomID int64 }

func (m messagesChannel) Identifier() string {
	return fmt.Sprintf(`{"channel":"MessagesChannel", "chat_room_id":"%d"}`, m.roomID)
}

// MessagesChannel subscribes to messages and video-processed events for
// a single chat room.
func MessagesChannel(roomID int64) Channel { return messagesChannel{roomID: roomID} }

type groupUpdatesChannel struct{}

func (groupUpdatesChannel) Identifier() string { return `{"channel":"GroupUpdatesChannel"}` }

// GroupUpdatesChannel subscribes to group-membership / group-state
// updates across all groups the user belongs to.
func GroupUpdatesChannel() Channel { return groupUpdatesChannel{} }

type groupPostsChannel struct{ groupID int64 }

func (g groupPostsChannel) Identifier() string {
	return fmt.Sprintf(`{"channel":"GroupPostsChannel", "group_id":"%d"}`, g.groupID)
}

// GroupPostsChannel subscribes to new posts within a single group's
// timeline.
func GroupPostsChannel(groupID int64) Channel { return groupPostsChannel{groupID: groupID} }

// EventStream is a multiplexed real-time event channel from the Yay!
// server. It maintains a single underlying connection (a WebSocket
// against wss://cable.yay.space) and dispatches server-pushed events
// to per-channel Subscriptions.
//
// EventStream is safe for concurrent use. When the underlying socket
// drops the stream auto-reconnects and re-subscribes per
// ReconnectPolicy; subscribers continue to read from the same Events()
// channels with no interruption visible past the dropped frames.
type EventStream struct {
	client *Client
	base   string
	opts   EventStreamOptions

	wsMu sync.RWMutex
	ws   *websocket.Conn

	subsMu sync.Mutex
	subs   map[string]*Subscription

	ctx    context.Context
	cancel context.CancelFunc

	closed atomic.Bool
	doneCh chan struct{}

	errMu    sync.Mutex
	finalErr error
}

// Subscription is one active channel subscription. Read events from
// Events(); the channel closes when the subscription terminates (either
// via Unsubscribe or because the parent EventStream closed). Done() is
// available for callers that want to wait on termination without
// consuming events themselves — the parent EventStream's Err() carries
// the cause when the subscription died because the stream died.
type Subscription struct {
	conn  *EventStream
	ident string

	events chan Event

	confirmOnce sync.Once
	confirmCh   chan struct{}
	confirmErr  error

	// sendMu pairs every send on `events` with the close in
	// Unsubscribe / terminate. Without it, a dispatch goroutine that
	// has already obtained the *Subscription from findSub races with
	// Unsubscribe and panics with "send on closed channel".
	sendMu sync.Mutex
	closed atomic.Bool
	done   chan struct{}
}

// Events returns the buffered channel of incoming events. The channel
// closes when the subscription terminates.
func (s *Subscription) Events() <-chan Event { return s.events }

// Done is closed when the subscription has terminated, either through
// Unsubscribe or because the parent EventStream shut down. Useful when
// a different goroutine is consuming Events() and you just want to
// block on termination.
func (s *Subscription) Done() <-chan struct{} { return s.done }

// Unsubscribe sends an unsubscribe command and terminates the
// subscription. It is safe to call multiple times.
func (s *Subscription) Unsubscribe(ctx context.Context) error {
	if !s.closed.CompareAndSwap(false, true) {
		return nil
	}
	s.conn.removeSub(s.ident)
	// Best effort: server may already be gone.
	_ = s.conn.writeCommand(ctx, "unsubscribe", s.ident)
	s.closeChannels()
	return nil
}

func (s *Subscription) terminate(err error) {
	if !s.closed.CompareAndSwap(false, true) {
		return
	}
	s.confirmOnce.Do(func() {
		s.confirmErr = err
		close(s.confirmCh)
	})
	s.closeChannels()
}

// closeChannels closes events and done under sendMu so an in-flight
// deliver() can observe `closed == true` and bail before it would
// otherwise send on a closed channel.
func (s *Subscription) closeChannels() {
	s.sendMu.Lock()
	close(s.events)
	close(s.done)
	s.sendMu.Unlock()
}

// deliver is the only safe path that writes to s.events. The lock
// pairs with closeChannels; the closed check inside the lock makes the
// "still open at lock time" guarantee that the subsequent send relies on.
func (s *Subscription) deliver(ev Event) {
	s.sendMu.Lock()
	defer s.sendMu.Unlock()
	if s.closed.Load() {
		return
	}
	select {
	case s.events <- ev:
	default:
		// Buffer full — drop. Caller is responsible for sizing
		// EventBuffer big enough for their traffic.
	}
}

// OpenEventStream opens a real-time event channel against the
// configured Yay! event-stream endpoint and returns a stream ready for
// Subscribe. The provided ctx governs the initial connect only; the
// resulting EventStream lives until Close (or a permanent reconnect
// failure). It requires an authenticated Client — make sure the access
// token is set (e.g., via LoginWithEmail) before calling.
//
// Internally the stream is a WebSocket; that's an implementation
// detail, but it does mean http(s) proxy / firewall configuration that
// blocks WebSocket upgrades will block this call.
func (c *Client) OpenEventStream(ctx context.Context, opts ...EventStreamOptions) (*EventStream, error) {
	o := EventStreamOptions{}
	if len(opts) > 0 {
		o = opts[0]
	}
	if o.Reconnect.InitialDelay <= 0 {
		o.Reconnect.InitialDelay = 500 * time.Millisecond
	}
	if o.Reconnect.MaxDelay <= 0 {
		o.Reconnect.MaxDelay = 30 * time.Second
	}
	if o.EventBuffer <= 0 {
		o.EventBuffer = defaultEventBuffer
	}
	if o.SubscribeTimeout <= 0 {
		o.SubscribeTimeout = defaultSubscribeTimeout
	}

	connCtx, cancel := context.WithCancel(context.Background())
	conn := &EventStream{
		client: c,
		base:   c.EventStreamURL,
		opts:   o,
		subs:   make(map[string]*Subscription),
		ctx:    connCtx,
		cancel: cancel,
		doneCh: make(chan struct{}),
	}
	if conn.base == "" {
		conn.base = DefaultEventStreamURL
	}

	if err := conn.connect(ctx); err != nil {
		cancel()
		return nil, err
	}

	go conn.runLoop()
	return conn, nil
}

// connect performs one dial + welcome handshake, populating conn.ws on
// success. It does not start the read loop — runLoop owns that.
func (conn *EventStream) connect(ctx context.Context) error {
	tokenResp, _, err := conn.client.GetWebSocketToken(ctx).Execute()
	if err != nil {
		return fmt.Errorf("fetch ws token: %w", err)
	}
	token := tokenResp.GetToken()

	u, err := url.Parse(conn.base)
	if err != nil {
		return fmt.Errorf("parse ws url: %w", err)
	}
	q := u.Query()
	q.Set("token", token)
	q.Set("app_version", conn.client.APIVersionName)
	u.RawQuery = q.Encode()

	ws, _, err := websocket.Dial(ctx, u.String(), &websocket.DialOptions{
		HTTPClient: conn.client.wsHTTPClient(),
	})
	if err != nil {
		return fmt.Errorf("dial: %w", err)
	}
	// Server-pushed messages can be larger than coder/websocket's tiny
	// default; chat messages with attachments easily exceed 32KB.
	ws.SetReadLimit(1 << 20)

	// Welcome handshake. The server sends {"type":"welcome"} immediately;
	// drain frames until we see it (any preceding pings are ignored).
	for {
		_, data, err := ws.Read(ctx)
		if err != nil {
			ws.Close(websocket.StatusInternalError, "welcome read failed")
			return fmt.Errorf("read welcome: %w", err)
		}
		var f wireFrame
		if err := json.Unmarshal(data, &f); err != nil {
			continue
		}
		if f.Type == "welcome" {
			break
		}
		if f.Type == "disconnect" {
			ws.Close(websocket.StatusNormalClosure, "")
			return fmt.Errorf("server disconnect during welcome")
		}
	}

	conn.wsMu.Lock()
	old := conn.ws
	conn.ws = ws
	conn.wsMu.Unlock()
	if old != nil {
		_ = old.Close(websocket.StatusNormalClosure, "")
	}
	return nil
}

// wireFrame is the ActionCable channel envelope. `message` carries
// per-event payloads when present; it stays empty for control frames
// like welcome / ping / confirm_subscription.
type wireFrame struct {
	Identifier string          `json:"identifier,omitempty"`
	Type       string          `json:"type,omitempty"`
	Message    json.RawMessage `json:"message,omitempty"`
	Sid        string          `json:"sid,omitempty"`
}

// eventPayload mirrors the inner shape the server packs into `message`
// for actual events. The deserializer accepts either `data` or
// `message` as the payload key.
type eventPayload struct {
	Event   string          `json:"event"`
	Data    json.RawMessage `json:"data"`
	Message json.RawMessage `json:"message"`
}

// runLoop owns the read pump for the active socket and the reconnect
// state machine. It exits only after the conn is closed (cleanly via
// Close, or after exhausting reconnect attempts).
func (conn *EventStream) runLoop() {
	defer close(conn.doneCh)

	conn.readUntilDisconnect()

	if conn.opts.Reconnect.Disabled || conn.closed.Load() {
		conn.terminateAllSubs(conn.getFinalErr())
		return
	}

	attempt := 0
	for {
		if conn.closed.Load() {
			return
		}
		attempt++
		if max := conn.opts.Reconnect.MaxAttempts; max > 0 && attempt > max {
			err := fmt.Errorf("reconnect: exhausted after %d attempts: %w",
				max, conn.getFinalErr())
			conn.setFinalErr(err)
			conn.terminateAllSubs(err)
			return
		}

		delay := conn.backoff(attempt)
		if !sleepWithCtx(conn.ctx, delay) {
			conn.setFinalErr(conn.ctx.Err())
			conn.terminateAllSubs(conn.ctx.Err())
			return
		}

		dialCtx, cancel := context.WithTimeout(conn.ctx, 30*time.Second)
		err := conn.connect(dialCtx)
		cancel()
		if err != nil {
			conn.setFinalErr(err)
			continue
		}

		// Re-subscribe everything we knew about before. Failures here
		// surface to per-sub Err but don't abort the whole conn.
		conn.resubscribeAll()
		if conn.readUntilDisconnect() {
			// Got at least one post-welcome frame, treat as stable —
			// reset the failure budget so a future blip starts fresh.
			attempt = 0
		}

		if conn.closed.Load() {
			return
		}
	}
}

func (conn *EventStream) backoff(attempt int) time.Duration {
	pol := conn.opts.Reconnect
	exp := pol.InitialDelay << uint(attempt-1)
	if exp <= 0 || exp > pol.MaxDelay {
		exp = pol.MaxDelay
	}
	if exp <= 0 {
		return 0
	}
	delay := time.Duration(rand.Int63n(int64(exp)) + int64(pol.InitialDelay))
	// Final clamp: rand(exp) + InitialDelay can exceed MaxDelay when
	// InitialDelay is non-trivial, but MaxDelay is the documented ceiling.
	if delay > pol.MaxDelay {
		delay = pol.MaxDelay
	}
	return delay
}

// readUntilDisconnect reads frames from the active socket and
// dispatches them. It returns when the socket closes or errors,
// reporting whether at least one post-welcome frame arrived (the
// caller uses that as a proxy for "the connection was actually
// healthy"). Any error is recorded as the conn's final error.
func (conn *EventStream) readUntilDisconnect() (gotFrame bool) {
	conn.wsMu.RLock()
	ws := conn.ws
	conn.wsMu.RUnlock()
	if ws == nil {
		return
	}
	for {
		_, data, err := ws.Read(conn.ctx)
		if err != nil {
			conn.setFinalErr(err)
			return
		}
		gotFrame = true
		conn.dispatch(data)
	}
}

func (conn *EventStream) dispatch(data []byte) {
	var f wireFrame
	if err := json.Unmarshal(data, &f); err != nil {
		return
	}
	switch f.Type {
	case "welcome":
		// Reconnect path. The runLoop already handled the initial welcome.
		return
	case "ping":
		return
	case "disconnect":
		// Server-requested disconnect. Bubble up via a Close on the ws so
		// the read loop exits and the reconnect logic kicks in.
		conn.wsMu.RLock()
		ws := conn.ws
		conn.wsMu.RUnlock()
		if ws != nil {
			ws.Close(websocket.StatusNormalClosure, "server disconnect")
		}
		return
	case "confirm_subscription":
		if sub := conn.findSub(f.Identifier); sub != nil {
			sub.confirmOnce.Do(func() { close(sub.confirmCh) })
		}
		return
	case "reject_subscription":
		if sub := conn.findSub(f.Identifier); sub != nil {
			sub.confirmOnce.Do(func() {
				sub.confirmErr = fmt.Errorf("subscription rejected")
				close(sub.confirmCh)
			})
			sub.terminate(fmt.Errorf("subscription rejected"))
		}
		return
	}

	// Anything else is an event — dispatch by identifier.
	if len(f.Message) == 0 {
		return
	}
	var p eventPayload
	if err := json.Unmarshal(f.Message, &p); err != nil {
		return
	}
	body := p.Data
	if len(body) == 0 {
		body = p.Message
	}
	if p.Event == "" || len(body) == 0 {
		return
	}
	ev, err := decodeEvent(p.Event, body)
	if err != nil || ev == nil {
		return
	}
	if sub := conn.findSub(f.Identifier); sub != nil {
		sub.deliver(ev)
	}
}

func (conn *EventStream) findSub(ident string) *Subscription {
	conn.subsMu.Lock()
	defer conn.subsMu.Unlock()
	return conn.subs[ident]
}

func (conn *EventStream) removeSub(ident string) {
	conn.subsMu.Lock()
	delete(conn.subs, ident)
	conn.subsMu.Unlock()
}

func (conn *EventStream) resubscribeAll() {
	conn.subsMu.Lock()
	subs := make([]*Subscription, 0, len(conn.subs))
	for _, s := range conn.subs {
		subs = append(subs, s)
	}
	conn.subsMu.Unlock()
	for _, s := range subs {
		_ = conn.writeCommand(conn.ctx, "subscribe", s.ident)
	}
}

func (conn *EventStream) terminateAllSubs(err error) {
	conn.subsMu.Lock()
	subs := make([]*Subscription, 0, len(conn.subs))
	for _, s := range conn.subs {
		subs = append(subs, s)
	}
	conn.subs = map[string]*Subscription{}
	conn.subsMu.Unlock()
	for _, s := range subs {
		s.terminate(err)
	}
}

func (conn *EventStream) setFinalErr(err error) {
	conn.errMu.Lock()
	conn.finalErr = err
	conn.errMu.Unlock()
}

func (conn *EventStream) getFinalErr() error {
	conn.errMu.Lock()
	defer conn.errMu.Unlock()
	return conn.finalErr
}

// writeCommand sends a subscribe/unsubscribe command for an identifier.
// It returns an error only when no socket is currently available, or
// when the underlying Write fails — in either case the caller treats it
// as best-effort: a re-subscribe will run on the next reconnect.
func (conn *EventStream) writeCommand(ctx context.Context, cmd, identifier string) error {
	conn.wsMu.RLock()
	ws := conn.ws
	conn.wsMu.RUnlock()
	if ws == nil {
		return errors.New("ws not connected")
	}
	body, _ := json.Marshal(map[string]string{
		"command":    cmd,
		"identifier": identifier,
	})
	return ws.Write(ctx, websocket.MessageText, body)
}

// Subscribe registers ch on the conn and blocks until the server
// confirms (or rejects) the subscription. Returns the *Subscription
// once confirmed.
func (conn *EventStream) Subscribe(ctx context.Context, ch Channel) (*Subscription, error) {
	if conn.closed.Load() {
		return nil, errors.New("event stream closed")
	}
	ident := ch.Identifier()

	conn.subsMu.Lock()
	if existing, ok := conn.subs[ident]; ok {
		conn.subsMu.Unlock()
		return existing, nil
	}
	sub := &Subscription{
		conn:      conn,
		ident:     ident,
		events:    make(chan Event, conn.opts.EventBuffer),
		confirmCh: make(chan struct{}),
		done:      make(chan struct{}),
	}
	conn.subs[ident] = sub
	conn.subsMu.Unlock()

	_ = conn.writeCommand(ctx, "subscribe", ident)

	timeout := conn.opts.SubscribeTimeout
	t := time.NewTimer(timeout)
	defer t.Stop()

	select {
	case <-sub.confirmCh:
		if sub.confirmErr != nil {
			conn.removeSub(ident)
			return nil, sub.confirmErr
		}
		return sub, nil
	case <-ctx.Done():
		conn.removeSub(ident)
		return nil, ctx.Err()
	case <-conn.doneCh:
		conn.removeSub(ident)
		if err := conn.getFinalErr(); err != nil {
			return nil, err
		}
		return nil, errors.New("event stream closed")
	case <-t.C:
		conn.removeSub(ident)
		return nil, fmt.Errorf("subscribe %q: timed out after %v", ident, timeout)
	}
}

// Done returns a channel that is closed once the stream has fully shut
// down — either via Close, or after a permanent failure (reconnect
// exhausted, ctx canceled). Pair with Err() to inspect the cause.
func (conn *EventStream) Done() <-chan struct{} { return conn.doneCh }

// Err returns the stream's terminal error after Done fires. It is nil
// for a clean Close (caller-initiated shutdown), non-nil if the stream
// died on its own — reconnect exhausted, network gone, etc. Calling
// Err before Done has fired returns nil; you must wait on Done first
// if you want a definitive answer.
func (conn *EventStream) Err() error {
	select {
	case <-conn.doneCh:
	default:
		return nil
	}
	if conn.closed.Load() {
		// Caller-initiated Close — finalErr is just the synthetic
		// context.Canceled that flushed the read pump, not a useful
		// signal for the caller.
		return nil
	}
	return conn.getFinalErr()
}

// Close terminates the stream, all subscriptions, and the underlying
// socket. It is safe to call multiple times. Pending Send() / Write
// loops are unblocked via context cancellation.
func (conn *EventStream) Close() error {
	if !conn.closed.CompareAndSwap(false, true) {
		<-conn.doneCh
		return nil
	}
	conn.wsMu.RLock()
	ws := conn.ws
	conn.wsMu.RUnlock()
	if ws != nil {
		_ = ws.Close(websocket.StatusNormalClosure, "")
	}
	conn.cancel()
	<-conn.doneCh
	return nil
}
