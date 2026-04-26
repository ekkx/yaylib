
// Code generated from the Yay! API schema; DO NOT EDIT.

package gen

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
	"strings"
	"reflect"
)


// GamesAPIService GamesAPI service
type GamesAPIService service

type ApiListGameAppsRequest struct {
	ctx context.Context
	ApiService *GamesAPIService
	number *int32
	ids *[]int64
	fromId *int64
}

func (r ApiListGameAppsRequest) Number(number int32) ApiListGameAppsRequest {
	r.number = &number
	return r
}

func (r ApiListGameAppsRequest) Ids(ids []int64) ApiListGameAppsRequest {
	r.ids = &ids
	return r
}

func (r ApiListGameAppsRequest) FromId(fromId int64) ApiListGameAppsRequest {
	r.fromId = &fromId
	return r
}

func (r ApiListGameAppsRequest) Execute() (*GamesResponse, *http.Response, error) {
	return r.ApiService.ListGameAppsExecute(r)
}

/*
ListGameApps Method for ListGameApps

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiListGameAppsRequest
*/
func (a *GamesAPIService) ListGameApps(ctx context.Context) ApiListGameAppsRequest {
	return ApiListGameAppsRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return GamesResponse
func (a *GamesAPIService) ListGameAppsExecute(r ApiListGameAppsRequest) (*GamesResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *GamesResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "GamesAPIService.ListGameApps")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/games/apps"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.number == nil {
		return localVarReturnValue, nil, reportError("number is required and must be specified")
	}
	if r.ids == nil {
		return localVarReturnValue, nil, reportError("ids is required and must be specified")
	}

	parameterAddToHeaderOrQuery(localVarQueryParams, "number", r.number, "form", "")
	if r.fromId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "from_id", r.fromId, "form", "")
	}
	{
		t := *r.ids
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				parameterAddToHeaderOrQuery(localVarQueryParams, "ids[]", s.Index(i).Interface(), "form", "multi")
			}
		} else {
			parameterAddToHeaderOrQuery(localVarQueryParams, "ids[]", t, "form", "multi")
		}
	}
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

type ApiListGameWalkthroughsRequest struct {
	ctx context.Context
	ApiService *GamesAPIService
	appId int64
}

func (r ApiListGameWalkthroughsRequest) Execute() ([]Walkthrough, *http.Response, error) {
	return r.ApiService.ListGameWalkthroughsExecute(r)
}

/*
ListGameWalkthroughs Method for ListGameWalkthroughs

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param appId
 @return ApiListGameWalkthroughsRequest
*/
func (a *GamesAPIService) ListGameWalkthroughs(ctx context.Context, appId int64) ApiListGameWalkthroughsRequest {
	return ApiListGameWalkthroughsRequest{
		ApiService: a,
		ctx: ctx,
		appId: appId,
	}
}

// Execute executes the request
//  @return []Walkthrough
func (a *GamesAPIService) ListGameWalkthroughsExecute(r ApiListGameWalkthroughsRequest) ([]Walkthrough, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []Walkthrough
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "GamesAPIService.ListGameWalkthroughs")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/games/apps/{app_id}/walkthroughs"
	localVarPath = strings.Replace(localVarPath, "{"+"app_id"+"}", url.PathEscape(parameterValueToString(r.appId, "appId")), -1)

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
