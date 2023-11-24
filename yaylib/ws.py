from __future__ import annotations

from . import client

from typing import Optional, List, Any


class WebSocketInteractor(object):
    def __init__(self, base: client.BaseClient, **options: Any) -> None:
        self.__base: client.BaseClient = base
        self.__intents: Optional[List[str]] = options.pop("intents", None)
