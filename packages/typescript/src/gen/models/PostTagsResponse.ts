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
import type { PostTag } from './PostTag.js';
import {
    PostTagFromJSON,
    PostTagFromJSONTyped,
    PostTagToJSON,
    PostTagToJSONTyped,
} from './PostTag.js';

/**
 * 
 * @export
 * @interface PostTagsResponse
 */
export interface PostTagsResponse {
    /**
     * 
     * @type {Array<PostTag>}
     * @memberof PostTagsResponse
     */
    tags?: Array<PostTag> | null;
}

/**
 * Check if a given object implements the PostTagsResponse interface.
 */
export function instanceOfPostTagsResponse(value: object): value is PostTagsResponse {
    return true;
}

export function PostTagsResponseFromJSON(json: any): PostTagsResponse {
    return PostTagsResponseFromJSONTyped(json, false);
}

export function PostTagsResponseFromJSONTyped(json: any, ignoreDiscriminator: boolean): PostTagsResponse {
    if (json == null) {
        return json;
    }
    return {
        
        'tags': json['tags'] == null ? undefined : ((json['tags'] as Array<any>).map(PostTagFromJSON)),
    };
}

export function PostTagsResponseToJSON(json: any): PostTagsResponse {
    return PostTagsResponseToJSONTyped(json, false);
}

export function PostTagsResponseToJSONTyped(value?: PostTagsResponse | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'tags': value['tags'] == null ? undefined : ((value['tags'] as Array<any>).map(PostTagToJSON)),
    };
}

