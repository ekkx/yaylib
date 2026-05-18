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
 * @interface InvitationCode
 */
export interface InvitationCode {
    /**
     * 
     * @type {number}
     * @memberof InvitationCode
     */
    possibleToEditAtMillis?: number | null;
    /**
     * 
     * @type {string}
     * @memberof InvitationCode
     */
    value?: string | null;
}

/**
 * Check if a given object implements the InvitationCode interface.
 */
export function instanceOfInvitationCode(value: object): value is InvitationCode {
    return true;
}

export function InvitationCodeFromJSON(json: any): InvitationCode {
    return InvitationCodeFromJSONTyped(json, false);
}

export function InvitationCodeFromJSONTyped(json: any, ignoreDiscriminator: boolean): InvitationCode {
    if (json == null) {
        return json;
    }
    return {
        
        'possibleToEditAtMillis': json['possible_to_edit_at_millis'] == null ? undefined : json['possible_to_edit_at_millis'],
        'value': json['value'] == null ? undefined : json['value'],
    };
}

export function InvitationCodeToJSON(json: any): InvitationCode {
    return InvitationCodeToJSONTyped(json, false);
}

export function InvitationCodeToJSONTyped(value?: InvitationCode | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'possible_to_edit_at_millis': value['possibleToEditAtMillis'],
        'value': value['value'],
    };
}

