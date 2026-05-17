// LoginWithEmail error-path parity (PORTING.md §15 #3 / #5).
//
//   #3  a session-store load failure that is NOT a cache miss must
//       surface the error WITHOUT any HTTP being issued — this is the
//       rate-limit protection: a corrupt/locked store must not turn into
//       a flood of login attempts.
//   #5  a 2FA-required login returns the server's error envelope; the
//       caller can branch on it via codeOf(err) === ErrCodeRequired2FA.

import { createServer, type Server } from "node:http";
import type { AddressInfo } from "node:net";

import { Client } from "../src/client";
import { codeOf } from "../src/errors";
import { ErrCodeRequired2FA } from "../src/error_codes";
import type { Session, SessionStore } from "../src/session";

let failed = 0;
function assert(name: string, cond: boolean, detail = ""): void {
  if (cond) console.log(`PASS ${name}`);
  else {
    failed++;
    console.log(`FAIL ${name}${detail ? `\n  ${detail}` : ""}`);
  }
}

async function withServer(
  handler: (path: string, body: string) => { status: number; body: string },
  fn: (baseURL: string, hits: () => number) => Promise<void>,
): Promise<void> {
  let hitCount = 0;
  const server: Server = createServer((req, res) => {
    const chunks: Buffer[] = [];
    req.on("data", (c) => chunks.push(c));
    req.on("end", () => {
      hitCount++;
      const out = handler(req.url ?? "", Buffer.concat(chunks).toString("utf8"));
      res.writeHead(out.status, { "Content-Type": "application/json" });
      res.end(out.body);
    });
  });
  await new Promise<void>((r) => server.listen(0, r));
  const baseURL = `http://127.0.0.1:${(server.address() as AddressInfo).port}`;
  try {
    await fn(baseURL, () => hitCount);
  } finally {
    server.closeAllConnections?.();
    await new Promise<void>((r) => server.close(() => r()));
  }
}

// #3 — a store whose load() throws a non-"not found" error. The login
// must reject and never touch the network.
async function loadErrorReturnsErrorWithoutHTTP(): Promise<void> {
  await withServer(
    () => ({ status: 500, body: '{"error_code":-1}' }),
    async (baseURL, hits) => {
      // A SessionStore whose load fails with something other than a
      // NoSessionError (e.g. a corrupt/locked store).
      const brokenStore: SessionStore = {
        load: async (): Promise<Session> => {
          throw new Error("corrupt JSON");
        },
        save: async () => undefined,
        delete: async () => undefined,
      };
      const c = new Client({ baseURL, sessionStore: brokenStore });

      let caught: unknown = null;
      try {
        await c
          .loginWithEmail()
          .email("foo@example.com")
          .password("pw")
          .execute();
      } catch (err) {
        caught = err;
      }

      assert("load error surfaces an error", caught instanceof Error);
      assert(
        "load error message is not swallowed",
        caught instanceof Error && /corrupt JSON/.test(caught.message),
        caught instanceof Error ? caught.message : String(caught),
      );
      assert(
        "no HTTP issued on store error (rate-limit protection)",
        hits() === 0,
        `server hits=${hits()}`,
      );
      assert(
        "tokens not set after a broken-store login",
        c.tokens().access === "",
        c.tokens().access,
      );
    },
  );
}

// #5 — the login endpoint returns the server's error envelope with the
// 2FA code; codeOf(err) resolves it to the exported ErrCodeRequired2FA
// constant (value sourced from the SDK, not hardcoded here).
async function twoFARequiredSurfacesErrorCode(): Promise<void> {
  await withServer(
    (path) => {
      if (!path.endsWith("/login_with_email")) {
        return { status: 404, body: '{"status":404,"error":"Not Found"}' };
      }
      return {
        status: 401,
        body: JSON.stringify({
          error_code: ErrCodeRequired2FA,
          message: "2FA required",
        }),
      };
    },
    async (baseURL) => {
      const c = new Client({
        baseURL,
        retryPolicy: { maxAttempts: 0, baseDelay: 1, maxDelay: 1, retryOnPOST: false },
      });

      let caught: unknown = null;
      try {
        await c
          .loginWithEmail()
          .email("u@example.com")
          .password("pw")
          .execute();
      } catch (err) {
        caught = err;
      }

      assert("2FA-required login rejects", caught != null);
      assert(
        "codeOf(err) === ErrCodeRequired2FA",
        codeOf(caught) === ErrCodeRequired2FA,
        `codeOf=${codeOf(caught)} want=${ErrCodeRequired2FA}`,
      );
      assert(
        "tokens not set on 2FA-required",
        c.tokens().access === "",
        c.tokens().access,
      );
    },
  );
}

(async () => {
  await loadErrorReturnsErrorWithoutHTTP();
  await twoFARequiredSurfacesErrorCode();
  process.exit(failed === 0 ? 0 : 1);
})().catch((e) => {
  console.error("test crashed:", e);
  process.exit(2);
});
