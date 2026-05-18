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
 * @interface TimelinePost
 */
export interface TimelinePost {
    /**
     * 
     * @type {Array<ModelPost>}
     * @memberof TimelinePost
     */
    posts?: Array<ModelPost> | null;
    /**
     * 
     * @type {Array<ModelPost>}
     * @memberof TimelinePost
     */
    rootPosts?: Array<ModelPost> | null;
}

/**
 * Check if a given object implements the TimelinePost interface.
 */
export function instanceOfTimelinePost(value: object): value is TimelinePost {
    return true;
}

export function TimelinePostFromJSON(json: any): TimelinePost {
    return TimelinePostFromJSONTyped(json, false);
}

export function TimelinePostFromJSONTyped(json: any, ignoreDiscriminator: boolean): TimelinePost {
    if (json == null) {
        return json;
    }
    return {
        
        'posts': json['posts'] == null ? undefined : ((json['posts'] as Array<any>).map(ModelPostFromJSON)),
        'rootPosts': json['root_posts'] == null ? undefined : ((json['root_posts'] as Array<any>).map(ModelPostFromJSON)),
    };
}

export function TimelinePostToJSON(json: any): TimelinePost {
    return TimelinePostToJSONTyped(json, false);
}

export function TimelinePostToJSONTyped(value?: TimelinePost | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'posts': value['posts'] == null ? undefined : ((value['posts'] as Array<any>).map(ModelPostToJSON)),
        'root_posts': value['rootPosts'] == null ? undefined : ((value['rootPosts'] as Array<any>).map(ModelPostToJSON)),
    };
}

