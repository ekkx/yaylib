import yaylib


api = yaylib.Client()

api.login(email="your_email", password="your_password")


class ChatBot(yaylib.ChatRoomEventHandler):
    def on_request(self, total_count: int):
        # チャットリクエストを承認する
        chat_rooms = api.get_chat_requests()
        for room in chat_rooms.chat_rooms:
            api.accept_chat_requests(chat_room_ids=[room.id])
        self.on_message(chat_rooms.chat_rooms[0])

    def on_message(self, chat_room):
        # 受信したメッセージをオウム返しする
        api.send_message(
            chat_room_id=chat_room.id,
            message_type=chat_room.last_message.message_type,
            text=chat_room.last_message.text,
            font_size=chat_room.last_message.font_size,
        )


ws_token = api.get_web_socket_token()

bot = ChatBot()
bot.run(ws_token)
