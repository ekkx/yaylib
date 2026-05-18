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
import type { TrustScore } from './TrustScore.js';
import {
    TrustScoreFromJSON,
    TrustScoreFromJSONTyped,
    TrustScoreToJSON,
    TrustScoreToJSONTyped,
} from './TrustScore.js';
import type { ActivityScore } from './ActivityScore.js';
import {
    ActivityScoreFromJSON,
    ActivityScoreFromJSONTyped,
    ActivityScoreToJSON,
    ActivityScoreToJSONTyped,
} from './ActivityScore.js';

/**
 * 
 * @export
 * @interface UserReputations
 */
export interface UserReputations {
    /**
     * 
     * @type {ActivityScore}
     * @memberof UserReputations
     */
    activityScore?: ActivityScore | null;
    /**
     * 
     * @type {TrustScore}
     * @memberof UserReputations
     */
    trustScore?: TrustScore | null;
}

/**
 * Check if a given object implements the UserReputations interface.
 */
export function instanceOfUserReputations(value: object): value is UserReputations {
    return true;
}

export function UserReputationsFromJSON(json: any): UserReputations {
    return UserReputationsFromJSONTyped(json, false);
}

export function UserReputationsFromJSONTyped(json: any, ignoreDiscriminator: boolean): UserReputations {
    if (json == null) {
        return json;
    }
    return {
        
        'activityScore': json['activity_score'] == null ? undefined : ActivityScoreFromJSON(json['activity_score']),
        'trustScore': json['trust_score'] == null ? undefined : TrustScoreFromJSON(json['trust_score']),
    };
}

export function UserReputationsToJSON(json: any): UserReputations {
    return UserReputationsToJSONTyped(json, false);
}

export function UserReputationsToJSONTyped(value?: UserReputations | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'activity_score': ActivityScoreToJSON(value['activityScore']),
        'trust_score': TrustScoreToJSON(value['trustScore']),
    };
}

