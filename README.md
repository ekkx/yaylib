<!-- <p align="center">
<img src=https://img.shields.io/github/stars/qualia-5w4/yaybot?style=for-the-badge&logo=appveyor&color=blue />
<img src=https://img.shields.io/github/forks/qualia-5w4/yaybot?style=for-the-badge&logo=appveyor&color=blue />
<img src=https://img.shields.io/github/issues/qualia-5w4/yaybot?style=for-the-badge&logo=appveyor&color=informational />
<img src=https://img.shields.io/github/issues-pr/qualia-5w4/yaybot?style=for-the-badge&logo=appveyor&color=informational />
</p>
<br /> -->
<p align="center">
  <a href="https://yay.space">
    <img src="https://yay.space/images/home-group-categories-background-3.jpg" alt="Logo">
    <!-- <img src="https://yay.space/images/announcement-banner-background.svg" alt="Logo"> -->
    <!-- <img src="https://yay.space/images/app-logo-3.svg" alt="Logo" width="150px" height="150px"> -->
  </a>
  
  <p align="center">
    <img src="https://img.shields.io/github/release/qualia-5w4/yaybot">
    <img src="https://img.shields.io/badge/python-3.11-blue.svg">
    <img src="https://img.shields.io/badge/License-MIT-blue.svg">
  </p>

  <h3 align="center">Yay! - 非公式ライブラリ</h3>

  <p align="center">
    汎用的な機能を備えた、SNS「Yay!」の非公式APIです。<br />
    とてもシンプルな記述で実装できます。
    <br />
    <a href="https://github.com/qualia-5w4/yaybot/tree/master/examples"><strong>詳細はドキュメントにて »</strong></a>
    <br />
    <br />
    <a href="https://github.com/qualia-5w4/yaybot/issues">バグを報告する</a>
    ·
    <a href="https://yay.space/user/3851771">Yay!アカウント</a>
      ·
    <a href="mailto:nikola.desuga@gmail.com">メールアドレス</a>
  </p>
</p>

<details>
  <summary>- - - 目次 - - -</summary>
  <ol>
    <li>
      <a href="#-使用方法">使用方法</a>
      <ul>
        <li><a href="#-ダウンロード">ダウンロード</a></li>
        <li><a href="#%EF%B8%8F-始め方">始め方</a></li>
      </ul>
    </li>
    <li><a href="#trophy-機能--特徴">機能 & 特徴</a></li>
    <li><a href="#scroll-ライセンス--免責事項">ライセンス & 免責事項</a></li>
  </ol>
</details>

## ☕ 使用方法

### 💻 ダウンロード

次のコマンドをターミナル上で実行してください:

```bash
pip install git+https://github.com/qualia-5w4/yaybot
```

もしインストールが正常に終了しない場合は、次のコマンドを実行してください:

```bash
git clone https://github.com/qualia-5w4/yaybot

cd yaybot

pip install -r requirements.txt

pip install .
```

### 🖥️ 始め方

1. `from yaybot import Yay` でインポートします。
2. ログインは任意です。※短時間に何度もログインすると制限される場合があります。
3. アクセストークンを使用する場合は、`yay = Yay(token='トークン')` と記述してください。

```python
from yaybot import Yay


yay = Yay()

yay.login(email='メールアドレス', password='パスワード') # ログイン

# print(yay.access_token) ※アクセストークンを取得する場合

user = yay.get_user('123')  # IDが'123'のユーザーのプロフィールを取得
print(user.screen_name)  # 取得したユーザーの名前を表示
print(user.bio)  # 取得したユーザーの自己紹介を表示


post = yay.get_post('456')  # IDが'456'の投稿を取得
print(post.author_screen_name)  # 投稿者の名前を表示
print(post.text)  # 投稿本文を表示
yay.like_post(post.id)  # 投稿をいいねする


followers = yay.get_user_followers('123')  # IDが'123'のユーザーのフォロワーを取得する
for follower in followers:
    yay.follow_user(follower.id)  # 取得したユーザーをフォローする

```

---

## :trophy: 機能 & 特徴

- シンプルな記述で実装可能
- 汎用的な機能
- プロキシをサポート: http/s

---

## :scroll: ライセンス & 免責事項

このコードは株式会社ナナメウエ、またはその関連会社といかなる関係も持っていません。これは独立した非公式 API です。**自己責任で使用してください。**

This code is in no way affiliated with, authorized, maintained, sponsored or endorsed by Nanameue inc. or any of its affiliates or subsidiaries. This is an independent and unofficial API. **USE IT AT YOUR OWN RISK.**

Licensed under the [MIT License](LICENSE)
