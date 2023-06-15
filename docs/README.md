<div><a id="readme-top"></a></div>

# yaylib - API ドキュメント | Docs

「yaylib」を使用すると、独自かつ高度な方法でプログラムを使用して Yay!にアクセスできます。投稿、個人チャット、サークル、通話、ユーザーなど、Yay!の主要な機能を活用しましょう。

<details open>
    <summary>目録</summary>
    <ol>
        <li><a href="#インストール">インストール</a></li>
        <li><a href="#使用方法">使用方法</a></li>
        <!-- <li><a href="async-client">Async Client</a></li> -->
        <li><a href="#api-一覧">API 一覧</a></li>
        <li><a href="#モデル">モデル</a></li>
        <li><a href="#レスポンス">レスポンス</a></li>
        <li><a href="#yay-エラーコード一覧">Yay! エラーコード一覧</a></li>
    </ol>
</details>

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

## 使用方法

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
		<td><code>refresh_token</code></td>
		<td><code>str</code></td>
		<td><code>None</code></td>
		<td>アクセストークンの再発行に必要なリフレッシュトークン</td>
	</tr>
    <tr>
		<td><code>proxy</code></td>
		<td><code>str</code></td>
		<td><code>None</code></td>
		<td>プロキシサーバーのアドレス</td>
	</tr>
    <tr>
		<td><code>timeout</code></td>
		<td><code>int</code></td>
		<td><code>60</code></td>
		<td>通信接続を待機する時間</td>
	</tr>
    <tr>
		<td><code>base_path</code></td>
		<td><code>str</code></td>
		<td><code>作業ディレクトリ</code></td>
		<td>ベースパス</td>
	</tr>
    <tr>
		<td><code>lang</code></td>
		<td><code>str</code></td>
		<td><code>ja</code></td>
		<td>ログ出力などの言語（ja もしくは、en）</td>
	</tr>
    <tr>
		<td><code>loglevel_stream</code></td>
		<td><code>int</code></td>
		<td><code>logging.INFO</code></td>
		<td>ログレベルの設定</td>
	</tr>
    <tr>
		<td><code>host</code></td>
		<td><code>str</code></td>
		<td><code>api.yay.space</code></td>
		<td>yay.spaceのプロダクションホスト</td>
	</tr>
</table>

### 使用例

① 以下の例を使用して、API クライアントを初期化します。

```python
import yaylib

api = yaylib.Client(access_token="[アクセストークン]")
```

② アクセストークンはログイン時に発行されます。

```python
import yaylib

api = yaylib.Client()
api.login(email="[メールアドレス]", password="[パスワード]")

api.login_data.access_token # アクセストークンを取得
api.login_data.refresh_token # リフレッシュトークンを取得（アクセストークン再発行時に使用）
```

一定時間内に何度もログインすると、**429 Too Many Requests エラー**が発生します。
よって、初回実行するときは、② の方法でアクセストークンを取得し、以後は ① のように**アクセストークンを使用してクライアントを初期化すること**をお勧めします。

※ アクセストークンには有効期限が設けられており、24 時間が経過する前に<a href="#get-token">`get_token()`</a>関数を使用するか、ログインして再発行してください。

<p align="right">(<a href="#readme-top">トップに戻る</a>)</p>

## API 一覧

<ul>
    <li><a href="#ログイン-api">ログイン</a></li>
    <li><a href="#ユーザー-api">ユーザー</a></li>
    <li><a href="#投稿-api">投稿</a></li>
    <li><a href="#スレッド-api">スレッド</a></li>
    <li><a href="#レター-api">レター</a></li>
    <li><a href="#チャット-api">チャット</a></li>
    <li><a href="#グループ-api">グループ</a></li>
    <li><a href="#通話-api">通話</a></li>
    <li><a href="#通知-api">通知</a></li>
    <li><a href="#その他-api">その他</a></li>
</ul>

### ログイン API

<table>
    <tr>
		<th>メソッド</th>
		<th>引数</th>
		<th>型</th>
		<th>返り値</th>
		<th>説明</th>
	</tr>
    <!-- change_email -->
    <tr>
		<td rowspan="3"><code>change_email()</code></td>
		<td>email</td>
		<td><code>str</code></td>
		<td rowspan="3"><code><a href="#loginupdateresponse">LoginUpdateResponse</a></code></td>
		<td rowspan="3">メールアドレスを変更します</td>
	</tr>
	<tr>
		<td>password</td>
		<td><code>str</code></td>
	</tr>
	<tr>
		<td>email_grant_token</td>
		<td><code>str</code></td>
	</tr>
    <!-- change_password -->
    <tr>
		<td rowspan="2"><code>change_password()</code></td>
		<td>current_password</td>
		<td><code>str</code></td>
		<td rowspan="2"><code><a href="#loginupdateresponse">LoginUpdateResponse</a></code></td>
		<td rowspan="2">パスワードを変更します</td>
	</tr>
	<tr>
		<td>new_password</td>
		<td><code>str</code></td>
	</tr>
    <!-- get_token -->
	<tr id="get-token">
		<td rowspan="4"><code>get_token()</code></td>
		<td>grant_type</td>
		<td><code>str</code></td>
        <td rowspan="4"><code><a href="#tokenresponse">TokenResponse</a></code></td>
		<td rowspan="4">
        アクセストークンを再発行します。<br>
        ※ <code>grant_type</code>が"refresh_token"の場合は、emailとpasswordは必要ありません。
        </td>
	</tr>
	<tr>
		<td>refresh_token</td>
		<td><code>str</code></td>
	</tr>
	<tr>
		<td>email</td>
		<td><code>str</code></td>
	</tr>
	<tr>
		<td>password</td>
		<td><code>str</code></td>
	</tr>
    <!-- login -->
    <tr>
		<td rowspan="2"><code>login()</code></td>
		<td>email</td>
		<td><code>str</code></td>
		<td rowspan="2"><code><a href="#loginuserresponse">LoginUserResponse</a></code></td>
		<td rowspan="2">メールアドレスでログインします</td>
	</tr>
	<tr>
		<td>password</td>
		<td><code>str</code></td>
	</tr>
    <!-- logout -->
    <tr>
		<td><code>logout()</code></td>
		<td></td>
		<td></td>
		<td><code>dict</code></td>
		<td>ログアウトします</td>
	</tr>
    <!-- resend_confirm_email -->
    <tr>
		<td><code>resend_confirm_email()</code></td>
		<td></td>
		<td></td>
		<td><code>dict</code></td>
		<td>確認メールを再送信します</td>
	</tr>
    <!-- restore_user -->
    <tr>
		<td><code>restore_user()</code></td>
		<td>user_id</td>
		<td><code>int</code></td>
		<td><code><a href="#loginuserresponse">LoginUserResponse</a></code></td>
		<td>ユーザーを復元します</td>
	</tr>
    <!-- revoke_tokens -->
    <tr>
		<td><code>revoke_tokens()</code></td>
		<td></td>
		<td></td>
		<td><code>dict</code></td>
		<td>トークンを無効化します</td>
	</tr>
    <!-- save_account_with_email -->
    <tr>
		<td rowspan="4"><code>save_account_with_email()</code></td>
		<td>email</td>
		<td><code>str</code></td>
		<td rowspan="4"><code><a href="#loginupdateresponse">LoginUpdateResponse</a></code></td>
		<td rowspan="4">メールアドレスでアカウントを保存します</td>
	</tr>
	<tr>
		<td>password</td>
		<td><code>str</code></td>
	</tr>
	<tr>
		<td>current_password</td>
		<td><code>str</code></td>
	</tr>
	<tr>
		<td>email_grant_token</td>
		<td><code>str</code></td>
	</tr>
</table>

<p align="right">(<a href="#readme-top">トップに戻る</a>)</p>

### ユーザー API

<p align="right">(<a href="#readme-top">トップに戻る</a>)</p>

### 投稿 API

<table>
    <tr>
		<th>メソッド</th>
		<th>引数</th>
		<th>型</th>
		<th>返り値</th>
		<th>説明</th>
	</tr>
    <!-- add_bookmark -->
    <tr>
		<td rowspan="2"><code>add_bookmark()</code></td>
		<td>user_id</td>
		<td><code>int</code></td>
		<td rowspan="2"><code><a href="#bookmarkpostresponse">BookmarkPostResponse</a></code></td>
		<td rowspan="2">ブックマークに追加します</td>
	</tr>
	<tr>
		<td>post_id</td>
		<td><code>int</code></td>
	</tr>
    <!-- add_group_highlight_post -->
    <tr>
		<td rowspan="2"><code>add_group_highlight_post()</code></td>
		<td>group_id</td>
		<td><code>int</code></td>
		<td rowspan="2"><code>dict</code></td>
		<td rowspan="2">投稿をグループハイライトに追加します</td>
	</tr>
	<tr>
		<td>post_id</td>
		<td><code>int</code></td>
	</tr>
    <!-- create_call_post -->
	<tr>
		<td rowspan="18"><code>create_call_post()</code></td>
		<td>text</td>
		<td><code>str</code></td>
        <td rowspan="18"><code><a href="#conferencecall">ConferenceCall</a></code></td>
		<td rowspan="18">通話の投稿を作成します</td>
	</tr>
	<tr>
		<td>font_size</td>
		<td><code>int</code></td>
	</tr>
	<tr>
		<td>color</td>
		<td><code>int</code></td>
	</tr>
	<tr>
		<td>group_id</td>
		<td><code>int</code></td>
	</tr>
	<tr>
		<td>call_type</td>
		<td><code>str</code></td>
	</tr>
	<tr>
		<td>category_id</td>
		<td><code>int</code></td>
	</tr>
	<tr>
		<td>game_title</td>
		<td><code>str</code></td>
	</tr>
	<tr>
		<td>joinable_by</td>
		<td><code>str</code></td>
	</tr>
	<tr>
		<td>message_tags</td>
		<td><code>str</code></td>
	</tr>
	<tr>
		<td>attachment_filename</td>
		<td><code>str</code></td>
	</tr>
	<tr>
		<td>attachment_2_filename</td>
		<td><code>str</code></td>
	</tr>
	<tr>
		<td>attachment_3_filename</td>
		<td><code>str</code></td>
	</tr>
	<tr>
		<td>attachment_4_filename</td>
		<td><code>str</code></td>
	</tr>
	<tr>
		<td>attachment_5_filename</td>
		<td><code>str</code></td>
	</tr>
	<tr>
		<td>attachment_6_filename</td>
		<td><code>str</code></td>
	</tr>
	<tr>
		<td>attachment_7_filename</td>
		<td><code>str</code></td>
	</tr>
	<tr>
		<td>attachment_8_filename</td>
		<td><code>str</code></td>
	</tr>
	<tr>
		<td>attachment_9_filename</td>
		<td><code>str</code></td>
	</tr>
    <!-- pin_group_post -->
    <tr>
		<td rowspan="2"><code>pin_group_post()</code></td>
		<td>post_id</td>
		<td><code>int</code></td>
		<td rowspan="2"><code>dict</code></td>
		<td rowspan="2">グループの投稿をピンします</td>
	</tr>
	<tr>
		<td>group_id</td>
		<td><code>int</code></td>
	</tr>
    <!-- pin_post -->
    <tr>
		<td><code>pin_post()</code></td>
		<td>post_id</td>
		<td><code>int</code></td>
		<td><code>dict</code></td>
		<td>投稿をピンします</td>
	</tr>
    <!-- mention -->
    <tr>
		<td><code>mention()</code></td>
		<td>user_id</td>
		<td><code>int</code></td>
		<td><code>str</code></td>
		<td>メンション用文字列を返します</td>
	</tr>
    <!-- create_post -->
	<tr>
		<td rowspan="19"><code>create_post()</code></td>
		<td>text</td>
		<td><code>str</code></td>
        <td rowspan="19"><code><a href="#post">Post</a></code></td>
		<td rowspan="19">投稿を作成します</td>
	</tr>
	<tr>
		<td>font_size</td>
		<td><code>int</code></td>
	</tr>
	<tr>
		<td>color</td>
		<td><code>int</code></td>
	</tr>
	<tr>
		<td>in_reply_to</td>
		<td><code>int</code></td>
	</tr>
	<tr>
		<td>group_id</td>
		<td><code>int</code></td>
	</tr>
	<tr>
		<td>mention_ids</td>
		<td><code>list of int</code></td>
	</tr>
	<tr>
		<td>choices</td>
		<td><code>list of str</code></td>
	</tr>
	<tr>
		<td>shared_url</td>
		<td><code>str</code></td>
	</tr>
	<tr>
		<td>message_tags</td>
		<td><code>str</code></td>
	</tr>
	<tr>
		<td>attachment_filename</td>
		<td><code>str</code></td>
	</tr>
	<tr>
		<td>attachment_2_filename</td>
		<td><code>str</code></td>
	</tr>
	<tr>
		<td>attachment_3_filename</td>
		<td><code>str</code></td>
	</tr>
	<tr>
		<td>attachment_4_filename</td>
		<td><code>str</code></td>
	</tr>
	<tr>
		<td>attachment_5_filename</td>
		<td><code>str</code></td>
	</tr>
	<tr>
		<td>attachment_6_filename</td>
		<td><code>str</code></td>
	</tr>
	<tr>
		<td>attachment_7_filename</td>
		<td><code>str</code></td>
	</tr>
	<tr>
		<td>attachment_8_filename</td>
		<td><code>str</code></td>
	</tr>
	<tr>
		<td>attachment_9_filename</td>
		<td><code>str</code></td>
	</tr>
	<tr>
		<td>video_file_name</td>
		<td><code>str</code></td>
	</tr>
    <!-- create_repost -->
	<tr>
		<td rowspan="20"><code>create_repost()</code></td>
		<td>post_id</td>
		<td><code>int</code></td>
        <td rowspan="20"><code><a href="#post">Post</a></code></td>
		<td rowspan="20">投稿を(´∀｀∩)↑age↑します</td>
	</tr>
	<tr>
		<td>text</td>
		<td><code>str</code></td>
	</tr>
	<tr>
		<td>font_size</td>
		<td><code>int</code></td>
	</tr>
	<tr>
		<td>color</td>
		<td><code>int</code></td>
	</tr>
	<tr>
		<td>in_reply_to</td>
		<td><code>int</code></td>
	</tr>
	<tr>
		<td>group_id</td>
		<td><code>int</code></td>
	</tr>
	<tr>
		<td>mention_ids</td>
		<td><code>list of int</code></td>
	</tr>
	<tr>
		<td>choices</td>
		<td><code>list of str</code></td>
	</tr>
	<tr>
		<td>shared_url</td>
		<td><code>str</code></td>
	</tr>
	<tr>
		<td>message_tags</td>
		<td><code>str</code></td>
	</tr>
	<tr>
		<td>attachment_filename</td>
		<td><code>str</code></td>
	</tr>
	<tr>
		<td>attachment_2_filename</td>
		<td><code>str</code></td>
	</tr>
	<tr>
		<td>attachment_3_filename</td>
		<td><code>str</code></td>
	</tr>
	<tr>
		<td>attachment_4_filename</td>
		<td><code>str</code></td>
	</tr>
	<tr>
		<td>attachment_5_filename</td>
		<td><code>str</code></td>
	</tr>
	<tr>
		<td>attachment_6_filename</td>
		<td><code>str</code></td>
	</tr>
	<tr>
		<td>attachment_7_filename</td>
		<td><code>str</code></td>
	</tr>
	<tr>
		<td>attachment_8_filename</td>
		<td><code>str</code></td>
	</tr>
	<tr>
		<td>attachment_9_filename</td>
		<td><code>str</code></td>
	</tr>
	<tr>
		<td>video_file_name</td>
		<td><code>str</code></td>
	</tr>
        <!-- create_share_post -->
	<tr>
		<td rowspan="6"><code>create_share_post()</code></td>
		<td>shareable_type</td>
		<td><code>str</code></td>
        <td rowspan="6"><code><a href="#post">Post</a></code></td>
		<td rowspan="6">シェア投稿を作成します</td>
	</tr>
	<tr>
		<td>shareable_id</td>
		<td><code>int</code></td>
	</tr>
	<tr>
		<td>text</td>
		<td><code>str</code></td>
	</tr>
	<tr>
		<td>font_size</td>
		<td><code>int</code></td>
	</tr>
	<tr>
		<td>color</td>
		<td><code>int</code></td>
	</tr>
	<tr>
		<td>group_id</td>
		<td><code>int</code></td>
	</tr>
    <!-- create_thread_post -->
	<tr>
		<td rowspan="20"><code>create_thread_post()</code></td>
		<td>post_id</td>
		<td><code>int</code></td>
        <td rowspan="20"><code><a href="#post">Post</a></code></td>
		<td rowspan="20">スレッドの投稿を作成します</td>
	</tr>
	<tr>
		<td>text</td>
		<td><code>str</code></td>
	</tr>
	<tr>
		<td>font_size</td>
		<td><code>int</code></td>
	</tr>
	<tr>
		<td>color</td>
		<td><code>int</code></td>
	</tr>
	<tr>
		<td>in_reply_to</td>
		<td><code>int</code></td>
	</tr>
	<tr>
		<td>group_id</td>
		<td><code>int</code></td>
	</tr>
	<tr>
		<td>mention_ids</td>
		<td><code>list of int</code></td>
	</tr>
	<tr>
		<td>choices</td>
		<td><code>list of str</code></td>
	</tr>
	<tr>
		<td>shared_url</td>
		<td><code>str</code></td>
	</tr>
	<tr>
		<td>message_tags</td>
		<td><code>str</code></td>
	</tr>
	<tr>
		<td>attachment_filename</td>
		<td><code>str</code></td>
	</tr>
	<tr>
		<td>attachment_2_filename</td>
		<td><code>str</code></td>
	</tr>
	<tr>
		<td>attachment_3_filename</td>
		<td><code>str</code></td>
	</tr>
	<tr>
		<td>attachment_4_filename</td>
		<td><code>str</code></td>
	</tr>
	<tr>
		<td>attachment_5_filename</td>
		<td><code>str</code></td>
	</tr>
	<tr>
		<td>attachment_6_filename</td>
		<td><code>str</code></td>
	</tr>
	<tr>
		<td>attachment_7_filename</td>
		<td><code>str</code></td>
	</tr>
	<tr>
		<td>attachment_8_filename</td>
		<td><code>str</code></td>
	</tr>
	<tr>
		<td>attachment_9_filename</td>
		<td><code>str</code></td>
	</tr>
	<tr>
		<td>video_file_name</td>
		<td><code>str</code></td>
	</tr>
    <!-- delete_all_post -->
    <tr>
		<td><code>delete_all_post()</code></td>
		<td></td>
		<td></td>
		<td><code>dict</code></td>
		<td>すべての投稿を削除します</td>
	</tr>
    <!-- unpin_group_post -->
    <tr>
		<td><code>unpin_group_post()</code></td>
		<td>group_id</td>
		<td><code>int</code></td>
		<td><code>dict</code></td>
		<td>グループのピン投稿を解除します</td>
	</tr>
    <!-- unpin_post -->
    <tr>
		<td><code>unpin_post()</code></td>
		<td>post_id</td>
		<td><code>int</code></td>
		<td><code>dict</code></td>
		<td>ピン投稿を削除します</td>
	</tr>
    <!-- get_bookmark -->
    <tr>
		<td rowspan="2"><code>get_bookmark()</code></td>
		<td>user_id</td>
		<td><code>int</code></td>
		<td rowspan="2"><code><a href="#postsresponse">PostsResponse</a></code></td>
		<td rowspan="2">ブックマークを取得します</td>
	</tr>
	<tr>
		<td>from_str</td>
		<td><code>str</code></td>
	</tr>
    <!-- get_timeline_calls -->
    <tr>
		<td rowspan="9"><code>get_timeline_calls()</code></td>
		<td>group_id</td>
		<td><code>int</code></td>
		<td rowspan="9"><code><a href="#postsresponse">PostsResponse</a></code></td>
		<td rowspan="9">誰でも通話を取得します</td>
	</tr>
	<tr>
		<td>from_timestamp</td>
		<td><code>int</code></td>
	</tr>
	<tr>
		<td>number</td>
		<td><code>int</code></td>
	</tr>
	<tr>
		<td>category_id</td>
		<td><code>int</code></td>
	</tr>
	<tr>
		<td>call_type</td>
		<td><code>str</code></td>
	</tr>
	<tr>
		<td>include_circle_call</td>
		<td><code>bool</code></td>
	</tr>
	<tr>
		<td>cross_generation</td>
		<td><code>bool</code></td>
	</tr>
	<tr>
		<td>exclude_recent_gomimushi</td>
		<td><code>bool</code></td>
	</tr>
	<tr>
		<td>shared_interest_categories</td>
		<td><code>bool</code></td>
	</tr>
    <!-- get_conversation -->
    <tr>
		<td rowspan="6"><code>get_conversation()</code></td>
		<td>conversation_id</td>
		<td><code>int</code></td>
		<td rowspan="6"><code><a href="#postsresponse">PostsResponse</a></code></td>
		<td rowspan="6">会話を取得します</td>
	</tr>
	<tr>
		<td>group_id</td>
		<td><code>int</code></td>
	</tr>
	<tr>
		<td>thread_id</td>
		<td><code>int</code></td>
	</tr>
	<tr>
		<td>from_post_id</td>
		<td><code>int</code></td>
	</tr>
	<tr>
		<td>number</td>
		<td><code>int</code></td>
	</tr>
	<tr>
		<td>reverse</td>
		<td><code>bool</code></td>
	</tr>
    <!-- get_conversation_root_posts -->
    <tr>
		<td><code>get_conversation_root_posts()</code></td>
		<td>post_ids</td>
		<td><code>list of int</code></td>
		<td><code><a href="#postsresponse">PostsResponse</a></code></td>
		<td>会話の原点の投稿を取得します</td>
	</tr>
    <!-- get_following_call_timeline -->
    <tr>
		<td rowspan="6"><code>get_following_call_timeline()</code></td>
		<td>from_timestamp</td>
		<td><code>int</code></td>
		<td rowspan="6"><code><a href="#postsresponse">PostsResponse</a></code></td>
		<td rowspan="6">フォロー中の通話を取得します</td>
	</tr>
	<tr>
		<td>number</td>
		<td><code>int</code></td>
	</tr>
	<tr>
		<td>category_id</td>
		<td><code>int</code></td>
	</tr>
	<tr>
		<td>call_type</td>
		<td><code>str</code></td>
	</tr>
	<tr>
		<td>include_circle_call</td>
		<td><code>bool</code></td>
	</tr>
	<tr>
		<td>exclude_recent_gomimushi</td>
		<td><code>bool</code></td>
	</tr>
    <!-- get_following_timeline -->
    <tr>
		<td rowspan="7"><code>get_following_timeline()</code></td>
		<td>from_str</td>
		<td><code>str</code></td>
		<td rowspan="7"><code><a href="#postsresponse">PostsResponse</a></code></td>
		<td rowspan="7">フォロー中のタイムラインを取得します</td>
	</tr>
	<tr>
		<td>only_root</td>
		<td><code>bool</code></td>
	</tr>
	<tr>
		<td>order_by</td>
		<td><code>str</code></td>
	</tr>
	<tr>
		<td>number</td>
		<td><code>int</code></td>
	</tr>
	<tr>
		<td>mxn</td>
		<td><code>int</code></td>
	</tr>
	<tr>
		<td>reduce_selfie</td>
		<td><code>bool</code></td>
	</tr>
	<tr>
		<td>custom_generation_range</td>
		<td><code>bool</code></td>
	</tr>
    <!-- get_group_highlight_posts -->
    <tr>
		<td rowspan="3"><code>get_group_highlight_posts()</code></td>
		<td>group_id</td>
		<td><code>int</code></td>
		<td rowspan="3"><code><a href="#postsresponse">PostsResponse</a></code></td>
		<td rowspan="3">グループのハイライト投稿を取得します</td>
	</tr>
	<tr>
		<td>from_post</td>
		<td><code>int</code></td>
	</tr>
	<tr>
		<td>number</td>
		<td><code>int</code></td>
	</tr>
    <!-- get_group_timeline_by_keyword -->
    <tr>
		<td rowspan="5"><code>get_group_timeline_by_keyword()</code></td>
		<td>group_id</td>
		<td><code>int</code></td>
		<td rowspan="5"><code><a href="#postsresponse">PostsResponse</a></code></td>
		<td rowspan="5">グループの投稿をキーワードで検索します</td>
	</tr>
	<tr>
		<td>keyword</td>
		<td><code>str</code></td>
	</tr>
	<tr>
		<td>from_post_id</td>
		<td><code>int</code></td>
	</tr>
	<tr>
		<td>number</td>
		<td><code>int</code></td>
	</tr>
	<tr>
		<td>only_thread_posts</td>
		<td><code>bool</code></td>
	</tr>
    <!-- get_group_timeline -->
    <tr>
		<td rowspan="6"><code>get_group_timeline()</code></td>
		<td>group_id</td>
		<td><code>int</code></td>
		<td rowspan="6"><code><a href="#postsresponse">PostsResponse</a></code></td>
		<td rowspan="6">グループのタイムラインを取得します</td>
	</tr>
	<tr>
		<td>from_post_id</td>
		<td><code>int</code></td>
	</tr>
	<tr>
		<td>reverse</td>
		<td><code>bool</code></td>
	</tr>
	<tr>
		<td>post_type</td>
		<td><code>str</code></td>
	</tr>
	<tr>
		<td>number</td>
		<td><code>int</code></td>
	</tr>
	<tr>
		<td>only_root</td>
		<td><code>bool</code></td>
	</tr>
    <!-- get_timeline -->
    <tr>
		<td rowspan="11"><code>get_timeline()</code></td>
		<td>noreply_mode</td>
		<td><code>bool</code></td>
		<td rowspan="11"><code><a href="#postsresponse">PostsResponse</a></code></td>
		<td rowspan="11">タイムラインを取得します</td>
	</tr>
	<tr>
		<td>from_post_id</td>
		<td><code>int</code></td>
	</tr>
	<tr>
		<td>number</td>
		<td><code>int</code></td>
	</tr>
	<tr>
		<td>order_by</td>
		<td><code>str</code></td>
	</tr>
	<tr>
		<td>experiment_older_age_rules</td>
		<td><code>bool</code></td>
	</tr>
	<tr>
		<td>shared_interest_categories</td>
		<td><code>bool</code></td>
	</tr>
	<tr>
		<td>mxn</td>
		<td><code>int</code></td>
	</tr>
	<tr>
		<td>en</td>
		<td><code>int</code></td>
	</tr>
	<tr>
		<td>vn</td>
		<td><code>int</code></td>
	</tr>
	<tr>
		<td>reduce_selfie</td>
		<td><code>bool</code></td>
	</tr>
	<tr>
		<td>custom_generation_range</td>
		<td><code>bool</code></td>
	</tr>
    <!-- get_timeline_by_hashtag -->
    <tr>
		<td rowspan="3"><code>get_timeline_by_hashtag()</code></td>
		<td>hashtag</td>
		<td><code>str</code></td>
		<td rowspan="3"><code><a href="#postsresponse">PostsResponse</a></code></td>
		<td rowspan="3">ハッシュタグでタイムラインを検索します</td>
	</tr>
	<tr>
		<td>from_post_id</td>
		<td><code>int</code></td>
	</tr>
	<tr>
		<td>number</td>
		<td><code>int</code></td>
	</tr>
    <!-- get_timeline_by_keyword -->
    <tr>
		<td rowspan="3"><code>get_timeline_by_keyword()</code></td>
		<td>keyword</td>
		<td><code>str</code></td>
		<td rowspan="3"><code><a href="#postsresponse">PostsResponse</a></code></td>
		<td rowspan="3">キーワードでタイムラインを検索します</td>
	</tr>
	<tr>
		<td>from_post_id</td>
		<td><code>int</code></td>
	</tr>
	<tr>
		<td>number</td>
		<td><code>int</code></td>
	</tr>
    <!-- get_user_timeline -->
    <tr>
		<td rowspan="3"><code>get_user_timeline()</code></td>
		<td>from_post_id</td>
		<td><code>int</code></td>
		<td rowspan="3"><code><a href="#postsresponse">PostsResponse</a></code></td>
		<td rowspan="3">ユーザーのタイムラインを取得します</td>
	</tr>
	<tr>
		<td>number</td>
		<td><code>int</code></td>
	</tr>
    <tr>
		<td>post_type</td>
		<td><code>str</code></td>
	</tr>
    <!-- get_my_posts -->
    <tr>
		<td rowspan="3"><code>get_my_posts()</code></td>
		<td>from_post_id</td>
		<td><code>int</code></td>
		<td rowspan="3"><code><a href="#postsresponse">PostsResponse</a></code></td>
		<td rowspan="3">自分の投稿を取得します</td>
	</tr>
	<tr>
		<td>number</td>
		<td><code>int</code></td>
	</tr>
	<tr>
		<td>include_group_post</td>
		<td><code>bool</code></td>
	</tr>
    <!-- get_post -->
    <tr>
		<td><code>get_post()</code></td>
		<td>post_id</td>
		<td><code>int</code></td>
		<td><code><a href="#post">Post</a></code></td>
		<td>投稿の詳細を取得します</td>
	</tr>
    <!-- get_posts -->
    <tr>
		<td><code>get_posts()</code></td>
		<td>post_ids</td>
		<td><code>list of int</code></td>
		<td><code><a href="#postspesponse">PostsResponse</a></code></td>
		<td>複数の投稿を取得します</td>
	</tr>
    <!-- get_post_likers -->
    <tr>
		<td rowspan="2"><code>get_post_likers()</code></td>
		<td>from_id</td>
		<td><code>int</code></td>
		<td rowspan="2"><code><a href="#postlikersresponse">PostLikersResponse</a></code></td>
		<td rowspan="2">投稿にいいねしたユーザーを取得します</td>
	</tr>
	<tr>
		<td>number</td>
		<td><code>int</code></td>
	</tr>
    <!-- get_reposts -->
    <tr>
		<td rowspan="3"><code>get_reposts()</code></td>
		<td>post_id</td>
		<td><code>int</code></td>
		<td rowspan="3"><code><a href="#postsresponse">PostsResponse</a></code></td>
		<td rowspan="3">投稿の(´∀｀∩)↑age↑を取得します</td>
	</tr>
	<tr>
		<td>from_post_id</td>
		<td><code>int</code></td>
	</tr>
	<tr>
		<td>number</td>
		<td><code>int</code></td>
	</tr>
    <!-- get_recommended_post_tags -->
    <tr>
		<td rowspan="2"><code>get_recommended_post_tags()</code></td>
		<td>tag</td>
		<td><code>str</code></td>
		<td rowspan="2"><code><a href="#posttagsresponse">PostTagsResponse</a></code></td>
		<td rowspan="2">おすすめのタグ候補を取得します</td>
	</tr>
	<tr>
		<td>save_recent_search</td>
		<td><code>bool</code></td>
	</tr>
    <!-- get_recommended_posts -->
    <tr>
		<td rowspan="3"><code>get_recommended_posts()</code></td>
		<td>experiment_num</td>
		<td><code>int</code></td>
		<td rowspan="3"><code><a href="#postsresponse">PostsResponse</a></code></td>
		<td rowspan="3">おすすめの投稿を取得します</td>
	</tr>
	<tr>
		<td>variant_num</td>
		<td><code>int</code></td>
	</tr>
	<tr>
		<td>number</td>
		<td><code>int</code></td>
	</tr>
    <!-- get_url_metadata -->
    <tr>
		<td><code>get_url_metadata()</code></td>
		<td>url</td>
		<td><code>str</code></td>
		<td><code><a href="#sharedurl">SharedUrl</a></code></td>
		<td>URLのメタデータを取得します</td>
	</tr>
    <!-- like_posts -->
    <tr>
		<td><code>like_posts()</code></td>
		<td>post_ids</td>
		<td><code>list of int</code></td>
		<td><code><a href="#likepostsresponse">LikePostsResponse</a></code></td>
		<td>投稿にいいねします</td>
	</tr>
    <!-- unlike_post -->
    <tr>
		<td><code>unlike_post()</code></td>
		<td>post_id</td>
		<td><code>int</code></td>
		<td><code>dict</code></td>
		<td>投稿のいいねを解除します</td>
	</tr>
    <!-- remove_bookmark -->
    <tr>
		<td rowspan="2"><code>remove_bookmark()</code></td>
		<td>user_id</td>
		<td><code>int</code></td>
		<td rowspan="2"><code>dict</code></td>
		<td rowspan="2">ブックマークを削除します</td>
	</tr>
	<tr>
		<td>post_id</td>
		<td><code>int</code></td>
	</tr>
    <!-- remove_group_highlight_post -->
    <tr>
		<td rowspan="2"><code>remove_group_highlight_post()</code></td>
		<td>group_id</td>
		<td><code>int</code></td>
		<td rowspan="2"><code>dict</code></td>
		<td rowspan="2">サークルのハイライト投稿を解除します</td>
	</tr>
	<tr>
		<td>post_id</td>
		<td><code>int</code></td>
	</tr>
    <!-- remove_posts -->
    <tr>
		<td><code>remove_posts()</code></td>
		<td>post_ids</td>
		<td><code>list of int</code></td>
		<td><code>dict</code></td>
		<td>複数の投稿を削除します</td>
	</tr>
    <!-- report_post -->
    <tr>
		<td rowspan="8"><code>report_post()</code></td>
		<td>post_id</td>
		<td><code>int</code></td>
		<td rowspan="8"><code>dict</code></td>
		<td rowspan="8">投稿を通報します</td>
	</tr>
	<tr>
		<td>opponent_id</td>
		<td><code>int</code></td>
	</tr>
	<tr>
		<td>category_id</td>
		<td><code>int</code></td>
	</tr>
	<tr>
		<td>reason</td>
		<td><code>str</code></td>
	</tr>
	<tr>
		<td>screenshot_filename</td>
		<td><code>str</code></td>
	</tr>
	<tr>
		<td>screenshot_2_filename</td>
		<td><code>str</code></td>
	</tr>
	<tr>
		<td>screenshot_3_filename</td>
		<td><code>str</code></td>
	</tr>
	<tr>
		<td>screenshot_4_filename</td>
		<td><code>str</code></td>
	</tr>
    <!-- update_post -->
    <tr>
		<td rowspan="5"><code>update_post()</code></td>
		<td>post_id</td>
		<td><code>int</code></td>
		<td rowspan="5"><code><a href="#post">Post</a></code></td>
		<td rowspan="5">投稿を編集します</td>
	</tr>
	<tr>
		<td>text</td>
		<td><code>str</code></td>
	</tr>
	<tr>
		<td>font_size</td>
		<td><code>int</code></td>
	</tr>
	<tr>
		<td>color</td>
		<td><code>int</code></td>
	</tr>
	<tr>
		<td>message_tags</td>
		<td><code>str</code></td>
	</tr>
    <!-- update_recommendation_feedback -->
    <tr>
		<td rowspan="4"><code>update_recommendation_feedback()</code></td>
		<td>post_id</td>
		<td><code>int</code></td>
		<td rowspan="4"><code>dict</code></td>
		<td rowspan="4">おすすめのフィードバックを更新します</td>
	</tr>
	<tr>
		<td>feedback_result</td>
		<td><code>str</code></td>
	</tr>
	<tr>
		<td>experiment_num</td>
		<td><code>int</code></td>
	</tr>
	<tr>
		<td>variant_num</td>
		<td><code>int</code></td>
	</tr>
    <!-- validate_post -->
    <tr>
		<td rowspan="3"><code>validate_post()</code></td>
		<td>text</td>
		<td><code>str</code></td>
		<td rowspan="3"><code><a href="#validationpostresponse">ValidationPostResponse</a></code></td>
		<td rowspan="3">与えられたテキストが有効な投稿であるかどうかを検証します</td>
	</tr>
	<tr>
		<td>group_id</td>
		<td><code>int</code></td>
	</tr>
	<tr>
		<td>thread_id</td>
		<td><code>int</code></td>
	</tr>
    <!-- view_video -->
    <tr>
		<td><code>view_video()</code></td>
		<td>video_id</td>
		<td><code>int</code></td>
		<td><code>dict</code></td>
		<td>動画を視聴します</td>
	</tr>
    <!-- vote_survey -->
    <tr>
		<td rowspan="2"><code>vote_survey()</code></td>
		<td>survey_id</td>
		<td><code>int</code></td>
		<td rowspan="2"><code><a href="#survey">Survey</a></code></td>
		<td rowspan="2">アンケートに投票します</td>
	</tr>
	<tr>
		<td>choice_id</td>
		<td><code>int</code></td>
	</tr>
</table>

<p align="right">(<a href="#readme-top">トップに戻る</a>)</p>

### スレッド API

<p align="right">(<a href="#readme-top">トップに戻る</a>)</p>

### レター API

<p align="right">(<a href="#readme-top">トップに戻る</a>)</p>

### チャット API

<p align="right">(<a href="#readme-top">トップに戻る</a>)</p>

### グループ API

<p align="right">(<a href="#readme-top">トップに戻る</a>)</p>

### 通話 API

<p align="right">(<a href="#readme-top">トップに戻る</a>)</p>

### 通知 API

<p align="right">(<a href="#readme-top">トップに戻る</a>)</p>

### その他 API

<p align="right">(<a href="#readme-top">トップに戻る</a>)</p>

## モデル

<ul>
    <li><a href="#activity">Activity</a></li>
    <li><a href="#banword">BanWord</a></li>
    <li><a href="#bgm">Bgm</a></li>
    <li><a href="#callgifthistory">CallGiftHistory</a></li>
    <li><a href="#chatroom">ChatRoom</a></li>
    <li><a href="#chatroomdraft">ChatRoomDraft</a></li>
    <li><a href="#choice">Choice</a></li>
    <li><a href="#coinamount">CoinAmount</a></li>
    <li><a href="#coinexpiration">CoinExpiration</a></li>
    <li><a href="#coinproduct">CoinProduct</a></li>
    <li><a href="#coinproductquota">CoinProductQuota</a></li>
    <li><a href="#conferencecall">ConferenceCall</a></li>
    <li><a href="#conferencecalluserrole">ConferenceCallUserRole</a></li>
    <li><a href="#contactstatus">ContactStatus</a></li>
</ul>

### Activity

通知データのオブジェクトモデル

<table>
    <tr>
		<th>フィールド</th>
		<th>型</th>
		<th>概要</th>
	</tr>
    <tr>
		<td><code>data</code></td>
		<td><code>dict</code></td>
		<td>レスポンス</td>
	</tr>
    <tr>
		<td><code>id</code></td>
		<td><code>int</code></td>
		<td>通知のID</td>
	</tr>
    <tr>
		<td><code>created_at</code></td>
		<td><code>int</code></td>
		<td>タイムスタンブ</td>
	</tr>
    <tr>
		<td><code>created_at_parsed</code></td>
		<td><code>str</code></td>
		<td>パースされたタイムスタンブ</td>
	</tr>
    <tr>
		<td><code>type</code></td>
		<td><code>str</code></td>
		<td>通知のタイプ</td>
	</tr>
    <tr>
		<td><code>user</code></td>
		<td><code><a href="#user">User</a></code></td>
		<td>ユーザーの詳細</td>
	</tr>
    <tr>
		<td><code>from_post</code></td>
		<td><code><a href="#post">Post</a></code></td>
		<td>投稿の詳細</td>
	</tr>
    <tr>
		<td><code>to_post</code></td>
		<td><code><a href="#post">Post</a></code></td>
		<td>投稿の詳細</td>
	</tr>
    <tr>
		<td><code>group</code></td>
		<td><code><a href="#group">Group</a></code></td>
		<td>サークルの詳細</td>
	</tr>
    <tr>
		<td><code>followers</code></td>
		<td><code>list of <a href="#user">User</a></code></td>
		<td>ユーザーの詳細のリスト</td>
	</tr>
    <tr>
		<td><code>from_post_ids</code></td>
		<td><code>list of int</code></td>
		<td>投稿IDのリスト</td>
	</tr>
    <tr>
		<td><code>vip_reward</code></td>
		<td><code>int</code></td>
		<td>VIPのリワード</td>
	</tr>
    <tr>
		<td><code>is_bulk_invitation</code></td>
		<td><code>bool</code></td>
		<td>通話への招待かどうか</td>
	</tr>
</table>

### BanWord

禁止ワードのオブジェクトモデル

<table>
    <tr>
		<th>フィールド</th>
		<th>型</th>
		<th>概要</th>
	</tr>
    <tr>
		<td><code>data</code></td>
		<td><code>dict</code></td>
		<td>レスポンス</td>
	</tr>
    <tr>
		<td><code>id</code></td>
		<td><code>int</code></td>
		<td>ID</td>
	</tr>
    <tr>
		<td><code>type</code></td>
		<td><code>str</code></td>
		<td>タイプ</td>
	</tr>
    <tr>
		<td><code>word</code></td>
		<td><code>str</code></td>
		<td>単語</td>
	</tr>
</table>

### Bgm

Bgm のオブジェクトモデル

<table>
    <tr>
		<th>フィールド</th>
		<th>型</th>
		<th>概要</th>
	</tr>
    <tr>
		<td><code>data</code></td>
		<td><code>dict</code></td>
		<td>レスポンス</td>
	</tr>
    <tr>
		<td><code>id</code></td>
		<td><code>int</code></td>
		<td>ID</td>
	</tr>
    <tr>
		<td><code>title</code></td>
		<td><code>str</code></td>
		<td>タイトル</td>
	</tr>
    <tr>
		<td><code>music_url</code></td>
		<td><code>str</code></td>
		<td></td>
	</tr>
    <tr>
		<td><code>order</code></td>
		<td><code>int</code></td>
		<td></td>
	</tr>
</table>

### CallGiftHistory

通話のギフト履歴のオブジェクトモデル

<table>
    <tr>
		<th>フィールド</th>
		<th>型</th>
		<th>概要</th>
	</tr>
    <tr>
		<td><code>data</code></td>
		<td><code>dict</code></td>
		<td>レスポンス</td>
	</tr>
    <tr>
		<td><code>gifts_count</code></td>
		<td><code>list of <a href="#giftcount">GiftCount</a></code></td>
		<td></td>
	</tr>
    <tr>
		<td><code>sent_at</code></td>
		<td><code>int</code></td>
		<td></td>
	</tr>
    <tr>
		<td><code>sent_at_parsed</code></td>
		<td><code>str</code></td>
		<td></td>
	</tr>
    <tr>
		<td><code>sender</code></td>
		<td><code><a href="#user">User</a></code></td>
		<td></td>
	</tr>
</table>

### ChatRoom

個人チャットのオブジェクトモデル

<table>
    <tr>
		<th>フィールド</th>
		<th>型</th>
		<th>概要</th>
	</tr>
    <tr>
		<td><code>data</code></td>
		<td><code>dict</code></td>
		<td>レスポンス</td>
	</tr>
    <tr>
		<td><code>id</code></td>
		<td><code>int</code></td>
		<td></td>
	</tr>
    <tr>
		<td><code>unread_count</code></td>
		<td><code>int</code></td>
		<td></td>
	</tr>
    <tr>
		<td><code>updated_at</code></td>
		<td><code>int</code></td>
		<td></td>
	</tr>
    <tr>
		<td><code>updated_at_parsed</code></td>
		<td><code>int</code></td>
		<td></td>
	</tr>
    <tr>
		<td><code>members</code></td>
		<td><code>list of <a href="#user">User</a></code></td>
		<td></td>
	</tr>
    <tr>
		<td><code>background</code></td>
		<td><code>str</code></td>
		<td></td>
	</tr>
    <tr>
		<td><code>last_message</code></td>
		<td><code><a href="#message">Message</a></code></td>
		<td></td>
	</tr>
    <tr>
		<td><code>name</code></td>
		<td><code>str</code></td>
		<td></td>
	</tr>
    <tr>
		<td><code>is_group</code></td>
		<td><code>bool</code></td>
		<td></td>
	</tr>
    <tr>
		<td><code>owner</code></td>
		<td><code><a href="#user">User</a></code></td>
		<td></td>
	</tr>
    <tr>
		<td><code>is_request</code></td>
		<td><code>bool</code></td>
		<td></td>
	</tr>
    <tr>
		<td><code>user_setting</code></td>
		<td><code><a href="#usersetting">UserSetting</a></code></td>
		<td></td>
	</tr>
</table>

### ChatRoomDraft

保留中の個人チャットのオブジェクトモデル

<table>
    <tr>
		<th>フィールド</th>
		<th>型</th>
		<th>概要</th>
	</tr>
    <tr>
		<td><code>data</code></td>
		<td><code>dict</code></td>
		<td>レスポンス</td>
	</tr>
    <tr>
		<td><code>id</code></td>
		<td><code>int</code></td>
		<td></td>
	</tr>
    <tr>
		<td><code>text</code></td>
		<td><code>str</code></td>
		<td></td>
	</tr>
</table>

### Choice

アンケート選択肢のオブジェクトモデル

<table>
    <tr>
		<th>フィールド</th>
		<th>型</th>
		<th>概要</th>
	</tr>
    <tr>
		<td><code>data</code></td>
		<td><code>dict</code></td>
		<td>レスポンス</td>
	</tr>
    <tr>
		<td><code>id</code></td>
		<td><code>int</code></td>
		<td></td>
	</tr>
    <tr>
		<td><code>label</code></td>
		<td><code>str</code></td>
		<td></td>
	</tr>
    <tr>
		<td><code>votes_count</code></td>
		<td><code>int</code></td>
		<td></td>
	</tr>
</table>

### CoinAmount

コインアマウントのオブジェクトモデル

<table>
    <tr>
		<th>フィールド</th>
		<th>型</th>
		<th>概要</th>
	</tr>
    <tr>
		<td><code>data</code></td>
		<td><code>dict</code></td>
		<td>レスポンス</td>
	</tr>
    <tr>
		<td><code>paid</code></td>
		<td><code>float</code></td>
		<td></td>
	</tr>
    <tr>
		<td><code>free</code></td>
		<td><code>float</code></td>
		<td></td>
	</tr>
    <tr>
		<td><code>total</code></td>
		<td><code>float</code></td>
		<td></td>
	</tr>
</table>

### CoinExpiration

コインの期限のオブジェクトモデル

<table>
    <tr>
		<th>フィールド</th>
		<th>型</th>
		<th>概要</th>
	</tr>
    <tr>
		<td><code>data</code></td>
		<td><code>dict</code></td>
		<td>レスポンス</td>
	</tr>
    <tr>
		<td><code>expired_at</code></td>
		<td><code>int</code></td>
		<td></td>
	</tr>
    <tr>
		<td><code>amount</code></td>
		<td><code>float</code></td>
		<td></td>
	</tr>
</table>

### CoinProduct

コインプロダクトのオブジェクトモデル

<table>
    <tr>
		<th>フィールド</th>
		<th>型</th>
		<th>概要</th>
	</tr>
    <tr>
		<td><code>data</code></td>
		<td><code>dict</code></td>
		<td>レスポンス</td>
	</tr>
    <tr>
		<td><code>id</code></td>
		<td><code>int</code></td>
		<td></td>
	</tr>
    <tr>
		<td><code>purchasable</code></td>
		<td><code>bool</code></td>
		<td></td>
	</tr>
    <tr>
		<td><code>amount</code></td>
		<td><code>float</code></td>
		<td></td>
	</tr>
</table>

### CoinProductQuota

コインプロダクトの割当量のオブジェクトモデル

<table>
    <tr>
		<th>フィールド</th>
		<th>型</th>
		<th>概要</th>
	</tr>
    <tr>
		<td><code>data</code></td>
		<td><code>dict</code></td>
		<td>レスポンス</td>
	</tr>
    <tr>
		<td><code>bought</code></td>
		<td><code>float</code></td>
		<td></td>
	</tr>
    <tr>
		<td><code>limit</code></td>
		<td><code>float</code></td>
		<td></td>
	</tr>
</table>

### ConferenceCall

通話投稿のオブジェクトモデル

<table>
    <tr>
		<th>フィールド</th>
		<th>型</th>
		<th>概要</th>
	</tr>
    <tr>
		<td><code>data</code></td>
		<td><code>dict</code></td>
		<td>レスポンス</td>
	</tr>
    <tr>
		<td><code>id</code></td>
		<td><code>int</code></td>
		<td></td>
	</tr>
    <tr>
		<td><code>post_id</code></td>
		<td><code>int</code></td>
		<td></td>
	</tr>
    <tr>
		<td><code>group_id</code></td>
		<td><code>int</code></td>
		<td></td>
	</tr>
    <tr>
		<td><code>is_active</code></td>
		<td><code>bool</code></td>
		<td></td>
	</tr>
    <tr>
		<td><code>anonymous_call_users_count</code></td>
		<td><code>int</code></td>
		<td></td>
	</tr>
    <tr>
		<td><code>agora_channel</code></td>
		<td><code>str</code></td>
		<td></td>
	</tr>
    <tr>
		<td><code>agora_token</code></td>
		<td><code>str</code></td>
		<td></td>
	</tr>
    <tr>
		<td><code>call_type</code></td>
		<td><code>str</code></td>
		<td></td>
	</tr>
    <tr>
		<td><code>joinable_by</code></td>
		<td><code>str</code></td>
		<td></td>
	</tr>
    <tr>
		<td><code>game</code></td>
		<td><code><a href="#game">Game</a></code></td>
		<td></td>
	</tr>
    <tr>
		<td><code>genre</code></td>
		<td><code><a href="#genre">Genre</a></code></td>
		<td></td>
	</tr>
    <tr>
		<td><code>duration_seconds</code></td>
		<td><code>int</code></td>
		<td></td>
	</tr>
    <tr>
		<td><code>max_participants</code></td>
		<td><code>int</code></td>
		<td></td>
	</tr>
    <tr>
		<td><code>conference_call_users</code></td>
		<td><code>list of <a href="#user">User</a></code></td>
		<td></td>
	</tr>
    <tr>
		<td><code>bump_params</code></td>
		<td><code>dict</code></td>
		<td></td>
	</tr>
    <tr>
		<td><code>conference_call_users_role</code></td>
		<td><code>list of <a href="#conferencecalluserrole">ConferenceCallUserRole</a></code></td>
		<td></td>
	</tr>
</table>

### ConferenceCallUserRole

通話のユーザー権限のオブジェクトモデル

<table>
    <tr>
		<th>フィールド</th>
		<th>型</th>
		<th>概要</th>
	</tr>
    <tr>
		<td><code>data</code></td>
		<td><code>dict</code></td>
		<td>レスポンス</td>
	</tr>
    <tr>
		<td><code>id</code></td>
		<td><code>int</code></td>
		<td></td>
	</tr>
    <tr>
		<td><code>role</code></td>
		<td><code>str</code></td>
		<td></td>
	</tr>
</table>

### ContactStatus

コンタクトステータスのオブジェクトモデル

<table>
    <tr>
		<th>フィールド</th>
		<th>型</th>
		<th>概要</th>
	</tr>
    <tr>
		<td><code>data</code></td>
		<td><code>dict</code></td>
		<td>レスポンス</td>
	</tr>
    <tr>
		<td><code>status</code></td>
		<td><code>str</code></td>
		<td></td>
	</tr>
    <tr>
		<td><code>user_id</code></td>
		<td><code>int</code></td>
		<td></td>
	</tr>
</table>

### CreateGroupQuota

サークルの作成可能割当量のオブジェクトモデル

<table>
    <tr>
		<th>フィールド</th>
		<th>型</th>
		<th>概要</th>
	</tr>
    <tr>
		<td><code>data</code></td>
		<td><code>dict</code></td>
		<td>レスポンス</td>
	</tr>
    <tr>
		<td><code>used_quota</code></td>
		<td><code>int</code></td>
		<td></td>
	</tr>
    <tr>
		<td><code>remaining_quota</code></td>
		<td><code>int</code></td>
		<td></td>
	</tr>
</table>

<p align="right">(<a href="#readme-top">トップに戻る</a>)</p>

## レスポンス

<ul>
    <li><a href="#activefollowingsresponse">ActiveFollowingsResponse</a></li>
</ul>

### ActiveFollowingsResponse

フォロー中のユーザーのオンラインステータス

<table>
    <tr>
		<th>フィールド</th>
		<th>型</th>
		<th>概要</th>
	</tr>
    <tr>
		<td><code>data</code></td>
		<td><code>dict</code></td>
		<td>レスポンス</td>
	</tr>
    <tr>
		<td><code>last_loggedin_at</code></td>
		<td><code>int</code></td>
		<td>最終ログイン日時</td>
	</tr>
    <tr>
		<td><code>users</code></td>
		<td><code>list of <a href="#user">User</a></code></td>
		<td>ユーザーの詳細のリスト</td>
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
		<td>プライベートユーザー</td>
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
		<td>スパムメールの可能性</td>
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
