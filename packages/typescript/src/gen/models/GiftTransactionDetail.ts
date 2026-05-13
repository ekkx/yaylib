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
import type { GiftTransaction } from './GiftTransaction';
import {
    GiftTransactionFromJSON,
    GiftTransactionFromJSONTyped,
    GiftTransactionToJSON,
    GiftTransactionToJSONTyped,
} from './GiftTransaction';

/**
 * 
 * @export
 * @interface GiftTransactionDetail
 */
export interface GiftTransactionDetail {
    /**
     * 
     * @type {Array<GiftTransaction>}
     * @memberof GiftTransactionDetail
     */
    gifts?: Array<GiftTransaction> | null;
}

/**
 * Check if a given object implements the GiftTransactionDetail interface.
 */
export function instanceOfGiftTransactionDetail(value: object): value is GiftTransactionDetail {
    return true;
}

export function GiftTransactionDetailFromJSON(json: any): GiftTransactionDetail {
    return GiftTransactionDetailFromJSONTyped(json, false);
}

export function GiftTransactionDetailFromJSONTyped(json: any, ignoreDiscriminator: boolean): GiftTransactionDetail {
    if (json == null) {
        return json;
    }
    return {
        
        'gifts': json['gifts'] == null ? undefined : ((json['gifts'] as Array<any>).map(GiftTransactionFromJSON)),
    };
}

export function GiftTransactionDetailToJSON(json: any): GiftTransactionDetail {
    return GiftTransactionDetailToJSONTyped(json, false);
}

export function GiftTransactionDetailToJSONTyped(value?: GiftTransactionDetail | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'gifts': value['gifts'] == null ? undefined : ((value['gifts'] as Array<any>).map(GiftTransactionToJSON)),
    };
}

