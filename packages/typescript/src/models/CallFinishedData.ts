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
 * @interface CallFinishedData
 */
export interface CallFinishedData {
    /**
     * 
     * @type {number}
     * @memberof CallFinishedData
     */
    id?: number | null;
}

/**
 * Check if a given object implements the CallFinishedData interface.
 */
export function instanceOfCallFinishedData(value: object): value is CallFinishedData {
    return true;
}

export function CallFinishedDataFromJSON(json: any): CallFinishedData {
    return CallFinishedDataFromJSONTyped(json, false);
}

export function CallFinishedDataFromJSONTyped(json: any, ignoreDiscriminator: boolean): CallFinishedData {
    if (json == null) {
        return json;
    }
    return {
        
        'id': json['id'] == null ? undefined : json['id'],
    };
}

export function CallFinishedDataToJSON(json: any): CallFinishedData {
    return CallFinishedDataToJSONTyped(json, false);
}

export function CallFinishedDataToJSONTyped(value?: CallFinishedData | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'id': value['id'],
    };
}

