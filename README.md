<div><a id="readme-top"></a></div>
<div align="center">
    <img src="https://img.shields.io/github/stars/ekkx/yaylib?style=for-the-badge&logo=appveyor&color=blue" />
    <img src="https://img.shields.io/github/forks/ekkx/yaylib?style=for-the-badge&logo=appveyor&color=blue" />
    <img src="https://img.shields.io/github/issues/ekkx/yaylib?style=for-the-badge&logo=appveyor&color=informational" />
    <img src="https://img.shields.io/github/issues-pr/ekkx/yaylib?style=for-the-badge&logo=appveyor&color=informational" />
</div>
<br />
<p align="center">
    <a href="https://ekkx.github.io/yaylib">
        <img src="https://github.com/ekkx/yaylib/assets/77382767/45c45b21-d812-4cad-8f27-315ffef53201" alt="Logo" height="300px">
    </a>
    <h3 align="center">yaylib</h3>
    <p align="center">
        好きでつながるバーチャルワールド - Yay!（イェイ）の API ライブラリ<br />
        あらゆる操作の自動化や、ボットの開発が可能です。
        <br />
        <br />
        <a href="https://github.com/ekkx/yay.js">
            <strong>Node.js 版はこちらから »</strong>
        </a>
        <br />
        <br />
        <a href="https://ekkx.github.io/yaylib">ドキュメント</a>
        ·
        <a href="https://github.com/ekkx/yaylib/issues/new">バグを報告</a>
        ·
        <a href="https://discord.gg/MEuBfNtqRN">Discord に参加</a>
    </p>
</p>

<br>

<!-- インストール -->

## [<img src="https://github.com/ekkx/yaylib/assets/77382767/2f632349-0cbc-4c81-bc19-11d24c8c142b" width="30" height="30" />](https://github.com/ekkx) Installation

**yaylib** は `pip` コマンドからインストールします。

```shell
pip install yaylib
```

> [!TIP]
> 動作条件は `Python 3.10` 以上からです。

<br>

<!-- 使用例 -->

## [<img src="https://github.com/ekkx/yaylib/assets/77382767/dc7dcea0-c581-4039-8fc2-3994884d2ba3" width="30" height="30" />](https://github.com/ekkx) Quick Example

#### ✨ 投稿を作成する

```python
import yaylib

bot = yaylib.Client()
bot.login('your_email', 'your_password')

bot.create_post('Hello with yaylib!')
```

#### ✨ タイムラインを取得する

```python
import yaylib

bot = yaylib.Client()

timeline = bot.get_timeline(number=100)

for post in timeline.posts:
    print(post.user.nickname)  # 投稿者名
    print(post.text)  # 本文
    print(post.likes_count)  # いいね数
    print(post.reposts_count)  # (´∀｀∩)↑age↑の数
    print(post.in_reply_to_post_count)  # 返信の数
```

#### ✨ タイムラインをキーワードで検索して「いいね」する

```python
import yaylib

bot = yaylib.Client()
bot.login('your_email', 'your_password')

timeline = bot.get_timeline_by_keyword(
    keyword='プログラミング',
    number=15
)

for post in timeline.posts:
    bot.like(post.id)
```

#### ✨ 画像と一緒に投稿を作成する

```python
import yaylib

bot = yaylib.Client()
bot.login('your_email', 'your_password')

# 画像のパスを指定
image_paths = [
    './test1.jpg',
    './test2.jpg',
    './test3.jpg',
]

# 画像の使い道を指定
image_type = yaylib.ImageType.POST

# サーバー上にアップロード
attachments = bot.upload_image(image_paths, image_type)

# サーバー上のファイル名を指定する
# attachmentsが一つ飛ばしなのはオリジナル品質の画像のみを指定するため
bot.create_post(
    'Hello with yaylib!',
    attachment_filename=attachments[0].filename,
    attachment_2_filename=attachments[2].filename,
    attachment_3_filename=attachments[4].filename,
)
```

#### ✨ 新規ユーザーをフォローする

```python
import yaylib

bot = yaylib.Client()
bot.login('your_email', 'your_password')

new_users = bot.search_users(recently_created=True)

for new_user in new_users.users:
    bot.follow_user(new_user.id)
```

#### ✨ リアルタイムでチャットを取得する

```python
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
```

より詳しい使用例については、[ドキュメント](https://ekkx.github.io/yaylib/demo.html)を参照してください。

<br>

<!-- yaylib で誕生したボットの一覧 -->

## 👑 yaylib で誕生したロボットたち

「yaylib」を用いて開発したロボットがある場合は、ぜひ教えてください！

<table align="center">
    <thead>
        <tr>
            <th><a href="https://yay.space/user/5855987">MindReader AI</a></th>
            <th><a href="https://yay.space/user/8271084">めいく</a></th>
            <th><a href="https://yay.space/user/7406336">GIGAZINE</a></th>
        </tr>
    </thead>
    <tbody>
        <tr>
            <td align="center">
                <img src="https://github.com/ekkx/yaylib/assets/77382767/cc41ce3c-0e11-4ec5-be99-ff7090a95667" width="200px">
                <br />
                <p>開発者: <a href="https://yay.space/user/35152">毛の可能性</a></p>
            </td>
            <td align="center">
                <img src="https://github.com/user-attachments/assets/201cb490-29b7-4dd9-a10f-1b27d999787a" width="200px">
                <br />
                <p>開発者: <a href="https://yay.space/user/1173972">まぐ</a></p>
            </td>
            <td align="center">
                <img src="https://github.com/ekkx/yaylib/assets/77382767/65fcb885-4fbe-4170-9378-6f8d9af61ff8" width="200px">
                <br />
                <p>開発者: <a href="https://yay.space/user/1298298">ぺゅー</a></p>
            </td>
        </tr>
    </tbody>
</table>

<br>

<!-- 共同開発について -->

## 🤝 共同開発について

詳しい **yaylib** の開発参加手順については、[こちら](https://github.com/ekkx/yaylib/blob/develop/CONTRIBUTING.md)を参照してください。

<br>

<!-- 免責事項 -->

## 📜 免責事項

yaylib は、API の公式なサポートやメンテナンスを提供するものではありません。このクライアントを使用する場合、**利用者はリスクや責任を自己負担できるもの**とします。このクライアントによって提供される情報やデータの正確性、信頼性、完全性、適時性について、いかなる保証も行いません。また、このクライアントの使用によって生じた損害や不利益について、一切の責任を負いかねます。利用者は自己の責任において、このクライアントを使用し、API にアクセスするものとします。なお、この免責事項は予告なく変更される場合があります。

<br>

<!-- ライセンス -->

## ⚖️ ライセンス

<p align="center">
  <a href="https://github.com/ekkx">
    <img src="https://github.com/ekkx/yaylib/assets/77382767/5d6aef18-5d98-4c9b-9f54-791308b393af" width="256" height="256">
  </a>
</p>

<p align="center">
  <strong>MIT © <a href="https://github.com/ekkx">ekkx</a></strong>
</p>

フルライセンスは [こちら](https://github.com/ekkx/yaylib/blob/master/LICENSE) からご確認いただけます。  
このプロジェクトは、 **【MIT ライセンス】** の条件の下でライセンスされています。

<p align="right">(<a href="#readme-top">トップに戻る</a>)</p>
