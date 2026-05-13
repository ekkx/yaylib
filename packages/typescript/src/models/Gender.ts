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
 * @interface Gender
 */
export interface Gender {
    /**
     * 
     * @type {number}
     * @memberof Gender
     */
    apiIntValue?: number | null;
    /**
     * 
     * @type {string}
     * @memberof Gender
     */
    apiStringValue?: string | null;
}

/**
 * Check if a given object implements the Gender interface.
 */
export function instanceOfGender(value: object): value is Gender {
    return true;
}

export function GenderFromJSON(json: any): Gender {
    return GenderFromJSONTyped(json, false);
}

export function GenderFromJSONTyped(json: any, ignoreDiscriminator: boolean): Gender {
    if (json == null) {
        return json;
    }
    return {
        
        'apiIntValue': json['api_int_value'] == null ? undefined : json['api_int_value'],
        'apiStringValue': json['api_string_value'] == null ? undefined : json['api_string_value'],
    };
}

export function GenderToJSON(json: any): Gender {
    return GenderToJSONTyped(json, false);
}

export function GenderToJSONTyped(value?: Gender | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'api_int_value': value['apiIntValue'],
        'api_string_value': value['apiStringValue'],
    };
}

