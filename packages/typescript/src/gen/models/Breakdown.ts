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
 * @interface Breakdown
 */
export interface Breakdown {
    /**
     * 
     * @type {number}
     * @memberof Breakdown
     */
    activity?: number | null;
    /**
     * 
     * @type {number}
     * @memberof Breakdown
     */
    originalPoints?: number | null;
    /**
     * 
     * @type {number}
     * @memberof Breakdown
     */
    specialMission?: number | null;
    /**
     * 
     * @type {number}
     * @memberof Breakdown
     */
    vipMultiplier?: number | null;
}

/**
 * Check if a given object implements the Breakdown interface.
 */
export function instanceOfBreakdown(value: object): value is Breakdown {
    return true;
}

export function BreakdownFromJSON(json: any): Breakdown {
    return BreakdownFromJSONTyped(json, false);
}

export function BreakdownFromJSONTyped(json: any, ignoreDiscriminator: boolean): Breakdown {
    if (json == null) {
        return json;
    }
    return {
        
        'activity': json['activity'] == null ? undefined : json['activity'],
        'originalPoints': json['original_points'] == null ? undefined : json['original_points'],
        'specialMission': json['special_mission'] == null ? undefined : json['special_mission'],
        'vipMultiplier': json['vip_multiplier'] == null ? undefined : json['vip_multiplier'],
    };
}

export function BreakdownToJSON(json: any): Breakdown {
    return BreakdownToJSONTyped(json, false);
}

export function BreakdownToJSONTyped(value?: Breakdown | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'activity': value['activity'],
        'original_points': value['originalPoints'],
        'special_mission': value['specialMission'],
        'vip_multiplier': value['vipMultiplier'],
    };
}

