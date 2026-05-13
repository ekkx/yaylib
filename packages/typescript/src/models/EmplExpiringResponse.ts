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
 * @interface EmplExpiringResponse
 */
export interface EmplExpiringResponse {
    /**
     * 
     * @type {string}
     * @memberof EmplExpiringResponse
     */
    expiringAmount?: string | null;
    /**
     * 
     * @type {number}
     * @memberof EmplExpiringResponse
     */
    expiringDate?: number | null;
}

/**
 * Check if a given object implements the EmplExpiringResponse interface.
 */
export function instanceOfEmplExpiringResponse(value: object): value is EmplExpiringResponse {
    return true;
}

export function EmplExpiringResponseFromJSON(json: any): EmplExpiringResponse {
    return EmplExpiringResponseFromJSONTyped(json, false);
}

export function EmplExpiringResponseFromJSONTyped(json: any, ignoreDiscriminator: boolean): EmplExpiringResponse {
    if (json == null) {
        return json;
    }
    return {
        
        'expiringAmount': json['expiring_amount'] == null ? undefined : json['expiring_amount'],
        'expiringDate': json['expiring_date'] == null ? undefined : json['expiring_date'],
    };
}

export function EmplExpiringResponseToJSON(json: any): EmplExpiringResponse {
    return EmplExpiringResponseToJSONTyped(json, false);
}

export function EmplExpiringResponseToJSONTyped(value?: EmplExpiringResponse | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'expiring_amount': value['expiringAmount'],
        'expiring_date': value['expiringDate'],
    };
}

