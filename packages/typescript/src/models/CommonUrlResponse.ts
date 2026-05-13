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
 * @interface CommonUrlResponse
 */
export interface CommonUrlResponse {
    /**
     * 
     * @type {string}
     * @memberof CommonUrlResponse
     */
    url?: string | null;
}

/**
 * Check if a given object implements the CommonUrlResponse interface.
 */
export function instanceOfCommonUrlResponse(value: object): value is CommonUrlResponse {
    return true;
}

export function CommonUrlResponseFromJSON(json: any): CommonUrlResponse {
    return CommonUrlResponseFromJSONTyped(json, false);
}

export function CommonUrlResponseFromJSONTyped(json: any, ignoreDiscriminator: boolean): CommonUrlResponse {
    if (json == null) {
        return json;
    }
    return {
        
        'url': json['url'] == null ? undefined : json['url'],
    };
}

export function CommonUrlResponseToJSON(json: any): CommonUrlResponse {
    return CommonUrlResponseToJSONTyped(json, false);
}

export function CommonUrlResponseToJSONTyped(value?: CommonUrlResponse | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'url': value['url'],
    };
}

