
// Code generated from the Yay! API schema; DO NOT EDIT.

package gen

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
	"strings"
)


// ModerationAPIService ModerationAPI service
type ModerationAPIService service

type ApiGetBannedWordsRequest struct {
	ctx context.Context
	ApiService *ModerationAPIService
	countryApiValue string
}

func (r ApiGetBannedWordsRequest) Execute() (*BanWordsResponse, *http.Response, error) {
	return r.ApiService.GetBannedWordsExecute(r)
}

func (r ApiGetBannedWordsRequest) ExecuteRaw() ([]byte, *http.Response, error) {
	_, httpResp, err := r.Execute()
	return executeRaw(httpResp, err)
}

/*
GetBannedWords Method for GetBannedWords

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param countryApiValue
 @return ApiGetBannedWordsRequest
*/
func (a *ModerationAPIService) GetBannedWords(ctx context.Context, countryApiValue string) ApiGetBannedWordsRequest {
	return ApiGetBannedWordsRequest{
		ApiService: a,
		ctx: ctx,
		countryApiValue: countryApiValue,
	}
}

// Execute executes the request
//  @return BanWordsResponse
func (a *ModerationAPIService) GetBannedWordsExecute(r ApiGetBannedWordsRequest) (*BanWordsResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *BanWordsResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ModerationAPIService.GetBannedWords")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/{countryApiValue}/api/v2/banned_words"
	localVarPath = strings.Replace(localVarPath, "{"+"countryApiValue"+"}", url.PathEscape(parameterValueToString(r.countryApiValue, "countryApiValue")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

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

type ApiGetPopularWordsRequest struct {
	ctx context.Context
	ApiService *ModerationAPIService
	countryApiValue string
	app string
}

func (r ApiGetPopularWordsRequest) Execute() (*PopularWordsResponse, *http.Response, error) {
	return r.ApiService.GetPopularWordsExecute(r)
}

func (r ApiGetPopularWordsRequest) ExecuteRaw() ([]byte, *http.Response, error) {
	_, httpResp, err := r.Execute()
	return executeRaw(httpResp, err)
}

/*
GetPopularWords Method for GetPopularWords

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param countryApiValue
 @param app
 @return ApiGetPopularWordsRequest
*/
func (a *ModerationAPIService) GetPopularWords(ctx context.Context, countryApiValue string, app string) ApiGetPopularWordsRequest {
	return ApiGetPopularWordsRequest{
		ApiService: a,
		ctx: ctx,
		countryApiValue: countryApiValue,
		app: app,
	}
}

// Execute executes the request
//  @return PopularWordsResponse
func (a *ModerationAPIService) GetPopularWordsExecute(r ApiGetPopularWordsRequest) (*PopularWordsResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *PopularWordsResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ModerationAPIService.GetPopularWords")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/{countryApiValue}/api/apps/{app}/popular_words"
	localVarPath = strings.Replace(localVarPath, "{"+"countryApiValue"+"}", url.PathEscape(parameterValueToString(r.countryApiValue, "countryApiValue")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"app"+"}", url.PathEscape(parameterValueToString(r.app, "app")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

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
