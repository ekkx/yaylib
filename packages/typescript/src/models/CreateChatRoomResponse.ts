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
/**
 * 
 * @export
 * @interface CreateChatRoomResponse
 */
export interface CreateChatRoomResponse {
    /**
     * 
     * @type {number}
     * @memberof CreateChatRoomResponse
     */
    roomId?: number | null;
}

/**
 * Check if a given object implements the CreateChatRoomResponse interface.
 */
export function instanceOfCreateChatRoomResponse(value: object): value is CreateChatRoomResponse {
    return true;
}

export function CreateChatRoomResponseFromJSON(json: any): CreateChatRoomResponse {
    return CreateChatRoomResponseFromJSONTyped(json, false);
}

export function CreateChatRoomResponseFromJSONTyped(json: any, ignoreDiscriminator: boolean): CreateChatRoomResponse {
    if (json == null) {
        return json;
    }
    return {
        
        'roomId': json['room_id'] == null ? undefined : json['room_id'],
    };
}

export function CreateChatRoomResponseToJSON(json: any): CreateChatRoomResponse {
    return CreateChatRoomResponseToJSONTyped(json, false);
}

export function CreateChatRoomResponseToJSONTyped(value?: CreateChatRoomResponse | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'room_id': value['roomId'],
    };
}

