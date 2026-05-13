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
 * @interface CommonIdsRequest
 */
export interface CommonIdsRequest {
    /**
     * 
     * @type {Array<number>}
     * @memberof CommonIdsRequest
     */
    ids?: Array<number> | null;
}

/**
 * Check if a given object implements the CommonIdsRequest interface.
 */
export function instanceOfCommonIdsRequest(value: object): value is CommonIdsRequest {
    return true;
}

export function CommonIdsRequestFromJSON(json: any): CommonIdsRequest {
    return CommonIdsRequestFromJSONTyped(json, false);
}

export function CommonIdsRequestFromJSONTyped(json: any, ignoreDiscriminator: boolean): CommonIdsRequest {
    if (json == null) {
        return json;
    }
    return {
        
        'ids': json['ids'] == null ? undefined : json['ids'],
    };
}

export function CommonIdsRequestToJSON(json: any): CommonIdsRequest {
    return CommonIdsRequestToJSONTyped(json, false);
}

export function CommonIdsRequestToJSONTyped(value?: CommonIdsRequest | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'ids': value['ids'],
    };
}

