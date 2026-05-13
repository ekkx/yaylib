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
import type { UserRank } from './UserRank';
import {
    UserRankFromJSON,
    UserRankFromJSONTyped,
    UserRankToJSON,
    UserRankToJSONTyped,
} from './UserRank';

/**
 * 
 * @export
 * @interface GroupInCircleUserLeaderboardResponse
 */
export interface GroupInCircleUserLeaderboardResponse {
    /**
     * 
     * @type {Array<UserRank>}
     * @memberof GroupInCircleUserLeaderboardResponse
     */
    userLeaderboard?: Array<UserRank> | null;
}

/**
 * Check if a given object implements the GroupInCircleUserLeaderboardResponse interface.
 */
export function instanceOfGroupInCircleUserLeaderboardResponse(value: object): value is GroupInCircleUserLeaderboardResponse {
    return true;
}

export function GroupInCircleUserLeaderboardResponseFromJSON(json: any): GroupInCircleUserLeaderboardResponse {
    return GroupInCircleUserLeaderboardResponseFromJSONTyped(json, false);
}

export function GroupInCircleUserLeaderboardResponseFromJSONTyped(json: any, ignoreDiscriminator: boolean): GroupInCircleUserLeaderboardResponse {
    if (json == null) {
        return json;
    }
    return {
        
        'userLeaderboard': json['user_leaderboard'] == null ? undefined : ((json['user_leaderboard'] as Array<any>).map(UserRankFromJSON)),
    };
}

export function GroupInCircleUserLeaderboardResponseToJSON(json: any): GroupInCircleUserLeaderboardResponse {
    return GroupInCircleUserLeaderboardResponseToJSONTyped(json, false);
}

export function GroupInCircleUserLeaderboardResponseToJSONTyped(value?: GroupInCircleUserLeaderboardResponse | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'user_leaderboard': value['userLeaderboard'] == null ? undefined : ((value['userLeaderboard'] as Array<any>).map(UserRankToJSON)),
    };
}

