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
 * @interface Web3WalletGasFee
 */
export interface Web3WalletGasFee {
    /**
     * 
     * @type {number}
     * @memberof Web3WalletGasFee
     */
    gasLimit?: number | null;
    /**
     * 
     * @type {number}
     * @memberof Web3WalletGasFee
     */
    gasPriceWei?: number | null;
}

/**
 * Check if a given object implements the Web3WalletGasFee interface.
 */
export function instanceOfWeb3WalletGasFee(value: object): value is Web3WalletGasFee {
    return true;
}

export function Web3WalletGasFeeFromJSON(json: any): Web3WalletGasFee {
    return Web3WalletGasFeeFromJSONTyped(json, false);
}

export function Web3WalletGasFeeFromJSONTyped(json: any, ignoreDiscriminator: boolean): Web3WalletGasFee {
    if (json == null) {
        return json;
    }
    return {
        
        'gasLimit': json['gas_limit'] == null ? undefined : json['gas_limit'],
        'gasPriceWei': json['gas_price_wei'] == null ? undefined : json['gas_price_wei'],
    };
}

export function Web3WalletGasFeeToJSON(json: any): Web3WalletGasFee {
    return Web3WalletGasFeeToJSONTyped(json, false);
}

export function Web3WalletGasFeeToJSONTyped(value?: Web3WalletGasFee | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'gas_limit': value['gasLimit'],
        'gas_price_wei': value['gasPriceWei'],
    };
}

