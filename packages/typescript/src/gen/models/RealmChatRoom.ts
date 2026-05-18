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
import type { ChatRoomLastMessage } from './ChatRoomLastMessage.js';
import {
    ChatRoomLastMessageFromJSON,
    ChatRoomLastMessageFromJSONTyped,
    ChatRoomLastMessageToJSON,
    ChatRoomLastMessageToJSONTyped,
} from './ChatRoomLastMessage.js';
import type { RealmUser } from './RealmUser.js';
import {
    RealmUserFromJSON,
    RealmUserFromJSONTyped,
    RealmUserToJSON,
    RealmUserToJSONTyped,
} from './RealmUser.js';
import type { UserSetting } from './UserSetting.js';
import {
    UserSettingFromJSON,
    UserSettingFromJSONTyped,
    UserSettingToJSON,
    UserSettingToJSONTyped,
} from './UserSetting.js';

/**
 * 
 * @export
 * @interface RealmChatRoom
 */
export interface RealmChatRoom {
    /**
     * 
     * @type {string}
     * @memberof RealmChatRoom
     */
    background?: string | null;
    /**
     * 
     * @type {number}
     * @memberof RealmChatRoom
     */
    id?: number | null;
    /**
     * 
     * @type {boolean}
     * @memberof RealmChatRoom
     */
    isGroup?: boolean | null;
    /**
     * 
     * @type {boolean}
     * @memberof RealmChatRoom
     */
    isRequest?: boolean | null;
    /**
     * 
     * @type {ChatRoomLastMessage}
     * @memberof RealmChatRoom
     */
    lastMessage?: ChatRoomLastMessage | null;
    /**
     * 
     * @type {Array<RealmUser>}
     * @memberof RealmChatRoom
     */
    members?: Array<RealmUser> | null;
    /**
     * 
     * @type {string}
     * @memberof RealmChatRoom
     */
    name?: string | null;
    /**
     * 
     * @type {RealmUser}
     * @memberof RealmChatRoom
     */
    owner?: RealmUser | null;
    /**
     * 
     * @type {number}
     * @memberof RealmChatRoom
     */
    unreadCount?: number | null;
    /**
     * 
     * @type {number}
     * @memberof RealmChatRoom
     */
    updatedAt?: number | null;
    /**
     * 
     * @type {UserSetting}
     * @memberof RealmChatRoom
     */
    userSetting?: UserSetting | null;
}

/**
 * Check if a given object implements the RealmChatRoom interface.
 */
export function instanceOfRealmChatRoom(value: object): value is RealmChatRoom {
    return true;
}

export function RealmChatRoomFromJSON(json: any): RealmChatRoom {
    return RealmChatRoomFromJSONTyped(json, false);
}

export function RealmChatRoomFromJSONTyped(json: any, ignoreDiscriminator: boolean): RealmChatRoom {
    if (json == null) {
        return json;
    }
    return {
        
        'background': json['background'] == null ? undefined : json['background'],
        'id': json['id'] == null ? undefined : json['id'],
        'isGroup': json['is_group'] == null ? undefined : json['is_group'],
        'isRequest': json['is_request'] == null ? undefined : json['is_request'],
        'lastMessage': json['last_message'] == null ? undefined : ChatRoomLastMessageFromJSON(json['last_message']),
        'members': json['members'] == null ? undefined : ((json['members'] as Array<any>).map(RealmUserFromJSON)),
        'name': json['name'] == null ? undefined : json['name'],
        'owner': json['owner'] == null ? undefined : RealmUserFromJSON(json['owner']),
        'unreadCount': json['unread_count'] == null ? undefined : json['unread_count'],
        'updatedAt': json['updated_at'] == null ? undefined : json['updated_at'],
        'userSetting': json['user_setting'] == null ? undefined : UserSettingFromJSON(json['user_setting']),
    };
}

export function RealmChatRoomToJSON(json: any): RealmChatRoom {
    return RealmChatRoomToJSONTyped(json, false);
}

export function RealmChatRoomToJSONTyped(value?: RealmChatRoom | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'background': value['background'],
        'id': value['id'],
        'is_group': value['isGroup'],
        'is_request': value['isRequest'],
        'last_message': ChatRoomLastMessageToJSON(value['lastMessage']),
        'members': value['members'] == null ? undefined : ((value['members'] as Array<any>).map(RealmUserToJSON)),
        'name': value['name'],
        'owner': RealmUserToJSON(value['owner']),
        'unread_count': value['unreadCount'],
        'updated_at': value['updatedAt'],
        'user_setting': UserSettingToJSON(value['userSetting']),
    };
}

