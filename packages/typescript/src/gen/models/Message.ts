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
import type { Type } from './Type.js';
import {
    TypeFromJSON,
    TypeFromJSONTyped,
    TypeToJSON,
    TypeToJSONTyped,
} from './Type.js';
import type { ChatInvitation } from './ChatInvitation.js';
import {
    ChatInvitationFromJSON,
    ChatInvitationFromJSONTyped,
    ChatInvitationToJSON,
    ChatInvitationToJSONTyped,
} from './ChatInvitation.js';
import type { ConferenceCall } from './ConferenceCall.js';
import {
    ConferenceCallFromJSON,
    ConferenceCallFromJSONTyped,
    ConferenceCallToJSON,
    ConferenceCallToJSONTyped,
} from './ConferenceCall.js';

/**
 * 
 * @export
 * @interface Message
 */
export interface Message {
    /**
     * 
     * @type {ConferenceCall}
     * @memberof Message
     */
    conferenceCall?: ConferenceCall | null;
    /**
     * 
     * @type {number}
     * @memberof Message
     */
    id?: number | null;
    /**
     * 
     * @type {ChatInvitation}
     * @memberof Message
     */
    invitation?: ChatInvitation | null;
    /**
     * 
     * @type {string}
     * @memberof Message
     */
    text?: string | null;
    /**
     * 
     * @type {Type}
     * @memberof Message
     */
    type?: Type | null;
    /**
     * 
     * @type {number}
     * @memberof Message
     */
    userId?: number | null;
}

/**
 * Check if a given object implements the Message interface.
 */
export function instanceOfMessage(value: object): value is Message {
    return true;
}

export function MessageFromJSON(json: any): Message {
    return MessageFromJSONTyped(json, false);
}

export function MessageFromJSONTyped(json: any, ignoreDiscriminator: boolean): Message {
    if (json == null) {
        return json;
    }
    return {
        
        'conferenceCall': json['conference_call'] == null ? undefined : ConferenceCallFromJSON(json['conference_call']),
        'id': json['id'] == null ? undefined : json['id'],
        'invitation': json['invitation'] == null ? undefined : ChatInvitationFromJSON(json['invitation']),
        'text': json['text'] == null ? undefined : json['text'],
        'type': json['type'] == null ? undefined : TypeFromJSON(json['type']),
        'userId': json['user_id'] == null ? undefined : json['user_id'],
    };
}

export function MessageToJSON(json: any): Message {
    return MessageToJSONTyped(json, false);
}

export function MessageToJSONTyped(value?: Message | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'conference_call': ConferenceCallToJSON(value['conferenceCall']),
        'id': value['id'],
        'invitation': ChatInvitationToJSON(value['invitation']),
        'text': value['text'],
        'type': TypeToJSON(value['type']),
        'user_id': value['userId'],
    };
}

