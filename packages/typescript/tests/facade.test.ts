// Flat operation facade (PORTING.md §15 S33 / §2).
//
// Every operation is reachable as client.<op> with no service prefix,
// and a hand-written wrapper of the same name shadows the generated op.
// Pure local check — no server is involved.

import { Client } from "../src/index";

let failed = 0;
function assert(name: string, cond: boolean, detail = ""): void {
  if (cond) console.log(`PASS ${name}`);
  else {
    failed++;
    console.log(`FAIL ${name}${detail ? `\n  ${detail}` : ""}`);
  }
}

// PORTING:S33
async function flatOpsAndWrapperShadow(): Promise<void> {
  const c = new Client();

  // Operations are promoted onto the client directly (no client.postsAPI.).
  assert(
    "client.getTimeline is a flat method",
    typeof (c as unknown as Record<string, unknown>).getTimeline === "function",
  );
  assert(
    "client.getRecommendedTimeline is a flat method",
    typeof (c as unknown as Record<string, unknown>).getRecommendedTimeline ===
      "function",
  );

  // The hand-written cached wrapper shadows the generated loginWithEmail
  // op: empty credentials fail fast with the wrapper's validation error,
  // never an HTTP attempt.
  let shadow = false;
  try {
    await c.loginWithEmail({ email: "", password: "" });
  } catch (e) {
    shadow = /requires email and password/.test(String((e as Error)?.message));
  }
  assert("loginWithEmail resolves to the cached wrapper (shadow)", shadow);
}

(async () => {
  await flatOpsAndWrapperShadow();
  process.exit(failed === 0 ? 0 : 1);
})().catch((e) => {
  console.error("test crashed:", e);
  process.exit(2);
});
