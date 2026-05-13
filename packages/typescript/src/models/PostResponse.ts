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
 * @interface PostResponse
 */
export interface PostResponse {
    /**
     * 
     * @type {Post}
     * @memberof PostResponse
     */
    post?: Post | null;
}

/**
 * Check if a given object implements the PostResponse interface.
 */
export function instanceOfPostResponse(value: object): value is PostResponse {
    return true;
}

export function PostResponseFromJSON(json: any): PostResponse {
    return PostResponseFromJSONTyped(json, false);
}

export function PostResponseFromJSONTyped(json: any, ignoreDiscriminator: boolean): PostResponse {
    if (json == null) {
        return json;
    }
    return {
        
        'post': json['post'] == null ? undefined : PostFromJSON(json['post']),
    };
}

export function PostResponseToJSON(json: any): PostResponse {
    return PostResponseToJSONTyped(json, false);
}

export function PostResponseToJSONTyped(value?: PostResponse | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'post': PostToJSON(value['post']),
    };
}

