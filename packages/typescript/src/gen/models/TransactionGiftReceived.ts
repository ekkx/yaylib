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
import type { TransactionGiftReceivedItem } from './TransactionGiftReceivedItem.js';
import {
    TransactionGiftReceivedItemFromJSON,
    TransactionGiftReceivedItemFromJSONTyped,
    TransactionGiftReceivedItemToJSON,
    TransactionGiftReceivedItemToJSONTyped,
} from './TransactionGiftReceivedItem.js';

/**
 * 
 * @export
 * @interface TransactionGiftReceived
 */
export interface TransactionGiftReceived {
    /**
     * 
     * @type {string}
     * @memberof TransactionGiftReceived
     */
    icon?: string | null;
    /**
     * 
     * @type {TransactionGiftReceivedItem}
     * @memberof TransactionGiftReceived
     */
    item?: TransactionGiftReceivedItem | null;
    /**
     * 
     * @type {number}
     * @memberof TransactionGiftReceived
     */
    itemId?: number | null;
    /**
     * 
     * @type {number}
     * @memberof TransactionGiftReceived
     */
    quantity?: number | null;
}

/**
 * Check if a given object implements the TransactionGiftReceived interface.
 */
export function instanceOfTransactionGiftReceived(value: object): value is TransactionGiftReceived {
    return true;
}

export function TransactionGiftReceivedFromJSON(json: any): TransactionGiftReceived {
    return TransactionGiftReceivedFromJSONTyped(json, false);
}

export function TransactionGiftReceivedFromJSONTyped(json: any, ignoreDiscriminator: boolean): TransactionGiftReceived {
    if (json == null) {
        return json;
    }
    return {
        
        'icon': json['icon'] == null ? undefined : json['icon'],
        'item': json['item'] == null ? undefined : TransactionGiftReceivedItemFromJSON(json['item']),
        'itemId': json['item_id'] == null ? undefined : json['item_id'],
        'quantity': json['quantity'] == null ? undefined : json['quantity'],
    };
}

export function TransactionGiftReceivedToJSON(json: any): TransactionGiftReceived {
    return TransactionGiftReceivedToJSONTyped(json, false);
}

export function TransactionGiftReceivedToJSONTyped(value?: TransactionGiftReceived | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'icon': value['icon'],
        'item': TransactionGiftReceivedItemToJSON(value['item']),
        'item_id': value['itemId'],
        'quantity': value['quantity'],
    };
}

