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
 * @interface PostTag
 */
export interface PostTag {
    /**
     * 
     * @type {number}
     * @memberof PostTag
     */
    id?: number | null;
    /**
     * 
     * @type {number}
     * @memberof PostTag
     */
    postHashtagsCount?: number | null;
    /**
     * 
     * @type {string}
     * @memberof PostTag
     */
    tag?: string | null;
}

/**
 * Check if a given object implements the PostTag interface.
 */
export function instanceOfPostTag(value: object): value is PostTag {
    return true;
}

export function PostTagFromJSON(json: any): PostTag {
    return PostTagFromJSONTyped(json, false);
}

export function PostTagFromJSONTyped(json: any, ignoreDiscriminator: boolean): PostTag {
    if (json == null) {
        return json;
    }
    return {
        
        'id': json['id'] == null ? undefined : json['id'],
        'postHashtagsCount': json['post_hashtags_count'] == null ? undefined : json['post_hashtags_count'],
        'tag': json['tag'] == null ? undefined : json['tag'],
    };
}

export function PostTagToJSON(json: any): PostTag {
    return PostTagToJSONTyped(json, false);
}

export function PostTagToJSONTyped(value?: PostTag | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'id': value['id'],
        'post_hashtags_count': value['postHashtagsCount'],
        'tag': value['tag'],
    };
}

