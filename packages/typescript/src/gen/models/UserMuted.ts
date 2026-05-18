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
import type { User } from './User.js';
import {
    UserFromJSON,
    UserFromJSONTyped,
    UserToJSON,
    UserToJSONTyped,
} from './User.js';

/**
 * 
 * @export
 * @interface UserMuted
 */
export interface UserMuted {
    /**
     * 
     * @type {boolean}
     * @memberof UserMuted
     */
    isMuted?: boolean | null;
    /**
     * 
     * @type {User}
     * @memberof UserMuted
     */
    user?: User | null;
}

/**
 * Check if a given object implements the UserMuted interface.
 */
export function instanceOfUserMuted(value: object): value is UserMuted {
    return true;
}

export function UserMutedFromJSON(json: any): UserMuted {
    return UserMutedFromJSONTyped(json, false);
}

export function UserMutedFromJSONTyped(json: any, ignoreDiscriminator: boolean): UserMuted {
    if (json == null) {
        return json;
    }
    return {
        
        'isMuted': json['is_muted'] == null ? undefined : json['is_muted'],
        'user': json['user'] == null ? undefined : UserFromJSON(json['user']),
    };
}

export function UserMutedToJSON(json: any): UserMuted {
    return UserMutedToJSONTyped(json, false);
}

export function UserMutedToJSONTyped(value?: UserMuted | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'is_muted': value['isMuted'],
        'user': UserToJSON(value['user']),
    };
}

