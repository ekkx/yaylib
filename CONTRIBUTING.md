# コントリビューションガイド

yaylib への貢献を歓迎します。yaylib は Yay! API の Go / TypeScript /
Python クライアントで、**API スキーマから機械生成されるコード**と、その
上に載る**手書きレイヤ**の 2 層で構成されています。貢献の受け入れ方は
この 2 層で異なるため、まず下記をご確認ください。

## 機械生成ファイル（直接編集できません）

次のパスは API スキーマから機械生成され、まるごと再生成されます。
ファイル先頭に `Code generated; DO NOT EDIT.`（または同等の注記）が
入っています。ここへの手編集は次回生成時に上書きされて失われるため、
**これらに触れる PR はマージできません**。

- Go: `gen/`, `aliases.go`, `error_codes.go`, `host_routes.go`
- TypeScript: `packages/typescript/src/gen/`
- Python: `packages/python/yaylib/models/`,
  `packages/python/yaylib/api/`, `packages/python/yaylib/__init__.py`,
  `packages/python/yaylib/api_client.py`,
  `packages/python/yaylib/configuration.py`,
  `packages/python/yaylib/exceptions.py`,
  `packages/python/yaylib/rest.py`,
  `packages/python/yaylib/_host_routes.py`

**API の型・enum 値・ステータス・挙動がドキュメントと食い違う**のを
見つけた場合は、コードではなく **Issue** でご報告ください。サーバが
実際に返したリクエスト/レスポンス（再現手順・生のボディ）を添えて
いただけると、生成元を調整して次のリリースに反映できます。これは
最も価値の高い貢献です。

## 手書きレイヤ（PR を歓迎します）

次は手書きで保守しているコードで、通常の OSS と同じく PR を歓迎します。

- Go ルートの `client.go` / `auth.go` / `transport.go` / `signing.go` /
  `jwt.go` / `upload.go` / `event_stream.go` / `events.go` /
  `retry.go` / `session*.go` / `tokens`(各言語) / `errors.go` /
  `uuid.go` と対応する `*_test.go`
- TypeScript: `packages/typescript/src/` 直下の `gen/` 以外
- Python: `packages/python/yaylib/` 直下の手書きラッパ
  （`client.py` / `auth.py` / `transport.py` / `signing.py` /
  `upload.py` / `_image.py` / `event_stream.py` / `retry.py` /
  `session*.py` / `tokens.py` / `errors.py` / `api_response.py`）

受け入れの基準:

1. **言語横断契約を壊さない** — API サーフェスの形・認証フロー・署名・
   アップロード・イベントストリームの不変項は `PORTING.md` に定義され
   ています。これに反する変更は差し戻します。
2. **テストを通し、必要なら追加する** — 各言語のテストはネットワーク
   不要（hermetic）で動きます。挙動を変える PR にはテストを添えて
   ください。
3. **3 言語で挙動が揃うこと** — ある言語だけのその場しのぎではなく、
   3 言語で同じ挙動になる変更が望まれます。1 言語固有のバグ修正は
   その言語のみで構いません。

## 開発・テスト

すべてネットワーク不要で実行できます。

```bash
# Go
go test ./.

# TypeScript
cd packages/typescript && npm ci && npm test

# Python（Pillow を使うため venv 推奨）
cd packages/python
python -m venv .venv && . .venv/bin/activate
pip install -e . && pip install pytest
pytest -q
```

## PR の心得

- 1 PR は 1 つの目的に絞り、小さく保つ
- 挙動変更にはテストを添える
- マージは squash で行います

ライセンスは MIT です。提出いただいた貢献も同ライセンスで取り込まれます。
ご質問・雑談は Issue テンプレートに記載の Discord へどうぞ。
