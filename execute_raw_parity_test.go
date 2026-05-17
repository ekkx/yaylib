package yaylib

import (
	"context"
	"errors"
	"testing"
)

// ExecuteRaw parity, driven against the shared server's
// body-malformed contract: the endpoint answers HTTP 200 with a JSON
// scalar no typed model can decode. The typed Execute() path must
// surface that as an *APIError carrying the raw body, while the
// sibling ExecuteRaw() path must hand back the unparsed bytes and the
// successful response untouched — the per-call escape hatch for
// production wire that the typed decoder cannot absorb. The contract
// is server-enforced and identical across the three languages; the
// test auto-skips when the shared server is not configured.
func TestExecuteRaw_AbsorbsTypedDecodeFailure(t *testing.T) {
	c := mockClient(t, "body-malformed")
	ctx := context.Background()

	if _, _, err := c.GetBucketPresignedUrls(ctx).Execute(); err == nil {
		t.Fatal("typed Execute() on malformed body: expected error, got nil")
	} else {
		var apiErr *APIError
		if !errors.As(err, &apiErr) {
			t.Fatalf("typed Execute() error = %T, want *APIError", err)
		}
		if len(apiErr.Body()) == 0 {
			t.Fatal("typed Execute() *APIError carried no body")
		}
	}

	body, resp, rawErr := c.GetBucketPresignedUrls(ctx).ExecuteRaw()
	if rawErr != nil {
		t.Fatalf("ExecuteRaw() on malformed body: unexpected error: %v", rawErr)
	}
	if resp == nil {
		t.Fatal("ExecuteRaw() returned nil response")
	}
	if resp.StatusCode != 200 {
		t.Fatalf("ExecuteRaw() status = %d, want 200", resp.StatusCode)
	}
	if len(body) == 0 {
		t.Fatal("ExecuteRaw() returned empty body")
	}
}
