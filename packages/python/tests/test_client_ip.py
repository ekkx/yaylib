# Lazy X-Client-IP parity test (PORTING.md §12 + §15). Translated from
# the Go reference's TestTransport_LazyFetchClientIPPopulatesHeader: the
# first request goes out WITHOUT X-Client-IP (the fetch is async); the
# background GET /v2/users/timestamp caches ip_address; the second
# request carries it.

import asyncio

import pytest

from yaylib.client import Client

from ._server import serve

FETCHED_IP = "203.0.113.7"


async def test_lazy_fetch_client_ip_populates_header():
    block_ips = []  # X-Client-IP seen on each /v1/users/1/block hit

    def handler(path, method, body, headers):
        if path.endswith("/v2/users/timestamp"):
            return (
                200,
                '{"time":0,"ip_address":"%s"}' % FETCHED_IP,
                {},
            )
        if path == "/v1/users/1/block":
            block_ips.append(headers.get("X-Client-IP", ""))
            return 200, "{}", {}
        return 404, "", {}

    async with serve(handler) as base_url:
        client = Client(base_url=base_url)
        try:
            # First request kicks off the lazy IP fetch in the
            # background; the request itself goes out without the header.
            await client.users_api.block_user(id=1)

            deadline = asyncio.get_event_loop().time() + 2.0
            while asyncio.get_event_loop().time() < deadline:
                if client.client_ip == FETCHED_IP:
                    break
                await asyncio.sleep(0.02)
            assert client.client_ip == FETCHED_IP

            await client.users_api.block_user(id=1)

            assert block_ips[0] == ""  # async fetch — first goes bare
            assert block_ips[1] == FETCHED_IP
        finally:
            await client.close()


async def test_close_cancels_pending_background_work():
    # PORTING.md §3: close() cancels background work owned by the Client.
    # Simulate a still-running lazy fetch with a long-lived task and
    # assert close() cancels it (and returns promptly, not in ~100s).
    client = Client(base_url="http://127.0.0.1:1")

    async def _never():
        await asyncio.sleep(100)

    task = asyncio.ensure_future(_never())
    client._bg_tasks.add(task)
    task.add_done_callback(client._bg_tasks.discard)

    loop = asyncio.get_event_loop()
    start = loop.time()
    await client.close()
    elapsed = loop.time() - start

    assert task.cancelled()
    assert client._bg_tasks == set()
    assert elapsed < 5  # did not wait the 100s sleep out
