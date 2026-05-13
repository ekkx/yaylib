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
import type { Group } from './Group';
import {
    GroupFromJSON,
    GroupFromJSONTyped,
    GroupToJSON,
    GroupToJSONTyped,
} from './Group';

/**
 * 
 * @export
 * @interface GroupsResponse
 */
export interface GroupsResponse {
    /**
     * 
     * @type {Array<Group>}
     * @memberof GroupsResponse
     */
    groups?: Array<Group> | null;
    /**
     * 
     * @type {Array<Group>}
     * @memberof GroupsResponse
     */
    pinnedGroups?: Array<Group> | null;
}

/**
 * Check if a given object implements the GroupsResponse interface.
 */
export function instanceOfGroupsResponse(value: object): value is GroupsResponse {
    return true;
}

export function GroupsResponseFromJSON(json: any): GroupsResponse {
    return GroupsResponseFromJSONTyped(json, false);
}

export function GroupsResponseFromJSONTyped(json: any, ignoreDiscriminator: boolean): GroupsResponse {
    if (json == null) {
        return json;
    }
    return {
        
        'groups': json['groups'] == null ? undefined : ((json['groups'] as Array<any>).map(GroupFromJSON)),
        'pinnedGroups': json['pinned_groups'] == null ? undefined : ((json['pinned_groups'] as Array<any>).map(GroupFromJSON)),
    };
}

export function GroupsResponseToJSON(json: any): GroupsResponse {
    return GroupsResponseToJSONTyped(json, false);
}

export function GroupsResponseToJSONTyped(value?: GroupsResponse | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'groups': value['groups'] == null ? undefined : ((value['groups'] as Array<any>).map(GroupToJSON)),
        'pinned_groups': value['pinnedGroups'] == null ? undefined : ((value['pinnedGroups'] as Array<any>).map(GroupToJSON)),
    };
}

