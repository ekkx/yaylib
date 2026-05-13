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
 * @interface VerifyDeviceResponse
 */
export interface VerifyDeviceResponse {
    /**
     * 
     * @type {boolean}
     * @memberof VerifyDeviceResponse
     */
    verified?: boolean | null;
    /**
     * 
     * @type {string}
     * @memberof VerifyDeviceResponse
     */
    verifiedAt?: string | null;
}

/**
 * Check if a given object implements the VerifyDeviceResponse interface.
 */
export function instanceOfVerifyDeviceResponse(value: object): value is VerifyDeviceResponse {
    return true;
}

export function VerifyDeviceResponseFromJSON(json: any): VerifyDeviceResponse {
    return VerifyDeviceResponseFromJSONTyped(json, false);
}

export function VerifyDeviceResponseFromJSONTyped(json: any, ignoreDiscriminator: boolean): VerifyDeviceResponse {
    if (json == null) {
        return json;
    }
    return {
        
        'verified': json['verified'] == null ? undefined : json['verified'],
        'verifiedAt': json['verified_at'] == null ? undefined : json['verified_at'],
    };
}

export function VerifyDeviceResponseToJSON(json: any): VerifyDeviceResponse {
    return VerifyDeviceResponseToJSONTyped(json, false);
}

export function VerifyDeviceResponseToJSONTyped(value?: VerifyDeviceResponse | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'verified': value['verified'],
        'verified_at': value['verifiedAt'],
    };
}

