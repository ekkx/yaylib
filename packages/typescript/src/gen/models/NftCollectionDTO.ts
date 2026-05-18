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
 * @interface NftCollectionDTO
 */
export interface NftCollectionDTO {
    /**
     * 
     * @type {string}
     * @memberof NftCollectionDTO
     */
    contractAddress?: string | null;
    /**
     * 
     * @type {string}
     * @memberof NftCollectionDTO
     */
    image?: string | null;
    /**
     * 
     * @type {string}
     * @memberof NftCollectionDTO
     */
    mediaType?: string | null;
    /**
     * 
     * @type {string}
     * @memberof NftCollectionDTO
     */
    name?: string | null;
    /**
     * 
     * @type {number}
     * @memberof NftCollectionDTO
     */
    priority?: number | null;
    /**
     * 
     * @type {string}
     * @memberof NftCollectionDTO
     */
    status?: string | null;
}

/**
 * Check if a given object implements the NftCollectionDTO interface.
 */
export function instanceOfNftCollectionDTO(value: object): value is NftCollectionDTO {
    return true;
}

export function NftCollectionDTOFromJSON(json: any): NftCollectionDTO {
    return NftCollectionDTOFromJSONTyped(json, false);
}

export function NftCollectionDTOFromJSONTyped(json: any, ignoreDiscriminator: boolean): NftCollectionDTO {
    if (json == null) {
        return json;
    }
    return {
        
        'contractAddress': json['contract_address'] == null ? undefined : json['contract_address'],
        'image': json['image'] == null ? undefined : json['image'],
        'mediaType': json['media_type'] == null ? undefined : json['media_type'],
        'name': json['name'] == null ? undefined : json['name'],
        'priority': json['priority'] == null ? undefined : json['priority'],
        'status': json['status'] == null ? undefined : json['status'],
    };
}

export function NftCollectionDTOToJSON(json: any): NftCollectionDTO {
    return NftCollectionDTOToJSONTyped(json, false);
}

export function NftCollectionDTOToJSONTyped(value?: NftCollectionDTO | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'contract_address': value['contractAddress'],
        'image': value['image'],
        'media_type': value['mediaType'],
        'name': value['name'],
        'priority': value['priority'],
        'status': value['status'],
    };
}

