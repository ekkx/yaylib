# メールアドレスとパスワードを用いてログインするサンプルコード
# ログイン出来たことを確認するため、「Hello with yaylib!」と投稿します。


email = "your_email"
password = "your_password"


import yaylib

client = yaylib.Client()

client.login(email, password)

client.create_post("Hello with yaylib!")
