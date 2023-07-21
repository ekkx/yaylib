"""
MIT License

Copyright (c) 2023-present Qvco, Konn

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

import json
import websocket

from .config import Configs
from .models import Message, ChatRoom, GroupUpdatesEvent
from .responses import ChannelResponse


class WebSocketBaseHandler(object):
    """イベントハンドラーの基底クラス"""

    def __init__(self):
        self.ws = None

    def _on_open(self, ws):
        pass

    def _on_message(self, ws, message):
        pass

    def _on_error(self, ws, error):
        print(error)

    def _on_close(self, ws):
        print("WebSocket Closed.")

    def on_connect(self, sid: str):
        pass

    def run(self, ws_token: str):
        self.ws = websocket.WebSocketApp(
            url=f"wss://{Configs.YAY_CABLE_HOST}/?token={ws_token}&app_version={Configs.YAY_VERSION_NAME}",
            on_open=self._on_open,
            on_message=self._on_message,
            on_error=self._on_error,
            on_close=self._on_close,
        )
        self.ws.run_forever()


class MessageEventHandler(WebSocketBaseHandler):
    """

    特定のチャットルームのメッセージイベントを取得します

    Methods
    -------

        - on_message(message: Message): 新しいメッセージを受信したときに呼び出されます

    Example
    -------

        >>> import yaylib
        >>>
        >>> class MyHandler(yaylib.MessageEventHandler):
        >>>     def on_message(self, message):
        >>>         print(message.text)
        >>>
        >>> api = yaylib.Client()
        >>>
        >>> ws_token = api.get_web_socket_token()
        >>> bot = MyHandler()
        >>> bot.run(ws_token)

    """

    def __init__(self, chat_room_id: int):
        super().__init__()
        self.chat_room_id = chat_room_id

    def _on_open(self, ws):
        ws.send(
            json.dumps(
                {
                    "command": "subscribe",
                    "identifier": f'{{"channel":"MessagesChannel", "chat_room_id": {self.chat_room_id}}}',
                }
            )
        )

    def _on_message(self, ws, message):
        message = ChannelResponse(json.loads(message))

        if message.identifier is not None and message.type is None:
            self.on_message(Message(message.message.data))

    def on_message(self, message: Message):
        pass


class ChatRoomEventHandler(WebSocketBaseHandler):
    """

    チャットルームのイベントを取得します

    Methods
    -------

        - on_message(chat_room: ChatRoom): 新しいメッセージを受信したときに呼び出されます
        - on_delete(room_id: int): チャットルームが削除されたときに呼び出されます

    Example
    -------

        >>> import yaylib
        >>>
        >>> class MyHandler(yaylib.ChatRoomEventHandler):
        >>>     def on_message(self, chat_room):
        >>>         print(chat_room.last_message.text)
        >>>
        >>>     def on_delete(self, room_id):
        >>>         print(room_id)
        >>>
        >>> api = yaylib.Client()
        >>>
        >>> ws_token = api.get_web_socket_token()
        >>> bot = MyHandler()
        >>> bot.run(ws_token)

    """

    def __init__(self):
        super().__init__()

    def _on_open(self, ws):
        ws.send(
            json.dumps(
                {
                    "command": "subscribe",
                    "identifier": '{"channel":"ChatRoomChannel"}',
                }
            )
        )

    def _on_message(self, ws, message):
        message = ChannelResponse(json.loads(message))

        if message.identifier is not None and message.type is None:
            if "event" not in message.message:
                self.on_message(ChatRoom(message.message.data.get("chat")))
            elif message.get("event") == "chat_deleted":
                self.on_delete(message.get("data").get("room_id"))

    def on_message(self, chat_room: ChatRoom):
        pass

    def on_delete(self, room_id: int):
        pass


class GroupUpdateEventHandler(WebSocketBaseHandler):
    """

    Yay!に存在する全てのサークルのイベントを取得します

    ※ イベントが発生してから約1分遅れて送信されます。

    Methods
    -------

        - on_post(group_id: int): サークルに投稿されたときに呼び出されます

    Example
    -------

        >>> import yaylib
        >>>
        >>> class MyHandler(yaylib.GroupUpdateEventHandler):
        >>>     def on_post(self, group_id):
        >>>         print(group_id)
        >>>
        >>> api = yaylib.Client()
        >>>
        >>> ws_token = api.get_web_socket_token()
        >>> bot = MyHandler()
        >>> bot.run(ws_token)

    """

    def __init__(self):
        super().__init__()

    def _on_open(self, ws):
        ws.send(
            json.dumps(
                {
                    "command": "subscribe",
                    "identifier": '{"channel":"GroupUpdatesChannel"}',
                }
            )
        )

    def _on_message(self, ws, message):
        message = ChannelResponse(json.loads(message))

        if message.type == "welcome":
            self.on_connect(message.sid)

        if message.identifier is not None and message.type is None:
            message = GroupUpdatesEvent(message.message.response)

            if message.event == "new_post":
                self.on_post(message.data.get("group_id"))

    def on_connect(self, sid: str):
        pass

    def on_post(self, group_id: int):
        pass


class GroupPostEventHandler(WebSocketBaseHandler):
    """

    特定のサークルの投稿イベントを取得します

    ※ イベントが発生してから約1分遅れて送信されます。

    Methods
    -------

        - on_post(group_id: int): サークルに投稿されたときに呼び出されます

    Example
    -------

        >>> import yaylib
        >>>
        >>> class MyHandler(yaylib.GroupPostEventHandler):
        >>>     def on_post(self, group_id):
        >>>         print(group_id)
        >>>
        >>> api = yaylib.Client()
        >>>
        >>> ws_token = api.get_web_socket_token()
        >>> bot = MyHandler()
        >>> bot.run(ws_token)

    """

    def __init__(self, group_id: int):
        super().__init__()
        self.group_id = group_id

    def _on_open(self, ws):
        ws.send(
            json.dumps(
                {
                    "command": "subscribe",
                    "identifier": f'{{"channel":"GroupPostsChannel", "group_id": {self.group_id}}}',
                }
            )
        )

    def _on_message(self, ws, message):
        message = json.loads(message)

        # if "identifier" in message and "type" not in message:
        #     message = GroupUpdateEvent(WebSocketResponse(message).message)
        #     if message.event == "new_post":
        #         self.on_post(message.group_id)

    def on_post(self, group_id: int):
        pass
