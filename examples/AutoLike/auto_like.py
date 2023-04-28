import os
import time

from dotenv import load_dotenv
from yaybot import Yay


class LikeBot(object):

    def __init__(self, token=None):
        self.token = token
        self.yay = Yay(token)

    def login(self, email, password):
        if email is None or password is None:
            return
        self.yay.login(email, password)

    def start(self, amount=None):
        amount = float('inf') if amount is None else amount
        min_collect = 50
        liked = 0

        while liked < amount:

            ids = []

            try:
                while len(ids) < min_collect:
                    posts = self.yay.get_timeline()
                    ids = [post.id for post in posts if not post.liked]

                    print(f'Collected {len(ids)} posts.')

                    if len(ids) < min_collect:
                        time.sleep(30)

                while len(ids) > 0:
                    sliced_ids = ids[:25]
                    del ids[:25]

                    liked += len(sliced_ids)
                    print(f'{liked}: {self.yay.like_posts(sliced_ids)}')

                ids.clear

            except Exception as e:
                print(str(e))
                ids.clear
                time.sleep(300)

        print(f'Liked {liked} posts in total.')


if __name__ == '__main__':

    # .envファイルにメールアドレスとパスワードを設定してください。
    load_dotenv()
    EMAIL = os.getenv('EMAIL')
    PASSWORD = os.getenv('PASSWORD')

    bot = LikeBot()

    bot.login(EMAIL, PASSWORD)  # ログイン
    bot.start()  # amount引数に回数を指定しない場合、永久的に実行し続けます。
