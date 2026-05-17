# PORTING.md §6.1 contract verification against an in-process server.
# Mirrors tests/auth_refresh.test.ts: the three documented outcomes plus
# the concurrency collapse.

import asyncio
import json

import pytest

from yaylib.client import Client
from yaylib.errors import APIError

from ._server import serve


# PORTING:S7
async def test_refresh_success():
    oauth_hits = 0
    timeline_hits = 0

    def handler(path, method, body):
        nonlocal oauth_hits, timeline_hits
        if path == "/api/v1/oauth/token":
            oauth_hits += 1
            if "grant_type=refresh_token" not in body:
                return 400, '{"error":"bad_grant"}', {}
            return (
                200,
                json.dumps(
                    {
                        "access_token": "fresh-access",
                        "refresh_token": "fresh-refresh",
                        "user_id": 7,
                    }
                ),
                {},
            )
        if "/v2/users/timestamp" in path:
            timeline_hits += 1
            if timeline_hits == 1:
                return 401, '{"error_code":-1}', {}
            return (
                200,
                json.dumps({"time": 1700000000, "ip_address": "203.0.113.5"}),
                {},
            )
        return 404, "", {}

    async with serve(handler) as base_url:
        client = Client(base_url=base_url)
        client.set_tokens("stale-access", "stale-refresh")
        try:
            res = await client.users_api.get_user_timestamp()
            assert res.time == 1700000000
            assert oauth_hits == 1
            assert timeline_hits == 2
            assert client.tokens.access == "fresh-access"
            assert client.tokens.refresh == "fresh-refresh"
        finally:
            await client.close()


# PORTING:S8
async def test_refresh_fails_surfaces_original_401():
    def handler(path, method, body):
        if path == "/api/v1/oauth/token":
            return 401, '{"error_code":-100,"message":"refresh expired"}', {}
        if "/v2/users/timestamp" in path:
            return 401, '{"error_code":-1,"message":"unauthorized"}', {}
        return 404, "", {}

    async with serve(handler) as base_url:
        client = Client(base_url=base_url)
        client.set_tokens("stale-access", "stale-refresh")
        try:
            with pytest.raises(APIError) as ei:
                await client.users_api.get_user_timestamp()
            assert ei.value.status == 401
            assert client.tokens.access == "stale-access"
        finally:
            await client.close()


# PORTING:S8
async def test_no_refresh_token_skips_refresh():
    oauth_hits = 0

    def handler(path, method, body):
        nonlocal oauth_hits
        if path == "/api/v1/oauth/token":
            oauth_hits += 1
            return 200, "{}", {}
        if "/v2/users/timestamp" in path:
            return 401, '{"error_code":-1}', {}
        return 404, "", {}

    async with serve(handler) as base_url:
        client = Client(base_url=base_url)
        client.set_tokens("only-access", "")
        try:
            with pytest.raises(APIError) as ei:
                await client.users_api.get_user_timestamp()
            assert ei.value.status == 401
            assert oauth_hits == 0
        finally:
            await client.close()


# PORTING:S11
async def test_concurrent_401s_collapse_to_one_refresh():
    oauth_hits = 0
    timeline_hits = 0

    def handler(path, method, body):
        nonlocal oauth_hits, timeline_hits
        if path == "/api/v1/oauth/token":
            oauth_hits += 1
            if "grant_type=refresh_token" not in body:
                return 400, '{"error":"bad_grant"}', {}
            return (
                200,
                json.dumps(
                    {
                        "access_token": "fresh",
                        "refresh_token": "fresh-r",
                        "user_id": 7,
                    }
                ),
                {},
            )
        if "/v2/users/timestamp" in path:
            timeline_hits += 1
            if timeline_hits <= 3:
                return 401, '{"error_code":-1}', {}
            return (
                200,
                json.dumps({"time": 1700000001, "ip_address": "203.0.113.5"}),
                {},
            )
        return 404, "", {}

    async with serve(handler) as base_url:
        client = Client(base_url=base_url)
        client.set_tokens("stale", "stale-r")
        try:
            results = await asyncio.gather(
                client.users_api.get_user_timestamp(),
                client.users_api.get_user_timestamp(),
                client.users_api.get_user_timestamp(),
            )
            assert oauth_hits == 1
            assert all(r.time == 1700000001 for r in results)
        finally:
            await client.close()
