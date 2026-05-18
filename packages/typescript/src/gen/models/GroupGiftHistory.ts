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
 * @interface GroupGiftHistory
 */
export interface GroupGiftHistory {
    /**
     * 
     * @type {Array<GiftCount>}
     * @memberof GroupGiftHistory
     */
    giftsCount?: Array<GiftCount> | null;
    /**
     * 
     * @type {number}
     * @memberof GroupGiftHistory
     */
    receivedDate?: number | null;
    /**
     * 
     * @type {RealmUser}
     * @memberof GroupGiftHistory
     */
    user?: RealmUser | null;
}

/**
 * Check if a given object implements the GroupGiftHistory interface.
 */
export function instanceOfGroupGiftHistory(value: object): value is GroupGiftHistory {
    return true;
}

export function GroupGiftHistoryFromJSON(json: any): GroupGiftHistory {
    return GroupGiftHistoryFromJSONTyped(json, false);
}

export function GroupGiftHistoryFromJSONTyped(json: any, ignoreDiscriminator: boolean): GroupGiftHistory {
    if (json == null) {
        return json;
    }
    return {
        
        'giftsCount': json['gifts_count'] == null ? undefined : ((json['gifts_count'] as Array<any>).map(GiftCountFromJSON)),
        'receivedDate': json['received_date'] == null ? undefined : json['received_date'],
        'user': json['user'] == null ? undefined : RealmUserFromJSON(json['user']),
    };
}

export function GroupGiftHistoryToJSON(json: any): GroupGiftHistory {
    return GroupGiftHistoryToJSONTyped(json, false);
}

export function GroupGiftHistoryToJSONTyped(value?: GroupGiftHistory | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'gifts_count': value['giftsCount'] == null ? undefined : ((value['giftsCount'] as Array<any>).map(GiftCountToJSON)),
        'received_date': value['receivedDate'],
        'user': RealmUserToJSON(value['user']),
    };
}

