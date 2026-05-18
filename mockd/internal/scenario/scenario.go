// Package scenario parses the X-Yay-Mock-Scenario header into a typed
// behavior directive and holds the per-(session, scenario) state that
// stateful scenarios (retry counters, refresh latch) need.
//
// The header grammar is a cross-language contract: the Go, TypeScript
// and Python test suites all emit these exact strings, so the accepted
// forms below must not change without bumping the contract in README.md.
//
// Grammar:
//
//	fail-<status>-times-<k>   first k requests answer <status> with a
//	                          CommonErrorResponse body, then fall through
//	delay-ms-<n>              sleep n ms, then fall through
//	retry-after-<n>           first request answers 429 with Retry-After:
//	                          <n> and body retry_in=<n>, then falls through
//	expired-token             protected requests answer 401 until the
//	                          token endpoint is hit (simulated refresh),
//	                          then fall through
//	ratelimit-quota-zero      every request answers 429 with
//	                          remaining_quota=0 (no fall-through)
//	body-malformed            200 with a JSON body that cannot decode
//	                          into any typed model (exercises the
//	                          typed-parse-fail / ExecuteRaw contract)
//	body-unknown-enum         200 with an object carrying an
//	                          out-of-spec enum value (exercises
//	                          tolerant enum decoding)
//	ok-status-<2xx>           a success answered with a non-200 2xx
//	                          (e.g. ok-status-201) + an empty object
//	                          body that decodes into any all-optional
//	                          success model (exercises "any 2xx is a
//	                          success")
//
// State is keyed by the X-Yay-Mock-Session header. Tests generate a
// fresh session id per case, so parallel runs across all three
// languages never share counters.
package scenario

import (
	"fmt"
	"regexp"
	"strconv"
	"sync"
	"time"
)

// Kind enumerates the supported scenario behaviors.
type Kind int

const (
	// KindNone means no scenario header (or an empty one) was sent.
	KindNone Kind = iota
	KindFailTimes
	KindDelayMS
	KindRetryAfter
	KindExpiredToken
	KindRatelimitQuotaZero
	KindBodyMalformed
	KindBodyUnknownEnum
	KindOKStatus
)

// Wire error codes mirrored from the upstream error catalog. They let
// SDK error-mapping assertions resolve to the same constants the real
// server would produce.
const (
	CodeAccessTokenExpired = -3
	CodeQuotaLimitExceeded = -343
)

// Spec is a parsed X-Yay-Mock-Scenario directive.
type Spec struct {
	Raw       string        // original header value (state key component)
	Kind      Kind          // behavior selector
	Status    int           // HTTP status for KindFailTimes
	Times     int           // request count for KindFailTimes
	Delay     time.Duration // sleep for KindDelayMS
	RetrySecs int           // seconds for KindRetryAfter (Retry-After + retry_in)
}

var (
	reFailTimes  = regexp.MustCompile(`^fail-(\d{3})-times-(\d+)$`)
	reDelayMS    = regexp.MustCompile(`^delay-ms-(\d+)$`)
	reRetryAfter = regexp.MustCompile(`^retry-after-(\d+)$`)
	reOKStatus   = regexp.MustCompile(`^ok-status-(2\d\d)$`)
)

// Parse turns a header value into a Spec. An empty value yields a
// KindNone Spec and no error. An unrecognized value is an error so the
// middleware can fail loud rather than silently behave as happy-path.
func Parse(raw string) (Spec, error) {
	s := Spec{Raw: raw, Kind: KindNone}
	if raw == "" {
		return s, nil
	}
	switch raw {
	case "expired-token":
		s.Kind = KindExpiredToken
		return s, nil
	case "ratelimit-quota-zero":
		s.Kind = KindRatelimitQuotaZero
		return s, nil
	case "body-malformed":
		s.Kind = KindBodyMalformed
		return s, nil
	case "body-unknown-enum":
		s.Kind = KindBodyUnknownEnum
		return s, nil
	}
	if m := reFailTimes.FindStringSubmatch(raw); m != nil {
		status, _ := strconv.Atoi(m[1])
		times, _ := strconv.Atoi(m[2])
		if status < 100 || status > 599 {
			return s, fmt.Errorf("scenario %q: status %d out of range", raw, status)
		}
		if times < 1 {
			return s, fmt.Errorf("scenario %q: times must be >= 1", raw)
		}
		s.Kind, s.Status, s.Times = KindFailTimes, status, times
		return s, nil
	}
	if m := reDelayMS.FindStringSubmatch(raw); m != nil {
		n, _ := strconv.Atoi(m[1])
		s.Kind, s.Delay = KindDelayMS, time.Duration(n)*time.Millisecond
		return s, nil
	}
	if m := reRetryAfter.FindStringSubmatch(raw); m != nil {
		n, _ := strconv.Atoi(m[1])
		s.Kind, s.RetrySecs = KindRetryAfter, n
		return s, nil
	}
	if m := reOKStatus.FindStringSubmatch(raw); m != nil {
		status, _ := strconv.Atoi(m[1])
		s.Kind, s.Status = KindOKStatus, status
		return s, nil
	}
	return s, fmt.Errorf("unrecognized scenario %q", raw)
}

// Store holds mutable per-session scenario state. The zero value is
// not usable; construct with NewStore.
type Store struct {
	mu        sync.Mutex
	counts    map[string]int  // (session\x00scenario) -> requests seen
	refreshed map[string]bool // session -> token refreshed at least once
}

// NewStore returns an empty Store.
func NewStore() *Store {
	return &Store{counts: map[string]int{}, refreshed: map[string]bool{}}
}

func counterKey(session, scenario string) string {
	return session + "\x00" + scenario
}

// Next records one request for (session, scenario) and returns its
// 1-based ordinal. The first request returns 1.
func (s *Store) Next(session, scenario string) int {
	s.mu.Lock()
	defer s.mu.Unlock()
	k := counterKey(session, scenario)
	s.counts[k]++
	return s.counts[k]
}

// MarkRefreshed latches that the session completed a token refresh.
func (s *Store) MarkRefreshed(session string) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.refreshed[session] = true
}

// IsRefreshed reports whether the session has completed a refresh.
func (s *Store) IsRefreshed(session string) bool {
	s.mu.Lock()
	defer s.mu.Unlock()
	return s.refreshed[session]
}

// Reset clears all state. Exposed for the /__mock/reset admin route so
// a suite can isolate cases without minting new session ids.
func (s *Store) Reset() {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.counts = map[string]int{}
	s.refreshed = map[string]bool{}
}
