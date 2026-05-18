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
import type { Group } from './Group.js';
import {
    GroupFromJSON,
    GroupFromJSONTyped,
    GroupToJSON,
    GroupToJSONTyped,
} from './Group.js';
import type { RealmGift } from './RealmGift.js';
import {
    RealmGiftFromJSON,
    RealmGiftFromJSONTyped,
    RealmGiftToJSON,
    RealmGiftToJSONTyped,
} from './RealmGift.js';
import type { Post } from './Post.js';
import {
    PostFromJSON,
    PostFromJSONTyped,
    PostToJSON,
    PostToJSONTyped,
} from './Post.js';
import type { Metadata } from './Metadata.js';
import {
    MetadataFromJSON,
    MetadataFromJSONTyped,
    MetadataToJSON,
    MetadataToJSONTyped,
} from './Metadata.js';
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
 * @interface Activity
 */
export interface Activity {
    /**
     * 
     * @type {Array<RealmUser>}
     * @memberof Activity
     */
    birthdayUsers?: Array<RealmUser> | null;
    /**
     * 
     * @type {number}
     * @memberof Activity
     */
    birthdayUsersCount?: number | null;
    /**
     * 
     * @type {number}
     * @memberof Activity
     */
    createdAt?: number | null;
    /**
     * 
     * @type {number}
     * @memberof Activity
     */
    emplAmount?: number | null;
    /**
     * 
     * @type {Array<RealmUser>}
     * @memberof Activity
     */
    followers?: Array<RealmUser> | null;
    /**
     * 
     * @type {number}
     * @memberof Activity
     */
    followersCount?: number | null;
    /**
     * 
     * @type {Post}
     * @memberof Activity
     */
    fromPost?: Post | null;
    /**
     * 
     * @type {Array<number>}
     * @memberof Activity
     */
    fromPostIds?: Array<number> | null;
    /**
     * 
     * @type {RealmGift}
     * @memberof Activity
     */
    giftItem?: RealmGift | null;
    /**
     * 
     * @type {Group}
     * @memberof Activity
     */
    group?: Group | null;
    /**
     * 
     * @type {number}
     * @memberof Activity
     */
    id?: number | null;
    /**
     * 
     * @type {Metadata}
     * @memberof Activity
     */
    metadata?: Metadata | null;
    /**
     * 
     * @type {Post}
     * @memberof Activity
     */
    toPost?: Post | null;
    /**
     * 
     * @type {string}
     * @memberof Activity
     */
    type?: string | null;
    /**
     * 
     * @type {RealmUser}
     * @memberof Activity
     */
    user?: RealmUser | null;
    /**
     * 
     * @type {number}
     * @memberof Activity
     */
    vipReward?: number | null;
}

/**
 * Check if a given object implements the Activity interface.
 */
export function instanceOfActivity(value: object): value is Activity {
    return true;
}

export function ActivityFromJSON(json: any): Activity {
    return ActivityFromJSONTyped(json, false);
}

export function ActivityFromJSONTyped(json: any, ignoreDiscriminator: boolean): Activity {
    if (json == null) {
        return json;
    }
    return {
        
        'birthdayUsers': json['birthday_users'] == null ? undefined : ((json['birthday_users'] as Array<any>).map(RealmUserFromJSON)),
        'birthdayUsersCount': json['birthday_users_count'] == null ? undefined : json['birthday_users_count'],
        'createdAt': json['created_at'] == null ? undefined : json['created_at'],
        'emplAmount': json['empl_amount'] == null ? undefined : json['empl_amount'],
        'followers': json['followers'] == null ? undefined : ((json['followers'] as Array<any>).map(RealmUserFromJSON)),
        'followersCount': json['followers_count'] == null ? undefined : json['followers_count'],
        'fromPost': json['from_post'] == null ? undefined : PostFromJSON(json['from_post']),
        'fromPostIds': json['from_post_ids'] == null ? undefined : json['from_post_ids'],
        'giftItem': json['gift_item'] == null ? undefined : RealmGiftFromJSON(json['gift_item']),
        'group': json['group'] == null ? undefined : GroupFromJSON(json['group']),
        'id': json['id'] == null ? undefined : json['id'],
        'metadata': json['metadata'] == null ? undefined : MetadataFromJSON(json['metadata']),
        'toPost': json['to_post'] == null ? undefined : PostFromJSON(json['to_post']),
        'type': json['type'] == null ? undefined : json['type'],
        'user': json['user'] == null ? undefined : RealmUserFromJSON(json['user']),
        'vipReward': json['vip_reward'] == null ? undefined : json['vip_reward'],
    };
}

export function ActivityToJSON(json: any): Activity {
    return ActivityToJSONTyped(json, false);
}

export function ActivityToJSONTyped(value?: Activity | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'birthday_users': value['birthdayUsers'] == null ? undefined : ((value['birthdayUsers'] as Array<any>).map(RealmUserToJSON)),
        'birthday_users_count': value['birthdayUsersCount'],
        'created_at': value['createdAt'],
        'empl_amount': value['emplAmount'],
        'followers': value['followers'] == null ? undefined : ((value['followers'] as Array<any>).map(RealmUserToJSON)),
        'followers_count': value['followersCount'],
        'from_post': PostToJSON(value['fromPost']),
        'from_post_ids': value['fromPostIds'],
        'gift_item': RealmGiftToJSON(value['giftItem']),
        'group': GroupToJSON(value['group']),
        'id': value['id'],
        'metadata': MetadataToJSON(value['metadata']),
        'to_post': PostToJSON(value['toPost']),
        'type': value['type'],
        'user': RealmUserToJSON(value['user']),
        'vip_reward': value['vipReward'],
    };
}

