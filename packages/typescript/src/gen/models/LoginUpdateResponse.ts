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
/**
 * 
 * @export
 * @interface LoginUpdateResponse
 */
export interface LoginUpdateResponse {
    /**
     * 
     * @type {string}
     * @memberof LoginUpdateResponse
     */
    accessToken?: string | null;
    /**
     * 
     * @type {number}
     * @memberof LoginUpdateResponse
     */
    expiresIn?: number | null;
    /**
     * 
     * @type {string}
     * @memberof LoginUpdateResponse
     */
    refreshToken?: string | null;
    /**
     * 
     * @type {number}
     * @memberof LoginUpdateResponse
     */
    userId?: number | null;
    /**
     * 
     * @type {string}
     * @memberof LoginUpdateResponse
     */
    username?: string | null;
}

/**
 * Check if a given object implements the LoginUpdateResponse interface.
 */
export function instanceOfLoginUpdateResponse(value: object): value is LoginUpdateResponse {
    return true;
}

export function LoginUpdateResponseFromJSON(json: any): LoginUpdateResponse {
    return LoginUpdateResponseFromJSONTyped(json, false);
}

export function LoginUpdateResponseFromJSONTyped(json: any, ignoreDiscriminator: boolean): LoginUpdateResponse {
    if (json == null) {
        return json;
    }
    return {
        
        'accessToken': json['access_token'] == null ? undefined : json['access_token'],
        'expiresIn': json['expires_in'] == null ? undefined : json['expires_in'],
        'refreshToken': json['refresh_token'] == null ? undefined : json['refresh_token'],
        'userId': json['user_id'] == null ? undefined : json['user_id'],
        'username': json['username'] == null ? undefined : json['username'],
    };
}

export function LoginUpdateResponseToJSON(json: any): LoginUpdateResponse {
    return LoginUpdateResponseToJSONTyped(json, false);
}

export function LoginUpdateResponseToJSONTyped(value?: LoginUpdateResponse | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'access_token': value['accessToken'],
        'expires_in': value['expiresIn'],
        'refresh_token': value['refreshToken'],
        'user_id': value['userId'],
        'username': value['username'],
    };
}

