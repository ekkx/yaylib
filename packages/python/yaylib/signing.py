# PORTING.md §7: signed_info / signed_version / X-Jwt.
#
# All three helpers produce byte-identical output to the Go reference
# against the test vectors pinned in PORTING.md §7.1. Python targets
# server-side runtimes (PORTING.md "Runtime scope"), so ``hashlib.md5``
# is always available — no polyfill needed.

from __future__ import annotations

import base64
import hashlib
import hmac
import time
from dataclasses import dataclass

from yaylib._config import SIGNED_INFO_SHARED_KEY, SIGNED_VERSION_PLATFORM


@dataclass(frozen=True)
class SignedInfo:
    """Bundles the timestamp and MD5 hash a Yay! request needs to pass
    server-side validation. The two values are bound: changing one
    without the other invalidates the signature.
    """

    # Unix-second timestamp the hash is computed against. Send this as
    # the request's ``timestamp`` field.
    timestamp: int
    # Lowercase 32-char MD5 hex digest. Send this as ``signed_info``.
    value: str


def generate_signed_info_at(
    *, api_key: str, device_uuid: str, timestamp: int
) -> SignedInfo:
    """Compute the signed_info MD5 for ``timestamp`` with no network I/O.

    Use this when you already have a synchronized timestamp; otherwise
    prefer ``Client.generate_signed_info`` which fetches the server's
    view of the clock to avoid device-clock drift.
    """
    payload = f"{api_key}{device_uuid}{timestamp}{SIGNED_INFO_SHARED_KEY}"
    value = hashlib.md5(payload.encode("utf-8")).hexdigest()
    return SignedInfo(timestamp=timestamp, value=value)


def generate_signed_version(
    *, api_version_key: str, api_version_name: str
) -> str:
    """The signed_version HMAC token some endpoints require alongside
    signed_info. The platform identifier is hard-coded to ``yay_android``
    to stay inside the server's platform allowlist.
    """
    payload = f"{SIGNED_VERSION_PLATFORM}/{api_version_name}"
    digest = hmac.new(
        api_version_key.encode("utf-8"),
        payload.encode("utf-8"),
        hashlib.sha256,
    ).digest()
    # Standard base64 with padding; no trailing newline (stdlib does not
    # add one, but be explicit about the contract).
    return base64.b64encode(digest).decode("ascii")


def generate_x_jwt(
    *, api_version_key: str, now: int | None = None, ttl_seconds: int = 5
) -> str:
    """The short-lived HS256 JWT required by some write endpoints (most
    visibly create_post). ``iat`` is the current Unix-second time;
    ``exp`` = iat + ttl_seconds; the HS256 key is APIVersionKey. TTL is
    5 seconds, so callers MUST regenerate per request.
    """
    iat = int(time.time()) if now is None else now
    exp = iat + ttl_seconds
    return _sign_jwt(api_version_key, iat, exp)


def _b64url_no_pad(raw: bytes) -> str:
    # RFC 4648 §5, padding stripped.
    return base64.urlsafe_b64encode(raw).rstrip(b"=").decode("ascii")


def _sign_jwt(key: str, iat: int, exp: int) -> str:
    # Header JSON has no ``typ`` and no spaces; payload JSON key order is
    # ``iat`` then ``exp``, also no spaces.
    header = _b64url_no_pad(b'{"alg":"HS256"}')
    payload = _b64url_no_pad(
        f'{{"iat":{iat},"exp":{exp}}}'.encode("utf-8")
    )
    signing_input = f"{header}.{payload}"
    sig = hmac.new(
        key.encode("utf-8"), signing_input.encode("utf-8"), hashlib.sha256
    ).digest()
    return f"{signing_input}.{_b64url_no_pad(sig)}"
