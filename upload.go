package yaylib

import (
	"bytes"
	"context"
	"crypto/rand"
	"fmt"
	"image"
	"image/gif"
	"image/jpeg"
	_ "image/png"
	"io"
	"net/http"
	"path/filepath"
	"strings"
	"sync"
	"time"

	"golang.org/x/image/draw"
)

// MaxImagesPerUpload is the maximum number of images accepted in a
// single multi-image upload call (UploadPostImages, UploadChatMessageImages,
// UploadReportImages).
const MaxImagesPerUpload = 9

// MediaCDNBase is the public CDN prefix that serves files uploaded
// via the Upload* image / video methods. The Yay! API sometimes
// returns stale or differently-prefixed URLs in fields like
// profile_icon, so callers should treat the upload-returned filename
// as the source of truth and call MediaURL to derive the live URL.
const MediaCDNBase = "https://cdn.yay.space/uploads"

// MediaURL returns the public CDN URL for a filename returned by any
// Upload* method. Use this rather than reading profile_icon /
// cover_image / attachment URLs from API responses, which can carry a
// legacy prefix that no longer resolves.
func MediaURL(filename string) string {
	if filename == "" {
		return ""
	}
	return MediaCDNBase + "/" + strings.TrimLeft(filename, "/")
}

// Upload is a single file passed to one of the Upload* methods.
//
// Filename is used to derive the file extension and the Content-Type
// sent on PUT. If empty, the upload is treated as JPEG. Body is read
// once and fully buffered in memory before the request is issued.
type Upload struct {
	Filename string
	Body     io.Reader
}

// uploadCategory selects the bucket-side path prefix and thumbnail
// dimensions for a given upload kind. Concrete values live below as
// unexported types — callers reach them through the typed
// UploadXxxImage[s] methods on Client, which know which category they
// correspond to and fill in any required IDs.
type uploadCategory interface {
	// uploadCategoryPath returns the category-specific path prefix.
	// now is the timestamp uploadImages used for the filename body so
	// date-bucketed categories share a single timestamp across the
	// whole call.
	uploadCategoryPath(now time.Time) string

	// thumbnailSize returns the maximum dimensions of the thumbnail
	// that should be uploaded alongside each main image. ok=false
	// means the category never produces a thumbnail (e.g. video-call
	// snapshots).
	//
	// The thumbnail fits inside the returned box while preserving the
	// source aspect ratio. GIF sources stay animated; only their
	// thumbnail is resized.
	thumbnailSize() (w, h int, ok bool)
}

type userAvatarUpload struct{ userID int64 }

func (c userAvatarUpload) uploadCategoryPath(_ time.Time) string {
	return fmt.Sprintf("user/%d/avatar", c.userID)
}
func (userAvatarUpload) thumbnailSize() (int, int, bool) { return 200, 200, true }

type userCoverUpload struct{ userID int64 }

func (c userCoverUpload) uploadCategoryPath(_ time.Time) string {
	return fmt.Sprintf("user/%d/cover", c.userID)
}
func (userCoverUpload) thumbnailSize() (int, int, bool) { return 300, 150, true }

type userPostUpload struct{ userID int64 }

func (c userPostUpload) uploadCategoryPath(now time.Time) string {
	return fmt.Sprintf("user/%d/posts/%s", c.userID, formatYMD(now))
}
func (userPostUpload) thumbnailSize() (int, int, bool) { return 450, 450, true }

type chatMessageUpload struct {
	roomID int64
	userID int64
}

func (c chatMessageUpload) uploadCategoryPath(now time.Time) string {
	return fmt.Sprintf("chatroom/%d/user/%d/messages/%s", c.roomID, c.userID, formatYMD(now))
}
func (chatMessageUpload) thumbnailSize() (int, int, bool) { return 450, 450, true }

type chatBackgroundUpload struct{ roomID int64 }

func (c chatBackgroundUpload) uploadCategoryPath(_ time.Time) string {
	return fmt.Sprintf("chatroom/%d/background", c.roomID)
}
func (chatBackgroundUpload) thumbnailSize() (int, int, bool) { return 200, 200, true }

type groupIconUpload struct{ groupID int64 }

func (c groupIconUpload) uploadCategoryPath(_ time.Time) string {
	return fmt.Sprintf("group/%d/icon", c.groupID)
}
func (groupIconUpload) thumbnailSize() (int, int, bool) { return 300, 300, true }

type groupCoverUpload struct{ groupID int64 }

func (c groupCoverUpload) uploadCategoryPath(_ time.Time) string {
	return fmt.Sprintf("group/%d/cover", c.groupID)
}
func (groupCoverUpload) thumbnailSize() (int, int, bool) { return 300, 300, true }

type groupCreationIconUpload struct{}

func (groupCreationIconUpload) uploadCategoryPath(now time.Time) string {
	return "group/create/icon/" + formatYMD(now)
}
func (groupCreationIconUpload) thumbnailSize() (int, int, bool) { return 300, 300, true }

type groupCreationCoverUpload struct{}

func (groupCreationCoverUpload) uploadCategoryPath(now time.Time) string {
	return "group/create/cover/" + formatYMD(now)
}
func (groupCreationCoverUpload) thumbnailSize() (int, int, bool) { return 300, 300, true }

type signupAvatarUpload struct{}

func (signupAvatarUpload) uploadCategoryPath(now time.Time) string {
	return "user/create/" + formatYMD(now)
}
func (signupAvatarUpload) thumbnailSize() (int, int, bool) { return 200, 200, true }

type threadIconUpload struct{ groupID int64 }

func (c threadIconUpload) uploadCategoryPath(now time.Time) string {
	return fmt.Sprintf("group/%d/threads/%s", c.groupID, formatYMD(now))
}
func (threadIconUpload) thumbnailSize() (int, int, bool) { return 300, 300, true }

type reportUpload struct{}

func (reportUpload) uploadCategoryPath(now time.Time) string {
	return "report/" + formatYMD(now)
}
func (reportUpload) thumbnailSize() (int, int, bool) { return 450, 450, true }

type videoCallSnapshotUpload struct{}

func (videoCallSnapshotUpload) uploadCategoryPath(now time.Time) string {
	return "video_call_snapshot/" + formatYMD(now)
}
func (videoCallSnapshotUpload) thumbnailSize() (int, int, bool) { return 0, 0, false }

// requireUserID returns Client.UserID if it has been populated by a
// successful login (or WithUserID), otherwise an error suitable to
// surface from a user-bound upload method.
func (c *Client) requireUserID() (int64, error) {
	if c.UserID == 0 {
		return 0, fmt.Errorf("yaylib: not logged in (call LoginWithEmail before user-bound uploads)")
	}
	return c.UserID, nil
}

// UploadAvatarImage uploads the authenticated user's avatar image and
// returns the server-canonical filename to pass to EditUser as
// profile_icon_filename. JPEG, PNG, and GIF (animated) sources are
// supported.
func (c *Client) UploadAvatarImage(ctx context.Context, file Upload) (string, error) {
	uid, err := c.requireUserID()
	if err != nil {
		return "", err
	}
	return c.uploadSingleImage(ctx, userAvatarUpload{userID: uid}, file)
}

// UploadCoverImage uploads the authenticated user's profile cover
// image and returns the server-canonical filename to pass to EditUser
// as cover_image_filename.
func (c *Client) UploadCoverImage(ctx context.Context, file Upload) (string, error) {
	uid, err := c.requireUserID()
	if err != nil {
		return "", err
	}
	return c.uploadSingleImage(ctx, userCoverUpload{userID: uid}, file)
}

// UploadPostImages uploads up to MaxImagesPerUpload images for a new
// post by the authenticated user and returns the server-canonical
// filenames to pass to CreatePost as attachment_filename /
// attachment_2_filename / ….
func (c *Client) UploadPostImages(ctx context.Context, files []Upload) ([]string, error) {
	uid, err := c.requireUserID()
	if err != nil {
		return nil, err
	}
	return c.uploadImages(ctx, userPostUpload{userID: uid}, files)
}

// UploadChatMessageImages uploads up to MaxImagesPerUpload images for
// a chat message in the given room and returns the server-canonical
// filenames.
func (c *Client) UploadChatMessageImages(ctx context.Context, roomID int64, files []Upload) ([]string, error) {
	uid, err := c.requireUserID()
	if err != nil {
		return nil, err
	}
	return c.uploadImages(ctx, chatMessageUpload{roomID: roomID, userID: uid}, files)
}

// UploadChatBackgroundImage uploads the background image of the chat
// room with the given ID and returns the server-canonical filename.
func (c *Client) UploadChatBackgroundImage(ctx context.Context, roomID int64, file Upload) (string, error) {
	return c.uploadSingleImage(ctx, chatBackgroundUpload{roomID: roomID}, file)
}

// UploadGroupIconImage uploads the icon of the group with the given
// ID and returns the server-canonical filename.
func (c *Client) UploadGroupIconImage(ctx context.Context, groupID int64, file Upload) (string, error) {
	return c.uploadSingleImage(ctx, groupIconUpload{groupID: groupID}, file)
}

// UploadGroupCoverImage uploads the cover image of the group with the
// given ID and returns the server-canonical filename.
func (c *Client) UploadGroupCoverImage(ctx context.Context, groupID int64, file Upload) (string, error) {
	return c.uploadSingleImage(ctx, groupCoverUpload{groupID: groupID}, file)
}

// UploadThreadIconImage uploads the icon of a thread under the given
// group and returns the server-canonical filename.
func (c *Client) UploadThreadIconImage(ctx context.Context, groupID int64, file Upload) (string, error) {
	return c.uploadSingleImage(ctx, threadIconUpload{groupID: groupID}, file)
}

// UploadSignupAvatarImage uploads an avatar image during the new-user
// signup flow, before the account exists. The returned filename is
// passed to CreateUser as profile_icon_filename.
func (c *Client) UploadSignupAvatarImage(ctx context.Context, file Upload) (string, error) {
	return c.uploadSingleImage(ctx, signupAvatarUpload{}, file)
}

// UploadGroupCreationIconImage uploads the icon image for a group
// that is about to be created, before the group ID exists.
func (c *Client) UploadGroupCreationIconImage(ctx context.Context, file Upload) (string, error) {
	return c.uploadSingleImage(ctx, groupCreationIconUpload{}, file)
}

// UploadGroupCreationCoverImage uploads the cover image for a group
// that is about to be created.
func (c *Client) UploadGroupCreationCoverImage(ctx context.Context, file Upload) (string, error) {
	return c.uploadSingleImage(ctx, groupCreationCoverUpload{}, file)
}

// UploadReportImages uploads screenshots attached to a moderation
// report. The Yay! report endpoint accepts up to four images.
func (c *Client) UploadReportImages(ctx context.Context, files []Upload) ([]string, error) {
	return c.uploadImages(ctx, reportUpload{}, files)
}

// UploadVideoCallSnapshotImage uploads a snapshot taken from a video
// call. No thumbnail is generated for this category.
func (c *Client) UploadVideoCallSnapshotImage(ctx context.Context, file Upload) (string, error) {
	return c.uploadSingleImage(ctx, videoCallSnapshotUpload{}, file)
}

// uploadSingleImage is a thin wrapper for the single-image methods.
func (c *Client) uploadSingleImage(ctx context.Context, category uploadCategory, file Upload) (string, error) {
	names, err := c.uploadImages(ctx, category, []Upload{file})
	if err != nil {
		return "", err
	}
	return names[0], nil
}

// uploadImages performs the full presigned-URL + parallel-PUT dance
// shared by every image-upload method. For each main image (JPEG,
// PNG, or animated GIF) it generates a properly-sized thumbnail —
// matching the source extension and Content-Type — and uploads both
// to the per-category bucket path. Categories whose thumbnailSize
// returns ok=false skip the thumbnail step entirely.
//
// Filenames take the form
//
//	<category-path>/<random16>_<unix_millis>_<index>_size_<W>x<H>.<ext>
//
// with the same shape for the thumbnail except for a leading
// "thumb_" on the file body. The size suffix carries the original
// dimensions, not the thumbnail's, matching the Yay! server's
// expectations.
func (c *Client) uploadImages(ctx context.Context, category uploadCategory, files []Upload) ([]string, error) {
	if category == nil {
		return nil, fmt.Errorf("yaylib: upload requires a category")
	}
	if len(files) == 0 {
		return nil, fmt.Errorf("yaylib: upload requires at least one file")
	}
	if len(files) > MaxImagesPerUpload {
		return nil, fmt.Errorf("yaylib: UploadImages accepts at most %d files (got %d)", MaxImagesPerUpload, len(files))
	}

	now := time.Now()
	prefix := randomFilenamePrefix(16)
	ts := now.UnixMilli()
	categoryPath := category.uploadCategoryPath(now)
	thumbW, thumbH, hasThumb := category.thumbnailSize()

	type job struct {
		body        []byte
		filename    string
		contentType string
	}
	mainJobs := make([]job, len(files))
	var thumbJobs []job

	for i, f := range files {
		body, err := readFully(f.Body)
		if err != nil {
			return nil, fmt.Errorf("yaylib: read file %d: %w", i, err)
		}
		ext := normalizeImageExt(f.Filename)
		// A failed decode means we can't size the filename correctly,
		// and uploading body bytes whose extension doesn't match the
		// payload would just trigger server-side moderation deletion.
		// Surface the error instead of silently uploading a `_size_0x0`
		// file or, worse, a main image without its required thumbnail.
		w, h, err := decodeImageSize(body)
		if err != nil {
			return nil, fmt.Errorf("yaylib: decode image %d: %w", i, err)
		}
		mainName := imageFilename(categoryPath, prefix, ts, i, ext, w, h)
		mainJobs[i] = job{body: body, filename: mainName, contentType: contentTypeFor(mainName)}

		if hasThumb {
			thumbBytes, thumbExt, thumbCT, err := makeThumbnailFor(body, ext, thumbW, thumbH)
			if err != nil {
				// The server deletes the main asset if its thumbnail
				// is missing or has a mismatched extension, so a
				// silently-skipped thumbnail effectively breaks the
				// whole upload — fail loudly instead.
				return nil, fmt.Errorf("yaylib: make thumbnail for image %d: %w", i, err)
			}
			thumbName := thumbnailFilename(categoryPath, prefix, ts, i, thumbExt, w, h)
			thumbJobs = append(thumbJobs, job{
				body:        thumbBytes,
				filename:    thumbName,
				contentType: thumbCT,
			})
		}
	}

	allJobs := append(mainJobs, thumbJobs...)
	allFileNames := make([]string, len(allJobs))
	for i, j := range allJobs {
		allFileNames[i] = j.filename
	}

	resp, _, err := c.GetBucketPresignedUrls(ctx).
		FileNames(allFileNames).
		Execute()
	if err != nil {
		return nil, fmt.Errorf("yaylib: get presigned urls: %w", err)
	}
	urls := resp.GetPresignedUrls()
	if len(urls) != len(allJobs) {
		return nil, fmt.Errorf("yaylib: presigned urls count mismatch (want %d got %d)", len(allJobs), len(urls))
	}

	hc := c.presignedHTTPClient()
	results := make([]error, len(allJobs))
	var wg sync.WaitGroup
	for i := range allJobs {
		wg.Add(1)
		go func() {
			defer wg.Done()
			rawURL := urls[i].GetUrl()
			if rawURL == "" {
				results[i] = fmt.Errorf("empty presigned url")
				return
			}
			results[i] = putToPresignedURL(ctx, hc, rawURL, allJobs[i].body, allJobs[i].contentType)
		}()
	}
	wg.Wait()
	for i, err := range results {
		if err != nil {
			return nil, fmt.Errorf("yaylib: upload %s: %w", allJobs[i].filename, err)
		}
	}

	// The Yay! servers normalize each filename (typically by adding an
	// "s3/" prefix that marks it as a freshly-uploaded object). That
	// canonical value is the one EditUser, CreatePost, and friends
	// expect back in profile_icon_filename / attachment_filename — the
	// raw client-side path won't be recognized as a pending upload.
	// Only the main-image canonical names are returned; the thumbnails
	// live at sibling paths the server resolves automatically.
	out := make([]string, len(files))
	for i := range files {
		out[i] = urls[i].GetFilename()
	}
	return out, nil
}

// UploadVideo uploads a single MP4 file and returns the server-side
// filename to pass back as video_file_name when creating a post or
// chat message.
//
// Unlike images, the video presigned-URL endpoint does not return a
// canonical filename, so the raw filename passed to the server is
// what callers should hand back to CreatePost.
func (c *Client) UploadVideo(ctx context.Context, file Upload) (string, error) {
	body, err := readFully(file.Body)
	if err != nil {
		return "", fmt.Errorf("yaylib: read video: %w", err)
	}
	name := videoFilename(randomFilenamePrefix(16), time.Now().UnixMilli())

	resp, _, err := c.GetUserPresignedUrl(ctx).
		VideoFileName(name).
		Execute()
	if err != nil {
		return "", fmt.Errorf("yaylib: get presigned url: %w", err)
	}
	rawURL := resp.GetPresignedUrl()
	if rawURL == "" {
		return "", fmt.Errorf("yaylib: empty presigned url")
	}
	if err := putToPresignedURL(ctx, c.presignedHTTPClient(), rawURL, body, "video/mp4"); err != nil {
		return "", fmt.Errorf("yaylib: upload video: %w", err)
	}
	return name, nil
}

// presignedHTTPClient returns an *http.Client suitable for S3 presigned
// PUTs. It strips the yaylib *Transport (which would otherwise inject
// Yay! headers and `Authorization: Bearer …` that conflict with the
// presigned URL's AWS signature scheme), but inherits the underlying
// base transport — so proxy / TLS / dialer customizations from
// WithHTTPClient still apply — along with the configured timeout.
func (c *Client) presignedHTTPClient() *http.Client {
	var base http.RoundTripper = http.DefaultTransport
	if t, ok := c.httpClient.Transport.(*Transport); ok && t.Base != nil {
		base = t.Base
	}
	return &http.Client{
		Transport: base,
		Timeout:   c.httpClient.Timeout,
	}
}

func putToPresignedURL(ctx context.Context, hc *http.Client, rawURL string, body []byte, contentType string) error {
	req, err := http.NewRequestWithContext(ctx, http.MethodPut, rawURL, bytes.NewReader(body))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", contentType)
	req.ContentLength = int64(len(body))

	resp, err := hc.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode/100 != 2 {
		preview, _ := io.ReadAll(io.LimitReader(resp.Body, 512))
		return fmt.Errorf("PUT %d: %s", resp.StatusCode, strings.TrimSpace(string(preview)))
	}
	return nil
}

func readFully(r io.Reader) ([]byte, error) {
	if r == nil {
		return nil, fmt.Errorf("body is nil")
	}
	return io.ReadAll(r)
}

func extOf(filename string) string {
	return strings.ToLower(filepath.Ext(filename))
}

func contentTypeFor(filename string) string {
	switch extOf(filename) {
	case ".png":
		return "image/png"
	case ".gif":
		return "image/gif"
	case ".mp4":
		return "video/mp4"
	}
	return "image/jpeg"
}

func imageFilename(categoryPath, prefix string, unixMillis int64, index int, ext string, width, height int) string {
	if ext == "" {
		ext = ".jpg"
	}
	sizeSuffix := ""
	if width > 0 && height > 0 {
		sizeSuffix = fmt.Sprintf("_size_%dx%d", width, height)
	}
	body := fmt.Sprintf("%s_%d_%d%s%s", prefix, unixMillis, index, sizeSuffix, ext)
	if categoryPath == "" {
		return body
	}
	return categoryPath + "/" + body
}

// thumbnailFilename mirrors imageFilename but adds the "thumb_"
// prefix the server expects for the resized companion image. The size
// suffix carries the *original* image dimensions, not the thumbnail's.
func thumbnailFilename(categoryPath, prefix string, unixMillis int64, index int, ext string, origWidth, origHeight int) string {
	if ext == "" {
		ext = ".jpg"
	}
	sizeSuffix := ""
	if origWidth > 0 && origHeight > 0 {
		sizeSuffix = fmt.Sprintf("_size_%dx%d", origWidth, origHeight)
	}
	body := fmt.Sprintf("thumb_%s_%d_%d%s%s", prefix, unixMillis, index, sizeSuffix, ext)
	if categoryPath == "" {
		return body
	}
	return categoryPath + "/" + body
}

// makeThumbnailFor produces the right kind of thumbnail for the
// source format. GIF inputs become a smaller animated GIF (so the
// avatar/icon stays animated when the server displays the thumbnail),
// everything else becomes a JPEG. The returned ext / contentType
// match the produced bytes.
//
// Sources already inside maxW x maxH are returned at their original
// size — upscaling would only blur the result.
func makeThumbnailFor(body []byte, sourceExt string, maxW, maxH int) (out []byte, ext, contentType string, err error) {
	if sourceExt == ".gif" {
		out, err := makeGIFThumbnail(body, maxW, maxH)
		if err != nil {
			return nil, "", "", err
		}
		return out, ".gif", "image/gif", nil
	}
	out, err = makeJPEGThumbnail(body, maxW, maxH)
	if err != nil {
		return nil, "", "", err
	}
	return out, ".jpeg", "image/jpeg", nil
}

// makeJPEGThumbnail decodes any standard image format and re-encodes
// it as a shrunk JPEG that fits inside maxW x maxH while preserving
// the source aspect ratio.
func makeJPEGThumbnail(body []byte, maxW, maxH int) ([]byte, error) {
	if maxW <= 0 || maxH <= 0 {
		return nil, fmt.Errorf("invalid target size %dx%d", maxW, maxH)
	}
	src, _, err := image.Decode(bytes.NewReader(body))
	if err != nil {
		return nil, fmt.Errorf("decode source image: %w", err)
	}
	srcBounds := src.Bounds()
	dstW, dstH := fitInside(srcBounds.Dx(), srcBounds.Dy(), maxW, maxH)
	if dstW == 0 {
		return nil, fmt.Errorf("empty source image")
	}

	dst := image.NewRGBA(image.Rect(0, 0, dstW, dstH))
	draw.BiLinear.Scale(dst, dst.Bounds(), src, srcBounds, draw.Over, nil)

	var buf bytes.Buffer
	if err := jpeg.Encode(&buf, dst, &jpeg.Options{Quality: 85}); err != nil {
		return nil, fmt.Errorf("encode thumbnail: %w", err)
	}
	return buf.Bytes(), nil
}

// makeGIFThumbnail decodes an animated GIF, scales every frame down
// to fit inside maxW x maxH (preserving aspect ratio), and re-encodes
// the result as a GIF that retains the original animation. If the
// source already fits inside the box, the original bytes are returned
// untouched.
func makeGIFThumbnail(body []byte, maxW, maxH int) ([]byte, error) {
	if maxW <= 0 || maxH <= 0 {
		return nil, fmt.Errorf("invalid target size %dx%d", maxW, maxH)
	}
	all, err := gif.DecodeAll(bytes.NewReader(body))
	if err != nil {
		return nil, fmt.Errorf("decode gif: %w", err)
	}
	srcW, srcH := all.Config.Width, all.Config.Height
	dstW, dstH := fitInside(srcW, srcH, maxW, maxH)
	if dstW == 0 {
		return nil, fmt.Errorf("empty gif")
	}
	if dstW == srcW && dstH == srcH {
		return body, nil
	}

	// Composite each frame onto a full-canvas RGBA so resizing covers
	// any partial-update frames correctly, then re-quantize back to a
	// paletted image of the new size.
	canvas := image.NewRGBA(image.Rect(0, 0, srcW, srcH))
	dstBounds := image.Rect(0, 0, dstW, dstH)
	for i, frame := range all.Image {
		draw.Draw(canvas, frame.Bounds(), frame, frame.Bounds().Min, draw.Over)

		scaled := image.NewRGBA(dstBounds)
		draw.BiLinear.Scale(scaled, dstBounds, canvas, canvas.Bounds(), draw.Over, nil)

		paletted := image.NewPaletted(dstBounds, frame.Palette)
		draw.FloydSteinberg.Draw(paletted, dstBounds, scaled, image.Point{})
		all.Image[i] = paletted
		// Disposal flags may reference sub-rects of the original size
		// that no longer apply once every frame is full-canvas — fall
		// back to "no disposal" so the next frame paints over cleanly.
		if i < len(all.Disposal) {
			all.Disposal[i] = gif.DisposalNone
		}
	}
	all.Config.Width = dstW
	all.Config.Height = dstH

	var buf bytes.Buffer
	if err := gif.EncodeAll(&buf, all); err != nil {
		return nil, fmt.Errorf("encode gif thumbnail: %w", err)
	}
	return buf.Bytes(), nil
}

// fitInside returns the width/height that fit srcW x srcH inside
// maxW x maxH while preserving aspect ratio. Sources already smaller
// than the box are returned unchanged. (0, 0) signals invalid input.
func fitInside(srcW, srcH, maxW, maxH int) (int, int) {
	if srcW <= 0 || srcH <= 0 {
		return 0, 0
	}
	scale := 1.0
	if s := float64(maxW) / float64(srcW); s < scale {
		scale = s
	}
	if s := float64(maxH) / float64(srcH); s < scale {
		scale = s
	}
	dstW := int(float64(srcW) * scale)
	dstH := int(float64(srcH) * scale)
	if dstW < 1 {
		dstW = 1
	}
	if dstH < 1 {
		dstH = 1
	}
	return dstW, dstH
}

// normalizeImageExt returns one of .jpg, .jpeg, .png, .gif. Anything
// else (no extension, .heic, weird casings, …) collapses to .jpg.
func normalizeImageExt(filename string) string {
	switch extOf(filename) {
	case ".png":
		return ".png"
	case ".gif":
		return ".gif"
	case ".jpeg":
		return ".jpeg"
	case ".jpg":
		return ".jpg"
	}
	return ".jpg"
}

func videoFilename(prefix string, unixMillis int64) string {
	return fmt.Sprintf("%s_%d.mp4", prefix, unixMillis)
}

func randomFilenamePrefix(n int) string {
	const chars = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	b := make([]byte, n)
	if _, err := rand.Read(b); err != nil {
		// crypto/rand on a healthy host does not fail; keep going with the
		// zero-initialized buffer rather than aborting at filename time.
	}
	for i := range b {
		b[i] = chars[int(b[i])%len(chars)]
	}
	return string(b)
}

// decodeImageSize peeks at the first few bytes of body to read the
// width/height of a JPEG, PNG, or GIF without fully decoding the image.
// On any error it returns 0,0 and the caller drops the size suffix.
func decodeImageSize(body []byte) (int, int, error) {
	cfg, _, err := image.DecodeConfig(bytes.NewReader(body))
	if err != nil {
		return 0, 0, err
	}
	return cfg.Width, cfg.Height, nil
}

// formatYMD matches the Calendar-derived "YYYY/M/D" path the server
// expects in date-bucketed categories — note the un-padded month and
// day.
func formatYMD(t time.Time) string {
	return fmt.Sprintf("%d/%d/%d", t.Year(), int(t.Month()), t.Day())
}
