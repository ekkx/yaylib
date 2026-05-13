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
import type { ModelPost } from './ModelPost';
import {
    ModelPostFromJSON,
    ModelPostFromJSONTyped,
    ModelPostToJSON,
    ModelPostToJSONTyped,
} from './ModelPost';
import type { User } from './User';
import {
    UserFromJSON,
    UserFromJSONTyped,
    UserToJSON,
    UserToJSONTyped,
} from './User';

/**
 * 
 * @export
 * @interface ModelThreadInfo
 */
export interface ModelThreadInfo {
    /**
     * 
     * @type {number}
     * @memberof ModelThreadInfo
     */
    createdAtSeconds?: number | null;
    /**
     * 
     * @type {boolean}
     * @memberof ModelThreadInfo
     */
    hasNewUpdates?: boolean | null;
    /**
     * 
     * @type {number}
     * @memberof ModelThreadInfo
     */
    id?: number | null;
    /**
     * 
     * @type {boolean}
     * @memberof ModelThreadInfo
     */
    isJoined?: boolean | null;
    /**
     * 
     * @type {ModelPost}
     * @memberof ModelThreadInfo
     */
    lastPost?: ModelPost | null;
    /**
     * 
     * @type {User}
     * @memberof ModelThreadInfo
     */
    owner?: User | null;
    /**
     * 
     * @type {number}
     * @memberof ModelThreadInfo
     */
    postsCount?: number | null;
    /**
     * 
     * @type {string}
     * @memberof ModelThreadInfo
     */
    threadIcon?: string | null;
    /**
     * 
     * @type {string}
     * @memberof ModelThreadInfo
     */
    title?: string | null;
    /**
     * 
     * @type {number}
     * @memberof ModelThreadInfo
     */
    unreadCount?: number | null;
    /**
     * 
     * @type {number}
     * @memberof ModelThreadInfo
     */
    updatedAtSeconds?: number | null;
}

/**
 * Check if a given object implements the ModelThreadInfo interface.
 */
export function instanceOfModelThreadInfo(value: object): value is ModelThreadInfo {
    return true;
}

export function ModelThreadInfoFromJSON(json: any): ModelThreadInfo {
    return ModelThreadInfoFromJSONTyped(json, false);
}

export function ModelThreadInfoFromJSONTyped(json: any, ignoreDiscriminator: boolean): ModelThreadInfo {
    if (json == null) {
        return json;
    }
    return {
        
        'createdAtSeconds': json['created_at_seconds'] == null ? undefined : json['created_at_seconds'],
        'hasNewUpdates': json['has_new_updates'] == null ? undefined : json['has_new_updates'],
        'id': json['id'] == null ? undefined : json['id'],
        'isJoined': json['is_joined'] == null ? undefined : json['is_joined'],
        'lastPost': json['last_post'] == null ? undefined : ModelPostFromJSON(json['last_post']),
        'owner': json['owner'] == null ? undefined : UserFromJSON(json['owner']),
        'postsCount': json['posts_count'] == null ? undefined : json['posts_count'],
        'threadIcon': json['thread_icon'] == null ? undefined : json['thread_icon'],
        'title': json['title'] == null ? undefined : json['title'],
        'unreadCount': json['unread_count'] == null ? undefined : json['unread_count'],
        'updatedAtSeconds': json['updated_at_seconds'] == null ? undefined : json['updated_at_seconds'],
    };
}

export function ModelThreadInfoToJSON(json: any): ModelThreadInfo {
    return ModelThreadInfoToJSONTyped(json, false);
}

export function ModelThreadInfoToJSONTyped(value?: ModelThreadInfo | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'created_at_seconds': value['createdAtSeconds'],
        'has_new_updates': value['hasNewUpdates'],
        'id': value['id'],
        'is_joined': value['isJoined'],
        'last_post': ModelPostToJSON(value['lastPost']),
        'owner': UserToJSON(value['owner']),
        'posts_count': value['postsCount'],
        'thread_icon': value['threadIcon'],
        'title': value['title'],
        'unread_count': value['unreadCount'],
        'updated_at_seconds': value['updatedAtSeconds'],
    };
}

