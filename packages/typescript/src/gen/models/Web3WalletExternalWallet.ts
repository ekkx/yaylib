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
import type { Balance } from './Balance.js';
import {
    BalanceFromJSON,
    BalanceFromJSONTyped,
    BalanceToJSON,
    BalanceToJSONTyped,
} from './Balance.js';

/**
 * 
 * @export
 * @interface Web3WalletExternalWallet
 */
export interface Web3WalletExternalWallet {
    /**
     * 
     * @type {Balance}
     * @memberof Web3WalletExternalWallet
     */
    balance?: Balance | null;
    /**
     * 
     * @type {string}
     * @memberof Web3WalletExternalWallet
     */
    walletAddress?: string | null;
}

/**
 * Check if a given object implements the Web3WalletExternalWallet interface.
 */
export function instanceOfWeb3WalletExternalWallet(value: object): value is Web3WalletExternalWallet {
    return true;
}

export function Web3WalletExternalWalletFromJSON(json: any): Web3WalletExternalWallet {
    return Web3WalletExternalWalletFromJSONTyped(json, false);
}

export function Web3WalletExternalWalletFromJSONTyped(json: any, ignoreDiscriminator: boolean): Web3WalletExternalWallet {
    if (json == null) {
        return json;
    }
    return {
        
        'balance': json['balance'] == null ? undefined : BalanceFromJSON(json['balance']),
        'walletAddress': json['wallet_address'] == null ? undefined : json['wallet_address'],
    };
}

export function Web3WalletExternalWalletToJSON(json: any): Web3WalletExternalWallet {
    return Web3WalletExternalWalletToJSONTyped(json, false);
}

export function Web3WalletExternalWalletToJSONTyped(value?: Web3WalletExternalWallet | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'balance': BalanceToJSON(value['balance']),
        'wallet_address': value['walletAddress'],
    };
}

