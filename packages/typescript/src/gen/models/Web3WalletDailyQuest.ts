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
/**
 * 
 * @export
 * @interface Web3WalletDailyQuest
 */
export interface Web3WalletDailyQuest {
    /**
     * 
     * @type {boolean}
     * @memberof Web3WalletDailyQuest
     */
    isCalculated?: boolean | null;
    /**
     * 
     * @type {number}
     * @memberof Web3WalletDailyQuest
     */
    palCount?: number | null;
    /**
     * 
     * @type {number}
     * @memberof Web3WalletDailyQuest
     */
    spentDurationSeconds?: number | null;
    /**
     * 
     * @type {number}
     * @memberof Web3WalletDailyQuest
     */
    totalDurationNeededSeconds?: number | null;
    /**
     * 
     * @type {number}
     * @memberof Web3WalletDailyQuest
     */
    updatedAtMillis?: number | null;
}

/**
 * Check if a given object implements the Web3WalletDailyQuest interface.
 */
export function instanceOfWeb3WalletDailyQuest(value: object): value is Web3WalletDailyQuest {
    return true;
}

export function Web3WalletDailyQuestFromJSON(json: any): Web3WalletDailyQuest {
    return Web3WalletDailyQuestFromJSONTyped(json, false);
}

export function Web3WalletDailyQuestFromJSONTyped(json: any, ignoreDiscriminator: boolean): Web3WalletDailyQuest {
    if (json == null) {
        return json;
    }
    return {
        
        'isCalculated': json['is_calculated'] == null ? undefined : json['is_calculated'],
        'palCount': json['pal_count'] == null ? undefined : json['pal_count'],
        'spentDurationSeconds': json['spent_duration_seconds'] == null ? undefined : json['spent_duration_seconds'],
        'totalDurationNeededSeconds': json['total_duration_needed_seconds'] == null ? undefined : json['total_duration_needed_seconds'],
        'updatedAtMillis': json['updated_at_millis'] == null ? undefined : json['updated_at_millis'],
    };
}

export function Web3WalletDailyQuestToJSON(json: any): Web3WalletDailyQuest {
    return Web3WalletDailyQuestToJSONTyped(json, false);
}

export function Web3WalletDailyQuestToJSONTyped(value?: Web3WalletDailyQuest | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'is_calculated': value['isCalculated'],
        'pal_count': value['palCount'],
        'spent_duration_seconds': value['spentDurationSeconds'],
        'total_duration_needed_seconds': value['totalDurationNeededSeconds'],
        'updated_at_millis': value['updatedAtMillis'],
    };
}

