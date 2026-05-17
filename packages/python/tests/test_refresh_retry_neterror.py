# PORTING.md §6.1 case 3 / §15: refresh succeeds but the post-refresh
# retry hits a connection-level failure. The caller MUST see the network
# error, NOT a stale 401 (restoring the 401 would mislead them into
# re-auth when the real problem is the network). Tokens stay rotated.
#
# Translated from the Go reference's
# TestTransport_401RefreshSucceedsButRetryNetworkErrorSurfacesError.

import json

import pytest

from yaylib.client import Client
from yaylib.errors import APIError
from yaylib.retry import RetryPolicy

from ._raw_server import serve_raw


# PORTING:S9
async def test_refresh_ok_retry_neterror_surfaces_error_not_stale_401():
    data_attempts = 0

    def decide(path, method, body):
        if "/api/v1/oauth/token" in path:
            return (
                "respond",
                200,
                json.dumps(
                    {"access_token": "NEW_ACC", "refresh_token": "NEW_REF"}
                ),
            )
        nonlocal data_attempts
        data_attempts += 1
        if data_attempts == 1:
            return ("respond", 401, '{"error_code":-1}')
        # Post-refresh retry: drop the connection.
        return ("drop",)

    async with serve_raw(decide) as base_url:
        # Zero-value retry policy (mirrors Go's RetryPolicy{}) so the
        # 401-refresh replay path is isolated from transient retry.
        client = Client(
            base_url=base_url,
            retry_policy=RetryPolicy(
                max_attempts=0, base_delay=0.0, max_delay=0.0
            ),
        )
        client.set_tokens("STALE", "REF")
        try:
            with pytest.raises(APIError) as ei:
                await client.users_api.get_user_timestamp()
            # The surfaced error is the network failure, NOT a 401.
            assert ei.value.status != 401
            # Tokens were still rotated by the successful refresh.
            assert client.tokens.access == "NEW_ACC"
            assert client.tokens.refresh == "NEW_REF"
        finally:
            await client.close()
