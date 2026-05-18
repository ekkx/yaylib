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
import type { GroupUser } from './GroupUser.js';
import {
    GroupUserFromJSON,
    GroupUserFromJSONTyped,
    GroupUserToJSON,
    GroupUserToJSONTyped,
} from './GroupUser.js';

/**
 * 
 * @export
 * @interface GroupMuteUsersResponse
 */
export interface GroupMuteUsersResponse {
    /**
     * 
     * @type {Array<GroupUser>}
     * @memberof GroupMuteUsersResponse
     */
    groupUsers?: Array<GroupUser> | null;
    /**
     * 
     * @type {number}
     * @memberof GroupMuteUsersResponse
     */
    nextCursor?: number | null;
    /**
     * 
     * @type {number}
     * @memberof GroupMuteUsersResponse
     */
    total?: number | null;
}

/**
 * Check if a given object implements the GroupMuteUsersResponse interface.
 */
export function instanceOfGroupMuteUsersResponse(value: object): value is GroupMuteUsersResponse {
    return true;
}

export function GroupMuteUsersResponseFromJSON(json: any): GroupMuteUsersResponse {
    return GroupMuteUsersResponseFromJSONTyped(json, false);
}

export function GroupMuteUsersResponseFromJSONTyped(json: any, ignoreDiscriminator: boolean): GroupMuteUsersResponse {
    if (json == null) {
        return json;
    }
    return {
        
        'groupUsers': json['group_users'] == null ? undefined : ((json['group_users'] as Array<any>).map(GroupUserFromJSON)),
        'nextCursor': json['next_cursor'] == null ? undefined : json['next_cursor'],
        'total': json['total'] == null ? undefined : json['total'],
    };
}

export function GroupMuteUsersResponseToJSON(json: any): GroupMuteUsersResponse {
    return GroupMuteUsersResponseToJSONTyped(json, false);
}

export function GroupMuteUsersResponseToJSONTyped(value?: GroupMuteUsersResponse | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'group_users': value['groupUsers'] == null ? undefined : ((value['groupUsers'] as Array<any>).map(GroupUserToJSON)),
        'next_cursor': value['nextCursor'],
        'total': value['total'],
    };
}

