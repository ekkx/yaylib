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
 * @interface Generation
 */
export interface Generation {
    /**
     * 
     * @type {number}
     * @memberof Generation
     */
    apiValue?: number | null;
}

/**
 * Check if a given object implements the Generation interface.
 */
export function instanceOfGeneration(value: object): value is Generation {
    return true;
}

export function GenerationFromJSON(json: any): Generation {
    return GenerationFromJSONTyped(json, false);
}

export function GenerationFromJSONTyped(json: any, ignoreDiscriminator: boolean): Generation {
    if (json == null) {
        return json;
    }
    return {
        
        'apiValue': json['api_value'] == null ? undefined : json['api_value'],
    };
}

export function GenerationToJSON(json: any): Generation {
    return GenerationToJSONTyped(json, false);
}

export function GenerationToJSONTyped(value?: Generation | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'api_value': value['apiValue'],
    };
}

