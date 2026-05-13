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
import type { InvitedUserDetailsDTO } from './InvitedUserDetailsDTO';
import {
    InvitedUserDetailsDTOFromJSON,
    InvitedUserDetailsDTOFromJSONTyped,
    InvitedUserDetailsDTOToJSON,
    InvitedUserDetailsDTOToJSONTyped,
} from './InvitedUserDetailsDTO';

/**
 * 
 * @export
 * @interface InvitedUserDTO
 */
export interface InvitedUserDTO {
    /**
     * 
     * @type {InvitedUserDetailsDTO}
     * @memberof InvitedUserDTO
     */
    referred?: InvitedUserDetailsDTO | null;
    /**
     * 
     * @type {number}
     * @memberof InvitedUserDTO
     */
    totalEarnedPoints?: number | null;
}

/**
 * Check if a given object implements the InvitedUserDTO interface.
 */
export function instanceOfInvitedUserDTO(value: object): value is InvitedUserDTO {
    return true;
}

export function InvitedUserDTOFromJSON(json: any): InvitedUserDTO {
    return InvitedUserDTOFromJSONTyped(json, false);
}

export function InvitedUserDTOFromJSONTyped(json: any, ignoreDiscriminator: boolean): InvitedUserDTO {
    if (json == null) {
        return json;
    }
    return {
        
        'referred': json['referred'] == null ? undefined : InvitedUserDetailsDTOFromJSON(json['referred']),
        'totalEarnedPoints': json['total_earned_points'] == null ? undefined : json['total_earned_points'],
    };
}

export function InvitedUserDTOToJSON(json: any): InvitedUserDTO {
    return InvitedUserDTOToJSONTyped(json, false);
}

export function InvitedUserDTOToJSONTyped(value?: InvitedUserDTO | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'referred': InvitedUserDetailsDTOToJSON(value['referred']),
        'total_earned_points': value['totalEarnedPoints'],
    };
}

