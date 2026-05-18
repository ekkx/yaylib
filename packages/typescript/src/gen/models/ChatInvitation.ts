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
 * @interface ChatInvitation
 */
export interface ChatInvitation {
    /**
     * 
     * @type {string}
     * @memberof ChatInvitation
     */
    firstUser?: string | null;
    /**
     * 
     * @type {string}
     * @memberof ChatInvitation
     */
    inviter?: string | null;
    /**
     * 
     * @type {number}
     * @memberof ChatInvitation
     */
    usersCount?: number | null;
}

/**
 * Check if a given object implements the ChatInvitation interface.
 */
export function instanceOfChatInvitation(value: object): value is ChatInvitation {
    return true;
}

export function ChatInvitationFromJSON(json: any): ChatInvitation {
    return ChatInvitationFromJSONTyped(json, false);
}

export function ChatInvitationFromJSONTyped(json: any, ignoreDiscriminator: boolean): ChatInvitation {
    if (json == null) {
        return json;
    }
    return {
        
        'firstUser': json['first_user'] == null ? undefined : json['first_user'],
        'inviter': json['inviter'] == null ? undefined : json['inviter'],
        'usersCount': json['users_count'] == null ? undefined : json['users_count'],
    };
}

export function ChatInvitationToJSON(json: any): ChatInvitation {
    return ChatInvitationToJSONTyped(json, false);
}

export function ChatInvitationToJSONTyped(value?: ChatInvitation | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'first_user': value['firstUser'],
        'inviter': value['inviter'],
        'users_count': value['usersCount'],
    };
}

