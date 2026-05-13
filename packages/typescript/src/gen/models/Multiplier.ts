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
 * @interface Multiplier
 */
export interface Multiplier {
    /**
     * 
     * @type {string}
     * @memberof Multiplier
     */
    activity?: string | null;
    /**
     * 
     * @type {string}
     * @memberof Multiplier
     */
    base?: string | null;
    /**
     * 
     * @type {string}
     * @memberof Multiplier
     */
    date?: string | null;
    /**
     * 
     * @type {string}
     * @memberof Multiplier
     */
    mission?: string | null;
    /**
     * 
     * @type {string}
     * @memberof Multiplier
     */
    multiplier?: string | null;
    /**
     * 
     * @type {string}
     * @memberof Multiplier
     */
    vip?: string | null;
}

/**
 * Check if a given object implements the Multiplier interface.
 */
export function instanceOfMultiplier(value: object): value is Multiplier {
    return true;
}

export function MultiplierFromJSON(json: any): Multiplier {
    return MultiplierFromJSONTyped(json, false);
}

export function MultiplierFromJSONTyped(json: any, ignoreDiscriminator: boolean): Multiplier {
    if (json == null) {
        return json;
    }
    return {
        
        'activity': json['activity'] == null ? undefined : json['activity'],
        'base': json['base'] == null ? undefined : json['base'],
        'date': json['date'] == null ? undefined : json['date'],
        'mission': json['mission'] == null ? undefined : json['mission'],
        'multiplier': json['multiplier'] == null ? undefined : json['multiplier'],
        'vip': json['vip'] == null ? undefined : json['vip'],
    };
}

export function MultiplierToJSON(json: any): Multiplier {
    return MultiplierToJSONTyped(json, false);
}

export function MultiplierToJSONTyped(value?: Multiplier | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'activity': value['activity'],
        'base': value['base'],
        'date': value['date'],
        'mission': value['mission'],
        'multiplier': value['multiplier'],
        'vip': value['vip'],
    };
}

