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
 * @interface Hatch
 */
export interface Hatch {
    /**
     * 
     * @type {string}
     * @memberof Hatch
     */
    palImageUrl?: string | null;
    /**
     * 
     * @type {string}
     * @memberof Hatch
     */
    palName?: string | null;
}

/**
 * Check if a given object implements the Hatch interface.
 */
export function instanceOfHatch(value: object): value is Hatch {
    return true;
}

export function HatchFromJSON(json: any): Hatch {
    return HatchFromJSONTyped(json, false);
}

export function HatchFromJSONTyped(json: any, ignoreDiscriminator: boolean): Hatch {
    if (json == null) {
        return json;
    }
    return {
        
        'palImageUrl': json['pal_image_url'] == null ? undefined : json['pal_image_url'],
        'palName': json['pal_name'] == null ? undefined : json['pal_name'],
    };
}

export function HatchToJSON(json: any): Hatch {
    return HatchToJSONTyped(json, false);
}

export function HatchToJSONTyped(value?: Hatch | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'pal_image_url': value['palImageUrl'],
        'pal_name': value['palName'],
    };
}

