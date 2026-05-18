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
 * @interface GroupUsersResponse
 */
export interface GroupUsersResponse {
    /**
     * 
     * @type {Array<GroupUser>}
     * @memberof GroupUsersResponse
     */
    groupUsers?: Array<GroupUser> | null;
}

/**
 * Check if a given object implements the GroupUsersResponse interface.
 */
export function instanceOfGroupUsersResponse(value: object): value is GroupUsersResponse {
    return true;
}

export function GroupUsersResponseFromJSON(json: any): GroupUsersResponse {
    return GroupUsersResponseFromJSONTyped(json, false);
}

export function GroupUsersResponseFromJSONTyped(json: any, ignoreDiscriminator: boolean): GroupUsersResponse {
    if (json == null) {
        return json;
    }
    return {
        
        'groupUsers': json['group_users'] == null ? undefined : ((json['group_users'] as Array<any>).map(GroupUserFromJSON)),
    };
}

export function GroupUsersResponseToJSON(json: any): GroupUsersResponse {
    return GroupUsersResponseToJSONTyped(json, false);
}

export function GroupUsersResponseToJSONTyped(value?: GroupUsersResponse | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'group_users': value['groupUsers'] == null ? undefined : ((value['groupUsers'] as Array<any>).map(GroupUserToJSON)),
    };
}

