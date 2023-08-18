# タイムラインのいいねボット用のサンプルコード


email = "your_email"
password = "your_password"


import time
import random
import yaylib


class LikeBot:
    def __init__(self, email=None, password=None, secret_token=None):
        self.api = yaylib.Client()
        self.api.login(email, password, secret_token)

    def delay(self):
        # レート制限を緩和するために遅延を挿入する
        sleep_time = random.uniform(0.2, 1.0)
        time.sleep(sleep_time)

    def run(self, amount=None):
        amount = float("inf") if amount is None else amount
        min_collect = 30
        liked = 0

        while liked < amount:
            ids = []

            self.api.logger.info("投稿を取得しています...")

            while len(ids) < min_collect:
                timeline = self.api.get_timeline(number=100)

                new_ids = [post.id for post in timeline.posts if not post.liked]

                ids.extend(new_ids)
                self.api.logger.info(f"取得済み投稿数: {len(ids)}")

                if len(ids) < min_collect:
                    time.sleep(5)

            for id in ids:
                self.api.like(id)
                self.delay()

            liked += len(ids)
            self.api.logger.info(f"いいね数: {liked}")
            ids.clear()

        self.api.logger.info(f"合計{liked}個の投稿にいいねしました。")


if __name__ == "__main__":
    bot = LikeBot(email, password)
    bot.run()
