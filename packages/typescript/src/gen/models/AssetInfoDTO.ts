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
 * @interface AssetInfoDTO
 */
export interface AssetInfoDTO {
    /**
     * 
     * @type {string}
     * @memberof AssetInfoDTO
     */
    image?: string | null;
    /**
     * 
     * @type {string}
     * @memberof AssetInfoDTO
     */
    name?: string | null;
    /**
     * 
     * @type {string}
     * @memberof AssetInfoDTO
     */
    tokenId?: string | null;
}

/**
 * Check if a given object implements the AssetInfoDTO interface.
 */
export function instanceOfAssetInfoDTO(value: object): value is AssetInfoDTO {
    return true;
}

export function AssetInfoDTOFromJSON(json: any): AssetInfoDTO {
    return AssetInfoDTOFromJSONTyped(json, false);
}

export function AssetInfoDTOFromJSONTyped(json: any, ignoreDiscriminator: boolean): AssetInfoDTO {
    if (json == null) {
        return json;
    }
    return {
        
        'image': json['image'] == null ? undefined : json['image'],
        'name': json['name'] == null ? undefined : json['name'],
        'tokenId': json['token_id'] == null ? undefined : json['token_id'],
    };
}

export function AssetInfoDTOToJSON(json: any): AssetInfoDTO {
    return AssetInfoDTOToJSONTyped(json, false);
}

export function AssetInfoDTOToJSONTyped(value?: AssetInfoDTO | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'image': value['image'],
        'name': value['name'],
        'token_id': value['tokenId'],
    };
}

