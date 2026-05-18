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
import type { PostSearchScope } from './PostSearchScope.js';
import {
    PostSearchScopeFromJSON,
    PostSearchScopeFromJSONTyped,
    PostSearchScopeToJSON,
    PostSearchScopeToJSONTyped,
} from './PostSearchScope.js';

/**
 * 
 * @export
 * @interface PostSearchQuery
 */
export interface PostSearchQuery {
    /**
     * 
     * @type {string}
     * @memberof PostSearchQuery
     */
    keyword?: string | null;
    /**
     * 
     * @type {boolean}
     * @memberof PostSearchQuery
     */
    onlyMedia?: boolean | null;
    /**
     * 
     * @type {PostSearchScope}
     * @memberof PostSearchQuery
     */
    scope?: PostSearchScope | null;
}

/**
 * Check if a given object implements the PostSearchQuery interface.
 */
export function instanceOfPostSearchQuery(value: object): value is PostSearchQuery {
    return true;
}

export function PostSearchQueryFromJSON(json: any): PostSearchQuery {
    return PostSearchQueryFromJSONTyped(json, false);
}

export function PostSearchQueryFromJSONTyped(json: any, ignoreDiscriminator: boolean): PostSearchQuery {
    if (json == null) {
        return json;
    }
    return {
        
        'keyword': json['keyword'] == null ? undefined : json['keyword'],
        'onlyMedia': json['only_media'] == null ? undefined : json['only_media'],
        'scope': json['scope'] == null ? undefined : PostSearchScopeFromJSON(json['scope']),
    };
}

export function PostSearchQueryToJSON(json: any): PostSearchQuery {
    return PostSearchQueryToJSONTyped(json, false);
}

export function PostSearchQueryToJSONTyped(value?: PostSearchQuery | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'keyword': value['keyword'],
        'only_media': value['onlyMedia'],
        'scope': PostSearchScopeToJSON(value['scope']),
    };
}

