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
import type { ModelPost } from './ModelPost';
import {
    ModelPostFromJSON,
    ModelPostFromJSONTyped,
    ModelPostToJSON,
    ModelPostToJSONTyped,
} from './ModelPost';

/**
 * 
 * @export
 * @interface UnavailableRootPost
 */
export interface UnavailableRootPost {
    /**
     * 
     * @type {ModelPost}
     * @memberof UnavailableRootPost
     */
    post?: ModelPost | null;
}

/**
 * Check if a given object implements the UnavailableRootPost interface.
 */
export function instanceOfUnavailableRootPost(value: object): value is UnavailableRootPost {
    return true;
}

export function UnavailableRootPostFromJSON(json: any): UnavailableRootPost {
    return UnavailableRootPostFromJSONTyped(json, false);
}

export function UnavailableRootPostFromJSONTyped(json: any, ignoreDiscriminator: boolean): UnavailableRootPost {
    if (json == null) {
        return json;
    }
    return {
        
        'post': json['post'] == null ? undefined : ModelPostFromJSON(json['post']),
    };
}

export function UnavailableRootPostToJSON(json: any): UnavailableRootPost {
    return UnavailableRootPostToJSONTyped(json, false);
}

export function UnavailableRootPostToJSONTyped(value?: UnavailableRootPost | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'post': ModelPostToJSON(value['post']),
    };
}

