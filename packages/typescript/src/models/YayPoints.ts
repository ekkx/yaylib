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
 * @interface YayPoints
 */
export interface YayPoints {
    /**
     * 
     * @type {number}
     * @memberof YayPoints
     */
    total?: number | null;
}

/**
 * Check if a given object implements the YayPoints interface.
 */
export function instanceOfYayPoints(value: object): value is YayPoints {
    return true;
}

export function YayPointsFromJSON(json: any): YayPoints {
    return YayPointsFromJSONTyped(json, false);
}

export function YayPointsFromJSONTyped(json: any, ignoreDiscriminator: boolean): YayPoints {
    if (json == null) {
        return json;
    }
    return {
        
        'total': json['total'] == null ? undefined : json['total'],
    };
}

export function YayPointsToJSON(json: any): YayPoints {
    return YayPointsToJSONTyped(json, false);
}

export function YayPointsToJSONTyped(value?: YayPoints | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'total': value['total'],
    };
}

