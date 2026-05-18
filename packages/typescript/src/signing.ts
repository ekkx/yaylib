// PORTING.md §7: signed_info / signed_version / X-Jwt.
//
// All three helpers produce byte-identical output to the Go reference
// against the test vectors pinned in PORTING.md §7.1.
//
// Runtime requirement: Node 18+. The current implementation reaches into
// `node:crypto` because Web Crypto has no MD5 primitive. Browsers /
// Workers / Deno can be supported by swapping in a pure-JS MD5 (e.g. a
// small WASM build or a hand-rolled implementation); the wire formulas
// stay the same.

import { createHash, createHmac } from "node:crypto";

import { SIGNED_INFO_SHARED_KEY, SIGNED_VERSION_PLATFORM } from "./config.js";

// SignedInfo bundles the timestamp and MD5 hash a Yay! request needs to
// pass server-side validation. The two values are bound: changing one
// without the other invalidates the signature.
export interface SignedInfo {
  // Unix-second timestamp the hash is computed against. Send this as the
  // request's `timestamp` field.
  readonly timestamp: number;
  // Lowercase 32-char MD5 hex digest. Send this as `signed_info`.
  readonly value: string;
}

// generateSignedInfoAt computes the signed_info MD5 for the given
// Unix-second timestamp without any network I/O. Use this when you
// already have a synchronized timestamp; otherwise prefer
// generateSignedInfo on the Client, which fetches the server's view of
// the clock to avoid device-clock drift.
export function generateSignedInfoAt(
  opts: { apiKey: string; deviceUUID: string },
  timestamp: number,
): SignedInfo {
  const payload =
    opts.apiKey + opts.deviceUUID + String(timestamp) + SIGNED_INFO_SHARED_KEY;
  const value = createHash("md5").update(payload, "utf8").digest("hex");
  return { timestamp, value };
}

// generateSignedVersion returns the signed_version HMAC token some
// endpoints require alongside signed_info. The platform identifier is
// hard-coded to `yay_android` to stay inside the server's platform
// allowlist; other values risk rejection.
export function generateSignedVersion(opts: {
  apiVersionKey: string;
  apiVersionName: string;
}): string {
  const payload = `${SIGNED_VERSION_PLATFORM}/${opts.apiVersionName}`;
  return createHmac("sha256", opts.apiVersionKey)
    .update(payload, "utf8")
    .digest("base64");
}

// generateXJwt returns the short-lived HS256 JWT required by some write
// endpoints (most visibly CreatePost). The token's iat is the current
// Unix-second time; exp = iat + 5; the HS256 key is APIVersionKey. TTL
// is 5 seconds, so callers MUST regenerate per-request — a token
// prepared offline and sent later will be rejected.
export function generateXJwt(opts: {
  apiVersionKey: string;
  now?: number;
  ttlSeconds?: number;
}): string {
  const iat = opts.now ?? Math.floor(Date.now() / 1000);
  const exp = iat + (opts.ttlSeconds ?? 5);
  return signJWT(opts.apiVersionKey, iat, exp);
}

function signJWT(key: string, iat: number, exp: number): string {
  // Header JSON has no `typ` and no spaces; payload JSON key order is
  // `iat` then `exp`, also no spaces. The base64 form is URL-safe with
  // `=` padding stripped (RFC 4648 §5, no_pad).
  const header = base64UrlNoPad(Buffer.from('{"alg":"HS256"}', "utf8"));
  const payload = base64UrlNoPad(
    Buffer.from(`{"iat":${iat},"exp":${exp}}`, "utf8"),
  );
  const signingInput = `${header}.${payload}`;
  const sig = createHmac("sha256", key).update(signingInput, "utf8").digest();
  return `${signingInput}.${base64UrlNoPad(sig)}`;
}

function base64UrlNoPad(buf: Buffer): string {
  return buf
    .toString("base64")
    .replace(/=+$/g, "")
    .replace(/\+/g, "-")
    .replace(/\//g, "_");
}
