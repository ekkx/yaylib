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
import type { ReceivedGift } from './ReceivedGift.js';
import {
    ReceivedGiftFromJSON,
    ReceivedGiftFromJSONTyped,
    ReceivedGiftToJSON,
    ReceivedGiftToJSONTyped,
} from './ReceivedGift.js';

/**
 * 
 * @export
 * @interface ModelGiftHistory
 */
export interface ModelGiftHistory {
    /**
     * 
     * @type {Gift}
     * @memberof ModelGiftHistory
     */
    giftItem?: Gift | null;
    /**
     * 
     * @type {Array<ReceivedGift>}
     * @memberof ModelGiftHistory
     */
    gifts?: Array<ReceivedGift> | null;
    /**
     * 
     * @type {boolean}
     * @memberof ModelGiftHistory
     */
    isSenderAnonymous?: boolean | null;
    /**
     * 
     * @type {User}
     * @memberof ModelGiftHistory
     */
    receiver?: User | null;
    /**
     * 
     * @type {User}
     * @memberof ModelGiftHistory
     */
    sender?: User | null;
    /**
     * 
     * @type {number}
     * @memberof ModelGiftHistory
     */
    transactionAtSeconds?: number | null;
}

/**
 * Check if a given object implements the ModelGiftHistory interface.
 */
export function instanceOfModelGiftHistory(value: object): value is ModelGiftHistory {
    return true;
}

export function ModelGiftHistoryFromJSON(json: any): ModelGiftHistory {
    return ModelGiftHistoryFromJSONTyped(json, false);
}

export function ModelGiftHistoryFromJSONTyped(json: any, ignoreDiscriminator: boolean): ModelGiftHistory {
    if (json == null) {
        return json;
    }
    return {
        
        'giftItem': json['gift_item'] == null ? undefined : GiftFromJSON(json['gift_item']),
        'gifts': json['gifts'] == null ? undefined : ((json['gifts'] as Array<any>).map(ReceivedGiftFromJSON)),
        'isSenderAnonymous': json['is_sender_anonymous'] == null ? undefined : json['is_sender_anonymous'],
        'receiver': json['receiver'] == null ? undefined : UserFromJSON(json['receiver']),
        'sender': json['sender'] == null ? undefined : UserFromJSON(json['sender']),
        'transactionAtSeconds': json['transaction_at_seconds'] == null ? undefined : json['transaction_at_seconds'],
    };
}

export function ModelGiftHistoryToJSON(json: any): ModelGiftHistory {
    return ModelGiftHistoryToJSONTyped(json, false);
}

export function ModelGiftHistoryToJSONTyped(value?: ModelGiftHistory | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'gift_item': GiftToJSON(value['giftItem']),
        'gifts': value['gifts'] == null ? undefined : ((value['gifts'] as Array<any>).map(ReceivedGiftToJSON)),
        'is_sender_anonymous': value['isSenderAnonymous'],
        'receiver': UserToJSON(value['receiver']),
        'sender': UserToJSON(value['sender']),
        'transaction_at_seconds': value['transactionAtSeconds'],
    };
}

