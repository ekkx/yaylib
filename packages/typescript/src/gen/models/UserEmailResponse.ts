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
 * @interface UserEmailResponse
 */
export interface UserEmailResponse {
    /**
     * 
     * @type {string}
     * @memberof UserEmailResponse
     */
    email?: string | null;
}

/**
 * Check if a given object implements the UserEmailResponse interface.
 */
export function instanceOfUserEmailResponse(value: object): value is UserEmailResponse {
    return true;
}

export function UserEmailResponseFromJSON(json: any): UserEmailResponse {
    return UserEmailResponseFromJSONTyped(json, false);
}

export function UserEmailResponseFromJSONTyped(json: any, ignoreDiscriminator: boolean): UserEmailResponse {
    if (json == null) {
        return json;
    }
    return {
        
        'email': json['email'] == null ? undefined : json['email'],
    };
}

export function UserEmailResponseToJSON(json: any): UserEmailResponse {
    return UserEmailResponseToJSONTyped(json, false);
}

export function UserEmailResponseToJSONTyped(value?: UserEmailResponse | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'email': value['email'],
    };
}

