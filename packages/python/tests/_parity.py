# Behavior-parity harness (PORTING.md §15).
#
# These helpers drive the wrapper's retry / refresh / rate-limit /
# event-stream behavior against a shared behavior-contract server — the
# same instance the Go and TypeScript suites use — so the three
# languages verify identical behavior instead of three divergent
# in-process fakes. Translated from the Go reference's parity_test.go;
# the scenario grammar and header names are byte-identical across ports.
#
# The server is discovered from the environment. When the variables are
# unset every parity test skips, so a standalone ``pytest`` still passes
# on the in-process fixtures alone.
#
#   YAYLIB_MOCK_BASE_URL   http://host:port   (REST)
#   YAYLIB_MOCK_WS_URL     ws://host:port     (event stream)
#
# Behavior is selected per request with two headers the server reads:
# X-Yay-Mock-Scenario picks the failure/latency behavior and
# X-Yay-Mock-Session isolates each test's stateful counters. The server
# echoes diagnostics back on every response it owns: X-Mock-Branch (the
# decision taken) and X-Mock-Attempt (the per-session attempt ordinal
# for counter-based scenarios).
#
# Seam (PORTING.md §6.1 / §12): the SDK injects its own headers, runs
# the 401 refresh-and-replay and the retry loop, and dials the
# WebSocket — all through one ``aiohttp.ClientSession`` the Transport
# owns. We stamp the two scenario headers as that session's *default*
# headers, which aiohttp applies last (request-level headers never
# collide with them) to every real send: the first attempt, the
# token-refresh round-trip, every retry replay, and the WS upgrade.
# That mirrors the Go harness installing ``scenarioRT`` as the base
# transport hop. Re-sends ride this same bare session rather than
# re-entering the wrapped client, so there is no refresh/retry
# recursion (PORTING.md §6.1).

from __future__ import annotations

import os
import secrets

import aiohttp
import pytest

from yaylib.client import Client

ENV_MOCK_BASE_URL = "YAYLIB_MOCK_BASE_URL"
ENV_MOCK_WS_URL = "YAYLIB_MOCK_WS_URL"

HDR_SCENARIO = "X-Yay-Mock-Scenario"
HDR_SESSION = "X-Yay-Mock-Session"

HDR_DIAG_BRANCH = "X-Mock-Branch"
HDR_DIAG_ATTEMPT = "X-Mock-Attempt"

# Real spec routes used to drive generic transport behavior. The happy
# fall-through needs a route the server knows so it answers 200; failure
# scenarios are path-independent. Neither is host-routed, so requests
# stay on the primary base URL.
PARITY_GET_PATH = "/v1/buckets/presigned_urls"
PARITY_POST_PATH = "/v1/calls/leave_agora_channel"


def mock_base_url() -> str:
    return os.environ.get(ENV_MOCK_BASE_URL, "")


def mock_ws_url() -> str:
    return os.environ.get(ENV_MOCK_WS_URL, "")


def require_mock() -> str:
    """Skip the calling test unless the shared server is configured;
    return its REST base URL.
    """
    base = mock_base_url()
    if not base:
        pytest.skip(f"{ENV_MOCK_BASE_URL} unset; skipping behavior-parity test")
    return base


def new_mock_session() -> str:
    """An opaque per-test key so the server's stateful scenario counters
    never collide across tests (or across the parallel three-language
    run).
    """
    return secrets.token_hex(16)


def _install_scenario(client: Client, scenario: str, session_id: str) -> None:
    """Stamp the scenario + session headers on every outbound request by
    making them the Transport session's default headers (see the module
    docstring for why this is the correct seam).
    """
    transport = client._transport

    def _ensure() -> aiohttp.ClientSession:
        if transport._session is None or transport._session.closed:
            hdrs = {HDR_SESSION: session_id}
            if scenario:
                hdrs[HDR_SCENARIO] = scenario
            transport._session = aiohttp.ClientSession(headers=hdrs)
        return transport._session

    transport._ensure_session = _ensure  # type: ignore[method-assign]


def mock_client(scenario: str, **kwargs) -> Client:
    """A Client pointed at the shared server with ``scenario`` active for
    this test's isolated session. Extra kwargs are forwarded to Client
    so a test can override (e.g. ``retry_policy=``).
    """
    base = require_mock()
    ws = mock_ws_url()
    if ws and "event_stream_url" not in kwargs:
        kwargs["event_stream_url"] = ws
    c = Client(base_url=base, **kwargs)
    # Keep the lazy X-Client-IP fetch off the test path — otherwise the
    # background timestamp request would also carry the scenario header
    # and consume the counter (Go pins this with WithClientIP).
    c._client_ip = "127.0.0.1"
    _install_scenario(c, scenario, new_mock_session())
    return c


def mock_stream_client(mode: str, **kwargs) -> Client:
    """A Client whose event stream targets the shared server's WebSocket
    with the given per-dial mode: "" confirms each subscribe and pushes
    the channel's representative event; "reject" rejects every
    subscribe; "no-confirm" stays silent (drives the subscribe
    timeout); "drop-after-confirm" confirms + pushes then closes (drives
    reconnect / re-subscribe). The mode rides the event-stream URL
    query, which the SDK preserves while appending the token and
    app_version.
    """
    require_mock()
    ws = mock_ws_url()
    if not ws:
        pytest.skip(f"{ENV_MOCK_WS_URL} unset; skipping event-stream parity test")
    if mode:
        ws += "/?mock=" + mode
    kwargs["event_stream_url"] = ws
    return mock_client("", **kwargs)


async def mock_get(client: Client):
    """Drive a GET against the shared server's generic primary route
    through the SDK transport. The returned BufferedResponse keeps its
    headers (including the X-Mock-* diagnostics) for assertion.
    """
    return await client._transport.request(
        "GET", mock_base_url() + PARITY_GET_PATH
    )


async def mock_post(client: Client):
    """Drive a POST against the shared server's generic primary route."""
    return await client._transport.request(
        "POST", mock_base_url() + PARITY_POST_PATH, body="payload"
    )
