#!/usr/bin/env bash
#
# Cross-language behavior-parity coverage matrix (PORTING.md §15.1).
#
# PORTING.md §15 is the cross-language behavior contract. Each scenario
# has a stable ID `S<n>`; the test covering it in each language carries
# a `PORTING:S<n>` tag comment. This script reads the IDs from §15 and
# greps the Go / TypeScript / Python test trees for the tags, printing a
# scenario × language matrix so drift between the contract and the three
# suites is mechanically visible instead of found by manual audit.
#
#   scripts/parity-matrix.sh            # print the matrix (always exit 0)
#   scripts/parity-matrix.sh --strict   # exit 1 if any S<n> is not
#                                        # tagged in all three languages
#
# Untagged-but-existing tests show up as gaps to backfill, not as
# failures (tagging of pre-existing tests is incremental — this matrix
# is the running TODO).
#
# A reporting tool: `set -u` only — greps with no match are expected
# (untagged scenarios) and must not abort the run.
set -u

ROOT="$(cd "$(dirname "${BASH_SOURCE[0]}")/.." && pwd)"
PORTING="$ROOT/PORTING.md"
strict=0
[ "${1:-}" = "--strict" ] && strict=1

# Scenario IDs in declaration order: lines like `- **S12** ...`.
ids=$(grep -oE '\*\*S[0-9]+\*\*' "$PORTING" | tr -d '*' | awk '!seen[$0]++')
if [ -z "$ids" ]; then
  echo "no S<n> scenario IDs found in $PORTING" >&2
  exit 2
fi

# Collect the set of tagged IDs per language. A tag is `PORTING:S1` or
# `PORTING:S7,S8` in a test file; one comment may list several.
tags_for() { # <glob-root> <find-args...>
  local root="$1"; shift
  [ -d "$root" ] || return 0
  grep -rhoE 'PORTING:S[0-9]+(,S[0-9]+)*' "$root" "$@" 2>/dev/null \
    | sed 's/PORTING://' | tr ',' '\n' | sort -u
}

go_tags="$(cd "$ROOT" && grep -rhoE 'PORTING:S[0-9]+(,S[0-9]+)*' --include='*_test.go' . 2>/dev/null | sed 's/PORTING://' | tr ',' '\n' | sort -u)"
ts_tags="$(tags_for "$ROOT/packages/typescript/tests" --include='*.ts')"
py_tags="$(tags_for "$ROOT/packages/python/tests" --include='*.py')"

has() { printf '%s\n' "$2" | grep -qx "$1"; }

printf '%-5s | %-3s | %-3s | %-3s\n' "ID" "go" "ts" "py"
printf -- '------+-----+-----+-----\n'
missing=0
for id in $ids; do
  g=·; t=·; p=·
  if has "$id" "$go_tags"; then g=✓; fi
  if has "$id" "$ts_tags"; then t=✓; fi
  if has "$id" "$py_tags"; then p=✓; fi
  printf '%-5s |  %s  |  %s  |  %s\n' "$id" "$g" "$t" "$p"
  if [ "$g$t$p" != "✓✓✓" ]; then missing=$((missing + 1)); fi
done

total=$(printf '%s\n' "$ids" | wc -l | tr -d ' ')
echo
echo "scenarios: $total | fully covered: $((total - missing)) | needs tagging/coverage: $missing"
if [ "$strict" = 1 ] && [ "$missing" -gt 0 ]; then
  exit 1
fi
exit 0
