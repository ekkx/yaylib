// 認証してタイムラインを取得する最小例。
//
//   YAY_EMAIL=... YAY_PASSWORD=... npx tsx examples/typescript/01_timeline.ts
import { Client, NoreplyMode } from "yaylib";

const client = new Client();

await client.loginWithEmail({
  email: process.env.YAY_EMAIL!,
  password: process.env.YAY_PASSWORD!,
});

const timeline = await client.getTimeline({
  noreplyMode: NoreplyMode.False,
  number: 20,
});

for (const post of timeline.posts ?? []) {
  console.log(post.id, post.text);
}
