
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


// GroupMuteAPIService GroupMuteAPI service
type GroupMuteAPIService service

type ApiListMutedGroupUsersRequest struct {
	ctx context.Context
	ApiService *GroupMuteAPIService
	id int64
	cursor *string
	comArthenicaFfmpegkitMediaInformationKEYSIZE *int32
	keyword *string
}

func (r ApiListMutedGroupUsersRequest) Cursor(cursor string) ApiListMutedGroupUsersRequest {
	r.cursor = &cursor
	return r
}

func (r ApiListMutedGroupUsersRequest) ComArthenicaFfmpegkitMediaInformationKEYSIZE(comArthenicaFfmpegkitMediaInformationKEYSIZE int32) ApiListMutedGroupUsersRequest {
	r.comArthenicaFfmpegkitMediaInformationKEYSIZE = &comArthenicaFfmpegkitMediaInformationKEYSIZE
	return r
}

func (r ApiListMutedGroupUsersRequest) Keyword(keyword string) ApiListMutedGroupUsersRequest {
	r.keyword = &keyword
	return r
}

func (r ApiListMutedGroupUsersRequest) Execute() (*GroupMuteUsersResponse, *http.Response, error) {
	return r.ApiService.ListMutedGroupUsersExecute(r)
}

/*
ListMutedGroupUsers Method for ListMutedGroupUsers

Implemented as a Kotlin suspend function on Android.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id
 @return ApiListMutedGroupUsersRequest
*/
func (a *GroupMuteAPIService) ListMutedGroupUsers(ctx context.Context, id int64) ApiListMutedGroupUsersRequest {
	return ApiListMutedGroupUsersRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
//  @return GroupMuteUsersResponse
func (a *GroupMuteAPIService) ListMutedGroupUsersExecute(r ApiListMutedGroupUsersRequest) (*GroupMuteUsersResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *GroupMuteUsersResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "GroupMuteAPIService.ListMutedGroupUsers")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/group_mute/{id}/muted_users"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterValueToString(r.id, "id")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.cursor == nil {
		return localVarReturnValue, nil, reportError("cursor is required and must be specified")
	}
	if r.comArthenicaFfmpegkitMediaInformationKEYSIZE == nil {
		return localVarReturnValue, nil, reportError("comArthenicaFfmpegkitMediaInformationKEYSIZE is required and must be specified")
	}

	if r.keyword != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "keyword", r.keyword, "form", "")
	}
	parameterAddToHeaderOrQuery(localVarQueryParams, "cursor", r.cursor, "form", "")
	parameterAddToHeaderOrQuery(localVarQueryParams, "com.arthenica.ffmpegkit.MediaInformation.KEY_SIZE", r.comArthenicaFfmpegkitMediaInformationKEYSIZE, "form", "")
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

type ApiMuteGroupUserRequest struct {
	ctx context.Context
	ApiService *GroupMuteAPIService
	id int64
	userId int64
}

func (r ApiMuteGroupUserRequest) Execute() (*http.Response, error) {
	return r.ApiService.MuteGroupUserExecute(r)
}

/*
MuteGroupUser Method for MuteGroupUser

Implemented as a Kotlin suspend function on Android.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id
 @param userId
 @return ApiMuteGroupUserRequest
*/
func (a *GroupMuteAPIService) MuteGroupUser(ctx context.Context, id int64, userId int64) ApiMuteGroupUserRequest {
	return ApiMuteGroupUserRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
		userId: userId,
	}
}

// Execute executes the request
func (a *GroupMuteAPIService) MuteGroupUserExecute(r ApiMuteGroupUserRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "GroupMuteAPIService.MuteGroupUser")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/group_mute/{id}/mute/{user_id}"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterValueToString(r.id, "id")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"user_id"+"}", url.PathEscape(parameterValueToString(r.userId, "userId")), -1)

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
	localVarHTTPHeaderAccepts := []string{}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type ApiUnmuteGroupUserRequest struct {
	ctx context.Context
	ApiService *GroupMuteAPIService
	id int64
	userId int64
}

func (r ApiUnmuteGroupUserRequest) Execute() (*http.Response, error) {
	return r.ApiService.UnmuteGroupUserExecute(r)
}

/*
UnmuteGroupUser Method for UnmuteGroupUser

Implemented as a Kotlin suspend function on Android.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id
 @param userId
 @return ApiUnmuteGroupUserRequest
*/
func (a *GroupMuteAPIService) UnmuteGroupUser(ctx context.Context, id int64, userId int64) ApiUnmuteGroupUserRequest {
	return ApiUnmuteGroupUserRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
		userId: userId,
	}
}

// Execute executes the request
func (a *GroupMuteAPIService) UnmuteGroupUserExecute(r ApiUnmuteGroupUserRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "GroupMuteAPIService.UnmuteGroupUser")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/group_mute/{id}/unmute/{user_id}"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterValueToString(r.id, "id")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"user_id"+"}", url.PathEscape(parameterValueToString(r.userId, "userId")), -1)

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
	localVarHTTPHeaderAccepts := []string{}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}
