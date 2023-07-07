# yaylib へようこそ！

<image src="https://github.com/qvco/yaylib/assets/77382767/96f5795d-1ba5-4ec6-b00c-ffd245968dc3" width="550px">

ここでは「yaylib」のインストール方法から基本的な使い方を紹介します。

「yaylib」の主な機能は、以下の通りです。

- **投稿、スレッド、グループの作成**
- **指定したユーザー、投稿、グループの詳細情報を入手**
- **投稿、ユーザーの検索**
- **フォロー、いいね、リプ、age↑、ブロック**
- **指定したユーザーの投稿、レター、サークル、フォロワーの取得**

などなど、すべて合わせると、**200 を超える機能**があります。

ここでは紹介しきれないので、より詳細を知りたい方は、<a href="https://github.com/qvco/yaylib/blob/main/docs/README.md">ドキュメント</a>をご覧ください。

<details open>
    <summary>もくじ</summary>
    <ol>
        <li><a href="前提条件">前提条件</a></li>
        <li>
            <a href=#インストール>インストール</a>
            <ul>
                <li><a href="#python-をインストール">Python をインストール</a></li>
                <li><a href="#yaylib-をインストール">yaylib をインストール</a></li>
            </ul>
        </li>
        <li>
            <a href="コードを書く準備">コードを書く準備</a>
            <ul>
                <li><a href="#本記事のゴール">目標</a></li>
                <li><a href="#コードエディター">コードエディター</a></li>
            </ul>
        </li>
        <li>
            <a href="#実際にコードを書く">実際にコードを書く</a>
            <ul>
                <li><a href="#タイムラインの取得">タイムラインの取得</a></li>
                <li><a href="#投稿する方法">投稿する方法</a></li>
                <li><a href="#さらに詳しい機能ドキュメント">さらに詳しい機能&ドキュメント</a></li>
            </ul>
        </li>
        <li><a href="コンタクト">コンタクト</a></li>
    </ol>
</details>

## 前提条件

Yay!アカウントを持っている前提で、記載しています。
まだの方は、アカウント登録を済ませておいて下さい。

また、このライブラリを使用する前に、以下の点に注意してください。

- **絶対に過度なリクエストは送信しない**

  膨大なデータの取得、いいね爆やレター爆の過度な実行はサーバーに負担をかけます。
  ユーザーが気持ちよく利用できるよう最低限の節度を守って使用しましょう。

- **何があっても自己責任**

  運営の手によってアカウントが BAN されたとしてもそれは自己責任です。

## インストール

### Python をインストール

yaylib を使用するには、Python 3.11 かそれ以上のバージョンが必要になります。
Python 3.11 のインストール方法については、以下の記事を参考にしてください。

<a href=https://www.javadrive.jp/python/install/index1.html>
記事: Pythonのダウンロードとインストール | Let'sプログラミング
</a>

### yaylib をインストール

「yaylib」をダウンロード・インストールするのは超簡単。

コマンドプロンプトを起動します。

コマンドプロンプトに「pip install yaylib」と入力するだけです。最後にエンターキーをおします。

```bash
pip install yaylib
```

<image src="https://github.com/qvco/yaylib/assets/77382767/63c75259-b7de-47ba-a115-d1f14f34864a" alt="コマンドプロンプト<pip install yaylibの画像>">

すぐに「yaylib」のダウンロードがはじまります。そして続けて、インストールが始まります。

<image src="https://github.com/qvco/yaylib/assets/77382767/2751de59-96ac-4b89-9113-48cbbf8bf9d7" alt="コマンドプロンプト<yaylibインストール完了画像>">

「**Successfully installed yaylib**」と表示されていれば、「yaylib」のインストールは完了です。

## コードを書く準備

#### 【本記事のゴール】

- yaylib を使用して、
- 投稿を検索したり、
- Yay!に投稿する（文字だけじゃなくて画像投稿も可）
- それ以外の機能も紹介しています。

### コードエディター

コードエディターはなんでも構いませんが、ここでは Visual Studio Code（以下、VS Code という）を使用します。
VS Code は、多くのプログラミング言語に対応しているテキストエディターです。

<a href="https://code.visualstudio.com/download">
    Visual Studio Code のダウンロードはこちらから
</a>

## 実際にコードを書く

場所はどこでも構いませんので、プロジェクトのフォルダーを作成します。

フォルダー名は「my_bot」にでもしておきましょう。

<image src="https://github.com/qvco/yaylib/assets/77382767/56fa9a2d-b9d4-446b-b937-d0f8d22a99c4" alt="フォルダ作成のgif">

次に、作成したフォルダーを VS Code で開き、Python ファイルを作成します。

ファイル名はなんでも構いませんが拡張子を必ず「.py」にしましょう。

<image src="https://github.com/qvco/yaylib/assets/77382767/a798d90b-2ae4-428f-8653-ad420082edde" alt="bot.pyの作成gif" width="800px">

### タイムラインの取得

投稿を検索して、取得する方法を記します。

投稿を検索するにあたり、下記条件が設定できます。

（よく使用するものを、ピックアップしてみました。）

- キーワード
- ハッシュタグ
- 取得件数

具体的な使い方は、下記コードをご確認ください。（注釈を入れました。）

```python
import yaylib


bot = yaylib.Client()

bot.login(email="メールアドレス", password="パスワード")


# タイムラインを取得
timeline = bot.get_timeline(
    noreply_mode=True, # 返信も取得するか
    number=10 # 取得件数
)
# 代表的な取得情報を表示
for post in timeline.posts:
    print(post.user.nickname) # 投稿者の名前
    print(post.text) # 投稿本文
    print(post.likes_count) # いいね数
    print(post.reposts_count) # (´∀｀∩)↑age↑の数


# キーワードでタイムラインを取得
timeline = bot.get_timeline_by_keyword(
    keyword="おはようございます", # 検索キーワード
    number=10 # 取得件数
)
# 代表的な取得情報を表示
for post in timeline.posts:
    print(post.user.nickname) # 投稿者の名前
    print(post.text) # 投稿本文
    print(post.likes_count) # いいね数
    print(post.reposts_count) # (´∀｀∩)↑age↑の数


# ハッシュタグでタイムラインを取得
timeline = bot.get_timeline_by_hashtag(
    hashtag="いいねでレター", # 検索ハッシュタグ
    number=10 # 取得件数
)
# 代表的な取得情報を表示
for post in timeline.posts:
    print(post.user.nickname) # 投稿者の名前
    print(post.text) # 投稿本文
    print(post.likes_count) # いいね数
    print(post.reposts_count) # (´∀｀∩)↑age↑の数
```

### 投稿する方法

#### 文字のみを投稿する場合

```python
import yaylib

# 認証情報
email = "メールアドレス"
password = "パスワード"

bot = yaylib.Client()

# ログインする
bot.login(email, password)

# 投稿する
bot.create_post(text="投稿したい内容")
```

#### 文字に加え、画像も投稿する場合

ローカルに保存している画像ファイルを、送信するケースになります。

先ほどのコード最終行を、以下のように変更します。

```python
# test.pngファイルをサーバーにアップロードする
filename = bot.upload_image(
    image_type="post", # 画像の使い道を「投稿」に指定
    image_path="./test.png"
)

# アップロードした画像の名前を指定して投稿する
bot.create_post(text="投稿したい内容", attachment_filename=filename)
```

### さらに詳しい機能&ドキュメント

ここまでで紹介した機能以外にももっとたくさんの便利な機能があります。

詳しい情報は<a href="https://github.com/qvco/yaylib/blob/main/docs/README.md">ドキュメントを参照</a>するか、コードエディターのドット入力補完を使用してください。

<image src="https://github.com/qvco/yaylib/assets/77382767/666e6eb9-bfeb-4b4c-8f6c-033555436fc3" alt="ドット補完の画像">

<!-- さらに詳しい機能の種類や詳細についてはこちらを確認してください。 -->

## コンタクト

ここまでの説明がわからなかったり、バグを発見したり、もし一緒に開発したいと思ってくださる方がいましたら、ぜひ 「yaylib」 の <a href="https://discord.gg/MEuBfNtqRN">Discord サーバー</a>に参加して声をかけてください。
