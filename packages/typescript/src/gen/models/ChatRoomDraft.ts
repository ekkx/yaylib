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
 * @interface ChatRoomDraft
 */
export interface ChatRoomDraft {
    /**
     * 
     * @type {number}
     * @memberof ChatRoomDraft
     */
    id?: number | null;
    /**
     * 
     * @type {string}
     * @memberof ChatRoomDraft
     */
    text?: string | null;
}

/**
 * Check if a given object implements the ChatRoomDraft interface.
 */
export function instanceOfChatRoomDraft(value: object): value is ChatRoomDraft {
    return true;
}

export function ChatRoomDraftFromJSON(json: any): ChatRoomDraft {
    return ChatRoomDraftFromJSONTyped(json, false);
}

export function ChatRoomDraftFromJSONTyped(json: any, ignoreDiscriminator: boolean): ChatRoomDraft {
    if (json == null) {
        return json;
    }
    return {
        
        'id': json['id'] == null ? undefined : json['id'],
        'text': json['text'] == null ? undefined : json['text'],
    };
}

export function ChatRoomDraftToJSON(json: any): ChatRoomDraft {
    return ChatRoomDraftToJSONTyped(json, false);
}

export function ChatRoomDraftToJSONTyped(value?: ChatRoomDraft | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'id': value['id'],
        'text': value['text'],
    };
}

