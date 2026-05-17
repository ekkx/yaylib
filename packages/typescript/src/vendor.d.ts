// Ambient declarations for the pure-JS image codecs that ship no types.
// Only the slices of each API the upload pipeline uses are declared.

declare module "jpeg-js" {
  export interface RawImageData {
    width: number;
    height: number;
    data: Uint8Array;
  }
  export function decode(
    data: Uint8Array | ArrayBuffer | Buffer,
    opts?: { useTArray?: boolean; formatAsRGBA?: boolean },
  ): RawImageData;
  export function encode(
    img: { width: number; height: number; data: Uint8Array | Buffer },
    quality?: number,
  ): { width: number; height: number; data: Uint8Array };
}

declare module "omggif" {
  export class GifReader {
    constructor(buf: Uint8Array);
    width: number;
    height: number;
    numFrames(): number;
    loopCount(): number;
    frameInfo(frame: number): {
      x: number;
      y: number;
      width: number;
      height: number;
      delay: number;
      disposal: number;
      transparent_index: number | null;
    };
    decodeAndBlitFrameRGBA(frame: number, pixels: Uint8Array | Uint8ClampedArray): void;
  }
  export class GifWriter {
    constructor(
      buf: Uint8Array | number[],
      width: number,
      height: number,
      gopts?: { loop?: number; palette?: number[] },
    );
    addFrame(
      x: number,
      y: number,
      width: number,
      height: number,
      indexedPixels: Uint8Array | number[],
      opts?: { palette?: number[]; delay?: number; disposal?: number; transparent?: number },
    ): number;
    end(): number;
  }
}
