<div><a id="readme-top"></a></div>
<p align="center">
    <img src="https://github.com/qvco/yaylib/assets/77382767/5265b956-55b7-466c-8cdb-cf0f3abed946" alt="Logo" height="300px">
    <h3 align="center">yaylib - Docs | ドキュメント</h3>
    <p align="center">
        <br />
        <a href="https://github.com/qvco/yaylib/blob/main/docs/API-Reference/README.md">
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
        <li><a href="#インストール">インストール</a></li>
        <li><a href="#使用方法">使用方法</a></li>
        <!-- <li><a href="async-client">Async Client</a></li> -->
        <li><a href="#メソッド一覧">メソッド一覧</a></li>
        <li><a href="#モデル">モデル</a></li>
        <li><a href="#レスポンス">レスポンス</a></li>
        <li><a href="#yay-エラーコード一覧">Yay! エラーコード一覧</a></li>
    </ol>
</details>

## インストール / Install

**※ Python 3.11 かそれ以上のバージョンが必要です。**

ライブラリをインストールするには、以下のコマンドを実行します:

```bash
pip install yaylib
```

開発バージョンをインストールする場合は、以下の手順を実行します:

```bash
git clone https://github.com/qvco/yaylib

cd yaylib

pip install -r requirements.txt

pip install -e .
```

## 使用方法 / Usage

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
		<td><code>proxy</code></td>
		<td><code>str</code></td>
		<td><code>None</code></td>
		<td>プロキシサーバーのアドレス</td>
	</tr>
    <tr>
		<td><code>max_retries</code></td>
		<td><code>int</code></td>
		<td><code>None</code></td>
		<td>リクエストに失敗した際のリトライ回数</td>
	</tr>
    <tr>
		<td><code>backoff_factor</code></td>
		<td><code>float</code></td>
		<td><code>None</code></td>
		<td>リトライ待機時間の増加割合を指定する係数</td>
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
		<td><code>lang</code></td>
		<td><code>str</code></td>
		<td><code>ja</code></td>
		<td>エラー出力言語（ja もしくは、en）</td>
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

以下の例を使用して、API クライアントを初期化します。

```python
import yaylib

api = yaylib.Client()
api.login(email="メールアドレス", password="パスワード")
```

アクセストークンを直接設定することも可能です。

```python
import yaylib

api = yaylib.Client(access_token="アクセストークン")

api.create_post("こんにちは")
```

<p align="right">(<a href="#readme-top">トップに戻る</a>)</p>

## メソッド一覧

### 投稿

<table>
    <tr>
		<th>メソッド</th>
		<th>概要</th>
	</tr>
	<tr>
		<td><code><a href="https://github.com/qvco/yaylib/blob/main/docs/API-Reference/post/add_bookmark.md">add_bookmark()</a></code></td>
		<td>ブックマークに追加します</td>
	</tr>
	<tr>
		<td><code><a href="https://github.com/qvco/yaylib/blob/main/docs/API-Reference/post/add_group_highlight_post.md">add_group_highlight_post()</a></code></td>
		<td>投稿をグループのまとめに追加します</td>
	</tr>
	<tr>
		<td><code><a href="https://github.com/qvco/yaylib/blob/main/docs/API-Reference/post/create_call_post.md">create_call_post()</a></code></td>
		<td>通話の投稿を作成します</td>
	</tr>
	<tr>
		<td><code><a href="https://github.com/qvco/yaylib/blob/main/docs/API-Reference/post/pin_group_post.md">pin_group_post()</a></code></td>
		<td>グループの投稿をピンします</td>
	</tr>
	<tr>
		<td><code><a href="https://github.com/qvco/yaylib/blob/main/docs/API-Reference/post/pin_post.md">pin_post()</a></code></td>
		<td>投稿をピンします</td>
	</tr>
	<tr>
		<td><code><a href="https://github.com/qvco/yaylib/blob/main/docs/API-Reference/post/mention.md">mention()</a></code></td>
		<td>メンション用文字列を返します</td>
	</tr>
	<tr>
		<td><code><a href="https://github.com/qvco/yaylib/blob/main/docs/API-Reference/post/create_post.md">create_post()</a></code></td>
		<td>投稿を作成します</td>
	</tr>
	<tr>
		<td><code><a href="https://github.com/qvco/yaylib/blob/main/docs/API-Reference/post/create_repost.md">create_repost()</a></code></td>
		<td>投稿を(´∀｀∩)↑age↑します</td>
	</tr>
	<tr>
		<td><code><a href="https://github.com/qvco/yaylib/blob/main/docs/API-Reference/post/create_share_post.md">create_share_post()</a></code></td>
		<td>シェア投稿を作成します</td>
	</tr>
	<tr>
		<td><code><a href="https://github.com/qvco/yaylib/blob/main/docs/API-Reference/post/create_thread_post.md">create_thread_post()</a></code></td>
		<td>スレッドの投稿を作成します</td>
	</tr>
	<tr>
		<td><code><a href="https://github.com/qvco/yaylib/blob/main/docs/API-Reference/post/delete_all_post.md">delete_all_post()</a></code></td>
		<td>すべての自分の投稿を削除します</td>
	</tr>
	<tr>
		<td><code><a href="https://github.com/qvco/yaylib/blob/main/docs/API-Reference/post/unpin_group_post.md">unpin_group_post()</a></code></td>
		<td>グループのピン投稿を解除します</td>
	</tr>
	<tr>
		<td><code><a href="https://github.com/qvco/yaylib/blob/main/docs/API-Reference/post/unpin_post.md">unpin_post()</a></code></td>
		<td>ピン投稿を削除します</td>
	</tr>
	<tr>
		<td><code><a href="https://github.com/qvco/yaylib/blob/main/docs/API-Reference/post/get_bookmark.md">get_bookmark()</a></code></td>
		<td>ブックマークを取得します</td>
	</tr>
	<tr>
		<td><code><a href="https://github.com/qvco/yaylib/blob/main/docs/API-Reference/post/get_timeline_calls.md">get_timeline_calls()</a></code></td>
		<td>誰でも通話を取得します</td>
	</tr>
	<tr>
		<td><code><a href="https://github.com/qvco/yaylib/blob/main/docs/API-Reference/post/get_conversation.md">get_conversation()</a></code></td>
		<td>リプライを含める投稿の会話を取得します</td>
	</tr>
	<tr>
		<td><code><a href="https://github.com/qvco/yaylib/blob/main/docs/API-Reference/post/get_following_call_timeline.md">get_following_call_timeline()</a></code></td>
		<td>フォロー中の通話を取得します</td>
	</tr>
	<tr>
		<td><code><a href="https://github.com/qvco/yaylib/blob/main/docs/API-Reference/post/get_following_timeline.md">get_following_timeline()</a></code></td>
		<td>フォロー中のタイムラインを取得します</td>
	</tr>
	<tr>
		<td><code><a href="https://github.com/qvco/yaylib/blob/main/docs/API-Reference/post/get_group_highlight_posts.md">get_group_highlight_posts()</a></code></td>
		<td>グループのハイライト投稿を取得します</td>
	</tr>
	<tr>
		<td><code><a href="https://github.com/qvco/yaylib/blob/main/docs/API-Reference/post/get_group_timeline_by_keyword.md">get_group_timeline_by_keyword()</a></code></td>
		<td>グループの投稿をキーワードで検索します</td>
	</tr>
	<tr>
		<td><code><a href="https://github.com/qvco/yaylib/blob/main/docs/API-Reference/post/get_group_timeline.md">get_group_timeline()</a></code></td>
		<td>グループのタイムラインを取得します</td>
	</tr>
	<tr>
		<td><code><a href="https://github.com/qvco/yaylib/blob/main/docs/API-Reference/post/get_timeline_by_hashtag.md">get_timeline_by_hashtag()</a></code></td>
		<td>ハッシュタグでタイムラインを検索します</td>
	</tr>
	<tr>
		<td><code><a href="https://github.com/qvco/yaylib/blob/main/docs/API-Reference/post/get_my_posts.md">get_my_posts()</a></code></td>
		<td>自分の投稿を取得します</td>
	</tr>
	<tr>
		<td><code><a href="https://github.com/qvco/yaylib/blob/main/docs/API-Reference/post/get_post.md">get_post()</a></code></td>
		<td>投稿の詳細を取得します</td>
	</tr>
	<tr>
		<td><code><a href="https://github.com/qvco/yaylib/blob/main/docs/API-Reference/post/get_post_likers.md">get_post_likers()</a></code></td>
		<td>投稿にいいねしたユーザーを取得します</td>
	</tr>
	<tr>
		<td><code><a href="https://github.com/qvco/yaylib/blob/main/docs/API-Reference/post/get_reposts.md">get_reposts()</a></code></td>
		<td>投稿の(´∀｀∩)↑age↑を取得します</td>
	</tr>
	<tr>
		<td><code><a href="https://github.com/qvco/yaylib/blob/main/docs/API-Reference/post/get_posts.md">get_posts()</a></code></td>
		<td>複数の投稿を取得します</td>
	</tr>
	<tr>
		<td><code><a href="https://github.com/qvco/yaylib/blob/main/docs/API-Reference/post/get_recommended_post_tags.md">get_recommended_post_tags()</a></code></td>
		<td>おすすめのタグ候補を取得します</td>
	</tr>
	<tr>
		<td><code><a href="https://github.com/qvco/yaylib/blob/main/docs/API-Reference/post/get_recommended_posts.md">get_recommended_posts()</a></code></td>
		<td>おすすめの投稿を取得します</td>
	</tr>
	<tr>
		<td><code><a href="https://github.com/qvco/yaylib/blob/main/docs/API-Reference/post/get_timeline_by_keyword.md">get_timeline_by_keyword()</a></code></td>
		<td>キーワードでタイムラインを検索します</td>
	</tr>
	<tr>
		<td><code><a href="https://github.com/qvco/yaylib/blob/main/docs/API-Reference/post/get_timeline.md">get_timeline()</a></code></td>
		<td>タイムラインを取得します</td>
	</tr>
	<tr>
		<td><code><a href="https://github.com/qvco/yaylib/blob/main/docs/API-Reference/post/get_url_metadata.md">get_url_metadata()</a></code></td>
		<td>URLのメタデータを取得します</td>
	</tr>
	<tr>
		<td><code><a href="https://github.com/qvco/yaylib/blob/main/docs/API-Reference/post/get_user_timeline.md">get_user_timeline()</a></code></td>
		<td>ユーザーのタイムラインを取得します</td>
	</tr>
	<tr>
		<td><code><a href="https://github.com/qvco/yaylib/blob/main/docs/API-Reference/post/like_posts.md">like_posts()</a></code></td>
		<td>投稿にいいねします</td>
	</tr>
	<tr>
		<td><code><a href="https://github.com/qvco/yaylib/blob/main/docs/API-Reference/post/remove_bookmark.md">remove_bookmark()</a></code></td>
		<td>ブックマークを削除します</td>
	</tr>
	<tr>
		<td><code><a href="https://github.com/qvco/yaylib/blob/main/docs/API-Reference/post/remove_group_highlight_post.md">remove_group_highlight_post()</a></code></td>
		<td>サークルのハイライト投稿を解除します</td>
	</tr>
	<tr>
		<td><code><a href="https://github.com/qvco/yaylib/blob/main/docs/API-Reference/post/remove_posts.md">remove_posts()</a></code></td>
		<td>複数の投稿を削除します</td>
	</tr>
	<tr>
		<td><code><a href="https://github.com/qvco/yaylib/blob/main/docs/API-Reference/post/report_post.md">report_post()</a></code></td>
		<td>投稿を通報します</td>
	</tr>
	<tr>
		<td><code><a href="https://github.com/qvco/yaylib/blob/main/docs/API-Reference/post/unlike_post.md">unlike_post()</a></code></td>
		<td>投稿のいいねを解除します</td>
	</tr>
	<tr>
		<td><code><a href="https://github.com/qvco/yaylib/blob/main/docs/API-Reference/post/update_post.md">update_post()</a></code></td>
		<td>投稿を編集します</td>
	</tr>
	<tr>
		<td><code><a href="https://github.com/qvco/yaylib/blob/main/docs/API-Reference/post/view_video.md">view_video()</a></code></td>
		<td>動画を視聴します</td>
	</tr>
	<tr>
		<td><code><a href="https://github.com/qvco/yaylib/blob/main/docs/API-Reference/post/vote_survey.md">vote_survey()</a></code></td>
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
		<td><code><a href="https://github.com/qvco/yaylib/blob/main/docs/API-Reference/thread/add_post_to_thread.md">add_post_to_thread()</a></code></td>
		<td>投稿をスレッドに追加します</td>
	</tr>
	<tr>
		<td><code><a href="https://github.com/qvco/yaylib/blob/main/docs/API-Reference/thread/convert_post_to_thread.md">convert_post_to_thread()</a></code></td>
		<td>投稿をスレッドに変換します</td>
	</tr>
	<tr>
		<td><code><a href="https://github.com/qvco/yaylib/blob/main/docs/API-Reference/thread/create_thread.md">create_thread()</a></code></td>
		<td>スレッドを作成します</td>
	</tr>
	<tr>
		<td><code><a href="https://github.com/qvco/yaylib/blob/main/docs/API-Reference/thread/get_group_thread_list.md">get_group_thread_list()</a></code></td>
		<td>グループのスレッド一覧を取得します</td>
	</tr>
	<tr>
		<td><code><a href="https://github.com/qvco/yaylib/blob/main/docs/API-Reference/thread/get_thread_joined_statuses.md">get_thread_joined_statuses()</a></code></td>
		<td>スレッド参加ステータスを取得します</td>
	</tr>
	<tr>
		<td><code><a href="https://github.com/qvco/yaylib/blob/main/docs/API-Reference/thread/get_thread_posts.md">get_thread_posts()</a></code></td>
		<td>スレッドの投稿を取得します</td>
	</tr>
	<tr>
		<td><code><a href="https://github.com/qvco/yaylib/blob/main/docs/API-Reference/thread/join_thread.md">join_thread()</a></code></td>
		<td>スレッドに参加します</td>
	</tr>
	<tr>
		<td><code><a href="https://github.com/qvco/yaylib/blob/main/docs/API-Reference/thread/leave_thread.md">leave_thread()</a></code></td>
		<td>スレッドから脱退します</td>
	</tr>
	<tr>
		<td><code><a href="https://github.com/qvco/yaylib/blob/main/docs/API-Reference/thread/remove_thread.md">remove_thread()</a></code></td>
		<td>スレッドを削除します</td>
	</tr>
	<tr>
		<td><code><a href="https://github.com/qvco/yaylib/blob/main/docs/API-Reference/thread/update_thread.md">update_thread()</a></code></td>
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
		<td><code><a href="https://github.com/qvco/yaylib/blob/main/docs/API-Reference/user/delete_footprint.md">delete_footprint()</a></code></td>
		<td>足跡を削除します</td>
	</tr>
	<tr>
		<td><code><a href="https://github.com/qvco/yaylib/blob/main/docs/API-Reference/user/destroy_user.md">destroy_user()</a></code></td>
		<td>アカウントを削除します</td>
	</tr>
	<tr>
		<td><code><a href="https://github.com/qvco/yaylib/blob/main/docs/API-Reference/user/follow_user.md">follow_user()</a></code></td>
		<td>ユーザーをフォローします</td>
	</tr>
	<tr>
		<td><code><a href="https://github.com/qvco/yaylib/blob/main/docs/API-Reference/user/follow_users.md">follow_users()</a></code></td>
		<td>複数のユーザーをフォローします</td>
	</tr>
	<tr>
		<td><code><a href="https://github.com/qvco/yaylib/blob/main/docs/API-Reference/user/get_active_followings.md">get_active_followings()</a></code></td>
		<td>アクティブなフォロー中のユーザーを取得します</td>
	</tr>
	<tr>
		<td><code><a href="https://github.com/qvco/yaylib/blob/main/docs/API-Reference/user/get_follow_recommendations.md">get_follow_recommendations()</a></code></td>
		<td>フォローするのにおすすめのユーザーを取得します</td>
	</tr>
	<tr>
		<td><code><a href="https://github.com/qvco/yaylib/blob/main/docs/API-Reference/user/get_follow_request.md">get_follow_request()</a></code></td>
		<td>フォローリクエストを取得します</td>
	</tr>
	<tr>
		<td><code><a href="https://github.com/qvco/yaylib/blob/main/docs/API-Reference/user/get_follow_request_count.md">get_follow_request_count()</a></code></td>
		<td>フォローリクエストの数を取得します</td>
	</tr>
	<tr>
		<td><code><a href="https://github.com/qvco/yaylib/blob/main/docs/API-Reference/user/get_following_users_born.md">get_following_users_born()</a></code></td>
		<td>フォロー中のユーザーの誕生日を取得します</td>
	</tr>
	<tr>
		<td><code><a href="https://github.com/qvco/yaylib/blob/main/docs/API-Reference/user/get_footprints.md">get_footprints()</a></code></td>
		<td>足跡を取得します</td>
	</tr>
	<tr>
		<td><code><a href="https://github.com/qvco/yaylib/blob/main/docs/API-Reference/user/get_fresh_user.md">get_fresh_user()</a></code></td>
		<td>認証情報などを含んだユーザー情報を取得します</td>
	</tr>
	<tr>
		<td><code><a href="https://github.com/qvco/yaylib/blob/main/docs/API-Reference/user/get_hima_users.md">get_hima_users()</a></code></td>
		<td>暇なユーザーを取得します</td>
	</tr>
	<tr>
		<td><code><a href="https://github.com/qvco/yaylib/blob/main/docs/API-Reference/user/get_user_ranking.md">helget_user_rankinglo()</a></code></td>
		<td>ユーザーのフォロワーランキングを取得します</td>
	</tr>
	<tr>
		<td><code><a href="https://github.com/qvco/yaylib/blob/main/docs/API-Reference/user/get_refresh_counter_requests.md">get_refresh_counter_requests()</a></code></td>
		<td>カウンター更新のリクエストを取得します</td>
	</tr>
	<tr>
		<td><code><a href="https://github.com/qvco/yaylib/blob/main/docs/API-Reference/user/get_social_shared_users.md">get_social_shared_users()</a></code></td>
		<td>SNS共有をしたユーザーを取得します</td>
	</tr>
	<tr>
		<td><code><a href="https://github.com/qvco/yaylib/blob/main/docs/API-Reference/user/get_timestamp.md">get_timestamp()</a></code></td>
		<td>タイムスタンプを取得します</td>
	</tr>
	<tr>
		<td><code><a href="https://github.com/qvco/yaylib/blob/main/docs/API-Reference/user/get_user.md">get_user()</a></code></td>
		<td>ユーザーの情報を取得します</td>
	</tr>
	<tr>
		<td><code><a href="https://github.com/qvco/yaylib/blob/main/docs/API-Reference/user/get_user_email.md">get_user_email()</a></code></td>
		<td>ユーザーのメールアドレスを取得します</td>
	</tr>
	<tr>
		<td><code><a href="https://github.com/qvco/yaylib/blob/main/docs/API-Reference/user/get_user_followers.md">get_user_followers()</a></code></td>
		<td>ユーザーのフォロワーを取得します</td>
	</tr>
	<tr>
		<td><code><a href="https://github.com/qvco/yaylib/blob/main/docs/API-Reference/user/get_user_followings.md">get_user_followings()</a></code></td>
		<td>フォロー中のユーザーを取得します</td>
	</tr>
	<tr>
		<td><code><a href="https://github.com/qvco/yaylib/blob/main/docs/API-Reference/user/get_user_from_qr.md">get_user_from_qr()</a></code></td>
		<td>QRコードからユーザーを取得します</td>
	</tr>
	<tr>
		<td><code><a href="https://github.com/qvco/yaylib/blob/main/docs/API-Reference/user/get_user_without_leaving_footprint.md">get_user_without_leaving_footprint()</a></code></td>
		<td>足跡をつけずにユーザーの情報を取得します</td>
	</tr>
	<tr>
		<td><code><a href="https://github.com/qvco/yaylib/blob/main/docs/API-Reference/user/get_users.md">get_users()</a></code></td>
		<td>複数のユーザーの情報を取得します</td>
	</tr>
	<tr>
		<td><code><a href="https://github.com/qvco/yaylib/blob/main/docs/API-Reference/user/reduce_kenta_penalty.md">reduce_kenta_penalty()</a></code></td>
		<td>ペナルティーを緩和します</td>
	</tr>
	<tr>
		<td><code><a href="https://github.com/qvco/yaylib/blob/main/docs/API-Reference/user/refresh_counter.md">refresh_counter()</a></code></td>
		<td>カウンターを更新します</td>
	</tr>
	<tr>
		<td><code><a href="https://github.com/qvco/yaylib/blob/main/docs/API-Reference/user/remove_user_avatar.md">remove_user_avatar()</a></code></td>
		<td>ユーザーのアイコンを削除します</td>
	</tr>
	<tr>
		<td><code><a href="https://github.com/qvco/yaylib/blob/main/docs/API-Reference/user/remove_user_cover.md">remove_user_cover()</a></code></td>
		<td>ユーザーのカバー画像を削除します</td>
	</tr>
	<tr>
		<td><code><a href="https://github.com/qvco/yaylib/blob/main/docs/API-Reference/user/report_user.md">report_user()</a></code></td>
		<td>ユーザーを通報します</td>
	</tr>
	<tr>
		<td><code><a href="https://github.com/qvco/yaylib/blob/main/docs/API-Reference/user/reset_password.md">reset_password()</a></code></td>
		<td>パスワードをリセットします</td>
	</tr>
	<tr>
		<td><code><a href="https://github.com/qvco/yaylib/blob/main/docs/API-Reference/user/search_lobi_users.md">search_lobi_users()</a></code></td>
		<td>Lobiのユーザーを検索します</td>
	</tr>
	<tr>
		<td><code><a href="https://github.com/qvco/yaylib/blob/main/docs/API-Reference/user/search_users.md">search_users()</a></code></td>
		<td>ユーザーを検索します</td>
	</tr>
	<tr>
		<td><code><a href="https://github.com/qvco/yaylib/blob/main/docs/API-Reference/user/set_follow_permission_enabled.md">set_follow_permission_enabled()</a></code></td>
		<td>フォローを許可制に設定します</td>
	</tr>
	<tr>
		<td><code><a href="https://github.com/qvco/yaylib/blob/main/docs/API-Reference/user/take_action_follow_request.md">take_action_follow_request()</a></code></td>
		<td></td>
	</tr>
	<tr>
		<td><code><a href="https://github.com/qvco/yaylib/blob/main/docs/API-Reference/user/turn_on_hima.md">turn_on_hima()</a></code></td>
		<td>ひまなうします</td>
	</tr>
	<tr>
		<td><code><a href="https://github.com/qvco/yaylib/blob/main/docs/API-Reference/user/unfollow_user.md">unfollow_user()</a></code></td>
		<td>ユーザーをアンフォローします</td>
	</tr>
	<tr>
		<td><code><a href="https://github.com/qvco/yaylib/blob/main/docs/API-Reference/user/update_user.md">update_user()</a></code></td>
		<td>プロフィールを更新します</td>
	</tr>
	<tr>
		<td><code><a href="https://github.com/qvco/yaylib/blob/main/docs/API-Reference/user/block_user.md">block_user()</a></code></td>
		<td>ユーザーをブロックします</td>
	</tr>
	<tr>
		<td><code><a href="https://github.com/qvco/yaylib/blob/main/docs/API-Reference/user/get_blocked_user_ids.md">get_blocked_user_ids()</a></code></td>
		<td>あなたをブロックしたユーザーを取得します</td>
	</tr>
	<tr>
		<td><code><a href="https://github.com/qvco/yaylib/blob/main/docs/API-Reference/user/get_blocked_users.md">get_blocked_users()</a></code></td>
		<td>ブロックしたユーザーを取得します</td>
	</tr>
	<tr>
		<td><code><a href="https://github.com/qvco/yaylib/blob/main/docs/API-Reference/user/unblock_user.md">unblock_user()</a></code></td>
		<td>ユーザーをアンブロックします</td>
	</tr>
	<tr>
		<td><code><a href="https://github.com/qvco/yaylib/blob/main/docs/API-Reference/user/get_hidden_users_list.md">get_hidden_users_list()</a></code></td>
		<td>非表示のユーザー一覧を取得します</td>
	</tr>
	<tr>
		<td><code><a href="https://github.com/qvco/yaylib/blob/main/docs/API-Reference/user/hide_user.md">hide_user()</a></code></td>
		<td>ユーザーを非表示にします</td>
	</tr>
	<tr>
		<td><code><a href="https://github.com/qvco/yaylib/blob/main/docs/API-Reference/user/unhide_users.md">unhide_users()</a></code></td>
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
		<td><code><a href="https://github.com/qvco/yaylib/blob/main/docs/API-Reference/group/accept_moderator_offer.md">accept_moderator_offer()</a></code></td>
		<td>サークル副管理人の権限オファーを引き受けます</td>
	</tr>
	<tr>
		<td><code><a href="https://github.com/qvco/yaylib/blob/main/docs/API-Reference/group/accept_ownership_offer.md">accept_ownership_offer()</a></code></td>
		<td>サークル管理人の権限オファーを引き受けます</td>
	</tr>
	<tr>
		<td><code><a href="https://github.com/qvco/yaylib/blob/main/docs/API-Reference/group/accept_group_join_request.md">accept_group_join_request()</a></code></td>
		<td>サークル参加リクエストを承認します</td>
	</tr>
	<tr>
		<td><code><a href="https://github.com/qvco/yaylib/blob/main/docs/API-Reference/group/add_related_groups.md">add_related_groups()</a></code></td>
		<td>関連サークルを追加します</td>
	</tr>
	<tr>
		<td><code><a href="https://github.com/qvco/yaylib/blob/main/docs/API-Reference/group/ban_group_user.md">ban_group_user()</a></code></td>
		<td>サークルからユーザーを追放します</td>
	</tr>
	<tr>
		<td><code><a href="https://github.com/qvco/yaylib/blob/main/docs/API-Reference/group/check_unread_status.md">check_unread_status()</a></code></td>
		<td>サークルの未読ステータスを取得します</td>
	</tr>
	<tr>
		<td><code><a href="https://github.com/qvco/yaylib/blob/main/docs/API-Reference/group/create_group.md">create_group()</a></code></td>
		<td>サークルを作成します</td>
	</tr>
	<tr>
		<td><code><a href="https://github.com/qvco/yaylib/blob/main/docs/API-Reference/group/pin_group.md">pin_group()</a></code></td>
		<td>サークルをピンします</td>
	</tr>
	<tr>
		<td><code><a href="https://github.com/qvco/yaylib/blob/main/docs/API-Reference/group/decline_moderator_offer.md">decline_moderator_offer()</a></code></td>
		<td>サークル副管理人の権限オファーを断ります</td>
	</tr>
	<tr>
		<td><code><a href="https://github.com/qvco/yaylib/blob/main/docs/API-Reference/group/decline_ownership_offer.md">decline_ownership_offer()</a></code></td>
		<td>サークル管理人の権限オファーを断ります</td>
	</tr>
	<tr>
		<td><code><a href="https://github.com/qvco/yaylib/blob/main/docs/API-Reference/group/decline_group_join_request.md">decline_group_join_request()</a></code></td>
		<td>サークル参加リクエストを断ります</td>
	</tr>
	<tr>
		<td><code><a href="https://github.com/qvco/yaylib/blob/main/docs/API-Reference/group/unpin_group.md">unpin_group()</a></code></td>
		<td>サークルのピン止めを解除します</td>
	</tr>
	<tr>
		<td><code><a href="https://github.com/qvco/yaylib/blob/main/docs/API-Reference/group/get_banned_group_members.md">get_banned_group_members()</a></code></td>
		<td>追放されたサークルメンバーを取得します</td>
	</tr>
	<tr>
		<td><code><a href="https://github.com/qvco/yaylib/blob/main/docs/API-Reference/group/get_group_categories.md">get_group_categories()</a></code></td>
		<td>サークルのカテゴリーを取得します</td>
	</tr>
	<tr>
		<td><code><a href="https://github.com/qvco/yaylib/blob/main/docs/API-Reference/group/get_create_group_quota.md">get_create_group_quota()</a></code></td>
		<td>サークル作成可能な割当量を取得します</td>
	</tr>
	<tr>
		<td><code><a href="https://github.com/qvco/yaylib/blob/main/docs/API-Reference/group/get_group.md">get_group()</a></code></td>
		<td>サークルの詳細を取得します</td>
	</tr>
	<tr>
		<td><code><a href="https://github.com/qvco/yaylib/blob/main/docs/API-Reference/group/get_groups.md">get_groups()</a></code></td>
		<td>複数のサークルの詳細を取得します</td>
	</tr>
	<tr>
		<td><code><a href="https://github.com/qvco/yaylib/blob/main/docs/API-Reference/group/get_invitable_users.md">get_invitable_users()</a></code></td>
		<td>サークルに招待可能なユーザーを取得します</td>
	</tr>
	<tr>
		<td><code><a href="https://github.com/qvco/yaylib/blob/main/docs/API-Reference/group/get_joined_statuses.md">get_joined_statuses()</a></code></td>
		<td>サークルの参加ステータスを取得します</td>
	</tr>
	<tr>
		<td><code><a href="https://github.com/qvco/yaylib/blob/main/docs/API-Reference/group/get_group_member.md">get_group_member()</a></code></td>
		<td>特定のサークルメンバーの情報を取得します</td>
	</tr>
	<tr>
		<td><code><a href="https://github.com/qvco/yaylib/blob/main/docs/API-Reference/group/get_group_members.md">get_group_members()</a></code></td>
		<td>サークルメンバーを取得します</td>
	</tr>
	<tr>
		<td><code><a href="https://github.com/qvco/yaylib/blob/main/docs/API-Reference/group/get_my_groups.md">get_my_groups()</a></code></td>
		<td>自分のサークルを取得します</td>
	</tr>
	<tr>
		<td><code><a href="https://github.com/qvco/yaylib/blob/main/docs/API-Reference/group/get_relatable_groups.md">get_relatable_groups()</a></code></td>
		<td>関連がある可能性があるサークルを取得します</td>
	</tr>
	<tr>
		<td><code><a href="https://github.com/qvco/yaylib/blob/main/docs/API-Reference/group/get_related_groups.md">get_related_groups()</a></code></td>
		<td>関連があるサークルを取得します</td>
	</tr>
	<tr>
		<td><code><a href="https://github.com/qvco/yaylib/blob/main/docs/API-Reference/group/get_user_groups.md">get_user_groups()</a></code></td>
		<td>特定のユーザーが参加しているサークルを取得します</td>
	</tr>
	<tr>
		<td><code><a href="https://github.com/qvco/yaylib/blob/main/docs/API-Reference/group/invite_users_to_group.md">invite_users_to_group()</a></code></td>
		<td>サークルにユーザーを招待します</td>
	</tr>
	<tr>
		<td><code><a href="https://github.com/qvco/yaylib/blob/main/docs/API-Reference/group/join_group.md">join_group()</a></code></td>
		<td>サークルに参加します</td>
	</tr>
	<tr>
		<td><code><a href="https://github.com/qvco/yaylib/blob/main/docs/API-Reference/group/leave_group.md">leave_group()</a></code></td>
		<td>サークルから脱退します</td>
	</tr>
	<tr>
		<td><code><a href="https://github.com/qvco/yaylib/blob/main/docs/API-Reference/group/post_gruop_social_shared.md">post_gruop_social_shared()</a></code></td>
		<td>サークルのソーシャルシェアを投稿します</td>
	</tr>
	<tr>
		<td><code><a href="https://github.com/qvco/yaylib/blob/main/docs/API-Reference/group/remove_group_cover.md">remove_group_cover()</a></code></td>
		<td>サークルのカバー画像を削除します</td>
	</tr>
	<tr>
		<td><code><a href="https://github.com/qvco/yaylib/blob/main/docs/API-Reference/group/remove_moderator.md">remove_moderator()</a></code></td>
		<td>サークルの副管理人を削除します</td>
	</tr>
	<tr>
		<td><code><a href="https://github.com/qvco/yaylib/blob/main/docs/API-Reference/group/report_group.md">report_group()</a></code></td>
		<td>サークルを通報します</td>
	</tr>
	<tr>
		<td><code><a href="https://github.com/qvco/yaylib/blob/main/docs/API-Reference/group/send_moderator_offers.md">send_moderator_offers()</a></code></td>
		<td>複数人にサークル副管理人のオファーを送信します</td>
	</tr>
	<tr>
		<td><code><a href="https://github.com/qvco/yaylib/blob/main/docs/API-Reference/group/send_ownership_offer.md">send_ownership_offer()</a></code></td>
		<td>サークル管理人権限のオファーを送信します</td>
	</tr>
	<tr>
		<td><code><a href="https://github.com/qvco/yaylib/blob/main/docs/API-Reference/group/set_group_title.md">set_group_title()</a></code></td>
		<td>サークルのタイトルを設定します</td>
	</tr>
	<tr>
		<td><code><a href="https://github.com/qvco/yaylib/blob/main/docs/API-Reference/group/take_over_group_ownership.md">take_over_group_ownership()</a></code></td>
		<td>サークル管理人の権限を引き継ぎます</td>
	</tr>
	<tr>
		<td><code><a href="https://github.com/qvco/yaylib/blob/main/docs/API-Reference/group/unban_group_member.md">unban_group_member()</a></code></td>
		<td>特定のサークルメンバーの追放を解除します</td>
	</tr>
	<tr>
		<td><code><a href="https://github.com/qvco/yaylib/blob/main/docs/API-Reference/group/update_group.md">update_group()</a></code></td>
		<td>サークルを編集します</td>
	</tr>
	<tr>
		<td><code><a href="https://github.com/qvco/yaylib/blob/main/docs/API-Reference/group/withdraw_moderator_offer.md">withdraw_moderator_offer()</a></code></td>
		<td>サークル副管理人のオファーを取り消します</td>
	</tr>
	<tr>
		<td><code><a href="https://github.com/qvco/yaylib/blob/main/docs/API-Reference/group/withdraw_ownership_offer.md">withdraw_ownership_offer()</a></code></td>
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
		<td><code><a href="https://github.com/qvco/yaylib/blob/main/docs/API-Reference/chat/accept_chat_request.md">accept_chat_request()</a></code></td>
		<td>チャットリクエストを承認します</td>
	</tr>
	<tr>
		<td><code><a href="https://github.com/qvco/yaylib/blob/main/docs/API-Reference/chat/check_unread_status.md">check_unread_status()</a></code></td>
		<td>チャットの未読ステータスを確認します</td>
	</tr>
	<tr>
		<td><code><a href="https://github.com/qvco/yaylib/blob/main/docs/API-Reference/chat/create_group_chat.md">create_group_chat()</a></code></td>
		<td>グループチャットを作成します</td>
	</tr>
	<tr>
		<td><code><a href="https://github.com/qvco/yaylib/blob/main/docs/API-Reference/chat/create_private_chat.md">create_private_chat()</a></code></td>
		<td>個人チャットを作成します</td>
	</tr>
	<tr>
		<td><code><a href="https://github.com/qvco/yaylib/blob/main/docs/API-Reference/chat/delete_background.md">delete_background()</a></code></td>
		<td>チャットの背景画像を削除します</td>
	</tr>
	<tr>
		<td><code><a href="https://github.com/qvco/yaylib/blob/main/docs/API-Reference/chat/delete_message.md">delete_message()</a></code></td>
		<td>メッセージを削除します</td>
	</tr>
	<tr>
		<td><code><a href="https://github.com/qvco/yaylib/blob/main/docs/API-Reference/chat/edit_chat_room.md">edit_chat_room()</a></code></td>
		<td>チャットルームを編集します</td>
	</tr>
	<tr>
		<td><code><a href="https://github.com/qvco/yaylib/blob/main/docs/API-Reference/chat/get_chatable_users.md">get_chatable_users()</a></code></td>
		<td>チャット可能なユーザーを取得します</td>
	</tr>
	<tr>
		<td><code><a href="https://github.com/qvco/yaylib/blob/main/docs/API-Reference/chat/get_gifs_data.md">get_gifs_data()</a></code></td>
		<td>チャットルームのGIFデータを取得します</td>
	</tr>
	<tr>
		<td><code><a href="https://github.com/qvco/yaylib/blob/main/docs/API-Reference/chat/get_hidden_chat_rooms.md">get_hidden_chat_rooms()</a></code></td>
		<td>非表示のチャットルームを取得します</td>
	</tr>
	<tr>
		<td><code><a href="https://github.com/qvco/yaylib/blob/main/docs/API-Reference/chat/get_main_chat_rooms.md">get_main_chat_rooms()</a></code></td>
		<td>メインのチャットルームを取得します</td>
	</tr>
	<tr>
		<td><code><a href="https://github.com/qvco/yaylib/blob/main/docs/API-Reference/chat/get_messages.md">get_messages()</a></code></td>
		<td>メッセージを取得します</td>
	</tr>
	<tr>
		<td><code><a href="https://github.com/qvco/yaylib/blob/main/docs/API-Reference/chat/get_request_chat_rooms.md">get_request_chat_rooms()</a></code></td>
		<td>チャットリクエストを取得します</td>
	</tr>
	<tr>
		<td><code><a href="https://github.com/qvco/yaylib/blob/main/docs/API-Reference/chat/get_chat_room.md">get_chat_room()</a></code></td>
		<td>チャットルームの詳細を取得します</td>
	</tr>
	<tr>
		<td><code><a href="https://github.com/qvco/yaylib/blob/main/docs/API-Reference/chat/get_sticker_packs.md">get_sticker_packs()</a></code></td>
		<td>スタンプを取得します</td>
	</tr>
	<tr>
		<td><code><a href="https://github.com/qvco/yaylib/blob/main/docs/API-Reference/chat/get_chat_requests_count.md">get_chat_requests_count()</a></code></td>
		<td>チャットリクエストの数を取得します</td>
	</tr>
	<tr>
		<td><code><a href="https://github.com/qvco/yaylib/blob/main/docs/API-Reference/chat/hide_chat.md">hide_chat()</a></code></td>
		<td>チャットルームを非表示にします</td>
	</tr>
	<tr>
		<td><code><a href="https://github.com/qvco/yaylib/blob/main/docs/API-Reference/chat/invite_to_chat.md">invite_to_chat()</a></code></td>
		<td>チャットルームにユーザーを招待します</td>
	</tr>
	<tr>
		<td><code><a href="https://github.com/qvco/yaylib/blob/main/docs/API-Reference/chat/kick_users_from_chat.md">kick_users_from_chat()</a></code></td>
		<td>チャットルームからユーザーを追放します</td>
	</tr>
	<tr>
		<td><code><a href="https://github.com/qvco/yaylib/blob/main/docs/API-Reference/chat/kick_users_from_chat.md">kick_users_from_chat()</a></code></td>
		<td>チャットルームからユーザーを追放します</td>
	</tr>
	<tr>
		<td><code><a href="https://github.com/qvco/yaylib/blob/main/docs/API-Reference/chat/pin_chat.md">pin_chat()</a></code></td>
		<td>チャットルームをピン止めします</td>
	</tr>
	<tr>
		<td><code><a href="https://github.com/qvco/yaylib/blob/main/docs/API-Reference/chat/read_message.md">read_message()</a></code></td>
		<td>メッセージを既読にします</td>
	</tr>
	<tr>
		<td><code><a href="https://github.com/qvco/yaylib/blob/main/docs/API-Reference/chat/refresh_chat_rooms.md">refresh_chat_rooms()</a></code></td>
		<td>チャットルームを更新します</td>
	</tr>
	<tr>
		<td><code><a href="https://github.com/qvco/yaylib/blob/main/docs/API-Reference/chat/remove_chat_rooms.md">remove_chat_rooms()</a></code></td>
		<td>チャットルームを削除します</td>
	</tr>
	<tr>
		<td><code><a href="https://github.com/qvco/yaylib/blob/main/docs/API-Reference/chat/report_chat_room.md">report_chat_room()</a></code></td>
		<td>チャットルームを通報します</td>
	</tr>
	<tr>
		<td><code><a href="https://github.com/qvco/yaylib/blob/main/docs/API-Reference/chat/send_message.md">send_message()</a></code></td>
		<td>チャットルームにメッセージを送信します</td>
	</tr>
	<tr>
		<td><code><a href="https://github.com/qvco/yaylib/blob/main/docs/API-Reference/chat/unhide_chat.md">unhide_chat()</a></code></td>
		<td>チャットの非表示を解除します</td>
	</tr>
	<tr>
		<td><code><a href="https://github.com/qvco/yaylib/blob/main/docs/API-Reference/chat/unpin_chat.md">unpin_chat()</a></code></td>
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
		<td><code><a href="https://github.com/qvco/yaylib/blob/main/docs/API-Reference/review/hello.md">hello()</a></code></td>
		<td>説明文</td>
	</tr>
	<tr>
		<td><code><a href="https://github.com/qvco/yaylib/blob/main/docs/API-Reference/review/hello.md">hello()</a></code></td>
		<td>説明文</td>
	</tr>
	<tr>
		<td><code><a href="https://github.com/qvco/yaylib/blob/main/docs/API-Reference/review/hello.md">hello()</a></code></td>
		<td>説明文</td>
	</tr>
	<tr>
		<td><code><a href="https://github.com/qvco/yaylib/blob/main/docs/API-Reference/review/hello.md">hello()</a></code></td>
		<td>説明文</td>
	</tr>
	<tr>
		<td><code><a href="https://github.com/qvco/yaylib/blob/main/docs/API-Reference/review/hello.md">hello()</a></code></td>
		<td>説明文</td>
	</tr>
	<tr>
		<td><code><a href="https://github.com/qvco/yaylib/blob/main/docs/API-Reference/review/hello.md">hello()</a></code></td>
		<td>説明文</td>
	</tr>
	<tr>
		<td><code><a href="https://github.com/qvco/yaylib/blob/main/docs/API-Reference/review/hello.md">hello()</a></code></td>
		<td>説明文</td>
	</tr>
	<tr>
		<td><code><a href="https://github.com/qvco/yaylib/blob/main/docs/API-Reference/review/hello.md">hello()</a></code></td>
		<td>説明文</td>
	</tr>
	<tr>
		<td><code><a href="https://github.com/qvco/yaylib/blob/main/docs/API-Reference/review/hello.md">hello()</a></code></td>
		<td>説明文</td>
	</tr>
	<tr>
		<td><code><a href="https://github.com/qvco/yaylib/blob/main/docs/API-Reference/review/hello.md">hello()</a></code></td>
		<td>説明文</td>
	</tr>
	<tr>
		<td><code><a href="https://github.com/qvco/yaylib/blob/main/docs/API-Reference/review/hello.md">hello()</a></code></td>
		<td>説明文</td>
	</tr>
	<tr>
		<td><code><a href="https://github.com/qvco/yaylib/blob/main/docs/API-Reference/review/hello.md">hello()</a></code></td>
		<td>説明文</td>
	</tr>
	<tr>
		<td><code><a href="https://github.com/qvco/yaylib/blob/main/docs/API-Reference/review/hello.md">hello()</a></code></td>
		<td>説明文</td>
	</tr>
	<tr>
		<td><code><a href="https://github.com/qvco/yaylib/blob/main/docs/API-Reference/review/hello.md">hello()</a></code></td>
		<td>説明文</td>
	</tr>
	<tr>
		<td><code><a href="https://github.com/qvco/yaylib/blob/main/docs/API-Reference/review/hello.md">hello()</a></code></td>
		<td>説明文</td>
	</tr>
	<tr>
		<td><code><a href="https://github.com/qvco/yaylib/blob/main/docs/API-Reference/review/hello.md">hello()</a></code></td>
		<td>説明文</td>
	</tr>
</table>

### 通話

<table>
    <tr>
		<th>メソッド</th>
		<th>概要</th>
	</tr>
	<tr>
		<td><code><a href="https://github.com/qvco/yaylib/blob/main/docs/API-Reference/call/hello.md">hello()</a></code></td>
		<td>説明文</td>
	</tr>
</table>

### ログイン

<table>
    <tr>
		<th>メソッド</th>
		<th>概要</th>
	</tr>
	<tr>
		<td><code><a href="https://github.com/qvco/yaylib/blob/main/docs/API-Reference/login/change_email.md">change_email()</a></code></td>
		<td>メールアドレスを変更します。</td>
	</tr>
</table>

### 通知

<table>
    <tr>
		<th>メソッド</th>
		<th>概要</th>
	</tr>
	<tr>
		<td><code><a href="https://github.com/qvco/yaylib/blob/main/docs/API-Reference/cassandra/hello.md">hello()</a></code></td>
		<td>説明文</td>
	</tr>
</table>

### その他

<table>
    <tr>
		<th>メソッド</th>
		<th>概要</th>
	</tr>
	<tr>
		<td><code><a href="https://github.com/qvco/yaylib/blob/main/docs/API-Reference/misc/hello.md">hello()</a></code></td>
		<td>説明文</td>
	</tr>
</table>

<p align="right">(<a href="#readme-top">トップに戻る</a>)</p>

## モデル

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
