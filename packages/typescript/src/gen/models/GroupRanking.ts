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
import type { ModelGroup } from './ModelGroup.js';
import {
    ModelGroupFromJSON,
    ModelGroupFromJSONTyped,
    ModelGroupToJSON,
    ModelGroupToJSONTyped,
} from './ModelGroup.js';

/**
 * 
 * @export
 * @interface GroupRanking
 */
export interface GroupRanking {
    /**
     * 
     * @type {ModelGroup}
     * @memberof GroupRanking
     */
    group?: ModelGroup | null;
    /**
     * 
     * @type {number}
     * @memberof GroupRanking
     */
    ranking?: number | null;
    /**
     * 
     * @type {number}
     * @memberof GroupRanking
     */
    totalScores?: number | null;
}

/**
 * Check if a given object implements the GroupRanking interface.
 */
export function instanceOfGroupRanking(value: object): value is GroupRanking {
    return true;
}

export function GroupRankingFromJSON(json: any): GroupRanking {
    return GroupRankingFromJSONTyped(json, false);
}

export function GroupRankingFromJSONTyped(json: any, ignoreDiscriminator: boolean): GroupRanking {
    if (json == null) {
        return json;
    }
    return {
        
        'group': json['group'] == null ? undefined : ModelGroupFromJSON(json['group']),
        'ranking': json['ranking'] == null ? undefined : json['ranking'],
        'totalScores': json['total_scores'] == null ? undefined : json['total_scores'],
    };
}

export function GroupRankingToJSON(json: any): GroupRanking {
    return GroupRankingToJSONTyped(json, false);
}

export function GroupRankingToJSONTyped(value?: GroupRanking | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'group': ModelGroupToJSON(value['group']),
        'ranking': value['ranking'],
        'total_scores': value['totalScores'],
    };
}

