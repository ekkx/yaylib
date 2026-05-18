"""イベントストリームで新着チャットメッセージを受信する簡単なボット。
Ctrl-C で終了します。

    YAY_EMAIL=... YAY_PASSWORD=... python examples/03_event_stream.py
"""

import asyncio
import contextlib
import os

import yaylib


async def main() -> None:
    async with yaylib.Client() as client:
        await client.login_with_email(
            email=os.environ["YAY_EMAIL"],
            password=os.environ["YAY_PASSWORD"],
        )

        async with client.open_event_stream() as stream:
            sub = await stream.subscribe(yaylib.chat_room_channel())
            print("listening for events (Ctrl-C to stop)")

            @sub.on_new_message
            async def _(event):
                print("new message:", event)

            @sub.on_event
            async def _any(event):
                print("event:", type(event).__name__)

            await sub.done()


with contextlib.suppress(KeyboardInterrupt):
    asyncio.run(main())
