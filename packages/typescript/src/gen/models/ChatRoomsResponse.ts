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
 * @interface ChatRoomsResponse
 */
export interface ChatRoomsResponse {
    /**
     * 
     * @type {Array<RealmChatRoom>}
     * @memberof ChatRoomsResponse
     */
    chatRooms?: Array<RealmChatRoom> | null;
    /**
     * 
     * @type {string}
     * @memberof ChatRoomsResponse
     */
    nextPageValue?: string | null;
    /**
     * 
     * @type {Array<RealmChatRoom>}
     * @memberof ChatRoomsResponse
     */
    pinnedChatRooms?: Array<RealmChatRoom> | null;
}

/**
 * Check if a given object implements the ChatRoomsResponse interface.
 */
export function instanceOfChatRoomsResponse(value: object): value is ChatRoomsResponse {
    return true;
}

export function ChatRoomsResponseFromJSON(json: any): ChatRoomsResponse {
    return ChatRoomsResponseFromJSONTyped(json, false);
}

export function ChatRoomsResponseFromJSONTyped(json: any, ignoreDiscriminator: boolean): ChatRoomsResponse {
    if (json == null) {
        return json;
    }
    return {
        
        'chatRooms': json['chat_rooms'] == null ? undefined : ((json['chat_rooms'] as Array<any>).map(RealmChatRoomFromJSON)),
        'nextPageValue': json['next_page_value'] == null ? undefined : String(json['next_page_value']),
        'pinnedChatRooms': json['pinned_chat_rooms'] == null ? undefined : ((json['pinned_chat_rooms'] as Array<any>).map(RealmChatRoomFromJSON)),
    };
}

export function ChatRoomsResponseToJSON(json: any): ChatRoomsResponse {
    return ChatRoomsResponseToJSONTyped(json, false);
}

export function ChatRoomsResponseToJSONTyped(value?: ChatRoomsResponse | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'chat_rooms': value['chatRooms'] == null ? undefined : ((value['chatRooms'] as Array<any>).map(RealmChatRoomToJSON)),
        'next_page_value': value['nextPageValue'],
        'pinned_chat_rooms': value['pinnedChatRooms'] == null ? undefined : ((value['pinnedChatRooms'] as Array<any>).map(RealmChatRoomToJSON)),
    };
}

