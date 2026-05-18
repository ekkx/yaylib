// イベントストリームで新着チャットメッセージを受信する簡単なボット。
// Ctrl-C で終了します。
//
//   YAY_EMAIL=... YAY_PASSWORD=... npx tsx examples/03_event_stream.ts
import { Client, chatRoomChannel } from "yaylib";

const client = new Client();

await client.loginWithEmail({
  email: process.env.YAY_EMAIL!,
  password: process.env.YAY_PASSWORD!,
});

const stream = await client.openEventStream();
const sub = await stream.subscribe(chatRoomChannel());
console.log("listening for events (Ctrl-C to stop)");

sub
  .onNewMessage((ev) => console.log("new message:", ev))
  .onEvent((ev) => console.log("event:", ev?.type));

process.on("SIGINT", () => {
  void stream.close().then(() => process.exit(0));
});

await sub.done;
