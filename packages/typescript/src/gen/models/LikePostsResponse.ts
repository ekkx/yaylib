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
 * @interface LikePostsResponse
 */
export interface LikePostsResponse {
    /**
     * 
     * @type {Array<number>}
     * @memberof LikePostsResponse
     */
    likeIds?: Array<number> | null;
}

/**
 * Check if a given object implements the LikePostsResponse interface.
 */
export function instanceOfLikePostsResponse(value: object): value is LikePostsResponse {
    return true;
}

export function LikePostsResponseFromJSON(json: any): LikePostsResponse {
    return LikePostsResponseFromJSONTyped(json, false);
}

export function LikePostsResponseFromJSONTyped(json: any, ignoreDiscriminator: boolean): LikePostsResponse {
    if (json == null) {
        return json;
    }
    return {
        
        'likeIds': json['like_ids'] == null ? undefined : json['like_ids'],
    };
}

export function LikePostsResponseToJSON(json: any): LikePostsResponse {
    return LikePostsResponseToJSONTyped(json, false);
}

export function LikePostsResponseToJSONTyped(value?: LikePostsResponse | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'like_ids': value['likeIds'],
    };
}

