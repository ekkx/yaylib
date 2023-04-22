import re
import os
import time

from dotenv import load_dotenv
from mirai_translate import Client  # 翻訳ライブラリ
from yaybot import Yay


class TranslatorBot(object):

    def __init__(self, token=None):
        self.token = token
        self.yay = Yay(token)
        self.client = Client()

    def login(self, email, password):
        if email is None or password is None:
            return
        self.yay.login(email, password)

    def get_tagged(self, now):
        notices = self.yay.get_notification(important=True)
        ret = []
        for notice in notices:
            # 通知の種類がメンションで5分以内のものを取得
            if notice.type == 'tagged' and (now - notice.created_at) < 60 * 5:
                ret.append(notice)
        return ret

    def start(self, timeout=30):

        ids = set()
        langs = ['en', 'ja']

        while True:
            try:
                now = int(time.time())
                notices = self.get_tagged(now)
                for notice in notices:
                    post_id = notice.from_post_id

                    if post_id in ids:
                        continue

                    ids.add(post_id)

                    post = self.yay.get_post(post_id)
                    raw_text = post.text
                    text = ' '.join(raw_text.split()[3:])
                    matches = re.findall(r'\b\w{2}\b', raw_text)

                    for match in matches:
                        if match not in langs:
                            continue

                    translated = self.client.translate(
                        text, matches[0], matches[1])

                    self.yay.create_reply(translated, post.id)
                    print(f'Reply Successful: {translated}')

                time.sleep(timeout)

            except Exception as e:
                print(str(e))


if __name__ == '__main__':

    # .envファイルにメールアドレスとパスワードを設定してください。
    load_dotenv()
    EMAIL = os.getenv('EMAIL')
    PASSWORD = os.getenv('PASSWORD')

    bot = TranslatorBot()
    bot.login(EMAIL, PASSWORD)  # ログイン
    bot.start()  # timeout引数は再度通知を取得するまでの待ち時間
