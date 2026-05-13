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
 * @interface UnsubscribedData
 */
export interface UnsubscribedData {
    /**
     * 
     * @type {Array<number>}
     * @memberof UnsubscribedData
     */
    userIds?: Array<number> | null;
}

/**
 * Check if a given object implements the UnsubscribedData interface.
 */
export function instanceOfUnsubscribedData(value: object): value is UnsubscribedData {
    return true;
}

export function UnsubscribedDataFromJSON(json: any): UnsubscribedData {
    return UnsubscribedDataFromJSONTyped(json, false);
}

export function UnsubscribedDataFromJSONTyped(json: any, ignoreDiscriminator: boolean): UnsubscribedData {
    if (json == null) {
        return json;
    }
    return {
        
        'userIds': json['user_ids'] == null ? undefined : json['user_ids'],
    };
}

export function UnsubscribedDataToJSON(json: any): UnsubscribedData {
    return UnsubscribedDataToJSONTyped(json, false);
}

export function UnsubscribedDataToJSONTyped(value?: UnsubscribedData | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'user_ids': value['userIds'],
    };
}

