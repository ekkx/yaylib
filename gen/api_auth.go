
// Code generated from the Yay! API schema; DO NOT EDIT.

package gen

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
)


// AuthAPIService AuthAPI service
type AuthAPIService service

type ApiOauthTokenRequest struct {
	ctx context.Context
	ApiService *AuthAPIService
	email *string
	grantType *string
	password *string
	refreshToken *string
}

func (r ApiOauthTokenRequest) Email(email string) ApiOauthTokenRequest {
	r.email = &email
	return r
}

func (r ApiOauthTokenRequest) GrantType(grantType string) ApiOauthTokenRequest {
	r.grantType = &grantType
	return r
}

func (r ApiOauthTokenRequest) Password(password string) ApiOauthTokenRequest {
	r.password = &password
	return r
}

func (r ApiOauthTokenRequest) RefreshToken(refreshToken string) ApiOauthTokenRequest {
	r.refreshToken = &refreshToken
	return r
}

func (r ApiOauthTokenRequest) Execute() (*TokenResponse, *http.Response, error) {
	return r.ApiService.OauthTokenExecute(r)
}

func (r ApiOauthTokenRequest) ExecuteRaw() ([]byte, *http.Response, error) {
	_, httpResp, err := r.Execute()
	return executeRaw(httpResp, err)
}

/*
OauthToken Method for OauthToken

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiOauthTokenRequest
*/
func (a *AuthAPIService) OauthToken(ctx context.Context) ApiOauthTokenRequest {
	return ApiOauthTokenRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return TokenResponse
func (a *AuthAPIService) OauthTokenExecute(r ApiOauthTokenRequest) (*TokenResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *TokenResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AuthAPIService.OauthToken")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/oauth/token"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/x-www-form-urlencoded"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.email != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "email", r.email, "", "")
	}
	if r.grantType != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "grant_type", r.grantType, "", "")
	}
	if r.password != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "password", r.password, "", "")
	}
	if r.refreshToken != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "refresh_token", r.refreshToken, "", "")
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}
