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
 * @interface PalGacha
 */
export interface PalGacha {
    /**
     * 
     * @type {string}
     * @memberof PalGacha
     */
    bannerCollectionUrl?: string | null;
    /**
     * 
     * @type {string}
     * @memberof PalGacha
     */
    bannerUrl?: string | null;
    /**
     * 
     * @type {string}
     * @memberof PalGacha
     */
    detailUrl?: string | null;
    /**
     * 
     * @type {string}
     * @memberof PalGacha
     */
    gachaType?: string | null;
    /**
     * 
     * @type {string}
     * @memberof PalGacha
     */
    id?: string | null;
    /**
     * 
     * @type {number}
     * @memberof PalGacha
     */
    price?: number | null;
}

/**
 * Check if a given object implements the PalGacha interface.
 */
export function instanceOfPalGacha(value: object): value is PalGacha {
    return true;
}

export function PalGachaFromJSON(json: any): PalGacha {
    return PalGachaFromJSONTyped(json, false);
}

export function PalGachaFromJSONTyped(json: any, ignoreDiscriminator: boolean): PalGacha {
    if (json == null) {
        return json;
    }
    return {
        
        'bannerCollectionUrl': json['banner_collection_url'] == null ? undefined : json['banner_collection_url'],
        'bannerUrl': json['banner_url'] == null ? undefined : json['banner_url'],
        'detailUrl': json['detail_url'] == null ? undefined : json['detail_url'],
        'gachaType': json['gacha_type'] == null ? undefined : json['gacha_type'],
        'id': json['id'] == null ? undefined : json['id'],
        'price': json['price'] == null ? undefined : json['price'],
    };
}

export function PalGachaToJSON(json: any): PalGacha {
    return PalGachaToJSONTyped(json, false);
}

export function PalGachaToJSONTyped(value?: PalGacha | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'banner_collection_url': value['bannerCollectionUrl'],
        'banner_url': value['bannerUrl'],
        'detail_url': value['detailUrl'],
        'gacha_type': value['gachaType'],
        'id': value['id'],
        'price': value['price'],
    };
}

