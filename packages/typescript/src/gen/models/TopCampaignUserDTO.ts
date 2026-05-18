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
import type { UserDTO } from './UserDTO.js';
import {
    UserDTOFromJSON,
    UserDTOFromJSONTyped,
    UserDTOToJSON,
    UserDTOToJSONTyped,
} from './UserDTO.js';

/**
 * 
 * @export
 * @interface TopCampaignUserDTO
 */
export interface TopCampaignUserDTO {
    /**
     * 
     * @type {number}
     * @memberof TopCampaignUserDTO
     */
    points?: number | null;
    /**
     * 
     * @type {number}
     * @memberof TopCampaignUserDTO
     */
    ranking?: number | null;
    /**
     * 
     * @type {UserDTO}
     * @memberof TopCampaignUserDTO
     */
    user?: UserDTO | null;
}

/**
 * Check if a given object implements the TopCampaignUserDTO interface.
 */
export function instanceOfTopCampaignUserDTO(value: object): value is TopCampaignUserDTO {
    return true;
}

export function TopCampaignUserDTOFromJSON(json: any): TopCampaignUserDTO {
    return TopCampaignUserDTOFromJSONTyped(json, false);
}

export function TopCampaignUserDTOFromJSONTyped(json: any, ignoreDiscriminator: boolean): TopCampaignUserDTO {
    if (json == null) {
        return json;
    }
    return {
        
        'points': json['points'] == null ? undefined : json['points'],
        'ranking': json['ranking'] == null ? undefined : json['ranking'],
        'user': json['user'] == null ? undefined : UserDTOFromJSON(json['user']),
    };
}

export function TopCampaignUserDTOToJSON(json: any): TopCampaignUserDTO {
    return TopCampaignUserDTOToJSONTyped(json, false);
}

export function TopCampaignUserDTOToJSONTyped(value?: TopCampaignUserDTO | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'points': value['points'],
        'ranking': value['ranking'],
        'user': UserDTOToJSON(value['user']),
    };
}

