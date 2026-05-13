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
import type { Reward } from './Reward';
import {
    RewardFromJSON,
    RewardFromJSONTyped,
    RewardToJSON,
    RewardToJSONTyped,
} from './Reward';

/**
 * 
 * @export
 * @interface RaceReward
 */
export interface RaceReward {
    /**
     * 
     * @type {Reward}
     * @memberof RaceReward
     */
    firstPrize?: Reward | null;
    /**
     * 
     * @type {Reward}
     * @memberof RaceReward
     */
    secondPrize?: Reward | null;
    /**
     * 
     * @type {Reward}
     * @memberof RaceReward
     */
    thirdPrize?: Reward | null;
}

/**
 * Check if a given object implements the RaceReward interface.
 */
export function instanceOfRaceReward(value: object): value is RaceReward {
    return true;
}

export function RaceRewardFromJSON(json: any): RaceReward {
    return RaceRewardFromJSONTyped(json, false);
}

export function RaceRewardFromJSONTyped(json: any, ignoreDiscriminator: boolean): RaceReward {
    if (json == null) {
        return json;
    }
    return {
        
        'firstPrize': json['first_prize'] == null ? undefined : RewardFromJSON(json['first_prize']),
        'secondPrize': json['second_prize'] == null ? undefined : RewardFromJSON(json['second_prize']),
        'thirdPrize': json['third_prize'] == null ? undefined : RewardFromJSON(json['third_prize']),
    };
}

export function RaceRewardToJSON(json: any): RaceReward {
    return RaceRewardToJSONTyped(json, false);
}

export function RaceRewardToJSONTyped(value?: RaceReward | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'first_prize': RewardToJSON(value['firstPrize']),
        'second_prize': RewardToJSON(value['secondPrize']),
        'third_prize': RewardToJSON(value['thirdPrize']),
    };
}

