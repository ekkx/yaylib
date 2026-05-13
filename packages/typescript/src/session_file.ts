// PORTING.md §5.1: file-backed SessionStore. Persists sessions as a JSON
// document keyed by email so one file can hold multiple identities.
// Writes are atomic within a single process (temp-file + rename); the
// file is created with mode 0o600 because it carries access tokens.
// Multi-process writers are out of scope (PORTING.md §5).

import { mkdir, readFile, rename, unlink, writeFile } from "node:fs/promises";
import { dirname } from "node:path";

import {
  type Session,
  type SessionStore,
  NoSessionError,
} from "./session";

interface FileStoreData {
  sessions: Record<string, Session>;
}

// AsyncMutex serializes the read-modify-write cycle inside one process
// so concurrent save / delete calls don't lose updates. Cross-process
// safety is not provided — use a separate session file per process.
class AsyncMutex {
  private current: Promise<void> = Promise.resolve();
  async run<T>(fn: () => Promise<T>): Promise<T> {
    const next = this.current.then(fn, fn);
    this.current = next.then(
      () => undefined,
      () => undefined,
    );
    return next;
  }
}

export class FileSessionStore implements SessionStore {
  private readonly mu = new AsyncMutex();

  constructor(private readonly path: string) {}

  async load(email: string): Promise<Session> {
    return this.mu.run(async () => {
      const data = await this.readAll();
      const s = data.sessions[email];
      if (!s) throw new NoSessionError(email);
      return { ...s, email: s.email || email };
    });
  }

  async save(session: Session): Promise<void> {
    if (!session.email) {
      throw new Error("yaylib: SessionStore.save: email is required");
    }
    await this.mu.run(async () => {
      const data = await this.readAll();
      data.sessions[session.email] = {
        ...session,
        updatedAt: session.updatedAt || new Date().toISOString(),
      };
      await this.writeAll(data);
    });
  }

  async delete(email: string): Promise<void> {
    await this.mu.run(async () => {
      const data = await this.readAll();
      if (!(email in data.sessions)) return;
      delete data.sessions[email];
      await this.writeAll(data);
    });
  }

  private async readAll(): Promise<FileStoreData> {
    let text: string;
    try {
      text = await readFile(this.path, "utf8");
    } catch (err) {
      if ((err as NodeJS.ErrnoException).code === "ENOENT") {
        return { sessions: {} };
      }
      throw err;
    }
    let parsed: Partial<FileStoreData>;
    try {
      parsed = JSON.parse(text) as Partial<FileStoreData>;
    } catch (err) {
      throw new Error(
        `yaylib: failed to parse session file ${this.path}: ${
          err instanceof Error ? err.message : String(err)
        }`,
      );
    }
    return { sessions: parsed.sessions ?? {} };
  }

  private async writeAll(data: FileStoreData): Promise<void> {
    const dir = dirname(this.path);
    if (dir && dir !== ".") {
      await mkdir(dir, { recursive: true, mode: 0o700 });
    }
    const tmp = `${this.path}.tmp`;
    await writeFile(tmp, JSON.stringify(data, null, 2), { mode: 0o600 });
    try {
      await rename(tmp, this.path);
    } catch (err) {
      try {
        await unlink(tmp);
      } catch {
        /* ignore */
      }
      throw err;
    }
  }
}

// newSessionStore is the Node-side default factory. Creates the parent
// directory eagerly so subsequent saves never fail on a missing path.
export async function newSessionStore(path: string): Promise<SessionStore> {
  const dir = dirname(path);
  if (dir && dir !== ".") {
    await mkdir(dir, { recursive: true, mode: 0o700 });
  }
  return new FileSessionStore(path);
}
