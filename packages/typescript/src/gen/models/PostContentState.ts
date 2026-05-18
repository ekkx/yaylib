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
 * @interface PostContentState
 */
export interface PostContentState {
    /**
     * 
     * @type {number}
     * @memberof PostContentState
     */
    placeHolderId?: number | null;
}

/**
 * Check if a given object implements the PostContentState interface.
 */
export function instanceOfPostContentState(value: object): value is PostContentState {
    return true;
}

export function PostContentStateFromJSON(json: any): PostContentState {
    return PostContentStateFromJSONTyped(json, false);
}

export function PostContentStateFromJSONTyped(json: any, ignoreDiscriminator: boolean): PostContentState {
    if (json == null) {
        return json;
    }
    return {
        
        'placeHolderId': json['place_holder_id'] == null ? undefined : json['place_holder_id'],
    };
}

export function PostContentStateToJSON(json: any): PostContentState {
    return PostContentStateToJSONTyped(json, false);
}

export function PostContentStateToJSONTyped(value?: PostContentState | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'place_holder_id': value['placeHolderId'],
    };
}

