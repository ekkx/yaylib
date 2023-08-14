# リアルタイムで個人チャットを取得するサンプルコード
# 例): メッセージが届いたチャットルームにオウム返しする場合


email = "your_email"
password = "your_password"


import yaylib


class ChatBot(yaylib.ChatRoomEventHandler):
    def __init__(self, api: yaylib.Client):
        self.api = api

    def on_request(self, total_count: int):
        # チャットリクエストを承認する
        chat_rooms = self.api.get_chat_requests()
        for room in chat_rooms.chat_rooms:
            self.api.accept_chat_requests(chat_room_ids=[room.id])
        self.on_message(chat_rooms.chat_rooms[0])

    def on_message(self, chat_room):
        # 受信したメッセージをオウム返しする
        self.api.send_message(
            chat_room_id=chat_room.id,
            message_type=chat_room.last_message.message_type,
            text=chat_room.last_message.text,
            font_size=chat_room.last_message.font_size,
        )


api = yaylib.Client()
api.login(email, password)

ws_token = api.get_web_socket_token()

bot = ChatBot(api)
bot.run(ws_token)
