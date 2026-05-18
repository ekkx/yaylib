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
import type { InvitedUserDTO } from './InvitedUserDTO.js';
import {
    InvitedUserDTOFromJSON,
    InvitedUserDTOFromJSONTyped,
    InvitedUserDTOToJSON,
    InvitedUserDTOToJSONTyped,
} from './InvitedUserDTO.js';

/**
 * 
 * @export
 * @interface CampaignInvitedUsersResponse
 */
export interface CampaignInvitedUsersResponse {
    /**
     * 
     * @type {Array<InvitedUserDTO>}
     * @memberof CampaignInvitedUsersResponse
     */
    invitedUsers?: Array<InvitedUserDTO> | null;
    /**
     * 
     * @type {number}
     * @memberof CampaignInvitedUsersResponse
     */
    totalPoints?: number | null;
}

/**
 * Check if a given object implements the CampaignInvitedUsersResponse interface.
 */
export function instanceOfCampaignInvitedUsersResponse(value: object): value is CampaignInvitedUsersResponse {
    return true;
}

export function CampaignInvitedUsersResponseFromJSON(json: any): CampaignInvitedUsersResponse {
    return CampaignInvitedUsersResponseFromJSONTyped(json, false);
}

export function CampaignInvitedUsersResponseFromJSONTyped(json: any, ignoreDiscriminator: boolean): CampaignInvitedUsersResponse {
    if (json == null) {
        return json;
    }
    return {
        
        'invitedUsers': json['invited_users'] == null ? undefined : ((json['invited_users'] as Array<any>).map(InvitedUserDTOFromJSON)),
        'totalPoints': json['total_points'] == null ? undefined : json['total_points'],
    };
}

export function CampaignInvitedUsersResponseToJSON(json: any): CampaignInvitedUsersResponse {
    return CampaignInvitedUsersResponseToJSONTyped(json, false);
}

export function CampaignInvitedUsersResponseToJSONTyped(value?: CampaignInvitedUsersResponse | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'invited_users': value['invitedUsers'] == null ? undefined : ((value['invitedUsers'] as Array<any>).map(InvitedUserDTOToJSON)),
        'total_points': value['totalPoints'],
    };
}

