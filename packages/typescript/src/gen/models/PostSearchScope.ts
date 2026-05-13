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
 * @interface PostSearchScope
 */
export interface PostSearchScope {
    /**
     * 
     * @type {string}
     * @memberof PostSearchScope
     */
    apiValue?: string | null;
    /**
     * 
     * @type {number}
     * @memberof PostSearchScope
     */
    titleId?: number | null;
}

/**
 * Check if a given object implements the PostSearchScope interface.
 */
export function instanceOfPostSearchScope(value: object): value is PostSearchScope {
    return true;
}

export function PostSearchScopeFromJSON(json: any): PostSearchScope {
    return PostSearchScopeFromJSONTyped(json, false);
}

export function PostSearchScopeFromJSONTyped(json: any, ignoreDiscriminator: boolean): PostSearchScope {
    if (json == null) {
        return json;
    }
    return {
        
        'apiValue': json['api_value'] == null ? undefined : json['api_value'],
        'titleId': json['title_id'] == null ? undefined : json['title_id'],
    };
}

export function PostSearchScopeToJSON(json: any): PostSearchScope {
    return PostSearchScopeToJSONTyped(json, false);
}

export function PostSearchScopeToJSONTyped(value?: PostSearchScope | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'api_value': value['apiValue'],
        'title_id': value['titleId'],
    };
}

