# 投稿でユーザーをメンションするサンプルコード
# 例): ユーザーIDが 12345 のユーザーをメンションする場合


email = "your_email"
password = "your_password"


import yaylib

api = yaylib.Client()

api.login(email, password)

api.create_post(f"{api.mention(12345)}さん、お元気ですか？")
