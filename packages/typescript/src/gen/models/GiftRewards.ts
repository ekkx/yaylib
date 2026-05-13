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
import type { Requirement } from './Requirement';
import {
    RequirementFromJSON,
    RequirementFromJSONTyped,
    RequirementToJSON,
    RequirementToJSONTyped,
} from './Requirement';

/**
 * 
 * @export
 * @interface GiftRewards
 */
export interface GiftRewards {
    /**
     * 
     * @type {number}
     * @memberof GiftRewards
     */
    id?: number | null;
    /**
     * 
     * @type {Array<Requirement>}
     * @memberof GiftRewards
     */
    requirements?: Array<Requirement> | null;
    /**
     * 
     * @type {number}
     * @memberof GiftRewards
     */
    rewardedAmount?: number | null;
}

/**
 * Check if a given object implements the GiftRewards interface.
 */
export function instanceOfGiftRewards(value: object): value is GiftRewards {
    return true;
}

export function GiftRewardsFromJSON(json: any): GiftRewards {
    return GiftRewardsFromJSONTyped(json, false);
}

export function GiftRewardsFromJSONTyped(json: any, ignoreDiscriminator: boolean): GiftRewards {
    if (json == null) {
        return json;
    }
    return {
        
        'id': json['id'] == null ? undefined : json['id'],
        'requirements': json['requirements'] == null ? undefined : ((json['requirements'] as Array<any>).map(RequirementFromJSON)),
        'rewardedAmount': json['rewarded_amount'] == null ? undefined : json['rewarded_amount'],
    };
}

export function GiftRewardsToJSON(json: any): GiftRewards {
    return GiftRewardsToJSONTyped(json, false);
}

export function GiftRewardsToJSONTyped(value?: GiftRewards | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'id': value['id'],
        'requirements': value['requirements'] == null ? undefined : ((value['requirements'] as Array<any>).map(RequirementToJSON)),
        'rewarded_amount': value['rewardedAmount'],
    };
}

