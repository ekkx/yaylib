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
 * @interface Title
 */
export interface Title {
    /**
     * 
     * @type {string}
     * @memberof Title
     */
    apiValue?: string | null;
}

/**
 * Check if a given object implements the Title interface.
 */
export function instanceOfTitle(value: object): value is Title {
    return true;
}

export function TitleFromJSON(json: any): Title {
    return TitleFromJSONTyped(json, false);
}

export function TitleFromJSONTyped(json: any, ignoreDiscriminator: boolean): Title {
    if (json == null) {
        return json;
    }
    return {
        
        'apiValue': json['api_value'] == null ? undefined : json['api_value'],
    };
}

export function TitleToJSON(json: any): Title {
    return TitleToJSONTyped(json, false);
}

export function TitleToJSONTyped(value?: Title | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'api_value': value['apiValue'],
    };
}

