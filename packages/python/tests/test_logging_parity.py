# Logging redaction parity, driven against the shared behavior-contract
# server (PORTING.md §12.2 / §15). Translated from the Go reference's
# TestLogging_RefreshPersistFailEventAndRedaction: the expired-token
# scenario serves 401 until the token endpoint is hit, then a fresh
# token response, then the happy path — exactly the 401 -> refresh ->
# persist chain. With a save-failing store the persist branch fires, so
# the contracted events appear and we assert no banned value (seeded
# tokens, the server-issued fresh tokens, or the API key) reaches the
# logger.
#
# The "default client is silent" construction check stays in
# test_logging.py (a pure, network-free fixture). This skips when
# YAYLIB_MOCK_BASE_URL is unset.

from __future__ import annotations

import logging

from yaylib.retry import RetryPolicy

from ._parity import mock_client, mock_get

NO_RETRY = RetryPolicy(max_attempts=0, base_delay=0.0, max_delay=0.0)


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


# PORTING:S32
async def test_refresh_persist_fail_event_and_redaction():
    cap = _Capture()
    log = logging.getLogger("yaylib.test.parity")
    log.setLevel(logging.DEBUG)
    log.propagate = False
    log.handlers = [cap]

    c = mock_client(
        "expired-token",
        retry_policy=NO_RETRY,
        session_store=_FailingStore(),
        logger=log,
    )
    c.set_tokens("STALE", "REF")
    c._current_email = "foo@example.com"  # enable the persist branch
    try:
        resp = await mock_get(c)
        assert resp.status == 200, "refresh + retry succeeded"

        # Contracted events present.
        persist = cap.by_event("token_persist_fail")
        assert persist is not None and persist.levelno == logging.WARNING, (
            "token_persist_fail WARN record missing"
        )
        refresh = cap.by_event("token_refresh")
        assert refresh is not None and getattr(refresh, "outcome", None) == "ok"
        assert cap.by_event("http_request") is not None, (
            "no http_request debug record emitted"
        )

        # Redaction: no record (message or any attr) carries a secret —
        # neither the stale tokens we seeded nor the fresh ones the
        # server issued on refresh, nor the API key.
        fresh = c.tokens
        banned = [
            b
            for b in ("STALE", "REF", fresh.access, fresh.refresh, c.api_key)
            if b
        ]
        for r in cap.records:
            hay = r.getMessage() + " " + " ".join(
                f"{k}={v}" for k, v in r.__dict__.items()
            )
            for b in banned:
                assert b not in hay, (
                    f"record {r.getMessage()!r} leaked banned value {b!r}"
                )
    finally:
        await c.close()
