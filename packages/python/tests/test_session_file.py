# FileSessionStore round-trip / atomicity / concurrency parity tests.
# Mirrors the contract in PORTING.md §5 + §5.1 and tests/session_file.test.ts.

import asyncio
import json
import os
import stat
import sys

import pytest

from yaylib.session import NoSessionError, Session
from yaylib.session_file import FileSessionStore, new_session_store


def make_session(email: str, user_id: int) -> Session:
    return Session(
        email=email,
        user_id=user_id,
        access_token=f"access-{user_id}",
        refresh_token=f"refresh-{user_id}",
        device_uuid=f"00000000-0000-0000-0000-{user_id:012d}",
        updated_at="",
    )


async def test_roundtrip(tmp_path):
    path = str(tmp_path / "sessions.json")
    store = await new_session_store(path)
    await store.save(make_session("a@example.com", 1))
    loaded = await store.load("a@example.com")
    assert loaded.user_id == 1
    assert loaded.access_token == "access-1"
    assert loaded.device_uuid == "00000000-0000-0000-0000-000000000001"
    assert loaded.updated_at != ""  # auto-populated


async def test_load_missing_raises(tmp_path):
    store = await new_session_store(str(tmp_path / "sessions.json"))
    with pytest.raises(NoSessionError):
        await store.load("ghost@example.com")


async def test_delete_idempotent(tmp_path):
    store = await new_session_store(str(tmp_path / "sessions.json"))
    # No-op even when the file doesn't exist yet.
    await store.delete("ghost@example.com")
    await store.save(make_session("real@example.com", 7))
    await store.delete("real@example.com")
    with pytest.raises(NoSessionError):
        await store.load("real@example.com")
    # Second delete on the now-empty record is still a no-op.
    await store.delete("real@example.com")


async def test_multiple_sessions(tmp_path):
    path = str(tmp_path / "sessions.json")
    store = await new_session_store(path)
    await store.save(make_session("a@example.com", 1))
    await store.save(make_session("b@example.com", 2))
    await store.save(make_session("c@example.com", 3))
    assert (await store.load("a@example.com")).user_id == 1
    assert (await store.load("b@example.com")).user_id == 2
    assert (await store.load("c@example.com")).user_id == 3
    with open(path, encoding="utf-8") as fh:
        data = json.load(fh)
    assert len(data["sessions"]) == 3


@pytest.mark.skipif(sys.platform == "win32", reason="POSIX perms only")
async def test_file_permissions(tmp_path):
    path = str(tmp_path / "sessions.json")
    store = await new_session_store(path)
    await store.save(make_session("perm@example.com", 9))
    mode = stat.S_IMODE(os.stat(path).st_mode)
    assert mode == 0o600


async def test_concurrent_writes(tmp_path):
    path = str(tmp_path / "sessions.json")
    store = FileSessionStore(path)
    # 10 concurrent saves under distinct emails. The asyncio lock inside
    # FileSessionStore serializes the read-modify-write cycles, so all 10
    # records MUST be present in the final file.
    await asyncio.gather(
        *(store.save(make_session(f"u{i}@example.com", i + 1)) for i in range(10))
    )
    with open(path, encoding="utf-8") as fh:
        data = json.load(fh)
    assert len(data["sessions"]) == 10
