
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


// ChatRoomsAPIService ChatRoomsAPI service
type ChatRoomsAPIService service

type ApiAcceptChatRequestRequest struct {
	ctx context.Context
	ApiService *ChatRoomsAPIService
	chatRoomIds *[]int64
}

func (r ApiAcceptChatRequestRequest) ChatRoomIds(chatRoomIds []int64) ApiAcceptChatRequestRequest {
	r.chatRoomIds = &chatRoomIds
	return r
}

func (r ApiAcceptChatRequestRequest) Execute() (*http.Response, error) {
	return r.ApiService.AcceptChatRequestExecute(r)
}

func (r ApiAcceptChatRequestRequest) ExecuteRaw() ([]byte, *http.Response, error) {
	httpResp, err := r.Execute()
	return executeRaw(httpResp, err)
}

/*
AcceptChatRequest Method for AcceptChatRequest

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiAcceptChatRequestRequest
*/
func (a *ChatRoomsAPIService) AcceptChatRequest(ctx context.Context) ApiAcceptChatRequestRequest {
	return ApiAcceptChatRequestRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
func (a *ChatRoomsAPIService) AcceptChatRequestExecute(r ApiAcceptChatRequestRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ChatRoomsAPIService.AcceptChatRequest")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/chat_rooms/accept_chat_request"

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
	localVarHTTPHeaderAccepts := []string{}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.chatRoomIds != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "chat_room_ids[]", r.chatRoomIds, "", "csv")
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

type ApiCreateChatRoomRequest struct {
	ctx context.Context
	ApiService *ChatRoomsAPIService
	backgroundFilename *string
	iconFilename *string
	name *string
	withUserIds *[]int64
}

func (r ApiCreateChatRoomRequest) BackgroundFilename(backgroundFilename string) ApiCreateChatRoomRequest {
	r.backgroundFilename = &backgroundFilename
	return r
}

func (r ApiCreateChatRoomRequest) IconFilename(iconFilename string) ApiCreateChatRoomRequest {
	r.iconFilename = &iconFilename
	return r
}

func (r ApiCreateChatRoomRequest) Name(name string) ApiCreateChatRoomRequest {
	r.name = &name
	return r
}

func (r ApiCreateChatRoomRequest) WithUserIds(withUserIds []int64) ApiCreateChatRoomRequest {
	r.withUserIds = &withUserIds
	return r
}

func (r ApiCreateChatRoomRequest) Execute() (*CreateChatRoomResponse, *http.Response, error) {
	return r.ApiService.CreateChatRoomExecute(r)
}

func (r ApiCreateChatRoomRequest) ExecuteRaw() ([]byte, *http.Response, error) {
	_, httpResp, err := r.Execute()
	return executeRaw(httpResp, err)
}

/*
CreateChatRoom Method for CreateChatRoom

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiCreateChatRoomRequest
*/
func (a *ChatRoomsAPIService) CreateChatRoom(ctx context.Context) ApiCreateChatRoomRequest {
	return ApiCreateChatRoomRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return CreateChatRoomResponse
func (a *ChatRoomsAPIService) CreateChatRoomExecute(r ApiCreateChatRoomRequest) (*CreateChatRoomResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *CreateChatRoomResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ChatRoomsAPIService.CreateChatRoom")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v3/chat_rooms/new"

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
	if r.backgroundFilename != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "background_filename", r.backgroundFilename, "", "")
	}
	if r.iconFilename != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "icon_filename", r.iconFilename, "", "")
	}
	if r.name != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "name", r.name, "", "")
	}
	if r.withUserIds != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "with_user_ids[]", r.withUserIds, "", "csv")
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

type ApiCreateChatRoomV1Request struct {
	ctx context.Context
	ApiService *ChatRoomsAPIService
	himaChat *bool
	matchingId *int64
	withUserId *int64
}

func (r ApiCreateChatRoomV1Request) HimaChat(himaChat bool) ApiCreateChatRoomV1Request {
	r.himaChat = &himaChat
	return r
}

func (r ApiCreateChatRoomV1Request) MatchingId(matchingId int64) ApiCreateChatRoomV1Request {
	r.matchingId = &matchingId
	return r
}

func (r ApiCreateChatRoomV1Request) WithUserId(withUserId int64) ApiCreateChatRoomV1Request {
	r.withUserId = &withUserId
	return r
}

func (r ApiCreateChatRoomV1Request) Execute() (*CreateChatRoomResponse, *http.Response, error) {
	return r.ApiService.CreateChatRoomV1Execute(r)
}

func (r ApiCreateChatRoomV1Request) ExecuteRaw() ([]byte, *http.Response, error) {
	_, httpResp, err := r.Execute()
	return executeRaw(httpResp, err)
}

/*
CreateChatRoomV1 Method for CreateChatRoomV1

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiCreateChatRoomV1Request
*/
func (a *ChatRoomsAPIService) CreateChatRoomV1(ctx context.Context) ApiCreateChatRoomV1Request {
	return ApiCreateChatRoomV1Request{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return CreateChatRoomResponse
func (a *ChatRoomsAPIService) CreateChatRoomV1Execute(r ApiCreateChatRoomV1Request) (*CreateChatRoomResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *CreateChatRoomResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ChatRoomsAPIService.CreateChatRoomV1")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/chat_rooms/new"

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
	if r.himaChat != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "hima_chat", r.himaChat, "", "")
	}
	if r.matchingId != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "matching_id", r.matchingId, "", "")
	}
	if r.withUserId != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "with_user_id", r.withUserId, "", "")
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

type ApiDeleteChatMessageRequest struct {
	ctx context.Context
	ApiService *ChatRoomsAPIService
	roomId int64
	messageId int64
}

func (r ApiDeleteChatMessageRequest) Execute() (*http.Response, error) {
	return r.ApiService.DeleteChatMessageExecute(r)
}

func (r ApiDeleteChatMessageRequest) ExecuteRaw() ([]byte, *http.Response, error) {
	httpResp, err := r.Execute()
	return executeRaw(httpResp, err)
}

/*
DeleteChatMessage Method for DeleteChatMessage

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param roomId
 @param messageId
 @return ApiDeleteChatMessageRequest
*/
func (a *ChatRoomsAPIService) DeleteChatMessage(ctx context.Context, roomId int64, messageId int64) ApiDeleteChatMessageRequest {
	return ApiDeleteChatMessageRequest{
		ApiService: a,
		ctx: ctx,
		roomId: roomId,
		messageId: messageId,
	}
}

// Execute executes the request
func (a *ChatRoomsAPIService) DeleteChatMessageExecute(r ApiDeleteChatMessageRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ChatRoomsAPIService.DeleteChatMessage")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/chat_rooms/{room_id}/messages/{message_id}/delete"
	localVarPath = strings.Replace(localVarPath, "{"+"room_id"+"}", url.PathEscape(parameterValueToString(r.roomId, "roomId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"message_id"+"}", url.PathEscape(parameterValueToString(r.messageId, "messageId")), -1)

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

type ApiDeleteChatRoomsRequest struct {
	ctx context.Context
	ApiService *ChatRoomsAPIService
	chatRoomIds *[]int64
}

func (r ApiDeleteChatRoomsRequest) ChatRoomIds(chatRoomIds []int64) ApiDeleteChatRoomsRequest {
	r.chatRoomIds = &chatRoomIds
	return r
}

func (r ApiDeleteChatRoomsRequest) Execute() (*http.Response, error) {
	return r.ApiService.DeleteChatRoomsExecute(r)
}

func (r ApiDeleteChatRoomsRequest) ExecuteRaw() ([]byte, *http.Response, error) {
	httpResp, err := r.Execute()
	return executeRaw(httpResp, err)
}

/*
DeleteChatRooms Method for DeleteChatRooms

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiDeleteChatRoomsRequest
*/
func (a *ChatRoomsAPIService) DeleteChatRooms(ctx context.Context) ApiDeleteChatRoomsRequest {
	return ApiDeleteChatRoomsRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
func (a *ChatRoomsAPIService) DeleteChatRoomsExecute(r ApiDeleteChatRoomsRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ChatRoomsAPIService.DeleteChatRooms")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/chat_rooms/mass_destroy"

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
	localVarHTTPHeaderAccepts := []string{}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.chatRoomIds != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "chat_room_ids[]", r.chatRoomIds, "", "csv")
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

type ApiGetChatMessagesRequest struct {
	ctx context.Context
	ApiService *ChatRoomsAPIService
	id int64
	number *int32
	fromMessageId *int64
	toMessageId *int64
}

func (r ApiGetChatMessagesRequest) Number(number int32) ApiGetChatMessagesRequest {
	r.number = &number
	return r
}

func (r ApiGetChatMessagesRequest) FromMessageId(fromMessageId int64) ApiGetChatMessagesRequest {
	r.fromMessageId = &fromMessageId
	return r
}

func (r ApiGetChatMessagesRequest) ToMessageId(toMessageId int64) ApiGetChatMessagesRequest {
	r.toMessageId = &toMessageId
	return r
}

func (r ApiGetChatMessagesRequest) Execute() (*MessagesResponse, *http.Response, error) {
	return r.ApiService.GetChatMessagesExecute(r)
}

func (r ApiGetChatMessagesRequest) ExecuteRaw() ([]byte, *http.Response, error) {
	_, httpResp, err := r.Execute()
	return executeRaw(httpResp, err)
}

/*
GetChatMessages Method for GetChatMessages

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id
 @return ApiGetChatMessagesRequest
*/
func (a *ChatRoomsAPIService) GetChatMessages(ctx context.Context, id int64) ApiGetChatMessagesRequest {
	return ApiGetChatMessagesRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
//  @return MessagesResponse
func (a *ChatRoomsAPIService) GetChatMessagesExecute(r ApiGetChatMessagesRequest) (*MessagesResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *MessagesResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ChatRoomsAPIService.GetChatMessages")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v2/chat_rooms/{id}/messages"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterValueToString(r.id, "id")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.number != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "number", r.number, "form", "")
	}
	if r.fromMessageId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "from_message_id", r.fromMessageId, "form", "")
	}
	if r.toMessageId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "to_message_id", r.toMessageId, "form", "")
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

type ApiGetChatRequestCountRequest struct {
	ctx context.Context
	ApiService *ChatRoomsAPIService
}

func (r ApiGetChatRequestCountRequest) Execute() (*TotalChatRequestResponse, *http.Response, error) {
	return r.ApiService.GetChatRequestCountExecute(r)
}

func (r ApiGetChatRequestCountRequest) ExecuteRaw() ([]byte, *http.Response, error) {
	_, httpResp, err := r.Execute()
	return executeRaw(httpResp, err)
}

/*
GetChatRequestCount Method for GetChatRequestCount

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiGetChatRequestCountRequest
*/
func (a *ChatRoomsAPIService) GetChatRequestCount(ctx context.Context) ApiGetChatRequestCountRequest {
	return ApiGetChatRequestCountRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return TotalChatRequestResponse
func (a *ChatRoomsAPIService) GetChatRequestCountExecute(r ApiGetChatRequestCountRequest) (*TotalChatRequestResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *TotalChatRequestResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ChatRoomsAPIService.GetChatRequestCount")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/chat_rooms/total_chat_request"

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

type ApiGetChatRequestsRequest struct {
	ctx context.Context
	ApiService *ChatRoomsAPIService
	fromTimestamp *int64
}

func (r ApiGetChatRequestsRequest) FromTimestamp(fromTimestamp int64) ApiGetChatRequestsRequest {
	r.fromTimestamp = &fromTimestamp
	return r
}

func (r ApiGetChatRequestsRequest) Execute() (*ChatRoomsResponse, *http.Response, error) {
	return r.ApiService.GetChatRequestsExecute(r)
}

func (r ApiGetChatRequestsRequest) ExecuteRaw() ([]byte, *http.Response, error) {
	_, httpResp, err := r.Execute()
	return executeRaw(httpResp, err)
}

/*
GetChatRequests Method for GetChatRequests

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiGetChatRequestsRequest
*/
func (a *ChatRoomsAPIService) GetChatRequests(ctx context.Context) ApiGetChatRequestsRequest {
	return ApiGetChatRequestsRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return ChatRoomsResponse
func (a *ChatRoomsAPIService) GetChatRequestsExecute(r ApiGetChatRequestsRequest) (*ChatRoomsResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ChatRoomsResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ChatRoomsAPIService.GetChatRequests")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/chat_rooms/request_list"

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

type ApiGetChatRoomRequest struct {
	ctx context.Context
	ApiService *ChatRoomsAPIService
	id int64
}

func (r ApiGetChatRoomRequest) Execute() (*ChatRoomResponse, *http.Response, error) {
	return r.ApiService.GetChatRoomExecute(r)
}

func (r ApiGetChatRoomRequest) ExecuteRaw() ([]byte, *http.Response, error) {
	_, httpResp, err := r.Execute()
	return executeRaw(httpResp, err)
}

/*
GetChatRoom Method for GetChatRoom

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id
 @return ApiGetChatRoomRequest
*/
func (a *ChatRoomsAPIService) GetChatRoom(ctx context.Context, id int64) ApiGetChatRoomRequest {
	return ApiGetChatRoomRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
//  @return ChatRoomResponse
func (a *ChatRoomsAPIService) GetChatRoomExecute(r ApiGetChatRoomRequest) (*ChatRoomResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ChatRoomResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ChatRoomsAPIService.GetChatRoom")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v2/chat_rooms/{id}"
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

type ApiGetChatUnreadStatusRequest struct {
	ctx context.Context
	ApiService *ChatRoomsAPIService
	fromTime *int64
}

func (r ApiGetChatUnreadStatusRequest) FromTime(fromTime int64) ApiGetChatUnreadStatusRequest {
	r.fromTime = &fromTime
	return r
}

func (r ApiGetChatUnreadStatusRequest) Execute() (*UnreadStatusResponse, *http.Response, error) {
	return r.ApiService.GetChatUnreadStatusExecute(r)
}

func (r ApiGetChatUnreadStatusRequest) ExecuteRaw() ([]byte, *http.Response, error) {
	_, httpResp, err := r.Execute()
	return executeRaw(httpResp, err)
}

/*
GetChatUnreadStatus Method for GetChatUnreadStatus

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiGetChatUnreadStatusRequest
*/
func (a *ChatRoomsAPIService) GetChatUnreadStatus(ctx context.Context) ApiGetChatUnreadStatusRequest {
	return ApiGetChatUnreadStatusRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return UnreadStatusResponse
func (a *ChatRoomsAPIService) GetChatUnreadStatusExecute(r ApiGetChatUnreadStatusRequest) (*UnreadStatusResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *UnreadStatusResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ChatRoomsAPIService.GetChatUnreadStatus")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/chat_rooms/unread_status"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.fromTime != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "from_time", r.fromTime, "form", "")
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

type ApiGetMainChatRoomsRequest struct {
	ctx context.Context
	ApiService *ChatRoomsAPIService
	fromTimestamp *int64
}

func (r ApiGetMainChatRoomsRequest) FromTimestamp(fromTimestamp int64) ApiGetMainChatRoomsRequest {
	r.fromTimestamp = &fromTimestamp
	return r
}

func (r ApiGetMainChatRoomsRequest) Execute() (*ChatRoomsResponse, *http.Response, error) {
	return r.ApiService.GetMainChatRoomsExecute(r)
}

func (r ApiGetMainChatRoomsRequest) ExecuteRaw() ([]byte, *http.Response, error) {
	_, httpResp, err := r.Execute()
	return executeRaw(httpResp, err)
}

/*
GetMainChatRooms Method for GetMainChatRooms

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiGetMainChatRoomsRequest
*/
func (a *ChatRoomsAPIService) GetMainChatRooms(ctx context.Context) ApiGetMainChatRoomsRequest {
	return ApiGetMainChatRoomsRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return ChatRoomsResponse
func (a *ChatRoomsAPIService) GetMainChatRoomsExecute(r ApiGetMainChatRoomsRequest) (*ChatRoomsResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ChatRoomsResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ChatRoomsAPIService.GetMainChatRooms")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/chat_rooms/main_list"

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

type ApiGetUpdatedChatRoomsRequest struct {
	ctx context.Context
	ApiService *ChatRoomsAPIService
	fromTime *int64
}

func (r ApiGetUpdatedChatRoomsRequest) FromTime(fromTime int64) ApiGetUpdatedChatRoomsRequest {
	r.fromTime = &fromTime
	return r
}

func (r ApiGetUpdatedChatRoomsRequest) Execute() (*ChatRoomsResponse, *http.Response, error) {
	return r.ApiService.GetUpdatedChatRoomsExecute(r)
}

func (r ApiGetUpdatedChatRoomsRequest) ExecuteRaw() ([]byte, *http.Response, error) {
	_, httpResp, err := r.Execute()
	return executeRaw(httpResp, err)
}

/*
GetUpdatedChatRooms Method for GetUpdatedChatRooms

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiGetUpdatedChatRoomsRequest
*/
func (a *ChatRoomsAPIService) GetUpdatedChatRooms(ctx context.Context) ApiGetUpdatedChatRoomsRequest {
	return ApiGetUpdatedChatRoomsRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return ChatRoomsResponse
func (a *ChatRoomsAPIService) GetUpdatedChatRoomsExecute(r ApiGetUpdatedChatRoomsRequest) (*ChatRoomsResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ChatRoomsResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ChatRoomsAPIService.GetUpdatedChatRooms")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v2/chat_rooms/update"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.fromTime != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "from_time", r.fromTime, "form", "")
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

type ApiInviteToChatRoomRequest struct {
	ctx context.Context
	ApiService *ChatRoomsAPIService
	id int64
	withUserIds *[]int64
}

func (r ApiInviteToChatRoomRequest) WithUserIds(withUserIds []int64) ApiInviteToChatRoomRequest {
	r.withUserIds = &withUserIds
	return r
}

func (r ApiInviteToChatRoomRequest) Execute() (*http.Response, error) {
	return r.ApiService.InviteToChatRoomExecute(r)
}

func (r ApiInviteToChatRoomRequest) ExecuteRaw() ([]byte, *http.Response, error) {
	httpResp, err := r.Execute()
	return executeRaw(httpResp, err)
}

/*
InviteToChatRoom Method for InviteToChatRoom

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id
 @return ApiInviteToChatRoomRequest
*/
func (a *ChatRoomsAPIService) InviteToChatRoom(ctx context.Context, id int64) ApiInviteToChatRoomRequest {
	return ApiInviteToChatRoomRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
func (a *ChatRoomsAPIService) InviteToChatRoomExecute(r ApiInviteToChatRoomRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ChatRoomsAPIService.InviteToChatRoom")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v2/chat_rooms/{id}/invite"
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
	localVarHTTPHeaderAccepts := []string{}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.withUserIds != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "with_user_ids[]", r.withUserIds, "", "csv")
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

type ApiKickFromChatRoomRequest struct {
	ctx context.Context
	ApiService *ChatRoomsAPIService
	id int64
	withUserIds *[]int64
}

func (r ApiKickFromChatRoomRequest) WithUserIds(withUserIds []int64) ApiKickFromChatRoomRequest {
	r.withUserIds = &withUserIds
	return r
}

func (r ApiKickFromChatRoomRequest) Execute() (*http.Response, error) {
	return r.ApiService.KickFromChatRoomExecute(r)
}

func (r ApiKickFromChatRoomRequest) ExecuteRaw() ([]byte, *http.Response, error) {
	httpResp, err := r.Execute()
	return executeRaw(httpResp, err)
}

/*
KickFromChatRoom Method for KickFromChatRoom

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id
 @return ApiKickFromChatRoomRequest
*/
func (a *ChatRoomsAPIService) KickFromChatRoom(ctx context.Context, id int64) ApiKickFromChatRoomRequest {
	return ApiKickFromChatRoomRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
func (a *ChatRoomsAPIService) KickFromChatRoomExecute(r ApiKickFromChatRoomRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ChatRoomsAPIService.KickFromChatRoom")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v2/chat_rooms/{id}/kick"
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
	localVarHTTPHeaderAccepts := []string{}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.withUserIds != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "with_user_ids[]", r.withUserIds, "", "csv")
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

type ApiPinChatRoomRequest struct {
	ctx context.Context
	ApiService *ChatRoomsAPIService
	id int64
}

func (r ApiPinChatRoomRequest) Execute() (*http.Response, error) {
	return r.ApiService.PinChatRoomExecute(r)
}

func (r ApiPinChatRoomRequest) ExecuteRaw() ([]byte, *http.Response, error) {
	httpResp, err := r.Execute()
	return executeRaw(httpResp, err)
}

/*
PinChatRoom Method for PinChatRoom

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id
 @return ApiPinChatRoomRequest
*/
func (a *ChatRoomsAPIService) PinChatRoom(ctx context.Context, id int64) ApiPinChatRoomRequest {
	return ApiPinChatRoomRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
func (a *ChatRoomsAPIService) PinChatRoomExecute(r ApiPinChatRoomRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ChatRoomsAPIService.PinChatRoom")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/chat_rooms/{id}/pinned"
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

type ApiReadChatAttachmentsRequest struct {
	ctx context.Context
	ApiService *ChatRoomsAPIService
	id int64
	readAttachmentRequest *ReadAttachmentRequest
}

func (r ApiReadChatAttachmentsRequest) ReadAttachmentRequest(readAttachmentRequest ReadAttachmentRequest) ApiReadChatAttachmentsRequest {
	r.readAttachmentRequest = &readAttachmentRequest
	return r
}

func (r ApiReadChatAttachmentsRequest) Execute() (*http.Response, error) {
	return r.ApiService.ReadChatAttachmentsExecute(r)
}

func (r ApiReadChatAttachmentsRequest) ExecuteRaw() ([]byte, *http.Response, error) {
	httpResp, err := r.Execute()
	return executeRaw(httpResp, err)
}

/*
ReadChatAttachments Method for ReadChatAttachments

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id
 @return ApiReadChatAttachmentsRequest
*/
func (a *ChatRoomsAPIService) ReadChatAttachments(ctx context.Context, id int64) ApiReadChatAttachmentsRequest {
	return ApiReadChatAttachmentsRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
func (a *ChatRoomsAPIService) ReadChatAttachmentsExecute(r ApiReadChatAttachmentsRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ChatRoomsAPIService.ReadChatAttachments")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/chat_rooms/{id}/attachments_read"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterValueToString(r.id, "id")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.readAttachmentRequest == nil {
		return nil, reportError("readAttachmentRequest is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

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
	// body params
	localVarPostBody = r.readAttachmentRequest
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

type ApiReadChatMessageRequest struct {
	ctx context.Context
	ApiService *ChatRoomsAPIService
	id int64
	messageId int64
}

func (r ApiReadChatMessageRequest) Execute() (*http.Response, error) {
	return r.ApiService.ReadChatMessageExecute(r)
}

func (r ApiReadChatMessageRequest) ExecuteRaw() ([]byte, *http.Response, error) {
	httpResp, err := r.Execute()
	return executeRaw(httpResp, err)
}

/*
ReadChatMessage Method for ReadChatMessage

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id
 @param messageId
 @return ApiReadChatMessageRequest
*/
func (a *ChatRoomsAPIService) ReadChatMessage(ctx context.Context, id int64, messageId int64) ApiReadChatMessageRequest {
	return ApiReadChatMessageRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
		messageId: messageId,
	}
}

// Execute executes the request
func (a *ChatRoomsAPIService) ReadChatMessageExecute(r ApiReadChatMessageRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ChatRoomsAPIService.ReadChatMessage")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v2/chat_rooms/{id}/messages/{message_id}/read"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterValueToString(r.id, "id")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"message_id"+"}", url.PathEscape(parameterValueToString(r.messageId, "messageId")), -1)

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

type ApiReadChatVideosRequest struct {
	ctx context.Context
	ApiService *ChatRoomsAPIService
	id int64
	videoMsgIds *[]int64
}

func (r ApiReadChatVideosRequest) VideoMsgIds(videoMsgIds []int64) ApiReadChatVideosRequest {
	r.videoMsgIds = &videoMsgIds
	return r
}

func (r ApiReadChatVideosRequest) Execute() (*http.Response, error) {
	return r.ApiService.ReadChatVideosExecute(r)
}

func (r ApiReadChatVideosRequest) ExecuteRaw() ([]byte, *http.Response, error) {
	httpResp, err := r.Execute()
	return executeRaw(httpResp, err)
}

/*
ReadChatVideos Method for ReadChatVideos

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id
 @return ApiReadChatVideosRequest
*/
func (a *ChatRoomsAPIService) ReadChatVideos(ctx context.Context, id int64) ApiReadChatVideosRequest {
	return ApiReadChatVideosRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
func (a *ChatRoomsAPIService) ReadChatVideosExecute(r ApiReadChatVideosRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ChatRoomsAPIService.ReadChatVideos")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/chat_rooms/{id}/videos_read"
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
	localVarHTTPHeaderAccepts := []string{}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.videoMsgIds != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "video_msg_ids[]", r.videoMsgIds, "", "csv")
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

type ApiRemoveChatRoomBackgroundRequest struct {
	ctx context.Context
	ApiService *ChatRoomsAPIService
	id int64
}

func (r ApiRemoveChatRoomBackgroundRequest) Execute() (*http.Response, error) {
	return r.ApiService.RemoveChatRoomBackgroundExecute(r)
}

func (r ApiRemoveChatRoomBackgroundRequest) ExecuteRaw() ([]byte, *http.Response, error) {
	httpResp, err := r.Execute()
	return executeRaw(httpResp, err)
}

/*
RemoveChatRoomBackground Method for RemoveChatRoomBackground

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id
 @return ApiRemoveChatRoomBackgroundRequest
*/
func (a *ChatRoomsAPIService) RemoveChatRoomBackground(ctx context.Context, id int64) ApiRemoveChatRoomBackgroundRequest {
	return ApiRemoveChatRoomBackgroundRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
func (a *ChatRoomsAPIService) RemoveChatRoomBackgroundExecute(r ApiRemoveChatRoomBackgroundRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ChatRoomsAPIService.RemoveChatRoomBackground")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v2/chat_rooms/{id}/background"
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

type ApiReportChatRoomRequest struct {
	ctx context.Context
	ApiService *ChatRoomsAPIService
	chatRoomId int64
	categoryId *int64
	opponentId *int64
	reason *string
	screenshot2Filename *string
	screenshot3Filename *string
	screenshot4Filename *string
	screenshotFilename *string
}

func (r ApiReportChatRoomRequest) CategoryId(categoryId int64) ApiReportChatRoomRequest {
	r.categoryId = &categoryId
	return r
}

func (r ApiReportChatRoomRequest) OpponentId(opponentId int64) ApiReportChatRoomRequest {
	r.opponentId = &opponentId
	return r
}

func (r ApiReportChatRoomRequest) Reason(reason string) ApiReportChatRoomRequest {
	r.reason = &reason
	return r
}

func (r ApiReportChatRoomRequest) Screenshot2Filename(screenshot2Filename string) ApiReportChatRoomRequest {
	r.screenshot2Filename = &screenshot2Filename
	return r
}

func (r ApiReportChatRoomRequest) Screenshot3Filename(screenshot3Filename string) ApiReportChatRoomRequest {
	r.screenshot3Filename = &screenshot3Filename
	return r
}

func (r ApiReportChatRoomRequest) Screenshot4Filename(screenshot4Filename string) ApiReportChatRoomRequest {
	r.screenshot4Filename = &screenshot4Filename
	return r
}

func (r ApiReportChatRoomRequest) ScreenshotFilename(screenshotFilename string) ApiReportChatRoomRequest {
	r.screenshotFilename = &screenshotFilename
	return r
}

func (r ApiReportChatRoomRequest) Execute() (*http.Response, error) {
	return r.ApiService.ReportChatRoomExecute(r)
}

func (r ApiReportChatRoomRequest) ExecuteRaw() ([]byte, *http.Response, error) {
	httpResp, err := r.Execute()
	return executeRaw(httpResp, err)
}

/*
ReportChatRoom Method for ReportChatRoom

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param chatRoomId
 @return ApiReportChatRoomRequest
*/
func (a *ChatRoomsAPIService) ReportChatRoom(ctx context.Context, chatRoomId int64) ApiReportChatRoomRequest {
	return ApiReportChatRoomRequest{
		ApiService: a,
		ctx: ctx,
		chatRoomId: chatRoomId,
	}
}

// Execute executes the request
func (a *ChatRoomsAPIService) ReportChatRoomExecute(r ApiReportChatRoomRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ChatRoomsAPIService.ReportChatRoom")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v3/chat_rooms/{chat_room_id}/report"
	localVarPath = strings.Replace(localVarPath, "{"+"chat_room_id"+"}", url.PathEscape(parameterValueToString(r.chatRoomId, "chatRoomId")), -1)

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
	localVarHTTPHeaderAccepts := []string{}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.categoryId != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "category_id", r.categoryId, "", "")
	}
	if r.opponentId != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "opponent_id", r.opponentId, "", "")
	}
	if r.reason != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "reason", r.reason, "", "")
	}
	if r.screenshot2Filename != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "screenshot_2_filename", r.screenshot2Filename, "", "")
	}
	if r.screenshot3Filename != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "screenshot_3_filename", r.screenshot3Filename, "", "")
	}
	if r.screenshot4Filename != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "screenshot_4_filename", r.screenshot4Filename, "", "")
	}
	if r.screenshotFilename != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "screenshot_filename", r.screenshotFilename, "", "")
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

type ApiSendChatMessageRequest struct {
	ctx context.Context
	ApiService *ChatRoomsAPIService
	id int64
	attachmentFileName *string
	callType *string
	fontSize *int32
	gifImageId *int64
	messageType *MessageType
	parentId *int64
	stickerId *int64
	stickerPackId *int64
	text *string
	videoFileName *string
}

func (r ApiSendChatMessageRequest) AttachmentFileName(attachmentFileName string) ApiSendChatMessageRequest {
	r.attachmentFileName = &attachmentFileName
	return r
}

func (r ApiSendChatMessageRequest) CallType(callType string) ApiSendChatMessageRequest {
	r.callType = &callType
	return r
}

func (r ApiSendChatMessageRequest) FontSize(fontSize int32) ApiSendChatMessageRequest {
	r.fontSize = &fontSize
	return r
}

func (r ApiSendChatMessageRequest) GifImageId(gifImageId int64) ApiSendChatMessageRequest {
	r.gifImageId = &gifImageId
	return r
}

func (r ApiSendChatMessageRequest) MessageType(messageType MessageType) ApiSendChatMessageRequest {
	r.messageType = &messageType
	return r
}

func (r ApiSendChatMessageRequest) ParentId(parentId int64) ApiSendChatMessageRequest {
	r.parentId = &parentId
	return r
}

func (r ApiSendChatMessageRequest) StickerId(stickerId int64) ApiSendChatMessageRequest {
	r.stickerId = &stickerId
	return r
}

func (r ApiSendChatMessageRequest) StickerPackId(stickerPackId int64) ApiSendChatMessageRequest {
	r.stickerPackId = &stickerPackId
	return r
}

func (r ApiSendChatMessageRequest) Text(text string) ApiSendChatMessageRequest {
	r.text = &text
	return r
}

func (r ApiSendChatMessageRequest) VideoFileName(videoFileName string) ApiSendChatMessageRequest {
	r.videoFileName = &videoFileName
	return r
}

func (r ApiSendChatMessageRequest) Execute() (*MessageResponse, *http.Response, error) {
	return r.ApiService.SendChatMessageExecute(r)
}

func (r ApiSendChatMessageRequest) ExecuteRaw() ([]byte, *http.Response, error) {
	_, httpResp, err := r.Execute()
	return executeRaw(httpResp, err)
}

/*
SendChatMessage Method for SendChatMessage

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id
 @return ApiSendChatMessageRequest
*/
func (a *ChatRoomsAPIService) SendChatMessage(ctx context.Context, id int64) ApiSendChatMessageRequest {
	return ApiSendChatMessageRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
//  @return MessageResponse
func (a *ChatRoomsAPIService) SendChatMessageExecute(r ApiSendChatMessageRequest) (*MessageResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *MessageResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ChatRoomsAPIService.SendChatMessage")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v3/chat_rooms/{id}/messages/new"
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
	if r.attachmentFileName != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "attachment_file_name", r.attachmentFileName, "", "")
	}
	if r.callType != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "call_type", r.callType, "", "")
	}
	if r.fontSize != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "font_size", r.fontSize, "", "")
	}
	if r.gifImageId != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "gif_image_id", r.gifImageId, "", "")
	}
	if r.messageType != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "message_type", r.messageType, "", "")
	}
	if r.parentId != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "parent_id", r.parentId, "", "")
	}
	if r.stickerId != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "sticker_id", r.stickerId, "", "")
	}
	if r.stickerPackId != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "sticker_pack_id", r.stickerPackId, "", "")
	}
	if r.text != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "text", r.text, "", "")
	}
	if r.videoFileName != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "video_file_name", r.videoFileName, "", "")
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

type ApiUnpinChatRoomRequest struct {
	ctx context.Context
	ApiService *ChatRoomsAPIService
	id int64
}

func (r ApiUnpinChatRoomRequest) Execute() (*http.Response, error) {
	return r.ApiService.UnpinChatRoomExecute(r)
}

func (r ApiUnpinChatRoomRequest) ExecuteRaw() ([]byte, *http.Response, error) {
	httpResp, err := r.Execute()
	return executeRaw(httpResp, err)
}

/*
UnpinChatRoom Method for UnpinChatRoom

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id
 @return ApiUnpinChatRoomRequest
*/
func (a *ChatRoomsAPIService) UnpinChatRoom(ctx context.Context, id int64) ApiUnpinChatRoomRequest {
	return ApiUnpinChatRoomRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
func (a *ChatRoomsAPIService) UnpinChatRoomExecute(r ApiUnpinChatRoomRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ChatRoomsAPIService.UnpinChatRoom")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/chat_rooms/{id}/pinned"
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

type ApiUpdateChatRoomRequest struct {
	ctx context.Context
	ApiService *ChatRoomsAPIService
	id int64
	backgroundFilename *string
	iconFilename *string
	name *string
}

func (r ApiUpdateChatRoomRequest) BackgroundFilename(backgroundFilename string) ApiUpdateChatRoomRequest {
	r.backgroundFilename = &backgroundFilename
	return r
}

func (r ApiUpdateChatRoomRequest) IconFilename(iconFilename string) ApiUpdateChatRoomRequest {
	r.iconFilename = &iconFilename
	return r
}

func (r ApiUpdateChatRoomRequest) Name(name string) ApiUpdateChatRoomRequest {
	r.name = &name
	return r
}

func (r ApiUpdateChatRoomRequest) Execute() (*http.Response, error) {
	return r.ApiService.UpdateChatRoomExecute(r)
}

func (r ApiUpdateChatRoomRequest) ExecuteRaw() ([]byte, *http.Response, error) {
	httpResp, err := r.Execute()
	return executeRaw(httpResp, err)
}

/*
UpdateChatRoom Method for UpdateChatRoom

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id
 @return ApiUpdateChatRoomRequest
*/
func (a *ChatRoomsAPIService) UpdateChatRoom(ctx context.Context, id int64) ApiUpdateChatRoomRequest {
	return ApiUpdateChatRoomRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
func (a *ChatRoomsAPIService) UpdateChatRoomExecute(r ApiUpdateChatRoomRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ChatRoomsAPIService.UpdateChatRoom")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v3/chat_rooms/{id}/edit"
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
	localVarHTTPHeaderAccepts := []string{}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.backgroundFilename != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "background_filename", r.backgroundFilename, "", "")
	}
	if r.iconFilename != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "icon_filename", r.iconFilename, "", "")
	}
	if r.name != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "name", r.name, "", "")
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
