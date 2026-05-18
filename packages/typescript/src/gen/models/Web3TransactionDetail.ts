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
 * @interface Web3TransactionDetail
 */
export interface Web3TransactionDetail {
    /**
     * 
     * @type {string}
     * @memberof Web3TransactionDetail
     */
    amount?: string | null;
    /**
     * 
     * @type {number}
     * @memberof Web3TransactionDetail
     */
    chainId?: number | null;
    /**
     * 
     * @type {string}
     * @memberof Web3TransactionDetail
     */
    currency?: string | null;
    /**
     * 
     * @type {number}
     * @memberof Web3TransactionDetail
     */
    emplAmount?: number | null;
    /**
     * 
     * @type {object}
     * @memberof Web3TransactionDetail
     */
    error?: object | null;
    /**
     * 
     * @type {number}
     * @memberof Web3TransactionDetail
     */
    finalStatusUpdatedAtMillis?: number | null;
    /**
     * 
     * @type {number}
     * @memberof Web3TransactionDetail
     */
    id?: number | null;
    /**
     * 
     * @type {string}
     * @memberof Web3TransactionDetail
     */
    maxTokenAmount?: string | null;
    /**
     * 
     * @type {string}
     * @memberof Web3TransactionDetail
     */
    minTokenAmount?: string | null;
    /**
     * 
     * @type {object}
     * @memberof Web3TransactionDetail
     */
    status?: object | null;
    /**
     * 
     * @type {string}
     * @memberof Web3TransactionDetail
     */
    tokenOutAddress?: string | null;
    /**
     * 
     * @type {string}
     * @memberof Web3TransactionDetail
     */
    tokenOutAmount?: string | null;
    /**
     * 
     * @type {string}
     * @memberof Web3TransactionDetail
     */
    tokenSymbol?: string | null;
    /**
     * 
     * @type {number}
     * @memberof Web3TransactionDetail
     */
    transactionFee?: number | null;
    /**
     * 
     * @type {string}
     * @memberof Web3TransactionDetail
     */
    walletAddress?: string | null;
}

/**
 * Check if a given object implements the Web3TransactionDetail interface.
 */
export function instanceOfWeb3TransactionDetail(value: object): value is Web3TransactionDetail {
    return true;
}

export function Web3TransactionDetailFromJSON(json: any): Web3TransactionDetail {
    return Web3TransactionDetailFromJSONTyped(json, false);
}

export function Web3TransactionDetailFromJSONTyped(json: any, ignoreDiscriminator: boolean): Web3TransactionDetail {
    if (json == null) {
        return json;
    }
    return {
        
        'amount': json['amount'] == null ? undefined : json['amount'],
        'chainId': json['chain_id'] == null ? undefined : json['chain_id'],
        'currency': json['currency'] == null ? undefined : json['currency'],
        'emplAmount': json['empl_amount'] == null ? undefined : json['empl_amount'],
        'error': json['error'] == null ? undefined : json['error'],
        'finalStatusUpdatedAtMillis': json['final_status_updated_at_millis'] == null ? undefined : json['final_status_updated_at_millis'],
        'id': json['id'] == null ? undefined : json['id'],
        'maxTokenAmount': json['max_token_amount'] == null ? undefined : json['max_token_amount'],
        'minTokenAmount': json['min_token_amount'] == null ? undefined : json['min_token_amount'],
        'status': json['status'] == null ? undefined : json['status'],
        'tokenOutAddress': json['token_out_address'] == null ? undefined : json['token_out_address'],
        'tokenOutAmount': json['token_out_amount'] == null ? undefined : json['token_out_amount'],
        'tokenSymbol': json['token_symbol'] == null ? undefined : json['token_symbol'],
        'transactionFee': json['transaction_fee'] == null ? undefined : json['transaction_fee'],
        'walletAddress': json['wallet_address'] == null ? undefined : json['wallet_address'],
    };
}

export function Web3TransactionDetailToJSON(json: any): Web3TransactionDetail {
    return Web3TransactionDetailToJSONTyped(json, false);
}

export function Web3TransactionDetailToJSONTyped(value?: Web3TransactionDetail | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'amount': value['amount'],
        'chain_id': value['chainId'],
        'currency': value['currency'],
        'empl_amount': value['emplAmount'],
        'error': value['error'],
        'final_status_updated_at_millis': value['finalStatusUpdatedAtMillis'],
        'id': value['id'],
        'max_token_amount': value['maxTokenAmount'],
        'min_token_amount': value['minTokenAmount'],
        'status': value['status'],
        'token_out_address': value['tokenOutAddress'],
        'token_out_amount': value['tokenOutAmount'],
        'token_symbol': value['tokenSymbol'],
        'transaction_fee': value['transactionFee'],
        'wallet_address': value['walletAddress'],
    };
}

