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
 * @interface CodeResponse
 */
export interface CodeResponse {
    /**
     * 
     * @type {string}
     * @memberof CodeResponse
     */
    code?: string | null;
    /**
     * 
     * @type {number}
     * @memberof CodeResponse
     */
    nextEditableDatetime?: number | null;
}

/**
 * Check if a given object implements the CodeResponse interface.
 */
export function instanceOfCodeResponse(value: object): value is CodeResponse {
    return true;
}

export function CodeResponseFromJSON(json: any): CodeResponse {
    return CodeResponseFromJSONTyped(json, false);
}

export function CodeResponseFromJSONTyped(json: any, ignoreDiscriminator: boolean): CodeResponse {
    if (json == null) {
        return json;
    }
    return {
        
        'code': json['code'] == null ? undefined : json['code'],
        'nextEditableDatetime': json['next_editable_datetime'] == null ? undefined : json['next_editable_datetime'],
    };
}

export function CodeResponseToJSON(json: any): CodeResponse {
    return CodeResponseToJSONTyped(json, false);
}

export function CodeResponseToJSONTyped(value?: CodeResponse | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'code': value['code'],
        'next_editable_datetime': value['nextEditableDatetime'],
    };
}

