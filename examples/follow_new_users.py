# 新規ユーザーをフォローするサンプルコード


email = "your_email"
password = "your_password"


import yaylib


class FollowBot:
    def __init__(self, email=None, password=None):
        self.api = yaylib.Client()
        self.api.login(email, password)

    def main(self):
        followed = 0

        while True:
            ids = []

            print("新規ユーザーを取得しています...")

            new_users = self.api.search_users(recently_created=True, number=100)

            ids = [user.id for user in new_users.users if not user.is_following]

            print(f"取得済みユーザー数: {len(ids)}")

            for id in ids:
                try:
                    self.api.follow_user(id)
                    print("ユーザーをフォローしました。")
                except yaylib.ForbiddenError:
                    print("ユーザーのフォローに失敗しました。")

            followed += len(ids)
            print(f"フォロー済み: {followed}")


if __name__ == "__main__":
    bot = FollowBot(email, password)
    bot.main()
