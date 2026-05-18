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
import type { GiftSlugItem } from './GiftSlugItem.js';
import {
    GiftSlugItemFromJSON,
    GiftSlugItemFromJSONTyped,
    GiftSlugItemToJSON,
    GiftSlugItemToJSONTyped,
} from './GiftSlugItem.js';
import type { GiftCount } from './GiftCount.js';
import {
    GiftCountFromJSON,
    GiftCountFromJSONTyped,
    GiftCountToJSON,
    GiftCountToJSONTyped,
} from './GiftCount.js';
import type { RealmUser } from './RealmUser.js';
import {
    RealmUserFromJSON,
    RealmUserFromJSONTyped,
    RealmUserToJSON,
    RealmUserToJSONTyped,
} from './RealmUser.js';

/**
 * 
 * @export
 * @interface GiftHistory
 */
export interface GiftHistory {
    /**
     * 
     * @type {boolean}
     * @memberof GiftHistory
     */
    anonymous?: boolean | null;
    /**
     * 
     * @type {GiftSlugItem}
     * @memberof GiftHistory
     */
    giftItem?: GiftSlugItem | null;
    /**
     * 
     * @type {Array<GiftCount>}
     * @memberof GiftHistory
     */
    giftsCount?: Array<GiftCount> | null;
    /**
     * 
     * @type {RealmUser}
     * @memberof GiftHistory
     */
    receiver?: RealmUser | null;
    /**
     * 
     * @type {RealmUser}
     * @memberof GiftHistory
     */
    sender?: RealmUser | null;
    /**
     * 
     * @type {number}
     * @memberof GiftHistory
     */
    sentAt?: number | null;
}

/**
 * Check if a given object implements the GiftHistory interface.
 */
export function instanceOfGiftHistory(value: object): value is GiftHistory {
    return true;
}

export function GiftHistoryFromJSON(json: any): GiftHistory {
    return GiftHistoryFromJSONTyped(json, false);
}

export function GiftHistoryFromJSONTyped(json: any, ignoreDiscriminator: boolean): GiftHistory {
    if (json == null) {
        return json;
    }
    return {
        
        'anonymous': json['anonymous'] == null ? undefined : json['anonymous'],
        'giftItem': json['gift_item'] == null ? undefined : GiftSlugItemFromJSON(json['gift_item']),
        'giftsCount': json['gifts_count'] == null ? undefined : ((json['gifts_count'] as Array<any>).map(GiftCountFromJSON)),
        'receiver': json['receiver'] == null ? undefined : RealmUserFromJSON(json['receiver']),
        'sender': json['sender'] == null ? undefined : RealmUserFromJSON(json['sender']),
        'sentAt': json['sent_at'] == null ? undefined : json['sent_at'],
    };
}

export function GiftHistoryToJSON(json: any): GiftHistory {
    return GiftHistoryToJSONTyped(json, false);
}

export function GiftHistoryToJSONTyped(value?: GiftHistory | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'anonymous': value['anonymous'],
        'gift_item': GiftSlugItemToJSON(value['giftItem']),
        'gifts_count': value['giftsCount'] == null ? undefined : ((value['giftsCount'] as Array<any>).map(GiftCountToJSON)),
        'receiver': RealmUserToJSON(value['receiver']),
        'sender': RealmUserToJSON(value['sender']),
        'sent_at': value['sentAt'],
    };
}

