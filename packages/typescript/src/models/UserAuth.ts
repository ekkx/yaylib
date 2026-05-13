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
/**
 * 
 * @export
 * @interface UserAuth
 */
export interface UserAuth {
    /**
     * 
     * @type {string}
     * @memberof UserAuth
     */
    accessToken?: string | null;
    /**
     * 
     * @type {number}
     * @memberof UserAuth
     */
    expiresIn?: number | null;
    /**
     * 
     * @type {string}
     * @memberof UserAuth
     */
    refreshToken?: string | null;
    /**
     * 
     * @type {number}
     * @memberof UserAuth
     */
    userId?: number | null;
}

/**
 * Check if a given object implements the UserAuth interface.
 */
export function instanceOfUserAuth(value: object): value is UserAuth {
    return true;
}

export function UserAuthFromJSON(json: any): UserAuth {
    return UserAuthFromJSONTyped(json, false);
}

export function UserAuthFromJSONTyped(json: any, ignoreDiscriminator: boolean): UserAuth {
    if (json == null) {
        return json;
    }
    return {
        
        'accessToken': json['access_token'] == null ? undefined : json['access_token'],
        'expiresIn': json['expires_in'] == null ? undefined : json['expires_in'],
        'refreshToken': json['refresh_token'] == null ? undefined : json['refresh_token'],
        'userId': json['user_id'] == null ? undefined : json['user_id'],
    };
}

export function UserAuthToJSON(json: any): UserAuth {
    return UserAuthToJSONTyped(json, false);
}

export function UserAuthToJSONTyped(value?: UserAuth | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'access_token': value['accessToken'],
        'expires_in': value['expiresIn'],
        'refresh_token': value['refreshToken'],
        'user_id': value['userId'],
    };
}

