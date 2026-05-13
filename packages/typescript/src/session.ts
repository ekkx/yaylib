// PORTING.md §5: SessionStore persists (email -> Session) pairs so tokens
// survive across process restarts. This module ships the interface and
// an in-memory implementation; the file-backed default lives in
// session_file.ts (Node-only) and is loaded lazily by NewSessionStore.

export interface Session {
  email: string;
  userID: number;
  accessToken: string;
  refreshToken: string;
  deviceUUID: string;
  updatedAt: string; // RFC3339
}

export class NoSessionError extends Error {
  constructor(email: string) {
    super(`yaylib: no session for ${email}`);
    this.name = "NoSessionError";
  }
}

export class SessionSaveFailedError extends Error {
  readonly cause: unknown;
  constructor(cause: unknown) {
    super(
      `yaylib: failed to persist session: ${
        cause instanceof Error ? cause.message : String(cause)
      }`,
    );
    this.name = "SessionSaveFailedError";
    this.cause = cause;
  }
}

export interface SessionStore {
  load(email: string): Promise<Session>;
  save(session: Session): Promise<void>;
  delete(email: string): Promise<void>;
}

// ---- in-memory implementation ----

export class MemorySessionStore implements SessionStore {
  private readonly sessions = new Map<string, Session>();

  async load(email: string): Promise<Session> {
    const s = this.sessions.get(email);
    if (!s) throw new NoSessionError(email);
    return { ...s };
  }

  async save(session: Session): Promise<void> {
    if (!session.email) {
      throw new Error("yaylib: SessionStore.save: email is required");
    }
    const stored: Session = {
      ...session,
      updatedAt: session.updatedAt || new Date().toISOString(),
    };
    this.sessions.set(stored.email, stored);
  }

  async delete(email: string): Promise<void> {
    this.sessions.delete(email);
  }
}

export function newMemoryStore(): SessionStore {
  return new MemorySessionStore();
}
