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
 * @interface ModelCoinExpiration
 */
export interface ModelCoinExpiration {
    /**
     * 
     * @type {number}
     * @memberof ModelCoinExpiration
     */
    amount?: number | null;
    /**
     * 
     * @type {number}
     * @memberof ModelCoinExpiration
     */
    expiredAtSeconds?: number | null;
}

/**
 * Check if a given object implements the ModelCoinExpiration interface.
 */
export function instanceOfModelCoinExpiration(value: object): value is ModelCoinExpiration {
    return true;
}

export function ModelCoinExpirationFromJSON(json: any): ModelCoinExpiration {
    return ModelCoinExpirationFromJSONTyped(json, false);
}

export function ModelCoinExpirationFromJSONTyped(json: any, ignoreDiscriminator: boolean): ModelCoinExpiration {
    if (json == null) {
        return json;
    }
    return {
        
        'amount': json['amount'] == null ? undefined : json['amount'],
        'expiredAtSeconds': json['expired_at_seconds'] == null ? undefined : json['expired_at_seconds'],
    };
}

export function ModelCoinExpirationToJSON(json: any): ModelCoinExpiration {
    return ModelCoinExpirationToJSONTyped(json, false);
}

export function ModelCoinExpirationToJSONTyped(value?: ModelCoinExpiration | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'amount': value['amount'],
        'expired_at_seconds': value['expiredAtSeconds'],
    };
}

