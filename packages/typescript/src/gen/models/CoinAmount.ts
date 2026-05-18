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
 * @interface CoinAmount
 */
export interface CoinAmount {
    /**
     * 
     * @type {number}
     * @memberof CoinAmount
     */
    free?: number | null;
    /**
     * 
     * @type {number}
     * @memberof CoinAmount
     */
    paid?: number | null;
    /**
     * 
     * @type {number}
     * @memberof CoinAmount
     */
    total?: number | null;
}

/**
 * Check if a given object implements the CoinAmount interface.
 */
export function instanceOfCoinAmount(value: object): value is CoinAmount {
    return true;
}

export function CoinAmountFromJSON(json: any): CoinAmount {
    return CoinAmountFromJSONTyped(json, false);
}

export function CoinAmountFromJSONTyped(json: any, ignoreDiscriminator: boolean): CoinAmount {
    if (json == null) {
        return json;
    }
    return {
        
        'free': json['free'] == null ? undefined : json['free'],
        'paid': json['paid'] == null ? undefined : json['paid'],
        'total': json['total'] == null ? undefined : json['total'],
    };
}

export function CoinAmountToJSON(json: any): CoinAmount {
    return CoinAmountToJSONTyped(json, false);
}

export function CoinAmountToJSONTyped(value?: CoinAmount | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'free': value['free'],
        'paid': value['paid'],
        'total': value['total'],
    };
}

