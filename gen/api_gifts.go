
// Code generated from the Yay! API schema; DO NOT EDIT.

package gen

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
)


// GiftsAPIService GiftsAPI service
type GiftsAPIService service

type ApiListGiftsRequest struct {
	ctx context.Context
	ApiService *GiftsAPIService
	currency *string
	number *int32
	from *string
}

func (r ApiListGiftsRequest) Currency(currency string) ApiListGiftsRequest {
	r.currency = &currency
	return r
}

func (r ApiListGiftsRequest) Number(number int32) ApiListGiftsRequest {
	r.number = &number
	return r
}

func (r ApiListGiftsRequest) From(from string) ApiListGiftsRequest {
	r.from = &from
	return r
}

func (r ApiListGiftsRequest) Execute() (*GiftsResponse, *http.Response, error) {
	return r.ApiService.ListGiftsExecute(r)
}

func (r ApiListGiftsRequest) ExecuteRaw() ([]byte, *http.Response, error) {
	_, httpResp, err := r.Execute()
	return executeRaw(httpResp, err)
}

/*
ListGifts Method for ListGifts

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiListGiftsRequest
*/
func (a *GiftsAPIService) ListGifts(ctx context.Context) ApiListGiftsRequest {
	return ApiListGiftsRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return GiftsResponse
func (a *GiftsAPIService) ListGiftsExecute(r ApiListGiftsRequest) (*GiftsResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *GiftsResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "GiftsAPIService.ListGifts")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v2/gifts"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.currency != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "currency", r.currency, "form", "")
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
