<div><a id="readme-top"></a></div>

# yaylib - API ドキュメント

### yaylib について

「yaylib」は、Yay!（イェイ）内データのプログラムによる取得や解析、あらゆる操作の自動化や、ボットの開発に使用することができます。

<details open>
    <summary>もくじ</summary>
    <ol>
        <li><a href="#インストール">インストール</a></li>
        <li><a href="#使用方法">使用方法</a></li>
        <!-- <li><a href="async-client">Async Client</a></li> -->
        <li><a href="#api-一覧">API 一覧</a></li>
        <li><a href="#モデル">モデル</a></li>
        <li><a href="#レスポンス">レスポンス</a></li>
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

### 具体的な使用例

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
		<td rowspan="3"><code><a href="#login-update-response">LoginUpdateResponse</a></code></td>
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
		<td rowspan="2"><code><a href="#login-update-response">LoginUpdateResponse</a></code></td>
		<td rowspan="2">パスワードを変更します</td>
	</tr>
	<tr>
		<td>new_password</td>
		<td><code>str</code></td>
	</tr>
    <!-- get_token -->
	<tr id="user-content-get-token">
		<td rowspan="4"><code>get_token()</code></td>
		<td>grant_type</td>
		<td><code>str</code></td>
        <td rowspan="4"><code><a href="#token-response">TokenResponse</a></code></td>
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
		<td rowspan="2"><code><a href="#login-user-response">LoginUserResponse</a></code></td>
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
		<td><code><a href="#login-user-response">LoginUserResponse</a></code></td>
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
		<td rowspan="4"><code><a href="#login-update-response">LoginUpdateResponse</a></code></td>
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
		<td rowspan="2"><code><a href="#bookmark-post-response">BookmarkPostResponse</a></code></td>
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
        <td rowspan="18"><code><a href="#conference-call">ConferenceCall</a></code></td>
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
		<td rowspan="2"><code><a href="#posts-response">PostsResponse</a></code></td>
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
		<td rowspan="9"><code><a href="#posts-response">PostsResponse</a></code></td>
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
		<td rowspan="6"><code><a href="#posts-response">PostsResponse</a></code></td>
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
		<td><code><a href="#posts-response">PostsResponse</a></code></td>
		<td>会話の原点の投稿を取得します</td>
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
		<td><code><a href="#">User</a></code></td>
		<td>ユーザーの詳細</td>
	</tr>
    <tr>
		<td><code>from_post</code></td>
		<td><code><a href="#">Post</a></code></td>
		<td>投稿の詳細</td>
	</tr>
    <tr>
		<td><code>to_post</code></td>
		<td><code><a href="#">Post</a></code></td>
		<td>投稿の詳細</td>
	</tr>
    <tr>
		<td><code>group</code></td>
		<td><code><a href="#">Group</a></code></td>
		<td>サークルの詳細</td>
	</tr>
    <tr>
		<td><code>followers</code></td>
		<td><code>list of <a href="#">User</a></code></td>
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
		<td><code>list of <a href="#">User</a></code></td>
		<td>ユーザーの詳細のリスト</td>
	</tr>
</table>

<p align="right">(<a href="#readme-top">トップに戻る</a>)</p>
