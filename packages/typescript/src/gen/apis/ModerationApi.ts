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
  BanWordsResponse,
  PopularWordsResponse,
} from '../models/index.js';
import {
    BanWordsResponseFromJSON,
    BanWordsResponseToJSON,
    PopularWordsResponseFromJSON,
    PopularWordsResponseToJSON,
} from '../models/index.js';

export interface GetBannedWordsRequest {
    countryApiValue: string;
}

export interface GetPopularWordsRequest {
    countryApiValue: string;
    app: string;
}

/**
 * 
 */
export class ModerationApi extends runtime.BaseAPI {

    /**
     */
    async getBannedWordsRaw(requestParameters: GetBannedWordsRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<BanWordsResponse>> {
        if (requestParameters['countryApiValue'] == null) {
            throw new runtime.RequiredError(
                'countryApiValue',
                'Required parameter "countryApiValue" was null or undefined when calling getBannedWords().'
            );
        }

        const queryParameters: any = {};

        const headerParameters: runtime.HTTPHeaders = {};

        const response = await this.request({
            path: `/{countryApiValue}/api/v2/banned_words`.replace(`{${"countryApiValue"}}`, encodeURIComponent(String(requestParameters['countryApiValue']))),
            method: 'GET',
            headers: headerParameters,
            query: queryParameters,
        }, initOverrides);

        return new runtime.JSONApiResponse(response, (jsonValue) => BanWordsResponseFromJSON(jsonValue));
    }

    /**
     */
    async getBannedWords(requestParameters: GetBannedWordsRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<BanWordsResponse> {
        const response = await this.getBannedWordsRaw(requestParameters, initOverrides);
        return await response.value();
    }

    /**
     */
    async getPopularWordsRaw(requestParameters: GetPopularWordsRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<PopularWordsResponse>> {
        if (requestParameters['countryApiValue'] == null) {
            throw new runtime.RequiredError(
                'countryApiValue',
                'Required parameter "countryApiValue" was null or undefined when calling getPopularWords().'
            );
        }

        if (requestParameters['app'] == null) {
            throw new runtime.RequiredError(
                'app',
                'Required parameter "app" was null or undefined when calling getPopularWords().'
            );
        }

        const queryParameters: any = {};

        const headerParameters: runtime.HTTPHeaders = {};

        const response = await this.request({
            path: `/{countryApiValue}/api/apps/{app}/popular_words`.replace(`{${"countryApiValue"}}`, encodeURIComponent(String(requestParameters['countryApiValue']))).replace(`{${"app"}}`, encodeURIComponent(String(requestParameters['app']))),
            method: 'GET',
            headers: headerParameters,
            query: queryParameters,
        }, initOverrides);

        return new runtime.JSONApiResponse(response, (jsonValue) => PopularWordsResponseFromJSON(jsonValue));
    }

    /**
     */
    async getPopularWords(requestParameters: GetPopularWordsRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<PopularWordsResponse> {
        const response = await this.getPopularWordsRaw(requestParameters, initOverrides);
        return await response.value();
    }

}