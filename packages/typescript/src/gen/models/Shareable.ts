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
import type { Group } from './Group.js';
import {
    GroupFromJSON,
    GroupFromJSONTyped,
    GroupToJSON,
    GroupToJSONTyped,
} from './Group.js';
import type { ThreadInfo } from './ThreadInfo.js';
import {
    ThreadInfoFromJSON,
    ThreadInfoFromJSONTyped,
    ThreadInfoToJSON,
    ThreadInfoToJSONTyped,
} from './ThreadInfo.js';
import type { Post } from './Post.js';
import {
    PostFromJSON,
    PostFromJSONTyped,
    PostToJSON,
    PostToJSONTyped,
} from './Post.js';

/**
 * 
 * @export
 * @interface Shareable
 */
export interface Shareable {
    /**
     * 
     * @type {Group}
     * @memberof Shareable
     */
    group?: Group | null;
    /**
     * 
     * @type {Post}
     * @memberof Shareable
     */
    post?: Post | null;
    /**
     * 
     * @type {ThreadInfo}
     * @memberof Shareable
     */
    thread?: ThreadInfo | null;
}

/**
 * Check if a given object implements the Shareable interface.
 */
export function instanceOfShareable(value: object): value is Shareable {
    return true;
}

export function ShareableFromJSON(json: any): Shareable {
    return ShareableFromJSONTyped(json, false);
}

export function ShareableFromJSONTyped(json: any, ignoreDiscriminator: boolean): Shareable {
    if (json == null) {
        return json;
    }
    return {
        
        'group': json['group'] == null ? undefined : GroupFromJSON(json['group']),
        'post': json['post'] == null ? undefined : PostFromJSON(json['post']),
        'thread': json['thread'] == null ? undefined : ThreadInfoFromJSON(json['thread']),
    };
}

export function ShareableToJSON(json: any): Shareable {
    return ShareableToJSONTyped(json, false);
}

export function ShareableToJSONTyped(value?: Shareable | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'group': GroupToJSON(value['group']),
        'post': PostToJSON(value['post']),
        'thread': ThreadInfoToJSON(value['thread']),
    };
}

