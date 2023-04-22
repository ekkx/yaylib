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
        min_collect = 30
        liked = 0

        while liked < amount:

            ids = set()

            try:
                while len(ids) < min_collect:
                    posts = self.yay.get_timeline()
                    for post in posts:
                        if post.liked == False:
                            ids.add(post.id)

                    print(f'Collected {len(ids)} posts.')

                    if len(ids) < min_collect:
                        time.sleep(15)

                for post_id in ids:
                    if liked >= amount:
                        break

                    print(self.yay.like_post(post_id))
                    liked += 1

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
