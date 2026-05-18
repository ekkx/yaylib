// PORTING.md §8: the thirteen typed upload categories + uploadVideo. The
// category table (method ↔ bucket path ↔ thumbnail size ↔ file cap) is
// replicated from upload.go with the same names and values; the wire
// pieces that aren't visible from the table (filename shaping, the s3/
// canonical-name contract, the animated-GIF thumbnail rule) live in
// image.ts.
//
// The presigned PUT bypasses the generated middleware chain entirely: the
// headers/auth middleware would inject `Authorization: Bearer …` and the
// Yay! headers, which conflict with the presigned URL's own signature and
// (per the leak test) must never reach S3. Presigned-URL issuance and the
// ws-token fetch still go through the wrapped client so they carry auth
// and retry, exactly as upload.go routes them through the normal client.

import {
  MAX_IMAGES_PER_UPLOAD,
  MAX_REPORT_IMAGES_PER_UPLOAD,
  contentTypeFor,
  formatYMD,
  imageFilename,
  makeThumbnailFor,
  normalizeImageExt,
  probeImageSize,
  randomFilenamePrefix,
  thumbnailFilename,
  videoFilename,
} from "./image.js";

export { MAX_IMAGES_PER_UPLOAD, MAX_REPORT_IMAGES_PER_UPLOAD };

// Upload is a single file passed to one of the upload<Xxx>Image[s]
// methods. `filename` is used to derive the extension and Content-Type;
// `body` is buffered fully before the request is issued (the thumbnail
// step needs the whole image anyway).
export interface Upload {
  filename: string;
  body: Uint8Array | ArrayBuffer | Blob;
}

// VideoBody is the single byte stream uploadVideo accepts. A Blob carries
// its own length and streams directly; the other shapes are buffered
// because the presigned PUT needs a known Content-Length.
export type VideoBody = Blob | Uint8Array | ArrayBuffer | ReadableStream<Uint8Array>;

// UploadCategory selects the bucket-side path prefix and thumbnail
// dimensions for one upload kind. Callers never construct these directly —
// the typed Client methods know which category they correspond to.
export interface UploadCategory {
  // path returns the category-specific prefix. `now` is shared across
  // every file in a multi-image call so date-bucketed categories get one
  // consistent timestamp.
  path(now: Date): string;
  // thumbnail returns [w, h, true] for the thumbnail box, or [0, 0,
  // false] for categories that never produce a thumbnail.
  thumbnail(): [number, number, boolean];
  // maxFiles caps how many images this category accepts in one call.
  maxFiles(): number;
}

const DEFAULT_MAX_FILES = MAX_IMAGES_PER_UPLOAD;

export function userAvatarUpload(userID: number): UploadCategory {
  return {
    path: () => `user/${userID}/avatar`,
    thumbnail: () => [200, 200, true],
    maxFiles: () => DEFAULT_MAX_FILES,
  };
}

export function userCoverUpload(userID: number): UploadCategory {
  return {
    path: () => `user/${userID}/cover`,
    thumbnail: () => [300, 150, true],
    maxFiles: () => DEFAULT_MAX_FILES,
  };
}

export function userPostUpload(userID: number): UploadCategory {
  return {
    path: (now) => `user/${userID}/posts/${formatYMD(now)}`,
    thumbnail: () => [450, 450, true],
    maxFiles: () => DEFAULT_MAX_FILES,
  };
}

export function chatMessageUpload(roomID: number, userID: number): UploadCategory {
  return {
    path: (now) => `chatroom/${roomID}/user/${userID}/messages/${formatYMD(now)}`,
    thumbnail: () => [450, 450, true],
    maxFiles: () => DEFAULT_MAX_FILES,
  };
}

export function chatBackgroundUpload(roomID: number): UploadCategory {
  return {
    path: () => `chatroom/${roomID}/background`,
    thumbnail: () => [200, 200, true],
    maxFiles: () => DEFAULT_MAX_FILES,
  };
}

export function groupIconUpload(groupID: number): UploadCategory {
  return {
    path: () => `group/${groupID}/icon`,
    thumbnail: () => [300, 300, true],
    maxFiles: () => DEFAULT_MAX_FILES,
  };
}

export function groupCoverUpload(groupID: number): UploadCategory {
  return {
    path: () => `group/${groupID}/cover`,
    thumbnail: () => [300, 300, true],
    maxFiles: () => DEFAULT_MAX_FILES,
  };
}

export function groupCreationIconUpload(): UploadCategory {
  return {
    path: (now) => `group/create/icon/${formatYMD(now)}`,
    thumbnail: () => [300, 300, true],
    maxFiles: () => DEFAULT_MAX_FILES,
  };
}

export function groupCreationCoverUpload(): UploadCategory {
  return {
    path: (now) => `group/create/cover/${formatYMD(now)}`,
    thumbnail: () => [300, 300, true],
    maxFiles: () => DEFAULT_MAX_FILES,
  };
}

export function signupAvatarUpload(): UploadCategory {
  return {
    path: (now) => `user/create/${formatYMD(now)}`,
    thumbnail: () => [200, 200, true],
    maxFiles: () => DEFAULT_MAX_FILES,
  };
}

export function threadIconUpload(groupID: number): UploadCategory {
  return {
    path: (now) => `group/${groupID}/threads/${formatYMD(now)}`,
    thumbnail: () => [300, 300, true],
    maxFiles: () => DEFAULT_MAX_FILES,
  };
}

export function reportUpload(): UploadCategory {
  return {
    path: (now) => `report/${formatYMD(now)}`,
    thumbnail: () => [450, 450, true],
    maxFiles: () => MAX_REPORT_IMAGES_PER_UPLOAD,
  };
}

export function videoCallSnapshotUpload(): UploadCategory {
  return {
    path: (now) => `video_call_snapshot/${formatYMD(now)}`,
    thumbnail: () => [0, 0, false],
    maxFiles: () => DEFAULT_MAX_FILES,
  };
}

// PresignedURL is the shape both presigned endpoints hand back per file.
export interface PresignedURL {
  filename?: string | null;
  url?: string | null;
}

// UploadDeps is the slice of Client behaviour the upload core needs. The
// Client supplies it; keeping it abstract avoids a client.ts <-> upload.ts
// import cycle and makes the flow unit-testable.
export interface UploadDeps {
  userID(): number;
  getPresignedURLs(fileNames: string[]): Promise<PresignedURL[]>;
  getVideoPresignedURL(videoFileName: string): Promise<string>;
  // rawFetch bypasses the generated middleware chain — see file header.
  rawFetch(url: string, init: RequestInit): Promise<Response>;
}

function requireUserID(deps: UploadDeps): number {
  const uid = deps.userID();
  if (!uid) {
    throw new Error("yaylib: not logged in (call loginWithEmail before user-bound uploads)");
  }
  return uid;
}

async function toBytes(body: Uint8Array | ArrayBuffer | Blob): Promise<Uint8Array> {
  if (body instanceof Uint8Array) return body;
  if (body instanceof ArrayBuffer) return new Uint8Array(body);
  // Blob
  return new Uint8Array(await body.arrayBuffer());
}

interface PutJob {
  body: Uint8Array;
  filename: string;
  contentType: string;
}

async function putToPresignedURL(
  deps: UploadDeps,
  rawURL: string,
  body: Uint8Array | Blob,
  size: number,
  contentType: string,
  signal: AbortSignal,
): Promise<void> {
  let res: Response;
  try {
    // size is unused at the header level: fetch derives Content-Length
    // from the Uint8Array/Blob body, and setting it manually trips
    // undici's duplicate-header guard. The presigned URL signs only the
    // host, so a matching Content-Type is all S3 intermediaries need.
    void size;
    res = await deps.rawFetch(rawURL, {
      method: "PUT",
      headers: { "Content-Type": contentType },
      body: body as unknown as BodyInit,
      signal,
    });
  } catch (err) {
    throw new Error(`PUT failed: ${err instanceof Error ? err.message : String(err)}`);
  }
  if (Math.floor(res.status / 100) !== 2) {
    let preview = "";
    try {
      preview = (await res.text()).slice(0, 512).trim();
    } catch {
      /* ignore */
    }
    throw new Error(`PUT ${res.status}: ${preview}`);
  }
}

// uploadImages performs the full presigned-URL + parallel-PUT dance shared
// by every image-upload method, mirroring upload.go's uploadImages. The
// returned names are the server-canonical (s3/-prefixed) main-image names;
// thumbnails live at sibling paths the server resolves automatically.
export async function uploadImages(
  deps: UploadDeps,
  category: UploadCategory,
  files: Upload[],
): Promise<string[]> {
  if (!category) throw new Error("yaylib: upload requires a category");
  if (!files || files.length === 0) {
    throw new Error("yaylib: upload requires at least one file");
  }
  const cap = category.maxFiles();
  if (files.length > cap) {
    throw new Error(`yaylib: this upload accepts at most ${cap} files (got ${files.length})`);
  }

  const now = new Date();
  const prefix = randomFilenamePrefix(16);
  const ts = now.getTime();
  const categoryPath = category.path(now);
  const [thumbW, thumbH, hasThumb] = category.thumbnail();

  const mainJobs: PutJob[] = [];
  const thumbJobs: PutJob[] = [];

  for (let i = 0; i < files.length; i++) {
    let body: Uint8Array;
    try {
      body = await toBytes(files[i].body);
    } catch (err) {
      throw new Error(`yaylib: read file ${i}: ${err instanceof Error ? err.message : String(err)}`);
    }
    const ext = normalizeImageExt(files[i].filename);
    // A failed decode means the size suffix would be wrong and the body
    // can't be thumbnailed; uploading it anyway just triggers server-side
    // moderation deletion, so surface the error.
    let w: number;
    let h: number;
    try {
      const dim = probeImageSize(body);
      w = dim.width;
      h = dim.height;
    } catch (err) {
      throw new Error(`yaylib: decode image ${i}: ${err instanceof Error ? err.message : String(err)}`);
    }
    const mainName = imageFilename(categoryPath, prefix, ts, i, ext, w, h);
    mainJobs.push({ body, filename: mainName, contentType: contentTypeFor(mainName) });

    if (hasThumb) {
      let thumb;
      try {
        thumb = makeThumbnailFor(body, ext, thumbW, thumbH);
      } catch (err) {
        // A missing/mismatched thumbnail makes the server delete the main
        // asset too, so fail loudly instead of silently skipping it.
        throw new Error(`yaylib: make thumbnail for image ${i}: ${err instanceof Error ? err.message : String(err)}`);
      }
      const thumbName = thumbnailFilename(categoryPath, prefix, ts, i, thumb.ext, w, h);
      thumbJobs.push({ body: thumb.body, filename: thumbName, contentType: thumb.contentType });
    }
  }

  const allJobs = [...mainJobs, ...thumbJobs];
  const allFileNames = allJobs.map((j) => j.filename);

  let urls: PresignedURL[];
  try {
    urls = await deps.getPresignedURLs(allFileNames);
  } catch (err) {
    throw new Error(`yaylib: get presigned urls: ${err instanceof Error ? err.message : String(err)}`);
  }
  if (urls.length !== allJobs.length) {
    throw new Error(`yaylib: presigned urls count mismatch (want ${allJobs.length} got ${urls.length})`);
  }

  // Parallel PUT. The first failure aborts the still-running PUTs so they
  // exit promptly instead of finishing their now-pointless bodies.
  const controller = new AbortController();
  let firstErr: Error | undefined;
  await Promise.all(
    allJobs.map(async (job, i) => {
      const rawURL = urls[i].url ?? "";
      try {
        if (rawURL === "") {
          throw new Error(`yaylib: upload ${job.filename}: empty presigned url`);
        }
        await putToPresignedURL(deps, rawURL, job.body, job.body.byteLength, job.contentType, controller.signal);
      } catch (err) {
        if (!firstErr) {
          firstErr = err instanceof Error ? err : new Error(String(err));
          if (!firstErr.message.startsWith("yaylib:")) {
            firstErr = new Error(`yaylib: upload ${job.filename}: ${firstErr.message}`);
          }
          controller.abort();
        }
      }
    }),
  );
  if (firstErr) throw firstErr;

  // The server normalizes each filename (typically an "s3/" prefix
  // marking it as a freshly-uploaded object). That canonical value — not
  // the client-side path — is what editUser / createPost expect back.
  const out: string[] = [];
  for (let i = 0; i < files.length; i++) {
    out.push(urls[i].filename ?? "");
  }
  return out;
}

export async function uploadSingleImage(
  deps: UploadDeps,
  category: UploadCategory,
  file: Upload,
): Promise<string> {
  const names = await uploadImages(deps, category, [file]);
  return names[0];
}

async function sizedBody(body: VideoBody): Promise<{ body: Uint8Array | Blob; size: number }> {
  if (body instanceof Blob) return { body, size: body.size };
  if (body instanceof Uint8Array) return { body, size: body.byteLength };
  if (body instanceof ArrayBuffer) {
    const u = new Uint8Array(body);
    return { body: u, size: u.byteLength };
  }
  // ReadableStream — buffer fully so the PUT has a known Content-Length.
  const chunks: Uint8Array[] = [];
  const reader = body.getReader();
  for (;;) {
    const { done, value } = await reader.read();
    if (done) break;
    if (value) chunks.push(value);
  }
  let total = 0;
  for (const c of chunks) total += c.byteLength;
  const merged = new Uint8Array(total);
  let off = 0;
  for (const c of chunks) {
    merged.set(c, off);
    off += c.byteLength;
  }
  return { body: merged, size: total };
}

// uploadVideo uploads a single MP4 stream and returns the filename to pass
// back as video_file_name. Unlike images the video presigned endpoint
// does not return a canonical filename, so the random name the SDK
// generates is what callers hand to createPost / sendChatMessage
// (PORTING.md §8).
export async function uploadVideo(deps: UploadDeps, body: VideoBody | null | undefined): Promise<string> {
  if (body == null) throw new Error("yaylib: uploadVideo: body is nil");
  const { body: stream, size } = await sizedBody(body);
  const name = videoFilename(randomFilenamePrefix(16), Date.now());

  let rawURL: string;
  try {
    rawURL = await deps.getVideoPresignedURL(name);
  } catch (err) {
    throw new Error(`yaylib: get presigned url: ${err instanceof Error ? err.message : String(err)}`);
  }
  if (rawURL === "") throw new Error("yaylib: empty presigned url");

  const controller = new AbortController();
  try {
    await putToPresignedURL(deps, rawURL, stream, size, "video/mp4", controller.signal);
  } catch (err) {
    throw new Error(`yaylib: upload video: ${err instanceof Error ? err.message : String(err)}`);
  }
  return name;
}

// requireUserIDFor is exported so the Client's user-bound methods can
// enforce the "not logged in" contract before doing any work.
export { requireUserID };
