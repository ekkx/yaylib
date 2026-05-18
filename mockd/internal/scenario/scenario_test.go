package scenario

import (
	"sync"
	"testing"
	"time"
)

func TestParseValid(t *testing.T) {
	cases := []struct {
		in   string
		want Spec
	}{
		{"", Spec{Raw: "", Kind: KindNone}},
		{"fail-503-times-2", Spec{Raw: "fail-503-times-2", Kind: KindFailTimes, Status: 503, Times: 2}},
		{"fail-401-times-1", Spec{Raw: "fail-401-times-1", Kind: KindFailTimes, Status: 401, Times: 1}},
		{"delay-ms-250", Spec{Raw: "delay-ms-250", Kind: KindDelayMS, Delay: 250 * time.Millisecond}},
		{"retry-after-30", Spec{Raw: "retry-after-30", Kind: KindRetryAfter, RetrySecs: 30}},
		{"expired-token", Spec{Raw: "expired-token", Kind: KindExpiredToken}},
		{"ratelimit-quota-zero", Spec{Raw: "ratelimit-quota-zero", Kind: KindRatelimitQuotaZero}},
		{"body-malformed", Spec{Raw: "body-malformed", Kind: KindBodyMalformed}},
		{"body-unknown-enum", Spec{Raw: "body-unknown-enum", Kind: KindBodyUnknownEnum}},
	}
	for _, c := range cases {
		got, err := Parse(c.in)
		if err != nil {
			t.Fatalf("Parse(%q) unexpected error: %v", c.in, err)
		}
		if got != c.want {
			t.Errorf("Parse(%q) = %+v, want %+v", c.in, got, c.want)
		}
	}
}

func TestParseInvalid(t *testing.T) {
	for _, in := range []string{
		"fail-999-times-1", // status out of range
		"fail-503-times-0", // times must be >= 1
		"fail-503",         // incomplete
		"delay-ms-",        // missing n
		"retry-after-abc",  // non-numeric
		"bogus",            // unknown
		"expired-token-x",  // not an exact literal
	} {
		if _, err := Parse(in); err == nil {
			t.Errorf("Parse(%q) expected error, got nil", in)
		}
	}
}

func TestStoreNextIsPerSessionAndScenario(t *testing.T) {
	s := NewStore()
	if n := s.Next("sessA", "fail-503-times-2"); n != 1 {
		t.Fatalf("first Next = %d, want 1", n)
	}
	if n := s.Next("sessA", "fail-503-times-2"); n != 2 {
		t.Fatalf("second Next = %d, want 2", n)
	}
	// Different session does not share the counter.
	if n := s.Next("sessB", "fail-503-times-2"); n != 1 {
		t.Fatalf("other-session Next = %d, want 1", n)
	}
	// Different scenario under the same session is independent.
	if n := s.Next("sessA", "retry-after-5"); n != 1 {
		t.Fatalf("other-scenario Next = %d, want 1", n)
	}
}

func TestStoreRefreshLatch(t *testing.T) {
	s := NewStore()
	if s.IsRefreshed("x") {
		t.Fatal("fresh session should not be refreshed")
	}
	s.MarkRefreshed("x")
	if !s.IsRefreshed("x") {
		t.Fatal("session should be refreshed after MarkRefreshed")
	}
	if s.IsRefreshed("y") {
		t.Fatal("refresh latch must be per-session")
	}
	s.Reset()
	if s.IsRefreshed("x") {
		t.Fatal("Reset must clear the refresh latch")
	}
}

func TestStoreConcurrent(t *testing.T) {
	s := NewStore()
	const goroutines = 50
	const per = 100
	var wg sync.WaitGroup
	wg.Add(goroutines)
	for g := 0; g < goroutines; g++ {
		go func() {
			defer wg.Done()
			for i := 0; i < per; i++ {
				s.Next("shared", "fail-500-times-1")
			}
		}()
	}
	wg.Wait()
	// The next ordinal must equal total+1 with no lost updates.
	if n := s.Next("shared", "fail-500-times-1"); n != goroutines*per+1 {
		t.Fatalf("after concurrent increments Next = %d, want %d", n, goroutines*per+1)
	}
}
