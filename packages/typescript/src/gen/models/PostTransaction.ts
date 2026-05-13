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
 * @interface PostTransaction
 */
export interface PostTransaction {
    /**
     * 
     * @type {string}
     * @memberof PostTransaction
     */
    receiptJsonStr?: string | null;
}

/**
 * Check if a given object implements the PostTransaction interface.
 */
export function instanceOfPostTransaction(value: object): value is PostTransaction {
    return true;
}

export function PostTransactionFromJSON(json: any): PostTransaction {
    return PostTransactionFromJSONTyped(json, false);
}

export function PostTransactionFromJSONTyped(json: any, ignoreDiscriminator: boolean): PostTransaction {
    if (json == null) {
        return json;
    }
    return {
        
        'receiptJsonStr': json['receipt_json_str'] == null ? undefined : json['receipt_json_str'],
    };
}

export function PostTransactionToJSON(json: any): PostTransaction {
    return PostTransactionToJSONTyped(json, false);
}

export function PostTransactionToJSONTyped(value?: PostTransaction | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'receipt_json_str': value['receiptJsonStr'],
    };
}

