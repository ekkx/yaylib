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
 * @interface GroupsRelatedResponse
 */
export interface GroupsRelatedResponse {
    /**
     * 
     * @type {Array<Group>}
     * @memberof GroupsRelatedResponse
     */
    groups?: Array<Group> | null;
    /**
     * 
     * @type {string}
     * @memberof GroupsRelatedResponse
     */
    nextPageValue?: string | null;
}

/**
 * Check if a given object implements the GroupsRelatedResponse interface.
 */
export function instanceOfGroupsRelatedResponse(value: object): value is GroupsRelatedResponse {
    return true;
}

export function GroupsRelatedResponseFromJSON(json: any): GroupsRelatedResponse {
    return GroupsRelatedResponseFromJSONTyped(json, false);
}

export function GroupsRelatedResponseFromJSONTyped(json: any, ignoreDiscriminator: boolean): GroupsRelatedResponse {
    if (json == null) {
        return json;
    }
    return {
        
        'groups': json['groups'] == null ? undefined : ((json['groups'] as Array<any>).map(GroupFromJSON)),
        'nextPageValue': json['next_page_value'] == null ? undefined : String(json['next_page_value']),
    };
}

export function GroupsRelatedResponseToJSON(json: any): GroupsRelatedResponse {
    return GroupsRelatedResponseToJSONTyped(json, false);
}

export function GroupsRelatedResponseToJSONTyped(value?: GroupsRelatedResponse | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'groups': value['groups'] == null ? undefined : ((value['groups'] as Array<any>).map(GroupToJSON)),
        'next_page_value': value['nextPageValue'],
    };
}

