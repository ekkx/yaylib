
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


// ReceivedGiftsAPIService ReceivedGiftsAPI service
type ReceivedGiftsAPIService service

type ApiGetReceivedGiftSendersRequest struct {
	ctx context.Context
	ApiService *ReceivedGiftsAPIService
	giftId int64
	number *int32
}

func (r ApiGetReceivedGiftSendersRequest) Number(number int32) ApiGetReceivedGiftSendersRequest {
	r.number = &number
	return r
}

func (r ApiGetReceivedGiftSendersRequest) Execute() (*GiftSendersResponse, *http.Response, error) {
	return r.ApiService.GetReceivedGiftSendersExecute(r)
}

func (r ApiGetReceivedGiftSendersRequest) ExecuteRaw() ([]byte, *http.Response, error) {
	_, httpResp, err := r.Execute()
	return executeRaw(httpResp, err)
}

/*
GetReceivedGiftSenders Method for GetReceivedGiftSenders

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param giftId
 @return ApiGetReceivedGiftSendersRequest
*/
func (a *ReceivedGiftsAPIService) GetReceivedGiftSenders(ctx context.Context, giftId int64) ApiGetReceivedGiftSendersRequest {
	return ApiGetReceivedGiftSendersRequest{
		ApiService: a,
		ctx: ctx,
		giftId: giftId,
	}
}

// Execute executes the request
//  @return GiftSendersResponse
func (a *ReceivedGiftsAPIService) GetReceivedGiftSendersExecute(r ApiGetReceivedGiftSendersRequest) (*GiftSendersResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *GiftSendersResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ReceivedGiftsAPIService.GetReceivedGiftSenders")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/received_gifts/{gift_id}/senders"
	localVarPath = strings.Replace(localVarPath, "{"+"gift_id"+"}", url.PathEscape(parameterValueToString(r.giftId, "giftId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

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

type ApiGetReceivedGiftTransactionRequest struct {
	ctx context.Context
	ApiService *ReceivedGiftsAPIService
	giftTransactionUuid string
}

func (r ApiGetReceivedGiftTransactionRequest) Execute() (*GiftReceivedTransactionResponse, *http.Response, error) {
	return r.ApiService.GetReceivedGiftTransactionExecute(r)
}

func (r ApiGetReceivedGiftTransactionRequest) ExecuteRaw() ([]byte, *http.Response, error) {
	_, httpResp, err := r.Execute()
	return executeRaw(httpResp, err)
}

/*
GetReceivedGiftTransaction Method for GetReceivedGiftTransaction

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param giftTransactionUuid
 @return ApiGetReceivedGiftTransactionRequest
*/
func (a *ReceivedGiftsAPIService) GetReceivedGiftTransaction(ctx context.Context, giftTransactionUuid string) ApiGetReceivedGiftTransactionRequest {
	return ApiGetReceivedGiftTransactionRequest{
		ApiService: a,
		ctx: ctx,
		giftTransactionUuid: giftTransactionUuid,
	}
}

// Execute executes the request
//  @return GiftReceivedTransactionResponse
func (a *ReceivedGiftsAPIService) GetReceivedGiftTransactionExecute(r ApiGetReceivedGiftTransactionRequest) (*GiftReceivedTransactionResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *GiftReceivedTransactionResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ReceivedGiftsAPIService.GetReceivedGiftTransaction")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/received_gifts/{gift_transaction_uuid}"
	localVarPath = strings.Replace(localVarPath, "{"+"gift_transaction_uuid"+"}", url.PathEscape(parameterValueToString(r.giftTransactionUuid, "giftTransactionUuid")), -1)

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

type ApiListReceivedGiftsRequest struct {
	ctx context.Context
	ApiService *ReceivedGiftsAPIService
}

func (r ApiListReceivedGiftsRequest) Execute() (*GiftReceivedResponse, *http.Response, error) {
	return r.ApiService.ListReceivedGiftsExecute(r)
}

func (r ApiListReceivedGiftsRequest) ExecuteRaw() ([]byte, *http.Response, error) {
	_, httpResp, err := r.Execute()
	return executeRaw(httpResp, err)
}

/*
ListReceivedGifts Method for ListReceivedGifts

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiListReceivedGiftsRequest
*/
func (a *ReceivedGiftsAPIService) ListReceivedGifts(ctx context.Context) ApiListReceivedGiftsRequest {
	return ApiListReceivedGiftsRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return GiftReceivedResponse
func (a *ReceivedGiftsAPIService) ListReceivedGiftsExecute(r ApiListReceivedGiftsRequest) (*GiftReceivedResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *GiftReceivedResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ReceivedGiftsAPIService.ListReceivedGifts")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v2/received_gifts"

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

type ApiListReceivedGiftsV1Request struct {
	ctx context.Context
	ApiService *ReceivedGiftsAPIService
	from *string
	number *int32
}

func (r ApiListReceivedGiftsV1Request) From(from string) ApiListReceivedGiftsV1Request {
	r.from = &from
	return r
}

func (r ApiListReceivedGiftsV1Request) Number(number int32) ApiListReceivedGiftsV1Request {
	r.number = &number
	return r
}

func (r ApiListReceivedGiftsV1Request) Execute() (*GiftReceivedResponse, *http.Response, error) {
	return r.ApiService.ListReceivedGiftsV1Execute(r)
}

func (r ApiListReceivedGiftsV1Request) ExecuteRaw() ([]byte, *http.Response, error) {
	_, httpResp, err := r.Execute()
	return executeRaw(httpResp, err)
}

/*
ListReceivedGiftsV1 Method for ListReceivedGiftsV1

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiListReceivedGiftsV1Request
*/
func (a *ReceivedGiftsAPIService) ListReceivedGiftsV1(ctx context.Context) ApiListReceivedGiftsV1Request {
	return ApiListReceivedGiftsV1Request{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return GiftReceivedResponse
func (a *ReceivedGiftsAPIService) ListReceivedGiftsV1Execute(r ApiListReceivedGiftsV1Request) (*GiftReceivedResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *GiftReceivedResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ReceivedGiftsAPIService.ListReceivedGiftsV1")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/received_gifts"

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
