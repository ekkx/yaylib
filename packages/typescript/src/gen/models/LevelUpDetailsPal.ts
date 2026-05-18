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
 * @interface LevelUpDetailsPal
 */
export interface LevelUpDetailsPal {
    /**
     * 
     * @type {string}
     * @memberof LevelUpDetailsPal
     */
    image?: string | null;
}

/**
 * Check if a given object implements the LevelUpDetailsPal interface.
 */
export function instanceOfLevelUpDetailsPal(value: object): value is LevelUpDetailsPal {
    return true;
}

export function LevelUpDetailsPalFromJSON(json: any): LevelUpDetailsPal {
    return LevelUpDetailsPalFromJSONTyped(json, false);
}

export function LevelUpDetailsPalFromJSONTyped(json: any, ignoreDiscriminator: boolean): LevelUpDetailsPal {
    if (json == null) {
        return json;
    }
    return {
        
        'image': json['image'] == null ? undefined : json['image'],
    };
}

export function LevelUpDetailsPalToJSON(json: any): LevelUpDetailsPal {
    return LevelUpDetailsPalToJSONTyped(json, false);
}

export function LevelUpDetailsPalToJSONTyped(value?: LevelUpDetailsPal | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'image': value['image'],
    };
}

