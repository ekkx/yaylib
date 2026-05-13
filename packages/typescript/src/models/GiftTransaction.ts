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
import type { Gift } from './Gift';
import {
    GiftFromJSON,
    GiftFromJSONTyped,
    GiftToJSON,
    GiftToJSONTyped,
} from './Gift';

/**
 * 
 * @export
 * @interface GiftTransaction
 */
export interface GiftTransaction {
    /**
     * 
     * @type {Gift}
     * @memberof GiftTransaction
     */
    gift?: Gift | null;
    /**
     * 
     * @type {number}
     * @memberof GiftTransaction
     */
    quantity?: number | null;
}

/**
 * Check if a given object implements the GiftTransaction interface.
 */
export function instanceOfGiftTransaction(value: object): value is GiftTransaction {
    return true;
}

export function GiftTransactionFromJSON(json: any): GiftTransaction {
    return GiftTransactionFromJSONTyped(json, false);
}

export function GiftTransactionFromJSONTyped(json: any, ignoreDiscriminator: boolean): GiftTransaction {
    if (json == null) {
        return json;
    }
    return {
        
        'gift': json['gift'] == null ? undefined : GiftFromJSON(json['gift']),
        'quantity': json['quantity'] == null ? undefined : json['quantity'],
    };
}

export function GiftTransactionToJSON(json: any): GiftTransaction {
    return GiftTransactionToJSONTyped(json, false);
}

export function GiftTransactionToJSONTyped(value?: GiftTransaction | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'gift': GiftToJSON(value['gift']),
        'quantity': value['quantity'],
    };
}

