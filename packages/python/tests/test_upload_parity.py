# Upload behavior parity, translated from the Go reference's
# upload_parity_test.go (and the TS upload_parity.test.ts). Driven
# against the shared server's upload contract: GET
# /v1/buckets/presigned_urls returns one server-canonical
# ("s3/"-prefixed) filename + an in-process S3 URL per requested name;
# GET /v1/users/presigned_url does the same for video; the S3 receiver
# answers 200, or 403 if the PUT carries Authorization.
#
# The shared S3 receiver intentionally does not echo the Content-Type,
# the stored bytes, or the file_names[] it saw (behavioral-fidelity
# only). So the parity assertions are the client-observable ones:
# success/error, the returned canonical filename shape (the server just
# prepends "s3/" to the SDK-built name), and the server-enforced
# no-bearer rule on the presigned PUT. PUT content inspection, image
# processing, and the pure filename / path / MIME helpers stay as local
# fixtures in test_upload.py.

import io
import re

import pytest
from PIL import Image

from yaylib.retry import RetryPolicy
from yaylib.upload import Upload

from ._parity import mock_client

# Go's zero-value RetryPolicy{} (retry disabled). The Python default is
# max_attempts=3, so the presigned-failure case must request "no retry".
NO_RETRY = RetryPolicy(max_attempts=0, base_delay=0.0, max_delay=0.0)


def _jpeg(w: int, h: int) -> bytes:
    buf = io.BytesIO()
    Image.new("RGB", (w, h), (90, 140, 200)).save(buf, format="JPEG", quality=75)
    return buf.getvalue()


def _png(w: int, h: int) -> bytes:
    buf = io.BytesIO()
    Image.new("RGB", (w, h), (30, 90, 160)).save(buf, format="PNG")
    return buf.getvalue()


# PORTING:S15
async def test_upload_avatar_image_happy_path():
    c = mock_client("")
    c.set_login_identity("", 42)
    try:
        name = await c.upload_avatar_image(
            Upload(filename="me.jpg", body=_jpeg(100, 100))
        )
        assert name.startswith("s3/user/42/avatar/"), name
        assert "_size_100x100" in name, name
    finally:
        await c.close()


# PORTING:S15
async def test_upload_avatar_image_filename_shape():
    c = mock_client("")
    c.set_login_identity("", 1)
    try:
        name = await c.upload_avatar_image(
            Upload(filename="x.png", body=_png(12, 13))
        )
        assert re.match(
            r"^s3/user/1/avatar/[0-9A-Za-z]{16}_\d{13}_0_size_12x13\.png$", name
        ), name
    finally:
        await c.close()


async def test_upload_avatar_image_presigned_failure():
    # fail-503-times-1 makes the first request in the session — the
    # presigned-URL GET — fail; with retry disabled the upload errors.
    c = mock_client("fail-503-times-1", retry_policy=NO_RETRY)
    c.set_login_identity("", 1)
    try:
        with pytest.raises(Exception):
            await c.upload_avatar_image(
                Upload(filename="a.png", body=_png(5, 5))
            )
    finally:
        await c.close()


async def test_upload_video_happy_path():
    c = mock_client("")
    try:
        name = await c.upload_video(b"\x00\x00\x00 ftypisom--mp4-bytes--")
        assert name.endswith(".mp4"), name
    finally:
        await c.close()


# PORTING:S16
async def test_presigned_put_does_not_leak_bearer():
    # Tokens are set, but the presigned PUT must authenticate via the
    # signed query only. The shared S3 receiver answers 403 if the PUT
    # carries Authorization, so a successful upload proves no bearer
    # leaked — the check is server-enforced and identical across the
    # three languages.
    c = mock_client("")
    c.set_login_identity("", 1)
    c.set_tokens("SECRET-ACCESS-TOKEN", "REF")
    try:
        name = await c.upload_avatar_image(
            Upload(filename="x.png", body=_png(5, 5))
        )
        assert name.startswith("s3/user/1/avatar/"), name
    finally:
        await c.close()
