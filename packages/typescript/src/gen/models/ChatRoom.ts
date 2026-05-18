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
import type { User } from './User.js';
import {
    UserFromJSON,
    UserFromJSONTyped,
    UserToJSON,
    UserToJSONTyped,
} from './User.js';
import type { Message } from './Message.js';
import {
    MessageFromJSON,
    MessageFromJSONTyped,
    MessageToJSON,
    MessageToJSONTyped,
} from './Message.js';

/**
 * 
 * @export
 * @interface ChatRoom
 */
export interface ChatRoom {
    /**
     * 
     * @type {string}
     * @memberof ChatRoom
     */
    background?: string | null;
    /**
     * 
     * @type {number}
     * @memberof ChatRoom
     */
    id?: number | null;
    /**
     * 
     * @type {boolean}
     * @memberof ChatRoom
     */
    isGroup?: boolean | null;
    /**
     * 
     * @type {boolean}
     * @memberof ChatRoom
     */
    isNotificationOn?: boolean | null;
    /**
     * 
     * @type {boolean}
     * @memberof ChatRoom
     */
    isPinned?: boolean | null;
    /**
     * 
     * @type {boolean}
     * @memberof ChatRoom
     */
    isRequest?: boolean | null;
    /**
     * 
     * @type {Message}
     * @memberof ChatRoom
     */
    lastMessage?: Message | null;
    /**
     * 
     * @type {Array<User>}
     * @memberof ChatRoom
     */
    members?: Array<User> | null;
    /**
     * 
     * @type {string}
     * @memberof ChatRoom
     */
    name?: string | null;
    /**
     * 
     * @type {User}
     * @memberof ChatRoom
     */
    owner?: User | null;
    /**
     * 
     * @type {number}
     * @memberof ChatRoom
     */
    unreadCount?: number | null;
    /**
     * 
     * @type {number}
     * @memberof ChatRoom
     */
    updatedAtMillis?: number | null;
}

/**
 * Check if a given object implements the ChatRoom interface.
 */
export function instanceOfChatRoom(value: object): value is ChatRoom {
    return true;
}

export function ChatRoomFromJSON(json: any): ChatRoom {
    return ChatRoomFromJSONTyped(json, false);
}

export function ChatRoomFromJSONTyped(json: any, ignoreDiscriminator: boolean): ChatRoom {
    if (json == null) {
        return json;
    }
    return {
        
        'background': json['background'] == null ? undefined : json['background'],
        'id': json['id'] == null ? undefined : json['id'],
        'isGroup': json['is_group'] == null ? undefined : json['is_group'],
        'isNotificationOn': json['is_notification_on'] == null ? undefined : json['is_notification_on'],
        'isPinned': json['is_pinned'] == null ? undefined : json['is_pinned'],
        'isRequest': json['is_request'] == null ? undefined : json['is_request'],
        'lastMessage': json['last_message'] == null ? undefined : MessageFromJSON(json['last_message']),
        'members': json['members'] == null ? undefined : ((json['members'] as Array<any>).map(UserFromJSON)),
        'name': json['name'] == null ? undefined : json['name'],
        'owner': json['owner'] == null ? undefined : UserFromJSON(json['owner']),
        'unreadCount': json['unread_count'] == null ? undefined : json['unread_count'],
        'updatedAtMillis': json['updated_at_millis'] == null ? undefined : json['updated_at_millis'],
    };
}

export function ChatRoomToJSON(json: any): ChatRoom {
    return ChatRoomToJSONTyped(json, false);
}

export function ChatRoomToJSONTyped(value?: ChatRoom | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'background': value['background'],
        'id': value['id'],
        'is_group': value['isGroup'],
        'is_notification_on': value['isNotificationOn'],
        'is_pinned': value['isPinned'],
        'is_request': value['isRequest'],
        'last_message': MessageToJSON(value['lastMessage']),
        'members': value['members'] == null ? undefined : ((value['members'] as Array<any>).map(UserToJSON)),
        'name': value['name'],
        'owner': UserToJSON(value['owner']),
        'unread_count': value['unreadCount'],
        'updated_at_millis': value['updatedAtMillis'],
    };
}

