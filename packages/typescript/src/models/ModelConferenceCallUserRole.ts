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
 * @interface ModelConferenceCallUserRole
 */
export interface ModelConferenceCallUserRole {
    /**
     * 
     * @type {number}
     * @memberof ModelConferenceCallUserRole
     */
    id?: number | null;
    /**
     * 
     * @type {object}
     * @memberof ModelConferenceCallUserRole
     */
    role?: object | null;
}

/**
 * Check if a given object implements the ModelConferenceCallUserRole interface.
 */
export function instanceOfModelConferenceCallUserRole(value: object): value is ModelConferenceCallUserRole {
    return true;
}

export function ModelConferenceCallUserRoleFromJSON(json: any): ModelConferenceCallUserRole {
    return ModelConferenceCallUserRoleFromJSONTyped(json, false);
}

export function ModelConferenceCallUserRoleFromJSONTyped(json: any, ignoreDiscriminator: boolean): ModelConferenceCallUserRole {
    if (json == null) {
        return json;
    }
    return {
        
        'id': json['id'] == null ? undefined : json['id'],
        'role': json['role'] == null ? undefined : json['role'],
    };
}

export function ModelConferenceCallUserRoleToJSON(json: any): ModelConferenceCallUserRole {
    return ModelConferenceCallUserRoleToJSONTyped(json, false);
}

export function ModelConferenceCallUserRoleToJSONTyped(value?: ModelConferenceCallUserRole | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'id': value['id'],
        'role': value['role'],
    };
}

