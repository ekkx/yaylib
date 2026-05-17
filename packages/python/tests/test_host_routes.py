"""Host-routing parity (mirrors the Go host_routes_test.go and the TS
host_routes.test.ts). A few operations are served from an auxiliary
host; the transport must send those — and only those — there, leaving
every other request on the primary host. Two in-process servers pin the
routing offline.
"""

from __future__ import annotations

import json

import pytest

from tests._server import serve
from yaylib.client import Client


@pytest.mark.asyncio
async def test_activity_feed_routed_to_auxiliary_host():
    primary_hits: list[str] = []
    aux_hits: list[str] = []

    def primary_handler(path, method, body):
        primary_hits.append(f"{method} {path.split('?')[0]}")
        return 200, json.dumps({"time": 1700000000}), {}

    def aux_handler(path, method, body):
        aux_hits.append(f"{method} {path.split('?')[0]}")
        return 200, "{}", {}

    async with serve(primary_handler) as primary_url:
        async with serve(aux_handler) as aux_url:
            client = Client(base_url=primary_url, cassandra_base_url=aux_url)
            client.set_tokens("ACC", "REF")
            try:
                await client.activities_api.get_user_activities()
            finally:
                await client.close()

    assert "GET /api/v2/user_activities" in aux_hits, aux_hits
    assert "GET /api/v2/user_activities" not in primary_hits, primary_hits


@pytest.mark.asyncio
async def test_unrouted_request_stays_on_primary_host():
    primary_hits: list[str] = []
    aux_hits: list[str] = []

    def primary_handler(path, method, body):
        primary_hits.append(f"{method} {path.split('?')[0]}")
        return 200, json.dumps({"time": 1700000000}), {}

    def aux_handler(path, method, body):
        aux_hits.append(f"{method} {path.split('?')[0]}")
        return 200, "{}", {}

    async with serve(primary_handler) as primary_url:
        async with serve(aux_handler) as aux_url:
            client = Client(base_url=primary_url, cassandra_base_url=aux_url)
            client.set_tokens("ACC", "REF")
            try:
                await client.users_api.get_user_timestamp()
            finally:
                await client.close()

    assert any(h.endswith("/v2/users/timestamp") for h in primary_hits), primary_hits
    assert not any(h.endswith("/v2/users/timestamp") for h in aux_hits), aux_hits
