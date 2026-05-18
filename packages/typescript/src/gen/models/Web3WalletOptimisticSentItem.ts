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
 * @interface Web3WalletOptimisticSentItem
 */
export interface Web3WalletOptimisticSentItem {
    /**
     * 
     * @type {number}
     * @memberof Web3WalletOptimisticSentItem
     */
    chainId?: number | null;
    /**
     * 
     * @type {string}
     * @memberof Web3WalletOptimisticSentItem
     */
    contractAddress?: string | null;
    /**
     * 
     * @type {number}
     * @memberof Web3WalletOptimisticSentItem
     */
    createdAtMillis?: number | null;
    /**
     * 
     * @type {string}
     * @memberof Web3WalletOptimisticSentItem
     */
    id?: string | null;
    /**
     * 
     * @type {string}
     * @memberof Web3WalletOptimisticSentItem
     */
    tokenId?: string | null;
}

/**
 * Check if a given object implements the Web3WalletOptimisticSentItem interface.
 */
export function instanceOfWeb3WalletOptimisticSentItem(value: object): value is Web3WalletOptimisticSentItem {
    return true;
}

export function Web3WalletOptimisticSentItemFromJSON(json: any): Web3WalletOptimisticSentItem {
    return Web3WalletOptimisticSentItemFromJSONTyped(json, false);
}

export function Web3WalletOptimisticSentItemFromJSONTyped(json: any, ignoreDiscriminator: boolean): Web3WalletOptimisticSentItem {
    if (json == null) {
        return json;
    }
    return {
        
        'chainId': json['chain_id'] == null ? undefined : json['chain_id'],
        'contractAddress': json['contract_address'] == null ? undefined : json['contract_address'],
        'createdAtMillis': json['created_at_millis'] == null ? undefined : json['created_at_millis'],
        'id': json['id'] == null ? undefined : json['id'],
        'tokenId': json['token_id'] == null ? undefined : json['token_id'],
    };
}

export function Web3WalletOptimisticSentItemToJSON(json: any): Web3WalletOptimisticSentItem {
    return Web3WalletOptimisticSentItemToJSONTyped(json, false);
}

export function Web3WalletOptimisticSentItemToJSONTyped(value?: Web3WalletOptimisticSentItem | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'chain_id': value['chainId'],
        'contract_address': value['contractAddress'],
        'created_at_millis': value['createdAtMillis'],
        'id': value['id'],
        'token_id': value['tokenId'],
    };
}

