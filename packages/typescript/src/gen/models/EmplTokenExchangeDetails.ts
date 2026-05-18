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
 * @interface EmplTokenExchangeDetails
 */
export interface EmplTokenExchangeDetails {
    /**
     * 
     * @type {number}
     * @memberof EmplTokenExchangeDetails
     */
    chainId?: number | null;
    /**
     * 
     * @type {string}
     * @memberof EmplTokenExchangeDetails
     */
    maxTokenAmount?: string | null;
    /**
     * 
     * @type {string}
     * @memberof EmplTokenExchangeDetails
     */
    minTokenAmount?: string | null;
    /**
     * 
     * @type {string}
     * @memberof EmplTokenExchangeDetails
     */
    tokenOutAddress?: string | null;
    /**
     * 
     * @type {string}
     * @memberof EmplTokenExchangeDetails
     */
    tokenSymbol?: string | null;
    /**
     * 
     * @type {number}
     * @memberof EmplTokenExchangeDetails
     */
    transactionFee?: number | null;
}

/**
 * Check if a given object implements the EmplTokenExchangeDetails interface.
 */
export function instanceOfEmplTokenExchangeDetails(value: object): value is EmplTokenExchangeDetails {
    return true;
}

export function EmplTokenExchangeDetailsFromJSON(json: any): EmplTokenExchangeDetails {
    return EmplTokenExchangeDetailsFromJSONTyped(json, false);
}

export function EmplTokenExchangeDetailsFromJSONTyped(json: any, ignoreDiscriminator: boolean): EmplTokenExchangeDetails {
    if (json == null) {
        return json;
    }
    return {
        
        'chainId': json['chain_id'] == null ? undefined : json['chain_id'],
        'maxTokenAmount': json['max_token_amount'] == null ? undefined : json['max_token_amount'],
        'minTokenAmount': json['min_token_amount'] == null ? undefined : json['min_token_amount'],
        'tokenOutAddress': json['token_out_address'] == null ? undefined : json['token_out_address'],
        'tokenSymbol': json['token_symbol'] == null ? undefined : json['token_symbol'],
        'transactionFee': json['transaction_fee'] == null ? undefined : json['transaction_fee'],
    };
}

export function EmplTokenExchangeDetailsToJSON(json: any): EmplTokenExchangeDetails {
    return EmplTokenExchangeDetailsToJSONTyped(json, false);
}

export function EmplTokenExchangeDetailsToJSONTyped(value?: EmplTokenExchangeDetails | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'chain_id': value['chainId'],
        'max_token_amount': value['maxTokenAmount'],
        'min_token_amount': value['minTokenAmount'],
        'token_out_address': value['tokenOutAddress'],
        'token_symbol': value['tokenSymbol'],
        'transaction_fee': value['transactionFee'],
    };
}

