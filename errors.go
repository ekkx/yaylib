package yaylib

import (
	"encoding/json"
	"errors"

	"github.com/ekkx/yaylib/v2/gen"
)

// APIError is the error returned by every Client method when the Yay! server
// responds with a non-2xx status. It exposes the raw response body via Body()
// and the message via Error().
//
// Most callers don't need to touch APIError directly — use CodeOf(err) to
// extract the typed error code and ErrorResponseOf(err) to access fields like
// BanUntil / RetryIn / RemainingQuota:
//
//	_, _, err := client.CreatePost(ctx).Execute()
//	switch yaylib.CodeOf(err) {
//	case yaylib.ErrCodeRequired2FA:
//	    // show 2FA prompt
//	case yaylib.ErrCodeUserTemporaryBanned:
//	    if r := yaylib.ErrorResponseOf(err); r != nil && r.BanUntil != nil {
//	        log.Printf("banned until %d", *r.BanUntil)
//	    }
//	}
//
// When you need the raw HTTP response body, reach for APIError directly via
// errors.As:
//
//	var apiErr *yaylib.APIError
//	if errors.As(err, &apiErr) {
//	    log.Printf("body=%s", apiErr.Body())
//	}
type APIError = gen.GenericOpenAPIError

// ErrorResponse is the parsed Yay! server error payload — the JSON body
// returned alongside a non-2xx HTTP status.
type ErrorResponse struct {
	Code           ErrorCode `json:"error_code,omitempty"`
	Message        string    `json:"message,omitempty"`
	BanUntil       *int64    `json:"ban_until,omitempty"`
	RetryIn        *int64    `json:"retry_in,omitempty"`
	RemainingQuota *int64    `json:"remaining_quota,omitempty"`
}

// ErrorResponseOf extracts and parses the ErrorResponse payload from err.
// Returns nil when err is not an APIError or the body does not decode as
// JSON. A malformed body (e.g. empty, HTML from an upstream proxy) is
// reported as nil rather than an error; callers that care about the raw
// bytes can fall back to errors.As(err, &apiErr) + apiErr.Body().
func ErrorResponseOf(err error) *ErrorResponse {
	var apiErr *APIError
	if !errors.As(err, &apiErr) {
		return nil
	}
	body := apiErr.Body()
	if len(body) == 0 {
		return nil
	}
	var r ErrorResponse
	if err := json.Unmarshal(body, &r); err != nil {
		return nil
	}
	return &r
}

// CodeOf returns the Yay! error code from err, or ErrCodeUnknown when err
// carries no recognizable code (not an APIError, missing body, or body with
// no "error_code" field).
func CodeOf(err error) ErrorCode {
	if r := ErrorResponseOf(err); r != nil {
		return r.Code
	}
	return ErrCodeUnknown
}
