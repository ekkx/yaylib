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
/**
 * 
 * @export
 * @interface EmplDetails
 */
export interface EmplDetails {
    /**
     * 
     * @type {number}
     * @memberof EmplDetails
     */
    regular?: number | null;
    /**
     * 
     * @type {number}
     * @memberof EmplDetails
     */
    rewarded?: number | null;
}

/**
 * Check if a given object implements the EmplDetails interface.
 */
export function instanceOfEmplDetails(value: object): value is EmplDetails {
    return true;
}

export function EmplDetailsFromJSON(json: any): EmplDetails {
    return EmplDetailsFromJSONTyped(json, false);
}

export function EmplDetailsFromJSONTyped(json: any, ignoreDiscriminator: boolean): EmplDetails {
    if (json == null) {
        return json;
    }
    return {
        
        'regular': json['regular'] == null ? undefined : json['regular'],
        'rewarded': json['rewarded'] == null ? undefined : json['rewarded'],
    };
}

export function EmplDetailsToJSON(json: any): EmplDetails {
    return EmplDetailsToJSONTyped(json, false);
}

export function EmplDetailsToJSONTyped(value?: EmplDetails | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'regular': value['regular'],
        'rewarded': value['rewarded'],
    };
}

