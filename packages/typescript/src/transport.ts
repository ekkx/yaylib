// PORTING.md §12: required headers injected on every API call, plus the
// 401 auto-refresh chain (§6.1). The headers middleware sets the headers
// the Yay! server validates and picks the right Authorization scheme;
// the refresh middleware detects a 401 response, asks the client to
// refresh the access token, and retries the original request once.

import type {
  Middleware,
  RequestContext,
  ResponseContext,
  FetchParams,
} from "./gen/runtime";
import { hostRoutes } from "./gen/hostRoutes";

export interface TransportContext {
  userAgent: string;
  appVersion: string; // used as X-App-Version
  deviceInfo: string;
  deviceUUID: string;
  connectionType: string;
  connectionSpeed: string;
  acceptLanguage: string;
  apiKey: string;
  clientIP(): string;
  accessToken(): string;
}

// RefreshFn attempts a one-shot refresh of the access token using the
// stale value carried by the original 401 request. Implementations MUST
// collapse concurrent calls — when multiple in-flight requests fail with
// 401 at the same time, only one OAuth /token round-trip should happen
// and the others MUST await the same outcome (PORTING.md §6.1).
//
// Return values:
//   - true  : the access token is now fresh; the caller should retry the
//             original request with the new Bearer.
//   - false : refresh failed (no refresh token, refresh endpoint
//             returned non-2xx, or the credentials store is empty). The
//             caller MUST surface the original 401 to the user.
export type RefreshFn = (staleAccess: string) => Promise<boolean>;

const OAUTH_TOKEN_PATH = "/api/v1/oauth/token";

function isOAuthTokenPath(url: string): boolean {
  return url.includes(OAUTH_TOKEN_PATH);
}

// HostRoutingContext supplies the configurable auxiliary base URLs the
// host-routing middleware resolves symbolic hosts to.
export interface HostRoutingContext {
  cassandraBaseURL: string;
}

function resolveAuxHost(alias: string, ctx: HostRoutingContext): string | undefined {
  switch (alias) {
    case "cassandra":
      return ctx.cassandraBaseURL;
    default:
      // Unrecognized alias from the generated table: keep the request
      // on the primary host rather than stranding it.
      return undefined;
  }
}

// buildHostRoutingMiddleware rewrites the request's origin when its
// operation is served from an auxiliary host (see gen/hostRoutes).
// Match is by exact "METHOD path" — the host-routed operations are all
// parameter-free, so the request pathname equals the OpenAPI path
// template. It must run before every other middleware so headers /
// auth / retry all observe the final URL (mirrors the Go transport's
// routeHost running first in RoundTrip).
export function buildHostRoutingMiddleware(ctx: HostRoutingContext): Middleware {
  return {
    pre: async (req: RequestContext): Promise<FetchParams | void> => {
      let url: URL;
      try {
        url = new URL(req.url);
      } catch {
        return; // non-absolute URL — nothing to rewrite
      }
      const method = (req.init.method ?? "GET").toUpperCase();
      const alias = hostRoutes[`${method} ${url.pathname}`];
      if (!alias) return;
      const base = resolveAuxHost(alias, ctx);
      if (!base) return;
      let target: URL;
      try {
        target = new URL(base);
      } catch {
        return;
      }
      url.protocol = target.protocol;
      url.host = target.host;
      return { url: url.toString(), init: req.init };
    },
  };
}

function setIfAbsent(headers: Record<string, string>, key: string, value: string) {
  // Header keys in fetch's Record<string,string> form are case-sensitive in
  // practice; the runtime canonicalizes on send. We mirror the Go transport
  // which canonicalizes via http.Header.
  for (const k of Object.keys(headers)) {
    if (k.toLowerCase() === key.toLowerCase()) return;
  }
  if (value !== "") headers[key] = value;
}

// Base64-encode the API key for `Authorization: Basic`. Node 18+ exposes
// btoa as a global; the SDK's engine requirement (^18) makes the fallback
// unnecessary. For the Yay! API key the input is ASCII so binary-safe
// btoa is correct without TextEncoder gymnastics.
function base64Encode(input: string): string {
  return btoa(input);
}

export function buildHeadersMiddleware(ctx: TransportContext): Middleware {
  return {
    pre: async (req: RequestContext): Promise<FetchParams | void> => {
      const init = req.init;
      const headers: Record<string, string> =
        (init.headers as Record<string, string>) ?? {};

      setIfAbsent(headers, "User-Agent", ctx.userAgent);
      setIfAbsent(headers, "X-App-Version", ctx.appVersion);
      setIfAbsent(headers, "X-Device-Info", ctx.deviceInfo);
      setIfAbsent(headers, "X-Device-UUID", ctx.deviceUUID);
      setIfAbsent(headers, "X-Client-IP", ctx.clientIP());
      setIfAbsent(headers, "X-Connection-Type", ctx.connectionType);
      setIfAbsent(headers, "X-Connection-Speed", ctx.connectionSpeed);
      setIfAbsent(headers, "Accept-Language", ctx.acceptLanguage);

      // X-Timestamp must be fresh per request.
      headers["X-Timestamp"] = String(Math.floor(Date.now() / 1000));

      // Normalize JSON Content-Type to match the server's exact casing /
      // spacing expectations.
      for (const k of Object.keys(headers)) {
        if (k.toLowerCase() === "content-type") {
          const v = headers[k];
          if (typeof v === "string" && v.toLowerCase().startsWith("application/json")) {
            headers[k] = "application/json;charset=UTF-8";
          }
        }
      }

      // Authorization: Basic for oauth/token*, Bearer for everything else.
      const hasAuth = Object.keys(headers).some(
        (k) => k.toLowerCase() === "authorization",
      );
      if (!hasAuth) {
        if (isOAuthTokenPath(req.url)) {
          headers["Authorization"] = `Basic ${base64Encode(ctx.apiKey)}`;
        } else {
          const access = ctx.accessToken();
          if (access) headers["Authorization"] = `Bearer ${access}`;
        }
      }

      return {
        url: req.url,
        init: { ...init, headers },
      };
    },
  };
}

// buildAuthRefreshMiddleware watches every response, kicks off a single
// OAuth /token refresh on 401 (excluding the OAuth endpoint itself),
// and retries the original request once with the new Bearer. The
// post-middleware semantics mirror the Go transport's handle401:
//
//   - Refresh success: re-issue the request, return the retry response
//     (the original 401 body is dropped).
//   - Refresh failure: return undefined so the original 401 propagates
//     to the caller with its body intact.
//   - Refresh success + retry transport error: the error bubbles out of
//     the post-middleware as if the original fetch had failed at the
//     transport layer — callers see the retry's network error rather
//     than a misleading 401.
//
// The middleware reads the Authorization header set by the headers
// middleware to learn which access token was actually sent, so the
// RefreshFn can no-op when a concurrent caller has already rotated the
// token by the time control returns here.
export function buildAuthRefreshMiddleware(ctx: {
  refresh: RefreshFn;
  accessToken(): string;
  fetchApi?: typeof fetch;
}): Middleware {
  return {
    post: async (rc: ResponseContext): Promise<Response | void> => {
      const { response, init, url, fetch: contextFetch } = rc;
      if (response.status !== 401) return undefined;
      if (isOAuthTokenPath(url)) return undefined;

      const headers = (init.headers ?? {}) as Record<string, string>;
      const authHeader =
        Object.entries(headers).find(([k]) => k.toLowerCase() === "authorization")?.[1] ?? "";
      const staleAccess = authHeader.startsWith("Bearer ") ? authHeader.slice(7) : "";

      const refreshOK = await ctx.refresh(staleAccess);
      if (!refreshOK) return undefined;

      // Drain the original 401 body so the underlying connection can be
      // released; we're about to discard the response in favour of the
      // retry.
      try {
        await response.arrayBuffer();
      } catch {
        /* ignore */
      }

      const retryHeaders: Record<string, string> = { ...headers };
      const newAccess = ctx.accessToken();
      if (newAccess) {
        retryHeaders["Authorization"] = `Bearer ${newAccess}`;
      }
      // X-Timestamp MUST stay fresh — the original value is now stale.
      retryHeaders["X-Timestamp"] = String(Math.floor(Date.now() / 1000));

      const fetchImpl = ctx.fetchApi ?? contextFetch ?? globalThis.fetch;
      return await fetchImpl(url, { ...init, headers: retryHeaders });
    },
  };
}
