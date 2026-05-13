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
 * @interface DotMoneyUrlResponse
 */
export interface DotMoneyUrlResponse {
    /**
     * 
     * @type {string}
     * @memberof DotMoneyUrlResponse
     */
    url?: string | null;
}

/**
 * Check if a given object implements the DotMoneyUrlResponse interface.
 */
export function instanceOfDotMoneyUrlResponse(value: object): value is DotMoneyUrlResponse {
    return true;
}

export function DotMoneyUrlResponseFromJSON(json: any): DotMoneyUrlResponse {
    return DotMoneyUrlResponseFromJSONTyped(json, false);
}

export function DotMoneyUrlResponseFromJSONTyped(json: any, ignoreDiscriminator: boolean): DotMoneyUrlResponse {
    if (json == null) {
        return json;
    }
    return {
        
        'url': json['url'] == null ? undefined : json['url'],
    };
}

export function DotMoneyUrlResponseToJSON(json: any): DotMoneyUrlResponse {
    return DotMoneyUrlResponseToJSONTyped(json, false);
}

export function DotMoneyUrlResponseToJSONTyped(value?: DotMoneyUrlResponse | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'url': value['url'],
    };
}

