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
  ApplicationConfigResponse,
} from '../models/index';
import {
    ApplicationConfigResponseFromJSON,
    ApplicationConfigResponseToJSON,
} from '../models/index';

export interface GetAppConfigRequest {
    app: string;
}

/**
 * 
 */
export class AppsApi extends runtime.BaseAPI {

    /**
     */
    async getAppConfigRaw(requestParameters: GetAppConfigRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<ApplicationConfigResponse>> {
        if (requestParameters['app'] == null) {
            throw new runtime.RequiredError(
                'app',
                'Required parameter "app" was null or undefined when calling getAppConfig().'
            );
        }

        const queryParameters: any = {};

        const headerParameters: runtime.HTTPHeaders = {};

        const response = await this.request({
            path: `/api/apps/{app}`.replace(`{${"app"}}`, encodeURIComponent(String(requestParameters['app']))),
            method: 'GET',
            headers: headerParameters,
            query: queryParameters,
        }, initOverrides);

        return new runtime.JSONApiResponse(response, (jsonValue) => ApplicationConfigResponseFromJSON(jsonValue));
    }

    /**
     */
    async getAppConfig(requestParameters: GetAppConfigRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<ApplicationConfigResponse> {
        const response = await this.getAppConfigRaw(requestParameters, initOverrides);
        return await response.value();
    }

}