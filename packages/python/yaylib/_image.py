"""PORTING.md §8: the upload-protocol pieces that aren't visible from the
category table — filename shaping, source-size probing, and the thumbnail
rules. The Yay! server pattern-matches the filenames and DELETES the main
image if its thumbnail is missing, statically encoded, or
extension-mismatched, so this module mirrors upload.go / image.ts exactly:
JPEG/PNG sources produce a shrunk JPEG thumbnail, animated GIF sources
stay animated through the resize.
"""

from __future__ import annotations

import io
import secrets
import time
from datetime import datetime, timezone
from typing import Tuple

from PIL import Image, ImageSequence

# Default per-call cap on multi-image uploads (upload_post_images,
# upload_chat_message_images). upload_report_images is tighter.
MAX_IMAGES_PER_UPLOAD = 9

# Per-call cap on upload_report_images (the server rejects the 5th).
MAX_REPORT_IMAGES_PER_UPLOAD = 4

_PREFIX_ALPHABET = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"


def ext_of(filename: str) -> str:
    """Lowercased extension including the leading dot, or "" when the
    filename has no extension. Mirrors Go's filepath.Ext semantics."""
    base = filename[filename.rfind("/") + 1 :]
    dot = base.rfind(".")
    if dot <= 0:
        return ""
    return base[dot:].lower()


def content_type_for(filename: str) -> str:
    """Map a filename extension to the Content-Type sent on the presigned
    PUT. Unknown extensions fall back to image/jpeg."""
    return {
        ".png": "image/png",
        ".gif": "image/gif",
        ".mp4": "video/mp4",
    }.get(ext_of(filename), "image/jpeg")


def normalize_image_ext(filename: str) -> str:
    """Collapse any input filename to one of .jpg, .jpeg, .png, .gif.
    Anything else becomes .jpg, matching upload.go."""
    e = ext_of(filename)
    return e if e in (".png", ".gif", ".jpeg", ".jpg") else ".jpg"


def _size_suffix(width: int, height: int) -> str:
    return f"_size_{width}x{height}" if width > 0 and height > 0 else ""


def image_filename(
    category_path: str,
    prefix: str,
    unix_millis: int,
    index: int,
    ext: str,
    width: int,
    height: int,
) -> str:
    """Main-image object key the Yay! server pattern-matches. Empty ext
    collapses to .jpg; the size suffix is omitted when either dimension
    is non-positive."""
    e = ext or ".jpg"
    body = f"{prefix}_{unix_millis}_{index}{_size_suffix(width, height)}{e}"
    return body if category_path == "" else f"{category_path}/{body}"


def thumbnail_filename(
    category_path: str,
    prefix: str,
    unix_millis: int,
    index: int,
    ext: str,
    orig_width: int,
    orig_height: int,
) -> str:
    """Like image_filename but with the "thumb_" prefix the server
    expects for the resized companion. The size suffix carries the
    *original* image dimensions, not the thumbnail's."""
    e = ext or ".jpg"
    body = f"thumb_{prefix}_{unix_millis}_{index}{_size_suffix(orig_width, orig_height)}{e}"
    return body if category_path == "" else f"{category_path}/{body}"


def video_filename(prefix: str, unix_millis: int) -> str:
    """The flat key videos use — no category path, no size suffix."""
    return f"{prefix}_{unix_millis}.mp4"


def random_filename_prefix(n: int) -> str:
    """n characters from [0-9A-Za-z], freshly generated per upload call.
    upload.go panics rather than fabricate a fixed prefix that would
    guarantee bucket collisions; secrets is always available here."""
    return "".join(secrets.choice(_PREFIX_ALPHABET) for _ in range(n))


def now_unix_millis() -> int:
    return int(time.time() * 1000)


def format_ymd(d: datetime) -> str:
    """"YYYY/M/D" with un-padded month and day. UTC fields make the
    bucket path deterministic regardless of host timezone (PORTING.md
    §8: the server treats the path as a hint)."""
    u = d.astimezone(timezone.utc)
    return f"{u.year}/{u.month}/{u.day}"


def fit_inside(src_w: int, src_h: int, max_w: int, max_h: int) -> Tuple[int, int]:
    """Width/height that fit src_w x src_h inside max_w x max_h while
    preserving aspect ratio. Sources already smaller than the box are
    returned unchanged. (0, 0) signals invalid input."""
    if src_w <= 0 or src_h <= 0:
        return (0, 0)
    scale = 1.0
    scale = min(scale, max_w / src_w)
    scale = min(scale, max_h / src_h)
    dst_w = max(1, int(src_w * scale))
    dst_h = max(1, int(src_h * scale))
    return (dst_w, dst_h)


def probe_image_size(body: bytes) -> Tuple[int, int]:
    """Native (width, height) of a JPEG, PNG, or GIF. An undecodable
    body raises — uploading a body whose size suffix is wrong (or whose
    required thumbnail can't be built) just triggers server-side
    moderation deletion, so we surface the failure."""
    try:
        with Image.open(io.BytesIO(body)) as im:
            w, h = im.size
    except Exception as err:  # noqa: BLE001
        raise ValueError(f"decode image: {err}") from err
    if w <= 0 or h <= 0:
        raise ValueError("decode image: could not read dimensions")
    return (w, h)


class Thumbnail:
    __slots__ = ("body", "ext", "content_type")

    def __init__(self, body: bytes, ext: str, content_type: str) -> None:
        self.body = body
        self.ext = ext
        self.content_type = content_type


def make_thumbnail_for(
    body: bytes, source_ext: str, max_w: int, max_h: int
) -> Thumbnail:
    """The right kind of thumbnail for the source format. GIF inputs stay
    animated GIFs (so the avatar/icon keeps moving when the server
    renders the thumbnail); everything else becomes a JPEG. The returned
    ext/content_type match the produced bytes, mirroring upload.go."""
    if source_ext == ".gif":
        return Thumbnail(_make_gif_thumbnail(body, max_w, max_h), ".gif", "image/gif")
    return Thumbnail(_make_jpeg_thumbnail(body, max_w, max_h), ".jpeg", "image/jpeg")


def _make_jpeg_thumbnail(body: bytes, max_w: int, max_h: int) -> bytes:
    if max_w <= 0 or max_h <= 0:
        raise ValueError(f"make thumbnail: invalid target size {max_w}x{max_h}")
    try:
        with Image.open(io.BytesIO(body)) as im:
            im = im.convert("RGB")
            sw, sh = im.size
            dst_w, dst_h = fit_inside(sw, sh, max_w, max_h)
            if dst_w == 0:
                raise ValueError("make thumbnail: empty source image")
            if (dst_w, dst_h) != (sw, sh):
                im = im.resize((dst_w, dst_h), Image.BILINEAR)
            out = io.BytesIO()
            im.save(out, format="JPEG", quality=85)
            return out.getvalue()
    except ValueError:
        raise
    except Exception as err:  # noqa: BLE001
        raise ValueError(f"make thumbnail: decode source image: {err}") from err


def _make_gif_thumbnail(body: bytes, max_w: int, max_h: int) -> bytes:
    if max_w <= 0 or max_h <= 0:
        raise ValueError(f"make thumbnail: invalid target size {max_w}x{max_h}")
    try:
        src = Image.open(io.BytesIO(body))
    except Exception as err:  # noqa: BLE001
        raise ValueError(f"make thumbnail: decode gif: {err}") from err
    sw, sh = src.size
    dst_w, dst_h = fit_inside(sw, sh, max_w, max_h)
    if dst_w == 0:
        raise ValueError("make thumbnail: empty gif")
    if (dst_w, dst_h) == (sw, sh):
        # Already fits — return the original bytes untouched (upload.go).
        return body

    frames = []
    durations = []
    for frame in ImageSequence.Iterator(src):
        # Compositing onto a full RGBA canvas before scaling mirrors the
        # full-canvas invariant upload.go establishes.
        rgba = frame.convert("RGBA")
        durations.append(frame.info.get("duration", src.info.get("duration", 0)))
        frames.append(rgba.resize((dst_w, dst_h), Image.BILINEAR).convert("P", palette=Image.ADAPTIVE))

    out = io.BytesIO()
    loop = src.info.get("loop", 0)
    frames[0].save(
        out,
        format="GIF",
        save_all=True,
        append_images=frames[1:],
        loop=loop if isinstance(loop, int) and loop >= 0 else 0,
        duration=durations,
        disposal=2,
        optimize=False,
    )
    return out.getvalue()
