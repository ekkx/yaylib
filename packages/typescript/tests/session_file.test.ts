// FileSessionStore round-trip / atomicity / concurrency parity tests.
// Mirrors the contract in PORTING.md §5 + §5.1.

import { mkdtempSync, statSync, rmSync, readFileSync } from "node:fs";
import { tmpdir } from "node:os";
import { join } from "node:path";

import { newSessionStore, FileSessionStore } from "../src/session_file";
import { NoSessionError, type Session } from "../src/session";

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

function makeSession(email: string, userID: number): Session {
  return {
    email,
    userID,
    accessToken: `access-${userID}`,
    refreshToken: `refresh-${userID}`,
    deviceUUID: `00000000-0000-0000-0000-${String(userID).padStart(12, "0")}`,
    updatedAt: "",
  };
}

async function withTempDir<T>(fn: (dir: string) => Promise<T>): Promise<T> {
  const dir = mkdtempSync(join(tmpdir(), "yaylib-sess-"));
  try {
    return await fn(dir);
  } finally {
    rmSync(dir, { recursive: true, force: true });
  }
}

async function roundTrip(): Promise<void> {
  await withTempDir(async (dir) => {
    const path = join(dir, "sessions.json");
    const store = await newSessionStore(path);
    await store.save(makeSession("a@example.com", 1));
    const loaded = await store.load("a@example.com");
    assertEq("roundtrip: userID preserved", loaded.userID, 1);
    assertEq(
      "roundtrip: access token preserved",
      loaded.accessToken,
      "access-1",
    );
    assertEq(
      "roundtrip: deviceUUID preserved",
      loaded.deviceUUID,
      "00000000-0000-0000-0000-000000000001",
    );
    assertEq("roundtrip: updatedAt was auto-populated", loaded.updatedAt !== "", true);
  });
}

async function loadMissingThrows(): Promise<void> {
  await withTempDir(async (dir) => {
    const store = await newSessionStore(join(dir, "sessions.json"));
    let caught: unknown = null;
    try {
      await store.load("ghost@example.com");
    } catch (err) {
      caught = err;
    }
    assertEq(
      "load missing: throws NoSessionError",
      caught instanceof NoSessionError,
      true,
    );
  });
}

async function deleteIdempotent(): Promise<void> {
  await withTempDir(async (dir) => {
    const store = await newSessionStore(join(dir, "sessions.json"));
    // Should be a no-op even when the file doesn't exist yet.
    await store.delete("ghost@example.com");
    await store.save(makeSession("real@example.com", 7));
    await store.delete("real@example.com");
    let caught: unknown = null;
    try {
      await store.load("real@example.com");
    } catch (err) {
      caught = err;
    }
    assertEq(
      "delete: load after delete throws NoSessionError",
      caught instanceof NoSessionError,
      true,
    );
    // Second delete on the now-empty record is still a no-op.
    await store.delete("real@example.com");
    assertEq("delete: second delete didn't throw", true, true);
  });
}

async function multipleSessions(): Promise<void> {
  await withTempDir(async (dir) => {
    const path = join(dir, "sessions.json");
    const store = await newSessionStore(path);
    await store.save(makeSession("a@example.com", 1));
    await store.save(makeSession("b@example.com", 2));
    await store.save(makeSession("c@example.com", 3));
    const a = await store.load("a@example.com");
    const b = await store.load("b@example.com");
    const c = await store.load("c@example.com");
    assertEq("multi: userID a", a.userID, 1);
    assertEq("multi: userID b", b.userID, 2);
    assertEq("multi: userID c", c.userID, 3);
    const raw = readFileSync(path, "utf8");
    const data = JSON.parse(raw) as { sessions: Record<string, unknown> };
    assertEq(
      "multi: file holds 3 distinct keys",
      Object.keys(data.sessions).length,
      3,
    );
  });
}

async function filePermissions(): Promise<void> {
  if (process.platform === "win32") {
    // eslint-disable-next-line no-console
    console.log("SKIP file permissions check on win32");
    return;
  }
  await withTempDir(async (dir) => {
    const path = join(dir, "sessions.json");
    const store = await newSessionStore(path);
    await store.save(makeSession("perm@example.com", 9));
    const st = statSync(path);
    const mode = st.mode & 0o777;
    assertEq("permissions: file mode is 0o600", mode, 0o600);
  });
}

async function concurrentWrites(): Promise<void> {
  await withTempDir(async (dir) => {
    const path = join(dir, "sessions.json");
    const store = new FileSessionStore(path);
    // Fire 10 concurrent saves under distinct emails. The mutex inside
    // FileSessionStore serializes the read-modify-write cycles, so all 10
    // records MUST be present in the final file.
    await Promise.all(
      Array.from({ length: 10 }, (_, i) =>
        store.save(makeSession(`u${i}@example.com`, i + 1)),
      ),
    );
    const raw = readFileSync(path, "utf8");
    const data = JSON.parse(raw) as { sessions: Record<string, unknown> };
    assertEq(
      "concurrent: all 10 sessions persisted",
      Object.keys(data.sessions).length,
      10,
    );
  });
}

(async () => {
  await roundTrip();
  await loadMissingThrows();
  await deleteIdempotent();
  await multipleSessions();
  await filePermissions();
  await concurrentWrites();
  process.exit(failed === 0 ? 0 : 1);
})().catch((e) => {
  // eslint-disable-next-line no-console
  console.error("test crashed:", e);
  process.exit(2);
});
