<div><a id="readme-top"></a></div>
<p align="center">
    <img src=https://img.shields.io/github/stars/qvco/yaylib?style=for-the-badge&logo=appveyor&color=blue />
    <img src=https://img.shields.io/github/forks/qvco/yaylib?style=for-the-badge&logo=appveyor&color=blue />
    <img src=https://img.shields.io/github/issues/qvco/yaylib?style=for-the-badge&logo=appveyor&color=informational />
    <img src=https://img.shields.io/github/issues-pr/qvco/yaylib?style=for-the-badge&logo=appveyor&color=informational />
</p>
<br />
<p align="center">
    <a href="https://github.com/othneildrew/Best-README-Template">
        <img src="https://github.com/qvco/yaylib/assets/77382767/6e72ec90-b8e9-40bf-a7ad-34fb2ccea0f9" alt="Logo" height="300px">
    </a>
    <h3 align="center">yaylib</h3>
    <p align="center">
        「<strong>yaylib</strong>」は同世代でつながるチャットアプリ、Yay!（イェイ）の API クライアントです。<br />
        このライブラリを使用することで、あらゆる操作の自動化や、ボットの開発が可能です。
        <br />
        <br />
        <a href="https://github.com/qvco/yaylib/blob/main/docs/README.md">
            <strong>詳しい機能の詳細や使い方はこちらから »</strong>
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
    <li><a href="#buy-me-a-coffee">Buy me a coffee</a></li>
    <li><a href="#インストール">インストール</a></li>
    <li><a href="#使用例">使用例</a></li>
    <li><a href="#yaylib-で誕生したロボットたち">yaylib で誕生したロボットたち</a></li>
    <li><a href="#共同開発について">共同開発について</a></li>
    <li><a href="#免責事項">免責事項</a></li>
    <li><a href="#利用許諾">利用許諾</a></li>
  </ol>
</details>

<!-- Buy me a coffee -->

## Buy me a coffee

もしこのライブラリが気に入っていただけたら  
私たちに ↓ コーヒー ↓ をお恵みくださいませ！！❤

<a href="https://www.buymeacoffee.com/qvco" target="_blank"><img src="https://www.buymeacoffee.com/assets/img/custom_images/orange_img.png" alt="Buy Me A Coffee" style="height: auto !important;width: auto !important;" ></a>

<!-- インストール -->

## インストール

**※ Python 3.11 かそれ以上のバージョンが必要です。**

ライブラリをインストールするには、以下のコマンドを実行します:

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

「yaylib」の始め方や機能&特徴については、[こちら](https://github.com/qvco/yaylib/blob/main/docs/TUTORIAL.md) を参照してください。

<!-- 使用例 -->

## 使用例

メールアドレスとパスワードを用いてログイン後、新しく投稿を作成するコードです。

```python
import yaylib

api = yaylib.Client()

api.login(email="メールアドレス", password="パスワード")

api.create_post(text="初めての投稿！", color=2)
```

より詳細な使用例については、[こちら](https://github.com/qvco/yaylib/blob/main/examples) を参照してください。

<p align="right">(<a href="#readme-top">トップに戻る</a>)</p>

<!-- yaylib で誕生したボットの一覧 -->

## yaylib で誕生したロボットたち

yaylib を用いて開発したロボットがある場合は、ぜひ教えてください！

<table align="center">
    <thead>
        <tr>
            <th><a href="https://yay.space/user/5855987">MindReader AI</a></th>
            <th><a href="https://yay.space/user/0">Funktion (架空)</a></th>
            <th><a href="https://yay.space/user/0">香ばしいボット (架空)</a></th>
        </tr>
    </thead>
    <tbody>
        <tr>
            <td align="center">
                <img src="https://github.com/qvco/yaylib/assets/77382767/472febe4-4c5f-490c-8417-de0d5dbbbc72" width="200px">
                <br />
                <p>開発者: <a href="https://yay.space/user/35152">毛の可能性</a></p>
            </td>
            <td align="center">
                <img src="https://github.com/qvco/yaylib/assets/77382767/ff207016-21bf-4e76-b0e0-f70ebc4a121f" width="200px">
                <br />
                <p>開発者: <a href="https://yay.space/user/0">ぺゅー</a></p>
            </td>
            <td align="center">
                <img src="https://github.com/qvco/yaylib/assets/77382767/2324e518-b2c8-43cd-95e5-90ee2383aec1" width="200px">
                <br />
                <p>開発者: <a href="https://yay.space/user/0">めんぶれ天然水。</a></p>
            </td>
        </tr>
    </tbody>
</table>

<!-- 共同開発について -->

## 共同開発について

私たちと一緒に開発することに興味を持っていただけているなら大歓迎です。

- <a href="https://github.com/qvco/yaylib/pulls">プルリクエストを送信する</a>
- nikola.desuga@gmail.com にメールを送信する
- <a href="https://discord.gg/MEuBfNtqRN">Discord サーバーに参加する</a>

のいずれかの方法でコンタクトしてください！

<!-- 免責事項 -->

## 免責事項

yaylib は、API の公式なサポートやメンテナンスを提供するものではありません。このクライアントを使用する場合、利用者は**リスクや責任を自己負担**できるものとします。このクライアントによって提供される情報やデータの正確性、信頼性、完全性、適時性について、いかなる保証も行いません。また、このクライアントの使用によって生じた損害や不利益について、一切の責任を負いかねます。利用者は自己の責任において、このクライアントを使用し、API にアクセスするものとします。なお、この免責事項は予告なく変更される場合があります。

<!-- 利用許諾 -->

## 利用許諾

フルライセンスは [こちら](https://github.com/qvco/yaylib/blob/main/LICENSE) からご確認いただけます。  
このプロジェクトは、 **【MIT ライセンス】** の条件の下でライセンスされています。

<p align="right">(<a href="#readme-top">トップに戻る</a>)</p>
