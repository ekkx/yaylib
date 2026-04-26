package yaylib

import (
	"context"
	"crypto/hmac"
	"crypto/md5"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGenerateSignedInfoAt_Deterministic(t *testing.T) {
	c := NewClient()
	c.APIKey = "test-api-key"
	c.DeviceUUID = "test-uuid"

	got := c.GenerateSignedInfoAt(1700000000)
	want := md5HexHash("test-api-key" + "test-uuid" + "1700000000" + "yayZ1")
	if got != want {
		t.Errorf("GenerateSignedInfoAt(1700000000) = %q, want %q", got, want)
	}
	if len(got) != 32 {
		t.Errorf("hash length = %d, want 32", len(got))
	}
}

func TestGenerateSignedInfoAt_DependsOnClientCredentials(t *testing.T) {
	a := NewClient()
	a.APIKey = "key-A"
	a.DeviceUUID = "uuid-shared"

	b := NewClient()
	b.APIKey = "key-B"
	b.DeviceUUID = "uuid-shared"

	if a.GenerateSignedInfoAt(1) == b.GenerateSignedInfoAt(1) {
		t.Error("different APIKeys must produce different hashes")
	}
}

func TestGenerateSignedInfo_UsesServerTimestamp(t *testing.T) {
	const serverTime int64 = 1750000000
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/v2/users/timestamp" {
			http.Error(w, "unexpected path", http.StatusNotFound)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(map[string]any{
			"time":       serverTime,
			"ip_address": "203.0.113.5",
		})
	}))
	defer srv.Close()

	c := NewClient(WithBaseURL(srv.URL))
	c.APIKey = "test-api-key"
	c.DeviceUUID = "test-uuid"

	ts, sig, err := c.GenerateSignedInfo(context.Background())
	if err != nil {
		t.Fatalf("GenerateSignedInfo: %v", err)
	}
	if ts != serverTime {
		t.Errorf("ts = %d, want %d", ts, serverTime)
	}
	wantSig := md5HexHash("test-api-key" + "test-uuid" + "1750000000" + "yayZ1")
	if sig != wantSig {
		t.Errorf("sig = %q, want %q", sig, wantSig)
	}
	if got := c.GenerateSignedInfoAt(ts); got != sig {
		t.Errorf("GenerateSignedInfoAt(%d) = %q, but GenerateSignedInfo returned %q", ts, got, sig)
	}
}

func TestGenerateSignedInfo_PropagatesServerError(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, `{"error_code": 1}`, http.StatusInternalServerError)
	}))
	defer srv.Close()

	c := NewClient(WithBaseURL(srv.URL), WithRetryPolicy(RetryPolicy{}))
	_, _, err := c.GenerateSignedInfo(context.Background())
	if err == nil {
		t.Fatal("expected error from server 500")
	}
}

// TestGenerateSignedVersion_Golden pins the value the SDK produces with
// the default APIVersionKey + APIVersionName so any accidental change
// to the formula (encoding, message body, key) is caught immediately.
func TestGenerateSignedVersion_Golden(t *testing.T) {
	c := NewClient()
	got := c.GenerateSignedVersion()
	const want = "zyKEqdZqaEqJjSRItHvchU9HmgWHNLtn02sWXwzQ4iQ="
	if got != want {
		t.Errorf("GenerateSignedVersion = %q, want %q", got, want)
	}
}

func TestGenerateSignedVersion_FollowsAPIVersionName(t *testing.T) {
	c := NewClient(WithAPIVersionName("9.99"))

	got := c.GenerateSignedVersion()
	want := hmacSHA256Base64(c.APIVersionKey, "yay_android/9.99")
	if got != want {
		t.Errorf("GenerateSignedVersion(9.99) = %q, want %q", got, want)
	}
}

func TestGenerateSignedVersion_DependsOnAPIVersionKey(t *testing.T) {
	a := NewClient()
	b := NewClient()
	b.APIVersionKey = "00000000000000000000000000000000"

	if a.GenerateSignedVersion() == b.GenerateSignedVersion() {
		t.Error("different APIVersionKeys must produce different signed_version values")
	}
}

func TestGenerateSignedVersion_NoTrailingWhitespace(t *testing.T) {
	got := NewClient().GenerateSignedVersion()
	if got == "" {
		t.Fatal("got empty signed_version")
	}
	if got[len(got)-1] == '\n' || got[len(got)-1] == ' ' {
		t.Errorf("trailing whitespace in %q", got)
	}
	// HMAC-SHA256 → 32 bytes → 44-char standard base64 with one '=' pad.
	if len(got) != 44 {
		t.Errorf("len = %d, want 44 (32-byte HMAC, std base64)", len(got))
	}
}

func md5HexHash(s string) string {
	sum := md5.Sum([]byte(s))
	return hex.EncodeToString(sum[:])
}

func hmacSHA256Base64(key, msg string) string {
	h := hmac.New(sha256.New, []byte(key))
	h.Write([]byte(msg))
	return base64.StdEncoding.EncodeToString(h.Sum(nil))
}
