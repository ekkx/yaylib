
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


// NotificationSettingsAPIService NotificationSettingsAPI service
type NotificationSettingsAPIService service

type ApiGetGroupNotificationSettingsRequest struct {
	ctx context.Context
	ApiService *NotificationSettingsAPIService
	id int64
}

func (r ApiGetGroupNotificationSettingsRequest) Execute() (*GroupNotificationSettingsResponse, *http.Response, error) {
	return r.ApiService.GetGroupNotificationSettingsExecute(r)
}

func (r ApiGetGroupNotificationSettingsRequest) ExecuteRaw() ([]byte, *http.Response, error) {
	_, httpResp, err := r.Execute()
	return executeRaw(httpResp, err)
}

/*
GetGroupNotificationSettings Method for GetGroupNotificationSettings

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id
 @return ApiGetGroupNotificationSettingsRequest
*/
func (a *NotificationSettingsAPIService) GetGroupNotificationSettings(ctx context.Context, id int64) ApiGetGroupNotificationSettingsRequest {
	return ApiGetGroupNotificationSettingsRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
//  @return GroupNotificationSettingsResponse
func (a *NotificationSettingsAPIService) GetGroupNotificationSettingsExecute(r ApiGetGroupNotificationSettingsRequest) (*GroupNotificationSettingsResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *GroupNotificationSettingsResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "NotificationSettingsAPIService.GetGroupNotificationSettings")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v2/notification_settings/groups/{id}"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterValueToString(r.id, "id")), -1)

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

type ApiUpdateChatRoomNotificationSettingsRequest struct {
	ctx context.Context
	ApiService *NotificationSettingsAPIService
	id int64
	notificationChat *int32
}

func (r ApiUpdateChatRoomNotificationSettingsRequest) NotificationChat(notificationChat int32) ApiUpdateChatRoomNotificationSettingsRequest {
	r.notificationChat = &notificationChat
	return r
}

func (r ApiUpdateChatRoomNotificationSettingsRequest) Execute() (*NotificationSettingResponse, *http.Response, error) {
	return r.ApiService.UpdateChatRoomNotificationSettingsExecute(r)
}

func (r ApiUpdateChatRoomNotificationSettingsRequest) ExecuteRaw() ([]byte, *http.Response, error) {
	_, httpResp, err := r.Execute()
	return executeRaw(httpResp, err)
}

/*
UpdateChatRoomNotificationSettings Method for UpdateChatRoomNotificationSettings

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id
 @return ApiUpdateChatRoomNotificationSettingsRequest
*/
func (a *NotificationSettingsAPIService) UpdateChatRoomNotificationSettings(ctx context.Context, id int64) ApiUpdateChatRoomNotificationSettingsRequest {
	return ApiUpdateChatRoomNotificationSettingsRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
//  @return NotificationSettingResponse
func (a *NotificationSettingsAPIService) UpdateChatRoomNotificationSettingsExecute(r ApiUpdateChatRoomNotificationSettingsRequest) (*NotificationSettingResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *NotificationSettingResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "NotificationSettingsAPIService.UpdateChatRoomNotificationSettings")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v2/notification_settings/chat_rooms/{id}"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterValueToString(r.id, "id")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"multipart/form-data"}

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
	if r.notificationChat != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "notification_chat", r.notificationChat, "", "")
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

type ApiUpdateGroupNotificationSettingsRequest struct {
	ctx context.Context
	ApiService *NotificationSettingsAPIService
	id int64
	notificationGroupJoin *int32
	notificationGroupMessageTagAll *int32
	notificationGroupPost *int32
	notificationGroupRequest *int32
}

func (r ApiUpdateGroupNotificationSettingsRequest) NotificationGroupJoin(notificationGroupJoin int32) ApiUpdateGroupNotificationSettingsRequest {
	r.notificationGroupJoin = &notificationGroupJoin
	return r
}

func (r ApiUpdateGroupNotificationSettingsRequest) NotificationGroupMessageTagAll(notificationGroupMessageTagAll int32) ApiUpdateGroupNotificationSettingsRequest {
	r.notificationGroupMessageTagAll = &notificationGroupMessageTagAll
	return r
}

func (r ApiUpdateGroupNotificationSettingsRequest) NotificationGroupPost(notificationGroupPost int32) ApiUpdateGroupNotificationSettingsRequest {
	r.notificationGroupPost = &notificationGroupPost
	return r
}

func (r ApiUpdateGroupNotificationSettingsRequest) NotificationGroupRequest(notificationGroupRequest int32) ApiUpdateGroupNotificationSettingsRequest {
	r.notificationGroupRequest = &notificationGroupRequest
	return r
}

func (r ApiUpdateGroupNotificationSettingsRequest) Execute() (*AdditionalSettingsResponse, *http.Response, error) {
	return r.ApiService.UpdateGroupNotificationSettingsExecute(r)
}

func (r ApiUpdateGroupNotificationSettingsRequest) ExecuteRaw() ([]byte, *http.Response, error) {
	_, httpResp, err := r.Execute()
	return executeRaw(httpResp, err)
}

/*
UpdateGroupNotificationSettings Method for UpdateGroupNotificationSettings

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id
 @return ApiUpdateGroupNotificationSettingsRequest
*/
func (a *NotificationSettingsAPIService) UpdateGroupNotificationSettings(ctx context.Context, id int64) ApiUpdateGroupNotificationSettingsRequest {
	return ApiUpdateGroupNotificationSettingsRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
//  @return AdditionalSettingsResponse
func (a *NotificationSettingsAPIService) UpdateGroupNotificationSettingsExecute(r ApiUpdateGroupNotificationSettingsRequest) (*AdditionalSettingsResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *AdditionalSettingsResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "NotificationSettingsAPIService.UpdateGroupNotificationSettings")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v2/notification_settings/groups/{id}"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterValueToString(r.id, "id")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"multipart/form-data"}

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
	if r.notificationGroupJoin != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "notification_group_join", r.notificationGroupJoin, "", "")
	}
	if r.notificationGroupMessageTagAll != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "notification_group_message_tag_all", r.notificationGroupMessageTagAll, "", "")
	}
	if r.notificationGroupPost != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "notification_group_post", r.notificationGroupPost, "", "")
	}
	if r.notificationGroupRequest != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "notification_group_request", r.notificationGroupRequest, "", "")
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
