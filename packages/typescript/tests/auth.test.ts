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
import { codeOf } from "../src/errors";
import { ErrCodeUnauthorized, ErrCodeRequired2FA } from "../src/error_codes";

const DEFAULT_WEB_KEY =
  "e9f1ae4c4470f29a92c0168dc42b13637cc332692103f23e626bc2b016f66603";
const reject401 = (code = 401) => ({
  status: 401,
  body: JSON.stringify({ result: "error", message: "unauthorized", error_code: code }),
});

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
        const res = await client.loginWithEmail({ email: "alice@example.com", password: "p4ss" });
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
      const res = await client.loginWithEmail({ email: "bob@example.com", password: "ignored" });
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
      const res = await client.loginWithEmail({ email: "carol@example.com", password: "p4ss", noCache: true });
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
        await client.loginWithEmail({ email: "dave@example.com", password: "p4ss" });
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
        await client1.loginWithEmail({ email: "eve@example.com", password: "p4ss" });
        assertEq("file store: first login hits server", loginHits, 1);

        // Fresh client + fresh store backed by the same file should
        // see the cached session.
        const store2 = new FileSessionStore(path);
        const client2 = new Client({ baseURL, sessionStore: store2 });
        await client2.loginWithEmail({ email: "eve@example.com", password: "p4ss" });
        assertEq("file store: second login is a cache hit", loginHits, 1);
        assertEq("file store: tokens restored", client2.tokens().access, "T");
        assertEq("file store: userID restored", client2.userID, 11);
      },
    );
  });
}

// --- Browser-login fallback (PORTING.md §6.2) --------------------------------

async function scenarioFallbackToWeb(): Promise<void> {
  let seenKey: unknown;
  await withServer(
    (path, _method, body) => {
      if (path === "/v3/users/login_with_email") return reject401();
      if (path === "/v2/users/login_with_email") {
        seenKey = (JSON.parse(body || "{}") as Record<string, unknown>).api_key;
        return {
          status: 200,
          body: JSON.stringify({
            access_token: "WEBACC",
            refresh_token: "WEBREF",
            user_id: 7777,
          }),
        };
      }
      return { status: 404, body: "" };
    },
    async (baseURL) => {
      const client = new Client({ baseURL });
      const resp = await client.loginWithEmail({ email: "u@example.com", password: "pw" });
      assertEq("fallback: userId from web", resp.userId, 7777);
      assertEq("fallback: access token", client.tokens().access, "WEBACC");
      assertEq("fallback: web api_key used", seenKey, DEFAULT_WEB_KEY);
    },
  );
}

async function scenarioNoFallbackOn2FA(): Promise<void> {
  let v2Hits = 0;
  await withServer(
    (path) => {
      if (path === "/v3/users/login_with_email") return reject401(ErrCodeRequired2FA);
      if (path === "/v2/users/login_with_email") {
        v2Hits++;
        return { status: 200, body: "{}" };
      }
      return { status: 404, body: "" };
    },
    async (baseURL) => {
      const client = new Client({ baseURL });
      let caught: unknown;
      try {
        await client.loginWithEmail({ email: "u@example.com", password: "pw" });
      } catch (e) {
        caught = e;
      }
      assertEq("2FA: code is Required2FA", codeOf(caught), ErrCodeRequired2FA);
      assertEq("2FA: no v2 fallback attempted", v2Hits, 0);
      assertEq("2FA: tokens stay empty", client.tokens().access, "");
    },
  );
}

async function scenarioNoFallbackWhenApiKeyPinned(): Promise<void> {
  let v2Hits = 0;
  await withServer(
    (path) => {
      if (path === "/v3/users/login_with_email") return reject401();
      if (path === "/v2/users/login_with_email") {
        v2Hits++;
        return { status: 200, body: "{}" };
      }
      return { status: 404, body: "" };
    },
    async (baseURL) => {
      const client = new Client({ baseURL });
      let caught: unknown;
      try {
        await client.loginWithEmail({
          email: "u@example.com",
          password: "pw",
          apiKey: "caller-pinned-key",
        });
      } catch (e) {
        caught = e;
      }
      assertEq("pinned: code is Unauthorized", codeOf(caught), ErrCodeUnauthorized);
      assertEq("pinned: no v2 fallback attempted", v2Hits, 0);
    },
  );
}

async function scenarioFallbackDisabled(): Promise<void> {
  let v2Hits = 0;
  await withServer(
    (path) => {
      if (path === "/v3/users/login_with_email") return reject401();
      if (path === "/v2/users/login_with_email") {
        v2Hits++;
        return { status: 200, body: "{}" };
      }
      return { status: 404, body: "" };
    },
    async (baseURL) => {
      const client = new Client({ baseURL, webApiKey: "" });
      let caught: unknown;
      try {
        await client.loginWithEmail({ email: "u@example.com", password: "pw" });
      } catch (e) {
        caught = e;
      }
      assertEq("disabled: code is Unauthorized", codeOf(caught), ErrCodeUnauthorized);
      assertEq("disabled: no v2 fallback attempted", v2Hits, 0);
    },
  );
}

async function scenarioWebDecodeFailureSurfaced(): Promise<void> {
  await withServer(
    (path) => {
      if (path === "/v3/users/login_with_email") return reject401();
      if (path === "/v2/users/login_with_email") {
        return { status: 200, body: "{ this is not valid json" };
      }
      return { status: 404, body: "" };
    },
    async (baseURL) => {
      const client = new Client({ baseURL });
      let caught: unknown;
      try {
        await client.loginWithEmail({ email: "u@example.com", password: "pw" });
      } catch (e) {
        caught = e;
      }
      assertEq(
        "decode-fail: not masked as Unauthorized",
        codeOf(caught) !== ErrCodeUnauthorized,
        true,
      );
      assertEq("decode-fail: tokens stay empty", client.tokens().access, "");
    },
  );
}

(async () => {
  await scenarioFreshLogin();
  await scenarioCacheHit();
  await scenarioNoCacheBypass();
  await scenarioSavePersistFailure();
  await scenarioFileStoreRoundTrip();
  await scenarioFallbackToWeb();
  await scenarioNoFallbackOn2FA();
  await scenarioNoFallbackWhenApiKeyPinned();
  await scenarioFallbackDisabled();
  await scenarioWebDecodeFailureSurfaced();
  process.exit(failed === 0 ? 0 : 1);
})().catch((e) => {
  // eslint-disable-next-line no-console
  console.error("test crashed:", e);
  process.exit(2);
});
