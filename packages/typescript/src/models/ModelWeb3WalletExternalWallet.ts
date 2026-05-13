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
 * @interface ModelWeb3WalletExternalWallet
 */
export interface ModelWeb3WalletExternalWallet {
    /**
     * 
     * @type {number}
     * @memberof ModelWeb3WalletExternalWallet
     */
    eggsCount?: number | null;
    /**
     * 
     * @type {number}
     * @memberof ModelWeb3WalletExternalWallet
     */
    palsCount?: number | null;
    /**
     * 
     * @type {string}
     * @memberof ModelWeb3WalletExternalWallet
     */
    walletAddress?: string | null;
}

/**
 * Check if a given object implements the ModelWeb3WalletExternalWallet interface.
 */
export function instanceOfModelWeb3WalletExternalWallet(value: object): value is ModelWeb3WalletExternalWallet {
    return true;
}

export function ModelWeb3WalletExternalWalletFromJSON(json: any): ModelWeb3WalletExternalWallet {
    return ModelWeb3WalletExternalWalletFromJSONTyped(json, false);
}

export function ModelWeb3WalletExternalWalletFromJSONTyped(json: any, ignoreDiscriminator: boolean): ModelWeb3WalletExternalWallet {
    if (json == null) {
        return json;
    }
    return {
        
        'eggsCount': json['eggs_count'] == null ? undefined : json['eggs_count'],
        'palsCount': json['pals_count'] == null ? undefined : json['pals_count'],
        'walletAddress': json['wallet_address'] == null ? undefined : json['wallet_address'],
    };
}

export function ModelWeb3WalletExternalWalletToJSON(json: any): ModelWeb3WalletExternalWallet {
    return ModelWeb3WalletExternalWalletToJSONTyped(json, false);
}

export function ModelWeb3WalletExternalWalletToJSONTyped(value?: ModelWeb3WalletExternalWallet | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'eggs_count': value['eggsCount'],
        'pals_count': value['palsCount'],
        'wallet_address': value['walletAddress'],
    };
}

