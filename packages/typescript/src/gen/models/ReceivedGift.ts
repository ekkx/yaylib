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
import type { Gift } from './Gift.js';
import {
    GiftFromJSON,
    GiftFromJSONTyped,
    GiftToJSON,
    GiftToJSONTyped,
} from './Gift.js';
import type { User } from './User.js';
import {
    UserFromJSON,
    UserFromJSONTyped,
    UserToJSON,
    UserToJSONTyped,
} from './User.js';

/**
 * 
 * @export
 * @interface ReceivedGift
 */
export interface ReceivedGift {
    /**
     * 
     * @type {Gift}
     * @memberof ReceivedGift
     */
    gift?: Gift | null;
    /**
     * 
     * @type {number}
     * @memberof ReceivedGift
     */
    receivedCount?: number | null;
    /**
     * 
     * @type {Array<User>}
     * @memberof ReceivedGift
     */
    senders?: Array<User> | null;
    /**
     * 
     * @type {number}
     * @memberof ReceivedGift
     */
    totalSendersCount?: number | null;
}

/**
 * Check if a given object implements the ReceivedGift interface.
 */
export function instanceOfReceivedGift(value: object): value is ReceivedGift {
    return true;
}

export function ReceivedGiftFromJSON(json: any): ReceivedGift {
    return ReceivedGiftFromJSONTyped(json, false);
}

export function ReceivedGiftFromJSONTyped(json: any, ignoreDiscriminator: boolean): ReceivedGift {
    if (json == null) {
        return json;
    }
    return {
        
        'gift': json['gift'] == null ? undefined : GiftFromJSON(json['gift']),
        'receivedCount': json['received_count'] == null ? undefined : json['received_count'],
        'senders': json['senders'] == null ? undefined : ((json['senders'] as Array<any>).map(UserFromJSON)),
        'totalSendersCount': json['total_senders_count'] == null ? undefined : json['total_senders_count'],
    };
}

export function ReceivedGiftToJSON(json: any): ReceivedGift {
    return ReceivedGiftToJSONTyped(json, false);
}

export function ReceivedGiftToJSONTyped(value?: ReceivedGift | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'gift': GiftToJSON(value['gift']),
        'received_count': value['receivedCount'],
        'senders': value['senders'] == null ? undefined : ((value['senders'] as Array<any>).map(UserToJSON)),
        'total_senders_count': value['totalSendersCount'],
    };
}

