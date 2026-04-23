
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


// UsersAPIService UsersAPI service
type UsersAPIService service

type ApiAgreePolicyRequest struct {
	ctx context.Context
	ApiService *UsersAPIService
	type_ string
}

func (r ApiAgreePolicyRequest) Execute() (*http.Response, error) {
	return r.ApiService.AgreePolicyExecute(r)
}

/*
AgreePolicy Method for AgreePolicy

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param type_
 @return ApiAgreePolicyRequest
*/
func (a *UsersAPIService) AgreePolicy(ctx context.Context, type_ string) ApiAgreePolicyRequest {
	return ApiAgreePolicyRequest{
		ApiService: a,
		ctx: ctx,
		type_: type_,
	}
}

// Execute executes the request
func (a *UsersAPIService) AgreePolicyExecute(r ApiAgreePolicyRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "UsersAPIService.AgreePolicy")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/users/policy_agreements/{type}"
	localVarPath = strings.Replace(localVarPath, "{"+"type"+"}", url.PathEscape(parameterValueToString(r.type_, "type_")), -1)

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

type ApiBlockUserRequest struct {
	ctx context.Context
	ApiService *UsersAPIService
	id int64
}

func (r ApiBlockUserRequest) Execute() (*http.Response, error) {
	return r.ApiService.BlockUserExecute(r)
}

/*
BlockUser Method for BlockUser

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id
 @return ApiBlockUserRequest
*/
func (a *UsersAPIService) BlockUser(ctx context.Context, id int64) ApiBlockUserRequest {
	return ApiBlockUserRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
func (a *UsersAPIService) BlockUserExecute(r ApiBlockUserRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "UsersAPIService.BlockUser")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/users/{id}/block"
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

type ApiChangeEmailRequest struct {
	ctx context.Context
	ApiService *UsersAPIService
	androidxAutofillHintConstantsAUTOFILLHINTPASSWORD *string
	apiKey *string
	email *string
	emailGrantToken *string
}

func (r ApiChangeEmailRequest) AndroidxAutofillHintConstantsAUTOFILLHINTPASSWORD(androidxAutofillHintConstantsAUTOFILLHINTPASSWORD string) ApiChangeEmailRequest {
	r.androidxAutofillHintConstantsAUTOFILLHINTPASSWORD = &androidxAutofillHintConstantsAUTOFILLHINTPASSWORD
	return r
}

func (r ApiChangeEmailRequest) ApiKey(apiKey string) ApiChangeEmailRequest {
	r.apiKey = &apiKey
	return r
}

func (r ApiChangeEmailRequest) Email(email string) ApiChangeEmailRequest {
	r.email = &email
	return r
}

func (r ApiChangeEmailRequest) EmailGrantToken(emailGrantToken string) ApiChangeEmailRequest {
	r.emailGrantToken = &emailGrantToken
	return r
}

func (r ApiChangeEmailRequest) Execute() (*LoginUpdateResponse, *http.Response, error) {
	return r.ApiService.ChangeEmailExecute(r)
}

/*
ChangeEmail Method for ChangeEmail

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiChangeEmailRequest
*/
func (a *UsersAPIService) ChangeEmail(ctx context.Context) ApiChangeEmailRequest {
	return ApiChangeEmailRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return LoginUpdateResponse
func (a *UsersAPIService) ChangeEmailExecute(r ApiChangeEmailRequest) (*LoginUpdateResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPut
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *LoginUpdateResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "UsersAPIService.ChangeEmail")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/users/change_email"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.androidxAutofillHintConstantsAUTOFILLHINTPASSWORD == nil {
		return localVarReturnValue, nil, reportError("androidxAutofillHintConstantsAUTOFILLHINTPASSWORD is required and must be specified")
	}
	if r.apiKey == nil {
		return localVarReturnValue, nil, reportError("apiKey is required and must be specified")
	}
	if r.email == nil {
		return localVarReturnValue, nil, reportError("email is required and must be specified")
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
	parameterAddToHeaderOrQuery(localVarFormParams, "androidx.autofill.HintConstants.AUTOFILL_HINT_PASSWORD", r.androidxAutofillHintConstantsAUTOFILLHINTPASSWORD, "", "")
	parameterAddToHeaderOrQuery(localVarFormParams, "api_key", r.apiKey, "", "")
	parameterAddToHeaderOrQuery(localVarFormParams, "email", r.email, "", "")
	if r.emailGrantToken != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "email_grant_token", r.emailGrantToken, "", "")
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

type ApiChangePasswordRequest struct {
	ctx context.Context
	ApiService *UsersAPIService
	androidxAutofillHintConstantsAUTOFILLHINTPASSWORD *string
	apiKey *string
	currentPassword *string
}

func (r ApiChangePasswordRequest) AndroidxAutofillHintConstantsAUTOFILLHINTPASSWORD(androidxAutofillHintConstantsAUTOFILLHINTPASSWORD string) ApiChangePasswordRequest {
	r.androidxAutofillHintConstantsAUTOFILLHINTPASSWORD = &androidxAutofillHintConstantsAUTOFILLHINTPASSWORD
	return r
}

func (r ApiChangePasswordRequest) ApiKey(apiKey string) ApiChangePasswordRequest {
	r.apiKey = &apiKey
	return r
}

func (r ApiChangePasswordRequest) CurrentPassword(currentPassword string) ApiChangePasswordRequest {
	r.currentPassword = &currentPassword
	return r
}

func (r ApiChangePasswordRequest) Execute() (*LoginUpdateResponse, *http.Response, error) {
	return r.ApiService.ChangePasswordExecute(r)
}

/*
ChangePassword Method for ChangePassword

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiChangePasswordRequest
*/
func (a *UsersAPIService) ChangePassword(ctx context.Context) ApiChangePasswordRequest {
	return ApiChangePasswordRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return LoginUpdateResponse
func (a *UsersAPIService) ChangePasswordExecute(r ApiChangePasswordRequest) (*LoginUpdateResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPut
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *LoginUpdateResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "UsersAPIService.ChangePassword")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/users/change_password"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.androidxAutofillHintConstantsAUTOFILLHINTPASSWORD == nil {
		return localVarReturnValue, nil, reportError("androidxAutofillHintConstantsAUTOFILLHINTPASSWORD is required and must be specified")
	}
	if r.apiKey == nil {
		return localVarReturnValue, nil, reportError("apiKey is required and must be specified")
	}
	if r.currentPassword == nil {
		return localVarReturnValue, nil, reportError("currentPassword is required and must be specified")
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
	parameterAddToHeaderOrQuery(localVarFormParams, "androidx.autofill.HintConstants.AUTOFILL_HINT_PASSWORD", r.androidxAutofillHintConstantsAUTOFILLHINTPASSWORD, "", "")
	parameterAddToHeaderOrQuery(localVarFormParams, "api_key", r.apiKey, "", "")
	parameterAddToHeaderOrQuery(localVarFormParams, "current_password", r.currentPassword, "", "")
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

type ApiCreateBookmarkRequest struct {
	ctx context.Context
	ApiService *UsersAPIService
	userId int64
	id int64
}

func (r ApiCreateBookmarkRequest) Execute() (*BookmarkPostResponse, *http.Response, error) {
	return r.ApiService.CreateBookmarkExecute(r)
}

/*
CreateBookmark Method for CreateBookmark

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param userId
 @param id
 @return ApiCreateBookmarkRequest
*/
func (a *UsersAPIService) CreateBookmark(ctx context.Context, userId int64, id int64) ApiCreateBookmarkRequest {
	return ApiCreateBookmarkRequest{
		ApiService: a,
		ctx: ctx,
		userId: userId,
		id: id,
	}
}

// Execute executes the request
//  @return BookmarkPostResponse
func (a *UsersAPIService) CreateBookmarkExecute(r ApiCreateBookmarkRequest) (*BookmarkPostResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPut
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *BookmarkPostResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "UsersAPIService.CreateBookmark")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/users/{user_id}/bookmarks/{id}"
	localVarPath = strings.Replace(localVarPath, "{"+"user_id"+"}", url.PathEscape(parameterValueToString(r.userId, "userId")), -1)
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

type ApiCreateReviewRequest struct {
	ctx context.Context
	ApiService *UsersAPIService
	apiKey *string
	comMbridgeMsdkFoundationEntityCampaignExJSONKEYTIMESTAMP *int64
	comment *string
	signedInfo *string
	userIds *[]int64
	uuid *string
}

func (r ApiCreateReviewRequest) ApiKey(apiKey string) ApiCreateReviewRequest {
	r.apiKey = &apiKey
	return r
}

func (r ApiCreateReviewRequest) ComMbridgeMsdkFoundationEntityCampaignExJSONKEYTIMESTAMP(comMbridgeMsdkFoundationEntityCampaignExJSONKEYTIMESTAMP int64) ApiCreateReviewRequest {
	r.comMbridgeMsdkFoundationEntityCampaignExJSONKEYTIMESTAMP = &comMbridgeMsdkFoundationEntityCampaignExJSONKEYTIMESTAMP
	return r
}

func (r ApiCreateReviewRequest) Comment(comment string) ApiCreateReviewRequest {
	r.comment = &comment
	return r
}

func (r ApiCreateReviewRequest) SignedInfo(signedInfo string) ApiCreateReviewRequest {
	r.signedInfo = &signedInfo
	return r
}

func (r ApiCreateReviewRequest) UserIds(userIds []int64) ApiCreateReviewRequest {
	r.userIds = &userIds
	return r
}

func (r ApiCreateReviewRequest) Uuid(uuid string) ApiCreateReviewRequest {
	r.uuid = &uuid
	return r
}

func (r ApiCreateReviewRequest) Execute() (*http.Response, error) {
	return r.ApiService.CreateReviewExecute(r)
}

/*
CreateReview Method for CreateReview

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiCreateReviewRequest
*/
func (a *UsersAPIService) CreateReview(ctx context.Context) ApiCreateReviewRequest {
	return ApiCreateReviewRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
func (a *UsersAPIService) CreateReviewExecute(r ApiCreateReviewRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "UsersAPIService.CreateReview")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/users/reviews"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.apiKey == nil {
		return nil, reportError("apiKey is required and must be specified")
	}
	if r.comMbridgeMsdkFoundationEntityCampaignExJSONKEYTIMESTAMP == nil {
		return nil, reportError("comMbridgeMsdkFoundationEntityCampaignExJSONKEYTIMESTAMP is required and must be specified")
	}
	if r.comment == nil {
		return nil, reportError("comment is required and must be specified")
	}
	if r.signedInfo == nil {
		return nil, reportError("signedInfo is required and must be specified")
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
	parameterAddToHeaderOrQuery(localVarFormParams, "com.mbridge.msdk.foundation.entity.CampaignEx.JSON_KEY_TIMESTAMP", r.comMbridgeMsdkFoundationEntityCampaignExJSONKEYTIMESTAMP, "", "")
	parameterAddToHeaderOrQuery(localVarFormParams, "comment", r.comment, "", "")
	parameterAddToHeaderOrQuery(localVarFormParams, "signed_info", r.signedInfo, "", "")
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

type ApiDeleteBookmarkRequest struct {
	ctx context.Context
	ApiService *UsersAPIService
	userId int64
	id int64
}

func (r ApiDeleteBookmarkRequest) Execute() (*http.Response, error) {
	return r.ApiService.DeleteBookmarkExecute(r)
}

/*
DeleteBookmark Method for DeleteBookmark

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param userId
 @param id
 @return ApiDeleteBookmarkRequest
*/
func (a *UsersAPIService) DeleteBookmark(ctx context.Context, userId int64, id int64) ApiDeleteBookmarkRequest {
	return ApiDeleteBookmarkRequest{
		ApiService: a,
		ctx: ctx,
		userId: userId,
		id: id,
	}
}

// Execute executes the request
func (a *UsersAPIService) DeleteBookmarkExecute(r ApiDeleteBookmarkRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "UsersAPIService.DeleteBookmark")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/users/{user_id}/bookmarks/{id}"
	localVarPath = strings.Replace(localVarPath, "{"+"user_id"+"}", url.PathEscape(parameterValueToString(r.userId, "userId")), -1)
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

type ApiDeleteFootprintRequest struct {
	ctx context.Context
	ApiService *UsersAPIService
	userId int64
	footprintId int64
}

func (r ApiDeleteFootprintRequest) Execute() (*http.Response, error) {
	return r.ApiService.DeleteFootprintExecute(r)
}

/*
DeleteFootprint Method for DeleteFootprint

Implemented as a Kotlin suspend function on Android.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param userId
 @param footprintId
 @return ApiDeleteFootprintRequest
*/
func (a *UsersAPIService) DeleteFootprint(ctx context.Context, userId int64, footprintId int64) ApiDeleteFootprintRequest {
	return ApiDeleteFootprintRequest{
		ApiService: a,
		ctx: ctx,
		userId: userId,
		footprintId: footprintId,
	}
}

// Execute executes the request
func (a *UsersAPIService) DeleteFootprintExecute(r ApiDeleteFootprintRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "UsersAPIService.DeleteFootprint")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v2/users/{user_id}/footprints/{footprint_id}"
	localVarPath = strings.Replace(localVarPath, "{"+"user_id"+"}", url.PathEscape(parameterValueToString(r.userId, "userId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"footprint_id"+"}", url.PathEscape(parameterValueToString(r.footprintId, "footprintId")), -1)

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

type ApiDeleteMyReviewsRequest struct {
	ctx context.Context
	ApiService *UsersAPIService
	reviewIds *[]int64
}

func (r ApiDeleteMyReviewsRequest) ReviewIds(reviewIds []int64) ApiDeleteMyReviewsRequest {
	r.reviewIds = &reviewIds
	return r
}

func (r ApiDeleteMyReviewsRequest) Execute() (*http.Response, error) {
	return r.ApiService.DeleteMyReviewsExecute(r)
}

/*
DeleteMyReviews Method for DeleteMyReviews

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiDeleteMyReviewsRequest
*/
func (a *UsersAPIService) DeleteMyReviews(ctx context.Context) ApiDeleteMyReviewsRequest {
	return ApiDeleteMyReviewsRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
func (a *UsersAPIService) DeleteMyReviewsExecute(r ApiDeleteMyReviewsRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "UsersAPIService.DeleteMyReviews")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/users/reviews"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.reviewIds == nil {
		return nil, reportError("reviewIds is required and must be specified")
	}

	{
		t := *r.reviewIds
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				parameterAddToHeaderOrQuery(localVarQueryParams, "review_ids[]", s.Index(i).Interface(), "form", "multi")
			}
		} else {
			parameterAddToHeaderOrQuery(localVarQueryParams, "review_ids[]", t, "form", "multi")
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

type ApiDisableTwoFactorAuthRequest struct {
	ctx context.Context
	ApiService *UsersAPIService
	code *string
}

func (r ApiDisableTwoFactorAuthRequest) Code(code string) ApiDisableTwoFactorAuthRequest {
	r.code = &code
	return r
}

func (r ApiDisableTwoFactorAuthRequest) Execute() (*http.Response, error) {
	return r.ApiService.DisableTwoFactorAuthExecute(r)
}

/*
DisableTwoFactorAuth Method for DisableTwoFactorAuth

Implemented as a Kotlin suspend function on Android.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiDisableTwoFactorAuthRequest
*/
func (a *UsersAPIService) DisableTwoFactorAuth(ctx context.Context) ApiDisableTwoFactorAuthRequest {
	return ApiDisableTwoFactorAuthRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
func (a *UsersAPIService) DisableTwoFactorAuthExecute(r ApiDisableTwoFactorAuthRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPut
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "UsersAPIService.DisableTwoFactorAuth")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/users/secret/disable"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.code == nil {
		return nil, reportError("code is required and must be specified")
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
	parameterAddToHeaderOrQuery(localVarFormParams, "code", r.code, "", "")
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

type ApiEditUserRequest struct {
	ctx context.Context
	ApiService *UsersAPIService
	apiKey *string
	comMbridgeMsdkFoundationEntityCampaignExJSONKEYTIMESTAMP *int64
	nickname *string
	signedInfo *string
	uuid *string
	biography *string
	countryCode *string
	coverImageFilename *string
	gender *int32
	prefecture *string
	profileIconFilename *string
	username *string
}

func (r ApiEditUserRequest) ApiKey(apiKey string) ApiEditUserRequest {
	r.apiKey = &apiKey
	return r
}

func (r ApiEditUserRequest) ComMbridgeMsdkFoundationEntityCampaignExJSONKEYTIMESTAMP(comMbridgeMsdkFoundationEntityCampaignExJSONKEYTIMESTAMP int64) ApiEditUserRequest {
	r.comMbridgeMsdkFoundationEntityCampaignExJSONKEYTIMESTAMP = &comMbridgeMsdkFoundationEntityCampaignExJSONKEYTIMESTAMP
	return r
}

func (r ApiEditUserRequest) Nickname(nickname string) ApiEditUserRequest {
	r.nickname = &nickname
	return r
}

func (r ApiEditUserRequest) SignedInfo(signedInfo string) ApiEditUserRequest {
	r.signedInfo = &signedInfo
	return r
}

func (r ApiEditUserRequest) Uuid(uuid string) ApiEditUserRequest {
	r.uuid = &uuid
	return r
}

func (r ApiEditUserRequest) Biography(biography string) ApiEditUserRequest {
	r.biography = &biography
	return r
}

func (r ApiEditUserRequest) CountryCode(countryCode string) ApiEditUserRequest {
	r.countryCode = &countryCode
	return r
}

func (r ApiEditUserRequest) CoverImageFilename(coverImageFilename string) ApiEditUserRequest {
	r.coverImageFilename = &coverImageFilename
	return r
}

func (r ApiEditUserRequest) Gender(gender int32) ApiEditUserRequest {
	r.gender = &gender
	return r
}

func (r ApiEditUserRequest) Prefecture(prefecture string) ApiEditUserRequest {
	r.prefecture = &prefecture
	return r
}

func (r ApiEditUserRequest) ProfileIconFilename(profileIconFilename string) ApiEditUserRequest {
	r.profileIconFilename = &profileIconFilename
	return r
}

func (r ApiEditUserRequest) Username(username string) ApiEditUserRequest {
	r.username = &username
	return r
}

func (r ApiEditUserRequest) Execute() (*http.Response, error) {
	return r.ApiService.EditUserExecute(r)
}

/*
EditUser Method for EditUser

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiEditUserRequest
*/
func (a *UsersAPIService) EditUser(ctx context.Context) ApiEditUserRequest {
	return ApiEditUserRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
func (a *UsersAPIService) EditUserExecute(r ApiEditUserRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "UsersAPIService.EditUser")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v3/users/edit"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.apiKey == nil {
		return nil, reportError("apiKey is required and must be specified")
	}
	if r.comMbridgeMsdkFoundationEntityCampaignExJSONKEYTIMESTAMP == nil {
		return nil, reportError("comMbridgeMsdkFoundationEntityCampaignExJSONKEYTIMESTAMP is required and must be specified")
	}
	if r.nickname == nil {
		return nil, reportError("nickname is required and must be specified")
	}
	if r.signedInfo == nil {
		return nil, reportError("signedInfo is required and must be specified")
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
	if r.biography != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "biography", r.biography, "", "")
	}
	parameterAddToHeaderOrQuery(localVarFormParams, "com.mbridge.msdk.foundation.entity.CampaignEx.JSON_KEY_TIMESTAMP", r.comMbridgeMsdkFoundationEntityCampaignExJSONKEYTIMESTAMP, "", "")
	if r.countryCode != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "country_code", r.countryCode, "", "")
	}
	if r.coverImageFilename != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "cover_image_filename", r.coverImageFilename, "", "")
	}
	if r.gender != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "gender", r.gender, "", "")
	}
	parameterAddToHeaderOrQuery(localVarFormParams, "nickname", r.nickname, "", "")
	if r.prefecture != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "prefecture", r.prefecture, "", "")
	}
	if r.profileIconFilename != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "profile_icon_filename", r.profileIconFilename, "", "")
	}
	parameterAddToHeaderOrQuery(localVarFormParams, "signed_info", r.signedInfo, "", "")
	if r.username != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "username", r.username, "", "")
	}
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

type ApiEditUserV2Request struct {
	ctx context.Context
	ApiService *UsersAPIService
	apiKey *string
	comMbridgeMsdkFoundationEntityCampaignExJSONKEYTIMESTAMP *int64
	isPrivate *bool
	nickname *string
	signedInfo *string
	uuid *string
}

func (r ApiEditUserV2Request) ApiKey(apiKey string) ApiEditUserV2Request {
	r.apiKey = &apiKey
	return r
}

func (r ApiEditUserV2Request) ComMbridgeMsdkFoundationEntityCampaignExJSONKEYTIMESTAMP(comMbridgeMsdkFoundationEntityCampaignExJSONKEYTIMESTAMP int64) ApiEditUserV2Request {
	r.comMbridgeMsdkFoundationEntityCampaignExJSONKEYTIMESTAMP = &comMbridgeMsdkFoundationEntityCampaignExJSONKEYTIMESTAMP
	return r
}

func (r ApiEditUserV2Request) IsPrivate(isPrivate bool) ApiEditUserV2Request {
	r.isPrivate = &isPrivate
	return r
}

func (r ApiEditUserV2Request) Nickname(nickname string) ApiEditUserV2Request {
	r.nickname = &nickname
	return r
}

func (r ApiEditUserV2Request) SignedInfo(signedInfo string) ApiEditUserV2Request {
	r.signedInfo = &signedInfo
	return r
}

func (r ApiEditUserV2Request) Uuid(uuid string) ApiEditUserV2Request {
	r.uuid = &uuid
	return r
}

func (r ApiEditUserV2Request) Execute() (*http.Response, error) {
	return r.ApiService.EditUserV2Execute(r)
}

/*
EditUserV2 Method for EditUserV2

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiEditUserV2Request
*/
func (a *UsersAPIService) EditUserV2(ctx context.Context) ApiEditUserV2Request {
	return ApiEditUserV2Request{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
func (a *UsersAPIService) EditUserV2Execute(r ApiEditUserV2Request) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "UsersAPIService.EditUserV2")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v2/users/edit"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.apiKey == nil {
		return nil, reportError("apiKey is required and must be specified")
	}
	if r.comMbridgeMsdkFoundationEntityCampaignExJSONKEYTIMESTAMP == nil {
		return nil, reportError("comMbridgeMsdkFoundationEntityCampaignExJSONKEYTIMESTAMP is required and must be specified")
	}
	if r.isPrivate == nil {
		return nil, reportError("isPrivate is required and must be specified")
	}
	if r.nickname == nil {
		return nil, reportError("nickname is required and must be specified")
	}
	if r.signedInfo == nil {
		return nil, reportError("signedInfo is required and must be specified")
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
	parameterAddToHeaderOrQuery(localVarFormParams, "com.mbridge.msdk.foundation.entity.CampaignEx.JSON_KEY_TIMESTAMP", r.comMbridgeMsdkFoundationEntityCampaignExJSONKEYTIMESTAMP, "", "")
	parameterAddToHeaderOrQuery(localVarFormParams, "is_private", r.isPrivate, "", "")
	parameterAddToHeaderOrQuery(localVarFormParams, "nickname", r.nickname, "", "")
	parameterAddToHeaderOrQuery(localVarFormParams, "signed_info", r.signedInfo, "", "")
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

type ApiEnableTwoFactorAuthRequest struct {
	ctx context.Context
	ApiService *UsersAPIService
	code *string
	type_ *string
}

func (r ApiEnableTwoFactorAuthRequest) Code(code string) ApiEnableTwoFactorAuthRequest {
	r.code = &code
	return r
}

func (r ApiEnableTwoFactorAuthRequest) Type_(type_ string) ApiEnableTwoFactorAuthRequest {
	r.type_ = &type_
	return r
}

func (r ApiEnableTwoFactorAuthRequest) Execute() (*TwoStepAuthEnabledResponse, *http.Response, error) {
	return r.ApiService.EnableTwoFactorAuthExecute(r)
}

/*
EnableTwoFactorAuth Method for EnableTwoFactorAuth

Implemented as a Kotlin suspend function on Android.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiEnableTwoFactorAuthRequest
*/
func (a *UsersAPIService) EnableTwoFactorAuth(ctx context.Context) ApiEnableTwoFactorAuthRequest {
	return ApiEnableTwoFactorAuthRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return TwoStepAuthEnabledResponse
func (a *UsersAPIService) EnableTwoFactorAuthExecute(r ApiEnableTwoFactorAuthRequest) (*TwoStepAuthEnabledResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPut
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *TwoStepAuthEnabledResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "UsersAPIService.EnableTwoFactorAuth")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/users/secret/enable"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.code == nil {
		return localVarReturnValue, nil, reportError("code is required and must be specified")
	}
	if r.type_ == nil {
		return localVarReturnValue, nil, reportError("type_ is required and must be specified")
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
	parameterAddToHeaderOrQuery(localVarFormParams, "code", r.code, "", "")
	parameterAddToHeaderOrQuery(localVarFormParams, "type", r.type_, "", "")
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

type ApiFollowUserRequest struct {
	ctx context.Context
	ApiService *UsersAPIService
	id int64
}

func (r ApiFollowUserRequest) Execute() (*http.Response, error) {
	return r.ApiService.FollowUserExecute(r)
}

/*
FollowUser Method for FollowUser

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id
 @return ApiFollowUserRequest
*/
func (a *UsersAPIService) FollowUser(ctx context.Context, id int64) ApiFollowUserRequest {
	return ApiFollowUserRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
func (a *UsersAPIService) FollowUserExecute(r ApiFollowUserRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "UsersAPIService.FollowUser")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v2/users/{id}/follow"
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

type ApiFollowUsersRequest struct {
	ctx context.Context
	ApiService *UsersAPIService
	userIds *[]int64
}

func (r ApiFollowUsersRequest) UserIds(userIds []int64) ApiFollowUsersRequest {
	r.userIds = &userIds
	return r
}

func (r ApiFollowUsersRequest) Execute() (*http.Response, error) {
	return r.ApiService.FollowUsersExecute(r)
}

/*
FollowUsers Method for FollowUsers

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiFollowUsersRequest
*/
func (a *UsersAPIService) FollowUsers(ctx context.Context) ApiFollowUsersRequest {
	return ApiFollowUsersRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
func (a *UsersAPIService) FollowUsersExecute(r ApiFollowUsersRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "UsersAPIService.FollowUsers")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v2/users/follow"

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

type ApiGetActiveFollowingsRequest struct {
	ctx context.Context
	ApiService *UsersAPIService
	onlyOnline *bool
	fromLoggedinAt *int64
}

func (r ApiGetActiveFollowingsRequest) OnlyOnline(onlyOnline bool) ApiGetActiveFollowingsRequest {
	r.onlyOnline = &onlyOnline
	return r
}

func (r ApiGetActiveFollowingsRequest) FromLoggedinAt(fromLoggedinAt int64) ApiGetActiveFollowingsRequest {
	r.fromLoggedinAt = &fromLoggedinAt
	return r
}

func (r ApiGetActiveFollowingsRequest) Execute() (*ActiveFollowingsResponse, *http.Response, error) {
	return r.ApiService.GetActiveFollowingsExecute(r)
}

/*
GetActiveFollowings Method for GetActiveFollowings

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiGetActiveFollowingsRequest
*/
func (a *UsersAPIService) GetActiveFollowings(ctx context.Context) ApiGetActiveFollowingsRequest {
	return ApiGetActiveFollowingsRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return ActiveFollowingsResponse
func (a *UsersAPIService) GetActiveFollowingsExecute(r ApiGetActiveFollowingsRequest) (*ActiveFollowingsResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ActiveFollowingsResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "UsersAPIService.GetActiveFollowings")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/users/active_followings"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.onlyOnline == nil {
		return localVarReturnValue, nil, reportError("onlyOnline is required and must be specified")
	}

	parameterAddToHeaderOrQuery(localVarQueryParams, "only_online", r.onlyOnline, "form", "")
	if r.fromLoggedinAt != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "from_loggedin_at", r.fromLoggedinAt, "form", "")
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

type ApiGetAdditionalNotificationSettingRequest struct {
	ctx context.Context
	ApiService *UsersAPIService
}

func (r ApiGetAdditionalNotificationSettingRequest) Execute() (*AdditionalSettingsResponse, *http.Response, error) {
	return r.ApiService.GetAdditionalNotificationSettingExecute(r)
}

/*
GetAdditionalNotificationSetting Method for GetAdditionalNotificationSetting

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiGetAdditionalNotificationSettingRequest
*/
func (a *UsersAPIService) GetAdditionalNotificationSetting(ctx context.Context) ApiGetAdditionalNotificationSettingRequest {
	return ApiGetAdditionalNotificationSettingRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return AdditionalSettingsResponse
func (a *UsersAPIService) GetAdditionalNotificationSettingExecute(r ApiGetAdditionalNotificationSettingRequest) (*AdditionalSettingsResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *AdditionalSettingsResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "UsersAPIService.GetAdditionalNotificationSetting")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/users/additonal_notification_setting"

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

type ApiGetBlockedUserIdsRequest struct {
	ctx context.Context
	ApiService *UsersAPIService
}

func (r ApiGetBlockedUserIdsRequest) Execute() (*BlockedUserIdsResponse, *http.Response, error) {
	return r.ApiService.GetBlockedUserIdsExecute(r)
}

/*
GetBlockedUserIds Method for GetBlockedUserIds

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiGetBlockedUserIdsRequest
*/
func (a *UsersAPIService) GetBlockedUserIds(ctx context.Context) ApiGetBlockedUserIdsRequest {
	return ApiGetBlockedUserIdsRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return BlockedUserIdsResponse
func (a *UsersAPIService) GetBlockedUserIdsExecute(r ApiGetBlockedUserIdsRequest) (*BlockedUserIdsResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *BlockedUserIdsResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "UsersAPIService.GetBlockedUserIds")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/users/block_ids"

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

type ApiGetBlockedUsersRequest struct {
	ctx context.Context
	ApiService *UsersAPIService
	searchUsersRequest *SearchUsersRequest
	fromId *int64
}

func (r ApiGetBlockedUsersRequest) SearchUsersRequest(searchUsersRequest SearchUsersRequest) ApiGetBlockedUsersRequest {
	r.searchUsersRequest = &searchUsersRequest
	return r
}

func (r ApiGetBlockedUsersRequest) FromId(fromId int64) ApiGetBlockedUsersRequest {
	r.fromId = &fromId
	return r
}

func (r ApiGetBlockedUsersRequest) Execute() (*BlockedUsersResponse, *http.Response, error) {
	return r.ApiService.GetBlockedUsersExecute(r)
}

/*
GetBlockedUsers Method for GetBlockedUsers

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiGetBlockedUsersRequest
*/
func (a *UsersAPIService) GetBlockedUsers(ctx context.Context) ApiGetBlockedUsersRequest {
	return ApiGetBlockedUsersRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return BlockedUsersResponse
func (a *UsersAPIService) GetBlockedUsersExecute(r ApiGetBlockedUsersRequest) (*BlockedUsersResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *BlockedUsersResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "UsersAPIService.GetBlockedUsers")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v2/users/blocked"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.searchUsersRequest == nil {
		return localVarReturnValue, nil, reportError("searchUsersRequest is required and must be specified")
	}

	if r.fromId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "from_id", r.fromId, "form", "")
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
	localVarPostBody = r.searchUsersRequest
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

type ApiGetBookmarkedPostsRequest struct {
	ctx context.Context
	ApiService *UsersAPIService
	userId int64
	androidxConstraintlayoutCoreMotionUtilsTypedValuesTransitionTypeSFROM *string
}

func (r ApiGetBookmarkedPostsRequest) AndroidxConstraintlayoutCoreMotionUtilsTypedValuesTransitionTypeSFROM(androidxConstraintlayoutCoreMotionUtilsTypedValuesTransitionTypeSFROM string) ApiGetBookmarkedPostsRequest {
	r.androidxConstraintlayoutCoreMotionUtilsTypedValuesTransitionTypeSFROM = &androidxConstraintlayoutCoreMotionUtilsTypedValuesTransitionTypeSFROM
	return r
}

func (r ApiGetBookmarkedPostsRequest) Execute() (*PostsResponse, *http.Response, error) {
	return r.ApiService.GetBookmarkedPostsExecute(r)
}

/*
GetBookmarkedPosts Method for GetBookmarkedPosts

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param userId
 @return ApiGetBookmarkedPostsRequest
*/
func (a *UsersAPIService) GetBookmarkedPosts(ctx context.Context, userId int64) ApiGetBookmarkedPostsRequest {
	return ApiGetBookmarkedPostsRequest{
		ApiService: a,
		ctx: ctx,
		userId: userId,
	}
}

// Execute executes the request
//  @return PostsResponse
func (a *UsersAPIService) GetBookmarkedPostsExecute(r ApiGetBookmarkedPostsRequest) (*PostsResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *PostsResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "UsersAPIService.GetBookmarkedPosts")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/users/{user_id}/bookmarks"
	localVarPath = strings.Replace(localVarPath, "{"+"user_id"+"}", url.PathEscape(parameterValueToString(r.userId, "userId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.androidxConstraintlayoutCoreMotionUtilsTypedValuesTransitionTypeSFROM != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "androidx.constraintlayout.core.motion.utils.TypedValues.TransitionType.S_FROM", r.androidxConstraintlayoutCoreMotionUtilsTypedValuesTransitionTypeSFROM, "form", "")
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

type ApiGetChatableFollowingsRequest struct {
	ctx context.Context
	ApiService *UsersAPIService
	searchUsersRequest *SearchUsersRequest
	fromFollowId *int64
	fromTimestamp *int64
	orderBy *string
}

func (r ApiGetChatableFollowingsRequest) SearchUsersRequest(searchUsersRequest SearchUsersRequest) ApiGetChatableFollowingsRequest {
	r.searchUsersRequest = &searchUsersRequest
	return r
}

func (r ApiGetChatableFollowingsRequest) FromFollowId(fromFollowId int64) ApiGetChatableFollowingsRequest {
	r.fromFollowId = &fromFollowId
	return r
}

func (r ApiGetChatableFollowingsRequest) FromTimestamp(fromTimestamp int64) ApiGetChatableFollowingsRequest {
	r.fromTimestamp = &fromTimestamp
	return r
}

func (r ApiGetChatableFollowingsRequest) OrderBy(orderBy string) ApiGetChatableFollowingsRequest {
	r.orderBy = &orderBy
	return r
}

func (r ApiGetChatableFollowingsRequest) Execute() (*FollowUsersResponse, *http.Response, error) {
	return r.ApiService.GetChatableFollowingsExecute(r)
}

/*
GetChatableFollowings Method for GetChatableFollowings

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiGetChatableFollowingsRequest
*/
func (a *UsersAPIService) GetChatableFollowings(ctx context.Context) ApiGetChatableFollowingsRequest {
	return ApiGetChatableFollowingsRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return FollowUsersResponse
func (a *UsersAPIService) GetChatableFollowingsExecute(r ApiGetChatableFollowingsRequest) (*FollowUsersResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *FollowUsersResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "UsersAPIService.GetChatableFollowings")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/users/followings/chatable"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.searchUsersRequest == nil {
		return localVarReturnValue, nil, reportError("searchUsersRequest is required and must be specified")
	}

	if r.fromFollowId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "from_follow_id", r.fromFollowId, "form", "")
	}
	if r.fromTimestamp != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "from_timestamp", r.fromTimestamp, "form", "")
	}
	if r.orderBy != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "order_by", r.orderBy, "form", "")
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
	localVarPostBody = r.searchUsersRequest
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

type ApiGetDefaultSettingsRequest struct {
	ctx context.Context
	ApiService *UsersAPIService
}

func (r ApiGetDefaultSettingsRequest) Execute() (*DefaultSettingsResponse, *http.Response, error) {
	return r.ApiService.GetDefaultSettingsExecute(r)
}

/*
GetDefaultSettings Method for GetDefaultSettings

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiGetDefaultSettingsRequest
*/
func (a *UsersAPIService) GetDefaultSettings(ctx context.Context) ApiGetDefaultSettingsRequest {
	return ApiGetDefaultSettingsRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return DefaultSettingsResponse
func (a *UsersAPIService) GetDefaultSettingsExecute(r ApiGetDefaultSettingsRequest) (*DefaultSettingsResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *DefaultSettingsResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "UsersAPIService.GetDefaultSettings")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/users/default_settings"

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

type ApiGetFollowRequestCountRequest struct {
	ctx context.Context
	ApiService *UsersAPIService
}

func (r ApiGetFollowRequestCountRequest) Execute() (*FollowRequestCountResponse, *http.Response, error) {
	return r.ApiService.GetFollowRequestCountExecute(r)
}

/*
GetFollowRequestCount Method for GetFollowRequestCount

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiGetFollowRequestCountRequest
*/
func (a *UsersAPIService) GetFollowRequestCount(ctx context.Context) ApiGetFollowRequestCountRequest {
	return ApiGetFollowRequestCountRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return FollowRequestCountResponse
func (a *UsersAPIService) GetFollowRequestCountExecute(r ApiGetFollowRequestCountRequest) (*FollowRequestCountResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *FollowRequestCountResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "UsersAPIService.GetFollowRequestCount")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v2/users/follow_requests_count"

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

type ApiGetFollowRequestsRequest struct {
	ctx context.Context
	ApiService *UsersAPIService
	fromTimestamp *int64
}

func (r ApiGetFollowRequestsRequest) FromTimestamp(fromTimestamp int64) ApiGetFollowRequestsRequest {
	r.fromTimestamp = &fromTimestamp
	return r
}

func (r ApiGetFollowRequestsRequest) Execute() (*UsersByTimestampResponse, *http.Response, error) {
	return r.ApiService.GetFollowRequestsExecute(r)
}

/*
GetFollowRequests Method for GetFollowRequests

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiGetFollowRequestsRequest
*/
func (a *UsersAPIService) GetFollowRequests(ctx context.Context) ApiGetFollowRequestsRequest {
	return ApiGetFollowRequestsRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return UsersByTimestampResponse
func (a *UsersAPIService) GetFollowRequestsExecute(r ApiGetFollowRequestsRequest) (*UsersByTimestampResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *UsersByTimestampResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "UsersAPIService.GetFollowRequests")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v2/users/follow_requests"

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

type ApiGetFollowingsBornTodayRequest struct {
	ctx context.Context
	ApiService *UsersAPIService
	birthdate *int64
}

func (r ApiGetFollowingsBornTodayRequest) Birthdate(birthdate int64) ApiGetFollowingsBornTodayRequest {
	r.birthdate = &birthdate
	return r
}

func (r ApiGetFollowingsBornTodayRequest) Execute() (*UsersResponse, *http.Response, error) {
	return r.ApiService.GetFollowingsBornTodayExecute(r)
}

/*
GetFollowingsBornToday Method for GetFollowingsBornToday

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiGetFollowingsBornTodayRequest
*/
func (a *UsersAPIService) GetFollowingsBornToday(ctx context.Context) ApiGetFollowingsBornTodayRequest {
	return ApiGetFollowingsBornTodayRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return UsersResponse
func (a *UsersAPIService) GetFollowingsBornTodayExecute(r ApiGetFollowingsBornTodayRequest) (*UsersResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *UsersResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "UsersAPIService.GetFollowingsBornToday")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/users/following_born_today"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.birthdate != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "birthdate", r.birthdate, "form", "")
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

type ApiGetFootprintsRequest struct {
	ctx context.Context
	ApiService *UsersAPIService
	number *int32
	mode *string
	androidxConstraintlayoutCoreMotionUtilsTypedValuesTransitionTypeSFROM *string
}

func (r ApiGetFootprintsRequest) Number(number int32) ApiGetFootprintsRequest {
	r.number = &number
	return r
}

func (r ApiGetFootprintsRequest) Mode(mode string) ApiGetFootprintsRequest {
	r.mode = &mode
	return r
}

func (r ApiGetFootprintsRequest) AndroidxConstraintlayoutCoreMotionUtilsTypedValuesTransitionTypeSFROM(androidxConstraintlayoutCoreMotionUtilsTypedValuesTransitionTypeSFROM string) ApiGetFootprintsRequest {
	r.androidxConstraintlayoutCoreMotionUtilsTypedValuesTransitionTypeSFROM = &androidxConstraintlayoutCoreMotionUtilsTypedValuesTransitionTypeSFROM
	return r
}

func (r ApiGetFootprintsRequest) Execute() (*FootprintsResponse, *http.Response, error) {
	return r.ApiService.GetFootprintsExecute(r)
}

/*
GetFootprints Method for GetFootprints

Implemented as a Kotlin suspend function on Android.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiGetFootprintsRequest
*/
func (a *UsersAPIService) GetFootprints(ctx context.Context) ApiGetFootprintsRequest {
	return ApiGetFootprintsRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return FootprintsResponse
func (a *UsersAPIService) GetFootprintsExecute(r ApiGetFootprintsRequest) (*FootprintsResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *FootprintsResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "UsersAPIService.GetFootprints")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v3/users/footprints"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.number == nil {
		return localVarReturnValue, nil, reportError("number is required and must be specified")
	}

	if r.mode != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "mode", r.mode, "form", "")
	}
	parameterAddToHeaderOrQuery(localVarQueryParams, "number", r.number, "form", "")
	if r.androidxConstraintlayoutCoreMotionUtilsTypedValuesTransitionTypeSFROM != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "androidx.constraintlayout.core.motion.utils.TypedValues.TransitionType.S_FROM", r.androidxConstraintlayoutCoreMotionUtilsTypedValuesTransitionTypeSFROM, "form", "")
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

type ApiGetFreshUserRequest struct {
	ctx context.Context
	ApiService *UsersAPIService
	id int64
}

func (r ApiGetFreshUserRequest) Execute() (*UserResponse, *http.Response, error) {
	return r.ApiService.GetFreshUserExecute(r)
}

/*
GetFreshUser Method for GetFreshUser

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id
 @return ApiGetFreshUserRequest
*/
func (a *UsersAPIService) GetFreshUser(ctx context.Context, id int64) ApiGetFreshUserRequest {
	return ApiGetFreshUserRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
//  @return UserResponse
func (a *UsersAPIService) GetFreshUserExecute(r ApiGetFreshUserRequest) (*UserResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *UserResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "UsersAPIService.GetFreshUser")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v2/users/fresh/{id}"
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

type ApiGetHimaUsersRequest struct {
	ctx context.Context
	ApiService *UsersAPIService
	fromHimaId *int64
	number *int32
}

func (r ApiGetHimaUsersRequest) FromHimaId(fromHimaId int64) ApiGetHimaUsersRequest {
	r.fromHimaId = &fromHimaId
	return r
}

func (r ApiGetHimaUsersRequest) Number(number int32) ApiGetHimaUsersRequest {
	r.number = &number
	return r
}

func (r ApiGetHimaUsersRequest) Execute() (*HimaUsersResponse, *http.Response, error) {
	return r.ApiService.GetHimaUsersExecute(r)
}

/*
GetHimaUsers Method for GetHimaUsers

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiGetHimaUsersRequest
*/
func (a *UsersAPIService) GetHimaUsers(ctx context.Context) ApiGetHimaUsersRequest {
	return ApiGetHimaUsersRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return HimaUsersResponse
func (a *UsersAPIService) GetHimaUsersExecute(r ApiGetHimaUsersRequest) (*HimaUsersResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *HimaUsersResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "UsersAPIService.GetHimaUsers")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v2/users/hima_users"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.fromHimaId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "from_hima_id", r.fromHimaId, "form", "")
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

type ApiGetMyReviewsRequest struct {
	ctx context.Context
	ApiService *UsersAPIService
	fromId *int64
}

func (r ApiGetMyReviewsRequest) FromId(fromId int64) ApiGetMyReviewsRequest {
	r.fromId = &fromId
	return r
}

func (r ApiGetMyReviewsRequest) Execute() (*ReviewsResponse, *http.Response, error) {
	return r.ApiService.GetMyReviewsExecute(r)
}

/*
GetMyReviews Method for GetMyReviews

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiGetMyReviewsRequest
*/
func (a *UsersAPIService) GetMyReviews(ctx context.Context) ApiGetMyReviewsRequest {
	return ApiGetMyReviewsRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return ReviewsResponse
func (a *UsersAPIService) GetMyReviewsExecute(r ApiGetMyReviewsRequest) (*ReviewsResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ReviewsResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "UsersAPIService.GetMyReviews")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/users/reviews/mine"

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

type ApiGetPolicyAgreementsRequest struct {
	ctx context.Context
	ApiService *UsersAPIService
}

func (r ApiGetPolicyAgreementsRequest) Execute() (*PolicyAgreementsResponse, *http.Response, error) {
	return r.ApiService.GetPolicyAgreementsExecute(r)
}

/*
GetPolicyAgreements Method for GetPolicyAgreements

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiGetPolicyAgreementsRequest
*/
func (a *UsersAPIService) GetPolicyAgreements(ctx context.Context) ApiGetPolicyAgreementsRequest {
	return ApiGetPolicyAgreementsRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return PolicyAgreementsResponse
func (a *UsersAPIService) GetPolicyAgreementsExecute(r ApiGetPolicyAgreementsRequest) (*PolicyAgreementsResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *PolicyAgreementsResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "UsersAPIService.GetPolicyAgreements")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/users/policy_agreements"

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

type ApiGetRecommendedFollowUsersRequest struct {
	ctx context.Context
	ApiService *UsersAPIService
	id int64
	number *int32
	page *int32
}

func (r ApiGetRecommendedFollowUsersRequest) Number(number int32) ApiGetRecommendedFollowUsersRequest {
	r.number = &number
	return r
}

func (r ApiGetRecommendedFollowUsersRequest) Page(page int32) ApiGetRecommendedFollowUsersRequest {
	r.page = &page
	return r
}

func (r ApiGetRecommendedFollowUsersRequest) Execute() (*UsersResponse, *http.Response, error) {
	return r.ApiService.GetRecommendedFollowUsersExecute(r)
}

/*
GetRecommendedFollowUsers Method for GetRecommendedFollowUsers

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id
 @return ApiGetRecommendedFollowUsersRequest
*/
func (a *UsersAPIService) GetRecommendedFollowUsers(ctx context.Context, id int64) ApiGetRecommendedFollowUsersRequest {
	return ApiGetRecommendedFollowUsersRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
//  @return UsersResponse
func (a *UsersAPIService) GetRecommendedFollowUsersExecute(r ApiGetRecommendedFollowUsersRequest) (*UsersResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *UsersResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "UsersAPIService.GetRecommendedFollowUsers")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/users/{id}/follow_recommended"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterValueToString(r.id, "id")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.number != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "number", r.number, "form", "")
	}
	if r.page != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "page", r.page, "form", "")
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

type ApiGetResetCountersRequest struct {
	ctx context.Context
	ApiService *UsersAPIService
}

func (r ApiGetResetCountersRequest) Execute() (*RefreshCounterRequestsResponse, *http.Response, error) {
	return r.ApiService.GetResetCountersExecute(r)
}

/*
GetResetCounters Method for GetResetCounters

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiGetResetCountersRequest
*/
func (a *UsersAPIService) GetResetCounters(ctx context.Context) ApiGetResetCountersRequest {
	return ApiGetResetCountersRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return RefreshCounterRequestsResponse
func (a *UsersAPIService) GetResetCountersExecute(r ApiGetResetCountersRequest) (*RefreshCounterRequestsResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *RefreshCounterRequestsResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "UsersAPIService.GetResetCounters")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/users/reset_counters"

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

type ApiGetTwoFactorAuthRequestInfoRequest struct {
	ctx context.Context
	ApiService *UsersAPIService
}

func (r ApiGetTwoFactorAuthRequestInfoRequest) Execute() (*TwoStepAuthRequestInfoResponse, *http.Response, error) {
	return r.ApiService.GetTwoFactorAuthRequestInfoExecute(r)
}

/*
GetTwoFactorAuthRequestInfo Method for GetTwoFactorAuthRequestInfo

Implemented as a Kotlin suspend function on Android.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiGetTwoFactorAuthRequestInfoRequest
*/
func (a *UsersAPIService) GetTwoFactorAuthRequestInfo(ctx context.Context) ApiGetTwoFactorAuthRequestInfoRequest {
	return ApiGetTwoFactorAuthRequestInfoRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return TwoStepAuthRequestInfoResponse
func (a *UsersAPIService) GetTwoFactorAuthRequestInfoExecute(r ApiGetTwoFactorAuthRequestInfoRequest) (*TwoStepAuthRequestInfoResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *TwoStepAuthRequestInfoResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "UsersAPIService.GetTwoFactorAuthRequestInfo")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/users/secret"

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

type ApiGetTwoFactorAuthStatusRequest struct {
	ctx context.Context
	ApiService *UsersAPIService
}

func (r ApiGetTwoFactorAuthStatusRequest) Execute() (*TwoFAStatusResponse, *http.Response, error) {
	return r.ApiService.GetTwoFactorAuthStatusExecute(r)
}

/*
GetTwoFactorAuthStatus Method for GetTwoFactorAuthStatus

Implemented as a Kotlin suspend function on Android.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiGetTwoFactorAuthStatusRequest
*/
func (a *UsersAPIService) GetTwoFactorAuthStatus(ctx context.Context) ApiGetTwoFactorAuthStatusRequest {
	return ApiGetTwoFactorAuthStatusRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return TwoFAStatusResponse
func (a *UsersAPIService) GetTwoFactorAuthStatusExecute(r ApiGetTwoFactorAuthStatusRequest) (*TwoFAStatusResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *TwoFAStatusResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "UsersAPIService.GetTwoFactorAuthStatus")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/users/secret/status"

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

type ApiGetUserRequest struct {
	ctx context.Context
	ApiService *UsersAPIService
	id int64
}

func (r ApiGetUserRequest) Execute() (*UserResponse, *http.Response, error) {
	return r.ApiService.GetUserExecute(r)
}

/*
GetUser Method for GetUser

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id
 @return ApiGetUserRequest
*/
func (a *UsersAPIService) GetUser(ctx context.Context, id int64) ApiGetUserRequest {
	return ApiGetUserRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
//  @return UserResponse
func (a *UsersAPIService) GetUserExecute(r ApiGetUserRequest) (*UserResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *UserResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "UsersAPIService.GetUser")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v2/users/{id}"
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

type ApiGetUserByQrRequest struct {
	ctx context.Context
	ApiService *UsersAPIService
	qr string
}

func (r ApiGetUserByQrRequest) Execute() (*UserResponse, *http.Response, error) {
	return r.ApiService.GetUserByQrExecute(r)
}

/*
GetUserByQr Method for GetUserByQr

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param qr
 @return ApiGetUserByQrRequest
*/
func (a *UsersAPIService) GetUserByQr(ctx context.Context, qr string) ApiGetUserByQrRequest {
	return ApiGetUserByQrRequest{
		ApiService: a,
		ctx: ctx,
		qr: qr,
	}
}

// Execute executes the request
//  @return UserResponse
func (a *UsersAPIService) GetUserByQrExecute(r ApiGetUserByQrRequest) (*UserResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *UserResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "UsersAPIService.GetUserByQr")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/users/qr_codes/{qr}"
	localVarPath = strings.Replace(localVarPath, "{"+"qr"+"}", url.PathEscape(parameterValueToString(r.qr, "qr")), -1)

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

type ApiGetUserCustomDefinitionsRequest struct {
	ctx context.Context
	ApiService *UsersAPIService
}

func (r ApiGetUserCustomDefinitionsRequest) Execute() (*UserCustomDefinitionsResponse, *http.Response, error) {
	return r.ApiService.GetUserCustomDefinitionsExecute(r)
}

/*
GetUserCustomDefinitions Method for GetUserCustomDefinitions

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiGetUserCustomDefinitionsRequest
*/
func (a *UsersAPIService) GetUserCustomDefinitions(ctx context.Context) ApiGetUserCustomDefinitionsRequest {
	return ApiGetUserCustomDefinitionsRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return UserCustomDefinitionsResponse
func (a *UsersAPIService) GetUserCustomDefinitionsExecute(r ApiGetUserCustomDefinitionsRequest) (*UserCustomDefinitionsResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *UserCustomDefinitionsResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "UsersAPIService.GetUserCustomDefinitions")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/users/custom_definitions"

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

type ApiGetUserFollowersRequest struct {
	ctx context.Context
	ApiService *UsersAPIService
	id int64
	androidxConstraintlayoutCoreMotionUtilsTypedValuesTransitionTypeSFROM *int64
	followedByMe *bool
	userNickname *string
}

func (r ApiGetUserFollowersRequest) AndroidxConstraintlayoutCoreMotionUtilsTypedValuesTransitionTypeSFROM(androidxConstraintlayoutCoreMotionUtilsTypedValuesTransitionTypeSFROM int64) ApiGetUserFollowersRequest {
	r.androidxConstraintlayoutCoreMotionUtilsTypedValuesTransitionTypeSFROM = &androidxConstraintlayoutCoreMotionUtilsTypedValuesTransitionTypeSFROM
	return r
}

func (r ApiGetUserFollowersRequest) FollowedByMe(followedByMe bool) ApiGetUserFollowersRequest {
	r.followedByMe = &followedByMe
	return r
}

func (r ApiGetUserFollowersRequest) UserNickname(userNickname string) ApiGetUserFollowersRequest {
	r.userNickname = &userNickname
	return r
}

func (r ApiGetUserFollowersRequest) Execute() (*FollowUsersResponse, *http.Response, error) {
	return r.ApiService.GetUserFollowersExecute(r)
}

/*
GetUserFollowers Method for GetUserFollowers

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id
 @return ApiGetUserFollowersRequest
*/
func (a *UsersAPIService) GetUserFollowers(ctx context.Context, id int64) ApiGetUserFollowersRequest {
	return ApiGetUserFollowersRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
//  @return FollowUsersResponse
func (a *UsersAPIService) GetUserFollowersExecute(r ApiGetUserFollowersRequest) (*FollowUsersResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *FollowUsersResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "UsersAPIService.GetUserFollowers")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v3/users/{id}/followers"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterValueToString(r.id, "id")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.androidxConstraintlayoutCoreMotionUtilsTypedValuesTransitionTypeSFROM != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "androidx.constraintlayout.core.motion.utils.TypedValues.TransitionType.S_FROM", r.androidxConstraintlayoutCoreMotionUtilsTypedValuesTransitionTypeSFROM, "form", "")
	}
	if r.followedByMe != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "followed_by_me", r.followedByMe, "form", "")
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

type ApiGetUserFollowingsRequest struct {
	ctx context.Context
	ApiService *UsersAPIService
	id int64
	androidxConstraintlayoutCoreMotionUtilsTypedValuesTransitionTypeSFROM *int64
	fromTimestamp *int64
	orderBy *string
	userNickname *string
}

func (r ApiGetUserFollowingsRequest) AndroidxConstraintlayoutCoreMotionUtilsTypedValuesTransitionTypeSFROM(androidxConstraintlayoutCoreMotionUtilsTypedValuesTransitionTypeSFROM int64) ApiGetUserFollowingsRequest {
	r.androidxConstraintlayoutCoreMotionUtilsTypedValuesTransitionTypeSFROM = &androidxConstraintlayoutCoreMotionUtilsTypedValuesTransitionTypeSFROM
	return r
}

func (r ApiGetUserFollowingsRequest) FromTimestamp(fromTimestamp int64) ApiGetUserFollowingsRequest {
	r.fromTimestamp = &fromTimestamp
	return r
}

func (r ApiGetUserFollowingsRequest) OrderBy(orderBy string) ApiGetUserFollowingsRequest {
	r.orderBy = &orderBy
	return r
}

func (r ApiGetUserFollowingsRequest) UserNickname(userNickname string) ApiGetUserFollowingsRequest {
	r.userNickname = &userNickname
	return r
}

func (r ApiGetUserFollowingsRequest) Execute() (*FollowUsersResponse, *http.Response, error) {
	return r.ApiService.GetUserFollowingsExecute(r)
}

/*
GetUserFollowings Method for GetUserFollowings

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id
 @return ApiGetUserFollowingsRequest
*/
func (a *UsersAPIService) GetUserFollowings(ctx context.Context, id int64) ApiGetUserFollowingsRequest {
	return ApiGetUserFollowingsRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
//  @return FollowUsersResponse
func (a *UsersAPIService) GetUserFollowingsExecute(r ApiGetUserFollowingsRequest) (*FollowUsersResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *FollowUsersResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "UsersAPIService.GetUserFollowings")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v3/users/{id}/followings"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterValueToString(r.id, "id")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.androidxConstraintlayoutCoreMotionUtilsTypedValuesTransitionTypeSFROM != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "androidx.constraintlayout.core.motion.utils.TypedValues.TransitionType.S_FROM", r.androidxConstraintlayoutCoreMotionUtilsTypedValuesTransitionTypeSFROM, "form", "")
	}
	if r.fromTimestamp != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "from_timestamp", r.fromTimestamp, "form", "")
	}
	if r.orderBy != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "order_by", r.orderBy, "form", "")
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

type ApiGetUserGiftTransactionsRequest struct {
	ctx context.Context
	ApiService *UsersAPIService
	userId int64
	number *int32
	androidxConstraintlayoutCoreMotionUtilsTypedValuesTransitionTypeSFROM *string
}

func (r ApiGetUserGiftTransactionsRequest) Number(number int32) ApiGetUserGiftTransactionsRequest {
	r.number = &number
	return r
}

func (r ApiGetUserGiftTransactionsRequest) AndroidxConstraintlayoutCoreMotionUtilsTypedValuesTransitionTypeSFROM(androidxConstraintlayoutCoreMotionUtilsTypedValuesTransitionTypeSFROM string) ApiGetUserGiftTransactionsRequest {
	r.androidxConstraintlayoutCoreMotionUtilsTypedValuesTransitionTypeSFROM = &androidxConstraintlayoutCoreMotionUtilsTypedValuesTransitionTypeSFROM
	return r
}

func (r ApiGetUserGiftTransactionsRequest) Execute() (*GiftTransactionsResponse, *http.Response, error) {
	return r.ApiService.GetUserGiftTransactionsExecute(r)
}

/*
GetUserGiftTransactions Method for GetUserGiftTransactions

Implemented as a Kotlin suspend function on Android.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param userId
 @return ApiGetUserGiftTransactionsRequest
*/
func (a *UsersAPIService) GetUserGiftTransactions(ctx context.Context, userId int64) ApiGetUserGiftTransactionsRequest {
	return ApiGetUserGiftTransactionsRequest{
		ApiService: a,
		ctx: ctx,
		userId: userId,
	}
}

// Execute executes the request
//  @return GiftTransactionsResponse
func (a *UsersAPIService) GetUserGiftTransactionsExecute(r ApiGetUserGiftTransactionsRequest) (*GiftTransactionsResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *GiftTransactionsResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "UsersAPIService.GetUserGiftTransactions")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/users/{user_id}/gift_transactions"
	localVarPath = strings.Replace(localVarPath, "{"+"user_id"+"}", url.PathEscape(parameterValueToString(r.userId, "userId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.number != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "number", r.number, "form", "")
	}
	if r.androidxConstraintlayoutCoreMotionUtilsTypedValuesTransitionTypeSFROM != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "androidx.constraintlayout.core.motion.utils.TypedValues.TransitionType.S_FROM", r.androidxConstraintlayoutCoreMotionUtilsTypedValuesTransitionTypeSFROM, "form", "")
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

type ApiGetUserInfoRequest struct {
	ctx context.Context
	ApiService *UsersAPIService
	id int64
}

func (r ApiGetUserInfoRequest) Execute() (*UserResponse, *http.Response, error) {
	return r.ApiService.GetUserInfoExecute(r)
}

/*
GetUserInfo Method for GetUserInfo

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id
 @return ApiGetUserInfoRequest
*/
func (a *UsersAPIService) GetUserInfo(ctx context.Context, id int64) ApiGetUserInfoRequest {
	return ApiGetUserInfoRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
//  @return UserResponse
func (a *UsersAPIService) GetUserInfoExecute(r ApiGetUserInfoRequest) (*UserResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *UserResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "UsersAPIService.GetUserInfo")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v2/users/info/{id}"
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

type ApiGetUserInterestsRequest struct {
	ctx context.Context
	ApiService *UsersAPIService
}

func (r ApiGetUserInterestsRequest) Execute() (*UserInterestsResponse, *http.Response, error) {
	return r.ApiService.GetUserInterestsExecute(r)
}

/*
GetUserInterests Method for GetUserInterests

Implemented as a Kotlin suspend function on Android.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiGetUserInterestsRequest
*/
func (a *UsersAPIService) GetUserInterests(ctx context.Context) ApiGetUserInterestsRequest {
	return ApiGetUserInterestsRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return UserInterestsResponse
func (a *UsersAPIService) GetUserInterestsExecute(r ApiGetUserInterestsRequest) (*UserInterestsResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *UserInterestsResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "UsersAPIService.GetUserInterests")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/users/interests"

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

type ApiGetUserPresignedUrlRequest struct {
	ctx context.Context
	ApiService *UsersAPIService
	videoFileName *string
}

func (r ApiGetUserPresignedUrlRequest) VideoFileName(videoFileName string) ApiGetUserPresignedUrlRequest {
	r.videoFileName = &videoFileName
	return r
}

func (r ApiGetUserPresignedUrlRequest) Execute() (*PresignedUrlResponse, *http.Response, error) {
	return r.ApiService.GetUserPresignedUrlExecute(r)
}

/*
GetUserPresignedUrl Method for GetUserPresignedUrl

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiGetUserPresignedUrlRequest
*/
func (a *UsersAPIService) GetUserPresignedUrl(ctx context.Context) ApiGetUserPresignedUrlRequest {
	return ApiGetUserPresignedUrlRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return PresignedUrlResponse
func (a *UsersAPIService) GetUserPresignedUrlExecute(r ApiGetUserPresignedUrlRequest) (*PresignedUrlResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *PresignedUrlResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "UsersAPIService.GetUserPresignedUrl")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/users/presigned_url"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.videoFileName == nil {
		return localVarReturnValue, nil, reportError("videoFileName is required and must be specified")
	}

	parameterAddToHeaderOrQuery(localVarQueryParams, "video_file_name", r.videoFileName, "form", "")
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

type ApiGetUserReviewsRequest struct {
	ctx context.Context
	ApiService *UsersAPIService
	id int64
	fromId *int64
}

func (r ApiGetUserReviewsRequest) FromId(fromId int64) ApiGetUserReviewsRequest {
	r.fromId = &fromId
	return r
}

func (r ApiGetUserReviewsRequest) Execute() (*ReviewsResponse, *http.Response, error) {
	return r.ApiService.GetUserReviewsExecute(r)
}

/*
GetUserReviews Method for GetUserReviews

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id
 @return ApiGetUserReviewsRequest
*/
func (a *UsersAPIService) GetUserReviews(ctx context.Context, id int64) ApiGetUserReviewsRequest {
	return ApiGetUserReviewsRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
//  @return ReviewsResponse
func (a *UsersAPIService) GetUserReviewsExecute(r ApiGetUserReviewsRequest) (*ReviewsResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ReviewsResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "UsersAPIService.GetUserReviews")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v2/users/reviews/{id}"
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

type ApiGetUserTimestampRequest struct {
	ctx context.Context
	ApiService *UsersAPIService
}

func (r ApiGetUserTimestampRequest) Execute() (*UserTimestampResponse, *http.Response, error) {
	return r.ApiService.GetUserTimestampExecute(r)
}

/*
GetUserTimestamp Method for GetUserTimestamp

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiGetUserTimestampRequest
*/
func (a *UsersAPIService) GetUserTimestamp(ctx context.Context) ApiGetUserTimestampRequest {
	return ApiGetUserTimestampRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return UserTimestampResponse
func (a *UsersAPIService) GetUserTimestampExecute(r ApiGetUserTimestampRequest) (*UserTimestampResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *UserTimestampResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "UsersAPIService.GetUserTimestamp")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v2/users/timestamp"

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

type ApiGetUsersByIdsRequest struct {
	ctx context.Context
	ApiService *UsersAPIService
	xJwt *string
	userIds *[]int64
}

func (r ApiGetUsersByIdsRequest) XJwt(xJwt string) ApiGetUsersByIdsRequest {
	r.xJwt = &xJwt
	return r
}

func (r ApiGetUsersByIdsRequest) UserIds(userIds []int64) ApiGetUsersByIdsRequest {
	r.userIds = &userIds
	return r
}

func (r ApiGetUsersByIdsRequest) Execute() (*UsersResponse, *http.Response, error) {
	return r.ApiService.GetUsersByIdsExecute(r)
}

/*
GetUsersByIds Method for GetUsersByIds

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiGetUsersByIdsRequest
*/
func (a *UsersAPIService) GetUsersByIds(ctx context.Context) ApiGetUsersByIdsRequest {
	return ApiGetUsersByIdsRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return UsersResponse
func (a *UsersAPIService) GetUsersByIdsExecute(r ApiGetUsersByIdsRequest) (*UsersResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *UsersResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "UsersAPIService.GetUsersByIds")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/users/list_id"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.xJwt == nil {
		return localVarReturnValue, nil, reportError("xJwt is required and must be specified")
	}
	if r.userIds == nil {
		return localVarReturnValue, nil, reportError("userIds is required and must be specified")
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
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	parameterAddToHeaderOrQuery(localVarHeaderParams, "X-Jwt", r.xJwt, "simple", "")
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

type ApiGetWebSocketTokenRequest struct {
	ctx context.Context
	ApiService *UsersAPIService
}

func (r ApiGetWebSocketTokenRequest) Execute() (*WebSocketTokenResponse, *http.Response, error) {
	return r.ApiService.GetWebSocketTokenExecute(r)
}

/*
GetWebSocketToken Method for GetWebSocketToken

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiGetWebSocketTokenRequest
*/
func (a *UsersAPIService) GetWebSocketToken(ctx context.Context) ApiGetWebSocketTokenRequest {
	return ApiGetWebSocketTokenRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return WebSocketTokenResponse
func (a *UsersAPIService) GetWebSocketTokenExecute(r ApiGetWebSocketTokenRequest) (*WebSocketTokenResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *WebSocketTokenResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "UsersAPIService.GetWebSocketToken")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/users/ws_token"

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

type ApiLoginWithEmailRequest struct {
	ctx context.Context
	ApiService *UsersAPIService
	loginEmailUserRequest *LoginEmailUserRequest
}

func (r ApiLoginWithEmailRequest) LoginEmailUserRequest(loginEmailUserRequest LoginEmailUserRequest) ApiLoginWithEmailRequest {
	r.loginEmailUserRequest = &loginEmailUserRequest
	return r
}

func (r ApiLoginWithEmailRequest) Execute() (*LoginUserResponse, *http.Response, error) {
	return r.ApiService.LoginWithEmailExecute(r)
}

/*
LoginWithEmail Method for LoginWithEmail

Implemented as a Kotlin suspend function on Android.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiLoginWithEmailRequest
*/
func (a *UsersAPIService) LoginWithEmail(ctx context.Context) ApiLoginWithEmailRequest {
	return ApiLoginWithEmailRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return LoginUserResponse
func (a *UsersAPIService) LoginWithEmailExecute(r ApiLoginWithEmailRequest) (*LoginUserResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *LoginUserResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "UsersAPIService.LoginWithEmail")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v3/users/login_with_email"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.loginEmailUserRequest == nil {
		return localVarReturnValue, nil, reportError("loginEmailUserRequest is required and must be specified")
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
	localVarPostBody = r.loginEmailUserRequest
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

type ApiLogoutRequest struct {
	ctx context.Context
	ApiService *UsersAPIService
	uuid *string
}

func (r ApiLogoutRequest) Uuid(uuid string) ApiLogoutRequest {
	r.uuid = &uuid
	return r
}

func (r ApiLogoutRequest) Execute() (*http.Response, error) {
	return r.ApiService.LogoutExecute(r)
}

/*
Logout Method for Logout

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiLogoutRequest
*/
func (a *UsersAPIService) Logout(ctx context.Context) ApiLogoutRequest {
	return ApiLogoutRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
func (a *UsersAPIService) LogoutExecute(r ApiLogoutRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "UsersAPIService.Logout")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/users/logout"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
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

type ApiPingAliveRequest struct {
	ctx context.Context
	ApiService *UsersAPIService
}

func (r ApiPingAliveRequest) Execute() (*http.Response, error) {
	return r.ApiService.PingAliveExecute(r)
}

/*
PingAlive Method for PingAlive

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiPingAliveRequest
*/
func (a *UsersAPIService) PingAlive(ctx context.Context) ApiPingAliveRequest {
	return ApiPingAliveRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
func (a *UsersAPIService) PingAliveExecute(r ApiPingAliveRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "UsersAPIService.PingAlive")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/users/alive"

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

type ApiRemoveCoverImageRequest struct {
	ctx context.Context
	ApiService *UsersAPIService
}

func (r ApiRemoveCoverImageRequest) Execute() (*http.Response, error) {
	return r.ApiService.RemoveCoverImageExecute(r)
}

/*
RemoveCoverImage Method for RemoveCoverImage

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiRemoveCoverImageRequest
*/
func (a *UsersAPIService) RemoveCoverImage(ctx context.Context) ApiRemoveCoverImageRequest {
	return ApiRemoveCoverImageRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
func (a *UsersAPIService) RemoveCoverImageExecute(r ApiRemoveCoverImageRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "UsersAPIService.RemoveCoverImage")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v2/users/remove_cover_image"

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

type ApiRemoveProfilePhotoRequest struct {
	ctx context.Context
	ApiService *UsersAPIService
}

func (r ApiRemoveProfilePhotoRequest) Execute() (*http.Response, error) {
	return r.ApiService.RemoveProfilePhotoExecute(r)
}

/*
RemoveProfilePhoto Method for RemoveProfilePhoto

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiRemoveProfilePhotoRequest
*/
func (a *UsersAPIService) RemoveProfilePhoto(ctx context.Context) ApiRemoveProfilePhotoRequest {
	return ApiRemoveProfilePhotoRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
func (a *UsersAPIService) RemoveProfilePhotoExecute(r ApiRemoveProfilePhotoRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "UsersAPIService.RemoveProfilePhoto")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v2/users/remove_profile_photo"

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

type ApiReplyToReviewRequest struct {
	ctx context.Context
	ApiService *UsersAPIService
	id int64
	apiKey *string
	comMbridgeMsdkFoundationEntityCampaignExJSONKEYTIMESTAMP *int64
	comment *string
	signedInfo *string
	uuid *string
	context *string
}

func (r ApiReplyToReviewRequest) ApiKey(apiKey string) ApiReplyToReviewRequest {
	r.apiKey = &apiKey
	return r
}

func (r ApiReplyToReviewRequest) ComMbridgeMsdkFoundationEntityCampaignExJSONKEYTIMESTAMP(comMbridgeMsdkFoundationEntityCampaignExJSONKEYTIMESTAMP int64) ApiReplyToReviewRequest {
	r.comMbridgeMsdkFoundationEntityCampaignExJSONKEYTIMESTAMP = &comMbridgeMsdkFoundationEntityCampaignExJSONKEYTIMESTAMP
	return r
}

func (r ApiReplyToReviewRequest) Comment(comment string) ApiReplyToReviewRequest {
	r.comment = &comment
	return r
}

func (r ApiReplyToReviewRequest) SignedInfo(signedInfo string) ApiReplyToReviewRequest {
	r.signedInfo = &signedInfo
	return r
}

func (r ApiReplyToReviewRequest) Uuid(uuid string) ApiReplyToReviewRequest {
	r.uuid = &uuid
	return r
}

func (r ApiReplyToReviewRequest) Context(context string) ApiReplyToReviewRequest {
	r.context = &context
	return r
}

func (r ApiReplyToReviewRequest) Execute() (*http.Response, error) {
	return r.ApiService.ReplyToReviewExecute(r)
}

/*
ReplyToReview Method for ReplyToReview

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id
 @return ApiReplyToReviewRequest
*/
func (a *UsersAPIService) ReplyToReview(ctx context.Context, id int64) ApiReplyToReviewRequest {
	return ApiReplyToReviewRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
func (a *UsersAPIService) ReplyToReviewExecute(r ApiReplyToReviewRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "UsersAPIService.ReplyToReview")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v2/users/reviews/{id}"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterValueToString(r.id, "id")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.apiKey == nil {
		return nil, reportError("apiKey is required and must be specified")
	}
	if r.comMbridgeMsdkFoundationEntityCampaignExJSONKEYTIMESTAMP == nil {
		return nil, reportError("comMbridgeMsdkFoundationEntityCampaignExJSONKEYTIMESTAMP is required and must be specified")
	}
	if r.comment == nil {
		return nil, reportError("comment is required and must be specified")
	}
	if r.signedInfo == nil {
		return nil, reportError("signedInfo is required and must be specified")
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
	parameterAddToHeaderOrQuery(localVarFormParams, "com.mbridge.msdk.foundation.entity.CampaignEx.JSON_KEY_TIMESTAMP", r.comMbridgeMsdkFoundationEntityCampaignExJSONKEYTIMESTAMP, "", "")
	parameterAddToHeaderOrQuery(localVarFormParams, "comment", r.comment, "", "")
	if r.context != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "context", r.context, "", "")
	}
	parameterAddToHeaderOrQuery(localVarFormParams, "signed_info", r.signedInfo, "", "")
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

type ApiReportUserRequest struct {
	ctx context.Context
	ApiService *UsersAPIService
	userId int64
	categoryId *int64
	reason *string
	screenshot2Filename *string
	screenshot3Filename *string
	screenshot4Filename *string
	screenshotFilename *string
}

func (r ApiReportUserRequest) CategoryId(categoryId int64) ApiReportUserRequest {
	r.categoryId = &categoryId
	return r
}

func (r ApiReportUserRequest) Reason(reason string) ApiReportUserRequest {
	r.reason = &reason
	return r
}

func (r ApiReportUserRequest) Screenshot2Filename(screenshot2Filename string) ApiReportUserRequest {
	r.screenshot2Filename = &screenshot2Filename
	return r
}

func (r ApiReportUserRequest) Screenshot3Filename(screenshot3Filename string) ApiReportUserRequest {
	r.screenshot3Filename = &screenshot3Filename
	return r
}

func (r ApiReportUserRequest) Screenshot4Filename(screenshot4Filename string) ApiReportUserRequest {
	r.screenshot4Filename = &screenshot4Filename
	return r
}

func (r ApiReportUserRequest) ScreenshotFilename(screenshotFilename string) ApiReportUserRequest {
	r.screenshotFilename = &screenshotFilename
	return r
}

func (r ApiReportUserRequest) Execute() (*http.Response, error) {
	return r.ApiService.ReportUserExecute(r)
}

/*
ReportUser Method for ReportUser

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param userId
 @return ApiReportUserRequest
*/
func (a *UsersAPIService) ReportUser(ctx context.Context, userId int64) ApiReportUserRequest {
	return ApiReportUserRequest{
		ApiService: a,
		ctx: ctx,
		userId: userId,
	}
}

// Execute executes the request
func (a *UsersAPIService) ReportUserExecute(r ApiReportUserRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "UsersAPIService.ReportUser")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v3/users/{user_id}/report"
	localVarPath = strings.Replace(localVarPath, "{"+"user_id"+"}", url.PathEscape(parameterValueToString(r.userId, "userId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.categoryId == nil {
		return nil, reportError("categoryId is required and must be specified")
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

type ApiRequestFollowRequest struct {
	ctx context.Context
	ApiService *UsersAPIService
	targetId int64
	comUnity3dAdsCoreDomainHandleInvocationsFromAdViewerKEYACTION *string
}

func (r ApiRequestFollowRequest) ComUnity3dAdsCoreDomainHandleInvocationsFromAdViewerKEYACTION(comUnity3dAdsCoreDomainHandleInvocationsFromAdViewerKEYACTION string) ApiRequestFollowRequest {
	r.comUnity3dAdsCoreDomainHandleInvocationsFromAdViewerKEYACTION = &comUnity3dAdsCoreDomainHandleInvocationsFromAdViewerKEYACTION
	return r
}

func (r ApiRequestFollowRequest) Execute() (*http.Response, error) {
	return r.ApiService.RequestFollowExecute(r)
}

/*
RequestFollow Method for RequestFollow

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param targetId
 @return ApiRequestFollowRequest
*/
func (a *UsersAPIService) RequestFollow(ctx context.Context, targetId int64) ApiRequestFollowRequest {
	return ApiRequestFollowRequest{
		ApiService: a,
		ctx: ctx,
		targetId: targetId,
	}
}

// Execute executes the request
func (a *UsersAPIService) RequestFollowExecute(r ApiRequestFollowRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "UsersAPIService.RequestFollow")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v2/users/{target_id}/follow_request"
	localVarPath = strings.Replace(localVarPath, "{"+"target_id"+"}", url.PathEscape(parameterValueToString(r.targetId, "targetId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.comUnity3dAdsCoreDomainHandleInvocationsFromAdViewerKEYACTION == nil {
		return nil, reportError("comUnity3dAdsCoreDomainHandleInvocationsFromAdViewerKEYACTION is required and must be specified")
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
	parameterAddToHeaderOrQuery(localVarFormParams, "com.unity3d.ads.core.domain.HandleInvocationsFromAdViewer.KEY_ACTION", r.comUnity3dAdsCoreDomainHandleInvocationsFromAdViewerKEYACTION, "", "")
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

type ApiResendConfirmEmailRequest struct {
	ctx context.Context
	ApiService *UsersAPIService
}

func (r ApiResendConfirmEmailRequest) Execute() (*http.Response, error) {
	return r.ApiService.ResendConfirmEmailExecute(r)
}

/*
ResendConfirmEmail Method for ResendConfirmEmail

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiResendConfirmEmailRequest
*/
func (a *UsersAPIService) ResendConfirmEmail(ctx context.Context) ApiResendConfirmEmailRequest {
	return ApiResendConfirmEmailRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
func (a *UsersAPIService) ResendConfirmEmailExecute(r ApiResendConfirmEmailRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "UsersAPIService.ResendConfirmEmail")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v2/users/resend_confirm_email"

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

type ApiResetCountersRequest struct {
	ctx context.Context
	ApiService *UsersAPIService
	counter *string
}

func (r ApiResetCountersRequest) Counter(counter string) ApiResetCountersRequest {
	r.counter = &counter
	return r
}

func (r ApiResetCountersRequest) Execute() (*http.Response, error) {
	return r.ApiService.ResetCountersExecute(r)
}

/*
ResetCounters Method for ResetCounters

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiResetCountersRequest
*/
func (a *UsersAPIService) ResetCounters(ctx context.Context) ApiResetCountersRequest {
	return ApiResetCountersRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
func (a *UsersAPIService) ResetCountersExecute(r ApiResetCountersRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "UsersAPIService.ResetCounters")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/users/reset_counters"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.counter == nil {
		return nil, reportError("counter is required and must be specified")
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
	parameterAddToHeaderOrQuery(localVarFormParams, "counter", r.counter, "", "")
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

type ApiResetPasswordRequest struct {
	ctx context.Context
	ApiService *UsersAPIService
	androidxAutofillHintConstantsAUTOFILLHINTPASSWORD *string
	email *string
	emailGrantToken *string
}

func (r ApiResetPasswordRequest) AndroidxAutofillHintConstantsAUTOFILLHINTPASSWORD(androidxAutofillHintConstantsAUTOFILLHINTPASSWORD string) ApiResetPasswordRequest {
	r.androidxAutofillHintConstantsAUTOFILLHINTPASSWORD = &androidxAutofillHintConstantsAUTOFILLHINTPASSWORD
	return r
}

func (r ApiResetPasswordRequest) Email(email string) ApiResetPasswordRequest {
	r.email = &email
	return r
}

func (r ApiResetPasswordRequest) EmailGrantToken(emailGrantToken string) ApiResetPasswordRequest {
	r.emailGrantToken = &emailGrantToken
	return r
}

func (r ApiResetPasswordRequest) Execute() (*http.Response, error) {
	return r.ApiService.ResetPasswordExecute(r)
}

/*
ResetPassword Method for ResetPassword

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiResetPasswordRequest
*/
func (a *UsersAPIService) ResetPassword(ctx context.Context) ApiResetPasswordRequest {
	return ApiResetPasswordRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
func (a *UsersAPIService) ResetPasswordExecute(r ApiResetPasswordRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPut
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "UsersAPIService.ResetPassword")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/users/reset_password"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.androidxAutofillHintConstantsAUTOFILLHINTPASSWORD == nil {
		return nil, reportError("androidxAutofillHintConstantsAUTOFILLHINTPASSWORD is required and must be specified")
	}
	if r.email == nil {
		return nil, reportError("email is required and must be specified")
	}
	if r.emailGrantToken == nil {
		return nil, reportError("emailGrantToken is required and must be specified")
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
	parameterAddToHeaderOrQuery(localVarFormParams, "androidx.autofill.HintConstants.AUTOFILL_HINT_PASSWORD", r.androidxAutofillHintConstantsAUTOFILLHINTPASSWORD, "", "")
	parameterAddToHeaderOrQuery(localVarFormParams, "email", r.email, "", "")
	parameterAddToHeaderOrQuery(localVarFormParams, "email_grant_token", r.emailGrantToken, "", "")
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

type ApiSearchUsersRequest struct {
	ctx context.Context
	ApiService *UsersAPIService
	gender *int32
	nickname *string
	title *string
	biography *string
	fromTimestamp *int64
	similarAge *bool
	notRecentGomimushi *bool
	recentlyCreated *bool
	samePrefecture *bool
	saveRecentSearch *bool
}

func (r ApiSearchUsersRequest) Gender(gender int32) ApiSearchUsersRequest {
	r.gender = &gender
	return r
}

func (r ApiSearchUsersRequest) Nickname(nickname string) ApiSearchUsersRequest {
	r.nickname = &nickname
	return r
}

func (r ApiSearchUsersRequest) Title(title string) ApiSearchUsersRequest {
	r.title = &title
	return r
}

func (r ApiSearchUsersRequest) Biography(biography string) ApiSearchUsersRequest {
	r.biography = &biography
	return r
}

func (r ApiSearchUsersRequest) FromTimestamp(fromTimestamp int64) ApiSearchUsersRequest {
	r.fromTimestamp = &fromTimestamp
	return r
}

func (r ApiSearchUsersRequest) SimilarAge(similarAge bool) ApiSearchUsersRequest {
	r.similarAge = &similarAge
	return r
}

func (r ApiSearchUsersRequest) NotRecentGomimushi(notRecentGomimushi bool) ApiSearchUsersRequest {
	r.notRecentGomimushi = &notRecentGomimushi
	return r
}

func (r ApiSearchUsersRequest) RecentlyCreated(recentlyCreated bool) ApiSearchUsersRequest {
	r.recentlyCreated = &recentlyCreated
	return r
}

func (r ApiSearchUsersRequest) SamePrefecture(samePrefecture bool) ApiSearchUsersRequest {
	r.samePrefecture = &samePrefecture
	return r
}

func (r ApiSearchUsersRequest) SaveRecentSearch(saveRecentSearch bool) ApiSearchUsersRequest {
	r.saveRecentSearch = &saveRecentSearch
	return r
}

func (r ApiSearchUsersRequest) Execute() (*UsersResponse, *http.Response, error) {
	return r.ApiService.SearchUsersExecute(r)
}

/*
SearchUsers Method for SearchUsers

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiSearchUsersRequest
*/
func (a *UsersAPIService) SearchUsers(ctx context.Context) ApiSearchUsersRequest {
	return ApiSearchUsersRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return UsersResponse
func (a *UsersAPIService) SearchUsersExecute(r ApiSearchUsersRequest) (*UsersResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *UsersResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "UsersAPIService.SearchUsers")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/users/search"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.gender != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "gender", r.gender, "form", "")
	}
	if r.nickname != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "nickname", r.nickname, "form", "")
	}
	if r.title != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "title", r.title, "form", "")
	}
	if r.biography != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "biography", r.biography, "form", "")
	}
	if r.fromTimestamp != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "from_timestamp", r.fromTimestamp, "form", "")
	}
	if r.similarAge != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "similar_age", r.similarAge, "form", "")
	}
	if r.notRecentGomimushi != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "not_recent_gomimushi", r.notRecentGomimushi, "form", "")
	}
	if r.recentlyCreated != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "recently_created", r.recentlyCreated, "form", "")
	}
	if r.samePrefecture != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "same_prefecture", r.samePrefecture, "form", "")
	}
	if r.saveRecentSearch != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "save_recent_search", r.saveRecentSearch, "form", "")
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

type ApiSetHimaRequest struct {
	ctx context.Context
	ApiService *UsersAPIService
}

func (r ApiSetHimaRequest) Execute() (*http.Response, error) {
	return r.ApiService.SetHimaExecute(r)
}

/*
SetHima Method for SetHima

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiSetHimaRequest
*/
func (a *UsersAPIService) SetHima(ctx context.Context) ApiSetHimaRequest {
	return ApiSetHimaRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
func (a *UsersAPIService) SetHimaExecute(r ApiSetHimaRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "UsersAPIService.SetHima")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/users/hima"

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

type ApiUnblockUserRequest struct {
	ctx context.Context
	ApiService *UsersAPIService
	id int64
}

func (r ApiUnblockUserRequest) Execute() (*http.Response, error) {
	return r.ApiService.UnblockUserExecute(r)
}

/*
UnblockUser Method for UnblockUser

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id
 @return ApiUnblockUserRequest
*/
func (a *UsersAPIService) UnblockUser(ctx context.Context, id int64) ApiUnblockUserRequest {
	return ApiUnblockUserRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
func (a *UsersAPIService) UnblockUserExecute(r ApiUnblockUserRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "UsersAPIService.UnblockUser")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v2/users/{id}/unblock"
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

type ApiUnfollowUserRequest struct {
	ctx context.Context
	ApiService *UsersAPIService
	id int64
}

func (r ApiUnfollowUserRequest) Execute() (*http.Response, error) {
	return r.ApiService.UnfollowUserExecute(r)
}

/*
UnfollowUser Method for UnfollowUser

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id
 @return ApiUnfollowUserRequest
*/
func (a *UsersAPIService) UnfollowUser(ctx context.Context, id int64) ApiUnfollowUserRequest {
	return ApiUnfollowUserRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
func (a *UsersAPIService) UnfollowUserExecute(r ApiUnfollowUserRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "UsersAPIService.UnfollowUser")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v2/users/{id}/unfollow"
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

type ApiUpdateAdditionalNotificationSettingRequest struct {
	ctx context.Context
	ApiService *UsersAPIService
	mode *string
	on *int32
}

func (r ApiUpdateAdditionalNotificationSettingRequest) Mode(mode string) ApiUpdateAdditionalNotificationSettingRequest {
	r.mode = &mode
	return r
}

func (r ApiUpdateAdditionalNotificationSettingRequest) On(on int32) ApiUpdateAdditionalNotificationSettingRequest {
	r.on = &on
	return r
}

func (r ApiUpdateAdditionalNotificationSettingRequest) Execute() (*http.Response, error) {
	return r.ApiService.UpdateAdditionalNotificationSettingExecute(r)
}

/*
UpdateAdditionalNotificationSetting Method for UpdateAdditionalNotificationSetting

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiUpdateAdditionalNotificationSettingRequest
*/
func (a *UsersAPIService) UpdateAdditionalNotificationSetting(ctx context.Context) ApiUpdateAdditionalNotificationSettingRequest {
	return ApiUpdateAdditionalNotificationSettingRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
func (a *UsersAPIService) UpdateAdditionalNotificationSettingExecute(r ApiUpdateAdditionalNotificationSettingRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "UsersAPIService.UpdateAdditionalNotificationSetting")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/users/additonal_notification_setting"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.mode == nil {
		return nil, reportError("mode is required and must be specified")
	}
	if r.on == nil {
		return nil, reportError("on is required and must be specified")
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
	parameterAddToHeaderOrQuery(localVarFormParams, "mode", r.mode, "", "")
	parameterAddToHeaderOrQuery(localVarFormParams, "on", r.on, "", "")
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

type ApiUpdateLanguageRequest struct {
	ctx context.Context
	ApiService *UsersAPIService
	apiKey *string
	comMbridgeMsdkFoundationEntityCampaignExJSONKEYTIMESTAMP *int64
	language *string
	signedInfo *string
	uuid *string
}

func (r ApiUpdateLanguageRequest) ApiKey(apiKey string) ApiUpdateLanguageRequest {
	r.apiKey = &apiKey
	return r
}

func (r ApiUpdateLanguageRequest) ComMbridgeMsdkFoundationEntityCampaignExJSONKEYTIMESTAMP(comMbridgeMsdkFoundationEntityCampaignExJSONKEYTIMESTAMP int64) ApiUpdateLanguageRequest {
	r.comMbridgeMsdkFoundationEntityCampaignExJSONKEYTIMESTAMP = &comMbridgeMsdkFoundationEntityCampaignExJSONKEYTIMESTAMP
	return r
}

func (r ApiUpdateLanguageRequest) Language(language string) ApiUpdateLanguageRequest {
	r.language = &language
	return r
}

func (r ApiUpdateLanguageRequest) SignedInfo(signedInfo string) ApiUpdateLanguageRequest {
	r.signedInfo = &signedInfo
	return r
}

func (r ApiUpdateLanguageRequest) Uuid(uuid string) ApiUpdateLanguageRequest {
	r.uuid = &uuid
	return r
}

func (r ApiUpdateLanguageRequest) Execute() (*http.Response, error) {
	return r.ApiService.UpdateLanguageExecute(r)
}

/*
UpdateLanguage Method for UpdateLanguage

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiUpdateLanguageRequest
*/
func (a *UsersAPIService) UpdateLanguage(ctx context.Context) ApiUpdateLanguageRequest {
	return ApiUpdateLanguageRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
func (a *UsersAPIService) UpdateLanguageExecute(r ApiUpdateLanguageRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPut
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "UsersAPIService.UpdateLanguage")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/users/language"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.apiKey == nil {
		return nil, reportError("apiKey is required and must be specified")
	}
	if r.comMbridgeMsdkFoundationEntityCampaignExJSONKEYTIMESTAMP == nil {
		return nil, reportError("comMbridgeMsdkFoundationEntityCampaignExJSONKEYTIMESTAMP is required and must be specified")
	}
	if r.language == nil {
		return nil, reportError("language is required and must be specified")
	}
	if r.signedInfo == nil {
		return nil, reportError("signedInfo is required and must be specified")
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
	parameterAddToHeaderOrQuery(localVarFormParams, "com.mbridge.msdk.foundation.entity.CampaignEx.JSON_KEY_TIMESTAMP", r.comMbridgeMsdkFoundationEntityCampaignExJSONKEYTIMESTAMP, "", "")
	parameterAddToHeaderOrQuery(localVarFormParams, "language", r.language, "", "")
	parameterAddToHeaderOrQuery(localVarFormParams, "signed_info", r.signedInfo, "", "")
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

type ApiUpdateLoginRequest struct {
	ctx context.Context
	ApiService *UsersAPIService
	apiKey *string
	email *string
	androidxAutofillHintConstantsAUTOFILLHINTPASSWORD *string
	currentPassword *string
	emailGrantToken *string
}

func (r ApiUpdateLoginRequest) ApiKey(apiKey string) ApiUpdateLoginRequest {
	r.apiKey = &apiKey
	return r
}

func (r ApiUpdateLoginRequest) Email(email string) ApiUpdateLoginRequest {
	r.email = &email
	return r
}

func (r ApiUpdateLoginRequest) AndroidxAutofillHintConstantsAUTOFILLHINTPASSWORD(androidxAutofillHintConstantsAUTOFILLHINTPASSWORD string) ApiUpdateLoginRequest {
	r.androidxAutofillHintConstantsAUTOFILLHINTPASSWORD = &androidxAutofillHintConstantsAUTOFILLHINTPASSWORD
	return r
}

func (r ApiUpdateLoginRequest) CurrentPassword(currentPassword string) ApiUpdateLoginRequest {
	r.currentPassword = &currentPassword
	return r
}

func (r ApiUpdateLoginRequest) EmailGrantToken(emailGrantToken string) ApiUpdateLoginRequest {
	r.emailGrantToken = &emailGrantToken
	return r
}

func (r ApiUpdateLoginRequest) Execute() (*LoginUpdateResponse, *http.Response, error) {
	return r.ApiService.UpdateLoginExecute(r)
}

/*
UpdateLogin Method for UpdateLogin

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiUpdateLoginRequest
*/
func (a *UsersAPIService) UpdateLogin(ctx context.Context) ApiUpdateLoginRequest {
	return ApiUpdateLoginRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return LoginUpdateResponse
func (a *UsersAPIService) UpdateLoginExecute(r ApiUpdateLoginRequest) (*LoginUpdateResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *LoginUpdateResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "UsersAPIService.UpdateLogin")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v3/users/login_update"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.apiKey == nil {
		return localVarReturnValue, nil, reportError("apiKey is required and must be specified")
	}
	if r.email == nil {
		return localVarReturnValue, nil, reportError("email is required and must be specified")
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
	if r.androidxAutofillHintConstantsAUTOFILLHINTPASSWORD != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "androidx.autofill.HintConstants.AUTOFILL_HINT_PASSWORD", r.androidxAutofillHintConstantsAUTOFILLHINTPASSWORD, "", "")
	}
	parameterAddToHeaderOrQuery(localVarFormParams, "api_key", r.apiKey, "", "")
	if r.currentPassword != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "current_password", r.currentPassword, "", "")
	}
	parameterAddToHeaderOrQuery(localVarFormParams, "email", r.email, "", "")
	if r.emailGrantToken != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "email_grant_token", r.emailGrantToken, "", "")
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

type ApiUpdateUserInterestsRequest struct {
	ctx context.Context
	ApiService *UsersAPIService
	commonIdsRequest *CommonIdsRequest
}

func (r ApiUpdateUserInterestsRequest) CommonIdsRequest(commonIdsRequest CommonIdsRequest) ApiUpdateUserInterestsRequest {
	r.commonIdsRequest = &commonIdsRequest
	return r
}

func (r ApiUpdateUserInterestsRequest) Execute() (*http.Response, error) {
	return r.ApiService.UpdateUserInterestsExecute(r)
}

/*
UpdateUserInterests Method for UpdateUserInterests

Implemented as a Kotlin suspend function on Android.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiUpdateUserInterestsRequest
*/
func (a *UsersAPIService) UpdateUserInterests(ctx context.Context) ApiUpdateUserInterestsRequest {
	return ApiUpdateUserInterestsRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
func (a *UsersAPIService) UpdateUserInterestsExecute(r ApiUpdateUserInterestsRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPut
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "UsersAPIService.UpdateUserInterests")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/users/interests"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.commonIdsRequest == nil {
		return nil, reportError("commonIdsRequest is required and must be specified")
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
	localVarPostBody = r.commonIdsRequest
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
