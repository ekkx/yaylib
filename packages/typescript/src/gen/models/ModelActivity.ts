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
import type { Gift } from './Gift';
import {
    GiftFromJSON,
    GiftFromJSONTyped,
    GiftToJSON,
    GiftToJSONTyped,
} from './Gift';
import type { ModelGroup } from './ModelGroup';
import {
    ModelGroupFromJSON,
    ModelGroupFromJSONTyped,
    ModelGroupToJSON,
    ModelGroupToJSONTyped,
} from './ModelGroup';
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
 * @interface ModelActivity
 */
export interface ModelActivity {
    /**
     * 
     * @type {Array<User>}
     * @memberof ModelActivity
     */
    birthdayUsers?: Array<User> | null;
    /**
     * 
     * @type {number}
     * @memberof ModelActivity
     */
    birthdayUsersCount?: number | null;
    /**
     * 
     * @type {string}
     * @memberof ModelActivity
     */
    body?: string | null;
    /**
     * 
     * @type {string}
     * @memberof ModelActivity
     */
    contentPreview?: string | null;
    /**
     * 
     * @type {number}
     * @memberof ModelActivity
     */
    createdAtMillis?: number | null;
    /**
     * 
     * @type {number}
     * @memberof ModelActivity
     */
    emplAmount?: number | null;
    /**
     * 
     * @type {Array<User>}
     * @memberof ModelActivity
     */
    followers?: Array<User> | null;
    /**
     * 
     * @type {number}
     * @memberof ModelActivity
     */
    followersCount?: number | null;
    /**
     * 
     * @type {ModelPost}
     * @memberof ModelActivity
     */
    fromPost?: ModelPost | null;
    /**
     * 
     * @type {Array<number>}
     * @memberof ModelActivity
     */
    fromPostIds?: Array<number> | null;
    /**
     * 
     * @type {Gift}
     * @memberof ModelActivity
     */
    giftItem?: Gift | null;
    /**
     * 
     * @type {ModelGroup}
     * @memberof ModelActivity
     */
    group?: ModelGroup | null;
    /**
     * 
     * @type {string}
     * @memberof ModelActivity
     */
    groupTopic?: string | null;
    /**
     * 
     * @type {number}
     * @memberof ModelActivity
     */
    id?: number | null;
    /**
     * 
     * @type {boolean}
     * @memberof ModelActivity
     */
    isAnonymous?: boolean | null;
    /**
     * 
     * @type {boolean}
     * @memberof ModelActivity
     */
    isBulkInvitation?: boolean | null;
    /**
     * 
     * @type {string}
     * @memberof ModelActivity
     */
    title?: string | null;
    /**
     * 
     * @type {ModelPost}
     * @memberof ModelActivity
     */
    toPost?: ModelPost | null;
    /**
     * 
     * @type {object}
     * @memberof ModelActivity
     */
    type?: object | null;
    /**
     * 
     * @type {string}
     * @memberof ModelActivity
     */
    url?: string | null;
    /**
     * 
     * @type {User}
     * @memberof ModelActivity
     */
    user?: User | null;
    /**
     * 
     * @type {number}
     * @memberof ModelActivity
     */
    vipReward?: number | null;
}

/**
 * Check if a given object implements the ModelActivity interface.
 */
export function instanceOfModelActivity(value: object): value is ModelActivity {
    return true;
}

export function ModelActivityFromJSON(json: any): ModelActivity {
    return ModelActivityFromJSONTyped(json, false);
}

export function ModelActivityFromJSONTyped(json: any, ignoreDiscriminator: boolean): ModelActivity {
    if (json == null) {
        return json;
    }
    return {
        
        'birthdayUsers': json['birthday_users'] == null ? undefined : ((json['birthday_users'] as Array<any>).map(UserFromJSON)),
        'birthdayUsersCount': json['birthday_users_count'] == null ? undefined : json['birthday_users_count'],
        'body': json['body'] == null ? undefined : json['body'],
        'contentPreview': json['content_preview'] == null ? undefined : json['content_preview'],
        'createdAtMillis': json['created_at_millis'] == null ? undefined : json['created_at_millis'],
        'emplAmount': json['empl_amount'] == null ? undefined : json['empl_amount'],
        'followers': json['followers'] == null ? undefined : ((json['followers'] as Array<any>).map(UserFromJSON)),
        'followersCount': json['followers_count'] == null ? undefined : json['followers_count'],
        'fromPost': json['from_post'] == null ? undefined : ModelPostFromJSON(json['from_post']),
        'fromPostIds': json['from_post_ids'] == null ? undefined : json['from_post_ids'],
        'giftItem': json['gift_item'] == null ? undefined : GiftFromJSON(json['gift_item']),
        'group': json['group'] == null ? undefined : ModelGroupFromJSON(json['group']),
        'groupTopic': json['group_topic'] == null ? undefined : json['group_topic'],
        'id': json['id'] == null ? undefined : json['id'],
        'isAnonymous': json['is_anonymous'] == null ? undefined : json['is_anonymous'],
        'isBulkInvitation': json['is_bulk_invitation'] == null ? undefined : json['is_bulk_invitation'],
        'title': json['title'] == null ? undefined : json['title'],
        'toPost': json['to_post'] == null ? undefined : ModelPostFromJSON(json['to_post']),
        'type': json['type'] == null ? undefined : json['type'],
        'url': json['url'] == null ? undefined : json['url'],
        'user': json['user'] == null ? undefined : UserFromJSON(json['user']),
        'vipReward': json['vip_reward'] == null ? undefined : json['vip_reward'],
    };
}

export function ModelActivityToJSON(json: any): ModelActivity {
    return ModelActivityToJSONTyped(json, false);
}

export function ModelActivityToJSONTyped(value?: ModelActivity | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'birthday_users': value['birthdayUsers'] == null ? undefined : ((value['birthdayUsers'] as Array<any>).map(UserToJSON)),
        'birthday_users_count': value['birthdayUsersCount'],
        'body': value['body'],
        'content_preview': value['contentPreview'],
        'created_at_millis': value['createdAtMillis'],
        'empl_amount': value['emplAmount'],
        'followers': value['followers'] == null ? undefined : ((value['followers'] as Array<any>).map(UserToJSON)),
        'followers_count': value['followersCount'],
        'from_post': ModelPostToJSON(value['fromPost']),
        'from_post_ids': value['fromPostIds'],
        'gift_item': GiftToJSON(value['giftItem']),
        'group': ModelGroupToJSON(value['group']),
        'group_topic': value['groupTopic'],
        'id': value['id'],
        'is_anonymous': value['isAnonymous'],
        'is_bulk_invitation': value['isBulkInvitation'],
        'title': value['title'],
        'to_post': ModelPostToJSON(value['toPost']),
        'type': value['type'],
        'url': value['url'],
        'user': UserToJSON(value['user']),
        'vip_reward': value['vipReward'],
    };
}

