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
import type { Web3EmplWithdrawalTokensDTO } from './Web3EmplWithdrawalTokensDTO.js';
import {
    Web3EmplWithdrawalTokensDTOFromJSON,
    Web3EmplWithdrawalTokensDTOFromJSONTyped,
    Web3EmplWithdrawalTokensDTOToJSON,
    Web3EmplWithdrawalTokensDTOToJSONTyped,
} from './Web3EmplWithdrawalTokensDTO.js';

/**
 * 
 * @export
 * @interface Web3EmplWithdrawalTokensResponse
 */
export interface Web3EmplWithdrawalTokensResponse {
    /**
     * 
     * @type {Web3EmplWithdrawalTokensDTO}
     * @memberof Web3EmplWithdrawalTokensResponse
     */
    result?: Web3EmplWithdrawalTokensDTO | null;
}

/**
 * Check if a given object implements the Web3EmplWithdrawalTokensResponse interface.
 */
export function instanceOfWeb3EmplWithdrawalTokensResponse(value: object): value is Web3EmplWithdrawalTokensResponse {
    return true;
}

export function Web3EmplWithdrawalTokensResponseFromJSON(json: any): Web3EmplWithdrawalTokensResponse {
    return Web3EmplWithdrawalTokensResponseFromJSONTyped(json, false);
}

export function Web3EmplWithdrawalTokensResponseFromJSONTyped(json: any, ignoreDiscriminator: boolean): Web3EmplWithdrawalTokensResponse {
    if (json == null) {
        return json;
    }
    return {
        
        'result': json['result'] == null ? undefined : Web3EmplWithdrawalTokensDTOFromJSON(json['result']),
    };
}

export function Web3EmplWithdrawalTokensResponseToJSON(json: any): Web3EmplWithdrawalTokensResponse {
    return Web3EmplWithdrawalTokensResponseToJSONTyped(json, false);
}

export function Web3EmplWithdrawalTokensResponseToJSONTyped(value?: Web3EmplWithdrawalTokensResponse | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'result': Web3EmplWithdrawalTokensDTOToJSON(value['result']),
    };
}

