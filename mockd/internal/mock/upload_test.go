package mock

import (
	"encoding/json"
	"io"
	"net/http"
	"strings"
	"testing"
)

func TestWSTokenOwned(t *testing.T) {
	ts := newTestServer(t)
	defer ts.Close()
	r := do(t, ts, http.MethodGet, "/v1/users/ws_token", "", "")
	var tok wsTokenResponse
	if err := json.Unmarshal([]byte(body(t, r)), &tok); err != nil {
		t.Fatalf("decode: %v", err)
	}
	if tok.Token != mockWSToken {
		t.Fatalf("token = %q, want %q", tok.Token, mockWSToken)
	}
}

func TestBucketPresignedURLs(t *testing.T) {
	ts := newTestServer(t)
	defer ts.Close()
	r := do(t, ts, http.MethodGet,
		"/v1/buckets/presigned_urls?file_names[]=user/1/avatar/main.jpg&file_names[]=user/1/avatar/thumb.jpg",
		"", "")
	var resp presignedURLsResponse
	if err := json.Unmarshal([]byte(body(t, r)), &resp); err != nil {
		t.Fatalf("decode: %v", err)
	}
	if len(resp.PresignedUrls) != 2 {
		t.Fatalf("got %d urls, want 2", len(resp.PresignedUrls))
	}
	first := resp.PresignedUrls[0]
	if first.Filename != "s3/user/1/avatar/main.jpg" {
		t.Fatalf("canonical filename = %q, want s3/ prefix", first.Filename)
	}
	if !strings.Contains(first.URL, s3Prefix+"user/1/avatar/main.jpg") {
		t.Fatalf("url = %q, want S3 receiver path", first.URL)
	}
}

func TestUserPresignedURLVideo(t *testing.T) {
	ts := newTestServer(t)
	defer ts.Close()
	r := do(t, ts, http.MethodGet, "/v1/users/presigned_url?video_file_name=clip_123.mp4", "", "")
	var resp presignedURLResponse
	if err := json.Unmarshal([]byte(body(t, r)), &resp); err != nil {
		t.Fatalf("decode: %v", err)
	}
	if !strings.HasSuffix(resp.PresignedURL, s3Prefix+"clip_123.mp4") {
		t.Fatalf("presigned_url = %q", resp.PresignedURL)
	}
}

func TestS3PutAcceptsBodyWithoutAuthorization(t *testing.T) {
	ts := newTestServer(t)
	defer ts.Close()
	req, _ := http.NewRequest(http.MethodPut, ts.URL+s3Prefix+"user/1/avatar/main.jpg",
		strings.NewReader("binary-bytes"))
	resp, err := ts.Client().Do(req)
	if err != nil {
		t.Fatalf("put: %v", err)
	}
	io.Copy(io.Discard, resp.Body)
	resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		t.Fatalf("put status = %d, want 200", resp.StatusCode)
	}
}

func TestS3PutRejectsBearer(t *testing.T) {
	ts := newTestServer(t)
	defer ts.Close()
	req, _ := http.NewRequest(http.MethodPut, ts.URL+s3Prefix+"user/1/avatar/main.jpg",
		strings.NewReader("x"))
	req.Header.Set("Authorization", "Bearer leaked")
	resp, err := ts.Client().Do(req)
	if err != nil {
		t.Fatalf("put: %v", err)
	}
	io.Copy(io.Discard, resp.Body)
	resp.Body.Close()
	if resp.StatusCode != http.StatusForbidden {
		t.Fatalf("put-with-bearer status = %d, want 403", resp.StatusCode)
	}
}

func TestPresignedRespectsScenario(t *testing.T) {
	ts := newTestServer(t)
	defer ts.Close()
	// A failing scenario still applies to the presigned GET.
	r := do(t, ts, http.MethodGet, "/v1/buckets/presigned_urls?file_names[]=a",
		"fail-503-times-1", "ps1")
	if r.StatusCode != 503 {
		t.Fatalf("scenario should apply to presigned GET: got %d", r.StatusCode)
	}
	r = do(t, ts, http.MethodGet, "/v1/buckets/presigned_urls?file_names[]=a",
		"fail-503-times-1", "ps1")
	if r.StatusCode != 200 {
		t.Fatalf("after budget: got %d, want 200", r.StatusCode)
	}
}
