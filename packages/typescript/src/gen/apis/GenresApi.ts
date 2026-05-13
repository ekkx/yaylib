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
  GenresResponse,
} from '../models/index';
import {
    GenresResponseFromJSON,
    GenresResponseToJSON,
} from '../models/index';

export interface ListGenresRequest {
    number?: number;
    from?: string | null;
}

/**
 * 
 */
export class GenresApi extends runtime.BaseAPI {

    /**
     */
    async listGenresRaw(requestParameters: ListGenresRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<GenresResponse>> {
        const queryParameters: any = {};

        if (requestParameters['number'] != null) {
            queryParameters['number'] = requestParameters['number'];
        }

        if (requestParameters['from'] != null) {
            queryParameters['from'] = requestParameters['from'];
        }

        const headerParameters: runtime.HTTPHeaders = {};

        const response = await this.request({
            path: `/v1/genres`,
            method: 'GET',
            headers: headerParameters,
            query: queryParameters,
        }, initOverrides);

        return new runtime.JSONApiResponse(response, (jsonValue) => GenresResponseFromJSON(jsonValue));
    }

    /**
     */
    async listGenres(requestParameters: ListGenresRequest = {}, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<GenresResponse> {
        const response = await this.listGenresRaw(requestParameters, initOverrides);
        return await response.value();
    }

}