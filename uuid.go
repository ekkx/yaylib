package yaylib

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"time"
)

// NewUUIDv4 returns a random UUID v4 in canonical 8-4-4-4-12 hex form.
//
// This is used as the default DeviceUUID when NewClient is called without
// WithDeviceUUID. A UUID generated here is only stable for the lifetime of
// the process — persist it (e.g. via the session store) to keep device
// identity continuous across runs.
func NewUUIDv4() string {
	var b [16]byte
	if _, err := rand.Read(b[:]); err != nil {
		// crypto/rand should never fail on a healthy system; degrade to a
		// time-seeded value rather than panic, because the identifier is
		// not a security token.
		t := time.Now().UnixNano()
		for i := 0; i < 16; i++ {
			b[i] = byte(t >> (i * 4))
		}
	}
	b[6] = (b[6] & 0x0f) | 0x40 // version 4
	b[8] = (b[8] & 0x3f) | 0x80 // variant 10
	h := hex.EncodeToString(b[:])
	return fmt.Sprintf("%s-%s-%s-%s-%s", h[0:8], h[8:12], h[12:16], h[16:20], h[20:32])
}
