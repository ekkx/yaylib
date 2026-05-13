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
 * @interface TokenResponse
 */
export interface TokenResponse {
    /**
     * 
     * @type {string}
     * @memberof TokenResponse
     */
    accessToken?: string | null;
    /**
     * 
     * @type {number}
     * @memberof TokenResponse
     */
    createdAt?: number | null;
    /**
     * 
     * @type {number}
     * @memberof TokenResponse
     */
    expiresIn?: number | null;
    /**
     * 
     * @type {number}
     * @memberof TokenResponse
     */
    id?: number | null;
    /**
     * 
     * @type {string}
     * @memberof TokenResponse
     */
    refreshToken?: string | null;
}

/**
 * Check if a given object implements the TokenResponse interface.
 */
export function instanceOfTokenResponse(value: object): value is TokenResponse {
    return true;
}

export function TokenResponseFromJSON(json: any): TokenResponse {
    return TokenResponseFromJSONTyped(json, false);
}

export function TokenResponseFromJSONTyped(json: any, ignoreDiscriminator: boolean): TokenResponse {
    if (json == null) {
        return json;
    }
    return {
        
        'accessToken': json['access_token'] == null ? undefined : json['access_token'],
        'createdAt': json['created_at'] == null ? undefined : json['created_at'],
        'expiresIn': json['expires_in'] == null ? undefined : json['expires_in'],
        'id': json['id'] == null ? undefined : json['id'],
        'refreshToken': json['refresh_token'] == null ? undefined : json['refresh_token'],
    };
}

export function TokenResponseToJSON(json: any): TokenResponse {
    return TokenResponseToJSONTyped(json, false);
}

export function TokenResponseToJSONTyped(value?: TokenResponse | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'access_token': value['accessToken'],
        'created_at': value['createdAt'],
        'expires_in': value['expiresIn'],
        'id': value['id'],
        'refresh_token': value['refreshToken'],
    };
}

