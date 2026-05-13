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
 * @interface LevelUp
 */
export interface LevelUp {
    /**
     * 
     * @type {number}
     * @memberof LevelUp
     */
    level?: number | null;
    /**
     * 
     * @type {string}
     * @memberof LevelUp
     */
    palImageUrl?: string | null;
    /**
     * 
     * @type {string}
     * @memberof LevelUp
     */
    palName?: string | null;
    /**
     * 
     * @type {number}
     * @memberof LevelUp
     */
    price?: number | null;
}

/**
 * Check if a given object implements the LevelUp interface.
 */
export function instanceOfLevelUp(value: object): value is LevelUp {
    return true;
}

export function LevelUpFromJSON(json: any): LevelUp {
    return LevelUpFromJSONTyped(json, false);
}

export function LevelUpFromJSONTyped(json: any, ignoreDiscriminator: boolean): LevelUp {
    if (json == null) {
        return json;
    }
    return {
        
        'level': json['level'] == null ? undefined : json['level'],
        'palImageUrl': json['pal_image_url'] == null ? undefined : json['pal_image_url'],
        'palName': json['pal_name'] == null ? undefined : json['pal_name'],
        'price': json['price'] == null ? undefined : json['price'],
    };
}

export function LevelUpToJSON(json: any): LevelUp {
    return LevelUpToJSONTyped(json, false);
}

export function LevelUpToJSONTyped(value?: LevelUp | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'level': value['level'],
        'pal_image_url': value['palImageUrl'],
        'pal_name': value['palName'],
        'price': value['price'],
    };
}

