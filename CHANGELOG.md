# Changelog

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
