package yaylib

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
)

// NewUUIDv4 returns a random UUID v4 in canonical 8-4-4-4-12 hex form.
//
// This is used as the default DeviceUUID when NewClient is called without
// WithDeviceUUID. A UUID generated here is only stable for the lifetime of
// the process — persist it (e.g. via the session store) to keep device
// identity continuous across runs.
//
// Panics if crypto/rand cannot return entropy. On a healthy host this
// does not happen; the panic surfaces an unrecoverable environment
// failure rather than silently producing weak, collision-prone IDs.
func NewUUIDv4() string {
	var b [16]byte
	if _, err := rand.Read(b[:]); err != nil {
		panic(fmt.Sprintf("yaylib: crypto/rand: %v", err))
	}
	b[6] = (b[6] & 0x0f) | 0x40 // version 4
	b[8] = (b[8] & 0x3f) | 0x80 // variant 10
	h := hex.EncodeToString(b[:])
	return fmt.Sprintf("%s-%s-%s-%s-%s", h[0:8], h[8:12], h[12:16], h[16:20], h[20:32])
}
