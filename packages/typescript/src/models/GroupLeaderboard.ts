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
import type { Group } from './Group';
import {
    GroupFromJSON,
    GroupFromJSONTyped,
    GroupToJSON,
    GroupToJSONTyped,
} from './Group';

/**
 * 
 * @export
 * @interface GroupLeaderboard
 */
export interface GroupLeaderboard {
    /**
     * 
     * @type {Group}
     * @memberof GroupLeaderboard
     */
    group?: Group | null;
    /**
     * 
     * @type {number}
     * @memberof GroupLeaderboard
     */
    ranking?: number | null;
    /**
     * 
     * @type {number}
     * @memberof GroupLeaderboard
     */
    totalScores?: number | null;
}

/**
 * Check if a given object implements the GroupLeaderboard interface.
 */
export function instanceOfGroupLeaderboard(value: object): value is GroupLeaderboard {
    return true;
}

export function GroupLeaderboardFromJSON(json: any): GroupLeaderboard {
    return GroupLeaderboardFromJSONTyped(json, false);
}

export function GroupLeaderboardFromJSONTyped(json: any, ignoreDiscriminator: boolean): GroupLeaderboard {
    if (json == null) {
        return json;
    }
    return {
        
        'group': json['group'] == null ? undefined : GroupFromJSON(json['group']),
        'ranking': json['ranking'] == null ? undefined : json['ranking'],
        'totalScores': json['total_scores'] == null ? undefined : json['total_scores'],
    };
}

export function GroupLeaderboardToJSON(json: any): GroupLeaderboard {
    return GroupLeaderboardToJSONTyped(json, false);
}

export function GroupLeaderboardToJSONTyped(value?: GroupLeaderboard | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'group': GroupToJSON(value['group']),
        'ranking': value['ranking'],
        'total_scores': value['totalScores'],
    };
}

