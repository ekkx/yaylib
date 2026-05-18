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
import type { ModelUserRank } from './ModelUserRank.js';
import {
    ModelUserRankFromJSON,
    ModelUserRankFromJSONTyped,
    ModelUserRankToJSON,
    ModelUserRankToJSONTyped,
} from './ModelUserRank.js';

/**
 * 
 * @export
 * @interface ActivityScore
 */
export interface ActivityScore {
    /**
     * 
     * @type {boolean}
     * @memberof ActivityScore
     */
    calculated?: boolean | null;
    /**
     * 
     * @type {ModelUserRank}
     * @memberof ActivityScore
     */
    grade?: ModelUserRank | null;
    /**
     * 
     * @type {number}
     * @memberof ActivityScore
     */
    score?: number | null;
}

/**
 * Check if a given object implements the ActivityScore interface.
 */
export function instanceOfActivityScore(value: object): value is ActivityScore {
    return true;
}

export function ActivityScoreFromJSON(json: any): ActivityScore {
    return ActivityScoreFromJSONTyped(json, false);
}

export function ActivityScoreFromJSONTyped(json: any, ignoreDiscriminator: boolean): ActivityScore {
    if (json == null) {
        return json;
    }
    return {
        
        'calculated': json['calculated'] == null ? undefined : json['calculated'],
        'grade': json['grade'] == null ? undefined : ModelUserRankFromJSON(json['grade']),
        'score': json['score'] == null ? undefined : json['score'],
    };
}

export function ActivityScoreToJSON(json: any): ActivityScore {
    return ActivityScoreToJSONTyped(json, false);
}

export function ActivityScoreToJSONTyped(value?: ActivityScore | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'calculated': value['calculated'],
        'grade': ModelUserRankToJSON(value['grade']),
        'score': value['score'],
    };
}

