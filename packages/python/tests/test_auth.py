# login_with_email wrapper parity tests (PORTING.md §6).
# Mirrors tests/auth.test.ts: cache hit / miss / no_cache / save failure /
# file-store round trip.

import json

import pytest

from yaylib.client import Client
from yaylib.session import (
    MemorySessionStore,
    NoSessionError,
    Session,
    SessionSaveFailed,
)
from yaylib.session_file import FileSessionStore, new_session_store

from ._server import serve


# PORTING:S1
async def test_fresh_login_persists(tmp_path):
    login_hits = 0

    def handler(path, method, body):
        nonlocal login_hits
        # Web-first: the primary browser endpoint (/v2) answers here; the
        # suffix match keeps this path-agnostic like the Go fake server.
        if path.endswith("/users/login_with_email"):
            login_hits += 1
            parsed = json.loads(body or "{}")
            if (
                parsed.get("email") != "alice@example.com"
                or parsed.get("password") != "p4ss"
            ):
                return 400, '{"error_code":-1}', {}
            return (
                200,
                json.dumps(
                    {
                        "access_token": "ACC",
                        "refresh_token": "REF",
                        "user_id": 42,
                    }
                ),
                {},
            )
        return 404, "", {}

    async with serve(handler) as base_url:
        store = await new_session_store(str(tmp_path / "sessions.json"))
        client = Client(base_url=base_url, session_store=store)
        try:
            res = await client.login_with_email(
                email="alice@example.com", password="p4ss"
            )
            assert login_hits == 1
            assert res.access_token == "ACC"
            assert client.tokens.access == "ACC"
            assert client.user_id == 42
            persisted = await store.load("alice@example.com")
            assert persisted.user_id == 42
        finally:
            await client.close()


# PORTING:S2
async def test_cache_hit_no_http():
    login_hits = 0

    def handler(path, method, body):
        nonlocal login_hits
        login_hits += 1
        return 500, "{}", {}

    async with serve(handler) as base_url:
        store = MemorySessionStore()
        await store.save(
            Session(
                email="bob@example.com",
                user_id=99,
                access_token="cached-acc",
                refresh_token="cached-ref",
                device_uuid="11111111-1111-1111-1111-111111111111",
                updated_at="2024-01-01T00:00:00Z",
            )
        )
        client = Client(base_url=base_url, session_store=store)
        try:
            res = await client.login_with_email(
                email="bob@example.com", password="ignored"
            )
            assert login_hits == 0
            assert client.tokens.access == "cached-acc"
            assert client.user_id == 99
            assert client.device_uuid == "11111111-1111-1111-1111-111111111111"
            assert res.user_id == 99
        finally:
            await client.close()


async def test_no_cache_bypass():
    login_hits = 0

    def handler(path, method, body):
        nonlocal login_hits
        if path.endswith("/users/login_with_email"):
            login_hits += 1
            return (
                200,
                json.dumps(
                    {
                        "access_token": "FRESH",
                        "refresh_token": "FRESH-R",
                        "user_id": 7,
                    }
                ),
                {},
            )
        return 404, "", {}

    async with serve(handler) as base_url:
        store = MemorySessionStore()
        await store.save(
            Session(
                email="carol@example.com",
                user_id=1,
                access_token="OLD",
                refresh_token="OLD-R",
                device_uuid="22222222-2222-2222-2222-222222222222",
                updated_at="2024-01-01T00:00:00Z",
            )
        )
        client = Client(base_url=base_url, session_store=store)
        try:
            res = await client.login_with_email(
                email="carol@example.com", password="p4ss", no_cache=True
            )
            assert login_hits == 1
            assert client.tokens.access == "FRESH"
            assert res.access_token == "FRESH"
            persisted = await store.load("carol@example.com")
            assert persisted.access_token == "FRESH"
        finally:
            await client.close()


class _FailingStore(MemorySessionStore):
    async def load(self, email):
        raise NoSessionError(email)

    async def save(self, session):
        raise RuntimeError("disk full")


# PORTING:S4
async def test_save_failure_wraps_but_keeps_tokens():
    def handler(path, method, body):
        if path.endswith("/users/login_with_email"):
            return (
                200,
                json.dumps(
                    {"access_token": "A", "refresh_token": "R", "user_id": 5}
                ),
                {},
            )
        return 404, "", {}

    async with serve(handler) as base_url:
        client = Client(base_url=base_url, session_store=_FailingStore())
        try:
            with pytest.raises(SessionSaveFailed):
                await client.login_with_email(
                    email="dave@example.com", password="p4ss"
                )
            assert client.tokens.access == "A"
            assert client.user_id == 5
        finally:
            await client.close()


async def test_file_store_round_trip(tmp_path):
    login_hits = 0

    def handler(path, method, body):
        nonlocal login_hits
        if path.endswith("/users/login_with_email"):
            login_hits += 1
            return (
                200,
                json.dumps(
                    {"access_token": "T", "refresh_token": "T-R", "user_id": 11}
                ),
                {},
            )
        return 404, "", {}

    async with serve(handler) as base_url:
        path = str(tmp_path / "sessions.json")
        client1 = Client(base_url=base_url, session_store=FileSessionStore(path))
        try:
            await client1.login_with_email(
                email="eve@example.com", password="p4ss"
            )
            assert login_hits == 1
        finally:
            await client1.close()

        client2 = Client(base_url=base_url, session_store=FileSessionStore(path))
        try:
            await client2.login_with_email(
                email="eve@example.com", password="p4ss"
            )
            assert login_hits == 1  # cache hit, no second HTTP login
            assert client2.tokens.access == "T"
            assert client2.user_id == 11
        finally:
            await client2.close()



# --- Web-first login (browser endpoint primary, app fallback) ----------------
# PORTING.md §6.2. Mirrors the Go auth_test.go web-first scenarios.

import base64 as _b64

from yaylib import (
    ErrCodeRequired2FA,
    ErrCodeUnauthorized,
    ErrCodeUserTemporaryBanned,
    code_of,
    error_response_of,
)
from yaylib.exceptions import ApiException

DEFAULT_WEB_KEY = "e9f1ae4c4470f29a92c0168dc42b13637cc332692103f23e626bc2b016f66603"
APP_KEY = "ccd59ee269c01511ba763467045c115779fcae3050238a252f1bd1a4b65cfec6"


def _reject_401(code=401):
    return 401, json.dumps({"result": "error", "message": "unauthorized",
                            "error_code": code}), {}


def _basic(key):
    return "Basic " + _b64.b64encode(key.encode("utf-8")).decode("ascii")


async def test_web_primary_succeeds():
    """Web (v2) is primary: it succeeds directly and the app (v3) endpoint
    is never touched; the session tracks the web key for refresh."""
    seen = {}
    app_hits = 0

    def handler(path, method, body):
        nonlocal app_hits
        if path == "/v2/users/login_with_email":
            seen["key"] = json.loads(body or "{}").get("api_key")
            return (200, json.dumps(
                {"access_token": "WEBACC", "refresh_token": "WEBREF",
                 "user_id": 7777}), {})
        if path == "/v3/users/login_with_email":
            app_hits += 1
            return (200, json.dumps(
                {"access_token": "APPACC", "refresh_token": "APPREF",
                 "user_id": 4242}), {})
        return 404, "", {}

    async with serve(handler) as base_url:
        client = Client(base_url=base_url)
        try:
            res = await client.login_with_email(
                email="u@example.com", password="pw")
            assert res.user_id == 7777
            assert client.tokens.access == "WEBACC"
            assert client.tokens.refresh == "WEBREF"
            assert seen["key"] == DEFAULT_WEB_KEY
            assert client._oauth_key_snapshot() == DEFAULT_WEB_KEY
            assert app_hits == 0
        finally:
            await client.close()


async def test_falls_back_to_app_on_401():
    """A plain 401 from the web login falls back to the app (v3) endpoint
    under the app key; both endpoints are hit and refresh tracks app key."""
    seen = {}
    web_hits = 0
    app_hits = 0

    def handler(path, method, body):
        nonlocal web_hits, app_hits
        if path == "/v2/users/login_with_email":
            web_hits += 1
            return _reject_401()
        if path == "/v3/users/login_with_email":
            app_hits += 1
            seen["key"] = json.loads(body or "{}").get("api_key")
            return (200, json.dumps(
                {"access_token": "APPACC", "refresh_token": "APPREF",
                 "user_id": 4242}), {})
        return 404, "", {}

    async with serve(handler) as base_url:
        client = Client(base_url=base_url)
        try:
            res = await client.login_with_email(
                email="u@example.com", password="pw")
            assert res.user_id == 4242
            assert client.tokens.access == "APPACC"
            assert seen["key"] == APP_KEY
            assert client._oauth_key_snapshot() == APP_KEY
            assert web_hits == 1 and app_hits == 1
        finally:
            await client.close()


async def test_web_disabled_uses_app_only():
    """With the web path off (web_api_key=""), login goes straight to the
    app endpoint; the v2 request is never attempted and an app 401
    surfaces unchanged."""
    web_hits = 0

    def handler(path, method, body):
        nonlocal web_hits
        if path == "/v2/users/login_with_email":
            web_hits += 1
            return (200, json.dumps(
                {"access_token": "WEBACC", "user_id": 1}), {})
        if path == "/v3/users/login_with_email":
            return _reject_401()
        return 404, "", {}

    async with serve(handler) as base_url:
        client = Client(base_url=base_url, web_api_key="")
        try:
            with pytest.raises(ApiException) as ei:
                await client.login_with_email(
                    email="u@example.com", password="pw")
            assert code_of(ei.value) == ErrCodeUnauthorized
            assert web_hits == 0
            assert client.tokens.access == ""
        finally:
            await client.close()


async def test_both_fail_surfaces_web_error():
    """When both the web login and the app fallback fail, the primary
    (web) 401 is surfaced."""
    def handler(path, method, body):
        if path == "/v2/users/login_with_email":
            return _reject_401()
        if path == "/v3/users/login_with_email":
            return _reject_401()
        return 404, "", {}

    async with serve(handler) as base_url:
        client = Client(base_url=base_url)
        try:
            with pytest.raises(ApiException) as ei:
                await client.login_with_email(
                    email="u@example.com", password="pw")
            assert code_of(ei.value) == ErrCodeUnauthorized
            assert client.tokens.access == ""
        finally:
            await client.close()


async def test_no_fallback_on_2fa():
    """A 2FA-required rejection is HTTP 401 but carries Required2FA; the
    app fallback must NOT fire, so the web 2FA error surfaces and the app
    (v3) endpoint is never called."""
    v3_hits = 0

    def handler(path, method, body):
        nonlocal v3_hits
        if path == "/v2/users/login_with_email":
            return _reject_401(code=ErrCodeRequired2FA)
        if path == "/v3/users/login_with_email":
            v3_hits += 1
            return (200, json.dumps(
                {"access_token": "X", "refresh_token": "Y", "user_id": 1}), {})
        return 404, "", {}

    async with serve(handler) as base_url:
        client = Client(base_url=base_url)
        try:
            with pytest.raises(ApiException) as ei:
                await client.login_with_email(
                    email="u@example.com", password="pw")
            assert code_of(ei.value) == ErrCodeRequired2FA
            assert v3_hits == 0
            assert client.tokens.access == ""
        finally:
            await client.close()


async def test_no_fallback_when_api_key_pinned():
    """A caller-pinned api_key skips the web-first path entirely: login
    goes straight to the app endpoint under the pinned key and the browser
    (v2) endpoint is never called."""
    seen = {}
    web_hits = 0

    def handler(path, method, body):
        nonlocal web_hits
        if path == "/v2/users/login_with_email":
            web_hits += 1
            return (200, "{}", {})
        if path == "/v3/users/login_with_email":
            seen["key"] = json.loads(body or "{}").get("api_key")
            return _reject_401()
        return 404, "", {}

    async with serve(handler) as base_url:
        client = Client(base_url=base_url)
        try:
            with pytest.raises(ApiException) as ei:
                await client.login_with_email(
                    email="u@example.com", password="pw",
                    api_key="caller-pinned-key")
            assert code_of(ei.value) == ErrCodeUnauthorized
            assert web_hits == 0
            assert seen["key"] == "caller-pinned-key"
        finally:
            await client.close()


async def test_web_decode_failure_surfaced():
    """A 2xx web response that fails to decode surfaces as a typed
    ApiException carrying the raw body; it is not Unauthorized so no app
    fallback fires."""
    app_hits = 0

    def handler(path, method, body):
        nonlocal app_hits
        if path == "/v2/users/login_with_email":
            return (200, "{ this is not valid json", {})
        if path == "/v3/users/login_with_email":
            app_hits += 1
            return (200, json.dumps(
                {"access_token": "APPACC", "refresh_token": "APPREF",
                 "user_id": 1}), {})
        return 404, "", {}

    async with serve(handler) as base_url:
        client = Client(base_url=base_url)
        try:
            with pytest.raises(ApiException) as ei:
                await client.login_with_email(
                    email="u@example.com", password="pw")
            # Not masked as an auth rejection, and no app fallback.
            assert code_of(ei.value) != ErrCodeUnauthorized
            assert app_hits == 0
            # The raw web body rides on the typed error.
            assert "not valid json" in (ei.value.body or "")
            assert client.tokens.access == ""
        finally:
            await client.close()


async def test_web_primary_typed_error_surfaced():
    """A NON-Unauthorized typed web rejection (temporary ban with
    ban_until) surfaces directly as a typed ApiException — no app fallback
    — and error_response_of reads ban_until with the raw body accessible."""
    app_hits = 0

    def handler(path, method, body):
        nonlocal app_hits
        if path == "/v2/users/login_with_email":
            return (403, json.dumps({
                "result": "error", "message": "temporarily banned",
                "error_code": ErrCodeUserTemporaryBanned,
                "ban_until": 1893456000}), {})
        if path == "/v3/users/login_with_email":
            app_hits += 1
            return (200, json.dumps(
                {"access_token": "APPACC", "user_id": 1}), {})
        return 404, "", {}

    async with serve(handler) as base_url:
        client = Client(base_url=base_url)
        try:
            with pytest.raises(ApiException) as ei:
                await client.login_with_email(
                    email="u@example.com", password="pw")
            assert code_of(ei.value) == ErrCodeUserTemporaryBanned
            assert app_hits == 0
            parsed = error_response_of(ei.value)
            assert parsed is not None and parsed.ban_until == 1893456000
            assert "temporarily banned" in (ei.value.body or "")
        finally:
            await client.close()


# --- Public load_session / save_session activation (PORTING.md §6.2) ----------


async def test_load_session_restores_web_oauth_key():
    """A web session loaded straight from disk via the public load_session
    must refresh with the web key — the regression the oauth_key restore
    fixes."""
    seen = {}

    def handler(path, method, body, headers):
        if "/api/v1/oauth/token" in path:
            seen["auth"] = headers.get("Authorization")
            return (200, json.dumps(
                {"access_token": "A2", "refresh_token": "R2", "id": 1}), {})
        return 404, "", {}

    async with serve(handler) as base_url:
        store = MemorySessionStore()
        await store.save(Session(
            email="u@example.com", user_id=1,
            access_token="A1", refresh_token="R1",
            oauth_key=DEFAULT_WEB_KEY))
        client = Client(base_url=base_url, session_store=store)
        try:
            sess = await client.load_session("u@example.com")
            assert sess.oauth_key == DEFAULT_WEB_KEY
            assert client._oauth_key_snapshot() == DEFAULT_WEB_KEY
            assert await client._do_refresh(client.tokens.access) is True
            assert seen["auth"] == _basic(DEFAULT_WEB_KEY)
            assert client.tokens.access == "A2"
        finally:
            await client.close()


async def test_load_session_app_session_keeps_app_key():
    """App sessions carry an empty oauth_key; load_session leaves refresh
    on the app key (the empty -> app-key default)."""
    store = MemorySessionStore()
    await store.save(Session(
        email="u@example.com", user_id=1,
        access_token="A1", refresh_token="R1"))  # no oauth_key
    client = Client(session_store=store)
    try:
        await client.load_session("u@example.com")
        assert client._oauth_key_snapshot() == APP_KEY
    finally:
        await client.close()


async def test_save_session_activates_web_oauth_key():
    """A web session injected directly via save_session also activates its
    establishing key, so its refresh uses the web key."""
    seen = {}

    def handler(path, method, body, headers):
        if "/api/v1/oauth/token" in path:
            seen["auth"] = headers.get("Authorization")
            return (200, json.dumps(
                {"access_token": "A2", "refresh_token": "R2", "id": 1}), {})
        return 404, "", {}

    async with serve(handler) as base_url:
        client = Client(base_url=base_url, session_store=MemorySessionStore())
        try:
            await client.save_session(Session(
                email="u@example.com", user_id=1,
                access_token="A1", refresh_token="R1",
                oauth_key=DEFAULT_WEB_KEY))
            assert client._oauth_key_snapshot() == DEFAULT_WEB_KEY
            assert await client._do_refresh(client.tokens.access) is True
            assert seen["auth"] == _basic(DEFAULT_WEB_KEY)
        finally:
            await client.close()


# --- OAuth refresh uses the session's establishing key (PORTING.md §6.2) ------


async def test_refresh_web_session_uses_web_key():
    seen = {}

    def handler(path, method, body, headers):
        if path == "/v3/users/login_with_email":
            return _reject_401()
        if path == "/v2/users/login_with_email":
            return (200, json.dumps(
                {"access_token": "A1", "refresh_token": "R1", "user_id": 1}), {})
        if "/api/v1/oauth/token" in path:
            seen["auth"] = headers.get("Authorization")
            return (200, json.dumps(
                {"access_token": "A2", "refresh_token": "R2", "id": 1}), {})
        return 404, "", {}

    async with serve(handler) as base_url:
        client = Client(base_url=base_url)
        try:
            await client.login_with_email(email="u@example.com", password="pw")
            assert client._oauth_key_snapshot() == DEFAULT_WEB_KEY
            assert await client._do_refresh(client.tokens.access) is True
            assert seen["auth"] == _basic(DEFAULT_WEB_KEY)
            assert client.tokens.access == "A2"
        finally:
            await client.close()


async def test_refresh_app_session_uses_app_key(tmp_path):
    # App login reached by disabling the web-first path with web_api_key="".
    seen = {}

    def handler(path, method, body, headers):
        if path.endswith("/login_with_email"):
            return (200, json.dumps(
                {"access_token": "A1", "refresh_token": "R1", "user_id": 1}), {})
        if "/api/v1/oauth/token" in path:
            seen["auth"] = headers.get("Authorization")
            return (200, json.dumps(
                {"access_token": "A2", "refresh_token": "R2", "id": 1}), {})
        return 404, "", {}

    async with serve(handler) as base_url:
        store = MemorySessionStore()
        client = Client(base_url=base_url, session_store=store, web_api_key="")
        try:
            await client.login_with_email(email="u@example.com", password="pw")
            assert client._oauth_key_snapshot() == APP_KEY
            sess = await store.load("u@example.com")
            assert sess.oauth_key == ""  # app sessions don't persist a key
            assert await client._do_refresh(client.tokens.access) is True
            assert seen["auth"] == _basic(APP_KEY)
        finally:
            await client.close()


async def test_refresh_restored_web_session_uses_web_key():
    seen = {}

    def handler(path, method, body, headers):
        if path == "/v3/users/login_with_email":
            return _reject_401()
        if path == "/v2/users/login_with_email":
            return (200, json.dumps(
                {"access_token": "A1", "refresh_token": "R1", "user_id": 1}), {})
        if "/api/v1/oauth/token" in path:
            seen["auth"] = headers.get("Authorization")
            return (200, json.dumps(
                {"access_token": "A2", "refresh_token": "R2", "id": 1}), {})
        return 404, "", {}

    async with serve(handler) as base_url:
        store = MemorySessionStore()
        c1 = Client(base_url=base_url, session_store=store)
        try:
            await c1.login_with_email(email="u@example.com", password="pw")
            assert (await store.load("u@example.com")).oauth_key == DEFAULT_WEB_KEY
        finally:
            await c1.close()

        c2 = Client(base_url=base_url, session_store=store)
        try:
            await c2.login_with_email(email="u@example.com", password="pw")
            assert c2._oauth_key_snapshot() == DEFAULT_WEB_KEY  # restored
            assert await c2._do_refresh(c2.tokens.access) is True
            assert seen["auth"] == _basic(DEFAULT_WEB_KEY)
        finally:
            await c2.close()


# --- signed_info stays on api_key, ignoring the OAuth key (PORTING.md §6.2) ---


def test_signed_info_ignores_oauth_key():
    # Same api_key + device_uuid; only the OAuth key differs. signed_info
    # is keyed on api_key, so setting the OAuth key to the web key must
    # not change the hash (the app-key/OAuth-key split).
    ref = Client(api_key="app-key", device_uuid="uuid-shared")
    web = Client(api_key="app-key", device_uuid="uuid-shared")
    web.set_oauth_key("web-key")  # establish a web session
    assert (
        web.generate_signed_info_at(1700000000).value
        == ref.generate_signed_info_at(1700000000).value
    )
