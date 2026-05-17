# PORTING.md §10 / §10.1: the multiplexed real-time event channel — one
# WebSocket against wss://cable.yay.space, N subscriptions on top, Rails
# ActionCable-style JSON framing. Reference: Go event_stream.go +
# event_stream_test.go. The wire constants (dial URL shape, channel
# identifier strings incl. the significant space after ",", frame
# inventory, event→DTO map) are pattern-matched server-side and MUST be
# byte-identical across ports.

from __future__ import annotations

import asyncio
import contextlib
import json
import random
from dataclasses import dataclass, field
from typing import Any, Callable, Dict, List, Optional

from aiohttp import WSMsgType

DEFAULT_EVENT_STREAM_URL = "wss://cable.yay.space"

_DEFAULT_SUBSCRIBE_TIMEOUT = 10.0
_DEFAULT_EVENT_BUFFER = 64

# How long a freshly-dialed connection must stay alive before the run
# loop treats it as healthy and resets the failure budget. Module-level
# so tests can shrink it; production never rewrites it (mirrors the Go
# reference's reconnectStableThreshold var).
reconnect_stable_threshold = 30.0


# --------------------------------------------------------------------------
# Events (events.go parity)
# --------------------------------------------------------------------------


class Event:
    """Base marker for a server-pushed event."""


@dataclass
class NewMessageEvent(Event):
    # The full chat-message DTO. Kept as the raw dict so a wire field the
    # generated RealmMessage model doesn't know (e.g. the WS-only
    # video_thumbnail_big_url, §10.1) stays observable.
    raw: Dict[str, Any]


@dataclass
class VideoProcessedEvent(Event):
    id: int = 0
    video_processed: bool = False
    video_url: str = ""
    video_thumbnail_url: str = ""
    video_thumbnail_big_url: str = ""


@dataclass
class ChatDeletedEvent(Event):
    room_id: int = 0


@dataclass
class TotalChatRequestEvent(Event):
    total_count: int = 0


@dataclass
class UnsubscribedEvent(Event):
    user_ids: List[int] = field(default_factory=list)


@dataclass
class GroupUpdatedEvent(Event):
    # Wire event name is `new_post` but the payload is a group-id update
    # (PORTING.md §10).
    group_id: int = 0


@dataclass
class CallFinishedEvent(Event):
    id: int = 0


@dataclass
class RawEvent(Event):
    """An event whose discriminator the SDK does not recognize. Lets
    callers observe new server signals before the SDK is updated.
    """

    name: str
    data: Dict[str, Any]


def decode_event(name: str, data: Dict[str, Any]) -> Optional[Event]:
    if not isinstance(data, dict):
        return None
    if name == "new_message":
        return NewMessageEvent(raw=data)
    if name == "video_processed":
        return VideoProcessedEvent(
            id=int(data.get("id") or 0),
            video_processed=bool(data.get("video_processed") or False),
            video_url=data.get("video_url") or "",
            video_thumbnail_url=data.get("video_thumbnail_url") or "",
            video_thumbnail_big_url=data.get("video_thumbnail_big_url") or "",
        )
    if name == "chat_deleted":
        return ChatDeletedEvent(room_id=int(data.get("room_id") or 0))
    if name == "total_chat_request":
        return TotalChatRequestEvent(
            total_count=int(data.get("total_count") or 0)
        )
    if name == "unsubscribed":
        return UnsubscribedEvent(user_ids=list(data.get("user_ids") or []))
    if name == "new_post":
        return GroupUpdatedEvent(group_id=int(data.get("group_id") or 0))
    if name == "conference_call_finished":
        return CallFinishedEvent(id=int(data.get("id") or 0))
    return RawEvent(name=name, data=data)


# Concrete event type -> its wire `event` name, for routing decoded
# events to the right decorator handler bucket.
_TYPE_TO_WIRE = {
    NewMessageEvent: "new_message",
    VideoProcessedEvent: "video_processed",
    ChatDeletedEvent: "chat_deleted",
    TotalChatRequestEvent: "total_chat_request",
    UnsubscribedEvent: "unsubscribed",
    GroupUpdatedEvent: "new_post",
    CallFinishedEvent: "conference_call_finished",
}


# --------------------------------------------------------------------------
# Channels (PORTING.md §10.1 — identifier strings are wire-significant:
# the space after "," matters and the IDs are JSON STRINGS not numbers)
# --------------------------------------------------------------------------


class Channel:
    def identifier(self) -> str:  # pragma: no cover - interface
        raise NotImplementedError


class _ChatRoomChannel(Channel):
    def identifier(self) -> str:
        return '{"channel":"ChatRoomChannel"}'


class _MessagesChannel(Channel):
    def __init__(self, room_id: int) -> None:
        self._room_id = room_id

    def identifier(self) -> str:
        return (
            '{"channel":"MessagesChannel", "chat_room_id":"%d"}'
            % self._room_id
        )


class _GroupUpdatesChannel(Channel):
    def identifier(self) -> str:
        return '{"channel":"GroupUpdatesChannel"}'


class _GroupPostsChannel(Channel):
    def __init__(self, group_id: int) -> None:
        self._group_id = group_id

    def identifier(self) -> str:
        return (
            '{"channel":"GroupPostsChannel", "group_id":"%d"}'
            % self._group_id
        )


def chat_room_channel() -> Channel:
    """Global chat-room signals (new requests, deletions, forced unsub)."""
    return _ChatRoomChannel()


def messages_channel(room_id: int) -> Channel:
    """Messages + video_processed for a single chat room."""
    return _MessagesChannel(room_id)


def group_updates_channel() -> Channel:
    """Group-membership / group-state updates across all groups."""
    return _GroupUpdatesChannel()


def group_posts_channel(group_id: int) -> Channel:
    """New posts within a single group's timeline."""
    return _GroupPostsChannel(group_id)


# --------------------------------------------------------------------------
# Policy / options
# --------------------------------------------------------------------------


@dataclass
class ReconnectPolicy:
    # Stop after the first disconnect; subscriptions terminate with the
    # connection error.
    disabled: bool = False
    # Caps total reconnect attempts (excluding the initial dial). 0 means
    # UNLIMITED — the OPPOSITE of RetryPolicy.max_attempts, because an
    # unattended event stream wants to keep recovering. Documented loudly
    # (PORTING.md §13).
    max_attempts: int = 0
    # First backoff sleep (seconds); each attempt doubles up to
    # max_delay, with full jitter.
    initial_delay: float = 0.5
    # Ceiling (seconds) on any single sleep between attempts.
    max_delay: float = 30.0


def default_reconnect_policy() -> ReconnectPolicy:
    return ReconnectPolicy(initial_delay=0.5, max_delay=30.0)


@dataclass
class EventStreamOptions:
    reconnect: ReconnectPolicy = field(default_factory=default_reconnect_policy)
    event_buffer: int = _DEFAULT_EVENT_BUFFER
    subscribe_timeout: float = _DEFAULT_SUBSCRIBE_TIMEOUT
    # Invoked when a subscription's buffer is full and an event must be
    # dropped — keeps backpressure observable (PORTING.md §10). Keep it
    # fast; it runs on the read pump.
    on_drop: Optional[Callable[[Event], None]] = None


# --------------------------------------------------------------------------
# Subscription
# --------------------------------------------------------------------------


_CLOSED = object()  # queue sentinel signalling terminal


class Subscription:
    def __init__(self, conn: "EventStream", ident: str, buffer: int) -> None:
        self._conn = conn
        self.ident = ident
        self._queue: "asyncio.Queue[Any]" = asyncio.Queue(maxsize=buffer)
        self._confirm: "asyncio.Future[None]" = (
            asyncio.get_event_loop().create_future()
        )
        self._done = asyncio.Event()
        self._closed = False
        self._handlers: Dict[str, List[Callable[[Event], Any]]] = {}
        self._handler_tasks: "set[asyncio.Task]" = set()
        self._pump_task: Optional[asyncio.Task] = None

    # ---- consumption: async-iterator path ----

    def __aiter__(self):
        return self._aiter()

    async def _aiter(self):
        while True:
            item = await self._queue.get()
            if item is _CLOSED:
                return
            yield item

    async def next_event(self, timeout: Optional[float] = None) -> Event:
        """Await the next event. Raises ConnectionError once the
        subscription has terminated and the buffer is drained.
        """
        if timeout is None:
            item = await self._queue.get()
        else:
            item = await asyncio.wait_for(self._queue.get(), timeout)
        if item is _CLOSED:
            raise ConnectionError("yaylib: subscription closed")
        return item

    async def done(self) -> None:
        await self._done.wait()

    @property
    def closed(self) -> bool:
        return self._closed

    # ---- consumption: decorator path (PORTING.md §10 idiom) ----

    def _register(self, wire: str, fn: Callable[[Event], Any]):
        self._handlers.setdefault(wire, []).append(fn)
        # Registering a handler switches the subscription to push mode:
        # a pump drains the queue into handlers. Starting it here (vs.
        # at delivery) means events that arrived between subscribe() and
        # this registration are still dispatched — no race for callers
        # that wire handlers up right after subscribe.
        if self._pump_task is None:
            self._pump_task = asyncio.ensure_future(self._pump())
        return fn

    async def _pump(self) -> None:
        try:
            while True:
                item = await self._queue.get()
                if item is _CLOSED:
                    return
                wire = _TYPE_TO_WIRE.get(type(item))
                for cb in self._handlers.get(wire or "", []):
                    self._spawn_handler(cb, item)
        except asyncio.CancelledError:
            raise

    def on_new_message(self, fn):
        return self._register("new_message", fn)

    def on_video_processed(self, fn):
        return self._register("video_processed", fn)

    def on_chat_deleted(self, fn):
        return self._register("chat_deleted", fn)

    def on_total_chat_request(self, fn):
        return self._register("total_chat_request", fn)

    def on_unsubscribed(self, fn):
        return self._register("unsubscribed", fn)

    def on_group_updated(self, fn):
        return self._register("new_post", fn)

    def on_call_finished(self, fn):
        return self._register("conference_call_finished", fn)

    # ---- delivery (read-pump side) ----

    def _deliver(self, ev: Event) -> None:
        if self._closed:
            return
        try:
            self._queue.put_nowait(ev)
        except asyncio.QueueFull:
            # Buffer full — drop. Surface via on_drop so backpressure is
            # observable (PORTING.md §10).
            cb = self._conn._opts.on_drop
            if cb is not None:
                cb(ev)

    def _spawn_handler(self, cb: Callable[[Event], Any], ev: Event) -> None:
        res = cb(ev)
        if asyncio.iscoroutine(res):
            task = asyncio.ensure_future(res)
            self._handler_tasks.add(task)
            task.add_done_callback(self._handler_tasks.discard)

    def _terminate(self, err: Optional[BaseException]) -> None:
        if self._closed:
            return
        self._closed = True
        if not self._confirm.done():
            if err is not None:
                self._confirm.set_exception(err)
            else:
                self._confirm.set_exception(
                    ConnectionError("yaylib: event stream closed")
                )
            self._confirm.add_done_callback(
                lambda f: not f.cancelled() and f.exception()
            )
        with contextlib.suppress(asyncio.QueueFull):
            self._queue.put_nowait(_CLOSED)
        if self._pump_task is not None:
            self._pump_task.cancel()
        self._done.set()

    async def unsubscribe(self) -> None:
        """Send an unsubscribe command and terminate. Idempotent."""
        if self._closed:
            return
        self._closed = True
        self._conn._remove_sub(self.ident)
        with contextlib.suppress(Exception):
            await self._conn._write_command("unsubscribe", self.ident)
        if not self._confirm.done():
            self._confirm.set_exception(
                ConnectionError("yaylib: unsubscribed")
            )
            # Guard against an "exception never retrieved" warning when
            # nothing is awaiting confirm (unsubscribe before confirm).
            self._confirm.add_done_callback(
                lambda f: not f.cancelled() and f.exception()
            )
        with contextlib.suppress(asyncio.QueueFull):
            self._queue.put_nowait(_CLOSED)
        if self._pump_task is not None:
            self._pump_task.cancel()
        self._done.set()


# --------------------------------------------------------------------------
# EventStream
# --------------------------------------------------------------------------


class EventStream:
    def __init__(
        self, client, base: str, opts: EventStreamOptions
    ) -> None:
        self._client = client
        self._base = base or DEFAULT_EVENT_STREAM_URL
        self._opts = opts
        self._ws = None
        self._subs: Dict[str, Subscription] = {}
        self._closed = False
        self._done = asyncio.Event()
        self._final_err: Optional[BaseException] = None
        self._run_task: Optional[asyncio.Task] = None

    # ---- connect / handshake ----

    async def _connect(self) -> None:
        token_resp = await self._client.users_api.get_web_socket_token()
        token = getattr(token_resp, "token", "") or ""

        sep = "&" if "?" in self._base else "?"
        url = (
            f"{self._base}{sep}token={token}"
            f"&app_version={self._client.api_version_name}"
        )

        session = self._client._transport._ensure_session()
        # No Authorization header on the dial — auth is exclusively the
        # `token` query parameter (PORTING.md §10.1 / §12).
        ws = await session.ws_connect(url, heartbeat=None)

        # Welcome handshake: the server sends {"type":"welcome"}
        # immediately; drain frames until we see it (preceding pings
        # ignored). A disconnect during welcome is an error.
        while True:
            msg = await ws.receive()
            if msg.type in (
                WSMsgType.TEXT,
                WSMsgType.BINARY,
            ):
                try:
                    f = json.loads(msg.data)
                except (ValueError, TypeError):
                    continue
                t = f.get("type")
                if t == "welcome":
                    break
                if t == "disconnect":
                    await ws.close()
                    raise ConnectionError(
                        "yaylib: server disconnect during welcome"
                    )
            else:
                await ws.close()
                raise ConnectionError("yaylib: ws closed during welcome")

        old = self._ws
        self._ws = ws
        if old is not None:
            with contextlib.suppress(Exception):
                await old.close()

    # ---- run loop / reconnect ----

    async def _run_loop(self) -> None:
        try:
            await self._read_until_disconnect()
            if self._opts.reconnect.disabled or self._closed:
                self._terminate_all(self._final_err)
                return

            attempt = 0
            while True:
                if self._closed:
                    return
                attempt += 1
                pol = self._opts.reconnect
                if pol.max_attempts > 0 and attempt > pol.max_attempts:
                    err = ConnectionError(
                        "yaylib: reconnect exhausted after "
                        f"{pol.max_attempts} attempts: {self._final_err}"
                    )
                    self._final_err = err
                    self._terminate_all(err)
                    return

                delay = self._backoff(attempt)
                try:
                    await asyncio.sleep(delay)
                except asyncio.CancelledError:
                    raise

                if self._closed:
                    return
                try:
                    await self._connect()
                except asyncio.CancelledError:
                    raise
                except Exception as exc:  # noqa: BLE001
                    self._final_err = exc
                    continue

                await self._resubscribe_all()
                started = asyncio.get_event_loop().time()
                await self._read_until_disconnect()
                if (
                    asyncio.get_event_loop().time() - started
                    >= reconnect_stable_threshold
                ):
                    # Survived long enough to be healthy — reset the
                    # budget so the next blip starts fresh. A faster drop
                    # counts toward max_attempts so flaps can't loop
                    # forever.
                    attempt = 0
                if self._closed:
                    return
        except asyncio.CancelledError:
            pass
        finally:
            self._done.set()

    def _backoff(self, attempt: int) -> float:
        pol = self._opts.reconnect
        exp = pol.initial_delay * (2 ** max(0, attempt - 1))
        if exp <= 0 or exp > pol.max_delay:
            exp = pol.max_delay
        if exp <= 0:
            return 0.0
        delay = random.random() * exp + pol.initial_delay
        return min(delay, pol.max_delay)

    async def _read_until_disconnect(self) -> None:
        ws = self._ws
        if ws is None:
            return
        try:
            async for msg in ws:
                if msg.type in (WSMsgType.TEXT, WSMsgType.BINARY):
                    self._dispatch(msg.data)
                else:
                    break
        except asyncio.CancelledError:
            raise
        except Exception as exc:  # noqa: BLE001
            self._final_err = exc
            return
        # Loop exit = socket closed/errored.
        if self._final_err is None and not self._closed:
            self._final_err = ConnectionError("yaylib: connection dropped")

    # ---- dispatch ----

    def _dispatch(self, data) -> None:
        try:
            f = json.loads(data)
        except (ValueError, TypeError):
            return
        if not isinstance(f, dict):
            return
        ftype = f.get("type")
        if ftype == "welcome":
            return
        if ftype == "ping":
            return
        if ftype == "disconnect":
            # Bubble up by closing the ws so the read loop exits and
            # reconnect kicks in.
            ws = self._ws
            if ws is not None:
                asyncio.ensure_future(self._safe_close_ws(ws))
            return
        if ftype == "confirm_subscription":
            sub = self._subs.get(f.get("identifier", ""))
            if sub is not None and not sub._confirm.done():
                sub._confirm.set_result(None)
            return
        if ftype == "reject_subscription":
            sub = self._subs.get(f.get("identifier", ""))
            if sub is not None:
                err = ConnectionError("yaylib: subscription rejected")
                if not sub._confirm.done():
                    sub._confirm.set_exception(err)
                sub._terminate(err)
            return

        message = f.get("message")
        if not message:
            return
        if isinstance(message, str):
            try:
                message = json.loads(message)
            except (ValueError, TypeError):
                return
        if not isinstance(message, dict):
            return
        name = message.get("event") or ""
        body = message.get("data")
        if body is None:
            body = message.get("message")
        if not name or body is None:
            return
        ev = decode_event(name, body)
        if ev is None:
            return
        sub = self._subs.get(f.get("identifier", ""))
        if sub is not None:
            sub._deliver(ev)

    async def _safe_close_ws(self, ws) -> None:
        with contextlib.suppress(Exception):
            await ws.close()

    # ---- subscription management ----

    def _remove_sub(self, ident: str) -> None:
        self._subs.pop(ident, None)

    async def _resubscribe_all(self) -> None:
        for ident in list(self._subs.keys()):
            with contextlib.suppress(Exception):
                await self._write_command("subscribe", ident)

    def _terminate_all(self, err: Optional[BaseException]) -> None:
        subs = list(self._subs.values())
        self._subs = {}
        for s in subs:
            s._terminate(err)

    async def _write_command(self, cmd: str, identifier: str) -> None:
        ws = self._ws
        if ws is None:
            raise ConnectionError("yaylib: ws not connected")
        await ws.send_str(
            json.dumps({"command": cmd, "identifier": identifier})
        )

    async def subscribe(self, channel: Channel) -> Subscription:
        if self._closed:
            raise ConnectionError("yaylib: event stream closed")
        ident = channel.identifier()
        existing = self._subs.get(ident)
        if existing is not None:
            return await self._await_confirm(existing, owned=False)

        sub = Subscription(self, ident, self._opts.event_buffer)
        self._subs[ident] = sub
        try:
            await self._write_command("subscribe", ident)
        except Exception as exc:  # noqa: BLE001
            self._remove_sub(ident)
            raise ConnectionError(
                f"yaylib: subscribe {ident}: {exc}"
            ) from exc
        return await self._await_confirm(sub, owned=True)

    async def _await_confirm(
        self, sub: Subscription, owned: bool
    ) -> Subscription:
        timeout = self._opts.subscribe_timeout

        async def _wait_done():
            await self._done.wait()

        done_wait = asyncio.ensure_future(_wait_done())
        try:
            await asyncio.wait(
                {sub._confirm, done_wait},
                timeout=timeout,
                return_when=asyncio.FIRST_COMPLETED,
            )
            if not sub._confirm.done() and not self._done.is_set():
                if owned:
                    self._remove_sub(sub.ident)
                raise ConnectionError(
                    f"yaylib: subscribe {sub.ident}: timed out after "
                    f"{timeout}s"
                )
            if sub._confirm.done():
                exc = sub._confirm.exception()
                if exc is not None:
                    if owned:
                        self._remove_sub(sub.ident)
                    raise exc
                return sub
            # stream done fired first
            if owned:
                self._remove_sub(sub.ident)
            if self._final_err is not None:
                raise self._final_err
            raise ConnectionError("yaylib: event stream closed")
        finally:
            done_wait.cancel()
            with contextlib.suppress(asyncio.CancelledError, Exception):
                await done_wait

    # ---- lifecycle ----

    async def _open(self) -> "EventStream":
        await self._connect()
        self._run_task = asyncio.ensure_future(self._run_loop())
        return self

    async def close(self) -> None:
        if self._closed:
            await self._done.wait()
            return
        self._closed = True
        ws = self._ws
        if ws is not None:
            with contextlib.suppress(Exception):
                await ws.close()
        if self._run_task is not None:
            self._run_task.cancel()
            with contextlib.suppress(asyncio.CancelledError, Exception):
                await self._run_task
        self._terminate_all(None)
        self._done.set()

    async def wait_done(self) -> None:
        await self._done.wait()

    def err(self) -> Optional[BaseException]:
        """Terminal error after the stream is done. None for a clean,
        caller-initiated close; non-nil if the stream died on its own.
        """
        if not self._done.is_set():
            return None
        if self._closed:
            return None
        return self._final_err

    async def __aenter__(self) -> "EventStream":
        return self

    async def __aexit__(self, exc_type, exc, tb) -> None:
        await self.close()


class _OpenEventStreamCM:
    """Returned by ``client.open_event_stream(...)``. Works both as
    ``async with client.open_event_stream(opts) as stream`` and
    ``stream = await client.open_event_stream(opts)``.
    """

    def __init__(self, client, opts: Optional[EventStreamOptions]) -> None:
        self._client = client
        self._opts = opts or EventStreamOptions()
        self._stream: Optional[EventStream] = None

    async def _start(self) -> EventStream:
        opts = self._opts
        if opts.reconnect.initial_delay <= 0:
            opts.reconnect.initial_delay = 0.5
        if opts.reconnect.max_delay <= 0:
            opts.reconnect.max_delay = 30.0
        if opts.event_buffer <= 0:
            opts.event_buffer = _DEFAULT_EVENT_BUFFER
        if opts.subscribe_timeout <= 0:
            opts.subscribe_timeout = _DEFAULT_SUBSCRIBE_TIMEOUT
        stream = EventStream(
            self._client, self._client.event_stream_url, opts
        )
        await stream._open()
        self._client._event_streams.add(stream)
        self._stream = stream
        return stream

    def __await__(self):
        return self._start().__await__()

    async def __aenter__(self) -> EventStream:
        return await self._start()

    async def __aexit__(self, exc_type, exc, tb) -> None:
        if self._stream is not None:
            await self._stream.close()
            self._client._event_streams.discard(self._stream)
