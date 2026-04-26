
// Code generated from the Yay! API schema; DO NOT EDIT.

package gen

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
	"reflect"
)


// HiddenAPIService HiddenAPI service
type HiddenAPIService service

type ApiCreateMuteKeywordRequest struct {
	ctx context.Context
	ApiService *HiddenAPIService
	muteKeywordRequest *MuteKeywordRequest
}

func (r ApiCreateMuteKeywordRequest) MuteKeywordRequest(muteKeywordRequest MuteKeywordRequest) ApiCreateMuteKeywordRequest {
	r.muteKeywordRequest = &muteKeywordRequest
	return r
}

func (r ApiCreateMuteKeywordRequest) Execute() (*CreateMuteKeywordResponse, *http.Response, error) {
	return r.ApiService.CreateMuteKeywordExecute(r)
}

/*
CreateMuteKeyword Method for CreateMuteKeyword

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiCreateMuteKeywordRequest
*/
func (a *HiddenAPIService) CreateMuteKeyword(ctx context.Context) ApiCreateMuteKeywordRequest {
	return ApiCreateMuteKeywordRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return CreateMuteKeywordResponse
func (a *HiddenAPIService) CreateMuteKeywordExecute(r ApiCreateMuteKeywordRequest) (*CreateMuteKeywordResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *CreateMuteKeywordResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "HiddenAPIService.CreateMuteKeyword")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/hidden/words"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.muteKeywordRequest == nil {
		return localVarReturnValue, nil, reportError("muteKeywordRequest is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

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
	// body params
	localVarPostBody = r.muteKeywordRequest
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

type ApiDeleteMuteKeywordRequest struct {
	ctx context.Context
	ApiService *HiddenAPIService
	ids *[]int64
}

func (r ApiDeleteMuteKeywordRequest) Ids(ids []int64) ApiDeleteMuteKeywordRequest {
	r.ids = &ids
	return r
}

func (r ApiDeleteMuteKeywordRequest) Execute() (*http.Response, error) {
	return r.ApiService.DeleteMuteKeywordExecute(r)
}

/*
DeleteMuteKeyword Method for DeleteMuteKeyword

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiDeleteMuteKeywordRequest
*/
func (a *HiddenAPIService) DeleteMuteKeyword(ctx context.Context) ApiDeleteMuteKeywordRequest {
	return ApiDeleteMuteKeywordRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
func (a *HiddenAPIService) DeleteMuteKeywordExecute(r ApiDeleteMuteKeywordRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "HiddenAPIService.DeleteMuteKeyword")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/hidden/words"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.ids == nil {
		return nil, reportError("ids is required and must be specified")
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

type ApiHideChatsRequest struct {
	ctx context.Context
	ApiService *HiddenAPIService
	chatRoomId *int64
}

func (r ApiHideChatsRequest) ChatRoomId(chatRoomId int64) ApiHideChatsRequest {
	r.chatRoomId = &chatRoomId
	return r
}

func (r ApiHideChatsRequest) Execute() (*http.Response, error) {
	return r.ApiService.HideChatsExecute(r)
}

/*
HideChats Method for HideChats

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiHideChatsRequest
*/
func (a *HiddenAPIService) HideChats(ctx context.Context) ApiHideChatsRequest {
	return ApiHideChatsRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
func (a *HiddenAPIService) HideChatsExecute(r ApiHideChatsRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "HiddenAPIService.HideChats")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/hidden/chats"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.chatRoomId == nil {
		return nil, reportError("chatRoomId is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"multipart/form-data"}

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
	parameterAddToHeaderOrQuery(localVarFormParams, "chat_room_id", r.chatRoomId, "", "")
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

type ApiHideUsersRequest struct {
	ctx context.Context
	ApiService *HiddenAPIService
	userId *int64
}

func (r ApiHideUsersRequest) UserId(userId int64) ApiHideUsersRequest {
	r.userId = &userId
	return r
}

func (r ApiHideUsersRequest) Execute() (*http.Response, error) {
	return r.ApiService.HideUsersExecute(r)
}

/*
HideUsers Method for HideUsers

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiHideUsersRequest
*/
func (a *HiddenAPIService) HideUsers(ctx context.Context) ApiHideUsersRequest {
	return ApiHideUsersRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
func (a *HiddenAPIService) HideUsersExecute(r ApiHideUsersRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "HiddenAPIService.HideUsers")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/hidden/users"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.userId == nil {
		return nil, reportError("userId is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"multipart/form-data"}

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
	parameterAddToHeaderOrQuery(localVarFormParams, "user_id", r.userId, "", "")
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

type ApiListHiddenChatsRequest struct {
	ctx context.Context
	ApiService *HiddenAPIService
	number *int32
	fromTimestamp *int64
}

func (r ApiListHiddenChatsRequest) Number(number int32) ApiListHiddenChatsRequest {
	r.number = &number
	return r
}

func (r ApiListHiddenChatsRequest) FromTimestamp(fromTimestamp int64) ApiListHiddenChatsRequest {
	r.fromTimestamp = &fromTimestamp
	return r
}

func (r ApiListHiddenChatsRequest) Execute() (*ChatRoomsResponse, *http.Response, error) {
	return r.ApiService.ListHiddenChatsExecute(r)
}

/*
ListHiddenChats Method for ListHiddenChats

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiListHiddenChatsRequest
*/
func (a *HiddenAPIService) ListHiddenChats(ctx context.Context) ApiListHiddenChatsRequest {
	return ApiListHiddenChatsRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return ChatRoomsResponse
func (a *HiddenAPIService) ListHiddenChatsExecute(r ApiListHiddenChatsRequest) (*ChatRoomsResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ChatRoomsResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "HiddenAPIService.ListHiddenChats")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/hidden/chats"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.number == nil {
		return localVarReturnValue, nil, reportError("number is required and must be specified")
	}

	parameterAddToHeaderOrQuery(localVarQueryParams, "number", r.number, "form", "")
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

type ApiListHiddenUsersRequest struct {
	ctx context.Context
	ApiService *HiddenAPIService
	from *string
	number *int32
}

func (r ApiListHiddenUsersRequest) From(from string) ApiListHiddenUsersRequest {
	r.from = &from
	return r
}

func (r ApiListHiddenUsersRequest) Number(number int32) ApiListHiddenUsersRequest {
	r.number = &number
	return r
}

func (r ApiListHiddenUsersRequest) Execute() (*HiddenResponse, *http.Response, error) {
	return r.ApiService.ListHiddenUsersExecute(r)
}

/*
ListHiddenUsers Method for ListHiddenUsers

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiListHiddenUsersRequest
*/
func (a *HiddenAPIService) ListHiddenUsers(ctx context.Context) ApiListHiddenUsersRequest {
	return ApiListHiddenUsersRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return HiddenResponse
func (a *HiddenAPIService) ListHiddenUsersExecute(r ApiListHiddenUsersRequest) (*HiddenResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *HiddenResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "HiddenAPIService.ListHiddenUsers")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/hidden/users"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.from != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "from", r.from, "form", "")
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

type ApiListMuteKeywordsRequest struct {
	ctx context.Context
	ApiService *HiddenAPIService
}

func (r ApiListMuteKeywordsRequest) Execute() (*MuteKeywordResponse, *http.Response, error) {
	return r.ApiService.ListMuteKeywordsExecute(r)
}

/*
ListMuteKeywords Method for ListMuteKeywords

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiListMuteKeywordsRequest
*/
func (a *HiddenAPIService) ListMuteKeywords(ctx context.Context) ApiListMuteKeywordsRequest {
	return ApiListMuteKeywordsRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return MuteKeywordResponse
func (a *HiddenAPIService) ListMuteKeywordsExecute(r ApiListMuteKeywordsRequest) (*MuteKeywordResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *MuteKeywordResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "HiddenAPIService.ListMuteKeywords")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/hidden/words"

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

type ApiUnhideChatsRequest struct {
	ctx context.Context
	ApiService *HiddenAPIService
	chatRoomIds *int64
}

func (r ApiUnhideChatsRequest) ChatRoomIds(chatRoomIds int64) ApiUnhideChatsRequest {
	r.chatRoomIds = &chatRoomIds
	return r
}

func (r ApiUnhideChatsRequest) Execute() (*http.Response, error) {
	return r.ApiService.UnhideChatsExecute(r)
}

/*
UnhideChats Method for UnhideChats

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiUnhideChatsRequest
*/
func (a *HiddenAPIService) UnhideChats(ctx context.Context) ApiUnhideChatsRequest {
	return ApiUnhideChatsRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
func (a *HiddenAPIService) UnhideChatsExecute(r ApiUnhideChatsRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "HiddenAPIService.UnhideChats")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/hidden/chats"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.chatRoomIds == nil {
		return nil, reportError("chatRoomIds is required and must be specified")
	}

	parameterAddToHeaderOrQuery(localVarQueryParams, "chat_room_ids", r.chatRoomIds, "form", "")
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

type ApiUnhideUsersRequest struct {
	ctx context.Context
	ApiService *HiddenAPIService
	userIds *[]int64
}

func (r ApiUnhideUsersRequest) UserIds(userIds []int64) ApiUnhideUsersRequest {
	r.userIds = &userIds
	return r
}

func (r ApiUnhideUsersRequest) Execute() (*http.Response, error) {
	return r.ApiService.UnhideUsersExecute(r)
}

/*
UnhideUsers Method for UnhideUsers

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiUnhideUsersRequest
*/
func (a *HiddenAPIService) UnhideUsers(ctx context.Context) ApiUnhideUsersRequest {
	return ApiUnhideUsersRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
func (a *HiddenAPIService) UnhideUsersExecute(r ApiUnhideUsersRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "HiddenAPIService.UnhideUsers")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/hidden/users"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.userIds == nil {
		return nil, reportError("userIds is required and must be specified")
	}

	{
		t := *r.userIds
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				parameterAddToHeaderOrQuery(localVarQueryParams, "user_ids[]", s.Index(i).Interface(), "form", "multi")
			}
		} else {
			parameterAddToHeaderOrQuery(localVarQueryParams, "user_ids[]", t, "form", "multi")
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
