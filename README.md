<div><a id="readme-top"></a></div>
<div align="center">
    <img src="https://img.shields.io/github/stars/qvco/yaylib?style=for-the-badge&logo=appveyor&color=blue" />
    <img src="https://img.shields.io/github/forks/qvco/yaylib?style=for-the-badge&logo=appveyor&color=blue" />
    <img src="https://img.shields.io/github/issues/qvco/yaylib?style=for-the-badge&logo=appveyor&color=informational" />
    <img src="https://img.shields.io/github/issues-pr/qvco/yaylib?style=for-the-badge&logo=appveyor&color=informational" />
</div>
<br />
<p align="center">
    <a href="https://github.com/othneildrew/Best-README-Template">
        <img src="https://github.com/qvco/yaylib/assets/77382767/45c45b21-d812-4cad-8f27-315ffef53201" alt="Logo" height="300px">
    </a>
    <h3 align="center">yaylib</h3>
    <p align="center">
        同世代でつながるチャットアプリ、Yay!（イェイ）の API ラッパーです。<br />
        あらゆる操作の自動化や、ボットの開発が可能です。
        <br />
        <br />
        <a href="https://github.com/qvco/yaylib/blob/master/docs/README.md">
            <strong>ドキュメントはこちらから »</strong>
        </a>
        <br />
        <br />
        <a href="https://github.com/qvco/yaylib/issues">Report Bug</a>
        ·
        <a href="https://github.com/qvco/yaylib/issues">Request Feature</a>
        ·
        <a href="https://discord.gg/MEuBfNtqRN">Join the discord</a>
    </p>
</p>

<!-- TABLE OF CONTENTS -->

<details>
  <summary>Table of Contents</summary>
  <ol>
    <li><a href="#-installation">インストール</a></li>
    <li><a href="#-quick-example">使用例</a></li>
    <li><a href="#crown-yaylib-で誕生したロボットたち">yaylib で誕生したロボットたち</a></li>
    <li><a href="#buy-me-a-coffee">Buy me a coffee</a></li>
    <li><a href="#handshake-共同開発について">共同開発について</a></li>
    <li><a href="#免責事項">免責事項</a></li>
    <li><a href="#利用許諾">利用許諾</a></li>
  </ol>
</details>

<!-- インストール -->

## [<img src="https://github.com/qvco/yaylib/assets/77382767/2f632349-0cbc-4c81-bc19-11d24c8c142b" width="30" height="30" />](https://github.com/qvco) Installation

**※ Python 3.10 以上のバージョンが必要です。**

「yaylib」をインストールするには、以下のコマンドをターミナル上で実行します:

```bash
pip install yaylib
```

<br>

> **Note**
> 開発バージョンをインストールする場合は、以下の手順を実行します:

```bash
git clone https://github.com/qvco/yaylib.git

cd yaylib

pip install -r requirements.txt

pip install -e .
```

「yaylib」の始め方については、[こちら](https://github.com/qvco/yaylib/blob/master/docs/TUTORIAL.md) を確認してください。

<!-- 使用例 -->

## [<img src="https://github.com/qvco/yaylib/assets/77382767/dc7dcea0-c581-4039-8fc2-3994884d2ba3" width="30" height="30" />](https://github.com/qvco) Quick Example

<!-- ### yaylib 利用で実現できること

<table>
    <tr>
        <th>カテゴリ</th>
        <th>「yaylib」で自動化できること</th>
        <th>応用先</th>
    </tr>
    <tr>
		<td>データ収集の効率化</td>
		<td>
            <li>投稿情報取得</li>
            <li>ユーザー情報取得</li>
            <li>人気のワード取得</li>
            <li>フォロー/フォロワー情報取得</li>
        </td>
		<td>
            <li>SNSデータマーケティング</li>
            <li>トレンド解析</li>
            <li>ニーズ調査</li>
        </td>
	</tr>
    <tr>
		<td>アカウント操作<br>(投稿関連)</td>
		<td>
            <li>投稿する</li>
            <li>リツイートする</li>
            <li>いいねする</li>
        </td>
		<td>
            <li>Yay! 自動運用</li>
            <li>Yay! Bot開発</li>
        </td>
	</tr>
    <tr>
		<td>アカウント操作<br>(ユーザー関連)</td>
		<td>
            <li>フォローする/フォロー解除する</li>
            <li>ブロックする/ブロック解除する</li>
            <li>ミュートする/ミュート解除する</li>
            <li>DMを送る</li>
        </td>
		<td>
            <li>Yay! 自動運用</li>
            <li>Yay! Bot開発</li>
        </td>
	</tr>
</table> -->

#### ✨ 投稿を作成する

```python
import yaylib

api = yaylib.Client()
api.login(email="メールアドレス", password="パスワード")

api.create_post("Hello with yaylib!")
```

#### ✨ 埋め込みリンクの投稿を作成する

```python
import yaylib

api = yaylib.Client()
api.login(email="メールアドレス", password="パスワード")

api.create_post("Hello with yaylib!", shared_url="https://github.com/qvco/yaylib")
```

#### ✨ 画像と一緒に投稿を作成する

```python
import yaylib

api = yaylib.Client()
api.login(email="メールアドレス", password="パスワード")

# サーバー上にアップロード
filenames = api.upload_image(
    image_type=yaylib.IMAGE_TYPE_POST,
    image_paths=["./test1.jpg", "./test2.jpg", "./test3.jpg"],
)

# サーバー上のファイル名を指定する
# filenamesが一つ飛ばしなのはオリジナル品質の画像のみを指定するため
api.create_post(
    "Hello with yaylib!",
    attachment_filename=filenames[0],
    attachment_2_filename=filenames[2],
    attachment_3_filename=filenames[4],
)
```

#### ✨ 投稿に返信する

```python
import yaylib

api = yaylib.Client()
api.login(email="メールアドレス", password="パスワード")

api.create_post(
    "Hello with yaylib!",
    in_reply_to=返信先の投稿ID,
    mention_ids=[返信先のユーザーID]
)
```

#### ✨ タイムラインを 100 件取得する

```python
import yaylib

api = yaylib.Client()

timeline = api.get_timeline(number=100)

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

api = yaylib.Client()
api.login(email="メールアドレス", password="パスワード")

timeline = api.get_timeline_by_keyword(
    keyword="プログラミング",
    number=15
)

for post in timeline.posts:
    response = api.like(post.id)
    print(post.id, response.data) # 実行結果を出力
```

#### ✨ タイムラインのユーザーをフォローする

```python
import yaylib

api = yaylib.Client()
api.login(email="メールアドレス", password="パスワード")

timeline = api.get_timeline(number=15)

for post in timeline.posts:
    api.follow_user(post.user.id)
```

#### ✨ リアルタイムでチャットを取得する

```python
import yaylib

api = yaylib.Client()
api.login(email="メールアドレス", password="パスワード")


class ChatBot(yaylib.ChatRoomEventHandler):

    def on_request(self, total_count: int):
        # チャットリクエストを承認する
        chat_room = api.get_chat_requests(number=1).chat_rooms[0]
        api.accept_chat_requests(chat_room_ids=[chat_room.id])
        self.on_message(chat_room)

    def on_message(self, chat_room):
        # メッセージを出力する
        print(chat_room.last_message.text)


ws_token = api.get_web_socket_token()

bot = ChatBot()
bot.run(ws_token)
```

より詳しい使用例については、[こちら](https://github.com/qvco/yaylib/blob/master/examples) を参照してください。

<p align="right">(<a href="#readme-top">トップに戻る</a>)</p>

<!-- yaylib で誕生したボットの一覧 -->

## :crown: yaylib で誕生したロボットたち

「yaylib」を用いて開発したロボットがある場合は、ぜひ教えてください！

<table align="center">
    <thead>
        <tr>
            <th><a href="https://yay.space/user/5855987">MindReader AI</a></th>
            <th><a href="https://yay.space/user/7293290">香ばしいボット</a></th>
            <th><a href="https://yay.space/user/7406336">GIGAZINE</a></th>
        </tr>
    </thead>
    <tbody>
        <tr>
            <td align="center">
                <img src="https://github.com/qvco/yaylib/assets/77382767/cc41ce3c-0e11-4ec5-be99-ff7090a95667" width="200px">
                <br />
                <p>開発者: <a href="https://yay.space/user/35152">毛の可能性</a></p>
            </td>
            <td align="center">
                <img src="https://github.com/qvco/yaylib/assets/77382767/cbffdc25-7873-4242-b065-e6a686bade54" width="200px">
                <br />
                <p>開発者: <a href="https://yay.space/user/93923">めんぶれ天然水。</a></p>
            </td>
            <td align="center">
                <img src="https://github.com/qvco/yaylib/assets/77382767/65fcb885-4fbe-4170-9378-6f8d9af61ff8" width="200px">
                <br />
                <p>開発者: <a href="https://yay.space/user/1298298">ぺゅー</a></p>
            </td>
        </tr>
    </tbody>
</table>

<!-- 共同開発について -->

## :handshake: 共同開発について

私たちと開発することに興味を持っていただけているなら、ぜひ参加して頂きたいです！

- <a href="https://github.com/qvco/yaylib/pulls">プルリクエストを送信する</a>
- <a href="https://discord.gg/MEuBfNtqRN">Discord サーバーに参加する</a>
- <a href="mailto:nikola.desuga@gmail.com">nikola.desuga@gmail.com にメールを送信する</a>

のいずれかの方法で繋がりましょう。詳しくは[こちらから](https://github.com/qvco/yaylib/blob/master/CONTRIBUTING.md)！

<!-- Buy me a coffee -->

## Buy me a coffee

このライブラリが気に入っていただけたら、**リポジトリに<a href="https://github.com/qvco/yaylib/">スターをお願いします</a>(⭐)**  
また、Buy Me a Coffee からご支援いただけますと幸いです。

<a href="https://www.buymeacoffee.com/qvco" target="_blank"><img src="https://www.buymeacoffee.com/assets/img/custom_images/orange_img.png" alt="Buy Me A Coffee" style="height: auto !important;width: auto !important;" ></a>

<!-- 免責事項 -->

## 免責事項

yaylib は、API の公式なサポートやメンテナンスを提供するものではありません。このクライアントを使用する場合、**利用者はリスクや責任を自己負担できるもの**とします。このクライアントによって提供される情報やデータの正確性、信頼性、完全性、適時性について、いかなる保証も行いません。また、このクライアントの使用によって生じた損害や不利益について、一切の責任を負いかねます。利用者は自己の責任において、このクライアントを使用し、API にアクセスするものとします。なお、この免責事項は予告なく変更される場合があります。

<!-- ライセンス -->

## ライセンス

<p align="center">
  <a href="https://github.com/qvco">
    <img src="https://github.com/qvco/yaylib/assets/77382767/5d6aef18-5d98-4c9b-9f54-791308b393af" width="256" height="256">
  </a>
</p>

<p align="center">
  <strong>MIT © <a href="https://github.com/qvco">Qvco</a> & <a href="https://github.com/konn-koko">Konn</a></strong>
</p>

フルライセンスは [こちら](https://github.com/qvco/yaylib/blob/master/LICENSE) からご確認いただけます。  
このプロジェクトは、 **【MIT ライセンス】** の条件の下でライセンスされています。

<p align="right">(<a href="#readme-top">トップに戻る</a>)</p>
