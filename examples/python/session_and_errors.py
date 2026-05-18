"""セッション永続化とエラー処理の例。

FileSessionStore を渡すと、ログイン状態がファイルに保存され、次回以降は
キャッシュされたセッションが再利用されます（毎回ログインしない）。
非 2xx は yaylib.APIError として送出され、code_of / error_response_of
で中身を判定できます。

    YAY_EMAIL=... YAY_PASSWORD=... python examples/python/session_and_errors.py
"""

import asyncio
import os

import yaylib


async def main() -> None:
    store = yaylib.FileSessionStore("yay-session.json")

    # 2 回目以降の実行では保存済みセッションが使われます。
    async with yaylib.Client(session_store=store) as client:
        await client.login_with_email(
            email=os.environ["YAY_EMAIL"],
            password=os.environ["YAY_PASSWORD"],
        )

        try:
            timeline = await client.get_timeline(
                noreply_mode=yaylib.NoreplyMode.EMPTY, number=20
            )
            print("posts:", len(timeline.posts or []))
        except yaylib.APIError as err:
            print("error code:", yaylib.code_of(err))  # grpc 的なコード分岐
            res = yaylib.error_response_of(err)  # ban_until / retry_in など
            if res is not None:
                print("message:", res.message)
            print("status:", err.status)  # 生のステータス


asyncio.run(main())
