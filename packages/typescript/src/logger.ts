// PORTING.md §12.2: logging is optional, injected, and silent by
// default. The SDK is a library — with no logger configured it MUST
// produce zero output. A logger is supplied once via ClientOptions and
// the SDK emits structured records to it from then on.
//
// Every record carries a stable `event` key whose value is the
// cross-language contract (the message text MAY differ per language;
// the `event` value and field keys MUST NOT). Banned values — access /
// refresh tokens, password, API key, signed_info/version, X-Jwt, the
// Authorization header, request/response bodies — MUST NEVER be passed
// to a logger, at any level.

export type LogFields = Record<string, unknown>;

export interface Logger {
  debug(message: string, fields?: LogFields): void;
  info(message: string, fields?: LogFields): void;
  warn(message: string, fields?: LogFields): void;
  error(message: string, fields?: LogFields): void;
}

// The silent default. Every method is a no-op so a library consumer
// that never opts in sees nothing on stdout/stderr.
export const noopLogger: Logger = {
  debug() {},
  info() {},
  warn() {},
  error() {},
};
