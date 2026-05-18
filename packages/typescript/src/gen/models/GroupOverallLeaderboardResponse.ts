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
import type { GroupLeaderboard } from './GroupLeaderboard.js';
import {
    GroupLeaderboardFromJSON,
    GroupLeaderboardFromJSONTyped,
    GroupLeaderboardToJSON,
    GroupLeaderboardToJSONTyped,
} from './GroupLeaderboard.js';

/**
 * 
 * @export
 * @interface GroupOverallLeaderboardResponse
 */
export interface GroupOverallLeaderboardResponse {
    /**
     * 
     * @type {Array<GroupLeaderboard>}
     * @memberof GroupOverallLeaderboardResponse
     */
    groupLeaderboard?: Array<GroupLeaderboard> | null;
    /**
     * 
     * @type {number}
     * @memberof GroupOverallLeaderboardResponse
     */
    updatedAt?: number | null;
}

/**
 * Check if a given object implements the GroupOverallLeaderboardResponse interface.
 */
export function instanceOfGroupOverallLeaderboardResponse(value: object): value is GroupOverallLeaderboardResponse {
    return true;
}

export function GroupOverallLeaderboardResponseFromJSON(json: any): GroupOverallLeaderboardResponse {
    return GroupOverallLeaderboardResponseFromJSONTyped(json, false);
}

export function GroupOverallLeaderboardResponseFromJSONTyped(json: any, ignoreDiscriminator: boolean): GroupOverallLeaderboardResponse {
    if (json == null) {
        return json;
    }
    return {
        
        'groupLeaderboard': json['group_leaderboard'] == null ? undefined : ((json['group_leaderboard'] as Array<any>).map(GroupLeaderboardFromJSON)),
        'updatedAt': json['updated_at'] == null ? undefined : json['updated_at'],
    };
}

export function GroupOverallLeaderboardResponseToJSON(json: any): GroupOverallLeaderboardResponse {
    return GroupOverallLeaderboardResponseToJSONTyped(json, false);
}

export function GroupOverallLeaderboardResponseToJSONTyped(value?: GroupOverallLeaderboardResponse | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'group_leaderboard': value['groupLeaderboard'] == null ? undefined : ((value['groupLeaderboard'] as Array<any>).map(GroupLeaderboardToJSON)),
        'updated_at': value['updatedAt'],
    };
}

