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
 * @interface TXNLogData
 */
export interface TXNLogData {
    /**
     * 
     * @type {string}
     * @memberof TXNLogData
     */
    contractAddress?: string | null;
    /**
     * 
     * @type {string}
     * @memberof TXNLogData
     */
    inputsHex?: string | null;
    /**
     * 
     * @type {string}
     * @memberof TXNLogData
     */
    txUniqueId?: string | null;
}

/**
 * Check if a given object implements the TXNLogData interface.
 */
export function instanceOfTXNLogData(value: object): value is TXNLogData {
    return true;
}

export function TXNLogDataFromJSON(json: any): TXNLogData {
    return TXNLogDataFromJSONTyped(json, false);
}

export function TXNLogDataFromJSONTyped(json: any, ignoreDiscriminator: boolean): TXNLogData {
    if (json == null) {
        return json;
    }
    return {
        
        'contractAddress': json['contract_address'] == null ? undefined : json['contract_address'],
        'inputsHex': json['inputs_hex'] == null ? undefined : json['inputs_hex'],
        'txUniqueId': json['tx_unique_id'] == null ? undefined : json['tx_unique_id'],
    };
}

export function TXNLogDataToJSON(json: any): TXNLogData {
    return TXNLogDataToJSONTyped(json, false);
}

export function TXNLogDataToJSONTyped(value?: TXNLogData | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'contract_address': value['contractAddress'],
        'inputs_hex': value['inputsHex'],
        'tx_unique_id': value['txUniqueId'],
    };
}

