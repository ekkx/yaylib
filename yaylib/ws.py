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

import asyncio
from typing import Optional

import aiohttp

from . import config
from .models import Message, WSChannelMessage, WSMessage


class Intents:
    def __init__(self):
        self.chat_message: bool = False
        self.group_update: bool = False

    @classmethod
    def all(cls):
        self = cls()
        for attr in vars(self):
            setattr(self, attr, True)
        return self

    @classmethod
    def none(cls):
        return cls()


class WSEventListener:
    async def on_ready(self):
        pass

    # ---------- ChatRoomChannel ----------

    async def on_message(self, message: Message):
        pass

    async def on_chat_delete(self, room_id: int):
        pass

    async def on_chat_request(self, total_requests: int):
        pass

    # ---------- GroupUpdatesChannel ----------

    async def on_group_update(self, group_id: int):
        pass


class WebSocketInteractor(WSEventListener):
    INTENT_MAP = {
        "chat_message": "ChatRoomChannel",
        "group_update": "GroupUpdatesChannel",
    }

    def __init__(self, client, intents: Intents):
        # pylint: disable=import-outside-toplevel
        from .client import Client

        self.__client: Client = client
        self.__intents = intents
        self.__session: Optional[aiohttp.ClientSession] = None
        self.__ws: Optional[aiohttp.ClientWebSocketResponse] = None
        self.__ws_token: Optional[str] = None

    def set_ws_token(self, token: str):
        self.__ws_token = token

    async def __send_channel_command(self, command: str, channel: str) -> None:
        if self.__ws is not None:
            await self.__ws.send_json(
                {
                    "command": command,
                    "identifier": f'{{"channel":"{channel}"}}',
                }
            )

    # ---------- event handlers ----------

    async def __on_ping_event(self, channel_msg: WSChannelMessage):
        pass  # ignore ping messages

    # pylint: disable=unused-argument
    async def __on_welcome_event(self, channel_msg: WSChannelMessage):
        intents_vars = vars(self.__intents)
        channels: list[str] = [
            channel for channel, value in intents_vars.items() if value
        ]
        for channel in channels:
            channel = self.INTENT_MAP.get(channel)
            if channel is not None:
                await self.__send_channel_command("subscribe", channel)

        await self.on_ready()

    async def __on_confirm_subscription_event(self, channel_msg: WSChannelMessage):
        if channel_msg.identifier:
            self.__client.logger.info(
                f"Connected to Gateway -> {channel_msg.identifier.channel}"
            )

    async def __on_disconnect_event(self, channel_msg: WSChannelMessage):
        self.__client.logger.error(
            f"WebSocket disconnected! reason: {channel_msg.reason}"
        )
        await self.stop()

    # ---------- channel handlers ----------

    async def __on_chat_room_channel_event(self, ws_msg: WSMessage):
        match ws_msg.event:
            case "new_message":
                if ws_msg.message is None:
                    return
                await self.on_message(Message(ws_msg.message))
            case "chat_deleted":
                await self.on_chat_delete(ws_msg.data.get("room_id"))
            case "total_chat_request":
                await self.on_chat_request(ws_msg.data.get("total_count"))
            case _:
                self.__client.logger.error(f"Unknown event: {ws_msg.event}")
                return

    async def __on_group_updates_channel_event(self, ws_msg: WSMessage):
        match ws_msg.event:
            case "new_post":
                await self.on_group_update(ws_msg.data.get("group_id"))
            case _:
                self.__client.logger.error(f"Unknown event: {ws_msg.event}")
                return

    # ---------- low level events ----------

    async def __on_open(self):
        self.__client.logger.debug("ws: __on_open()")

    async def __on_message(self, data: dict):
        channel_msg = WSChannelMessage(data)
        self.__client.logger.debug(f"ws: __on_message({channel_msg})")

        event_handlers = {
            "ping": self.__on_ping_event,
            "welcome": self.__on_welcome_event,
            "confirm_subscription": self.__on_confirm_subscription_event,
            "disconnect": self.__on_disconnect_event,
        }
        event_handler = event_handlers.get(channel_msg.type)
        if event_handler:
            await event_handler(channel_msg)
            return

        content = channel_msg.message
        if not content or not content.event or not channel_msg.identifier:
            return

        channel_handlers = {
            "ChatRoomChannel": self.__on_chat_room_channel_event,
            "GroupUpdatesChannel": self.__on_group_updates_channel_event,
        }
        channel_handler = channel_handlers.get(channel_msg.identifier.channel)
        if channel_handler:
            await channel_handler(content)
        else:
            self.__client.logger.debug(
                f"Unknown channel: {channel_msg.identifier.channel}"
            )

    async def __on_error(self, data: dict):
        self.__client.logger.error(f"ws: __on_error(), data={data}")

    async def __on_close(self, data: dict):
        self.__client.logger.debug("ws: __on_close()")
        print(data)

    async def __start_ws(
        self, email: Optional[str] = None, password: Optional[str] = None
    ) -> None:
        """WebSocket の接続を確立する

        Args:
            email (str, optional):
            password (str, optional):
        """
        if email and password:
            await self.__client.auth.login(email, password)

        if self.__ws_token is None:
            self.__ws_token = (await self.__client.misc.get_web_socket_token()).token

        self.__session = aiohttp.ClientSession()
        async with self.__session.ws_connect(
            f"wss://{config.CABLE_HOST}/?token={self.__ws_token}&app_version={config.VERSION_NAME}"
        ) as ws:
            self.__ws = ws

            await self.__on_open()

            async for msg in ws:
                match msg.type:
                    case aiohttp.WSMsgType.TEXT:
                        await self.__on_message(msg.json())
                    case aiohttp.WSMsgType.ERROR:
                        await self.__on_error(msg.json())
                    case aiohttp.WSMsgType.CLOSE:
                        await self.__on_close(msg.json())

    def run(self, email: Optional[str] = None, password: Optional[str] = None) -> None:
        """WebSocket の接続を確立する

        Args:
            email (str, optional):
            password (str, optional):
        """
        asyncio.run(self.__start_ws(email, password))

    async def stop(self) -> None:
        """WebSocket の接続を終了する"""
        if self.__ws is not None:
            await self.__ws.close()

        if self.__session is not None:
            await self.__session.close()
