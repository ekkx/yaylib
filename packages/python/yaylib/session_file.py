# PORTING.md §5.1: file-backed SessionStore. Persists sessions as a JSON
# document keyed by email so one file can hold multiple identities.
# Writes are atomic within a single process (temp-file + rename); the
# file is created with mode 0o600 because it carries access tokens.
# Multi-process writers are out of scope (PORTING.md §5).

from __future__ import annotations

import asyncio
import dataclasses
import json
import os
from typing import Optional

from yaylib.session import (
    NoSessionError,
    Session,
    SessionStore,
    _now_rfc3339,
)

_FIELDS = (
    "email",
    "user_id",
    "access_token",
    "refresh_token",
    "device_uuid",
    "updated_at",
)


def _session_from_dict(email: str, raw: dict) -> Session:
    return Session(
        email=raw.get("email") or email,
        user_id=int(raw.get("user_id") or 0),
        access_token=raw.get("access_token") or "",
        refresh_token=raw.get("refresh_token") or "",
        device_uuid=raw.get("device_uuid") or "",
        updated_at=raw.get("updated_at") or "",
    )


class FileSessionStore(SessionStore):
    """One-process-safe file store. The asyncio lock serializes the
    read-modify-write cycle so concurrent save / delete calls don't lose
    updates. Cross-process safety is not provided (PORTING.md §5).
    """

    def __init__(self, path: str) -> None:
        self._path = path
        # Created lazily so construction works outside a running loop
        # (Client is built synchronously).
        self._lock: Optional[asyncio.Lock] = None

    def _mu(self) -> asyncio.Lock:
        if self._lock is None:
            self._lock = asyncio.Lock()
        return self._lock

    async def load(self, email: str) -> Session:
        async with self._mu():
            data = self._read_all()
            raw = data["sessions"].get(email)
            if raw is None:
                raise NoSessionError(email)
            return _session_from_dict(email, raw)

    async def save(self, session: Session) -> None:
        if not session.email:
            raise ValueError("yaylib: SessionStore.save: email is required")
        async with self._mu():
            data = self._read_all()
            record = dataclasses.asdict(session)
            record["updated_at"] = session.updated_at or _now_rfc3339()
            data["sessions"][session.email] = {
                k: record[k] for k in _FIELDS
            }
            self._write_all(data)

    async def delete(self, email: str) -> None:
        async with self._mu():
            data = self._read_all()
            if email not in data["sessions"]:
                return
            del data["sessions"][email]
            self._write_all(data)

    # ---- blocking helpers, only ever called inside the lock ----

    def _read_all(self) -> dict:
        try:
            with open(self._path, "r", encoding="utf-8") as fh:
                text = fh.read()
        except FileNotFoundError:
            return {"sessions": {}}
        try:
            parsed = json.loads(text) if text else {}
        except ValueError as err:
            raise ValueError(
                f"yaylib: failed to parse session file {self._path}: {err}"
            ) from err
        sessions = parsed.get("sessions") if isinstance(parsed, dict) else None
        return {"sessions": sessions if isinstance(sessions, dict) else {}}

    def _write_all(self, data: dict) -> None:
        directory = os.path.dirname(self._path)
        if directory and directory != ".":
            os.makedirs(directory, mode=0o700, exist_ok=True)
        tmp = f"{self._path}.tmp"
        # Open with explicit mode so the token-bearing file is 0o600 even
        # if it didn't exist yet.
        fd = os.open(tmp, os.O_WRONLY | os.O_CREAT | os.O_TRUNC, 0o600)
        try:
            with os.fdopen(fd, "w", encoding="utf-8") as fh:
                json.dump(data, fh, indent=2)
        except BaseException:
            try:
                os.unlink(tmp)
            except OSError:
                pass
            raise
        try:
            os.replace(tmp, self._path)
        except BaseException:
            try:
                os.unlink(tmp)
            except OSError:
                pass
            raise
        # os.replace preserves the destination's prior mode when it
        # already existed; re-assert 0o600 unconditionally.
        try:
            os.chmod(self._path, 0o600)
        except OSError:
            pass


async def new_session_store(path: str) -> SessionStore:
    """Default factory. Creates the parent directory eagerly so
    subsequent saves never fail on a missing path.
    """
    directory = os.path.dirname(path)
    if directory and directory != ".":
        os.makedirs(directory, mode=0o700, exist_ok=True)
    return FileSessionStore(path)
