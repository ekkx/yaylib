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
 * @interface ModelInterest
 */
export interface ModelInterest {
    /**
     * 
     * @type {string}
     * @memberof ModelInterest
     */
    iconUrl?: string | null;
    /**
     * 
     * @type {number}
     * @memberof ModelInterest
     */
    id?: number | null;
    /**
     * 
     * @type {boolean}
     * @memberof ModelInterest
     */
    isSelected?: boolean | null;
    /**
     * 
     * @type {string}
     * @memberof ModelInterest
     */
    name?: string | null;
}

/**
 * Check if a given object implements the ModelInterest interface.
 */
export function instanceOfModelInterest(value: object): value is ModelInterest {
    return true;
}

export function ModelInterestFromJSON(json: any): ModelInterest {
    return ModelInterestFromJSONTyped(json, false);
}

export function ModelInterestFromJSONTyped(json: any, ignoreDiscriminator: boolean): ModelInterest {
    if (json == null) {
        return json;
    }
    return {
        
        'iconUrl': json['icon_url'] == null ? undefined : json['icon_url'],
        'id': json['id'] == null ? undefined : json['id'],
        'isSelected': json['is_selected'] == null ? undefined : json['is_selected'],
        'name': json['name'] == null ? undefined : json['name'],
    };
}

export function ModelInterestToJSON(json: any): ModelInterest {
    return ModelInterestToJSONTyped(json, false);
}

export function ModelInterestToJSONTyped(value?: ModelInterest | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'icon_url': value['iconUrl'],
        'id': value['id'],
        'is_selected': value['isSelected'],
        'name': value['name'],
    };
}

