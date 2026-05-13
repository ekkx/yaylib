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
/**
 * 
 * @export
 * @interface Details
 */
export interface Details {
    /**
     * 
     * @type {number}
     * @memberof Details
     */
    chainId?: number | null;
    /**
     * 
     * @type {string}
     * @memberof Details
     */
    error?: string | null;
    /**
     * 
     * @type {string}
     * @memberof Details
     */
    maxTokenAmount?: string | null;
    /**
     * 
     * @type {string}
     * @memberof Details
     */
    minTokenAmount?: string | null;
    /**
     * 
     * @type {string}
     * @memberof Details
     */
    tokenOutAddress?: string | null;
    /**
     * 
     * @type {string}
     * @memberof Details
     */
    tokenOutAmount?: string | null;
    /**
     * 
     * @type {string}
     * @memberof Details
     */
    tokenSymbol?: string | null;
    /**
     * 
     * @type {string}
     * @memberof Details
     */
    transactionFee?: string | null;
}

/**
 * Check if a given object implements the Details interface.
 */
export function instanceOfDetails(value: object): value is Details {
    return true;
}

export function DetailsFromJSON(json: any): Details {
    return DetailsFromJSONTyped(json, false);
}

export function DetailsFromJSONTyped(json: any, ignoreDiscriminator: boolean): Details {
    if (json == null) {
        return json;
    }
    return {
        
        'chainId': json['chain_id'] == null ? undefined : json['chain_id'],
        'error': json['error'] == null ? undefined : json['error'],
        'maxTokenAmount': json['max_token_amount'] == null ? undefined : json['max_token_amount'],
        'minTokenAmount': json['min_token_amount'] == null ? undefined : json['min_token_amount'],
        'tokenOutAddress': json['token_out_address'] == null ? undefined : json['token_out_address'],
        'tokenOutAmount': json['token_out_amount'] == null ? undefined : json['token_out_amount'],
        'tokenSymbol': json['token_symbol'] == null ? undefined : json['token_symbol'],
        'transactionFee': json['transaction_fee'] == null ? undefined : json['transaction_fee'],
    };
}

export function DetailsToJSON(json: any): Details {
    return DetailsToJSONTyped(json, false);
}

export function DetailsToJSONTyped(value?: Details | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'chain_id': value['chainId'],
        'error': value['error'],
        'max_token_amount': value['maxTokenAmount'],
        'min_token_amount': value['minTokenAmount'],
        'token_out_address': value['tokenOutAddress'],
        'token_out_amount': value['tokenOutAmount'],
        'token_symbol': value['tokenSymbol'],
        'transaction_fee': value['transactionFee'],
    };
}

