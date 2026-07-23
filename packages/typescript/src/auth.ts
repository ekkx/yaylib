// PORTING.md §6: LoginWithEmail is the only wrapped auth endpoint. The
// wrapper is the cache boundary — "log in or use the cached session" is
// one call. Other auth-flow body endpoints (CreateUser, EditUser,
// LoginWithSns, OauthToken / refresh) stay as plain calls on the
// generated client.

import type { LoginUserResponse } from "./gen/models/LoginUserResponse.js";

import type { Client } from "./client.js";
import { APIError, asAPIError, codeOf } from "./errors.js";
import { ErrCodeUnauthorized } from "./error_codes.js";
import { type Session, SessionSaveFailedError } from "./session.js";

// LoginWithEmailParams is the options object — the same call shape as
// every other generated TS operation (a single params object), not a
// builder. The generated layer accepts only one body parameter
// (`loginEmailUserRequest`); the wrapper lets callers pass each field
// directly and auto-fills apiKey / uuid from the client.
export interface LoginWithEmailParams {
  email: string;
  password: string;
  /** Override the API key for this single login call. Defaults to client.apiKey. */
  apiKey?: string;
  /** Override the device UUID for this single login call. Defaults to client.deviceUUID. */
  uuid?: string;
  twoFACode?: string;
  /** Force a fresh HTTP login, bypassing any cached session. */
  noCache?: boolean;
}

interface LoginParams {
  email?: string;
  password?: string;
  apiKey?: string;
  uuid?: string;
  twoFACode?: string;
  bypassCache?: boolean;
}

// loginWithEmail performs the wrapped login. Exposed on Client as
// `await client.loginWithEmail({ email, password })`; see client.ts.
export async function loginWithEmail(
  client: Client,
  params: LoginWithEmailParams,
): Promise<LoginUserResponse> {
  return executeLogin(client, {
    email: params.email,
    password: params.password,
    apiKey: params.apiKey,
    uuid: params.uuid,
    twoFACode: params.twoFACode,
    bypassCache: params.noCache,
  });
}

// LoginBody is the resolved login request (apiKey / uuid filled in),
// shared by the web and app login helpers.
interface LoginBody {
  email: string;
  password: string;
  apiKey: string;
  uuid: string;
  twoFACode?: string;
}

async function executeLogin(
  client: Client,
  params: LoginParams,
): Promise<LoginUserResponse> {
  if (!params.email || !params.password) {
    throw new Error("yaylib: loginWithEmail requires email and password");
  }

  // Cache lookup — only when a session store is configured and the caller
  // has not opted out via .noCache(). loadSession fully activates the
  // session (tokens, device UUID, OAuth key, identity), matching the Go
  // maybeRestore path.
  if (client.sessionStore && !params.bypassCache && params.email) {
    try {
      const cached = await client.loadSession(params.email);
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
  const body: LoginBody = {
    email: params.email,
    password: params.password,
    apiKey: params.apiKey ?? client.apiKey,
    uuid: params.uuid ?? client.deviceUUID,
    twoFACode: params.twoFACode,
  };

  const { resp, oauthKey } = await doFreshLogin(client, body);

  // Activate tokens / identity on the client.
  const access = resp.accessToken ?? "";
  const refresh = resp.refreshToken ?? "";
  const userID = resp.userId ?? 0;
  client.setTokens(access, refresh);
  client.setLoginIdentity(params.email, userID);
  client.setOAuthKey(oauthKey);

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
      // Omit for a default (app) session so its file stays unchanged.
      oauthKey: oauthKey === client.apiKey ? undefined : oauthKey,
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

// doFreshLogin performs a fresh (non-cached) email login and reports which
// key established the session — the Basic credential later token refreshes
// must use (PORTING.md §6.2). Web-first: the browser login is primary; the
// app login is the fallback, the sole path when the browser endpoint is
// switched off (webApiKey === "") or the caller pinned a specific apiKey,
// and never silently replaced by webApiKey. Mirrors Go doFreshLogin.
async function doFreshLogin(
  client: Client,
  body: LoginBody,
): Promise<{ resp: LoginUserResponse; oauthKey: string }> {
  // A caller-pinned api_key (resolved key differs from the client default)
  // means "use the app login with exactly this key"; skip the browser path
  // so we never authenticate under webApiKey instead.
  const callerPinnedKey = body.apiKey !== client.apiKey;

  if (!client.webApiKey || callerPinnedKey) {
    const resp = await loginWithEmailApp(client, body);
    return { resp, oauthKey: body.apiKey };
  }

  // Primary: browser login with the web key.
  let webErr: APIError;
  try {
    const resp = await client.loginWithEmailWeb({
      email: body.email,
      password: body.password,
      uuid: body.uuid,
      twoFACode: body.twoFACode,
    });
    return { resp, oauthKey: client.webApiKey };
  } catch (err) {
    webErr = await asAPIError(err);
  }

  // Fall back to the app login ONLY on a plain unauthorized. Any other
  // typed rejection surfaces unchanged — notably Required2FA (the caller
  // must supply a code, not switch endpoints), and a 2xx that failed to
  // decode (a genuine login result whose raw body rides on the error, §15),
  // which is not Unauthorized.
  if (codeOf(webErr) !== ErrCodeUnauthorized) {
    throw webErr;
  }
  client.logger.debug("app login fallback", {
    event: "login_fallback",
    outcome: "attempt",
  });
  try {
    const resp = await loginWithEmailApp(client, body);
    client.logger.debug("app login fallback ok", {
      event: "login_fallback",
      outcome: "ok",
    });
    return { resp, oauthKey: body.apiKey };
  } catch {
    client.logger.debug("app login fallback failed", {
      event: "login_fallback",
      outcome: "failed",
    });
    // Both paths failed: surface the primary (web) error, which carries the
    // parsed error code and raw body.
    throw webErr;
  }
}

// loginWithEmailApp issues the app login through the wrapped generated
// client (header injection, retry, host routing), converting the thrown
// generated error into our stable APIError so codeOf / errorResponseOf
// work on it uniformly.
async function loginWithEmailApp(
  client: Client,
  body: LoginBody,
): Promise<LoginUserResponse> {
  try {
    return await client.usersAPI.loginWithEmail({ loginEmailUserRequest: body });
  } catch (err) {
    throw await asAPIError(err);
  }
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
