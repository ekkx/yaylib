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
 * @interface BanWordType
 */
export interface BanWordType {
    /**
     * 
     * @type {string}
     * @memberof BanWordType
     */
    apiValue?: string | null;
}

/**
 * Check if a given object implements the BanWordType interface.
 */
export function instanceOfBanWordType(value: object): value is BanWordType {
    return true;
}

export function BanWordTypeFromJSON(json: any): BanWordType {
    return BanWordTypeFromJSONTyped(json, false);
}

export function BanWordTypeFromJSONTyped(json: any, ignoreDiscriminator: boolean): BanWordType {
    if (json == null) {
        return json;
    }
    return {
        
        'apiValue': json['api_value'] == null ? undefined : json['api_value'],
    };
}

export function BanWordTypeToJSON(json: any): BanWordType {
    return BanWordTypeToJSONTyped(json, false);
}

export function BanWordTypeToJSONTyped(value?: BanWordType | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'api_value': value['apiValue'],
    };
}

