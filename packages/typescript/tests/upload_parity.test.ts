// Upload behavior parity (port of upload_parity_test.go). Driven against
// the shared server's upload contract: GET /v1/buckets/presigned_urls
// returns one server-canonical ("s3/"-prefixed) filename + an in-process
// S3 URL per requested name; GET /v1/users/presigned_url does the same
// for video; the S3 receiver answers 200, or 403 if the PUT carries
// Authorization.
//
// The shared S3 receiver intentionally does not echo the Content-Type,
// the stored bytes, or the file_names[] it saw (behavioral-fidelity
// only). So the parity assertions are the client-observable ones:
// success/error, the returned canonical filename shape (the server just
// prepends "s3/" to the SDK-built name), and the server-enforced
// no-bearer rule on the presigned PUT. PUT content inspection, image
// processing, the empty-presigned / non-image guards and the pure
// filename / path / MIME helpers stay as local fixtures in
// upload.test.ts.

import { encode as jpegEncode } from "jpeg-js";
import { PNG } from "pngjs";

import { assert, mockClient, run, skipUnlessMock } from "./parity_harness";

function encodeJpeg(w: number, h: number): Uint8Array {
  const data = new Uint8Array(w * h * 4);
  for (let i = 0; i < w * h; i++) {
    data[i * 4] = i % 256;
    data[i * 4 + 1] = (i * 3) % 256;
    data[i * 4 + 2] = 200;
    data[i * 4 + 3] = 255;
  }
  return new Uint8Array(jpegEncode({ width: w, height: h, data }, 75).data);
}

function encodePng(w: number, h: number): Uint8Array {
  const png = new PNG({ width: w, height: h });
  for (let i = 0; i < w * h; i++) {
    png.data[i * 4] = i % 256;
    png.data[i * 4 + 1] = (i * 5) % 256;
    png.data[i * 4 + 2] = 120;
    png.data[i * 4 + 3] = 255;
  }
  return new Uint8Array(PNG.sync.write(png));
}

async function uploadAvatarHappyPath(): Promise<void> {
  const c = mockClient("");
  c.setLoginIdentity("", 42);
  const name = await c.uploadAvatarImage({
    filename: "me.jpg",
    body: encodeJpeg(100, 100),
  });
  assert(
    "uploadAvatar: name has s3/user/42/avatar/ prefix",
    name.startsWith("s3/user/42/avatar/"),
    `name=${name}`,
  );
  assert(
    "uploadAvatar: name has _size_100x100 suffix",
    name.includes("_size_100x100"),
    `name=${name}`,
  );
}

async function uploadAvatarFilenameShape(): Promise<void> {
  const c = mockClient("");
  c.setLoginIdentity("", 1);
  const name = await c.uploadAvatarImage({
    filename: "x.png",
    body: encodePng(12, 13),
  });
  const re = /^s3\/user\/1\/avatar\/[0-9A-Za-z]{16}_\d{13}_0_size_12x13\.png$/;
  assert(
    "uploadAvatar: canonical filename shape matches the contract",
    re.test(name),
    `name=${name}`,
  );
}

async function uploadAvatarPresignedFailure(): Promise<void> {
  // fail-503-times-1 makes the first request in the session — the
  // presigned-URL GET — fail; with retry disabled the upload errors.
  const c = mockClient("fail-503-times-1", {
    retryPolicy: { maxAttempts: 1, baseDelay: 1, maxDelay: 1, retryOnPOST: false },
  });
  c.setLoginIdentity("", 1);
  let threw = false;
  try {
    await c.uploadAvatarImage({ filename: "a.png", body: encodePng(5, 5) });
  } catch {
    threw = true;
  }
  assert("uploadAvatar: failing presigned fetch surfaces an error", threw);
}

async function uploadVideoHappyPath(): Promise<void> {
  const c = mockClient("");
  const mp4 = new TextEncoder().encode("\x00\x00\x00 ftypisom--mp4-bytes--");
  const name = await c.uploadVideo(mp4);
  assert(
    "uploadVideo: name has .mp4 suffix",
    name.endsWith(".mp4"),
    `name=${name}`,
  );
}

async function presignedPUTDoesNotLeakBearer(): Promise<void> {
  // Tokens are set, but the presigned PUT must authenticate via the
  // signed query only. The shared S3 receiver answers 403 if the PUT
  // carries Authorization, so a successful upload proves no bearer
  // leaked — the check is server-enforced and identical across the
  // three languages.
  const c = mockClient("");
  c.setLoginIdentity("", 1);
  c.setTokens("SECRET-ACCESS-TOKEN", "REF");
  const name = await c.uploadAvatarImage({
    filename: "x.png",
    body: encodePng(5, 5),
  });
  assert(
    "presignedPUT: upload succeeded, so no bearer leaked to the PUT",
    name.startsWith("s3/user/1/avatar/"),
    `name=${name}`,
  );
}

skipUnlessMock("upload parity");
run(async () => {
  await uploadAvatarHappyPath();
  await uploadAvatarFilenameShape();
  await uploadAvatarPresignedFailure();
  await uploadVideoHappyPath();
  await presignedPUTDoesNotLeakBearer();
});
