# Retry-policy behavior parity, driven against the shared
# behavior-contract server (PORTING.md §13 / §15). Translated from the
# Go reference's retry_test.go — the scenario engine enforces the
# failure pattern and the test asserts the observable outcome (final
# status, the echoed X-Mock-Attempt ordinal, elapsed time, cancellation)
# rather than a local request counter.
#
# These skip when YAYLIB_MOCK_BASE_URL is unset; the in-process
# equivalents live in test_retry.py and stay as the standalone fixtures.

import asyncio
import time

import pytest

from yaylib.retry import RetryPolicy

from ._parity import HDR_DIAG_ATTEMPT, HDR_DIAG_BRANCH, mock_get, mock_post

# Go's zero-value RetryPolicy{} (retry disabled). The Python default is
# max_attempts=3, so the "no retry" cases must request it explicitly.
NO_RETRY = RetryPolicy(max_attempts=0, base_delay=0.0, max_delay=0.0)


async def test_retries_on_5xx():
    c = mock_client_with("fail-503-times-2", max_attempts=5,
                          base_delay=0.001, max_delay=0.010)
    try:
        resp = await mock_get(c)
        assert resp.status == 200, "retried past the two 503s"
        assert resp.getheader(HDR_DIAG_ATTEMPT) == "3", (
            "two failures + the success"
        )
        assert resp.getheader(HDR_DIAG_BRANCH) == "fail-exhausted-happy"
    finally:
        await c.close()


async def test_retries_post_on_429():
    # retry_on_post=False (default): POST 5xx is not retried, but a 429
    # is retried regardless because the server explicitly asked and there
    # is no duplicate-creation risk.
    c = mock_client_with("fail-429-times-1", max_attempts=3, base_delay=0.001)
    try:
        resp = await mock_post(c)
        assert resp.status == 200
        assert resp.getheader(HDR_DIAG_ATTEMPT) == "2", (
            "POST 429 retried once"
        )
    finally:
        await c.close()


async def test_no_retry_post_by_default():
    # fail-503-times-1 would succeed on a retry, but POST 5xx is not
    # retried by default, so the single 503 surfaces.
    c = mock_client_with("fail-503-times-1", max_attempts=5, base_delay=0.001)
    try:
        resp = await mock_post(c)
        assert resp.status == 503, "POST not retried"
        assert resp.getheader(HDR_DIAG_ATTEMPT) == "1", "no retry"
    finally:
        await c.close()


async def test_retry_post_when_enabled():
    c = mock_client_with("fail-503-times-1", max_attempts=3,
                          base_delay=0.001, retry_on_post=True)
    try:
        resp = await mock_post(c)
        assert resp.status == 200
        assert resp.getheader(HDR_DIAG_ATTEMPT) == "2", (
            "POST retried when enabled"
        )
    finally:
        await c.close()


async def test_respects_max_attempts():
    c = mock_client_with("fail-503-times-99", max_attempts=4,
                          base_delay=0.001, max_delay=0.005)
    try:
        resp = await mock_get(c)
        assert resp.status == 503, "final response surfaced"
        assert resp.getheader(HDR_DIAG_ATTEMPT) == "4", (
            "stopped at max_attempts"
        )
    finally:
        await c.close()


async def test_disabled_by_zero_policy():
    c = mock_client_with_policy("fail-503-times-99", NO_RETRY)
    try:
        resp = await mock_get(c)
        assert resp.getheader(HDR_DIAG_ATTEMPT) == "1", (
            "zero policy disables retry"
        )
    finally:
        await c.close()


async def test_honors_retry_in_body():
    c = mock_client_with("retry-after-1", max_attempts=2,
                          base_delay=0.001, max_delay=5.0)
    try:
        start = time.monotonic()
        resp = await mock_get(c)
        elapsed = time.monotonic() - start
        assert resp.status == 200, "retried after the 429"
        # base/max delay are tiny; honoring retry_in (~1s) is the only
        # way the gap clears 0.9s.
        assert elapsed >= 0.9, f"elapsed={elapsed}, want >=1s"
    finally:
        await c.close()


async def test_retry_respects_cancellation():
    # Long backoff + a tight outer timeout: the cancellation must
    # propagate out of the retry sleep, not be swallowed.
    c = mock_client_with("fail-503-times-99", max_attempts=10,
                          base_delay=0.5, max_delay=5.0)
    try:
        with pytest.raises((asyncio.TimeoutError, asyncio.CancelledError)):
            await asyncio.wait_for(mock_get(c), timeout=0.1)
    finally:
        await c.close()


# ---- local helpers (kept out of _parity.py: RetryPolicy is the only
# knob these scenarios vary) ----


def mock_client_with(scenario, **policy_kwargs):
    from ._parity import mock_client

    return mock_client(scenario, retry_policy=RetryPolicy(**policy_kwargs))


def mock_client_with_policy(scenario, policy):
    from ._parity import mock_client

    return mock_client(scenario, retry_policy=policy)
