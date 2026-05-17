"""PORTING.md §8: the typed upload categories + upload_video. The category
table (method <-> bucket path <-> thumbnail size <-> file cap) lives here;
the filename shaping, the s3/ canonical-name contract and the animated-GIF
thumbnail rule live in _image.py.

The presigned PUT bypasses the wrapped client entirely (the transport's
Yay! headers / Bearer would conflict with the presigned URL's own
signature), exactly as upload.go routes it through the bare client. The
presigned-URL fetch and the ws-token fetch still go through the wrapped
client so they carry auth + retry.
"""

from __future__ import annotations

import asyncio
from dataclasses import dataclass
from datetime import datetime, timezone
from typing import Callable, List, Optional, Protocol, Tuple, Union

from yaylib._image import (
    MAX_IMAGES_PER_UPLOAD,
    MAX_REPORT_IMAGES_PER_UPLOAD,
    content_type_for,
    format_ymd,
    image_filename,
    make_thumbnail_for,
    normalize_image_ext,
    now_unix_millis,
    probe_image_size,
    random_filename_prefix,
    thumbnail_filename,
    video_filename,
)

__all__ = [
    "Upload",
    "UploadCategory",
    "MAX_IMAGES_PER_UPLOAD",
    "MAX_REPORT_IMAGES_PER_UPLOAD",
    "user_avatar_upload",
    "user_cover_upload",
    "user_post_upload",
    "chat_message_upload",
    "chat_background_upload",
    "group_icon_upload",
    "group_cover_upload",
    "group_creation_icon_upload",
    "group_creation_cover_upload",
    "signup_avatar_upload",
    "thread_icon_upload",
    "report_upload",
    "video_call_snapshot_upload",
]

_DEFAULT_MAX_FILES = MAX_IMAGES_PER_UPLOAD

# A file body is raw bytes or any object exposing a .read() -> bytes.
Body = Union[bytes, bytearray, "SupportsRead"]


class SupportsRead(Protocol):
    def read(self) -> bytes: ...


@dataclass
class Upload:
    """A single file passed to one of the upload_<xxx>_image[s] methods.
    ``filename`` derives the extension and Content-Type; ``body`` is
    buffered fully before the request is issued."""

    filename: str
    body: Body


def _to_bytes(body: Body) -> bytes:
    if isinstance(body, (bytes, bytearray)):
        return bytes(body)
    return bytes(body.read())


class UploadCategory:
    """Selects the bucket-side path prefix, the thumbnail box, and the
    per-call file cap."""

    __slots__ = ("path", "thumbnail", "max_files")

    def __init__(
        self,
        path: Callable[[datetime], str],
        thumbnail: Callable[[], Tuple[int, int, bool]],
        max_files: Callable[[], int],
    ) -> None:
        self.path = path
        self.thumbnail = thumbnail
        self.max_files = max_files


def _cat(
    path: Callable[[datetime], str],
    thumb: Tuple[int, int, bool],
    max_files: int = _DEFAULT_MAX_FILES,
) -> UploadCategory:
    return UploadCategory(path, lambda: thumb, lambda: max_files)


def user_avatar_upload(user_id: int) -> UploadCategory:
    return _cat(lambda _n: f"user/{user_id}/avatar", (200, 200, True))


def user_cover_upload(user_id: int) -> UploadCategory:
    return _cat(lambda _n: f"user/{user_id}/cover", (300, 150, True))


def user_post_upload(user_id: int) -> UploadCategory:
    return _cat(lambda n: f"user/{user_id}/posts/{format_ymd(n)}", (450, 450, True))


def chat_message_upload(room_id: int, user_id: int) -> UploadCategory:
    return _cat(
        lambda n: f"chatroom/{room_id}/user/{user_id}/messages/{format_ymd(n)}",
        (450, 450, True),
    )


def chat_background_upload(room_id: int) -> UploadCategory:
    return _cat(lambda _n: f"chatroom/{room_id}/background", (200, 200, True))


def group_icon_upload(group_id: int) -> UploadCategory:
    return _cat(lambda _n: f"group/{group_id}/icon", (300, 300, True))


def group_cover_upload(group_id: int) -> UploadCategory:
    return _cat(lambda _n: f"group/{group_id}/cover", (300, 300, True))


def group_creation_icon_upload() -> UploadCategory:
    return _cat(lambda n: f"group/create/icon/{format_ymd(n)}", (300, 300, True))


def group_creation_cover_upload() -> UploadCategory:
    return _cat(lambda n: f"group/create/cover/{format_ymd(n)}", (300, 300, True))


def signup_avatar_upload() -> UploadCategory:
    return _cat(lambda n: f"user/create/{format_ymd(n)}", (200, 200, True))


def thread_icon_upload(group_id: int) -> UploadCategory:
    return _cat(lambda n: f"group/{group_id}/threads/{format_ymd(n)}", (300, 300, True))


def report_upload() -> UploadCategory:
    return _cat(
        lambda n: f"report/{format_ymd(n)}",
        (450, 450, True),
        MAX_REPORT_IMAGES_PER_UPLOAD,
    )


def video_call_snapshot_upload() -> UploadCategory:
    return _cat(lambda n: f"video_call_snapshot/{format_ymd(n)}", (0, 0, False))


class _PresignedURL(Protocol):
    filename: Optional[str]
    url: Optional[str]


class UploadDeps(Protocol):
    """The slice of Client behaviour the upload core needs. The Client
    supplies it; keeping it abstract avoids a client <-> upload import
    cycle and makes the flow unit-testable."""

    def user_id(self) -> int: ...

    async def get_presigned_urls(
        self, file_names: List[str]
    ) -> List[_PresignedURL]: ...

    async def get_video_presigned_url(self, video_file_name: str) -> str: ...

    # raw_put bypasses the wrapped client — see the module header. Raises
    # on a non-2xx response or transport error.
    async def raw_put(self, url: str, body: bytes, content_type: str) -> None: ...


def require_user_id(deps: UploadDeps) -> int:
    uid = deps.user_id()
    if not uid:
        raise ValueError(
            "yaylib: not logged in (call login_with_email before user-bound uploads)"
        )
    return uid


class _Job:
    __slots__ = ("body", "filename", "content_type")

    def __init__(self, body: bytes, filename: str, content_type: str) -> None:
        self.body = body
        self.filename = filename
        self.content_type = content_type


async def upload_images(
    deps: UploadDeps, category: UploadCategory, files: List[Upload]
) -> List[str]:
    """The full presigned-URL + parallel-PUT dance shared by every
    image-upload method, mirroring upload.go's uploadImages. The returned
    names are the server-canonical (s3/-prefixed) main-image names;
    thumbnails live at sibling paths the server resolves automatically."""
    if category is None:
        raise ValueError("yaylib: upload requires a category")
    if not files:
        raise ValueError("yaylib: upload requires at least one file")
    cap = category.max_files()
    if len(files) > cap:
        raise ValueError(
            f"yaylib: this upload accepts at most {cap} files (got {len(files)})"
        )

    now = datetime.now(timezone.utc)
    prefix = random_filename_prefix(16)
    ts = now_unix_millis()
    category_path = category.path(now)
    thumb_w, thumb_h, has_thumb = category.thumbnail()

    main_jobs: List[_Job] = []
    thumb_jobs: List[_Job] = []

    for i, f in enumerate(files):
        try:
            body = _to_bytes(f.body)
        except Exception as err:  # noqa: BLE001
            raise ValueError(f"yaylib: read file {i}: {err}") from err
        ext = normalize_image_ext(f.filename)
        # A failed decode means the size suffix would be wrong and the
        # body can't be thumbnailed; uploading it anyway just triggers
        # server-side moderation deletion, so surface the error.
        try:
            w, h = probe_image_size(body)
        except Exception as err:  # noqa: BLE001
            raise ValueError(f"yaylib: decode image {i}: {err}") from err
        main_name = image_filename(category_path, prefix, ts, i, ext, w, h)
        main_jobs.append(_Job(body, main_name, content_type_for(main_name)))

        if has_thumb:
            try:
                thumb = make_thumbnail_for(body, ext, thumb_w, thumb_h)
            except Exception as err:  # noqa: BLE001
                # A missing/mismatched thumbnail makes the server delete
                # the main asset too, so fail loudly.
                raise ValueError(
                    f"yaylib: make thumbnail for image {i}: {err}"
                ) from err
            thumb_name = thumbnail_filename(
                category_path, prefix, ts, i, thumb.ext, w, h
            )
            thumb_jobs.append(_Job(thumb.body, thumb_name, thumb.content_type))

    all_jobs = main_jobs + thumb_jobs
    all_file_names = [j.filename for j in all_jobs]

    try:
        urls = await deps.get_presigned_urls(all_file_names)
    except Exception as err:  # noqa: BLE001
        raise ValueError(f"yaylib: get presigned urls: {err}") from err
    if len(urls) != len(all_jobs):
        raise ValueError(
            f"yaylib: presigned urls count mismatch "
            f"(want {len(all_jobs)} got {len(urls)})"
        )

    async def _put(job: _Job, raw_url: str) -> None:
        if not raw_url:
            raise ValueError(f"yaylib: upload {job.filename}: empty presigned url")
        try:
            await deps.raw_put(raw_url, job.body, job.content_type)
        except ValueError:
            raise
        except Exception as err:  # noqa: BLE001
            raise ValueError(f"yaylib: upload {job.filename}: {err}") from err

    results = await asyncio.gather(
        *[_put(job, urls[i].url or "") for i, job in enumerate(all_jobs)],
        return_exceptions=True,
    )
    for r in results:
        if isinstance(r, BaseException):
            raise r

    # The server normalizes each filename (typically an "s3/" prefix
    # marking it as a freshly-uploaded object). That canonical value —
    # not the client-side path — is what edit_user / create_post expect.
    return [urls[i].filename or "" for i in range(len(files))]


async def upload_single_image(
    deps: UploadDeps, category: UploadCategory, file: Upload
) -> str:
    names = await upload_images(deps, category, [file])
    return names[0]


async def upload_video(deps: UploadDeps, body: Optional[Body]) -> str:
    """Upload a single MP4 stream and return the filename to pass back as
    video_file_name. The video presigned endpoint does not return a
    canonical filename, so the random name the SDK generates is what
    callers hand to create_post / send_chat_message (PORTING.md §8)."""
    if body is None:
        raise ValueError("yaylib: upload_video: body is None")
    data = _to_bytes(body)
    name = video_filename(random_filename_prefix(16), now_unix_millis())

    try:
        raw_url = await deps.get_video_presigned_url(name)
    except Exception as err:  # noqa: BLE001
        raise ValueError(f"yaylib: get presigned url: {err}") from err
    if not raw_url:
        raise ValueError("yaylib: empty presigned url")

    try:
        await deps.raw_put(raw_url, data, "video/mp4")
    except Exception as err:  # noqa: BLE001
        raise ValueError(f"yaylib: upload video: {err}") from err
    return name
