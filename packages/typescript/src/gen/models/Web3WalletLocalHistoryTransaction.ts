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
 * @interface Web3WalletLocalHistoryTransaction
 */
export interface Web3WalletLocalHistoryTransaction {
    /**
     * 
     * @type {number}
     * @memberof Web3WalletLocalHistoryTransaction
     */
    amount?: number | null;
    /**
     * 
     * @type {string}
     * @memberof Web3WalletLocalHistoryTransaction
     */
    assetTypeName?: string | null;
    /**
     * 
     * @type {string}
     * @memberof Web3WalletLocalHistoryTransaction
     */
    chain?: string | null;
    /**
     * 
     * @type {number}
     * @memberof Web3WalletLocalHistoryTransaction
     */
    chainId?: number | null;
    /**
     * 
     * @type {string}
     * @memberof Web3WalletLocalHistoryTransaction
     */
    chainUrl?: string | null;
    /**
     * 
     * @type {string}
     * @memberof Web3WalletLocalHistoryTransaction
     */
    contractAddress?: string | null;
    /**
     * 
     * @type {number}
     * @memberof Web3WalletLocalHistoryTransaction
     */
    gasLimit?: number | null;
    /**
     * 
     * @type {number}
     * @memberof Web3WalletLocalHistoryTransaction
     */
    gasPercentage?: number | null;
    /**
     * 
     * @type {number}
     * @memberof Web3WalletLocalHistoryTransaction
     */
    gasPrice?: number | null;
    /**
     * 
     * @type {string}
     * @memberof Web3WalletLocalHistoryTransaction
     */
    inputsHex?: string | null;
    /**
     * 
     * @type {number}
     * @memberof Web3WalletLocalHistoryTransaction
     */
    nonce?: number | null;
    /**
     * 
     * @type {string}
     * @memberof Web3WalletLocalHistoryTransaction
     */
    palId?: string | null;
    /**
     * 
     * @type {number}
     * @memberof Web3WalletLocalHistoryTransaction
     */
    timestamp?: number | null;
    /**
     * 
     * @type {string}
     * @memberof Web3WalletLocalHistoryTransaction
     */
    transactionHashHex?: string | null;
    /**
     * 
     * @type {string}
     * @memberof Web3WalletLocalHistoryTransaction
     */
    transactionHistoryStatusName?: string | null;
    /**
     * 
     * @type {string}
     * @memberof Web3WalletLocalHistoryTransaction
     */
    transferModeKey?: string | null;
    /**
     * 
     * @type {string}
     * @memberof Web3WalletLocalHistoryTransaction
     */
    txUniqueId?: string | null;
}

/**
 * Check if a given object implements the Web3WalletLocalHistoryTransaction interface.
 */
export function instanceOfWeb3WalletLocalHistoryTransaction(value: object): value is Web3WalletLocalHistoryTransaction {
    return true;
}

export function Web3WalletLocalHistoryTransactionFromJSON(json: any): Web3WalletLocalHistoryTransaction {
    return Web3WalletLocalHistoryTransactionFromJSONTyped(json, false);
}

export function Web3WalletLocalHistoryTransactionFromJSONTyped(json: any, ignoreDiscriminator: boolean): Web3WalletLocalHistoryTransaction {
    if (json == null) {
        return json;
    }
    return {
        
        'amount': json['amount'] == null ? undefined : json['amount'],
        'assetTypeName': json['asset_type_name'] == null ? undefined : json['asset_type_name'],
        'chain': json['chain'] == null ? undefined : json['chain'],
        'chainId': json['chain_id'] == null ? undefined : json['chain_id'],
        'chainUrl': json['chain_url'] == null ? undefined : json['chain_url'],
        'contractAddress': json['contract_address'] == null ? undefined : json['contract_address'],
        'gasLimit': json['gas_limit'] == null ? undefined : json['gas_limit'],
        'gasPercentage': json['gas_percentage'] == null ? undefined : json['gas_percentage'],
        'gasPrice': json['gas_price'] == null ? undefined : json['gas_price'],
        'inputsHex': json['inputs_hex'] == null ? undefined : json['inputs_hex'],
        'nonce': json['nonce'] == null ? undefined : json['nonce'],
        'palId': json['pal_id'] == null ? undefined : json['pal_id'],
        'timestamp': json['timestamp'] == null ? undefined : json['timestamp'],
        'transactionHashHex': json['transaction_hash_hex'] == null ? undefined : json['transaction_hash_hex'],
        'transactionHistoryStatusName': json['transaction_history_status_name'] == null ? undefined : json['transaction_history_status_name'],
        'transferModeKey': json['transfer_mode_key'] == null ? undefined : json['transfer_mode_key'],
        'txUniqueId': json['tx_unique_id'] == null ? undefined : json['tx_unique_id'],
    };
}

export function Web3WalletLocalHistoryTransactionToJSON(json: any): Web3WalletLocalHistoryTransaction {
    return Web3WalletLocalHistoryTransactionToJSONTyped(json, false);
}

export function Web3WalletLocalHistoryTransactionToJSONTyped(value?: Web3WalletLocalHistoryTransaction | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'amount': value['amount'],
        'asset_type_name': value['assetTypeName'],
        'chain': value['chain'],
        'chain_id': value['chainId'],
        'chain_url': value['chainUrl'],
        'contract_address': value['contractAddress'],
        'gas_limit': value['gasLimit'],
        'gas_percentage': value['gasPercentage'],
        'gas_price': value['gasPrice'],
        'inputs_hex': value['inputsHex'],
        'nonce': value['nonce'],
        'pal_id': value['palId'],
        'timestamp': value['timestamp'],
        'transaction_hash_hex': value['transactionHashHex'],
        'transaction_history_status_name': value['transactionHistoryStatusName'],
        'transfer_mode_key': value['transferModeKey'],
        'tx_unique_id': value['txUniqueId'],
    };
}

