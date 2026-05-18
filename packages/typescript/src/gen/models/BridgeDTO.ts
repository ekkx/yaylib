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
 * @interface BridgeDTO
 */
export interface BridgeDTO {
    /**
     * 
     * @type {string}
     * @memberof BridgeDTO
     */
    url?: string | null;
}

/**
 * Check if a given object implements the BridgeDTO interface.
 */
export function instanceOfBridgeDTO(value: object): value is BridgeDTO {
    return true;
}

export function BridgeDTOFromJSON(json: any): BridgeDTO {
    return BridgeDTOFromJSONTyped(json, false);
}

export function BridgeDTOFromJSONTyped(json: any, ignoreDiscriminator: boolean): BridgeDTO {
    if (json == null) {
        return json;
    }
    return {
        
        'url': json['url'] == null ? undefined : json['url'],
    };
}

export function BridgeDTOToJSON(json: any): BridgeDTO {
    return BridgeDTOToJSONTyped(json, false);
}

export function BridgeDTOToJSONTyped(value?: BridgeDTO | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'url': value['url'],
    };
}

