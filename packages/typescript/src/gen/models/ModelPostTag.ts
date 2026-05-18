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
 * @interface ModelPostTag
 */
export interface ModelPostTag {
    /**
     * 
     * @type {number}
     * @memberof ModelPostTag
     */
    id?: number | null;
    /**
     * 
     * @type {number}
     * @memberof ModelPostTag
     */
    postHashtagsCount?: number | null;
    /**
     * 
     * @type {string}
     * @memberof ModelPostTag
     */
    tag?: string | null;
}

/**
 * Check if a given object implements the ModelPostTag interface.
 */
export function instanceOfModelPostTag(value: object): value is ModelPostTag {
    return true;
}

export function ModelPostTagFromJSON(json: any): ModelPostTag {
    return ModelPostTagFromJSONTyped(json, false);
}

export function ModelPostTagFromJSONTyped(json: any, ignoreDiscriminator: boolean): ModelPostTag {
    if (json == null) {
        return json;
    }
    return {
        
        'id': json['id'] == null ? undefined : json['id'],
        'postHashtagsCount': json['post_hashtags_count'] == null ? undefined : json['post_hashtags_count'],
        'tag': json['tag'] == null ? undefined : json['tag'],
    };
}

export function ModelPostTagToJSON(json: any): ModelPostTag {
    return ModelPostTagToJSONTyped(json, false);
}

export function ModelPostTagToJSONTyped(value?: ModelPostTag | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'id': value['id'],
        'post_hashtags_count': value['postHashtagsCount'],
        'tag': value['tag'],
    };
}

