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
 * @interface LoginEmailUserRequest
 */
export interface LoginEmailUserRequest {
    /**
     * 
     * @type {string}
     * @memberof LoginEmailUserRequest
     */
    apiKey?: string | null;
    /**
     * 
     * @type {string}
     * @memberof LoginEmailUserRequest
     */
    email?: string | null;
    /**
     * 
     * @type {string}
     * @memberof LoginEmailUserRequest
     */
    password?: string | null;
    /**
     * 
     * @type {string}
     * @memberof LoginEmailUserRequest
     */
    twoFACode?: string | null;
    /**
     * 
     * @type {string}
     * @memberof LoginEmailUserRequest
     */
    uuid?: string | null;
}

/**
 * Check if a given object implements the LoginEmailUserRequest interface.
 */
export function instanceOfLoginEmailUserRequest(value: object): value is LoginEmailUserRequest {
    return true;
}

export function LoginEmailUserRequestFromJSON(json: any): LoginEmailUserRequest {
    return LoginEmailUserRequestFromJSONTyped(json, false);
}

export function LoginEmailUserRequestFromJSONTyped(json: any, ignoreDiscriminator: boolean): LoginEmailUserRequest {
    if (json == null) {
        return json;
    }
    return {
        
        'apiKey': json['api_key'] == null ? undefined : json['api_key'],
        'email': json['email'] == null ? undefined : json['email'],
        'password': json['password'] == null ? undefined : json['password'],
        'twoFACode': json['two_f_a_code'] == null ? undefined : json['two_f_a_code'],
        'uuid': json['uuid'] == null ? undefined : json['uuid'],
    };
}

export function LoginEmailUserRequestToJSON(json: any): LoginEmailUserRequest {
    return LoginEmailUserRequestToJSONTyped(json, false);
}

export function LoginEmailUserRequestToJSONTyped(value?: LoginEmailUserRequest | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'api_key': value['apiKey'],
        'email': value['email'],
        'password': value['password'],
        'two_f_a_code': value['twoFACode'],
        'uuid': value['uuid'],
    };
}

