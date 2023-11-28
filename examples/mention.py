# 投稿でユーザーをメンションするサンプルコード
# 例): ユーザーIDが 15184 のユーザーを「アルパカ」という名前でメンションする場合


email = "your_email"
password = "your_password"


import yaylib
from yaylib import mention

api = yaylib.Client()
api.login(email, password)

api.create_post(f"こんにちは、{mention(user_id=15184, display_name='アルパカ')}さん！")
