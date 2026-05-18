# examples

3 言語の実行できるサンプルです。いずれも `YAY_EMAIL` / `YAY_PASSWORD`
を環境変数で渡します。

| | 内容 |
|---|---|
| 認証 + タイムライン取得 | `go/timeline` · `typescript/01_timeline.ts` · `python/01_timeline.py` |
| テキスト投稿 | `go/post` · `typescript/02_post.ts` · `python/02_post.py` |
| イベントストリームの簡単なボット | `go/eventstream` · `typescript/03_event_stream.ts` · `python/03_event_stream.py` |
| セッション永続化とエラー処理 | `go/session_and_errors` · `typescript/04_session_and_errors.ts` · `python/04_session_and_errors.py` |

```bash
# Go
YAY_EMAIL=... YAY_PASSWORD=... go run ./examples/go/timeline

# TypeScript
YAY_EMAIL=... YAY_PASSWORD=... npx tsx examples/typescript/01_timeline.ts

# Python
YAY_EMAIL=... YAY_PASSWORD=... python examples/python/01_timeline.py
```
