# Error-layer parity (PORTING.md §9 + §15). Translated from the Go
# reference's auth_test.go: 2FA-required surfaces as the typed code, and
# a session-store load failure returns the error WITHOUT any HTTP
# (rate-limit protection). Plus the §9 robustness invariants for
# code_of / error_response_of / raw-body retention.

import json

import pytest

from yaylib.client import Client
from yaylib.errors import APIError, code_of, error_response_of
from yaylib.session import MemorySessionStore, NoSessionError

from ._server import serve

# ErrCodeRequired2FA from the Yay! error catalog (error_codes.go).
ERR_CODE_REQUIRED_2FA = -31


async def test_2fa_required_surfaces_typed_code():
    def handler(path, method, body):
        if path.endswith("/login_with_email"):
            return (
                401,
                json.dumps(
                    {
                        "error_code": ERR_CODE_REQUIRED_2FA,
                        "message": "2FA required",
                    }
                ),
                {},
            )
        return 404, "", {}

    async with serve(handler) as base_url:
        client = Client(base_url=base_url)
        try:
            with pytest.raises(APIError) as ei:
                await (
                    client.login_with_email()
                    .email("u@example.com")
                    .password("pw")
                    .execute()
                )
            err = ei.value
            assert code_of(err) == ERR_CODE_REQUIRED_2FA
            parsed = error_response_of(err)
            assert parsed is not None
            assert parsed.message == "2FA required"
            # Tokens must not be set on a 2FA-required response.
            assert client.tokens.access == ""
        finally:
            await client.close()


class _LoadErrorStore(MemorySessionStore):
    async def load(self, email):
        raise RuntimeError("corrupt JSON")


async def test_login_store_load_error_returns_without_http():
    hits = 0

    def handler(path, method, body):
        nonlocal hits
        if path.endswith("/login_with_email"):
            hits += 1
        return 500, "{}", {}

    async with serve(handler) as base_url:
        client = Client(
            base_url=base_url, session_store=_LoadErrorStore()
        )
        try:
            with pytest.raises(RuntimeError, match="corrupt JSON"):
                await (
                    client.login_with_email()
                    .email("foo@example.com")
                    .password("pw")
                    .execute()
                )
            # A non-NoSessionError load failure must NOT fall through to
            # an HTTP login (rate-limit protection, PORTING.md §15).
            assert hits == 0
        finally:
            await client.close()


def test_code_of_and_error_response_of_on_non_api_error():
    # Never raise; return the zero-ish code / None for unrecognized
    # errors (PORTING.md §9).
    assert code_of(ValueError("nope")) == 0
    assert code_of(None) == 0
    assert error_response_of(ValueError("nope")) is None
    assert error_response_of(None) is None


def test_api_error_retains_raw_body():
    # An APIError stays usable for raw-body access; it is not collapsed
    # into ErrorResponse only (PORTING.md §9).
    err = APIError(status=400, reason="Bad Request")
    err.body = json.dumps(
        {"error_code": -1, "message": "bad", "ban_until": 123}
    )
    assert code_of(err) == -1
    parsed = error_response_of(err)
    assert parsed.message == "bad"
    assert parsed.ban_until == 123
    # Raw body still readable directly.
    assert "error_code" in err.body
