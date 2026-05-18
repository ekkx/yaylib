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
import type { ProfileImageDTO } from './ProfileImageDTO.js';
import {
    ProfileImageDTOFromJSON,
    ProfileImageDTOFromJSONTyped,
    ProfileImageDTOToJSON,
    ProfileImageDTOToJSONTyped,
} from './ProfileImageDTO.js';

/**
 * 
 * @export
 * @interface PalRaceItemDTO
 */
export interface PalRaceItemDTO {
    /**
     * 
     * @type {number}
     * @memberof PalRaceItemDTO
     */
    currentLevel?: number | null;
    /**
     * 
     * @type {string}
     * @memberof PalRaceItemDTO
     */
    fatigueLevel?: string | null;
    /**
     * 
     * @type {string}
     * @memberof PalRaceItemDTO
     */
    nickname?: string | null;
    /**
     * 
     * @type {number}
     * @memberof PalRaceItemDTO
     */
    overallStrength?: number | null;
    /**
     * 
     * @type {ProfileImageDTO}
     * @memberof PalRaceItemDTO
     */
    profileImage?: ProfileImageDTO | null;
    /**
     * 
     * @type {string}
     * @memberof PalRaceItemDTO
     */
    rank?: string | null;
    /**
     * 
     * @type {string}
     * @memberof PalRaceItemDTO
     */
    tokenId?: string | null;
    /**
     * 
     * @type {string}
     * @memberof PalRaceItemDTO
     */
    type?: string | null;
    /**
     * 
     * @type {string}
     * @memberof PalRaceItemDTO
     */
    winRate?: string | null;
    /**
     * 
     * @type {number}
     * @memberof PalRaceItemDTO
     */
    winningStreak?: number | null;
}

/**
 * Check if a given object implements the PalRaceItemDTO interface.
 */
export function instanceOfPalRaceItemDTO(value: object): value is PalRaceItemDTO {
    return true;
}

export function PalRaceItemDTOFromJSON(json: any): PalRaceItemDTO {
    return PalRaceItemDTOFromJSONTyped(json, false);
}

export function PalRaceItemDTOFromJSONTyped(json: any, ignoreDiscriminator: boolean): PalRaceItemDTO {
    if (json == null) {
        return json;
    }
    return {
        
        'currentLevel': json['current_level'] == null ? undefined : json['current_level'],
        'fatigueLevel': json['fatigue_level'] == null ? undefined : json['fatigue_level'],
        'nickname': json['nickname'] == null ? undefined : json['nickname'],
        'overallStrength': json['overall_strength'] == null ? undefined : json['overall_strength'],
        'profileImage': json['profile_image'] == null ? undefined : ProfileImageDTOFromJSON(json['profile_image']),
        'rank': json['rank'] == null ? undefined : json['rank'],
        'tokenId': json['token_id'] == null ? undefined : json['token_id'],
        'type': json['type'] == null ? undefined : json['type'],
        'winRate': json['win_rate'] == null ? undefined : json['win_rate'],
        'winningStreak': json['winning_streak'] == null ? undefined : json['winning_streak'],
    };
}

export function PalRaceItemDTOToJSON(json: any): PalRaceItemDTO {
    return PalRaceItemDTOToJSONTyped(json, false);
}

export function PalRaceItemDTOToJSONTyped(value?: PalRaceItemDTO | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'current_level': value['currentLevel'],
        'fatigue_level': value['fatigueLevel'],
        'nickname': value['nickname'],
        'overall_strength': value['overallStrength'],
        'profile_image': ProfileImageDTOToJSON(value['profileImage']),
        'rank': value['rank'],
        'token_id': value['tokenId'],
        'type': value['type'],
        'win_rate': value['winRate'],
        'winning_streak': value['winningStreak'],
    };
}

