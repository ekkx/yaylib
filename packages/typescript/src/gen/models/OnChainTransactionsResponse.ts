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
import type { OnChainTransactionDTO } from './OnChainTransactionDTO';
import {
    OnChainTransactionDTOFromJSON,
    OnChainTransactionDTOFromJSONTyped,
    OnChainTransactionDTOToJSON,
    OnChainTransactionDTOToJSONTyped,
} from './OnChainTransactionDTO';

/**
 * 
 * @export
 * @interface OnChainTransactionsResponse
 */
export interface OnChainTransactionsResponse {
    /**
     * 
     * @type {string}
     * @memberof OnChainTransactionsResponse
     */
    cursor?: string | null;
    /**
     * 
     * @type {Array<OnChainTransactionDTO>}
     * @memberof OnChainTransactionsResponse
     */
    result?: Array<OnChainTransactionDTO> | null;
}

/**
 * Check if a given object implements the OnChainTransactionsResponse interface.
 */
export function instanceOfOnChainTransactionsResponse(value: object): value is OnChainTransactionsResponse {
    return true;
}

export function OnChainTransactionsResponseFromJSON(json: any): OnChainTransactionsResponse {
    return OnChainTransactionsResponseFromJSONTyped(json, false);
}

export function OnChainTransactionsResponseFromJSONTyped(json: any, ignoreDiscriminator: boolean): OnChainTransactionsResponse {
    if (json == null) {
        return json;
    }
    return {
        
        'cursor': json['cursor'] == null ? undefined : json['cursor'],
        'result': json['result'] == null ? undefined : ((json['result'] as Array<any>).map(OnChainTransactionDTOFromJSON)),
    };
}

export function OnChainTransactionsResponseToJSON(json: any): OnChainTransactionsResponse {
    return OnChainTransactionsResponseToJSONTyped(json, false);
}

export function OnChainTransactionsResponseToJSONTyped(value?: OnChainTransactionsResponse | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'cursor': value['cursor'],
        'result': value['result'] == null ? undefined : ((value['result'] as Array<any>).map(OnChainTransactionDTOToJSON)),
    };
}

