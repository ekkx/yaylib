# Changelog

## v2.3.0

- メールログインを **web 優先** に変更しました。ブラウザ用ログイン
  エンドポイント（web 用キー）を主経路とし、アプリ用ログインは
  フォールバックになりました。アプリ用ログイン経路はサーバ側の端末検証
  により SDK が発行する端末 UUID を 401 で拒否するため、毎回の無駄な
  往復を避け、実際に通る経路を先に使います。
  - web 用キーを空にした場合、または呼び出し側が API キーを明示指定した
    場合はアプリ用ログインのみを使います（従来の脱出口を維持）
  - フォールバックは純粋な unauthorized のときだけ発火します。2FA が
    必要なケースや、成功応答（2xx）のデコード失敗はそのまま surface し、
    アプリ用ログインへは落としません
  - web 用ログインの失敗も、生成クライアントと同じ型のエラー（生の
    レスポンスボディ付き）で返るようになりました
  - 詳細な契約は `PORTING.md` の 6.2 節を参照
- セッションをディスクから直接復元する `LoadSession` が、セッションを
  確立したキー（`oauth_key`）を復元しない不具合を修正しました。従来は
  トークンとデバイス UUID・ログイン情報だけを適用していたため、web 発行
  セッションを `LoadSession` で復元するとトークン更新がアプリ用キーで
  行われ 401 になっていました。復元処理を一本化し、キャッシュ経由の復元
  と `LoadSession` の双方で確立キーを適用します（アプリ発行セッションの
  挙動は不変）。
- あわせて TypeScript / Python に公開の `loadSession` / `load_session`・
  `saveSession` / `save_session` を追加し（Go / Rust と同等）、セッションを
  ディスクから直接復元・保存できるようにしました。
- 全言語（Go / TypeScript / Python / Rust）に反映済みです。
- Yay! API 4.26.5 にモデル／クライアントを同期しました。あわせて、どの API
  からも参照されていなかった未使用の `I0A` enum を削除しています。

## v2.2.1

- 代替ログイン（v2.2.0 で追加）で確立したセッションのトークン更新が、
  正しい API キーで行われるように修正しました。従来はトークン更新が
  常に既定のキーで実行されていたため、代替ログインで得たセッションでは
  更新が 401 になることがありました。セッションを確立したキーを記憶し、
  更新時とセッション復元後にもそのキーを使います（署名処理には影響
  しません）。全言語（Go / TypeScript / Python / Rust）に反映済みです。

## v2.2.0

通話（カンファレンスコール）関連のモデルとエンドポイントを追加しました
（Go / TypeScript / Python / Rust 共通）。

- `ConferenceCall` / `RealmConferenceCall` モデルと `call_type`（`vdo` /
  `voice`）・`joinable_by`（参加可能範囲）enum を追加
- 通話・チャット・投稿の各サービスに関連エンドポイントを追加

また、メールログインが 401（unauthorized）で拒否された場合に、代替の
ログインエンドポイントへ自動でフォールバックするようになりました。

- 一部のアカウントは代替エンドポイントでのみ受け付けられます
- フォールバックは純粋な unauthorized のときだけ発火し、2FA が必要な
  ケースや、呼び出し側が明示的に API キーを指定した場合、フォールバック用
  キーを空にした場合は発火しません
- 詳細な契約は `PORTING.md` の 6.2 節を参照

## v2.1.0

通話・グループ・投稿まわりのエンドポイントを追加しました（Go / TypeScript / Python 共通）。

- `getCallParticipant` — 通話の参加者 UUID からユーザー情報を取得
- `getGroupCalls` — グループ通話の投稿一覧を取得
- `getCommunityCampaignStatus` / `getCommunityCampaignGroupLeaderboard` /
  `getCommunityCampaignUserLeaderboard` — コミュニティキャンペーン
  （グループ対抗ランキング）の参加資格・リーダーボードを取得
- `updateMuteKeyword` — ミュートキーワードの更新
- `unpinPost` — 投稿のピン留めを解除

## v2.0.1

- TypeScript / Python の `Client` に署名・JWT ヘルパーをメソッドとして追加
  （`generateXJwt` / `generateSignedVersion` / `generateSignedInfoAt` /
  `generateSignedInfo` / `generateCallActionSignature` /
  `validateCallActionSignature`）。クライアントが保持している API キーや
  UUID を呼び出し側で渡す必要がなくなり、Go と同じ使い心地になりました
  （従来のスタンドアロン関数も引き続き利用できます）。
- 実行できるサンプルを全言語ぶん追加し、`examples/` に集約しました。

> Go の API に変更はありません（このリリースでの新規追加は TypeScript /
> Python のみ）。

## v2.0.0

最初の安定版。Go / TypeScript / Python の 3 言語で同一の API 契約を提供します。
Python は旧 `1.x` 系からの全面書き換えで、**破壊的変更を含みます**。

### インストール

```sh
go get github.com/ekkx/yaylib/v2     # Go (module path に /v2 が必要)
npm install yaylib                   # TypeScript / JavaScript (ESM)
pip install yaylib                   # Python (>=3.10)
```

### ハイライト

- **フラットなクライアント**: 全オペレーションを `client.<operation>` で直接呼び出し（3 言語共通の命名）
- **透過的なセッション**: `login_with_email` がセッションをキャッシュし、トークンの自動リフレッシュに対応
- **署名ヘルパー**: `signed_info` / `signed_version` / `x_jwt` / 通話アクション署名を内蔵
- **リアルタイム**: 多重化された `EventStream`（複数チャンネル購読・自動再接続）
- **メディアアップロード**: 用途別アップロード API（アバター / 投稿画像 / 動画 ほか）
- **堅牢なトランスポート**: 自動リトライ（指数バックオフ + ジッタ）、必須ヘッダ自動付与
- **緩い wire 許容**: 数値↔文字列カーソル、未知 enum、2xx 揺れを吸収しデコード失敗を回避
- **エスケープハッチ**: 任意の呼び出しを生バイトで受け取る raw 実行

### 破壊的変更（Python `1.x` からの移行）

- クライアントは `async with` での利用を推奨（基盤接続を確実にクローズ）
- API バージョンは SDK バージョンと分離（`APIVersionName` 定数で別管理）
- メソッド名を動詞ファーストの統一命名へ変更

### バージョニング方針

major のみ 3 言語を揃え、minor / patch は言語ごとに独立して上げます。
