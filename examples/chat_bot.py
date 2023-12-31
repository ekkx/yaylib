# リアルタイムで個人チャットを取得するサンプルコード


email = "your_email"
password = "your_password"


import yaylib
from yaylib import Message


class MyBot(yaylib.Client):
    def on_ready(self):
        print("ボットがオンラインになりました！")

    def on_chat_request(self, total_count):
        # チャットリクエストはすべて承認する
        chat_requests = self.get_chat_requests()
        for chat_room in chat_requests.chat_rooms:
            self.accept_chat_requests([chat_room.id])

        # 最新のメッセージをon_message関数に送信
        message = self.get_messages(chat_requests.chat_rooms[0].id)
        self.on_message(message[0])

    def on_message(self, message: Message):
        # 「ping」というメッセージに対して「pong」と返信する
        if message.text == "ping":
            self.send_message(message.room_id, text="pong")

    def on_chat_room_delete(self, room_id):
        print(f"チャットルームが削除されました。ルームID: {room_id}")


intents = yaylib.Intents.none()
intents.chat_message = True

bot = MyBot(intents=intents)
bot.run(email, password)
