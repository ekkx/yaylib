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
import type { RecentSearchType } from './RecentSearchType.js';
import {
    RecentSearchTypeFromJSON,
    RecentSearchTypeFromJSONTyped,
    RecentSearchTypeToJSON,
    RecentSearchTypeToJSONTyped,
} from './RecentSearchType.js';
import type { User } from './User.js';
import {
    UserFromJSON,
    UserFromJSONTyped,
    UserToJSON,
    UserToJSONTyped,
} from './User.js';
import type { ModelPostTag } from './ModelPostTag.js';
import {
    ModelPostTagFromJSON,
    ModelPostTagFromJSONTyped,
    ModelPostTagToJSON,
    ModelPostTagToJSONTyped,
} from './ModelPostTag.js';

/**
 * 
 * @export
 * @interface ModelRecentSearch
 */
export interface ModelRecentSearch {
    /**
     * 
     * @type {ModelPostTag}
     * @memberof ModelRecentSearch
     */
    hashtag?: ModelPostTag | null;
    /**
     * 
     * @type {number}
     * @memberof ModelRecentSearch
     */
    id?: number | null;
    /**
     * 
     * @type {string}
     * @memberof ModelRecentSearch
     */
    keyword?: string | null;
    /**
     * 
     * @type {RecentSearchType}
     * @memberof ModelRecentSearch
     */
    type?: RecentSearchType | null;
    /**
     * 
     * @type {User}
     * @memberof ModelRecentSearch
     */
    user?: User | null;
}

/**
 * Check if a given object implements the ModelRecentSearch interface.
 */
export function instanceOfModelRecentSearch(value: object): value is ModelRecentSearch {
    return true;
}

export function ModelRecentSearchFromJSON(json: any): ModelRecentSearch {
    return ModelRecentSearchFromJSONTyped(json, false);
}

export function ModelRecentSearchFromJSONTyped(json: any, ignoreDiscriminator: boolean): ModelRecentSearch {
    if (json == null) {
        return json;
    }
    return {
        
        'hashtag': json['hashtag'] == null ? undefined : ModelPostTagFromJSON(json['hashtag']),
        'id': json['id'] == null ? undefined : json['id'],
        'keyword': json['keyword'] == null ? undefined : json['keyword'],
        'type': json['type'] == null ? undefined : RecentSearchTypeFromJSON(json['type']),
        'user': json['user'] == null ? undefined : UserFromJSON(json['user']),
    };
}

export function ModelRecentSearchToJSON(json: any): ModelRecentSearch {
    return ModelRecentSearchToJSONTyped(json, false);
}

export function ModelRecentSearchToJSONTyped(value?: ModelRecentSearch | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'hashtag': ModelPostTagToJSON(value['hashtag']),
        'id': value['id'],
        'keyword': value['keyword'],
        'type': RecentSearchTypeToJSON(value['type']),
        'user': UserToJSON(value['user']),
    };
}

