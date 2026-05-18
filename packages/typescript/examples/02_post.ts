// テキスト投稿の例。投稿系のエンドポイントは X-Jwt を必要とします
// （client.generateXJwt() で取得し、そのまま渡します）。
//
//   YAY_EMAIL=... YAY_PASSWORD=... npx tsx examples/02_post.ts
import { Client, generateXJwt } from "yaylib";

const client = new Client();

await client.loginWithEmail({
  email: process.env.YAY_EMAIL!,
  password: process.env.YAY_PASSWORD!,
});

const post = await client.createPost({
  xJwt: generateXJwt({ apiVersionKey: client.apiVersionKey }),
  postType: "text",
  text: "hello from yaylib",
});

console.log("posted:", post.id);
