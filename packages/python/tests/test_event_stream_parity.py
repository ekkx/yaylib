# Event-stream behavior parity, driven against the shared server's
# WebSocket contract (PORTING.md §10 / §15). Translated from the Go
# reference's event_stream_parity_test.go. The per-dial mode is selected
# by mock_stream_client (see _parity.py). The shared server pushes one
# representative event per channel: ChatRoomChannel ->
# chat_deleted{room_id:123}, GroupUpdatesChannel -> new_post{group_id:1}.
#
# The reconnect-budget mechanics (exhaustion, the 30s stable-connection
# reset, the per-sub on_drop overflow) and the unsubscribe lifecycle
# stay as in-process fixtures in test_event_stream.py: the shared WS
# contract pushes a single event per subscribe and exposes no
# socket-fault knob, so those are not expressible here (§15 covers
# reconnect/reject/timeout, which are). These skip when YAYLIB_MOCK_WS_URL
# is unset.

import asyncio

import pytest

from yaylib.event_stream import (
    ChatDeletedEvent,
    EventStreamOptions,
    GroupUpdatedEvent,
    ReconnectPolicy,
    chat_room_channel,
    group_updates_channel,
)

from ._parity import mock_stream_client, ws_close_all

_DISABLED = EventStreamOptions(reconnect=ReconnectPolicy(disabled=True))


async def test_subscribe_and_receive_event():
    c = mock_stream_client("")
    c.set_tokens("stub", "")
    try:
        async with c.open_event_stream(_DISABLED) as stream:
            sub = await stream.subscribe(chat_room_channel())
            ev = await sub.next_event(timeout=2.0)
            assert isinstance(ev, ChatDeletedEvent)
            assert ev.room_id == 123, "server's representative payload"
    finally:
        await c.close()


async def test_rejected_subscription():
    c = mock_stream_client("reject")
    c.set_tokens("stub", "")
    try:
        async with c.open_event_stream(_DISABLED) as stream:
            with pytest.raises(Exception):
                await stream.subscribe(chat_room_channel())
    finally:
        await c.close()


async def test_multiple_channels():
    c = mock_stream_client("")
    c.set_tokens("stub", "")
    try:
        async with c.open_event_stream(_DISABLED) as stream:
            sub1 = await stream.subscribe(chat_room_channel())
            sub2 = await stream.subscribe(group_updates_channel())

            # Per-sub queues are independent, so a sequential read of
            # each channel's representative event is race-free (no shared
            # channel to drain, unlike the Go select).
            ev1 = await sub1.next_event(timeout=3.0)
            ev2 = await sub2.next_event(timeout=3.0)
            assert isinstance(ev1, ChatDeletedEvent) and ev1.room_id == 123
            assert isinstance(ev2, GroupUpdatedEvent) and ev2.group_id == 1
    finally:
        await c.close()


async def test_reconnect_after_server_close():
    # drop-after-confirm closes the socket right after pushing the event,
    # so the client must reconnect and re-subscribe to keep receiving.
    # Seeing the representative event more than once proves the
    # reconnect + re-subscribe cycle.
    c = mock_stream_client("drop-after-confirm")
    c.set_tokens("stub", "")
    opts = EventStreamOptions(
        reconnect=ReconnectPolicy(initial_delay=0.01, max_delay=0.03)
    )
    try:
        async with c.open_event_stream(opts) as stream:
            sub = await stream.subscribe(chat_room_channel())
            for i in range(2):
                ev = await sub.next_event(timeout=3.0)
                assert isinstance(ev, ChatDeletedEvent), (
                    f"cycle {i}: want ChatDeletedEvent"
                )
    finally:
        await c.close()


async def test_subscribe_timeout():
    c = mock_stream_client("no-confirm")
    c.set_tokens("stub", "")
    opts = EventStreamOptions(
        reconnect=ReconnectPolicy(disabled=True), subscribe_timeout=0.2
    )
    try:
        async with c.open_event_stream(opts) as stream:
            with pytest.raises(Exception):
                await stream.subscribe(chat_room_channel())
    finally:
        await c.close()


async def test_done_and_err_on_clean_close():
    c = mock_stream_client("")
    c.set_tokens("stub", "")
    try:
        stream = await c.open_event_stream(_DISABLED)
        assert stream.err() is None, "Err before Done should be nil"
        await stream.close()
        await asyncio.wait_for(stream.wait_done(), timeout=2.0)
        assert stream.err() is None, "Err after a clean Close should be nil"
    finally:
        await c.close()


async def test_multiple_subs_resubscribe_after_reconnect():
    # Mode "" keeps the socket open so BOTH subscribes confirm on one
    # stable connection (drop-after-confirm closes after the first
    # confirm, which the Go reference only survives via Go's random
    # map-iteration on resubscribe — not reproducible under Python's
    # insertion-ordered dict). The reconnect is driven deterministically
    # by the server's admin close route, which the mockserver exposes
    # exactly for multi-subscription reconnect parity.
    c = mock_stream_client("")
    c.set_tokens("stub", "")
    opts = EventStreamOptions(
        reconnect=ReconnectPolicy(initial_delay=0.01, max_delay=0.03)
    )
    try:
        async with c.open_event_stream(opts) as stream:
            sub1 = await stream.subscribe(chat_room_channel())
            sub2 = await stream.subscribe(group_updates_channel())

            # Initial representative event on each channel.
            assert isinstance(
                await sub1.next_event(timeout=3.0), ChatDeletedEvent
            )
            assert isinstance(
                await sub2.next_event(timeout=3.0), GroupUpdatedEvent
            )

            # Force a server-side close; the client must reconnect and
            # re-subscribe EVERY sub, after which both receive again.
            await ws_close_all()
            assert isinstance(
                await sub1.next_event(timeout=4.0), ChatDeletedEvent
            )
            assert isinstance(
                await sub2.next_event(timeout=4.0), GroupUpdatedEvent
            )
    finally:
        await c.close()


async def test_ws_dial_does_not_leak_bearer():
    c = mock_stream_client("")
    # Tokens are set, but the WS dial must authenticate via the query
    # only. The shared server refuses the upgrade (401) if the dial
    # carries Authorization, so a successful open + subscribe proves no
    # bearer rode the handshake.
    c.set_tokens("ACCESS_THAT_MUST_NOT_LEAK", "REF")
    try:
        async with c.open_event_stream(_DISABLED) as stream:
            sub = await stream.subscribe(chat_room_channel())
            ev = await sub.next_event(timeout=2.0)
            assert isinstance(ev, ChatDeletedEvent)
    finally:
        await c.close()
