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

/**
 * 
 * @export
 * @interface GroupResponse
 */
export interface GroupResponse {
    /**
     * 
     * @type {Group}
     * @memberof GroupResponse
     */
    group?: Group | null;
}

/**
 * Check if a given object implements the GroupResponse interface.
 */
export function instanceOfGroupResponse(value: object): value is GroupResponse {
    return true;
}

export function GroupResponseFromJSON(json: any): GroupResponse {
    return GroupResponseFromJSONTyped(json, false);
}

export function GroupResponseFromJSONTyped(json: any, ignoreDiscriminator: boolean): GroupResponse {
    if (json == null) {
        return json;
    }
    return {
        
        'group': json['group'] == null ? undefined : GroupFromJSON(json['group']),
    };
}

export function GroupResponseToJSON(json: any): GroupResponse {
    return GroupResponseToJSONTyped(json, false);
}

export function GroupResponseToJSONTyped(value?: GroupResponse | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'group': GroupToJSON(value['group']),
    };
}

