# Embedded wire constants. PORTING.md §3.1 fixes the defaults so a
# freshly-constructed Client speaks to the Yay! servers without any
# explicit option. Each constant is overridable via the corresponding
# Client kwarg.
#
# This module is hand-written (the generated layer owns ``configuration.py``;
# the underscore prefix keeps the two from colliding on import).

from __future__ import annotations

DEFAULT_API_KEY = (
    "ccd59ee269c01511ba763467045c115779fcae3050238a252f1bd1a4b65cfec6"
)

# HS256 key used by signed_version and X-Jwt.
DEFAULT_API_VERSION_KEY = "34c8c1cdf29b46a492e8ec58e6db37ec"

# Value sent as X-App-Version and embedded in signed_version.
DEFAULT_API_VERSION_NAME = "4.26"

# Build version embedded in X-Device-Info.
DEFAULT_APP_VERSION = "4.26.1"

DEFAULT_DEVICE_TYPE = "android"
DEFAULT_DEVICE_OS = "11"
DEFAULT_DEVICE_DENSITY = "3.5"
DEFAULT_DEVICE_SCREEN = "1440x2960"
DEFAULT_DEVICE_MODEL = "Galaxy S9"

DEFAULT_CONNECTION_TYPE = "wifi"
DEFAULT_CONNECTION_SPEED = "0 kbps"
DEFAULT_ACCEPT_LANGUAGE = "ja"

DEFAULT_BASE_URL = "https://api.yay.space"

# Auxiliary host a few activity-feed operations are served from
# (see yaylib._host_routes). Requests for those operations are routed
# here instead of DEFAULT_BASE_URL; override via the
# ``cassandra_base_url`` Client option (e.g. to point at a test server).
DEFAULT_CASSANDRA_BASE_URL = "https://cas.yay.space"

DEFAULT_EVENT_STREAM_URL = "wss://cable.yay.space"
MEDIA_CDN_BASE = "https://cdn.yay.space/uploads"

# Wire pepper concatenated after APIKey + DeviceUUID + timestamp before
# MD5-hashing for signed_info.
SIGNED_INFO_SHARED_KEY = "yayZ1"

# Platform identifier prepended to APIVersionName inside the
# signed_version HMAC payload. Hard-coded; deviating risks per-platform
# allowlist rejection.
SIGNED_VERSION_PLATFORM = "yay_android"


def build_user_agent(
    *,
    device_type: str,
    device_os: str,
    device_density: str,
    device_screen: str,
    device_model: str,
) -> str:
    return (
        f"{device_type} {device_os} "
        f"({device_density}x {device_screen} {device_model})"
    )


def build_device_info(app_version: str, user_agent: str) -> str:
    return f"yay {app_version} {user_agent}"


def media_url(filename: str) -> str:
    """Public CDN URL for a filename returned by any ``upload_*`` method.

    Prefer this over reading ``profile_icon`` / ``cover_image`` /
    ``attachment_url`` off API responses, which can carry a legacy
    prefix that no longer resolves.
    """
    if not filename:
        return ""
    return f"{MEDIA_CDN_BASE}/{filename.lstrip('/')}"
