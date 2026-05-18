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
 * @interface ModelUserRank
 */
export interface ModelUserRank {
    /**
     * 
     * @type {number}
     * @memberof ModelUserRank
     */
    disabledIconId?: number | null;
    /**
     * 
     * @type {number}
     * @memberof ModelUserRank
     */
    enabledIconId?: number | null;
    /**
     * 
     * @type {string}
     * @memberof ModelUserRank
     */
    value?: string | null;
}

/**
 * Check if a given object implements the ModelUserRank interface.
 */
export function instanceOfModelUserRank(value: object): value is ModelUserRank {
    return true;
}

export function ModelUserRankFromJSON(json: any): ModelUserRank {
    return ModelUserRankFromJSONTyped(json, false);
}

export function ModelUserRankFromJSONTyped(json: any, ignoreDiscriminator: boolean): ModelUserRank {
    if (json == null) {
        return json;
    }
    return {
        
        'disabledIconId': json['disabled_icon_id'] == null ? undefined : json['disabled_icon_id'],
        'enabledIconId': json['enabled_icon_id'] == null ? undefined : json['enabled_icon_id'],
        'value': json['value'] == null ? undefined : json['value'],
    };
}

export function ModelUserRankToJSON(json: any): ModelUserRank {
    return ModelUserRankToJSONTyped(json, false);
}

export function ModelUserRankToJSONTyped(value?: ModelUserRank | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'disabled_icon_id': value['disabledIconId'],
        'enabled_icon_id': value['enabledIconId'],
        'value': value['value'],
    };
}

