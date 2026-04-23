package yaylib

import "github.com/ekkx/yaylib/gen"

// APIError is the error returned by every Client method when the Yay! server
// responds with a non-2xx status. It exposes the raw response body via Body()
// and the message via Error().
//
// Example:
//
//	_, _, err := client.CreatePost(ctx)./* ... */ Execute()
//	if err != nil {
//	    var ae *yaylib.APIError
//	    if errors.As(err, &ae) {
//	        log.Printf("status: %s, body: %s", ae.Error(), ae.Body())
//	    }
//	}
type APIError = gen.GenericOpenAPIError
