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
 * @interface PopularWordType
 */
export interface PopularWordType {
    /**
     * 
     * @type {string}
     * @memberof PopularWordType
     */
    apiValue?: string | null;
}

/**
 * Check if a given object implements the PopularWordType interface.
 */
export function instanceOfPopularWordType(value: object): value is PopularWordType {
    return true;
}

export function PopularWordTypeFromJSON(json: any): PopularWordType {
    return PopularWordTypeFromJSONTyped(json, false);
}

export function PopularWordTypeFromJSONTyped(json: any, ignoreDiscriminator: boolean): PopularWordType {
    if (json == null) {
        return json;
    }
    return {
        
        'apiValue': json['api_value'] == null ? undefined : json['api_value'],
    };
}

export function PopularWordTypeToJSON(json: any): PopularWordType {
    return PopularWordTypeToJSONTyped(json, false);
}

export function PopularWordTypeToJSONTyped(value?: PopularWordType | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'api_value': value['apiValue'],
    };
}

