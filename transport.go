package yaylib

import (
	"net/http"
	"strconv"
	"time"
)

// TokenStore holds the active OAuth credentials. The Transport reads Access
// for every request; Refresh is reserved for a future auto-refresh feature.
type TokenStore struct {
	Access  string
	Refresh string
}

// Transport is an http.RoundTripper that injects every header the Yay!
// servers expect from the official Android app, plus a Bearer token when
// one has been set.
//
// It is not intended for direct use by callers; NewClient wires one in front
// of the client's http.Client.
type Transport struct {
	Base   http.RoundTripper
	client *Client
}

func (t *Transport) RoundTrip(r *http.Request) (*http.Response, error) {
	// Clone the outbound request so we don't mutate state owned by the
	// caller; this is standard practice for a custom RoundTripper.
	r = r.Clone(r.Context())

	set := func(k, v string) {
		if r.Header.Get(k) == "" {
			r.Header.Set(k, v)
		}
	}
	c := t.client
	set("User-Agent", c.UserAgent)
	set("X-App-Version", c.APIVersionName)
	set("X-Device-Info", c.DeviceInfo)
	set("X-Device-UUID", c.DeviceUUID)
	set("X-Connection-Type", c.ConnectionType)
	set("X-Connection-Speed", c.ConnectionSpeed)
	set("Accept-Language", c.AcceptLanguage)

	// X-Timestamp must be fresh per request — don't let a caller pin a
	// stale value.
	r.Header.Set("X-Timestamp", strconv.FormatInt(time.Now().Unix(), 10))

	if c.Tokens != nil && c.Tokens.Access != "" {
		set("Authorization", "Bearer "+c.Tokens.Access)
	}
	return t.Base.RoundTrip(r)
}
