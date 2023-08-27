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
        collected_ids = set()

        while True:
            ids = []

            self.api.logger.info("新規ユーザーを取得しています...")

            new_users = self.api.search_users(recently_created=True, number=100)

            new_ids = [
                user.id
                for user in new_users.users
                if not user.is_following and user.id not in collected_ids
            ]

            ids.extend(new_ids)
            collected_ids.update(new_ids)

            self.api.logger.info(f"取得済みユーザー数: {len(ids)}")

            for id in ids:
                self.api.follow_user(id)

            followed += len(ids)
            self.api.logger.info(f"フォロー済み: {followed}")

            ids.clear()
            collected_ids.clear()


if __name__ == "__main__":
    bot = FollowBot(email, password)
    bot.run()
