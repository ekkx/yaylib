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
import type { ModelUserRank } from './ModelUserRank';
import {
    ModelUserRankFromJSON,
    ModelUserRankFromJSONTyped,
    ModelUserRankToJSON,
    ModelUserRankToJSONTyped,
} from './ModelUserRank';

/**
 * 
 * @export
 * @interface Web3WalletUserStatusDetails
 */
export interface Web3WalletUserStatusDetails {
    /**
     * 
     * @type {ModelUserRank}
     * @memberof Web3WalletUserStatusDetails
     */
    activityRank?: ModelUserRank | null;
    /**
     * 
     * @type {boolean}
     * @memberof Web3WalletUserStatusDetails
     */
    isRankCalculated?: boolean | null;
    /**
     * 
     * @type {boolean}
     * @memberof Web3WalletUserStatusDetails
     */
    isScoreCalculated?: boolean | null;
    /**
     * 
     * @type {number}
     * @memberof Web3WalletUserStatusDetails
     */
    score?: number | null;
    /**
     * 
     * @type {string}
     * @memberof Web3WalletUserStatusDetails
     */
    scoreMaxFormat?: string | null;
    /**
     * 
     * @type {ModelUserRank}
     * @memberof Web3WalletUserStatusDetails
     */
    userRank?: ModelUserRank | null;
}

/**
 * Check if a given object implements the Web3WalletUserStatusDetails interface.
 */
export function instanceOfWeb3WalletUserStatusDetails(value: object): value is Web3WalletUserStatusDetails {
    return true;
}

export function Web3WalletUserStatusDetailsFromJSON(json: any): Web3WalletUserStatusDetails {
    return Web3WalletUserStatusDetailsFromJSONTyped(json, false);
}

export function Web3WalletUserStatusDetailsFromJSONTyped(json: any, ignoreDiscriminator: boolean): Web3WalletUserStatusDetails {
    if (json == null) {
        return json;
    }
    return {
        
        'activityRank': json['activity_rank'] == null ? undefined : ModelUserRankFromJSON(json['activity_rank']),
        'isRankCalculated': json['is_rank_calculated'] == null ? undefined : json['is_rank_calculated'],
        'isScoreCalculated': json['is_score_calculated'] == null ? undefined : json['is_score_calculated'],
        'score': json['score'] == null ? undefined : json['score'],
        'scoreMaxFormat': json['score_max_format'] == null ? undefined : json['score_max_format'],
        'userRank': json['user_rank'] == null ? undefined : ModelUserRankFromJSON(json['user_rank']),
    };
}

export function Web3WalletUserStatusDetailsToJSON(json: any): Web3WalletUserStatusDetails {
    return Web3WalletUserStatusDetailsToJSONTyped(json, false);
}

export function Web3WalletUserStatusDetailsToJSONTyped(value?: Web3WalletUserStatusDetails | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'activity_rank': ModelUserRankToJSON(value['activityRank']),
        'is_rank_calculated': value['isRankCalculated'],
        'is_score_calculated': value['isScoreCalculated'],
        'score': value['score'],
        'score_max_format': value['scoreMaxFormat'],
        'user_rank': ModelUserRankToJSON(value['userRank']),
    };
}

