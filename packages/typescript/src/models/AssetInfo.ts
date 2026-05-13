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
 * @interface AssetInfo
 */
export interface AssetInfo {
    /**
     * 
     * @type {string}
     * @memberof AssetInfo
     */
    image?: string | null;
    /**
     * 
     * @type {string}
     * @memberof AssetInfo
     */
    name?: string | null;
    /**
     * 
     * @type {string}
     * @memberof AssetInfo
     */
    tokenId?: string | null;
}

/**
 * Check if a given object implements the AssetInfo interface.
 */
export function instanceOfAssetInfo(value: object): value is AssetInfo {
    return true;
}

export function AssetInfoFromJSON(json: any): AssetInfo {
    return AssetInfoFromJSONTyped(json, false);
}

export function AssetInfoFromJSONTyped(json: any, ignoreDiscriminator: boolean): AssetInfo {
    if (json == null) {
        return json;
    }
    return {
        
        'image': json['image'] == null ? undefined : json['image'],
        'name': json['name'] == null ? undefined : json['name'],
        'tokenId': json['token_id'] == null ? undefined : json['token_id'],
    };
}

export function AssetInfoToJSON(json: any): AssetInfo {
    return AssetInfoToJSONTyped(json, false);
}

export function AssetInfoToJSONTyped(value?: AssetInfo | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'image': value['image'],
        'name': value['name'],
        'token_id': value['tokenId'],
    };
}

