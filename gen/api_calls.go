
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


// CallsAPIService CallsAPI service
type CallsAPIService service

type ApiBulkInviteToCallRequest struct {
	ctx context.Context
	ApiService *CallsAPIService
	callId int64
	groupId *int64
}

func (r ApiBulkInviteToCallRequest) GroupId(groupId int64) ApiBulkInviteToCallRequest {
	r.groupId = &groupId
	return r
}

func (r ApiBulkInviteToCallRequest) Execute() (*http.Response, error) {
	return r.ApiService.BulkInviteToCallExecute(r)
}

func (r ApiBulkInviteToCallRequest) ExecuteRaw() ([]byte, *http.Response, error) {
	httpResp, err := r.Execute()
	return executeRaw(httpResp, err)
}

/*
BulkInviteToCall Method for BulkInviteToCall

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param callId
 @return ApiBulkInviteToCallRequest
*/
func (a *CallsAPIService) BulkInviteToCall(ctx context.Context, callId int64) ApiBulkInviteToCallRequest {
	return ApiBulkInviteToCallRequest{
		ApiService: a,
		ctx: ctx,
		callId: callId,
	}
}

// Execute executes the request
func (a *CallsAPIService) BulkInviteToCallExecute(r ApiBulkInviteToCallRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CallsAPIService.BulkInviteToCall")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/calls/{call_id}/bulk_invite"
	localVarPath = strings.Replace(localVarPath, "{"+"call_id"+"}", url.PathEscape(parameterValueToString(r.callId, "callId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.groupId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "group_id", r.groupId, "form", "")
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

type ApiBumpCallRequest struct {
	ctx context.Context
	ApiService *CallsAPIService
	callId int64
	participantLimit *int32
}

func (r ApiBumpCallRequest) ParticipantLimit(participantLimit int32) ApiBumpCallRequest {
	r.participantLimit = &participantLimit
	return r
}

func (r ApiBumpCallRequest) Execute() (*http.Response, error) {
	return r.ApiService.BumpCallExecute(r)
}

func (r ApiBumpCallRequest) ExecuteRaw() ([]byte, *http.Response, error) {
	httpResp, err := r.Execute()
	return executeRaw(httpResp, err)
}

/*
BumpCall Method for BumpCall

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param callId
 @return ApiBumpCallRequest
*/
func (a *CallsAPIService) BumpCall(ctx context.Context, callId int64) ApiBumpCallRequest {
	return ApiBumpCallRequest{
		ApiService: a,
		ctx: ctx,
		callId: callId,
	}
}

// Execute executes the request
func (a *CallsAPIService) BumpCallExecute(r ApiBumpCallRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CallsAPIService.BumpCall")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/calls/{call_id}/bump"
	localVarPath = strings.Replace(localVarPath, "{"+"call_id"+"}", url.PathEscape(parameterValueToString(r.callId, "callId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.participantLimit != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "participant_limit", r.participantLimit, "form", "")
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

type ApiGenerateCallActionSignatureRequest struct {
	ctx context.Context
	ApiService *CallsAPIService
	action *string
	conferenceId *int64
	targetUserUuid *string
}

func (r ApiGenerateCallActionSignatureRequest) Action(action string) ApiGenerateCallActionSignatureRequest {
	r.action = &action
	return r
}

func (r ApiGenerateCallActionSignatureRequest) ConferenceId(conferenceId int64) ApiGenerateCallActionSignatureRequest {
	r.conferenceId = &conferenceId
	return r
}

func (r ApiGenerateCallActionSignatureRequest) TargetUserUuid(targetUserUuid string) ApiGenerateCallActionSignatureRequest {
	r.targetUserUuid = &targetUserUuid
	return r
}

func (r ApiGenerateCallActionSignatureRequest) Execute() (*CallActionSignatureResponse, *http.Response, error) {
	return r.ApiService.GenerateCallActionSignatureExecute(r)
}

func (r ApiGenerateCallActionSignatureRequest) ExecuteRaw() ([]byte, *http.Response, error) {
	_, httpResp, err := r.Execute()
	return executeRaw(httpResp, err)
}

/*
GenerateCallActionSignature Method for GenerateCallActionSignature

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiGenerateCallActionSignatureRequest
*/
func (a *CallsAPIService) GenerateCallActionSignature(ctx context.Context) ApiGenerateCallActionSignatureRequest {
	return ApiGenerateCallActionSignatureRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return CallActionSignatureResponse
func (a *CallsAPIService) GenerateCallActionSignatureExecute(r ApiGenerateCallActionSignatureRequest) (*CallActionSignatureResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *CallActionSignatureResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CallsAPIService.GenerateCallActionSignature")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/calls/action_signature/generate"

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
	if r.action != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "action", r.action, "", "")
	}
	if r.conferenceId != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "conference_id", r.conferenceId, "", "")
	}
	if r.targetUserUuid != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "target_user_uuid", r.targetUserUuid, "", "")
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

type ApiGetAgoraRtmTokenRequest struct {
	ctx context.Context
	ApiService *CallsAPIService
	callId int64
}

func (r ApiGetAgoraRtmTokenRequest) Execute() (*RtmTokenResponse, *http.Response, error) {
	return r.ApiService.GetAgoraRtmTokenExecute(r)
}

func (r ApiGetAgoraRtmTokenRequest) ExecuteRaw() ([]byte, *http.Response, error) {
	_, httpResp, err := r.Execute()
	return executeRaw(httpResp, err)
}

/*
GetAgoraRtmToken Method for GetAgoraRtmToken

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param callId
 @return ApiGetAgoraRtmTokenRequest
*/
func (a *CallsAPIService) GetAgoraRtmToken(ctx context.Context, callId int64) ApiGetAgoraRtmTokenRequest {
	return ApiGetAgoraRtmTokenRequest{
		ApiService: a,
		ctx: ctx,
		callId: callId,
	}
}

// Execute executes the request
//  @return RtmTokenResponse
func (a *CallsAPIService) GetAgoraRtmTokenExecute(r ApiGetAgoraRtmTokenRequest) (*RtmTokenResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *RtmTokenResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CallsAPIService.GetAgoraRtmToken")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v3/calls/{call_id}/agora_rtm_token"
	localVarPath = strings.Replace(localVarPath, "{"+"call_id"+"}", url.PathEscape(parameterValueToString(r.callId, "callId")), -1)

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

type ApiGetCallBgmsRequest struct {
	ctx context.Context
	ApiService *CallsAPIService
}

func (r ApiGetCallBgmsRequest) Execute() (*BgmsResponse, *http.Response, error) {
	return r.ApiService.GetCallBgmsExecute(r)
}

func (r ApiGetCallBgmsRequest) ExecuteRaw() ([]byte, *http.Response, error) {
	_, httpResp, err := r.Execute()
	return executeRaw(httpResp, err)
}

/*
GetCallBgms Method for GetCallBgms

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiGetCallBgmsRequest
*/
func (a *CallsAPIService) GetCallBgms(ctx context.Context) ApiGetCallBgmsRequest {
	return ApiGetCallBgmsRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return BgmsResponse
func (a *CallsAPIService) GetCallBgmsExecute(r ApiGetCallBgmsRequest) (*BgmsResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *BgmsResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CallsAPIService.GetCallBgms")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/calls/bgm"

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

type ApiGetCallGiftHistoryRequest struct {
	ctx context.Context
	ApiService *CallsAPIService
	callId int64
	from *string
	number *int32
}

func (r ApiGetCallGiftHistoryRequest) From(from string) ApiGetCallGiftHistoryRequest {
	r.from = &from
	return r
}

func (r ApiGetCallGiftHistoryRequest) Number(number int32) ApiGetCallGiftHistoryRequest {
	r.number = &number
	return r
}

func (r ApiGetCallGiftHistoryRequest) Execute() (*CallGiftHistoryResponse, *http.Response, error) {
	return r.ApiService.GetCallGiftHistoryExecute(r)
}

func (r ApiGetCallGiftHistoryRequest) ExecuteRaw() ([]byte, *http.Response, error) {
	_, httpResp, err := r.Execute()
	return executeRaw(httpResp, err)
}

/*
GetCallGiftHistory Method for GetCallGiftHistory

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param callId
 @return ApiGetCallGiftHistoryRequest
*/
func (a *CallsAPIService) GetCallGiftHistory(ctx context.Context, callId int64) ApiGetCallGiftHistoryRequest {
	return ApiGetCallGiftHistoryRequest{
		ApiService: a,
		ctx: ctx,
		callId: callId,
	}
}

// Execute executes the request
//  @return CallGiftHistoryResponse
func (a *CallsAPIService) GetCallGiftHistoryExecute(r ApiGetCallGiftHistoryRequest) (*CallGiftHistoryResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *CallGiftHistoryResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CallsAPIService.GetCallGiftHistory")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/calls/{call_id}/gift_transactions"
	localVarPath = strings.Replace(localVarPath, "{"+"call_id"+"}", url.PathEscape(parameterValueToString(r.callId, "callId")), -1)

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

type ApiGetConferenceCallRequest struct {
	ctx context.Context
	ApiService *CallsAPIService
	callId int64
}

func (r ApiGetConferenceCallRequest) Execute() (*ConferenceCallResponse, *http.Response, error) {
	return r.ApiService.GetConferenceCallExecute(r)
}

func (r ApiGetConferenceCallRequest) ExecuteRaw() ([]byte, *http.Response, error) {
	_, httpResp, err := r.Execute()
	return executeRaw(httpResp, err)
}

/*
GetConferenceCall Method for GetConferenceCall

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param callId
 @return ApiGetConferenceCallRequest
*/
func (a *CallsAPIService) GetConferenceCall(ctx context.Context, callId int64) ApiGetConferenceCallRequest {
	return ApiGetConferenceCallRequest{
		ApiService: a,
		ctx: ctx,
		callId: callId,
	}
}

// Execute executes the request
//  @return ConferenceCallResponse
func (a *CallsAPIService) GetConferenceCallExecute(r ApiGetConferenceCallRequest) (*ConferenceCallResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ConferenceCallResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CallsAPIService.GetConferenceCall")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v2/calls/conferences/{call_id}"
	localVarPath = strings.Replace(localVarPath, "{"+"call_id"+"}", url.PathEscape(parameterValueToString(r.callId, "callId")), -1)

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

type ApiGetInvitableCallUsersRequest struct {
	ctx context.Context
	ApiService *CallsAPIService
	callId int64
	fromTimestamp *int64
	userNickname *string
}

func (r ApiGetInvitableCallUsersRequest) FromTimestamp(fromTimestamp int64) ApiGetInvitableCallUsersRequest {
	r.fromTimestamp = &fromTimestamp
	return r
}

func (r ApiGetInvitableCallUsersRequest) UserNickname(userNickname string) ApiGetInvitableCallUsersRequest {
	r.userNickname = &userNickname
	return r
}

func (r ApiGetInvitableCallUsersRequest) Execute() (*UsersByTimestampResponse, *http.Response, error) {
	return r.ApiService.GetInvitableCallUsersExecute(r)
}

func (r ApiGetInvitableCallUsersRequest) ExecuteRaw() ([]byte, *http.Response, error) {
	_, httpResp, err := r.Execute()
	return executeRaw(httpResp, err)
}

/*
GetInvitableCallUsers Method for GetInvitableCallUsers

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param callId
 @return ApiGetInvitableCallUsersRequest
*/
func (a *CallsAPIService) GetInvitableCallUsers(ctx context.Context, callId int64) ApiGetInvitableCallUsersRequest {
	return ApiGetInvitableCallUsersRequest{
		ApiService: a,
		ctx: ctx,
		callId: callId,
	}
}

// Execute executes the request
//  @return UsersByTimestampResponse
func (a *CallsAPIService) GetInvitableCallUsersExecute(r ApiGetInvitableCallUsersRequest) (*UsersByTimestampResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *UsersByTimestampResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CallsAPIService.GetInvitableCallUsers")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/calls/{call_id}/users/invitable"
	localVarPath = strings.Replace(localVarPath, "{"+"call_id"+"}", url.PathEscape(parameterValueToString(r.callId, "callId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.fromTimestamp != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "from_timestamp", r.fromTimestamp, "form", "")
	}
	if r.userNickname != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "user[nickname]", r.userNickname, "form", "")
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

type ApiGetPhoneStatusRequest struct {
	ctx context.Context
	ApiService *CallsAPIService
	opponentId int64
}

func (r ApiGetPhoneStatusRequest) Execute() (*CallStatusResponse, *http.Response, error) {
	return r.ApiService.GetPhoneStatusExecute(r)
}

func (r ApiGetPhoneStatusRequest) ExecuteRaw() ([]byte, *http.Response, error) {
	_, httpResp, err := r.Execute()
	return executeRaw(httpResp, err)
}

/*
GetPhoneStatus Method for GetPhoneStatus

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param opponentId
 @return ApiGetPhoneStatusRequest
*/
func (a *CallsAPIService) GetPhoneStatus(ctx context.Context, opponentId int64) ApiGetPhoneStatusRequest {
	return ApiGetPhoneStatusRequest{
		ApiService: a,
		ctx: ctx,
		opponentId: opponentId,
	}
}

// Execute executes the request
//  @return CallStatusResponse
func (a *CallsAPIService) GetPhoneStatusExecute(r ApiGetPhoneStatusRequest) (*CallStatusResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *CallStatusResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CallsAPIService.GetPhoneStatus")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/calls/phone_status/{opponent_id}"
	localVarPath = strings.Replace(localVarPath, "{"+"opponent_id"+"}", url.PathEscape(parameterValueToString(r.opponentId, "opponentId")), -1)

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

type ApiInviteToCallRequest struct {
	ctx context.Context
	ApiService *CallsAPIService
	chatRoomId *int64
	roomId *int64
	roomUrl *string
}

func (r ApiInviteToCallRequest) ChatRoomId(chatRoomId int64) ApiInviteToCallRequest {
	r.chatRoomId = &chatRoomId
	return r
}

func (r ApiInviteToCallRequest) RoomId(roomId int64) ApiInviteToCallRequest {
	r.roomId = &roomId
	return r
}

func (r ApiInviteToCallRequest) RoomUrl(roomUrl string) ApiInviteToCallRequest {
	r.roomUrl = &roomUrl
	return r
}

func (r ApiInviteToCallRequest) Execute() (*http.Response, error) {
	return r.ApiService.InviteToCallExecute(r)
}

func (r ApiInviteToCallRequest) ExecuteRaw() ([]byte, *http.Response, error) {
	httpResp, err := r.Execute()
	return executeRaw(httpResp, err)
}

/*
InviteToCall Method for InviteToCall

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiInviteToCallRequest
*/
func (a *CallsAPIService) InviteToCall(ctx context.Context) ApiInviteToCallRequest {
	return ApiInviteToCallRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
func (a *CallsAPIService) InviteToCallExecute(r ApiInviteToCallRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CallsAPIService.InviteToCall")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v2/calls/invite"

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
	if r.chatRoomId != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "chat_room_id", r.chatRoomId, "", "")
	}
	if r.roomId != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "room_id", r.roomId, "", "")
	}
	if r.roomUrl != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "room_url", r.roomUrl, "", "")
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

type ApiInviteToConferenceCallRequest struct {
	ctx context.Context
	ApiService *CallsAPIService
	callId int64
	userIds *[]int64
}

func (r ApiInviteToConferenceCallRequest) UserIds(userIds []int64) ApiInviteToConferenceCallRequest {
	r.userIds = &userIds
	return r
}

func (r ApiInviteToConferenceCallRequest) Execute() (*http.Response, error) {
	return r.ApiService.InviteToConferenceCallExecute(r)
}

func (r ApiInviteToConferenceCallRequest) ExecuteRaw() ([]byte, *http.Response, error) {
	httpResp, err := r.Execute()
	return executeRaw(httpResp, err)
}

/*
InviteToConferenceCall Method for InviteToConferenceCall

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param callId
 @return ApiInviteToConferenceCallRequest
*/
func (a *CallsAPIService) InviteToConferenceCall(ctx context.Context, callId int64) ApiInviteToConferenceCallRequest {
	return ApiInviteToConferenceCallRequest{
		ApiService: a,
		ctx: ctx,
		callId: callId,
	}
}

// Execute executes the request
func (a *CallsAPIService) InviteToConferenceCallExecute(r ApiInviteToConferenceCallRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CallsAPIService.InviteToConferenceCall")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/calls/conference_calls/{call_id}/invite"
	localVarPath = strings.Replace(localVarPath, "{"+"call_id"+"}", url.PathEscape(parameterValueToString(r.callId, "callId")), -1)

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
	if r.userIds != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "user_ids[]", r.userIds, "", "csv")
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

type ApiKickFromConferenceCallRequest struct {
	ctx context.Context
	ApiService *CallsAPIService
	callId int64
	ban *bool
	uuid *string
}

func (r ApiKickFromConferenceCallRequest) Ban(ban bool) ApiKickFromConferenceCallRequest {
	r.ban = &ban
	return r
}

func (r ApiKickFromConferenceCallRequest) Uuid(uuid string) ApiKickFromConferenceCallRequest {
	r.uuid = &uuid
	return r
}

func (r ApiKickFromConferenceCallRequest) Execute() (*CallActionSignatureResponse, *http.Response, error) {
	return r.ApiService.KickFromConferenceCallExecute(r)
}

func (r ApiKickFromConferenceCallRequest) ExecuteRaw() ([]byte, *http.Response, error) {
	_, httpResp, err := r.Execute()
	return executeRaw(httpResp, err)
}

/*
KickFromConferenceCall Method for KickFromConferenceCall

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param callId
 @return ApiKickFromConferenceCallRequest
*/
func (a *CallsAPIService) KickFromConferenceCall(ctx context.Context, callId int64) ApiKickFromConferenceCallRequest {
	return ApiKickFromConferenceCallRequest{
		ApiService: a,
		ctx: ctx,
		callId: callId,
	}
}

// Execute executes the request
//  @return CallActionSignatureResponse
func (a *CallsAPIService) KickFromConferenceCallExecute(r ApiKickFromConferenceCallRequest) (*CallActionSignatureResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *CallActionSignatureResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CallsAPIService.KickFromConferenceCall")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v3/calls/conference_calls/{call_id}/kick"
	localVarPath = strings.Replace(localVarPath, "{"+"call_id"+"}", url.PathEscape(parameterValueToString(r.callId, "callId")), -1)

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
	if r.ban != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "ban", r.ban, "", "")
	}
	if r.uuid != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "uuid", r.uuid, "", "")
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

type ApiLeaveAgoraChannelRequest struct {
	ctx context.Context
	ApiService *CallsAPIService
	conferenceId *int64
	userId *int64
}

func (r ApiLeaveAgoraChannelRequest) ConferenceId(conferenceId int64) ApiLeaveAgoraChannelRequest {
	r.conferenceId = &conferenceId
	return r
}

func (r ApiLeaveAgoraChannelRequest) UserId(userId int64) ApiLeaveAgoraChannelRequest {
	r.userId = &userId
	return r
}

func (r ApiLeaveAgoraChannelRequest) Execute() (*http.Response, error) {
	return r.ApiService.LeaveAgoraChannelExecute(r)
}

func (r ApiLeaveAgoraChannelRequest) ExecuteRaw() ([]byte, *http.Response, error) {
	httpResp, err := r.Execute()
	return executeRaw(httpResp, err)
}

/*
LeaveAgoraChannel Method for LeaveAgoraChannel

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiLeaveAgoraChannelRequest
*/
func (a *CallsAPIService) LeaveAgoraChannel(ctx context.Context) ApiLeaveAgoraChannelRequest {
	return ApiLeaveAgoraChannelRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
func (a *CallsAPIService) LeaveAgoraChannelExecute(r ApiLeaveAgoraChannelRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CallsAPIService.LeaveAgoraChannel")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/calls/leave_agora_channel"

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
	if r.conferenceId != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "conference_id", r.conferenceId, "", "")
	}
	if r.userId != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "user_id", r.userId, "", "")
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

type ApiLeaveConferenceCallRequest struct {
	ctx context.Context
	ApiService *CallsAPIService
	callSid *string
	conferenceId *int64
	duration *int64
	totalUsersInCall *int32
}

func (r ApiLeaveConferenceCallRequest) CallSid(callSid string) ApiLeaveConferenceCallRequest {
	r.callSid = &callSid
	return r
}

func (r ApiLeaveConferenceCallRequest) ConferenceId(conferenceId int64) ApiLeaveConferenceCallRequest {
	r.conferenceId = &conferenceId
	return r
}

func (r ApiLeaveConferenceCallRequest) Duration(duration int64) ApiLeaveConferenceCallRequest {
	r.duration = &duration
	return r
}

func (r ApiLeaveConferenceCallRequest) TotalUsersInCall(totalUsersInCall int32) ApiLeaveConferenceCallRequest {
	r.totalUsersInCall = &totalUsersInCall
	return r
}

func (r ApiLeaveConferenceCallRequest) Execute() (*http.Response, error) {
	return r.ApiService.LeaveConferenceCallExecute(r)
}

func (r ApiLeaveConferenceCallRequest) ExecuteRaw() ([]byte, *http.Response, error) {
	httpResp, err := r.Execute()
	return executeRaw(httpResp, err)
}

/*
LeaveConferenceCall Method for LeaveConferenceCall

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiLeaveConferenceCallRequest
*/
func (a *CallsAPIService) LeaveConferenceCall(ctx context.Context) ApiLeaveConferenceCallRequest {
	return ApiLeaveConferenceCallRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
func (a *CallsAPIService) LeaveConferenceCallExecute(r ApiLeaveConferenceCallRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CallsAPIService.LeaveConferenceCall")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/calls/leave_conference_call"

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
	if r.callSid != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "call_sid", r.callSid, "", "")
	}
	if r.conferenceId != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "conference_id", r.conferenceId, "", "")
	}
	if r.duration != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "duration", r.duration, "", "")
	}
	if r.totalUsersInCall != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "total_users_in_call", r.totalUsersInCall, "", "")
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

type ApiStartConferenceCallRequest struct {
	ctx context.Context
	ApiService *CallsAPIService
	callSid *string
	conferenceId *int64
}

func (r ApiStartConferenceCallRequest) CallSid(callSid string) ApiStartConferenceCallRequest {
	r.callSid = &callSid
	return r
}

func (r ApiStartConferenceCallRequest) ConferenceId(conferenceId int64) ApiStartConferenceCallRequest {
	r.conferenceId = &conferenceId
	return r
}

func (r ApiStartConferenceCallRequest) Execute() (*ConferenceCallResponse, *http.Response, error) {
	return r.ApiService.StartConferenceCallExecute(r)
}

func (r ApiStartConferenceCallRequest) ExecuteRaw() ([]byte, *http.Response, error) {
	_, httpResp, err := r.Execute()
	return executeRaw(httpResp, err)
}

/*
StartConferenceCall Method for StartConferenceCall

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiStartConferenceCallRequest
*/
func (a *CallsAPIService) StartConferenceCall(ctx context.Context) ApiStartConferenceCallRequest {
	return ApiStartConferenceCallRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return ConferenceCallResponse
func (a *CallsAPIService) StartConferenceCallExecute(r ApiStartConferenceCallRequest) (*ConferenceCallResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ConferenceCallResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CallsAPIService.StartConferenceCall")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v2/calls/start_conference_call"

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
	if r.callSid != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "call_sid", r.callSid, "", "")
	}
	if r.conferenceId != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "conference_id", r.conferenceId, "", "")
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

type ApiUpdateCallRequest struct {
	ctx context.Context
	ApiService *CallsAPIService
	callId int64
	categoryId *string
	gameTitle *string
	joinableBy *string
}

func (r ApiUpdateCallRequest) CategoryId(categoryId string) ApiUpdateCallRequest {
	r.categoryId = &categoryId
	return r
}

func (r ApiUpdateCallRequest) GameTitle(gameTitle string) ApiUpdateCallRequest {
	r.gameTitle = &gameTitle
	return r
}

func (r ApiUpdateCallRequest) JoinableBy(joinableBy string) ApiUpdateCallRequest {
	r.joinableBy = &joinableBy
	return r
}

func (r ApiUpdateCallRequest) Execute() (*http.Response, error) {
	return r.ApiService.UpdateCallExecute(r)
}

func (r ApiUpdateCallRequest) ExecuteRaw() ([]byte, *http.Response, error) {
	httpResp, err := r.Execute()
	return executeRaw(httpResp, err)
}

/*
UpdateCall Method for UpdateCall

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param callId
 @return ApiUpdateCallRequest
*/
func (a *CallsAPIService) UpdateCall(ctx context.Context, callId int64) ApiUpdateCallRequest {
	return ApiUpdateCallRequest{
		ApiService: a,
		ctx: ctx,
		callId: callId,
	}
}

// Execute executes the request
func (a *CallsAPIService) UpdateCallExecute(r ApiUpdateCallRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPut
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CallsAPIService.UpdateCall")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/calls/{call_id}"
	localVarPath = strings.Replace(localVarPath, "{"+"call_id"+"}", url.PathEscape(parameterValueToString(r.callId, "callId")), -1)

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
	if r.gameTitle != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "game_title", r.gameTitle, "", "")
	}
	if r.joinableBy != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "joinable_by", r.joinableBy, "", "")
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

type ApiUpdateCallUserRequest struct {
	ctx context.Context
	ApiService *CallsAPIService
	callId int64
	userId int64
	role *string
}

func (r ApiUpdateCallUserRequest) Role(role string) ApiUpdateCallUserRequest {
	r.role = &role
	return r
}

func (r ApiUpdateCallUserRequest) Execute() (*http.Response, error) {
	return r.ApiService.UpdateCallUserExecute(r)
}

func (r ApiUpdateCallUserRequest) ExecuteRaw() ([]byte, *http.Response, error) {
	httpResp, err := r.Execute()
	return executeRaw(httpResp, err)
}

/*
UpdateCallUser Method for UpdateCallUser

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param callId
 @param userId
 @return ApiUpdateCallUserRequest
*/
func (a *CallsAPIService) UpdateCallUser(ctx context.Context, callId int64, userId int64) ApiUpdateCallUserRequest {
	return ApiUpdateCallUserRequest{
		ApiService: a,
		ctx: ctx,
		callId: callId,
		userId: userId,
	}
}

// Execute executes the request
func (a *CallsAPIService) UpdateCallUserExecute(r ApiUpdateCallUserRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPut
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CallsAPIService.UpdateCallUser")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/calls/{call_id}/users/{user_id}"
	localVarPath = strings.Replace(localVarPath, "{"+"call_id"+"}", url.PathEscape(parameterValueToString(r.callId, "callId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"user_id"+"}", url.PathEscape(parameterValueToString(r.userId, "userId")), -1)

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
	if r.role != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "role", r.role, "", "")
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

type ApiValidateCallActionSignatureRequest struct {
	ctx context.Context
	ApiService *CallsAPIService
	callId *int64
	senderUuid *string
	receiverUuid *string
	action *string
	timestamp *int64
	signature *string
}

func (r ApiValidateCallActionSignatureRequest) CallId(callId int64) ApiValidateCallActionSignatureRequest {
	r.callId = &callId
	return r
}

func (r ApiValidateCallActionSignatureRequest) SenderUuid(senderUuid string) ApiValidateCallActionSignatureRequest {
	r.senderUuid = &senderUuid
	return r
}

func (r ApiValidateCallActionSignatureRequest) ReceiverUuid(receiverUuid string) ApiValidateCallActionSignatureRequest {
	r.receiverUuid = &receiverUuid
	return r
}

func (r ApiValidateCallActionSignatureRequest) Action(action string) ApiValidateCallActionSignatureRequest {
	r.action = &action
	return r
}

func (r ApiValidateCallActionSignatureRequest) Timestamp(timestamp int64) ApiValidateCallActionSignatureRequest {
	r.timestamp = &timestamp
	return r
}

func (r ApiValidateCallActionSignatureRequest) Signature(signature string) ApiValidateCallActionSignatureRequest {
	r.signature = &signature
	return r
}

func (r ApiValidateCallActionSignatureRequest) Execute() (*http.Response, error) {
	return r.ApiService.ValidateCallActionSignatureExecute(r)
}

func (r ApiValidateCallActionSignatureRequest) ExecuteRaw() ([]byte, *http.Response, error) {
	httpResp, err := r.Execute()
	return executeRaw(httpResp, err)
}

/*
ValidateCallActionSignature Method for ValidateCallActionSignature

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiValidateCallActionSignatureRequest
*/
func (a *CallsAPIService) ValidateCallActionSignature(ctx context.Context) ApiValidateCallActionSignatureRequest {
	return ApiValidateCallActionSignatureRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
func (a *CallsAPIService) ValidateCallActionSignatureExecute(r ApiValidateCallActionSignatureRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CallsAPIService.ValidateCallActionSignature")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v2/calls/action_signature/validate"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.callId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "call_id", r.callId, "form", "")
	}
	if r.senderUuid != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "sender_uuid", r.senderUuid, "form", "")
	}
	if r.receiverUuid != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "receiver_uuid", r.receiverUuid, "form", "")
	}
	if r.action != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "action", r.action, "form", "")
	}
	if r.timestamp != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "timestamp", r.timestamp, "form", "")
	}
	if r.signature != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "signature", r.signature, "form", "")
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
