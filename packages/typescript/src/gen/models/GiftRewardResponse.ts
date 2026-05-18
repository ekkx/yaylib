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
import type { GiftRewards } from './GiftRewards.js';
import {
    GiftRewardsFromJSON,
    GiftRewardsFromJSONTyped,
    GiftRewardsToJSON,
    GiftRewardsToJSONTyped,
} from './GiftRewards.js';

/**
 * 
 * @export
 * @interface GiftRewardResponse
 */
export interface GiftRewardResponse {
    /**
     * 
     * @type {Array<GiftRewards>}
     * @memberof GiftRewardResponse
     */
    giftRewards?: Array<GiftRewards> | null;
}

/**
 * Check if a given object implements the GiftRewardResponse interface.
 */
export function instanceOfGiftRewardResponse(value: object): value is GiftRewardResponse {
    return true;
}

export function GiftRewardResponseFromJSON(json: any): GiftRewardResponse {
    return GiftRewardResponseFromJSONTyped(json, false);
}

export function GiftRewardResponseFromJSONTyped(json: any, ignoreDiscriminator: boolean): GiftRewardResponse {
    if (json == null) {
        return json;
    }
    return {
        
        'giftRewards': json['gift_rewards'] == null ? undefined : ((json['gift_rewards'] as Array<any>).map(GiftRewardsFromJSON)),
    };
}

export function GiftRewardResponseToJSON(json: any): GiftRewardResponse {
    return GiftRewardResponseToJSONTyped(json, false);
}

export function GiftRewardResponseToJSONTyped(value?: GiftRewardResponse | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'gift_rewards': value['giftRewards'] == null ? undefined : ((value['giftRewards'] as Array<any>).map(GiftRewardsToJSON)),
    };
}

