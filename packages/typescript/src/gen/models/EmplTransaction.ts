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
/**
 * 
 * @export
 * @interface EmplTransaction
 */
export interface EmplTransaction {
    /**
     * 
     * @type {string}
     * @memberof EmplTransaction
     */
    amount?: string | null;
    /**
     * 
     * @type {number}
     * @memberof EmplTransaction
     */
    id?: number | null;
    /**
     * 
     * @type {number}
     * @memberof EmplTransaction
     */
    occurredAt?: number | null;
    /**
     * 
     * @type {string}
     * @memberof EmplTransaction
     */
    type?: string | null;
}

/**
 * Check if a given object implements the EmplTransaction interface.
 */
export function instanceOfEmplTransaction(value: object): value is EmplTransaction {
    return true;
}

export function EmplTransactionFromJSON(json: any): EmplTransaction {
    return EmplTransactionFromJSONTyped(json, false);
}

export function EmplTransactionFromJSONTyped(json: any, ignoreDiscriminator: boolean): EmplTransaction {
    if (json == null) {
        return json;
    }
    return {
        
        'amount': json['amount'] == null ? undefined : json['amount'],
        'id': json['id'] == null ? undefined : json['id'],
        'occurredAt': json['occurred_at'] == null ? undefined : json['occurred_at'],
        'type': json['type'] == null ? undefined : json['type'],
    };
}

export function EmplTransactionToJSON(json: any): EmplTransaction {
    return EmplTransactionToJSONTyped(json, false);
}

export function EmplTransactionToJSONTyped(value?: EmplTransaction | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'amount': value['amount'],
        'id': value['id'],
        'occurred_at': value['occurredAt'],
        'type': value['type'],
    };
}

