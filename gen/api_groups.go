
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


// GroupsAPIService GroupsAPI service
type GroupsAPIService service

type ApiAcceptGroupJoinRequestRequest struct {
	ctx context.Context
	ApiService *GroupsAPIService
	id int64
	userId int64
}

func (r ApiAcceptGroupJoinRequestRequest) Execute() (*http.Response, error) {
	return r.ApiService.AcceptGroupJoinRequestExecute(r)
}

/*
AcceptGroupJoinRequest Method for AcceptGroupJoinRequest

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id
 @param userId
 @return ApiAcceptGroupJoinRequestRequest
*/
func (a *GroupsAPIService) AcceptGroupJoinRequest(ctx context.Context, id int64, userId int64) ApiAcceptGroupJoinRequestRequest {
	return ApiAcceptGroupJoinRequestRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
		userId: userId,
	}
}

// Execute executes the request
func (a *GroupsAPIService) AcceptGroupJoinRequestExecute(r ApiAcceptGroupJoinRequestRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "GroupsAPIService.AcceptGroupJoinRequest")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/groups/{id}/accept/{userId}"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterValueToString(r.id, "id")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"userId"+"}", url.PathEscape(parameterValueToString(r.userId, "userId")), -1)

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

type ApiAcceptGroupTransferRequest struct {
	ctx context.Context
	ApiService *GroupsAPIService
	id int64
}

func (r ApiAcceptGroupTransferRequest) Execute() (*http.Response, error) {
	return r.ApiService.AcceptGroupTransferExecute(r)
}

/*
AcceptGroupTransfer Method for AcceptGroupTransfer

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id
 @return ApiAcceptGroupTransferRequest
*/
func (a *GroupsAPIService) AcceptGroupTransfer(ctx context.Context, id int64) ApiAcceptGroupTransferRequest {
	return ApiAcceptGroupTransferRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
func (a *GroupsAPIService) AcceptGroupTransferExecute(r ApiAcceptGroupTransferRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPut
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "GroupsAPIService.AcceptGroupTransfer")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/groups/{id}/transfer"
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

type ApiBanGroupUserRequest struct {
	ctx context.Context
	ApiService *GroupsAPIService
	id int64
	userId int64
}

func (r ApiBanGroupUserRequest) Execute() (*http.Response, error) {
	return r.ApiService.BanGroupUserExecute(r)
}

/*
BanGroupUser Method for BanGroupUser

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id
 @param userId
 @return ApiBanGroupUserRequest
*/
func (a *GroupsAPIService) BanGroupUser(ctx context.Context, id int64, userId int64) ApiBanGroupUserRequest {
	return ApiBanGroupUserRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
		userId: userId,
	}
}

// Execute executes the request
func (a *GroupsAPIService) BanGroupUserExecute(r ApiBanGroupUserRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "GroupsAPIService.BanGroupUser")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/groups/{id}/ban/{userId}"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterValueToString(r.id, "id")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"userId"+"}", url.PathEscape(parameterValueToString(r.userId, "userId")), -1)

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

type ApiCancelGroupTransferRequest struct {
	ctx context.Context
	ApiService *GroupsAPIService
	id int64
}

func (r ApiCancelGroupTransferRequest) Execute() (*http.Response, error) {
	return r.ApiService.CancelGroupTransferExecute(r)
}

/*
CancelGroupTransfer Method for CancelGroupTransfer

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id
 @return ApiCancelGroupTransferRequest
*/
func (a *GroupsAPIService) CancelGroupTransfer(ctx context.Context, id int64) ApiCancelGroupTransferRequest {
	return ApiCancelGroupTransferRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
func (a *GroupsAPIService) CancelGroupTransferExecute(r ApiCancelGroupTransferRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "GroupsAPIService.CancelGroupTransfer")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/groups/{id}/transfer"
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

type ApiCreateGroupRequest struct {
	ctx context.Context
	ApiService *GroupsAPIService
	apiKey *string
	hideFromGameEight *bool
	signedInfo *string
	timestamp *int64
	topic *string
	uuid *string
	allowMembersToPostImageAndVideo *bool
	allowMembersToPostUrl *bool
	allowOwnershipTransfer *bool
	allowThreadCreationBy *string
	callTimelineDisplay *bool
	coverImageFilename *string
	description *string
	gender *int32
	generationGroupsLimit *int32
	groupCategoryId *int64
	groupIconFilename *string
	guidelines *string
	hideConferenceCall *bool
	hideReportedPosts *bool
	isPrivate *bool
	onlyMobileVerified *bool
	onlyVerifiedAge *bool
	secret *bool
	subCategoryId *string
}

func (r ApiCreateGroupRequest) ApiKey(apiKey string) ApiCreateGroupRequest {
	r.apiKey = &apiKey
	return r
}

func (r ApiCreateGroupRequest) HideFromGameEight(hideFromGameEight bool) ApiCreateGroupRequest {
	r.hideFromGameEight = &hideFromGameEight
	return r
}

func (r ApiCreateGroupRequest) SignedInfo(signedInfo string) ApiCreateGroupRequest {
	r.signedInfo = &signedInfo
	return r
}

func (r ApiCreateGroupRequest) Timestamp(timestamp int64) ApiCreateGroupRequest {
	r.timestamp = &timestamp
	return r
}

func (r ApiCreateGroupRequest) Topic(topic string) ApiCreateGroupRequest {
	r.topic = &topic
	return r
}

func (r ApiCreateGroupRequest) Uuid(uuid string) ApiCreateGroupRequest {
	r.uuid = &uuid
	return r
}

func (r ApiCreateGroupRequest) AllowMembersToPostImageAndVideo(allowMembersToPostImageAndVideo bool) ApiCreateGroupRequest {
	r.allowMembersToPostImageAndVideo = &allowMembersToPostImageAndVideo
	return r
}

func (r ApiCreateGroupRequest) AllowMembersToPostUrl(allowMembersToPostUrl bool) ApiCreateGroupRequest {
	r.allowMembersToPostUrl = &allowMembersToPostUrl
	return r
}

func (r ApiCreateGroupRequest) AllowOwnershipTransfer(allowOwnershipTransfer bool) ApiCreateGroupRequest {
	r.allowOwnershipTransfer = &allowOwnershipTransfer
	return r
}

func (r ApiCreateGroupRequest) AllowThreadCreationBy(allowThreadCreationBy string) ApiCreateGroupRequest {
	r.allowThreadCreationBy = &allowThreadCreationBy
	return r
}

func (r ApiCreateGroupRequest) CallTimelineDisplay(callTimelineDisplay bool) ApiCreateGroupRequest {
	r.callTimelineDisplay = &callTimelineDisplay
	return r
}

func (r ApiCreateGroupRequest) CoverImageFilename(coverImageFilename string) ApiCreateGroupRequest {
	r.coverImageFilename = &coverImageFilename
	return r
}

func (r ApiCreateGroupRequest) Description(description string) ApiCreateGroupRequest {
	r.description = &description
	return r
}

func (r ApiCreateGroupRequest) Gender(gender int32) ApiCreateGroupRequest {
	r.gender = &gender
	return r
}

func (r ApiCreateGroupRequest) GenerationGroupsLimit(generationGroupsLimit int32) ApiCreateGroupRequest {
	r.generationGroupsLimit = &generationGroupsLimit
	return r
}

func (r ApiCreateGroupRequest) GroupCategoryId(groupCategoryId int64) ApiCreateGroupRequest {
	r.groupCategoryId = &groupCategoryId
	return r
}

func (r ApiCreateGroupRequest) GroupIconFilename(groupIconFilename string) ApiCreateGroupRequest {
	r.groupIconFilename = &groupIconFilename
	return r
}

func (r ApiCreateGroupRequest) Guidelines(guidelines string) ApiCreateGroupRequest {
	r.guidelines = &guidelines
	return r
}

func (r ApiCreateGroupRequest) HideConferenceCall(hideConferenceCall bool) ApiCreateGroupRequest {
	r.hideConferenceCall = &hideConferenceCall
	return r
}

func (r ApiCreateGroupRequest) HideReportedPosts(hideReportedPosts bool) ApiCreateGroupRequest {
	r.hideReportedPosts = &hideReportedPosts
	return r
}

func (r ApiCreateGroupRequest) IsPrivate(isPrivate bool) ApiCreateGroupRequest {
	r.isPrivate = &isPrivate
	return r
}

func (r ApiCreateGroupRequest) OnlyMobileVerified(onlyMobileVerified bool) ApiCreateGroupRequest {
	r.onlyMobileVerified = &onlyMobileVerified
	return r
}

func (r ApiCreateGroupRequest) OnlyVerifiedAge(onlyVerifiedAge bool) ApiCreateGroupRequest {
	r.onlyVerifiedAge = &onlyVerifiedAge
	return r
}

func (r ApiCreateGroupRequest) Secret(secret bool) ApiCreateGroupRequest {
	r.secret = &secret
	return r
}

func (r ApiCreateGroupRequest) SubCategoryId(subCategoryId string) ApiCreateGroupRequest {
	r.subCategoryId = &subCategoryId
	return r
}

func (r ApiCreateGroupRequest) Execute() (*CreateGroupResponse, *http.Response, error) {
	return r.ApiService.CreateGroupExecute(r)
}

/*
CreateGroup Method for CreateGroup

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiCreateGroupRequest
*/
func (a *GroupsAPIService) CreateGroup(ctx context.Context) ApiCreateGroupRequest {
	return ApiCreateGroupRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return CreateGroupResponse
func (a *GroupsAPIService) CreateGroupExecute(r ApiCreateGroupRequest) (*CreateGroupResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *CreateGroupResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "GroupsAPIService.CreateGroup")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v3/groups/new"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.apiKey == nil {
		return localVarReturnValue, nil, reportError("apiKey is required and must be specified")
	}
	if r.hideFromGameEight == nil {
		return localVarReturnValue, nil, reportError("hideFromGameEight is required and must be specified")
	}
	if r.signedInfo == nil {
		return localVarReturnValue, nil, reportError("signedInfo is required and must be specified")
	}
	if r.timestamp == nil {
		return localVarReturnValue, nil, reportError("timestamp is required and must be specified")
	}
	if r.topic == nil {
		return localVarReturnValue, nil, reportError("topic is required and must be specified")
	}
	if r.uuid == nil {
		return localVarReturnValue, nil, reportError("uuid is required and must be specified")
	}

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
	if r.allowMembersToPostImageAndVideo != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "allow_members_to_post_image_and_video", r.allowMembersToPostImageAndVideo, "", "")
	}
	if r.allowMembersToPostUrl != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "allow_members_to_post_url", r.allowMembersToPostUrl, "", "")
	}
	if r.allowOwnershipTransfer != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "allow_ownership_transfer", r.allowOwnershipTransfer, "", "")
	}
	if r.allowThreadCreationBy != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "allow_thread_creation_by", r.allowThreadCreationBy, "", "")
	}
	parameterAddToHeaderOrQuery(localVarFormParams, "api_key", r.apiKey, "", "")
	if r.callTimelineDisplay != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "call_timeline_display", r.callTimelineDisplay, "", "")
	}
	if r.coverImageFilename != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "cover_image_filename", r.coverImageFilename, "", "")
	}
	if r.description != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "description", r.description, "", "")
	}
	if r.gender != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "gender", r.gender, "", "")
	}
	if r.generationGroupsLimit != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "generation_groups_limit", r.generationGroupsLimit, "", "")
	}
	if r.groupCategoryId != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "group_category_id", r.groupCategoryId, "", "")
	}
	if r.groupIconFilename != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "group_icon_filename", r.groupIconFilename, "", "")
	}
	if r.guidelines != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "guidelines", r.guidelines, "", "")
	}
	if r.hideConferenceCall != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "hide_conference_call", r.hideConferenceCall, "", "")
	}
	parameterAddToHeaderOrQuery(localVarFormParams, "hide_from_game_eight", r.hideFromGameEight, "", "")
	if r.hideReportedPosts != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "hide_reported_posts", r.hideReportedPosts, "", "")
	}
	if r.isPrivate != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "is_private", r.isPrivate, "", "")
	}
	if r.onlyMobileVerified != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "only_mobile_verified", r.onlyMobileVerified, "", "")
	}
	if r.onlyVerifiedAge != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "only_verified_age", r.onlyVerifiedAge, "", "")
	}
	if r.secret != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "secret", r.secret, "", "")
	}
	parameterAddToHeaderOrQuery(localVarFormParams, "signed_info", r.signedInfo, "", "")
	if r.subCategoryId != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "sub_category_id", r.subCategoryId, "", "")
	}
	parameterAddToHeaderOrQuery(localVarFormParams, "timestamp", r.timestamp, "", "")
	parameterAddToHeaderOrQuery(localVarFormParams, "topic", r.topic, "", "")
	parameterAddToHeaderOrQuery(localVarFormParams, "uuid", r.uuid, "", "")
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

type ApiDeclineGroupJoinRequestRequest struct {
	ctx context.Context
	ApiService *GroupsAPIService
	id int64
	userId int64
}

func (r ApiDeclineGroupJoinRequestRequest) Execute() (*http.Response, error) {
	return r.ApiService.DeclineGroupJoinRequestExecute(r)
}

/*
DeclineGroupJoinRequest Method for DeclineGroupJoinRequest

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id
 @param userId
 @return ApiDeclineGroupJoinRequestRequest
*/
func (a *GroupsAPIService) DeclineGroupJoinRequest(ctx context.Context, id int64, userId int64) ApiDeclineGroupJoinRequestRequest {
	return ApiDeclineGroupJoinRequestRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
		userId: userId,
	}
}

// Execute executes the request
func (a *GroupsAPIService) DeclineGroupJoinRequestExecute(r ApiDeclineGroupJoinRequestRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "GroupsAPIService.DeclineGroupJoinRequest")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/groups/{id}/decline/{userId}"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterValueToString(r.id, "id")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"userId"+"}", url.PathEscape(parameterValueToString(r.userId, "userId")), -1)

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

type ApiDeputizeGroupUsersRequest struct {
	ctx context.Context
	ApiService *GroupsAPIService
	id int64
}

func (r ApiDeputizeGroupUsersRequest) Execute() (*http.Response, error) {
	return r.ApiService.DeputizeGroupUsersExecute(r)
}

/*
DeputizeGroupUsers Method for DeputizeGroupUsers

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id
 @return ApiDeputizeGroupUsersRequest
*/
func (a *GroupsAPIService) DeputizeGroupUsers(ctx context.Context, id int64) ApiDeputizeGroupUsersRequest {
	return ApiDeputizeGroupUsersRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
func (a *GroupsAPIService) DeputizeGroupUsersExecute(r ApiDeputizeGroupUsersRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPut
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "GroupsAPIService.DeputizeGroupUsers")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/groups/{id}/deputize"
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

type ApiDeputizeGroupUsersMassRequest struct {
	ctx context.Context
	ApiService *GroupsAPIService
	groupId int64
	apiKey *string
	signedInfo *string
	timestamp *int64
	userIds *[]int64
	uuid *string
}

func (r ApiDeputizeGroupUsersMassRequest) ApiKey(apiKey string) ApiDeputizeGroupUsersMassRequest {
	r.apiKey = &apiKey
	return r
}

func (r ApiDeputizeGroupUsersMassRequest) SignedInfo(signedInfo string) ApiDeputizeGroupUsersMassRequest {
	r.signedInfo = &signedInfo
	return r
}

func (r ApiDeputizeGroupUsersMassRequest) Timestamp(timestamp int64) ApiDeputizeGroupUsersMassRequest {
	r.timestamp = &timestamp
	return r
}

func (r ApiDeputizeGroupUsersMassRequest) UserIds(userIds []int64) ApiDeputizeGroupUsersMassRequest {
	r.userIds = &userIds
	return r
}

func (r ApiDeputizeGroupUsersMassRequest) Uuid(uuid string) ApiDeputizeGroupUsersMassRequest {
	r.uuid = &uuid
	return r
}

func (r ApiDeputizeGroupUsersMassRequest) Execute() (*http.Response, error) {
	return r.ApiService.DeputizeGroupUsersMassExecute(r)
}

/*
DeputizeGroupUsersMass Method for DeputizeGroupUsersMass

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param groupId
 @return ApiDeputizeGroupUsersMassRequest
*/
func (a *GroupsAPIService) DeputizeGroupUsersMass(ctx context.Context, groupId int64) ApiDeputizeGroupUsersMassRequest {
	return ApiDeputizeGroupUsersMassRequest{
		ApiService: a,
		ctx: ctx,
		groupId: groupId,
	}
}

// Execute executes the request
func (a *GroupsAPIService) DeputizeGroupUsersMassExecute(r ApiDeputizeGroupUsersMassRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "GroupsAPIService.DeputizeGroupUsersMass")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v3/groups/{group_id}/deputize/mass"
	localVarPath = strings.Replace(localVarPath, "{"+"group_id"+"}", url.PathEscape(parameterValueToString(r.groupId, "groupId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.apiKey == nil {
		return nil, reportError("apiKey is required and must be specified")
	}
	if r.signedInfo == nil {
		return nil, reportError("signedInfo is required and must be specified")
	}
	if r.timestamp == nil {
		return nil, reportError("timestamp is required and must be specified")
	}
	if r.userIds == nil {
		return nil, reportError("userIds is required and must be specified")
	}
	if r.uuid == nil {
		return nil, reportError("uuid is required and must be specified")
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
	parameterAddToHeaderOrQuery(localVarFormParams, "api_key", r.apiKey, "", "")
	parameterAddToHeaderOrQuery(localVarFormParams, "signed_info", r.signedInfo, "", "")
	parameterAddToHeaderOrQuery(localVarFormParams, "timestamp", r.timestamp, "", "")
	parameterAddToHeaderOrQuery(localVarFormParams, "user_ids[]", r.userIds, "", "csv")
	parameterAddToHeaderOrQuery(localVarFormParams, "uuid", r.uuid, "", "")
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

type ApiFireGroupUserRequest struct {
	ctx context.Context
	ApiService *GroupsAPIService
	groupId int64
	userId int64
}

func (r ApiFireGroupUserRequest) Execute() (*http.Response, error) {
	return r.ApiService.FireGroupUserExecute(r)
}

/*
FireGroupUser Method for FireGroupUser

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param groupId
 @param userId
 @return ApiFireGroupUserRequest
*/
func (a *GroupsAPIService) FireGroupUser(ctx context.Context, groupId int64, userId int64) ApiFireGroupUserRequest {
	return ApiFireGroupUserRequest{
		ApiService: a,
		ctx: ctx,
		groupId: groupId,
		userId: userId,
	}
}

// Execute executes the request
func (a *GroupsAPIService) FireGroupUserExecute(r ApiFireGroupUserRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "GroupsAPIService.FireGroupUser")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/groups/{group_id}/fire/{user_id}"
	localVarPath = strings.Replace(localVarPath, "{"+"group_id"+"}", url.PathEscape(parameterValueToString(r.groupId, "groupId")), -1)
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

type ApiGetGroupRequest struct {
	ctx context.Context
	ApiService *GroupsAPIService
	id int64
}

func (r ApiGetGroupRequest) Execute() (*GroupResponse, *http.Response, error) {
	return r.ApiService.GetGroupExecute(r)
}

/*
GetGroup Method for GetGroup

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id
 @return ApiGetGroupRequest
*/
func (a *GroupsAPIService) GetGroup(ctx context.Context, id int64) ApiGetGroupRequest {
	return ApiGetGroupRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
//  @return GroupResponse
func (a *GroupsAPIService) GetGroupExecute(r ApiGetGroupRequest) (*GroupResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *GroupResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "GroupsAPIService.GetGroup")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/groups/{id}"
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

type ApiGetGroupBanListRequest struct {
	ctx context.Context
	ApiService *GroupsAPIService
	id int64
	page *int32
}

func (r ApiGetGroupBanListRequest) Page(page int32) ApiGetGroupBanListRequest {
	r.page = &page
	return r
}

func (r ApiGetGroupBanListRequest) Execute() (*UsersResponse, *http.Response, error) {
	return r.ApiService.GetGroupBanListExecute(r)
}

/*
GetGroupBanList Method for GetGroupBanList

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id
 @return ApiGetGroupBanListRequest
*/
func (a *GroupsAPIService) GetGroupBanList(ctx context.Context, id int64) ApiGetGroupBanListRequest {
	return ApiGetGroupBanListRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
//  @return UsersResponse
func (a *GroupsAPIService) GetGroupBanListExecute(r ApiGetGroupBanListRequest) (*UsersResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *UsersResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "GroupsAPIService.GetGroupBanList")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/groups/{id}/ban_list"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterValueToString(r.id, "id")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.page == nil {
		return localVarReturnValue, nil, reportError("page is required and must be specified")
	}

	parameterAddToHeaderOrQuery(localVarQueryParams, "page", r.page, "form", "")
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

type ApiGetGroupCreateQuotaRequest struct {
	ctx context.Context
	ApiService *GroupsAPIService
}

func (r ApiGetGroupCreateQuotaRequest) Execute() (*CreateQuotaResponse, *http.Response, error) {
	return r.ApiService.GetGroupCreateQuotaExecute(r)
}

/*
GetGroupCreateQuota Method for GetGroupCreateQuota

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiGetGroupCreateQuotaRequest
*/
func (a *GroupsAPIService) GetGroupCreateQuota(ctx context.Context) ApiGetGroupCreateQuotaRequest {
	return ApiGetGroupCreateQuotaRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return CreateQuotaResponse
func (a *GroupsAPIService) GetGroupCreateQuotaExecute(r ApiGetGroupCreateQuotaRequest) (*CreateQuotaResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *CreateQuotaResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "GroupsAPIService.GetGroupCreateQuota")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/groups/created_quota"

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

type ApiGetGroupGiftHistoryRequest struct {
	ctx context.Context
	ApiService *GroupsAPIService
	groupId int64
	number *int32
	from *string
}

func (r ApiGetGroupGiftHistoryRequest) Number(number int32) ApiGetGroupGiftHistoryRequest {
	r.number = &number
	return r
}

func (r ApiGetGroupGiftHistoryRequest) From(from string) ApiGetGroupGiftHistoryRequest {
	r.from = &from
	return r
}

func (r ApiGetGroupGiftHistoryRequest) Execute() (*GroupGiftHistoryResponse, *http.Response, error) {
	return r.ApiService.GetGroupGiftHistoryExecute(r)
}

/*
GetGroupGiftHistory Method for GetGroupGiftHistory

Implemented as a Kotlin suspend function on Android.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param groupId
 @return ApiGetGroupGiftHistoryRequest
*/
func (a *GroupsAPIService) GetGroupGiftHistory(ctx context.Context, groupId int64) ApiGetGroupGiftHistoryRequest {
	return ApiGetGroupGiftHistoryRequest{
		ApiService: a,
		ctx: ctx,
		groupId: groupId,
	}
}

// Execute executes the request
//  @return GroupGiftHistoryResponse
func (a *GroupsAPIService) GetGroupGiftHistoryExecute(r ApiGetGroupGiftHistoryRequest) (*GroupGiftHistoryResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *GroupGiftHistoryResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "GroupsAPIService.GetGroupGiftHistory")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/groups/{group_id}/gift_history"
	localVarPath = strings.Replace(localVarPath, "{"+"group_id"+"}", url.PathEscape(parameterValueToString(r.groupId, "groupId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.number != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "number", r.number, "form", "")
	}
	if r.from != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "from", r.from, "form", "")
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

type ApiGetGroupGiftTransactionsRequest struct {
	ctx context.Context
	ApiService *GroupsAPIService
	groupId int64
	number *int32
	from *string
}

func (r ApiGetGroupGiftTransactionsRequest) Number(number int32) ApiGetGroupGiftTransactionsRequest {
	r.number = &number
	return r
}

func (r ApiGetGroupGiftTransactionsRequest) From(from string) ApiGetGroupGiftTransactionsRequest {
	r.from = &from
	return r
}

func (r ApiGetGroupGiftTransactionsRequest) Execute() (*GiftTransactionsResponse, *http.Response, error) {
	return r.ApiService.GetGroupGiftTransactionsExecute(r)
}

/*
GetGroupGiftTransactions Method for GetGroupGiftTransactions

Implemented as a Kotlin suspend function on Android.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param groupId
 @return ApiGetGroupGiftTransactionsRequest
*/
func (a *GroupsAPIService) GetGroupGiftTransactions(ctx context.Context, groupId int64) ApiGetGroupGiftTransactionsRequest {
	return ApiGetGroupGiftTransactionsRequest{
		ApiService: a,
		ctx: ctx,
		groupId: groupId,
	}
}

// Execute executes the request
//  @return GiftTransactionsResponse
func (a *GroupsAPIService) GetGroupGiftTransactionsExecute(r ApiGetGroupGiftTransactionsRequest) (*GiftTransactionsResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *GiftTransactionsResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "GroupsAPIService.GetGroupGiftTransactions")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/groups/{group_id}/gift_transactions"
	localVarPath = strings.Replace(localVarPath, "{"+"group_id"+"}", url.PathEscape(parameterValueToString(r.groupId, "groupId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.number != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "number", r.number, "form", "")
	}
	if r.from != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "from", r.from, "form", "")
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

type ApiGetGroupHighlightsRequest struct {
	ctx context.Context
	ApiService *GroupsAPIService
	groupId int64
	from *int64
	number *int32
}

func (r ApiGetGroupHighlightsRequest) From(from int64) ApiGetGroupHighlightsRequest {
	r.from = &from
	return r
}

func (r ApiGetGroupHighlightsRequest) Number(number int32) ApiGetGroupHighlightsRequest {
	r.number = &number
	return r
}

func (r ApiGetGroupHighlightsRequest) Execute() (*PostsResponse, *http.Response, error) {
	return r.ApiService.GetGroupHighlightsExecute(r)
}

/*
GetGroupHighlights Method for GetGroupHighlights

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param groupId
 @return ApiGetGroupHighlightsRequest
*/
func (a *GroupsAPIService) GetGroupHighlights(ctx context.Context, groupId int64) ApiGetGroupHighlightsRequest {
	return ApiGetGroupHighlightsRequest{
		ApiService: a,
		ctx: ctx,
		groupId: groupId,
	}
}

// Execute executes the request
//  @return PostsResponse
func (a *GroupsAPIService) GetGroupHighlightsExecute(r ApiGetGroupHighlightsRequest) (*PostsResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *PostsResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "GroupsAPIService.GetGroupHighlights")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/groups/{group_id}/highlights"
	localVarPath = strings.Replace(localVarPath, "{"+"group_id"+"}", url.PathEscape(parameterValueToString(r.groupId, "groupId")), -1)

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

type ApiGetGroupMemberRequest struct {
	ctx context.Context
	ApiService *GroupsAPIService
	id int64
	userId int64
}

func (r ApiGetGroupMemberRequest) Execute() (*GroupUserResponse, *http.Response, error) {
	return r.ApiService.GetGroupMemberExecute(r)
}

/*
GetGroupMember Method for GetGroupMember

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id
 @param userId
 @return ApiGetGroupMemberRequest
*/
func (a *GroupsAPIService) GetGroupMember(ctx context.Context, id int64, userId int64) ApiGetGroupMemberRequest {
	return ApiGetGroupMemberRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
		userId: userId,
	}
}

// Execute executes the request
//  @return GroupUserResponse
func (a *GroupsAPIService) GetGroupMemberExecute(r ApiGetGroupMemberRequest) (*GroupUserResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *GroupUserResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "GroupsAPIService.GetGroupMember")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/groups/{id}/members/{userId}"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterValueToString(r.id, "id")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"userId"+"}", url.PathEscape(parameterValueToString(r.userId, "userId")), -1)

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

type ApiGetGroupMembersRequest struct {
	ctx context.Context
	ApiService *GroupsAPIService
	id int64
	number *int32
	fromId *int64
}

func (r ApiGetGroupMembersRequest) Number(number int32) ApiGetGroupMembersRequest {
	r.number = &number
	return r
}

func (r ApiGetGroupMembersRequest) FromId(fromId int64) ApiGetGroupMembersRequest {
	r.fromId = &fromId
	return r
}

func (r ApiGetGroupMembersRequest) Execute() (*GroupUsersResponse, *http.Response, error) {
	return r.ApiService.GetGroupMembersExecute(r)
}

/*
GetGroupMembers Method for GetGroupMembers

Implemented as a Kotlin suspend function on Android.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id
 @return ApiGetGroupMembersRequest
*/
func (a *GroupsAPIService) GetGroupMembers(ctx context.Context, id int64) ApiGetGroupMembersRequest {
	return ApiGetGroupMembersRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
//  @return GroupUsersResponse
func (a *GroupsAPIService) GetGroupMembersExecute(r ApiGetGroupMembersRequest) (*GroupUsersResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *GroupUsersResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "GroupsAPIService.GetGroupMembers")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v2/groups/{id}/members"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterValueToString(r.id, "id")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.number != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "number", r.number, "form", "")
	}
	if r.fromId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "from_id", r.fromId, "form", "")
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

type ApiGetGroupReceivedGiftSendersRequest struct {
	ctx context.Context
	ApiService *GroupsAPIService
	groupId int64
	giftId int64
	userId *int64
	timestamp *int64
	number *int32
}

func (r ApiGetGroupReceivedGiftSendersRequest) UserId(userId int64) ApiGetGroupReceivedGiftSendersRequest {
	r.userId = &userId
	return r
}

func (r ApiGetGroupReceivedGiftSendersRequest) Timestamp(timestamp int64) ApiGetGroupReceivedGiftSendersRequest {
	r.timestamp = &timestamp
	return r
}

func (r ApiGetGroupReceivedGiftSendersRequest) Number(number int32) ApiGetGroupReceivedGiftSendersRequest {
	r.number = &number
	return r
}

func (r ApiGetGroupReceivedGiftSendersRequest) Execute() (*GiftSendersResponse, *http.Response, error) {
	return r.ApiService.GetGroupReceivedGiftSendersExecute(r)
}

/*
GetGroupReceivedGiftSenders Method for GetGroupReceivedGiftSenders

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param groupId
 @param giftId
 @return ApiGetGroupReceivedGiftSendersRequest
*/
func (a *GroupsAPIService) GetGroupReceivedGiftSenders(ctx context.Context, groupId int64, giftId int64) ApiGetGroupReceivedGiftSendersRequest {
	return ApiGetGroupReceivedGiftSendersRequest{
		ApiService: a,
		ctx: ctx,
		groupId: groupId,
		giftId: giftId,
	}
}

// Execute executes the request
//  @return GiftSendersResponse
func (a *GroupsAPIService) GetGroupReceivedGiftSendersExecute(r ApiGetGroupReceivedGiftSendersRequest) (*GiftSendersResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *GiftSendersResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "GroupsAPIService.GetGroupReceivedGiftSenders")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/groups/{group_id}/received_gifts/{gift_id}/senders"
	localVarPath = strings.Replace(localVarPath, "{"+"group_id"+"}", url.PathEscape(parameterValueToString(r.groupId, "groupId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"gift_id"+"}", url.PathEscape(parameterValueToString(r.giftId, "giftId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.userId == nil {
		return localVarReturnValue, nil, reportError("userId is required and must be specified")
	}
	if r.timestamp == nil {
		return localVarReturnValue, nil, reportError("timestamp is required and must be specified")
	}

	parameterAddToHeaderOrQuery(localVarQueryParams, "user_id", r.userId, "form", "")
	parameterAddToHeaderOrQuery(localVarQueryParams, "timestamp", r.timestamp, "form", "")
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

type ApiGetGroupUnreadStatusRequest struct {
	ctx context.Context
	ApiService *GroupsAPIService
	fromTime *int64
}

func (r ApiGetGroupUnreadStatusRequest) FromTime(fromTime int64) ApiGetGroupUnreadStatusRequest {
	r.fromTime = &fromTime
	return r
}

func (r ApiGetGroupUnreadStatusRequest) Execute() (*UnreadStatusResponse, *http.Response, error) {
	return r.ApiService.GetGroupUnreadStatusExecute(r)
}

/*
GetGroupUnreadStatus Method for GetGroupUnreadStatus

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiGetGroupUnreadStatusRequest
*/
func (a *GroupsAPIService) GetGroupUnreadStatus(ctx context.Context) ApiGetGroupUnreadStatusRequest {
	return ApiGetGroupUnreadStatusRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return UnreadStatusResponse
func (a *GroupsAPIService) GetGroupUnreadStatusExecute(r ApiGetGroupUnreadStatusRequest) (*UnreadStatusResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *UnreadStatusResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "GroupsAPIService.GetGroupUnreadStatus")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/groups/unread_status"

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

type ApiGetInvitableGroupUsersRequest struct {
	ctx context.Context
	ApiService *GroupsAPIService
	groupId int64
	fromTimestamp *int64
	userNickname *string
}

func (r ApiGetInvitableGroupUsersRequest) FromTimestamp(fromTimestamp int64) ApiGetInvitableGroupUsersRequest {
	r.fromTimestamp = &fromTimestamp
	return r
}

func (r ApiGetInvitableGroupUsersRequest) UserNickname(userNickname string) ApiGetInvitableGroupUsersRequest {
	r.userNickname = &userNickname
	return r
}

func (r ApiGetInvitableGroupUsersRequest) Execute() (*UsersByTimestampResponse, *http.Response, error) {
	return r.ApiService.GetInvitableGroupUsersExecute(r)
}

/*
GetInvitableGroupUsers Method for GetInvitableGroupUsers

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param groupId
 @return ApiGetInvitableGroupUsersRequest
*/
func (a *GroupsAPIService) GetInvitableGroupUsers(ctx context.Context, groupId int64) ApiGetInvitableGroupUsersRequest {
	return ApiGetInvitableGroupUsersRequest{
		ApiService: a,
		ctx: ctx,
		groupId: groupId,
	}
}

// Execute executes the request
//  @return UsersByTimestampResponse
func (a *GroupsAPIService) GetInvitableGroupUsersExecute(r ApiGetInvitableGroupUsersRequest) (*UsersByTimestampResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *UsersByTimestampResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "GroupsAPIService.GetInvitableGroupUsers")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/groups/{group_id}/users/invitable"
	localVarPath = strings.Replace(localVarPath, "{"+"group_id"+"}", url.PathEscape(parameterValueToString(r.groupId, "groupId")), -1)

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

type ApiGetJoinedGroupStatusesRequest struct {
	ctx context.Context
	ApiService *GroupsAPIService
	ids *[]int64
}

func (r ApiGetJoinedGroupStatusesRequest) Ids(ids []int64) ApiGetJoinedGroupStatusesRequest {
	r.ids = &ids
	return r
}

func (r ApiGetJoinedGroupStatusesRequest) Execute() (map[string]string, *http.Response, error) {
	return r.ApiService.GetJoinedGroupStatusesExecute(r)
}

/*
GetJoinedGroupStatuses Method for GetJoinedGroupStatuses

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiGetJoinedGroupStatusesRequest
*/
func (a *GroupsAPIService) GetJoinedGroupStatuses(ctx context.Context) ApiGetJoinedGroupStatusesRequest {
	return ApiGetJoinedGroupStatusesRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return map[string]string
func (a *GroupsAPIService) GetJoinedGroupStatusesExecute(r ApiGetJoinedGroupStatusesRequest) (map[string]string, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  map[string]string
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "GroupsAPIService.GetJoinedGroupStatuses")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/groups/joined_statuses"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.ids == nil {
		return localVarReturnValue, nil, reportError("ids is required and must be specified")
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

type ApiGetRelatableGroupsRequest struct {
	ctx context.Context
	ApiService *GroupsAPIService
	id int64
	keyword *string
	from *string
}

func (r ApiGetRelatableGroupsRequest) Keyword(keyword string) ApiGetRelatableGroupsRequest {
	r.keyword = &keyword
	return r
}

func (r ApiGetRelatableGroupsRequest) From(from string) ApiGetRelatableGroupsRequest {
	r.from = &from
	return r
}

func (r ApiGetRelatableGroupsRequest) Execute() (*GroupsRelatedResponse, *http.Response, error) {
	return r.ApiService.GetRelatableGroupsExecute(r)
}

/*
GetRelatableGroups Method for GetRelatableGroups

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id
 @return ApiGetRelatableGroupsRequest
*/
func (a *GroupsAPIService) GetRelatableGroups(ctx context.Context, id int64) ApiGetRelatableGroupsRequest {
	return ApiGetRelatableGroupsRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
//  @return GroupsRelatedResponse
func (a *GroupsAPIService) GetRelatableGroupsExecute(r ApiGetRelatableGroupsRequest) (*GroupsRelatedResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *GroupsRelatedResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "GroupsAPIService.GetRelatableGroups")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/groups/{id}/relatable"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterValueToString(r.id, "id")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.keyword != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "keyword", r.keyword, "form", "")
	}
	if r.from != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "from", r.from, "form", "")
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

type ApiGetRelatedGroupsRequest struct {
	ctx context.Context
	ApiService *GroupsAPIService
	id int64
	keyword *string
	from *string
}

func (r ApiGetRelatedGroupsRequest) Keyword(keyword string) ApiGetRelatedGroupsRequest {
	r.keyword = &keyword
	return r
}

func (r ApiGetRelatedGroupsRequest) From(from string) ApiGetRelatedGroupsRequest {
	r.from = &from
	return r
}

func (r ApiGetRelatedGroupsRequest) Execute() (*GroupsRelatedResponse, *http.Response, error) {
	return r.ApiService.GetRelatedGroupsExecute(r)
}

/*
GetRelatedGroups Method for GetRelatedGroups

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id
 @return ApiGetRelatedGroupsRequest
*/
func (a *GroupsAPIService) GetRelatedGroups(ctx context.Context, id int64) ApiGetRelatedGroupsRequest {
	return ApiGetRelatedGroupsRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
//  @return GroupsRelatedResponse
func (a *GroupsAPIService) GetRelatedGroupsExecute(r ApiGetRelatedGroupsRequest) (*GroupsRelatedResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *GroupsRelatedResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "GroupsAPIService.GetRelatedGroups")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/groups/{id}/related"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterValueToString(r.id, "id")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.keyword != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "keyword", r.keyword, "form", "")
	}
	if r.from != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "from", r.from, "form", "")
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

type ApiGetUserGroupListRequest struct {
	ctx context.Context
	ApiService *GroupsAPIService
	page *int32
	userId *int64
}

func (r ApiGetUserGroupListRequest) Page(page int32) ApiGetUserGroupListRequest {
	r.page = &page
	return r
}

func (r ApiGetUserGroupListRequest) UserId(userId int64) ApiGetUserGroupListRequest {
	r.userId = &userId
	return r
}

func (r ApiGetUserGroupListRequest) Execute() (*GroupsResponse, *http.Response, error) {
	return r.ApiService.GetUserGroupListExecute(r)
}

/*
GetUserGroupList Method for GetUserGroupList

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiGetUserGroupListRequest
*/
func (a *GroupsAPIService) GetUserGroupList(ctx context.Context) ApiGetUserGroupListRequest {
	return ApiGetUserGroupListRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return GroupsResponse
func (a *GroupsAPIService) GetUserGroupListExecute(r ApiGetUserGroupListRequest) (*GroupsResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *GroupsResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "GroupsAPIService.GetUserGroupList")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/groups/user_group_list"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.page == nil {
		return localVarReturnValue, nil, reportError("page is required and must be specified")
	}
	if r.userId == nil {
		return localVarReturnValue, nil, reportError("userId is required and must be specified")
	}

	parameterAddToHeaderOrQuery(localVarQueryParams, "page", r.page, "form", "")
	parameterAddToHeaderOrQuery(localVarQueryParams, "user_id", r.userId, "form", "")
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

type ApiInviteToGroupRequest struct {
	ctx context.Context
	ApiService *GroupsAPIService
	id int64
	userIds *[]int64
}

func (r ApiInviteToGroupRequest) UserIds(userIds []int64) ApiInviteToGroupRequest {
	r.userIds = &userIds
	return r
}

func (r ApiInviteToGroupRequest) Execute() (*http.Response, error) {
	return r.ApiService.InviteToGroupExecute(r)
}

/*
InviteToGroup Method for InviteToGroup

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id
 @return ApiInviteToGroupRequest
*/
func (a *GroupsAPIService) InviteToGroup(ctx context.Context, id int64) ApiInviteToGroupRequest {
	return ApiInviteToGroupRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
func (a *GroupsAPIService) InviteToGroupExecute(r ApiInviteToGroupRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "GroupsAPIService.InviteToGroup")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/groups/{id}/invite"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterValueToString(r.id, "id")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.userIds == nil {
		return nil, reportError("userIds is required and must be specified")
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
	parameterAddToHeaderOrQuery(localVarFormParams, "user_ids[]", r.userIds, "", "csv")
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

type ApiJoinGroupRequest struct {
	ctx context.Context
	ApiService *GroupsAPIService
	id int64
}

func (r ApiJoinGroupRequest) Execute() (*http.Response, error) {
	return r.ApiService.JoinGroupExecute(r)
}

/*
JoinGroup Method for JoinGroup

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id
 @return ApiJoinGroupRequest
*/
func (a *GroupsAPIService) JoinGroup(ctx context.Context, id int64) ApiJoinGroupRequest {
	return ApiJoinGroupRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
func (a *GroupsAPIService) JoinGroupExecute(r ApiJoinGroupRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "GroupsAPIService.JoinGroup")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/groups/{id}/join"
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

type ApiLeaveGroupRequest struct {
	ctx context.Context
	ApiService *GroupsAPIService
	id int64
}

func (r ApiLeaveGroupRequest) Execute() (*http.Response, error) {
	return r.ApiService.LeaveGroupExecute(r)
}

/*
LeaveGroup Method for LeaveGroup

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id
 @return ApiLeaveGroupRequest
*/
func (a *GroupsAPIService) LeaveGroup(ctx context.Context, id int64) ApiLeaveGroupRequest {
	return ApiLeaveGroupRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
func (a *GroupsAPIService) LeaveGroupExecute(r ApiLeaveGroupRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "GroupsAPIService.LeaveGroup")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/groups/{id}/leave"
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

type ApiListGroupCategoriesRequest struct {
	ctx context.Context
	ApiService *GroupsAPIService
	page *int32
	number *int32
}

func (r ApiListGroupCategoriesRequest) Page(page int32) ApiListGroupCategoriesRequest {
	r.page = &page
	return r
}

func (r ApiListGroupCategoriesRequest) Number(number int32) ApiListGroupCategoriesRequest {
	r.number = &number
	return r
}

func (r ApiListGroupCategoriesRequest) Execute() (*GroupCategoriesResponse, *http.Response, error) {
	return r.ApiService.ListGroupCategoriesExecute(r)
}

/*
ListGroupCategories Method for ListGroupCategories

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiListGroupCategoriesRequest
*/
func (a *GroupsAPIService) ListGroupCategories(ctx context.Context) ApiListGroupCategoriesRequest {
	return ApiListGroupCategoriesRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return GroupCategoriesResponse
func (a *GroupsAPIService) ListGroupCategoriesExecute(r ApiListGroupCategoriesRequest) (*GroupCategoriesResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *GroupCategoriesResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "GroupsAPIService.ListGroupCategories")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/groups/categories"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.page == nil {
		return localVarReturnValue, nil, reportError("page is required and must be specified")
	}

	parameterAddToHeaderOrQuery(localVarQueryParams, "page", r.page, "form", "")
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

type ApiListGroupsRequest struct {
	ctx context.Context
	ApiService *GroupsAPIService
	groupCategoryId *int64
	keyword *string
	fromTimestamp *int64
	subCategoryId *int64
}

func (r ApiListGroupsRequest) GroupCategoryId(groupCategoryId int64) ApiListGroupsRequest {
	r.groupCategoryId = &groupCategoryId
	return r
}

func (r ApiListGroupsRequest) Keyword(keyword string) ApiListGroupsRequest {
	r.keyword = &keyword
	return r
}

func (r ApiListGroupsRequest) FromTimestamp(fromTimestamp int64) ApiListGroupsRequest {
	r.fromTimestamp = &fromTimestamp
	return r
}

func (r ApiListGroupsRequest) SubCategoryId(subCategoryId int64) ApiListGroupsRequest {
	r.subCategoryId = &subCategoryId
	return r
}

func (r ApiListGroupsRequest) Execute() (*GroupsResponse, *http.Response, error) {
	return r.ApiService.ListGroupsExecute(r)
}

/*
ListGroups Method for ListGroups

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiListGroupsRequest
*/
func (a *GroupsAPIService) ListGroups(ctx context.Context) ApiListGroupsRequest {
	return ApiListGroupsRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return GroupsResponse
func (a *GroupsAPIService) ListGroupsExecute(r ApiListGroupsRequest) (*GroupsResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *GroupsResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "GroupsAPIService.ListGroups")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v2/groups"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.groupCategoryId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "group_category_id", r.groupCategoryId, "form", "")
	}
	if r.keyword != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "keyword", r.keyword, "form", "")
	}
	if r.fromTimestamp != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "from_timestamp", r.fromTimestamp, "form", "")
	}
	if r.subCategoryId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "sub_category_id", r.subCategoryId, "form", "")
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

type ApiListMyGroupsRequest struct {
	ctx context.Context
	ApiService *GroupsAPIService
	fromTimestamp *int64
}

func (r ApiListMyGroupsRequest) FromTimestamp(fromTimestamp int64) ApiListMyGroupsRequest {
	r.fromTimestamp = &fromTimestamp
	return r
}

func (r ApiListMyGroupsRequest) Execute() (*GroupsResponse, *http.Response, error) {
	return r.ApiService.ListMyGroupsExecute(r)
}

/*
ListMyGroups Method for ListMyGroups

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiListMyGroupsRequest
*/
func (a *GroupsAPIService) ListMyGroups(ctx context.Context) ApiListMyGroupsRequest {
	return ApiListMyGroupsRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return GroupsResponse
func (a *GroupsAPIService) ListMyGroupsExecute(r ApiListMyGroupsRequest) (*GroupsResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *GroupsResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "GroupsAPIService.ListMyGroups")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v2/groups/mine"

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

type ApiPinGroupHighlightPostRequest struct {
	ctx context.Context
	ApiService *GroupsAPIService
	groupId int64
	postId int64
}

func (r ApiPinGroupHighlightPostRequest) Execute() (*http.Response, error) {
	return r.ApiService.PinGroupHighlightPostExecute(r)
}

/*
PinGroupHighlightPost Method for PinGroupHighlightPost

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param groupId
 @param postId
 @return ApiPinGroupHighlightPostRequest
*/
func (a *GroupsAPIService) PinGroupHighlightPost(ctx context.Context, groupId int64, postId int64) ApiPinGroupHighlightPostRequest {
	return ApiPinGroupHighlightPostRequest{
		ApiService: a,
		ctx: ctx,
		groupId: groupId,
		postId: postId,
	}
}

// Execute executes the request
func (a *GroupsAPIService) PinGroupHighlightPostExecute(r ApiPinGroupHighlightPostRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPut
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "GroupsAPIService.PinGroupHighlightPost")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/groups/{group_id}/highlights/{post_id}"
	localVarPath = strings.Replace(localVarPath, "{"+"group_id"+"}", url.PathEscape(parameterValueToString(r.groupId, "groupId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"post_id"+"}", url.PathEscape(parameterValueToString(r.postId, "postId")), -1)

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

type ApiRemoveGroupCoverRequest struct {
	ctx context.Context
	ApiService *GroupsAPIService
	id int64
}

func (r ApiRemoveGroupCoverRequest) Execute() (*http.Response, error) {
	return r.ApiService.RemoveGroupCoverExecute(r)
}

/*
RemoveGroupCover Method for RemoveGroupCover

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id
 @return ApiRemoveGroupCoverRequest
*/
func (a *GroupsAPIService) RemoveGroupCover(ctx context.Context, id int64) ApiRemoveGroupCoverRequest {
	return ApiRemoveGroupCoverRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
func (a *GroupsAPIService) RemoveGroupCoverExecute(r ApiRemoveGroupCoverRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "GroupsAPIService.RemoveGroupCover")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v3/groups/{id}/cover"
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

type ApiRemoveGroupDeputiesRequest struct {
	ctx context.Context
	ApiService *GroupsAPIService
	id int64
}

func (r ApiRemoveGroupDeputiesRequest) Execute() (*http.Response, error) {
	return r.ApiService.RemoveGroupDeputiesExecute(r)
}

/*
RemoveGroupDeputies Method for RemoveGroupDeputies

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id
 @return ApiRemoveGroupDeputiesRequest
*/
func (a *GroupsAPIService) RemoveGroupDeputies(ctx context.Context, id int64) ApiRemoveGroupDeputiesRequest {
	return ApiRemoveGroupDeputiesRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
func (a *GroupsAPIService) RemoveGroupDeputiesExecute(r ApiRemoveGroupDeputiesRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "GroupsAPIService.RemoveGroupDeputies")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/groups/{id}/deputize"
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

type ApiRemoveGroupIconRequest struct {
	ctx context.Context
	ApiService *GroupsAPIService
	id int64
}

func (r ApiRemoveGroupIconRequest) Execute() (*http.Response, error) {
	return r.ApiService.RemoveGroupIconExecute(r)
}

/*
RemoveGroupIcon Method for RemoveGroupIcon

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id
 @return ApiRemoveGroupIconRequest
*/
func (a *GroupsAPIService) RemoveGroupIcon(ctx context.Context, id int64) ApiRemoveGroupIconRequest {
	return ApiRemoveGroupIconRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
func (a *GroupsAPIService) RemoveGroupIconExecute(r ApiRemoveGroupIconRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "GroupsAPIService.RemoveGroupIcon")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v3/groups/{id}/icon"
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

type ApiRemoveRelatedGroupsRequest struct {
	ctx context.Context
	ApiService *GroupsAPIService
	id int64
	relatedGroupId *[]int64
}

func (r ApiRemoveRelatedGroupsRequest) RelatedGroupId(relatedGroupId []int64) ApiRemoveRelatedGroupsRequest {
	r.relatedGroupId = &relatedGroupId
	return r
}

func (r ApiRemoveRelatedGroupsRequest) Execute() (*http.Response, error) {
	return r.ApiService.RemoveRelatedGroupsExecute(r)
}

/*
RemoveRelatedGroups Method for RemoveRelatedGroups

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id
 @return ApiRemoveRelatedGroupsRequest
*/
func (a *GroupsAPIService) RemoveRelatedGroups(ctx context.Context, id int64) ApiRemoveRelatedGroupsRequest {
	return ApiRemoveRelatedGroupsRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
func (a *GroupsAPIService) RemoveRelatedGroupsExecute(r ApiRemoveRelatedGroupsRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "GroupsAPIService.RemoveRelatedGroups")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/groups/{id}/related"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterValueToString(r.id, "id")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.relatedGroupId == nil {
		return nil, reportError("relatedGroupId is required and must be specified")
	}

	{
		t := *r.relatedGroupId
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				parameterAddToHeaderOrQuery(localVarQueryParams, "related_group_id[]", s.Index(i).Interface(), "form", "multi")
			}
		} else {
			parameterAddToHeaderOrQuery(localVarQueryParams, "related_group_id[]", t, "form", "multi")
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

type ApiReportGroupRequest struct {
	ctx context.Context
	ApiService *GroupsAPIService
	groupId int64
	categoryId *int64
	opponentId *int64
	reason *string
	screenshot2Filename *string
	screenshot3Filename *string
	screenshot4Filename *string
	screenshotFilename *string
}

func (r ApiReportGroupRequest) CategoryId(categoryId int64) ApiReportGroupRequest {
	r.categoryId = &categoryId
	return r
}

func (r ApiReportGroupRequest) OpponentId(opponentId int64) ApiReportGroupRequest {
	r.opponentId = &opponentId
	return r
}

func (r ApiReportGroupRequest) Reason(reason string) ApiReportGroupRequest {
	r.reason = &reason
	return r
}

func (r ApiReportGroupRequest) Screenshot2Filename(screenshot2Filename string) ApiReportGroupRequest {
	r.screenshot2Filename = &screenshot2Filename
	return r
}

func (r ApiReportGroupRequest) Screenshot3Filename(screenshot3Filename string) ApiReportGroupRequest {
	r.screenshot3Filename = &screenshot3Filename
	return r
}

func (r ApiReportGroupRequest) Screenshot4Filename(screenshot4Filename string) ApiReportGroupRequest {
	r.screenshot4Filename = &screenshot4Filename
	return r
}

func (r ApiReportGroupRequest) ScreenshotFilename(screenshotFilename string) ApiReportGroupRequest {
	r.screenshotFilename = &screenshotFilename
	return r
}

func (r ApiReportGroupRequest) Execute() (*http.Response, error) {
	return r.ApiService.ReportGroupExecute(r)
}

/*
ReportGroup Method for ReportGroup

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param groupId
 @return ApiReportGroupRequest
*/
func (a *GroupsAPIService) ReportGroup(ctx context.Context, groupId int64) ApiReportGroupRequest {
	return ApiReportGroupRequest{
		ApiService: a,
		ctx: ctx,
		groupId: groupId,
	}
}

// Execute executes the request
func (a *GroupsAPIService) ReportGroupExecute(r ApiReportGroupRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "GroupsAPIService.ReportGroup")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v3/groups/{group_id}/report"
	localVarPath = strings.Replace(localVarPath, "{"+"group_id"+"}", url.PathEscape(parameterValueToString(r.groupId, "groupId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.categoryId == nil {
		return nil, reportError("categoryId is required and must be specified")
	}
	if r.opponentId == nil {
		return nil, reportError("opponentId is required and must be specified")
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
	parameterAddToHeaderOrQuery(localVarFormParams, "category_id", r.categoryId, "", "")
	parameterAddToHeaderOrQuery(localVarFormParams, "opponent_id", r.opponentId, "", "")
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

type ApiRequestGroupWalkthroughRequest struct {
	ctx context.Context
	ApiService *GroupsAPIService
	id int64
}

func (r ApiRequestGroupWalkthroughRequest) Execute() (*http.Response, error) {
	return r.ApiService.RequestGroupWalkthroughExecute(r)
}

/*
RequestGroupWalkthrough Method for RequestGroupWalkthrough

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id
 @return ApiRequestGroupWalkthroughRequest
*/
func (a *GroupsAPIService) RequestGroupWalkthrough(ctx context.Context, id int64) ApiRequestGroupWalkthroughRequest {
	return ApiRequestGroupWalkthroughRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
func (a *GroupsAPIService) RequestGroupWalkthroughExecute(r ApiRequestGroupWalkthroughRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "GroupsAPIService.RequestGroupWalkthrough")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/groups/{id}/request_walkthrough"
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

type ApiSearchGroupPostsRequest struct {
	ctx context.Context
	ApiService *GroupsAPIService
	id int64
	keyword *string
	onlyThreadPosts *bool
	fromPostId *int64
	number *int32
}

func (r ApiSearchGroupPostsRequest) Keyword(keyword string) ApiSearchGroupPostsRequest {
	r.keyword = &keyword
	return r
}

func (r ApiSearchGroupPostsRequest) OnlyThreadPosts(onlyThreadPosts bool) ApiSearchGroupPostsRequest {
	r.onlyThreadPosts = &onlyThreadPosts
	return r
}

func (r ApiSearchGroupPostsRequest) FromPostId(fromPostId int64) ApiSearchGroupPostsRequest {
	r.fromPostId = &fromPostId
	return r
}

func (r ApiSearchGroupPostsRequest) Number(number int32) ApiSearchGroupPostsRequest {
	r.number = &number
	return r
}

func (r ApiSearchGroupPostsRequest) Execute() (*PostsResponse, *http.Response, error) {
	return r.ApiService.SearchGroupPostsExecute(r)
}

/*
SearchGroupPosts Method for SearchGroupPosts

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id
 @return ApiSearchGroupPostsRequest
*/
func (a *GroupsAPIService) SearchGroupPosts(ctx context.Context, id int64) ApiSearchGroupPostsRequest {
	return ApiSearchGroupPostsRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
//  @return PostsResponse
func (a *GroupsAPIService) SearchGroupPostsExecute(r ApiSearchGroupPostsRequest) (*PostsResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *PostsResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "GroupsAPIService.SearchGroupPosts")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v2/groups/{id}/posts/search"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterValueToString(r.id, "id")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.keyword == nil {
		return localVarReturnValue, nil, reportError("keyword is required and must be specified")
	}
	if r.onlyThreadPosts == nil {
		return localVarReturnValue, nil, reportError("onlyThreadPosts is required and must be specified")
	}

	parameterAddToHeaderOrQuery(localVarQueryParams, "keyword", r.keyword, "form", "")
	if r.fromPostId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "from_post_id", r.fromPostId, "form", "")
	}
	if r.number != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "number", r.number, "form", "")
	}
	parameterAddToHeaderOrQuery(localVarQueryParams, "only_thread_posts", r.onlyThreadPosts, "form", "")
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

type ApiSetGroupTitleRequest struct {
	ctx context.Context
	ApiService *GroupsAPIService
	id int64
	title *string
}

func (r ApiSetGroupTitleRequest) Title(title string) ApiSetGroupTitleRequest {
	r.title = &title
	return r
}

func (r ApiSetGroupTitleRequest) Execute() (*http.Response, error) {
	return r.ApiService.SetGroupTitleExecute(r)
}

/*
SetGroupTitle Method for SetGroupTitle

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id
 @return ApiSetGroupTitleRequest
*/
func (a *GroupsAPIService) SetGroupTitle(ctx context.Context, id int64) ApiSetGroupTitleRequest {
	return ApiSetGroupTitleRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
func (a *GroupsAPIService) SetGroupTitleExecute(r ApiSetGroupTitleRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "GroupsAPIService.SetGroupTitle")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/groups/{id}/set_title"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterValueToString(r.id, "id")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.title == nil {
		return nil, reportError("title is required and must be specified")
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
	parameterAddToHeaderOrQuery(localVarFormParams, "title", r.title, "", "")
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

type ApiTakeOverGroupRequest struct {
	ctx context.Context
	ApiService *GroupsAPIService
	id int64
}

func (r ApiTakeOverGroupRequest) Execute() (*http.Response, error) {
	return r.ApiService.TakeOverGroupExecute(r)
}

/*
TakeOverGroup Method for TakeOverGroup

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id
 @return ApiTakeOverGroupRequest
*/
func (a *GroupsAPIService) TakeOverGroup(ctx context.Context, id int64) ApiTakeOverGroupRequest {
	return ApiTakeOverGroupRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
func (a *GroupsAPIService) TakeOverGroupExecute(r ApiTakeOverGroupRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "GroupsAPIService.TakeOverGroup")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/groups/{id}/take_over"
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

type ApiTransferGroupRequest struct {
	ctx context.Context
	ApiService *GroupsAPIService
	id int64
	apiKey *string
	signedInfo *string
	timestamp *int64
	userId *int64
	uuid *string
}

func (r ApiTransferGroupRequest) ApiKey(apiKey string) ApiTransferGroupRequest {
	r.apiKey = &apiKey
	return r
}

func (r ApiTransferGroupRequest) SignedInfo(signedInfo string) ApiTransferGroupRequest {
	r.signedInfo = &signedInfo
	return r
}

func (r ApiTransferGroupRequest) Timestamp(timestamp int64) ApiTransferGroupRequest {
	r.timestamp = &timestamp
	return r
}

func (r ApiTransferGroupRequest) UserId(userId int64) ApiTransferGroupRequest {
	r.userId = &userId
	return r
}

func (r ApiTransferGroupRequest) Uuid(uuid string) ApiTransferGroupRequest {
	r.uuid = &uuid
	return r
}

func (r ApiTransferGroupRequest) Execute() (*http.Response, error) {
	return r.ApiService.TransferGroupExecute(r)
}

/*
TransferGroup Method for TransferGroup

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id
 @return ApiTransferGroupRequest
*/
func (a *GroupsAPIService) TransferGroup(ctx context.Context, id int64) ApiTransferGroupRequest {
	return ApiTransferGroupRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
func (a *GroupsAPIService) TransferGroupExecute(r ApiTransferGroupRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "GroupsAPIService.TransferGroup")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v3/groups/{id}/transfer"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterValueToString(r.id, "id")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.apiKey == nil {
		return nil, reportError("apiKey is required and must be specified")
	}
	if r.signedInfo == nil {
		return nil, reportError("signedInfo is required and must be specified")
	}
	if r.timestamp == nil {
		return nil, reportError("timestamp is required and must be specified")
	}
	if r.userId == nil {
		return nil, reportError("userId is required and must be specified")
	}
	if r.uuid == nil {
		return nil, reportError("uuid is required and must be specified")
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
	parameterAddToHeaderOrQuery(localVarFormParams, "api_key", r.apiKey, "", "")
	parameterAddToHeaderOrQuery(localVarFormParams, "signed_info", r.signedInfo, "", "")
	parameterAddToHeaderOrQuery(localVarFormParams, "timestamp", r.timestamp, "", "")
	parameterAddToHeaderOrQuery(localVarFormParams, "user_id", r.userId, "", "")
	parameterAddToHeaderOrQuery(localVarFormParams, "uuid", r.uuid, "", "")
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

type ApiUnbanGroupUserRequest struct {
	ctx context.Context
	ApiService *GroupsAPIService
	id int64
	userId int64
}

func (r ApiUnbanGroupUserRequest) Execute() (*http.Response, error) {
	return r.ApiService.UnbanGroupUserExecute(r)
}

/*
UnbanGroupUser Method for UnbanGroupUser

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id
 @param userId
 @return ApiUnbanGroupUserRequest
*/
func (a *GroupsAPIService) UnbanGroupUser(ctx context.Context, id int64, userId int64) ApiUnbanGroupUserRequest {
	return ApiUnbanGroupUserRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
		userId: userId,
	}
}

// Execute executes the request
func (a *GroupsAPIService) UnbanGroupUserExecute(r ApiUnbanGroupUserRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "GroupsAPIService.UnbanGroupUser")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/groups/{id}/unban/{userId}"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterValueToString(r.id, "id")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"userId"+"}", url.PathEscape(parameterValueToString(r.userId, "userId")), -1)

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

type ApiUnpinGroupHighlightPostRequest struct {
	ctx context.Context
	ApiService *GroupsAPIService
	groupId int64
	postId int64
}

func (r ApiUnpinGroupHighlightPostRequest) Execute() (*http.Response, error) {
	return r.ApiService.UnpinGroupHighlightPostExecute(r)
}

/*
UnpinGroupHighlightPost Method for UnpinGroupHighlightPost

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param groupId
 @param postId
 @return ApiUnpinGroupHighlightPostRequest
*/
func (a *GroupsAPIService) UnpinGroupHighlightPost(ctx context.Context, groupId int64, postId int64) ApiUnpinGroupHighlightPostRequest {
	return ApiUnpinGroupHighlightPostRequest{
		ApiService: a,
		ctx: ctx,
		groupId: groupId,
		postId: postId,
	}
}

// Execute executes the request
func (a *GroupsAPIService) UnpinGroupHighlightPostExecute(r ApiUnpinGroupHighlightPostRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "GroupsAPIService.UnpinGroupHighlightPost")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/groups/{group_id}/highlights/{post_id}"
	localVarPath = strings.Replace(localVarPath, "{"+"group_id"+"}", url.PathEscape(parameterValueToString(r.groupId, "groupId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"post_id"+"}", url.PathEscape(parameterValueToString(r.postId, "postId")), -1)

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

type ApiUpdateGroupRequest struct {
	ctx context.Context
	ApiService *GroupsAPIService
	id int64
	apiKey *string
	signedInfo *string
	timestamp *int64
	uuid *string
	allowMembersToPostImageAndVideo *bool
	allowMembersToPostUrl *bool
	allowOwnershipTransfer *bool
	allowThreadCreationBy *string
	callTimelineDisplay *bool
	coverImageFilename *string
	description *string
	gender *int32
	generationGroupsLimit *int32
	groupCategoryId *int64
	groupIconFilename *string
	guidelines *string
	hideConferenceCall *bool
	hideFromGameEight *bool
	hideReportedPosts *bool
	isPrivate *bool
	onlyMobileVerified *bool
	onlyVerifiedAge *bool
	secret *bool
	subCategoryId *string
	topic *string
}

func (r ApiUpdateGroupRequest) ApiKey(apiKey string) ApiUpdateGroupRequest {
	r.apiKey = &apiKey
	return r
}

func (r ApiUpdateGroupRequest) SignedInfo(signedInfo string) ApiUpdateGroupRequest {
	r.signedInfo = &signedInfo
	return r
}

func (r ApiUpdateGroupRequest) Timestamp(timestamp int64) ApiUpdateGroupRequest {
	r.timestamp = &timestamp
	return r
}

func (r ApiUpdateGroupRequest) Uuid(uuid string) ApiUpdateGroupRequest {
	r.uuid = &uuid
	return r
}

func (r ApiUpdateGroupRequest) AllowMembersToPostImageAndVideo(allowMembersToPostImageAndVideo bool) ApiUpdateGroupRequest {
	r.allowMembersToPostImageAndVideo = &allowMembersToPostImageAndVideo
	return r
}

func (r ApiUpdateGroupRequest) AllowMembersToPostUrl(allowMembersToPostUrl bool) ApiUpdateGroupRequest {
	r.allowMembersToPostUrl = &allowMembersToPostUrl
	return r
}

func (r ApiUpdateGroupRequest) AllowOwnershipTransfer(allowOwnershipTransfer bool) ApiUpdateGroupRequest {
	r.allowOwnershipTransfer = &allowOwnershipTransfer
	return r
}

func (r ApiUpdateGroupRequest) AllowThreadCreationBy(allowThreadCreationBy string) ApiUpdateGroupRequest {
	r.allowThreadCreationBy = &allowThreadCreationBy
	return r
}

func (r ApiUpdateGroupRequest) CallTimelineDisplay(callTimelineDisplay bool) ApiUpdateGroupRequest {
	r.callTimelineDisplay = &callTimelineDisplay
	return r
}

func (r ApiUpdateGroupRequest) CoverImageFilename(coverImageFilename string) ApiUpdateGroupRequest {
	r.coverImageFilename = &coverImageFilename
	return r
}

func (r ApiUpdateGroupRequest) Description(description string) ApiUpdateGroupRequest {
	r.description = &description
	return r
}

func (r ApiUpdateGroupRequest) Gender(gender int32) ApiUpdateGroupRequest {
	r.gender = &gender
	return r
}

func (r ApiUpdateGroupRequest) GenerationGroupsLimit(generationGroupsLimit int32) ApiUpdateGroupRequest {
	r.generationGroupsLimit = &generationGroupsLimit
	return r
}

func (r ApiUpdateGroupRequest) GroupCategoryId(groupCategoryId int64) ApiUpdateGroupRequest {
	r.groupCategoryId = &groupCategoryId
	return r
}

func (r ApiUpdateGroupRequest) GroupIconFilename(groupIconFilename string) ApiUpdateGroupRequest {
	r.groupIconFilename = &groupIconFilename
	return r
}

func (r ApiUpdateGroupRequest) Guidelines(guidelines string) ApiUpdateGroupRequest {
	r.guidelines = &guidelines
	return r
}

func (r ApiUpdateGroupRequest) HideConferenceCall(hideConferenceCall bool) ApiUpdateGroupRequest {
	r.hideConferenceCall = &hideConferenceCall
	return r
}

func (r ApiUpdateGroupRequest) HideFromGameEight(hideFromGameEight bool) ApiUpdateGroupRequest {
	r.hideFromGameEight = &hideFromGameEight
	return r
}

func (r ApiUpdateGroupRequest) HideReportedPosts(hideReportedPosts bool) ApiUpdateGroupRequest {
	r.hideReportedPosts = &hideReportedPosts
	return r
}

func (r ApiUpdateGroupRequest) IsPrivate(isPrivate bool) ApiUpdateGroupRequest {
	r.isPrivate = &isPrivate
	return r
}

func (r ApiUpdateGroupRequest) OnlyMobileVerified(onlyMobileVerified bool) ApiUpdateGroupRequest {
	r.onlyMobileVerified = &onlyMobileVerified
	return r
}

func (r ApiUpdateGroupRequest) OnlyVerifiedAge(onlyVerifiedAge bool) ApiUpdateGroupRequest {
	r.onlyVerifiedAge = &onlyVerifiedAge
	return r
}

func (r ApiUpdateGroupRequest) Secret(secret bool) ApiUpdateGroupRequest {
	r.secret = &secret
	return r
}

func (r ApiUpdateGroupRequest) SubCategoryId(subCategoryId string) ApiUpdateGroupRequest {
	r.subCategoryId = &subCategoryId
	return r
}

func (r ApiUpdateGroupRequest) Topic(topic string) ApiUpdateGroupRequest {
	r.topic = &topic
	return r
}

func (r ApiUpdateGroupRequest) Execute() (*GroupResponse, *http.Response, error) {
	return r.ApiService.UpdateGroupExecute(r)
}

/*
UpdateGroup Method for UpdateGroup

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id
 @return ApiUpdateGroupRequest
*/
func (a *GroupsAPIService) UpdateGroup(ctx context.Context, id int64) ApiUpdateGroupRequest {
	return ApiUpdateGroupRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
//  @return GroupResponse
func (a *GroupsAPIService) UpdateGroupExecute(r ApiUpdateGroupRequest) (*GroupResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *GroupResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "GroupsAPIService.UpdateGroup")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v3/groups/{id}/update"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterValueToString(r.id, "id")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.apiKey == nil {
		return localVarReturnValue, nil, reportError("apiKey is required and must be specified")
	}
	if r.signedInfo == nil {
		return localVarReturnValue, nil, reportError("signedInfo is required and must be specified")
	}
	if r.timestamp == nil {
		return localVarReturnValue, nil, reportError("timestamp is required and must be specified")
	}
	if r.uuid == nil {
		return localVarReturnValue, nil, reportError("uuid is required and must be specified")
	}

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
	if r.allowMembersToPostImageAndVideo != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "allow_members_to_post_image_and_video", r.allowMembersToPostImageAndVideo, "", "")
	}
	if r.allowMembersToPostUrl != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "allow_members_to_post_url", r.allowMembersToPostUrl, "", "")
	}
	if r.allowOwnershipTransfer != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "allow_ownership_transfer", r.allowOwnershipTransfer, "", "")
	}
	if r.allowThreadCreationBy != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "allow_thread_creation_by", r.allowThreadCreationBy, "", "")
	}
	parameterAddToHeaderOrQuery(localVarFormParams, "api_key", r.apiKey, "", "")
	if r.callTimelineDisplay != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "call_timeline_display", r.callTimelineDisplay, "", "")
	}
	if r.coverImageFilename != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "cover_image_filename", r.coverImageFilename, "", "")
	}
	if r.description != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "description", r.description, "", "")
	}
	if r.gender != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "gender", r.gender, "", "")
	}
	if r.generationGroupsLimit != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "generation_groups_limit", r.generationGroupsLimit, "", "")
	}
	if r.groupCategoryId != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "group_category_id", r.groupCategoryId, "", "")
	}
	if r.groupIconFilename != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "group_icon_filename", r.groupIconFilename, "", "")
	}
	if r.guidelines != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "guidelines", r.guidelines, "", "")
	}
	if r.hideConferenceCall != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "hide_conference_call", r.hideConferenceCall, "", "")
	}
	if r.hideFromGameEight != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "hide_from_game_eight", r.hideFromGameEight, "", "")
	}
	if r.hideReportedPosts != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "hide_reported_posts", r.hideReportedPosts, "", "")
	}
	if r.isPrivate != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "is_private", r.isPrivate, "", "")
	}
	if r.onlyMobileVerified != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "only_mobile_verified", r.onlyMobileVerified, "", "")
	}
	if r.onlyVerifiedAge != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "only_verified_age", r.onlyVerifiedAge, "", "")
	}
	if r.secret != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "secret", r.secret, "", "")
	}
	parameterAddToHeaderOrQuery(localVarFormParams, "signed_info", r.signedInfo, "", "")
	if r.subCategoryId != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "sub_category_id", r.subCategoryId, "", "")
	}
	parameterAddToHeaderOrQuery(localVarFormParams, "timestamp", r.timestamp, "", "")
	if r.topic != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "topic", r.topic, "", "")
	}
	parameterAddToHeaderOrQuery(localVarFormParams, "uuid", r.uuid, "", "")
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

type ApiUpdateRelatedGroupsRequest struct {
	ctx context.Context
	ApiService *GroupsAPIService
	id int64
	relatedGroupId *[]int64
}

func (r ApiUpdateRelatedGroupsRequest) RelatedGroupId(relatedGroupId []int64) ApiUpdateRelatedGroupsRequest {
	r.relatedGroupId = &relatedGroupId
	return r
}

func (r ApiUpdateRelatedGroupsRequest) Execute() (*http.Response, error) {
	return r.ApiService.UpdateRelatedGroupsExecute(r)
}

/*
UpdateRelatedGroups Method for UpdateRelatedGroups

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id
 @return ApiUpdateRelatedGroupsRequest
*/
func (a *GroupsAPIService) UpdateRelatedGroups(ctx context.Context, id int64) ApiUpdateRelatedGroupsRequest {
	return ApiUpdateRelatedGroupsRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
func (a *GroupsAPIService) UpdateRelatedGroupsExecute(r ApiUpdateRelatedGroupsRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPut
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "GroupsAPIService.UpdateRelatedGroups")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/groups/{id}/related"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterValueToString(r.id, "id")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.relatedGroupId == nil {
		return nil, reportError("relatedGroupId is required and must be specified")
	}

	{
		t := *r.relatedGroupId
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				parameterAddToHeaderOrQuery(localVarQueryParams, "related_group_id[]", s.Index(i).Interface(), "form", "multi")
			}
		} else {
			parameterAddToHeaderOrQuery(localVarQueryParams, "related_group_id[]", t, "form", "multi")
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

type ApiVisitGroupRequest struct {
	ctx context.Context
	ApiService *GroupsAPIService
	id int64
}

func (r ApiVisitGroupRequest) Execute() (*http.Response, error) {
	return r.ApiService.VisitGroupExecute(r)
}

/*
VisitGroup Method for VisitGroup

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id
 @return ApiVisitGroupRequest
*/
func (a *GroupsAPIService) VisitGroup(ctx context.Context, id int64) ApiVisitGroupRequest {
	return ApiVisitGroupRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
func (a *GroupsAPIService) VisitGroupExecute(r ApiVisitGroupRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "GroupsAPIService.VisitGroup")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/groups/{id}/visit"
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

type ApiWithdrawGroupDeputyRequest struct {
	ctx context.Context
	ApiService *GroupsAPIService
	groupId int64
	userId int64
}

func (r ApiWithdrawGroupDeputyRequest) Execute() (*http.Response, error) {
	return r.ApiService.WithdrawGroupDeputyExecute(r)
}

/*
WithdrawGroupDeputy Method for WithdrawGroupDeputy

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param groupId
 @param userId
 @return ApiWithdrawGroupDeputyRequest
*/
func (a *GroupsAPIService) WithdrawGroupDeputy(ctx context.Context, groupId int64, userId int64) ApiWithdrawGroupDeputyRequest {
	return ApiWithdrawGroupDeputyRequest{
		ApiService: a,
		ctx: ctx,
		groupId: groupId,
		userId: userId,
	}
}

// Execute executes the request
func (a *GroupsAPIService) WithdrawGroupDeputyExecute(r ApiWithdrawGroupDeputyRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPut
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "GroupsAPIService.WithdrawGroupDeputy")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/groups/{group_id}/deputize/{user_id}/withdraw"
	localVarPath = strings.Replace(localVarPath, "{"+"group_id"+"}", url.PathEscape(parameterValueToString(r.groupId, "groupId")), -1)
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

type ApiWithdrawGroupTransferRequest struct {
	ctx context.Context
	ApiService *GroupsAPIService
	id int64
	userId *int64
}

func (r ApiWithdrawGroupTransferRequest) UserId(userId int64) ApiWithdrawGroupTransferRequest {
	r.userId = &userId
	return r
}

func (r ApiWithdrawGroupTransferRequest) Execute() (*http.Response, error) {
	return r.ApiService.WithdrawGroupTransferExecute(r)
}

/*
WithdrawGroupTransfer Method for WithdrawGroupTransfer

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id
 @return ApiWithdrawGroupTransferRequest
*/
func (a *GroupsAPIService) WithdrawGroupTransfer(ctx context.Context, id int64) ApiWithdrawGroupTransferRequest {
	return ApiWithdrawGroupTransferRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
func (a *GroupsAPIService) WithdrawGroupTransferExecute(r ApiWithdrawGroupTransferRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPut
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "GroupsAPIService.WithdrawGroupTransfer")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/groups/{id}/transfer/withdraw"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterValueToString(r.id, "id")), -1)

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
