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
 * @interface TransactionGiftReceivedItem
 */
export interface TransactionGiftReceivedItem {
    /**
     * 
     * @type {string}
     * @memberof TransactionGiftReceivedItem
     */
    currency?: string | null;
    /**
     * 
     * @type {number}
     * @memberof TransactionGiftReceivedItem
     */
    id?: number | null;
    /**
     * 
     * @type {number}
     * @memberof TransactionGiftReceivedItem
     */
    price?: number | null;
    /**
     * 
     * @type {string}
     * @memberof TransactionGiftReceivedItem
     */
    title?: string | null;
}

/**
 * Check if a given object implements the TransactionGiftReceivedItem interface.
 */
export function instanceOfTransactionGiftReceivedItem(value: object): value is TransactionGiftReceivedItem {
    return true;
}

export function TransactionGiftReceivedItemFromJSON(json: any): TransactionGiftReceivedItem {
    return TransactionGiftReceivedItemFromJSONTyped(json, false);
}

export function TransactionGiftReceivedItemFromJSONTyped(json: any, ignoreDiscriminator: boolean): TransactionGiftReceivedItem {
    if (json == null) {
        return json;
    }
    return {
        
        'currency': json['currency'] == null ? undefined : json['currency'],
        'id': json['id'] == null ? undefined : json['id'],
        'price': json['price'] == null ? undefined : json['price'],
        'title': json['title'] == null ? undefined : json['title'],
    };
}

export function TransactionGiftReceivedItemToJSON(json: any): TransactionGiftReceivedItem {
    return TransactionGiftReceivedItemToJSONTyped(json, false);
}

export function TransactionGiftReceivedItemToJSONTyped(value?: TransactionGiftReceivedItem | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'currency': value['currency'],
        'id': value['id'],
        'price': value['price'],
        'title': value['title'],
    };
}

