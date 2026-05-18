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
 * @interface GiftRewardGift
 */
export interface GiftRewardGift {
    /**
     * 
     * @type {string}
     * @memberof GiftRewardGift
     */
    icon?: string | null;
    /**
     * 
     * @type {number}
     * @memberof GiftRewardGift
     */
    receivedCount?: number | null;
    /**
     * 
     * @type {number}
     * @memberof GiftRewardGift
     */
    requiredCount?: number | null;
    /**
     * 
     * @type {string}
     * @memberof GiftRewardGift
     */
    title?: string | null;
}

/**
 * Check if a given object implements the GiftRewardGift interface.
 */
export function instanceOfGiftRewardGift(value: object): value is GiftRewardGift {
    return true;
}

export function GiftRewardGiftFromJSON(json: any): GiftRewardGift {
    return GiftRewardGiftFromJSONTyped(json, false);
}

export function GiftRewardGiftFromJSONTyped(json: any, ignoreDiscriminator: boolean): GiftRewardGift {
    if (json == null) {
        return json;
    }
    return {
        
        'icon': json['icon'] == null ? undefined : json['icon'],
        'receivedCount': json['received_count'] == null ? undefined : json['received_count'],
        'requiredCount': json['required_count'] == null ? undefined : json['required_count'],
        'title': json['title'] == null ? undefined : json['title'],
    };
}

export function GiftRewardGiftToJSON(json: any): GiftRewardGift {
    return GiftRewardGiftToJSONTyped(json, false);
}

export function GiftRewardGiftToJSONTyped(value?: GiftRewardGift | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'icon': value['icon'],
        'received_count': value['receivedCount'],
        'required_count': value['requiredCount'],
        'title': value['title'],
    };
}

