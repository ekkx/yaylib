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
 * @interface ApiN
 */
export interface ApiN {
    /**
     * 
     * @type {string}
     * @memberof ApiN
     */
    d?: string | null;
    /**
     * 
     * @type {string}
     * @memberof ApiN
     */
    f?: string | null;
    /**
     * 
     * @type {number}
     * @memberof ApiN
     */
    h?: number | null;
    /**
     * 
     * @type {number}
     * @memberof ApiN
     */
    i?: number | null;
}

/**
 * Check if a given object implements the ApiN interface.
 */
export function instanceOfApiN(value: object): value is ApiN {
    return true;
}

export function ApiNFromJSON(json: any): ApiN {
    return ApiNFromJSONTyped(json, false);
}

export function ApiNFromJSONTyped(json: any, ignoreDiscriminator: boolean): ApiN {
    if (json == null) {
        return json;
    }
    return {
        
        'd': json['d'] == null ? undefined : json['d'],
        'f': json['f'] == null ? undefined : json['f'],
        'h': json['h'] == null ? undefined : json['h'],
        'i': json['i'] == null ? undefined : json['i'],
    };
}

export function ApiNToJSON(json: any): ApiN {
    return ApiNToJSONTyped(json, false);
}

export function ApiNToJSONTyped(value?: ApiN | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'd': value['d'],
        'f': value['f'],
        'h': value['h'],
        'i': value['i'],
    };
}

