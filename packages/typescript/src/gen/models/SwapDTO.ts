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
 * @interface SwapDTO
 */
export interface SwapDTO {
    /**
     * 
     * @type {string}
     * @memberof SwapDTO
     */
    url?: string | null;
}

/**
 * Check if a given object implements the SwapDTO interface.
 */
export function instanceOfSwapDTO(value: object): value is SwapDTO {
    return true;
}

export function SwapDTOFromJSON(json: any): SwapDTO {
    return SwapDTOFromJSONTyped(json, false);
}

export function SwapDTOFromJSONTyped(json: any, ignoreDiscriminator: boolean): SwapDTO {
    if (json == null) {
        return json;
    }
    return {
        
        'url': json['url'] == null ? undefined : json['url'],
    };
}

export function SwapDTOToJSON(json: any): SwapDTO {
    return SwapDTOToJSONTyped(json, false);
}

export function SwapDTOToJSONTyped(value?: SwapDTO | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'url': value['url'],
    };
}

