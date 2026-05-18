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
 * @interface GroupRole
 */
export interface GroupRole {
    /**
     * 
     * @type {string}
     * @memberof GroupRole
     */
    apiValue?: string | null;
}

/**
 * Check if a given object implements the GroupRole interface.
 */
export function instanceOfGroupRole(value: object): value is GroupRole {
    return true;
}

export function GroupRoleFromJSON(json: any): GroupRole {
    return GroupRoleFromJSONTyped(json, false);
}

export function GroupRoleFromJSONTyped(json: any, ignoreDiscriminator: boolean): GroupRole {
    if (json == null) {
        return json;
    }
    return {
        
        'apiValue': json['api_value'] == null ? undefined : json['api_value'],
    };
}

export function GroupRoleToJSON(json: any): GroupRole {
    return GroupRoleToJSONTyped(json, false);
}

export function GroupRoleToJSONTyped(value?: GroupRole | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'api_value': value['apiValue'],
    };
}

