# タイムラインのいいねボット用のサンプルコード


email = "your_email"
password = "your_password"


import time
import yaylib


class LikeBot:
    def __init__(self, email=None, password=None):
        self.api = yaylib.Client()
        self.api.login(email, password)

    def run(self):
        min_collect = 30
        liked = 0

        while True:
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

            liked += len(ids)
            self.api.logger.info(f"いいね数: {liked}")
            ids.clear()


if __name__ == "__main__":
    bot = LikeBot(email, password)
    bot.run()
