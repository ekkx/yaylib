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
 * @interface Deposit
 */
export interface Deposit {
    /**
     * 
     * @type {AssetType}
     * @memberof Deposit
     */
    assetType?: AssetType | null;
    /**
     * 
     * @type {string}
     * @memberof Deposit
     */
    palImageUrl?: string | null;
    /**
     * 
     * @type {string}
     * @memberof Deposit
     */
    palName?: string | null;
}

/**
 * Check if a given object implements the Deposit interface.
 */
export function instanceOfDeposit(value: object): value is Deposit {
    return true;
}

export function DepositFromJSON(json: any): Deposit {
    return DepositFromJSONTyped(json, false);
}

export function DepositFromJSONTyped(json: any, ignoreDiscriminator: boolean): Deposit {
    if (json == null) {
        return json;
    }
    return {
        
        'assetType': json['asset_type'] == null ? undefined : AssetTypeFromJSON(json['asset_type']),
        'palImageUrl': json['pal_image_url'] == null ? undefined : json['pal_image_url'],
        'palName': json['pal_name'] == null ? undefined : json['pal_name'],
    };
}

export function DepositToJSON(json: any): Deposit {
    return DepositToJSONTyped(json, false);
}

export function DepositToJSONTyped(value?: Deposit | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'asset_type': AssetTypeToJSON(value['assetType']),
        'pal_image_url': value['palImageUrl'],
        'pal_name': value['palName'],
    };
}

