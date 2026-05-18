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
import type { Details } from './Details.js';
import {
    DetailsFromJSON,
    DetailsFromJSONTyped,
    DetailsToJSON,
    DetailsToJSONTyped,
} from './Details.js';
import type { EmplTransaction } from './EmplTransaction.js';
import {
    EmplTransactionFromJSON,
    EmplTransactionFromJSONTyped,
    EmplTransactionToJSON,
    EmplTransactionToJSONTyped,
} from './EmplTransaction.js';

/**
 * 
 * @export
 * @interface Result
 */
export interface Result {
    /**
     * 
     * @type {string}
     * @memberof Result
     */
    amount?: string | null;
    /**
     * 
     * @type {string}
     * @memberof Result
     */
    currency?: string | null;
    /**
     * 
     * @type {Details}
     * @memberof Result
     */
    details?: Details | null;
    /**
     * 
     * @type {EmplTransaction}
     * @memberof Result
     */
    emplTransaction?: EmplTransaction | null;
    /**
     * 
     * @type {number}
     * @memberof Result
     */
    finalStatusUpdatedAt?: number | null;
    /**
     * 
     * @type {number}
     * @memberof Result
     */
    id?: number | null;
    /**
     * 
     * @type {string}
     * @memberof Result
     */
    status?: string | null;
    /**
     * 
     * @type {string}
     * @memberof Result
     */
    transactionCode?: string | null;
    /**
     * 
     * @type {string}
     * @memberof Result
     */
    walletAddress?: string | null;
}

/**
 * Check if a given object implements the Result interface.
 */
export function instanceOfResult(value: object): value is Result {
    return true;
}

export function ResultFromJSON(json: any): Result {
    return ResultFromJSONTyped(json, false);
}

export function ResultFromJSONTyped(json: any, ignoreDiscriminator: boolean): Result {
    if (json == null) {
        return json;
    }
    return {
        
        'amount': json['amount'] == null ? undefined : json['amount'],
        'currency': json['currency'] == null ? undefined : json['currency'],
        'details': json['details'] == null ? undefined : DetailsFromJSON(json['details']),
        'emplTransaction': json['empl_transaction'] == null ? undefined : EmplTransactionFromJSON(json['empl_transaction']),
        'finalStatusUpdatedAt': json['final_status_updated_at'] == null ? undefined : json['final_status_updated_at'],
        'id': json['id'] == null ? undefined : json['id'],
        'status': json['status'] == null ? undefined : json['status'],
        'transactionCode': json['transaction_code'] == null ? undefined : json['transaction_code'],
        'walletAddress': json['wallet_address'] == null ? undefined : json['wallet_address'],
    };
}

export function ResultToJSON(json: any): Result {
    return ResultToJSONTyped(json, false);
}

export function ResultToJSONTyped(value?: Result | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'amount': value['amount'],
        'currency': value['currency'],
        'details': DetailsToJSON(value['details']),
        'empl_transaction': EmplTransactionToJSON(value['emplTransaction']),
        'final_status_updated_at': value['finalStatusUpdatedAt'],
        'id': value['id'],
        'status': value['status'],
        'transaction_code': value['transactionCode'],
        'wallet_address': value['walletAddress'],
    };
}

