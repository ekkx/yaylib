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
 * @interface Evolution
 */
export interface Evolution {
    /**
     * 
     * @type {object}
     * @memberof Evolution
     */
    grade?: object | null;
    /**
     * 
     * @type {number}
     * @memberof Evolution
     */
    level?: number | null;
    /**
     * 
     * @type {string}
     * @memberof Evolution
     */
    palImageUrl?: string | null;
    /**
     * 
     * @type {string}
     * @memberof Evolution
     */
    palName?: string | null;
}

/**
 * Check if a given object implements the Evolution interface.
 */
export function instanceOfEvolution(value: object): value is Evolution {
    return true;
}

export function EvolutionFromJSON(json: any): Evolution {
    return EvolutionFromJSONTyped(json, false);
}

export function EvolutionFromJSONTyped(json: any, ignoreDiscriminator: boolean): Evolution {
    if (json == null) {
        return json;
    }
    return {
        
        'grade': json['grade'] == null ? undefined : json['grade'],
        'level': json['level'] == null ? undefined : json['level'],
        'palImageUrl': json['pal_image_url'] == null ? undefined : json['pal_image_url'],
        'palName': json['pal_name'] == null ? undefined : json['pal_name'],
    };
}

export function EvolutionToJSON(json: any): Evolution {
    return EvolutionToJSONTyped(json, false);
}

export function EvolutionToJSONTyped(value?: Evolution | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'grade': value['grade'],
        'level': value['level'],
        'pal_image_url': value['palImageUrl'],
        'pal_name': value['palName'],
    };
}

