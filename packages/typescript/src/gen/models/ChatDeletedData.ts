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
 * @interface ChatDeletedData
 */
export interface ChatDeletedData {
    /**
     * 
     * @type {number}
     * @memberof ChatDeletedData
     */
    roomId?: number | null;
}

/**
 * Check if a given object implements the ChatDeletedData interface.
 */
export function instanceOfChatDeletedData(value: object): value is ChatDeletedData {
    return true;
}

export function ChatDeletedDataFromJSON(json: any): ChatDeletedData {
    return ChatDeletedDataFromJSONTyped(json, false);
}

export function ChatDeletedDataFromJSONTyped(json: any, ignoreDiscriminator: boolean): ChatDeletedData {
    if (json == null) {
        return json;
    }
    return {
        
        'roomId': json['room_id'] == null ? undefined : json['room_id'],
    };
}

export function ChatDeletedDataToJSON(json: any): ChatDeletedData {
    return ChatDeletedDataToJSONTyped(json, false);
}

export function ChatDeletedDataToJSONTyped(value?: ChatDeletedData | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'room_id': value['roomId'],
    };
}

