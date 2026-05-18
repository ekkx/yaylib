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
import type { Post } from './Post.js';
import {
    PostFromJSON,
    PostFromJSONTyped,
    PostToJSON,
    PostToJSONTyped,
} from './Post.js';
import type { RealmUser } from './RealmUser.js';
import {
    RealmUserFromJSON,
    RealmUserFromJSONTyped,
    RealmUserToJSON,
    RealmUserToJSONTyped,
} from './RealmUser.js';

/**
 * 
 * @export
 * @interface ThreadInfo
 */
export interface ThreadInfo {
    /**
     * 
     * @type {number}
     * @memberof ThreadInfo
     */
    createdAt?: number | null;
    /**
     * 
     * @type {number}
     * @memberof ThreadInfo
     */
    id?: number | null;
    /**
     * 
     * @type {boolean}
     * @memberof ThreadInfo
     */
    isJoined?: boolean | null;
    /**
     * 
     * @type {Post}
     * @memberof ThreadInfo
     */
    lastPost?: Post | null;
    /**
     * 
     * @type {boolean}
     * @memberof ThreadInfo
     */
    newUpdates?: boolean | null;
    /**
     * 
     * @type {RealmUser}
     * @memberof ThreadInfo
     */
    owner?: RealmUser | null;
    /**
     * 
     * @type {number}
     * @memberof ThreadInfo
     */
    postsCount?: number | null;
    /**
     * 
     * @type {string}
     * @memberof ThreadInfo
     */
    threadIcon?: string | null;
    /**
     * 
     * @type {string}
     * @memberof ThreadInfo
     */
    title?: string | null;
    /**
     * 
     * @type {number}
     * @memberof ThreadInfo
     */
    unreadCount?: number | null;
    /**
     * 
     * @type {number}
     * @memberof ThreadInfo
     */
    updatedAt?: number | null;
}

/**
 * Check if a given object implements the ThreadInfo interface.
 */
export function instanceOfThreadInfo(value: object): value is ThreadInfo {
    return true;
}

export function ThreadInfoFromJSON(json: any): ThreadInfo {
    return ThreadInfoFromJSONTyped(json, false);
}

export function ThreadInfoFromJSONTyped(json: any, ignoreDiscriminator: boolean): ThreadInfo {
    if (json == null) {
        return json;
    }
    return {
        
        'createdAt': json['created_at'] == null ? undefined : json['created_at'],
        'id': json['id'] == null ? undefined : json['id'],
        'isJoined': json['is_joined'] == null ? undefined : json['is_joined'],
        'lastPost': json['last_post'] == null ? undefined : PostFromJSON(json['last_post']),
        'newUpdates': json['new_updates'] == null ? undefined : json['new_updates'],
        'owner': json['owner'] == null ? undefined : RealmUserFromJSON(json['owner']),
        'postsCount': json['posts_count'] == null ? undefined : json['posts_count'],
        'threadIcon': json['thread_icon'] == null ? undefined : json['thread_icon'],
        'title': json['title'] == null ? undefined : json['title'],
        'unreadCount': json['unread_count'] == null ? undefined : json['unread_count'],
        'updatedAt': json['updated_at'] == null ? undefined : json['updated_at'],
    };
}

export function ThreadInfoToJSON(json: any): ThreadInfo {
    return ThreadInfoToJSONTyped(json, false);
}

export function ThreadInfoToJSONTyped(value?: ThreadInfo | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'created_at': value['createdAt'],
        'id': value['id'],
        'is_joined': value['isJoined'],
        'last_post': PostToJSON(value['lastPost']),
        'new_updates': value['newUpdates'],
        'owner': RealmUserToJSON(value['owner']),
        'posts_count': value['postsCount'],
        'thread_icon': value['threadIcon'],
        'title': value['title'],
        'unread_count': value['unreadCount'],
        'updated_at': value['updatedAt'],
    };
}

