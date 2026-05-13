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
 * @interface CreateUserResponse
 */
export interface CreateUserResponse {
    /**
     * 
     * @type {string}
     * @memberof CreateUserResponse
     */
    accessToken?: string | null;
    /**
     * 
     * @type {number}
     * @memberof CreateUserResponse
     */
    expiresIn?: number | null;
    /**
     * 
     * @type {number}
     * @memberof CreateUserResponse
     */
    id?: number | null;
    /**
     * 
     * @type {string}
     * @memberof CreateUserResponse
     */
    refreshToken?: string | null;
}

/**
 * Check if a given object implements the CreateUserResponse interface.
 */
export function instanceOfCreateUserResponse(value: object): value is CreateUserResponse {
    return true;
}

export function CreateUserResponseFromJSON(json: any): CreateUserResponse {
    return CreateUserResponseFromJSONTyped(json, false);
}

export function CreateUserResponseFromJSONTyped(json: any, ignoreDiscriminator: boolean): CreateUserResponse {
    if (json == null) {
        return json;
    }
    return {
        
        'accessToken': json['access_token'] == null ? undefined : json['access_token'],
        'expiresIn': json['expires_in'] == null ? undefined : json['expires_in'],
        'id': json['id'] == null ? undefined : json['id'],
        'refreshToken': json['refresh_token'] == null ? undefined : json['refresh_token'],
    };
}

export function CreateUserResponseToJSON(json: any): CreateUserResponse {
    return CreateUserResponseToJSONTyped(json, false);
}

export function CreateUserResponseToJSONTyped(value?: CreateUserResponse | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'access_token': value['accessToken'],
        'expires_in': value['expiresIn'],
        'id': value['id'],
        'refresh_token': value['refreshToken'],
    };
}

