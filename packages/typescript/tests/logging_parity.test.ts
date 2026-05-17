// Logging behavior parity (port of logging_test.go's parity scenario).
// PORTING.md §12.2 + §15: with a logger injected the
// 401 → refresh → persist-failure path emits the contracted events and
// never leaks a token / password / API key into any record. Driven by
// the shared server's expired-token scenario (401 until the token
// endpoint is hit, then a fresh TokenResponse, then the happy path —
// exactly the refresh chain). The default-is-silent assertion is a pure
// local fixture and stays in logging.test.ts.

import type { Logger, LogFields } from "../src/logger";
import type { Session, SessionStore } from "../src/session";
import { assert, mockClient, mockGet, run, skipUnlessMock } from "./parity_harness";

interface Rec {
  level: string;
  msg: string;
  fields: LogFields;
}

class CaptureLogger implements Logger {
  records: Rec[] = [];
  private push(level: string, msg: string, fields?: LogFields): void {
    this.records.push({ level, msg, fields: fields ?? {} });
  }
  debug(m: string, f?: LogFields): void {
    this.push("debug", m, f);
  }
  info(m: string, f?: LogFields): void {
    this.push("info", m, f);
  }
  warn(m: string, f?: LogFields): void {
    this.push("warn", m, f);
  }
  error(m: string, f?: LogFields): void {
    this.push("error", m, f);
  }
  byEvent(event: string): Rec | undefined {
    return this.records.find((r) => r.fields["event"] === event);
  }
}

// SessionStore whose save always rejects, to drive token_persist_fail
// during the post-refresh persist.
const failingStore: SessionStore = {
  load: async (): Promise<Session> => {
    throw new Error("no session");
  },
  save: async (): Promise<void> => {
    throw new Error("disk full");
  },
  delete: async (): Promise<void> => {},
};

async function refreshPersistFailEventAndRedaction(): Promise<void> {
  const h = new CaptureLogger();
  const c = mockClient("expired-token", {
    retryPolicy: { maxAttempts: 0, baseDelay: 1, maxDelay: 1, retryOnPOST: false },
    sessionStore: failingStore,
    logger: h,
  });
  c.setTokens("STALE", "REF");
  c.setLoginIdentity("foo@example.com", 1234); // enable the persist branch

  const r = await mockGet(c);
  assert(
    "logging: request completed via refresh chain (status 200)",
    r.status === 200,
    `status=${r.status}`,
  );

  // Contracted events present.
  const persist = h.byEvent("token_persist_fail");
  assert("logging: token_persist_fail record emitted", persist !== undefined);
  assert(
    "logging: token_persist_fail level = warn",
    persist?.level === "warn",
    `level=${persist?.level}`,
  );
  const refresh = h.byEvent("token_refresh");
  assert(
    "logging: token_refresh ok record present",
    refresh !== undefined && refresh.fields["outcome"] === "ok",
    `rec=${JSON.stringify(refresh)}`,
  );
  assert(
    "logging: http_request debug record emitted",
    h.byEvent("http_request") !== undefined,
  );

  // Redaction: no record (message or any field) carries a secret —
  // neither the stale tokens we seeded nor the fresh ones the server
  // issued on refresh, nor the API key.
  const fresh = c.tokens();
  const banned = ["STALE", "REF", fresh.access, fresh.refresh, c.apiKey];
  let leaked = "";
  for (const rec of h.records) {
    let hay = rec.msg;
    for (const [k, v] of Object.entries(rec.fields)) {
      hay += " " + k + "=" + String(v);
    }
    for (const b of banned) {
      if (b && hay.includes(b)) {
        leaked = `record ${JSON.stringify(rec.msg)} leaked ${JSON.stringify(b)}: ${JSON.stringify(rec.fields)}`;
      }
    }
  }
  assert("logging: no banned value reached the logger", leaked === "", leaked);
}

skipUnlessMock("logging parity");
run(async () => {
  await refreshPersistFailEventAndRedaction();
});
