import yaylib

api = yaylib.Client()

email = ""
password = ""

api.login(email, password)


class ChatBot(yaylib.ChatRoomEventHandler):

    def on_request(self, total_count: int):
        # チャットリクエストを承認する
        chat_room = api.get_request_chat_rooms().chat_rooms[0]
        api.accept_chat_request(chat_room_ids=[chat_room.id])
        self.on_message(chat_room)

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
