
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


// PostsAPIService PostsAPI service
type PostsAPIService service

type ApiCreateConferenceCallPostRequest struct {
	ctx context.Context
	ApiService *PostsAPIService
	apiKey *string
	signedInfo *string
	timestamp *int64
	uuid *string
	attachment2Filename *string
	attachment3Filename *string
	attachment4Filename *string
	attachment5Filename *string
	attachment6Filename *string
	attachment7Filename *string
	attachment8Filename *string
	attachment9Filename *string
	attachmentFilename *string
	callType *string
	categoryId *int64
	color *int32
	fontSize *int32
	gameTitle *string
	groupId *int64
	joinableBy *string
	language *string
	messageTags *map[string]interface{}
	text *string
}

func (r ApiCreateConferenceCallPostRequest) ApiKey(apiKey string) ApiCreateConferenceCallPostRequest {
	r.apiKey = &apiKey
	return r
}

func (r ApiCreateConferenceCallPostRequest) SignedInfo(signedInfo string) ApiCreateConferenceCallPostRequest {
	r.signedInfo = &signedInfo
	return r
}

func (r ApiCreateConferenceCallPostRequest) Timestamp(timestamp int64) ApiCreateConferenceCallPostRequest {
	r.timestamp = &timestamp
	return r
}

func (r ApiCreateConferenceCallPostRequest) Uuid(uuid string) ApiCreateConferenceCallPostRequest {
	r.uuid = &uuid
	return r
}

func (r ApiCreateConferenceCallPostRequest) Attachment2Filename(attachment2Filename string) ApiCreateConferenceCallPostRequest {
	r.attachment2Filename = &attachment2Filename
	return r
}

func (r ApiCreateConferenceCallPostRequest) Attachment3Filename(attachment3Filename string) ApiCreateConferenceCallPostRequest {
	r.attachment3Filename = &attachment3Filename
	return r
}

func (r ApiCreateConferenceCallPostRequest) Attachment4Filename(attachment4Filename string) ApiCreateConferenceCallPostRequest {
	r.attachment4Filename = &attachment4Filename
	return r
}

func (r ApiCreateConferenceCallPostRequest) Attachment5Filename(attachment5Filename string) ApiCreateConferenceCallPostRequest {
	r.attachment5Filename = &attachment5Filename
	return r
}

func (r ApiCreateConferenceCallPostRequest) Attachment6Filename(attachment6Filename string) ApiCreateConferenceCallPostRequest {
	r.attachment6Filename = &attachment6Filename
	return r
}

func (r ApiCreateConferenceCallPostRequest) Attachment7Filename(attachment7Filename string) ApiCreateConferenceCallPostRequest {
	r.attachment7Filename = &attachment7Filename
	return r
}

func (r ApiCreateConferenceCallPostRequest) Attachment8Filename(attachment8Filename string) ApiCreateConferenceCallPostRequest {
	r.attachment8Filename = &attachment8Filename
	return r
}

func (r ApiCreateConferenceCallPostRequest) Attachment9Filename(attachment9Filename string) ApiCreateConferenceCallPostRequest {
	r.attachment9Filename = &attachment9Filename
	return r
}

func (r ApiCreateConferenceCallPostRequest) AttachmentFilename(attachmentFilename string) ApiCreateConferenceCallPostRequest {
	r.attachmentFilename = &attachmentFilename
	return r
}

func (r ApiCreateConferenceCallPostRequest) CallType(callType string) ApiCreateConferenceCallPostRequest {
	r.callType = &callType
	return r
}

func (r ApiCreateConferenceCallPostRequest) CategoryId(categoryId int64) ApiCreateConferenceCallPostRequest {
	r.categoryId = &categoryId
	return r
}

func (r ApiCreateConferenceCallPostRequest) Color(color int32) ApiCreateConferenceCallPostRequest {
	r.color = &color
	return r
}

func (r ApiCreateConferenceCallPostRequest) FontSize(fontSize int32) ApiCreateConferenceCallPostRequest {
	r.fontSize = &fontSize
	return r
}

func (r ApiCreateConferenceCallPostRequest) GameTitle(gameTitle string) ApiCreateConferenceCallPostRequest {
	r.gameTitle = &gameTitle
	return r
}

func (r ApiCreateConferenceCallPostRequest) GroupId(groupId int64) ApiCreateConferenceCallPostRequest {
	r.groupId = &groupId
	return r
}

func (r ApiCreateConferenceCallPostRequest) JoinableBy(joinableBy string) ApiCreateConferenceCallPostRequest {
	r.joinableBy = &joinableBy
	return r
}

func (r ApiCreateConferenceCallPostRequest) Language(language string) ApiCreateConferenceCallPostRequest {
	r.language = &language
	return r
}

// Unresolved Java type: pq.g0
func (r ApiCreateConferenceCallPostRequest) MessageTags(messageTags map[string]interface{}) ApiCreateConferenceCallPostRequest {
	r.messageTags = &messageTags
	return r
}

func (r ApiCreateConferenceCallPostRequest) Text(text string) ApiCreateConferenceCallPostRequest {
	r.text = &text
	return r
}

func (r ApiCreateConferenceCallPostRequest) Execute() (*CreatePostResponse, *http.Response, error) {
	return r.ApiService.CreateConferenceCallPostExecute(r)
}

/*
CreateConferenceCallPost Method for CreateConferenceCallPost

Implemented as a Kotlin suspend function on Android.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiCreateConferenceCallPostRequest
*/
func (a *PostsAPIService) CreateConferenceCallPost(ctx context.Context) ApiCreateConferenceCallPostRequest {
	return ApiCreateConferenceCallPostRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return CreatePostResponse
func (a *PostsAPIService) CreateConferenceCallPostExecute(r ApiCreateConferenceCallPostRequest) (*CreatePostResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *CreatePostResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PostsAPIService.CreateConferenceCallPost")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v2/posts/new_conference_call"

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
	parameterAddToHeaderOrQuery(localVarFormParams, "api_key", r.apiKey, "", "")
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
	if r.callType != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "call_type", r.callType, "", "")
	}
	if r.categoryId != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "category_id", r.categoryId, "", "")
	}
	if r.color != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "color", r.color, "", "")
	}
	if r.fontSize != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "font_size", r.fontSize, "", "")
	}
	if r.gameTitle != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "game_title", r.gameTitle, "", "")
	}
	if r.groupId != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "group_id", r.groupId, "", "")
	}
	if r.joinableBy != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "joinable_by", r.joinableBy, "", "")
	}
	if r.language != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "language", r.language, "", "")
	}
	if r.messageTags != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "message_tags", r.messageTags, "", "")
	}
	parameterAddToHeaderOrQuery(localVarFormParams, "signed_info", r.signedInfo, "", "")
	if r.text != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "text", r.text, "", "")
	}
	parameterAddToHeaderOrQuery(localVarFormParams, "timestamp", r.timestamp, "", "")
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

type ApiCreatePostRequest struct {
	ctx context.Context
	ApiService *PostsAPIService
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

func (r ApiCreatePostRequest) XJwt(xJwt string) ApiCreatePostRequest {
	r.xJwt = &xJwt
	return r
}

func (r ApiCreatePostRequest) Attachment2Filename(attachment2Filename string) ApiCreatePostRequest {
	r.attachment2Filename = &attachment2Filename
	return r
}

func (r ApiCreatePostRequest) Attachment3Filename(attachment3Filename string) ApiCreatePostRequest {
	r.attachment3Filename = &attachment3Filename
	return r
}

func (r ApiCreatePostRequest) Attachment4Filename(attachment4Filename string) ApiCreatePostRequest {
	r.attachment4Filename = &attachment4Filename
	return r
}

func (r ApiCreatePostRequest) Attachment5Filename(attachment5Filename string) ApiCreatePostRequest {
	r.attachment5Filename = &attachment5Filename
	return r
}

func (r ApiCreatePostRequest) Attachment6Filename(attachment6Filename string) ApiCreatePostRequest {
	r.attachment6Filename = &attachment6Filename
	return r
}

func (r ApiCreatePostRequest) Attachment7Filename(attachment7Filename string) ApiCreatePostRequest {
	r.attachment7Filename = &attachment7Filename
	return r
}

func (r ApiCreatePostRequest) Attachment8Filename(attachment8Filename string) ApiCreatePostRequest {
	r.attachment8Filename = &attachment8Filename
	return r
}

func (r ApiCreatePostRequest) Attachment9Filename(attachment9Filename string) ApiCreatePostRequest {
	r.attachment9Filename = &attachment9Filename
	return r
}

func (r ApiCreatePostRequest) AttachmentFilename(attachmentFilename string) ApiCreatePostRequest {
	r.attachmentFilename = &attachmentFilename
	return r
}

func (r ApiCreatePostRequest) Choices(choices []string) ApiCreatePostRequest {
	r.choices = &choices
	return r
}

func (r ApiCreatePostRequest) Color(color int32) ApiCreatePostRequest {
	r.color = &color
	return r
}

func (r ApiCreatePostRequest) FontSize(fontSize int32) ApiCreatePostRequest {
	r.fontSize = &fontSize
	return r
}

func (r ApiCreatePostRequest) GroupId(groupId int64) ApiCreatePostRequest {
	r.groupId = &groupId
	return r
}

func (r ApiCreatePostRequest) InReplyTo(inReplyTo int64) ApiCreatePostRequest {
	r.inReplyTo = &inReplyTo
	return r
}

func (r ApiCreatePostRequest) Language(language string) ApiCreatePostRequest {
	r.language = &language
	return r
}

func (r ApiCreatePostRequest) MentionIds(mentionIds []int64) ApiCreatePostRequest {
	r.mentionIds = &mentionIds
	return r
}

// Unresolved Java type: pq.g0
func (r ApiCreatePostRequest) MessageTags(messageTags map[string]interface{}) ApiCreatePostRequest {
	r.messageTags = &messageTags
	return r
}

func (r ApiCreatePostRequest) PostType(postType PostType) ApiCreatePostRequest {
	r.postType = &postType
	return r
}

// Unresolved Java type: pq.g0
func (r ApiCreatePostRequest) SharedUrl(sharedUrl map[string]interface{}) ApiCreatePostRequest {
	r.sharedUrl = &sharedUrl
	return r
}

func (r ApiCreatePostRequest) Text(text string) ApiCreatePostRequest {
	r.text = &text
	return r
}

func (r ApiCreatePostRequest) VideoFileName(videoFileName string) ApiCreatePostRequest {
	r.videoFileName = &videoFileName
	return r
}

func (r ApiCreatePostRequest) Execute() (*Post, *http.Response, error) {
	return r.ApiService.CreatePostExecute(r)
}

/*
CreatePost Method for CreatePost

Implemented as a Kotlin suspend function on Android.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiCreatePostRequest
*/
func (a *PostsAPIService) CreatePost(ctx context.Context) ApiCreatePostRequest {
	return ApiCreatePostRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return Post
func (a *PostsAPIService) CreatePostExecute(r ApiCreatePostRequest) (*Post, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *Post
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PostsAPIService.CreatePost")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v3/posts/new"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.xJwt == nil {
		return localVarReturnValue, nil, reportError("xJwt is required and must be specified")
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
	parameterAddToHeaderOrQuery(localVarHeaderParams, "X-Jwt", r.xJwt, "simple", "")
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

type ApiCreateSharePostRequest struct {
	ctx context.Context
	ApiService *PostsAPIService
	apiKey *string
	shareableId *int64
	shareableType *string
	signedInfo *string
	timestamp *int64
	uuid *string
	color *int32
	fontSize *int32
	groupId *int64
	language *string
	messageTags *map[string]interface{}
	text *string
}

func (r ApiCreateSharePostRequest) ApiKey(apiKey string) ApiCreateSharePostRequest {
	r.apiKey = &apiKey
	return r
}

func (r ApiCreateSharePostRequest) ShareableId(shareableId int64) ApiCreateSharePostRequest {
	r.shareableId = &shareableId
	return r
}

func (r ApiCreateSharePostRequest) ShareableType(shareableType string) ApiCreateSharePostRequest {
	r.shareableType = &shareableType
	return r
}

func (r ApiCreateSharePostRequest) SignedInfo(signedInfo string) ApiCreateSharePostRequest {
	r.signedInfo = &signedInfo
	return r
}

func (r ApiCreateSharePostRequest) Timestamp(timestamp int64) ApiCreateSharePostRequest {
	r.timestamp = &timestamp
	return r
}

func (r ApiCreateSharePostRequest) Uuid(uuid string) ApiCreateSharePostRequest {
	r.uuid = &uuid
	return r
}

func (r ApiCreateSharePostRequest) Color(color int32) ApiCreateSharePostRequest {
	r.color = &color
	return r
}

func (r ApiCreateSharePostRequest) FontSize(fontSize int32) ApiCreateSharePostRequest {
	r.fontSize = &fontSize
	return r
}

func (r ApiCreateSharePostRequest) GroupId(groupId int64) ApiCreateSharePostRequest {
	r.groupId = &groupId
	return r
}

func (r ApiCreateSharePostRequest) Language(language string) ApiCreateSharePostRequest {
	r.language = &language
	return r
}

// Unresolved Java type: pq.g0
func (r ApiCreateSharePostRequest) MessageTags(messageTags map[string]interface{}) ApiCreateSharePostRequest {
	r.messageTags = &messageTags
	return r
}

func (r ApiCreateSharePostRequest) Text(text string) ApiCreateSharePostRequest {
	r.text = &text
	return r
}

func (r ApiCreateSharePostRequest) Execute() (*Post, *http.Response, error) {
	return r.ApiService.CreateSharePostExecute(r)
}

/*
CreateSharePost Method for CreateSharePost

Implemented as a Kotlin suspend function on Android.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiCreateSharePostRequest
*/
func (a *PostsAPIService) CreateSharePost(ctx context.Context) ApiCreateSharePostRequest {
	return ApiCreateSharePostRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return Post
func (a *PostsAPIService) CreateSharePostExecute(r ApiCreateSharePostRequest) (*Post, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *Post
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PostsAPIService.CreateSharePost")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v2/posts/new_share_post"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.apiKey == nil {
		return localVarReturnValue, nil, reportError("apiKey is required and must be specified")
	}
	if r.shareableId == nil {
		return localVarReturnValue, nil, reportError("shareableId is required and must be specified")
	}
	if r.shareableType == nil {
		return localVarReturnValue, nil, reportError("shareableType is required and must be specified")
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
	parameterAddToHeaderOrQuery(localVarFormParams, "api_key", r.apiKey, "", "")
	if r.color != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "color", r.color, "", "")
	}
	if r.fontSize != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "font_size", r.fontSize, "", "")
	}
	if r.groupId != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "group_id", r.groupId, "", "")
	}
	if r.language != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "language", r.language, "", "")
	}
	if r.messageTags != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "message_tags", r.messageTags, "", "")
	}
	parameterAddToHeaderOrQuery(localVarFormParams, "shareable_id", r.shareableId, "", "")
	parameterAddToHeaderOrQuery(localVarFormParams, "shareable_type", r.shareableType, "", "")
	parameterAddToHeaderOrQuery(localVarFormParams, "signed_info", r.signedInfo, "", "")
	if r.text != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "text", r.text, "", "")
	}
	parameterAddToHeaderOrQuery(localVarFormParams, "timestamp", r.timestamp, "", "")
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

type ApiDeleteAllPostsRequest struct {
	ctx context.Context
	ApiService *PostsAPIService
}

func (r ApiDeleteAllPostsRequest) Execute() (*http.Response, error) {
	return r.ApiService.DeleteAllPostsExecute(r)
}

/*
DeleteAllPosts Method for DeleteAllPosts

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiDeleteAllPostsRequest
*/
func (a *PostsAPIService) DeleteAllPosts(ctx context.Context) ApiDeleteAllPostsRequest {
	return ApiDeleteAllPostsRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
func (a *PostsAPIService) DeleteAllPostsExecute(r ApiDeleteAllPostsRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PostsAPIService.DeleteAllPosts")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/posts/delete_all_post"

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

type ApiDeletePostsRequest struct {
	ctx context.Context
	ApiService *PostsAPIService
	postsIds *[]int64
}

func (r ApiDeletePostsRequest) PostsIds(postsIds []int64) ApiDeletePostsRequest {
	r.postsIds = &postsIds
	return r
}

func (r ApiDeletePostsRequest) Execute() (*http.Response, error) {
	return r.ApiService.DeletePostsExecute(r)
}

/*
DeletePosts Method for DeletePosts

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiDeletePostsRequest
*/
func (a *PostsAPIService) DeletePosts(ctx context.Context) ApiDeletePostsRequest {
	return ApiDeletePostsRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
func (a *PostsAPIService) DeletePostsExecute(r ApiDeletePostsRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PostsAPIService.DeletePosts")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v2/posts/mass_destroy"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.postsIds == nil {
		return nil, reportError("postsIds is required and must be specified")
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
	parameterAddToHeaderOrQuery(localVarFormParams, "posts_ids[]", r.postsIds, "", "csv")
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

type ApiGetActiveCallPostRequest struct {
	ctx context.Context
	ApiService *PostsAPIService
	userId *int64
}

func (r ApiGetActiveCallPostRequest) UserId(userId int64) ApiGetActiveCallPostRequest {
	r.userId = &userId
	return r
}

func (r ApiGetActiveCallPostRequest) Execute() (*PostResponse, *http.Response, error) {
	return r.ApiService.GetActiveCallPostExecute(r)
}

/*
GetActiveCallPost Method for GetActiveCallPost

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiGetActiveCallPostRequest
*/
func (a *PostsAPIService) GetActiveCallPost(ctx context.Context) ApiGetActiveCallPostRequest {
	return ApiGetActiveCallPostRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return PostResponse
func (a *PostsAPIService) GetActiveCallPostExecute(r ApiGetActiveCallPostRequest) (*PostResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *PostResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PostsAPIService.GetActiveCallPost")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/posts/active_call"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.userId == nil {
		return localVarReturnValue, nil, reportError("userId is required and must be specified")
	}

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

type ApiGetCallFollowersTimelineRequest struct {
	ctx context.Context
	ApiService *PostsAPIService
	fromTimestamp *int64
	number *int32
	categoryId *int64
	callType *string
	includeCircleCall *bool
	excludeRecentGomimushi *bool
}

func (r ApiGetCallFollowersTimelineRequest) FromTimestamp(fromTimestamp int64) ApiGetCallFollowersTimelineRequest {
	r.fromTimestamp = &fromTimestamp
	return r
}

func (r ApiGetCallFollowersTimelineRequest) Number(number int32) ApiGetCallFollowersTimelineRequest {
	r.number = &number
	return r
}

func (r ApiGetCallFollowersTimelineRequest) CategoryId(categoryId int64) ApiGetCallFollowersTimelineRequest {
	r.categoryId = &categoryId
	return r
}

func (r ApiGetCallFollowersTimelineRequest) CallType(callType string) ApiGetCallFollowersTimelineRequest {
	r.callType = &callType
	return r
}

func (r ApiGetCallFollowersTimelineRequest) IncludeCircleCall(includeCircleCall bool) ApiGetCallFollowersTimelineRequest {
	r.includeCircleCall = &includeCircleCall
	return r
}

func (r ApiGetCallFollowersTimelineRequest) ExcludeRecentGomimushi(excludeRecentGomimushi bool) ApiGetCallFollowersTimelineRequest {
	r.excludeRecentGomimushi = &excludeRecentGomimushi
	return r
}

func (r ApiGetCallFollowersTimelineRequest) Execute() (*PostsResponse, *http.Response, error) {
	return r.ApiService.GetCallFollowersTimelineExecute(r)
}

/*
GetCallFollowersTimeline Method for GetCallFollowersTimeline

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiGetCallFollowersTimelineRequest
*/
func (a *PostsAPIService) GetCallFollowersTimeline(ctx context.Context) ApiGetCallFollowersTimelineRequest {
	return ApiGetCallFollowersTimelineRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return PostsResponse
func (a *PostsAPIService) GetCallFollowersTimelineExecute(r ApiGetCallFollowersTimelineRequest) (*PostsResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *PostsResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PostsAPIService.GetCallFollowersTimeline")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v2/posts/call_followers_timeline"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.fromTimestamp != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "from_timestamp", r.fromTimestamp, "form", "")
	}
	if r.number != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "number", r.number, "form", "")
	}
	if r.categoryId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "category_id", r.categoryId, "form", "")
	}
	if r.callType != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "call_type", r.callType, "form", "")
	}
	if r.includeCircleCall != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "include_circle_call", r.includeCircleCall, "form", "")
	}
	if r.excludeRecentGomimushi != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "exclude_recent_gomimushi", r.excludeRecentGomimushi, "form", "")
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

type ApiGetCallTimelineRequest struct {
	ctx context.Context
	ApiService *PostsAPIService
	groupId *int64
	fromTimestamp *int64
	number *int32
	categoryId *int64
	callType *string
	includeCircleCall *bool
	crossGeneration *bool
	excludeRecentGomimushi *bool
}

func (r ApiGetCallTimelineRequest) GroupId(groupId int64) ApiGetCallTimelineRequest {
	r.groupId = &groupId
	return r
}

func (r ApiGetCallTimelineRequest) FromTimestamp(fromTimestamp int64) ApiGetCallTimelineRequest {
	r.fromTimestamp = &fromTimestamp
	return r
}

func (r ApiGetCallTimelineRequest) Number(number int32) ApiGetCallTimelineRequest {
	r.number = &number
	return r
}

func (r ApiGetCallTimelineRequest) CategoryId(categoryId int64) ApiGetCallTimelineRequest {
	r.categoryId = &categoryId
	return r
}

func (r ApiGetCallTimelineRequest) CallType(callType string) ApiGetCallTimelineRequest {
	r.callType = &callType
	return r
}

func (r ApiGetCallTimelineRequest) IncludeCircleCall(includeCircleCall bool) ApiGetCallTimelineRequest {
	r.includeCircleCall = &includeCircleCall
	return r
}

func (r ApiGetCallTimelineRequest) CrossGeneration(crossGeneration bool) ApiGetCallTimelineRequest {
	r.crossGeneration = &crossGeneration
	return r
}

func (r ApiGetCallTimelineRequest) ExcludeRecentGomimushi(excludeRecentGomimushi bool) ApiGetCallTimelineRequest {
	r.excludeRecentGomimushi = &excludeRecentGomimushi
	return r
}

func (r ApiGetCallTimelineRequest) Execute() (*PostsResponse, *http.Response, error) {
	return r.ApiService.GetCallTimelineExecute(r)
}

/*
GetCallTimeline Method for GetCallTimeline

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiGetCallTimelineRequest
*/
func (a *PostsAPIService) GetCallTimeline(ctx context.Context) ApiGetCallTimelineRequest {
	return ApiGetCallTimelineRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return PostsResponse
func (a *PostsAPIService) GetCallTimelineExecute(r ApiGetCallTimelineRequest) (*PostsResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *PostsResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PostsAPIService.GetCallTimeline")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v2/posts/call_timeline"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.groupId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "group_id", r.groupId, "form", "")
	}
	if r.fromTimestamp != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "from_timestamp", r.fromTimestamp, "form", "")
	}
	if r.number != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "number", r.number, "form", "")
	}
	if r.categoryId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "category_id", r.categoryId, "form", "")
	}
	if r.callType != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "call_type", r.callType, "form", "")
	}
	if r.includeCircleCall != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "include_circle_call", r.includeCircleCall, "form", "")
	}
	if r.crossGeneration != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "cross_generation", r.crossGeneration, "form", "")
	}
	if r.excludeRecentGomimushi != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "exclude_recent_gomimushi", r.excludeRecentGomimushi, "form", "")
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

type ApiGetFollowingTimelineRequest struct {
	ctx context.Context
	ApiService *PostsAPIService
	orderBy *string
	from *string
	fromPostId *int64
	onlyRoot *bool
	number *int32
	mxn *int32
	reduceSelfie *bool
	customGenerationRange *bool
}

func (r ApiGetFollowingTimelineRequest) OrderBy(orderBy string) ApiGetFollowingTimelineRequest {
	r.orderBy = &orderBy
	return r
}

func (r ApiGetFollowingTimelineRequest) From(from string) ApiGetFollowingTimelineRequest {
	r.from = &from
	return r
}

func (r ApiGetFollowingTimelineRequest) FromPostId(fromPostId int64) ApiGetFollowingTimelineRequest {
	r.fromPostId = &fromPostId
	return r
}

func (r ApiGetFollowingTimelineRequest) OnlyRoot(onlyRoot bool) ApiGetFollowingTimelineRequest {
	r.onlyRoot = &onlyRoot
	return r
}

func (r ApiGetFollowingTimelineRequest) Number(number int32) ApiGetFollowingTimelineRequest {
	r.number = &number
	return r
}

func (r ApiGetFollowingTimelineRequest) Mxn(mxn int32) ApiGetFollowingTimelineRequest {
	r.mxn = &mxn
	return r
}

func (r ApiGetFollowingTimelineRequest) ReduceSelfie(reduceSelfie bool) ApiGetFollowingTimelineRequest {
	r.reduceSelfie = &reduceSelfie
	return r
}

func (r ApiGetFollowingTimelineRequest) CustomGenerationRange(customGenerationRange bool) ApiGetFollowingTimelineRequest {
	r.customGenerationRange = &customGenerationRange
	return r
}

func (r ApiGetFollowingTimelineRequest) Execute() (*PostsResponse, *http.Response, error) {
	return r.ApiService.GetFollowingTimelineExecute(r)
}

/*
GetFollowingTimeline Method for GetFollowingTimeline

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiGetFollowingTimelineRequest
*/
func (a *PostsAPIService) GetFollowingTimeline(ctx context.Context) ApiGetFollowingTimelineRequest {
	return ApiGetFollowingTimelineRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return PostsResponse
func (a *PostsAPIService) GetFollowingTimelineExecute(r ApiGetFollowingTimelineRequest) (*PostsResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *PostsResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PostsAPIService.GetFollowingTimeline")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v2/posts/following_timeline"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.orderBy == nil {
		return localVarReturnValue, nil, reportError("orderBy is required and must be specified")
	}

	if r.from != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "from", r.from, "form", "")
	}
	if r.fromPostId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "from_post_id", r.fromPostId, "form", "")
	}
	if r.onlyRoot != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "only_root", r.onlyRoot, "form", "")
	}
	parameterAddToHeaderOrQuery(localVarQueryParams, "order_by", r.orderBy, "form", "")
	if r.number != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "number", r.number, "form", "")
	}
	if r.mxn != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "mxn", r.mxn, "form", "")
	}
	if r.reduceSelfie != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "reduce_selfie", r.reduceSelfie, "form", "")
	}
	if r.customGenerationRange != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "custom_generation_range", r.customGenerationRange, "form", "")
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

type ApiGetGroupTimelineRequest struct {
	ctx context.Context
	ApiService *PostsAPIService
	groupId *int64
	fromPostId *int64
	reverse *bool
	postType *PostType
	number *int32
	onlyRoot *bool
}

func (r ApiGetGroupTimelineRequest) GroupId(groupId int64) ApiGetGroupTimelineRequest {
	r.groupId = &groupId
	return r
}

func (r ApiGetGroupTimelineRequest) FromPostId(fromPostId int64) ApiGetGroupTimelineRequest {
	r.fromPostId = &fromPostId
	return r
}

func (r ApiGetGroupTimelineRequest) Reverse(reverse bool) ApiGetGroupTimelineRequest {
	r.reverse = &reverse
	return r
}

func (r ApiGetGroupTimelineRequest) PostType(postType PostType) ApiGetGroupTimelineRequest {
	r.postType = &postType
	return r
}

func (r ApiGetGroupTimelineRequest) Number(number int32) ApiGetGroupTimelineRequest {
	r.number = &number
	return r
}

func (r ApiGetGroupTimelineRequest) OnlyRoot(onlyRoot bool) ApiGetGroupTimelineRequest {
	r.onlyRoot = &onlyRoot
	return r
}

func (r ApiGetGroupTimelineRequest) Execute() (*PostsResponse, *http.Response, error) {
	return r.ApiService.GetGroupTimelineExecute(r)
}

/*
GetGroupTimeline Method for GetGroupTimeline

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiGetGroupTimelineRequest
*/
func (a *PostsAPIService) GetGroupTimeline(ctx context.Context) ApiGetGroupTimelineRequest {
	return ApiGetGroupTimelineRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return PostsResponse
func (a *PostsAPIService) GetGroupTimelineExecute(r ApiGetGroupTimelineRequest) (*PostsResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *PostsResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PostsAPIService.GetGroupTimeline")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v2/posts/group_timeline"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.groupId == nil {
		return localVarReturnValue, nil, reportError("groupId is required and must be specified")
	}

	parameterAddToHeaderOrQuery(localVarQueryParams, "group_id", r.groupId, "form", "")
	if r.fromPostId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "from_post_id", r.fromPostId, "form", "")
	}
	if r.reverse != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "reverse", r.reverse, "form", "")
	}
	if r.postType != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "post_type", r.postType, "form", "")
	}
	if r.number != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "number", r.number, "form", "")
	}
	if r.onlyRoot != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "only_root", r.onlyRoot, "form", "")
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

type ApiGetMyPostsRequest struct {
	ctx context.Context
	ApiService *PostsAPIService
	number *int32
	fromPostId *int64
	includeGroupPost *bool
}

func (r ApiGetMyPostsRequest) Number(number int32) ApiGetMyPostsRequest {
	r.number = &number
	return r
}

func (r ApiGetMyPostsRequest) FromPostId(fromPostId int64) ApiGetMyPostsRequest {
	r.fromPostId = &fromPostId
	return r
}

func (r ApiGetMyPostsRequest) IncludeGroupPost(includeGroupPost bool) ApiGetMyPostsRequest {
	r.includeGroupPost = &includeGroupPost
	return r
}

func (r ApiGetMyPostsRequest) Execute() (*PostsResponse, *http.Response, error) {
	return r.ApiService.GetMyPostsExecute(r)
}

/*
GetMyPosts Method for GetMyPosts

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiGetMyPostsRequest
*/
func (a *PostsAPIService) GetMyPosts(ctx context.Context) ApiGetMyPostsRequest {
	return ApiGetMyPostsRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return PostsResponse
func (a *PostsAPIService) GetMyPostsExecute(r ApiGetMyPostsRequest) (*PostsResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *PostsResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PostsAPIService.GetMyPosts")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v2/posts/mine"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.number == nil {
		return localVarReturnValue, nil, reportError("number is required and must be specified")
	}

	if r.fromPostId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "from_post_id", r.fromPostId, "form", "")
	}
	parameterAddToHeaderOrQuery(localVarQueryParams, "number", r.number, "form", "")
	if r.includeGroupPost != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "include_group_post", r.includeGroupPost, "form", "")
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

type ApiGetPostRequest struct {
	ctx context.Context
	ApiService *PostsAPIService
	id int64
	cacheControl *string
}

func (r ApiGetPostRequest) CacheControl(cacheControl string) ApiGetPostRequest {
	r.cacheControl = &cacheControl
	return r
}

func (r ApiGetPostRequest) Execute() (*PostResponse, *http.Response, error) {
	return r.ApiService.GetPostExecute(r)
}

/*
GetPost Method for GetPost

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id
 @return ApiGetPostRequest
*/
func (a *PostsAPIService) GetPost(ctx context.Context, id int64) ApiGetPostRequest {
	return ApiGetPostRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
//  @return PostResponse
func (a *PostsAPIService) GetPostExecute(r ApiGetPostRequest) (*PostResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *PostResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PostsAPIService.GetPost")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v2/posts/{id}"
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
	if r.cacheControl != nil {
		parameterAddToHeaderOrQuery(localVarHeaderParams, "Cache-Control", r.cacheControl, "simple", "")
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

type ApiGetPostGiftTransactionsRequest struct {
	ctx context.Context
	ApiService *PostsAPIService
	postId int64
	number *int32
	from *string
}

func (r ApiGetPostGiftTransactionsRequest) Number(number int32) ApiGetPostGiftTransactionsRequest {
	r.number = &number
	return r
}

func (r ApiGetPostGiftTransactionsRequest) From(from string) ApiGetPostGiftTransactionsRequest {
	r.from = &from
	return r
}

func (r ApiGetPostGiftTransactionsRequest) Execute() (*GiftTransactionsResponse, *http.Response, error) {
	return r.ApiService.GetPostGiftTransactionsExecute(r)
}

/*
GetPostGiftTransactions Method for GetPostGiftTransactions

Implemented as a Kotlin suspend function on Android.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param postId
 @return ApiGetPostGiftTransactionsRequest
*/
func (a *PostsAPIService) GetPostGiftTransactions(ctx context.Context, postId int64) ApiGetPostGiftTransactionsRequest {
	return ApiGetPostGiftTransactionsRequest{
		ApiService: a,
		ctx: ctx,
		postId: postId,
	}
}

// Execute executes the request
//  @return GiftTransactionsResponse
func (a *PostsAPIService) GetPostGiftTransactionsExecute(r ApiGetPostGiftTransactionsRequest) (*GiftTransactionsResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *GiftTransactionsResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PostsAPIService.GetPostGiftTransactions")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/posts/{post_id}/gift_transactions"
	localVarPath = strings.Replace(localVarPath, "{"+"post_id"+"}", url.PathEscape(parameterValueToString(r.postId, "postId")), -1)

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

type ApiGetPostLikersRequest struct {
	ctx context.Context
	ApiService *PostsAPIService
	id int64
	fromId *int64
}

func (r ApiGetPostLikersRequest) FromId(fromId int64) ApiGetPostLikersRequest {
	r.fromId = &fromId
	return r
}

func (r ApiGetPostLikersRequest) Execute() (*PostLikersResponse, *http.Response, error) {
	return r.ApiService.GetPostLikersExecute(r)
}

/*
GetPostLikers Method for GetPostLikers

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id
 @return ApiGetPostLikersRequest
*/
func (a *PostsAPIService) GetPostLikers(ctx context.Context, id int64) ApiGetPostLikersRequest {
	return ApiGetPostLikersRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
//  @return PostLikersResponse
func (a *PostsAPIService) GetPostLikersExecute(r ApiGetPostLikersRequest) (*PostLikersResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *PostLikersResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PostsAPIService.GetPostLikers")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/posts/{id}/likers"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterValueToString(r.id, "id")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

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

type ApiGetPostRepostsRequest struct {
	ctx context.Context
	ApiService *PostsAPIService
	id int64
	fromPostId *int64
}

func (r ApiGetPostRepostsRequest) FromPostId(fromPostId int64) ApiGetPostRepostsRequest {
	r.fromPostId = &fromPostId
	return r
}

func (r ApiGetPostRepostsRequest) Execute() (*PostsResponse, *http.Response, error) {
	return r.ApiService.GetPostRepostsExecute(r)
}

/*
GetPostReposts Method for GetPostReposts

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id
 @return ApiGetPostRepostsRequest
*/
func (a *PostsAPIService) GetPostReposts(ctx context.Context, id int64) ApiGetPostRepostsRequest {
	return ApiGetPostRepostsRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
//  @return PostsResponse
func (a *PostsAPIService) GetPostRepostsExecute(r ApiGetPostRepostsRequest) (*PostsResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *PostsResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PostsAPIService.GetPostReposts")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v2/posts/{id}/reposts"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterValueToString(r.id, "id")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.fromPostId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "from_post_id", r.fromPostId, "form", "")
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

type ApiGetPostUrlMetadataRequest struct {
	ctx context.Context
	ApiService *PostsAPIService
	url *string
}

func (r ApiGetPostUrlMetadataRequest) Url(url string) ApiGetPostUrlMetadataRequest {
	r.url = &url
	return r
}

func (r ApiGetPostUrlMetadataRequest) Execute() (*SharedUrl, *http.Response, error) {
	return r.ApiService.GetPostUrlMetadataExecute(r)
}

/*
GetPostUrlMetadata Method for GetPostUrlMetadata

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiGetPostUrlMetadataRequest
*/
func (a *PostsAPIService) GetPostUrlMetadata(ctx context.Context) ApiGetPostUrlMetadataRequest {
	return ApiGetPostUrlMetadataRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return SharedUrl
func (a *PostsAPIService) GetPostUrlMetadataExecute(r ApiGetPostUrlMetadataRequest) (*SharedUrl, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *SharedUrl
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PostsAPIService.GetPostUrlMetadata")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v2/posts/url_metadata"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.url == nil {
		return localVarReturnValue, nil, reportError("url is required and must be specified")
	}

	parameterAddToHeaderOrQuery(localVarQueryParams, "url", r.url, "form", "")
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

type ApiGetPostsByIdsRequest struct {
	ctx context.Context
	ApiService *PostsAPIService
	postIds *[]int64
}

func (r ApiGetPostsByIdsRequest) PostIds(postIds []int64) ApiGetPostsByIdsRequest {
	r.postIds = &postIds
	return r
}

func (r ApiGetPostsByIdsRequest) Execute() (*PostsResponse, *http.Response, error) {
	return r.ApiService.GetPostsByIdsExecute(r)
}

/*
GetPostsByIds Method for GetPostsByIds

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiGetPostsByIdsRequest
*/
func (a *PostsAPIService) GetPostsByIds(ctx context.Context) ApiGetPostsByIdsRequest {
	return ApiGetPostsByIdsRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return PostsResponse
func (a *PostsAPIService) GetPostsByIdsExecute(r ApiGetPostsByIdsRequest) (*PostsResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *PostsResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PostsAPIService.GetPostsByIds")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v2/posts/multiple"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.postIds == nil {
		return localVarReturnValue, nil, reportError("postIds is required and must be specified")
	}

	{
		t := *r.postIds
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				parameterAddToHeaderOrQuery(localVarQueryParams, "post_ids[]", s.Index(i).Interface(), "form", "multi")
			}
		} else {
			parameterAddToHeaderOrQuery(localVarQueryParams, "post_ids[]", t, "form", "multi")
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

type ApiGetPostsByTagRequest struct {
	ctx context.Context
	ApiService *PostsAPIService
	tag string
	fromPostId *int64
	number *int32
}

func (r ApiGetPostsByTagRequest) FromPostId(fromPostId int64) ApiGetPostsByTagRequest {
	r.fromPostId = &fromPostId
	return r
}

func (r ApiGetPostsByTagRequest) Number(number int32) ApiGetPostsByTagRequest {
	r.number = &number
	return r
}

func (r ApiGetPostsByTagRequest) Execute() (*PostsResponse, *http.Response, error) {
	return r.ApiService.GetPostsByTagExecute(r)
}

/*
GetPostsByTag Method for GetPostsByTag

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param tag
 @return ApiGetPostsByTagRequest
*/
func (a *PostsAPIService) GetPostsByTag(ctx context.Context, tag string) ApiGetPostsByTagRequest {
	return ApiGetPostsByTagRequest{
		ApiService: a,
		ctx: ctx,
		tag: tag,
	}
}

// Execute executes the request
//  @return PostsResponse
func (a *PostsAPIService) GetPostsByTagExecute(r ApiGetPostsByTagRequest) (*PostsResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *PostsResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PostsAPIService.GetPostsByTag")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v2/posts/tags/{tag}"
	localVarPath = strings.Replace(localVarPath, "{"+"tag"+"}", url.PathEscape(parameterValueToString(r.tag, "tag")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.fromPostId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "from_post_id", r.fromPostId, "form", "")
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

type ApiGetRecentEngagementPostsRequest struct {
	ctx context.Context
	ApiService *PostsAPIService
	number *int32
}

func (r ApiGetRecentEngagementPostsRequest) Number(number int32) ApiGetRecentEngagementPostsRequest {
	r.number = &number
	return r
}

func (r ApiGetRecentEngagementPostsRequest) Execute() (*PostsResponse, *http.Response, error) {
	return r.ApiService.GetRecentEngagementPostsExecute(r)
}

/*
GetRecentEngagementPosts Method for GetRecentEngagementPosts

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiGetRecentEngagementPostsRequest
*/
func (a *PostsAPIService) GetRecentEngagementPosts(ctx context.Context) ApiGetRecentEngagementPostsRequest {
	return ApiGetRecentEngagementPostsRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return PostsResponse
func (a *PostsAPIService) GetRecentEngagementPostsExecute(r ApiGetRecentEngagementPostsRequest) (*PostsResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *PostsResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PostsAPIService.GetRecentEngagementPosts")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v2/posts/recent_engagement"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.number == nil {
		return localVarReturnValue, nil, reportError("number is required and must be specified")
	}

	parameterAddToHeaderOrQuery(localVarQueryParams, "number", r.number, "form", "")
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

type ApiGetRecommendedPostTagsRequest struct {
	ctx context.Context
	ApiService *PostsAPIService
	saveRecentSearch *bool
	tag *string
}

func (r ApiGetRecommendedPostTagsRequest) SaveRecentSearch(saveRecentSearch bool) ApiGetRecommendedPostTagsRequest {
	r.saveRecentSearch = &saveRecentSearch
	return r
}

func (r ApiGetRecommendedPostTagsRequest) Tag(tag string) ApiGetRecommendedPostTagsRequest {
	r.tag = &tag
	return r
}

func (r ApiGetRecommendedPostTagsRequest) Execute() (*PostTagsResponse, *http.Response, error) {
	return r.ApiService.GetRecommendedPostTagsExecute(r)
}

/*
GetRecommendedPostTags Method for GetRecommendedPostTags

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiGetRecommendedPostTagsRequest
*/
func (a *PostsAPIService) GetRecommendedPostTags(ctx context.Context) ApiGetRecommendedPostTagsRequest {
	return ApiGetRecommendedPostTagsRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return PostTagsResponse
func (a *PostsAPIService) GetRecommendedPostTagsExecute(r ApiGetRecommendedPostTagsRequest) (*PostTagsResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *PostTagsResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PostsAPIService.GetRecommendedPostTags")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/posts/recommended_tag"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.saveRecentSearch == nil {
		return localVarReturnValue, nil, reportError("saveRecentSearch is required and must be specified")
	}
	if r.tag == nil {
		return localVarReturnValue, nil, reportError("tag is required and must be specified")
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
	parameterAddToHeaderOrQuery(localVarFormParams, "save_recent_search", r.saveRecentSearch, "", "")
	parameterAddToHeaderOrQuery(localVarFormParams, "tag", r.tag, "", "")
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

type ApiGetRecommendedTimelineRequest struct {
	ctx context.Context
	ApiService *PostsAPIService
	experimentNum *int32
	variantNum *int32
	number *int32
}

func (r ApiGetRecommendedTimelineRequest) ExperimentNum(experimentNum int32) ApiGetRecommendedTimelineRequest {
	r.experimentNum = &experimentNum
	return r
}

func (r ApiGetRecommendedTimelineRequest) VariantNum(variantNum int32) ApiGetRecommendedTimelineRequest {
	r.variantNum = &variantNum
	return r
}

func (r ApiGetRecommendedTimelineRequest) Number(number int32) ApiGetRecommendedTimelineRequest {
	r.number = &number
	return r
}

func (r ApiGetRecommendedTimelineRequest) Execute() (*PostsResponse, *http.Response, error) {
	return r.ApiService.GetRecommendedTimelineExecute(r)
}

/*
GetRecommendedTimeline Method for GetRecommendedTimeline

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiGetRecommendedTimelineRequest
*/
func (a *PostsAPIService) GetRecommendedTimeline(ctx context.Context) ApiGetRecommendedTimelineRequest {
	return ApiGetRecommendedTimelineRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return PostsResponse
func (a *PostsAPIService) GetRecommendedTimelineExecute(r ApiGetRecommendedTimelineRequest) (*PostsResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *PostsResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PostsAPIService.GetRecommendedTimeline")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v2/posts/recommended_timeline"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.experimentNum == nil {
		return localVarReturnValue, nil, reportError("experimentNum is required and must be specified")
	}
	if r.variantNum == nil {
		return localVarReturnValue, nil, reportError("variantNum is required and must be specified")
	}
	if r.number == nil {
		return localVarReturnValue, nil, reportError("number is required and must be specified")
	}

	parameterAddToHeaderOrQuery(localVarQueryParams, "experiment_num", r.experimentNum, "form", "")
	parameterAddToHeaderOrQuery(localVarQueryParams, "variant_num", r.variantNum, "form", "")
	parameterAddToHeaderOrQuery(localVarQueryParams, "number", r.number, "form", "")
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

type ApiGetTimelineByModeRequest struct {
	ctx context.Context
	ApiService *PostsAPIService
	noreplyMode string
	orderBy *string
	experimentOlderAgeRules *bool
	from *string
	fromPostId *int64
	number *int32
	mxn *int32
	en *int32
	vn *int32
	reduceSelfie *bool
	customGenerationRange *bool
}

func (r ApiGetTimelineByModeRequest) OrderBy(orderBy string) ApiGetTimelineByModeRequest {
	r.orderBy = &orderBy
	return r
}

func (r ApiGetTimelineByModeRequest) ExperimentOlderAgeRules(experimentOlderAgeRules bool) ApiGetTimelineByModeRequest {
	r.experimentOlderAgeRules = &experimentOlderAgeRules
	return r
}

func (r ApiGetTimelineByModeRequest) From(from string) ApiGetTimelineByModeRequest {
	r.from = &from
	return r
}

func (r ApiGetTimelineByModeRequest) FromPostId(fromPostId int64) ApiGetTimelineByModeRequest {
	r.fromPostId = &fromPostId
	return r
}

func (r ApiGetTimelineByModeRequest) Number(number int32) ApiGetTimelineByModeRequest {
	r.number = &number
	return r
}

func (r ApiGetTimelineByModeRequest) Mxn(mxn int32) ApiGetTimelineByModeRequest {
	r.mxn = &mxn
	return r
}

func (r ApiGetTimelineByModeRequest) En(en int32) ApiGetTimelineByModeRequest {
	r.en = &en
	return r
}

func (r ApiGetTimelineByModeRequest) Vn(vn int32) ApiGetTimelineByModeRequest {
	r.vn = &vn
	return r
}

func (r ApiGetTimelineByModeRequest) ReduceSelfie(reduceSelfie bool) ApiGetTimelineByModeRequest {
	r.reduceSelfie = &reduceSelfie
	return r
}

func (r ApiGetTimelineByModeRequest) CustomGenerationRange(customGenerationRange bool) ApiGetTimelineByModeRequest {
	r.customGenerationRange = &customGenerationRange
	return r
}

func (r ApiGetTimelineByModeRequest) Execute() (*PostsResponse, *http.Response, error) {
	return r.ApiService.GetTimelineByModeExecute(r)
}

/*
GetTimelineByMode Method for GetTimelineByMode

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param noreplyMode
 @return ApiGetTimelineByModeRequest
*/
func (a *PostsAPIService) GetTimelineByMode(ctx context.Context, noreplyMode string) ApiGetTimelineByModeRequest {
	return ApiGetTimelineByModeRequest{
		ApiService: a,
		ctx: ctx,
		noreplyMode: noreplyMode,
	}
}

// Execute executes the request
//  @return PostsResponse
func (a *PostsAPIService) GetTimelineByModeExecute(r ApiGetTimelineByModeRequest) (*PostsResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *PostsResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PostsAPIService.GetTimelineByMode")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v2/posts/{noreply_mode}timeline"
	localVarPath = strings.Replace(localVarPath, "{"+"noreply_mode"+"}", url.PathEscape(parameterValueToString(r.noreplyMode, "noreplyMode")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.orderBy == nil {
		return localVarReturnValue, nil, reportError("orderBy is required and must be specified")
	}

	parameterAddToHeaderOrQuery(localVarQueryParams, "order_by", r.orderBy, "form", "")
	if r.experimentOlderAgeRules != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "experiment_older_age_rules", r.experimentOlderAgeRules, "form", "")
	}
	if r.from != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "from", r.from, "form", "")
	}
	if r.fromPostId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "from_post_id", r.fromPostId, "form", "")
	}
	if r.number != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "number", r.number, "form", "")
	}
	if r.mxn != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "mxn", r.mxn, "form", "")
	}
	if r.en != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "en", r.en, "form", "")
	}
	if r.vn != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "vn", r.vn, "form", "")
	}
	if r.reduceSelfie != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "reduce_selfie", r.reduceSelfie, "form", "")
	}
	if r.customGenerationRange != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "custom_generation_range", r.customGenerationRange, "form", "")
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

type ApiGetUserTimelineRequest struct {
	ctx context.Context
	ApiService *PostsAPIService
	userId *int64
	fromPostId *int64
	postType *PostType
	number *int32
}

func (r ApiGetUserTimelineRequest) UserId(userId int64) ApiGetUserTimelineRequest {
	r.userId = &userId
	return r
}

func (r ApiGetUserTimelineRequest) FromPostId(fromPostId int64) ApiGetUserTimelineRequest {
	r.fromPostId = &fromPostId
	return r
}

func (r ApiGetUserTimelineRequest) PostType(postType PostType) ApiGetUserTimelineRequest {
	r.postType = &postType
	return r
}

func (r ApiGetUserTimelineRequest) Number(number int32) ApiGetUserTimelineRequest {
	r.number = &number
	return r
}

func (r ApiGetUserTimelineRequest) Execute() (*PostsResponse, *http.Response, error) {
	return r.ApiService.GetUserTimelineExecute(r)
}

/*
GetUserTimeline Method for GetUserTimeline

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiGetUserTimelineRequest
*/
func (a *PostsAPIService) GetUserTimeline(ctx context.Context) ApiGetUserTimelineRequest {
	return ApiGetUserTimelineRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return PostsResponse
func (a *PostsAPIService) GetUserTimelineExecute(r ApiGetUserTimelineRequest) (*PostsResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *PostsResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PostsAPIService.GetUserTimeline")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v2/posts/user_timeline"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.userId == nil {
		return localVarReturnValue, nil, reportError("userId is required and must be specified")
	}

	parameterAddToHeaderOrQuery(localVarQueryParams, "user_id", r.userId, "form", "")
	if r.fromPostId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "from_post_id", r.fromPostId, "form", "")
	}
	if r.postType != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "post_type", r.postType, "form", "")
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

type ApiLikePostsRequest struct {
	ctx context.Context
	ApiService *PostsAPIService
	postIds *[]int64
}

func (r ApiLikePostsRequest) PostIds(postIds []int64) ApiLikePostsRequest {
	r.postIds = &postIds
	return r
}

func (r ApiLikePostsRequest) Execute() (*LikePostsResponse, *http.Response, error) {
	return r.ApiService.LikePostsExecute(r)
}

/*
LikePosts Method for LikePosts

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiLikePostsRequest
*/
func (a *PostsAPIService) LikePosts(ctx context.Context) ApiLikePostsRequest {
	return ApiLikePostsRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return LikePostsResponse
func (a *PostsAPIService) LikePostsExecute(r ApiLikePostsRequest) (*LikePostsResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *LikePostsResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PostsAPIService.LikePosts")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v2/posts/like"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.postIds == nil {
		return localVarReturnValue, nil, reportError("postIds is required and must be specified")
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
	parameterAddToHeaderOrQuery(localVarFormParams, "post_ids[]", r.postIds, "", "csv")
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

type ApiMovePostToSpecificThreadRequest struct {
	ctx context.Context
	ApiService *PostsAPIService
	id int64
	threadId int64
}

func (r ApiMovePostToSpecificThreadRequest) Execute() (*ThreadInfo, *http.Response, error) {
	return r.ApiService.MovePostToSpecificThreadExecute(r)
}

/*
MovePostToSpecificThread Method for MovePostToSpecificThread

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id
 @param threadId
 @return ApiMovePostToSpecificThreadRequest
*/
func (a *PostsAPIService) MovePostToSpecificThread(ctx context.Context, id int64, threadId int64) ApiMovePostToSpecificThreadRequest {
	return ApiMovePostToSpecificThreadRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
		threadId: threadId,
	}
}

// Execute executes the request
//  @return ThreadInfo
func (a *PostsAPIService) MovePostToSpecificThreadExecute(r ApiMovePostToSpecificThreadRequest) (*ThreadInfo, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPut
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ThreadInfo
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PostsAPIService.MovePostToSpecificThread")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v3/posts/{id}/move_to_thread/{thread_id}"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterValueToString(r.id, "id")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"thread_id"+"}", url.PathEscape(parameterValueToString(r.threadId, "threadId")), -1)

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

type ApiMovePostToThreadRequest struct {
	ctx context.Context
	ApiService *PostsAPIService
	id int64
	threadIconFilename *string
	title *string
}

func (r ApiMovePostToThreadRequest) ThreadIconFilename(threadIconFilename string) ApiMovePostToThreadRequest {
	r.threadIconFilename = &threadIconFilename
	return r
}

func (r ApiMovePostToThreadRequest) Title(title string) ApiMovePostToThreadRequest {
	r.title = &title
	return r
}

func (r ApiMovePostToThreadRequest) Execute() (*ThreadInfo, *http.Response, error) {
	return r.ApiService.MovePostToThreadExecute(r)
}

/*
MovePostToThread Method for MovePostToThread

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id
 @return ApiMovePostToThreadRequest
*/
func (a *PostsAPIService) MovePostToThread(ctx context.Context, id int64) ApiMovePostToThreadRequest {
	return ApiMovePostToThreadRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
//  @return ThreadInfo
func (a *PostsAPIService) MovePostToThreadExecute(r ApiMovePostToThreadRequest) (*ThreadInfo, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ThreadInfo
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PostsAPIService.MovePostToThread")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v3/posts/{id}/move_to_thread"
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
	if r.threadIconFilename != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "thread_icon_filename", r.threadIconFilename, "", "")
	}
	if r.title != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "title", r.title, "", "")
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

type ApiPinGroupPostRequest struct {
	ctx context.Context
	ApiService *PostsAPIService
	groupId *int64
	postId *int64
}

func (r ApiPinGroupPostRequest) GroupId(groupId int64) ApiPinGroupPostRequest {
	r.groupId = &groupId
	return r
}

func (r ApiPinGroupPostRequest) PostId(postId int64) ApiPinGroupPostRequest {
	r.postId = &postId
	return r
}

func (r ApiPinGroupPostRequest) Execute() (*http.Response, error) {
	return r.ApiService.PinGroupPostExecute(r)
}

/*
PinGroupPost Method for PinGroupPost

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiPinGroupPostRequest
*/
func (a *PostsAPIService) PinGroupPost(ctx context.Context) ApiPinGroupPostRequest {
	return ApiPinGroupPostRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
func (a *PostsAPIService) PinGroupPostExecute(r ApiPinGroupPostRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPut
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PostsAPIService.PinGroupPost")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v2/posts/group_pinned_post"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.groupId == nil {
		return nil, reportError("groupId is required and must be specified")
	}
	if r.postId == nil {
		return nil, reportError("postId is required and must be specified")
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
	parameterAddToHeaderOrQuery(localVarFormParams, "group_id", r.groupId, "", "")
	parameterAddToHeaderOrQuery(localVarFormParams, "post_id", r.postId, "", "")
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

type ApiReportPostRequest struct {
	ctx context.Context
	ApiService *PostsAPIService
	postId int64
	categoryId *int64
	opponentId *int64
	reason *string
	screenshot2Filename *string
	screenshot3Filename *string
	screenshot4Filename *string
	screenshotFilename *string
}

func (r ApiReportPostRequest) CategoryId(categoryId int64) ApiReportPostRequest {
	r.categoryId = &categoryId
	return r
}

func (r ApiReportPostRequest) OpponentId(opponentId int64) ApiReportPostRequest {
	r.opponentId = &opponentId
	return r
}

func (r ApiReportPostRequest) Reason(reason string) ApiReportPostRequest {
	r.reason = &reason
	return r
}

func (r ApiReportPostRequest) Screenshot2Filename(screenshot2Filename string) ApiReportPostRequest {
	r.screenshot2Filename = &screenshot2Filename
	return r
}

func (r ApiReportPostRequest) Screenshot3Filename(screenshot3Filename string) ApiReportPostRequest {
	r.screenshot3Filename = &screenshot3Filename
	return r
}

func (r ApiReportPostRequest) Screenshot4Filename(screenshot4Filename string) ApiReportPostRequest {
	r.screenshot4Filename = &screenshot4Filename
	return r
}

func (r ApiReportPostRequest) ScreenshotFilename(screenshotFilename string) ApiReportPostRequest {
	r.screenshotFilename = &screenshotFilename
	return r
}

func (r ApiReportPostRequest) Execute() (*http.Response, error) {
	return r.ApiService.ReportPostExecute(r)
}

/*
ReportPost Method for ReportPost

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param postId
 @return ApiReportPostRequest
*/
func (a *PostsAPIService) ReportPost(ctx context.Context, postId int64) ApiReportPostRequest {
	return ApiReportPostRequest{
		ApiService: a,
		ctx: ctx,
		postId: postId,
	}
}

// Execute executes the request
func (a *PostsAPIService) ReportPostExecute(r ApiReportPostRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PostsAPIService.ReportPost")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v3/posts/{post_id}/report"
	localVarPath = strings.Replace(localVarPath, "{"+"post_id"+"}", url.PathEscape(parameterValueToString(r.postId, "postId")), -1)

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

type ApiRepostRequest struct {
	ctx context.Context
	ApiService *PostsAPIService
	xJwt *string
	postId *int64
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

func (r ApiRepostRequest) XJwt(xJwt string) ApiRepostRequest {
	r.xJwt = &xJwt
	return r
}

func (r ApiRepostRequest) PostId(postId int64) ApiRepostRequest {
	r.postId = &postId
	return r
}

func (r ApiRepostRequest) Attachment2Filename(attachment2Filename string) ApiRepostRequest {
	r.attachment2Filename = &attachment2Filename
	return r
}

func (r ApiRepostRequest) Attachment3Filename(attachment3Filename string) ApiRepostRequest {
	r.attachment3Filename = &attachment3Filename
	return r
}

func (r ApiRepostRequest) Attachment4Filename(attachment4Filename string) ApiRepostRequest {
	r.attachment4Filename = &attachment4Filename
	return r
}

func (r ApiRepostRequest) Attachment5Filename(attachment5Filename string) ApiRepostRequest {
	r.attachment5Filename = &attachment5Filename
	return r
}

func (r ApiRepostRequest) Attachment6Filename(attachment6Filename string) ApiRepostRequest {
	r.attachment6Filename = &attachment6Filename
	return r
}

func (r ApiRepostRequest) Attachment7Filename(attachment7Filename string) ApiRepostRequest {
	r.attachment7Filename = &attachment7Filename
	return r
}

func (r ApiRepostRequest) Attachment8Filename(attachment8Filename string) ApiRepostRequest {
	r.attachment8Filename = &attachment8Filename
	return r
}

func (r ApiRepostRequest) Attachment9Filename(attachment9Filename string) ApiRepostRequest {
	r.attachment9Filename = &attachment9Filename
	return r
}

func (r ApiRepostRequest) AttachmentFilename(attachmentFilename string) ApiRepostRequest {
	r.attachmentFilename = &attachmentFilename
	return r
}

func (r ApiRepostRequest) Choices(choices []string) ApiRepostRequest {
	r.choices = &choices
	return r
}

func (r ApiRepostRequest) Color(color int32) ApiRepostRequest {
	r.color = &color
	return r
}

func (r ApiRepostRequest) FontSize(fontSize int32) ApiRepostRequest {
	r.fontSize = &fontSize
	return r
}

func (r ApiRepostRequest) GroupId(groupId int64) ApiRepostRequest {
	r.groupId = &groupId
	return r
}

func (r ApiRepostRequest) InReplyTo(inReplyTo int64) ApiRepostRequest {
	r.inReplyTo = &inReplyTo
	return r
}

func (r ApiRepostRequest) Language(language string) ApiRepostRequest {
	r.language = &language
	return r
}

func (r ApiRepostRequest) MentionIds(mentionIds []int64) ApiRepostRequest {
	r.mentionIds = &mentionIds
	return r
}

// Unresolved Java type: pq.g0
func (r ApiRepostRequest) MessageTags(messageTags map[string]interface{}) ApiRepostRequest {
	r.messageTags = &messageTags
	return r
}

func (r ApiRepostRequest) PostType(postType PostType) ApiRepostRequest {
	r.postType = &postType
	return r
}

// Unresolved Java type: pq.g0
func (r ApiRepostRequest) SharedUrl(sharedUrl map[string]interface{}) ApiRepostRequest {
	r.sharedUrl = &sharedUrl
	return r
}

func (r ApiRepostRequest) Text(text string) ApiRepostRequest {
	r.text = &text
	return r
}

func (r ApiRepostRequest) VideoFileName(videoFileName string) ApiRepostRequest {
	r.videoFileName = &videoFileName
	return r
}

func (r ApiRepostRequest) Execute() (*CreatePostResponse, *http.Response, error) {
	return r.ApiService.RepostExecute(r)
}

/*
Repost Method for Repost

Implemented as a Kotlin suspend function on Android.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiRepostRequest
*/
func (a *PostsAPIService) Repost(ctx context.Context) ApiRepostRequest {
	return ApiRepostRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return CreatePostResponse
func (a *PostsAPIService) RepostExecute(r ApiRepostRequest) (*CreatePostResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *CreatePostResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PostsAPIService.Repost")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v3/posts/repost"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.xJwt == nil {
		return localVarReturnValue, nil, reportError("xJwt is required and must be specified")
	}
	if r.postId == nil {
		return localVarReturnValue, nil, reportError("postId is required and must be specified")
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
	parameterAddToHeaderOrQuery(localVarHeaderParams, "X-Jwt", r.xJwt, "simple", "")
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
	parameterAddToHeaderOrQuery(localVarFormParams, "post_id", r.postId, "", "")
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

type ApiSearchPostsRequest struct {
	ctx context.Context
	ApiService *PostsAPIService
	keyword *string
	postOwnerScope *string
	onlyMedia *bool
	fromPostId *int64
	number *int32
}

func (r ApiSearchPostsRequest) Keyword(keyword string) ApiSearchPostsRequest {
	r.keyword = &keyword
	return r
}

func (r ApiSearchPostsRequest) PostOwnerScope(postOwnerScope string) ApiSearchPostsRequest {
	r.postOwnerScope = &postOwnerScope
	return r
}

func (r ApiSearchPostsRequest) OnlyMedia(onlyMedia bool) ApiSearchPostsRequest {
	r.onlyMedia = &onlyMedia
	return r
}

func (r ApiSearchPostsRequest) FromPostId(fromPostId int64) ApiSearchPostsRequest {
	r.fromPostId = &fromPostId
	return r
}

func (r ApiSearchPostsRequest) Number(number int32) ApiSearchPostsRequest {
	r.number = &number
	return r
}

func (r ApiSearchPostsRequest) Execute() (*PostsResponse, *http.Response, error) {
	return r.ApiService.SearchPostsExecute(r)
}

/*
SearchPosts Method for SearchPosts

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiSearchPostsRequest
*/
func (a *PostsAPIService) SearchPosts(ctx context.Context) ApiSearchPostsRequest {
	return ApiSearchPostsRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return PostsResponse
func (a *PostsAPIService) SearchPostsExecute(r ApiSearchPostsRequest) (*PostsResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *PostsResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PostsAPIService.SearchPosts")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v2/posts/search"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.keyword == nil {
		return localVarReturnValue, nil, reportError("keyword is required and must be specified")
	}
	if r.postOwnerScope == nil {
		return localVarReturnValue, nil, reportError("postOwnerScope is required and must be specified")
	}
	if r.onlyMedia == nil {
		return localVarReturnValue, nil, reportError("onlyMedia is required and must be specified")
	}

	parameterAddToHeaderOrQuery(localVarQueryParams, "keyword", r.keyword, "form", "")
	parameterAddToHeaderOrQuery(localVarQueryParams, "post_owner_scope", r.postOwnerScope, "form", "")
	parameterAddToHeaderOrQuery(localVarQueryParams, "only_media", r.onlyMedia, "form", "")
	if r.fromPostId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "from_post_id", r.fromPostId, "form", "")
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

type ApiUnlikePostRequest struct {
	ctx context.Context
	ApiService *PostsAPIService
	id int64
}

func (r ApiUnlikePostRequest) Execute() (*http.Response, error) {
	return r.ApiService.UnlikePostExecute(r)
}

/*
UnlikePost Method for UnlikePost

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id
 @return ApiUnlikePostRequest
*/
func (a *PostsAPIService) UnlikePost(ctx context.Context, id int64) ApiUnlikePostRequest {
	return ApiUnlikePostRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
func (a *PostsAPIService) UnlikePostExecute(r ApiUnlikePostRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PostsAPIService.UnlikePost")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/posts/{id}/unlike"
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

type ApiUnpinGroupPostRequest struct {
	ctx context.Context
	ApiService *PostsAPIService
	groupId *int64
}

func (r ApiUnpinGroupPostRequest) GroupId(groupId int64) ApiUnpinGroupPostRequest {
	r.groupId = &groupId
	return r
}

func (r ApiUnpinGroupPostRequest) Execute() (*http.Response, error) {
	return r.ApiService.UnpinGroupPostExecute(r)
}

/*
UnpinGroupPost Method for UnpinGroupPost

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiUnpinGroupPostRequest
*/
func (a *PostsAPIService) UnpinGroupPost(ctx context.Context) ApiUnpinGroupPostRequest {
	return ApiUnpinGroupPostRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
func (a *PostsAPIService) UnpinGroupPostExecute(r ApiUnpinGroupPostRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PostsAPIService.UnpinGroupPost")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v2/posts/group_pinned_post"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.groupId == nil {
		return nil, reportError("groupId is required and must be specified")
	}

	parameterAddToHeaderOrQuery(localVarQueryParams, "group_id", r.groupId, "form", "")
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

type ApiUpdatePostRequest struct {
	ctx context.Context
	ApiService *PostsAPIService
	id int64
	apiKey *string
	signedInfo *string
	timestamp *int64
	uuid *string
	color *int32
	fontSize *int32
	language *string
	messageTags *map[string]interface{}
	text *string
}

func (r ApiUpdatePostRequest) ApiKey(apiKey string) ApiUpdatePostRequest {
	r.apiKey = &apiKey
	return r
}

func (r ApiUpdatePostRequest) SignedInfo(signedInfo string) ApiUpdatePostRequest {
	r.signedInfo = &signedInfo
	return r
}

func (r ApiUpdatePostRequest) Timestamp(timestamp int64) ApiUpdatePostRequest {
	r.timestamp = &timestamp
	return r
}

func (r ApiUpdatePostRequest) Uuid(uuid string) ApiUpdatePostRequest {
	r.uuid = &uuid
	return r
}

func (r ApiUpdatePostRequest) Color(color int32) ApiUpdatePostRequest {
	r.color = &color
	return r
}

func (r ApiUpdatePostRequest) FontSize(fontSize int32) ApiUpdatePostRequest {
	r.fontSize = &fontSize
	return r
}

func (r ApiUpdatePostRequest) Language(language string) ApiUpdatePostRequest {
	r.language = &language
	return r
}

// Unresolved Java type: pq.g0
func (r ApiUpdatePostRequest) MessageTags(messageTags map[string]interface{}) ApiUpdatePostRequest {
	r.messageTags = &messageTags
	return r
}

func (r ApiUpdatePostRequest) Text(text string) ApiUpdatePostRequest {
	r.text = &text
	return r
}

func (r ApiUpdatePostRequest) Execute() (*Post, *http.Response, error) {
	return r.ApiService.UpdatePostExecute(r)
}

/*
UpdatePost Method for UpdatePost

Implemented as a Kotlin suspend function on Android.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id
 @return ApiUpdatePostRequest
*/
func (a *PostsAPIService) UpdatePost(ctx context.Context, id int64) ApiUpdatePostRequest {
	return ApiUpdatePostRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
//  @return Post
func (a *PostsAPIService) UpdatePostExecute(r ApiUpdatePostRequest) (*Post, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPut
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *Post
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PostsAPIService.UpdatePost")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v3/posts/{id}"
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
	parameterAddToHeaderOrQuery(localVarFormParams, "api_key", r.apiKey, "", "")
	if r.color != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "color", r.color, "", "")
	}
	if r.fontSize != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "font_size", r.fontSize, "", "")
	}
	if r.language != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "language", r.language, "", "")
	}
	if r.messageTags != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "message_tags", r.messageTags, "", "")
	}
	parameterAddToHeaderOrQuery(localVarFormParams, "signed_info", r.signedInfo, "", "")
	if r.text != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "text", r.text, "", "")
	}
	parameterAddToHeaderOrQuery(localVarFormParams, "timestamp", r.timestamp, "", "")
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

type ApiValidatePostRequest struct {
	ctx context.Context
	ApiService *PostsAPIService
	text *string
	groupId *int64
	threadId *int64
}

func (r ApiValidatePostRequest) Text(text string) ApiValidatePostRequest {
	r.text = &text
	return r
}

func (r ApiValidatePostRequest) GroupId(groupId int64) ApiValidatePostRequest {
	r.groupId = &groupId
	return r
}

func (r ApiValidatePostRequest) ThreadId(threadId int64) ApiValidatePostRequest {
	r.threadId = &threadId
	return r
}

func (r ApiValidatePostRequest) Execute() (*ValidationPostResponse, *http.Response, error) {
	return r.ApiService.ValidatePostExecute(r)
}

/*
ValidatePost Method for ValidatePost

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiValidatePostRequest
*/
func (a *PostsAPIService) ValidatePost(ctx context.Context) ApiValidatePostRequest {
	return ApiValidatePostRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return ValidationPostResponse
func (a *PostsAPIService) ValidatePostExecute(r ApiValidatePostRequest) (*ValidationPostResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ValidationPostResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PostsAPIService.ValidatePost")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/posts/validate"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.text == nil {
		return localVarReturnValue, nil, reportError("text is required and must be specified")
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
	if r.groupId != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "group_id", r.groupId, "", "")
	}
	parameterAddToHeaderOrQuery(localVarFormParams, "text", r.text, "", "")
	if r.threadId != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "thread_id", r.threadId, "", "")
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

type ApiViewPostVideoRequest struct {
	ctx context.Context
	ApiService *PostsAPIService
	videoId int64
}

func (r ApiViewPostVideoRequest) Execute() (*http.Response, error) {
	return r.ApiService.ViewPostVideoExecute(r)
}

/*
ViewPostVideo Method for ViewPostVideo

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param videoId
 @return ApiViewPostVideoRequest
*/
func (a *PostsAPIService) ViewPostVideo(ctx context.Context, videoId int64) ApiViewPostVideoRequest {
	return ApiViewPostVideoRequest{
		ApiService: a,
		ctx: ctx,
		videoId: videoId,
	}
}

// Execute executes the request
func (a *PostsAPIService) ViewPostVideoExecute(r ApiViewPostVideoRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PostsAPIService.ViewPostVideo")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/posts/videos/{videoId}/view"
	localVarPath = strings.Replace(localVarPath, "{"+"videoId"+"}", url.PathEscape(parameterValueToString(r.videoId, "videoId")), -1)

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
