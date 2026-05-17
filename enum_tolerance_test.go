package yaylib

import (
	"encoding/json"
	"testing"

	"github.com/ekkx/yaylib/v2/gen"
)

// Enum-typed fields are decoded leniently: a value the SDK does not
// know about is still accepted and kept verbatim on the typed field
// (no decode error), so a server that adds a new variant does not
// break older clients. IsValid() reports false for such a value, and
// the named constants stay usable for the values the SDK does know.
func TestUnknownEnumValueIsAccepted(t *testing.T) {
	var p gen.Post
	if err := json.Unmarshal([]byte(`{"post_type":"__mock_unknown_enum__"}`), &p); err != nil {
		t.Fatalf("unknown enum value: unmarshal failed: %v", err)
	}
	if got := string(p.GetPostType()); got != "__mock_unknown_enum__" {
		t.Fatalf("GetPostType() = %q, want %q", got, "__mock_unknown_enum__")
	}
	if p.GetPostType().IsValid() {
		t.Fatalf("IsValid() = true for unknown value, want false")
	}

	var known gen.Post
	if err := json.Unmarshal([]byte(`{"post_type":"image"}`), &known); err != nil {
		t.Fatalf("known enum value: unmarshal failed: %v", err)
	}
	if got := known.GetPostType(); got != gen.POSTTYPE_Image {
		t.Fatalf("GetPostType() = %q, want %q", got, gen.POSTTYPE_Image)
	}
	if !known.GetPostType().IsValid() {
		t.Fatalf("IsValid() = false for known value, want true")
	}
}
