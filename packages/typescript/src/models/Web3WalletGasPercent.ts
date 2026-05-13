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
 * @interface Web3WalletGasPercent
 */
export interface Web3WalletGasPercent {
    /**
     * 
     * @type {number}
     * @memberof Web3WalletGasPercent
     */
    fast?: number | null;
    /**
     * 
     * @type {number}
     * @memberof Web3WalletGasPercent
     */
    normal?: number | null;
    /**
     * 
     * @type {number}
     * @memberof Web3WalletGasPercent
     */
    slow?: number | null;
}

/**
 * Check if a given object implements the Web3WalletGasPercent interface.
 */
export function instanceOfWeb3WalletGasPercent(value: object): value is Web3WalletGasPercent {
    return true;
}

export function Web3WalletGasPercentFromJSON(json: any): Web3WalletGasPercent {
    return Web3WalletGasPercentFromJSONTyped(json, false);
}

export function Web3WalletGasPercentFromJSONTyped(json: any, ignoreDiscriminator: boolean): Web3WalletGasPercent {
    if (json == null) {
        return json;
    }
    return {
        
        'fast': json['fast'] == null ? undefined : json['fast'],
        'normal': json['normal'] == null ? undefined : json['normal'],
        'slow': json['slow'] == null ? undefined : json['slow'],
    };
}

export function Web3WalletGasPercentToJSON(json: any): Web3WalletGasPercent {
    return Web3WalletGasPercentToJSONTyped(json, false);
}

export function Web3WalletGasPercentToJSONTyped(value?: Web3WalletGasPercent | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'fast': value['fast'],
        'normal': value['normal'],
        'slow': value['slow'],
    };
}

