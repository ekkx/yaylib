# yaylib - API ドキュメント

ドキュメントのイントロダクション

<details open>
    <summary>もくじ</summary>
    <ol>
        <li><a href="#インストール">インストール</a></li>
        <li><a href="#client">Client</a></li>
        <!-- <li><a href="async-client">Async Client</a></li> -->
        <li><a href="#api-一覧">API 一覧</a></li>
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
    </ol>
</details>

## インストール

インストール手順の説明

## Client

クライアントの使い方の説明

## API 一覧

### ログイン API

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

### ユーザー API

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
