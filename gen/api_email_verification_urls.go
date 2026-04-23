
// Code generated from the Yay! API schema; DO NOT EDIT.

package gen

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
)


// EmailVerificationUrlsAPIService EmailVerificationUrlsAPI service
type EmailVerificationUrlsAPIService service

type ApiRequestEmailVerificationRequest struct {
	ctx context.Context
	ApiService *EmailVerificationUrlsAPIService
	deviceUuid *string
	email *string
	locale *string
	intent *string
}

func (r ApiRequestEmailVerificationRequest) DeviceUuid(deviceUuid string) ApiRequestEmailVerificationRequest {
	r.deviceUuid = &deviceUuid
	return r
}

func (r ApiRequestEmailVerificationRequest) Email(email string) ApiRequestEmailVerificationRequest {
	r.email = &email
	return r
}

func (r ApiRequestEmailVerificationRequest) Locale(locale string) ApiRequestEmailVerificationRequest {
	r.locale = &locale
	return r
}

func (r ApiRequestEmailVerificationRequest) Intent(intent string) ApiRequestEmailVerificationRequest {
	r.intent = &intent
	return r
}

func (r ApiRequestEmailVerificationRequest) Execute() (*CommonUrlResponse, *http.Response, error) {
	return r.ApiService.RequestEmailVerificationExecute(r)
}

/*
RequestEmailVerification Method for RequestEmailVerification

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiRequestEmailVerificationRequest
*/
func (a *EmailVerificationUrlsAPIService) RequestEmailVerification(ctx context.Context) ApiRequestEmailVerificationRequest {
	return ApiRequestEmailVerificationRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return CommonUrlResponse
func (a *EmailVerificationUrlsAPIService) RequestEmailVerificationExecute(r ApiRequestEmailVerificationRequest) (*CommonUrlResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *CommonUrlResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "EmailVerificationUrlsAPIService.RequestEmailVerification")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/email_verification_urls"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.deviceUuid == nil {
		return localVarReturnValue, nil, reportError("deviceUuid is required and must be specified")
	}
	if r.email == nil {
		return localVarReturnValue, nil, reportError("email is required and must be specified")
	}
	if r.locale == nil {
		return localVarReturnValue, nil, reportError("locale is required and must be specified")
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
	parameterAddToHeaderOrQuery(localVarFormParams, "device_uuid", r.deviceUuid, "", "")
	parameterAddToHeaderOrQuery(localVarFormParams, "email", r.email, "", "")
	if r.intent != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "intent", r.intent, "", "")
	}
	parameterAddToHeaderOrQuery(localVarFormParams, "locale", r.locale, "", "")
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
