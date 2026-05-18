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
 * @interface LoginSnsUserRequest
 */
export interface LoginSnsUserRequest {
    /**
     * 
     * @type {string}
     * @memberof LoginSnsUserRequest
     */
    accessToken?: string | null;
    /**
     * 
     * @type {string}
     * @memberof LoginSnsUserRequest
     */
    accessTokenSecret?: string | null;
    /**
     * 
     * @type {string}
     * @memberof LoginSnsUserRequest
     */
    apiKey?: string | null;
    /**
     * 
     * @type {string}
     * @memberof LoginSnsUserRequest
     */
    authorizationCode?: string | null;
    /**
     * 
     * @type {string}
     * @memberof LoginSnsUserRequest
     */
    consumerKey?: string | null;
    /**
     * 
     * @type {string}
     * @memberof LoginSnsUserRequest
     */
    consumerSecret?: string | null;
    /**
     * 
     * @type {string}
     * @memberof LoginSnsUserRequest
     */
    email?: string | null;
    /**
     * 
     * @type {string}
     * @memberof LoginSnsUserRequest
     */
    provider?: string | null;
    /**
     * 
     * @type {string}
     * @memberof LoginSnsUserRequest
     */
    twoFACode?: string | null;
}

/**
 * Check if a given object implements the LoginSnsUserRequest interface.
 */
export function instanceOfLoginSnsUserRequest(value: object): value is LoginSnsUserRequest {
    return true;
}

export function LoginSnsUserRequestFromJSON(json: any): LoginSnsUserRequest {
    return LoginSnsUserRequestFromJSONTyped(json, false);
}

export function LoginSnsUserRequestFromJSONTyped(json: any, ignoreDiscriminator: boolean): LoginSnsUserRequest {
    if (json == null) {
        return json;
    }
    return {
        
        'accessToken': json['access_token'] == null ? undefined : json['access_token'],
        'accessTokenSecret': json['access_token_secret'] == null ? undefined : json['access_token_secret'],
        'apiKey': json['api_key'] == null ? undefined : json['api_key'],
        'authorizationCode': json['authorization_code'] == null ? undefined : json['authorization_code'],
        'consumerKey': json['consumer_key'] == null ? undefined : json['consumer_key'],
        'consumerSecret': json['consumer_secret'] == null ? undefined : json['consumer_secret'],
        'email': json['email'] == null ? undefined : json['email'],
        'provider': json['provider'] == null ? undefined : json['provider'],
        'twoFACode': json['two_f_a_code'] == null ? undefined : json['two_f_a_code'],
    };
}

export function LoginSnsUserRequestToJSON(json: any): LoginSnsUserRequest {
    return LoginSnsUserRequestToJSONTyped(json, false);
}

export function LoginSnsUserRequestToJSONTyped(value?: LoginSnsUserRequest | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'access_token': value['accessToken'],
        'access_token_secret': value['accessTokenSecret'],
        'api_key': value['apiKey'],
        'authorization_code': value['authorizationCode'],
        'consumer_key': value['consumerKey'],
        'consumer_secret': value['consumerSecret'],
        'email': value['email'],
        'provider': value['provider'],
        'two_f_a_code': value['twoFACode'],
    };
}

