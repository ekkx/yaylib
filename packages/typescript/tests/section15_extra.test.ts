// Extra §15 behavior-parity scenarios driven against the shared
// behavior-contract server. Skips when the server is unconfigured so a
// standalone `npm test` still passes.
//
// Covers S29: a success answered with a non-200 2xx (e.g. login's 201)
// must still decode into the operation's all-optional success model. The
// generated runtime keys "success" off the 2xx class, not the exact
// documented status code, so the body is parsed rather than discarded.

import { assert, mockClient, run, skipUnlessMock } from "./parity_harness";

// PORTING:S29
// ok-status-201: the server answers the generic GET with HTTP 201 and an
// empty-object body. The typed call (getBucketPresignedUrls, which runs
// the full JSON decode via response.value()) must RESOLVE — proving the
// runtime treated the non-200 2xx as success and parsed the body instead
// of throwing on the unexpected status.
async function anyTwoXXIsSuccess(): Promise<void> {
  const c = mockClient("ok-status-201");
  let resolved = false;
  let value: unknown;
  let threw: unknown;
  try {
    value = await c.bucketsAPI.getBucketPresignedUrls({});
    resolved = true;
  } catch (e) {
    threw = e;
  }
  assert(
    "anyTwoXXIsSuccess: typed call resolved (201 treated as success, body parsed)",
    resolved,
    threw ? `threw: ${String(threw)}` : "",
  );
  assert(
    "anyTwoXXIsSuccess: a defined success object was returned",
    value !== undefined && value !== null,
    `value=${String(value)}`,
  );
}

skipUnlessMock("section15 extra parity");
run(async () => {
  await anyTwoXXIsSuccess();
});
