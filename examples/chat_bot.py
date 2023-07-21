import yaylib

api = yaylib.Client()

email = ""
password = ""

api.login(email, password)


class ChatBot(yaylib.ChatRoomEventHandler):
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
