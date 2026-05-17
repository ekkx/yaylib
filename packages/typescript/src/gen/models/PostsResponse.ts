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
import type { Post } from './Post';
import {
    PostFromJSON,
    PostFromJSONTyped,
    PostToJSON,
    PostToJSONTyped,
} from './Post';

/**
 * 
 * @export
 * @interface PostsResponse
 */
export interface PostsResponse {
    /**
     * 
     * @type {boolean}
     * @memberof PostsResponse
     */
    hasMoreHotPosts?: boolean | null;
    /**
     * 
     * @type {string}
     * @memberof PostsResponse
     */
    nextPageValue?: string | null;
    /**
     * 
     * @type {Array<Post>}
     * @memberof PostsResponse
     */
    pinnedPosts?: Array<Post> | null;
    /**
     * 
     * @type {Array<Post>}
     * @memberof PostsResponse
     */
    posts?: Array<Post> | null;
}

/**
 * Check if a given object implements the PostsResponse interface.
 */
export function instanceOfPostsResponse(value: object): value is PostsResponse {
    return true;
}

export function PostsResponseFromJSON(json: any): PostsResponse {
    return PostsResponseFromJSONTyped(json, false);
}

export function PostsResponseFromJSONTyped(json: any, ignoreDiscriminator: boolean): PostsResponse {
    if (json == null) {
        return json;
    }
    return {
        
        'hasMoreHotPosts': json['has_more_hot_posts'] == null ? undefined : json['has_more_hot_posts'],
        'nextPageValue': json['next_page_value'] == null ? undefined : json['next_page_value'],
        'pinnedPosts': json['pinned_posts'] == null ? undefined : ((json['pinned_posts'] as Array<any>).map(PostFromJSON)),
        'posts': json['posts'] == null ? undefined : ((json['posts'] as Array<any>).map(PostFromJSON)),
    };
}

export function PostsResponseToJSON(json: any): PostsResponse {
    return PostsResponseToJSONTyped(json, false);
}

export function PostsResponseToJSONTyped(value?: PostsResponse | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'has_more_hot_posts': value['hasMoreHotPosts'],
        'next_page_value': value['nextPageValue'],
        'pinned_posts': value['pinnedPosts'] == null ? undefined : ((value['pinnedPosts'] as Array<any>).map(PostToJSON)),
        'posts': value['posts'] == null ? undefined : ((value['posts'] as Array<any>).map(PostToJSON)),
    };
}

