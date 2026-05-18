# examples

すべての言語で同じサンプルを用意しています。

- 認証 + タイムライン取得
- テキスト投稿
- イベントストリームの簡単なボット
- セッション永続化とエラー処理

`go/` · `typescript/` · `python/` にそれぞれ同じ題材で入っています。
いずれも `YAY_EMAIL` / `YAY_PASSWORD` を環境変数で渡して実行します。

```bash
go run ./examples/go/timeline
npx tsx examples/typescript/timeline.ts
python examples/python/timeline.py
```
