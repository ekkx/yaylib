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
 * @interface Target
 */
export interface Target {
    /**
     * 
     * @type {boolean}
     * @memberof Target
     */
    followedBy?: boolean | null;
}

/**
 * Check if a given object implements the Target interface.
 */
export function instanceOfTarget(value: object): value is Target {
    return true;
}

export function TargetFromJSON(json: any): Target {
    return TargetFromJSONTyped(json, false);
}

export function TargetFromJSONTyped(json: any, ignoreDiscriminator: boolean): Target {
    if (json == null) {
        return json;
    }
    return {
        
        'followedBy': json['followed_by'] == null ? undefined : json['followed_by'],
    };
}

export function TargetToJSON(json: any): Target {
    return TargetToJSONTyped(json, false);
}

export function TargetToJSONTyped(value?: Target | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'followed_by': value['followedBy'],
    };
}

