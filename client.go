// Package yaylib is a Go client for Yay! (jp.nanameue.yay).
//
// Usage:
//
//	client, _ := yaylib.NewClient()
//	req := yaylib.NewLoginEmailUserRequest()
//	req.SetApiKey(client.APIKey)
//	req.SetEmail(email)
//	req.SetPassword(password)
//	req.SetUuid(client.DeviceUUID)
//	resp, _, err := client.
//	    LoginWithEmail(ctx).
//	    LoginEmailUserRequest(*req).
//	    Execute()
//
// Every Yay! operation is reachable directly on *Client — for example
// LoginWithEmail, GetRecommendedTimeline, CreatePost, BlockUser.
//
// The client also takes care of:
//
//   - the HTTP headers the Yay! servers expect from the official mobile app
//     (User-Agent, X-App-Version, X-Device-Info, X-Device-UUID, X-Timestamp,
//     X-Connection-*, Accept-Language) and a Bearer token once credentials
//     have been set,
//   - an optional SQLite session cache so you don't re-login on every run
//     (Yay! rate-limits the login endpoint),
//   - a helper that mints the short-lived HS256 `X-Jwt` token some write
//     endpoints require (see Client.NewXJwt).
//
// `APIKey`, `DeviceUUID`, and the related metadata are exposed as fields so
// callers can populate request bodies that need them (body builders replace
// the entire body, so these fields can't be injected transparently).
package yaylib

import (
	"fmt"
	"net/http"
	"sync"
	"time"

	"github.com/ekkx/yaylib/gen"
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
// Exported configuration fields are safe to read at any time and safe to
// mutate before any network call has been issued. Changing them after
// issuing calls is allowed but may be racy with in-flight requests.
type Client struct {
	// Wire configuration.
	APIKey         string
	APIVersionKey  string
	APIVersionName string
	AppVersion     string
	UserAgent      string
	DeviceInfo     string
	DeviceUUID     string
	BaseURL        string

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

	// Active credentials. The transport reads Access to set `Authorization:
	// Bearer`. Populate via SetTokens or by restoring a cached Session.
	Tokens *TokenStore

	// refreshMu serializes 401 auto-refresh: only one refresh runs at a time.
	refreshMu sync.Mutex

	// currentEmail is the email of the account the client is logged in as,
	// set by LoginWithEmail on success. It is used to key session-store
	// writes when an access token is refreshed.
	currentEmail string

	transport  *Transport
	httpClient *http.Client

	sessionStore SessionStore

	api *gen.APIClient

	// Every Yay! operation is promoted onto *Client through these embeds so
	// callers use e.g. client.LoginWithEmail(ctx) without a service prefix.
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
func WithHTTPClient(h *http.Client) Option {
	return func(c *Client) { c.httpClient = h }
}

// WithSessionStore attaches a SessionStore so tokens can be restored across
// runs. Use NewSessionStore for a JSON-file backed store, NewMemoryStore
// for an in-process one, or provide any type satisfying SessionStore.
func WithSessionStore(s SessionStore) Option {
	return func(c *Client) { c.sessionStore = s }
}

// NewClient constructs a Client with all Yay! defaults filled in. Override
// specific fields with With... options.
func NewClient(opts ...Option) (*Client, error) {
	userAgent := fmt.Sprintf("%s %s (%sx %s %s)",
		DefaultDeviceType, DefaultDeviceOS, DefaultDeviceDensity,
		DefaultDeviceScreen, DefaultDeviceModel)
	deviceInfo := fmt.Sprintf("yay %s %s", DefaultAppVersion, userAgent)

	c := &Client{
		APIKey:          DefaultAPIKey,
		APIVersionKey:   DefaultAPIVersionKey,
		APIVersionName:  DefaultAPIVersionName,
		AppVersion:      DefaultAppVersion,
		UserAgent:       userAgent,
		DeviceInfo:      deviceInfo,
		DeviceUUID:      NewUUIDv4(),
		BaseURL:         DefaultBaseURL,
		ConnectionType:  DefaultConnectionType,
		ConnectionSpeed: DefaultConnectionSpeed,
		AcceptLanguage:  DefaultAcceptLanguage,
		Tokens:          &TokenStore{},
	}

	for _, opt := range opts {
		opt(c)
	}

	// Build HTTP client if caller didn't supply one, and wrap the transport
	// either way so our headers always run.
	if c.httpClient == nil {
		c.httpClient = &http.Client{Timeout: DefaultHTTPTimeout}
	}
	inner := c.httpClient.Transport
	if inner == nil {
		inner = http.DefaultTransport
	}
	c.transport = &Transport{Base: inner, client: c}
	c.httpClient.Transport = c.transport

	// Build and wire the generated client.
	cfg := gen.NewConfiguration()
	cfg.HTTPClient = c.httpClient
	cfg.UserAgent = "" // our transport sets it; a default here would override us
	cfg.Servers = gen.ServerConfigurations{{URL: c.BaseURL}}
	c.api = gen.NewAPIClient(cfg)

	c.wireServices()
	return c, nil
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

// API returns the underlying generated client. Use this only when you need
// something not reachable through the promoted methods on Client (e.g. raw
// *gen.Configuration access or adding a request interceptor).
func (c *Client) API() *gen.APIClient { return c.api }

// SetTokens activates the given access / refresh tokens for subsequent calls.
// No persistence happens here; use SaveSession to write to the session cache.
func (c *Client) SetTokens(access, refresh string) {
	c.Tokens.Access = access
	c.Tokens.Refresh = refresh
}

// Close releases any resources held by the client. It is safe to call
// whether or not a session store was configured.
func (c *Client) Close() error {
	return nil
}
