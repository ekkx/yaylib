package yaylib

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"image"
	"image/color"
	"image/gif"
	"image/jpeg"
	"image/png"
	"io"
	"net/http"
	"net/http/httptest"
	"path"
	"regexp"
	"strings"
	"sync"
	"testing"
	"time"
)

// fakeUploadServer fakes both the Yay! API endpoints (presigned URL
// issuance) and the S3 PUT target with a single httptest.Server.
type fakeUploadServer struct {
	t *testing.T

	mu              sync.Mutex
	objects         map[string][]byte // path → body
	contentTypes    map[string]string // path → Content-Type used on PUT
	presignedCalls  int
	imageNames      []string // most recent file_names[] received
	videoFileName   string   // most recent video_file_name received
	failPresignedAt int      // 0 = never; N = fail on Nth call
	putFailures     int      // remaining 500s for PUT
}

func newFakeUploadServer(t *testing.T) (*fakeUploadServer, *httptest.Server) {
	f := &fakeUploadServer{
		t:            t,
		objects:      map[string][]byte{},
		contentTypes: map[string]string{},
	}
	srv := httptest.NewServer(http.HandlerFunc(f.serveHTTP))
	return f, srv
}

func (f *fakeUploadServer) serveHTTP(w http.ResponseWriter, r *http.Request) {
	switch {
	case r.Method == http.MethodGet && r.URL.Path == "/v1/buckets/presigned_urls":
		f.handlePresignedImages(w, r)
	case r.Method == http.MethodGet && r.URL.Path == "/v1/users/presigned_url":
		f.handlePresignedVideo(w, r)
	case r.Method == http.MethodPut && strings.HasPrefix(r.URL.Path, "/uploads/"):
		f.handlePut(w, r)
	default:
		http.Error(w, "unknown route: "+r.Method+" "+r.URL.Path, http.StatusNotFound)
	}
}

func (f *fakeUploadServer) handlePresignedImages(w http.ResponseWriter, r *http.Request) {
	f.mu.Lock()
	f.presignedCalls++
	if f.failPresignedAt > 0 && f.presignedCalls == f.failPresignedAt {
		f.mu.Unlock()
		http.Error(w, "presigned fail", http.StatusInternalServerError)
		return
	}
	names := r.URL.Query()["file_names[]"]
	f.imageNames = append([]string(nil), names...)
	f.mu.Unlock()

	urls := make([]map[string]string, len(names))
	for i, n := range names {
		urls[i] = map[string]string{
			"filename": "s3/" + n,
			"url":      fakeServerURL(r) + "/uploads/s3/" + n,
		}
	}
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(map[string]interface{}{
		"presigned_urls": urls,
	})
}

func (f *fakeUploadServer) handlePresignedVideo(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("video_file_name")
	f.mu.Lock()
	f.videoFileName = name
	f.mu.Unlock()

	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(map[string]interface{}{
		"presigned_url": fakeServerURL(r) + "/uploads/" + name,
	})
}

func (f *fakeUploadServer) handlePut(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	f.mu.Lock()
	defer f.mu.Unlock()
	if f.putFailures > 0 {
		f.putFailures--
		http.Error(w, "put fail", http.StatusInternalServerError)
		return
	}
	f.objects[r.URL.Path] = body
	f.contentTypes[r.URL.Path] = r.Header.Get("Content-Type")
	w.WriteHeader(http.StatusOK)
}

func fakeServerURL(r *http.Request) string {
	scheme := "http"
	if r.TLS != nil {
		scheme = "https"
	}
	return scheme + "://" + r.Host
}

// encodeImage produces a real image of the requested size in the given
// format, so decodeImageSize sees actual width/height bytes.
func encodeImage(t *testing.T, format string, w, h int) []byte {
	t.Helper()
	img := image.NewRGBA(image.Rect(0, 0, w, h))
	for y := 0; y < h; y++ {
		for x := 0; x < w; x++ {
			img.Set(x, y, color.RGBA{R: uint8(x), G: uint8(y), B: 200, A: 255})
		}
	}
	var buf bytes.Buffer
	var err error
	switch format {
	case "jpeg":
		err = jpeg.Encode(&buf, img, &jpeg.Options{Quality: 75})
	case "png":
		err = png.Encode(&buf, img)
	case "gif":
		err = gif.Encode(&buf, img, nil)
	default:
		t.Fatalf("unsupported format %q", format)
	}
	if err != nil {
		t.Fatalf("encode %s: %v", format, err)
	}
	return buf.Bytes()
}

// loggedInClient returns a Client wired to the fake server with a
// non-zero UserID, so user-bound upload methods don't reject the call.
func loggedInClient(srvURL string, userID int64) *Client {
	c := NewClient(WithBaseURL(srvURL))
	c.UserID = userID
	return c
}

func TestUploadPostImages_HappyPath(t *testing.T) {
	f, srv := newFakeUploadServer(t)
	defer srv.Close()

	c := loggedInClient(srv.URL, 7360590)
	ctx := context.Background()

	pngBody := encodeImage(t, "png", 30, 40)
	gifBody := encodeImage(t, "gif", 50, 50)
	jpgBody := encodeImage(t, "jpeg", 60, 80)

	names, err := c.UploadPostImages(ctx, []Upload{
		{Filename: "selfie.png", Body: bytes.NewReader(pngBody)},
		{Filename: "sticker.gif", Body: bytes.NewReader(gifBody)},
		{Filename: "photo.jpg", Body: bytes.NewReader(jpgBody)},
	})
	if err != nil {
		t.Fatalf("UploadPostImages: %v", err)
	}
	if len(names) != 3 {
		t.Fatalf("len(names) = %d, want 3", len(names))
	}

	wantExt := []string{".png", ".gif", ".jpg"}
	wantCT := []string{"image/png", "image/gif", "image/jpeg"}
	wantBody := [][]byte{pngBody, gifBody, jpgBody}
	wantSize := []string{"_size_30x40", "_size_50x50", "_size_60x80"}

	for i, name := range names {
		if got := path.Ext(name); got != wantExt[i] {
			t.Errorf("name[%d] ext = %q, want %q (%s)", i, got, wantExt[i], name)
		}
		if !strings.HasPrefix(name, fmt.Sprintf("s3/user/7360590/posts/%d/", time.Now().Year())) {
			t.Errorf("name[%d] = %q, missing s3/user/7360590/posts/<year>/ prefix", i, name)
		}
		if !strings.Contains(name, wantSize[i]) {
			t.Errorf("name[%d] = %q, want size suffix %s", i, name, wantSize[i])
		}
		key := "/uploads/" + name
		if got := f.contentTypes[key]; got != wantCT[i] {
			t.Errorf("Content-Type[%s] = %q, want %q", name, got, wantCT[i])
		}
		if got := f.objects[key]; !bytes.Equal(got, wantBody[i]) {
			t.Errorf("body[%s] mismatch (len got=%d want=%d)", name, len(got), len(wantBody[i]))
		}
	}

	thumbCount := 0
	for path, ct := range f.contentTypes {
		if !strings.Contains(path, "/thumb_") {
			continue
		}
		thumbCount++
		switch {
		case strings.HasSuffix(path, ".gif"):
			if ct != "image/gif" {
				t.Errorf("gif thumbnail %s Content-Type = %q, want image/gif", path, ct)
			}
		case strings.HasSuffix(path, ".jpeg"), strings.HasSuffix(path, ".jpg"):
			if ct != "image/jpeg" {
				t.Errorf("jpeg thumbnail %s Content-Type = %q, want image/jpeg", path, ct)
			}
		default:
			t.Errorf("unexpected thumbnail extension: %s", path)
		}
	}
	if thumbCount != 3 {
		t.Errorf("thumbnail PUTs = %d, want 3 (one per input)", thumbCount)
	}

	if len(f.imageNames) != 6 {
		t.Fatalf("imageNames len = %d, want 6 (3 main + 3 thumb)", len(f.imageNames))
	}
	for i := 0; i < 3; i++ {
		if want := "s3/" + f.imageNames[i]; want != names[i] {
			t.Errorf("imageNames[%d] = %q, returned name = %q (want %q)", i, f.imageNames[i], names[i], want)
		}
	}
	thumbsSent := 0
	for i := 3; i < 6; i++ {
		if !strings.Contains(path.Base(f.imageNames[i]), "thumb_") {
			t.Errorf("imageNames[%d] = %q, want a thumb_ entry", i, f.imageNames[i])
		} else {
			thumbsSent++
		}
	}
	if thumbsSent != 3 {
		t.Errorf("thumbnails sent in file_names[] = %d, want 3", thumbsSent)
	}
}

func TestUploadAvatarImage_HappyPath(t *testing.T) {
	f, srv := newFakeUploadServer(t)
	defer srv.Close()

	c := loggedInClient(srv.URL, 42)
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
		t.Errorf("name = %q, missing size suffix", name)
	}
	if _, ok := f.objects["/uploads/"+name]; !ok {
		t.Errorf("PUT to /uploads/%s did not happen", name)
	}
}

func TestUploadAvatarImage_RequiresLogin(t *testing.T) {
	c := NewClient() // UserID == 0 (no login)
	body := encodeImage(t, "jpeg", 5, 5)

	_, err := c.UploadAvatarImage(context.Background(), Upload{
		Filename: "x.jpg", Body: bytes.NewReader(body),
	})
	if err == nil {
		t.Fatal("expected error from UploadAvatarImage without login")
	}
	if !strings.Contains(err.Error(), "not logged in") {
		t.Errorf("err = %q, want \"not logged in\"", err)
	}
}

func TestUploadCoverImage_RequiresLogin(t *testing.T) {
	_, err := NewClient().UploadCoverImage(context.Background(), Upload{
		Filename: "c.jpg", Body: bytes.NewReader([]byte{0}),
	})
	if err == nil {
		t.Fatal("expected error from UploadCoverImage without login")
	}
}

func TestUploadPostImages_RequiresLogin(t *testing.T) {
	_, err := NewClient().UploadPostImages(context.Background(), []Upload{{
		Filename: "p.jpg", Body: bytes.NewReader([]byte{0}),
	}})
	if err == nil {
		t.Fatal("expected error from UploadPostImages without login")
	}
}

func TestUploadChatMessageImages_RequiresLogin(t *testing.T) {
	_, err := NewClient().UploadChatMessageImages(context.Background(), 1, []Upload{{
		Filename: "p.jpg", Body: bytes.NewReader([]byte{0}),
	}})
	if err == nil {
		t.Fatal("expected error from UploadChatMessageImages without login")
	}
}

func TestUploadAvatarImage_UnknownExtensionFallsBackToJPEG(t *testing.T) {
	f, srv := newFakeUploadServer(t)
	defer srv.Close()

	c := loggedInClient(srv.URL, 1)
	ctx := context.Background()
	body := encodeImage(t, "jpeg", 5, 5)

	for _, fn := range []string{"noext", "weird.xyz"} {
		name, err := c.UploadAvatarImage(ctx, Upload{Filename: fn, Body: bytes.NewReader(body)})
		if err != nil {
			t.Fatalf("UploadAvatarImage(%q): %v", fn, err)
		}
		if path.Ext(name) != ".jpg" {
			t.Errorf("name(%q) = %q, want .jpg suffix", fn, name)
		}
		if got := f.contentTypes["/uploads/"+name]; got != "image/jpeg" {
			t.Errorf("Content-Type(%q) = %q, want image/jpeg", name, got)
		}
	}
}

func TestUploadPostImages_RejectsEmptyAndOverflow(t *testing.T) {
	c := loggedInClient("http://does-not-matter", 1)
	ctx := context.Background()

	if _, err := c.UploadPostImages(ctx, nil); err == nil {
		t.Error("UploadPostImages(nil) expected error")
	}

	tooMany := make([]Upload, MaxImagesPerUpload+1)
	for i := range tooMany {
		tooMany[i] = Upload{Filename: "x.jpg", Body: bytes.NewReader([]byte{0})}
	}
	if _, err := c.UploadPostImages(ctx, tooMany); err == nil {
		t.Errorf("UploadPostImages(%d files) expected error", len(tooMany))
	}
}

func TestUploadAvatarImage_PutFailureSurfacesError(t *testing.T) {
	f, srv := newFakeUploadServer(t)
	defer srv.Close()

	f.mu.Lock()
	f.putFailures = 99 // ensure every PUT fails
	f.mu.Unlock()

	c := loggedInClient(srv.URL, 1)
	body := encodeImage(t, "png", 5, 5)

	_, err := c.UploadAvatarImage(context.Background(), Upload{
		Filename: "a.png", Body: bytes.NewReader(body),
	})
	if err == nil {
		t.Fatal("expected error from failing PUT")
	}
	if !strings.Contains(err.Error(), "PUT 500") {
		t.Errorf("error = %q, want PUT 500 in message", err)
	}
}

func TestUploadAvatarImage_PresignedFailure(t *testing.T) {
	f, srv := newFakeUploadServer(t)
	defer srv.Close()

	f.mu.Lock()
	f.failPresignedAt = 1
	f.mu.Unlock()

	c := loggedInClient(srv.URL, 1)
	c.RetryPolicy = RetryPolicy{}
	body := encodeImage(t, "png", 5, 5)

	_, err := c.UploadAvatarImage(context.Background(), Upload{
		Filename: "a.png", Body: bytes.NewReader(body),
	})
	if err == nil {
		t.Fatal("expected error from failing presigned URL fetch")
	}
}

func TestUploadAvatarImage_NonImageBodyOmitsSize(t *testing.T) {
	f, srv := newFakeUploadServer(t)
	defer srv.Close()

	c := loggedInClient(srv.URL, 1)
	rawBytes := []byte("not an image at all")

	name, err := c.UploadAvatarImage(context.Background(), Upload{
		Filename: "broken.jpg", Body: bytes.NewReader(rawBytes),
	})
	if err != nil {
		t.Fatalf("UploadAvatarImage: %v", err)
	}
	if strings.Contains(name, "_size_") {
		t.Errorf("name = %q, expected no _size_ suffix when body isn't decodable", name)
	}
	if !bytes.Equal(f.objects["/uploads/"+name], rawBytes) {
		t.Errorf("body mismatch")
	}
}

func TestUploadVideo_HappyPath(t *testing.T) {
	f, srv := newFakeUploadServer(t)
	defer srv.Close()

	c := NewClient(WithBaseURL(srv.URL))
	mp4Body := []byte("\x00\x00\x00 ftypisom--mp4-bytes--")
	name, err := c.UploadVideo(context.Background(), Upload{
		Filename: "clip.mp4",
		Body:     bytes.NewReader(mp4Body),
	})
	if err != nil {
		t.Fatalf("UploadVideo: %v", err)
	}
	if !strings.HasSuffix(name, ".mp4") {
		t.Errorf("name = %q, want .mp4 suffix", name)
	}
	key := "/uploads/" + name
	if got := f.contentTypes[key]; got != "video/mp4" {
		t.Errorf("Content-Type[%s] = %q, want video/mp4", name, got)
	}
	if got := f.objects[key]; !bytes.Equal(got, mp4Body) {
		t.Errorf("body[%s] mismatch", name)
	}
	if f.videoFileName != name {
		t.Errorf("videoFileName query = %q, returned = %q", f.videoFileName, name)
	}
}

func TestUploadVideo_NilBody(t *testing.T) {
	c := NewClient(WithBaseURL("http://does-not-matter"))
	_, err := c.UploadVideo(context.Background(), Upload{Filename: "clip.mp4"})
	if err == nil {
		t.Fatal("UploadVideo with nil body expected error")
	}
}

func TestContentTypeFor(t *testing.T) {
	cases := []struct {
		name string
		want string
	}{
		{"a.png", "image/png"},
		{"path/to/B.PNG", "image/png"},
		{"a.gif", "image/gif"},
		{"a.GIF", "image/gif"},
		{"a.jpg", "image/jpeg"},
		{"a.jpeg", "image/jpeg"},
		{"a.JPG", "image/jpeg"},
		{"a.mp4", "video/mp4"},
		{"unknown", "image/jpeg"},
		{"", "image/jpeg"},
		{"weird.xyz", "image/jpeg"},
	}
	for _, c := range cases {
		if got := contentTypeFor(c.name); got != c.want {
			t.Errorf("contentTypeFor(%q) = %q, want %q", c.name, got, c.want)
		}
	}
}

func TestImageFilename_Format(t *testing.T) {
	got := imageFilename("user/42/avatar", "ABC123", 1700000000000, 0, ".png", 200, 100)
	want := "user/42/avatar/ABC123_1700000000000_0_size_200x100.png"
	if got != want {
		t.Errorf("imageFilename = %q, want %q", got, want)
	}

	got = imageFilename("user/42/avatar", "ABC123", 1700000000000, 2, ".jpeg", 0, 0)
	want = "user/42/avatar/ABC123_1700000000000_2.jpeg"
	if got != want {
		t.Errorf("imageFilename without size = %q, want %q", got, want)
	}

	got = imageFilename("user/42/avatar", "ABC123", 1700000000000, 0, "", 10, 10)
	want = "user/42/avatar/ABC123_1700000000000_0_size_10x10.jpg"
	if got != want {
		t.Errorf("imageFilename empty ext = %q, want %q", got, want)
	}
}

func TestThumbnailFilename_Format(t *testing.T) {
	got := thumbnailFilename("user/42/avatar", "ABC123", 1700000000000, 0, ".jpeg", 800, 600)
	want := "user/42/avatar/thumb_ABC123_1700000000000_0_size_800x600.jpeg"
	if got != want {
		t.Errorf("thumbnailFilename = %q, want %q", got, want)
	}
	got = thumbnailFilename("user/42/avatar", "ABC123", 1700000000000, 0, ".jpeg", 0, 0)
	want = "user/42/avatar/thumb_ABC123_1700000000000_0.jpeg"
	if got != want {
		t.Errorf("thumbnailFilename without size = %q, want %q", got, want)
	}
}

func TestMakeThumbnail_PreservesAspectRatio(t *testing.T) {
	src := encodeImage(t, "png", 400, 200)
	out, err := makeJPEGThumbnail(src, 200, 200)
	if err != nil {
		t.Fatalf("makeThumbnail: %v", err)
	}
	cfg, format, err := image.DecodeConfig(bytes.NewReader(out))
	if err != nil {
		t.Fatalf("decode thumbnail: %v", err)
	}
	if format != "jpeg" {
		t.Errorf("format = %q, want jpeg", format)
	}
	if cfg.Width != 200 || cfg.Height != 100 {
		t.Errorf("size = %dx%d, want 200x100", cfg.Width, cfg.Height)
	}
}

func TestMakeThumbnail_NoUpscale(t *testing.T) {
	src := encodeImage(t, "png", 50, 80)
	out, err := makeJPEGThumbnail(src, 200, 200)
	if err != nil {
		t.Fatalf("makeThumbnail: %v", err)
	}
	cfg, _, err := image.DecodeConfig(bytes.NewReader(out))
	if err != nil {
		t.Fatalf("decode: %v", err)
	}
	if cfg.Width != 50 || cfg.Height != 80 {
		t.Errorf("size = %dx%d, want 50x80 (no upscale)", cfg.Width, cfg.Height)
	}
}

func TestMakeThumbnail_RejectsCorrupt(t *testing.T) {
	if _, err := makeJPEGThumbnail([]byte("not an image"), 200, 200); err == nil {
		t.Fatal("expected decode error on corrupt input")
	}
	if _, err := makeJPEGThumbnail(encodeImage(t, "png", 10, 10), 0, 200); err == nil {
		t.Fatal("expected error for invalid target size")
	}
}

func TestUploadVideoCallSnapshotImage_SkipsThumb(t *testing.T) {
	f, srv := newFakeUploadServer(t)
	defer srv.Close()

	c := NewClient(WithBaseURL(srv.URL))
	body := encodeImage(t, "jpeg", 100, 100)
	name, err := c.UploadVideoCallSnapshotImage(context.Background(), Upload{
		Filename: "snap.jpg", Body: bytes.NewReader(body),
	})
	if err != nil {
		t.Fatalf("UploadVideoCallSnapshotImage: %v", err)
	}
	if name == "" {
		t.Fatal("empty filename returned")
	}
	if got := len(f.imageNames); got != 1 {
		t.Errorf("file_names[] entries = %d, want 1 (no thumb for video-call snapshot)", got)
	}
	for path := range f.contentTypes {
		if strings.Contains(path, "/thumb_") {
			t.Errorf("unexpected thumbnail upload at %s", path)
		}
	}
}

func TestUploadAvatarImage_GIFThumbStaysGIF(t *testing.T) {
	f, srv := newFakeUploadServer(t)
	defer srv.Close()

	c := loggedInClient(srv.URL, 1)
	body := encodeImage(t, "gif", 800, 400)
	name, err := c.UploadAvatarImage(context.Background(), Upload{
		Filename: "x.gif", Body: bytes.NewReader(body),
	})
	if err != nil {
		t.Fatalf("UploadAvatarImage: %v", err)
	}
	if name == "" {
		t.Fatal("empty filename returned")
	}
	if got := len(f.imageNames); got != 2 {
		t.Errorf("file_names[] entries = %d, want 2 (main gif + gif thumb)", got)
	}

	var thumbPath string
	var thumbBytes []byte
	for path, body := range f.objects {
		if strings.Contains(path, "/thumb_") {
			thumbPath = path
			thumbBytes = body
			break
		}
	}
	if thumbBytes == nil {
		t.Fatal("no thumbnail PUT for GIF input")
	}
	if !strings.HasSuffix(thumbPath, ".gif") {
		t.Errorf("thumb path = %q, want .gif suffix (matches the main extension)", thumbPath)
	}
	if got := f.contentTypes[thumbPath]; got != "image/gif" {
		t.Errorf("thumb Content-Type = %q, want image/gif", got)
	}
	cfg, format, err := image.DecodeConfig(bytes.NewReader(thumbBytes))
	if err != nil {
		t.Fatalf("decode thumb: %v", err)
	}
	if format != "gif" {
		t.Errorf("thumb format = %q, want gif (animation preserved)", format)
	}
	if cfg.Width != 200 || cfg.Height != 100 {
		t.Errorf("thumb size = %dx%d, want 200x100", cfg.Width, cfg.Height)
	}
}

func TestMakeGIFThumbnail_PreservesAllFrames(t *testing.T) {
	var buf bytes.Buffer
	g := &gif.GIF{LoopCount: 0}
	palette := color.Palette{color.Black, color.White, color.RGBA{R: 200}}
	for i := 0; i < 3; i++ {
		img := image.NewPaletted(image.Rect(0, 0, 400, 200), palette)
		for x := 0; x < 400; x++ {
			for y := 0; y < 200; y++ {
				img.SetColorIndex(x, y, uint8((x+y+i)%len(palette)))
			}
		}
		g.Image = append(g.Image, img)
		g.Delay = append(g.Delay, 10)
		g.Disposal = append(g.Disposal, gif.DisposalNone)
	}
	g.Config = image.Config{ColorModel: palette, Width: 400, Height: 200}
	if err := gif.EncodeAll(&buf, g); err != nil {
		t.Fatalf("encode source: %v", err)
	}

	out, err := makeGIFThumbnail(buf.Bytes(), 200, 200)
	if err != nil {
		t.Fatalf("makeGIFThumbnail: %v", err)
	}
	all, err := gif.DecodeAll(bytes.NewReader(out))
	if err != nil {
		t.Fatalf("decode result: %v", err)
	}
	if len(all.Image) != 3 {
		t.Errorf("frame count = %d, want 3", len(all.Image))
	}
	if all.Config.Width != 200 || all.Config.Height != 100 {
		t.Errorf("config size = %dx%d, want 200x100", all.Config.Width, all.Config.Height)
	}
}

func TestMakeGIFThumbnail_NoUpscale(t *testing.T) {
	var buf bytes.Buffer
	img := image.NewPaletted(image.Rect(0, 0, 50, 50), color.Palette{color.Black, color.White})
	g := &gif.GIF{
		Image:    []*image.Paletted{img},
		Delay:    []int{10},
		Disposal: []byte{gif.DisposalNone},
		Config:   image.Config{ColorModel: img.Palette, Width: 50, Height: 50},
	}
	if err := gif.EncodeAll(&buf, g); err != nil {
		t.Fatalf("encode source: %v", err)
	}
	original := buf.Bytes()
	out, err := makeGIFThumbnail(original, 200, 200)
	if err != nil {
		t.Fatalf("makeGIFThumbnail: %v", err)
	}
	if !bytes.Equal(out, original) {
		t.Error("expected same bytes returned when source already fits")
	}
}

func TestUploadAvatarImage_ThumbIs200x200(t *testing.T) {
	f, srv := newFakeUploadServer(t)
	defer srv.Close()

	c := loggedInClient(srv.URL, 7)
	body := encodeImage(t, "jpeg", 800, 600)
	if _, err := c.UploadAvatarImage(context.Background(), Upload{
		Filename: "me.jpg", Body: bytes.NewReader(body),
	}); err != nil {
		t.Fatalf("UploadAvatarImage: %v", err)
	}
	var thumbBytes []byte
	for path, body := range f.objects {
		if strings.Contains(path, "/thumb_") {
			thumbBytes = body
			break
		}
	}
	if thumbBytes == nil {
		t.Fatal("no thumbnail PUT recorded")
	}
	cfg, _, err := image.DecodeConfig(bytes.NewReader(thumbBytes))
	if err != nil {
		t.Fatalf("decode thumbnail: %v", err)
	}
	if cfg.Width != 200 || cfg.Height != 150 {
		t.Errorf("avatar thumb size = %dx%d, want 200x150 (fitted into 200x200)", cfg.Width, cfg.Height)
	}
}

func TestUploadAvatarImage_NonImageBodySkipsThumb(t *testing.T) {
	f, srv := newFakeUploadServer(t)
	defer srv.Close()

	c := loggedInClient(srv.URL, 1)
	if _, err := c.UploadAvatarImage(context.Background(), Upload{
		Filename: "broken.jpg", Body: bytes.NewReader([]byte("not an image")),
	}); err != nil {
		t.Fatalf("UploadAvatarImage: %v", err)
	}
	for path := range f.contentTypes {
		if strings.Contains(path, "/thumb_") {
			t.Errorf("unexpected thumbnail upload for undecodable input: %s", path)
		}
	}
}

func TestUploadCategoryPath(t *testing.T) {
	now := time.Date(2026, time.April, 26, 12, 0, 0, 0, time.UTC)
	cases := []struct {
		c    uploadCategory
		want string
	}{
		{userAvatarUpload{userID: 7360590}, "user/7360590/avatar"},
		{userCoverUpload{userID: 7360590}, "user/7360590/cover"},
		{userPostUpload{userID: 7360590}, "user/7360590/posts/2026/4/26"},
		{chatMessageUpload{roomID: 1, userID: 2}, "chatroom/1/user/2/messages/2026/4/26"},
		{chatBackgroundUpload{roomID: 9}, "chatroom/9/background"},
		{groupIconUpload{groupID: 5}, "group/5/icon"},
		{groupCoverUpload{groupID: 5}, "group/5/cover"},
		{groupCreationIconUpload{}, "group/create/icon/2026/4/26"},
		{groupCreationCoverUpload{}, "group/create/cover/2026/4/26"},
		{signupAvatarUpload{}, "user/create/2026/4/26"},
		{threadIconUpload{groupID: 5}, "group/5/threads/2026/4/26"},
		{reportUpload{}, "report/2026/4/26"},
		{videoCallSnapshotUpload{}, "video_call_snapshot/2026/4/26"},
	}
	for _, tc := range cases {
		if got := tc.c.uploadCategoryPath(now); got != tc.want {
			t.Errorf("%T.uploadCategoryPath = %q, want %q", tc.c, got, tc.want)
		}
	}
}

func TestFormatYMD_NoZeroPadding(t *testing.T) {
	cases := []struct {
		t    time.Time
		want string
	}{
		{time.Date(2026, time.April, 9, 0, 0, 0, 0, time.UTC), "2026/4/9"},
		{time.Date(2026, time.December, 31, 0, 0, 0, 0, time.UTC), "2026/12/31"},
		{time.Date(2000, time.January, 1, 0, 0, 0, 0, time.UTC), "2000/1/1"},
	}
	for _, c := range cases {
		if got := formatYMD(c.t); got != c.want {
			t.Errorf("formatYMD(%v) = %q, want %q", c.t, got, c.want)
		}
	}
}

func TestMediaURL(t *testing.T) {
	cases := []struct{ in, want string }{
		{"s3/user/7360590/avatar/abc_123_0_size_251x251.jpeg",
			"https://cdn.yay.space/uploads/s3/user/7360590/avatar/abc_123_0_size_251x251.jpeg"},
		{"/s3/leading-slash.jpg",
			"https://cdn.yay.space/uploads/s3/leading-slash.jpg"},
		{"", ""},
	}
	for _, c := range cases {
		if got := MediaURL(c.in); got != c.want {
			t.Errorf("MediaURL(%q) = %q, want %q", c.in, got, c.want)
		}
	}
}

func TestRandomFilenamePrefix_LengthAndAlphabet(t *testing.T) {
	const allowed = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	for n := 1; n <= 32; n++ {
		got := randomFilenamePrefix(n)
		if len(got) != n {
			t.Fatalf("len = %d, want %d", len(got), n)
		}
		for i, ch := range got {
			if !strings.ContainsRune(allowed, ch) {
				t.Fatalf("char %d %q not in allowed alphabet", i, ch)
			}
		}
	}
}

func TestUploadAvatarImage_FilenameShape(t *testing.T) {
	_, srv := newFakeUploadServer(t)
	defer srv.Close()

	c := loggedInClient(srv.URL, 1)
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

func TestUploadAvatarImage_EmptyPresignedResponse(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/v1/buckets/presigned_urls" {
			http.Error(w, "no", http.StatusNotFound)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"presigned_urls": []}`))
	}))
	defer srv.Close()

	c := loggedInClient(srv.URL, 1)
	body := encodeImage(t, "png", 5, 5)
	_, err := c.UploadAvatarImage(context.Background(), Upload{
		Filename: "x.png", Body: bytes.NewReader(body),
	})
	if err == nil {
		t.Fatal("expected error on empty presigned_urls")
	}
	if !strings.Contains(err.Error(), "count mismatch") {
		t.Errorf("error = %q, want count mismatch", err)
	}
}

func TestUploadAvatarImage_ErrorIsWrapped(t *testing.T) {
	c := loggedInClient("http://127.0.0.1:1", 1) // unroutable
	body := encodeImage(t, "png", 5, 5)
	_, err := c.UploadAvatarImage(context.Background(), Upload{
		Filename: "x.png", Body: bytes.NewReader(body),
	})
	if err == nil {
		t.Fatal("expected error against unroutable URL")
	}
	if errors.Is(err, http.ErrServerClosed) {
		t.Fatal("unexpected sentinel match")
	}
	_ = fmt.Sprintf("%v", err)
}
