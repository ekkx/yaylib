
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


// ThreadsAPIService ThreadsAPI service
type ThreadsAPIService service

type ApiAddThreadMemberRequest struct {
	ctx context.Context
	ApiService *ThreadsAPIService
	threadId int64
	id int64
}

func (r ApiAddThreadMemberRequest) Execute() (*http.Response, error) {
	return r.ApiService.AddThreadMemberExecute(r)
}

func (r ApiAddThreadMemberRequest) ExecuteRaw() ([]byte, *http.Response, error) {
	httpResp, err := r.Execute()
	return executeRaw(httpResp, err)
}

/*
AddThreadMember Method for AddThreadMember

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param threadId
 @param id
 @return ApiAddThreadMemberRequest
*/
func (a *ThreadsAPIService) AddThreadMember(ctx context.Context, threadId int64, id int64) ApiAddThreadMemberRequest {
	return ApiAddThreadMemberRequest{
		ApiService: a,
		ctx: ctx,
		threadId: threadId,
		id: id,
	}
}

// Execute executes the request
func (a *ThreadsAPIService) AddThreadMemberExecute(r ApiAddThreadMemberRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ThreadsAPIService.AddThreadMember")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/threads/{thread_id}/members/{id}"
	localVarPath = strings.Replace(localVarPath, "{"+"thread_id"+"}", url.PathEscape(parameterValueToString(r.threadId, "threadId")), -1)
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

type ApiCreateThreadRequest struct {
	ctx context.Context
	ApiService *ThreadsAPIService
	createGroupThreadRequest *CreateGroupThreadRequest
}

func (r ApiCreateThreadRequest) CreateGroupThreadRequest(createGroupThreadRequest CreateGroupThreadRequest) ApiCreateThreadRequest {
	r.createGroupThreadRequest = &createGroupThreadRequest
	return r
}

func (r ApiCreateThreadRequest) Execute() (*ThreadInfo, *http.Response, error) {
	return r.ApiService.CreateThreadExecute(r)
}

func (r ApiCreateThreadRequest) ExecuteRaw() ([]byte, *http.Response, error) {
	_, httpResp, err := r.Execute()
	return executeRaw(httpResp, err)
}

/*
CreateThread Method for CreateThread

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiCreateThreadRequest
*/
func (a *ThreadsAPIService) CreateThread(ctx context.Context) ApiCreateThreadRequest {
	return ApiCreateThreadRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return ThreadInfo
func (a *ThreadsAPIService) CreateThreadExecute(r ApiCreateThreadRequest) (*ThreadInfo, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ThreadInfo
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ThreadsAPIService.CreateThread")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/threads"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.createGroupThreadRequest == nil {
		return localVarReturnValue, nil, reportError("createGroupThreadRequest is required and must be specified")
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
	localVarPostBody = r.createGroupThreadRequest
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

type ApiCreateThreadPostRequest struct {
	ctx context.Context
	ApiService *ThreadsAPIService
	id int64
	xJwt *string
	attachment2Filename *string
	attachment3Filename *string
	attachment4Filename *string
	attachment5Filename *string
	attachment6Filename *string
	attachment7Filename *string
	attachment8Filename *string
	attachment9Filename *string
	attachmentFilename *string
	choices *[]string
	color *int32
	fontSize *int32
	groupId *int64
	inReplyTo *int64
	language *string
	mentionIds *[]int64
	messageTags *map[string]interface{}
	postType *PostType
	sharedUrl *map[string]interface{}
	text *string
	videoFileName *string
}

func (r ApiCreateThreadPostRequest) XJwt(xJwt string) ApiCreateThreadPostRequest {
	r.xJwt = &xJwt
	return r
}

func (r ApiCreateThreadPostRequest) Attachment2Filename(attachment2Filename string) ApiCreateThreadPostRequest {
	r.attachment2Filename = &attachment2Filename
	return r
}

func (r ApiCreateThreadPostRequest) Attachment3Filename(attachment3Filename string) ApiCreateThreadPostRequest {
	r.attachment3Filename = &attachment3Filename
	return r
}

func (r ApiCreateThreadPostRequest) Attachment4Filename(attachment4Filename string) ApiCreateThreadPostRequest {
	r.attachment4Filename = &attachment4Filename
	return r
}

func (r ApiCreateThreadPostRequest) Attachment5Filename(attachment5Filename string) ApiCreateThreadPostRequest {
	r.attachment5Filename = &attachment5Filename
	return r
}

func (r ApiCreateThreadPostRequest) Attachment6Filename(attachment6Filename string) ApiCreateThreadPostRequest {
	r.attachment6Filename = &attachment6Filename
	return r
}

func (r ApiCreateThreadPostRequest) Attachment7Filename(attachment7Filename string) ApiCreateThreadPostRequest {
	r.attachment7Filename = &attachment7Filename
	return r
}

func (r ApiCreateThreadPostRequest) Attachment8Filename(attachment8Filename string) ApiCreateThreadPostRequest {
	r.attachment8Filename = &attachment8Filename
	return r
}

func (r ApiCreateThreadPostRequest) Attachment9Filename(attachment9Filename string) ApiCreateThreadPostRequest {
	r.attachment9Filename = &attachment9Filename
	return r
}

func (r ApiCreateThreadPostRequest) AttachmentFilename(attachmentFilename string) ApiCreateThreadPostRequest {
	r.attachmentFilename = &attachmentFilename
	return r
}

func (r ApiCreateThreadPostRequest) Choices(choices []string) ApiCreateThreadPostRequest {
	r.choices = &choices
	return r
}

func (r ApiCreateThreadPostRequest) Color(color int32) ApiCreateThreadPostRequest {
	r.color = &color
	return r
}

func (r ApiCreateThreadPostRequest) FontSize(fontSize int32) ApiCreateThreadPostRequest {
	r.fontSize = &fontSize
	return r
}

func (r ApiCreateThreadPostRequest) GroupId(groupId int64) ApiCreateThreadPostRequest {
	r.groupId = &groupId
	return r
}

func (r ApiCreateThreadPostRequest) InReplyTo(inReplyTo int64) ApiCreateThreadPostRequest {
	r.inReplyTo = &inReplyTo
	return r
}

func (r ApiCreateThreadPostRequest) Language(language string) ApiCreateThreadPostRequest {
	r.language = &language
	return r
}

func (r ApiCreateThreadPostRequest) MentionIds(mentionIds []int64) ApiCreateThreadPostRequest {
	r.mentionIds = &mentionIds
	return r
}

func (r ApiCreateThreadPostRequest) MessageTags(messageTags map[string]interface{}) ApiCreateThreadPostRequest {
	r.messageTags = &messageTags
	return r
}

func (r ApiCreateThreadPostRequest) PostType(postType PostType) ApiCreateThreadPostRequest {
	r.postType = &postType
	return r
}

func (r ApiCreateThreadPostRequest) SharedUrl(sharedUrl map[string]interface{}) ApiCreateThreadPostRequest {
	r.sharedUrl = &sharedUrl
	return r
}

func (r ApiCreateThreadPostRequest) Text(text string) ApiCreateThreadPostRequest {
	r.text = &text
	return r
}

func (r ApiCreateThreadPostRequest) VideoFileName(videoFileName string) ApiCreateThreadPostRequest {
	r.videoFileName = &videoFileName
	return r
}

func (r ApiCreateThreadPostRequest) Execute() (*Post, *http.Response, error) {
	return r.ApiService.CreateThreadPostExecute(r)
}

func (r ApiCreateThreadPostRequest) ExecuteRaw() ([]byte, *http.Response, error) {
	_, httpResp, err := r.Execute()
	return executeRaw(httpResp, err)
}

/*
CreateThreadPost Method for CreateThreadPost

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id
 @return ApiCreateThreadPostRequest
*/
func (a *ThreadsAPIService) CreateThreadPost(ctx context.Context, id int64) ApiCreateThreadPostRequest {
	return ApiCreateThreadPostRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
//  @return Post
func (a *ThreadsAPIService) CreateThreadPostExecute(r ApiCreateThreadPostRequest) (*Post, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *Post
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ThreadsAPIService.CreateThreadPost")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/threads/{id}/posts"
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
	if r.xJwt != nil {
		parameterAddToHeaderOrQuery(localVarHeaderParams, "X-Jwt", r.xJwt, "simple", "")
	}
	if r.attachment2Filename != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "attachment_2_filename", r.attachment2Filename, "", "")
	}
	if r.attachment3Filename != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "attachment_3_filename", r.attachment3Filename, "", "")
	}
	if r.attachment4Filename != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "attachment_4_filename", r.attachment4Filename, "", "")
	}
	if r.attachment5Filename != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "attachment_5_filename", r.attachment5Filename, "", "")
	}
	if r.attachment6Filename != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "attachment_6_filename", r.attachment6Filename, "", "")
	}
	if r.attachment7Filename != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "attachment_7_filename", r.attachment7Filename, "", "")
	}
	if r.attachment8Filename != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "attachment_8_filename", r.attachment8Filename, "", "")
	}
	if r.attachment9Filename != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "attachment_9_filename", r.attachment9Filename, "", "")
	}
	if r.attachmentFilename != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "attachment_filename", r.attachmentFilename, "", "")
	}
	if r.choices != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "choices[]", r.choices, "", "csv")
	}
	if r.color != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "color", r.color, "", "")
	}
	if r.fontSize != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "font_size", r.fontSize, "", "")
	}
	if r.groupId != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "group_id", r.groupId, "", "")
	}
	if r.inReplyTo != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "in_reply_to", r.inReplyTo, "", "")
	}
	if r.language != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "language", r.language, "", "")
	}
	if r.mentionIds != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "mention_ids[]", r.mentionIds, "", "csv")
	}
	if r.messageTags != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "message_tags", r.messageTags, "", "")
	}
	if r.postType != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "post_type", r.postType, "", "")
	}
	if r.sharedUrl != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "shared_url", r.sharedUrl, "", "")
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

type ApiDeleteThreadRequest struct {
	ctx context.Context
	ApiService *ThreadsAPIService
	id int64
}

func (r ApiDeleteThreadRequest) Execute() (*http.Response, error) {
	return r.ApiService.DeleteThreadExecute(r)
}

func (r ApiDeleteThreadRequest) ExecuteRaw() ([]byte, *http.Response, error) {
	httpResp, err := r.Execute()
	return executeRaw(httpResp, err)
}

/*
DeleteThread Method for DeleteThread

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id
 @return ApiDeleteThreadRequest
*/
func (a *ThreadsAPIService) DeleteThread(ctx context.Context, id int64) ApiDeleteThreadRequest {
	return ApiDeleteThreadRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
func (a *ThreadsAPIService) DeleteThreadExecute(r ApiDeleteThreadRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ThreadsAPIService.DeleteThread")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/threads/{id}"
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

type ApiGetJoinedThreadStatusesRequest struct {
	ctx context.Context
	ApiService *ThreadsAPIService
	ids *[]int64
}

func (r ApiGetJoinedThreadStatusesRequest) Ids(ids []int64) ApiGetJoinedThreadStatusesRequest {
	r.ids = &ids
	return r
}

func (r ApiGetJoinedThreadStatusesRequest) Execute() (map[string]string, *http.Response, error) {
	return r.ApiService.GetJoinedThreadStatusesExecute(r)
}

func (r ApiGetJoinedThreadStatusesRequest) ExecuteRaw() ([]byte, *http.Response, error) {
	_, httpResp, err := r.Execute()
	return executeRaw(httpResp, err)
}

/*
GetJoinedThreadStatuses Method for GetJoinedThreadStatuses

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiGetJoinedThreadStatusesRequest
*/
func (a *ThreadsAPIService) GetJoinedThreadStatuses(ctx context.Context) ApiGetJoinedThreadStatusesRequest {
	return ApiGetJoinedThreadStatusesRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return map[string]string
func (a *ThreadsAPIService) GetJoinedThreadStatusesExecute(r ApiGetJoinedThreadStatusesRequest) (map[string]string, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  map[string]string
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ThreadsAPIService.GetJoinedThreadStatuses")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/threads/joined_statuses"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.ids != nil {
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

type ApiGetThreadRequest struct {
	ctx context.Context
	ApiService *ThreadsAPIService
	id int64
}

func (r ApiGetThreadRequest) Execute() (*ThreadInfo, *http.Response, error) {
	return r.ApiService.GetThreadExecute(r)
}

func (r ApiGetThreadRequest) ExecuteRaw() ([]byte, *http.Response, error) {
	_, httpResp, err := r.Execute()
	return executeRaw(httpResp, err)
}

/*
GetThread Method for GetThread

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id
 @return ApiGetThreadRequest
*/
func (a *ThreadsAPIService) GetThread(ctx context.Context, id int64) ApiGetThreadRequest {
	return ApiGetThreadRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
//  @return ThreadInfo
func (a *ThreadsAPIService) GetThreadExecute(r ApiGetThreadRequest) (*ThreadInfo, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ThreadInfo
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ThreadsAPIService.GetThread")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/threads/{id}"
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

type ApiGetThreadPostsRequest struct {
	ctx context.Context
	ApiService *ThreadsAPIService
	id int64
	postType *PostType
	number *int32
	from *int64
}

func (r ApiGetThreadPostsRequest) PostType(postType PostType) ApiGetThreadPostsRequest {
	r.postType = &postType
	return r
}

func (r ApiGetThreadPostsRequest) Number(number int32) ApiGetThreadPostsRequest {
	r.number = &number
	return r
}

func (r ApiGetThreadPostsRequest) From(from int64) ApiGetThreadPostsRequest {
	r.from = &from
	return r
}

func (r ApiGetThreadPostsRequest) Execute() (*PostsResponse, *http.Response, error) {
	return r.ApiService.GetThreadPostsExecute(r)
}

func (r ApiGetThreadPostsRequest) ExecuteRaw() ([]byte, *http.Response, error) {
	_, httpResp, err := r.Execute()
	return executeRaw(httpResp, err)
}

/*
GetThreadPosts Method for GetThreadPosts

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id
 @return ApiGetThreadPostsRequest
*/
func (a *ThreadsAPIService) GetThreadPosts(ctx context.Context, id int64) ApiGetThreadPostsRequest {
	return ApiGetThreadPostsRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
//  @return PostsResponse
func (a *ThreadsAPIService) GetThreadPostsExecute(r ApiGetThreadPostsRequest) (*PostsResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *PostsResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ThreadsAPIService.GetThreadPosts")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/threads/{id}/posts"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterValueToString(r.id, "id")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.postType != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "post_type", r.postType, "form", "")
	}
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

type ApiListThreadsRequest struct {
	ctx context.Context
	ApiService *ThreadsAPIService
	groupId *int64
	from *string
	joinStatus *string
}

func (r ApiListThreadsRequest) GroupId(groupId int64) ApiListThreadsRequest {
	r.groupId = &groupId
	return r
}

func (r ApiListThreadsRequest) From(from string) ApiListThreadsRequest {
	r.from = &from
	return r
}

func (r ApiListThreadsRequest) JoinStatus(joinStatus string) ApiListThreadsRequest {
	r.joinStatus = &joinStatus
	return r
}

func (r ApiListThreadsRequest) Execute() (*GroupThreadListResponse, *http.Response, error) {
	return r.ApiService.ListThreadsExecute(r)
}

func (r ApiListThreadsRequest) ExecuteRaw() ([]byte, *http.Response, error) {
	_, httpResp, err := r.Execute()
	return executeRaw(httpResp, err)
}

/*
ListThreads Method for ListThreads

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiListThreadsRequest
*/
func (a *ThreadsAPIService) ListThreads(ctx context.Context) ApiListThreadsRequest {
	return ApiListThreadsRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return GroupThreadListResponse
func (a *ThreadsAPIService) ListThreadsExecute(r ApiListThreadsRequest) (*GroupThreadListResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *GroupThreadListResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ThreadsAPIService.ListThreads")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/threads"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.groupId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "group_id", r.groupId, "form", "")
	}
	if r.from != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "from", r.from, "form", "")
	}
	if r.joinStatus != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "join_status", r.joinStatus, "form", "")
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

type ApiRemoveThreadMemberRequest struct {
	ctx context.Context
	ApiService *ThreadsAPIService
	threadId int64
	id int64
}

func (r ApiRemoveThreadMemberRequest) Execute() (*http.Response, error) {
	return r.ApiService.RemoveThreadMemberExecute(r)
}

func (r ApiRemoveThreadMemberRequest) ExecuteRaw() ([]byte, *http.Response, error) {
	httpResp, err := r.Execute()
	return executeRaw(httpResp, err)
}

/*
RemoveThreadMember Method for RemoveThreadMember

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param threadId
 @param id
 @return ApiRemoveThreadMemberRequest
*/
func (a *ThreadsAPIService) RemoveThreadMember(ctx context.Context, threadId int64, id int64) ApiRemoveThreadMemberRequest {
	return ApiRemoveThreadMemberRequest{
		ApiService: a,
		ctx: ctx,
		threadId: threadId,
		id: id,
	}
}

// Execute executes the request
func (a *ThreadsAPIService) RemoveThreadMemberExecute(r ApiRemoveThreadMemberRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ThreadsAPIService.RemoveThreadMember")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/threads/{thread_id}/members/{id}"
	localVarPath = strings.Replace(localVarPath, "{"+"thread_id"+"}", url.PathEscape(parameterValueToString(r.threadId, "threadId")), -1)
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

type ApiUpdateThreadRequest struct {
	ctx context.Context
	ApiService *ThreadsAPIService
	id int64
	threadIconFilename *string
	title *string
}

func (r ApiUpdateThreadRequest) ThreadIconFilename(threadIconFilename string) ApiUpdateThreadRequest {
	r.threadIconFilename = &threadIconFilename
	return r
}

func (r ApiUpdateThreadRequest) Title(title string) ApiUpdateThreadRequest {
	r.title = &title
	return r
}

func (r ApiUpdateThreadRequest) Execute() (*http.Response, error) {
	return r.ApiService.UpdateThreadExecute(r)
}

func (r ApiUpdateThreadRequest) ExecuteRaw() ([]byte, *http.Response, error) {
	httpResp, err := r.Execute()
	return executeRaw(httpResp, err)
}

/*
UpdateThread Method for UpdateThread

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id
 @return ApiUpdateThreadRequest
*/
func (a *ThreadsAPIService) UpdateThread(ctx context.Context, id int64) ApiUpdateThreadRequest {
	return ApiUpdateThreadRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
func (a *ThreadsAPIService) UpdateThreadExecute(r ApiUpdateThreadRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPut
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ThreadsAPIService.UpdateThread")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/threads/{id}"
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
	if r.threadIconFilename != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "thread_icon_filename", r.threadIconFilename, "", "")
	}
	if r.title != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "title", r.title, "", "")
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
