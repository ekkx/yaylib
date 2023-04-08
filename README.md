# YayBot

SNS「Yay!」の非公式 API です。🚀  
使用言語は Python です。

https://yay.space/

### Important

まだ開発途中であり、今後も機能の追加を予定中です。  
機能の修正や追加要望等あれば、以下から連絡してください。

---

### [Yay!アカウント](https://yay.space/user/3851771) | [メールアドレス](nikola.desuga@gmail.com)

---

![Python 3.11](https://img.shields.io/badge/python-3.11-blue.svg)

### 使い方

```python
from api import Yay


yay = Yay()

yay.login(email='メールアドレス', password='パスワード') # ログイン


user_data = yay.get_user('123')  # IDが'123'のユーザーのプロフィールを取得
print(user_data.display_name)  # 取得したユーザーの名前を表示
print(user_data.biography)  # 取得したユーザーの自己紹介を表示


post_data = yay.get_post(post_id='456')  # IDが'456'の投稿を取得
print(post_data.author_display_name)  # 投稿者の名前を表示
print(post_data.text)  # 投稿本文を表示
yay.like_post(post_id='456')  # 投稿をいいねする


followers = yay.get_user_followers('123')  # IDが'123'のユーザーのフォロワーを取得する
for follower in followers:
    user_id = follower.id
    yay.follow_user(user_id)  # 取得したユーザーをフォローする

```

# 使用条件

- この API をマーケティング目的（スパム、ボット、嫌がらせ、大量の一斉送信など）に使用しないこと。

# 注意

このコードは、「Yay!」、株式会社ナナメウエ、またはその関連会社といかなる関係も持っていません。これは独立した非公式 API です。自己責任で使用してください。

This code is in no way affiliated with, authorized, maintained, sponsored or endorsed by Yay!, Nanameue inc. or any of its affiliates or subsidiaries. This is an independent and unofficial API. Use it at your own risk.
