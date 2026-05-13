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
 * @interface TXNData
 */
export interface TXNData {
    /**
     * 
     * @type {string}
     * @memberof TXNData
     */
    contractAddress?: string | null;
    /**
     * 
     * @type {string}
     * @memberof TXNData
     */
    inputsHex?: string | null;
}

/**
 * Check if a given object implements the TXNData interface.
 */
export function instanceOfTXNData(value: object): value is TXNData {
    return true;
}

export function TXNDataFromJSON(json: any): TXNData {
    return TXNDataFromJSONTyped(json, false);
}

export function TXNDataFromJSONTyped(json: any, ignoreDiscriminator: boolean): TXNData {
    if (json == null) {
        return json;
    }
    return {
        
        'contractAddress': json['contract_address'] == null ? undefined : json['contract_address'],
        'inputsHex': json['inputs_hex'] == null ? undefined : json['inputs_hex'],
    };
}

export function TXNDataToJSON(json: any): TXNData {
    return TXNDataToJSONTyped(json, false);
}

export function TXNDataToJSONTyped(value?: TXNData | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'contract_address': value['contractAddress'],
        'inputs_hex': value['inputsHex'],
    };
}

