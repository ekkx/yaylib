package yaylib

import (
	"context"
	"testing"
)

// PORTING:S33
// Flat surface: in Go every operation is promoted onto *Client via
// embed, and a hand-written wrapper shadows the embedded generated op
// of the same name. Both are guaranteed structurally — referencing the
// promoted method compiles only if promotion holds, and LoginWithEmail
// resolves to the cached wrapper (its builder type), not the generated
// request. Hermetic: no network is performed.
func TestFlatFacade_OpsPromotedAndWrapperShadows(t *testing.T) {
	c := NewClient()
	defer c.Close()

	// Taking these as method values compiles only if the operations
	// are promoted directly onto *Client (no service prefix) — the
	// embed-promotion that gives Go the flat surface for free.
	_ = c.GetTimeline
	_ = c.GetRecommendedTimeline

	// LoginWithEmail resolves to the cached wrapper (returns its
	// chain builder), shadowing the embedded generated op of the same
	// name — the embed-shadowing rule. Reaching .Email/.Password on
	// the result compiles only against the wrapper builder.
	_ = c.LoginWithEmail(context.Background()).Email("").Password("")

	if c.UserID != 0 {
		t.Fatalf("fresh client should not be logged in, got UserID=%d", c.UserID)
	}
}
