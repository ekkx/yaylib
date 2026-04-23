
// Code generated from the Yay! API schema; DO NOT EDIT.

package gen

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
	"reflect"
)


// BucketsAPIService BucketsAPI service
type BucketsAPIService service

type ApiGetBucketPresignedUrlsRequest struct {
	ctx context.Context
	ApiService *BucketsAPIService
	fileNames *[]string
}

func (r ApiGetBucketPresignedUrlsRequest) FileNames(fileNames []string) ApiGetBucketPresignedUrlsRequest {
	r.fileNames = &fileNames
	return r
}

func (r ApiGetBucketPresignedUrlsRequest) Execute() (*PresignedUrlsResponse, *http.Response, error) {
	return r.ApiService.GetBucketPresignedUrlsExecute(r)
}

/*
GetBucketPresignedUrls Method for GetBucketPresignedUrls

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiGetBucketPresignedUrlsRequest
*/
func (a *BucketsAPIService) GetBucketPresignedUrls(ctx context.Context) ApiGetBucketPresignedUrlsRequest {
	return ApiGetBucketPresignedUrlsRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return PresignedUrlsResponse
func (a *BucketsAPIService) GetBucketPresignedUrlsExecute(r ApiGetBucketPresignedUrlsRequest) (*PresignedUrlsResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *PresignedUrlsResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "BucketsAPIService.GetBucketPresignedUrls")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/buckets/presigned_urls"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.fileNames == nil {
		return localVarReturnValue, nil, reportError("fileNames is required and must be specified")
	}

	{
		t := *r.fileNames
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				parameterAddToHeaderOrQuery(localVarQueryParams, "file_names[]", s.Index(i).Interface(), "form", "multi")
			}
		} else {
			parameterAddToHeaderOrQuery(localVarQueryParams, "file_names[]", t, "form", "multi")
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
