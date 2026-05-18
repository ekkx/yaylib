<div align="center">

<img src="https://github.com/ekkx/yaylib/assets/77382767/45c45b21-d812-4cad-8f27-315ffef53201" alt="yaylib" height="160px">

### yaylib

**好きでつながるバーチャルワールド - Yay! の API ライブラリ**

<a href="https://github.com/ekkx/yaylib"><img src="https://img.shields.io/badge/Go-00ADD8?style=flat-square&logo=go&logoColor=white" alt="Go"></a>&nbsp;
<a href="https://github.com/ekkx/yaylib/tree/master/packages/typescript"><img src="https://img.shields.io/badge/TypeScript-3178C6?style=flat-square&logo=typescript&logoColor=white" alt="TypeScript"></a>&nbsp;
<a href="https://github.com/ekkx/yaylib/tree/master/packages/python"><img src="https://img.shields.io/badge/Python-3776AB?style=flat-square&logo=python&logoColor=white" alt="Python"></a>

<sub>あらゆる操作の自動化や、ボットの開発が可能です。</sub>

<a href="https://discord.gg/MEuBfNtqRN">Discord に参加</a>

</div>

Yay! の API を [Go](https://github.com/ekkx/yaylib) / [TypeScript](https://github.com/ekkx/yaylib/tree/master/packages/typescript) / [Python](https://github.com/ekkx/yaylib/tree/master/packages/python) から扱うための非公式 SDK です。

### インストール

```sh
go get github.com/ekkx/yaylib/v2   # Go
npm install yaylib                 # TypeScript
pip install yaylib                 # Python
```

### クイックスタート (Go)

```go
package main

import (
	"context"

	"github.com/ekkx/yaylib/v2"
)

func main() {
	ctx := context.Background()
	client := yaylib.NewClient()

	// ログイン（セッションは透過的にキャッシュされます）
	client.LoginWithEmail(ctx).Email("...").Password("...").Execute()

	// タイムラインを取得
	tl, _, _ := client.GetTimeline(ctx, yaylib.NoreplyModeFalse).Number(20).Execute()
	_ = tl

	// 投稿する
	client.CreatePost(ctx).
		XJwt(client.GenerateXJwt()).
		PostType("text").
		Text("hello from yaylib").
		Execute()
}
```

すべてのオペレーションは `client.<Operation>` として直接呼び出せます。

### イベントストリーム

```go
stream, _ := client.OpenEventStream(ctx)
sub, _ := stream.Subscribe(ctx, yaylib.ChatRoomChannel())

for ev := range sub.Events() {
	switch e := ev.(type) {
	case *yaylib.NewMessageEvent:
		// 新着メッセージを受信
		fmt.Println("new message:", e)
	}
}
```

### サンプル

実行できるサンプルを [`examples/`](https://github.com/ekkx/yaylib/tree/master/examples) に用意しています。

- `examples/go/timeline` — 認証 + タイムライン取得
- `examples/go/post` — テキスト投稿
- `examples/go/eventstream` — イベントストリームの簡単なボット
- `examples/go/session_and_errors` — セッション永続化とエラー処理

他言語のサンプルも [`examples/`](https://github.com/ekkx/yaylib/tree/master/examples) にまとまっています。

```bash
YAY_EMAIL=... YAY_PASSWORD=... go run ./examples/go/timeline
```

### ⚖️ ライセンス

<p align="center">
  <a href="https://github.com/ekkx">
    <img src="https://github.com/ekkx/yaylib/assets/77382767/5d6aef18-5d98-4c9b-9f54-791308b393af" width="256" height="256">
  </a>
</p>

<p align="center">
  <strong>MIT © <a href="https://github.com/ekkx">ekkx</a></strong>
</p>
