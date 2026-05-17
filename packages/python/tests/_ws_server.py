"""In-process HTTP+WebSocket server for event-stream parity tests.

Mirrors the Go reference's fakeServer / wsSession: a JSON
``/v1/users/ws_token`` handler plus a WebSocket upgrade on every other
path. ``on_connect`` (an async callable) drives behaviour per accepted
socket. Connections are force-closed on teardown so a lingering socket
can't hang the test process.
"""

from __future__ import annotations

import asyncio
import contextlib
import json
from typing import Awaitable, Callable

from aiohttp import WSMsgType, web


class WSSession:
    def __init__(self, ws: web.WebSocketResponse) -> None:
        self.ws = ws
        self.received: "asyncio.Queue[dict]" = asyncio.Queue()
        self._closed = asyncio.Event()

    async def _read_loop(self) -> None:
        try:
            async for msg in self.ws:
                if msg.type == WSMsgType.TEXT:
                    with contextlib.suppress(ValueError, TypeError):
                        self.received.put_nowait(json.loads(msg.data))
                elif msg.type in (WSMsgType.CLOSE, WSMsgType.CLOSING,
                                  WSMsgType.CLOSED, WSMsgType.ERROR):
                    break
        finally:
            self._closed.set()

    async def send_welcome(self) -> None:
        await self.ws.send_str('{"type":"welcome","sid":"test-sid"}')

    async def confirm(self, identifier: str) -> None:
        await self.ws.send_str(
            json.dumps(
                {"identifier": identifier, "type": "confirm_subscription"}
            )
        )

    async def reject(self, identifier: str) -> None:
        await self.ws.send_str(
            json.dumps(
                {"identifier": identifier, "type": "reject_subscription"}
            )
        )

    async def push_event(self, identifier: str, event: str, data) -> None:
        inner = json.dumps({"event": event, "data": data})
        await self.ws.send_str(
            json.dumps({"identifier": identifier, "message": inner})
        )

    async def close(self) -> None:
        with contextlib.suppress(Exception):
            await self.ws.close()
        self._closed.set()

    async def idle(self) -> None:
        """Block until the socket closes (the test-side analog of the Go
        reference's ``<-s.ctx.Done()``)."""
        await self._closed.wait()


OnConnect = Callable[[WSSession], Awaitable[None]]


@contextlib.asynccontextmanager
async def serve_ws(on_connect: OnConnect):
    open_sessions: "set[WSSession]" = set()

    async def ws_token(_req: web.Request) -> web.Response:
        return web.json_response({"token": "test-token"})

    async def timestamp(_req: web.Request) -> web.Response:
        # Benign response for the lazy X-Client-IP fetch so it doesn't
        # add error noise during stream tests.
        return web.json_response({"time": 0, "ip_address": "127.0.0.1"})

    async def ws_handler(request: web.Request) -> web.WebSocketResponse:
        ws = web.WebSocketResponse()
        await ws.prepare(request)
        sess = WSSession(ws)
        open_sessions.add(sess)
        reader = asyncio.ensure_future(sess._read_loop())
        try:
            await on_connect(sess)
        finally:
            reader.cancel()
            with contextlib.suppress(asyncio.CancelledError, Exception):
                await reader
            with contextlib.suppress(Exception):
                await ws.close()
            open_sessions.discard(sess)
        return ws

    app = web.Application()
    app.router.add_get("/v1/users/ws_token", ws_token)
    app.router.add_get("/v2/users/timestamp", timestamp)
    app.router.add_route("GET", "/{tail:.*}", ws_handler)

    runner = web.AppRunner(app)
    await runner.setup()
    site = web.TCPSite(runner, "127.0.0.1", 0)
    await site.start()
    port = site._server.sockets[0].getsockname()[1]
    try:
        yield f"http://127.0.0.1:{port}"
    finally:
        for s in list(open_sessions):
            await s.close()
        await runner.cleanup()
