"""テキスト投稿の例。投稿系のエンドポイントは X-Jwt を必要とします
（client.generate_x_jwt() で取得し、そのまま渡します。5 秒で失効するため
呼び出しごとに再生成してください）。

    YAY_EMAIL=... YAY_PASSWORD=... python examples/python/02_post.py
"""

import asyncio
import os

import yaylib


async def main() -> None:
    async with yaylib.Client() as client:
        await client.login_with_email(
            email=os.environ["YAY_EMAIL"],
            password=os.environ["YAY_PASSWORD"],
        )

        post = await client.create_post(
            x_jwt=client.generate_x_jwt(),
            post_type="text",
            text="hello from yaylib",
        )
        print("posted:", post.id)


asyncio.run(main())
