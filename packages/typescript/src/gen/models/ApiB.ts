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
 * @interface ApiB
 */
export interface ApiB {
    /**
     * 
     * @type {object}
     * @memberof ApiB
     */
    d?: object | null;
    /**
     * 
     * @type {object}
     * @memberof ApiB
     */
    f?: object | null;
}

/**
 * Check if a given object implements the ApiB interface.
 */
export function instanceOfApiB(value: object): value is ApiB {
    return true;
}

export function ApiBFromJSON(json: any): ApiB {
    return ApiBFromJSONTyped(json, false);
}

export function ApiBFromJSONTyped(json: any, ignoreDiscriminator: boolean): ApiB {
    if (json == null) {
        return json;
    }
    return {
        
        'd': json['d'] == null ? undefined : json['d'],
        'f': json['f'] == null ? undefined : json['f'],
    };
}

export function ApiBToJSON(json: any): ApiB {
    return ApiBToJSONTyped(json, false);
}

export function ApiBToJSONTyped(value?: ApiB | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'd': value['d'],
        'f': value['f'],
    };
}

