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
import type { Result } from './Result.js';
import {
    ResultFromJSON,
    ResultFromJSONTyped,
    ResultToJSON,
    ResultToJSONTyped,
} from './Result.js';

/**
 * 
 * @export
 * @interface TransactionDetail
 */
export interface TransactionDetail {
    /**
     * 
     * @type {Result}
     * @memberof TransactionDetail
     */
    result?: Result | null;
}

/**
 * Check if a given object implements the TransactionDetail interface.
 */
export function instanceOfTransactionDetail(value: object): value is TransactionDetail {
    return true;
}

export function TransactionDetailFromJSON(json: any): TransactionDetail {
    return TransactionDetailFromJSONTyped(json, false);
}

export function TransactionDetailFromJSONTyped(json: any, ignoreDiscriminator: boolean): TransactionDetail {
    if (json == null) {
        return json;
    }
    return {
        
        'result': json['result'] == null ? undefined : ResultFromJSON(json['result']),
    };
}

export function TransactionDetailToJSON(json: any): TransactionDetail {
    return TransactionDetailToJSONTyped(json, false);
}

export function TransactionDetailToJSONTyped(value?: TransactionDetail | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'result': ResultToJSON(value['result']),
    };
}

