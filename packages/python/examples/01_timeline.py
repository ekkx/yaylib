"""認証してタイムラインを取得する最小例。

    YAY_EMAIL=... YAY_PASSWORD=... python examples/01_timeline.py
"""

import asyncio
import os

import yaylib


async def main() -> None:
    # `async with` で基盤接続が確実にクローズされます。
    async with yaylib.Client() as client:
        await client.login_with_email(
            email=os.environ["YAY_EMAIL"],
            password=os.environ["YAY_PASSWORD"],
        )

        timeline = await client.get_timeline(
            noreply_mode=yaylib.NoreplyMode.EMPTY, number=20
        )
        for post in timeline.posts or []:
            print(post.id, post.text)


asyncio.run(main())
