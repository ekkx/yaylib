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
 * @interface UserWrapper
 */
export interface UserWrapper {
    /**
     * 
     * @type {number}
     * @memberof UserWrapper
     */
    id?: number | null;
    /**
     * 
     * @type {RealmUser}
     * @memberof UserWrapper
     */
    user?: RealmUser | null;
}

/**
 * Check if a given object implements the UserWrapper interface.
 */
export function instanceOfUserWrapper(value: object): value is UserWrapper {
    return true;
}

export function UserWrapperFromJSON(json: any): UserWrapper {
    return UserWrapperFromJSONTyped(json, false);
}

export function UserWrapperFromJSONTyped(json: any, ignoreDiscriminator: boolean): UserWrapper {
    if (json == null) {
        return json;
    }
    return {
        
        'id': json['id'] == null ? undefined : json['id'],
        'user': json['user'] == null ? undefined : RealmUserFromJSON(json['user']),
    };
}

export function UserWrapperToJSON(json: any): UserWrapper {
    return UserWrapperToJSONTyped(json, false);
}

export function UserWrapperToJSONTyped(value?: UserWrapper | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'id': value['id'],
        'user': RealmUserToJSON(value['user']),
    };
}

