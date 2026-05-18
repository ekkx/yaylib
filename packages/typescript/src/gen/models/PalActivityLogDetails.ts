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
import type { PalGrade } from './PalGrade.js';
import {
    PalGradeFromJSON,
    PalGradeFromJSONTyped,
    PalGradeToJSON,
    PalGradeToJSONTyped,
} from './PalGrade.js';

/**
 * 
 * @export
 * @interface PalActivityLogDetails
 */
export interface PalActivityLogDetails {
    /**
     * 
     * @type {number}
     * @memberof PalActivityLogDetails
     */
    attacking?: number | null;
    /**
     * 
     * @type {number}
     * @memberof PalActivityLogDetails
     */
    cost?: number | null;
    /**
     * 
     * @type {string}
     * @memberof PalActivityLogDetails
     */
    gachaType?: string | null;
    /**
     * 
     * @type {string}
     * @memberof PalActivityLogDetails
     */
    itemStatus?: string | null;
    /**
     * 
     * @type {string}
     * @memberof PalActivityLogDetails
     */
    itemSubtype?: string | null;
    /**
     * 
     * @type {string}
     * @memberof PalActivityLogDetails
     */
    itemType?: string | null;
    /**
     * 
     * @type {number}
     * @memberof PalActivityLogDetails
     */
    level?: number | null;
    /**
     * 
     * @type {string}
     * @memberof PalActivityLogDetails
     */
    nftType?: string | null;
    /**
     * 
     * @type {number}
     * @memberof PalActivityLogDetails
     */
    outcome?: number | null;
    /**
     * 
     * @type {PalGrade}
     * @memberof PalActivityLogDetails
     */
    palGrade?: PalGrade | null;
    /**
     * 
     * @type {string}
     * @memberof PalActivityLogDetails
     */
    palImage?: string | null;
    /**
     * 
     * @type {string}
     * @memberof PalActivityLogDetails
     */
    palName?: string | null;
    /**
     * 
     * @type {Array<number>}
     * @memberof PalActivityLogDetails
     */
    places?: Array<number> | null;
    /**
     * 
     * @type {string}
     * @memberof PalActivityLogDetails
     */
    raceId?: string | null;
    /**
     * 
     * @type {string}
     * @memberof PalActivityLogDetails
     */
    rank?: string | null;
    /**
     * 
     * @type {string}
     * @memberof PalActivityLogDetails
     */
    result?: string | null;
    /**
     * 
     * @type {string}
     * @memberof PalActivityLogDetails
     */
    tokenId?: string | null;
    /**
     * 
     * @type {string}
     * @memberof PalActivityLogDetails
     */
    transactionCode?: string | null;
}



/**
 * Check if a given object implements the PalActivityLogDetails interface.
 */
export function instanceOfPalActivityLogDetails(value: object): value is PalActivityLogDetails {
    return true;
}

export function PalActivityLogDetailsFromJSON(json: any): PalActivityLogDetails {
    return PalActivityLogDetailsFromJSONTyped(json, false);
}

export function PalActivityLogDetailsFromJSONTyped(json: any, ignoreDiscriminator: boolean): PalActivityLogDetails {
    if (json == null) {
        return json;
    }
    return {
        
        'attacking': json['attacking'] == null ? undefined : json['attacking'],
        'cost': json['cost'] == null ? undefined : json['cost'],
        'gachaType': json['gacha_type'] == null ? undefined : json['gacha_type'],
        'itemStatus': json['item_status'] == null ? undefined : json['item_status'],
        'itemSubtype': json['item_subtype'] == null ? undefined : json['item_subtype'],
        'itemType': json['item_type'] == null ? undefined : json['item_type'],
        'level': json['level'] == null ? undefined : json['level'],
        'nftType': json['nft_type'] == null ? undefined : json['nft_type'],
        'outcome': json['outcome'] == null ? undefined : json['outcome'],
        'palGrade': json['pal_grade'] == null ? undefined : PalGradeFromJSON(json['pal_grade']),
        'palImage': json['pal_image'] == null ? undefined : json['pal_image'],
        'palName': json['pal_name'] == null ? undefined : json['pal_name'],
        'places': json['places'] == null ? undefined : json['places'],
        'raceId': json['race_id'] == null ? undefined : json['race_id'],
        'rank': json['rank'] == null ? undefined : json['rank'],
        'result': json['result'] == null ? undefined : json['result'],
        'tokenId': json['token_id'] == null ? undefined : json['token_id'],
        'transactionCode': json['transaction_code'] == null ? undefined : json['transaction_code'],
    };
}

export function PalActivityLogDetailsToJSON(json: any): PalActivityLogDetails {
    return PalActivityLogDetailsToJSONTyped(json, false);
}

export function PalActivityLogDetailsToJSONTyped(value?: PalActivityLogDetails | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'attacking': value['attacking'],
        'cost': value['cost'],
        'gacha_type': value['gachaType'],
        'item_status': value['itemStatus'],
        'item_subtype': value['itemSubtype'],
        'item_type': value['itemType'],
        'level': value['level'],
        'nft_type': value['nftType'],
        'outcome': value['outcome'],
        'pal_grade': PalGradeToJSON(value['palGrade']),
        'pal_image': value['palImage'],
        'pal_name': value['palName'],
        'places': value['places'],
        'race_id': value['raceId'],
        'rank': value['rank'],
        'result': value['result'],
        'token_id': value['tokenId'],
        'transaction_code': value['transactionCode'],
    };
}

