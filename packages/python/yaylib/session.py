# PORTING.md §5: SessionStore persists (email -> Session) pairs so tokens
# survive across process restarts. This module ships the interface and an
# in-memory implementation; the file-backed default lives in
# session_file.py and is loaded via new_session_store.

from __future__ import annotations

import abc
from dataclasses import dataclass, replace
from datetime import datetime, timezone


def _now_rfc3339() -> str:
    # RFC3339 / ISO8601 with a trailing 'Z', matching the Go reference's
    # time.RFC3339 output for UTC.
    return (
        datetime.now(timezone.utc)
        .replace(microsecond=0)
        .strftime("%Y-%m-%dT%H:%M:%SZ")
    )


@dataclass
class Session:
    """One persisted identity. Field names mirror the §5.1 JSON schema
    so the file representation is a direct dump of this dataclass.
    """

    email: str = ""
    user_id: int = 0
    access_token: str = ""
    refresh_token: str = ""
    device_uuid: str = ""
    updated_at: str = ""  # RFC3339


class NoSessionError(Exception):
    """Sentinel raised by ``SessionStore.load`` on a cache miss."""

    def __init__(self, email: str) -> None:
        super().__init__(f"yaylib: no session for {email}")
        self.email = email


class SessionSaveFailed(Exception):
    """Raised by login_with_email when persisting a fresh session fails.

    The login response is still returned and the tokens are active in
    memory (PORTING.md §5); only the persistence step failed.
    """

    def __init__(self, cause: BaseException) -> None:
        super().__init__(f"yaylib: failed to persist session: {cause}")
        self.cause = cause


class SessionStore(abc.ABC):
    @abc.abstractmethod
    async def load(self, email: str) -> Session:
        """Return the cached session or raise NoSessionError."""

    @abc.abstractmethod
    async def save(self, session: Session) -> None:
        """Upsert. Called after fresh login and after each refresh."""

    @abc.abstractmethod
    async def delete(self, email: str) -> None:
        """Idempotent; a missing key is not an error."""


class MemorySessionStore(SessionStore):
    def __init__(self) -> None:
        self._sessions: dict[str, Session] = {}

    async def load(self, email: str) -> Session:
        s = self._sessions.get(email)
        if s is None:
            raise NoSessionError(email)
        return replace(s)

    async def save(self, session: Session) -> None:
        if not session.email:
            raise ValueError("yaylib: SessionStore.save: email is required")
        stored = replace(
            session,
            updated_at=session.updated_at or _now_rfc3339(),
        )
        self._sessions[stored.email] = stored

    async def delete(self, email: str) -> None:
        self._sessions.pop(email, None)


def new_memory_store() -> SessionStore:
    return MemorySessionStore()
