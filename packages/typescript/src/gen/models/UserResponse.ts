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
import type { RealmUser } from './RealmUser.js';
import {
    RealmUserFromJSON,
    RealmUserFromJSONTyped,
    RealmUserToJSON,
    RealmUserToJSONTyped,
} from './RealmUser.js';
import type { RealmGiftingAbility } from './RealmGiftingAbility.js';
import {
    RealmGiftingAbilityFromJSON,
    RealmGiftingAbilityFromJSONTyped,
    RealmGiftingAbilityToJSON,
    RealmGiftingAbilityToJSONTyped,
} from './RealmGiftingAbility.js';

/**
 * 
 * @export
 * @interface UserResponse
 */
export interface UserResponse {
    /**
     * 
     * @type {boolean}
     * @memberof UserResponse
     */
    appleConnected?: boolean | null;
    /**
     * 
     * @type {string}
     * @memberof UserResponse
     */
    birthDate?: string | null;
    /**
     * 
     * @type {number}
     * @memberof UserResponse
     */
    blockingLimit?: number | null;
    /**
     * 
     * @type {boolean}
     * @memberof UserResponse
     */
    emailConfirmed?: boolean | null;
    /**
     * 
     * @type {boolean}
     * @memberof UserResponse
     */
    facebookConnected?: boolean | null;
    /**
     * 
     * @type {RealmGiftingAbility}
     * @memberof UserResponse
     */
    giftingAbility?: RealmGiftingAbility | null;
    /**
     * 
     * @type {boolean}
     * @memberof UserResponse
     */
    googleConnected?: boolean | null;
    /**
     * 
     * @type {boolean}
     * @memberof UserResponse
     */
    groupPhoneOn?: boolean | null;
    /**
     * 
     * @type {boolean}
     * @memberof UserResponse
     */
    groupVideoOn?: boolean | null;
    /**
     * 
     * @type {number}
     * @memberof UserResponse
     */
    interestsCount?: number | null;
    /**
     * 
     * @type {boolean}
     * @memberof UserResponse
     */
    lineConnected?: boolean | null;
    /**
     * 
     * @type {string}
     * @memberof UserResponse
     */
    maskedEmail?: string | null;
    /**
     * 
     * @type {boolean}
     * @memberof UserResponse
     */
    phoneOn?: boolean | null;
    /**
     * 
     * @type {boolean}
     * @memberof UserResponse
     */
    pushNotification?: boolean | null;
    /**
     * 
     * @type {string}
     * @memberof UserResponse
     */
    twitterId?: string | null;
    /**
     * 
     * @type {RealmUser}
     * @memberof UserResponse
     */
    user?: RealmUser | null;
    /**
     * 
     * @type {string}
     * @memberof UserResponse
     */
    uuid?: string | null;
    /**
     * 
     * @type {boolean}
     * @memberof UserResponse
     */
    videoOn?: boolean | null;
    /**
     * 
     * @type {number}
     * @memberof UserResponse
     */
    vipUntilSeconds?: number | null;
    /**
     * 
     * @type {boolean}
     * @memberof UserResponse
     */
    worldIdConnected?: boolean | null;
}

/**
 * Check if a given object implements the UserResponse interface.
 */
export function instanceOfUserResponse(value: object): value is UserResponse {
    return true;
}

export function UserResponseFromJSON(json: any): UserResponse {
    return UserResponseFromJSONTyped(json, false);
}

export function UserResponseFromJSONTyped(json: any, ignoreDiscriminator: boolean): UserResponse {
    if (json == null) {
        return json;
    }
    return {
        
        'appleConnected': json['apple_connected'] == null ? undefined : json['apple_connected'],
        'birthDate': json['birth_date'] == null ? undefined : json['birth_date'],
        'blockingLimit': json['blocking_limit'] == null ? undefined : json['blocking_limit'],
        'emailConfirmed': json['email_confirmed'] == null ? undefined : json['email_confirmed'],
        'facebookConnected': json['facebook_connected'] == null ? undefined : json['facebook_connected'],
        'giftingAbility': json['gifting_ability'] == null ? undefined : RealmGiftingAbilityFromJSON(json['gifting_ability']),
        'googleConnected': json['google_connected'] == null ? undefined : json['google_connected'],
        'groupPhoneOn': json['group_phone_on'] == null ? undefined : json['group_phone_on'],
        'groupVideoOn': json['group_video_on'] == null ? undefined : json['group_video_on'],
        'interestsCount': json['interests_count'] == null ? undefined : json['interests_count'],
        'lineConnected': json['line_connected'] == null ? undefined : json['line_connected'],
        'maskedEmail': json['masked_email'] == null ? undefined : json['masked_email'],
        'phoneOn': json['phone_on'] == null ? undefined : json['phone_on'],
        'pushNotification': json['push_notification'] == null ? undefined : json['push_notification'],
        'twitterId': json['twitter_id'] == null ? undefined : json['twitter_id'],
        'user': json['user'] == null ? undefined : RealmUserFromJSON(json['user']),
        'uuid': json['uuid'] == null ? undefined : json['uuid'],
        'videoOn': json['video_on'] == null ? undefined : json['video_on'],
        'vipUntilSeconds': json['vip_until_seconds'] == null ? undefined : json['vip_until_seconds'],
        'worldIdConnected': json['world_id_connected'] == null ? undefined : json['world_id_connected'],
    };
}

export function UserResponseToJSON(json: any): UserResponse {
    return UserResponseToJSONTyped(json, false);
}

export function UserResponseToJSONTyped(value?: UserResponse | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'apple_connected': value['appleConnected'],
        'birth_date': value['birthDate'],
        'blocking_limit': value['blockingLimit'],
        'email_confirmed': value['emailConfirmed'],
        'facebook_connected': value['facebookConnected'],
        'gifting_ability': RealmGiftingAbilityToJSON(value['giftingAbility']),
        'google_connected': value['googleConnected'],
        'group_phone_on': value['groupPhoneOn'],
        'group_video_on': value['groupVideoOn'],
        'interests_count': value['interestsCount'],
        'line_connected': value['lineConnected'],
        'masked_email': value['maskedEmail'],
        'phone_on': value['phoneOn'],
        'push_notification': value['pushNotification'],
        'twitter_id': value['twitterId'],
        'user': RealmUserToJSON(value['user']),
        'uuid': value['uuid'],
        'video_on': value['videoOn'],
        'vip_until_seconds': value['vipUntilSeconds'],
        'world_id_connected': value['worldIdConnected'],
    };
}

