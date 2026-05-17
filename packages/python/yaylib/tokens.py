# PORTING.md §4: Tokens is the immutable snapshot of OAuth credentials.
# Mutate by going through Client.set_tokens, login_with_email,
# load_session, or the internal 401 refresh path — never by mutating a
# returned snapshot.

from __future__ import annotations

from dataclasses import dataclass


@dataclass(frozen=True)
class Tokens:
    access: str = ""
    refresh: str = ""


def empty_tokens() -> Tokens:
    return Tokens(access="", refresh="")
