// Upload parity tests (PORTING.md §8). Ports upload_test.go: filename
// shaping, the s3/ canonical-name contract, content-type mapping,
// per-category paths, the animated-GIF-stays-animated rule, login gating,
// caps, and the presigned PUT never leaking the Bearer.

import { createServer, type Server } from "node:http";
import type { AddressInfo } from "node:net";

import { encode as jpegEncode } from "jpeg-js";
import { GifWriter } from "omggif";
import { PNG } from "pngjs";
import { imageSize } from "image-size";

import { Client } from "../src/client";
import {
  contentTypeFor,
  formatYMD,
  imageFilename,
  randomFilenamePrefix,
  thumbnailFilename,
} from "../src/image";
import { mediaURL } from "../src/config";
import {
  userAvatarUpload,
  userPostUpload,
  chatMessageUpload,
  groupCreationIconUpload,
  reportUpload,
  videoCallSnapshotUpload,
} from "../src/upload";

let failed = 0;
function assert(name: string, cond: boolean, detail = ""): void {
  if (cond) {
    console.log(`PASS ${name}`);
  } else {
    failed++;
    console.log(`FAIL ${name}${detail ? `\n  ${detail}` : ""}`);
  }
}
function assertEq<T>(name: string, got: T, want: T): void {
  assert(name, got === want, `got=${String(got)} want=${String(want)}`);
}

// ---- image encoders (real bytes so image-size + the codecs work) ------

function encodeJpeg(w: number, h: number): Uint8Array {
  const data = new Uint8Array(w * h * 4);
  for (let i = 0; i < w * h; i++) {
    data[i * 4] = i % 256;
    data[i * 4 + 1] = (i * 3) % 256;
    data[i * 4 + 2] = 200;
    data[i * 4 + 3] = 255;
  }
  return new Uint8Array(jpegEncode({ width: w, height: h, data }, 80).data);
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

function encodeGif(w: number, h: number, frames = 1): Uint8Array {
  const buf = new Uint8Array(w * h * frames * 2 + 4096 + frames * 256 * 3);
  const gw = new GifWriter(buf, w, h, { loop: 0 });
  // omggif requires a power-of-two palette.
  const palette = [0x000000, 0xffffff, 0xc80000, 0x00c800];
  for (let f = 0; f < frames; f++) {
    const idx = new Uint8Array(w * h);
    for (let i = 0; i < w * h; i++) idx[i] = (i + f) % palette.length;
    gw.addFrame(0, 0, w, h, idx, { palette, delay: 10, disposal: 0 });
  }
  return buf.slice(0, gw.end());
}

// ---- fake Yay! API + S3 server ----------------------------------------

interface FakeState {
  objects: Map<string, Uint8Array>;
  contentTypes: Map<string, string>;
  putAuth: string[];
  imageNames: string[];
  videoFileName: string;
  presignedCalls: number;
  failPresigned: boolean;
  emptyPresigned: boolean;
  putFails: boolean;
}

async function withFakeServer<T>(
  init: Partial<FakeState>,
  fn: (baseURL: string, st: FakeState) => Promise<T>,
): Promise<T> {
  const st: FakeState = {
    objects: new Map(),
    contentTypes: new Map(),
    putAuth: [],
    imageNames: [],
    videoFileName: "",
    presignedCalls: 0,
    failPresigned: false,
    emptyPresigned: false,
    putFails: false,
    ...init,
  };
  const server: Server = createServer((req, res) => {
    const u = new URL(req.url ?? "", "http://x");
    const chunks: Buffer[] = [];
    req.on("data", (c) => chunks.push(c));
    req.on("end", () => {
      if (req.method === "GET" && u.pathname === "/v1/buckets/presigned_urls") {
        st.presignedCalls++;
        if (st.failPresigned) {
          res.writeHead(500).end("presigned fail");
          return;
        }
        const names = u.searchParams.getAll("file_names[]");
        st.imageNames = names.slice();
        const urls = st.emptyPresigned
          ? []
          : names.map((n) => ({ filename: `s3/${n}`, url: `${baseURL}/uploads/s3/${n}` }));
        res.writeHead(200, { "Content-Type": "application/json" }).end(
          JSON.stringify({ presigned_urls: urls }),
        );
        return;
      }
      if (req.method === "GET" && u.pathname === "/v1/users/presigned_url") {
        const name = u.searchParams.get("video_file_name") ?? "";
        st.videoFileName = name;
        res.writeHead(200, { "Content-Type": "application/json" }).end(
          JSON.stringify({ presigned_url: `${baseURL}/uploads/${name}` }),
        );
        return;
      }
      if (req.method === "PUT" && u.pathname.startsWith("/uploads/")) {
        st.putAuth.push(req.headers["authorization"] ?? "");
        if (st.putFails) {
          res.writeHead(500).end("put fail");
          return;
        }
        const key = u.pathname;
        st.objects.set(key, new Uint8Array(Buffer.concat(chunks)));
        st.contentTypes.set(key, req.headers["content-type"] ?? "");
        res.writeHead(200).end();
        return;
      }
      res.writeHead(404).end("unknown route");
    });
  });
  await new Promise<void>((r) => server.listen(0, r));
  const baseURL = `http://127.0.0.1:${(server.address() as AddressInfo).port}`;
  try {
    return await fn(baseURL, st);
  } finally {
    server.closeAllConnections?.();
    await new Promise<void>((r) => server.close(() => r()));
  }
}

const NO_RETRY = { maxAttempts: 0, baseDelay: 1, maxDelay: 1, retryOnPOST: false };

function loggedIn(baseURL: string, userID: number): Client {
  const c = new Client({ baseURL, retryPolicy: NO_RETRY });
  c.setLoginIdentity("u@example.com", userID);
  return c;
}

// ---- pure-function parity --------------------------------------------

function pureTests(): void {
  assertEq(
    "imageFilename: full shape",
    imageFilename("user/42/avatar", "ABC123", 1700000000000, 0, ".png", 200, 100),
    "user/42/avatar/ABC123_1700000000000_0_size_200x100.png",
  );
  assertEq(
    "imageFilename: no size suffix when 0",
    imageFilename("user/42/avatar", "ABC123", 1700000000000, 2, ".jpeg", 0, 0),
    "user/42/avatar/ABC123_1700000000000_2.jpeg",
  );
  assertEq(
    "imageFilename: empty ext -> .jpg",
    imageFilename("user/42/avatar", "ABC123", 1700000000000, 0, "", 10, 10),
    "user/42/avatar/ABC123_1700000000000_0_size_10x10.jpg",
  );
  assertEq(
    "thumbnailFilename: thumb_ prefix + orig size",
    thumbnailFilename("user/42/avatar", "ABC123", 1700000000000, 0, ".jpeg", 800, 600),
    "user/42/avatar/thumb_ABC123_1700000000000_0_size_800x600.jpeg",
  );

  for (const [fn, want] of [
    ["a.png", "image/png"],
    ["path/to/B.PNG", "image/png"],
    ["a.GIF", "image/gif"],
    ["a.jpg", "image/jpeg"],
    ["a.JPG", "image/jpeg"],
    ["a.mp4", "video/mp4"],
    ["unknown", "image/jpeg"],
    ["", "image/jpeg"],
    ["weird.xyz", "image/jpeg"],
  ] as const) {
    assertEq(`contentTypeFor(${fn})`, contentTypeFor(fn), want);
  }

  assertEq(
    "formatYMD: no zero padding",
    formatYMD(new Date(Date.UTC(2026, 3, 9))),
    "2026/4/9",
  );
  assertEq(
    "formatYMD: double digits",
    formatYMD(new Date(Date.UTC(2000, 0, 1))),
    "2000/1/1",
  );

  assertEq(
    "mediaURL: s3 prefix preserved",
    mediaURL("s3/user/7360590/avatar/abc_123_0_size_251x251.jpeg"),
    "https://cdn.yay.space/uploads/s3/user/7360590/avatar/abc_123_0_size_251x251.jpeg",
  );
  assertEq("mediaURL: leading slash trimmed", mediaURL("/s3/x.jpg"), "https://cdn.yay.space/uploads/s3/x.jpg");
  assertEq("mediaURL: empty", mediaURL(""), "");

  const allowed = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ";
  let prefixOK = true;
  for (let n = 1; n <= 32; n++) {
    const got = randomFilenamePrefix(n);
    if (got.length !== n || [...got].some((ch) => !allowed.includes(ch))) prefixOK = false;
  }
  assert("randomFilenamePrefix: length + alphabet", prefixOK);

  const now = new Date(Date.UTC(2026, 3, 26, 12));
  assertEq("category: userAvatar", userAvatarUpload(7360590).path(now), "user/7360590/avatar");
  assertEq("category: userPost (date bucketed)", userPostUpload(7360590).path(now), "user/7360590/posts/2026/4/26");
  assertEq(
    "category: chatMessage",
    chatMessageUpload(1, 2).path(now),
    "chatroom/1/user/2/messages/2026/4/26",
  );
  assertEq("category: groupCreationIcon", groupCreationIconUpload().path(now), "group/create/icon/2026/4/26");
  assertEq("category: report", reportUpload().path(now), "report/2026/4/26");
  assertEq("category: videoCallSnapshot", videoCallSnapshotUpload().path(now), "video_call_snapshot/2026/4/26");
  assertEq("category: report cap is 4", reportUpload().maxFiles(), 4);
  assertEq("category: videoCallSnapshot has no thumb", videoCallSnapshotUpload().thumbnail()[2], false);
}

// ---- scenarios --------------------------------------------------------

async function scenarioPostImagesHappyPath(): Promise<void> {
  await withFakeServer({}, async (baseURL, st) => {
    const c = loggedIn(baseURL, 7360590);
    const png = encodePng(30, 40);
    const gif = encodeGif(50, 50);
    const jpg = encodeJpeg(60, 80);
    const names = await c.uploadPostImages([
      { filename: "selfie.png", body: png },
      { filename: "sticker.gif", body: gif },
      { filename: "photo.jpg", body: jpg },
    ]);
    assertEq("postImages: 3 names returned", names.length, 3);

    const wantExt = [".png", ".gif", ".jpg"];
    const wantCT = ["image/png", "image/gif", "image/jpeg"];
    const wantSize = ["_size_30x40", "_size_50x50", "_size_60x80"];
    const year = new Date().getUTCFullYear();
    for (let i = 0; i < 3; i++) {
      const name = names[i];
      assert(`postImages[${i}] ext ${wantExt[i]}`, name.endsWith(wantExt[i]), name);
      assert(
        `postImages[${i}] s3/user/7360590/posts/<year>/ prefix`,
        name.startsWith(`s3/user/7360590/posts/${year}/`),
        name,
      );
      assert(`postImages[${i}] size suffix ${wantSize[i]}`, name.includes(wantSize[i]), name);
      assertEq(`postImages[${i}] content-type`, st.contentTypes.get(`/uploads/${name}`), wantCT[i]);
    }
    let thumbCount = 0;
    for (const [p, ct] of st.contentTypes) {
      if (!p.includes("/thumb_")) continue;
      thumbCount++;
      if (p.endsWith(".gif")) assertEq("gif thumb content-type", ct, "image/gif");
      else if (p.endsWith(".jpeg")) assertEq("jpeg thumb content-type", ct, "image/jpeg");
    }
    assertEq("postImages: 3 thumbnails uploaded", thumbCount, 3);
    assertEq("postImages: 6 file_names[] (3 main + 3 thumb)", st.imageNames.length, 6);
  });
}

async function scenarioAvatarHappyAndShape(): Promise<void> {
  await withFakeServer({}, async (baseURL, st) => {
    const c = loggedIn(baseURL, 1);
    const name = await c.uploadAvatarImage({ filename: "x.png", body: encodePng(12, 13) });
    assert(
      "avatar: filename shape",
      /^s3\/user\/1\/avatar\/[0-9A-Za-z]{16}_\d{13}_0_size_12x13\.png$/.test(name),
      name,
    );
    assert("avatar: main object PUT", st.objects.has(`/uploads/${name}`));
  });
}

async function scenarioRequiresLogin(): Promise<void> {
  const c = new Client({ baseURL: "http://127.0.0.1:1" });
  let threw = false;
  try {
    await c.uploadAvatarImage({ filename: "x.jpg", body: encodeJpeg(5, 5) });
  } catch (e) {
    threw = (e as Error).message.includes("not logged in");
  }
  assert("avatar without login: 'not logged in' error", threw);
}

async function scenarioUnknownExtFallsBackToJpeg(): Promise<void> {
  await withFakeServer({}, async (baseURL, st) => {
    const c = loggedIn(baseURL, 1);
    for (const fn of ["noext", "weird.xyz"]) {
      const name = await c.uploadAvatarImage({ filename: fn, body: encodeJpeg(5, 5) });
      assert(`unknown ext (${fn}) -> .jpg`, name.endsWith(".jpg"), name);
      assertEq(`unknown ext (${fn}) content-type`, st.contentTypes.get(`/uploads/${name}`), "image/jpeg");
    }
  });
}

async function scenarioEmptyAndOverflow(): Promise<void> {
  const c = loggedIn("http://does-not-matter", 1);
  let e1 = false;
  try {
    await c.uploadPostImages([]);
  } catch {
    e1 = true;
  }
  assert("postImages([]) rejects", e1);
  const tooMany = Array.from({ length: 10 }, () => ({ filename: "x.jpg", body: encodeJpeg(2, 2) }));
  let e2 = false;
  try {
    await c.uploadPostImages(tooMany);
  } catch {
    e2 = true;
  }
  assert("postImages(10) rejects (cap 9)", e2);
  let e3 = false;
  try {
    await c.uploadReportImages(Array.from({ length: 5 }, () => ({ filename: "x.jpg", body: encodeJpeg(2, 2) })));
  } catch {
    e3 = true;
  }
  assert("reportImages(5) rejects (cap 4)", e3);
}

async function scenarioPutFailure(): Promise<void> {
  await withFakeServer({ putFails: true }, async (baseURL) => {
    const c = loggedIn(baseURL, 1);
    let msg = "";
    try {
      await c.uploadAvatarImage({ filename: "a.png", body: encodePng(5, 5) });
    } catch (e) {
      msg = (e as Error).message;
    }
    assert("put failure: 'PUT 500' in message", msg.includes("PUT 500"), msg);
  });
}

async function scenarioPresignedFailure(): Promise<void> {
  await withFakeServer({ failPresigned: true }, async (baseURL) => {
    const c = loggedIn(baseURL, 1);
    let threw = false;
    try {
      await c.uploadAvatarImage({ filename: "a.png", body: encodePng(5, 5) });
    } catch {
      threw = true;
    }
    assert("presigned 500: surfaces error", threw);
  });
}

async function scenarioEmptyPresigned(): Promise<void> {
  await withFakeServer({ emptyPresigned: true }, async (baseURL) => {
    const c = loggedIn(baseURL, 1);
    let msg = "";
    try {
      await c.uploadAvatarImage({ filename: "x.png", body: encodePng(5, 5) });
    } catch (e) {
      msg = (e as Error).message;
    }
    assert("empty presigned: 'count mismatch'", msg.includes("count mismatch"), msg);
  });
}

async function scenarioNonImage(): Promise<void> {
  await withFakeServer({}, async (baseURL) => {
    const c = loggedIn(baseURL, 1);
    let msg = "";
    try {
      await c.uploadAvatarImage({ filename: "broken.jpg", body: new TextEncoder().encode("not an image at all") });
    } catch (e) {
      msg = (e as Error).message;
    }
    assert("non-image body: 'decode image' error", msg.includes("decode image"), msg);
  });
}

async function scenarioVideoHappyAndNil(): Promise<void> {
  await withFakeServer({}, async (baseURL, st) => {
    const c = new Client({ baseURL, retryPolicy: NO_RETRY });
    const mp4 = new TextEncoder().encode("\x00\x00\x00 ftypisom--mp4-bytes--");
    const name = await c.uploadVideo(mp4);
    assert("video: .mp4 suffix", name.endsWith(".mp4"), name);
    assertEq("video: content-type", st.contentTypes.get(`/uploads/${name}`), "video/mp4");
    assertEq("video: server saw same filename", st.videoFileName, name);

    let threw = false;
    try {
      await c.uploadVideo(null as unknown as Uint8Array);
    } catch {
      threw = true;
    }
    assert("video nil body: rejects", threw);
  });
}

async function scenarioGifThumbStaysGif(): Promise<void> {
  await withFakeServer({}, async (baseURL, st) => {
    const c = loggedIn(baseURL, 1);
    await c.uploadAvatarImage({ filename: "x.gif", body: encodeGif(800, 400, 3) });
    assertEq("gif avatar: 2 file_names[] (main gif + gif thumb)", st.imageNames.length, 2);
    let thumb: Uint8Array | undefined;
    let thumbPath = "";
    for (const [p, b] of st.objects) {
      if (p.includes("/thumb_")) {
        thumb = b;
        thumbPath = p;
        break;
      }
    }
    assert("gif avatar: thumbnail PUT exists", thumb != null);
    assert("gif thumb path ends .gif", thumbPath.endsWith(".gif"), thumbPath);
    assertEq("gif thumb content-type", st.contentTypes.get(thumbPath), "image/gif");
    if (thumb) {
      const dim = imageSize(thumb);
      assertEq("gif thumb format", dim.type, "gif");
      assertEq("gif thumb width", dim.width, 200);
      assertEq("gif thumb height", dim.height, 100);
    }
  });
}

async function scenarioJpegThumbDims(): Promise<void> {
  await withFakeServer({}, async (baseURL, st) => {
    const c = loggedIn(baseURL, 7);
    await c.uploadAvatarImage({ filename: "me.jpg", body: encodeJpeg(800, 600) });
    let thumb: Uint8Array | undefined;
    for (const [p, b] of st.objects) {
      if (p.includes("/thumb_")) {
        thumb = b;
        break;
      }
    }
    assert("jpeg avatar: thumbnail PUT exists", thumb != null);
    if (thumb) {
      const dim = imageSize(thumb);
      assertEq("jpeg thumb format", dim.type, "jpg");
      assertEq("jpeg thumb width (fit 200x200)", dim.width, 200);
      assertEq("jpeg thumb height (fit 200x200)", dim.height, 150);
    }
  });
}

async function scenarioVideoCallSnapshotSkipsThumb(): Promise<void> {
  await withFakeServer({}, async (baseURL, st) => {
    const c = new Client({ baseURL, retryPolicy: NO_RETRY });
    const name = await c.uploadVideoCallSnapshotImage({ filename: "snap.jpg", body: encodeJpeg(100, 100) });
    assert("snapshot: name returned", name.length > 0);
    assertEq("snapshot: only 1 file_names[] (no thumb)", st.imageNames.length, 1);
    let anyThumb = false;
    for (const p of st.contentTypes.keys()) if (p.includes("/thumb_")) anyThumb = true;
    assert("snapshot: no thumbnail uploaded", !anyThumb);
  });
}

async function scenarioNoBearerLeakOnPut(): Promise<void> {
  await withFakeServer({}, async (baseURL, st) => {
    const c = loggedIn(baseURL, 1);
    c.setTokens("SECRET-ACCESS-TOKEN", "REF");
    await c.uploadAvatarImage({ filename: "x.png", body: encodePng(5, 5) });
    assert("PUT requests captured", st.putAuth.length > 0);
    assert(
      "no PUT carried Authorization (presigned URL signs via query)",
      st.putAuth.every((a) => a === ""),
      st.putAuth.join(","),
    );
  });
}

(async () => {
  pureTests();
  await scenarioPostImagesHappyPath();
  await scenarioAvatarHappyAndShape();
  await scenarioRequiresLogin();
  await scenarioUnknownExtFallsBackToJpeg();
  await scenarioEmptyAndOverflow();
  await scenarioPutFailure();
  await scenarioPresignedFailure();
  await scenarioEmptyPresigned();
  await scenarioNonImage();
  await scenarioVideoHappyAndNil();
  await scenarioGifThumbStaysGif();
  await scenarioJpegThumbDims();
  await scenarioVideoCallSnapshotSkipsThumb();
  await scenarioNoBearerLeakOnPut();
  process.exit(failed === 0 ? 0 : 1);
})().catch((e) => {
  console.error("test crashed:", e);
  process.exit(2);
});
