// Package mock wires the scenario engine in front of the typed
// router. Requests carrying X-Yay-Mock-Scenario are fully owned here
// and never reach the typed router; everything else falls through
// to it (with the token endpoint owned so auth parity works without a
// scenario).
package mock

import (
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/ekkx/yaylib/mockd/internal/scenario"
)

const (
	headerScenario = "X-Yay-Mock-Scenario"
	headerSession  = "X-Yay-Mock-Session"

	tokenPath       = "/api/v1/oauth/token"
	wsTokenPath     = "/v1/users/ws_token"
	bucketPresigned = "/v1/buckets/presigned_urls"
	userPresigned   = "/v1/users/presigned_url"
	s3Prefix        = "/__mock/s3/"
)

// Server is the top-level mock http.Handler.
type Server struct {
	store *scenario.Store
	oas   http.Handler
	ws    *wsHub
}

// New constructs the mock server. The typed router is built once
// and reused.
func New() (*Server, error) {
	oasSrv, err := newOASServer()
	if err != nil {
		return nil, err
	}
	return &Server{store: scenario.NewStore(), oas: oasSrv, ws: newWSHub()}, nil
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// The event-stream base URL is configurable, so a WS dial can
	// land on any path; detect it by the upgrade headers, ahead of
	// everything else.
	if isWebSocketUpgrade(r) {
		s.serveWS(w, r)
		return
	}

	switch r.URL.Path {
	case "/__mock/health":
		diag(w, "", "admin-health")
		writeJSON(w, http.StatusOK, map[string]string{"status": "ok"})
		return
	case "/__mock/reset":
		s.store.Reset()
		diag(w, "", "admin-reset")
		w.WriteHeader(http.StatusNoContent)
		return
	case "/__mock/ws/close":
		s.ws.closeAll()
		diag(w, "", "admin-ws-close")
		w.WriteHeader(http.StatusNoContent)
		return
	}

	// The S3 receiver accepts the presigned PUT. The dial-equivalent
	// contract holds here too: a presigned PUT carries no bearer.
	if strings.HasPrefix(r.URL.Path, s3Prefix) {
		if r.Header.Get("Authorization") != "" {
			diag(w, "", "s3-put-rejected")
			writeError(w, http.StatusForbidden, commonError{
				Message: "presigned PUT must not carry Authorization",
			})
			return
		}
		diag(w, "", "s3-put")
		w.WriteHeader(http.StatusOK)
		return
	}

	session := r.Header.Get(headerSession)
	rawScenario := r.Header.Get(headerScenario)
	spec, err := scenario.Parse(rawScenario)
	if err != nil {
		diag(w, rawScenario, "bad-scenario")
		writeError(w, http.StatusBadRequest, commonError{
			Message: "invalid " + headerScenario + ": " + err.Error(),
		})
		return
	}

	switch spec.Kind {
	case scenario.KindNone:
		diag(w, spec.Raw, "happy")
		s.fallthroughHappy(w, r)

	case scenario.KindFailTimes:
		n := s.store.Next(session, spec.Raw)
		diagAttempt(w, n)
		if n <= spec.Times {
			diag(w, spec.Raw, "fail")
			writeError(w, spec.Status, failBody(spec.Status, spec.Raw))
			return
		}
		diag(w, spec.Raw, "fail-exhausted-happy")
		s.fallthroughHappy(w, r)

	case scenario.KindDelayMS:
		time.Sleep(spec.Delay)
		diag(w, spec.Raw, "delay-happy")
		s.fallthroughHappy(w, r)

	case scenario.KindRetryAfter:
		n := s.store.Next(session, spec.Raw)
		diagAttempt(w, n)
		if n == 1 {
			secs := int64(spec.RetrySecs)
			w.Header().Set("Retry-After", strconv.Itoa(spec.RetrySecs))
			diag(w, spec.Raw, "retry-after-429")
			writeError(w, http.StatusTooManyRequests, commonError{
				Message: "rate limited; retry after backoff",
				RetryIn: &secs,
			})
			return
		}
		diag(w, spec.Raw, "retry-after-happy")
		s.fallthroughHappy(w, r)

	case scenario.KindExpiredToken:
		if isTokenEndpoint(r) {
			// The SDK's 401 handler reaches here to refresh.
			s.store.MarkRefreshed(session)
			diag(w, spec.Raw, "expired-token-refresh")
			writeJSON(w, http.StatusOK, refreshToken())
			return
		}
		if !s.store.IsRefreshed(session) {
			diag(w, spec.Raw, "expired-token-401")
			writeError(w, http.StatusUnauthorized, commonError{
				ErrorCode: intp(scenario.CodeAccessTokenExpired),
				Message:   "access token expired",
			})
			return
		}
		diag(w, spec.Raw, "expired-token-happy")
		s.fallthroughHappy(w, r)

	case scenario.KindRatelimitQuotaZero:
		zero := int64(0)
		diag(w, spec.Raw, "ratelimit-429")
		writeError(w, http.StatusTooManyRequests, commonError{
			ErrorCode:      intp(scenario.CodeQuotaLimitExceeded),
			Message:        "quota exhausted",
			RemainingQuota: &zero,
		})

	case scenario.KindBodyMalformed:
		// 200 + a JSON scalar: it cannot decode into any typed
		// response struct, so the SDK's typed Execute fails while
		// ExecuteRaw returns these bytes with no error.
		diag(w, spec.Raw, "body-malformed")
		writeJSON(w, http.StatusOK, "mock-malformed-body")

	case scenario.KindBodyUnknownEnum:
		// 200 + an object whose enum-typed wire field carries an
		// out-of-spec value. The tolerant decoder accepts it (field
		// set to the unknown string, IsValid() false). The parity
		// test points this at an endpoint whose typed model has the
		// post_type enum.
		diag(w, spec.Raw, "body-unknown-enum")
		writeJSON(w, http.StatusOK, map[string]string{"post_type": unknownEnumValue})

	case scenario.KindOKStatus:
		// A successful response answered with a non-200 2xx (login is
		// 201). The empty object decodes into any all-optional success
		// model, so the SDK must treat the 2xx as success — not key on
		// the exact documented status and discard the body.
		diag(w, spec.Raw, "ok-status")
		writeJSON(w, spec.Status, map[string]any{})
	}
}

const unknownEnumValue = "__mock_unknown_enum__"

// Diagnostic response headers. Every HTTP response the mock owns
// carries which scenario was seen, the (session, scenario) attempt
// ordinal for the stateful scenarios, and which decision branch ran —
// so a cross-language parity failure can be diagnosed from the
// response alone. These are observational only; SDKs ignore them.
const (
	headerDiagScenario = "X-Mock-Scenario"
	headerDiagAttempt  = "X-Mock-Attempt"
	headerDiagBranch   = "X-Mock-Branch"
)

func diag(w http.ResponseWriter, scenarioRaw, branch string) {
	if scenarioRaw != "" {
		w.Header().Set(headerDiagScenario, scenarioRaw)
	}
	w.Header().Set(headerDiagBranch, branch)
}

func diagAttempt(w http.ResponseWriter, n int) {
	w.Header().Set(headerDiagAttempt, strconv.Itoa(n))
}

// fallthroughHappy serves the success path. A few endpoints are owned
// here because parity tests need a real body shape: the token endpoint
// (auth), the WS token, and the presigned-URL endpoints (upload). The
// presigned URLs point back at this server's S3 receiver so the SDK's
// PUT stays in-process. Everything else goes to the typed router.
func (s *Server) fallthroughHappy(w http.ResponseWriter, r *http.Request) {
	switch {
	case isTokenEndpoint(r):
		writeJSON(w, http.StatusOK, loginToken())
	case r.Method == http.MethodGet && r.URL.Path == wsTokenPath:
		writeJSON(w, http.StatusOK, wsTokenResponse{Token: mockWSToken})
	case r.Method == http.MethodGet && r.URL.Path == bucketPresigned:
		s.serveBucketPresigned(w, r)
	case r.Method == http.MethodGet && r.URL.Path == userPresigned:
		s.serveUserPresigned(w, r)
	default:
		s.oas.ServeHTTP(w, r)
	}
}

// serveBucketPresigned answers GetBucketPresignedUrls. Each requested
// file name yields a URL on this server's S3 receiver and the
// server-canonical filename (the "s3/" prefix the real server adds —
// the SDK threads this canonical value into the create/edit call).
func (s *Server) serveBucketPresigned(w http.ResponseWriter, r *http.Request) {
	names := r.URL.Query()["file_names[]"]
	resp := presignedURLsResponse{PresignedUrls: make([]presignedURL, 0, len(names))}
	for _, n := range names {
		resp.PresignedUrls = append(resp.PresignedUrls, presignedURL{
			Filename: "s3/" + n,
			URL:      baseURL(r) + s3Prefix + n,
		})
	}
	writeJSON(w, http.StatusOK, resp)
}

// serveUserPresigned answers GetUserPresignedUrl (video). Video names
// are flat and carry no canonical prefix.
func (s *Server) serveUserPresigned(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("video_file_name")
	writeJSON(w, http.StatusOK, presignedURLResponse{
		PresignedURL: baseURL(r) + s3Prefix + name,
	})
}

func isTokenEndpoint(r *http.Request) bool {
	return r.Method == http.MethodPost && r.URL.Path == tokenPath
}

func baseURL(r *http.Request) string {
	scheme := "http"
	if r.TLS != nil {
		scheme = "https"
	}
	return scheme + "://" + r.Host
}

// failBody builds the CommonErrorResponse for a fail-<status>-times-<k>
// hit. A 401 carries the access-token-expired code so SDK error
// mapping resolves the same constant the real server would.
func failBody(status int, raw string) commonError {
	e := commonError{Message: "mock scenario " + raw}
	if status == http.StatusUnauthorized {
		e.ErrorCode = intp(scenario.CodeAccessTokenExpired)
	}
	return e
}
