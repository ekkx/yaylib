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
 * @interface GroupUserResponse
 */
export interface GroupUserResponse {
    /**
     * 
     * @type {GroupUser}
     * @memberof GroupUserResponse
     */
    groupUser?: GroupUser | null;
}

/**
 * Check if a given object implements the GroupUserResponse interface.
 */
export function instanceOfGroupUserResponse(value: object): value is GroupUserResponse {
    return true;
}

export function GroupUserResponseFromJSON(json: any): GroupUserResponse {
    return GroupUserResponseFromJSONTyped(json, false);
}

export function GroupUserResponseFromJSONTyped(json: any, ignoreDiscriminator: boolean): GroupUserResponse {
    if (json == null) {
        return json;
    }
    return {
        
        'groupUser': json['group_user'] == null ? undefined : GroupUserFromJSON(json['group_user']),
    };
}

export function GroupUserResponseToJSON(json: any): GroupUserResponse {
    return GroupUserResponseToJSONTyped(json, false);
}

export function GroupUserResponseToJSONTyped(value?: GroupUserResponse | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'group_user': GroupUserToJSON(value['groupUser']),
    };
}

