from __future__ import annotations

from . import client

from typing import Optional, Any


__all__ = (
    "Intents",
    "WebSocketInteractor",
)


class Intents(object):
    pass


class WebSocketInteractor(object):
    def __init__(self, intents: Intents, **options: Any) -> None:
        self.__intents: Intents = intents
