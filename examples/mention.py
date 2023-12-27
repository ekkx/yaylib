# 投稿でユーザーをメンションするサンプルコード
# 例): ユーザーIDが 15184 のユーザーを「アルパカ」という名前でメンションする場合


email = "your_email"
password = "your_password"


import yaylib
from yaylib import mention

client = yaylib.Client()
client.login(email, password)

client.create_post(f"{mention(12345, 'アルパカ')}さん、お元気ですか？")
