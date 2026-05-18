<div align="center">

<img src="https://github.com/ekkx/yaylib/assets/77382767/45c45b21-d812-4cad-8f27-315ffef53201" alt="yaylib" height="160px">

### yaylib

**好きでつながるバーチャルワールド - Yay! の API ライブラリ**

<a href="https://github.com/ekkx/yaylib/tree/master/packages/typescript"><img src="https://img.shields.io/badge/TypeScript-3178C6?style=flat-square&logo=typescript&logoColor=white" alt="TypeScript"></a>&nbsp;
<a href="https://github.com/ekkx/yaylib/tree/master/packages/typescript"><img src="https://img.shields.io/badge/JavaScript-F7DF1E?style=flat-square&logo=javascript&logoColor=black" alt="JavaScript"></a>

<sub>あらゆる操作の自動化や、ボットの開発が可能です。</sub>

<a href="https://discord.gg/MEuBfNtqRN">Discord に参加</a>

</div>

Yay! の API を TypeScript / JavaScript から扱うための非公式 SDK です。

### インストール

```sh
npm install yaylib
```

Node.js 22 以上。

### クイックスタート

```ts
import { Client, NoreplyMode, chatRoomChannel } from "yaylib";

const client = new Client();

// ログイン（セッションは透過的にキャッシュされます）
await client.loginWithEmail({ email: "...", password: "..." });

// タイムラインを取得
const timeline = await client.getTimeline({
  noreplyMode: NoreplyMode.False,
  number: 20,
});
for (const post of timeline.posts) {
  console.log(post.id, post.text);
}

// 投稿する
await client.createPost({
  xJwt: client.generateXJwt(),
  postType: "text",
  text: "hello from yaylib",
});
```

すべてのオペレーションは `client.<operation>` として直接呼び出せます。

### イベントストリーム

```ts
const stream = client.openEventStream();
const sub = stream.subscribe(chatRoomChannel());

// 新着メッセージを受信
sub.onNewMessage((event) => {
  console.log("new message:", event);
});

await sub.closed;
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
