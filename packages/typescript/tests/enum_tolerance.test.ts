// Enum decode tolerance (PORTING.md §15 #26).
//
// The server may return an enum value the SDK doesn't know about. Typed
// JSON decode must NOT reject it — an unknown value round-trips through
// as the raw string, while a known value decodes to its typed constant.
// This is a pure local-decode check; no server is involved.

import { PostFromJSON } from "../src/gen/models/Post";
import { PostType } from "../src/gen/models/PostType";

let failed = 0;
function assert(name: string, cond: boolean, detail = ""): void {
  if (cond) console.log(`PASS ${name}`);
  else {
    failed++;
    console.log(`FAIL ${name}${detail ? `\n  ${detail}` : ""}`);
  }
}

// PORTING:S28
function unknownEnumDoesNotThrow(): void {
  let threw = false;
  let decoded: unknown;
  try {
    decoded = PostFromJSON({ post_type: "__mock_unknown_enum__" });
  } catch {
    threw = true;
  }
  assert("unknown enum value does not throw on decode", !threw);
  assert(
    "unknown enum value round-trips as the raw string",
    !threw && (decoded as { postType?: unknown }).postType === "__mock_unknown_enum__",
    String((decoded as { postType?: unknown })?.postType),
  );
}

// PORTING:S28
function knownEnumDecodesToConstant(): void {
  const post = PostFromJSON({ post_type: "image" });
  assert(
    "known enum value decodes to the typed constant",
    post.postType === PostType.Image,
    String(post.postType),
  );
  // The decoded value is usable in an equality switch against the
  // exported constant set.
  assert(
    "decoded constant compares equal to PostType.Image",
    post.postType === "image",
    String(post.postType),
  );
}

(async () => {
  unknownEnumDoesNotThrow();
  knownEnumDecodesToConstant();
  process.exit(failed === 0 ? 0 : 1);
})().catch((e) => {
  console.error("test crashed:", e);
  process.exit(2);
});
