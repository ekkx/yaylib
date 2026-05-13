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


import * as runtime from '../runtime';
import type {
  GiftsResponse,
} from '../models/index';
import {
    GiftsResponseFromJSON,
    GiftsResponseToJSON,
} from '../models/index';

export interface ListGiftsRequest {
    currency?: string | null;
    number?: number | null;
    from?: string | null;
}

/**
 * 
 */
export class GiftsApi extends runtime.BaseAPI {

    /**
     */
    async listGiftsRaw(requestParameters: ListGiftsRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<GiftsResponse>> {
        const queryParameters: any = {};

        if (requestParameters['currency'] != null) {
            queryParameters['currency'] = requestParameters['currency'];
        }

        if (requestParameters['number'] != null) {
            queryParameters['number'] = requestParameters['number'];
        }

        if (requestParameters['from'] != null) {
            queryParameters['from'] = requestParameters['from'];
        }

        const headerParameters: runtime.HTTPHeaders = {};

        const response = await this.request({
            path: `/v2/gifts`,
            method: 'GET',
            headers: headerParameters,
            query: queryParameters,
        }, initOverrides);

        return new runtime.JSONApiResponse(response, (jsonValue) => GiftsResponseFromJSON(jsonValue));
    }

    /**
     */
    async listGifts(requestParameters: ListGiftsRequest = {}, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<GiftsResponse> {
        const response = await this.listGiftsRaw(requestParameters, initOverrides);
        return await response.value();
    }

}