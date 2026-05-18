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
import type { AssetType } from './AssetType.js';
import {
    AssetTypeFromJSON,
    AssetTypeFromJSONTyped,
    AssetTypeToJSON,
    AssetTypeToJSONTyped,
} from './AssetType.js';

/**
 * 
 * @export
 * @interface Withdraw
 */
export interface Withdraw {
    /**
     * 
     * @type {AssetType}
     * @memberof Withdraw
     */
    assetType?: AssetType | null;
    /**
     * 
     * @type {string}
     * @memberof Withdraw
     */
    palImageUrl?: string | null;
    /**
     * 
     * @type {string}
     * @memberof Withdraw
     */
    palName?: string | null;
}

/**
 * Check if a given object implements the Withdraw interface.
 */
export function instanceOfWithdraw(value: object): value is Withdraw {
    return true;
}

export function WithdrawFromJSON(json: any): Withdraw {
    return WithdrawFromJSONTyped(json, false);
}

export function WithdrawFromJSONTyped(json: any, ignoreDiscriminator: boolean): Withdraw {
    if (json == null) {
        return json;
    }
    return {
        
        'assetType': json['asset_type'] == null ? undefined : AssetTypeFromJSON(json['asset_type']),
        'palImageUrl': json['pal_image_url'] == null ? undefined : json['pal_image_url'],
        'palName': json['pal_name'] == null ? undefined : json['pal_name'],
    };
}

export function WithdrawToJSON(json: any): Withdraw {
    return WithdrawToJSONTyped(json, false);
}

export function WithdrawToJSONTyped(value?: Withdraw | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'asset_type': AssetTypeToJSON(value['assetType']),
        'pal_image_url': value['palImageUrl'],
        'pal_name': value['palName'],
    };
}

