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
  ActivitiesResponse,
} from '../models/index.js';
import {
    ActivitiesResponseFromJSON,
    ActivitiesResponseToJSON,
} from '../models/index.js';

export interface GetUserActivitiesRequest {
    fromTimestamp?: number | null;
}

export interface GetUserActivitiesV1Request {
    important?: boolean;
    fromTimestamp?: number | null;
    number?: number;
}

/**
 * 
 */
export class ActivitiesApi extends runtime.BaseAPI {

    /**
     */
    async getUserActivitiesRaw(requestParameters: GetUserActivitiesRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<ActivitiesResponse>> {
        const queryParameters: any = {};

        if (requestParameters['fromTimestamp'] != null) {
            queryParameters['from_timestamp'] = requestParameters['fromTimestamp'];
        }

        const headerParameters: runtime.HTTPHeaders = {};

        const response = await this.request({
            path: `/api/v2/user_activities`,
            method: 'GET',
            headers: headerParameters,
            query: queryParameters,
        }, initOverrides);

        return new runtime.JSONApiResponse(response, (jsonValue) => ActivitiesResponseFromJSON(jsonValue));
    }

    /**
     */
    async getUserActivities(requestParameters: GetUserActivitiesRequest = {}, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<ActivitiesResponse> {
        const response = await this.getUserActivitiesRaw(requestParameters, initOverrides);
        return await response.value();
    }

    /**
     */
    async getUserActivitiesV1Raw(requestParameters: GetUserActivitiesV1Request, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<ActivitiesResponse>> {
        const queryParameters: any = {};

        if (requestParameters['important'] != null) {
            queryParameters['important'] = requestParameters['important'];
        }

        if (requestParameters['fromTimestamp'] != null) {
            queryParameters['from_timestamp'] = requestParameters['fromTimestamp'];
        }

        if (requestParameters['number'] != null) {
            queryParameters['number'] = requestParameters['number'];
        }

        const headerParameters: runtime.HTTPHeaders = {};

        const response = await this.request({
            path: `/api/user_activities`,
            method: 'GET',
            headers: headerParameters,
            query: queryParameters,
        }, initOverrides);

        return new runtime.JSONApiResponse(response, (jsonValue) => ActivitiesResponseFromJSON(jsonValue));
    }

    /**
     */
    async getUserActivitiesV1(requestParameters: GetUserActivitiesV1Request = {}, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<ActivitiesResponse> {
        const response = await this.getUserActivitiesV1Raw(requestParameters, initOverrides);
        return await response.value();
    }

}