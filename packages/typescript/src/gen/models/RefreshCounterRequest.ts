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

import { mapValues } from '../runtime';
/**
 * 
 * @export
 * @interface RefreshCounterRequest
 */
export interface RefreshCounterRequest {
    /**
     * 
     * @type {string}
     * @memberof RefreshCounterRequest
     */
    counter?: string | null;
    /**
     * 
     * @type {number}
     * @memberof RefreshCounterRequest
     */
    lastRequestedAt?: number | null;
    /**
     * 
     * @type {string}
     * @memberof RefreshCounterRequest
     */
    status?: string | null;
}

/**
 * Check if a given object implements the RefreshCounterRequest interface.
 */
export function instanceOfRefreshCounterRequest(value: object): value is RefreshCounterRequest {
    return true;
}

export function RefreshCounterRequestFromJSON(json: any): RefreshCounterRequest {
    return RefreshCounterRequestFromJSONTyped(json, false);
}

export function RefreshCounterRequestFromJSONTyped(json: any, ignoreDiscriminator: boolean): RefreshCounterRequest {
    if (json == null) {
        return json;
    }
    return {
        
        'counter': json['counter'] == null ? undefined : json['counter'],
        'lastRequestedAt': json['last_requested_at'] == null ? undefined : json['last_requested_at'],
        'status': json['status'] == null ? undefined : json['status'],
    };
}

export function RefreshCounterRequestToJSON(json: any): RefreshCounterRequest {
    return RefreshCounterRequestToJSONTyped(json, false);
}

export function RefreshCounterRequestToJSONTyped(value?: RefreshCounterRequest | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'counter': value['counter'],
        'last_requested_at': value['lastRequestedAt'],
        'status': value['status'],
    };
}

