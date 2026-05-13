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
import type { UserSearchScope } from './UserSearchScope';
import {
    UserSearchScopeFromJSON,
    UserSearchScopeFromJSONTyped,
    UserSearchScopeToJSON,
    UserSearchScopeToJSONTyped,
} from './UserSearchScope';

/**
 * 
 * @export
 * @interface UserSearchQuery
 */
export interface UserSearchQuery {
    /**
     * 
     * @type {string}
     * @memberof UserSearchQuery
     */
    keyword?: string | null;
    /**
     * 
     * @type {boolean}
     * @memberof UserSearchQuery
     */
    noRecentPenalty?: boolean | null;
    /**
     * 
     * @type {boolean}
     * @memberof UserSearchQuery
     */
    samePrefecture?: boolean | null;
    /**
     * 
     * @type {UserSearchScope}
     * @memberof UserSearchQuery
     */
    scope?: UserSearchScope | null;
    /**
     * 
     * @type {boolean}
     * @memberof UserSearchQuery
     */
    similarAge?: boolean | null;
}

/**
 * Check if a given object implements the UserSearchQuery interface.
 */
export function instanceOfUserSearchQuery(value: object): value is UserSearchQuery {
    return true;
}

export function UserSearchQueryFromJSON(json: any): UserSearchQuery {
    return UserSearchQueryFromJSONTyped(json, false);
}

export function UserSearchQueryFromJSONTyped(json: any, ignoreDiscriminator: boolean): UserSearchQuery {
    if (json == null) {
        return json;
    }
    return {
        
        'keyword': json['keyword'] == null ? undefined : json['keyword'],
        'noRecentPenalty': json['no_recent_penalty'] == null ? undefined : json['no_recent_penalty'],
        'samePrefecture': json['same_prefecture'] == null ? undefined : json['same_prefecture'],
        'scope': json['scope'] == null ? undefined : UserSearchScopeFromJSON(json['scope']),
        'similarAge': json['similar_age'] == null ? undefined : json['similar_age'],
    };
}

export function UserSearchQueryToJSON(json: any): UserSearchQuery {
    return UserSearchQueryToJSONTyped(json, false);
}

export function UserSearchQueryToJSONTyped(value?: UserSearchQuery | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'keyword': value['keyword'],
        'no_recent_penalty': value['noRecentPenalty'],
        'same_prefecture': value['samePrefecture'],
        'scope': UserSearchScopeToJSON(value['scope']),
        'similar_age': value['similarAge'],
    };
}

