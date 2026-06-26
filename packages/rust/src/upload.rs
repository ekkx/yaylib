use std::io::Cursor;
use std::time::{SystemTime, UNIX_EPOCH};

use image::codecs::gif::{GifDecoder, GifEncoder, Repeat};
use image::codecs::jpeg::JpegEncoder;
use image::{AnimationDecoder, DynamicImage, Frame, GenericImageView, ImageDecoder, RgbaImage};
use rand::Rng;

use crate::errors::ApiError;

pub const MAX_IMAGES_PER_UPLOAD: usize = 9;

pub const MAX_REPORT_IMAGES_PER_UPLOAD: usize = 4;

pub const MEDIA_CDN_BASE: &str = "https://cdn.yay.space/uploads";

/// Constructs a full CDN URL from a server-canonical filename.
pub fn media_url(filename: &str) -> String {
    if filename.is_empty() {
        return String::new();
    }
    format!("{}/{}", MEDIA_CDN_BASE, filename.trim_start_matches('/'))
}

/// A file to upload to the Yay! API (filename used for extension detection).
pub struct Upload {
    pub filename: String,
    pub body: Vec<u8>,
}

pub(crate) trait UploadCategory {
    fn category_path(&self, now: SystemTime) -> String;
    fn thumbnail_size(&self) -> Option<(u32, u32)>;
    fn max_files(&self) -> usize;
}

struct UserAvatarUpload {
    user_id: i64,
}
impl UploadCategory for UserAvatarUpload {
    fn category_path(&self, _now: SystemTime) -> String {
        format!("user/{}/avatar", self.user_id)
    }
    fn thumbnail_size(&self) -> Option<(u32, u32)> {
        Some((200, 200))
    }
    fn max_files(&self) -> usize {
        MAX_IMAGES_PER_UPLOAD
    }
}

struct UserCoverUpload {
    user_id: i64,
}
impl UploadCategory for UserCoverUpload {
    fn category_path(&self, _now: SystemTime) -> String {
        format!("user/{}/cover", self.user_id)
    }
    fn thumbnail_size(&self) -> Option<(u32, u32)> {
        Some((300, 150))
    }
    fn max_files(&self) -> usize {
        MAX_IMAGES_PER_UPLOAD
    }
}

struct UserPostUpload {
    user_id: i64,
}
impl UploadCategory for UserPostUpload {
    fn category_path(&self, now: SystemTime) -> String {
        format!("user/{}/posts/{}", self.user_id, format_ymd(now))
    }
    fn thumbnail_size(&self) -> Option<(u32, u32)> {
        Some((450, 450))
    }
    fn max_files(&self) -> usize {
        MAX_IMAGES_PER_UPLOAD
    }
}

struct ChatMessageUpload {
    room_id: i64,
    user_id: i64,
}
impl UploadCategory for ChatMessageUpload {
    fn category_path(&self, now: SystemTime) -> String {
        format!(
            "chatroom/{}/user/{}/messages/{}",
            self.room_id,
            self.user_id,
            format_ymd(now)
        )
    }
    fn thumbnail_size(&self) -> Option<(u32, u32)> {
        Some((450, 450))
    }
    fn max_files(&self) -> usize {
        MAX_IMAGES_PER_UPLOAD
    }
}

struct ChatBackgroundUpload {
    room_id: i64,
}
impl UploadCategory for ChatBackgroundUpload {
    fn category_path(&self, _now: SystemTime) -> String {
        format!("chatroom/{}/background", self.room_id)
    }
    fn thumbnail_size(&self) -> Option<(u32, u32)> {
        Some((200, 200))
    }
    fn max_files(&self) -> usize {
        MAX_IMAGES_PER_UPLOAD
    }
}

struct GroupIconUpload {
    group_id: i64,
}
impl UploadCategory for GroupIconUpload {
    fn category_path(&self, _now: SystemTime) -> String {
        format!("group/{}/icon", self.group_id)
    }
    fn thumbnail_size(&self) -> Option<(u32, u32)> {
        Some((300, 300))
    }
    fn max_files(&self) -> usize {
        MAX_IMAGES_PER_UPLOAD
    }
}

struct GroupCoverUpload {
    group_id: i64,
}
impl UploadCategory for GroupCoverUpload {
    fn category_path(&self, _now: SystemTime) -> String {
        format!("group/{}/cover", self.group_id)
    }
    fn thumbnail_size(&self) -> Option<(u32, u32)> {
        Some((300, 300))
    }
    fn max_files(&self) -> usize {
        MAX_IMAGES_PER_UPLOAD
    }
}

struct GroupCreationIconUpload;
impl UploadCategory for GroupCreationIconUpload {
    fn category_path(&self, now: SystemTime) -> String {
        format!("group/create/icon/{}", format_ymd(now))
    }
    fn thumbnail_size(&self) -> Option<(u32, u32)> {
        Some((300, 300))
    }
    fn max_files(&self) -> usize {
        MAX_IMAGES_PER_UPLOAD
    }
}

struct GroupCreationCoverUpload;
impl UploadCategory for GroupCreationCoverUpload {
    fn category_path(&self, now: SystemTime) -> String {
        format!("group/create/cover/{}", format_ymd(now))
    }
    fn thumbnail_size(&self) -> Option<(u32, u32)> {
        Some((300, 300))
    }
    fn max_files(&self) -> usize {
        MAX_IMAGES_PER_UPLOAD
    }
}

struct SignupAvatarUpload;
impl UploadCategory for SignupAvatarUpload {
    fn category_path(&self, now: SystemTime) -> String {
        format!("user/create/{}", format_ymd(now))
    }
    fn thumbnail_size(&self) -> Option<(u32, u32)> {
        Some((200, 200))
    }
    fn max_files(&self) -> usize {
        MAX_IMAGES_PER_UPLOAD
    }
}

struct ThreadIconUpload {
    group_id: i64,
}
impl UploadCategory for ThreadIconUpload {
    fn category_path(&self, now: SystemTime) -> String {
        format!("group/{}/threads/{}", self.group_id, format_ymd(now))
    }
    fn thumbnail_size(&self) -> Option<(u32, u32)> {
        Some((300, 300))
    }
    fn max_files(&self) -> usize {
        MAX_IMAGES_PER_UPLOAD
    }
}

struct ReportUpload;
impl UploadCategory for ReportUpload {
    fn category_path(&self, now: SystemTime) -> String {
        format!("report/{}", format_ymd(now))
    }
    fn thumbnail_size(&self) -> Option<(u32, u32)> {
        Some((450, 450))
    }
    fn max_files(&self) -> usize {
        MAX_REPORT_IMAGES_PER_UPLOAD
    }
}

struct VideoCallSnapshotUpload;
impl UploadCategory for VideoCallSnapshotUpload {
    fn category_path(&self, now: SystemTime) -> String {
        format!("video_call_snapshot/{}", format_ymd(now))
    }
    fn thumbnail_size(&self) -> Option<(u32, u32)> {
        None
    }
    fn max_files(&self) -> usize {
        MAX_IMAGES_PER_UPLOAD
    }
}

struct UploadJob {
    body: Vec<u8>,
    filename: String,
    content_type: &'static str,
}

/// Uploads images for the given category, obtaining presigned URLs, generating
/// thumbnails, and PUTting everything to S3. Returns the server-canonical
/// filenames (with `s3/` prefix).
///
/// This is the shared engine; public upload methods on Client delegate here.
pub(crate) async fn upload_images(
    http: &reqwest::Client,
    base_url: &str,
    access_token: Option<&str>,
    category: &dyn UploadCategory,
    files: &[Upload],
) -> Result<Vec<String>, ApiError> {
    if files.is_empty() {
        return Err(ApiError::Other("upload requires at least one file".into()));
    }
    let cap = category.max_files();
    if files.len() > cap {
        return Err(ApiError::Other(format!(
            "this upload accepts at most {} files (got {})",
            cap,
            files.len()
        )));
    }

    let now = SystemTime::now();
    let prefix = random_filename_prefix(16);
    let ts = now
        .duration_since(UNIX_EPOCH)
        .unwrap_or_default()
        .as_millis() as i64;
    let category_path = category.category_path(now);
    let thumb_size = category.thumbnail_size();

    let mut main_jobs: Vec<UploadJob> = Vec::with_capacity(files.len());
    let mut thumb_jobs: Vec<UploadJob> = Vec::new();

    for (i, f) in files.iter().enumerate() {
        let ext = normalize_image_ext(&f.filename);
        let (w, h) = decode_image_size(&f.body).map_err(|e| {
            ApiError::Other(format!("decode image {}: {}", i, e))
        })?;

        let main_name = image_filename(&category_path, &prefix, ts, i, &ext, w, h);
        let ct = content_type_for(&main_name);
        main_jobs.push(UploadJob {
            body: f.body.clone(),
            filename: main_name,
            content_type: ct,
        });

        if let Some((max_w, max_h)) = thumb_size {
            let (thumb_bytes, thumb_ext, thumb_ct) =
                make_thumbnail_for(&f.body, &ext, max_w, max_h).map_err(|e| {
                    ApiError::Other(format!("make thumbnail for image {}: {}", i, e))
                })?;
            let thumb_name =
                thumbnail_filename(&category_path, &prefix, ts, i, &thumb_ext, w, h);
            thumb_jobs.push(UploadJob {
                body: thumb_bytes,
                filename: thumb_name,
                content_type: thumb_ct,
            });
        }
    }

    let mut all_jobs: Vec<UploadJob> = Vec::with_capacity(main_jobs.len() + thumb_jobs.len());
    let main_count = main_jobs.len();
    all_jobs.extend(main_jobs);
    all_jobs.extend(thumb_jobs);

    let all_filenames: Vec<&str> = all_jobs.iter().map(|j| j.filename.as_str()).collect();

    let presigned = get_bucket_presigned_urls(http, base_url, access_token, &all_filenames).await?;

    if presigned.len() != all_jobs.len() {
        return Err(ApiError::Other(format!(
            "presigned urls count mismatch (want {} got {})",
            all_jobs.len(),
            presigned.len()
        )));
    }

    let mut handles = Vec::with_capacity(all_jobs.len());
    for (i, job) in all_jobs.into_iter().enumerate() {
        let url = presigned[i].url.clone();
        if url.is_empty() {
            return Err(ApiError::Other(format!(
                "upload {}: empty presigned url",
                job.filename
            )));
        }
        let client = http.clone();
        let filename = job.filename.clone();
        let ct = job.content_type;
        let body = job.body;
        handles.push(tokio::spawn(async move {
            put_to_presigned_url(&client, &url, body, ct)
                .await
                .map_err(|e| ApiError::Other(format!("upload {}: {}", filename, e)))
        }));
    }

    for handle in handles {
        handle.await.map_err(|e| ApiError::Other(format!("join: {}", e)))??;
    }

    let out: Vec<String> = presigned[..main_count]
        .iter()
        .map(|p| p.filename.clone())
        .collect();
    Ok(out)
}

/// Uploads a single MP4 video. Returns the filename to pass to CreatePost as
/// `video_file_name`. Unlike images, the server does NOT return a canonical
/// filename for videos, so the random name the SDK generates is what callers
/// hand back.
pub(crate) async fn upload_video(
    http: &reqwest::Client,
    base_url: &str,
    access_token: Option<&str>,
    body: Vec<u8>,
) -> Result<String, ApiError> {
    if body.is_empty() {
        return Err(ApiError::Other("video body is empty".into()));
    }
    let prefix = random_filename_prefix(16);
    let ts = SystemTime::now()
        .duration_since(UNIX_EPOCH)
        .unwrap_or_default()
        .as_millis() as i64;
    let name = video_filename(&prefix, ts);

    let presigned_url = get_user_presigned_url(http, base_url, access_token, &name).await?;
    if presigned_url.is_empty() {
        return Err(ApiError::Other("empty presigned url".into()));
    }

    put_to_presigned_url(http, &presigned_url, body, "video/mp4")
        .await
        .map_err(|e| ApiError::Other(format!("upload video: {}", e)))?;

    Ok(name)
}

// --- Upload helpers for Client convenience methods ---

pub(crate) fn avatar_category(user_id: i64) -> impl UploadCategory {
    UserAvatarUpload { user_id }
}

pub(crate) fn cover_category(user_id: i64) -> impl UploadCategory {
    UserCoverUpload { user_id }
}

pub(crate) fn post_category(user_id: i64) -> impl UploadCategory {
    UserPostUpload { user_id }
}

pub(crate) fn chat_message_category(room_id: i64, user_id: i64) -> impl UploadCategory {
    ChatMessageUpload { room_id, user_id }
}

pub(crate) fn chat_background_category(room_id: i64) -> impl UploadCategory {
    ChatBackgroundUpload { room_id }
}

pub(crate) fn group_icon_category(group_id: i64) -> impl UploadCategory {
    GroupIconUpload { group_id }
}

pub(crate) fn group_cover_category(group_id: i64) -> impl UploadCategory {
    GroupCoverUpload { group_id }
}

pub(crate) fn group_creation_icon_category() -> impl UploadCategory {
    GroupCreationIconUpload
}

pub(crate) fn group_creation_cover_category() -> impl UploadCategory {
    GroupCreationCoverUpload
}

pub(crate) fn signup_avatar_category() -> impl UploadCategory {
    SignupAvatarUpload
}

pub(crate) fn thread_icon_category(group_id: i64) -> impl UploadCategory {
    ThreadIconUpload { group_id }
}

pub(crate) fn report_category() -> impl UploadCategory {
    ReportUpload
}

pub(crate) fn video_call_snapshot_category() -> impl UploadCategory {
    VideoCallSnapshotUpload
}

// --- Presigned URL API ---

#[derive(Debug, serde::Deserialize)]
struct PresignedUrlEntry {
    #[serde(default)]
    url: String,
    #[serde(default)]
    filename: String,
}

#[derive(Debug, serde::Deserialize)]
struct PresignedUrlsResponse {
    #[serde(default)]
    presigned_urls: Vec<PresignedUrlEntry>,
}

async fn get_bucket_presigned_urls(
    http: &reqwest::Client,
    base_url: &str,
    access_token: Option<&str>,
    file_names: &[&str],
) -> Result<Vec<PresignedUrlEntry>, ApiError> {
    let url = format!("{}/v1/buckets/presigned_urls", base_url.trim_end_matches('/'));
    let mut req = http.get(&url);
    if let Some(token) = access_token {
        req = req.bearer_auth(token);
    }
    for name in file_names {
        req = req.query(&[("file_names[]", name)]);
    }
    let resp = req.send().await?;
    let status = resp.status();
    if !status.is_success() {
        let body = resp.bytes().await.unwrap_or_default();
        return Err(ApiError::from_response(status.as_u16(), body.to_vec()));
    }
    let parsed: PresignedUrlsResponse = resp.json().await?;
    Ok(parsed.presigned_urls)
}

#[derive(Debug, serde::Deserialize)]
struct UserPresignedUrlResponse {
    #[serde(default)]
    presigned_url: String,
}

async fn get_user_presigned_url(
    http: &reqwest::Client,
    base_url: &str,
    access_token: Option<&str>,
    video_file_name: &str,
) -> Result<String, ApiError> {
    let url = format!(
        "{}/v1/users/presigned_url",
        base_url.trim_end_matches('/')
    );
    let mut req = http.get(&url).query(&[("video_file_name", video_file_name)]);
    if let Some(token) = access_token {
        req = req.bearer_auth(token);
    }
    let resp = req.send().await?;
    let status = resp.status();
    if !status.is_success() {
        let body = resp.bytes().await.unwrap_or_default();
        return Err(ApiError::from_response(status.as_u16(), body.to_vec()));
    }
    let parsed: UserPresignedUrlResponse = resp.json().await?;
    Ok(parsed.presigned_url)
}

async fn put_to_presigned_url(
    http: &reqwest::Client,
    url: &str,
    body: Vec<u8>,
    content_type: &str,
) -> Result<(), String> {
    let resp = http
        .put(url)
        .header("Content-Type", content_type)
        .body(body)
        .send()
        .await
        .map_err(|e| e.to_string())?;

    if !resp.status().is_success() {
        let status = resp.status();
        let preview = resp
            .text()
            .await
            .unwrap_or_default()
            .chars()
            .take(512)
            .collect::<String>();
        return Err(format!("PUT {}: {}", status.as_u16(), preview.trim()));
    }
    Ok(())
}

// --- Filename generation ---

fn image_filename(
    category_path: &str,
    prefix: &str,
    unix_millis: i64,
    index: usize,
    ext: &str,
    width: u32,
    height: u32,
) -> String {
    let ext = if ext.is_empty() { ".jpg" } else { ext };
    let size_suffix = if width > 0 && height > 0 {
        format!("_size_{}x{}", width, height)
    } else {
        String::new()
    };
    let body = format!("{}_{}_{}{}{}",  prefix, unix_millis, index, size_suffix, ext);
    if category_path.is_empty() {
        body
    } else {
        format!("{}/{}", category_path, body)
    }
}

fn thumbnail_filename(
    category_path: &str,
    prefix: &str,
    unix_millis: i64,
    index: usize,
    ext: &str,
    orig_width: u32,
    orig_height: u32,
) -> String {
    let ext = if ext.is_empty() { ".jpg" } else { ext };
    let size_suffix = if orig_width > 0 && orig_height > 0 {
        format!("_size_{}x{}", orig_width, orig_height)
    } else {
        String::new()
    };
    let body = format!(
        "thumb_{}_{}_{}{}{}",
        prefix, unix_millis, index, size_suffix, ext
    );
    if category_path.is_empty() {
        body
    } else {
        format!("{}/{}", category_path, body)
    }
}

fn video_filename(prefix: &str, unix_millis: i64) -> String {
    format!("{}_{}.mp4", prefix, unix_millis)
}

fn random_filename_prefix(n: usize) -> String {
    const CHARS: &[u8] = b"0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ";
    let mut rng = rand::thread_rng();
    (0..n)
        .map(|_| CHARS[rng.gen_range(0..CHARS.len())] as char)
        .collect()
}

// --- Image utilities ---

fn normalize_image_ext(filename: &str) -> String {
    let ext = std::path::Path::new(filename)
        .extension()
        .map(|e| e.to_string_lossy().to_lowercase())
        .unwrap_or_default();
    match ext.as_str() {
        "png" => ".png".to_string(),
        "gif" => ".gif".to_string(),
        "jpeg" => ".jpeg".to_string(),
        "jpg" => ".jpg".to_string(),
        _ => ".jpg".to_string(),
    }
}

fn content_type_for(filename: &str) -> &'static str {
    let ext = std::path::Path::new(filename)
        .extension()
        .map(|e| e.to_string_lossy().to_lowercase())
        .unwrap_or_default();
    match ext.as_str() {
        "png" => "image/png",
        "gif" => "image/gif",
        "mp4" => "video/mp4",
        _ => "image/jpeg",
    }
}

fn decode_image_size(body: &[u8]) -> Result<(u32, u32), String> {
    let reader = image::ImageReader::new(Cursor::new(body))
        .with_guessed_format()
        .map_err(|e| format!("guess format: {}", e))?;
    let (w, h) = reader
        .into_dimensions()
        .map_err(|e| format!("read dimensions: {}", e))?;
    Ok((w, h))
}

fn make_thumbnail_for(
    body: &[u8],
    source_ext: &str,
    max_w: u32,
    max_h: u32,
) -> Result<(Vec<u8>, String, &'static str), String> {
    if source_ext == ".gif" {
        let out = make_gif_thumbnail(body, max_w, max_h)?;
        Ok((out, ".gif".to_string(), "image/gif"))
    } else {
        let out = make_jpeg_thumbnail(body, max_w, max_h)?;
        Ok((out, ".jpeg".to_string(), "image/jpeg"))
    }
}

fn make_jpeg_thumbnail(body: &[u8], max_w: u32, max_h: u32) -> Result<Vec<u8>, String> {
    if max_w == 0 || max_h == 0 {
        return Err(format!("invalid target size {}x{}", max_w, max_h));
    }
    let img = image::load_from_memory(body).map_err(|e| format!("decode source image: {}", e))?;
    let (src_w, src_h) = img.dimensions();
    let (dst_w, dst_h) = fit_inside(src_w, src_h, max_w, max_h);
    if dst_w == 0 {
        return Err("empty source image".into());
    }

    let resized = img.resize_exact(dst_w, dst_h, image::imageops::Triangle);

    let mut buf = Vec::new();
    let encoder = JpegEncoder::new_with_quality(&mut buf, 85);
    resized
        .write_with_encoder(encoder)
        .map_err(|e| format!("encode thumbnail: {}", e))?;
    Ok(buf)
}

/// Decodes an animated GIF, scales every frame to fit inside `max_w x max_h`
/// (preserving aspect ratio), and re-encodes as an animated GIF. If the source
/// already fits, the original bytes are returned untouched.
///
/// # Differences from the Go reference implementation
///
/// The Go version uses `draw.FloydSteinberg` (error-diffusion dithering) to
/// re-quantize each resized RGBA frame back to a paletted image, and sets
/// `Disposal = DisposalNone` on every frame.
///
/// The `image` crate's `GifEncoder` delegates quantization to the `gif` crate's
/// `Frame::from_rgba_speed`, which uses the NeuQuant algorithm (neural-network
/// based) instead of Floyd-Steinberg dithering.  NeuQuant generally produces
/// good palette selection but does not perform error diffusion, so smooth
/// gradients may show slight banding compared to Floyd-Steinberg.  At
/// `speed = 1` (best quality) the difference is minimal for thumbnail sizes.
///
/// The `image` crate also forces `DisposalMethod::Background` on every frame
/// (hardcoded in `encode_gif`).  This is acceptable because our output frames
/// are all full-canvas (composited and resized to the final dimensions), so
/// the disposal method has no visible effect -- each frame fully paints over
/// the previous one.
fn make_gif_thumbnail(body: &[u8], max_w: u32, max_h: u32) -> Result<Vec<u8>, String> {
    if max_w == 0 || max_h == 0 {
        return Err(format!("invalid target size {}x{}", max_w, max_h));
    }

    let decoder =
        GifDecoder::new(Cursor::new(body)).map_err(|e| format!("decode gif: {}", e))?;
    let (src_w, src_h) = decoder.dimensions();
    let (dst_w, dst_h) = fit_inside(src_w, src_h, max_w, max_h);
    if dst_w == 0 {
        return Err("empty gif".into());
    }
    if dst_w == src_w && dst_h == src_h {
        return Ok(body.to_vec());
    }

    let frames = decoder
        .into_frames()
        .collect_frames()
        .map_err(|e| format!("decode gif frames: {}", e))?;

    // Composite each frame onto a full-canvas RGBA so resizing covers
    // any partial-update frames correctly (same approach as Go).
    let mut canvas = RgbaImage::new(src_w, src_h);
    let mut out_frames: Vec<Frame> = Vec::with_capacity(frames.len());

    for frame in &frames {
        let frame_buf = frame.buffer();
        let (fx, fy) = (frame.left(), frame.top());
        for y in 0..frame_buf.height() {
            for x in 0..frame_buf.width() {
                let px = frame_buf.get_pixel(x, y);
                if px[3] > 0 {
                    let cx = fx + x;
                    let cy = fy + y;
                    if cx < src_w && cy < src_h {
                        canvas.put_pixel(cx, cy, *px);
                    }
                }
            }
        }

        let canvas_dyn = DynamicImage::ImageRgba8(canvas.clone());
        let resized = canvas_dyn.resize_exact(dst_w, dst_h, image::imageops::Triangle);
        let resized_rgba = resized.to_rgba8();

        out_frames.push(Frame::from_parts(resized_rgba, 0, 0, frame.delay()));
    }

    let mut buf = Vec::new();
    {
        // speed=1 gives the best NeuQuant palette quality (range 1-30).
        let mut encoder = GifEncoder::new_with_speed(&mut buf, 1);
        encoder
            .set_repeat(Repeat::Infinite)
            .map_err(|e| format!("set gif repeat: {}", e))?;
        encoder
            .encode_frames(out_frames)
            .map_err(|e| format!("encode gif thumbnail: {}", e))?;
    }
    Ok(buf)
}

fn fit_inside(src_w: u32, src_h: u32, max_w: u32, max_h: u32) -> (u32, u32) {
    if src_w == 0 || src_h == 0 {
        return (0, 0);
    }
    let mut scale: f64 = 1.0;
    let sw = max_w as f64 / src_w as f64;
    if sw < scale {
        scale = sw;
    }
    let sh = max_h as f64 / src_h as f64;
    if sh < scale {
        scale = sh;
    }
    let dst_w = (src_w as f64 * scale) as u32;
    let dst_h = (src_h as f64 * scale) as u32;
    (dst_w.max(1), dst_h.max(1))
}

/// Formats a timestamp as `YYYY/M/D` with NO zero-padding, matching the Yay!
/// server's date-bucketed category path convention.
fn format_ymd(t: SystemTime) -> String {
    let dur = t.duration_since(UNIX_EPOCH).unwrap_or_default();
    let secs = dur.as_secs() as i64;

    let days = secs / 86400;
    let mut y = 1970i32;
    let mut remaining_days = days;

    loop {
        let days_in_year = if is_leap_year(y) { 366 } else { 365 };
        if remaining_days < days_in_year {
            break;
        }
        remaining_days -= days_in_year;
        y += 1;
    }

    let month_days: [i64; 12] = if is_leap_year(y) {
        [31, 29, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31]
    } else {
        [31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31]
    };

    let mut m = 0u32;
    for (i, &md) in month_days.iter().enumerate() {
        if remaining_days < md {
            m = i as u32 + 1;
            break;
        }
        remaining_days -= md;
    }
    if m == 0 {
        m = 12;
    }

    let d = remaining_days + 1;
    format!("{}/{}/{}", y, m, d)
}

fn is_leap_year(y: i32) -> bool {
    (y % 4 == 0 && y % 100 != 0) || (y % 400 == 0)
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_format_ymd_no_zero_pad() {
        // 2026-04-26 00:00:00 UTC = day 20569 from epoch
        // (56 years 1970-2025 = 20454 days, then Jan31+Feb28+Mar31+25 days of Apr = 115)
        let t = UNIX_EPOCH + std::time::Duration::from_secs(20569 * 86400);
        let s = format_ymd(t);
        assert_eq!(s, "2026/4/26");
    }

    #[test]
    fn test_format_ymd_epoch() {
        let s = format_ymd(UNIX_EPOCH);
        assert_eq!(s, "1970/1/1");
    }

    #[test]
    fn test_fit_inside_no_upscale() {
        assert_eq!(fit_inside(100, 100, 200, 200), (100, 100));
    }

    #[test]
    fn test_fit_inside_shrink() {
        assert_eq!(fit_inside(400, 200, 200, 200), (200, 100));
    }

    #[test]
    fn test_fit_inside_zero() {
        assert_eq!(fit_inside(0, 0, 200, 200), (0, 0));
    }

    #[test]
    fn test_image_filename_format() {
        let name = image_filename("user/123/avatar", "abcdefgh12345678", 1700000000000, 0, ".jpg", 640, 480);
        assert_eq!(
            name,
            "user/123/avatar/abcdefgh12345678_1700000000000_0_size_640x480.jpg"
        );
    }

    #[test]
    fn test_thumbnail_filename_format() {
        let name = thumbnail_filename("user/123/avatar", "abcdefgh12345678", 1700000000000, 0, ".jpg", 640, 480);
        assert_eq!(
            name,
            "user/123/avatar/thumb_abcdefgh12345678_1700000000000_0_size_640x480.jpg"
        );
    }

    #[test]
    fn test_video_filename_format() {
        let name = video_filename("abcdefgh12345678", 1700000000000);
        assert_eq!(name, "abcdefgh12345678_1700000000000.mp4");
    }

    #[test]
    fn test_normalize_image_ext() {
        assert_eq!(normalize_image_ext("photo.png"), ".png");
        assert_eq!(normalize_image_ext("photo.GIF"), ".gif");
        assert_eq!(normalize_image_ext("photo.JPEG"), ".jpeg");
        assert_eq!(normalize_image_ext("photo.jpg"), ".jpg");
        assert_eq!(normalize_image_ext("photo.heic"), ".jpg");
        assert_eq!(normalize_image_ext("photo"), ".jpg");
    }

    #[test]
    fn test_content_type_for() {
        assert_eq!(content_type_for("test.png"), "image/png");
        assert_eq!(content_type_for("test.gif"), "image/gif");
        assert_eq!(content_type_for("test.mp4"), "video/mp4");
        assert_eq!(content_type_for("test.jpg"), "image/jpeg");
        assert_eq!(content_type_for("test.jpeg"), "image/jpeg");
        assert_eq!(content_type_for("test"), "image/jpeg");
    }

    #[test]
    fn test_media_url() {
        assert_eq!(
            media_url("s3/user/123/avatar/foo.jpg"),
            "https://cdn.yay.space/uploads/s3/user/123/avatar/foo.jpg"
        );
        assert_eq!(media_url(""), "");
        assert_eq!(
            media_url("/leading_slash.jpg"),
            "https://cdn.yay.space/uploads/leading_slash.jpg"
        );
    }

    #[test]
    fn test_random_filename_prefix_length() {
        let p = random_filename_prefix(16);
        assert_eq!(p.len(), 16);
        assert!(p.chars().all(|c| c.is_ascii_alphanumeric()));
    }
}
