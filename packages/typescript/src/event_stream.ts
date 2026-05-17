// PORTING.md §10: a multiplexed real-time channel — one underlying
// WebSocket against wss://cable.yay.space, N subscriptions on top. The
// handshake / subscribe / dispatch flow is Rails-ActionCable-style JSON;
// the wire constants in §10.1 (dial URL, channel identifier strings,
// frame inventory, event→DTO map) are pattern-matched server-side and
// replicated here verbatim — drift produces silent failures.
//
// Unlike Go's blocking read loop this port is event-driven (the platform
// WebSocket is callback-based), but the reconnect state machine,
// stability threshold, jittered backoff and OnDrop backpressure hook
// mirror event_stream.go exactly.

import { type Channel } from "./channels";
import { type Event, decodeEvent } from "./events";
import { type Logger } from "./logger";

export * from "./channels";
export * from "./events";

// WebSocketLike is the slice of the standard WebSocket the stream needs.
// The platform global satisfies it; tests inject an in-memory pair so the
// suite stays hermetic.
export interface WebSocketLike {
  send(data: string): void;
  close(code?: number, reason?: string): void;
  onopen: ((ev: unknown) => void) | null;
  onmessage: ((ev: { data: unknown }) => void) | null;
  onclose: ((ev: { code?: number; reason?: string }) => void) | null;
  onerror: ((ev: unknown) => void) | null;
}

export interface ReconnectPolicy {
  // Disabled stops the stream after the first disconnect; subscriptions
  // terminate with the connection error.
  disabled?: boolean;
  // MaxAttempts caps reconnect attempts (excluding the initial dial). 0
  // means unlimited — opposite of RetryPolicy, because an unattended
  // event stream wants to keep recovering.
  maxAttempts?: number;
  // InitialDelay is the first backoff sleep; each attempt doubles up to
  // MaxDelay, with full jitter.
  initialDelayMs?: number;
  maxDelayMs?: number;
}

export interface EventStreamOptions {
  reconnect?: ReconnectPolicy;
  // Per-subscription buffer used only while no listener is attached.
  // Overflow drops the incoming event and fires onDrop. Defaults to 64.
  eventBuffer?: number;
  // Caps how long subscribe() waits for confirm_subscription. Default 10s.
  subscribeTimeoutMs?: number;
  // Fires when a subscription's buffer is full and an event is dropped —
  // keeps backpressure observable. Keep it fast.
  onDrop?: (ev: Event) => void;
  // How long a fresh connection must stay up before the failure budget
  // resets (PORTING.md §10: default 30s). Exposed so tests can shrink it.
  stableThresholdMs?: number;
  // Caps the dial + welcome handshake. Default 30s.
  connectTimeoutMs?: number;
}

export interface EventStreamDeps {
  eventStreamURL: string;
  apiVersionName: string;
  logger: Logger;
  // Fetches the short-lived ws token (GET /v1/users/ws_token via the
  // wrapped client, so an unauthenticated client surfaces the 401 here).
  getWebSocketToken(): Promise<string>;
  webSocketFactory?: (url: string) => WebSocketLike;
}

interface ResolvedOptions {
  reconnect: Required<ReconnectPolicy>;
  eventBuffer: number;
  subscribeTimeoutMs: number;
  onDrop?: (ev: Event) => void;
  stableThresholdMs: number;
  connectTimeoutMs: number;
}

function resolveOptions(o: EventStreamOptions): ResolvedOptions {
  return {
    reconnect: {
      disabled: o.reconnect?.disabled ?? false,
      maxAttempts: o.reconnect?.maxAttempts ?? 0,
      initialDelayMs:
        o.reconnect?.initialDelayMs && o.reconnect.initialDelayMs > 0
          ? o.reconnect.initialDelayMs
          : 500,
      maxDelayMs:
        o.reconnect?.maxDelayMs && o.reconnect.maxDelayMs > 0
          ? o.reconnect.maxDelayMs
          : 30_000,
    },
    eventBuffer: o.eventBuffer && o.eventBuffer > 0 ? o.eventBuffer : 64,
    subscribeTimeoutMs:
      o.subscribeTimeoutMs && o.subscribeTimeoutMs > 0 ? o.subscribeTimeoutMs : 10_000,
    onDrop: o.onDrop,
    stableThresholdMs:
      o.stableThresholdMs && o.stableThresholdMs > 0 ? o.stableThresholdMs : 30_000,
    connectTimeoutMs:
      o.connectTimeoutMs && o.connectTimeoutMs > 0 ? o.connectTimeoutMs : 30_000,
  };
}

function defaultWebSocketFactory(url: string): WebSocketLike {
  const WS = (globalThis as { WebSocket?: new (u: string) => unknown }).WebSocket;
  if (!WS) {
    throw new Error(
      "yaylib: no WebSocket implementation available; pass webSocketFactory in ClientOptions",
    );
  }
  return new WS(url) as unknown as WebSocketLike;
}

interface WireFrame {
  identifier?: string;
  type?: string;
  message?: unknown;
}

function parseFrame(data: unknown): WireFrame | null {
  let text: string;
  if (typeof data === "string") text = data;
  else if (data instanceof Uint8Array) text = new TextDecoder().decode(data);
  else if (data instanceof ArrayBuffer) text = new TextDecoder().decode(new Uint8Array(data));
  else if (data && typeof (data as { toString?: () => string }).toString === "function")
    text = String(data);
  else return null;
  try {
    return JSON.parse(text) as WireFrame;
  } catch {
    return null;
  }
}

// Subscription is one active channel subscription. Read events with the
// callback API (sub.onNewMessage(...), sub.onEvent(...)); the buffer +
// onDrop hook only matter while no listener is attached.
export class Subscription {
  /** @internal */ readonly ident: string;
  private readonly _stream: EventStream;
  private readonly _bufferCap: number;
  private readonly _buffer: Event[] = [];
  private readonly _listeners = new Set<(ev: Event) => void>();

  private _confirmed = false;
  private _confirmResolve!: () => void;
  private _confirmReject!: (err: Error) => void;
  /** @internal */ readonly confirm: Promise<void>;

  private _closed = false;
  private _doneResolve!: () => void;
  readonly done: Promise<void>;

  /** @internal */
  constructor(stream: EventStream, ident: string, bufferCap: number) {
    this._stream = stream;
    this.ident = ident;
    this._bufferCap = bufferCap;
    this.confirm = new Promise<void>((resolve, reject) => {
      this._confirmResolve = resolve;
      this._confirmReject = reject;
    });
    // A confirm that rejects must not become an unhandled rejection if the
    // caller only ever awaits via awaitConfirm's race.
    this.confirm.catch(() => undefined);
    this.done = new Promise<void>((resolve) => {
      this._doneResolve = resolve;
    });
  }

  /** Register a listener for every event on this subscription. */
  onEvent(fn: (ev: Event) => void): this {
    this._listeners.add(fn);
    // Flush anything buffered before the first listener attached.
    if (this._buffer.length > 0) {
      const pending = this._buffer.splice(0, this._buffer.length);
      for (const ev of pending) fn(ev);
    }
    return this;
  }

  private _typed<K extends Event["type"]>(
    kind: K,
    fn: (ev: Extract<Event, { type: K }>) => void,
  ): this {
    return this.onEvent((ev) => {
      if (ev.type === kind) fn(ev as Extract<Event, { type: K }>);
    });
  }

  onNewMessage(fn: (ev: Extract<Event, { type: "NewMessageEvent" }>) => void): this {
    return this._typed("NewMessageEvent", fn);
  }
  onVideoProcessed(fn: (ev: Extract<Event, { type: "VideoProcessedEvent" }>) => void): this {
    return this._typed("VideoProcessedEvent", fn);
  }
  onChatDeleted(fn: (ev: Extract<Event, { type: "ChatDeletedEvent" }>) => void): this {
    return this._typed("ChatDeletedEvent", fn);
  }
  onTotalChatRequest(fn: (ev: Extract<Event, { type: "TotalChatRequestEvent" }>) => void): this {
    return this._typed("TotalChatRequestEvent", fn);
  }
  onUnsubscribed(fn: (ev: Extract<Event, { type: "UnsubscribedEvent" }>) => void): this {
    return this._typed("UnsubscribedEvent", fn);
  }
  onGroupUpdated(fn: (ev: Extract<Event, { type: "GroupUpdatedEvent" }>) => void): this {
    return this._typed("GroupUpdatedEvent", fn);
  }
  onCallFinished(fn: (ev: Extract<Event, { type: "CallFinishedEvent" }>) => void): this {
    return this._typed("CallFinishedEvent", fn);
  }
  onRaw(fn: (ev: Extract<Event, { type: "RawEvent" }>) => void): this {
    return this._typed("RawEvent", fn);
  }

  get closed(): boolean {
    return this._closed;
  }

  /** @internal — dispatcher delivers here; never blocks on slow consumers. */
  deliver(ev: Event): void {
    if (this._closed) return;
    if (this._listeners.size > 0) {
      for (const fn of this._listeners) fn(ev);
      return;
    }
    if (this._buffer.length >= this._bufferCap) {
      // Buffer full and nobody listening — drop, surface via onDrop.
      this._stream.onDropEvent(ev);
      return;
    }
    this._buffer.push(ev);
  }

  /** @internal */
  confirmOK(): void {
    if (this._confirmed) return;
    this._confirmed = true;
    this._confirmResolve();
  }

  /** @internal */
  confirmFail(err: Error): void {
    if (this._confirmed) return;
    this._confirmed = true;
    this._confirmReject(err);
  }

  /** @internal — terminate because the stream died or rejected the sub. */
  terminate(err: Error): void {
    if (this._closed) return;
    this._closed = true;
    this.confirmFail(err);
    this._doneResolve();
  }

  /**
   * Unsubscribe sends an unsubscribe command and terminates the
   * subscription. Safe to call multiple times.
   */
  async unsubscribe(): Promise<void> {
    if (this._closed) return;
    this._closed = true;
    this._stream.removeSub(this.ident);
    try {
      this._stream.writeCommand("unsubscribe", this.ident);
    } catch {
      // Best effort — the server may already be gone.
    }
    this._doneResolve();
  }

  /** @internal — clean shutdown driven by the parent stream's close(). */
  closeQuietly(): void {
    if (this._closed) return;
    this._closed = true;
    this.confirmFail(new Error("event stream closed"));
    this._doneResolve();
  }
}

export class EventStream {
  private readonly _deps: EventStreamDeps;
  private readonly _opts: ResolvedOptions;
  private readonly _base: string;
  private readonly _subs = new Map<string, Subscription>();

  private _ws: WebSocketLike | null = null;
  private _closed = false;
  private _attempt = 0;
  private _connectedAt = 0;
  private _finalErr: Error | undefined;
  private _sleepCancel: (() => void) | null = null;

  private _doneResolve!: () => void;
  private readonly _done: Promise<void>;

  /** @internal — use openEventStream(). */
  constructor(deps: EventStreamDeps, opts: ResolvedOptions) {
    this._deps = deps;
    this._opts = opts;
    this._base = deps.eventStreamURL || "wss://cable.yay.space";
    this._done = new Promise<void>((resolve) => {
      this._doneResolve = resolve;
    });
  }

  /** @internal */
  onDropEvent(ev: Event): void {
    this._opts.onDrop?.(ev);
  }

  /**
   * Done resolves once the stream has fully shut down — via close() or
   * after a permanent failure (reconnect exhausted). Pair with err().
   */
  done(): Promise<void> {
    return this._done;
  }

  /**
   * err returns the terminal error after done() resolves: undefined for a
   * caller-initiated close, the cause otherwise. Returns undefined before
   * shutdown.
   */
  err(): Error | undefined {
    if (this._closed) return undefined;
    return this._finalErr;
  }

  private _setFinalErr(err: Error | undefined): void {
    this._finalErr = err;
  }

  private _finalize(err: Error | undefined): void {
    this._setFinalErr(err);
    this._terminateAllSubs(err ?? new Error("event stream closed"));
    this._doneResolve();
  }

  private _wsURL(token: string): string {
    // Dial carries NO Authorization header — auth is exclusively the
    // token query param (PORTING.md §10.1).
    const u = new URL(this._base);
    u.searchParams.set("token", token);
    u.searchParams.set("app_version", this._deps.apiVersionName);
    return u.toString();
  }

  /** @internal — performs one dial + welcome handshake. */
  async connect(timeoutMs: number): Promise<void> {
    const token = await this._deps.getWebSocketToken();
    const factory = this._deps.webSocketFactory ?? defaultWebSocketFactory;
    const ws = factory(this._wsURL(token));

    await new Promise<void>((resolve, reject) => {
      let settled = false;
      const timer = setTimeout(() => {
        if (settled) return;
        settled = true;
        try {
          ws.close();
        } catch {
          /* ignore */
        }
        reject(new Error("event stream: connect timed out"));
      }, timeoutMs);
      const finish = (err?: Error) => {
        if (settled) return;
        settled = true;
        clearTimeout(timer);
        if (err) {
          try {
            ws.close();
          } catch {
            /* ignore */
          }
          reject(err);
        } else {
          resolve();
        }
      };
      ws.onopen = () => {
        /* wait for the welcome frame */
      };
      ws.onerror = () => finish(new Error("event stream: socket error during handshake"));
      ws.onclose = () => finish(new Error("event stream: closed during handshake"));
      ws.onmessage = (ev) => {
        const f = parseFrame(ev.data);
        if (!f) return;
        if (f.type === "welcome") finish();
        else if (f.type === "disconnect")
          finish(new Error("event stream: server disconnect during welcome"));
        // Any preceding pings are ignored.
      };
    });

    const old = this._ws;
    this._ws = ws;
    if (old) {
      try {
        old.close();
      } catch {
        /* ignore */
      }
    }
    this._connectedAt = Date.now();
    ws.onmessage = (ev) => this._dispatch(ev.data);
    ws.onerror = null;
    ws.onclose = () => this._onSocketClosed(new Error("event stream: connection closed"));
  }

  private _onSocketClosed(err: Error): void {
    if (this._closed) return;
    this._setFinalErr(err);
    if (this._opts.reconnect.disabled) {
      this._finalize(err);
      return;
    }
    // A connection that stayed up past the stability threshold resets the
    // failure budget; a flap (welcome + immediate drop) counts toward
    // MaxAttempts so a misconfigured peer can't reconnect forever.
    if (this._connectedAt && Date.now() - this._connectedAt >= this._opts.stableThresholdMs) {
      this._attempt = 0;
    }
    void this._reconnect();
  }

  private _backoff(attempt: number): number {
    const { initialDelayMs, maxDelayMs } = this._opts.reconnect;
    let exp = initialDelayMs * 2 ** (attempt - 1);
    if (!Number.isFinite(exp) || exp <= 0 || exp > maxDelayMs) exp = maxDelayMs;
    if (exp <= 0) return 0;
    let delay = Math.floor(Math.random() * exp) + initialDelayMs;
    if (delay > maxDelayMs) delay = maxDelayMs;
    return delay;
  }

  private _sleep(ms: number): Promise<boolean> {
    if (ms <= 0) return Promise.resolve(!this._closed);
    return new Promise<boolean>((resolve) => {
      const timer = setTimeout(() => {
        this._sleepCancel = null;
        resolve(!this._closed);
      }, ms);
      this._sleepCancel = () => {
        clearTimeout(timer);
        this._sleepCancel = null;
        resolve(false);
      };
    });
  }

  private async _reconnect(): Promise<void> {
    while (!this._closed) {
      this._attempt++;
      const max = this._opts.reconnect.maxAttempts;
      if (max > 0 && this._attempt > max) {
        const e = new Error(
          `event stream: reconnect exhausted after ${max} attempts: ${this._finalErr?.message ?? "unknown"}`,
        );
        this._finalize(e);
        return;
      }
      const delayMs = this._backoff(this._attempt);
      this._deps.logger.debug("reconnecting event stream", {
        event: "ws_reconnect",
        attempt: this._attempt,
        delay_ms: delayMs,
      });
      const ok = await this._sleep(delayMs);
      if (!ok) {
        this._finalize(this._finalErr);
        return;
      }
      try {
        await this.connect(this._opts.connectTimeoutMs);
      } catch (e) {
        this._setFinalErr(e instanceof Error ? e : new Error(String(e)));
        continue;
      }
      // Reconnected. Re-subscribe everything; the new socket's onclose
      // re-enters this state machine if it drops again.
      this._resubscribeAll();
      return;
    }
    this._finalize(this._finalErr);
  }

  private _dispatch(data: unknown): void {
    const f = parseFrame(data);
    if (!f) return;
    switch (f.type) {
      case "welcome":
      case "ping":
        return;
      case "disconnect":
        // Server-requested disconnect — close the socket so onclose
        // drives the reconnect logic.
        try {
          this._ws?.close(1000, "server disconnect");
        } catch {
          /* ignore */
        }
        return;
      case "confirm_subscription": {
        const sub = f.identifier ? this._subs.get(f.identifier) : undefined;
        sub?.confirmOK();
        return;
      }
      case "reject_subscription": {
        const sub = f.identifier ? this._subs.get(f.identifier) : undefined;
        if (sub) {
          const e = new Error("event stream: subscription rejected");
          sub.confirmFail(e);
          sub.terminate(e);
          this._subs.delete(f.identifier as string);
        }
        return;
      }
    }
    // Anything else is an event keyed by identifier.
    const msg = f.message;
    if (!msg || typeof msg !== "object") return;
    const payload = msg as { event?: string; data?: unknown; message?: unknown };
    const name = payload.event;
    const body = payload.data ?? payload.message;
    if (!name || body == null) return;
    const ev = decodeEvent(name, body);
    const sub = f.identifier ? this._subs.get(f.identifier) : undefined;
    sub?.deliver(ev);
  }

  /** @internal */
  removeSub(ident: string): void {
    this._subs.delete(ident);
  }

  /** @internal — best effort; a re-subscribe runs on the next reconnect. */
  writeCommand(cmd: "subscribe" | "unsubscribe", identifier: string): void {
    const ws = this._ws;
    if (!ws) throw new Error("event stream: not connected");
    ws.send(JSON.stringify({ command: cmd, identifier }));
  }

  private _resubscribeAll(): void {
    for (const sub of this._subs.values()) {
      try {
        this.writeCommand("subscribe", sub.ident);
      } catch {
        /* best effort */
      }
    }
  }

  private _terminateAllSubs(err: Error): void {
    const subs = [...this._subs.values()];
    this._subs.clear();
    for (const s of subs) s.terminate(err);
  }

  /**
   * Subscribe registers the channel and resolves once the server confirms
   * (or rejects). Concurrent calls for the same channel share the
   * underlying Subscription.
   */
  async subscribe(channel: Channel): Promise<Subscription> {
    if (this._closed) throw new Error("event stream closed");
    const ident = channel.identifier;
    const existing = this._subs.get(ident);
    if (existing) return this._awaitConfirm(existing, false);

    const sub = new Subscription(this, ident, this._opts.eventBuffer);
    this._subs.set(ident, sub);
    try {
      this.writeCommand("subscribe", ident);
    } catch (err) {
      this._subs.delete(ident);
      throw new Error(
        `event stream: subscribe ${ident}: ${err instanceof Error ? err.message : String(err)}`,
      );
    }
    return this._awaitConfirm(sub, true);
  }

  private async _awaitConfirm(sub: Subscription, owned: boolean): Promise<Subscription> {
    const cleanup = () => {
      if (owned) this._subs.delete(sub.ident);
    };
    let timer: ReturnType<typeof setTimeout> | undefined;
    const timeout = new Promise<never>((_, reject) => {
      timer = setTimeout(
        () => reject(new Error(`event stream: subscribe ${sub.ident}: timed out`)),
        this._opts.subscribeTimeoutMs,
      );
    });
    const streamDied = this._done.then(() => {
      throw this._finalErr ?? new Error("event stream closed");
    });
    try {
      await Promise.race([sub.confirm, streamDied, timeout]);
      return sub;
    } catch (err) {
      cleanup();
      throw err instanceof Error ? err : new Error(String(err));
    } finally {
      if (timer) clearTimeout(timer);
    }
  }

  /**
   * Close terminates the stream, all subscriptions and the socket. Safe
   * to call multiple times.
   */
  async close(): Promise<void> {
    if (this._closed) {
      await this._done;
      return;
    }
    this._closed = true;
    if (this._sleepCancel) this._sleepCancel();
    try {
      this._ws?.close(1000, "");
    } catch {
      /* ignore */
    }
    const subs = [...this._subs.values()];
    this._subs.clear();
    for (const s of subs) s.closeQuietly();
    this._doneResolve();
    await this._done;
  }
}

/** @internal — Client.openEventStream() is the public entry point. */
export async function openEventStream(
  deps: EventStreamDeps,
  opts: EventStreamOptions = {},
): Promise<EventStream> {
  const resolved = resolveOptions(opts);
  const stream = new EventStream(deps, resolved);
  await stream.connect(resolved.connectTimeoutMs);
  return stream;
}
