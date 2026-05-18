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
import type { Web3EmplWithdrawalTokenDTO } from './Web3EmplWithdrawalTokenDTO.js';
import {
    Web3EmplWithdrawalTokenDTOFromJSON,
    Web3EmplWithdrawalTokenDTOFromJSONTyped,
    Web3EmplWithdrawalTokenDTOToJSON,
    Web3EmplWithdrawalTokenDTOToJSONTyped,
} from './Web3EmplWithdrawalTokenDTO.js';

/**
 * 
 * @export
 * @interface Web3EmplWithdrawalTokensDTO
 */
export interface Web3EmplWithdrawalTokensDTO {
    /**
     * 
     * @type {Array<Web3EmplWithdrawalTokenDTO>}
     * @memberof Web3EmplWithdrawalTokensDTO
     */
    tokens?: Array<Web3EmplWithdrawalTokenDTO> | null;
}

/**
 * Check if a given object implements the Web3EmplWithdrawalTokensDTO interface.
 */
export function instanceOfWeb3EmplWithdrawalTokensDTO(value: object): value is Web3EmplWithdrawalTokensDTO {
    return true;
}

export function Web3EmplWithdrawalTokensDTOFromJSON(json: any): Web3EmplWithdrawalTokensDTO {
    return Web3EmplWithdrawalTokensDTOFromJSONTyped(json, false);
}

export function Web3EmplWithdrawalTokensDTOFromJSONTyped(json: any, ignoreDiscriminator: boolean): Web3EmplWithdrawalTokensDTO {
    if (json == null) {
        return json;
    }
    return {
        
        'tokens': json['tokens'] == null ? undefined : ((json['tokens'] as Array<any>).map(Web3EmplWithdrawalTokenDTOFromJSON)),
    };
}

export function Web3EmplWithdrawalTokensDTOToJSON(json: any): Web3EmplWithdrawalTokensDTO {
    return Web3EmplWithdrawalTokensDTOToJSONTyped(json, false);
}

export function Web3EmplWithdrawalTokensDTOToJSONTyped(value?: Web3EmplWithdrawalTokensDTO | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'tokens': value['tokens'] == null ? undefined : ((value['tokens'] as Array<any>).map(Web3EmplWithdrawalTokenDTOToJSON)),
    };
}

