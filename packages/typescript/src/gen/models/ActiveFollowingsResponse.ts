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
 * @interface ActiveFollowingsResponse
 */
export interface ActiveFollowingsResponse {
    /**
     * 
     * @type {number}
     * @memberof ActiveFollowingsResponse
     */
    lastLoggedinAt?: number | null;
    /**
     * 
     * @type {Array<RealmUser>}
     * @memberof ActiveFollowingsResponse
     */
    users?: Array<RealmUser> | null;
}

/**
 * Check if a given object implements the ActiveFollowingsResponse interface.
 */
export function instanceOfActiveFollowingsResponse(value: object): value is ActiveFollowingsResponse {
    return true;
}

export function ActiveFollowingsResponseFromJSON(json: any): ActiveFollowingsResponse {
    return ActiveFollowingsResponseFromJSONTyped(json, false);
}

export function ActiveFollowingsResponseFromJSONTyped(json: any, ignoreDiscriminator: boolean): ActiveFollowingsResponse {
    if (json == null) {
        return json;
    }
    return {
        
        'lastLoggedinAt': json['last_loggedin_at'] == null ? undefined : json['last_loggedin_at'],
        'users': json['users'] == null ? undefined : ((json['users'] as Array<any>).map(RealmUserFromJSON)),
    };
}

export function ActiveFollowingsResponseToJSON(json: any): ActiveFollowingsResponse {
    return ActiveFollowingsResponseToJSONTyped(json, false);
}

export function ActiveFollowingsResponseToJSONTyped(value?: ActiveFollowingsResponse | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'last_loggedin_at': value['lastLoggedinAt'],
        'users': value['users'] == null ? undefined : ((value['users'] as Array<any>).map(RealmUserToJSON)),
    };
}

