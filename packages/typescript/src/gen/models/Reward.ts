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
 * @interface Reward
 */
export interface Reward {
    /**
     * 
     * @type {number}
     * @memberof Reward
     */
    rewardedEmpl?: number | null;
}

/**
 * Check if a given object implements the Reward interface.
 */
export function instanceOfReward(value: object): value is Reward {
    return true;
}

export function RewardFromJSON(json: any): Reward {
    return RewardFromJSONTyped(json, false);
}

export function RewardFromJSONTyped(json: any, ignoreDiscriminator: boolean): Reward {
    if (json == null) {
        return json;
    }
    return {
        
        'rewardedEmpl': json['rewarded_empl'] == null ? undefined : json['rewarded_empl'],
    };
}

export function RewardToJSON(json: any): Reward {
    return RewardToJSONTyped(json, false);
}

export function RewardToJSONTyped(value?: Reward | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'rewarded_empl': value['rewardedEmpl'],
    };
}

