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
import type { Web3WalletTransactionHistory } from './Web3WalletTransactionHistory.js';
import {
    Web3WalletTransactionHistoryFromJSON,
    Web3WalletTransactionHistoryFromJSONTyped,
    Web3WalletTransactionHistoryToJSON,
    Web3WalletTransactionHistoryToJSONTyped,
} from './Web3WalletTransactionHistory.js';

/**
 * 
 * @export
 * @interface TransactionHistory
 */
export interface TransactionHistory {
    /**
     * 
     * @type {string}
     * @memberof TransactionHistory
     */
    cursor?: string | null;
    /**
     * 
     * @type {Array<Web3WalletTransactionHistory>}
     * @memberof TransactionHistory
     */
    result?: Array<Web3WalletTransactionHistory> | null;
}

/**
 * Check if a given object implements the TransactionHistory interface.
 */
export function instanceOfTransactionHistory(value: object): value is TransactionHistory {
    return true;
}

export function TransactionHistoryFromJSON(json: any): TransactionHistory {
    return TransactionHistoryFromJSONTyped(json, false);
}

export function TransactionHistoryFromJSONTyped(json: any, ignoreDiscriminator: boolean): TransactionHistory {
    if (json == null) {
        return json;
    }
    return {
        
        'cursor': json['cursor'] == null ? undefined : json['cursor'],
        'result': json['result'] == null ? undefined : ((json['result'] as Array<any>).map(Web3WalletTransactionHistoryFromJSON)),
    };
}

export function TransactionHistoryToJSON(json: any): TransactionHistory {
    return TransactionHistoryToJSONTyped(json, false);
}

export function TransactionHistoryToJSONTyped(value?: TransactionHistory | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'cursor': value['cursor'],
        'result': value['result'] == null ? undefined : ((value['result'] as Array<any>).map(Web3WalletTransactionHistoryToJSON)),
    };
}

