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

import { mapValues } from '../runtime.js';
import type { RefreshCounterRequest } from './RefreshCounterRequest.js';
import {
    RefreshCounterRequestFromJSON,
    RefreshCounterRequestFromJSONTyped,
    RefreshCounterRequestToJSON,
    RefreshCounterRequestToJSONTyped,
} from './RefreshCounterRequest.js';

/**
 * 
 * @export
 * @interface RefreshCounterRequestsResponse
 */
export interface RefreshCounterRequestsResponse {
    /**
     * 
     * @type {Array<RefreshCounterRequest>}
     * @memberof RefreshCounterRequestsResponse
     */
    resetCounterRequests?: Array<RefreshCounterRequest> | null;
}

/**
 * Check if a given object implements the RefreshCounterRequestsResponse interface.
 */
export function instanceOfRefreshCounterRequestsResponse(value: object): value is RefreshCounterRequestsResponse {
    return true;
}

export function RefreshCounterRequestsResponseFromJSON(json: any): RefreshCounterRequestsResponse {
    return RefreshCounterRequestsResponseFromJSONTyped(json, false);
}

export function RefreshCounterRequestsResponseFromJSONTyped(json: any, ignoreDiscriminator: boolean): RefreshCounterRequestsResponse {
    if (json == null) {
        return json;
    }
    return {
        
        'resetCounterRequests': json['reset_counter_requests'] == null ? undefined : ((json['reset_counter_requests'] as Array<any>).map(RefreshCounterRequestFromJSON)),
    };
}

export function RefreshCounterRequestsResponseToJSON(json: any): RefreshCounterRequestsResponse {
    return RefreshCounterRequestsResponseToJSONTyped(json, false);
}

export function RefreshCounterRequestsResponseToJSONTyped(value?: RefreshCounterRequestsResponse | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'reset_counter_requests': value['resetCounterRequests'] == null ? undefined : ((value['resetCounterRequests'] as Array<any>).map(RefreshCounterRequestToJSON)),
    };
}

