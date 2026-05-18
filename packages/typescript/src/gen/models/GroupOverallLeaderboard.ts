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
import type { GroupRanking } from './GroupRanking.js';
import {
    GroupRankingFromJSON,
    GroupRankingFromJSONTyped,
    GroupRankingToJSON,
    GroupRankingToJSONTyped,
} from './GroupRanking.js';

/**
 * 
 * @export
 * @interface GroupOverallLeaderboard
 */
export interface GroupOverallLeaderboard {
    /**
     * 
     * @type {Array<GroupRanking>}
     * @memberof GroupOverallLeaderboard
     */
    groupRankings?: Array<GroupRanking> | null;
    /**
     * 
     * @type {number}
     * @memberof GroupOverallLeaderboard
     */
    updatedAtMillis?: number | null;
}

/**
 * Check if a given object implements the GroupOverallLeaderboard interface.
 */
export function instanceOfGroupOverallLeaderboard(value: object): value is GroupOverallLeaderboard {
    return true;
}

export function GroupOverallLeaderboardFromJSON(json: any): GroupOverallLeaderboard {
    return GroupOverallLeaderboardFromJSONTyped(json, false);
}

export function GroupOverallLeaderboardFromJSONTyped(json: any, ignoreDiscriminator: boolean): GroupOverallLeaderboard {
    if (json == null) {
        return json;
    }
    return {
        
        'groupRankings': json['group_rankings'] == null ? undefined : ((json['group_rankings'] as Array<any>).map(GroupRankingFromJSON)),
        'updatedAtMillis': json['updated_at_millis'] == null ? undefined : json['updated_at_millis'],
    };
}

export function GroupOverallLeaderboardToJSON(json: any): GroupOverallLeaderboard {
    return GroupOverallLeaderboardToJSONTyped(json, false);
}

export function GroupOverallLeaderboardToJSONTyped(value?: GroupOverallLeaderboard | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'group_rankings': value['groupRankings'] == null ? undefined : ((value['groupRankings'] as Array<any>).map(GroupRankingToJSON)),
        'updated_at_millis': value['updatedAtMillis'],
    };
}

