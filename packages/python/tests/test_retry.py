# RetryPolicy parity tests (PORTING.md §13). Mirrors tests/retry.test.ts:
# 429 retries everywhere, 5xx retries only on safe methods (or POST when
# opted in), body retry_in honored, max-attempts cap, zero-value
# disables retry entirely.

import json
import time

import pytest

from yaylib.client import Client
from yaylib.errors import APIError
from yaylib.retry import RetryPolicy

from ._server import serve

# Tight backoff so tests stay snappy (seconds).
FAST = RetryPolicy(
    max_attempts=3, base_delay=0.001, max_delay=0.005, retry_on_post=False
)


async def test_get_5xx_retries_until_success():
    hits = 0

    def handler(path, method, body):
        nonlocal hits
        hits += 1
        if hits < 3:
            return 503, '{"error":"transient"}', {}
        return 200, json.dumps({"time": 1700000000, "ip_address": "1.2.3.4"}), {}

    async with serve(handler) as base_url:
        client = Client(base_url=base_url, retry_policy=FAST)
        try:
            res = await client.users_api.get_user_timestamp()
            assert res.time == 1700000000
            assert hits == 3
        finally:
            await client.close()


async def test_post_5xx_not_retried_by_default():
    hits = 0

    def handler(path, method, body):
        nonlocal hits
        hits += 1
        return 503, '{"error":"transient"}', {}

    async with serve(handler) as base_url:
        client = Client(base_url=base_url, retry_policy=FAST)
        try:
            with pytest.raises(APIError):
                await client.users_api.block_user(id=1)
            assert hits == 1
        finally:
            await client.close()


async def test_post_5xx_retries_when_opted_in():
    hits = 0

    def handler(path, method, body):
        nonlocal hits
        hits += 1
        if hits < 2:
            return 503, '{"error":"transient"}', {}
        return 200, '{"success":true}', {}

    policy = RetryPolicy(
        max_attempts=3, base_delay=0.001, max_delay=0.005, retry_on_post=True
    )
    async with serve(handler) as base_url:
        client = Client(base_url=base_url, retry_policy=policy)
        try:
            try:
                await client.users_api.block_user(id=1)
            except APIError:
                pass
            assert hits == 2
        finally:
            await client.close()


async def test_429_retries_on_post_even_without_opt_in():
    hits = 0

    def handler(path, method, body):
        nonlocal hits
        hits += 1
        if hits < 2:
            return 429, '{"error":"rate_limited"}', {}
        return 200, '{"success":true}', {}

    async with serve(handler) as base_url:
        client = Client(base_url=base_url, retry_policy=FAST)
        try:
            try:
                await client.users_api.block_user(id=1)
            except APIError:
                pass
            assert hits == 2
        finally:
            await client.close()


async def test_retry_in_body_honored():
    hits = 0
    first_at = 0.0
    second_at = 0.0

    def handler(path, method, body):
        nonlocal hits, first_at, second_at
        hits += 1
        now = time.monotonic()
        if hits == 1:
            first_at = now
            return 503, json.dumps({"error": "wait", "retry_in": 0.05}), {}
        second_at = now
        return 200, json.dumps({"time": 1700000000, "ip_address": "1.2.3.4"}), {}

    async with serve(handler) as base_url:
        client = Client(base_url=base_url, retry_policy=FAST)
        try:
            res = await client.users_api.get_user_timestamp()
            assert res.time == 1700000000
            # base/max delay are <= 5ms; honoring retry_in (~50ms) is the
            # only way the gap clears 30ms.
            assert (second_at - first_at) >= 0.03
        finally:
            await client.close()


async def test_max_attempts_reached_surfaces_last():
    hits = 0

    def handler(path, method, body):
        nonlocal hits
        hits += 1
        return 503, '{"error":"perma"}', {}

    async with serve(handler) as base_url:
        client = Client(base_url=base_url, retry_policy=FAST)
        try:
            with pytest.raises(APIError) as ei:
                await client.users_api.get_user_timestamp()
            assert hits == 3
            assert ei.value.status == 503
        finally:
            await client.close()


async def test_zero_policy_disables_retry():
    hits = 0

    def handler(path, method, body):
        nonlocal hits
        hits += 1
        return 503, '{"error":"perma"}', {}

    policy = RetryPolicy(
        max_attempts=0, base_delay=0.001, max_delay=0.001, retry_on_post=False
    )
    async with serve(handler) as base_url:
        client = Client(base_url=base_url, retry_policy=policy)
        try:
            with pytest.raises(APIError):
                await client.users_api.get_user_timestamp()
            assert hits == 1
        finally:
            await client.close()
