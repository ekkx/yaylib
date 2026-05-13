// PORTING.md §12: required headers injected on every API call. This
// middleware sets the headers the Yay! server validates and picks the
// right Authorization scheme. The 401 auto-refresh chain (§6.1) lives in
// auth.ts; this module covers only headers.

import type { Middleware, RequestContext, FetchParams } from "./gen/runtime";

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

const OAUTH_TOKEN_PATH = "/api/v1/oauth/token";

function isOAuthTokenPath(url: string): boolean {
  return url.includes(OAUTH_TOKEN_PATH);
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
