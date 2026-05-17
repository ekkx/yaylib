# PORTING.md §9: three layers of error access — the wire-level APIError,
# the parsed ErrorResponse body, and the typed error code for match /
# dispatch.
#
# The generated layer raises ``ApiException`` (and the status-specific
# subclasses ``UnauthorizedException`` / ``BadRequestException`` / ...)
# carrying the raw body + status. ``APIError`` is the stable, documented
# name the SDK promises callers — it aliases the generated base so
# ``except APIError`` recovers every wire error regardless of which
# subclass the generated code threw.

from __future__ import annotations

import json
from dataclasses import dataclass
from typing import Optional

from yaylib.exceptions import ApiException as APIError
from yaylib import _error_codes
from yaylib._error_codes import *  # noqa: F401,F403 — generated ErrCode* constants

__all__ = [
    "APIError",
    "ErrorResponse",
    "error_response_of",
    "code_of",
    "as_api_error",
] + list(_error_codes.__all__)


@dataclass
class ErrorResponse:
    """Parsed shape of the Yay! server's JSON error body. Every field is
    optional because the server emits a subset depending on the code.
    """

    error_code: Optional[int] = None
    message: Optional[str] = None
    ban_until: Optional[int] = None
    retry_in: Optional[int] = None
    remaining_quota: Optional[int] = None


def _raw_body(err: object) -> Optional[str]:
    if not isinstance(err, APIError):
        return None
    body = getattr(err, "body", None)
    if isinstance(body, (bytes, bytearray)):
        try:
            return body.decode("utf-8")
        except Exception:
            return None
    if isinstance(body, str) and body:
        return body
    # Some code paths leave the parsed object on ``data`` instead.
    data = getattr(err, "data", None)
    if isinstance(data, str) and data:
        return data
    return None


def error_response_of(err: object) -> Optional[ErrorResponse]:
    """Return the parsed payload when ``err`` is an APIError with a JSON
    body, otherwise ``None`` (never raises).
    """
    text = _raw_body(err)
    if not text:
        return None
    try:
        parsed = json.loads(text)
    except (ValueError, TypeError):
        return None
    if not isinstance(parsed, dict):
        return None

    def _int(key: str) -> Optional[int]:
        v = parsed.get(key)
        return v if isinstance(v, int) and not isinstance(v, bool) else None

    msg = parsed.get("message")
    return ErrorResponse(
        error_code=_int("error_code"),
        message=msg if isinstance(msg, str) else None,
        ban_until=_int("ban_until"),
        retry_in=_int("retry_in"),
        remaining_quota=_int("remaining_quota"),
    )


def code_of(err: object) -> int:
    """Surface the server's ``error_code`` for match dispatch. Returns 0
    (the unknown / zero-ish code) when ``err`` is not an APIError or
    carries no recognizable code. Never raises.
    """
    r = error_response_of(err)
    if r is None or r.error_code is None:
        return 0
    return r.error_code


def as_api_error(err: BaseException) -> APIError:
    """Normalize an arbitrary exception into the stable APIError shape.

    A generated ``ApiException`` (or subclass) passes through unchanged so
    its body / status stay intact; anything else (e.g. an aiohttp
    transport failure) is wrapped as a status-0 APIError.
    """
    if isinstance(err, APIError):
        return err
    return APIError(status=0, reason=str(err))
