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
 * @interface HatchGacha
 */
export interface HatchGacha {
    /**
     * 
     * @type {number}
     * @memberof HatchGacha
     */
    emplCost?: number | null;
    /**
     * 
     * @type {object}
     * @memberof HatchGacha
     */
    gachaType?: object | null;
    /**
     * 
     * @type {string}
     * @memberof HatchGacha
     */
    palImageUrl?: string | null;
    /**
     * 
     * @type {string}
     * @memberof HatchGacha
     */
    palName?: string | null;
    /**
     * 
     * @type {object}
     * @memberof HatchGacha
     */
    palRank?: object | null;
    /**
     * 
     * @type {string}
     * @memberof HatchGacha
     */
    tokenId?: string | null;
}

/**
 * Check if a given object implements the HatchGacha interface.
 */
export function instanceOfHatchGacha(value: object): value is HatchGacha {
    return true;
}

export function HatchGachaFromJSON(json: any): HatchGacha {
    return HatchGachaFromJSONTyped(json, false);
}

export function HatchGachaFromJSONTyped(json: any, ignoreDiscriminator: boolean): HatchGacha {
    if (json == null) {
        return json;
    }
    return {
        
        'emplCost': json['empl_cost'] == null ? undefined : json['empl_cost'],
        'gachaType': json['gacha_type'] == null ? undefined : json['gacha_type'],
        'palImageUrl': json['pal_image_url'] == null ? undefined : json['pal_image_url'],
        'palName': json['pal_name'] == null ? undefined : json['pal_name'],
        'palRank': json['pal_rank'] == null ? undefined : json['pal_rank'],
        'tokenId': json['token_id'] == null ? undefined : json['token_id'],
    };
}

export function HatchGachaToJSON(json: any): HatchGacha {
    return HatchGachaToJSONTyped(json, false);
}

export function HatchGachaToJSONTyped(value?: HatchGacha | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'empl_cost': value['emplCost'],
        'gacha_type': value['gachaType'],
        'pal_image_url': value['palImageUrl'],
        'pal_name': value['palName'],
        'pal_rank': value['palRank'],
        'token_id': value['tokenId'],
    };
}

