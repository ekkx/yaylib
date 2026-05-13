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
import type { RealmUser } from './RealmUser';
import {
    RealmUserFromJSON,
    RealmUserFromJSONTyped,
    RealmUserToJSON,
    RealmUserToJSONTyped,
} from './RealmUser';
import type { PostTag } from './PostTag';
import {
    PostTagFromJSON,
    PostTagFromJSONTyped,
    PostTagToJSON,
    PostTagToJSONTyped,
} from './PostTag';

/**
 * 
 * @export
 * @interface RecentSearch
 */
export interface RecentSearch {
    /**
     * 
     * @type {PostTag}
     * @memberof RecentSearch
     */
    hashtag?: PostTag | null;
    /**
     * 
     * @type {number}
     * @memberof RecentSearch
     */
    id?: number | null;
    /**
     * 
     * @type {string}
     * @memberof RecentSearch
     */
    keyword?: string | null;
    /**
     * 
     * @type {string}
     * @memberof RecentSearch
     */
    type?: string | null;
    /**
     * 
     * @type {RealmUser}
     * @memberof RecentSearch
     */
    user?: RealmUser | null;
}

/**
 * Check if a given object implements the RecentSearch interface.
 */
export function instanceOfRecentSearch(value: object): value is RecentSearch {
    return true;
}

export function RecentSearchFromJSON(json: any): RecentSearch {
    return RecentSearchFromJSONTyped(json, false);
}

export function RecentSearchFromJSONTyped(json: any, ignoreDiscriminator: boolean): RecentSearch {
    if (json == null) {
        return json;
    }
    return {
        
        'hashtag': json['hashtag'] == null ? undefined : PostTagFromJSON(json['hashtag']),
        'id': json['id'] == null ? undefined : json['id'],
        'keyword': json['keyword'] == null ? undefined : json['keyword'],
        'type': json['type'] == null ? undefined : json['type'],
        'user': json['user'] == null ? undefined : RealmUserFromJSON(json['user']),
    };
}

export function RecentSearchToJSON(json: any): RecentSearch {
    return RecentSearchToJSONTyped(json, false);
}

export function RecentSearchToJSONTyped(value?: RecentSearch | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'hashtag': PostTagToJSON(value['hashtag']),
        'id': value['id'],
        'keyword': value['keyword'],
        'type': value['type'],
        'user': RealmUserToJSON(value['user']),
    };
}

