<div><a id="readme-top"></a></div>
<div align="center">
    <img src="https://img.shields.io/github/stars/qvco/yaylib?style=for-the-badge&logo=appveyor&color=blue" />
    <img src="https://img.shields.io/github/forks/qvco/yaylib?style=for-the-badge&logo=appveyor&color=blue" />
    <img src="https://img.shields.io/github/issues/qvco/yaylib?style=for-the-badge&logo=appveyor&color=informational" />
    <img src="https://img.shields.io/github/issues-pr/qvco/yaylib?style=for-the-badge&logo=appveyor&color=informational" />
</div>
<br />
<p align="center">
    <a href="https://github.com/othneildrew/Best-README-Template">
        <img src="https://github.com/qvco/yaylib/assets/77382767/45c45b21-d812-4cad-8f27-315ffef53201" alt="Logo" height="300px">
    </a>
    <h3 align="center">yaylib</h3>
    <p align="center">
        「<strong>yaylib</strong>」は同世代でつながるチャットアプリ、Yay!（イェイ）の API ラッパーです。<br />
        あらゆる操作の自動化や、ボットの開発が可能です。
        <br />
        <br />
        <a href="https://github.com/qvco/yaylib/blob/master/docs/README.md">
            <strong>ドキュメントはこちらから »</strong>
        </a>
        <br />
        <br />
        <a href="https://github.com/qvco/yaylib/issues">Report Bug</a>
        ·
        <a href="https://github.com/qvco/yaylib/issues">Request Feature</a>
        ·
        <a href="https://discord.gg/MEuBfNtqRN">Join the discord</a>
    </p>
</p>

<!-- TABLE OF CONTENTS -->

<details>
  <summary>Table of Contents</summary>
  <ol>
    <li><a href="#インストール">インストール</a></li>
    <li><a href="#使用例">使用例</a></li>
    <li><a href="#yaylib-で誕生したロボットたち">yaylib で誕生したロボットたち</a></li>
    <li><a href="#buy-me-a-coffee">Buy me a coffee</a></li>
    <li><a href="#共同開発について">共同開発について</a></li>
    <li><a href="#免責事項">免責事項</a></li>
    <li><a href="#利用許諾">利用許諾</a></li>
  </ol>
</details>

<!-- インストール -->

## インストール

**※ Python 3.11 かそれ以上のバージョンが必要です。**

「yaylib」をインストールするには、以下のコマンドを実行します:

```bash
pip install yaylib
```

開発バージョンをインストールするには、以下の手順を実行します:

```bash
git clone https://github.com/qvco/yaylib

cd yaylib

pip install -r requirements.txt

pip install -e .
```

「yaylib」の始め方については、[こちら](https://github.com/qvco/yaylib/blob/master/docs/TUTORIAL.md) を確認してください。

<!-- 使用例 -->

## 使用例

メールアドレスとパスワードを使用してログインしたあと、タイムラインをキーワードで検索して「いいね」するコードです。

```python
import yaylib


api = yaylib.Client()
api.login(email="メールアドレス", password="パスワード")

timeline = api.get_timeline_by_keyword(
    keyword="プログラミング",
    number=15
)

for post in timeline.posts:
    response = api.like_post(post.id)
    print(post.id, response.data) # 実行結果を出力
```

より詳しい使用例については、[こちら](https://github.com/qvco/yaylib/blob/master/examples) を参照してください。

<p align="right">(<a href="#readme-top">トップに戻る</a>)</p>

<!-- yaylib で誕生したボットの一覧 -->

## yaylib で誕生したロボットたち

「yaylib」を用いて開発したロボットがある場合は、ぜひ教えてください！

<table align="center">
    <thead>
        <tr>
            <th><a href="https://yay.space/user/5855987">MindReader AI</a></th>
            <th><a href="https://yay.space/user/7293290">香ばしいボット</a></th>
            <th><a href="https://yay.space/user/7406336">GIGAZINE</a></th>
        </tr>
    </thead>
    <tbody>
        <tr>
            <td align="center">
                <img src="https://github.com/qvco/yaylib/assets/77382767/cc41ce3c-0e11-4ec5-be99-ff7090a95667" width="200px">
                <br />
                <p>開発者: <a href="https://yay.space/user/35152">毛の可能性</a></p>
            </td>
            <td align="center">
                <img src="https://github.com/qvco/yaylib/assets/77382767/cbffdc25-7873-4242-b065-e6a686bade54" width="200px">
                <br />
                <p>開発者: <a href="https://yay.space/user/93923">めんぶれ天然水。</a></p>
            </td>
            <td align="center">
                <img src="https://github.com/qvco/yaylib/assets/77382767/65fcb885-4fbe-4170-9378-6f8d9af61ff8" width="200px">
                <br />
                <p>開発者: <a href="https://yay.space/user/1298298">ぺゅー</a></p>
            </td>
        </tr>
    </tbody>
</table>

<!-- Buy me a coffee -->

## Buy me a coffee

このライブラリが気に入っていただけたら、<a href="https://github.com/qvco/yaylib/">スターをお願いします</a> ⭐️  
また、Buy Me a Coffee からご支援いただけますと幸いです。

<a href="https://www.buymeacoffee.com/qvco" target="_blank"><img src="https://www.buymeacoffee.com/assets/img/custom_images/orange_img.png" alt="Buy Me A Coffee" style="height: auto !important;width: auto !important;" ></a>

<!-- 共同開発について -->

## 共同開発について

私たちと開発することに興味を持っていただけているなら、ぜひ参加してください！

- <a href="https://github.com/qvco/yaylib/pulls">プルリクエストを送信する</a>
- <a href="mailto:nikola.desuga@gmail.com">nikola.desuga@gmail.com</a> にメールを送信する
- <a href="https://discord.gg/MEuBfNtqRN">Discord サーバーに参加する</a>

のいずれかの方法でコンタクトしてください。詳しくは[こちらから](https://github.com/qvco/yaylib/blob/master/CONTRIBUTING.md)！

<!-- 免責事項 -->

## 免責事項

yaylib は、API の公式なサポートやメンテナンスを提供するものではありません。このクライアントを使用する場合、利用者は**リスクや責任を自己負担**できるものとします。このクライアントによって提供される情報やデータの正確性、信頼性、完全性、適時性について、いかなる保証も行いません。また、このクライアントの使用によって生じた損害や不利益について、一切の責任を負いかねます。利用者は自己の責任において、このクライアントを使用し、API にアクセスするものとします。なお、この免責事項は予告なく変更される場合があります。

<!-- 利用許諾 -->

## 利用許諾

フルライセンスは [こちら](https://github.com/qvco/yaylib/blob/master/LICENSE) からご確認いただけます。  
このプロジェクトは、 **【MIT ライセンス】** の条件の下でライセンスされています。

<p align="right">(<a href="#readme-top">トップに戻る</a>)</p>
