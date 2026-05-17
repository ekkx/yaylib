# Required-header injection parity (PORTING.md §12 + §15). Translated
# from the Go reference's TestTransport_InjectsRequiredHeaders and
# TestTransport_OAuthEndpointUsesBasicAuth — the scenarios, not the Go
# net/http mechanics.

import base64
import json

from yaylib._config import DEFAULT_API_KEY
from yaylib.client import Client

from ._server import serve

_REQUIRED = [
    "User-Agent",
    "X-App-Version",
    "X-Device-Info",
    "X-Device-UUID",
    "X-Connection-Type",
    "X-Connection-Speed",
    "Accept-Language",
    "X-Timestamp",
]


async def test_injects_required_headers_with_bearer():
    captured = {}

    def handler(path, method, body, headers):
        captured.update(headers)
        return 200, json.dumps({"time": 0, "ip_address": "1.2.3.4"}), {}

    async with serve(handler) as base_url:
        client = Client(base_url=base_url)
        client.set_tokens("ACC", "REF")
        try:
            # get_user_timestamp is the /v2/users/timestamp path, so it
            # carries the full header set without spawning a lazy-IP
            # companion request — exactly one request to inspect.
            await client.users_api.get_user_timestamp()
        finally:
            await client.close()

    for h in _REQUIRED:
        assert captured.get(h), f"missing required header {h}"
    # Non-OAuth endpoint → Bearer.
    assert captured.get("Authorization") == "Bearer ACC"
    # PORTING.md §12 wire-significant literals.
    assert captured.get("X-Connection-Speed") == "0 kbps"


async def test_oauth_endpoint_uses_basic_auth():
    # The /api/v1/oauth/token call is reachable through the 401
    # auto-refresh path; capture the Authorization scheme the server
    # actually sees there.
    oauth_auth = {}

    def handler(path, method, body, headers):
        if path == "/api/v1/oauth/token":
            oauth_auth["value"] = headers.get("Authorization", "")
            return (
                200,
                json.dumps(
                    {"access_token": "new-a", "refresh_token": "new-r"}
                ),
                {},
            )
        if path.endswith("/v2/users/timestamp"):
            # 401 once → triggers refresh → retry succeeds.
            if "first" not in oauth_auth:
                oauth_auth["first"] = True
                return 401, '{"error_code":-1}', {}
            return 200, json.dumps({"time": 0, "ip_address": "1.2.3.4"}), {}
        return 404, "", {}

    async with serve(handler) as base_url:
        client = Client(base_url=base_url)
        client.set_tokens("stale", "stale-r")
        try:
            await client.users_api.get_user_timestamp()
        finally:
            await client.close()

    got = oauth_auth.get("value", "")
    assert got.startswith("Basic ")
    expected = base64.b64encode(DEFAULT_API_KEY.encode("utf-8")).decode("ascii")
    assert got == f"Basic {expected}"
