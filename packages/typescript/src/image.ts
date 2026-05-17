// PORTING.md §8: the upload protocol pieces that aren't visible from the
// category table — filename shaping, source-size probing, and the
// thumbnail rules. The Yay! server pattern-matches the filenames and
// DELETES the main image if its thumbnail is missing, statically encoded,
// or extension-mismatched, so this module mirrors upload.go's behaviour
// exactly: JPEG/PNG sources produce a shrunk JPEG thumbnail, animated GIF
// sources stay animated through the resize.

import { decode as jpegDecode, encode as jpegEncode } from "jpeg-js";
import { GifReader, GifWriter } from "omggif";
import { PNG } from "pngjs";
import { imageSize } from "image-size";

// MaxImagesPerUpload is the default per-call cap on multi-image uploads
// (uploadPostImages, uploadChatMessageImages). uploadReportImages is
// tighter — the server rejects the 5th image.
export const MAX_IMAGES_PER_UPLOAD = 9;

// MaxReportImagesPerUpload is the per-call cap on uploadReportImages.
export const MAX_REPORT_IMAGES_PER_UPLOAD = 4;

// extOf returns the lowercased extension including the leading dot, or ""
// when the filename has no extension. Mirrors Go's filepath.Ext.
export function extOf(filename: string): string {
  const base = filename.slice(filename.lastIndexOf("/") + 1);
  const dot = base.lastIndexOf(".");
  if (dot <= 0) return "";
  return base.slice(dot).toLowerCase();
}

// contentTypeFor maps a filename extension to the Content-Type sent on the
// presigned PUT. Unknown extensions fall back to image/jpeg.
export function contentTypeFor(filename: string): string {
  switch (extOf(filename)) {
    case ".png":
      return "image/png";
    case ".gif":
      return "image/gif";
    case ".mp4":
      return "video/mp4";
    default:
      return "image/jpeg";
  }
}

// normalizeImageExt collapses any input filename to one of .jpg, .jpeg,
// .png, .gif. Anything else (no extension, .heic, odd casings) becomes
// .jpg, matching upload.go.
export function normalizeImageExt(filename: string): string {
  switch (extOf(filename)) {
    case ".png":
      return ".png";
    case ".gif":
      return ".gif";
    case ".jpeg":
      return ".jpeg";
    case ".jpg":
      return ".jpg";
    default:
      return ".jpg";
  }
}

function sizeSuffix(width: number, height: number): string {
  return width > 0 && height > 0 ? `_size_${width}x${height}` : "";
}

// imageFilename builds the main-image object key the Yay! server
// pattern-matches. Empty ext collapses to .jpg; the size suffix is
// omitted when either dimension is non-positive.
export function imageFilename(
  categoryPath: string,
  prefix: string,
  unixMillis: number,
  index: number,
  ext: string,
  width: number,
  height: number,
): string {
  const e = ext === "" ? ".jpg" : ext;
  const body = `${prefix}_${unixMillis}_${index}${sizeSuffix(width, height)}${e}`;
  return categoryPath === "" ? body : `${categoryPath}/${body}`;
}

// thumbnailFilename mirrors imageFilename but adds the "thumb_" prefix the
// server expects for the resized companion. The size suffix carries the
// *original* image dimensions, not the thumbnail's.
export function thumbnailFilename(
  categoryPath: string,
  prefix: string,
  unixMillis: number,
  index: number,
  ext: string,
  origWidth: number,
  origHeight: number,
): string {
  const e = ext === "" ? ".jpg" : ext;
  const body = `thumb_${prefix}_${unixMillis}_${index}${sizeSuffix(origWidth, origHeight)}${e}`;
  return categoryPath === "" ? body : `${categoryPath}/${body}`;
}

// videoFilename is the flat key videos use — no category path, no size
// suffix (PORTING.md §8).
export function videoFilename(prefix: string, unixMillis: number): string {
  return `${prefix}_${unixMillis}.mp4`;
}

const PREFIX_ALPHABET =
  "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ";

// randomFilenamePrefix returns n characters from [0-9A-Za-z], freshly
// generated per upload call. A missing crypto source throws rather than
// fabricating a fixed prefix that would guarantee bucket collisions
// (upload.go panics for the same reason).
export function randomFilenamePrefix(n: number): string {
  const c = globalThis.crypto;
  if (!c || typeof c.getRandomValues !== "function") {
    throw new Error("yaylib: crypto.getRandomValues unavailable; refusing to fabricate upload filename");
  }
  const buf = new Uint8Array(n);
  c.getRandomValues(buf);
  let out = "";
  for (let i = 0; i < n; i++) {
    out += PREFIX_ALPHABET[buf[i] % PREFIX_ALPHABET.length];
  }
  return out;
}

// formatYMD matches the "YYYY/M/D" path the server expects in
// date-bucketed categories — note the un-padded month and day. UTC fields
// are used so the bucket path is deterministic regardless of host
// timezone (PORTING.md §8: the server treats the path as a hint).
export function formatYMD(d: Date): string {
  return `${d.getUTCFullYear()}/${d.getUTCMonth() + 1}/${d.getUTCDate()}`;
}

// fitInside returns the width/height that fit srcW x srcH inside
// maxW x maxH while preserving aspect ratio. Sources already smaller than
// the box are returned unchanged. [0, 0] signals invalid input.
export function fitInside(
  srcW: number,
  srcH: number,
  maxW: number,
  maxH: number,
): [number, number] {
  if (srcW <= 0 || srcH <= 0) return [0, 0];
  let scale = 1;
  const sw = maxW / srcW;
  if (sw < scale) scale = sw;
  const sh = maxH / srcH;
  if (sh < scale) scale = sh;
  let dstW = Math.floor(srcW * scale);
  let dstH = Math.floor(srcH * scale);
  if (dstW < 1) dstW = 1;
  if (dstH < 1) dstH = 1;
  return [dstW, dstH];
}

// probeImageSize reads the native width/height of a JPEG, PNG, or GIF from
// its header bytes. An undecodable body throws — uploading a body whose
// size suffix is wrong (or whose required thumbnail can't be built) just
// triggers server-side moderation deletion, so we surface the failure.
export function probeImageSize(body: Uint8Array): { width: number; height: number } {
  let dim: { width?: number; height?: number };
  try {
    dim = imageSize(body);
  } catch (err) {
    throw new Error(`decode image: ${err instanceof Error ? err.message : String(err)}`);
  }
  if (!dim || !dim.width || !dim.height) {
    throw new Error("decode image: could not read dimensions");
  }
  return { width: dim.width, height: dim.height };
}

interface RGBA {
  width: number;
  height: number;
  data: Uint8Array; // RGBA, length = width*height*4
}

function decodeToRGBA(body: Uint8Array): RGBA {
  let type: string | undefined;
  try {
    type = imageSize(body).type;
  } catch {
    type = undefined;
  }
  if (type === "png") {
    const png = PNG.sync.read(Buffer.from(body));
    return { width: png.width, height: png.height, data: new Uint8Array(png.data) };
  }
  // jpg / anything jpeg-js can read.
  const img = jpegDecode(body, { useTArray: true, formatAsRGBA: true });
  return { width: img.width, height: img.height, data: new Uint8Array(img.data) };
}

// resizeBilinear scales an RGBA buffer to dstW x dstH with bilinear
// sampling. Exact pixel parity with Go's draw.BiLinear is not required;
// the server cares about the dimensions and that a thumbnail exists.
function resizeBilinear(src: RGBA, dstW: number, dstH: number): RGBA {
  if (dstW === src.width && dstH === src.height) {
    return { width: dstW, height: dstH, data: new Uint8Array(src.data) };
  }
  const out = new Uint8Array(dstW * dstH * 4);
  const xRatio = src.width > 1 ? (src.width - 1) / Math.max(dstW - 1, 1) : 0;
  const yRatio = src.height > 1 ? (src.height - 1) / Math.max(dstH - 1, 1) : 0;
  for (let y = 0; y < dstH; y++) {
    const sy = y * yRatio;
    const y0 = Math.floor(sy);
    const y1 = Math.min(y0 + 1, src.height - 1);
    const wy = sy - y0;
    for (let x = 0; x < dstW; x++) {
      const sx = x * xRatio;
      const x0 = Math.floor(sx);
      const x1 = Math.min(x0 + 1, src.width - 1);
      const wx = sx - x0;
      const i00 = (y0 * src.width + x0) * 4;
      const i01 = (y0 * src.width + x1) * 4;
      const i10 = (y1 * src.width + x0) * 4;
      const i11 = (y1 * src.width + x1) * 4;
      const o = (y * dstW + x) * 4;
      for (let c = 0; c < 4; c++) {
        const top = src.data[i00 + c] * (1 - wx) + src.data[i01 + c] * wx;
        const bot = src.data[i10 + c] * (1 - wx) + src.data[i11 + c] * wx;
        out[o + c] = Math.round(top * (1 - wy) + bot * wy);
      }
    }
  }
  return { width: dstW, height: dstH, data: out };
}

export interface Thumbnail {
  body: Uint8Array;
  ext: string;
  contentType: string;
}

// makeThumbnailFor produces the right kind of thumbnail for the source
// format. GIF inputs stay animated GIFs (so the avatar/icon keeps moving
// when the server renders the thumbnail); everything else becomes a JPEG.
// The returned ext/contentType match the produced bytes, mirroring
// upload.go's makeThumbnailFor.
export function makeThumbnailFor(
  body: Uint8Array,
  sourceExt: string,
  maxW: number,
  maxH: number,
): Thumbnail {
  if (sourceExt === ".gif") {
    return { body: makeGifThumbnail(body, maxW, maxH), ext: ".gif", contentType: "image/gif" };
  }
  return { body: makeJpegThumbnail(body, maxW, maxH), ext: ".jpeg", contentType: "image/jpeg" };
}

// makeJpegThumbnail decodes a JPEG or PNG and re-encodes it as a shrunk
// JPEG that fits inside maxW x maxH while preserving aspect ratio. Sources
// already inside the box are returned at their original size — upscaling
// would only blur the result.
export function makeJpegThumbnail(body: Uint8Array, maxW: number, maxH: number): Uint8Array {
  if (maxW <= 0 || maxH <= 0) {
    throw new Error(`make thumbnail: invalid target size ${maxW}x${maxH}`);
  }
  let src: RGBA;
  try {
    src = decodeToRGBA(body);
  } catch (err) {
    throw new Error(`make thumbnail: decode source image: ${err instanceof Error ? err.message : String(err)}`);
  }
  const [dstW, dstH] = fitInside(src.width, src.height, maxW, maxH);
  if (dstW === 0) throw new Error("make thumbnail: empty source image");
  const scaled = resizeBilinear(src, dstW, dstH);
  const encoded = jpegEncode({ width: scaled.width, height: scaled.height, data: scaled.data }, 85);
  return new Uint8Array(encoded.data);
}

// quantize maps an RGBA buffer to an indexed palette using a popularity
// palette plus Floyd–Steinberg error diffusion. The palette is whatever
// fits in <=256 colours; GIF requires a power-of-two palette so it is
// padded up. Exact dither parity with Go is not required — the contract
// is "stays an animated GIF of the right size", which the GifWriter
// reconstruction below satisfies.
function quantize(
  rgba: RGBA,
): { indices: Uint8Array; palette: number[] } {
  // Build a popularity histogram on a 4-bit-per-channel grid so the
  // palette stays small and deterministic.
  const counts = new Map<number, number>();
  const px = rgba.data;
  for (let i = 0; i < px.length; i += 4) {
    const key =
      ((px[i] >> 4) << 8) | ((px[i + 1] >> 4) << 4) | (px[i + 2] >> 4);
    counts.set(key, (counts.get(key) ?? 0) + 1);
  }
  const ranked = [...counts.entries()]
    .sort((a, b) => b[1] - a[1])
    .slice(0, 256)
    .map(([key]) => {
      const r = ((key >> 8) & 0xf) * 17;
      const g = ((key >> 4) & 0xf) * 17;
      const b = (key & 0xf) * 17;
      return [r, g, b] as [number, number, number];
    });
  if (ranked.length === 0) ranked.push([0, 0, 0]);
  // Pad to a power of two (GifWriter requirement).
  let palSize = 2;
  while (palSize < ranked.length) palSize <<= 1;
  while (ranked.length < palSize) ranked.push([0, 0, 0]);
  const palette = ranked.map(([r, g, b]) => (r << 16) | (g << 8) | b);

  const nearest = (r: number, g: number, b: number): number => {
    let best = 0;
    let bestD = Infinity;
    for (let i = 0; i < ranked.length; i++) {
      const dr = r - ranked[i][0];
      const dg = g - ranked[i][1];
      const db = b - ranked[i][2];
      const d = dr * dr + dg * dg + db * db;
      if (d < bestD) {
        bestD = d;
        best = i;
      }
    }
    return best;
  };

  const w = rgba.width;
  const h = rgba.height;
  const buf = Float32Array.from(px);
  const indices = new Uint8Array(w * h);
  for (let y = 0; y < h; y++) {
    for (let x = 0; x < w; x++) {
      const o = (y * w + x) * 4;
      const r = Math.max(0, Math.min(255, buf[o]));
      const g = Math.max(0, Math.min(255, buf[o + 1]));
      const b = Math.max(0, Math.min(255, buf[o + 2]));
      const idx = nearest(r, g, b);
      indices[y * w + x] = idx;
      const er = r - ranked[idx][0];
      const eg = g - ranked[idx][1];
      const eb = b - ranked[idx][2];
      const spread = (dx: number, dy: number, f: number) => {
        const nx = x + dx;
        const ny = y + dy;
        if (nx < 0 || nx >= w || ny < 0 || ny >= h) return;
        const no = (ny * w + nx) * 4;
        buf[no] += er * f;
        buf[no + 1] += eg * f;
        buf[no + 2] += eb * f;
      };
      spread(1, 0, 7 / 16);
      spread(-1, 1, 3 / 16);
      spread(0, 1, 5 / 16);
      spread(1, 1, 1 / 16);
    }
  }
  return { indices, palette };
}

// makeGifThumbnail decodes an animated GIF, scales every frame down to fit
// inside maxW x maxH (preserving aspect ratio), and re-encodes the result
// as a GIF that retains the original animation. If the source already
// fits inside the box the original bytes are returned untouched, matching
// upload.go.
export function makeGifThumbnail(body: Uint8Array, maxW: number, maxH: number): Uint8Array {
  if (maxW <= 0 || maxH <= 0) {
    throw new Error(`make thumbnail: invalid target size ${maxW}x${maxH}`);
  }
  let reader: GifReader;
  try {
    reader = new GifReader(body);
  } catch (err) {
    throw new Error(`make thumbnail: decode gif: ${err instanceof Error ? err.message : String(err)}`);
  }
  const srcW = reader.width;
  const srcH = reader.height;
  const [dstW, dstH] = fitInside(srcW, srcH, maxW, maxH);
  if (dstW === 0) throw new Error("make thumbnail: empty gif");
  if (dstW === srcW && dstH === srcH) {
    // Already fits — return the original bytes untouched.
    return body;
  }

  const frameCount = reader.numFrames();
  // Over-allocate the output; GifWriter writes into a fixed buffer.
  const out = new Uint8Array(dstW * dstH * frameCount * 2 + 1024 + frameCount * 256 * 3);
  let loop = 0;
  try {
    loop = reader.loopCount();
  } catch {
    loop = 0;
  }
  const writer = new GifWriter(out, dstW, dstH, { loop: loop >= 0 ? loop : 0 });

  // omggif's decodeAndBlitFrameRGBA composites each frame onto the running
  // canvas honouring disposal, so we always get a full-canvas RGBA frame —
  // the same invariant upload.go establishes by drawing onto a full RGBA
  // canvas before resizing.
  const canvas = new Uint8Array(srcW * srcH * 4);
  for (let i = 0; i < frameCount; i++) {
    const info = reader.frameInfo(i);
    reader.decodeAndBlitFrameRGBA(i, canvas);
    const scaled = resizeBilinear(
      { width: srcW, height: srcH, data: canvas },
      dstW,
      dstH,
    );
    const { indices, palette } = quantize(scaled);
    writer.addFrame(0, 0, dstW, dstH, indices, {
      palette,
      delay: info.delay,
      disposal: 0,
    });
  }
  const len = writer.end();
  return out.slice(0, len);
}
