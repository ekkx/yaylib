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
 * @interface UserSearchScope
 */
export interface UserSearchScope {
    /**
     * 
     * @type {number}
     * @memberof UserSearchScope
     */
    titleId?: number | null;
}

/**
 * Check if a given object implements the UserSearchScope interface.
 */
export function instanceOfUserSearchScope(value: object): value is UserSearchScope {
    return true;
}

export function UserSearchScopeFromJSON(json: any): UserSearchScope {
    return UserSearchScopeFromJSONTyped(json, false);
}

export function UserSearchScopeFromJSONTyped(json: any, ignoreDiscriminator: boolean): UserSearchScope {
    if (json == null) {
        return json;
    }
    return {
        
        'titleId': json['title_id'] == null ? undefined : json['title_id'],
    };
}

export function UserSearchScopeToJSON(json: any): UserSearchScope {
    return UserSearchScopeToJSONTyped(json, false);
}

export function UserSearchScopeToJSONTyped(value?: UserSearchScope | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'title_id': value['titleId'],
    };
}

