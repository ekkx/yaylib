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
import type { ModelPost } from './ModelPost.js';
import {
    ModelPostFromJSON,
    ModelPostFromJSONTyped,
    ModelPostToJSON,
    ModelPostToJSONTyped,
} from './ModelPost.js';

/**
 * 
 * @export
 * @interface HiddenRecommendedPost
 */
export interface HiddenRecommendedPost {
    /**
     * 
     * @type {ModelPost}
     * @memberof HiddenRecommendedPost
     */
    post?: ModelPost | null;
}

/**
 * Check if a given object implements the HiddenRecommendedPost interface.
 */
export function instanceOfHiddenRecommendedPost(value: object): value is HiddenRecommendedPost {
    return true;
}

export function HiddenRecommendedPostFromJSON(json: any): HiddenRecommendedPost {
    return HiddenRecommendedPostFromJSONTyped(json, false);
}

export function HiddenRecommendedPostFromJSONTyped(json: any, ignoreDiscriminator: boolean): HiddenRecommendedPost {
    if (json == null) {
        return json;
    }
    return {
        
        'post': json['post'] == null ? undefined : ModelPostFromJSON(json['post']),
    };
}

export function HiddenRecommendedPostToJSON(json: any): HiddenRecommendedPost {
    return HiddenRecommendedPostToJSONTyped(json, false);
}

export function HiddenRecommendedPostToJSONTyped(value?: HiddenRecommendedPost | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'post': ModelPostToJSON(value['post']),
    };
}

