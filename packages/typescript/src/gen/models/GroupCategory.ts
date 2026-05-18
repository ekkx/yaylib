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
 * @interface GroupCategory
 */
export interface GroupCategory {
    /**
     * 
     * @type {string}
     * @memberof GroupCategory
     */
    icon?: string | null;
    /**
     * 
     * @type {number}
     * @memberof GroupCategory
     */
    id?: number | null;
    /**
     * 
     * @type {string}
     * @memberof GroupCategory
     */
    name?: string | null;
    /**
     * 
     * @type {number}
     * @memberof GroupCategory
     */
    rank?: number | null;
}

/**
 * Check if a given object implements the GroupCategory interface.
 */
export function instanceOfGroupCategory(value: object): value is GroupCategory {
    return true;
}

export function GroupCategoryFromJSON(json: any): GroupCategory {
    return GroupCategoryFromJSONTyped(json, false);
}

export function GroupCategoryFromJSONTyped(json: any, ignoreDiscriminator: boolean): GroupCategory {
    if (json == null) {
        return json;
    }
    return {
        
        'icon': json['icon'] == null ? undefined : json['icon'],
        'id': json['id'] == null ? undefined : json['id'],
        'name': json['name'] == null ? undefined : json['name'],
        'rank': json['rank'] == null ? undefined : json['rank'],
    };
}

export function GroupCategoryToJSON(json: any): GroupCategory {
    return GroupCategoryToJSONTyped(json, false);
}

export function GroupCategoryToJSONTyped(value?: GroupCategory | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'icon': value['icon'],
        'id': value['id'],
        'name': value['name'],
        'rank': value['rank'],
    };
}

