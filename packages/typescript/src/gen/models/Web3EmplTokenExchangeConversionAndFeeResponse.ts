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
import type { Web3EmplTokenExchangeConversionAndFeeDTO } from './Web3EmplTokenExchangeConversionAndFeeDTO.js';
import {
    Web3EmplTokenExchangeConversionAndFeeDTOFromJSON,
    Web3EmplTokenExchangeConversionAndFeeDTOFromJSONTyped,
    Web3EmplTokenExchangeConversionAndFeeDTOToJSON,
    Web3EmplTokenExchangeConversionAndFeeDTOToJSONTyped,
} from './Web3EmplTokenExchangeConversionAndFeeDTO.js';

/**
 * 
 * @export
 * @interface Web3EmplTokenExchangeConversionAndFeeResponse
 */
export interface Web3EmplTokenExchangeConversionAndFeeResponse {
    /**
     * 
     * @type {Web3EmplTokenExchangeConversionAndFeeDTO}
     * @memberof Web3EmplTokenExchangeConversionAndFeeResponse
     */
    result?: Web3EmplTokenExchangeConversionAndFeeDTO | null;
}

/**
 * Check if a given object implements the Web3EmplTokenExchangeConversionAndFeeResponse interface.
 */
export function instanceOfWeb3EmplTokenExchangeConversionAndFeeResponse(value: object): value is Web3EmplTokenExchangeConversionAndFeeResponse {
    return true;
}

export function Web3EmplTokenExchangeConversionAndFeeResponseFromJSON(json: any): Web3EmplTokenExchangeConversionAndFeeResponse {
    return Web3EmplTokenExchangeConversionAndFeeResponseFromJSONTyped(json, false);
}

export function Web3EmplTokenExchangeConversionAndFeeResponseFromJSONTyped(json: any, ignoreDiscriminator: boolean): Web3EmplTokenExchangeConversionAndFeeResponse {
    if (json == null) {
        return json;
    }
    return {
        
        'result': json['result'] == null ? undefined : Web3EmplTokenExchangeConversionAndFeeDTOFromJSON(json['result']),
    };
}

export function Web3EmplTokenExchangeConversionAndFeeResponseToJSON(json: any): Web3EmplTokenExchangeConversionAndFeeResponse {
    return Web3EmplTokenExchangeConversionAndFeeResponseToJSONTyped(json, false);
}

export function Web3EmplTokenExchangeConversionAndFeeResponseToJSONTyped(value?: Web3EmplTokenExchangeConversionAndFeeResponse | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'result': Web3EmplTokenExchangeConversionAndFeeDTOToJSON(value['result']),
    };
}

