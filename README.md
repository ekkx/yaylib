<div><a id="readme-top"></a></div>
<div align="center">
    <img src="https://img.shields.io/github/stars/ekkx/yaylib?style=for-the-badge&logo=appveyor&color=blue" />
    <img src="https://img.shields.io/github/forks/ekkx/yaylib?style=for-the-badge&logo=appveyor&color=blue" />
    <img src="https://img.shields.io/github/issues/ekkx/yaylib?style=for-the-badge&logo=appveyor&color=informational" />
    <img src="https://img.shields.io/github/issues-pr/ekkx/yaylib?style=for-the-badge&logo=appveyor&color=informational" />
</div>
<br />
<p align="center">
    <a href="https://github.com/othneildrew/Best-README-Template">
        <img src="https://github.com/ekkx/yaylib/assets/77382767/45c45b21-d812-4cad-8f27-315ffef53201" alt="Logo" height="300px">
    </a>
    <h3 align="center">yaylib</h3>
    <p align="center">
        同世代とつながる通話コミュニティ - Yay!（イェイ）の API ライブラリ<br />
        あらゆる操作の自動化や、ボットの開発が可能です。
        <br />
        <br />
        <a href="https://github.com/ekkx/yay.js">
            <strong>Node.js 版はこちらから »</strong>
        </a>
        <br />
        <br />
        <a href="https://github.com/ekkx/yaylib/issues">Report Bug</a>
        ·
        <a href="https://github.com/ekkx/yaylib/issues">Request Feature</a>
        ·
        <a href="https://discord.gg/MEuBfNtqRN">Join the discord</a>
    </p>
</p>

<!-- TABLE OF CONTENTS -->

<details>
  <summary>Table of Contents</summary>
  <ol>
    <li><a href="#buy-me-a-coffee">Buy me a coffee</a></li>
    <li><a href="#-installation">インストール</a></li>
    <li><a href="#-quick-example">使用例</a></li>
    <li><a href="#crown-yaylib-で誕生したロボットたち">yaylib で誕生したロボットたち</a></li>
    <li><a href="#handshake-共同開発について">共同開発について</a></li>
    <li><a href="#免責事項">免責事項</a></li>
    <li><a href="#利用許諾">利用許諾</a></li>
  </ol>
</details>

<!-- Buy me a coffee -->

## Buy me a coffee

このライブラリが気に入っていただけたら、<a href="https://github.com/ekkx/yaylib/">**リポジトリにスターをお願いします</a>(⭐)**  
また、Buy Me a Coffee からご支援いただけますと幸いです。

<a href="https://www.buymeacoffee.com/qvco" target="_blank"><img src="https://www.buymeacoffee.com/assets/img/custom_images/orange_img.png" alt="Buy Me A Coffee" style="height: auto !important;width: auto !important;" ></a>

<!-- インストール -->

## [<img src="https://github.com/ekkx/yaylib/assets/77382767/2f632349-0cbc-4c81-bc19-11d24c8c142b" width="30" height="30" />](https://github.com/ekkx) Installation

**※ Python 3.10 以上のバージョンが必要です。**

「yaylib」をインストールするには、以下のコマンドをターミナル上で実行します:

```shell
pip install yaylib
```

<br>

> [!TIP]
> 開発環境をインストールする場合は、以下の手順を実行します:

```shell
git clone https://github.com/ekkx/yaylib.git

cd yaylib

make up
```

<!-- 使用例 -->

## [<img src="https://github.com/ekkx/yaylib/assets/77382767/dc7dcea0-c581-4039-8fc2-3994884d2ba3" width="30" height="30" />](https://github.com/ekkx) Quick Example

#### ✨ 投稿を作成する

```python
import yaylib

client = yaylib.Client()
client.login(email="your_email", password="your_password")

client.create_post("Hello with yaylib!")
```

#### ✨ 埋め込みリンクの投稿を作成する

```python
import yaylib

client = yaylib.Client()
client.login(email="your_email", password="your_password")

client.create_post("Hello with yaylib!", shared_url="https://github.com/ekkx/yaylib")
```

#### ✨ 画像と一緒に投稿を作成する

```python
import yaylib

client = yaylib.Client()
client.login(email="your_email", password="your_password")

# 画像のパスを指定
image_paths = [
    "./test1.jpg",
    "./test2.jpg",
    "./test3.jpg",
]

# 画像の使い道を指定
image_type = yaylib.ImageType.post

# サーバー上にアップロード
attachments = client.upload_image(image_paths, image_type)

# サーバー上のファイル名を指定する
# attachmentsが一つ飛ばしなのはオリジナル品質の画像のみを指定するため
client.create_post(
    "Hello with yaylib!",
    attachment_filename=attachments[0].filename,
    attachment_2_filename=attachments[2].filename,
    attachment_3_filename=attachments[4].filename,
)
```

#### ✨ タイムラインを 100 件取得する

```python
import yaylib

client = yaylib.Client()

timeline = client.get_timeline(number=100)

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

client = yaylib.Client()
client.login(email="your_email", password="your_password")

timeline = client.get_timeline_by_keyword(
    keyword="プログラミング",
    number=15
)

for post in timeline.posts:
    client.like(post.id)
```

#### ✨ 新規ユーザーをフォローする

```python
import yaylib

client = yaylib.Client()
client.login(email="your_email", password="your_password")

new_users = client.search_users(recently_created=True)

for new_user in new_users.users:
    client.follow_user(new_user.id)
```

#### ✨ リアルタイムでチャットを取得する

```python
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

        # 最新のメッセージをon_message_create関数に送信
        message = self.get_messages(chat_requests.chat_rooms[0].id)
        self.on_message_create(message[0])

    def on_message_create(self, message: Message):
        # 「ping」というメッセージに対して「pong」と返信する
        if message.text == "ping":
            self.send_message(message.room_id, text="pong")

    def on_chat_room_delete(self, room_id):
        print(f"チャットルームが削除されました。ルームID: {room_id}")

intents = yaylib.Intents.none()
intents.chat_message = True

bot = MyBot(intents=intents)
bot.run("your_email", "your_password")
```

より詳しい使用例については、[こちら](https://github.com/ekkx/yaylib/blob/master/examples) を参照してください。

<p align="right">(<a href="#readme-top">トップに戻る</a>)</p>

<!-- yaylib で誕生したボットの一覧 -->

## :crown: yaylib で誕生したロボットたち

「yaylib」を用いて開発したロボットがある場合は、ぜひ教えてください！

<table align="center">
    <thead>
        <tr>
            <th><a href="https://yay.space/user/5855987">MindReader AI</a></th>
            <th><a href="https://yay.space/user/7874560">☀️気象くん☁️</a></th>
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
                <img src="https://github.com/ekkx/yaylib/assets/77382767/4fd728a0-9b3a-427f-ac1f-70e6d6538564" width="200px">
                <br />
                <p>開発者: <a href="https://yay.space/user/7520368">tori</a></p>
            </td>
            <td align="center">
                <img src="https://github.com/ekkx/yaylib/assets/77382767/65fcb885-4fbe-4170-9378-6f8d9af61ff8" width="200px">
                <br />
                <p>開発者: <a href="https://yay.space/user/1298298">ぺゅー</a></p>
            </td>
        </tr>
    </tbody>
</table>

<!-- 共同開発について -->

## :handshake: 共同開発について

私たちと開発することに興味を持っていただけているなら、ぜひ参加して頂きたいです！

- <a href="https://github.com/ekkx/yaylib/pulls">プルリクエストを送信する</a>
- <a href="https://discord.gg/MEuBfNtqRN">Discord サーバーに参加する</a>
- <a href="mailto:nikola.desuga@gmail.com">nikola.desuga@gmail.com にメールを送信する</a>

のいずれかの方法で繋がりましょう。詳しくは[こちらから](https://github.com/ekkx/yaylib/blob/master/CONTRIBUTING.md)！

<!-- 免責事項 -->

## 免責事項

yaylib は、API の公式なサポートやメンテナンスを提供するものではありません。このクライアントを使用する場合、**利用者はリスクや責任を自己負担できるもの**とします。このクライアントによって提供される情報やデータの正確性、信頼性、完全性、適時性について、いかなる保証も行いません。また、このクライアントの使用によって生じた損害や不利益について、一切の責任を負いかねます。利用者は自己の責任において、このクライアントを使用し、API にアクセスするものとします。なお、この免責事項は予告なく変更される場合があります。

<!-- ライセンス -->

## ライセンス

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
