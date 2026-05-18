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
import type { LevelUpDetailsPal } from './LevelUpDetailsPal.js';
import {
    LevelUpDetailsPalFromJSON,
    LevelUpDetailsPalFromJSONTyped,
    LevelUpDetailsPalToJSON,
    LevelUpDetailsPalToJSONTyped,
} from './LevelUpDetailsPal.js';
import type { StateChanges } from './StateChanges.js';
import {
    StateChangesFromJSON,
    StateChangesFromJSONTyped,
    StateChangesToJSON,
    StateChangesToJSONTyped,
} from './StateChanges.js';

/**
 * 
 * @export
 * @interface LevelUpDetails
 */
export interface LevelUpDetails {
    /**
     * 
     * @type {boolean}
     * @memberof LevelUpDetails
     */
    available?: boolean | null;
    /**
     * 
     * @type {string}
     * @memberof LevelUpDetails
     */
    cost?: string | null;
    /**
     * 
     * @type {string}
     * @memberof LevelUpDetails
     */
    emplEarnLimit?: string | null;
    /**
     * 
     * @type {number}
     * @memberof LevelUpDetails
     */
    level?: number | null;
    /**
     * 
     * @type {LevelUpDetailsPal}
     * @memberof LevelUpDetails
     */
    pal?: LevelUpDetailsPal | null;
    /**
     * 
     * @type {StateChanges}
     * @memberof LevelUpDetails
     */
    statChanges?: StateChanges | null;
}

/**
 * Check if a given object implements the LevelUpDetails interface.
 */
export function instanceOfLevelUpDetails(value: object): value is LevelUpDetails {
    return true;
}

export function LevelUpDetailsFromJSON(json: any): LevelUpDetails {
    return LevelUpDetailsFromJSONTyped(json, false);
}

export function LevelUpDetailsFromJSONTyped(json: any, ignoreDiscriminator: boolean): LevelUpDetails {
    if (json == null) {
        return json;
    }
    return {
        
        'available': json['available'] == null ? undefined : json['available'],
        'cost': json['cost'] == null ? undefined : json['cost'],
        'emplEarnLimit': json['empl_earn_limit'] == null ? undefined : json['empl_earn_limit'],
        'level': json['level'] == null ? undefined : json['level'],
        'pal': json['pal'] == null ? undefined : LevelUpDetailsPalFromJSON(json['pal']),
        'statChanges': json['stat_changes'] == null ? undefined : StateChangesFromJSON(json['stat_changes']),
    };
}

export function LevelUpDetailsToJSON(json: any): LevelUpDetails {
    return LevelUpDetailsToJSONTyped(json, false);
}

export function LevelUpDetailsToJSONTyped(value?: LevelUpDetails | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'available': value['available'],
        'cost': value['cost'],
        'empl_earn_limit': value['emplEarnLimit'],
        'level': value['level'],
        'pal': LevelUpDetailsPalToJSON(value['pal']),
        'stat_changes': StateChangesToJSON(value['statChanges']),
    };
}

