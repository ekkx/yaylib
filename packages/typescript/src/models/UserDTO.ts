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
 * @interface UserDTO
 */
export interface UserDTO {
    /**
     * 
     * @type {number}
     * @memberof UserDTO
     */
    id?: number | null;
    /**
     * 
     * @type {string}
     * @memberof UserDTO
     */
    nickname?: string | null;
    /**
     * 
     * @type {string}
     * @memberof UserDTO
     */
    profileIcon?: string | null;
}

/**
 * Check if a given object implements the UserDTO interface.
 */
export function instanceOfUserDTO(value: object): value is UserDTO {
    return true;
}

export function UserDTOFromJSON(json: any): UserDTO {
    return UserDTOFromJSONTyped(json, false);
}

export function UserDTOFromJSONTyped(json: any, ignoreDiscriminator: boolean): UserDTO {
    if (json == null) {
        return json;
    }
    return {
        
        'id': json['id'] == null ? undefined : json['id'],
        'nickname': json['nickname'] == null ? undefined : json['nickname'],
        'profileIcon': json['profile_icon'] == null ? undefined : json['profile_icon'],
    };
}

export function UserDTOToJSON(json: any): UserDTO {
    return UserDTOToJSONTyped(json, false);
}

export function UserDTOToJSONTyped(value?: UserDTO | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'id': value['id'],
        'nickname': value['nickname'],
        'profile_icon': value['profileIcon'],
    };
}

