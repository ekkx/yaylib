package yaylib

import (
	"bytes"
	"context"
	"regexp"
	"strings"
	"testing"
)

// Upload behavior parity, driven against the shared server's upload
// contract: GET /v1/buckets/presigned_urls returns one server-canonical
// ("s3/"-prefixed) filename + an in-process S3 URL per requested name;
// GET /v1/users/presigned_url does the same for video; the S3 receiver
// answers 200, or 403 if the PUT carries Authorization.
//
// The shared S3 receiver intentionally does not echo the Content-Type,
// the stored bytes, or the file_names[] it saw (behavioral-fidelity
// only — same design as the empty happy-path bodies). So the parity
// assertions are the client-observable ones: success/error, the
// returned canonical filename shape (the server just prepends "s3/" to
// the SDK-built name), and the server-enforced no-bearer rule on the
// presigned PUT. The PUT Content-Type / body / thumbnail-name
// inspection, the image-processing internals, the empty-presigned and
// non-image guards, and the pure filename/path/MIME helpers stay as
// local fixtures in upload_test.go.

func TestUploadAvatarImage_HappyPath(t *testing.T) {
	c := mockClient(t, "")
	c.UserID = 42
	body := encodeImage(t, "jpeg", 100, 100)

	name, err := c.UploadAvatarImage(context.Background(), Upload{
		Filename: "me.jpg",
		Body:     bytes.NewReader(body),
	})
	if err != nil {
		t.Fatalf("UploadAvatarImage: %v", err)
	}
	if !strings.HasPrefix(name, "s3/user/42/avatar/") {
		t.Errorf("name = %q, want s3/user/42/avatar/ prefix", name)
	}
	if !strings.Contains(name, "_size_100x100") {
		t.Errorf("name = %q, missing _size_100x100 suffix", name)
	}
}

func TestUploadAvatarImage_FilenameShape(t *testing.T) {
	c := mockClient(t, "")
	c.UserID = 1
	body := encodeImage(t, "png", 12, 13)

	name, err := c.UploadAvatarImage(context.Background(), Upload{
		Filename: "x.png", Body: bytes.NewReader(body),
	})
	if err != nil {
		t.Fatalf("UploadAvatarImage: %v", err)
	}
	re := regexp.MustCompile(`^s3/user/1/avatar/[0-9A-Za-z]{16}_\d{13}_0_size_12x13\.png$`)
	if !re.MatchString(name) {
		t.Errorf("name = %q, does not match %s", name, re)
	}
}

func TestUploadAvatarImage_PresignedFailure(t *testing.T) {
	// fail-503-times-1 makes the first request in the session — the
	// presigned-URL GET — fail; with retry disabled the upload errors.
	c := mockClient(t, "fail-503-times-1", WithRetryPolicy(RetryPolicy{}))
	c.UserID = 1
	body := encodeImage(t, "png", 5, 5)

	if _, err := c.UploadAvatarImage(context.Background(), Upload{
		Filename: "a.png", Body: bytes.NewReader(body),
	}); err == nil {
		t.Fatal("expected error from failing presigned-URL fetch")
	}
}

func TestUploadVideo_HappyPath(t *testing.T) {
	c := mockClient(t, "")
	mp4Body := []byte("\x00\x00\x00 ftypisom--mp4-bytes--")

	name, err := c.UploadVideo(context.Background(), bytes.NewReader(mp4Body))
	if err != nil {
		t.Fatalf("UploadVideo: %v", err)
	}
	if !strings.HasSuffix(name, ".mp4") {
		t.Errorf("name = %q, want .mp4 suffix", name)
	}
}

func TestPresignedPUT_DoesNotLeakBearer(t *testing.T) {
	// Tokens are set, but the presigned PUT must authenticate via the
	// signed query only. The shared S3 receiver answers 403 if the PUT
	// carries Authorization, so a successful upload proves no bearer
	// leaked — the check is server-enforced and identical across the
	// three languages.
	c := mockClient(t, "")
	c.UserID = 1
	c.SetTokens("SECRET-ACCESS-TOKEN", "REF")
	body := encodeImage(t, "png", 5, 5)

	if _, err := c.UploadAvatarImage(context.Background(), Upload{
		Filename: "x.png", Body: bytes.NewReader(body),
	}); err != nil {
		t.Fatalf("UploadAvatarImage failed; a leaked bearer would 403 the PUT: %v", err)
	}
}
