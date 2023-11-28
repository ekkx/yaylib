# 投稿でユーザーをメンションするサンプルコード
# 例): ユーザーIDが 12345 のユーザーをメンションする場合


email = "your_email"
password = "your_password"


import yaylib
from yaylib import mention

api = yaylib.Client()

api.login(email, password)

api.create_post(f"{mention(12345, 'アルパカ')}さん、お元気ですか？")
