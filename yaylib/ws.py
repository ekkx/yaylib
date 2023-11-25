"""
MIT License

Copyright (c) 2023 qvco

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
"""

from __future__ import annotations

import websocket
from typing import Optional, Any

from . import client
from .config import Configs


__all__ = (
    "Intents",
    "WebSocketInteractor",
)


class Intents(object):
    pass


class WebSocketInteractor(object):
    def __init__(
        self,
        intents: Intents,
        base: client.BaseClient,
    ) -> None:
        self.__intents: Intents = intents
        self.__base: client.BaseClient = base
        self.__ws_token: Optional[str] = None
        self.__ws: Optional[websocket.WebSocketApp] = None

    def __on_open(self, ws):
        pass

    def __on_message(self, ws, message):
        pass

    def __on_error(self, ws, error):
        self.__base.logger.error(error)

    def __on_close(self, ws, close_status_code, close_msg):
        pass

    def __on_connect(self, sid: str):
        pass

    def __connect(self):
        # connect to channels based on intents
        pass

    def run(self, email: str, password: str) -> None:
        self.__base._prepare(email, password)
        self.__ws_token = self.__base.MiscAPI.get_web_socket_token().token
        self.__ws = websocket.WebSocketApp(
            f"wss://{Configs.CABLE_HOST}/?token={self.__ws_token}&app_version={Configs.VERSION_NAME}",
            on_open=self.__on_open,
            on_message=self.__on_message,
            on_error=self.__on_error,
            on_close=self.__on_close,
        )
        self.__ws.run_forever()

    def stop(self) -> None:
        self.__ws.keep_running = False
