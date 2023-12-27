# タイムラインのいいねボット用のサンプルコード


email = "your_email"
password = "your_password"


import time
import yaylib


class LikeBot:
    def __init__(self, email=None, password=None):
        self.client = yaylib.Client()
        self.client.login(email, password)

    def main(self):
        min_collect = 30
        liked = 0
        collected_ids = set()

        while True:
            ids = []

            self.client.logger.info("投稿を取得しています...")

            while True:
                timeline = self.client.get_timeline(number=100)

                new_ids = [
                    post.id
                    for post in timeline.posts
                    if not post.liked and post.id not in collected_ids
                ]

                ids.extend(new_ids)
                collected_ids.update(new_ids)

                self.client.logger.info(f"取得済み投稿数: {len(ids)}")

                if len(ids) >= min_collect:
                    break

                time.sleep(10)

            for id in ids:
                self.client.like(id)
                self.client.logger.info("いいねしました。")

            liked += len(ids)
            self.client.logger.info(f"いいね数: {liked}")

            ids.clear()
            collected_ids.clear()


if __name__ == "__main__":
    bot = LikeBot(email, password)
    bot.main()
