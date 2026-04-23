
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


// ConversationsAPIService ConversationsAPI service
type ConversationsAPIService service

type ApiGetConversationRequest struct {
	ctx context.Context
	ApiService *ConversationsAPIService
	id int64
	reverse *bool
	groupId *int64
	threadId *int64
	number *int32
	fromPostId *int64
}

func (r ApiGetConversationRequest) Reverse(reverse bool) ApiGetConversationRequest {
	r.reverse = &reverse
	return r
}

func (r ApiGetConversationRequest) GroupId(groupId int64) ApiGetConversationRequest {
	r.groupId = &groupId
	return r
}

func (r ApiGetConversationRequest) ThreadId(threadId int64) ApiGetConversationRequest {
	r.threadId = &threadId
	return r
}

func (r ApiGetConversationRequest) Number(number int32) ApiGetConversationRequest {
	r.number = &number
	return r
}

func (r ApiGetConversationRequest) FromPostId(fromPostId int64) ApiGetConversationRequest {
	r.fromPostId = &fromPostId
	return r
}

func (r ApiGetConversationRequest) Execute() (*PostsResponse, *http.Response, error) {
	return r.ApiService.GetConversationExecute(r)
}

/*
GetConversation Method for GetConversation

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id
 @return ApiGetConversationRequest
*/
func (a *ConversationsAPIService) GetConversation(ctx context.Context, id int64) ApiGetConversationRequest {
	return ApiGetConversationRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
//  @return PostsResponse
func (a *ConversationsAPIService) GetConversationExecute(r ApiGetConversationRequest) (*PostsResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *PostsResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ConversationsAPIService.GetConversation")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v2/conversations/{id}"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterValueToString(r.id, "id")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.reverse == nil {
		return localVarReturnValue, nil, reportError("reverse is required and must be specified")
	}

	if r.groupId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "group_id", r.groupId, "form", "")
	}
	if r.threadId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "thread_id", r.threadId, "form", "")
	}
	if r.number != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "number", r.number, "form", "")
	}
	if r.fromPostId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "from_post_id", r.fromPostId, "form", "")
	}
	parameterAddToHeaderOrQuery(localVarQueryParams, "reverse", r.reverse, "form", "")
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

type ApiGetRootPostsRequest struct {
	ctx context.Context
	ApiService *ConversationsAPIService
	ids *[]int64
}

func (r ApiGetRootPostsRequest) Ids(ids []int64) ApiGetRootPostsRequest {
	r.ids = &ids
	return r
}

func (r ApiGetRootPostsRequest) Execute() (*PostsResponse, *http.Response, error) {
	return r.ApiService.GetRootPostsExecute(r)
}

/*
GetRootPosts Method for GetRootPosts

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiGetRootPostsRequest
*/
func (a *ConversationsAPIService) GetRootPosts(ctx context.Context) ApiGetRootPostsRequest {
	return ApiGetRootPostsRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return PostsResponse
func (a *ConversationsAPIService) GetRootPostsExecute(r ApiGetRootPostsRequest) (*PostsResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *PostsResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ConversationsAPIService.GetRootPosts")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v2/conversations/root_posts"

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
