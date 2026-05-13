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
/**
 * 
 * @export
 * @interface LastStatusScreenUserRank
 */
export interface LastStatusScreenUserRank {
    /**
     * 
     * @type {number}
     * @memberof LastStatusScreenUserRank
     */
    userId?: number | null;
    /**
     * 
     * @type {string}
     * @memberof LastStatusScreenUserRank
     */
    userRank?: string | null;
}

/**
 * Check if a given object implements the LastStatusScreenUserRank interface.
 */
export function instanceOfLastStatusScreenUserRank(value: object): value is LastStatusScreenUserRank {
    return true;
}

export function LastStatusScreenUserRankFromJSON(json: any): LastStatusScreenUserRank {
    return LastStatusScreenUserRankFromJSONTyped(json, false);
}

export function LastStatusScreenUserRankFromJSONTyped(json: any, ignoreDiscriminator: boolean): LastStatusScreenUserRank {
    if (json == null) {
        return json;
    }
    return {
        
        'userId': json['user_id'] == null ? undefined : json['user_id'],
        'userRank': json['user_rank'] == null ? undefined : json['user_rank'],
    };
}

export function LastStatusScreenUserRankToJSON(json: any): LastStatusScreenUserRank {
    return LastStatusScreenUserRankToJSONTyped(json, false);
}

export function LastStatusScreenUserRankToJSONTyped(value?: LastStatusScreenUserRank | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'user_id': value['userId'],
        'user_rank': value['userRank'],
    };
}

