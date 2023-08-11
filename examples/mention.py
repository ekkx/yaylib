import yaylib


api = yaylib.Client()

api.login(email="your_email", password="your_password")

api.create_post(
    # IDが 123 のユーザーをメンションする場合
    text=f"{api.mention(123)}さん、お元気ですか？"
)
