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
 * @interface UsersByTimestampResponse
 */
export interface UsersByTimestampResponse {
    /**
     * 
     * @type {number}
     * @memberof UsersByTimestampResponse
     */
    lastTimestamp?: number | null;
    /**
     * 
     * @type {Array<RealmUser>}
     * @memberof UsersByTimestampResponse
     */
    users?: Array<RealmUser> | null;
}

/**
 * Check if a given object implements the UsersByTimestampResponse interface.
 */
export function instanceOfUsersByTimestampResponse(value: object): value is UsersByTimestampResponse {
    return true;
}

export function UsersByTimestampResponseFromJSON(json: any): UsersByTimestampResponse {
    return UsersByTimestampResponseFromJSONTyped(json, false);
}

export function UsersByTimestampResponseFromJSONTyped(json: any, ignoreDiscriminator: boolean): UsersByTimestampResponse {
    if (json == null) {
        return json;
    }
    return {
        
        'lastTimestamp': json['last_timestamp'] == null ? undefined : json['last_timestamp'],
        'users': json['users'] == null ? undefined : ((json['users'] as Array<any>).map(RealmUserFromJSON)),
    };
}

export function UsersByTimestampResponseToJSON(json: any): UsersByTimestampResponse {
    return UsersByTimestampResponseToJSONTyped(json, false);
}

export function UsersByTimestampResponseToJSONTyped(value?: UsersByTimestampResponse | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'last_timestamp': value['lastTimestamp'],
        'users': value['users'] == null ? undefined : ((value['users'] as Array<any>).map(RealmUserToJSON)),
    };
}

