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
 * @interface RtmTokenResponse
 */
export interface RtmTokenResponse {
    /**
     * 
     * @type {string}
     * @memberof RtmTokenResponse
     */
    token?: string | null;
}

/**
 * Check if a given object implements the RtmTokenResponse interface.
 */
export function instanceOfRtmTokenResponse(value: object): value is RtmTokenResponse {
    return true;
}

export function RtmTokenResponseFromJSON(json: any): RtmTokenResponse {
    return RtmTokenResponseFromJSONTyped(json, false);
}

export function RtmTokenResponseFromJSONTyped(json: any, ignoreDiscriminator: boolean): RtmTokenResponse {
    if (json == null) {
        return json;
    }
    return {
        
        'token': json['token'] == null ? undefined : json['token'],
    };
}

export function RtmTokenResponseToJSON(json: any): RtmTokenResponse {
    return RtmTokenResponseToJSONTyped(json, false);
}

export function RtmTokenResponseToJSONTyped(value?: RtmTokenResponse | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'token': value['token'],
    };
}

