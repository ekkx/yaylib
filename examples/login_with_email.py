# メールアドレスとパスワードを用いてログインするサンプルコード
# ログイン出来たことを確認するため、「Hello with yaylib!」と投稿します。


email = "your_email"
password = "your_password"


import yaylib

api = yaylib.Client()

api.login(email, password)

api.create_post("Hello with yaylib!")
