// Golden-vector parity with PORTING.md §7.1 — same inputs as the Go
// reference's signing_test.go and the documented test vectors. A drift
// in any of the three outputs means the formula changed (encoding, key
// choice, message body, trailing whitespace) and the Go / TS / Python
// ports are no longer byte-compatible.
//
// Run with `npx tsx tests/signing.test.ts` (zero-config).

import {
  generateSignedInfoAt,
  generateSignedVersion,
  generateXJwt,
} from "../src/signing";
import {
  DEFAULT_API_KEY,
  DEFAULT_API_VERSION_KEY,
  DEFAULT_API_VERSION_NAME,
} from "../src/config";

const cases: Array<{ name: string; got: string; want: string }> = [
  {
    name: "signed_info (default APIKey, zero-UUID, ts=1700000000)",
    got: generateSignedInfoAt(
      {
        apiKey: DEFAULT_API_KEY,
        deviceUUID: "00000000-0000-0000-0000-000000000000",
      },
      1700000000,
    ).value,
    want: "253b7ca3e67590204034a73326d2be87",
  },
  {
    name: "signed_version (default key, name=4.26)",
    got: generateSignedVersion({
      apiVersionKey: DEFAULT_API_VERSION_KEY,
      apiVersionName: DEFAULT_API_VERSION_NAME,
    }),
    want: "zyKEqdZqaEqJjSRItHvchU9HmgWHNLtn02sWXwzQ4iQ=",
  },
  {
    name: "x_jwt (default key, iat=1700000000)",
    got: generateXJwt({ apiVersionKey: DEFAULT_API_VERSION_KEY, now: 1700000000 }),
    want:
      "eyJhbGciOiJIUzI1NiJ9.eyJpYXQiOjE3MDAwMDAwMDAsImV4cCI6MTcwMDAwMDAwNX0.-UYjGTM53-Uwe8tNGjBf4ZLx644_zUEK8fpW1w16KfA",
  },
];

let failed = 0;
for (const c of cases) {
  const ok = c.got === c.want;
  if (!ok) failed++;
  // eslint-disable-next-line no-console
  console.log(`${ok ? "PASS" : "FAIL"} ${c.name}`);
  if (!ok) {
    // eslint-disable-next-line no-console
    console.log(`  got=${c.got}\n  want=${c.want}`);
  }
}
process.exit(failed === 0 ? 0 : 1);
