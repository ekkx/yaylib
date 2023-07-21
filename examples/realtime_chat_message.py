import yaylib


email = ""
password = ""


api = yaylib.Client()

api.login(email, password)


class MyHandler(yaylib.ChatRoomEventHandler):
    def on_message(self, chat_room):
        # メッセージのテキストを出力
        print(chat_room.last_message.text)


ws_token = api.get_web_socket_token()

bot = MyHandler()
bot.run(ws_token)
