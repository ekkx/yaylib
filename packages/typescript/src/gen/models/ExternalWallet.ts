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
import type { Web3WalletExternalWallet } from './Web3WalletExternalWallet.js';
import {
    Web3WalletExternalWalletFromJSON,
    Web3WalletExternalWalletFromJSONTyped,
    Web3WalletExternalWalletToJSON,
    Web3WalletExternalWalletToJSONTyped,
} from './Web3WalletExternalWallet.js';

/**
 * 
 * @export
 * @interface ExternalWallet
 */
export interface ExternalWallet {
    /**
     * 
     * @type {Web3WalletExternalWallet}
     * @memberof ExternalWallet
     */
    result?: Web3WalletExternalWallet | null;
}

/**
 * Check if a given object implements the ExternalWallet interface.
 */
export function instanceOfExternalWallet(value: object): value is ExternalWallet {
    return true;
}

export function ExternalWalletFromJSON(json: any): ExternalWallet {
    return ExternalWalletFromJSONTyped(json, false);
}

export function ExternalWalletFromJSONTyped(json: any, ignoreDiscriminator: boolean): ExternalWallet {
    if (json == null) {
        return json;
    }
    return {
        
        'result': json['result'] == null ? undefined : Web3WalletExternalWalletFromJSON(json['result']),
    };
}

export function ExternalWalletToJSON(json: any): ExternalWallet {
    return ExternalWalletToJSONTyped(json, false);
}

export function ExternalWalletToJSONTyped(value?: ExternalWallet | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'result': Web3WalletExternalWalletToJSON(value['result']),
    };
}

