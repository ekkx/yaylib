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
 * @interface CoinExpiration
 */
export interface CoinExpiration {
    /**
     * 
     * @type {number}
     * @memberof CoinExpiration
     */
    amount?: number | null;
    /**
     * 
     * @type {number}
     * @memberof CoinExpiration
     */
    expiredAtSeconds?: number | null;
}

/**
 * Check if a given object implements the CoinExpiration interface.
 */
export function instanceOfCoinExpiration(value: object): value is CoinExpiration {
    return true;
}

export function CoinExpirationFromJSON(json: any): CoinExpiration {
    return CoinExpirationFromJSONTyped(json, false);
}

export function CoinExpirationFromJSONTyped(json: any, ignoreDiscriminator: boolean): CoinExpiration {
    if (json == null) {
        return json;
    }
    return {
        
        'amount': json['amount'] == null ? undefined : json['amount'],
        'expiredAtSeconds': json['expired_at_seconds'] == null ? undefined : json['expired_at_seconds'],
    };
}

export function CoinExpirationToJSON(json: any): CoinExpiration {
    return CoinExpirationToJSONTyped(json, false);
}

export function CoinExpirationToJSONTyped(value?: CoinExpiration | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'amount': value['amount'],
        'expired_at_seconds': value['expiredAtSeconds'],
    };
}

