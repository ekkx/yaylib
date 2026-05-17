"""Minimal raw-socket HTTP/1.1 server for tests that need to drop a
connection mid-exchange (simulating a post-refresh network failure).

aiohttp's web server always completes a response, so the
connection-level failure scenario from the Go reference
(``TestTransport_401RefreshSucceedsButRetryNetworkErrorSurfacesError``,
which hijacks and closes the TCP socket) needs a hand-rolled server.
``Connection: close`` is sent on every response so each request is its
own connection — no pipelining to parse.
"""

from __future__ import annotations

import asyncio
import contextlib
from typing import Callable, Tuple

# decide(path, method, body) -> ("respond", status, body_str) | ("drop",)
Decider = Callable[[str, str, str], Tuple]


@contextlib.asynccontextmanager
async def serve_raw(decide: Decider):
    async def handle(reader: asyncio.StreamReader, writer: asyncio.StreamWriter):
        try:
            header_blob = await reader.readuntil(b"\r\n\r\n")
        except (asyncio.IncompleteReadError, asyncio.LimitOverrunError):
            writer.close()
            return
        lines = header_blob.decode("latin1").split("\r\n")
        method, path, _ = lines[0].split(" ", 2)
        content_length = 0
        for ln in lines[1:]:
            if ln.lower().startswith("content-length:"):
                content_length = int(ln.split(":", 1)[1].strip() or "0")
        body = b""
        if content_length:
            body = await reader.readexactly(content_length)

        action = decide(path, method, body.decode("utf-8", "replace"))
        if action[0] == "drop":
            # Abruptly close without a response → client sees a
            # connection-level failure (aiohttp ClientError).
            writer.transport.abort()
            return

        _, status, payload = action
        data = payload.encode("utf-8")
        head = (
            f"HTTP/1.1 {status} X\r\n"
            f"Content-Type: application/json\r\n"
            f"Content-Length: {len(data)}\r\n"
            f"Connection: close\r\n\r\n"
        ).encode("latin1")
        writer.write(head + data)
        with contextlib.suppress(Exception):
            await writer.drain()
        writer.close()
        with contextlib.suppress(Exception):
            await writer.wait_closed()

    server = await asyncio.start_server(handle, "127.0.0.1", 0)
    port = server.sockets[0].getsockname()[1]
    async with server:
        await server.start_serving()
        try:
            yield f"http://127.0.0.1:{port}"
        finally:
            server.close()
            with contextlib.suppress(Exception):
                await server.wait_closed()
