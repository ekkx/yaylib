package yaylib

import (
	"context"
	"crypto/hmac"
	"crypto/md5"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"errors"
	"fmt"
	"strconv"
)

// signedInfoSharedKey is the suffix the Yay! servers concatenate after
// api_key + uuid + timestamp before MD5-hashing.
const signedInfoSharedKey = "yayZ1"

// signedVersionPlatform is the platform identifier the Yay! servers
// expect inside the signed_version HMAC payload. It is hard-coded to
// "yay_android" because the server is most permissive for Android
// clients; using a different value risks per-platform allowlist
// rejections we have no way to inspect.
const signedVersionPlatform = "yay_android"

// GenerateSignedInfo fetches the server's current timestamp and returns
// it together with the signed_info MD5 hash bound to that timestamp.
// Pass both to the request: the server validates the hash against the
// timestamp it sees in the request body, so the two must match.
// The round-trip to GET /v2/users/timestamp eliminates clock-drift
// risk — if the device clock is off, a locally-generated timestamp
// would produce a hash the server rejects. When you already have a
// trusted timestamp (tests, batched requests reusing one ts), use
// GenerateSignedInfoAt instead to skip the network call.
func (c *Client) GenerateSignedInfo(ctx context.Context) (timestamp int64, signedInfo string, err error) {
	resp, _, err := c.GetUserTimestamp(ctx).Execute()
	if err != nil {
		return 0, "", fmt.Errorf("yaylib: fetch server timestamp: %w", err)
	}
	ts := resp.GetTime()
	return ts, c.GenerateSignedInfoAt(ts), nil
}

// GenerateSignedInfoAt computes the signed_info MD5 hash for the given
// Unix-second timestamp without any network I/O. Use this when you
// already have a synchronized timestamp; otherwise prefer
// GenerateSignedInfo, which fetches the server's view of the clock.
func (c *Client) GenerateSignedInfoAt(timestamp int64) string {
	payload := c.APIKey + c.deviceUUIDSnapshot() + strconv.FormatInt(timestamp, 10) + signedInfoSharedKey
	sum := md5.Sum([]byte(payload))
	return hex.EncodeToString(sum[:])
}

// GenerateSignedVersion returns the signed_version HMAC token that
// some Yay! API endpoints require alongside signed_info.
func (c *Client) GenerateSignedVersion() string {
	payload := signedVersionPlatform + "/" + c.APIVersionName
	h := hmac.New(sha256.New, []byte(c.APIVersionKey))
	h.Write([]byte(payload))
	return base64.StdEncoding.EncodeToString(h.Sum(nil))
}

// GenerateCallActionSignature requests a server-issued signature
// authorising the given action against a conference call. The returned
// payload bundles every field a downstream verifier needs (call id,
// sender/receiver UUIDs, action, timestamp, signature); pass it as-is
// to ValidateCallActionSignature, or read individual fields when a
// different consumer needs them.
func (c *Client) GenerateCallActionSignature(ctx context.Context, conferenceID int64, targetUserUUID, action string) (*SignaturePayload, error) {
	resp, _, err := c.CallsAPIService.GenerateCallActionSignature(ctx).
		ConferenceId(conferenceID).
		TargetUserUuid(targetUserUUID).
		Action(action).
		Execute()
	if err != nil {
		return nil, err
	}
	payload := resp.GetSignaturePayload()
	return &payload, nil
}

// ValidateCallActionSignature replays a payload obtained from
// GenerateCallActionSignature against the validation endpoint, which
// returns success only when every field still matches the server's
// view. The wrapper unpacks the payload into the six query parameters
// the endpoint expects so callers don't have to.
func (c *Client) ValidateCallActionSignature(ctx context.Context, payload *SignaturePayload) error {
	if payload == nil {
		return errors.New("yaylib: ValidateCallActionSignature: payload is nil")
	}
	_, err := c.CallsAPIService.ValidateCallActionSignature(ctx).
		CallId(payload.GetCallId()).
		SenderUuid(payload.GetSenderUuid()).
		ReceiverUuid(payload.GetReceiverUuid()).
		Action(payload.GetAction()).
		Timestamp(payload.GetTimestamp()).
		Signature(payload.GetSignature()).
		Execute()
	return err
}
