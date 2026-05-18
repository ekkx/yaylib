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
import type { RealmChatRoom } from './RealmChatRoom.js';
import {
    RealmChatRoomFromJSON,
    RealmChatRoomFromJSONTyped,
    RealmChatRoomToJSON,
    RealmChatRoomToJSONTyped,
} from './RealmChatRoom.js';

/**
 * 
 * @export
 * @interface ChatRoomResponse
 */
export interface ChatRoomResponse {
    /**
     * 
     * @type {RealmChatRoom}
     * @memberof ChatRoomResponse
     */
    chat?: RealmChatRoom | null;
}

/**
 * Check if a given object implements the ChatRoomResponse interface.
 */
export function instanceOfChatRoomResponse(value: object): value is ChatRoomResponse {
    return true;
}

export function ChatRoomResponseFromJSON(json: any): ChatRoomResponse {
    return ChatRoomResponseFromJSONTyped(json, false);
}

export function ChatRoomResponseFromJSONTyped(json: any, ignoreDiscriminator: boolean): ChatRoomResponse {
    if (json == null) {
        return json;
    }
    return {
        
        'chat': json['chat'] == null ? undefined : RealmChatRoomFromJSON(json['chat']),
    };
}

export function ChatRoomResponseToJSON(json: any): ChatRoomResponse {
    return ChatRoomResponseToJSONTyped(json, false);
}

export function ChatRoomResponseToJSONTyped(value?: ChatRoomResponse | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'chat': RealmChatRoomToJSON(value['chat']),
    };
}

