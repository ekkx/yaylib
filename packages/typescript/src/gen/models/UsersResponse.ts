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
 * @interface UsersResponse
 */
export interface UsersResponse {
    /**
     * 
     * @type {string}
     * @memberof UsersResponse
     */
    nextPageValue?: string | null;
    /**
     * 
     * @type {Array<RealmUser>}
     * @memberof UsersResponse
     */
    users?: Array<RealmUser> | null;
}

/**
 * Check if a given object implements the UsersResponse interface.
 */
export function instanceOfUsersResponse(value: object): value is UsersResponse {
    return true;
}

export function UsersResponseFromJSON(json: any): UsersResponse {
    return UsersResponseFromJSONTyped(json, false);
}

export function UsersResponseFromJSONTyped(json: any, ignoreDiscriminator: boolean): UsersResponse {
    if (json == null) {
        return json;
    }
    return {
        
        'nextPageValue': json['next_page_value'] == null ? undefined : String(json['next_page_value']),
        'users': json['users'] == null ? undefined : ((json['users'] as Array<any>).map(RealmUserFromJSON)),
    };
}

export function UsersResponseToJSON(json: any): UsersResponse {
    return UsersResponseToJSONTyped(json, false);
}

export function UsersResponseToJSONTyped(value?: UsersResponse | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'next_page_value': value['nextPageValue'],
        'users': value['users'] == null ? undefined : ((value['users'] as Array<any>).map(RealmUserToJSON)),
    };
}

