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
        if path == "/v3/users/login_with_email":
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
        if path == "/v3/users/login_with_email":
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
        if path == "/v3/users/login_with_email":
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
        if path == "/v3/users/login_with_email":
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
