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
import type { RealmConferenceCall } from './RealmConferenceCall';
import {
    RealmConferenceCallFromJSON,
    RealmConferenceCallFromJSONTyped,
    RealmConferenceCallToJSON,
    RealmConferenceCallToJSONTyped,
} from './RealmConferenceCall';
import type { MessageType } from './MessageType';
import {
    MessageTypeFromJSON,
    MessageTypeFromJSONTyped,
    MessageTypeToJSON,
    MessageTypeToJSONTyped,
} from './MessageType';

/**
 * 
 * @export
 * @interface ChatRoomLastMessage
 */
export interface ChatRoomLastMessage {
    /**
     * 
     * @type {RealmConferenceCall}
     * @memberof ChatRoomLastMessage
     */
    conferenceCall?: RealmConferenceCall | null;
    /**
     * 
     * @type {number}
     * @memberof ChatRoomLastMessage
     */
    createdAt?: number | null;
    /**
     * 
     * @type {number}
     * @memberof ChatRoomLastMessage
     */
    id?: number | null;
    /**
     * 
     * @type {MessageType}
     * @memberof ChatRoomLastMessage
     */
    messageType?: MessageType | null;
    /**
     * 
     * @type {number}
     * @memberof ChatRoomLastMessage
     */
    roomId?: number | null;
    /**
     * 
     * @type {string}
     * @memberof ChatRoomLastMessage
     */
    text?: string | null;
    /**
     * 
     * @type {number}
     * @memberof ChatRoomLastMessage
     */
    userId?: number | null;
}



/**
 * Check if a given object implements the ChatRoomLastMessage interface.
 */
export function instanceOfChatRoomLastMessage(value: object): value is ChatRoomLastMessage {
    return true;
}

export function ChatRoomLastMessageFromJSON(json: any): ChatRoomLastMessage {
    return ChatRoomLastMessageFromJSONTyped(json, false);
}

export function ChatRoomLastMessageFromJSONTyped(json: any, ignoreDiscriminator: boolean): ChatRoomLastMessage {
    if (json == null) {
        return json;
    }
    return {
        
        'conferenceCall': json['conference_call'] == null ? undefined : RealmConferenceCallFromJSON(json['conference_call']),
        'createdAt': json['created_at'] == null ? undefined : json['created_at'],
        'id': json['id'] == null ? undefined : json['id'],
        'messageType': json['message_type'] == null ? undefined : MessageTypeFromJSON(json['message_type']),
        'roomId': json['room_id'] == null ? undefined : json['room_id'],
        'text': json['text'] == null ? undefined : json['text'],
        'userId': json['user_id'] == null ? undefined : json['user_id'],
    };
}

export function ChatRoomLastMessageToJSON(json: any): ChatRoomLastMessage {
    return ChatRoomLastMessageToJSONTyped(json, false);
}

export function ChatRoomLastMessageToJSONTyped(value?: ChatRoomLastMessage | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'conference_call': RealmConferenceCallToJSON(value['conferenceCall']),
        'created_at': value['createdAt'],
        'id': value['id'],
        'message_type': MessageTypeToJSON(value['messageType']),
        'room_id': value['roomId'],
        'text': value['text'],
        'user_id': value['userId'],
    };
}

