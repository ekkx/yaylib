// Code generated from the Yay! API schema; DO NOT EDIT.

package gen

import (
	"io"
	"net/http"
)

// executeRaw drives the per-request ExecuteRaw helpers. It reads the
// response body regardless of whether the corresponding typed Execute
// would have rejected the JSON.
//
// Error policy:
//   - Pure transport failure (httpResp == nil): the original error
//     from Execute is propagated. We can't read a body we never got.
//   - HTTP failure (4xx / 5xx): returned as APIError with the body
//     attached, so CodeOf / ErrorResponseOf still work on raw mode.
//   - 2xx with a typed-decode error inside Execute: that error is
//     absorbed — the whole point of the raw escape hatch is to let
//     callers see bytes the typed model would have rejected.
func executeRaw(httpResp *http.Response, executeErr error) ([]byte, *http.Response, error) {
	if httpResp == nil {
		return nil, nil, executeErr
	}
	if httpResp.Body == nil {
		return nil, httpResp, &GenericOpenAPIError{error: "no response body"}
	}
	body, err := io.ReadAll(httpResp.Body)
	if err != nil {
		return nil, httpResp, err
	}
	if httpResp.StatusCode/100 != 2 {
		return body, httpResp, &GenericOpenAPIError{body: body, error: httpResp.Status}
	}
	return body, httpResp, nil
}
