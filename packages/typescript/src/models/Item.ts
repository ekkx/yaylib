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
import type { RealmGift } from './RealmGift';
import {
    RealmGiftFromJSON,
    RealmGiftFromJSONTyped,
    RealmGiftToJSON,
    RealmGiftToJSONTyped,
} from './RealmGift';

/**
 * 
 * @export
 * @interface Item
 */
export interface Item {
    /**
     * 
     * @type {RealmGift}
     * @memberof Item
     */
    gift?: RealmGift | null;
    /**
     * 
     * @type {number}
     * @memberof Item
     */
    quantity?: number | null;
}

/**
 * Check if a given object implements the Item interface.
 */
export function instanceOfItem(value: object): value is Item {
    return true;
}

export function ItemFromJSON(json: any): Item {
    return ItemFromJSONTyped(json, false);
}

export function ItemFromJSONTyped(json: any, ignoreDiscriminator: boolean): Item {
    if (json == null) {
        return json;
    }
    return {
        
        'gift': json['gift'] == null ? undefined : RealmGiftFromJSON(json['gift']),
        'quantity': json['quantity'] == null ? undefined : json['quantity'],
    };
}

export function ItemToJSON(json: any): Item {
    return ItemToJSONTyped(json, false);
}

export function ItemToJSONTyped(value?: Item | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'gift': RealmGiftToJSON(value['gift']),
        'quantity': value['quantity'],
    };
}

