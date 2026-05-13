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
import type { GiftRewardGift } from './GiftRewardGift';
import {
    GiftRewardGiftFromJSON,
    GiftRewardGiftFromJSONTyped,
    GiftRewardGiftToJSON,
    GiftRewardGiftToJSONTyped,
} from './GiftRewardGift';

/**
 * 
 * @export
 * @interface GiftReward
 */
export interface GiftReward {
    /**
     * 
     * @type {number}
     * @memberof GiftReward
     */
    emplReward?: number | null;
    /**
     * 
     * @type {Array<GiftRewardGift>}
     * @memberof GiftReward
     */
    gifts?: Array<GiftRewardGift> | null;
    /**
     * 
     * @type {number}
     * @memberof GiftReward
     */
    id?: number | null;
}

/**
 * Check if a given object implements the GiftReward interface.
 */
export function instanceOfGiftReward(value: object): value is GiftReward {
    return true;
}

export function GiftRewardFromJSON(json: any): GiftReward {
    return GiftRewardFromJSONTyped(json, false);
}

export function GiftRewardFromJSONTyped(json: any, ignoreDiscriminator: boolean): GiftReward {
    if (json == null) {
        return json;
    }
    return {
        
        'emplReward': json['empl_reward'] == null ? undefined : json['empl_reward'],
        'gifts': json['gifts'] == null ? undefined : ((json['gifts'] as Array<any>).map(GiftRewardGiftFromJSON)),
        'id': json['id'] == null ? undefined : json['id'],
    };
}

export function GiftRewardToJSON(json: any): GiftReward {
    return GiftRewardToJSONTyped(json, false);
}

export function GiftRewardToJSONTyped(value?: GiftReward | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'empl_reward': value['emplReward'],
        'gifts': value['gifts'] == null ? undefined : ((value['gifts'] as Array<any>).map(GiftRewardGiftToJSON)),
        'id': value['id'],
    };
}

