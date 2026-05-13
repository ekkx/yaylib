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
import type { RealmUser } from './RealmUser';
import {
    RealmUserFromJSON,
    RealmUserFromJSONTyped,
    RealmUserToJSON,
    RealmUserToJSONTyped,
} from './RealmUser';

/**
 * 
 * @export
 * @interface FollowUsersResponse
 */
export interface FollowUsersResponse {
    /**
     * 
     * @type {number}
     * @memberof FollowUsersResponse
     */
    nextPageValue?: number | null;
    /**
     * 
     * @type {Array<RealmUser>}
     * @memberof FollowUsersResponse
     */
    users?: Array<RealmUser> | null;
}

/**
 * Check if a given object implements the FollowUsersResponse interface.
 */
export function instanceOfFollowUsersResponse(value: object): value is FollowUsersResponse {
    return true;
}

export function FollowUsersResponseFromJSON(json: any): FollowUsersResponse {
    return FollowUsersResponseFromJSONTyped(json, false);
}

export function FollowUsersResponseFromJSONTyped(json: any, ignoreDiscriminator: boolean): FollowUsersResponse {
    if (json == null) {
        return json;
    }
    return {
        
        'nextPageValue': json['next_page_value'] == null ? undefined : json['next_page_value'],
        'users': json['users'] == null ? undefined : ((json['users'] as Array<any>).map(RealmUserFromJSON)),
    };
}

export function FollowUsersResponseToJSON(json: any): FollowUsersResponse {
    return FollowUsersResponseToJSONTyped(json, false);
}

export function FollowUsersResponseToJSONTyped(value?: FollowUsersResponse | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'next_page_value': value['nextPageValue'],
        'users': value['users'] == null ? undefined : ((value['users'] as Array<any>).map(RealmUserToJSON)),
    };
}

