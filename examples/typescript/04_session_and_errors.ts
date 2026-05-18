// セッション永続化とエラー処理の例。
//
// newSessionStore を渡すと、ログイン状態がファイルに保存され、次回以降は
// キャッシュされたセッションが再利用されます（毎回ログインしない）。
// 非 2xx は APIError として throw され、codeOf / errorResponseOf で
// 中身を判定できます。
//
//   YAY_EMAIL=... YAY_PASSWORD=... npx tsx examples/typescript/04_session_and_errors.ts
import {
  Client,
  NoreplyMode,
  newSessionStore,
  codeOf,
  errorResponseOf,
  asAPIError,
} from "yaylib";

const store = await newSessionStore("yay-session.json");
const client = new Client({ sessionStore: store });

// 2 回目以降の実行では保存済みセッションが使われます。
await client.loginWithEmail({
  email: process.env.YAY_EMAIL!,
  password: process.env.YAY_PASSWORD!,
});

try {
  const timeline = await client.getTimeline({
    noreplyMode: NoreplyMode.False,
    number: 20,
  });
  console.log("posts:", timeline.posts?.length ?? 0);
} catch (err) {
  console.log("error code:", codeOf(err)); // grpc 的なコード分岐
  const res = errorResponseOf(err); // ban_until / retry_in など
  if (res) console.log("message:", res.message);
  const apiErr = await asAPIError(err); // 生のステータス / ボディ
  console.log("status:", apiErr.response?.status);
}
