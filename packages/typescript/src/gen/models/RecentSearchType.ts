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
 * @interface RecentSearchType
 */
export interface RecentSearchType {
    /**
     * 
     * @type {string}
     * @memberof RecentSearchType
     */
    apiValue?: string | null;
}

/**
 * Check if a given object implements the RecentSearchType interface.
 */
export function instanceOfRecentSearchType(value: object): value is RecentSearchType {
    return true;
}

export function RecentSearchTypeFromJSON(json: any): RecentSearchType {
    return RecentSearchTypeFromJSONTyped(json, false);
}

export function RecentSearchTypeFromJSONTyped(json: any, ignoreDiscriminator: boolean): RecentSearchType {
    if (json == null) {
        return json;
    }
    return {
        
        'apiValue': json['api_value'] == null ? undefined : json['api_value'],
    };
}

export function RecentSearchTypeToJSON(json: any): RecentSearchType {
    return RecentSearchTypeToJSONTyped(json, false);
}

export function RecentSearchTypeToJSONTyped(value?: RecentSearchType | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'api_value': value['apiValue'],
    };
}

