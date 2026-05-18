# Event-stream parity (PORTING.md §10 / §10.1 / §15). Scenarios
# translated from the Go reference's event_stream_test.go — the
# behaviour, not the Go channel mechanics.

import asyncio

import pytest
from aiohttp import web

from yaylib import event_stream as es
from yaylib.client import Client
from yaylib.event_stream import (
    ChatDeletedEvent,
    EventStreamOptions,
    GroupUpdatedEvent,
    ReconnectPolicy,
    TotalChatRequestEvent,
    chat_room_channel,
    group_updates_channel,
)

from ._ws_server import serve_ws


def _client(http_url: str) -> Client:
    ws_url = http_url.replace("http://", "ws://", 1)
    c = Client(base_url=http_url, event_stream_url=ws_url)
    c.set_tokens("stub", "")  # skip auth-refresh path
    return c


_DISABLED = EventStreamOptions(reconnect=ReconnectPolicy(disabled=True))


# PORTING:S18,S25
async def test_subscribe_and_receive_event():
    async def on_connect(s):
        await s.send_welcome()
        msg = await s.received.get()
        assert msg["command"] == "subscribe"
        await s.confirm(msg["identifier"])
        await s.push_event(msg["identifier"], "chat_deleted", {"room_id": 123})
        await s.idle()

    async with serve_ws(on_connect) as base_url:
        client = _client(base_url)
        try:
            async with client.open_event_stream(_DISABLED) as stream:
                sub = await stream.subscribe(chat_room_channel())
                ev = await sub.next_event(timeout=2)
                assert isinstance(ev, ChatDeletedEvent)
                assert ev.room_id == 123
        finally:
            await client.close()


# PORTING:S34
async def test_unknown_event_observable_via_on_raw_and_on_event():
    # An unmodelled wire event must still surface through the primary
    # subscription API as a RawEvent carrying its raw name and data.
    async def on_connect(s):
        await s.send_welcome()
        msg = await s.received.get()
        await s.confirm(msg["identifier"])
        await s.push_event(
            msg["identifier"],
            "__unknown_evt__",
            {"foo": "bar", "n": 7},
        )
        await s.idle()

    async with serve_ws(on_connect) as base_url:
        client = _client(base_url)
        try:
            async with client.open_event_stream(_DISABLED) as stream:
                sub = await stream.subscribe(chat_room_channel())

                raw_seen = asyncio.Queue()
                any_seen = asyncio.Queue()

                @sub.on_raw
                def _r(ev):
                    raw_seen.put_nowait(ev)

                @sub.on_event
                def _a(ev):
                    any_seen.put_nowait(ev)

                raw_ev = await asyncio.wait_for(raw_seen.get(), timeout=2)
                any_ev = await asyncio.wait_for(any_seen.get(), timeout=2)

                for ev in (raw_ev, any_ev):
                    assert isinstance(ev, es.RawEvent)
                    assert ev.name == "__unknown_evt__"
                    assert ev.data == {"foo": "bar", "n": 7}
        finally:
            await client.close()


# PORTING:S21
async def test_rejected_subscription():
    async def on_connect(s):
        await s.send_welcome()
        msg = await s.received.get()
        await s.reject(msg["identifier"])
        await s.idle()

    async with serve_ws(on_connect) as base_url:
        client = _client(base_url)
        try:
            async with client.open_event_stream(_DISABLED) as stream:
                with pytest.raises(Exception):
                    await stream.subscribe(chat_room_channel())
        finally:
            await client.close()


# PORTING:S18,S25
async def test_multiple_channels():
    async def on_connect(s):
        await s.send_welcome()
        for _ in range(2):
            msg = await s.received.get()
            await s.confirm(msg["identifier"])
        await s.push_event(
            '{"channel":"ChatRoomChannel"}',
            "total_chat_request",
            {"total_count": 7},
        )
        await s.push_event(
            '{"channel":"GroupUpdatesChannel"}',
            "new_post",
            {"group_id": 42},
        )
        await s.idle()

    async with serve_ws(on_connect) as base_url:
        client = _client(base_url)
        try:
            async with client.open_event_stream(_DISABLED) as stream:
                sub1 = await stream.subscribe(chat_room_channel())
                sub2 = await stream.subscribe(group_updates_channel())
                e1 = await sub1.next_event(timeout=3)
                e2 = await sub2.next_event(timeout=3)
                assert isinstance(e1, TotalChatRequestEvent)
                assert e1.total_count == 7
                assert isinstance(e2, GroupUpdatedEvent)
                assert e2.group_id == 42
        finally:
            await client.close()


# PORTING:S19
async def test_reconnect_after_server_close_resubscribes():
    attempts = 0

    async def on_connect(s):
        nonlocal attempts
        attempts += 1
        n = attempts
        await s.send_welcome()
        msg = await s.received.get()
        assert msg["command"] == "subscribe"
        await s.confirm(msg["identifier"])
        if n == 1:
            await s.close()  # drop right after confirm
            return
        await s.push_event(msg["identifier"], "chat_deleted", {"room_id": 999})
        await s.idle()

    opts = EventStreamOptions(
        reconnect=ReconnectPolicy(initial_delay=0.05, max_delay=0.2)
    )
    async with serve_ws(on_connect) as base_url:
        client = _client(base_url)
        try:
            async with client.open_event_stream(opts) as stream:
                sub = await stream.subscribe(chat_room_channel())
                ev = await sub.next_event(timeout=3)
                assert isinstance(ev, ChatDeletedEvent)
                assert ev.room_id == 999
                assert attempts >= 2
        finally:
            await client.close()


async def test_unsubscribe_closes_events():
    got_unsub = asyncio.Event()

    async def on_connect(s):
        await s.send_welcome()
        msg = await s.received.get()
        await s.confirm(msg["identifier"])
        while True:
            nxt = await s.received.get()
            if nxt["command"] == "unsubscribe":
                got_unsub.set()
                return

    async with serve_ws(on_connect) as base_url:
        client = _client(base_url)
        try:
            async with client.open_event_stream(_DISABLED) as stream:
                sub = await stream.subscribe(chat_room_channel())
                await sub.unsubscribe()
                await asyncio.wait_for(got_unsub.wait(), timeout=2)
                with pytest.raises(ConnectionError):
                    await sub.next_event(timeout=1)
        finally:
            await client.close()


# PORTING:S21
async def test_subscribe_timeout():
    async def on_connect(s):
        await s.send_welcome()
        await s.received.get()  # never confirm
        await s.idle()

    opts = EventStreamOptions(
        reconnect=ReconnectPolicy(disabled=True), subscribe_timeout=0.2
    )
    async with serve_ws(on_connect) as base_url:
        client = _client(base_url)
        try:
            async with client.open_event_stream(opts) as stream:
                with pytest.raises(ConnectionError, match="timed out"):
                    await stream.subscribe(chat_room_channel())
        finally:
            await client.close()


async def test_done_and_err_on_clean_close():
    async def on_connect(s):
        await s.send_welcome()
        await s.idle()

    async with serve_ws(on_connect) as base_url:
        client = _client(base_url)
        try:
            stream = await client.open_event_stream(_DISABLED)
            assert stream.err() is None
            await stream.close()
            await asyncio.wait_for(stream.wait_done(), timeout=2)
            assert stream.err() is None  # clean, caller-initiated
        finally:
            await client.close()


# PORTING:S35
async def test_err_after_reconnect_exhausted():
    async def on_connect(s):
        await s.send_welcome()
        await s.close()  # drop immediately

    opts = EventStreamOptions(
        reconnect=ReconnectPolicy(
            max_attempts=1, initial_delay=0.01, max_delay=0.02
        )
    )
    async with serve_ws(on_connect) as base_url:
        client = _client(base_url)
        try:
            stream = await client.open_event_stream(opts)
            await asyncio.wait_for(stream.wait_done(), timeout=3)
            assert stream.err() is not None
        finally:
            await client.close()


# PORTING:S23
async def test_on_drop_fires_when_buffer_full():
    release = asyncio.Event()

    async def on_connect(s):
        await s.send_welcome()
        msg = await s.received.get()
        await s.confirm(msg["identifier"])
        await release.wait()
        for i in range(5):
            await s.push_event(
                msg["identifier"], "chat_deleted", {"room_id": i}
            )
        await s.idle()

    dropped = 0

    def on_drop(_ev):
        nonlocal dropped
        dropped += 1

    opts = EventStreamOptions(
        reconnect=ReconnectPolicy(disabled=True),
        event_buffer=1,
        on_drop=on_drop,
    )
    async with serve_ws(on_connect) as base_url:
        client = _client(base_url)
        try:
            async with client.open_event_stream(opts) as stream:
                await stream.subscribe(chat_room_channel())
                release.set()  # burst — we never consume the buffer
                for _ in range(200):
                    if dropped > 0:
                        break
                    await asyncio.sleep(0.01)
                assert dropped >= 1
        finally:
            await client.close()


async def test_decorator_handler_idiom():
    async def on_connect(s):
        await s.send_welcome()
        msg = await s.received.get()
        await s.confirm(msg["identifier"])
        await s.push_event(msg["identifier"], "chat_deleted", {"room_id": 55})
        await s.idle()

    seen = asyncio.Queue()

    async with serve_ws(on_connect) as base_url:
        client = _client(base_url)
        try:
            async with client.open_event_stream(_DISABLED) as stream:
                sub = await stream.subscribe(chat_room_channel())

                @sub.on_chat_deleted
                def _h(ev):
                    seen.put_nowait(ev)

                ev = await asyncio.wait_for(seen.get(), timeout=2)
                assert isinstance(ev, ChatDeletedEvent)
                assert ev.room_id == 55
        finally:
            await client.close()


# PORTING:S20
async def test_multiple_subs_resubscribe_after_reconnect():
    attempts = 0
    ident_chat = '{"channel":"ChatRoomChannel"}'
    ident_group = '{"channel":"GroupUpdatesChannel"}'

    async def on_connect(s):
        nonlocal attempts
        attempts += 1
        n = attempts
        await s.send_welcome()
        seen = set()
        while len(seen) < 2:
            msg = await s.received.get()
            if msg.get("command") != "subscribe":
                continue
            seen.add(msg["identifier"])
            await s.confirm(msg["identifier"])
        assert ident_chat in seen and ident_group in seen
        if n == 1:
            await s.close()
            return
        await s.push_event(ident_chat, "chat_deleted", {"room_id": 11})
        await s.push_event(ident_group, "new_post", {"group_id": 22})
        await s.idle()

    opts = EventStreamOptions(
        reconnect=ReconnectPolicy(initial_delay=0.02, max_delay=0.05)
    )
    async with serve_ws(on_connect) as base_url:
        client = _client(base_url)
        try:
            async with client.open_event_stream(opts) as stream:
                sub_chat = await stream.subscribe(chat_room_channel())
                sub_group = await stream.subscribe(group_updates_channel())
                e1 = await sub_chat.next_event(timeout=4)
                e2 = await sub_group.next_event(timeout=4)
                assert isinstance(e1, ChatDeletedEvent) and e1.room_id == 11
                assert isinstance(e2, GroupUpdatedEvent) and e2.group_id == 22
                assert attempts >= 2
        finally:
            await client.close()


# PORTING:S24
async def test_stable_connection_resets_attempt_budget(monkeypatch):
    # Shrink the stability threshold so the test can cross it quickly.
    monkeypatch.setattr(es, "reconnect_stable_threshold", 0.05)
    attempts = 0

    async def on_connect(s):
        nonlocal attempts
        attempts += 1
        n = attempts
        await s.send_welcome()
        msg = await s.received.get()
        await s.confirm(msg["identifier"])
        if n <= 2:
            # Stay alive past the (shrunk) threshold, then drop. Without
            # the time-based reset, MaxAttempts=1 would exhaust after the
            # 2nd connection — only the reset lets a 3rd dial happen.
            await asyncio.sleep(0.08)
            await s.close()
            return
        await s.idle()

    opts = EventStreamOptions(
        reconnect=ReconnectPolicy(
            max_attempts=1, initial_delay=0.01, max_delay=0.02
        )
    )
    async with serve_ws(on_connect) as base_url:
        client = _client(base_url)
        try:
            async with client.open_event_stream(opts) as stream:
                await stream.subscribe(chat_room_channel())
                for _ in range(200):
                    if attempts >= 3:
                        break
                    await asyncio.sleep(0.01)
                assert attempts >= 3
        finally:
            await client.close()


# PORTING:S22
async def test_ws_dial_does_not_leak_bearer():
    captured = {}

    async def ws_token(_req):
        return web.json_response({"token": "stub"})

    async def upgrade(request):
        captured["auth"] = request.headers.get("Authorization", "")
        return web.Response(status=400, text="no upgrade in this test")

    app = web.Application()
    app.router.add_get("/v1/users/ws_token", ws_token)
    app.router.add_get("/v2/users/timestamp", lambda r: web.json_response(
        {"time": 0, "ip_address": "127.0.0.1"}
    ))
    app.router.add_route("GET", "/{tail:.*}", upgrade)
    runner = web.AppRunner(app)
    await runner.setup()
    site = web.TCPSite(runner, "127.0.0.1", 0)
    await site.start()
    port = site._server.sockets[0].getsockname()[1]
    try:
        client = _client(f"http://127.0.0.1:{port}")
        client.set_tokens("SECRET-ACCESS-TOKEN", "REF")
        try:
            with pytest.raises(Exception):
                await client.open_event_stream(_DISABLED)
            assert captured.get("auth", "") == ""
        finally:
            await client.close()
    finally:
        await runner.cleanup()
