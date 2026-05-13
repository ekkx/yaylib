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
  GamesResponse,
  Walkthrough,
} from '../models/index';
import {
    GamesResponseFromJSON,
    GamesResponseToJSON,
    WalkthroughFromJSON,
    WalkthroughToJSON,
} from '../models/index';

export interface ListGameAppsRequest {
    number?: number;
    fromId?: number | null;
    ids?: Array<number>;
}

export interface ListGameWalkthroughsRequest {
    appId: number;
}

/**
 * 
 */
export class GamesApi extends runtime.BaseAPI {

    /**
     */
    async listGameAppsRaw(requestParameters: ListGameAppsRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<GamesResponse>> {
        const queryParameters: any = {};

        if (requestParameters['number'] != null) {
            queryParameters['number'] = requestParameters['number'];
        }

        if (requestParameters['fromId'] != null) {
            queryParameters['from_id'] = requestParameters['fromId'];
        }

        if (requestParameters['ids'] != null) {
            queryParameters['ids[]'] = requestParameters['ids'];
        }

        const headerParameters: runtime.HTTPHeaders = {};

        const response = await this.request({
            path: `/v1/games/apps`,
            method: 'GET',
            headers: headerParameters,
            query: queryParameters,
        }, initOverrides);

        return new runtime.JSONApiResponse(response, (jsonValue) => GamesResponseFromJSON(jsonValue));
    }

    /**
     */
    async listGameApps(requestParameters: ListGameAppsRequest = {}, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<GamesResponse> {
        const response = await this.listGameAppsRaw(requestParameters, initOverrides);
        return await response.value();
    }

    /**
     */
    async listGameWalkthroughsRaw(requestParameters: ListGameWalkthroughsRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<Array<Walkthrough>>> {
        if (requestParameters['appId'] == null) {
            throw new runtime.RequiredError(
                'appId',
                'Required parameter "appId" was null or undefined when calling listGameWalkthroughs().'
            );
        }

        const queryParameters: any = {};

        const headerParameters: runtime.HTTPHeaders = {};

        const response = await this.request({
            path: `/v1/games/apps/{app_id}/walkthroughs`.replace(`{${"app_id"}}`, encodeURIComponent(String(requestParameters['appId']))),
            method: 'GET',
            headers: headerParameters,
            query: queryParameters,
        }, initOverrides);

        return new runtime.JSONApiResponse(response, (jsonValue) => jsonValue.map(WalkthroughFromJSON));
    }

    /**
     */
    async listGameWalkthroughs(requestParameters: ListGameWalkthroughsRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<Array<Walkthrough>> {
        const response = await this.listGameWalkthroughsRaw(requestParameters, initOverrides);
        return await response.value();
    }

}
