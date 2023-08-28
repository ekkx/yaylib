# 新規ユーザーをフォローするサンプルコード


email = "your_email"
password = "your_password"


import yaylib


class FollowBot:
    def __init__(self, email=None, password=None):
        self.api = yaylib.Client()
        self.api.login(email, password)

    def run(self):
        followed = 0

        while True:
            ids = []

            self.api.logger.info("新規ユーザーを取得しています...")

            new_users = self.api.search_users(recently_created=True, number=100)

            ids = [user.id for user in new_users.users if not user.is_following]

            self.api.logger.info(f"取得済みユーザー数: {len(ids)}")

            for id in ids:
                try:
                    self.api.follow_user(id)
                except yaylib.errors.ForbiddenError:
                    self.api.logger.error("ユーザーのフォローに失敗しました。")

            followed += len(ids)
            self.api.logger.info(f"フォロー済み: {followed}")


if __name__ == "__main__":
    bot = FollowBot(email, password)
    bot.run()
