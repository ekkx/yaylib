// Event-stream parity tests (PORTING.md §10). Ports event_stream_test.go.
// The WebSocket transport is injected (webSocketFactory) so the suite is
// hermetic — no real socket, no `ws` dependency — exactly as the upload /
// retry suites stand up in-process HTTP servers.

import { createServer, type Server } from "node:http";
import type { AddressInfo } from "node:net";

import { Client } from "../src/client";
import { noopLogger } from "../src/logger";
import {
  type EventStream,
  type WebSocketLike,
  openEventStream,
} from "../src/event_stream";
import {
  chatRoomChannel,
  groupUpdatesChannel,
} from "../src/channels";

let failed = 0;
function assert(name: string, cond: boolean, detail = ""): void {
  if (cond) console.log(`PASS ${name}`);
  else {
    failed++;
    console.log(`FAIL ${name}${detail ? `\n  ${detail}` : ""}`);
  }
}
const sleep = (ms: number) => new Promise<void>((r) => setTimeout(r, ms));
async function waitUntil(pred: () => boolean, timeoutMs = 3000): Promise<boolean> {
  const end = Date.now() + timeoutMs;
  while (Date.now() < end) {
    if (pred()) return true;
    await sleep(10);
  }
  return pred();
}

// ---- in-memory WebSocket pair -----------------------------------------

type Frame = Record<string, unknown>;

class Session {
  ws!: FakeWS;
  private received: Frame[] = [];
  private waiters: ((f: Frame) => void)[] = [];
  recv(obj: Frame): void {
    const w = this.waiters.shift();
    if (w) w(obj);
    else this.received.push(obj);
  }
  next(): Promise<Frame> {
    const r = this.received.shift();
    return r ? Promise.resolve(r) : new Promise((res) => this.waiters.push(res));
  }
  send(obj: Frame): void {
    setTimeout(() => this.ws._msg(JSON.stringify(obj)), 0);
  }
  welcome(): void {
    this.send({ type: "welcome", sid: "sid" });
  }
  confirm(id: string): void {
    this.send({ identifier: id, type: "confirm_subscription" });
  }
  reject(id: string): void {
    this.send({ identifier: id, type: "reject_subscription" });
  }
  pushEvent(id: string, event: string, data: unknown): void {
    this.send({ identifier: id, message: { event, data } });
  }
  close(): void {
    setTimeout(() => this.ws._drop(), 0);
  }
}

class FakeWS implements WebSocketLike {
  onopen: ((ev: unknown) => void) | null = null;
  onmessage: ((ev: { data: unknown }) => void) | null = null;
  onclose: ((ev: { code?: number; reason?: string }) => void) | null = null;
  onerror: ((ev: unknown) => void) | null = null;
  private _closed = false;
  constructor(
    public url: string,
    session: Session,
    onConnect: (s: Session) => void,
  ) {
    session.ws = this;
    setTimeout(() => {
      if (this._closed) return;
      this.onopen?.({});
      onConnect(session);
    }, 0);
  }
  send(data: string): void {
    try {
      (this._session as Session).recv(JSON.parse(data));
    } catch {
      /* ignore */
    }
  }
  private _session!: Session;
  bind(s: Session): void {
    this._session = s;
  }
  close(): void {
    if (this._closed) return;
    this._closed = true;
    setTimeout(() => this.onclose?.({}), 0);
  }
  _msg(text: string): void {
    if (!this._closed) this.onmessage?.({ data: text });
  }
  _drop(): void {
    if (this._closed) return;
    this._closed = true;
    this.onclose?.({});
  }
}

class Hub {
  urls: string[] = [];
  connections = 0;
  constructor(private onConnect: (s: Session) => void) {}
  factory = (url: string): WebSocketLike => {
    this.urls.push(url);
    this.connections++;
    const session = new Session();
    const ws = new FakeWS(url, session, this.onConnect);
    ws.bind(session);
    return ws;
  };
}

function streamDeps(hub: Hub, token = "test-token") {
  return {
    eventStreamURL: "ws://cable.test",
    apiVersionName: "4.26",
    logger: noopLogger,
    getWebSocketToken: async () => token,
    webSocketFactory: hub.factory,
  };
}

const FAST_RECONNECT = { initialDelayMs: 20, maxDelayMs: 50 };

// ---- scenarios --------------------------------------------------------

async function subscribeAndReceiveEvent(): Promise<void> {
  const hub = new Hub((s) => {
    s.welcome();
    void (async () => {
      const msg = await s.next();
      assert("subscribe&recv: first command is subscribe", msg.command === "subscribe");
      s.confirm(msg.identifier as string);
      s.pushEvent(msg.identifier as string, "chat_deleted", { room_id: 123 });
    })();
  });
  const conn = await openEventStream(streamDeps(hub), { reconnect: { disabled: true } });
  const sub = await conn.subscribe(chatRoomChannel());
  let roomId = -1;
  sub.onChatDeleted((ev) => {
    roomId = ev.roomId;
  });
  await waitUntil(() => roomId === 123);
  assert("subscribe&recv: ChatDeletedEvent room_id=123", roomId === 123, `roomId=${roomId}`);
  await conn.close();
}

async function rejectedSubscription(): Promise<void> {
  const hub = new Hub((s) => {
    s.welcome();
    void (async () => {
      const msg = await s.next();
      s.reject(msg.identifier as string);
    })();
  });
  const conn = await openEventStream(streamDeps(hub), { reconnect: { disabled: true } });
  let threw = false;
  try {
    await conn.subscribe(chatRoomChannel());
  } catch {
    threw = true;
  }
  assert("rejected subscription: subscribe() rejects", threw);
  await conn.close();
}

async function multipleChannels(): Promise<void> {
  const hub = new Hub((s) => {
    s.welcome();
    void (async () => {
      for (let i = 0; i < 2; i++) {
        const msg = await s.next();
        s.confirm(msg.identifier as string);
      }
      s.pushEvent(`{"channel":"ChatRoomChannel"}`, "total_chat_request", { total_count: 7 });
      s.pushEvent(`{"channel":"GroupUpdatesChannel"}`, "new_post", { group_id: 42 });
    })();
  });
  const conn = await openEventStream(streamDeps(hub), { reconnect: { disabled: true } });
  let chat = 0;
  let group = 0;
  const s1 = await conn.subscribe(chatRoomChannel());
  s1.onTotalChatRequest((ev) => {
    chat = ev.totalCount;
  });
  const s2 = await conn.subscribe(groupUpdatesChannel());
  s2.onGroupUpdated((ev) => {
    group = ev.groupId;
  });
  await waitUntil(() => chat === 7 && group === 42);
  assert("multi-channel: chat=7 group=42", chat === 7 && group === 42, `chat=${chat} group=${group}`);
  await conn.close();
}

async function reconnectAfterServerClose(): Promise<void> {
  let attempts = 0;
  const hub = new Hub((s) => {
    const n = ++attempts;
    s.welcome();
    void (async () => {
      const msg = await s.next();
      s.confirm(msg.identifier as string);
      if (n === 1) {
        s.close();
        return;
      }
      s.pushEvent(msg.identifier as string, "chat_deleted", { room_id: 999 });
    })();
  });
  const conn = await openEventStream(streamDeps(hub), { reconnect: FAST_RECONNECT });
  const sub = await conn.subscribe(chatRoomChannel());
  let roomId = -1;
  sub.onChatDeleted((ev) => {
    roomId = ev.roomId;
  });
  await waitUntil(() => roomId === 999, 5000);
  assert("reconnect: event after re-subscribe (room_id=999)", roomId === 999, `roomId=${roomId}`);
  assert("reconnect: server saw >=2 connections", attempts >= 2, `attempts=${attempts}`);
  await conn.close();
}

async function unsubscribe(): Promise<void> {
  let sawUnsub = false;
  const hub = new Hub((s) => {
    s.welcome();
    void (async () => {
      const msg = await s.next();
      s.confirm(msg.identifier as string);
      for (;;) {
        const n = await s.next();
        if (n.command === "unsubscribe") {
          sawUnsub = true;
          return;
        }
      }
    })();
  });
  const conn = await openEventStream(streamDeps(hub), { reconnect: { disabled: true } });
  const sub = await conn.subscribe(chatRoomChannel());
  await sub.unsubscribe();
  await sub.done;
  await waitUntil(() => sawUnsub);
  assert("unsubscribe: server saw unsubscribe command", sawUnsub);
  assert("unsubscribe: sub.done resolved", sub.closed);
  await conn.close();
}

async function subscribeTimeout(): Promise<void> {
  const hub = new Hub((s) => {
    s.welcome();
    void s.next(); // read the subscribe but never confirm
  });
  const conn = await openEventStream(streamDeps(hub), {
    reconnect: { disabled: true },
    subscribeTimeoutMs: 150,
  });
  let threw = false;
  try {
    await conn.subscribe(chatRoomChannel());
  } catch {
    threw = true;
  }
  assert("subscribe timeout: rejects when never confirmed", threw);
  await conn.close();
}

async function doneAndErrOnCleanClose(): Promise<void> {
  const hub = new Hub((s) => s.welcome());
  const conn = await openEventStream(streamDeps(hub), { reconnect: { disabled: true } });
  assert("clean close: err() undefined before close", conn.err() === undefined);
  await conn.close();
  await conn.done();
  assert("clean close: err() undefined after caller close", conn.err() === undefined);
}

async function errAfterReconnectExhausted(): Promise<void> {
  const hub = new Hub((s) => {
    s.welcome();
    s.close();
  });
  const conn = await openEventStream(streamDeps(hub), {
    reconnect: { maxAttempts: 1, initialDelayMs: 10, maxDelayMs: 20 },
  });
  await Promise.race([conn.done(), sleep(4000)]);
  assert("reconnect exhausted: err() is non-undefined", conn.err() !== undefined, String(conn.err()));
  await conn.close();
}

async function onDropFiresWhenBufferFull(): Promise<void> {
  let dropped = 0;
  const ident = `{"channel":"ChatRoomChannel"}`;
  const hub = new Hub((s) => {
    s.welcome();
    void (async () => {
      const msg = await s.next();
      s.confirm(msg.identifier as string);
      for (let i = 0; i < 5; i++) s.pushEvent(ident, "chat_deleted", { room_id: i });
    })();
  });
  const conn = await openEventStream(streamDeps(hub), {
    reconnect: { disabled: true },
    eventBuffer: 1,
    onDrop: () => {
      dropped++;
    },
  });
  // Subscribe but never attach a listener — the buffer fills and overflow
  // events must drop via onDrop.
  await conn.subscribe(chatRoomChannel());
  await waitUntil(() => dropped > 0);
  assert("onDrop: fired >=1 when buffer full + no listener", dropped >= 1, `dropped=${dropped}`);
  await conn.close();
}

async function stableConnectionResetsAttemptBudget(): Promise<void> {
  let attempts = 0;
  const hub = new Hub((s) => {
    const n = ++attempts;
    s.welcome();
    void (async () => {
      const msg = await s.next();
      s.confirm(msg.identifier as string);
      if (n <= 2) {
        await sleep(80); // stay alive past the (shrunk) stability threshold
        s.close();
      }
      // n>=3 stays open
    })();
  });
  const conn = await openEventStream(streamDeps(hub), {
    reconnect: { maxAttempts: 1, initialDelayMs: 10, maxDelayMs: 20 },
    stableThresholdMs: 50,
  });
  await conn.subscribe(chatRoomChannel());
  await waitUntil(() => attempts >= 3, 4000);
  assert(
    "stable reset: >=3 connections (stable conn reset the failure budget)",
    attempts >= 3,
    `attempts=${attempts}`,
  );
  await conn.close();
}

async function multipleSubsResubscribeAfterReconnect(): Promise<void> {
  const identChat = `{"channel":"ChatRoomChannel"}`;
  const identGroup = `{"channel":"GroupUpdatesChannel"}`;
  let attempts = 0;
  const hub = new Hub((s) => {
    const n = ++attempts;
    s.welcome();
    void (async () => {
      const seen = new Set<string>();
      while (seen.size < 2) {
        const msg = await s.next();
        if (msg.command !== "subscribe") continue;
        seen.add(msg.identifier as string);
        s.confirm(msg.identifier as string);
      }
      if (n === 1) {
        s.close();
        return;
      }
      s.pushEvent(identChat, "chat_deleted", { room_id: 11 });
      s.pushEvent(identGroup, "new_post", { group_id: 22 });
    })();
  });
  const conn = await openEventStream(streamDeps(hub), { reconnect: FAST_RECONNECT });
  const subChat = await conn.subscribe(chatRoomChannel());
  const subGroup = await conn.subscribe(groupUpdatesChannel());
  let chat = -1;
  let group = -1;
  subChat.onChatDeleted((ev) => {
    chat = ev.roomId;
  });
  subGroup.onGroupUpdated((ev) => {
    group = ev.groupId;
  });
  await waitUntil(() => chat === 11 && group === 22, 5000);
  assert(
    "resubscribe: both subs re-wired after reconnect (chat=11 group=22)",
    chat === 11 && group === 22,
    `chat=${chat} group=${group}`,
  );
  await conn.close();
}

// ---- Client-level: ws-token auth + dial does not leak Bearer ----------

async function withHTTP(
  handler: (path: string, auth: string) => { status: number; body: string },
  fn: (baseURL: string, lastAuth: () => string) => Promise<void>,
): Promise<void> {
  let lastAuth = "";
  const server: Server = createServer((req, res) => {
    const u = new URL(req.url ?? "", "http://x");
    if (u.pathname === "/v1/users/ws_token") lastAuth = req.headers["authorization"] ?? "";
    const out = handler(u.pathname, req.headers["authorization"] ?? "");
    res.writeHead(out.status, { "Content-Type": "application/json" }).end(out.body);
  });
  await new Promise<void>((r) => server.listen(0, r));
  const baseURL = `http://127.0.0.1:${(server.address() as AddressInfo).port}`;
  try {
    await fn(baseURL, () => lastAuth);
  } finally {
    server.closeAllConnections?.();
    await new Promise<void>((r) => server.close(() => r()));
  }
}

async function clientWsTokenAuthAndNoBearerLeak(): Promise<void> {
  await withHTTP(
    () => ({ status: 200, body: JSON.stringify({ token: "ws-token-xyz" }) }),
    async (baseURL) => {
      let dialURL = "";
      const factory = (url: string): WebSocketLike => {
        dialURL = url;
        const session = new Session();
        const ws = new FakeWS(url, session, (s) => s.welcome());
        ws.bind(session);
        return ws;
      };
      const c = new Client({
        baseURL,
        eventStreamURL: "ws://cable.test",
        webSocketFactory: factory,
        retryPolicy: { maxAttempts: 0, baseDelay: 1, maxDelay: 1, retryOnPOST: false },
      });
      c.setTokens("SECRET-ACCESS-TOKEN", "REF");
      const conn = await c.openEventStream({ reconnect: { disabled: true } });
      assert("client ws: dial URL has token query", dialURL.includes("token=ws-token-xyz"), dialURL);
      assert("client ws: dial URL has app_version query", dialURL.includes("app_version="), dialURL);
      assert(
        "client ws: dial URL does NOT leak the Bearer",
        !dialURL.includes("SECRET-ACCESS-TOKEN"),
        dialURL,
      );
      await conn.close();
    },
  );
  // Separately assert the ws_token fetch was authenticated.
  await withHTTP(
    () => ({ status: 200, body: JSON.stringify({ token: "t" }) }),
    async (baseURL, lastAuth) => {
      const c = new Client({
        baseURL,
        webSocketFactory: (url) => {
          const session = new Session();
          const ws = new FakeWS(url, session, (s) => s.welcome());
          ws.bind(session);
          return ws;
        },
        retryPolicy: { maxAttempts: 0, baseDelay: 1, maxDelay: 1, retryOnPOST: false },
      });
      c.setTokens("BEARERTOK", "REF");
      const conn = await c.openEventStream({ reconnect: { disabled: true } });
      assert(
        "client ws: ws_token request had Authorization: Bearer",
        lastAuth() === "Bearer BEARERTOK",
        lastAuth(),
      );
      await conn.close();
    },
  );
}

async function clientUnauthenticated401(): Promise<void> {
  await withHTTP(
    () => ({ status: 401, body: JSON.stringify({ error_code: -1, message: "unauthorized" }) }),
    async (baseURL) => {
      const c = new Client({
        baseURL,
        webSocketFactory: (url) => {
          const session = new Session();
          const ws = new FakeWS(url, session, (s) => s.welcome());
          ws.bind(session);
          return ws;
        },
        retryPolicy: { maxAttempts: 0, baseDelay: 1, maxDelay: 1, retryOnPOST: false },
      });
      let threw = false;
      try {
        await c.openEventStream({ reconnect: { disabled: true } });
      } catch {
        threw = true;
      }
      assert("client ws: unauthenticated open surfaces the 401", threw);
    },
  );
}

(async () => {
  await subscribeAndReceiveEvent();
  await rejectedSubscription();
  await multipleChannels();
  await reconnectAfterServerClose();
  await unsubscribe();
  await subscribeTimeout();
  await doneAndErrOnCleanClose();
  await errAfterReconnectExhausted();
  await onDropFiresWhenBufferFull();
  await stableConnectionResetsAttemptBudget();
  await multipleSubsResubscribeAfterReconnect();
  await clientWsTokenAuthAndNoBearerLeak();
  await clientUnauthenticated401();
  process.exit(failed === 0 ? 0 : 1);
})().catch((e) => {
  console.error("test crashed:", e);
  process.exit(2);
});
