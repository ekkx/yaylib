"""
MIT License

Copyright (c) 2023 ekkx

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

import json
import websocket

from typing import Optional, Dict, Type, Any

from . import client
from .config import Configs
from .models import Message


__all__ = (
    "Intents",
    "WebSocketInteractor",
)


class Intents:
    def __init__(self):
        self.chat_message: bool = False
        self.group_update: bool = False

    @classmethod
    def all(cls: Type[Intents]) -> Intents:
        self = cls()
        for attr in vars(self):
            setattr(self, attr, True)
        return self

    @classmethod
    def none(cls: Type[Intents]) -> Intents:
        return cls()


class Identifier(object):
    __slots__ = "channel"

    def __init__(self, data) -> None:
        self.channel: Optional[str] = data.get("channel")


class Content(object):
    __slots__ = ("event", "message", "data")

    def __init__(self, data) -> None:
        self.event: Optional[str] = data.get("event")
        self.message: Optional[Dict[str, Any]] = data.get("message")
        self.data: Optional[Dict[str, Any]] = data.get("data")


class ChannelMessage(object):
    __slots__ = ("type", "message", "identifier", "sid")

    def __init__(self, data) -> None:
        self.type: Optional[str] = data.get("type")

        self.message: Optional[Content] = data.get("message")
        if self.message is not None and isinstance(self.message, dict):
            self.message = Content(self.message)

        self.identifier: Optional[Identifier] = data.get("identifier")
        if self.identifier is not None:
            self.identifier = Identifier(json.loads(self.identifier))

        self.sid: Optional[str] = data.get("sid")


class WebSocketInteractor(object):
    intent_map: Dict[str, str] = {
        "chat_message": "ChatRoomChannel",
        "group_update": "GroupUpdatesChannel",
    }

    def __init__(
        self,
        base: client.BaseClient,
        intents: Intents,
    ) -> None:
        self.__intents: Intents = intents
        self.__base: client.BaseClient = base
        self.__ws_token: Optional[str] = None
        self.__ws: Optional[websocket.WebSocketApp] = None

    def set_ws_token(self, ws_token: str) -> None:
        self.__ws_token = ws_token

    def __on_open(self, ws):
        self.__base.logger.debug("on_open")

    def __on_message(self, ws, message: str):
        self.__base.logger.debug(f"on_message: {message}")

        events = ChannelMessage(json.loads(message))

        if events.type == "ping":
            return  # ignore ping events

        if events.type == "welcome":
            self.__connect()
            self.on_ready()
            return

        if events.type == "confirm_subscription" and events.identifier:
            self.__base.logger.debug(
                f"Connected to Gateway -> {events.identifier.channel}"
            )
            return

        content = events.message

        if content and events.identifier and content.event:
            if events.identifier.channel == "ChatRoomChannel":
                if content.event == "new_message":
                    if content.message is not None:
                        self.on_message(Message(content.message))

                elif content.event == "chat_deleted":
                    self.on_chat_room_delete(content.data.get("room_id"))

                elif content.event == "total_chat_request":
                    self.on_chat_request(content.data.get("total_count"))

            elif events.identifier.channel == "GroupUpdatesChannel":
                if content.event == "new_post":
                    self.on_group_update(content.data.get("group_id"))

    def on_chat_message(self, message: Message):
        """チャットメッセージが届いた時に発火される関数"""
        pass

    def on_chat_room_delete(self, room_id: int | None):
        """チャットルームが削除された時に発火される関数"""
        pass

    def on_chat_request(self, total_count: int | None):
        """チャットリクエストが届いた時に発火される関数"""
        pass

    def on_group_update(self, group_id: int | None):
        """サークルに投稿された時に発火される関数"""
        pass

    def __on_error(self, ws, error):
        self.__base.logger.error(error)

    def __on_close(self, ws, close_status_code, close_msg):
        pass

    def __send_channel_command(self, command: str, channel: str) -> None:
        if self.__ws is not None:
            self.__ws.send(
                json.dumps(
                    {
                        "command": command,
                        "identifier": f'{{"channel":"{channel}"}}',
                    }
                )
            )

    def __subscribe(self, channel: str) -> None:
        self.__send_channel_command("subscribe", channel)

    def __unsubscribe(self, channel: str) -> None:
        self.__send_channel_command("unsubscribe", channel)

    def __connect(self) -> None:
        # Connect to channels based on intents
        intents_vars = vars(self.__intents)
        channels: list[str] = [
            channel for channel, value in intents_vars.items() if value
        ]
        for channel in channels:
            channel = self.intent_map.get(channel)
            if channel is not None:
                self.__subscribe(channel)

    def on_ready(self) -> None:
        """クライアントの準備が完了すると呼び出されます"""
        self.__base.logger.debug("on_ready")

    def run(self, email: str = None, password: str = None) -> None:
        """クライアントを起動します"""
        if email is not None and password is not None:
            self.__base._prepare(email, password)

        if self.__ws_token is None:
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
        """クライアントを停止します"""
        self.__ws.keep_running = False
