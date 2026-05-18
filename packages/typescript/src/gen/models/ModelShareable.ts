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
import type { ModelGroup } from './ModelGroup.js';
import {
    ModelGroupFromJSON,
    ModelGroupFromJSONTyped,
    ModelGroupToJSON,
    ModelGroupToJSONTyped,
} from './ModelGroup.js';
import type { ModelPost } from './ModelPost.js';
import {
    ModelPostFromJSON,
    ModelPostFromJSONTyped,
    ModelPostToJSON,
    ModelPostToJSONTyped,
} from './ModelPost.js';
import type { ModelThreadInfo } from './ModelThreadInfo.js';
import {
    ModelThreadInfoFromJSON,
    ModelThreadInfoFromJSONTyped,
    ModelThreadInfoToJSON,
    ModelThreadInfoToJSONTyped,
} from './ModelThreadInfo.js';

/**
 * 
 * @export
 * @interface ModelShareable
 */
export interface ModelShareable {
    /**
     * 
     * @type {ModelGroup}
     * @memberof ModelShareable
     */
    group?: ModelGroup | null;
    /**
     * 
     * @type {ModelPost}
     * @memberof ModelShareable
     */
    post?: ModelPost | null;
    /**
     * 
     * @type {ModelThreadInfo}
     * @memberof ModelShareable
     */
    thread?: ModelThreadInfo | null;
}

/**
 * Check if a given object implements the ModelShareable interface.
 */
export function instanceOfModelShareable(value: object): value is ModelShareable {
    return true;
}

export function ModelShareableFromJSON(json: any): ModelShareable {
    return ModelShareableFromJSONTyped(json, false);
}

export function ModelShareableFromJSONTyped(json: any, ignoreDiscriminator: boolean): ModelShareable {
    if (json == null) {
        return json;
    }
    return {
        
        'group': json['group'] == null ? undefined : ModelGroupFromJSON(json['group']),
        'post': json['post'] == null ? undefined : ModelPostFromJSON(json['post']),
        'thread': json['thread'] == null ? undefined : ModelThreadInfoFromJSON(json['thread']),
    };
}

export function ModelShareableToJSON(json: any): ModelShareable {
    return ModelShareableToJSONTyped(json, false);
}

export function ModelShareableToJSONTyped(value?: ModelShareable | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'group': ModelGroupToJSON(value['group']),
        'post': ModelPostToJSON(value['post']),
        'thread': ModelThreadInfoToJSON(value['thread']),
    };
}

