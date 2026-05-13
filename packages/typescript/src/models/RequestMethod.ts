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
 * @interface RequestMethod
 */
export interface RequestMethod {
    /**
     * 
     * @type {string}
     * @memberof RequestMethod
     */
    apiValue?: string | null;
}

/**
 * Check if a given object implements the RequestMethod interface.
 */
export function instanceOfRequestMethod(value: object): value is RequestMethod {
    return true;
}

export function RequestMethodFromJSON(json: any): RequestMethod {
    return RequestMethodFromJSONTyped(json, false);
}

export function RequestMethodFromJSONTyped(json: any, ignoreDiscriminator: boolean): RequestMethod {
    if (json == null) {
        return json;
    }
    return {
        
        'apiValue': json['api_value'] == null ? undefined : json['api_value'],
    };
}

export function RequestMethodToJSON(json: any): RequestMethod {
    return RequestMethodToJSONTyped(json, false);
}

export function RequestMethodToJSONTyped(value?: RequestMethod | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'api_value': value['apiValue'],
    };
}

