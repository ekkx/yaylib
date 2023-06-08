## yaylib へようこそ！

ここでは yaylib のインストール方法から基本的な使い方を紹介します。

<details open>
    <summary>目次</summary>
    <ul>
        <li><a href="yaylibで出来ること">yaylib で出来ること</a></li>
        <li>
            <a href=#インストール>インストール</a>
            <ul>
                <li><a href="#Pythonをインストール">Python をインストール</a></li>
                <li><a href="#yaylibをインストール">yaylib をインストール</a></li>
            </ul>
        </li>
        <li>
            <a href="コードを書いてみる">コードを書いてみる</a>
            <ul>
                <li><a href="#コードエディタ">コードエディタ</a></li>
                <li><a href="#実際にコードを書く">実際にコードを書く</a></li>
                <li><a href="#タイムラインの取得">タイムラインの取得</a></li>
                <li><a href="#投稿する方法">投稿する方法</a></li>
                <li><a href="#さらに詳しい機能">さらに詳しい機能</a></li>
            </ul>
        </li>
        <li><a href="コンタクト">コンタクト</a></li>
    </ul>
</details>

## yaylib で出来ること

主な機能は、以下の通りです。

- **投稿、リプライ、スレッド、グループの作成**
- **指定したユーザー、投稿、グループの詳細情報を入手**
- **投稿、ユーザーの検索**
- **フォロー、いいね、age↑、ブロック**
- **指定したユーザーの投稿、レター、サークル、フォロワーの取得**

などです。すべて合わせると、**200 を超える機能**があります。
より詳細を知りたい方は、<a>ドキュメント</a>をご覧ください。

### 前提条件

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

<a href=#https://www.javadrive.jp/python/install/index1.html>
記事: Pythonのダウンロードとインストール | Let'sプログラミング
</a>

### yaylib をインストール

「yaylib」をダウンロード・インストールするのは超簡単です。

コマンドプロンプトを起動します。

コマンドプロンプトに「pip install yaylib」と入力するだけです。最後にエンターキーをおします。

```bash
pip install yaylib
```

<image src="https://github.com/qvco/yaylib/assets/77382767/63c75259-b7de-47ba-a115-d1f14f34864a" alt="コマンドプロンプト<pip install yaylibの画像>">

すぐに「yaylib」のダウンロードがはじまります。そして続けて、インストールが始まります。

<image src="https://github.com/qvco/yaylib/assets/77382767/2751de59-96ac-4b89-9113-48cbbf8bf9d7" alt="コマンドプロンプト<yaylibインストール完了画像>">

「**Successfully installed yaylib**」と表示されていれば、「yaylib」のインストールは完了です。

## コードを書いてみる

<h4>【本記事のゴール】</h4>

- yaylib を使用して、
- 投稿を検索したり、
- Yay!に、投稿する。（文字だけじゃなくて、画像投稿も可）
- それ以外にも便利な機能を紹介しています。

### コードエディタ

コードエディタはなんでも構いませんが、ここでは Visual Studio Code を使用します。
Visual Studio Code は、多くのプログラミング言語に対応しているテキストエディタです。

<a href="https://code.visualstudio.com/download">
Visual Studio Code のダウンロードはこちらから
</a>

### 実際にコードを書く

場所はどこでも構いませんので、プロジェクトのフォルダーを作成します。
フォルダー名は「my_bot」にでもしておきましょう。

<image alt="フォルダ作成のgif">

次に、作成したフォルダーを Visual Studio Code で開き、Python ファイルを作成します。
ファイル名はなんでも構いません。

<image alt="bot.pyの作成gif">

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
# test.pngファイルをアップロードする
uploaded_filename = bot.upload_image(
    image_type="post", # 画像の使い道を「投稿」に指定
    image_path="./test.png"
)

# アップロードした画像の名前を指定して投稿する
bot.create_post(text="投稿したい内容", attachment_filename=uploaded_filename)
```

### さらに詳しい機能

ここまでで紹介した機能以外にももっとたくさんの便利な機能があります。
さらに詳しい機能の種類や詳細についてはこちらを確認してください。

## コンタクト

ここまでの説明がわからなかったり、バグを発見したり、もし一緒に開発したいと思ってくださる方がいましたら、ぜひ 「yaylib」 の <a href="https://discord.gg/MEuBfNtqRN">Discord サーバー</a>に参加して声をかけてください。
