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
import type { ExpiredEmpl } from './ExpiredEmpl';
import {
    ExpiredEmplFromJSON,
    ExpiredEmplFromJSONTyped,
    ExpiredEmplToJSON,
    ExpiredEmplToJSONTyped,
} from './ExpiredEmpl';

/**
 * 
 * @export
 * @interface ExpiredEmplResponse
 */
export interface ExpiredEmplResponse {
    /**
     * 
     * @type {Array<ExpiredEmpl>}
     * @memberof ExpiredEmplResponse
     */
    expirations?: Array<ExpiredEmpl> | null;
    /**
     * 
     * @type {string}
     * @memberof ExpiredEmplResponse
     */
    totalAmount?: string | null;
}

/**
 * Check if a given object implements the ExpiredEmplResponse interface.
 */
export function instanceOfExpiredEmplResponse(value: object): value is ExpiredEmplResponse {
    return true;
}

export function ExpiredEmplResponseFromJSON(json: any): ExpiredEmplResponse {
    return ExpiredEmplResponseFromJSONTyped(json, false);
}

export function ExpiredEmplResponseFromJSONTyped(json: any, ignoreDiscriminator: boolean): ExpiredEmplResponse {
    if (json == null) {
        return json;
    }
    return {
        
        'expirations': json['expirations'] == null ? undefined : ((json['expirations'] as Array<any>).map(ExpiredEmplFromJSON)),
        'totalAmount': json['total_amount'] == null ? undefined : json['total_amount'],
    };
}

export function ExpiredEmplResponseToJSON(json: any): ExpiredEmplResponse {
    return ExpiredEmplResponseToJSONTyped(json, false);
}

export function ExpiredEmplResponseToJSONTyped(value?: ExpiredEmplResponse | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'expirations': value['expirations'] == null ? undefined : ((value['expirations'] as Array<any>).map(ExpiredEmplToJSON)),
        'total_amount': value['totalAmount'],
    };
}

