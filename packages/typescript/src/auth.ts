// PORTING.md §6: LoginWithEmail is the only wrapped auth endpoint. The
// wrapper is the cache boundary — "log in or use the cached session" is
// one call. Other auth-flow body endpoints (CreateUser, EditUser,
// LoginWithSns, OauthToken / refresh) stay as plain calls on the
// generated client.

import type { LoginUserResponse } from "./gen/models/LoginUserResponse";

import type { Client } from "./client";
import { APIError, asAPIError } from "./errors";
import { type Session, SessionSaveFailedError } from "./session";

// LoginWithEmailBuilder is the chain-setter form that mirrors the Go
// reference. The generated layer accepts only a single body parameter
// (`loginEmailUserRequest`); the wrapper lets callers populate each
// field individually and auto-fills APIKey / Uuid from the client
// before issuing the request.
export interface LoginWithEmailBuilder {
  email(v: string): LoginWithEmailBuilder;
  password(v: string): LoginWithEmailBuilder;
  /** Override the API key on this single login call. Defaults to client.apiKey. */
  apiKey(v: string): LoginWithEmailBuilder;
  /** Override the device UUID on this single login call. Defaults to client.deviceUUID. */
  uuid(v: string): LoginWithEmailBuilder;
  twoFACode(v: string): LoginWithEmailBuilder;
  /** Force a fresh HTTP login, bypassing any cached session. */
  noCache(): LoginWithEmailBuilder;
  execute(): Promise<LoginUserResponse>;
}

interface LoginParams {
  email?: string;
  password?: string;
  apiKey?: string;
  uuid?: string;
  twoFACode?: string;
  bypassCache?: boolean;
}

// loginWithEmail returns the builder. The wrapper is exposed on Client
// as `client.loginWithEmail()`; see client.ts.
export function loginWithEmail(client: Client): LoginWithEmailBuilder {
  const params: LoginParams = {};
  const builder: LoginWithEmailBuilder = {
    email: (v) => {
      params.email = v;
      return builder;
    },
    password: (v) => {
      params.password = v;
      return builder;
    },
    apiKey: (v) => {
      params.apiKey = v;
      return builder;
    },
    uuid: (v) => {
      params.uuid = v;
      return builder;
    },
    twoFACode: (v) => {
      params.twoFACode = v;
      return builder;
    },
    noCache: () => {
      params.bypassCache = true;
      return builder;
    },
    execute: async () => {
      return await executeLogin(client, params);
    },
  };
  return builder;
}

async function executeLogin(
  client: Client,
  params: LoginParams,
): Promise<LoginUserResponse> {
  if (!params.email || !params.password) {
    throw new Error("yaylib: LoginWithEmail requires email and password");
  }

  // Cache lookup — only when a session store is configured and the
  // caller has not opted out via .noCache().
  if (client.sessionStore && !params.bypassCache) {
    try {
      const cached = await client.sessionStore.load(params.email);
      activateCachedSession(client, cached);
      return synthesizeLoginResponse(cached);
    } catch (err) {
      // NoSessionError → cache miss, fall through to HTTP login. Other
      // errors propagate so the caller doesn't silently drop a rate-limit
      // or disk failure on the wire-bound load path.
      if (!isNoSessionError(err)) {
        throw err;
      }
    }
  }

  // Fresh HTTP login. APIKey / Uuid auto-fill matches the Go reference.
  const body = {
    email: params.email,
    password: params.password,
    apiKey: params.apiKey ?? client.apiKey,
    uuid: params.uuid ?? client.deviceUUID,
    twoFACode: params.twoFACode,
  };

  let resp: LoginUserResponse;
  try {
    resp = await client.usersAPI.loginWithEmail({ loginEmailUserRequest: body });
  } catch (err) {
    throw await asAPIError(err);
  }

  // Activate tokens / identity on the client.
  const access = resp.accessToken ?? "";
  const refresh = resp.refreshToken ?? "";
  const userID = resp.userId ?? 0;
  client.setTokens(access, refresh);
  client.setLoginIdentity(params.email, userID);

  // Persist the session. PORTING.md §5: a persist failure here returns
  // SessionSaveFailedError, but the tokens are still active in memory so
  // the caller can choose to proceed.
  if (client.sessionStore) {
    const session: Session = {
      email: params.email,
      userID,
      accessToken: access,
      refreshToken: refresh,
      deviceUUID: client.deviceUUID,
      updatedAt: new Date().toISOString(),
    };
    try {
      await client.sessionStore.save(session);
    } catch (err) {
      throw new SessionSaveFailedError(err);
    }
  }

  return resp;
}

function activateCachedSession(client: Client, session: Session): void {
  client.setTokens(session.accessToken, session.refreshToken);
  if (session.deviceUUID) {
    client.setDeviceUUID(session.deviceUUID);
  }
  client.setLoginIdentity(session.email, session.userID);
}

// Synthesize a LoginUserResponse from a cached Session so callers can
// observe the same shape regardless of cache hit vs. miss. Fields the
// server returns but the cache doesn't carry (e.g. login_radar metadata)
// default to undefined.
function synthesizeLoginResponse(session: Session): LoginUserResponse {
  return {
    accessToken: session.accessToken,
    refreshToken: session.refreshToken,
    userId: session.userID,
  } as LoginUserResponse;
}

function isNoSessionError(err: unknown): boolean {
  return (
    !!err &&
    typeof err === "object" &&
    "name" in err &&
    (err as { name: string }).name === "NoSessionError"
  );
}

// Re-export from index — keeps the wrapper accessible without callers
// having to know about the auth module directly.
export { APIError };
