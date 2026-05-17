// LoginWithEmail wrapper parity tests (PORTING.md §6).
// Exercises cache hit / miss / NoCache / save failure paths against an
// in-process HTTP server stub.

import { createServer, type Server } from "node:http";
import { mkdtempSync, rmSync } from "node:fs";
import type { AddressInfo } from "node:net";
import { tmpdir } from "node:os";
import { join } from "node:path";

import { Client } from "../src/client";
import { newSessionStore, FileSessionStore } from "../src/session_file";
import { MemorySessionStore, SessionSaveFailedError } from "../src/session";

type Handler = (path: string, method: string, body: string) => {
  status: number;
  body: string;
  headers?: Record<string, string>;
};

async function withServer<T>(
  handler: Handler,
  fn: (baseURL: string) => Promise<T>,
): Promise<T> {
  const server: Server = createServer((req, res) => {
    const chunks: Buffer[] = [];
    req.on("data", (c) => chunks.push(c));
    req.on("end", () => {
      const body = Buffer.concat(chunks).toString("utf8");
      const out = handler(req.url ?? "", req.method ?? "GET", body);
      res.writeHead(out.status, {
        "Content-Type": "application/json",
        ...(out.headers ?? {}),
      });
      res.end(out.body);
    });
  });
  await new Promise<void>((resolve) => server.listen(0, resolve));
  const addr = server.address() as AddressInfo;
  const baseURL = `http://127.0.0.1:${addr.port}`;
  try {
    return await fn(baseURL);
  } finally {
    await new Promise<void>((resolve) => server.close(() => resolve()));
  }
}

async function withTempDir<T>(fn: (dir: string) => Promise<T>): Promise<T> {
  const dir = mkdtempSync(join(tmpdir(), "yaylib-auth-"));
  try {
    return await fn(dir);
  } finally {
    rmSync(dir, { recursive: true, force: true });
  }
}

let failed = 0;
function assertEq<T>(name: string, got: T, want: T): void {
  if (got !== want) {
    failed++;
    // eslint-disable-next-line no-console
    console.log(`FAIL ${name}\n  got=${String(got)}\n  want=${String(want)}`);
  } else {
    // eslint-disable-next-line no-console
    console.log(`PASS ${name}`);
  }
}

// Scenario 1: fresh login (cache miss) issues HTTP + persists session.
// PORTING:S1
async function scenarioFreshLogin(): Promise<void> {
  let loginHits = 0;
  await withTempDir(async (dir) => {
    await withServer(
      (path, _method, body) => {
        if (path === "/v3/users/login_with_email") {
          loginHits++;
          const parsed = JSON.parse(body || "{}") as Record<string, unknown>;
          if (parsed.email !== "alice@example.com" || parsed.password !== "p4ss") {
            return { status: 400, body: '{"error_code":-1}' };
          }
          // Server returns tokens + user id.
          return {
            status: 200,
            body: JSON.stringify({
              access_token: "ACC",
              refresh_token: "REF",
              user_id: 42,
            }),
          };
        }
        return { status: 404, body: "" };
      },
      async (baseURL) => {
        const store = await newSessionStore(join(dir, "sessions.json"));
        const client = new Client({ baseURL, sessionStore: store });
        const res = await client
          .loginWithEmail()
          .email("alice@example.com")
          .password("p4ss")
          .execute();
        assertEq("fresh: server hit exactly once", loginHits, 1);
        assertEq("fresh: response carries access_token", res.accessToken, "ACC");
        assertEq("fresh: tokens activated on client", client.tokens().access, "ACC");
        assertEq("fresh: userID activated", client.userID, 42);
        // Persistence is observable: re-loading the store sees the session.
        const persisted = await store.load("alice@example.com");
        assertEq("fresh: session persisted", persisted.userID, 42);
      },
    );
  });
}

// Scenario 2: cache hit returns synthesized response without HTTP.
// PORTING:S2
async function scenarioCacheHit(): Promise<void> {
  let loginHits = 0;
  await withServer(
    () => {
      loginHits++;
      return { status: 500, body: "{}" };
    },
    async (baseURL) => {
      const store = new MemorySessionStore();
      await store.save({
        email: "bob@example.com",
        userID: 99,
        accessToken: "cached-acc",
        refreshToken: "cached-ref",
        deviceUUID: "11111111-1111-1111-1111-111111111111",
        updatedAt: new Date().toISOString(),
      });
      const client = new Client({ baseURL, sessionStore: store });
      const res = await client
        .loginWithEmail()
        .email("bob@example.com")
        .password("ignored")
        .execute();
      assertEq("cache hit: no server traffic", loginHits, 0);
      assertEq("cache hit: tokens applied", client.tokens().access, "cached-acc");
      assertEq("cache hit: userID applied", client.userID, 99);
      assertEq(
        "cache hit: deviceUUID restored",
        client.deviceUUID,
        "11111111-1111-1111-1111-111111111111",
      );
      assertEq("cache hit: response shape matches", res.userId, 99);
    },
  );
}

// Scenario 3: NoCache bypasses cache and forces HTTP login even when a
// session is stored.
async function scenarioNoCacheBypass(): Promise<void> {
  let loginHits = 0;
  await withServer(
    (path) => {
      if (path === "/v3/users/login_with_email") {
        loginHits++;
        return {
          status: 200,
          body: JSON.stringify({
            access_token: "FRESH",
            refresh_token: "FRESH-R",
            user_id: 7,
          }),
        };
      }
      return { status: 404, body: "" };
    },
    async (baseURL) => {
      const store = new MemorySessionStore();
      await store.save({
        email: "carol@example.com",
        userID: 1,
        accessToken: "OLD",
        refreshToken: "OLD-R",
        deviceUUID: "22222222-2222-2222-2222-222222222222",
        updatedAt: new Date().toISOString(),
      });
      const client = new Client({ baseURL, sessionStore: store });
      const res = await client
        .loginWithEmail()
        .email("carol@example.com")
        .password("p4ss")
        .noCache()
        .execute();
      assertEq("noCache: server hit once", loginHits, 1);
      assertEq("noCache: fresh tokens activated", client.tokens().access, "FRESH");
      assertEq("noCache: response from server", res.accessToken, "FRESH");
      // The persisted session is updated with the fresh tokens.
      const persisted = await store.load("carol@example.com");
      assertEq("noCache: persisted tokens replaced", persisted.accessToken, "FRESH");
    },
  );
}

// Scenario 4: persist failure during fresh login wraps the cause in
// SessionSaveFailedError; the in-memory tokens still get activated.
// PORTING:S4
async function scenarioSavePersistFailure(): Promise<void> {
  await withServer(
    (path) => {
      if (path === "/v3/users/login_with_email") {
        return {
          status: 200,
          body: JSON.stringify({
            access_token: "A",
            refresh_token: "R",
            user_id: 5,
          }),
        };
      }
      return { status: 404, body: "" };
    },
    async (baseURL) => {
      // Hand-built store that always fails save.
      const failingStore = {
        load: async () => {
          const e = new Error("not found") as Error & { name: string };
          e.name = "NoSessionError";
          throw e;
        },
        save: async () => {
          throw new Error("disk full");
        },
        delete: async () => undefined,
      };
      const client = new Client({ baseURL, sessionStore: failingStore });
      let caught: unknown = null;
      try {
        await client
          .loginWithEmail()
          .email("dave@example.com")
          .password("p4ss")
          .execute();
      } catch (err) {
        caught = err;
      }
      assertEq(
        "save failure: SessionSaveFailedError surfaced",
        caught instanceof SessionSaveFailedError,
        true,
      );
      assertEq("save failure: tokens still active", client.tokens().access, "A");
      assertEq("save failure: userID still active", client.userID, 5);
    },
  );
}

// Scenario 5: file store integration — second login from a fresh client
// instance hits the cache.
async function scenarioFileStoreRoundTrip(): Promise<void> {
  let loginHits = 0;
  await withTempDir(async (dir) => {
    await withServer(
      (path) => {
        if (path === "/v3/users/login_with_email") {
          loginHits++;
          return {
            status: 200,
            body: JSON.stringify({
              access_token: "T",
              refresh_token: "T-R",
              user_id: 11,
            }),
          };
        }
        return { status: 404, body: "" };
      },
      async (baseURL) => {
        const path = join(dir, "sessions.json");
        const store1 = new FileSessionStore(path);
        const client1 = new Client({ baseURL, sessionStore: store1 });
        await client1
          .loginWithEmail()
          .email("eve@example.com")
          .password("p4ss")
          .execute();
        assertEq("file store: first login hits server", loginHits, 1);

        // Fresh client + fresh store backed by the same file should
        // see the cached session.
        const store2 = new FileSessionStore(path);
        const client2 = new Client({ baseURL, sessionStore: store2 });
        await client2
          .loginWithEmail()
          .email("eve@example.com")
          .password("p4ss")
          .execute();
        assertEq("file store: second login is a cache hit", loginHits, 1);
        assertEq("file store: tokens restored", client2.tokens().access, "T");
        assertEq("file store: userID restored", client2.userID, 11);
      },
    );
  });
}

(async () => {
  await scenarioFreshLogin();
  await scenarioCacheHit();
  await scenarioNoCacheBypass();
  await scenarioSavePersistFailure();
  await scenarioFileStoreRoundTrip();
  process.exit(failed === 0 ? 0 : 1);
})().catch((e) => {
  // eslint-disable-next-line no-console
  console.error("test crashed:", e);
  process.exit(2);
});
