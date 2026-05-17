// Event-stream behavior parity (port of event_stream_parity_test.go),
// driven against the shared server's WebSocket contract over a real
// socket (Node's global WebSocket). The per-dial mode is selected by
// mockStreamClient. The server pushes one representative event per
// channel: ChatRoomChannel → chat_deleted{room_id:123},
// GroupUpdatesChannel → new_post{group_id:1}.
//
// The reconnect-budget mechanics (exhaustion, the stable-connection
// reset, the per-sub onDrop overflow) and the unsubscribe lifecycle stay
// local fixtures in event_stream.test.ts: the shared WS contract pushes a
// single event per subscribe and exposes no socket-fault knob, so those
// are not expressible here (§15 covers reconnect/reject/timeout, which
// are).
//
// Multiple-sub resubscribe (§15) also stays a local fixture
// (event_stream.test.ts:multipleSubsResubscribeAfterReconnect): the
// shared server's drop-after-confirm closes the socket after the *first*
// subscribe per connection, so observing both subs re-confirmed requires
// the resubscribe order to vary across reconnects. Go's reference passes
// only because Go map iteration is randomized; the TS SDK resubscribes
// in deterministic insertion order, so the multi-sub case cannot be
// satisfied against this contract and is left to the in-process fixture
// (which controls the fake server frame-by-frame). The single-sub
// reconnectAfterServerClose below already proves reconnect + re-subscribe
// per §15.

import { chatRoomChannel, groupUpdatesChannel } from "../src/channels";
import type { Event } from "../src/events";
import type { Subscription } from "../src/event_stream";
import { assert, mockStreamClient, run, skipUnlessMock } from "./parity_harness";

function awaitEvent(sub: Subscription, ms: number): Promise<Event> {
  return new Promise<Event>((resolve, reject) => {
    const t = setTimeout(() => reject(new Error(`no event within ${ms}ms`)), ms);
    sub.onEvent((ev) => {
      clearTimeout(t);
      resolve(ev);
    });
  });
}

async function expectThrow(name: string, fn: () => Promise<unknown>): Promise<void> {
  try {
    await fn();
    assert(name, false, "expected an error, got success");
  } catch {
    assert(name, true);
  }
}

// PORTING:S18,S25
async function subscribeAndReceiveEvent(): Promise<void> {
  const c = mockStreamClient("");
  c.setTokens("stub", "");
  const conn = await c.openEventStream({ reconnect: { disabled: true } });
  try {
    const sub = await conn.subscribe(chatRoomChannel());
    const ev = await awaitEvent(sub, 2000);
    assert(
      "stream: ChatRoomChannel delivers a ChatDeletedEvent",
      ev.type === "ChatDeletedEvent",
      `type=${ev.type}`,
    );
    assert(
      "stream: representative payload RoomID = 123",
      ev.type === "ChatDeletedEvent" && ev.roomId === 123,
      `ev=${JSON.stringify(ev)}`,
    );
  } finally {
    await conn.close();
  }
}

// PORTING:S21
async function rejectedSubscription(): Promise<void> {
  const c = mockStreamClient("reject");
  c.setTokens("stub", "");
  const conn = await c.openEventStream({ reconnect: { disabled: true } });
  try {
    await expectThrow("stream: rejected subscribe surfaces an error", () =>
      conn.subscribe(chatRoomChannel()),
    );
  } finally {
    await conn.close();
  }
}

// PORTING:S18,S25
async function multipleChannels(): Promise<void> {
  const c = mockStreamClient("");
  c.setTokens("stub", "");
  const conn = await c.openEventStream({ reconnect: { disabled: true } });
  try {
    const sub1 = await conn.subscribe(chatRoomChannel());
    const sub2 = await conn.subscribe(groupUpdatesChannel());
    let gotChat = false;
    let gotGroup = false;
    sub1.onEvent((ev) => {
      if (ev.type === "ChatDeletedEvent" && ev.roomId === 123) gotChat = true;
    });
    sub2.onEvent((ev) => {
      if (ev.type === "GroupUpdatedEvent" && ev.groupId === 1) gotGroup = true;
    });
    const ok = await waitUntil(() => gotChat && gotGroup, 3000);
    assert(
      "stream: both channels deliver their representative events",
      ok,
      `chat=${gotChat} group=${gotGroup}`,
    );
  } finally {
    await conn.close();
  }
}

// PORTING:S19
async function reconnectAfterServerClose(): Promise<void> {
  // drop-after-confirm closes the socket right after pushing the event,
  // so the client must reconnect and re-subscribe to keep receiving.
  // Seeing the representative event more than once proves the
  // reconnect + re-subscribe cycle.
  const c = mockStreamClient("drop-after-confirm");
  c.setTokens("stub", "");
  const conn = await c.openEventStream({
    reconnect: { initialDelayMs: 10, maxDelayMs: 30 },
  });
  try {
    const sub = await conn.subscribe(chatRoomChannel());
    let chat = 0;
    sub.onEvent((ev) => {
      if (ev.type === "ChatDeletedEvent") chat++;
    });
    const ok = await waitUntil(() => chat >= 2, 6000);
    assert(
      "stream: reconnect + re-subscribe re-delivers the event",
      ok,
      `chat=${chat}`,
    );
  } finally {
    await conn.close();
  }
}

// PORTING:S21
async function subscribeTimeout(): Promise<void> {
  const c = mockStreamClient("no-confirm");
  c.setTokens("stub", "");
  const conn = await c.openEventStream({
    reconnect: { disabled: true },
    subscribeTimeoutMs: 200,
  });
  try {
    await expectThrow("stream: unconfirmed subscribe times out", () =>
      conn.subscribe(chatRoomChannel()),
    );
  } finally {
    await conn.close();
  }
}

async function doneAndErrOnCleanClose(): Promise<void> {
  const c = mockStreamClient("");
  c.setTokens("stub", "");
  const conn = await c.openEventStream({ reconnect: { disabled: true } });
  assert(
    "stream: err() is undefined before close",
    conn.err() === undefined,
    `err=${String(conn.err())}`,
  );
  await conn.close();
  const done = await Promise.race([
    conn.done().then(() => true),
    new Promise<boolean>((r) => setTimeout(() => r(false), 2000)),
  ]);
  assert("stream: done() fires after close", done);
  assert(
    "stream: err() is undefined after a clean close",
    conn.err() === undefined,
    `err=${String(conn.err())}`,
  );
}

// PORTING:S22
async function wsDialDoesNotLeakBearer(): Promise<void> {
  const c = mockStreamClient("");
  // Tokens are set, but the WS dial must authenticate via the query
  // only. The shared server refuses the upgrade (401) if the dial
  // carries Authorization, so a successful open + subscribe proves no
  // bearer rode the handshake.
  c.setTokens("ACCESS_THAT_MUST_NOT_LEAK", "REF");
  const conn = await c.openEventStream({ reconnect: { disabled: true } });
  try {
    const sub = await conn.subscribe(chatRoomChannel());
    const ev = await awaitEvent(sub, 2000);
    assert(
      "stream: no-bearer dial opens and subscribes successfully",
      ev.type === "ChatDeletedEvent",
      `type=${ev.type}`,
    );
  } finally {
    await conn.close();
  }
}

const sleep = (ms: number) => new Promise<void>((r) => setTimeout(r, ms));
async function waitUntil(pred: () => boolean, timeoutMs: number): Promise<boolean> {
  const end = Date.now() + timeoutMs;
  while (Date.now() < end) {
    if (pred()) return true;
    await sleep(20);
  }
  return pred();
}

skipUnlessMock("event-stream parity", true);
run(async () => {
  await subscribeAndReceiveEvent();
  await rejectedSubscription();
  await multipleChannels();
  await reconnectAfterServerClose();
  await subscribeTimeout();
  await doneAndErrOnCleanClose();
  await wsDialDoesNotLeakBearer();
});
