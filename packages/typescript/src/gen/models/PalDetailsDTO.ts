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
 * @interface PalDetailsDTO
 */
export interface PalDetailsDTO {
    /**
     * 
     * @type {number}
     * @memberof PalDetailsDTO
     */
    cost?: number | null;
    /**
     * 
     * @type {number}
     * @memberof PalDetailsDTO
     */
    level?: number | null;
    /**
     * 
     * @type {string}
     * @memberof PalDetailsDTO
     */
    nftType?: string | null;
    /**
     * 
     * @type {string}
     * @memberof PalDetailsDTO
     */
    palImage?: string | null;
    /**
     * 
     * @type {string}
     * @memberof PalDetailsDTO
     */
    palName?: string | null;
    /**
     * 
     * @type {string}
     * @memberof PalDetailsDTO
     */
    rank?: string | null;
    /**
     * 
     * @type {string}
     * @memberof PalDetailsDTO
     */
    tokenId?: string | null;
}

/**
 * Check if a given object implements the PalDetailsDTO interface.
 */
export function instanceOfPalDetailsDTO(value: object): value is PalDetailsDTO {
    return true;
}

export function PalDetailsDTOFromJSON(json: any): PalDetailsDTO {
    return PalDetailsDTOFromJSONTyped(json, false);
}

export function PalDetailsDTOFromJSONTyped(json: any, ignoreDiscriminator: boolean): PalDetailsDTO {
    if (json == null) {
        return json;
    }
    return {
        
        'cost': json['cost'] == null ? undefined : json['cost'],
        'level': json['level'] == null ? undefined : json['level'],
        'nftType': json['nft_type'] == null ? undefined : json['nft_type'],
        'palImage': json['pal_image'] == null ? undefined : json['pal_image'],
        'palName': json['pal_name'] == null ? undefined : json['pal_name'],
        'rank': json['rank'] == null ? undefined : json['rank'],
        'tokenId': json['token_id'] == null ? undefined : json['token_id'],
    };
}

export function PalDetailsDTOToJSON(json: any): PalDetailsDTO {
    return PalDetailsDTOToJSONTyped(json, false);
}

export function PalDetailsDTOToJSONTyped(value?: PalDetailsDTO | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'cost': value['cost'],
        'level': value['level'],
        'nft_type': value['nftType'],
        'pal_image': value['palImage'],
        'pal_name': value['palName'],
        'rank': value['rank'],
        'token_id': value['tokenId'],
    };
}

