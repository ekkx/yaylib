# PORTING.md §12 / §6.1 / §13: the transport replaces the generated
# ``rest_client``. It injects the headers the Yay! server validates,
# picks the right Authorization scheme, runs the 401 auto-refresh chain,
# and applies the retry policy — all funnelled through one ``request``
# coroutine the generated ``ApiClient`` already calls.
#
# Re-sends (the post-401 replay and the transient-failure retry) go
# through the same bare ``aiohttp`` session used for the first attempt,
# NOT back through a wrapped client — there is no recursion hazard here
# because this object is the bottom of the stack (PORTING.md §6.1 note).

from __future__ import annotations

import asyncio
import base64
import json
import re
import time
from dataclasses import dataclass
from typing import Awaitable, Callable, Optional

from urllib.parse import urlsplit, urlunsplit

import aiohttp
from multidict import CIMultiDict

from yaylib._host_routes import HOST_ROUTES
from yaylib.exceptions import ApiException, ApiValueError
from yaylib.retry import (
    RetryPolicy,
    full_jitter_delay,
    method_allows_retry,
    should_retry_status,
)

_OAUTH_TOKEN_PATH = "/api/v1/oauth/token"
_TIMESTAMP_PATH = "/v2/users/timestamp"


def _is_oauth_token_path(url: str) -> bool:
    return _OAUTH_TOKEN_PATH in url


def _is_timestamp_path(url: str) -> bool:
    # The endpoint used to look up the caller's IP for X-Client-IP. The
    # lazy fetch must not re-trigger itself when calling it.
    path = url.split("?", 1)[0]
    return path.endswith(_TIMESTAMP_PATH)


def _resolve_aux_host(alias: str, ctx: "TransportContext") -> str:
    # Unrecognized alias from the generated table: keep the request on
    # the primary host rather than stranding it.
    if alias == "cassandra":
        return ctx.cassandra_base_url
    return ""


def _route_host(method: str, url: str, ctx: "TransportContext") -> str:
    """Rewrite the request origin when the operation is served from an
    auxiliary host (see HOST_ROUTES). Match is by exact "METHOD path" —
    the host-routed operations are all parameter-free, so the request
    path equals the OpenAPI path template. Unlisted requests are
    returned unchanged.
    """
    parts = urlsplit(url)
    alias = HOST_ROUTES.get(f"{method.upper()} {parts.path}")
    if not alias:
        return url
    base = _resolve_aux_host(alias, ctx)
    if not base:
        return url
    target = urlsplit(base)
    if not target.netloc:
        return url
    return urlunsplit(
        (target.scheme, target.netloc, parts.path, parts.query, parts.fragment)
    )


@dataclass
class TransportContext:
    user_agent: str
    app_version: str  # sent as X-App-Version
    device_info: str
    device_uuid: str
    connection_type: str
    connection_speed: str
    accept_language: str
    api_key: str
    client_ip: Callable[[], str]
    access_token: Callable[[], str]
    # Auxiliary host for activity-feed operations (see
    # yaylib._host_routes). The transport rewrites the request origin to
    # this for host-routed operations; everything else stays on the
    # primary host.
    cassandra_base_url: str = ""


# RefreshFn attempts a one-shot refresh of the access token using the
# stale value carried by the original 401 request. Implementations MUST
# collapse concurrent calls (PORTING.md §6.1). Returns True when the
# access token is now fresh (retry the original request), False when
# refresh failed (surface the original 401 to the caller).
RefreshFn = Callable[[str], Awaitable[bool]]


class BufferedResponse:
    """RESTResponse-compatible wrapper. The body is read fully up-front
    so the response is decoupled from the connection — required for the
    retry path (we discard intermediate responses) and so the session
    can close without keep-alive sockets hanging.
    """

    def __init__(self, status: int, reason: Optional[str], headers, data: bytes):
        self.status = status
        self.reason = reason
        self._headers = headers
        self.data = data

    async def read(self) -> bytes:
        return self.data

    def getheaders(self) -> "CIMultiDict[str]":
        return self._headers

    def getheader(self, name: str, default=None):
        return self._headers.get(name, default)


def _find_header(headers: dict, name: str) -> Optional[str]:
    lname = name.lower()
    for k in headers:
        if k.lower() == lname:
            return k
    return None


def _set_if_absent(headers: dict, key: str, value: str) -> None:
    if not value:
        return
    if _find_header(headers, key) is None:
        headers[key] = value


def _timeout(_request_timeout) -> aiohttp.ClientTimeout:
    if _request_timeout is None:
        return aiohttp.ClientTimeout(total=5 * 60)
    if isinstance(_request_timeout, (int, float)):
        return aiohttp.ClientTimeout(total=float(_request_timeout))
    if isinstance(_request_timeout, (tuple, list)) and len(_request_timeout) == 2:
        return aiohttp.ClientTimeout(
            sock_connect=float(_request_timeout[0]),
            sock_read=float(_request_timeout[1]),
        )
    return aiohttp.ClientTimeout(total=5 * 60)


class Transport:
    """Drop-in replacement for the generated ``RESTClientObject``."""

    def __init__(
        self,
        *,
        ctx: TransportContext,
        refresh: RefreshFn,
        policy: RetryPolicy,
        on_response: Optional[Callable[[], None]] = None,
    ) -> None:
        self._ctx = ctx
        self._refresh = refresh
        self._policy = policy
        # Fired after any non-timestamp request gets a response — drives
        # the Client's lazy X-Client-IP fetch (PORTING.md §12).
        self._on_response = on_response
        self._session: Optional[aiohttp.ClientSession] = None

    def _ensure_session(self) -> aiohttp.ClientSession:
        # Created lazily inside the running loop — Client is built
        # synchronously and aiohttp refuses a session without a loop.
        if self._session is None or self._session.closed:
            self._session = aiohttp.ClientSession()
        return self._session

    async def close(self) -> None:
        if self._session is not None and not self._session.closed:
            await self._session.close()
            self._session = None

    # ---- body serialization (mirrors the generated rest.py) ----

    @staticmethod
    def _serialize_body(headers: dict, body, post_params):
        if post_params and body:
            raise ApiValueError(
                "body parameter cannot be used with post_params parameter."
            )
        ct_key = _find_header(headers, "content-type")
        content_type = headers[ct_key] if ct_key else "application/json"
        if re.search("json", content_type, re.IGNORECASE):
            return json.dumps(body) if body is not None else None
        if content_type == "application/x-www-form-urlencoded":
            return aiohttp.FormData(post_params or [])
        if content_type == "multipart/form-data":
            data = aiohttp.FormData()
            for param in post_params or []:
                k, v = param
                if isinstance(v, tuple) and len(v) == 3:
                    data.add_field(
                        k, value=v[1], filename=v[0], content_type=v[2]
                    )
                else:
                    if isinstance(v, dict):
                        v = json.dumps(v)
                    elif isinstance(v, int):
                        v = str(v)
                    data.add_field(k, v)
            return data
        if isinstance(body, (str, bytes)):
            return body
        if body is None:
            return None
        raise ApiException(
            status=0,
            reason="Cannot prepare a request message for provided arguments.",
        )

    # ---- header injection (PORTING.md §12) ----

    def _build_headers(self, base: dict, url: str) -> dict:
        headers = dict(base)
        ctx = self._ctx
        _set_if_absent(headers, "User-Agent", ctx.user_agent)
        _set_if_absent(headers, "X-App-Version", ctx.app_version)
        _set_if_absent(headers, "X-Device-Info", ctx.device_info)
        _set_if_absent(headers, "X-Device-UUID", ctx.device_uuid)
        _set_if_absent(headers, "X-Client-IP", ctx.client_ip())
        _set_if_absent(headers, "X-Connection-Type", ctx.connection_type)
        _set_if_absent(headers, "X-Connection-Speed", ctx.connection_speed)
        _set_if_absent(headers, "Accept-Language", ctx.accept_language)

        # X-Timestamp must be fresh per request (never cached).
        ts_key = _find_header(headers, "x-timestamp")
        if ts_key:
            del headers[ts_key]
        headers["X-Timestamp"] = str(int(time.time()))

        # Normalize JSON Content-Type to the exact casing / spacing the
        # server expects.
        ct_key = _find_header(headers, "content-type")
        if ct_key is None:
            headers["Content-Type"] = "application/json;charset=UTF-8"
        elif headers[ct_key].lower().startswith("application/json"):
            headers[ct_key] = "application/json;charset=UTF-8"

        # Authorization: Basic for oauth/token*, Bearer for everything
        # else. Caller-provided Authorization wins (set-if-absent).
        if _find_header(headers, "authorization") is None:
            if _is_oauth_token_path(url):
                token = base64.b64encode(
                    ctx.api_key.encode("utf-8")
                ).decode("ascii")
                headers["Authorization"] = f"Basic {token}"
            else:
                access = ctx.access_token()
                if access:
                    headers["Authorization"] = f"Bearer {access}"
        return headers

    async def _raw(
        self, method: str, url: str, headers: dict, data, timeout
    ) -> BufferedResponse:
        session = self._ensure_session()
        async with session.request(
            method, url, headers=headers, data=data, timeout=timeout
        ) as resp:
            body = await resp.read()
            return BufferedResponse(
                resp.status,
                resp.reason,
                CIMultiDict(resp.headers),
                body,
            )

    async def send_unwrapped(
        self, method: str, url: str, headers: dict, data, timeout=None
    ) -> BufferedResponse:
        """Bare round-trip with no header injection / refresh / retry.

        Used by the OAuth refresh path (PORTING.md §6.1: re-issue with the
        UNWRAPPED HTTP client) so it can never recurse back into the
        refresh hook.
        """
        return await self._raw(
            method.upper(), url, headers, data, _timeout(timeout)
        )

    @staticmethod
    def _retry_in_seconds(resp: BufferedResponse) -> Optional[float]:
        if not resp.data:
            return None
        try:
            parsed = json.loads(resp.data)
        except (ValueError, TypeError):
            return None
        if not isinstance(parsed, dict):
            return None
        v = parsed.get("retry_in")
        if isinstance(v, (int, float)) and not isinstance(v, bool) and v >= 0:
            return float(v)
        return None

    async def request(
        self,
        method,
        url,
        headers=None,
        body=None,
        post_params=None,
        _request_timeout=None,
    ) -> BufferedResponse:
        method = method.upper()
        # A few operations are served from an auxiliary host. Rewrite the
        # origin before anything else so header injection, the 401
        # refresh-and-replay, and retries all act on the final URL.
        url = _route_host(method, url, self._ctx)
        base_headers = dict(headers or {})
        data = self._serialize_body(base_headers, body, post_params)
        timeout = _timeout(_request_timeout)

        policy = self._policy
        retry_enabled = policy.max_attempts > 1
        is_oauth = _is_oauth_token_path(url)

        attempt = 1  # number of sends performed (initial counts as 1)
        did_refresh = False

        while True:
            send_headers = self._build_headers(base_headers, url)
            auth = send_headers.get(_find_header(send_headers, "authorization") or "")
            stale_access = (
                auth[7:] if isinstance(auth, str) and auth.startswith("Bearer ")
                else ""
            )

            try:
                resp = await self._raw(
                    method, url, send_headers, data, timeout
                )
            except asyncio.CancelledError:
                raise
            except aiohttp.ClientError as exc:
                if (
                    retry_enabled
                    and method_allows_retry(method, policy)
                    and attempt < policy.max_attempts
                ):
                    delay = min(
                        full_jitter_delay(attempt, policy), policy.max_delay
                    )
                    await asyncio.sleep(delay)
                    attempt += 1
                    continue
                raise ApiException(status=0, reason=str(exc)) from exc

            # Kick the lazy X-Client-IP fetch on the first non-timestamp
            # request (PORTING.md §12). Skip the fetch endpoint itself to
            # avoid recursion; the Client side single-flights / no-ops
            # when the IP is already known.
            if self._on_response is not None and not _is_timestamp_path(url):
                self._on_response()

            # 401 auto-refresh (PORTING.md §6.1). The replay after a
            # successful refresh is OUTSIDE the retry budget — it mirrors
            # the Go/TS split where refresh and retry are separate hooks.
            if resp.status == 401 and not is_oauth and not did_refresh:
                if await self._refresh(stale_access):
                    did_refresh = True
                    continue
                return resp

            if (
                retry_enabled
                and should_retry_status(resp.status, method, policy)
                and attempt < policy.max_attempts
            ):
                body_wait = self._retry_in_seconds(resp)
                delay = (
                    body_wait
                    if body_wait is not None
                    else full_jitter_delay(attempt, policy)
                )
                await asyncio.sleep(delay)
                attempt += 1
                continue

            return resp
