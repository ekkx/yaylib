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
 * @interface TotalChatRequestData
 */
export interface TotalChatRequestData {
    /**
     * 
     * @type {number}
     * @memberof TotalChatRequestData
     */
    totalCount?: number | null;
}

/**
 * Check if a given object implements the TotalChatRequestData interface.
 */
export function instanceOfTotalChatRequestData(value: object): value is TotalChatRequestData {
    return true;
}

export function TotalChatRequestDataFromJSON(json: any): TotalChatRequestData {
    return TotalChatRequestDataFromJSONTyped(json, false);
}

export function TotalChatRequestDataFromJSONTyped(json: any, ignoreDiscriminator: boolean): TotalChatRequestData {
    if (json == null) {
        return json;
    }
    return {
        
        'totalCount': json['total_count'] == null ? undefined : json['total_count'],
    };
}

export function TotalChatRequestDataToJSON(json: any): TotalChatRequestData {
    return TotalChatRequestDataToJSONTyped(json, false);
}

export function TotalChatRequestDataToJSONTyped(value?: TotalChatRequestData | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'total_count': value['totalCount'],
    };
}

