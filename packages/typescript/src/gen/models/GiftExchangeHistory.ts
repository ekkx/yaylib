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
import type { Item } from './Item';
import {
    ItemFromJSON,
    ItemFromJSONTyped,
    ItemToJSON,
    ItemToJSONTyped,
} from './Item';

/**
 * 
 * @export
 * @interface GiftExchangeHistory
 */
export interface GiftExchangeHistory {
    /**
     * 
     * @type {number}
     * @memberof GiftExchangeHistory
     */
    createdAt?: number | null;
    /**
     * 
     * @type {number}
     * @memberof GiftExchangeHistory
     */
    id?: number | null;
    /**
     * 
     * @type {Array<Item>}
     * @memberof GiftExchangeHistory
     */
    items?: Array<Item> | null;
    /**
     * 
     * @type {number}
     * @memberof GiftExchangeHistory
     */
    quantity?: number | null;
    /**
     * 
     * @type {number}
     * @memberof GiftExchangeHistory
     */
    rewardedAmount?: number | null;
    /**
     * 
     * @type {string}
     * @memberof GiftExchangeHistory
     */
    status?: string | null;
    /**
     * 
     * @type {string}
     * @memberof GiftExchangeHistory
     */
    uuid?: string | null;
}

/**
 * Check if a given object implements the GiftExchangeHistory interface.
 */
export function instanceOfGiftExchangeHistory(value: object): value is GiftExchangeHistory {
    return true;
}

export function GiftExchangeHistoryFromJSON(json: any): GiftExchangeHistory {
    return GiftExchangeHistoryFromJSONTyped(json, false);
}

export function GiftExchangeHistoryFromJSONTyped(json: any, ignoreDiscriminator: boolean): GiftExchangeHistory {
    if (json == null) {
        return json;
    }
    return {
        
        'createdAt': json['created_at'] == null ? undefined : json['created_at'],
        'id': json['id'] == null ? undefined : json['id'],
        'items': json['items'] == null ? undefined : ((json['items'] as Array<any>).map(ItemFromJSON)),
        'quantity': json['quantity'] == null ? undefined : json['quantity'],
        'rewardedAmount': json['rewarded_amount'] == null ? undefined : json['rewarded_amount'],
        'status': json['status'] == null ? undefined : json['status'],
        'uuid': json['uuid'] == null ? undefined : json['uuid'],
    };
}

export function GiftExchangeHistoryToJSON(json: any): GiftExchangeHistory {
    return GiftExchangeHistoryToJSONTyped(json, false);
}

export function GiftExchangeHistoryToJSONTyped(value?: GiftExchangeHistory | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'created_at': value['createdAt'],
        'id': value['id'],
        'items': value['items'] == null ? undefined : ((value['items'] as Array<any>).map(ItemToJSON)),
        'quantity': value['quantity'],
        'rewarded_amount': value['rewardedAmount'],
        'status': value['status'],
        'uuid': value['uuid'],
    };
}

