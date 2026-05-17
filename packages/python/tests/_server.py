"""In-process HTTP server for parity tests.

Mirrors the Node ``createServer`` helper used by the TypeScript tests.
The server is force-shut on teardown (connections cancelled) so a
lingering keep-alive socket can't hang the test process.
"""

from __future__ import annotations

import contextlib
import inspect
from typing import Callable, Tuple

from aiohttp import web

# handler(path, method, body[, headers]) -> (status, body[, headers]).
# A 4-arg handler additionally receives the request headers (CIMultiDict).
Handler = Callable[..., Tuple]


@contextlib.asynccontextmanager
async def serve(handler: Handler):
    wants_headers = len(inspect.signature(handler).parameters) >= 4

    async def _dispatch(request: web.Request) -> web.Response:
        body = await request.text()
        if wants_headers:
            result = handler(
                request.path_qs, request.method, body, request.headers
            )
        else:
            result = handler(request.path_qs, request.method, body)
        if len(result) == 3:
            status, out_body, extra = result
        else:
            status, out_body = result
            extra = {}
        headers = {"Content-Type": "application/json"}
        headers.update(extra or {})
        return web.Response(status=status, text=out_body, headers=headers)

    app = web.Application()
    app.router.add_route("*", "/{tail:.*}", _dispatch)
    runner = web.AppRunner(app)
    await runner.setup()
    site = web.TCPSite(runner, "127.0.0.1", 0)
    await site.start()
    port = site._server.sockets[0].getsockname()[1]
    try:
        yield f"http://127.0.0.1:{port}"
    finally:
        # cleanup() closes the listening socket AND drops live
        # connections, so undici-style keep-alive lingering can't stall.
        await runner.cleanup()
