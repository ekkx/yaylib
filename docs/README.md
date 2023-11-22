<div><a id="readme-top"></a></div>
<p align="center">
    <img src="https://github.com/qvco/yaylib/assets/77382767/45c45b21-d812-4cad-8f27-315ffef53201" alt="Logo" height="300px">
    <h3 align="center">yaylib - Docs | ドキュメント</h3>
    <p align="center">
        <br />
        <a href="https://github.com/qvco/yaylib/blob/master/docs/API-Reference/README.md">
            <strong>API reference index »</strong>
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

<details open>
    <summary>目録</summary>
    <ol>
        <li><a href="#インストール--install">インストール</a></li>
        <li><a href="#使用方法--usage">使用方法</a></li>
        <!-- <li><a href="async-client">Async Client</a></li> -->
        <li><a href="#メソッド一覧">メソッド一覧</a></li>
        <li><a href="#オブジェクトモデル一覧">オブジェクトモデル一覧</a></li>
        <li><a href="#yay-エラーコード一覧">Yay! エラーコード一覧</a></li>
    </ol>
</details>

## インストール / Install

**※ Python 3.10 以上のバージョンが必要です。**

「yaylib」をインストールするには、以下のコマンドをターミナル上で実行します:

```bash
pip install yaylib
```

<br>

> **Note**
> 開発バージョンをインストールする場合は、以下の手順を実行します:

```bash
git clone https://github.com/qvco/yaylib.git

cd yaylib

pip install -r requirements.txt

pip install -e .
```

## 使用方法 / Usage

以下の例は、簡単なクライアントの初期化方法です。

```python
import yaylib

api = yaylib.Client()

email = "メールアドレス"
password = "パスワード"

api.login(email, password)
```

「yaylib」ではログインの回数制限を回避するために、ローカルストレージに保存された認証情報を再利用します。

他人に見られる可能性のある開発環境下などで、認証情報を暗号化して保存したい場合は、`Client`クラスを初期化する際に、`encrypt_cookie`引数を`True`に設定します。

また、メールアドレスやパスワードなどの機密情報を不正なアクセスから保護するために、.env ファイルなどを使用して環境変数とソースコードを分けるようにしてください。

**.env**

```sh
EMAIL=メールアドレス
PASSWORD=パスワード
```

**sample.py**

```python
import yaylib

# 必要モジュールのインポート
import os
from dotenv import load_dotenv

# .envファイルの内容を読み込む
load_dotenv()

# os.environを用いて環境変数を参照
email = os.environ.get("EMAIL")
password = os.environ.get("PASSWORD")

api = yaylib.Client(encrypt_cookie=True)

api.login(email, password)
```

そうすると、初回実行時に`secret_key`という鍵がターミナルに表示されます。

![image](https://github.com/qvco/yaylib/assets/77382767/32b3f6c6-833c-4e81-abbe-30a73b0233b9)

※ 緑色の文字列が`secret_key`。

保存した認証情報を復号化し、再利用するには`secret_key`を`login()`メソッドの引数に設定します。

**.env**

```sh
EMAIL=メールアドレス
PASSWORD=パスワード
SECRET_KEY=lCo3vhaCQOaBxRdMe-ZyTUiTUjiRkrPX7vQR2nDezxc=
```

**sample.py**

```python
import yaylib

# 必要モジュールのインポート
import os
from dotenv import load_dotenv

# .envファイルの内容を読み込む
load_dotenv()

# os.environを用いて環境変数を参照
email = os.environ.get("EMAIL")
password = os.environ.get("PASSWORD")
secret_key = os.environ.get("SECRET_KEY")

api = yaylib.Client(encrypt_cookie=True)

api.login(email, password, secret_key)
```

また、`secret_key`を取得する代替方法として、`Client`クラスの`secret_key`プロパティにアクセスすることで取得することも出来ます。

```python
api.login(email, password)

# ログインした後にsecret_keyを取得
api.secret_key
>>> lCo3vhaCQOaBxRdMe-ZyTUiTUjiRkrPX7vQR2nDezxc=
```

### Client クラスについて

Client クラスは、API クライアントのメインクラスです。

<table>
    <tr>
		<th>引数</th>
		<th>型</th>
		<th>初期値</th>
		<th>概要</th>
	</tr>
    <tr>
		<td><code>access_token</code></td>
		<td><code>str</code></td>
		<td><code>None</code></td>
		<td>操作の認証に必要なアクセストークン</td>
	</tr>
    <tr>
		<td><code>proxy_url</code></td>
		<td><code>str</code></td>
		<td><code>None</code></td>
		<td>プロキシサーバーのアドレス</td>
	</tr>
    <tr>
		<td><code>max_retries</code></td>
		<td><code>int</code></td>
		<td><code>3</code></td>
		<td>リクエストに失敗した際のリトライ回数</td>
	</tr>
    <tr>
		<td><code>backoff_factor</code></td>
		<td><code>float</code></td>
		<td><code>1.5</code></td>
		<td>リトライ待機時間の増加割合を指定する係数</td>
	</tr>
    <tr>
		<td><code>wait_on_rate_limit</code></td>
		<td><code>bool</code></td>
		<td><code>True</code></td>
		<td>レート制限を待機するかどうか</td>
	</tr>
    <tr>
		<td><code>min_delay</code></td>
		<td><code>float</code></td>
		<td><code>0.3</code></td>
		<td>最短遅延時間</td>
	</tr>
    <tr>
		<td><code>max_delay</code></td>
		<td><code>float</code></td>
		<td><code>1.2</code></td>
		<td>最長遅延時間</td>
	</tr>
    <tr>
		<td><code>timeout</code></td>
		<td><code>int</code></td>
		<td><code>30</code></td>
		<td>通信接続を待機する時間</td>
	</tr>
    <tr>
		<td><code>base_path</code></td>
		<td><code>str</code></td>
		<td><code>作業ディレクトリ</code></td>
		<td>ベースパス</td>
	</tr>
    <tr>
		<td><code>err_lang</code></td>
		<td><code>str</code></td>
		<td><code>ja</code></td>
		<td>エラー出力言語（ja もしくは en）</td>
	</tr>
    <tr>
		<td><code>save_cookie_file</code></td>
		<td><code>bool</code></td>
		<td><code>True</code></td>
		<td>クッキーを保存するかどうか</td>
	</tr>
    <tr>
		<td><code>encrypt_cookie</code></td>
		<td><code>bool</code></td>
		<td><code>False</code></td>
		<td>暗号化してクッキーを保存するかどうか</td>
	</tr>
    <tr>
		<td><code>cookie_filename</code></td>
		<td><code>str</code></td>
		<td><code>cookies</code></td>
		<td>クッキーを保存するファイル名</td>
	</tr>
    <tr>
		<td><code>loglevel</code></td>
		<td><code>int</code></td>
		<td><code>logging.INFO</code></td>
		<td>ログレベルの設定</td>
	</tr>
</table>

<p align="right">(<a href="#readme-top">トップに戻る</a>)</p>

## メソッド一覧

### 投稿

<table>
    <tr>
		<th>メソッド</th>
		<th>概要</th>
	</tr>
	<tr>
		<td><code><a href="https://github.com/qvco/yaylib/blob/master/docs/API-Reference/post/add_bookmark.md">add_bookmark()</a></code></td>
		<td>ブックマークに追加します</td>
	</tr>
	<tr>
		<td><code><a href="https://github.com/qvco/yaylib/blob/master/docs/API-Reference/post/add_group_highlight_post.md">add_group_highlight_post()</a></code></td>
		<td>投稿をグループのまとめに追加します</td>
	</tr>
	<tr>
		<td><code><a href="https://github.com/qvco/yaylib/blob/master/docs/API-Reference/post/create_call_post.md">create_call_post()</a></code></td>
		<td>通話の投稿を作成します</td>
	</tr>
	<tr>
		<td><code><a href="https://github.com/qvco/yaylib/blob/master/docs/API-Reference/post/pin_group_post.md">pin_group_post()</a></code></td>
		<td>グループの投稿をピンします</td>
	</tr>
	<tr>
		<td><code><a href="https://github.com/qvco/yaylib/blob/master/docs/API-Reference/post/pin_post.md">pin_post()</a></code></td>
		<td>投稿をピンします</td>
	</tr>
	<tr>
		<td><code><a href="https://github.com/qvco/yaylib/blob/master/docs/API-Reference/post/mention.md">mention()</a></code></td>
		<td>メンション用文字列を返します</td>
	</tr>
	<tr>
		<td><code><a href="https://github.com/qvco/yaylib/blob/master/docs/API-Reference/post/create_post.md">create_post()</a></code></td>
		<td>投稿を作成します</td>
	</tr>
	<tr>
		<td><code><a href="https://github.com/qvco/yaylib/blob/master/docs/API-Reference/post/create_repost.md">create_repost()</a></code></td>
		<td>投稿を(´∀｀∩)↑age↑します</td>
	</tr>
	<tr>
		<td><code><a href="https://github.com/qvco/yaylib/blob/master/docs/API-Reference/post/create_share_post.md">create_share_post()</a></code></td>
		<td>シェア投稿を作成します</td>
	</tr>
	<tr>
		<td><code><a href="https://github.com/qvco/yaylib/blob/master/docs/API-Reference/post/create_thread_post.md">create_thread_post()</a></code></td>
		<td>スレッドの投稿を作成します</td>
	</tr>
	<tr>
		<td><code><a href="https://github.com/qvco/yaylib/blob/master/docs/API-Reference/post/delete_all_post.md">delete_all_post()</a></code></td>
		<td>すべての自分の投稿を削除します</td>
	</tr>
	<tr>
		<td><code><a href="https://github.com/qvco/yaylib/blob/master/docs/API-Reference/post/unpin_group_post.md">unpin_group_post()</a></code></td>
		<td>グループのピン投稿を解除します</td>
	</tr>
	<tr>
		<td><code><a href="https://github.com/qvco/yaylib/blob/master/docs/API-Reference/post/unpin_post.md">unpin_post()</a></code></td>
		<td>ピン投稿を削除します</td>
	</tr>
	<tr>
		<td><code><a href="https://github.com/qvco/yaylib/blob/master/docs/API-Reference/post/get_bookmark.md">get_bookmark()</a></code></td>
		<td>ブックマークを取得します</td>
	</tr>
	<tr>
		<td><code><a href="https://github.com/qvco/yaylib/blob/master/docs/API-Reference/post/get_timeline_calls.md">get_timeline_calls()</a></code></td>
		<td>誰でも通話を取得します</td>
	</tr>
	<tr>
		<td><code><a href="https://github.com/qvco/yaylib/blob/master/docs/API-Reference/post/get_conversation.md">get_conversation()</a></code></td>
		<td>リプライを含める投稿の会話を取得します</td>
	</tr>
	<tr>
		<td><code><a href="https://github.com/qvco/yaylib/blob/master/docs/API-Reference/post/get_following_call_timeline.md">get_following_call_timeline()</a></code></td>
		<td>フォロー中の通話を取得します</td>
	</tr>
	<tr>
		<td><code><a href="https://github.com/qvco/yaylib/blob/master/docs/API-Reference/post/get_following_timeline.md">get_following_timeline()</a></code></td>
		<td>フォロー中のタイムラインを取得します</td>
	</tr>
	<tr>
		<td><code><a href="https://github.com/qvco/yaylib/blob/master/docs/API-Reference/post/get_group_highlight_posts.md">get_group_highlight_posts()</a></code></td>
		<td>グループのハイライト投稿を取得します</td>
	</tr>
	<tr>
		<td><code><a href="https://github.com/qvco/yaylib/blob/master/docs/API-Reference/post/get_group_timeline_by_keyword.md">get_group_timeline_by_keyword()</a></code></td>
		<td>グループの投稿をキーワードで検索します</td>
	</tr>
	<tr>
		<td><code><a href="https://github.com/qvco/yaylib/blob/master/docs/API-Reference/post/get_group_timeline.md">get_group_timeline()</a></code></td>
		<td>グループのタイムラインを取得します</td>
	</tr>
	<tr>
		<td><code><a href="https://github.com/qvco/yaylib/blob/master/docs/API-Reference/post/get_timeline_by_hashtag.md">get_timeline_by_hashtag()</a></code></td>
		<td>ハッシュタグでタイムラインを検索します</td>
	</tr>
	<tr>
		<td><code><a href="https://github.com/qvco/yaylib/blob/master/docs/API-Reference/post/get_my_posts.md">get_my_posts()</a></code></td>
		<td>自分の投稿を取得します</td>
	</tr>
	<tr>
		<td><code><a href="https://github.com/qvco/yaylib/blob/master/docs/API-Reference/post/get_post.md">get_post()</a></code></td>
		<td>投稿の詳細を取得します</td>
	</tr>
	<tr>
		<td><code><a href="https://github.com/qvco/yaylib/blob/master/docs/API-Reference/post/get_post_likers.md">get_post_likers()</a></code></td>
		<td>投稿にいいねしたユーザーを取得します</td>
	</tr>
	<tr>
		<td><code><a href="https://github.com/qvco/yaylib/blob/master/docs/API-Reference/post/get_reposts.md">get_reposts()</a></code></td>
		<td>投稿の(´∀｀∩)↑age↑を取得します</td>
	</tr>
	<tr>
		<td><code><a href="https://github.com/qvco/yaylib/blob/master/docs/API-Reference/post/get_posts.md">get_posts()</a></code></td>
		<td>複数の投稿を取得します</td>
	</tr>
	<tr>
		<td><code><a href="https://github.com/qvco/yaylib/blob/master/docs/API-Reference/post/get_recommended_post_tags.md">get_recommended_post_tags()</a></code></td>
		<td>おすすめのタグ候補を取得します</td>
	</tr>
	<tr>
		<td><code><a href="https://github.com/qvco/yaylib/blob/master/docs/API-Reference/post/get_recommended_posts.md">get_recommended_posts()</a></code></td>
		<td>おすすめの投稿を取得します</td>
	</tr>
	<tr>
		<td><code><a href="https://github.com/qvco/yaylib/blob/master/docs/API-Reference/post/get_timeline_by_keyword.md">get_timeline_by_keyword()</a></code></td>
		<td>キーワードでタイムラインを検索します</td>
	</tr>
	<tr>
		<td><code><a href="https://github.com/qvco/yaylib/blob/master/docs/API-Reference/post/get_timeline.md">get_timeline()</a></code></td>
		<td>タイムラインを取得します</td>
	</tr>
	<tr>
		<td><code><a href="https://github.com/qvco/yaylib/blob/master/docs/API-Reference/post/get_url_metadata.md">get_url_metadata()</a></code></td>
		<td>URLのメタデータを取得します</td>
	</tr>
	<tr>
		<td><code><a href="https://github.com/qvco/yaylib/blob/master/docs/API-Reference/post/get_user_timeline.md">get_user_timeline()</a></code></td>
		<td>ユーザーのタイムラインを取得します</td>
	</tr>
	<tr>
		<td><code><a href="https://github.com/qvco/yaylib/blob/master/docs/API-Reference/post/like.md">like()</a></code></td>
		<td>投稿にいいねします</td>
	</tr>
	<tr>
		<td><code><a href="https://github.com/qvco/yaylib/blob/master/docs/API-Reference/post/remove_bookmark.md">remove_bookmark()</a></code></td>
		<td>ブックマークを削除します</td>
	</tr>
	<tr>
		<td><code><a href="https://github.com/qvco/yaylib/blob/master/docs/API-Reference/post/remove_group_highlight_post.md">remove_group_highlight_post()</a></code></td>
		<td>サークルのハイライト投稿を解除します</td>
	</tr>
	<tr>
		<td><code><a href="https://github.com/qvco/yaylib/blob/master/docs/API-Reference/post/remove_posts.md">remove_posts()</a></code></td>
		<td>複数の投稿を削除します</td>
	</tr>
	<tr>
		<td><code><a href="https://github.com/qvco/yaylib/blob/master/docs/API-Reference/post/report_post.md">report_post()</a></code></td>
		<td>投稿を通報します</td>
	</tr>
	<tr>
		<td><code><a href="https://github.com/qvco/yaylib/blob/master/docs/API-Reference/post/unlike.md">unlike()</a></code></td>
		<td>投稿のいいねを解除します</td>
	</tr>
	<tr>
		<td><code><a href="https://github.com/qvco/yaylib/blob/master/docs/API-Reference/post/update_post.md">update_post()</a></code></td>
		<td>投稿を編集します</td>
	</tr>
	<tr>
		<td><code><a href="https://github.com/qvco/yaylib/blob/master/docs/API-Reference/post/view_video.md">view_video()</a></code></td>
		<td>動画を視聴します</td>
	</tr>
	<tr>
		<td><code><a href="https://github.com/qvco/yaylib/blob/master/docs/API-Reference/post/vote_survey.md">vote_survey()</a></code></td>
		<td>アンケートに投票します</td>
	</tr>
</table>

### スレッド

<table>
    <tr>
		<th>メソッド</th>
		<th>概要</th>
	</tr>
	<tr>
		<td><code><a href="https://github.com/qvco/yaylib/blob/master/docs/API-Reference/thread/add_post_to_thread.md">add_post_to_thread()</a></code></td>
		<td>投稿をスレッドに追加します</td>
	</tr>
	<tr>
		<td><code><a href="https://github.com/qvco/yaylib/blob/master/docs/API-Reference/thread/convert_post_to_thread.md">convert_post_to_thread()</a></code></td>
		<td>投稿をスレッドに変換します</td>
	</tr>
	<tr>
		<td><code><a href="https://github.com/qvco/yaylib/blob/master/docs/API-Reference/thread/create_thread.md">create_thread()</a></code></td>
		<td>スレッドを作成します</td>
	</tr>
	<tr>
		<td><code><a href="https://github.com/qvco/yaylib/blob/master/docs/API-Reference/thread/get_group_thread_list.md">get_group_thread_list()</a></code></td>
		<td>グループのスレッド一覧を取得します</td>
	</tr>
	<tr>
		<td><code><a href="https://github.com/qvco/yaylib/blob/master/docs/API-Reference/thread/get_thread_joined_statuses.md">get_thread_joined_statuses()</a></code></td>
		<td>スレッド参加ステータスを取得します</td>
	</tr>
	<tr>
		<td><code><a href="https://github.com/qvco/yaylib/blob/master/docs/API-Reference/thread/get_thread_posts.md">get_thread_posts()</a></code></td>
		<td>スレッドの投稿を取得します</td>
	</tr>
	<tr>
		<td><code><a href="https://github.com/qvco/yaylib/blob/master/docs/API-Reference/thread/join_thread.md">join_thread()</a></code></td>
		<td>スレッドに参加します</td>
	</tr>
	<tr>
		<td><code><a href="https://github.com/qvco/yaylib/blob/master/docs/API-Reference/thread/leave_thread.md">leave_thread()</a></code></td>
		<td>スレッドから脱退します</td>
	</tr>
	<tr>
		<td><code><a href="https://github.com/qvco/yaylib/blob/master/docs/API-Reference/thread/remove_thread.md">remove_thread()</a></code></td>
		<td>スレッドを削除します</td>
	</tr>
	<tr>
		<td><code><a href="https://github.com/qvco/yaylib/blob/master/docs/API-Reference/thread/update_thread.md">update_thread()</a></code></td>
		<td>スレッドを編集します</td>
	</tr>
</table>

### ユーザー

<table>
    <tr>
		<th>メソッド</th>
		<th>概要</th>
	</tr>
	<tr>
		<td><code><a href="https://github.com/qvco/yaylib/blob/master/docs/API-Reference/user/delete_footprint.md">delete_footprint()</a></code></td>
		<td>足跡を削除します</td>
	</tr>
	<tr>
		<td><code><a href="https://github.com/qvco/yaylib/blob/master/docs/API-Reference/user/destroy_user.md">destroy_user()</a></code></td>
		<td>アカウントを削除します</td>
	</tr>
	<tr>
		<td><code><a href="https://github.com/qvco/yaylib/blob/master/docs/API-Reference/user/follow_user.md">follow_user()</a></code></td>
		<td>ユーザーをフォローします</td>
	</tr>
	<tr>
		<td><code><a href="https://github.com/qvco/yaylib/blob/master/docs/API-Reference/user/follow_users.md">follow_users()</a></code></td>
		<td>複数のユーザーをフォローします</td>
	</tr>
	<tr>
		<td><code><a href="https://github.com/qvco/yaylib/blob/master/docs/API-Reference/user/get_active_followings.md">get_active_followings()</a></code></td>
		<td>アクティブなフォロー中のユーザーを取得します</td>
	</tr>
	<tr>
		<td><code><a href="https://github.com/qvco/yaylib/blob/master/docs/API-Reference/user/get_follow_recommendations.md">get_follow_recommendations()</a></code></td>
		<td>フォローするのにおすすめのユーザーを取得します</td>
	</tr>
	<tr>
		<td><code><a href="https://github.com/qvco/yaylib/blob/master/docs/API-Reference/user/get_follow_request.md">get_follow_request()</a></code></td>
		<td>フォローリクエストを取得します</td>
	</tr>
	<tr>
		<td><code><a href="https://github.com/qvco/yaylib/blob/master/docs/API-Reference/user/get_follow_request_count.md">get_follow_request_count()</a></code></td>
		<td>フォローリクエストの数を取得します</td>
	</tr>
	<tr>
		<td><code><a href="https://github.com/qvco/yaylib/blob/master/docs/API-Reference/user/get_following_users_born.md">get_following_users_born()</a></code></td>
		<td>フォロー中のユーザーの誕生日を取得します</td>
	</tr>
	<tr>
		<td><code><a href="https://github.com/qvco/yaylib/blob/master/docs/API-Reference/user/get_footprints.md">get_footprints()</a></code></td>
		<td>足跡を取得します</td>
	</tr>
	<tr>
		<td><code><a href="https://github.com/qvco/yaylib/blob/master/docs/API-Reference/user/get_fresh_user.md">get_fresh_user()</a></code></td>
		<td>認証情報などを含んだユーザー情報を取得します</td>
	</tr>
	<tr>
		<td><code><a href="https://github.com/qvco/yaylib/blob/master/docs/API-Reference/user/get_hima_users.md">get_hima_users()</a></code></td>
		<td>暇なユーザーを取得します</td>
	</tr>
	<tr>
		<td><code><a href="https://github.com/qvco/yaylib/blob/master/docs/API-Reference/user/get_user_ranking.md">helget_user_rankinglo()</a></code></td>
		<td>ユーザーのフォロワーランキングを取得します</td>
	</tr>
	<tr>
		<td><code><a href="https://github.com/qvco/yaylib/blob/master/docs/API-Reference/user/get_refresh_counter_requests.md">get_refresh_counter_requests()</a></code></td>
		<td>カウンター更新のリクエストを取得します</td>
	</tr>
	<tr>
		<td><code><a href="https://github.com/qvco/yaylib/blob/master/docs/API-Reference/user/get_social_shared_users.md">get_social_shared_users()</a></code></td>
		<td>SNS共有をしたユーザーを取得します</td>
	</tr>
	<tr>
		<td><code><a href="https://github.com/qvco/yaylib/blob/master/docs/API-Reference/user/get_timestamp.md">get_timestamp()</a></code></td>
		<td>タイムスタンプを取得します</td>
	</tr>
	<tr>
		<td><code><a href="https://github.com/qvco/yaylib/blob/master/docs/API-Reference/user/get_user.md">get_user()</a></code></td>
		<td>ユーザーの情報を取得します</td>
	</tr>
	<tr>
		<td><code><a href="https://github.com/qvco/yaylib/blob/master/docs/API-Reference/user/get_user_email.md">get_user_email()</a></code></td>
		<td>ユーザーのメールアドレスを取得します</td>
	</tr>
	<tr>
		<td><code><a href="https://github.com/qvco/yaylib/blob/master/docs/API-Reference/user/get_user_followers.md">get_user_followers()</a></code></td>
		<td>ユーザーのフォロワーを取得します</td>
	</tr>
	<tr>
		<td><code><a href="https://github.com/qvco/yaylib/blob/master/docs/API-Reference/user/get_user_followings.md">get_user_followings()</a></code></td>
		<td>フォロー中のユーザーを取得します</td>
	</tr>
	<tr>
		<td><code><a href="https://github.com/qvco/yaylib/blob/master/docs/API-Reference/user/get_user_from_qr.md">get_user_from_qr()</a></code></td>
		<td>QRコードからユーザーを取得します</td>
	</tr>
	<tr>
		<td><code><a href="https://github.com/qvco/yaylib/blob/master/docs/API-Reference/user/get_user_without_leaving_footprint.md">get_user_without_leaving_footprint()</a></code></td>
		<td>足跡をつけずにユーザーの情報を取得します</td>
	</tr>
	<tr>
		<td><code><a href="https://github.com/qvco/yaylib/blob/master/docs/API-Reference/user/get_users.md">get_users()</a></code></td>
		<td>複数のユーザーの情報を取得します</td>
	</tr>
	<tr>
		<td><code><a href="https://github.com/qvco/yaylib/blob/master/docs/API-Reference/user/reduce_kenta_penalty.md">reduce_kenta_penalty()</a></code></td>
		<td>ペナルティーを緩和します</td>
	</tr>
	<tr>
		<td><code><a href="https://github.com/qvco/yaylib/blob/master/docs/API-Reference/user/refresh_counter.md">refresh_counter()</a></code></td>
		<td>カウンターを更新します</td>
	</tr>
	<tr>
		<td><code><a href="https://github.com/qvco/yaylib/blob/master/docs/API-Reference/user/remove_user_avatar.md">remove_user_avatar()</a></code></td>
		<td>ユーザーのアイコンを削除します</td>
	</tr>
	<tr>
		<td><code><a href="https://github.com/qvco/yaylib/blob/master/docs/API-Reference/user/remove_user_cover.md">remove_user_cover()</a></code></td>
		<td>ユーザーのカバー画像を削除します</td>
	</tr>
	<tr>
		<td><code><a href="https://github.com/qvco/yaylib/blob/master/docs/API-Reference/user/report_user.md">report_user()</a></code></td>
		<td>ユーザーを通報します</td>
	</tr>
	<tr>
		<td><code><a href="https://github.com/qvco/yaylib/blob/master/docs/API-Reference/user/reset_password.md">reset_password()</a></code></td>
		<td>パスワードをリセットします</td>
	</tr>
	<tr>
		<td><code><a href="https://github.com/qvco/yaylib/blob/master/docs/API-Reference/user/search_lobi_users.md">search_lobi_users()</a></code></td>
		<td>Lobiのユーザーを検索します</td>
	</tr>
	<tr>
		<td><code><a href="https://github.com/qvco/yaylib/blob/master/docs/API-Reference/user/search_users.md">search_users()</a></code></td>
		<td>ユーザーを検索します</td>
	</tr>
	<tr>
		<td><code><a href="https://github.com/qvco/yaylib/blob/master/docs/API-Reference/user/set_follow_permission_enabled.md">set_follow_permission_enabled()</a></code></td>
		<td>フォローを許可制に設定します</td>
	</tr>
	<tr>
		<td><code><a href="https://github.com/qvco/yaylib/blob/master/docs/API-Reference/user/take_action_follow_request.md">take_action_follow_request()</a></code></td>
		<td></td>
	</tr>
	<tr>
		<td><code><a href="https://github.com/qvco/yaylib/blob/master/docs/API-Reference/user/turn_on_hima.md">turn_on_hima()</a></code></td>
		<td>ひまなうします</td>
	</tr>
	<tr>
		<td><code><a href="https://github.com/qvco/yaylib/blob/master/docs/API-Reference/user/unfollow_user.md">unfollow_user()</a></code></td>
		<td>ユーザーをアンフォローします</td>
	</tr>
	<tr>
		<td><code><a href="https://github.com/qvco/yaylib/blob/master/docs/API-Reference/user/update_user.md">update_user()</a></code></td>
		<td>プロフィールを更新します</td>
	</tr>
	<tr>
		<td><code><a href="https://github.com/qvco/yaylib/blob/master/docs/API-Reference/user/block_user.md">block_user()</a></code></td>
		<td>ユーザーをブロックします</td>
	</tr>
	<tr>
		<td><code><a href="https://github.com/qvco/yaylib/blob/master/docs/API-Reference/user/get_blocked_user_ids.md">get_blocked_user_ids()</a></code></td>
		<td>あなたをブロックしたユーザーを取得します</td>
	</tr>
	<tr>
		<td><code><a href="https://github.com/qvco/yaylib/blob/master/docs/API-Reference/user/get_blocked_users.md">get_blocked_users()</a></code></td>
		<td>ブロックしたユーザーを取得します</td>
	</tr>
	<tr>
		<td><code><a href="https://github.com/qvco/yaylib/blob/master/docs/API-Reference/user/unblock_user.md">unblock_user()</a></code></td>
		<td>ユーザーをアンブロックします</td>
	</tr>
	<tr>
		<td><code><a href="https://github.com/qvco/yaylib/blob/master/docs/API-Reference/user/get_hidden_users_list.md">get_hidden_users_list()</a></code></td>
		<td>非表示のユーザー一覧を取得します</td>
	</tr>
	<tr>
		<td><code><a href="https://github.com/qvco/yaylib/blob/master/docs/API-Reference/user/hide_user.md">hide_user()</a></code></td>
		<td>ユーザーを非表示にします</td>
	</tr>
	<tr>
		<td><code><a href="https://github.com/qvco/yaylib/blob/master/docs/API-Reference/user/unhide_users.md">unhide_users()</a></code></td>
		<td>ユーザーの非表示を解除します</td>
	</tr>
</table>

### サークル

<table>
    <tr>
		<th>メソッド</th>
		<th>概要</th>
	</tr>
	<tr>
		<td><code><a href="https://github.com/qvco/yaylib/blob/master/docs/API-Reference/group/accept_moderator_offer.md">accept_moderator_offer()</a></code></td>
		<td>サークル副管理人の権限オファーを引き受けます</td>
	</tr>
	<tr>
		<td><code><a href="https://github.com/qvco/yaylib/blob/master/docs/API-Reference/group/accept_ownership_offer.md">accept_ownership_offer()</a></code></td>
		<td>サークル管理人の権限オファーを引き受けます</td>
	</tr>
	<tr>
		<td><code><a href="https://github.com/qvco/yaylib/blob/master/docs/API-Reference/group/accept_group_join_request.md">accept_group_join_request()</a></code></td>
		<td>サークル参加リクエストを承認します</td>
	</tr>
	<tr>
		<td><code><a href="https://github.com/qvco/yaylib/blob/master/docs/API-Reference/group/add_related_groups.md">add_related_groups()</a></code></td>
		<td>関連サークルを追加します</td>
	</tr>
	<tr>
		<td><code><a href="https://github.com/qvco/yaylib/blob/master/docs/API-Reference/group/ban_group_user.md">ban_group_user()</a></code></td>
		<td>サークルからユーザーを追放します</td>
	</tr>
	<tr>
		<td><code><a href="https://github.com/qvco/yaylib/blob/master/docs/API-Reference/group/check_unread_status.md">check_unread_status()</a></code></td>
		<td>サークルの未読ステータスを取得します</td>
	</tr>
	<tr>
		<td><code><a href="https://github.com/qvco/yaylib/blob/master/docs/API-Reference/group/create_group.md">create_group()</a></code></td>
		<td>サークルを作成します</td>
	</tr>
	<tr>
		<td><code><a href="https://github.com/qvco/yaylib/blob/master/docs/API-Reference/group/pin_group.md">pin_group()</a></code></td>
		<td>サークルをピンします</td>
	</tr>
	<tr>
		<td><code><a href="https://github.com/qvco/yaylib/blob/master/docs/API-Reference/group/decline_moderator_offer.md">decline_moderator_offer()</a></code></td>
		<td>サークル副管理人の権限オファーを断ります</td>
	</tr>
	<tr>
		<td><code><a href="https://github.com/qvco/yaylib/blob/master/docs/API-Reference/group/decline_ownership_offer.md">decline_ownership_offer()</a></code></td>
		<td>サークル管理人の権限オファーを断ります</td>
	</tr>
	<tr>
		<td><code><a href="https://github.com/qvco/yaylib/blob/master/docs/API-Reference/group/decline_group_join_request.md">decline_group_join_request()</a></code></td>
		<td>サークル参加リクエストを断ります</td>
	</tr>
	<tr>
		<td><code><a href="https://github.com/qvco/yaylib/blob/master/docs/API-Reference/group/unpin_group.md">unpin_group()</a></code></td>
		<td>サークルのピン止めを解除します</td>
	</tr>
	<tr>
		<td><code><a href="https://github.com/qvco/yaylib/blob/master/docs/API-Reference/group/get_banned_group_members.md">get_banned_group_members()</a></code></td>
		<td>追放されたサークルメンバーを取得します</td>
	</tr>
	<tr>
		<td><code><a href="https://github.com/qvco/yaylib/blob/master/docs/API-Reference/group/get_group_categories.md">get_group_categories()</a></code></td>
		<td>サークルのカテゴリーを取得します</td>
	</tr>
	<tr>
		<td><code><a href="https://github.com/qvco/yaylib/blob/master/docs/API-Reference/group/get_create_group_quota.md">get_create_group_quota()</a></code></td>
		<td>サークル作成可能な割当量を取得します</td>
	</tr>
	<tr>
		<td><code><a href="https://github.com/qvco/yaylib/blob/master/docs/API-Reference/group/get_group.md">get_group()</a></code></td>
		<td>サークルの詳細を取得します</td>
	</tr>
	<tr>
		<td><code><a href="https://github.com/qvco/yaylib/blob/master/docs/API-Reference/group/get_groups.md">get_groups()</a></code></td>
		<td>複数のサークルの詳細を取得します</td>
	</tr>
	<tr>
		<td><code><a href="https://github.com/qvco/yaylib/blob/master/docs/API-Reference/group/get_invitable_users.md">get_invitable_users()</a></code></td>
		<td>サークルに招待可能なユーザーを取得します</td>
	</tr>
	<tr>
		<td><code><a href="https://github.com/qvco/yaylib/blob/master/docs/API-Reference/group/get_joined_statuses.md">get_joined_statuses()</a></code></td>
		<td>サークルの参加ステータスを取得します</td>
	</tr>
	<tr>
		<td><code><a href="https://github.com/qvco/yaylib/blob/master/docs/API-Reference/group/get_group_member.md">get_group_member()</a></code></td>
		<td>特定のサークルメンバーの情報を取得します</td>
	</tr>
	<tr>
		<td><code><a href="https://github.com/qvco/yaylib/blob/master/docs/API-Reference/group/get_group_members.md">get_group_members()</a></code></td>
		<td>サークルメンバーを取得します</td>
	</tr>
	<tr>
		<td><code><a href="https://github.com/qvco/yaylib/blob/master/docs/API-Reference/group/get_my_groups.md">get_my_groups()</a></code></td>
		<td>自分のサークルを取得します</td>
	</tr>
	<tr>
		<td><code><a href="https://github.com/qvco/yaylib/blob/master/docs/API-Reference/group/get_relatable_groups.md">get_relatable_groups()</a></code></td>
		<td>関連がある可能性があるサークルを取得します</td>
	</tr>
	<tr>
		<td><code><a href="https://github.com/qvco/yaylib/blob/master/docs/API-Reference/group/get_related_groups.md">get_related_groups()</a></code></td>
		<td>関連があるサークルを取得します</td>
	</tr>
	<tr>
		<td><code><a href="https://github.com/qvco/yaylib/blob/master/docs/API-Reference/group/get_user_groups.md">get_user_groups()</a></code></td>
		<td>特定のユーザーが参加しているサークルを取得します</td>
	</tr>
	<tr>
		<td><code><a href="https://github.com/qvco/yaylib/blob/master/docs/API-Reference/group/invite_users_to_group.md">invite_users_to_group()</a></code></td>
		<td>サークルにユーザーを招待します</td>
	</tr>
	<tr>
		<td><code><a href="https://github.com/qvco/yaylib/blob/master/docs/API-Reference/group/join_group.md">join_group()</a></code></td>
		<td>サークルに参加します</td>
	</tr>
	<tr>
		<td><code><a href="https://github.com/qvco/yaylib/blob/master/docs/API-Reference/group/leave_group.md">leave_group()</a></code></td>
		<td>サークルから脱退します</td>
	</tr>
	<tr>
		<td><code><a href="https://github.com/qvco/yaylib/blob/master/docs/API-Reference/group/post_gruop_social_shared.md">post_gruop_social_shared()</a></code></td>
		<td>サークルのソーシャルシェアを投稿します</td>
	</tr>
	<tr>
		<td><code><a href="https://github.com/qvco/yaylib/blob/master/docs/API-Reference/group/remove_group_cover.md">remove_group_cover()</a></code></td>
		<td>サークルのカバー画像を削除します</td>
	</tr>
	<tr>
		<td><code><a href="https://github.com/qvco/yaylib/blob/master/docs/API-Reference/group/remove_moderator.md">remove_moderator()</a></code></td>
		<td>サークルの副管理人を削除します</td>
	</tr>
	<tr>
		<td><code><a href="https://github.com/qvco/yaylib/blob/master/docs/API-Reference/group/report_group.md">report_group()</a></code></td>
		<td>サークルを通報します</td>
	</tr>
	<tr>
		<td><code><a href="https://github.com/qvco/yaylib/blob/master/docs/API-Reference/group/send_moderator_offers.md">send_moderator_offers()</a></code></td>
		<td>複数人にサークル副管理人のオファーを送信します</td>
	</tr>
	<tr>
		<td><code><a href="https://github.com/qvco/yaylib/blob/master/docs/API-Reference/group/send_ownership_offer.md">send_ownership_offer()</a></code></td>
		<td>サークル管理人権限のオファーを送信します</td>
	</tr>
	<tr>
		<td><code><a href="https://github.com/qvco/yaylib/blob/master/docs/API-Reference/group/set_group_title.md">set_group_title()</a></code></td>
		<td>サークルのタイトルを設定します</td>
	</tr>
	<tr>
		<td><code><a href="https://github.com/qvco/yaylib/blob/master/docs/API-Reference/group/take_over_group_ownership.md">take_over_group_ownership()</a></code></td>
		<td>サークル管理人の権限を引き継ぎます</td>
	</tr>
	<tr>
		<td><code><a href="https://github.com/qvco/yaylib/blob/master/docs/API-Reference/group/unban_group_member.md">unban_group_member()</a></code></td>
		<td>特定のサークルメンバーの追放を解除します</td>
	</tr>
	<tr>
		<td><code><a href="https://github.com/qvco/yaylib/blob/master/docs/API-Reference/group/update_group.md">update_group()</a></code></td>
		<td>サークルを編集します</td>
	</tr>
	<tr>
		<td><code><a href="https://github.com/qvco/yaylib/blob/master/docs/API-Reference/group/withdraw_moderator_offer.md">withdraw_moderator_offer()</a></code></td>
		<td>サークル副管理人のオファーを取り消します</td>
	</tr>
	<tr>
		<td><code><a href="https://github.com/qvco/yaylib/blob/master/docs/API-Reference/group/withdraw_ownership_offer.md">withdraw_ownership_offer()</a></code></td>
		<td>サークル管理人のオファーを取り消します</td>
	</tr>
</table>

### チャット

<table>
    <tr>
		<th>メソッド</th>
		<th>概要</th>
	</tr>
	<tr>
		<td><code><a href="https://github.com/qvco/yaylib/blob/master/docs/API-Reference/chat/accept_chat_request.md">accept_chat_request()</a></code></td>
		<td>チャットリクエストを承認します</td>
	</tr>
	<tr>
		<td><code><a href="https://github.com/qvco/yaylib/blob/master/docs/API-Reference/chat/check_unread_status.md">check_unread_status()</a></code></td>
		<td>チャットの未読ステータスを確認します</td>
	</tr>
	<tr>
		<td><code><a href="https://github.com/qvco/yaylib/blob/master/docs/API-Reference/chat/create_group_chat.md">create_group_chat()</a></code></td>
		<td>グループチャットを作成します</td>
	</tr>
	<tr>
		<td><code><a href="https://github.com/qvco/yaylib/blob/master/docs/API-Reference/chat/create_private_chat.md">create_private_chat()</a></code></td>
		<td>個人チャットを作成します</td>
	</tr>
	<tr>
		<td><code><a href="https://github.com/qvco/yaylib/blob/master/docs/API-Reference/chat/delete_background.md">delete_background()</a></code></td>
		<td>チャットの背景画像を削除します</td>
	</tr>
	<tr>
		<td><code><a href="https://github.com/qvco/yaylib/blob/master/docs/API-Reference/chat/delete_message.md">delete_message()</a></code></td>
		<td>メッセージを削除します</td>
	</tr>
	<tr>
		<td><code><a href="https://github.com/qvco/yaylib/blob/master/docs/API-Reference/chat/edit_chat_room.md">edit_chat_room()</a></code></td>
		<td>チャットルームを編集します</td>
	</tr>
	<tr>
		<td><code><a href="https://github.com/qvco/yaylib/blob/master/docs/API-Reference/chat/get_chatable_users.md">get_chatable_users()</a></code></td>
		<td>チャット可能なユーザーを取得します</td>
	</tr>
	<tr>
		<td><code><a href="https://github.com/qvco/yaylib/blob/master/docs/API-Reference/chat/get_gifs_data.md">get_gifs_data()</a></code></td>
		<td>チャットルームのGIFデータを取得します</td>
	</tr>
	<tr>
		<td><code><a href="https://github.com/qvco/yaylib/blob/master/docs/API-Reference/chat/get_hidden_chat_rooms.md">get_hidden_chat_rooms()</a></code></td>
		<td>非表示のチャットルームを取得します</td>
	</tr>
	<tr>
		<td><code><a href="https://github.com/qvco/yaylib/blob/master/docs/API-Reference/chat/get_main_chat_rooms.md">get_main_chat_rooms()</a></code></td>
		<td>メインのチャットルームを取得します</td>
	</tr>
	<tr>
		<td><code><a href="https://github.com/qvco/yaylib/blob/master/docs/API-Reference/chat/get_messages.md">get_messages()</a></code></td>
		<td>メッセージを取得します</td>
	</tr>
	<tr>
		<td><code><a href="https://github.com/qvco/yaylib/blob/master/docs/API-Reference/chat/get_request_chat_rooms.md">get_request_chat_rooms()</a></code></td>
		<td>チャットリクエストを取得します</td>
	</tr>
	<tr>
		<td><code><a href="https://github.com/qvco/yaylib/blob/master/docs/API-Reference/chat/get_chat_room.md">get_chat_room()</a></code></td>
		<td>チャットルームの詳細を取得します</td>
	</tr>
	<tr>
		<td><code><a href="https://github.com/qvco/yaylib/blob/master/docs/API-Reference/chat/get_sticker_packs.md">get_sticker_packs()</a></code></td>
		<td>スタンプを取得します</td>
	</tr>
	<tr>
		<td><code><a href="https://github.com/qvco/yaylib/blob/master/docs/API-Reference/chat/get_chat_requests_count.md">get_chat_requests_count()</a></code></td>
		<td>チャットリクエストの数を取得します</td>
	</tr>
	<tr>
		<td><code><a href="https://github.com/qvco/yaylib/blob/master/docs/API-Reference/chat/hide_chat.md">hide_chat()</a></code></td>
		<td>チャットルームを非表示にします</td>
	</tr>
	<tr>
		<td><code><a href="https://github.com/qvco/yaylib/blob/master/docs/API-Reference/chat/invite_to_chat.md">invite_to_chat()</a></code></td>
		<td>チャットルームにユーザーを招待します</td>
	</tr>
	<tr>
		<td><code><a href="https://github.com/qvco/yaylib/blob/master/docs/API-Reference/chat/kick_users_from_chat.md">kick_users_from_chat()</a></code></td>
		<td>チャットルームからユーザーを追放します</td>
	</tr>
	<tr>
		<td><code><a href="https://github.com/qvco/yaylib/blob/master/docs/API-Reference/chat/kick_users_from_chat.md">kick_users_from_chat()</a></code></td>
		<td>チャットルームからユーザーを追放します</td>
	</tr>
	<tr>
		<td><code><a href="https://github.com/qvco/yaylib/blob/master/docs/API-Reference/chat/pin_chat.md">pin_chat()</a></code></td>
		<td>チャットルームをピン止めします</td>
	</tr>
	<tr>
		<td><code><a href="https://github.com/qvco/yaylib/blob/master/docs/API-Reference/chat/read_message.md">read_message()</a></code></td>
		<td>メッセージを既読にします</td>
	</tr>
	<tr>
		<td><code><a href="https://github.com/qvco/yaylib/blob/master/docs/API-Reference/chat/refresh_chat_rooms.md">refresh_chat_rooms()</a></code></td>
		<td>チャットルームを更新します</td>
	</tr>
	<tr>
		<td><code><a href="https://github.com/qvco/yaylib/blob/master/docs/API-Reference/chat/remove_chat_rooms.md">remove_chat_rooms()</a></code></td>
		<td>チャットルームを削除します</td>
	</tr>
	<tr>
		<td><code><a href="https://github.com/qvco/yaylib/blob/master/docs/API-Reference/chat/report_chat_room.md">report_chat_room()</a></code></td>
		<td>チャットルームを通報します</td>
	</tr>
	<tr>
		<td><code><a href="https://github.com/qvco/yaylib/blob/master/docs/API-Reference/chat/send_message.md">send_message()</a></code></td>
		<td>チャットルームにメッセージを送信します</td>
	</tr>
	<tr>
		<td><code><a href="https://github.com/qvco/yaylib/blob/master/docs/API-Reference/chat/unhide_chat.md">unhide_chat()</a></code></td>
		<td>チャットの非表示を解除します</td>
	</tr>
	<tr>
		<td><code><a href="https://github.com/qvco/yaylib/blob/master/docs/API-Reference/chat/unpin_chat.md">unpin_chat()</a></code></td>
		<td>チャットルームのピン止めを解除します</td>
	</tr>
</table>

### レター

<table>
    <tr>
		<th>メソッド</th>
		<th>概要</th>
	</tr>
	<tr>
		<td><code><a href="https://github.com/qvco/yaylib/blob/master/docs/API-Reference/review/create_review.md">create_review()</a></code></td>
		<td>レターを送信します</td>
	</tr>
	<tr>
		<td><code><a href="https://github.com/qvco/yaylib/blob/master/docs/API-Reference/review/create_reviews.md">create_reviews()</a></code></td>
		<td>複数人にレターを送信します</td>
	</tr>
	<tr>
		<td><code><a href="https://github.com/qvco/yaylib/blob/master/docs/API-Reference/review/delete_reviews.md">delete_reviews()</a></code></td>
		<td>レターを削除します</td>
	</tr>
	<tr>
		<td><code><a href="https://github.com/qvco/yaylib/blob/master/docs/API-Reference/review/get_my_reviews.md">get_my_reviews()</a></code></td>
		<td>送信したレターを取得します</td>
	</tr>
	<tr>
		<td><code><a href="https://github.com/qvco/yaylib/blob/master/docs/API-Reference/review/get_reviews.md">get_reviews()</a></code></td>
		<td>ユーザーが受け取ったレターを取得します</td>
	</tr>
	<tr>
		<td><code><a href="https://github.com/qvco/yaylib/blob/master/docs/API-Reference/review/pin_review.md">pin_review()</a></code></td>
		<td>レターをピンします</td>
	</tr>
	<tr>
		<td><code><a href="https://github.com/qvco/yaylib/blob/master/docs/API-Reference/review/unpin_review.md">unpin_review()</a></code></td>
		<td>レターのピン止めを解除します</td>
	</tr>
</table>

### 通話

<table>
    <tr>
		<th>メソッド</th>
		<th>概要</th>
	</tr>
	<tr>
		<td><code><a href="https://github.com/qvco/yaylib/blob/master/docs/API-Reference/call/get_user_active_call.md">get_user_active_call()</a></code></td>
		<td>ユーザーが参加中の通話を取得します</td>
	</tr>
	<tr>
		<td><code><a href="https://github.com/qvco/yaylib/blob/master/docs/API-Reference/call/get_bgms.md">get_bgms()</a></code></td>
		<td>通話のBGMを取得します</td>
	</tr>
	<tr>
		<td><code><a href="https://github.com/qvco/yaylib/blob/master/docs/API-Reference/call/get_call.md">get_call()</a></code></td>
		<td>通話の詳細を取得します</td>
	</tr>
	<tr>
		<td><code><a href="https://github.com/qvco/yaylib/blob/master/docs/API-Reference/call/get_call_invitable_users.md">get_call_invitable_users()</a></code></td>
		<td>通話に招待可能なユーザーを取得します</td>
	</tr>
	<tr>
		<td><code><a href="https://github.com/qvco/yaylib/blob/master/docs/API-Reference/call/get_call_status.md">get_call_status()</a></code></td>
		<td>通話の状態を取得します</td>
	</tr>
	<tr>
		<td><code><a href="https://github.com/qvco/yaylib/blob/master/docs/API-Reference/call/get_games.md">get_games()</a></code></td>
		<td>ゲームを取得します</td>
	</tr>
	<tr>
		<td><code><a href="https://github.com/qvco/yaylib/blob/master/docs/API-Reference/call/get_genres.md">get_genres()</a></code></td>
		<td>通話のジャンルを取得します</td>
	</tr>
	<tr>
		<td><code><a href="https://github.com/qvco/yaylib/blob/master/docs/API-Reference/call/get_group_calls.md">get_group_calls()</a></code></td>
		<td>サークルの通話を取得します</td>
	</tr>
	<tr>
		<td><code><a href="https://github.com/qvco/yaylib/blob/master/docs/API-Reference/call/invite_online_followings_to_call.md">invite_online_followings_to_call()</a></code></td>
		<td>オンラインの友達をまとめて通話に招待します</td>
	</tr>
	<tr>
		<td><code><a href="https://github.com/qvco/yaylib/blob/master/docs/API-Reference/call/invite_users_to_call.md">invite_users_to_call()</a></code></td>
		<td>ユーザーを通話に招待します</td>
	</tr>
	<tr>
		<td><code><a href="https://github.com/qvco/yaylib/blob/master/docs/API-Reference/call/invite_users_to_chat_call.md">invite_users_to_chat_call()</a></code></td>
		<td>チャット通話にユーザーを招待します</td>
	</tr>
	<tr>
		<td><code><a href="https://github.com/qvco/yaylib/blob/master/docs/API-Reference/call/kick_user_from_call.md">kick_user_from_call()</a></code></td>
		<td>ユーザーを通話からキックします</td>
	</tr>
	<tr>
		<td><code><a href="https://github.com/qvco/yaylib/blob/master/docs/API-Reference/call/set_call.md">set_call()</a></code></td>
		<td>通話を開始します</td>
	</tr>
	<tr>
		<td><code><a href="https://github.com/qvco/yaylib/blob/master/docs/API-Reference/call/set_user_role.md">set_user_role()</a></code></td>
		<td>通話に参加中ののユーザーに役職を与えます</td>
	</tr>
	<tr>
		<td><code><a href="https://github.com/qvco/yaylib/blob/master/docs/API-Reference/call/join_call.md">join_call()</a></code></td>
		<td>通話に参加します</td>
	</tr>
	<tr>
		<td><code><a href="https://github.com/qvco/yaylib/blob/master/docs/API-Reference/call/leave_call.md">leave_call()</a></code></td>
		<td>通話から退出します</td>
	</tr>
</table>

### ログイン

<table>
    <tr>
		<th>メソッド</th>
		<th>概要</th>
	</tr>
	<tr>
		<td><code><a href="https://github.com/qvco/yaylib/blob/master/docs/API-Reference/login/is_valid_token.md">is_valid_token()</a></code></td>
		<td>アクセストークンが有効か検証します</td>
	</tr>
	<tr>
		<td><code><a href="https://github.com/qvco/yaylib/blob/master/docs/API-Reference/login/change_password.md">change_password()</a></code></td>
		<td>パスワードを変更します</td>
	</tr>
	<tr>
		<td><code><a href="https://github.com/qvco/yaylib/blob/master/docs/API-Reference/login/get_token.md">get_token()</a></code></td>
		<td>トークンを再発行します</td>
	</tr>
	<tr>
		<td><code><a href="https://github.com/qvco/yaylib/blob/master/docs/API-Reference/login/login.md">login()</a></code></td>
		<td>メールアドレスでログインします</td>
	</tr>
	<tr>
		<td><code><a href="https://github.com/qvco/yaylib/blob/master/docs/API-Reference/login/logout.md">logout()</a></code></td>
		<td>ログアウトします</td>
	</tr>
	<tr>
		<td><code><a href="https://github.com/qvco/yaylib/blob/master/docs/API-Reference/login/resend_confirm_email.md">resend_confirm_email()</a></code></td>
		<td>確認メールを再送信します</td>
	</tr>
	<tr>
		<td><code><a href="https://github.com/qvco/yaylib/blob/master/docs/API-Reference/login/restore_user.md">restore_user()</a></code></td>
		<td>ユーザーを復元します</td>
	</tr>
	<tr>
		<td><code><a href="https://github.com/qvco/yaylib/blob/master/docs/API-Reference/login/revoke_tokens.md">revoke_tokens()</a></code></td>
		<td>トークンを無効化します</td>
	</tr>
	<tr>
		<td><code><a href="https://github.com/qvco/yaylib/blob/master/docs/API-Reference/login/save_account_with_email.md">save_account_with_email()</a></code></td>
		<td>メールアドレスでアカウントを保存します</td>
	</tr>
</table>

### 通知

<table>
    <tr>
		<th>メソッド</th>
		<th>概要</th>
	</tr>
	<tr>
		<td><code><a href="https://github.com/qvco/yaylib/blob/master/docs/API-Reference/cassandra/get_activities.md">get_activities()</a></code></td>
		<td>通知を取得します</td>
	</tr>
	<tr>
		<td><code><a href="https://github.com/qvco/yaylib/blob/master/docs/API-Reference/cassandra/get_merged_activities.md">get_merged_activities()</a></code></td>
		<td>全種類の通知を取得します</td>
	</tr>
	<tr>
		<td><code><a href="https://github.com/qvco/yaylib/blob/master/docs/API-Reference/cassandra/received_notification.md">received_notification()</a></code></td>
		<td>謎です。</td>
	</tr>
</table>

### その他

<table>
    <tr>
		<th>メソッド</th>
		<th>概要</th>
	</tr>
	<tr>
		<td><code><a href="https://github.com/qvco/yaylib/blob/master/docs/API-Reference/misc/get_email_grant_token.md">get_email_grant_token()</a></code></td>
		<td>email_grant_tokenを取得します</td>
	</tr>
	<tr>
		<td><code><a href="https://github.com/qvco/yaylib/blob/master/docs/API-Reference/misc/get_email_verification_presigned_url.md">get_email_verification_presigned_url()</a></code></td>
		<td>メールアドレス確認用の署名付きURLを取得します</td>
	</tr>
	<tr>
		<td><code><a href="https://github.com/qvco/yaylib/blob/master/docs/API-Reference/misc/get_file_upload_presigned_urls.md">get_file_upload_presigned_urls()</a></code></td>
		<td>ファイルアップロード用の署名付きURLを取得します</td>
	</tr>
	<tr>
		<td><code><a href="https://github.com/qvco/yaylib/blob/master/docs/API-Reference/misc/get_old_file_upload_presigned_url.md">get_old_file_upload_presigned_url()</a></code></td>
		<td>動画ファイルアップロード用の署名付きURLを取得します</td>
	</tr>
	<tr>
		<td><code><a href="https://github.com/qvco/yaylib/blob/master/docs/API-Reference/misc/get_web_socket_token.md">get_web_socket_token()</a></code></td>
		<td>Web Socket Tokenを取得します</td>
	</tr>
	<tr>
		<td><code><a href="https://github.com/qvco/yaylib/blob/master/docs/API-Reference/misc/upload_image.md">upload_image()</a></code></td>
		<td>画像をアップロードしてattachment_filenameを返します。</td>
	</tr>
</table>

<p align="right">(<a href="#readme-top">トップに戻る</a>)</p>

## オブジェクトモデル一覧

<table>
    <tr>
		<th>モデル名</th>
		<th>概要</th>
	</tr>
    <tr>
		<td><code><a href="https://github.com/qvco/yaylib/blob/master/docs/Object-Models/ActiveFollowingsResponse.md">ActiveFollowingsResponse</a></code></td>
		<td>オンラインのフォロー中のレスポンスモデル</td>
	</tr>
    <tr>
		<td><code><a href="https://github.com/qvco/yaylib/blob/master/docs/Object-Models/ActivitiesResponse.md">ActivitiesResponse</a></code></td>
		<td>通知データのレスポンスモデル</td>
	</tr>
	<tr>
		<td><code><a href="https://github.com/qvco/yaylib/blob/master/docs/Object-Models/Activity.md">Activity</a></code></td>
		<td>通知データのオブジェクトモデル</td>
	</tr>
    <tr>
		<td><code><a href="https://github.com/qvco/yaylib/blob/master/docs/Object-Models/Bgm.md">Bgm</a></code></td>
		<td>通話のBGMのオブジェクトモデル</td>
	</tr>
    <tr>
		<td><code><a href="https://github.com/qvco/yaylib/blob/master/docs/Object-Models/BlockedUserIdsResponse.md">BlockedUserIdsResponse</a></code></td>
		<td>ブロックされたユーザーIDのレスポンスモデル</td>
	</tr>
    <tr>
		<td><code><a href="https://github.com/qvco/yaylib/blob/master/docs/Object-Models/BlockedUsersResponse.md">BlockedUsersResponse</a></code></td>
		<td>ブロックしたユーザーのレスポンスモデル</td>
	</tr>
    <tr>
		<td><code><a href="https://github.com/qvco/yaylib/blob/master/docs/Object-Models/BookmarkPostResponse.md">BookmarkPostResponse</a></code></td>
		<td>ブックマーク投稿のレスポンスモデル</td>
	</tr>
    <tr>
		<td><code><a href="https://github.com/qvco/yaylib/blob/master/docs/Object-Models/CallStatusResponse.md">CallStatusResponse</a></code></td>
		<td>通話ステータスのレスポンスモデル</td>
	</tr>
    <tr>
		<td><code><a href="https://github.com/qvco/yaylib/blob/master/docs/Object-Models/ChatRoom.md">ChatRoom</a></code></td>
		<td>チャットルームのオブジェクトモデル</td>
	</tr>
    <tr>
		<td><code><a href="https://github.com/qvco/yaylib/blob/master/docs/Object-Models/ChatRoomsResponse.md">ChatRoomsResponse</a></code></td>
		<td>チャットルームのレスポンスモデル</td>
	</tr>
    <tr>
		<td><code><a href="https://github.com/qvco/yaylib/blob/master/docs/Object-Models/ConferenceCall.md">ConferenceCall</a></code></td>
		<td>通話のオブジェクトモデル</td>
	</tr>
    <tr>
		<td><code><a href="https://github.com/qvco/yaylib/blob/master/docs/Object-Models/CreateChatRoomResponse.md">CreateChatRoomResponse</a></code></td>
		<td>チャットルーム作成時のレスポンスモデル</td>
	</tr>
    <tr>
		<td><code><a href="https://github.com/qvco/yaylib/blob/master/docs/Object-Models/CreateGroupQuota.md">CreateGroupQuota</a></code></td>
		<td>グループ作成可能割当量のオブジェクトモデル</td>
	</tr>
    <tr>
		<td><code><a href="https://github.com/qvco/yaylib/blob/master/docs/Object-Models/CreateGroupResponse.md">CreateGroupResponse</a></code></td>
		<td>グループ作成時のレスポンスモデル</td>
	</tr>
    <tr>
		<td><code><a href="https://github.com/qvco/yaylib/blob/master/docs/Object-Models/FollowRecommendationsResponse.md">FollowRecommendationsResponse</a></code></td>
		<td>おすすめのユーザーのレスポンスモデル</td>
	</tr>
    <tr>
		<td><code><a href="https://github.com/qvco/yaylib/blob/master/docs/Object-Models/FollowUsersResponse.md">FollowUsersResponse</a></code></td>
		<td>フォローに関するユーザーのレスポンスモデル</td>
	</tr>
    <tr>
		<td><code><a href="https://github.com/qvco/yaylib/blob/master/docs/Object-Models/Footprint.md">Footprint</a></code></td>
		<td>足跡のオブジェクトモデル</td>
	</tr>
    <tr>
		<td><code><a href="https://github.com/qvco/yaylib/blob/master/docs/Object-Models/GamesResponse.md">GamesResponse</a></code></td>
		<td>ゲームデータのレスポンスモデル</td>
	</tr>
    <tr>
		<td><code><a href="https://github.com/qvco/yaylib/blob/master/docs/Object-Models/GenresResponse.md">GenresResponse</a></code></td>
		<td>通話のジャンルデータのレスポンスモデル</td>
	</tr>
    <tr>
		<td><code><a href="https://github.com/qvco/yaylib/blob/master/docs/Object-Models/GifImageCategory.md">GifImageCategory</a></code></td>
		<td>GIT画像カテゴリーのオブジェクトモデル</td>
	</tr>
    <tr>
		<td><code><a href="https://github.com/qvco/yaylib/blob/master/docs/Object-Models/GroupCategoriesResponse.md">GroupCategoriesResponse</a></code></td>
		<td>サークルカテゴリーのレスポンスモデル</td>
	</tr>
    <tr>
		<td><code><a href="https://github.com/qvco/yaylib/blob/master/docs/Object-Models/GroupResponse.md">GroupResponse</a></code></td>
		<td>サークルデータのレスポンスモデル</td>
	</tr>
    <tr>
		<td><code><a href="https://github.com/qvco/yaylib/blob/master/docs/Object-Models/GroupThreadListResponse.md">GroupThreadListResponse</a></code></td>
		<td>サークルのスレッド一覧のレスポンスモデル</td>
	</tr>
    <tr>
		<td><code><a href="https://github.com/qvco/yaylib/blob/master/docs/Object-Models/GroupUserResponse.md">GroupUserResponse</a></code></td>
		<td>サークルメンバーのレスポンスモデル</td>
	</tr>
    <tr>
		<td><code><a href="https://github.com/qvco/yaylib/blob/master/docs/Object-Models/GroupUsersResponse.md">GroupUsersResponse</a></code></td>
		<td>複数のサークルメンバーのレスポンスモデル</td>
	</tr>
    <tr>
		<td><code><a href="https://github.com/qvco/yaylib/blob/master/docs/Object-Models/GroupsRelatedResponse.md">GroupsRelatedResponse</a></code></td>
		<td>関連のあるサークルのレスポンスモデル</td>
	</tr>
    <tr>
		<td><code><a href="https://github.com/qvco/yaylib/blob/master/docs/Object-Models/GroupsResponse.md">GroupsResponse</a></code></td>
		<td>複数のサークルのレスポンスモデル</td>
	</tr>
    <tr>
		<td><code><a href="https://github.com/qvco/yaylib/blob/master/docs/Object-Models/HiddenResponse.md">HiddenResponse</a></code></td>
		<td>非表示のレスポンスモデル</td>
	</tr>
    <tr>
		<td><code><a href="https://github.com/qvco/yaylib/blob/master/docs/Object-Models/LikePostsResponse.md">LikePostsResponse</a></code></td>
		<td>投稿にいいねしたときのレスポンスモデル</td>
	</tr>
    <tr>
		<td><code><a href="https://github.com/qvco/yaylib/blob/master/docs/Object-Models/LoginUpdateResponse.md">LoginUpdateResponse</a></code></td>
		<td>認証情報を更新したときのレスポンスモデル</td>
	</tr>
    <tr>
		<td><code><a href="https://github.com/qvco/yaylib/blob/master/docs/Object-Models/LoginUserResponse.md">LoginUserResponse</a></code></td>
		<td>認証情報のレスポンスモデル</td>
	</tr>
    <tr>
		<td><code><a href="https://github.com/qvco/yaylib/blob/master/docs/Object-Models/Message.md">Message</a></code></td>
		<td>メッセージのオブジェクトモデル</td>
	</tr>
    <tr>
		<td><code><a href="https://github.com/qvco/yaylib/blob/master/docs/Object-Models/MessageResponse.md">MessageResponse</a></code></td>
		<td>メッセージのレスポンスモデル</td>
	</tr>
    <tr>
		<td><code><a href="https://github.com/qvco/yaylib/blob/master/docs/Object-Models/Post.md">Post</a></code></td>
		<td>投稿のオブジェクトモデル</td>
	</tr>
    <tr>
		<td><code><a href="https://github.com/qvco/yaylib/blob/master/docs/Object-Models/PostLikersResponse.md">PostLikersResponse</a></code></td>
		<td>投稿にいいねしたユーザーのレスポンスモデル</td>
	</tr>
    <tr>
		<td><code><a href="https://github.com/qvco/yaylib/blob/master/docs/Object-Models/PostsResponse.md">PostsResponse</a></code></td>
		<td>複数の投稿のレスポンスモデル</td>
	</tr>
    <tr>
		<td><code><a href="https://github.com/qvco/yaylib/blob/master/docs/Object-Models/PresignedUrl.md">PresignedUrl</a></code></td>
		<td>署名付きURLのオブジェクトモデル</td>
	</tr>
    <tr>
		<td><code><a href="https://github.com/qvco/yaylib/blob/master/docs/Object-Models/RankingUsersResponse.md">RankingUsersResponse</a></code></td>
		<td>フォロワーランキンングのユーザーレスポンスモデル</td>
	</tr>
    <tr>
		<td><code><a href="https://github.com/qvco/yaylib/blob/master/docs/Object-Models/RefreshCounterRequestsResponse.md">RefreshCounterRequestsResponse</a></code></td>
		<td>表示数を更新をリクエストしたときのレスポンスモデル</td>
	</tr>
    <tr>
		<td><code><a href="https://github.com/qvco/yaylib/blob/master/docs/Object-Models/ReviewsResponse.md">ReviewsResponse</a></code></td>
		<td>複数のレターのレスポンスモデル</td>
	</tr>
    <tr>
		<td><code><a href="https://github.com/qvco/yaylib/blob/master/docs/Object-Models/SharedUrl.md">SharedUrl</a></code></td>
		<td>共有URLのオブジェクトモデル</td>
	</tr>
    <tr>
		<td><code><a href="https://github.com/qvco/yaylib/blob/master/docs/Object-Models/SocialShareUsersResponse.md">SocialShareUsersResponse</a></code></td>
		<td>複数のソーシャルシェアユーザーのレスポンスモデル</td>
	</tr>
    <tr>
		<td><code><a href="https://github.com/qvco/yaylib/blob/master/docs/Object-Models/StickerPack.md">StickerPack</a></code></td>
		<td>スタンプのオブジェクトモデル</td>
	</tr>
    <tr>
		<td><code><a href="https://github.com/qvco/yaylib/blob/master/docs/Object-Models/Survey.md">Survey</a></code></td>
		<td>投票のオブジェクトモデル</td>
	</tr>
    <tr>
		<td><code><a href="https://github.com/qvco/yaylib/blob/master/docs/Object-Models/ThreadInfo.md">ThreadInfo</a></code></td>
		<td>スレッド情報のオブジェクトモデル</td>
	</tr>
    <tr>
		<td><code><a href="https://github.com/qvco/yaylib/blob/master/docs/Object-Models/TokenResponse.md">TokenResponse</a></code></td>
		<td>トークンのレスポンスモデル</td>
	</tr>
    <tr>
		<td><code><a href="https://github.com/qvco/yaylib/blob/master/docs/Object-Models/UnreadStatusResponse.md">UnreadStatusResponse</a></code></td>
		<td>未読ステータスのレスポンスモデル</td>
	</tr>
    <tr>
		<td><code><a href="https://github.com/qvco/yaylib/blob/master/docs/Object-Models/User.md">User</a></code></td>
		<td>ユーザーのオブジェクトモデル</td>
	</tr>
    <tr>
		<td><code><a href="https://github.com/qvco/yaylib/blob/master/docs/Object-Models/UserResponse.md">UserResponse</a></code></td>
		<td>ユーザーのレスポンスモデル</td>
	</tr>
    <tr>
		<td><code><a href="https://github.com/qvco/yaylib/blob/master/docs/Object-Models/UserTimestampResponse.md">UserTimestampResponse</a></code></td>
		<td>タイムスタンプ付随のユーザーのレスポンスモデル</td>
	</tr>
    <tr>
		<td><code><a href="https://github.com/qvco/yaylib/blob/master/docs/Object-Models/UserWrapper.md">UserWrapper</a></code></td>
		<td>ユーザーラッパーのオブジェクトモデル</td>
	</tr>
    <tr>
		<td><code><a href="https://github.com/qvco/yaylib/blob/master/docs/Object-Models/UsersByTimestampResponse.md">UsersByTimestampResponse</a></code></td>
		<td>タイムスタンプ付随の複数のユーザーのレスポンスモデル</td>
	</tr>
    <tr>
		<td><code><a href="https://github.com/qvco/yaylib/blob/master/docs/Object-Models/UsersResponse.md">UsersResponse.md</a></code></td>
		<td>複数のユーザーのレスポンスモデル</td>
	</tr>
</table>

<p align="right">(<a href="#readme-top">トップに戻る</a>)</p>

## Yay! エラーコード一覧

<!-- ここではエラーのコードと概要を紹介します。 -->

<!-- 参照先 <https://yay.space/locales/ja/translation.json?v=3.18.0> -->

<table>
    <tr>
		<th>コード</th>
		<th>エラーの原因</th>
	</tr>
    <tr>
		<td><code>unknown</code></td>
		<td>原因不明</td>
	</tr>
    <tr>
		<td><code>-1</code></td>
		<td>パラメターが不正です</td>
	</tr>
    <tr>
		<td><code>-2</code></td>
		<td>このアカウントはすでに登録されています</td>
	</tr>
    <tr>
		<td><code>-3</code></td>
		<td>アクセストークンの有効期限切れ</td>
	</tr>
    <tr>
		<td><code>-4</code></td>
		<td>このIDはすでに使われています</td>
	</tr>
    <tr>
		<td><code>-5</code></td>
		<td>ユーザーが見つかりません</td>
	</tr>
    <tr>
		<td><code>-6</code></td>
		<td>投稿が見つかりません</td>
	</tr>
    <tr>
		<td><code>-7</code></td>
		<td>チャットルームが見つかりません</td>
	</tr>
    <tr>
		<td><code>-8</code></td>
		<td>メッセージが見つかりません</td>
	</tr>
    <tr>
		<td><code>-9</code></td>
		<td>チャットルームに特定のユーザーが見つかりません</td>
	</tr>
    <tr>
		<td><code>-10</code></td>
		<td>チャットルーム内のユーザーは2人以上の必要があります</td>
	</tr>
    <tr>
		<td><code>-11</code></td>
		<td>正しいパスワードを確認してください</td>
	</tr>
    <tr>
		<td><code>-12</code></td>
		<td>このユーザーはブロック中です</td>
	</tr>
    <tr>
		<td><code>-13</code></td>
		<td>プライベートユーザーです</td>
	</tr>
    <tr>
		<td><code>-14</code></td>
		<td>アプリケーションが見つかりません</td>
	</tr>
    <tr>
		<td><code>-15</code></td>
		<td>アカウント認証に失敗しました</td>
	</tr>
    <tr>
		<td><code>-16</code></td>
		<td>このSNSアカウントはすでに使われています</td>
	</tr>
    <tr>
		<td><code>-17</code></td>
		<td>SNSの連携を解除できません</td>
	</tr>
    <tr>
		<td><code>-18</code></td>
		<td>アクセストークンが不正です</td>
	</tr>
    <tr>
		<td><code>-19</code></td>
		<td>スポットが見つかりません</td>
	</tr>
    <tr>
		<td><code>-20</code></td>
		<td>このアカウントは永久的に凍結されました</td>
	</tr>
    <tr>
		<td><code>-21</code></td>
		<td>サービスの規約に抵触したため、アカウントを停止しています</td>
	</tr>
    <tr>
		<td><code>-22</code></td>
		<td>学校の情報が変更されています</td>
	</tr>
    <tr>
		<td><code>-26</code></td>
		<td>アカウントを作成して3日以内に削除はできません</td>
	</tr>
    <tr>
		<td><code>-29</code></td>
		<td>Captcha認証が必要です</td>
	</tr>
    <tr>
		<td><code>-30</code></td>
		<td>Captcha認証に失敗しました</td>
	</tr>
    <tr>
		<td><code>-100</code></td>
		<td>サークルが満員です</td>
	</tr>
    <tr>
		<td><code>-103</code></td>
		<td>このサークルから追放されています</td>
	</tr>
    <tr>
		<td><code>-200</code></td>
		<td>現在のパスワードを入力してください</td>
	</tr>
    <tr>
		<td><code>-201</code></td>
		<td>現在のパスワードを確認してください</td>
	</tr>
    <tr>
		<td><code>-202</code></td>
		<td>メールアドレス、もしくはパスワードが不正です</td>
	</tr>
    <tr>
		<td><code>-203</code></td>
		<td>このメールアドレスはすでに使われています</td>
	</tr>
    <tr>
		<td><code>-204</code></td>
		<td>スパムメールの可能性があり、操作を完了できません</td>
	</tr>
    <tr>
		<td><code>-308</code></td>
		<td>グルチャの最大人数は50人です</td>
	</tr>
    <tr>
		<td><code>-309</code></td>
		<td>通話が満員です</td>
	</tr>
    <tr>
		<td><code>-310</code></td>
		<td>通話は終了しました</td>
	</tr>
    <tr>
		<td><code>-312</code></td>
		<td>サークルの管理人にブロックされています</td>
	</tr>
    <tr>
		<td><code>-313</code></td>
		<td>チャットするにはフォローされる必要があります</td>
	</tr>
    <tr>
		<td><code>-315</code></td>
		<td>通話はロックされています</td>
	</tr>
    <tr>
		<td><code>-317</code></td>
		<td>枠主をフォローすることで参加できます</td>
	</tr>
    <tr>
		<td><code>-319</code></td>
		<td>正しいメールアドレスを確認してください</td>
	</tr>
    <tr>
		<td><code>-320</code></td>
		<td>このメールアドレスはすでに登録されています</td>
	</tr>
    <tr>
		<td><code>-321</code></td>
		<td>通話に参加できません。永久退出処分を受けています。</td>
	</tr>
    <tr>
		<td><code>-322</code></td>
		<td>通話のオーナーではありません</td>
	</tr>
    <tr>
		<td><code>-326</code></td>
		<td>VIPユーザーではありません</td>
	</tr>
    <tr>
		<td><code>-331</code></td>
		<td>ブロック数の上限に達しました</td>
	</tr>
    <tr>
		<td><code>-332</code></td>
		<td>認証コードが不正です</td>
	</tr>
    <tr>
		<td><code>-333</code></td>
		<td>認証コードの有効期限切れ</td>
	</tr>
    <tr>
		<td><code>-334</code></td>
		<td>アプリをアップデートしてください</td>
	</tr>
    <tr>
		<td><code>-335</code></td>
		<td>この機能を使用するためには電話番号認証が必要です</td>
	</tr>
    <tr>
		<td><code>-336</code></td>
		<td>制限に達したので、これ以上フォローすることができません。この制限はフォローワーを増やすことで、増やすことができます。</td>
	</tr>
    <tr>
		<td><code>-338</code></td>
		<td>レターを送れません。年齢が離れすぎています</td>
	</tr>
    <tr>
		<td><code>-339</code></td>
		<td>サークル管理人か副管理人のみ可能な操作です</td>
	</tr>
    <tr>
		<td><code>-340</code></td>
		<td>利用規約に基づき、現在アカウントを登録できません</td>
	</tr>
    <tr>
		<td><code>-342</code></td>
		<td>SnsShareRewardAlreadyBeenClaimed</td>
	</tr>
    <tr>
		<td><code>-343</code></td>
		<td>この機能の上限回数に達しました。1時間ほど時間を置いて再度お試しください。</td>
	</tr>
    <tr>
		<td><code>-346</code></td>
		<td>チャット送信に年齢確認が必要です</td>
	</tr>
    <tr>
		<td><code>-347</code></td>
		<td>このサークルに参加するには年齢確認をする必要があります</td>
	</tr>
    <tr>
		<td><code>-348</code></td>
		<td>チャットをするには電話番号認証が必要です</td>
	</tr>
    <tr>
		<td><code>-350</code></td>
		<td>編集は投稿の作成者のみ可能です</td>
	</tr>
    <tr>
		<td><code>-352</code></td>
		<td>特定の年齢層のみ参加が許可されているサークルです</td>
	</tr>
    <tr>
		<td><code>-355</code></td>
		<td>認証コードの送信回数が上限を越えました。1時間ほど時間をおいて再度お試しください</td>
	</tr>
    <tr>
		<td><code>-356</code></td>
		<td>チャットをするには電話番号認証が必要です</td>
	</tr>
    <tr>
		<td><code>-357</code></td>
		<td>サークルの招待は承諾されています</td>
	</tr>
    <tr>
		<td><code>-358</code></td>
		<td>サークルの招待は拒否されています</td>
	</tr>
    <tr>
		<td><code>-360</code></td>
		<td>IPがBanされました</td>
	</tr>
    <tr>
		<td><code>-361</code></td>
		<td>Twitterに接続されていません</td>
	</tr>
    <tr>
		<td><code>-363</code></td>
		<td>フォロワーにのみ投稿を公開しています</td>
	</tr>
    <tr>
		<td><code>-364</code></td>
		<td>カウンター更新の回数制限に達しました</td>
	</tr>
    <tr>
		<td><code>-367</code></td>
		<td>このユーザーがあなたをフォローする必要があります</td>
	</tr>
    <tr>
		<td><code>-369</code></td>
		<td>国設定は一度設定すると8時間は変更できません。時間をおいて再度お試しください</td>
	</tr>
    <tr>
		<td><code>-370</code></td>
		<td>サークルメンバーでないため、この通話に参加できません</td>
	</tr>
    <tr>
		<td><code>-371</code></td>
		<td>グループ保留中の移動</td>
	</tr>
    <tr>
		<td><code>-372</code></td>
		<td>グループ保留中の代理指名</td>
	</tr>
    <tr>
		<td><code>-373</code></td>
		<td>相手は危険なユーザーとのチャットを許可していません</td>
	</tr>
    <tr>
		<td><code>-374</code></td>
		<td>近日ペナルティなどを受けているため、新規のユーザーとのチャットルーム作成が制限されています</td>
	</tr>
    <tr>
		<td><code>-375</code></td>
		<td>このユーザーの投稿は(´∀｀∩)↑age↑できません</td>
	</tr>
    <tr>
		<td><code>-376</code></td>
		<td>これ以上アカウントを作成することはできません</td>
	</tr>
    <tr>
		<td><code>-377</code></td>
		<td>特定の性別のみ参加が許可されているサークルです</td>
	</tr>
    <tr>
		<td><code>-378</code></td>
		<td>このサークルを作成するには性別の設定が必要です</td>
	</tr>
    <tr>
		<td><code>-382</code></td>
		<td>サークルの追加上限に到達しました</td>
	</tr>
    <tr>
		<td><code>-383</code></td>
		<td>これ以上ピン留めできません。解除してからもう一度お試しください</td>
	</tr>
    <tr>
		<td><code>-384</code></td>
		<td>Twitterでのグループ共有の制限を超えています</td>
	</tr>
    <tr>
		<td><code>-385</code></td>
		<td>通報を受けているため、この投稿は(´∀｀∩)↑age↑できません</td>
	</tr>
    <!-- <tr>
		<td><code>-000</code></td>
		<td>InsufficientCoins</td>
	</tr> -->
    <tr>
		<td><code>-402</code></td>
		<td>この通話に参加するには枠主と相互にフォローする必要があります</td>
	</tr>
    <tr>
		<td><code>-403</code></td>
		<td>投稿の編集上限に達しています。これ以上の編集はできません</td>
	</tr>
    <tr>
		<td><code>-404</code></td>
		<td>サークルにおける招待の許容数を超えました。時間をあけてから再度行ってください。</td>
	</tr>
    <tr>
		<td><code>-405</code></td>
		<td>参加するには電話番号認証が必要です。Yay! アプリから電話番号認証を行なってください。</td>
	</tr>
    <tr>
		<td><code>-406</code></td>
		<td>一定期間の経過後、投稿の編集はできません</td>
	</tr>
    <tr>
		<td><code>-407</code></td>
		<td>パスワードが短すぎます</td>
	</tr>
    <tr>
		<td><code>-408</code></td>
		<td>パスワードが長すぎます</td>
	</tr>
    <tr>
		<td><code>-409</code></td>
		<td>他のパスワードを使用してください。無効な文字列が含まれています</td>
	</tr>
    <tr>
		<td><code>-410</code></td>
		<td>他のパスワードを使用してください。文字列や数字、または記号などの組み合わせをお試しください</td>
	</tr>
    <tr>
		<td><code>-411</code></td>
		<td>このメールアドレスを使用するには承認される必要があります</td>
	</tr>
    <tr>
		<td><code>-412</code></td>
		<td>投稿をスレッドに移動できません</td>
	</tr>
    <tr>
		<td><code>-413</code></td>
		<td>URLを投稿できません</td>
	</tr>
    <tr>
		<td><code>-977</code></td>
		<td>通話を開始できませんでした</td>
	</tr>
    <!-- <tr>
		<td><code>-000</code></td>
		<td>InvalidAppVersion</td>
	</tr> -->
    <tr>
		<td><code>-1000</code></td>
		<td>電話番号がBanされています</td>
	</tr>
</table>
