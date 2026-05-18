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
 * @interface MultiplierBreakdown
 */
export interface MultiplierBreakdown {
    /**
     * 
     * @type {number}
     * @memberof MultiplierBreakdown
     */
    activity?: number | null;
    /**
     * 
     * @type {number}
     * @memberof MultiplierBreakdown
     */
    specialMission?: number | null;
    /**
     * 
     * @type {number}
     * @memberof MultiplierBreakdown
     */
    vipMultiplier?: number | null;
}

/**
 * Check if a given object implements the MultiplierBreakdown interface.
 */
export function instanceOfMultiplierBreakdown(value: object): value is MultiplierBreakdown {
    return true;
}

export function MultiplierBreakdownFromJSON(json: any): MultiplierBreakdown {
    return MultiplierBreakdownFromJSONTyped(json, false);
}

export function MultiplierBreakdownFromJSONTyped(json: any, ignoreDiscriminator: boolean): MultiplierBreakdown {
    if (json == null) {
        return json;
    }
    return {
        
        'activity': json['activity'] == null ? undefined : json['activity'],
        'specialMission': json['special_mission'] == null ? undefined : json['special_mission'],
        'vipMultiplier': json['vip_multiplier'] == null ? undefined : json['vip_multiplier'],
    };
}

export function MultiplierBreakdownToJSON(json: any): MultiplierBreakdown {
    return MultiplierBreakdownToJSONTyped(json, false);
}

export function MultiplierBreakdownToJSONTyped(value?: MultiplierBreakdown | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'activity': value['activity'],
        'special_mission': value['specialMission'],
        'vip_multiplier': value['vipMultiplier'],
    };
}

