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
 * @interface Link
 */
export interface Link {
    /**
     * 
     * @type {string}
     * @memberof Link
     */
    result?: string | null;
}

/**
 * Check if a given object implements the Link interface.
 */
export function instanceOfLink(value: object): value is Link {
    return true;
}

export function LinkFromJSON(json: any): Link {
    return LinkFromJSONTyped(json, false);
}

export function LinkFromJSONTyped(json: any, ignoreDiscriminator: boolean): Link {
    if (json == null) {
        return json;
    }
    return {
        
        'result': json['result'] == null ? undefined : json['result'],
    };
}

export function LinkToJSON(json: any): Link {
    return LinkToJSONTyped(json, false);
}

export function LinkToJSONTyped(value?: Link | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'result': value['result'],
    };
}

