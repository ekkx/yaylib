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
 * @interface UserSearchFilter
 */
export interface UserSearchFilter {
    /**
     * 
     * @type {number}
     * @memberof UserSearchFilter
     */
    titleId?: number | null;
}

/**
 * Check if a given object implements the UserSearchFilter interface.
 */
export function instanceOfUserSearchFilter(value: object): value is UserSearchFilter {
    return true;
}

export function UserSearchFilterFromJSON(json: any): UserSearchFilter {
    return UserSearchFilterFromJSONTyped(json, false);
}

export function UserSearchFilterFromJSONTyped(json: any, ignoreDiscriminator: boolean): UserSearchFilter {
    if (json == null) {
        return json;
    }
    return {
        
        'titleId': json['title_id'] == null ? undefined : json['title_id'],
    };
}

export function UserSearchFilterToJSON(json: any): UserSearchFilter {
    return UserSearchFilterToJSONTyped(json, false);
}

export function UserSearchFilterToJSONTyped(value?: UserSearchFilter | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'title_id': value['titleId'],
    };
}

