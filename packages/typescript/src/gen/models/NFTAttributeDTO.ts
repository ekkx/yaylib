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
 * @interface NFTAttributeDTO
 */
export interface NFTAttributeDTO {
    /**
     * 
     * @type {string}
     * @memberof NFTAttributeDTO
     */
    traitType?: string | null;
    /**
     * 
     * @type {string}
     * @memberof NFTAttributeDTO
     */
    value?: string | null;
}

/**
 * Check if a given object implements the NFTAttributeDTO interface.
 */
export function instanceOfNFTAttributeDTO(value: object): value is NFTAttributeDTO {
    return true;
}

export function NFTAttributeDTOFromJSON(json: any): NFTAttributeDTO {
    return NFTAttributeDTOFromJSONTyped(json, false);
}

export function NFTAttributeDTOFromJSONTyped(json: any, ignoreDiscriminator: boolean): NFTAttributeDTO {
    if (json == null) {
        return json;
    }
    return {
        
        'traitType': json['trait_type'] == null ? undefined : json['trait_type'],
        'value': json['value'] == null ? undefined : json['value'],
    };
}

export function NFTAttributeDTOToJSON(json: any): NFTAttributeDTO {
    return NFTAttributeDTOToJSONTyped(json, false);
}

export function NFTAttributeDTOToJSONTyped(value?: NFTAttributeDTO | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'trait_type': value['traitType'],
        'value': value['value'],
    };
}

