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
 * @interface PendingEmplActivityDetails
 */
export interface PendingEmplActivityDetails {
    /**
     * 
     * @type {string}
     * @memberof PendingEmplActivityDetails
     */
    maxTokenAmount?: string | null;
    /**
     * 
     * @type {string}
     * @memberof PendingEmplActivityDetails
     */
    minTokenAmount?: string | null;
    /**
     * 
     * @type {string}
     * @memberof PendingEmplActivityDetails
     */
    tokenOutAddress?: string | null;
    /**
     * 
     * @type {number}
     * @memberof PendingEmplActivityDetails
     */
    transactionFee?: number | null;
}

/**
 * Check if a given object implements the PendingEmplActivityDetails interface.
 */
export function instanceOfPendingEmplActivityDetails(value: object): value is PendingEmplActivityDetails {
    return true;
}

export function PendingEmplActivityDetailsFromJSON(json: any): PendingEmplActivityDetails {
    return PendingEmplActivityDetailsFromJSONTyped(json, false);
}

export function PendingEmplActivityDetailsFromJSONTyped(json: any, ignoreDiscriminator: boolean): PendingEmplActivityDetails {
    if (json == null) {
        return json;
    }
    return {
        
        'maxTokenAmount': json['max_token_amount'] == null ? undefined : json['max_token_amount'],
        'minTokenAmount': json['min_token_amount'] == null ? undefined : json['min_token_amount'],
        'tokenOutAddress': json['token_out_address'] == null ? undefined : json['token_out_address'],
        'transactionFee': json['transaction_fee'] == null ? undefined : json['transaction_fee'],
    };
}

export function PendingEmplActivityDetailsToJSON(json: any): PendingEmplActivityDetails {
    return PendingEmplActivityDetailsToJSONTyped(json, false);
}

export function PendingEmplActivityDetailsToJSONTyped(value?: PendingEmplActivityDetails | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'max_token_amount': value['maxTokenAmount'],
        'min_token_amount': value['minTokenAmount'],
        'token_out_address': value['tokenOutAddress'],
        'transaction_fee': value['transactionFee'],
    };
}

