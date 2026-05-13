// Package yaylib is a Go client for the Yay! API.
//
// Usage:
//
//	store, _ := yaylib.NewSessionStore("sessions.json")
//	client := yaylib.NewClient(yaylib.WithSessionStore(store))
//
//	resp, _, err := client.LoginWithEmail(ctx).
//	    Email(email).
//	    Password(password).
//	    Execute()
//
// Every Yay! operation is reachable directly on *Client — for example
// LoginWithEmail, GetRecommendedTimeline, CreatePost, BlockUser.
//
// The client also takes care of:
//
//   - the HTTP headers Yay! servers require (User-Agent, X-App-Version,
//     X-Device-Info, X-Device-UUID, X-Timestamp, X-Connection-*,
//     Accept-Language) and the Authorization header once credentials have
//     been set,
//   - an optional session cache so you don't re-login on every run — Yay!
//     rate-limits the login endpoint. Use NewSessionStore for a JSON-file
//     backend or NewMemoryStore for an in-process one,
//   - automatic 401 → refresh-token retry, with the refreshed access token
//     persisted back to the session cache,
//   - a helper that mints the short-lived HS256 `X-Jwt` token some write
//     endpoints require (see Client.GenerateXJwt),
//   - typed error inspection: yaylib.CodeOf(err) returns an ErrorCode and
//     yaylib.ErrorResponseOf(err) returns the parsed *ErrorResponse,
//   - real-time server-pushed events via Client.OpenEventStream — chat
//     messages, group updates, video-processed signals, etc., delivered
//     through Subscriptions over a multiplexed connection.
//
// `APIKey`, `DeviceUUID`, and related metadata are exposed as fields so
// callers can populate request bodies that need them. (Body builders
// replace the entire body, so these fields can't be injected transparently.
// LoginWithEmail is the exception: its wrapper auto-fills APIKey / Uuid
// when the caller leaves them blank.)
package yaylib

import (
	"context"
	"fmt"
	"net/http"
	"sync"
	"time"

	"github.com/ekkx/yaylib/v2/gen"
)

// Default values baked in for the public Yay! production environment.
// Override any of them via the With* options on NewClient.
const (
	DefaultBaseURL        = "https://api.yay.space"
	DefaultAPIKey         = "ccd59ee269c01511ba763467045c115779fcae3050238a252f1bd1a4b65cfec6"
	DefaultAPIVersionKey  = "34c8c1cdf29b46a492e8ec58e6db37ec" // HS256 signing key for X-Jwt
	DefaultAPIVersionName = "4.26"
	DefaultAppVersion     = "4.26.1"

	DefaultDeviceType    = "android"
	DefaultDeviceOS      = "11"
	DefaultDeviceDensity = "3.5"
	DefaultDeviceScreen  = "1440x2960"
	DefaultDeviceModel   = "Galaxy S9"

	DefaultConnectionType  = "wifi"
	DefaultConnectionSpeed = "0 kbps"
	DefaultAcceptLanguage  = "ja"

	DefaultHTTPTimeout = 90 * time.Second
)

// Client is a high-level Yay! API client. Construct with NewClient(...).
//
// API surface contract:
//
// Every Yay! operation is reachable as a top-level method on *Client —
// LoginWithEmail, GetRecommendedTimeline, CreatePost, BlockUser, and
// the rest. Operation IDs are globally unique, so the embed promotion
// the SDK uses to expose them never collides; an IDE's autocomplete
// on the Client value lists every operation directly.
//
// For escape-hatch access to the raw generated layer — bypassing
// SDK-side wrappers (the LoginWithEmail cache, header injection
// expectations, etc.) or reaching new server fields the wrapper
// hasn't picked up yet — call into the service field by its type
// name: client.AuthAPIService.LoginWithEmail(ctx)..., or
// client.PostsAPIService.CreatePost(ctx)..., and so on. The same
// concurrency / timeout rules apply because these still go through
// the SDK's *Transport.
//
// Concurrency:
//
//   - Once NewClient returns, treat the exported "wire configuration"
//     fields (APIKey, APIVersionKey, APIVersionName, AppVersion,
//     UserAgent, DeviceInfo, ConnectionType, ConnectionSpeed,
//     AcceptLanguage, RetryPolicy, EventStreamURL) as read-only.
//     The transport reads them on every request without a lock, so
//     mutating them on a Client that already has in-flight requests
//     is a data race. Set them through the With* options on NewClient.
//   - BaseURL is captured into the underlying gen.Configuration during
//     NewClient and is NOT re-read on subsequent requests; mutating
//     it has no effect (use a fresh Client with WithBaseURL instead).
//   - DeviceUUID, Tokens, UserID, ClientIP — and the unexported
//     currentEmail — change at runtime (login, refresh, session
//     restore). Internal reads go through snapshot helpers that take
//     the relevant RWMutex. External callers MUST NOT touch these
//     fields directly while requests are in flight; use SetTokens,
//     LoadSession / SaveSession, and WithDeviceUUID / WithUserID /
//     WithClientIP for safe access.
type Client struct {
	// Wire configuration.
	APIKey         string
	APIVersionKey  string
	APIVersionName string
	AppVersion     string
	UserAgent      string
	DeviceInfo     string
	// DeviceUUID is the per-device identifier sent in X-Device-UUID
	// (and used in the signed_info MD5 input). Reads / writes from
	// inside the SDK go through uuidMu so a session restore that
	// updates this field can't race with the transport's header
	// injection. External callers are expected to set this only at
	// construction via WithDeviceUUID; runtime mutation from outside
	// the SDK is unsynchronized and discouraged.
	DeviceUUID string
	uuidMu     sync.RWMutex
	// BaseURL is captured into the underlying gen.Configuration during
	// NewClient and is NOT re-read on subsequent requests. Mutating
	// client.BaseURL after construction has no effect — to target a
	// different environment, construct a new Client with WithBaseURL
	// instead.
	BaseURL string

	// EventStreamURL is the wss:// endpoint used by OpenEventStream.
	// Defaults to wss://cable.yay.space.
	EventStreamURL string

	// Per-request metadata.
	ConnectionType  string
	ConnectionSpeed string
	AcceptLanguage  string

	// ClientIP is sent as `X-Client-IP` on every request. It is populated
	// lazily after the first API call by querying GetUserTimestamp and
	// caching the returned ip_address; callers may also set it explicitly
	// via WithClientIP.
	ClientIP         string
	clientIPMu       sync.Mutex
	clientIPFetching bool

	// tokens holds the active OAuth credentials. The field is unexported
	// so callers can't bypass authMu — read snapshots via (*Client).Tokens
	// and rotate via SetTokens or the login / refresh flows.
	tokens *Tokens

	// UserID is the numeric account ID of the currently logged-in user.
	// It is populated automatically after a successful LoginWithEmail
	// call (fresh login or session restore), and is used by the upload
	// helpers (UploadAvatarImage / UploadCoverImage / UploadPostImages)
	// to fill in the path component the Yay! servers expect. Callers
	// who skip the login wrappers and call SetTokens directly should
	// also set this field via WithUserID, otherwise user-bound upload
	// methods return an error.
	//
	// Reads from inside the SDK go through userIDSnapshot so a session
	// restore can't race with an in-flight upload; external callers
	// reading Client.UserID directly while requests are in flight is a
	// data race.
	UserID int64

	// authMu guards tokens.Access / tokens.Refresh, UserID, and
	// currentEmail. These fields move together (login → activate
	// tokens + record identity, refresh → rotate tokens, session
	// restore → install all three), so a single RWMutex avoids a
	// nested-lock ordering hazard.
	authMu sync.RWMutex

	// RetryPolicy controls 5xx / 429 / network-error retries. Defaults to
	// DefaultRetryPolicy(); override with WithRetryPolicy. Setting to a
	// zero value (RetryPolicy{}) disables retries entirely.
	//
	// Construction-time only: mutating RetryPolicy on an active Client
	// is a data race because the transport copies it on every request
	// without holding a lock.
	RetryPolicy RetryPolicy

	// refreshMu serializes 401 auto-refresh: only one refresh runs at a time.
	refreshMu sync.Mutex

	// currentEmail is the email of the account the client is logged in as,
	// set by LoginWithEmail on success. It is used to key session-store
	// writes when an access token is refreshed. Guarded by authMu.
	currentEmail string

	httpClient *http.Client

	sessionStore SessionStore

	api *gen.APIClient

	// lifecycleCtx is canceled when Close runs. Background work owned
	// by the Client — the lazy X-Client-IP fetch, every EventStream
	// opened by OpenEventStream — derives from this ctx so they stop
	// promptly instead of running until their own timeout / disconnect.
	lifecycleCtx    context.Context
	lifecycleCancel context.CancelFunc

	// Every Yay! operation is promoted onto *Client through these embeds so
	// callers use e.g. client.LoginWithEmail(ctx) without a service prefix.
	//
	// When a Client method shadows an embedded one (LoginWithEmail is wrapped
	// for transparent session caching, for instance), the original gen-level
	// method is still reachable via the embed's type-name field — e.g.
	// client.AuthAPIService.LoginWithEmail(ctx).LoginEmailUserRequest(req).
	// Execute(). Use this when the wrapper's chain setters lag behind a new
	// server-side field, or when you specifically don't want the wrapper's
	// caching / auto-injection behavior.
	*gen.ActivitiesAPIService
	*gen.AppsAPIService
	*gen.AuthAPIService
	*gen.BucketsAPIService
	*gen.CallsAPIService
	*gen.ChatRoomsAPIService
	*gen.ConversationsAPIService
	*gen.EmailVerificationUrlsAPIService
	*gen.GamesAPIService
	*gen.GenresAPIService
	*gen.GiftsAPIService
	*gen.GroupMuteAPIService
	*gen.GroupsAPIService
	*gen.HiddenAPIService
	*gen.ModerationAPIService
	*gen.NotificationSettingsAPIService
	*gen.PinnedAPIService
	*gen.PostsAPIService
	*gen.ReceivedGiftsAPIService
	*gen.StickerPacksAPIService
	*gen.SurveysAPIService
	*gen.ThreadsAPIService
	*gen.UsersAPIService
}

// Option customizes a Client at construction time.
type Option func(*Client)

// WithAPIKey overrides the default Yay! API key.
func WithAPIKey(k string) Option { return func(c *Client) { c.APIKey = k } }

// WithAPIVersionKey overrides the HS256 key used to sign the X-Jwt header.
func WithAPIVersionKey(k string) Option { return func(c *Client) { c.APIVersionKey = k } }

// WithAPIVersionName overrides the X-App-Version header value.
func WithAPIVersionName(v string) Option { return func(c *Client) { c.APIVersionName = v } }

// WithAppVersion overrides the app version embedded in X-Device-Info.
func WithAppVersion(v string) Option { return func(c *Client) { c.AppVersion = v } }

// WithBaseURL overrides the API base URL (default: https://api.yay.space).
func WithBaseURL(u string) Option { return func(c *Client) { c.BaseURL = u } }

// WithUserID pre-populates Client.UserID. Use this only when
// authenticating via SetTokens (no LoginWithEmail call); the login
// wrappers fill UserID automatically.
func WithUserID(uid int64) Option { return func(c *Client) { c.UserID = uid } }

// WithEventStreamURL overrides the event-stream endpoint used by
// OpenEventStream (default: wss://cable.yay.space).
func WithEventStreamURL(u string) Option {
	return func(c *Client) { c.EventStreamURL = u }
}

// WithDeviceUUID sets a specific device UUID instead of generating a fresh
// one. Use this to keep device identity stable across runs.
func WithDeviceUUID(u string) Option { return func(c *Client) { c.DeviceUUID = u } }

// WithUserAgent overrides the UA string. Default is the device-info style
// expected by Yay! ("android 11 (3.5x 1440x2960 Galaxy S9)").
func WithUserAgent(ua string) Option { return func(c *Client) { c.UserAgent = ua } }

// WithDeviceInfo overrides the X-Device-Info header.
func WithDeviceInfo(d string) Option { return func(c *Client) { c.DeviceInfo = d } }

// WithConnectionType overrides X-Connection-Type (default: "wifi").
func WithConnectionType(t string) Option { return func(c *Client) { c.ConnectionType = t } }

// WithAcceptLanguage overrides Accept-Language (default: "ja").
func WithAcceptLanguage(l string) Option { return func(c *Client) { c.AcceptLanguage = l } }

// WithClientIP pre-sets X-Client-IP so the client does not need to perform
// the lazy GetUserTimestamp lookup on first use.
func WithClientIP(ip string) Option { return func(c *Client) { c.ClientIP = ip } }

// WithConnectionSpeed overrides X-Connection-Speed (default: "0 kbps").
func WithConnectionSpeed(s string) Option { return func(c *Client) { c.ConnectionSpeed = s } }

// WithHTTPClient sets a custom *http.Client. The client's Transport is wrapped
// to add Yay!-required headers.
//
// NewClient takes a snapshot — it shallow-copies the supplied
// *http.Client and wraps the copy's Transport. Mutating the original
// after construction (e.g. changing Timeout, swapping Transport)
// does NOT affect the SDK; pass a fresh option-set to a new Client
// instead.
func WithHTTPClient(h *http.Client) Option {
	return func(c *Client) { c.httpClient = h }
}

// WithSessionStore attaches a SessionStore so tokens can be restored across
// runs. Use NewSessionStore for a JSON-file backed store, NewMemoryStore
// for an in-process one, or provide any type satisfying SessionStore.
func WithSessionStore(s SessionStore) Option {
	return func(c *Client) { c.sessionStore = s }
}

// WithRetryPolicy overrides the default retry behavior for transient
// failures (5xx, 429, network errors). Pass RetryPolicy{} (zero value)
// to disable retries entirely. See DefaultRetryPolicy for the built-in
// defaults.
func WithRetryPolicy(p RetryPolicy) Option {
	return func(c *Client) { c.RetryPolicy = p }
}

// NewClient constructs a Client with all Yay! defaults filled in. Override
// specific fields with With... options.
//
// Call (*Client).Close when the Client is no longer needed: it cancels
// the lifecycle context that backs every goroutine the SDK starts on
// behalf of the Client (lazy X-Client-IP fetch, open event streams).
// Letting a Client go out of scope without Close is a goroutine leak
// of bounded duration — the lazy IP fetch self-terminates after 30s
// and event streams stay alive until their socket drops.
func NewClient(opts ...Option) *Client {
	userAgent := fmt.Sprintf("%s %s (%sx %s %s)",
		DefaultDeviceType, DefaultDeviceOS, DefaultDeviceDensity,
		DefaultDeviceScreen, DefaultDeviceModel)
	deviceInfo := fmt.Sprintf("yay %s %s", DefaultAppVersion, userAgent)

	lifecycleCtx, lifecycleCancel := context.WithCancel(context.Background())
	c := &Client{
		lifecycleCtx:    lifecycleCtx,
		lifecycleCancel: lifecycleCancel,
		APIKey:          DefaultAPIKey,
		APIVersionKey:   DefaultAPIVersionKey,
		APIVersionName:  DefaultAPIVersionName,
		AppVersion:      DefaultAppVersion,
		UserAgent:       userAgent,
		DeviceInfo:      deviceInfo,
		DeviceUUID:      NewUUIDv4(),
		BaseURL:         DefaultBaseURL,
		EventStreamURL:  DefaultEventStreamURL,
		ConnectionType:  DefaultConnectionType,
		ConnectionSpeed: DefaultConnectionSpeed,
		AcceptLanguage:  DefaultAcceptLanguage,
		tokens:          &Tokens{},
		RetryPolicy:     DefaultRetryPolicy(),
	}

	for _, opt := range opts {
		opt(c)
	}

	// Build HTTP client if caller didn't supply one, and wrap the transport
	// either way so our headers always run. When the caller supplied one,
	// shallow-copy it first so we never mutate their *http.Client (sharing
	// it with another consumer would otherwise leak Yay! headers / Bearer
	// tokens, and passing the same instance to two yaylib.NewClient calls
	// would stack Transport wrappers).
	if c.httpClient == nil {
		c.httpClient = &http.Client{Timeout: DefaultHTTPTimeout}
	} else {
		copied := *c.httpClient
		c.httpClient = &copied
	}
	inner := c.httpClient.Transport
	if inner == nil {
		inner = http.DefaultTransport
	}
	c.httpClient.Transport = &Transport{Base: inner, client: c}

	// Build and wire the generated client.
	cfg := gen.NewConfiguration()
	cfg.HTTPClient = c.httpClient
	cfg.UserAgent = "" // our transport sets it; a default here would override us
	cfg.Servers = gen.ServerConfigurations{{URL: c.BaseURL}}
	c.api = gen.NewAPIClient(cfg)

	c.wireServices()
	return c
}

func (c *Client) wireServices() {
	a := c.api
	c.ActivitiesAPIService = a.ActivitiesAPI
	c.AppsAPIService = a.AppsAPI
	c.AuthAPIService = a.AuthAPI
	c.BucketsAPIService = a.BucketsAPI
	c.CallsAPIService = a.CallsAPI
	c.ChatRoomsAPIService = a.ChatRoomsAPI
	c.ConversationsAPIService = a.ConversationsAPI
	c.EmailVerificationUrlsAPIService = a.EmailVerificationUrlsAPI
	c.GamesAPIService = a.GamesAPI
	c.GenresAPIService = a.GenresAPI
	c.GiftsAPIService = a.GiftsAPI
	c.GroupMuteAPIService = a.GroupMuteAPI
	c.GroupsAPIService = a.GroupsAPI
	c.HiddenAPIService = a.HiddenAPI
	c.ModerationAPIService = a.ModerationAPI
	c.NotificationSettingsAPIService = a.NotificationSettingsAPI
	c.PinnedAPIService = a.PinnedAPI
	c.PostsAPIService = a.PostsAPI
	c.ReceivedGiftsAPIService = a.ReceivedGiftsAPI
	c.StickerPacksAPIService = a.StickerPacksAPI
	c.SurveysAPIService = a.SurveysAPI
	c.ThreadsAPIService = a.ThreadsAPI
	c.UsersAPIService = a.UsersAPI
}

// Close cancels the Client's lifecycle context, stopping every goroutine
// the SDK started on its behalf — the lazy X-Client-IP fetch and every
// EventStream opened via OpenEventStream. It is safe to call multiple
// times; subsequent calls are no-ops. After Close the Client must not
// issue new requests.
func (c *Client) Close() error {
	c.lifecycleCancel()
	return nil
}

// SetTokens activates the given access / refresh tokens for subsequent calls.
// No persistence happens here; use SaveSession to write to the session cache.
// Safe to call concurrently with in-flight requests — authMu serializes the
// rotation.
func (c *Client) SetTokens(access, refresh string) {
	c.authMu.Lock()
	defer c.authMu.Unlock()
	if c.tokens == nil {
		c.tokens = &Tokens{}
	}
	c.tokens.Access = access
	c.tokens.Refresh = refresh
}

// Tokens returns a snapshot of the access / refresh tokens currently
// active on the Client. The returned value is a copy — modifying it
// has no effect on the Client. Mutate via SetTokens.
func (c *Client) Tokens() Tokens {
	c.authMu.RLock()
	defer c.authMu.RUnlock()
	if c.tokens == nil {
		return Tokens{}
	}
	return *c.tokens
}

// userIDSnapshot returns the current account UserID under authMu.
func (c *Client) userIDSnapshot() int64 {
	c.authMu.RLock()
	defer c.authMu.RUnlock()
	return c.UserID
}

// currentEmailSnapshot returns the email recorded for the active session
// under authMu. Used by the 401 auto-refresh path to key session-store
// writes.
func (c *Client) currentEmailSnapshot() string {
	c.authMu.RLock()
	defer c.authMu.RUnlock()
	return c.currentEmail
}

// setLoginIdentity atomically records the email + UserID for the active
// session. Tokens are rotated via SetTokens; this helper covers the
// non-token half of the login state.
func (c *Client) setLoginIdentity(email string, userID int64) {
	c.authMu.Lock()
	if email != "" {
		c.currentEmail = email
	}
	if userID != 0 {
		c.UserID = userID
	}
	c.authMu.Unlock()
}

// deviceUUIDSnapshot returns the current device UUID under uuidMu so
// session restores that overwrite c.DeviceUUID don't race with the
// transport's header injection or signing-input assembly.
func (c *Client) deviceUUIDSnapshot() string {
	c.uuidMu.RLock()
	defer c.uuidMu.RUnlock()
	return c.DeviceUUID
}

// setDeviceUUID atomically replaces the device UUID. Used by session
// restore so cached tokens stay paired with the device identity that
// was active when they were issued.
func (c *Client) setDeviceUUID(u string) {
	c.uuidMu.Lock()
	c.DeviceUUID = u
	c.uuidMu.Unlock()
}
