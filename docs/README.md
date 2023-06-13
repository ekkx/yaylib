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

api = yaylib.Client
api.login(email="[メールアドレス]", password="[パスワード]")

api.login_data.access_token # アクセストークンを取得
api.login_data.refresh_token # リフレッシュトークンを取得（アクセストークン再発行時に使用）
```

一定時間内に何度もログインすると、**429 Too Many Requests エラー**が発生します。
よって、初回実行するときは、② の方法でアクセストークンを取得し、以後は ① のようにアクセストークンを使用してクライアントを初期化することをお勧めします。

※ アクセストークンには有効期限が設けられており、24 時間が経過する前に<a href="#get-token">`get_token()`</a>関数を使用するか、ログインして再発行してください。

<p align="right">(<a href="#readme-top">トップに戻る</a>)</p>

## API 一覧

<ul>
    <li><a href="#ログイン-api">ログイン API</a></li>
    <li><a href="#ユーザー-api">ユーザー API</a></li>
    <li><a href="#投稿-api">投稿 API</a></li>
    <li><a href="#スレッド-api">スレッド API</a></li>
    <li><a href="#レター-api">レター API</a></li>
    <li><a href="#チャット-api">チャット API</a></li>
    <li><a href="#グループ-api">グループ API</a></li>
    <li><a href="#通話-api">通話 API</a></li>
    <li><a href="#通知-api">通知 API</a></li>
    <li><a href="#その他-api">その他 API</a></li>
</ul>

### ログイン API

<table>
    <tr>
		<th>メソッド</th>
		<th>引数</th>
		<th>型</th>
		<th>説明</th>
	</tr>
    <!-- change_email -->
    <tr>
		<td rowspan="3"><code>change_email()</code></td>
		<td>email</td>
		<td><code>str</code></td>
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
    <!-- get_token -->
	<tr id="get-token">
		<td rowspan="4"><code>get_token()</code></td>
		<td>grant_type</td>
		<td><code>str</code></td>
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
</table>

### ユーザー API

<table>
    <tr>
		<th>メソッド</th>
		<th>引数</th>
		<th>型</th>
		<th>説明</th>
	</tr>
    <tr>
		<td id="get-token"><code>change_email()</code></td>
		<td>group_id</td>
		<td><code>int</code></td>
		<td>アクセストークンを再発行します</td>
	</tr>
</table>

### 投稿 API

<table>
    <tr>
		<th>メソッド</th>
		<th>引数</th>
		<th>説明</th>
	</tr>
    <tr>
		<td><code>accept_moderator_offer()</code></td>
		<td>
            <li>group_id: int</li>
        </td>
		<td>グループ副管理人の権限オファーを引き受けます</td>
	</tr>
</table>

### スレッド API

<table>
    <tr>
		<th>メソッド</th>
		<th>引数</th>
		<th>説明</th>
	</tr>
    <tr>
		<td><code>accept_moderator_offer()</code></td>
		<td>
            <li>group_id: int</li>
        </td>
		<td>グループ副管理人の権限オファーを引き受けます</td>
	</tr>
</table>

### レター API

<table>
    <tr>
		<th>メソッド</th>
		<th>引数</th>
		<th>説明</th>
	</tr>
    <tr>
		<td><code>accept_moderator_offer()</code></td>
		<td>
            <li>group_id: int</li>
        </td>
		<td>グループ副管理人の権限オファーを引き受けます</td>
	</tr>
</table>

### チャット API

<table>
    <tr>
		<th>メソッド</th>
		<th>引数</th>
		<th>説明</th>
	</tr>
    <tr>
		<td><code>accept_moderator_offer()</code></td>
		<td>
            <li>group_id: int</li>
        </td>
		<td>グループ副管理人の権限オファーを引き受けます</td>
	</tr>
</table>

### グループ API

<table>
	<tr>
		<th>メソッド</th>
		<th>引数</th>
		<th>説明</th>
	</tr>
	<tr>
		<td><code>accept_moderator_offer()</code></td>
		<td>
            <li>group_id: int</li>
        </td>
		<td>グループ副管理人の権限オファーを引き受けます</td>
	</tr>
	<tr>
		<td><code>accept_ownership_offer()</code></td>
		<td>
            <li>group_id: int</li>
        </td>
		<td>グループ管理人の権限オファーを引き受けます</td>
	</tr>
	<tr>
		<td><code>accept_group_join_request()</code></td>
		<td>
            <li>group_id: int</li>
            <li>user_id: int</li>
        </td>
		<td>グループ参加リクエストを承認します</td>
	</tr>
	<tr>
		<td><code>add_related_groups()</code></td>
		<td>
            <li>group_id: int</li>
            <li>related_group_id: List[int]</li>
        </td>
		<td>関連グループを追加します</td>
	</tr>
	<tr>
		<td><code>ban_group_user()</code></td>
		<td>
            <li>group_id: int</li>
            <li>user_id: int</li>
        </td>
		<td>グループからユーザーを追放します</td>
	</tr>
	<tr>
		<td><code>check_unread_status()</code></td>
		<td>
            <li>from_time: int (任意)</li>
        </td>
		<td>グループの未読ステータスを取得します</td>
	</tr>
	<tr>
		<td><code>create_group()</code></td>
		<td>
            <li>topic: str</li>
            <li>description: str (任意)</li>
            <li>secret: bool (任意)</li>
            <li>hide_reported_posts: bool (任意)</li>
            <li>hide_conference_call: bool (任意)</li>
            <li>is_private: bool (任意)</li>
            <li>only_verified_age: bool (任意)</li>
            <li>only_mobile_verified: bool (任意)</li>
            <li>call_timeline_display: bool (任意)</li>
            <li>allow_ownership_transfer: bool (任意)</li>
            <li>allow_thread_creation_by: str (任意)</li>
            <li>gender: int (任意)</li>
            <li>generation_groups_limit: int (任意)</li>
            <li>group_category_id: int (任意)</li>
            <li>cover_image_filename: str (任意)</li>
            <li>sub_category_id: str (任意)</li>
            <li>hide_from_game_eight: bool (任意)</li>
            <li>allow_members_to_post_media: bool (任意)</li>
            <li>allow_members_to_post_url: bool (任意)</li>
            <li>guidelines: str (任意)</li>
        </td>
		<td>グループを作成します</td>
	</tr>
    <tr>
		<td><code>create_pin_group()</code></td>
		<td>
            <li>group_id: int</li>
        </td>
		<td>グループの投稿をピンします</td>
	</tr>
    <tr>
		<td><code>decline_moderator_offer()</code></td>
		<td>
            <li>group_id: int</li>
        </td>
		<td>グループ副管理人の権限オファーを断ります</td>
	</tr>
    <tr>
		<td><code>decline_ownership_offer()</code></td>
		<td>
            <li>group_id: int</li>
        </td>
		<td>グループ管理人の権限オファーを断ります</td>
	</tr>
    <tr>
		<td><code>decline_group_join_request()</code></td>
		<td>
            <li>group_id: int</li>
            <li>user_id: int</li>
        </td>
		<td>グループ参加リクエストを断ります</td>
	</tr>
</table>

### 通話 API

<table>
    <tr>
		<th>メソッド</th>
		<th>引数</th>
		<th>説明</th>
	</tr>
    <tr>
		<td><code>accept_moderator_offer()</code></td>
		<td>
            <li>group_id: int</li>
        </td>
		<td>グループ副管理人の権限オファーを引き受けます</td>
	</tr>
</table>

### 通知 API

<table>
    <tr>
		<th>メソッド</th>
		<th>引数</th>
		<th>説明</th>
	</tr>
    <tr>
		<td><code>accept_moderator_offer()</code></td>
		<td>
            <li>group_id: int</li>
        </td>
		<td>グループ副管理人の権限オファーを引き受けます</td>
	</tr>
</table>

### その他 API

<table>
    <tr>
		<th>メソッド</th>
		<th>引数</th>
		<th>説明</th>
	</tr>
    <tr>
		<td><code>accept_moderator_offer()</code></td>
		<td>
            <li>group_id: int</li>
        </td>
		<td>グループ副管理人の権限オファーを引き受けます</td>
	</tr>
</table>

<p align="right">(<a href="#readme-top">トップに戻る</a>)</p>

## モデル

<ul>
    <li><a href="#activity">Activity</a></li>
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
		<td>dict</td>
		<td>レスポンス</td>
	</tr>
    <tr>
		<td><code>id</code></td>
		<td>int</td>
		<td>通知のID</td>
	</tr>
    <tr>
		<td><code>created_at</code></td>
		<td>int</td>
		<td>タイムスタンブ</td>
	</tr>
    <tr>
		<td><code>created_at_parsed</code></td>
		<td>str</td>
		<td>パースされたタイムスタンブ</td>
	</tr>
    <tr>
		<td><code>type</code></td>
		<td>str</td>
		<td>通知のタイプ</td>
	</tr>
    <tr>
		<td><code>user</code></td>
		<td>User</td>
		<td>ユーザーの詳細</td>
	</tr>
    <tr>
		<td><code>from_post</code></td>
		<td>Post</td>
		<td>投稿の詳細</td>
	</tr>
    <tr>
		<td><code>to_post</code></td>
		<td>Post</td>
		<td>投稿の詳細</td>
	</tr>
    <tr>
		<td><code>group</code></td>
		<td>int</td>
		<td>サークルの詳細</td>
	</tr>
    <tr>
		<td><code>followers</code></td>
		<td>list of User</td>
		<td>ユーザーの</td>
	</tr>
    <tr>
		<td><code>from_post_ids</code></td>
		<td>list of int</td>
		<td>投稿IDのリスト</td>
	</tr>
    <tr>
		<td><code>vip_reward</code></td>
		<td>int</td>
		<td>VIPのリワード</td>
	</tr>
    <tr>
		<td><code>is_bulk_invitation</code></td>
		<td>bool</td>
		<td>通話への招待かどうか</td>
	</tr>
</table>

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
		<td><code>User</code></td>
		<td>ユーザーの詳細</td>
	</tr>
    <tr>
		<td><code>from_post</code></td>
		<td><code>Post</code></td>
		<td>投稿の詳細</td>
	</tr>
    <tr>
		<td><code>to_post</code></td>
		<td><code>Post</code></td>
		<td>投稿の詳細</td>
	</tr>
    <tr>
		<td><code>group</code></td>
		<td><code>int</code></td>
		<td>サークルの詳細</td>
	</tr>
    <tr>
		<td><code>followers</code></td>
		<td><code>list of User</code></td>
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
		<td>dict</td>
		<td>レスポンス</td>
	</tr>
    <tr>
		<td><code>last_loggedin_at</code></td>
		<td>int</td>
		<td>最終ログイン日時</td>
	</tr>
    <tr>
		<td><code>users</code></td>
		<td>list of User</td>
		<td>ユーザーの詳細のリスト</td>
	</tr>
</table>

<p align="right">(<a href="#readme-top">トップに戻る</a>)</p>
