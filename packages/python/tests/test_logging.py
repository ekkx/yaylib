"""Logging parity (mirrors Go logging_test.go / TS logging.test.ts).

PORTING.md §12.2 + §15: the SDK is silent by default; with a logger
injected the 401 -> refresh -> persist-failure path emits the
contracted events and never leaks a token / password / API key.
"""

from __future__ import annotations

import json
import logging

import pytest

from tests._server import serve
from yaylib.client import Client


class _Capture(logging.Handler):
    def __init__(self) -> None:
        super().__init__(level=logging.DEBUG)
        self.records: list[logging.LogRecord] = []

    def emit(self, record: logging.LogRecord) -> None:
        self.records.append(record)

    def by_event(self, event: str):
        for r in self.records:
            if getattr(r, "event", None) == event:
                return r
        return None


class _FailingStore:
    async def load(self, email):  # noqa: ANN001
        raise RuntimeError("no session")

    async def save(self, session):  # noqa: ANN001
        raise RuntimeError("disk full")

    async def delete(self, email):  # noqa: ANN001
        return None


def _handler():
    def handler(path, method, body):
        if path == "/api/v1/oauth/token":
            return (
                200,
                json.dumps({"access_token": "NEW_ACC", "refresh_token": "NEW_REF"}),
                {},
            )
        if "/v2/users/timestamp" in path:
            _handler.hits += 1  # type: ignore[attr-defined]
            if _handler.hits == 1:  # type: ignore[attr-defined]
                return 401, '{"error_code":-1}', {}
            return 200, json.dumps({"time": 1700000000, "ip_address": "1.2.3.4"}), {}
        return 404, "", {}

    _handler.hits = 0  # type: ignore[attr-defined]
    return handler


async def _run(client: Client) -> None:
    client._current_email = "foo@example.com"  # enable the persist branch
    client.set_tokens("STALE_ACC", "STALE_REF")
    try:
        await client.users_api.get_user_timestamp()
    finally:
        await client.close()


@pytest.mark.asyncio
async def test_default_client_is_silent():
    # The default logger discards everything (NullHandler) and does not
    # propagate, so a library consumer that never opts in sees nothing.
    c = Client()
    lg = c._logger
    assert lg.propagate is False
    assert lg.handlers and all(
        isinstance(h, logging.NullHandler) for h in lg.handlers
    )


@pytest.mark.asyncio
async def test_logger_events_and_redaction():
    cap = _Capture()
    log = logging.getLogger("yaylib.test.parity")
    log.setLevel(logging.DEBUG)
    log.propagate = False
    log.handlers = [cap]

    async with serve(_handler()) as base_url:
        await _run(Client(base_url=base_url, session_store=_FailingStore(), logger=log))

    persist = cap.by_event("token_persist_fail")
    assert persist is not None and persist.levelno == logging.WARNING
    refresh = cap.by_event("token_refresh")
    assert refresh is not None and getattr(refresh, "outcome", None) == "ok"
    assert cap.by_event("http_request") is not None

    banned = ["STALE_ACC", "STALE_REF", "NEW_ACC", "NEW_REF"]
    for r in cap.records:
        hay = r.getMessage() + " " + " ".join(
            f"{k}={v}" for k, v in r.__dict__.items()
        )
        for b in banned:
            assert b not in hay, f"record {r.getMessage()!r} leaked {b}"
