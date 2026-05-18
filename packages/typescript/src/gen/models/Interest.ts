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
 * @interface Interest
 */
export interface Interest {
    /**
     * 
     * @type {string}
     * @memberof Interest
     */
    icon?: string | null;
    /**
     * 
     * @type {number}
     * @memberof Interest
     */
    id?: number | null;
    /**
     * 
     * @type {string}
     * @memberof Interest
     */
    name?: string | null;
    /**
     * 
     * @type {boolean}
     * @memberof Interest
     */
    selected?: boolean | null;
}

/**
 * Check if a given object implements the Interest interface.
 */
export function instanceOfInterest(value: object): value is Interest {
    return true;
}

export function InterestFromJSON(json: any): Interest {
    return InterestFromJSONTyped(json, false);
}

export function InterestFromJSONTyped(json: any, ignoreDiscriminator: boolean): Interest {
    if (json == null) {
        return json;
    }
    return {
        
        'icon': json['icon'] == null ? undefined : json['icon'],
        'id': json['id'] == null ? undefined : json['id'],
        'name': json['name'] == null ? undefined : json['name'],
        'selected': json['selected'] == null ? undefined : json['selected'],
    };
}

export function InterestToJSON(json: any): Interest {
    return InterestToJSONTyped(json, false);
}

export function InterestToJSONTyped(value?: Interest | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'icon': value['icon'],
        'id': value['id'],
        'name': value['name'],
        'selected': value['selected'],
    };
}

