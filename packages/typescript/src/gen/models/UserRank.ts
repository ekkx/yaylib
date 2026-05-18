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
import type { RealmUser } from './RealmUser.js';
import {
    RealmUserFromJSON,
    RealmUserFromJSONTyped,
    RealmUserToJSON,
    RealmUserToJSONTyped,
} from './RealmUser.js';

/**
 * 
 * @export
 * @interface UserRank
 */
export interface UserRank {
    /**
     * 
     * @type {number}
     * @memberof UserRank
     */
    ranking?: number | null;
    /**
     * 
     * @type {Array<string>}
     * @memberof UserRank
     */
    topGifts?: Array<string> | null;
    /**
     * 
     * @type {number}
     * @memberof UserRank
     */
    totalScores?: number | null;
    /**
     * 
     * @type {number}
     * @memberof UserRank
     */
    updatedAt?: number | null;
    /**
     * 
     * @type {RealmUser}
     * @memberof UserRank
     */
    user?: RealmUser | null;
}

/**
 * Check if a given object implements the UserRank interface.
 */
export function instanceOfUserRank(value: object): value is UserRank {
    return true;
}

export function UserRankFromJSON(json: any): UserRank {
    return UserRankFromJSONTyped(json, false);
}

export function UserRankFromJSONTyped(json: any, ignoreDiscriminator: boolean): UserRank {
    if (json == null) {
        return json;
    }
    return {
        
        'ranking': json['ranking'] == null ? undefined : json['ranking'],
        'topGifts': json['top_gifts'] == null ? undefined : json['top_gifts'],
        'totalScores': json['total_scores'] == null ? undefined : json['total_scores'],
        'updatedAt': json['updated_at'] == null ? undefined : json['updated_at'],
        'user': json['user'] == null ? undefined : RealmUserFromJSON(json['user']),
    };
}

export function UserRankToJSON(json: any): UserRank {
    return UserRankToJSONTyped(json, false);
}

export function UserRankToJSONTyped(value?: UserRank | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'ranking': value['ranking'],
        'top_gifts': value['topGifts'],
        'total_scores': value['totalScores'],
        'updated_at': value['updatedAt'],
        'user': RealmUserToJSON(value['user']),
    };
}

