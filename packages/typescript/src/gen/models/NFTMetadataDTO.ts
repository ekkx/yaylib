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
import type { NFTAttributeDTO } from './NFTAttributeDTO.js';
import {
    NFTAttributeDTOFromJSON,
    NFTAttributeDTOFromJSONTyped,
    NFTAttributeDTOToJSON,
    NFTAttributeDTOToJSONTyped,
} from './NFTAttributeDTO.js';

/**
 * 
 * @export
 * @interface NFTMetadataDTO
 */
export interface NFTMetadataDTO {
    /**
     * 
     * @type {Array<NFTAttributeDTO>}
     * @memberof NFTMetadataDTO
     */
    attributes?: Array<NFTAttributeDTO> | null;
}

/**
 * Check if a given object implements the NFTMetadataDTO interface.
 */
export function instanceOfNFTMetadataDTO(value: object): value is NFTMetadataDTO {
    return true;
}

export function NFTMetadataDTOFromJSON(json: any): NFTMetadataDTO {
    return NFTMetadataDTOFromJSONTyped(json, false);
}

export function NFTMetadataDTOFromJSONTyped(json: any, ignoreDiscriminator: boolean): NFTMetadataDTO {
    if (json == null) {
        return json;
    }
    return {
        
        'attributes': json['attributes'] == null ? undefined : ((json['attributes'] as Array<any>).map(NFTAttributeDTOFromJSON)),
    };
}

export function NFTMetadataDTOToJSON(json: any): NFTMetadataDTO {
    return NFTMetadataDTOToJSONTyped(json, false);
}

export function NFTMetadataDTOToJSONTyped(value?: NFTMetadataDTO | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'attributes': value['attributes'] == null ? undefined : ((value['attributes'] as Array<any>).map(NFTAttributeDTOToJSON)),
    };
}

