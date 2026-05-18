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
import type { RealmUser } from './RealmUser.js';
import {
    RealmUserFromJSON,
    RealmUserFromJSONTyped,
    RealmUserToJSON,
    RealmUserToJSONTyped,
} from './RealmUser.js';

/**
 * 
 * @export
 * @interface BlockedUsersResponse
 */
export interface BlockedUsersResponse {
    /**
     * 
     * @type {number}
     * @memberof BlockedUsersResponse
     */
    blockedCount?: number | null;
    /**
     * 
     * @type {number}
     * @memberof BlockedUsersResponse
     */
    lastId?: number | null;
    /**
     * 
     * @type {Array<RealmUser>}
     * @memberof BlockedUsersResponse
     */
    users?: Array<RealmUser> | null;
}

/**
 * Check if a given object implements the BlockedUsersResponse interface.
 */
export function instanceOfBlockedUsersResponse(value: object): value is BlockedUsersResponse {
    return true;
}

export function BlockedUsersResponseFromJSON(json: any): BlockedUsersResponse {
    return BlockedUsersResponseFromJSONTyped(json, false);
}

export function BlockedUsersResponseFromJSONTyped(json: any, ignoreDiscriminator: boolean): BlockedUsersResponse {
    if (json == null) {
        return json;
    }
    return {
        
        'blockedCount': json['blocked_count'] == null ? undefined : json['blocked_count'],
        'lastId': json['last_id'] == null ? undefined : json['last_id'],
        'users': json['users'] == null ? undefined : ((json['users'] as Array<any>).map(RealmUserFromJSON)),
    };
}

export function BlockedUsersResponseToJSON(json: any): BlockedUsersResponse {
    return BlockedUsersResponseToJSONTyped(json, false);
}

export function BlockedUsersResponseToJSONTyped(value?: BlockedUsersResponse | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'blocked_count': value['blockedCount'],
        'last_id': value['lastId'],
        'users': value['users'] == null ? undefined : ((value['users'] as Array<any>).map(RealmUserToJSON)),
    };
}

