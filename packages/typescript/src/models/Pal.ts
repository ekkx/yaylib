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
import type { ProfileImage } from './ProfileImage';
import {
    ProfileImageFromJSON,
    ProfileImageFromJSONTyped,
    ProfileImageToJSON,
    ProfileImageToJSONTyped,
} from './ProfileImage';

/**
 * 
 * @export
 * @interface Pal
 */
export interface Pal {
    /**
     * 
     * @type {number}
     * @memberof Pal
     */
    currentLevel?: number | null;
    /**
     * 
     * @type {string}
     * @memberof Pal
     */
    fatigueLevel?: string | null;
    /**
     * 
     * @type {string}
     * @memberof Pal
     */
    nickname?: string | null;
    /**
     * 
     * @type {number}
     * @memberof Pal
     */
    overallStrength?: number | null;
    /**
     * 
     * @type {ProfileImage}
     * @memberof Pal
     */
    profileImage?: ProfileImage | null;
    /**
     * 
     * @type {string}
     * @memberof Pal
     */
    rank?: string | null;
    /**
     * 
     * @type {string}
     * @memberof Pal
     */
    tokenId?: string | null;
    /**
     * 
     * @type {string}
     * @memberof Pal
     */
    tribe?: string | null;
    /**
     * 
     * @type {string}
     * @memberof Pal
     */
    type?: string | null;
    /**
     * 
     * @type {string}
     * @memberof Pal
     */
    winRate?: string | null;
    /**
     * 
     * @type {number}
     * @memberof Pal
     */
    winningStreak?: number | null;
}

/**
 * Check if a given object implements the Pal interface.
 */
export function instanceOfPal(value: object): value is Pal {
    return true;
}

export function PalFromJSON(json: any): Pal {
    return PalFromJSONTyped(json, false);
}

export function PalFromJSONTyped(json: any, ignoreDiscriminator: boolean): Pal {
    if (json == null) {
        return json;
    }
    return {
        
        'currentLevel': json['current_level'] == null ? undefined : json['current_level'],
        'fatigueLevel': json['fatigue_level'] == null ? undefined : json['fatigue_level'],
        'nickname': json['nickname'] == null ? undefined : json['nickname'],
        'overallStrength': json['overall_strength'] == null ? undefined : json['overall_strength'],
        'profileImage': json['profile_image'] == null ? undefined : ProfileImageFromJSON(json['profile_image']),
        'rank': json['rank'] == null ? undefined : json['rank'],
        'tokenId': json['token_id'] == null ? undefined : json['token_id'],
        'tribe': json['tribe'] == null ? undefined : json['tribe'],
        'type': json['type'] == null ? undefined : json['type'],
        'winRate': json['win_rate'] == null ? undefined : json['win_rate'],
        'winningStreak': json['winning_streak'] == null ? undefined : json['winning_streak'],
    };
}

export function PalToJSON(json: any): Pal {
    return PalToJSONTyped(json, false);
}

export function PalToJSONTyped(value?: Pal | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'current_level': value['currentLevel'],
        'fatigue_level': value['fatigueLevel'],
        'nickname': value['nickname'],
        'overall_strength': value['overallStrength'],
        'profile_image': ProfileImageToJSON(value['profileImage']),
        'rank': value['rank'],
        'token_id': value['tokenId'],
        'tribe': value['tribe'],
        'type': value['type'],
        'win_rate': value['winRate'],
        'winning_streak': value['winningStreak'],
    };
}

