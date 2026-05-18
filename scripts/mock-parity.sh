#!/usr/bin/env bash
#
# Cross-language behavior-parity runner.
#
# Builds and boots the bundled behavior-contract mock server on an
# ephemeral port, exports its address, then runs one or more of the
# language test suites against that single instance so Go / TypeScript /
# Python assert identical retry / refresh / rate-limit / upload / event
# -stream behavior instead of three divergent in-process fakes.
#
# The suites discover the server via two env vars and skip their parity
# tests when the vars are unset (so a standalone `go test ./.` /
# `npm test` / `pytest` still passes with only the pure-function and
# typed-decode fixtures):
#
#   YAYLIB_MOCK_BASE_URL   http://host:port   (REST)
#   YAYLIB_MOCK_WS_URL     ws://host:port     (event stream)
#
# Each language environment (Go toolchain, npm deps, Python deps) is
# expected to be prepared by the caller.
#
# Usage:
#   scripts/mock-parity.sh             # all three suites
#   scripts/mock-parity.sh go          # one suite (go | ts | py)
#   scripts/mock-parity.sh go ts       # a subset
#
set -euo pipefail

SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
ROOT_DIR="$(cd "$SCRIPT_DIR/.." && pwd)"

SUITES=("$@")
if [ ${#SUITES[@]} -eq 0 ]; then
  SUITES=(go ts py)
fi

TMPDIR_RUN="$(mktemp -d)"
BIN="$TMPDIR_RUN/mockd"
READY="$TMPDIR_RUN/addr"
SRV_PID=""

cleanup() {
  [ -n "$SRV_PID" ] && kill "$SRV_PID" 2>/dev/null || true
  rm -rf "$TMPDIR_RUN"
}
trap cleanup EXIT

echo "==> building mock server"
( cd "$ROOT_DIR/mockd" && go build -o "$BIN" . )

echo "==> starting mock server on an ephemeral port"
"$BIN" -addr 127.0.0.1:0 -ready-file "$READY" &
SRV_PID=$!

# Wait for the resolved address to appear (the server writes it once
# it is actually listening, so there is no port race).
for _ in $(seq 1 50); do
  [ -s "$READY" ] && break
  sleep 0.1
done
if [ ! -s "$READY" ]; then
  echo "mock server did not become ready" >&2
  exit 1
fi
ADDR="$(cat "$READY")"
export YAYLIB_MOCK_BASE_URL="http://$ADDR"
export YAYLIB_MOCK_WS_URL="ws://$ADDR"
echo "==> mock server ready at $ADDR"

# Confirm health before handing off to the suites.
curl -fsS -m 5 "$YAYLIB_MOCK_BASE_URL/__mock/health" >/dev/null
echo "==> health ok"

rc=0
for suite in "${SUITES[@]}"; do
  echo "==> parity suite: $suite"
  case "$suite" in
    go)
      ( cd "$ROOT_DIR" && go test ./. ) || rc=1
      ;;
    ts)
      ( cd "$ROOT_DIR/packages/typescript" && npm test --silent ) || rc=1
      ;;
    py)
      ( cd "$ROOT_DIR/packages/python" && PYTHONPATH="$PWD" python3 -m pytest -q ) || rc=1
      ;;
    *)
      echo "unknown suite: $suite (want: go | ts | py)" >&2
      rc=1
      ;;
  esac
done

exit "$rc"
