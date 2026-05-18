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
import type { User } from './User.js';
import {
    UserFromJSON,
    UserFromJSONTyped,
    UserToJSON,
    UserToJSONTyped,
} from './User.js';

/**
 * 
 * @export
 * @interface GroupInCircleRanking
 */
export interface GroupInCircleRanking {
    /**
     * 
     * @type {Array<string>}
     * @memberof GroupInCircleRanking
     */
    gifts?: Array<string> | null;
    /**
     * 
     * @type {User}
     * @memberof GroupInCircleRanking
     */
    user?: User | null;
}

/**
 * Check if a given object implements the GroupInCircleRanking interface.
 */
export function instanceOfGroupInCircleRanking(value: object): value is GroupInCircleRanking {
    return true;
}

export function GroupInCircleRankingFromJSON(json: any): GroupInCircleRanking {
    return GroupInCircleRankingFromJSONTyped(json, false);
}

export function GroupInCircleRankingFromJSONTyped(json: any, ignoreDiscriminator: boolean): GroupInCircleRanking {
    if (json == null) {
        return json;
    }
    return {
        
        'gifts': json['gifts'] == null ? undefined : json['gifts'],
        'user': json['user'] == null ? undefined : UserFromJSON(json['user']),
    };
}

export function GroupInCircleRankingToJSON(json: any): GroupInCircleRanking {
    return GroupInCircleRankingToJSONTyped(json, false);
}

export function GroupInCircleRankingToJSONTyped(value?: GroupInCircleRanking | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'gifts': value['gifts'],
        'user': UserToJSON(value['user']),
    };
}

