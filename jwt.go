package yaylib

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"time"
)

// NewXJwt returns the short-lived HS256 JWT required by some write endpoints
// (most visibly CreatePost). The token is signed with the client's
// APIVersionKey, carries only {iat, exp=iat+5}, and must be regenerated for
// each call because the 5-second validity window expires quickly.
//
// Usage:
//
//	resp, _, err := client.CreatePost(ctx).
//	    XJwt(client.NewXJwt()).
//	    Text("...").
//	    PostType("text").
//	    Execute()
func (c *Client) NewXJwt() string {
	return newXJwt(c.APIVersionKey, time.Now().Unix(), 5)
}

// NewXJwtWithTTL is NewXJwt but lets the caller pick the validity window
// (in seconds). Useful mainly for tests / debugging.
func (c *Client) NewXJwtWithTTL(ttl time.Duration) string {
	return newXJwt(c.APIVersionKey, time.Now().Unix(), int64(ttl.Seconds()))
}

func newXJwt(key string, iat, ttl int64) string {
	header := base64.RawURLEncoding.EncodeToString([]byte(`{"alg":"HS256"}`))
	payload := base64.RawURLEncoding.EncodeToString(
		[]byte(fmt.Sprintf(`{"iat":%d,"exp":%d}`, iat, iat+ttl)))
	signingInput := header + "." + payload
	mac := hmac.New(sha256.New, []byte(key))
	mac.Write([]byte(signingInput))
	sig := base64.RawURLEncoding.EncodeToString(mac.Sum(nil))
	return signingInput + "." + sig
}
