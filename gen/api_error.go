// Code generated from the Yay! API schema; DO NOT EDIT.

package gen

// NewAPIError builds a GenericOpenAPIError that carries a raw response body
// and message. Hand-written request paths that bypass the generated clients
// (for example a login issued through the unwrapped HTTP layer) use it so
// their non-2xx responses surface as the same typed error the generated
// clients return — CodeOf / ErrorResponseOf / errors.As keep working
// uniformly regardless of which path produced the error.
func NewAPIError(body []byte, message string) *GenericOpenAPIError {
	return &GenericOpenAPIError{body: body, error: message}
}
