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
 * @interface TrustScore
 */
export interface TrustScore {
    /**
     * 
     * @type {boolean}
     * @memberof TrustScore
     */
    calculated?: boolean | null;
    /**
     * 
     * @type {ModelUserRank}
     * @memberof TrustScore
     */
    grade?: ModelUserRank | null;
    /**
     * 
     * @type {number}
     * @memberof TrustScore
     */
    score?: number | null;
}

/**
 * Check if a given object implements the TrustScore interface.
 */
export function instanceOfTrustScore(value: object): value is TrustScore {
    return true;
}

export function TrustScoreFromJSON(json: any): TrustScore {
    return TrustScoreFromJSONTyped(json, false);
}

export function TrustScoreFromJSONTyped(json: any, ignoreDiscriminator: boolean): TrustScore {
    if (json == null) {
        return json;
    }
    return {
        
        'calculated': json['calculated'] == null ? undefined : json['calculated'],
        'grade': json['grade'] == null ? undefined : ModelUserRankFromJSON(json['grade']),
        'score': json['score'] == null ? undefined : json['score'],
    };
}

export function TrustScoreToJSON(json: any): TrustScore {
    return TrustScoreToJSONTyped(json, false);
}

export function TrustScoreToJSONTyped(value?: TrustScore | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'calculated': value['calculated'],
        'grade': ModelUserRankToJSON(value['grade']),
        'score': value['score'],
    };
}

