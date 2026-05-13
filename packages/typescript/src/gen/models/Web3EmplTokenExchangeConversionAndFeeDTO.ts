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
 * @interface Web3EmplTokenExchangeConversionAndFeeDTO
 */
export interface Web3EmplTokenExchangeConversionAndFeeDTO {
    /**
     * 
     * @type {string}
     * @memberof Web3EmplTokenExchangeConversionAndFeeDTO
     */
    emplAmount?: string | null;
    /**
     * 
     * @type {string}
     * @memberof Web3EmplTokenExchangeConversionAndFeeDTO
     */
    emplTxFee?: string | null;
    /**
     * 
     * @type {string}
     * @memberof Web3EmplTokenExchangeConversionAndFeeDTO
     */
    maxAmountOut?: string | null;
    /**
     * 
     * @type {string}
     * @memberof Web3EmplTokenExchangeConversionAndFeeDTO
     */
    minAmountOut?: string | null;
    /**
     * 
     * @type {string}
     * @memberof Web3EmplTokenExchangeConversionAndFeeDTO
     */
    tokenOutAddress?: string | null;
}

/**
 * Check if a given object implements the Web3EmplTokenExchangeConversionAndFeeDTO interface.
 */
export function instanceOfWeb3EmplTokenExchangeConversionAndFeeDTO(value: object): value is Web3EmplTokenExchangeConversionAndFeeDTO {
    return true;
}

export function Web3EmplTokenExchangeConversionAndFeeDTOFromJSON(json: any): Web3EmplTokenExchangeConversionAndFeeDTO {
    return Web3EmplTokenExchangeConversionAndFeeDTOFromJSONTyped(json, false);
}

export function Web3EmplTokenExchangeConversionAndFeeDTOFromJSONTyped(json: any, ignoreDiscriminator: boolean): Web3EmplTokenExchangeConversionAndFeeDTO {
    if (json == null) {
        return json;
    }
    return {
        
        'emplAmount': json['empl_amount'] == null ? undefined : json['empl_amount'],
        'emplTxFee': json['empl_tx_fee'] == null ? undefined : json['empl_tx_fee'],
        'maxAmountOut': json['max_amount_out'] == null ? undefined : json['max_amount_out'],
        'minAmountOut': json['min_amount_out'] == null ? undefined : json['min_amount_out'],
        'tokenOutAddress': json['token_out_address'] == null ? undefined : json['token_out_address'],
    };
}

export function Web3EmplTokenExchangeConversionAndFeeDTOToJSON(json: any): Web3EmplTokenExchangeConversionAndFeeDTO {
    return Web3EmplTokenExchangeConversionAndFeeDTOToJSONTyped(json, false);
}

export function Web3EmplTokenExchangeConversionAndFeeDTOToJSONTyped(value?: Web3EmplTokenExchangeConversionAndFeeDTO | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'empl_amount': value['emplAmount'],
        'empl_tx_fee': value['emplTxFee'],
        'max_amount_out': value['maxAmountOut'],
        'min_amount_out': value['minAmountOut'],
        'token_out_address': value['tokenOutAddress'],
    };
}

