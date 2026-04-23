# yaylib

Unofficial TypeScript client for the [Yay!](https://yay.space/) social
networking API.

> **Early preview (`v0.x`).** The full typed REST client is not shipped
> yet — this release only exposes a handful of primitives. The API surface
> will change without notice until `v1.0.0`. Pin an exact version if you
> depend on it.

## Install

```sh
npm install yaylib
```

## What's in `v0.x`

```ts
import { API_BASE_URL, WS_URL, createDeviceUuid } from "yaylib";

API_BASE_URL;       // "https://api.yay.space"
WS_URL;             // "wss://cable.yay.space"
createDeviceUuid(); // e.g. "0f9f7d3a-4a7e-4c1e-bf5b-2a0f9b2b1c3d"
```

`createDeviceUuid()` returns a v4 UUID suitable for the `X-Device-UUID`
header that Yay! requires on every request. Node.js >= 20 is required.

## Roadmap

- [ ] Add typed REST client
- [ ] Session / token management with refresh handling
- [ ] Real-time WebSocket (ActionCable) support

## License

MIT
