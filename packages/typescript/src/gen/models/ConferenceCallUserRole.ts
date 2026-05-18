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
 * @interface ConferenceCallUserRole
 */
export interface ConferenceCallUserRole {
    /**
     * 
     * @type {number}
     * @memberof ConferenceCallUserRole
     */
    id?: number | null;
    /**
     * 
     * @type {string}
     * @memberof ConferenceCallUserRole
     */
    role?: string | null;
    /**
     * 
     * @type {number}
     * @memberof ConferenceCallUserRole
     */
    userId?: number | null;
}

/**
 * Check if a given object implements the ConferenceCallUserRole interface.
 */
export function instanceOfConferenceCallUserRole(value: object): value is ConferenceCallUserRole {
    return true;
}

export function ConferenceCallUserRoleFromJSON(json: any): ConferenceCallUserRole {
    return ConferenceCallUserRoleFromJSONTyped(json, false);
}

export function ConferenceCallUserRoleFromJSONTyped(json: any, ignoreDiscriminator: boolean): ConferenceCallUserRole {
    if (json == null) {
        return json;
    }
    return {
        
        'id': json['id'] == null ? undefined : json['id'],
        'role': json['role'] == null ? undefined : json['role'],
        'userId': json['user_id'] == null ? undefined : json['user_id'],
    };
}

export function ConferenceCallUserRoleToJSON(json: any): ConferenceCallUserRole {
    return ConferenceCallUserRoleToJSONTyped(json, false);
}

export function ConferenceCallUserRoleToJSONTyped(value?: ConferenceCallUserRole | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'id': value['id'],
        'role': value['role'],
        'user_id': value['userId'],
    };
}

