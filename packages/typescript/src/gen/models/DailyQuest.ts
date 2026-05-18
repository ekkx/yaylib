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
 * @interface DailyQuest
 */
export interface DailyQuest {
    /**
     * 
     * @type {boolean}
     * @memberof DailyQuest
     */
    calculated?: boolean | null;
    /**
     * 
     * @type {number}
     * @memberof DailyQuest
     */
    spentDurationSeconds?: number | null;
    /**
     * 
     * @type {number}
     * @memberof DailyQuest
     */
    targetDurationSeconds?: number | null;
    /**
     * 
     * @type {number}
     * @memberof DailyQuest
     */
    timestamp?: number | null;
}

/**
 * Check if a given object implements the DailyQuest interface.
 */
export function instanceOfDailyQuest(value: object): value is DailyQuest {
    return true;
}

export function DailyQuestFromJSON(json: any): DailyQuest {
    return DailyQuestFromJSONTyped(json, false);
}

export function DailyQuestFromJSONTyped(json: any, ignoreDiscriminator: boolean): DailyQuest {
    if (json == null) {
        return json;
    }
    return {
        
        'calculated': json['calculated'] == null ? undefined : json['calculated'],
        'spentDurationSeconds': json['spent_duration_seconds'] == null ? undefined : json['spent_duration_seconds'],
        'targetDurationSeconds': json['target_duration_seconds'] == null ? undefined : json['target_duration_seconds'],
        'timestamp': json['timestamp'] == null ? undefined : json['timestamp'],
    };
}

export function DailyQuestToJSON(json: any): DailyQuest {
    return DailyQuestToJSONTyped(json, false);
}

export function DailyQuestToJSONTyped(value?: DailyQuest | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'calculated': value['calculated'],
        'spent_duration_seconds': value['spentDurationSeconds'],
        'target_duration_seconds': value['targetDurationSeconds'],
        'timestamp': value['timestamp'],
    };
}

