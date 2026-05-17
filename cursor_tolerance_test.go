package yaylib

import (
	"encoding/json"
	"testing"

	"github.com/ekkx/yaylib/v2/gen"
)

// The pagination cursor is a string in the schema, but the server
// sends it as a bare JSON number on some endpoints and a JSON string
// on others. Typed decoding must accept both and expose the string
// form either way.
func TestPaginationCursorToleratesNumberAndString(t *testing.T) {
	cases := []struct{ name, body, want string }{
		{"number wire", `{"next_page_value":585851614,"posts":[]}`, "585851614"},
		{"string wire", `{"next_page_value":"abc123","posts":[]}`, "abc123"},
		{"null wire", `{"next_page_value":null,"posts":[]}`, ""},
	}
	for _, c := range cases {
		var r gen.PostsResponse
		if err := json.Unmarshal([]byte(c.body), &r); err != nil {
			t.Fatalf("%s: unmarshal failed: %v", c.name, err)
		}
		if got := r.GetNextPageValue(); got != c.want {
			t.Fatalf("%s: GetNextPageValue() = %q, want %q", c.name, got, c.want)
		}
	}
}
