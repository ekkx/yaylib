package mock

import (
	"encoding/json"
	"net/http"
)

// commonError mirrors the upstream CommonErrorResponse wire schema
// (snake_case keys, every field nullable). Only the fields a scenario
// sets are emitted so SDK error decoders see a realistic body.
type commonError struct {
	ErrorCode      *int   `json:"error_code,omitempty"`
	Message        string `json:"message,omitempty"`
	BanUntil       *int64 `json:"ban_until,omitempty"`
	RetryIn        *int64 `json:"retry_in,omitempty"`
	RemainingQuota *int64 `json:"remaining_quota,omitempty"`
}

func writeJSON(w http.ResponseWriter, status int, body any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(body)
}

// writeError emits status with a CommonErrorResponse body.
func writeError(w http.ResponseWriter, status int, e commonError) {
	writeJSON(w, status, e)
}

func intp(v int) *int { return &v }

// tokenResponse mirrors the upstream TokenResponse schema. Values are
// fixed so login/refresh parity assertions are deterministic across
// languages; the access token differs from the refresh token and a
// refresh issues a distinct access token from a fresh login.
type tokenResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	ExpiresIn    int64  `json:"expires_in"`
	ID           int64  `json:"id"`
	CreatedAt    int64  `json:"created_at"`
}

const (
	freshAccessToken     = "mock-access-token"
	refreshedAccessToken = "mock-access-token-refreshed"
	mockRefreshToken     = "mock-refresh-token"
	mockUserID           = 1
	mockExpiresIn        = 3600
	mockCreatedAt        = 1700000000 // fixed epoch for deterministic bodies
)

// presignedURL mirrors the upstream PresignedUrl schema.
type presignedURL struct {
	Filename string `json:"filename"`
	URL      string `json:"url"`
}

type presignedURLsResponse struct {
	PresignedUrls []presignedURL `json:"presigned_urls"`
}

type presignedURLResponse struct {
	PresignedURL string `json:"presigned_url"`
}

type wsTokenResponse struct {
	Token string `json:"token"`
}

const mockWSToken = "mock-ws-token"

func loginToken() tokenResponse {
	return tokenResponse{
		AccessToken:  freshAccessToken,
		RefreshToken: mockRefreshToken,
		ExpiresIn:    mockExpiresIn,
		ID:           mockUserID,
		CreatedAt:    mockCreatedAt,
	}
}

func refreshToken() tokenResponse {
	t := loginToken()
	t.AccessToken = refreshedAccessToken
	return t
}
