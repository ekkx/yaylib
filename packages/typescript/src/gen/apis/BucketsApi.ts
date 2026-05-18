// @ts-nocheck
/* tslint:disable */
/* eslint-disable */
/**
 * Yay! API
 * No description provided
 *
 * Schema version: 4.26.1
 * 
 *
 * NOTE: Code generated; DO NOT EDIT.
 * 
 * Do not edit the class manually.
 */


import * as runtime from '../runtime.js';
import type {
  PresignedUrlsResponse,
} from '../models/index.js';
import {
    PresignedUrlsResponseFromJSON,
    PresignedUrlsResponseToJSON,
} from '../models/index.js';

export interface GetBucketPresignedUrlsRequest {
    fileNames?: Array<string>;
}

/**
 * 
 */
export class BucketsApi extends runtime.BaseAPI {

    /**
     */
    async getBucketPresignedUrlsRaw(requestParameters: GetBucketPresignedUrlsRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<PresignedUrlsResponse>> {
        const queryParameters: any = {};

        if (requestParameters['fileNames'] != null) {
            queryParameters['file_names[]'] = requestParameters['fileNames'];
        }

        const headerParameters: runtime.HTTPHeaders = {};

        const response = await this.request({
            path: `/v1/buckets/presigned_urls`,
            method: 'GET',
            headers: headerParameters,
            query: queryParameters,
        }, initOverrides);

        return new runtime.JSONApiResponse(response, (jsonValue) => PresignedUrlsResponseFromJSON(jsonValue));
    }

    /**
     */
    async getBucketPresignedUrls(requestParameters: GetBucketPresignedUrlsRequest = {}, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<PresignedUrlsResponse> {
        const response = await this.getBucketPresignedUrlsRaw(requestParameters, initOverrides);
        return await response.value();
    }

}