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
 * @interface PostSearchFilter
 */
export interface PostSearchFilter {
    /**
     * 
     * @type {number}
     * @memberof PostSearchFilter
     */
    messageId?: number | null;
    /**
     * 
     * @type {number}
     * @memberof PostSearchFilter
     */
    titleId?: number | null;
}

/**
 * Check if a given object implements the PostSearchFilter interface.
 */
export function instanceOfPostSearchFilter(value: object): value is PostSearchFilter {
    return true;
}

export function PostSearchFilterFromJSON(json: any): PostSearchFilter {
    return PostSearchFilterFromJSONTyped(json, false);
}

export function PostSearchFilterFromJSONTyped(json: any, ignoreDiscriminator: boolean): PostSearchFilter {
    if (json == null) {
        return json;
    }
    return {
        
        'messageId': json['message_id'] == null ? undefined : json['message_id'],
        'titleId': json['title_id'] == null ? undefined : json['title_id'],
    };
}

export function PostSearchFilterToJSON(json: any): PostSearchFilter {
    return PostSearchFilterToJSONTyped(json, false);
}

export function PostSearchFilterToJSONTyped(value?: PostSearchFilter | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'message_id': value['messageId'],
        'title_id': value['titleId'],
    };
}

