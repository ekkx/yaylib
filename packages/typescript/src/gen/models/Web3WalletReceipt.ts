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
import type { Log } from './Log';
import {
    LogFromJSON,
    LogFromJSONTyped,
    LogToJSON,
    LogToJSONTyped,
} from './Log';

/**
 * 
 * @export
 * @interface Web3WalletReceipt
 */
export interface Web3WalletReceipt {
    /**
     * 
     * @type {string}
     * @memberof Web3WalletReceipt
     */
    blockHash?: string | null;
    /**
     * 
     * @type {string}
     * @memberof Web3WalletReceipt
     */
    blockNumber?: string | null;
    /**
     * 
     * @type {string}
     * @memberof Web3WalletReceipt
     */
    contractAddress?: string | null;
    /**
     * 
     * @type {string}
     * @memberof Web3WalletReceipt
     */
    cumulativeGasUsed?: string | null;
    /**
     * 
     * @type {string}
     * @memberof Web3WalletReceipt
     */
    gasUsed?: string | null;
    /**
     * 
     * @type {Array<Log>}
     * @memberof Web3WalletReceipt
     */
    logs?: Array<Log> | null;
    /**
     * 
     * @type {string}
     * @memberof Web3WalletReceipt
     */
    logsBloom?: string | null;
    /**
     * 
     * @type {string}
     * @memberof Web3WalletReceipt
     */
    root?: string | null;
    /**
     * 
     * @type {string}
     * @memberof Web3WalletReceipt
     */
    status?: string | null;
    /**
     * 
     * @type {string}
     * @memberof Web3WalletReceipt
     */
    transactionHash?: string | null;
    /**
     * 
     * @type {string}
     * @memberof Web3WalletReceipt
     */
    transactionIndex?: string | null;
}

/**
 * Check if a given object implements the Web3WalletReceipt interface.
 */
export function instanceOfWeb3WalletReceipt(value: object): value is Web3WalletReceipt {
    return true;
}

export function Web3WalletReceiptFromJSON(json: any): Web3WalletReceipt {
    return Web3WalletReceiptFromJSONTyped(json, false);
}

export function Web3WalletReceiptFromJSONTyped(json: any, ignoreDiscriminator: boolean): Web3WalletReceipt {
    if (json == null) {
        return json;
    }
    return {
        
        'blockHash': json['block_hash'] == null ? undefined : json['block_hash'],
        'blockNumber': json['block_number'] == null ? undefined : json['block_number'],
        'contractAddress': json['contract_address'] == null ? undefined : json['contract_address'],
        'cumulativeGasUsed': json['cumulative_gas_used'] == null ? undefined : json['cumulative_gas_used'],
        'gasUsed': json['gas_used'] == null ? undefined : json['gas_used'],
        'logs': json['logs'] == null ? undefined : ((json['logs'] as Array<any>).map(LogFromJSON)),
        'logsBloom': json['logs_bloom'] == null ? undefined : json['logs_bloom'],
        'root': json['root'] == null ? undefined : json['root'],
        'status': json['status'] == null ? undefined : json['status'],
        'transactionHash': json['transaction_hash'] == null ? undefined : json['transaction_hash'],
        'transactionIndex': json['transaction_index'] == null ? undefined : json['transaction_index'],
    };
}

export function Web3WalletReceiptToJSON(json: any): Web3WalletReceipt {
    return Web3WalletReceiptToJSONTyped(json, false);
}

export function Web3WalletReceiptToJSONTyped(value?: Web3WalletReceipt | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'block_hash': value['blockHash'],
        'block_number': value['blockNumber'],
        'contract_address': value['contractAddress'],
        'cumulative_gas_used': value['cumulativeGasUsed'],
        'gas_used': value['gasUsed'],
        'logs': value['logs'] == null ? undefined : ((value['logs'] as Array<any>).map(LogToJSON)),
        'logs_bloom': value['logsBloom'],
        'root': value['root'],
        'status': value['status'],
        'transaction_hash': value['transactionHash'],
        'transaction_index': value['transactionIndex'],
    };
}

