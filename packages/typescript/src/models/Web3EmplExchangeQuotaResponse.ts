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
 * @interface Web3EmplExchangeQuotaResponse
 */
export interface Web3EmplExchangeQuotaResponse {
    /**
     * 
     * @type {number}
     * @memberof Web3EmplExchangeQuotaResponse
     */
    emplExchangeLimit?: number | null;
    /**
     * 
     * @type {number}
     * @memberof Web3EmplExchangeQuotaResponse
     */
    emplExchangeQuotaRemaining?: number | null;
    /**
     * 
     * @type {number}
     * @memberof Web3EmplExchangeQuotaResponse
     */
    timestamp?: number | null;
}

/**
 * Check if a given object implements the Web3EmplExchangeQuotaResponse interface.
 */
export function instanceOfWeb3EmplExchangeQuotaResponse(value: object): value is Web3EmplExchangeQuotaResponse {
    return true;
}

export function Web3EmplExchangeQuotaResponseFromJSON(json: any): Web3EmplExchangeQuotaResponse {
    return Web3EmplExchangeQuotaResponseFromJSONTyped(json, false);
}

export function Web3EmplExchangeQuotaResponseFromJSONTyped(json: any, ignoreDiscriminator: boolean): Web3EmplExchangeQuotaResponse {
    if (json == null) {
        return json;
    }
    return {
        
        'emplExchangeLimit': json['empl_exchange_limit'] == null ? undefined : json['empl_exchange_limit'],
        'emplExchangeQuotaRemaining': json['empl_exchange_quota_remaining'] == null ? undefined : json['empl_exchange_quota_remaining'],
        'timestamp': json['timestamp'] == null ? undefined : json['timestamp'],
    };
}

export function Web3EmplExchangeQuotaResponseToJSON(json: any): Web3EmplExchangeQuotaResponse {
    return Web3EmplExchangeQuotaResponseToJSONTyped(json, false);
}

export function Web3EmplExchangeQuotaResponseToJSONTyped(value?: Web3EmplExchangeQuotaResponse | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'empl_exchange_limit': value['emplExchangeLimit'],
        'empl_exchange_quota_remaining': value['emplExchangeQuotaRemaining'],
        'timestamp': value['timestamp'],
    };
}

