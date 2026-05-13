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
 * @interface Dead
 */
export interface Dead {
    /**
     * 
     * @type {string}
     * @memberof Dead
     */
    palImageUrl?: string | null;
    /**
     * 
     * @type {string}
     * @memberof Dead
     */
    palName?: string | null;
}

/**
 * Check if a given object implements the Dead interface.
 */
export function instanceOfDead(value: object): value is Dead {
    return true;
}

export function DeadFromJSON(json: any): Dead {
    return DeadFromJSONTyped(json, false);
}

export function DeadFromJSONTyped(json: any, ignoreDiscriminator: boolean): Dead {
    if (json == null) {
        return json;
    }
    return {
        
        'palImageUrl': json['pal_image_url'] == null ? undefined : json['pal_image_url'],
        'palName': json['pal_name'] == null ? undefined : json['pal_name'],
    };
}

export function DeadToJSON(json: any): Dead {
    return DeadToJSONTyped(json, false);
}

export function DeadToJSONTyped(value?: Dead | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'pal_image_url': value['palImageUrl'],
        'pal_name': value['palName'],
    };
}

