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
 * @interface Role
 */
export interface Role {
    /**
     * 
     * @type {boolean}
     * @memberof Role
     */
    isModerator?: boolean | null;
    /**
     * 
     * @type {boolean}
     * @memberof Role
     */
    isOwner?: boolean | null;
}

/**
 * Check if a given object implements the Role interface.
 */
export function instanceOfRole(value: object): value is Role {
    return true;
}

export function RoleFromJSON(json: any): Role {
    return RoleFromJSONTyped(json, false);
}

export function RoleFromJSONTyped(json: any, ignoreDiscriminator: boolean): Role {
    if (json == null) {
        return json;
    }
    return {
        
        'isModerator': json['is_moderator'] == null ? undefined : json['is_moderator'],
        'isOwner': json['is_owner'] == null ? undefined : json['is_owner'],
    };
}

export function RoleToJSON(json: any): Role {
    return RoleToJSONTyped(json, false);
}

export function RoleToJSONTyped(value?: Role | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'is_moderator': value['isModerator'],
        'is_owner': value['isOwner'],
    };
}

