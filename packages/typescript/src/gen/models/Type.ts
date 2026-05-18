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
 * @interface Type
 */
export interface Type {
    /**
     * 
     * @type {string}
     * @memberof Type
     */
    apiValue?: string | null;
}

/**
 * Check if a given object implements the Type interface.
 */
export function instanceOfType(value: object): value is Type {
    return true;
}

export function TypeFromJSON(json: any): Type {
    return TypeFromJSONTyped(json, false);
}

export function TypeFromJSONTyped(json: any, ignoreDiscriminator: boolean): Type {
    if (json == null) {
        return json;
    }
    return {
        
        'apiValue': json['api_value'] == null ? undefined : json['api_value'],
    };
}

export function TypeToJSON(json: any): Type {
    return TypeToJSONTyped(json, false);
}

export function TypeToJSONTyped(value?: Type | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'api_value': value['apiValue'],
    };
}

