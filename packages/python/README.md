<div align="center">

<img src="https://github.com/ekkx/yaylib/assets/77382767/45c45b21-d812-4cad-8f27-315ffef53201" alt="yaylib" height="160px">

### yaylib

**好きでつながるバーチャルワールド - Yay! の API ライブラリ**

<a href="https://github.com/ekkx/yaylib/tree/master/packages/python"><img src="https://img.shields.io/badge/Python-3776AB?style=flat-square&logo=python&logoColor=white" alt="Python"></a>

<sub>あらゆる操作の自動化や、ボットの開発が可能です。</sub>

<a href="https://discord.gg/MEuBfNtqRN">Discord に参加</a>

</div>

Yay! の API を Python から扱うための非公式 SDK です。

### インストール

```sh
pip install yaylib
```

対応バージョン: Python 3.10 / 3.11 / 3.12 / 3.13

### クイックスタート

```python
import asyncio
import yaylib


async def main():
    # `async with` で接続が自動的にクローズされます
    async with yaylib.Client() as client:
        # ログイン（セッションは透過的にキャッシュされます）
        await client.login_with_email(email="...", password="...")

        # タイムラインを取得
        timeline = await client.get_timeline(
            noreply_mode=yaylib.NoreplyMode.EMPTY, number=20
        )
        for post in timeline.posts:
            print(post.id, post.text)

        # 投稿する
        await client.create_post(
            x_jwt=client.generate_x_jwt(),
            post_type="text",
            text="hello from yaylib",
        )


asyncio.run(main())
```

すべてのオペレーションは `client.<operation>` として直接呼び出せます。

### イベントストリーム

```python
async with client.open_event_stream() as stream:
    sub = stream.subscribe(yaylib.chat_room_channel())

    # 新着メッセージを受信
    @sub.on_new_message
    async def _(event):
        print("new message:", event)

    await sub.wait_closed()
```

### サンプル

実行できるサンプルを [`examples/python/`](https://github.com/ekkx/yaylib/tree/master/examples/python) に用意しています。

- `01_timeline.py` — 認証 + タイムライン取得
- `02_post.py` — テキスト投稿
- `03_event_stream.py` — イベントストリームの簡単なボット
- `04_session_and_errors.py` — セッション永続化とエラー処理

```bash
YAY_EMAIL=... YAY_PASSWORD=... python examples/python/01_timeline.py
```

### ⚖️ ライセンス

<p align="center">
  <a href="https://github.com/ekkx">
    <img src="https://github.com/ekkx/yaylib/assets/77382767/5d6aef18-5d98-4c9b-9f54-791308b393af" width="256" height="256">
  </a>
</p>

<p align="center">
  <strong>MIT © <a href="https://github.com/ekkx">ekkx</a></strong>
</p>
