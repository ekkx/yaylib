:description: yaylib におけるデモンストレーションのコードを紹介します。

使い方
=====

.. rst-class:: lead

   ここでは、yaylib の使い方の例を簡単に紹介します 💁‍♀️

----

当プロジェクトでは、オブジェクト指向を採用しており、プログラミングが初めての方でも簡単に扱える設計になっています。

いくつか例を紹介しますので、気になったコードがあれば適宜書き換えたりして遊んでみてください。

Timeline
--------

.. sidebar:: 解説

    タイムラインの投稿本文を取得するコードです。

    5 行目では、最新のタイムラインを 100 件取得しています。

    9 行目では、``プログラミング`` というワードが含まれる投稿を取得しています。

    13 行目では、サークル内のタイムラインを取得しています。

.. code-block:: python
    :caption: main.py
    :linenos:

    import yaylib

    client = yaylib.Client()

    timeline = client.get_timeline(number=100)
    for post in timeline.posts:
        print(post.text)

    timeline = client.get_timeline_by_keyword('プログラミング')
    for post in timeline.posts:
        print(post.text)

    timeline = client.get_group_timeline(group_id=149956)
    for post in timeline.posts:
        print(post.text)

Chat Bot
--------

.. sidebar:: 解説

    リアルタイム通信機能を使用して簡単なチャットボットを実装しています。

    一度もチャットをしたことがないユーザーともチャットができるよう、自動的にチャットリクエストを承認し、そのメッセージを 15 行目で ``on_message()`` 関数に送信しています。

    18 行目では、「ping」というメッセージに対して「pong」と返信しています。

    28 行目では、チャットメッセージのイベントを取得するように ``Intents`` を設定しています。

.. code-block:: python
    :caption: main.py
    :linenos:

    import yaylib

    class ChatBot(yaylib.Client):
        async def on_ready():
            print('Botがオンラインになりました！')

        async def on_chat_request(self, total_count):
            # チャットリクエストを承認し on_message() に送信する
            chat_requests = await self.chat.get_chat_requests()
            for chat_room in chat_requests.chat_rooms:
                await self.chat.accept_chat_requests(chat_room_ids=[chat_room.id])
            message = await self.chat.get_messages(chat_requests.chat_rooms[0].id)
            await self.on_message(message[0])

        async def on_message(self, message: yaylib.Message):
            if message.text == 'ping':
                await self.chat.send_message(
                    message.room_id,
                    text='pong',
                )

        async def on_chat_delete(self, room_id):
            print(f'チャットルームが削除されました。{room_id}')

    intents = yaylib.Intents.none()
    intents.chat_message = True

    bot = ChatBot(intents=intents)
    bot.run('your_email', 'your_password')
