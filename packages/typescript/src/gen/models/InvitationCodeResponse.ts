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
import type { CodeResponse } from './CodeResponse.js';
import {
    CodeResponseFromJSON,
    CodeResponseFromJSONTyped,
    CodeResponseToJSON,
    CodeResponseToJSONTyped,
} from './CodeResponse.js';

/**
 * 
 * @export
 * @interface InvitationCodeResponse
 */
export interface InvitationCodeResponse {
    /**
     * 
     * @type {CodeResponse}
     * @memberof InvitationCodeResponse
     */
    invitationCode?: CodeResponse | null;
}

/**
 * Check if a given object implements the InvitationCodeResponse interface.
 */
export function instanceOfInvitationCodeResponse(value: object): value is InvitationCodeResponse {
    return true;
}

export function InvitationCodeResponseFromJSON(json: any): InvitationCodeResponse {
    return InvitationCodeResponseFromJSONTyped(json, false);
}

export function InvitationCodeResponseFromJSONTyped(json: any, ignoreDiscriminator: boolean): InvitationCodeResponse {
    if (json == null) {
        return json;
    }
    return {
        
        'invitationCode': json['invitation_code'] == null ? undefined : CodeResponseFromJSON(json['invitation_code']),
    };
}

export function InvitationCodeResponseToJSON(json: any): InvitationCodeResponse {
    return InvitationCodeResponseToJSONTyped(json, false);
}

export function InvitationCodeResponseToJSONTyped(value?: InvitationCodeResponse | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'invitation_code': CodeResponseToJSON(value['invitationCode']),
    };
}

