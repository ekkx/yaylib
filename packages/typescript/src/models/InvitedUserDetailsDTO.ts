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
 * @interface InvitedUserDetailsDTO
 */
export interface InvitedUserDetailsDTO {
    /**
     * 
     * @type {number}
     * @memberof InvitedUserDetailsDTO
     */
    id?: number | null;
    /**
     * 
     * @type {string}
     * @memberof InvitedUserDetailsDTO
     */
    nickname?: string | null;
    /**
     * 
     * @type {string}
     * @memberof InvitedUserDetailsDTO
     */
    profileIcon?: string | null;
}

/**
 * Check if a given object implements the InvitedUserDetailsDTO interface.
 */
export function instanceOfInvitedUserDetailsDTO(value: object): value is InvitedUserDetailsDTO {
    return true;
}

export function InvitedUserDetailsDTOFromJSON(json: any): InvitedUserDetailsDTO {
    return InvitedUserDetailsDTOFromJSONTyped(json, false);
}

export function InvitedUserDetailsDTOFromJSONTyped(json: any, ignoreDiscriminator: boolean): InvitedUserDetailsDTO {
    if (json == null) {
        return json;
    }
    return {
        
        'id': json['id'] == null ? undefined : json['id'],
        'nickname': json['nickname'] == null ? undefined : json['nickname'],
        'profileIcon': json['profile_icon'] == null ? undefined : json['profile_icon'],
    };
}

export function InvitedUserDetailsDTOToJSON(json: any): InvitedUserDetailsDTO {
    return InvitedUserDetailsDTOToJSONTyped(json, false);
}

export function InvitedUserDetailsDTOToJSONTyped(value?: InvitedUserDetailsDTO | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'id': value['id'],
        'nickname': value['nickname'],
        'profile_icon': value['profileIcon'],
    };
}

