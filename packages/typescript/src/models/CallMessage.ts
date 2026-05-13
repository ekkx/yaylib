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
 * @interface CallMessage
 */
export interface CallMessage {
    /**
     * 
     * @type {number}
     * @memberof CallMessage
     */
    createdAtMillis?: number | null;
    /**
     * 
     * @type {string}
     * @memberof CallMessage
     */
    id?: string | null;
    /**
     * 
     * @type {boolean}
     * @memberof CallMessage
     */
    isMessageDeleted?: boolean | null;
    /**
     * 
     * @type {string}
     * @memberof CallMessage
     */
    text?: string | null;
    /**
     * 
     * @type {string}
     * @memberof CallMessage
     */
    userId?: string | null;
}

/**
 * Check if a given object implements the CallMessage interface.
 */
export function instanceOfCallMessage(value: object): value is CallMessage {
    return true;
}

export function CallMessageFromJSON(json: any): CallMessage {
    return CallMessageFromJSONTyped(json, false);
}

export function CallMessageFromJSONTyped(json: any, ignoreDiscriminator: boolean): CallMessage {
    if (json == null) {
        return json;
    }
    return {
        
        'createdAtMillis': json['created_at_millis'] == null ? undefined : json['created_at_millis'],
        'id': json['id'] == null ? undefined : json['id'],
        'isMessageDeleted': json['is_message_deleted'] == null ? undefined : json['is_message_deleted'],
        'text': json['text'] == null ? undefined : json['text'],
        'userId': json['user_id'] == null ? undefined : json['user_id'],
    };
}

export function CallMessageToJSON(json: any): CallMessage {
    return CallMessageToJSONTyped(json, false);
}

export function CallMessageToJSONTyped(value?: CallMessage | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'created_at_millis': value['createdAtMillis'],
        'id': value['id'],
        'is_message_deleted': value['isMessageDeleted'],
        'text': value['text'],
        'user_id': value['userId'],
    };
}

