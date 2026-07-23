// LoginWithEmail wrapper parity tests (PORTING.md §6 / §6.2).
// Exercises the web-first fresh-login flow, cache hit / miss / NoCache /
// save failure, the OAuth-key round-trip through the session store, and
// the public loadSession / saveSession activation — mirroring the Go
// auth_test.go scenarios against an in-process HTTP server stub.

import { createServer, type Server } from "node:http";
import { mkdtempSync, rmSync } from "node:fs";
import type { AddressInfo } from "node:net";
import { tmpdir } from "node:os";
import { join } from "node:path";

import { Client } from "../src/client";
import { newSessionStore, FileSessionStore } from "../src/session_file";
import { MemorySessionStore, SessionSaveFailedError } from "../src/session";
import { APIError, codeOf, errorResponseOf } from "../src/errors";
import {
  ErrCodeUnauthorized,
  ErrCodeRequired2FA,
  ErrCodeUserTemporaryBanned,
} from "../src/error_codes";

const DEFAULT_WEB_KEY =
  "e9f1ae4c4470f29a92c0168dc42b13637cc332692103f23e626bc2b016f66603";
const APP_KEY =
  "ccd59ee269c01511ba763467045c115779fcae3050238a252f1bd1a4b65cfec6";
const basic = (key: string) => `Basic ${Buffer.from(key).toString("base64")}`;

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
    // Drop keep-alive sockets (bare-fetch web login + generated app login)
    // so server.close resolves promptly instead of hanging.
    server.closeAllConnections?.();
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

// webFirstServer models the web-first flow. The browser login (v2) is
// primary: it succeeds when webOK is true and otherwise rejects with 401.
// The app login (v3) is the fallback: it succeeds when appOK, else 401.
// Hits on each endpoint and the api_key each request carried are recorded
// so tests can assert which path ran and under which key.
interface WebFirstState {
  webHits: number;
  appHits: number;
  seenWebKey?: string;
  seenAppKey?: string;
}
function webFirstServer(
  state: WebFirstState,
  webOK: boolean,
  appOK: boolean,
): Handler {
  return (path, _method, body) => {
    const key = () => (JSON.parse(body || "{}") as Record<string, unknown>).api_key as string;
    if (path.startsWith("/v2/users/login_with_email")) {
      state.webHits++;
      state.seenWebKey = key();
      if (!webOK) return reject401();
      return {
        status: 200,
        body: JSON.stringify({
          access_token: "WEBACC",
          refresh_token: "WEBREF",
          user_id: 7777,
        }),
      };
    }
    if (path.startsWith("/v3/users/login_with_email")) {
      state.appHits++;
      state.seenAppKey = key();
      if (!appOK) return reject401();
      return {
        status: 200,
        body: JSON.stringify({
          access_token: "APPACC",
          refresh_token: "APPREF",
          user_id: 4242,
        }),
      };
    }
    return { status: 404, body: "" };
  };
}

// Scenario 1: fresh login (cache miss) issues HTTP + persists session.
// Web is primary, so the browser (v2) endpoint answers and the persisted
// session tracks the web key. PORTING:S1
async function scenarioFreshLogin(): Promise<void> {
  let loginHits = 0;
  await withTempDir(async (dir) => {
    await withServer(
      (path, _method, body) => {
        if (path.endsWith("/login_with_email")) {
          loginHits++;
          const parsed = JSON.parse(body || "{}") as Record<string, unknown>;
          if (parsed.email !== "alice@example.com" || parsed.password !== "p4ss") {
            return { status: 400, body: '{"error_code":-1}' };
          }
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
        const persisted = await store.load("alice@example.com");
        assertEq("fresh: session persisted", persisted.userID, 42);
        assertEq("fresh: persisted oauth_key is web key", persisted.oauthKey, DEFAULT_WEB_KEY);
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
      if (path.endsWith("/login_with_email")) {
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
      if (path.endsWith("/login_with_email")) {
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
        if (path.endsWith("/login_with_email")) {
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

// --- Web-first fresh-login flow (PORTING.md §6.2) ---------------------------

// Web is the primary path: it succeeds directly and the app (v3) endpoint
// is never touched; the session tracks the web key.
async function scenarioWebPrimarySucceeds(): Promise<void> {
  const state: WebFirstState = { webHits: 0, appHits: 0 };
  await withServer(webFirstServer(state, true, true), async (baseURL) => {
    const client = new Client({ baseURL });
    const resp = await client.loginWithEmail({ email: "u@example.com", password: "pw" });
    assertEq("web primary: userId from web", resp.userId, 7777);
    assertEq("web primary: access token", client.tokens().access, "WEBACC");
    assertEq("web primary: refresh token", client.tokens().refresh, "WEBREF");
    assertEq("web primary: web api_key used", state.seenWebKey, DEFAULT_WEB_KEY);
    // eslint-disable-next-line @typescript-eslint/no-explicit-any
    assertEq("web primary: oauthKey is web key", (client as any)._oauthKeySnapshot(), DEFAULT_WEB_KEY);
    assertEq("web primary: app endpoint not hit", state.appHits, 0);
  });
}

// Primary web login rejected with a plain 401 → fall back to the app (v3)
// endpoint under the app key; the session then tracks the app key.
async function scenarioFallsBackToAppOn401(): Promise<void> {
  const state: WebFirstState = { webHits: 0, appHits: 0 };
  await withServer(webFirstServer(state, false, true), async (baseURL) => {
    const client = new Client({ baseURL });
    const resp = await client.loginWithEmail({ email: "u@example.com", password: "pw" });
    assertEq("fallback: userId from app", resp.userId, 4242);
    assertEq("fallback: access token", client.tokens().access, "APPACC");
    assertEq("fallback: app api_key used", state.seenAppKey, APP_KEY);
    // eslint-disable-next-line @typescript-eslint/no-explicit-any
    assertEq("fallback: oauthKey is app key", (client as any)._oauthKeySnapshot(), APP_KEY);
    assertEq("fallback: web tried once", state.webHits, 1);
    assertEq("fallback: app tried once", state.appHits, 1);
  });
}

// Web switched off (webApiKey="") → straight to the app endpoint, no v2
// request; an app 401 surfaces unchanged.
async function scenarioWebDisabledUsesAppOnly(): Promise<void> {
  const state: WebFirstState = { webHits: 0, appHits: 0 };
  await withServer(webFirstServer(state, true, false), async (baseURL) => {
    const client = new Client({ baseURL, webApiKey: "" });
    let caught: unknown;
    try {
      await client.loginWithEmail({ email: "u@example.com", password: "pw" });
    } catch (e) {
      caught = e;
    }
    assertEq("web disabled: code is Unauthorized", codeOf(caught), ErrCodeUnauthorized);
    assertEq("web disabled: web endpoint not hit", state.webHits, 0);
    assertEq("web disabled: tokens stay empty", client.tokens().access, "");
  });
}

// Both the primary web login and the app fallback fail → the primary (web)
// 401 is surfaced.
async function scenarioBothFailSurfacesWeb401(): Promise<void> {
  const state: WebFirstState = { webHits: 0, appHits: 0 };
  await withServer(webFirstServer(state, false, false), async (baseURL) => {
    const client = new Client({ baseURL });
    let caught: unknown;
    try {
      await client.loginWithEmail({ email: "u@example.com", password: "pw" });
    } catch (e) {
      caught = e;
    }
    assertEq("both fail: code is Unauthorized", codeOf(caught), ErrCodeUnauthorized);
    assertEq("both fail: web tried", state.webHits, 1);
    assertEq("both fail: app tried", state.appHits, 1);
    assertEq("both fail: tokens stay empty", client.tokens().access, "");
  });
}

// A 2FA-required rejection is also HTTP 401 but carries ErrCodeRequired2FA;
// the app fallback must NOT fire and the primary web 2FA surfaces.
async function scenarioNoFallbackOn2FA(): Promise<void> {
  let appHits = 0;
  await withServer(
    (path) => {
      if (path.startsWith("/v2/users/login_with_email")) return reject401(ErrCodeRequired2FA);
      if (path.startsWith("/v3/users/login_with_email")) {
        appHits++;
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
      assertEq("2FA: no app fallback attempted", appHits, 0);
      assertEq("2FA: tokens stay empty", client.tokens().access, "");
    },
  );
}

// A caller-pinned api_key skips the web-first path; login goes straight to
// the app endpoint under the pinned key and the browser (v2) is never hit.
async function scenarioNoFallbackWhenApiKeyPinned(): Promise<void> {
  let webHits = 0;
  let seenAppKey: unknown;
  await withServer(
    (path, _method, body) => {
      if (path.startsWith("/v2/users/login_with_email")) {
        webHits++;
        return { status: 200, body: "{}" };
      }
      if (path.startsWith("/v3/users/login_with_email")) {
        seenAppKey = (JSON.parse(body || "{}") as Record<string, unknown>).api_key;
        return reject401();
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
      assertEq("pinned: web endpoint not hit", webHits, 0);
      assertEq("pinned: app used the pinned key", seenAppKey, "caller-pinned-key");
    },
  );
}

// A primary web 2xx response that fails to decode is a genuine login result:
// surfaced as a typed APIError carrying the raw body, and (not being
// Unauthorized) it does not fall back to the app login.
async function scenarioWebDecodeFailureSurfaced(): Promise<void> {
  let appHits = 0;
  await withServer(
    (path) => {
      if (path.startsWith("/v2/users/login_with_email")) {
        return { status: 200, body: "{ this is not valid json" };
      }
      if (path.startsWith("/v3/users/login_with_email")) {
        appHits++;
        return {
          status: 200,
          body: JSON.stringify({ access_token: "APPACC", refresh_token: "APPREF", user_id: 1 }),
        };
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
      assertEq("decode-fail: no app fallback", appHits, 0);
      assertEq("decode-fail: is APIError", caught instanceof APIError, true);
      assertEq(
        "decode-fail: raw body carried on error",
        caught instanceof APIError && caught.bodyText().includes("not valid json"),
        true,
      );
      assertEq("decode-fail: tokens stay empty", client.tokens().access, "");
    },
  );
}

// A primary web login rejected with a NON-Unauthorized typed error (a
// temporary ban carrying ban_until) surfaces directly as a typed APIError —
// no app fallback, and errorResponseOf reads the parsed fields / raw body.
async function scenarioWebPrimaryTypedErrorSurfaced(): Promise<void> {
  let appHits = 0;
  await withServer(
    (path) => {
      if (path.startsWith("/v2/users/login_with_email")) {
        return {
          status: 403,
          body: JSON.stringify({
            result: "error",
            message: "temporarily banned",
            error_code: ErrCodeUserTemporaryBanned,
            ban_until: 1893456000,
          }),
        };
      }
      if (path.startsWith("/v3/users/login_with_email")) {
        appHits++;
        return { status: 200, body: JSON.stringify({ access_token: "APPACC", user_id: 1 }) };
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
      assertEq("ban: code is UserTemporaryBanned", codeOf(caught), ErrCodeUserTemporaryBanned);
      assertEq("ban: no app fallback (non-401)", appHits, 0);
      const er = errorResponseOf(caught);
      assertEq("ban: ban_until parsed from web error", er?.banUntil, 1893456000);
      assertEq("ban: is APIError", caught instanceof APIError, true);
      assertEq(
        "ban: raw web body accessible",
        caught instanceof APIError && caught.bodyText().includes("temporarily banned"),
        true,
      );
    },
  );
}

// signed_info is keyed on apiKey, NOT the session's OAuth key. Setting the
// OAuth key to the web key must not change the signature.
async function scenarioSignedInfoIgnoresOAuthKey(): Promise<void> {
  const ref = new Client({ apiKey: "app-key", deviceUUID: "uuid-shared" });
  const web = new Client({ apiKey: "app-key", deviceUUID: "uuid-shared" });
  web.setOAuthKey("web-key"); // establish a web session
  assertEq(
    "signed_info: ignores oauth_key (stays on apiKey)",
    web.generateSignedInfoAt(1700000000).value,
    ref.generateSignedInfoAt(1700000000).value,
  );
}

// --- OAuth refresh uses the session's key (PORTING.md §6.2) ------------------

// Server that captures the Authorization header on the oauth/token path.
async function withRefreshServer<T>(
  loginHandler: (path: string) => { status: number; body: string } | null,
  fn: (baseURL: string, getOauthAuth: () => string) => Promise<T>,
): Promise<T> {
  let oauthAuth = "";
  const server: Server = createServer((req, res) => {
    const chunks: Buffer[] = [];
    req.on("data", (c) => chunks.push(c));
    req.on("end", () => {
      const path = req.url ?? "";
      res.setHeader("Content-Type", "application/json");
      if (path.includes("/api/v1/oauth/token")) {
        oauthAuth = req.headers["authorization"] ?? "";
        res.writeHead(200);
        res.end(JSON.stringify({ access_token: "A2", refresh_token: "R2", id: 1 }));
        return;
      }
      const out = loginHandler(path);
      if (out) {
        res.writeHead(out.status);
        res.end(out.body);
      } else {
        res.writeHead(404);
        res.end("");
      }
    });
  });
  await new Promise<void>((resolve) => server.listen(0, resolve));
  const addr = server.address() as AddressInfo;
  try {
    return await fn(`http://127.0.0.1:${addr.port}`, () => oauthAuth);
  } finally {
    server.closeAllConnections?.();
    await new Promise<void>((resolve) => server.close(() => resolve()));
  }
}

// Web is primary: v2 succeeds, v3 rejects.
const webLoginHandler = (path: string) => {
  if (path.startsWith("/v3/users/login_with_email")) return reject401();
  if (path.startsWith("/v2/users/login_with_email"))
    return {
      status: 200,
      body: JSON.stringify({ access_token: "A1", refresh_token: "R1", user_id: 1 }),
    };
  return null;
};

async function scenarioRefreshWebSessionUsesWebKey(): Promise<void> {
  await withRefreshServer(webLoginHandler, async (baseURL, getOauthAuth) => {
    const client = new Client({ baseURL });
    await client.loginWithEmail({ email: "u@example.com", password: "pw" });
    // eslint-disable-next-line @typescript-eslint/no-explicit-any
    assertEq("refresh web: oauthKey is web key", (client as any)._oauthKeySnapshot(), DEFAULT_WEB_KEY);
    // eslint-disable-next-line @typescript-eslint/no-explicit-any
    const ok = await (client as any)._doRefresh(client.tokens().access);
    assertEq("refresh web: succeeds", ok, true);
    assertEq("refresh web: uses web-key Basic", getOauthAuth(), basic(DEFAULT_WEB_KEY));
    assertEq("refresh web: tokens rotated", client.tokens().access, "A2");
  });
}

// An app login (reached by disabling web-first with webApiKey="") refreshes
// with the app key and does not persist an oauth_key.
async function scenarioRefreshAppSessionUsesAppKey(): Promise<void> {
  const appLogin = (path: string) =>
    path.endsWith("/login_with_email")
      ? {
          status: 200,
          body: JSON.stringify({ access_token: "A1", refresh_token: "R1", user_id: 1 }),
        }
      : null;
  await withRefreshServer(appLogin, async (baseURL, getOauthAuth) => {
    const store = new MemorySessionStore();
    const client = new Client({ baseURL, sessionStore: store, webApiKey: "" });
    await client.loginWithEmail({ email: "u@example.com", password: "pw" });
    // eslint-disable-next-line @typescript-eslint/no-explicit-any
    assertEq("refresh app: oauthKey is app key", (client as any)._oauthKeySnapshot(), APP_KEY);
    const sess = await store.load("u@example.com");
    assertEq("refresh app: no oauthKey persisted", sess.oauthKey, undefined);
    // eslint-disable-next-line @typescript-eslint/no-explicit-any
    await (client as any)._doRefresh(client.tokens().access);
    assertEq("refresh app: uses app-key Basic", getOauthAuth(), basic(APP_KEY));
  });
}

// A cached web session, restored on a fresh client via the cached-login
// path, still refreshes with the web key.
async function scenarioRefreshRestoredWebSessionUsesWebKey(): Promise<void> {
  await withRefreshServer(webLoginHandler, async (baseURL, getOauthAuth) => {
    const store = new MemorySessionStore();
    const c1 = new Client({ baseURL, sessionStore: store });
    await c1.loginWithEmail({ email: "u@example.com", password: "pw" });
    assertEq(
      "restore: web oauthKey persisted",
      (await store.load("u@example.com")).oauthKey,
      DEFAULT_WEB_KEY,
    );
    const c2 = new Client({ baseURL, sessionStore: store });
    await c2.loginWithEmail({ email: "u@example.com", password: "pw" }); // cache hit
    // eslint-disable-next-line @typescript-eslint/no-explicit-any
    await (c2 as any)._doRefresh(c2.tokens().access);
    assertEq("restore: refresh uses web-key Basic", getOauthAuth(), basic(DEFAULT_WEB_KEY));
  });
}

// A web session restored directly with the public loadSession (not through
// the cached-login path) must also refresh with the web key. This is the
// regression the oauth_key restore fixes.
async function scenarioLoadSessionRestoresWebOAuthKey(): Promise<void> {
  await withRefreshServer(
    () => null,
    async (baseURL, getOauthAuth) => {
      const store = new MemorySessionStore();
      await store.save({
        email: "u@example.com",
        userID: 1,
        accessToken: "A1",
        refreshToken: "R1",
        deviceUUID: "",
        oauthKey: DEFAULT_WEB_KEY,
        updatedAt: new Date().toISOString(),
      });
      const client = new Client({ baseURL, sessionStore: store });
      await client.loadSession("u@example.com");
      // eslint-disable-next-line @typescript-eslint/no-explicit-any
      assertEq("loadSession web: oauthKey is web key", (client as any)._oauthKeySnapshot(), DEFAULT_WEB_KEY);
      // eslint-disable-next-line @typescript-eslint/no-explicit-any
      await (client as any)._doRefresh(client.tokens().access);
      assertEq("loadSession web: refresh uses web-key Basic", getOauthAuth(), basic(DEFAULT_WEB_KEY));
      assertEq("loadSession web: tokens rotated", client.tokens().access, "A2");
    },
  );
}

// App sessions carry an empty oauth_key; loadSession must leave refresh on
// the app key (empty → app-key default).
async function scenarioLoadSessionAppSessionKeepsAppKey(): Promise<void> {
  const store = new MemorySessionStore();
  await store.save({
    email: "u@example.com",
    userID: 1,
    accessToken: "A1",
    refreshToken: "R1",
    deviceUUID: "",
    updatedAt: new Date().toISOString(),
  });
  const client = new Client({ sessionStore: store });
  await client.loadSession("u@example.com");
  // eslint-disable-next-line @typescript-eslint/no-explicit-any
  assertEq("loadSession app: oauthKey is app key", (client as any)._oauthKeySnapshot(), APP_KEY);
}

// A web session injected directly via the public saveSession (not through
// login) also activates its establishing key, so its refresh uses the web
// key. Mirrors loadSession — saveSession is a full activation, not a write.
async function scenarioSaveSessionActivatesWebOAuthKey(): Promise<void> {
  await withRefreshServer(
    () => null,
    async (baseURL, getOauthAuth) => {
      const client = new Client({ baseURL, sessionStore: new MemorySessionStore() });
      await client.saveSession({
        email: "u@example.com",
        userID: 1,
        accessToken: "A1",
        refreshToken: "R1",
        deviceUUID: "",
        oauthKey: DEFAULT_WEB_KEY,
        updatedAt: new Date().toISOString(),
      });
      // eslint-disable-next-line @typescript-eslint/no-explicit-any
      assertEq("saveSession web: oauthKey is web key", (client as any)._oauthKeySnapshot(), DEFAULT_WEB_KEY);
      // eslint-disable-next-line @typescript-eslint/no-explicit-any
      await (client as any)._doRefresh(client.tokens().access);
      assertEq("saveSession web: refresh uses web-key Basic", getOauthAuth(), basic(DEFAULT_WEB_KEY));
    },
  );
}

(async () => {
  await scenarioFreshLogin();
  await scenarioCacheHit();
  await scenarioNoCacheBypass();
  await scenarioSavePersistFailure();
  await scenarioFileStoreRoundTrip();
  await scenarioWebPrimarySucceeds();
  await scenarioFallsBackToAppOn401();
  await scenarioWebDisabledUsesAppOnly();
  await scenarioBothFailSurfacesWeb401();
  await scenarioNoFallbackOn2FA();
  await scenarioNoFallbackWhenApiKeyPinned();
  await scenarioWebDecodeFailureSurfaced();
  await scenarioWebPrimaryTypedErrorSurfaced();
  await scenarioSignedInfoIgnoresOAuthKey();
  await scenarioRefreshWebSessionUsesWebKey();
  await scenarioRefreshAppSessionUsesAppKey();
  await scenarioRefreshRestoredWebSessionUsesWebKey();
  await scenarioLoadSessionRestoresWebOAuthKey();
  await scenarioLoadSessionAppSessionKeepsAppKey();
  await scenarioSaveSessionActivatesWebOAuthKey();
  process.exit(failed === 0 ? 0 : 1);
})().catch((e) => {
  // eslint-disable-next-line no-console
  console.error("test crashed:", e);
  process.exit(2);
});
