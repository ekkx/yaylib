# Local upload fixtures (no shared server) — the pure-function and guard
# coverage that, per the parity decision boundary, stays local rather
# than going through the shared behavior server. Mirrors the kept-local
# half of the Go/TS upload tests.

import io
from datetime import datetime, timezone

import pytest
from PIL import Image, ImageSequence

from yaylib import _image
from yaylib.upload import (
    MAX_IMAGES_PER_UPLOAD,
    MAX_REPORT_IMAGES_PER_UPLOAD,
    Upload,
    report_upload,
    upload_images,
    upload_single_image,
    upload_video,
    user_avatar_upload,
    user_post_upload,
    video_call_snapshot_upload,
)


def _jpeg(w, h):
    b = io.BytesIO()
    Image.new("RGB", (w, h), (10, 20, 30)).save(b, format="JPEG", quality=75)
    return b.getvalue()


def _png(w, h):
    b = io.BytesIO()
    Image.new("RGB", (w, h), (40, 80, 120)).save(b, format="PNG")
    return b.getvalue()


def _gif(w, h, frames=3):
    # Visually distinct frames so Pillow's GIF encoder doesn't coalesce
    # them into a single still (which would make the source non-animated
    # and defeat the point of the test).
    b = io.BytesIO()
    imgs = [
        Image.new("RGB", (w, h), ((i * 60) % 256, 30, (200 - i * 40) % 256)).convert(
            "P", palette=Image.ADAPTIVE
        )
        for i in range(frames)
    ]
    imgs[0].save(
        b, format="GIF", save_all=True, append_images=imgs[1:], duration=80, loop=0
    )
    return b.getvalue()


# ---- pure helpers ----


def test_content_type_for():
    assert _image.content_type_for("a.png") == "image/png"
    assert _image.content_type_for("a.gif") == "image/gif"
    assert _image.content_type_for("a.mp4") == "video/mp4"
    assert _image.content_type_for("a.jpg") == "image/jpeg"
    assert _image.content_type_for("a.JPEG") == "image/jpeg"
    assert _image.content_type_for("a.heic") == "image/jpeg"  # fallback
    assert _image.content_type_for("noext") == "image/jpeg"


def test_normalize_image_ext():
    assert _image.normalize_image_ext("a.PNG") == ".png"
    assert _image.normalize_image_ext("a.gif") == ".gif"
    assert _image.normalize_image_ext("a.jpeg") == ".jpeg"
    assert _image.normalize_image_ext("a.jpg") == ".jpg"
    assert _image.normalize_image_ext("a.heic") == ".jpg"
    assert _image.normalize_image_ext("noext") == ".jpg"


def test_image_and_thumbnail_filename_format():
    assert (
        _image.image_filename("user/1/avatar", "abc", 1700000000000, 0, ".png", 12, 13)
        == "user/1/avatar/abc_1700000000000_0_size_12x13.png"
    )
    # thumb_ prefix, original dims in the size suffix
    assert (
        _image.thumbnail_filename("g/2", "P", 999, 1, ".jpeg", 60, 40)
        == "g/2/thumb_P_999_1_size_60x40.jpeg"
    )
    # empty ext collapses to .jpg; non-positive dims drop the suffix
    assert _image.image_filename("", "p", 5, 0, "", 0, 0) == "p_5_0.jpg"


def test_video_filename_is_flat():
    assert _image.video_filename("PFX", 12345) == "PFX_12345.mp4"


def test_format_ymd_no_zero_padding():
    d = datetime(2026, 4, 9, tzinfo=timezone.utc)
    assert _image.format_ymd(d) == "2026/4/9"


def test_fit_inside():
    assert _image.fit_inside(100, 50, 40, 40) == (40, 20)  # aspect preserved
    assert _image.fit_inside(30, 20, 200, 200) == (30, 20)  # no upscale
    assert _image.fit_inside(0, 10, 50, 50) == (0, 0)  # invalid


def test_random_filename_prefix_length_and_alphabet():
    p = _image.random_filename_prefix(16)
    assert len(p) == 16
    assert all(c.isalnum() and c.isascii() for c in p)
    assert _image.random_filename_prefix(16) != _image.random_filename_prefix(16)


def test_probe_image_size():
    assert _image.probe_image_size(_jpeg(64, 48)) == (64, 48)
    assert _image.probe_image_size(_png(12, 13)) == (12, 13)
    assert _image.probe_image_size(_gif(20, 10)) == (20, 10)
    with pytest.raises(ValueError, match="decode image"):
        _image.probe_image_size(b"not an image at all")


def test_make_jpeg_thumbnail_shrinks_and_is_jpeg():
    th = _image.make_thumbnail_for(_png(400, 200), ".png", 100, 100)
    assert th.ext == ".jpeg" and th.content_type == "image/jpeg"
    with Image.open(io.BytesIO(th.body)) as im:
        assert im.format == "JPEG"
        assert im.size == (100, 50)  # fit inside 100x100, aspect kept


def test_make_jpeg_thumbnail_no_upscale():
    small = _png(20, 10)
    th = _image.make_thumbnail_for(small, ".png", 200, 200)
    with Image.open(io.BytesIO(th.body)) as im:
        assert im.size == (20, 10)  # unchanged, no upscale


def test_make_gif_thumbnail_stays_animated_gif():
    th = _image.make_thumbnail_for(_gif(120, 60, frames=4), ".gif", 40, 40)
    assert th.ext == ".gif" and th.content_type == "image/gif"
    with Image.open(io.BytesIO(th.body)) as im:
        assert im.format == "GIF"
        assert getattr(im, "is_animated", False)
        assert im.n_frames == 4
        assert im.size == (40, 20)


def test_make_gif_thumbnail_no_upscale_returns_original_bytes():
    src = _gif(10, 10, frames=2)
    th = _image.make_thumbnail_for(src, ".gif", 300, 300)
    assert th.body == src  # already fits — untouched


# ---- categories ----


def test_category_paths_and_thumbnail_boxes():
    now = datetime(2026, 4, 9, tzinfo=timezone.utc)
    assert user_avatar_upload(42).path(now) == "user/42/avatar"
    assert user_avatar_upload(42).thumbnail() == (200, 200, True)
    assert user_post_upload(7).path(now) == "user/7/posts/2026/4/9"
    snap = video_call_snapshot_upload()
    assert snap.path(now) == "video_call_snapshot/2026/4/9"
    assert snap.thumbnail() == (0, 0, False)  # snapshot has no thumbnail
    assert report_upload().max_files() == MAX_REPORT_IMAGES_PER_UPLOAD


# ---- guards (raise before any network) ----


class _Deps:
    """A deps stub whose network methods explode — the guard tests must
    fail before reaching any of them."""

    def __init__(self, uid=1):
        self._uid = uid

    def user_id(self):
        return self._uid

    async def get_presigned_urls(self, file_names):
        raise AssertionError("network reached; a guard should have fired")

    async def get_video_presigned_url(self, video_file_name):
        raise AssertionError("network reached; a guard should have fired")

    async def raw_put(self, url, body, content_type):
        raise AssertionError("network reached; a guard should have fired")


async def test_empty_files_rejected():
    with pytest.raises(ValueError, match="at least one file"):
        await upload_images(_Deps(), user_avatar_upload(1), [])


async def test_overflow_rejected():
    files = [Upload(filename=f"{i}.jpg", body=_jpeg(4, 4)) for i in range(MAX_IMAGES_PER_UPLOAD + 1)]
    with pytest.raises(ValueError, match="at most"):
        await upload_images(_Deps(), user_post_upload(1), files)


async def test_report_overflow_rejected():
    files = [Upload(filename=f"{i}.jpg", body=_jpeg(4, 4)) for i in range(MAX_REPORT_IMAGES_PER_UPLOAD + 1)]
    with pytest.raises(ValueError, match="at most"):
        await upload_images(_Deps(), report_upload(), files)


async def test_non_image_body_surfaces_error():
    with pytest.raises(ValueError, match="decode image"):
        await upload_single_image(
            _Deps(), user_avatar_upload(1),
            Upload(filename="broken.jpg", body=b"not an image"),
        )


async def test_upload_video_none_body_rejected():
    with pytest.raises(ValueError, match="body is None"):
        await upload_video(_Deps(), None)


async def test_upload_requires_login_via_client():
    from yaylib.client import Client

    c = Client(base_url="http://unused.invalid")
    try:
        with pytest.raises(ValueError, match="not logged in"):
            await c.upload_avatar_image(Upload(filename="a.jpg", body=_jpeg(4, 4)))
    finally:
        await c.close()
