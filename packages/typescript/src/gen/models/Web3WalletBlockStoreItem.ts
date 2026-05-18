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
 * @interface Web3WalletBlockStoreItem
 */
export interface Web3WalletBlockStoreItem {
    /**
     * 
     * @type {number}
     * @memberof Web3WalletBlockStoreItem
     */
    createdAtMillis?: number | null;
    /**
     * 
     * @type {string}
     * @memberof Web3WalletBlockStoreItem
     */
    name?: string | null;
    /**
     * 
     * @type {string}
     * @memberof Web3WalletBlockStoreItem
     */
    privateKey?: string | null;
    /**
     * 
     * @type {string}
     * @memberof Web3WalletBlockStoreItem
     */
    seedPhrase?: string | null;
    /**
     * 
     * @type {string}
     * @memberof Web3WalletBlockStoreItem
     */
    walletAddress?: string | null;
}

/**
 * Check if a given object implements the Web3WalletBlockStoreItem interface.
 */
export function instanceOfWeb3WalletBlockStoreItem(value: object): value is Web3WalletBlockStoreItem {
    return true;
}

export function Web3WalletBlockStoreItemFromJSON(json: any): Web3WalletBlockStoreItem {
    return Web3WalletBlockStoreItemFromJSONTyped(json, false);
}

export function Web3WalletBlockStoreItemFromJSONTyped(json: any, ignoreDiscriminator: boolean): Web3WalletBlockStoreItem {
    if (json == null) {
        return json;
    }
    return {
        
        'createdAtMillis': json['created_at_millis'] == null ? undefined : json['created_at_millis'],
        'name': json['name'] == null ? undefined : json['name'],
        'privateKey': json['private_key'] == null ? undefined : json['private_key'],
        'seedPhrase': json['seed_phrase'] == null ? undefined : json['seed_phrase'],
        'walletAddress': json['wallet_address'] == null ? undefined : json['wallet_address'],
    };
}

export function Web3WalletBlockStoreItemToJSON(json: any): Web3WalletBlockStoreItem {
    return Web3WalletBlockStoreItemToJSONTyped(json, false);
}

export function Web3WalletBlockStoreItemToJSONTyped(value?: Web3WalletBlockStoreItem | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'created_at_millis': value['createdAtMillis'],
        'name': value['name'],
        'private_key': value['privateKey'],
        'seed_phrase': value['seedPhrase'],
        'wallet_address': value['walletAddress'],
    };
}

