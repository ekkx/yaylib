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
 * @interface AssetType
 */
export interface AssetType {
    /**
     * 
     * @type {string}
     * @memberof AssetType
     */
    apiValue?: string | null;
}

/**
 * Check if a given object implements the AssetType interface.
 */
export function instanceOfAssetType(value: object): value is AssetType {
    return true;
}

export function AssetTypeFromJSON(json: any): AssetType {
    return AssetTypeFromJSONTyped(json, false);
}

export function AssetTypeFromJSONTyped(json: any, ignoreDiscriminator: boolean): AssetType {
    if (json == null) {
        return json;
    }
    return {
        
        'apiValue': json['api_value'] == null ? undefined : json['api_value'],
    };
}

export function AssetTypeToJSON(json: any): AssetType {
    return AssetTypeToJSONTyped(json, false);
}

export function AssetTypeToJSONTyped(value?: AssetType | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'api_value': value['apiValue'],
    };
}

