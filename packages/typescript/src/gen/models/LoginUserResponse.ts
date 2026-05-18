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
import type { SnsInfo } from './SnsInfo.js';
import {
    SnsInfoFromJSON,
    SnsInfoFromJSONTyped,
    SnsInfoToJSON,
    SnsInfoToJSONTyped,
} from './SnsInfo.js';

/**
 * 
 * @export
 * @interface LoginUserResponse
 */
export interface LoginUserResponse {
    /**
     * 
     * @type {string}
     * @memberof LoginUserResponse
     */
    accessToken?: string | null;
    /**
     * 
     * @type {number}
     * @memberof LoginUserResponse
     */
    expiresIn?: number | null;
    /**
     * 
     * @type {boolean}
     * @memberof LoginUserResponse
     */
    isNew?: boolean | null;
    /**
     * 
     * @type {string}
     * @memberof LoginUserResponse
     */
    refreshToken?: string | null;
    /**
     * 
     * @type {SnsInfo}
     * @memberof LoginUserResponse
     */
    snsInfo?: SnsInfo | null;
    /**
     * 
     * @type {number}
     * @memberof LoginUserResponse
     */
    userId?: number | null;
    /**
     * 
     * @type {string}
     * @memberof LoginUserResponse
     */
    username?: string | null;
}

/**
 * Check if a given object implements the LoginUserResponse interface.
 */
export function instanceOfLoginUserResponse(value: object): value is LoginUserResponse {
    return true;
}

export function LoginUserResponseFromJSON(json: any): LoginUserResponse {
    return LoginUserResponseFromJSONTyped(json, false);
}

export function LoginUserResponseFromJSONTyped(json: any, ignoreDiscriminator: boolean): LoginUserResponse {
    if (json == null) {
        return json;
    }
    return {
        
        'accessToken': json['access_token'] == null ? undefined : json['access_token'],
        'expiresIn': json['expires_in'] == null ? undefined : json['expires_in'],
        'isNew': json['is_new'] == null ? undefined : json['is_new'],
        'refreshToken': json['refresh_token'] == null ? undefined : json['refresh_token'],
        'snsInfo': json['sns_info'] == null ? undefined : SnsInfoFromJSON(json['sns_info']),
        'userId': json['user_id'] == null ? undefined : json['user_id'],
        'username': json['username'] == null ? undefined : json['username'],
    };
}

export function LoginUserResponseToJSON(json: any): LoginUserResponse {
    return LoginUserResponseToJSONTyped(json, false);
}

export function LoginUserResponseToJSONTyped(value?: LoginUserResponse | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'access_token': value['accessToken'],
        'expires_in': value['expiresIn'],
        'is_new': value['isNew'],
        'refresh_token': value['refreshToken'],
        'sns_info': SnsInfoToJSON(value['snsInfo']),
        'user_id': value['userId'],
        'username': value['username'],
    };
}

