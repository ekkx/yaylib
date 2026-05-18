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
 * @interface PalGachaHistoryResponseItem
 */
export interface PalGachaHistoryResponseItem {
    /**
     * 
     * @type {number}
     * @memberof PalGachaHistoryResponseItem
     */
    createdAt?: number | null;
    /**
     * 
     * @type {string}
     * @memberof PalGachaHistoryResponseItem
     */
    gachaType?: string | null;
    /**
     * 
     * @type {string}
     * @memberof PalGachaHistoryResponseItem
     */
    id?: string | null;
    /**
     * 
     * @type {string}
     * @memberof PalGachaHistoryResponseItem
     */
    palImageUrl?: string | null;
    /**
     * 
     * @type {string}
     * @memberof PalGachaHistoryResponseItem
     */
    palName?: string | null;
    /**
     * 
     * @type {string}
     * @memberof PalGachaHistoryResponseItem
     */
    palState?: string | null;
    /**
     * 
     * @type {string}
     * @memberof PalGachaHistoryResponseItem
     */
    palTokenId?: string | null;
    /**
     * 
     * @type {number}
     * @memberof PalGachaHistoryResponseItem
     */
    price?: number | null;
    /**
     * 
     * @type {string}
     * @memberof PalGachaHistoryResponseItem
     */
    rarity?: string | null;
}

/**
 * Check if a given object implements the PalGachaHistoryResponseItem interface.
 */
export function instanceOfPalGachaHistoryResponseItem(value: object): value is PalGachaHistoryResponseItem {
    return true;
}

export function PalGachaHistoryResponseItemFromJSON(json: any): PalGachaHistoryResponseItem {
    return PalGachaHistoryResponseItemFromJSONTyped(json, false);
}

export function PalGachaHistoryResponseItemFromJSONTyped(json: any, ignoreDiscriminator: boolean): PalGachaHistoryResponseItem {
    if (json == null) {
        return json;
    }
    return {
        
        'createdAt': json['created_at'] == null ? undefined : json['created_at'],
        'gachaType': json['gacha_type'] == null ? undefined : json['gacha_type'],
        'id': json['id'] == null ? undefined : json['id'],
        'palImageUrl': json['pal_image_url'] == null ? undefined : json['pal_image_url'],
        'palName': json['pal_name'] == null ? undefined : json['pal_name'],
        'palState': json['pal_state'] == null ? undefined : json['pal_state'],
        'palTokenId': json['pal_token_id'] == null ? undefined : json['pal_token_id'],
        'price': json['price'] == null ? undefined : json['price'],
        'rarity': json['rarity'] == null ? undefined : json['rarity'],
    };
}

export function PalGachaHistoryResponseItemToJSON(json: any): PalGachaHistoryResponseItem {
    return PalGachaHistoryResponseItemToJSONTyped(json, false);
}

export function PalGachaHistoryResponseItemToJSONTyped(value?: PalGachaHistoryResponseItem | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'created_at': value['createdAt'],
        'gacha_type': value['gachaType'],
        'id': value['id'],
        'pal_image_url': value['palImageUrl'],
        'pal_name': value['palName'],
        'pal_state': value['palState'],
        'pal_token_id': value['palTokenId'],
        'price': value['price'],
        'rarity': value['rarity'],
    };
}

