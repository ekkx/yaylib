import time
import yaylib


class LikeBot:
    def __init__(self, email=None, password=None, secret_token=None):
        self.api = yaylib.Client()
        self.api.login(email, password, secret_token)

    def run(self, amount=None):
        amount = float("inf") if amount is None else amount
        min_collect = 30
        liked = 0

        while liked < amount:
            ids = []

            try:
                self.api.logger.info("投稿を取得しています...")

                while len(ids) < min_collect:
                    timeline = self.api.get_timeline(number=min_collect)
                    new_ids = [post.id for post in timeline.posts if not post.liked]

                    ids.extend(new_ids)
                    self.api.logger.info(f"取得済み投稿数: {len(ids)}")

                    if len(ids) < min_collect:
                        time.sleep(5)

                for id in ids:
                    self.api.like(id)

                liked += len(ids)
                self.api.logger.info(f"いいね数: {liked}")
                ids.clear

            except Exception as e:
                self.api.logger.warning(str(e))
                self.api.logger.info("休憩中...☕")
                ids.clear
                time.sleep(300)

        self.api.logger.info(f"合計{liked}個の投稿にいいねしました。")


if __name__ == "__main__":
    email = "メールアドレス"
    password = "パスワード"
    secret_token = None

    bot = LikeBot(email, password, secret_token)
    bot.run()
