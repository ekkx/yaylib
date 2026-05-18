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
 * @interface Web3WalletTransactionHistoryHeader
 */
export interface Web3WalletTransactionHistoryHeader {
    /**
     * 
     * @type {number}
     * @memberof Web3WalletTransactionHistoryHeader
     */
    titleId?: number | null;
}

/**
 * Check if a given object implements the Web3WalletTransactionHistoryHeader interface.
 */
export function instanceOfWeb3WalletTransactionHistoryHeader(value: object): value is Web3WalletTransactionHistoryHeader {
    return true;
}

export function Web3WalletTransactionHistoryHeaderFromJSON(json: any): Web3WalletTransactionHistoryHeader {
    return Web3WalletTransactionHistoryHeaderFromJSONTyped(json, false);
}

export function Web3WalletTransactionHistoryHeaderFromJSONTyped(json: any, ignoreDiscriminator: boolean): Web3WalletTransactionHistoryHeader {
    if (json == null) {
        return json;
    }
    return {
        
        'titleId': json['title_id'] == null ? undefined : json['title_id'],
    };
}

export function Web3WalletTransactionHistoryHeaderToJSON(json: any): Web3WalletTransactionHistoryHeader {
    return Web3WalletTransactionHistoryHeaderToJSONTyped(json, false);
}

export function Web3WalletTransactionHistoryHeaderToJSONTyped(value?: Web3WalletTransactionHistoryHeader | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'title_id': value['titleId'],
    };
}

