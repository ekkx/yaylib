
// Code generated from the Yay! API schema; DO NOT EDIT.

package gen

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
)


// ActivitiesAPIService ActivitiesAPI service
type ActivitiesAPIService service

type ApiGetUserActivitiesRequest struct {
	ctx context.Context
	ApiService *ActivitiesAPIService
	fromTimestamp *int64
}

func (r ApiGetUserActivitiesRequest) FromTimestamp(fromTimestamp int64) ApiGetUserActivitiesRequest {
	r.fromTimestamp = &fromTimestamp
	return r
}

func (r ApiGetUserActivitiesRequest) Execute() (*ActivitiesResponse, *http.Response, error) {
	return r.ApiService.GetUserActivitiesExecute(r)
}

func (r ApiGetUserActivitiesRequest) ExecuteRaw() ([]byte, *http.Response, error) {
	_, httpResp, err := r.Execute()
	return executeRaw(httpResp, err)
}

/*
GetUserActivities Method for GetUserActivities

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiGetUserActivitiesRequest
*/
func (a *ActivitiesAPIService) GetUserActivities(ctx context.Context) ApiGetUserActivitiesRequest {
	return ApiGetUserActivitiesRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return ActivitiesResponse
func (a *ActivitiesAPIService) GetUserActivitiesExecute(r ApiGetUserActivitiesRequest) (*ActivitiesResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ActivitiesResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ActivitiesAPIService.GetUserActivities")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/user_activities"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.fromTimestamp != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "from_timestamp", r.fromTimestamp, "form", "")
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

type ApiGetUserActivitiesV1Request struct {
	ctx context.Context
	ApiService *ActivitiesAPIService
	important *bool
	fromTimestamp *int64
	number *int32
}

func (r ApiGetUserActivitiesV1Request) Important(important bool) ApiGetUserActivitiesV1Request {
	r.important = &important
	return r
}

func (r ApiGetUserActivitiesV1Request) FromTimestamp(fromTimestamp int64) ApiGetUserActivitiesV1Request {
	r.fromTimestamp = &fromTimestamp
	return r
}

func (r ApiGetUserActivitiesV1Request) Number(number int32) ApiGetUserActivitiesV1Request {
	r.number = &number
	return r
}

func (r ApiGetUserActivitiesV1Request) Execute() (*ActivitiesResponse, *http.Response, error) {
	return r.ApiService.GetUserActivitiesV1Execute(r)
}

func (r ApiGetUserActivitiesV1Request) ExecuteRaw() ([]byte, *http.Response, error) {
	_, httpResp, err := r.Execute()
	return executeRaw(httpResp, err)
}

/*
GetUserActivitiesV1 Method for GetUserActivitiesV1

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiGetUserActivitiesV1Request
*/
func (a *ActivitiesAPIService) GetUserActivitiesV1(ctx context.Context) ApiGetUserActivitiesV1Request {
	return ApiGetUserActivitiesV1Request{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return ActivitiesResponse
func (a *ActivitiesAPIService) GetUserActivitiesV1Execute(r ApiGetUserActivitiesV1Request) (*ActivitiesResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ActivitiesResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ActivitiesAPIService.GetUserActivitiesV1")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/user_activities"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.important != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "important", r.important, "form", "")
	}
	if r.fromTimestamp != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "from_timestamp", r.fromTimestamp, "form", "")
	}
	if r.number != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "number", r.number, "form", "")
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
