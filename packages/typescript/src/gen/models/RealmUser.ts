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
import type { GroupUser } from './GroupUser.js';
import {
    GroupUserFromJSON,
    GroupUserFromJSONTyped,
    GroupUserToJSON,
    GroupUserToJSONTyped,
} from './GroupUser.js';
import type { OnlineStatusEnum } from './OnlineStatusEnum.js';
import {
    OnlineStatusEnumFromJSON,
    OnlineStatusEnumFromJSONTyped,
    OnlineStatusEnumToJSON,
    OnlineStatusEnumToJSONTyped,
} from './OnlineStatusEnum.js';
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
 * @interface RealmUser
 */
export interface RealmUser {
    /**
     * 
     * @type {boolean}
     * @memberof RealmUser
     */
    ageVerified?: boolean | null;
    /**
     * 
     * @type {boolean}
     * @memberof RealmUser
     */
    appleConnected?: boolean | null;
    /**
     * 
     * @type {string}
     * @memberof RealmUser
     */
    biography?: string | null;
    /**
     * 
     * @type {string}
     * @memberof RealmUser
     */
    birthDate?: string | null;
    /**
     * 
     * @type {number}
     * @memberof RealmUser
     */
    blockingLimit?: number | null;
    /**
     * 
     * @type {boolean}
     * @memberof RealmUser
     */
    chatRequest?: boolean | null;
    /**
     * 
     * @type {Array<string>}
     * @memberof RealmUser
     */
    connectedBy?: Array<string> | null;
    /**
     * 
     * @type {Array<string>}
     * @memberof RealmUser
     */
    contactPhones?: Array<string> | null;
    /**
     * 
     * @type {string}
     * @memberof RealmUser
     */
    countryCode?: string | null;
    /**
     * 
     * @type {string}
     * @memberof RealmUser
     */
    coverImage?: string | null;
    /**
     * 
     * @type {string}
     * @memberof RealmUser
     */
    coverImageThumbnail?: string | null;
    /**
     * 
     * @type {number}
     * @memberof RealmUser
     */
    createdAtSeconds?: number | null;
    /**
     * 
     * @type {boolean}
     * @memberof RealmUser
     */
    dangerousUser?: boolean | null;
    /**
     * 
     * @type {boolean}
     * @memberof RealmUser
     */
    emailConfirmed?: boolean | null;
    /**
     * 
     * @type {boolean}
     * @memberof RealmUser
     */
    facebookConnected?: boolean | null;
    /**
     * 
     * @type {boolean}
     * @memberof RealmUser
     */
    followPending?: boolean | null;
    /**
     * 
     * @type {boolean}
     * @memberof RealmUser
     */
    followedBy?: boolean | null;
    /**
     * 
     * @type {boolean}
     * @memberof RealmUser
     */
    followedByPending?: boolean | null;
    /**
     * 
     * @type {number}
     * @memberof RealmUser
     */
    followersCount?: number | null;
    /**
     * 
     * @type {boolean}
     * @memberof RealmUser
     */
    following?: boolean | null;
    /**
     * 
     * @type {number}
     * @memberof RealmUser
     */
    followingsCount?: number | null;
    /**
     * 
     * @type {string}
     * @memberof RealmUser
     */
    frame?: string | null;
    /**
     * 
     * @type {string}
     * @memberof RealmUser
     */
    frameThumbnail?: string | null;
    /**
     * 
     * @type {boolean}
     * @memberof RealmUser
     */
    fromDifferentGenerationAndTrusted?: boolean | null;
    /**
     * 
     * @type {number}
     * @memberof RealmUser
     */
    gender?: number | null;
    /**
     * 
     * @type {number}
     * @memberof RealmUser
     */
    generation?: number | null;
    /**
     * 
     * @type {RealmGiftingAbility}
     * @memberof RealmUser
     */
    giftingAbility?: RealmGiftingAbility | null;
    /**
     * 
     * @type {boolean}
     * @memberof RealmUser
     */
    googleConnected?: boolean | null;
    /**
     * 
     * @type {boolean}
     * @memberof RealmUser
     */
    groupPhoneOn?: boolean | null;
    /**
     * 
     * @type {GroupUser}
     * @memberof RealmUser
     */
    groupUser?: GroupUser | null;
    /**
     * 
     * @type {boolean}
     * @memberof RealmUser
     */
    groupVideoOn?: boolean | null;
    /**
     * 
     * @type {number}
     * @memberof RealmUser
     */
    groupsUsersCount?: number | null;
    /**
     * 
     * @type {boolean}
     * @memberof RealmUser
     */
    hidden?: boolean | null;
    /**
     * 
     * @type {boolean}
     * @memberof RealmUser
     */
    hideVip?: boolean | null;
    /**
     * 
     * @type {number}
     * @memberof RealmUser
     */
    id?: number | null;
    /**
     * 
     * @type {number}
     * @memberof RealmUser
     */
    interestsCount?: number | null;
    /**
     * 
     * @type {boolean}
     * @memberof RealmUser
     */
    interestsSelected?: boolean | null;
    /**
     * 
     * @type {boolean}
     * @memberof RealmUser
     */
    isPrivate?: boolean | null;
    /**
     * 
     * @type {number}
     * @memberof RealmUser
     */
    lastLoggedInAtSeconds?: number | null;
    /**
     * 
     * @type {boolean}
     * @memberof RealmUser
     */
    lineConnected?: boolean | null;
    /**
     * 
     * @type {number}
     * @memberof RealmUser
     */
    loginStreakCount?: number | null;
    /**
     * 
     * @type {string}
     * @memberof RealmUser
     */
    maskedEmail?: string | null;
    /**
     * 
     * @type {number}
     * @memberof RealmUser
     */
    mediaCount?: number | null;
    /**
     * 
     * @type {string}
     * @memberof RealmUser
     */
    mobileNumber?: string | null;
    /**
     * 
     * @type {boolean}
     * @memberof RealmUser
     */
    newUser?: boolean | null;
    /**
     * 
     * @type {string}
     * @memberof RealmUser
     */
    nickname?: string | null;
    /**
     * 
     * @type {OnlineStatusEnum}
     * @memberof RealmUser
     */
    onlineStatus?: OnlineStatusEnum | null;
    /**
     * 
     * @type {boolean}
     * @memberof RealmUser
     */
    phoneOn?: boolean | null;
    /**
     * 
     * @type {number}
     * @memberof RealmUser
     */
    postsCount?: number | null;
    /**
     * 
     * @type {string}
     * @memberof RealmUser
     */
    prefecture?: string | null;
    /**
     * 
     * @type {string}
     * @memberof RealmUser
     */
    profileIcon?: string | null;
    /**
     * 
     * @type {string}
     * @memberof RealmUser
     */
    profileIconThumbnail?: string | null;
    /**
     * 
     * @type {boolean}
     * @memberof RealmUser
     */
    pushNotification?: boolean | null;
    /**
     * 
     * @type {boolean}
     * @memberof RealmUser
     */
    recentlyKenta?: boolean | null;
    /**
     * 
     * @type {string}
     * @memberof RealmUser
     */
    restrictedReviewBy?: string | null;
    /**
     * 
     * @type {number}
     * @memberof RealmUser
     */
    reviewsCount?: number | null;
    /**
     * 
     * @type {string}
     * @memberof RealmUser
     */
    title?: string | null;
    /**
     * 
     * @type {number}
     * @memberof RealmUser
     */
    totalGiftsReceivedCount?: number | null;
    /**
     * 
     * @type {string}
     * @memberof RealmUser
     */
    twitterId?: string | null;
    /**
     * 
     * @type {boolean}
     * @memberof RealmUser
     */
    twoFaEnabled?: boolean | null;
    /**
     * 
     * @type {number}
     * @memberof RealmUser
     */
    updatedTimeMillis?: number | null;
    /**
     * 
     * @type {string}
     * @memberof RealmUser
     */
    username?: string | null;
    /**
     * 
     * @type {string}
     * @memberof RealmUser
     */
    uuid?: string | null;
    /**
     * 
     * @type {boolean}
     * @memberof RealmUser
     */
    videoOn?: boolean | null;
    /**
     * 
     * @type {boolean}
     * @memberof RealmUser
     */
    vip?: boolean | null;
    /**
     * 
     * @type {number}
     * @memberof RealmUser
     */
    vipUntilSeconds?: number | null;
    /**
     * 
     * @type {boolean}
     * @memberof RealmUser
     */
    worldIdConnected?: boolean | null;
}



/**
 * Check if a given object implements the RealmUser interface.
 */
export function instanceOfRealmUser(value: object): value is RealmUser {
    return true;
}

export function RealmUserFromJSON(json: any): RealmUser {
    return RealmUserFromJSONTyped(json, false);
}

export function RealmUserFromJSONTyped(json: any, ignoreDiscriminator: boolean): RealmUser {
    if (json == null) {
        return json;
    }
    return {
        
        'ageVerified': json['age_verified'] == null ? undefined : json['age_verified'],
        'appleConnected': json['apple_connected'] == null ? undefined : json['apple_connected'],
        'biography': json['biography'] == null ? undefined : json['biography'],
        'birthDate': json['birth_date'] == null ? undefined : json['birth_date'],
        'blockingLimit': json['blocking_limit'] == null ? undefined : json['blocking_limit'],
        'chatRequest': json['chat_request'] == null ? undefined : json['chat_request'],
        'connectedBy': json['connected_by'] == null ? undefined : json['connected_by'],
        'contactPhones': json['contact_phones'] == null ? undefined : json['contact_phones'],
        'countryCode': json['country_code'] == null ? undefined : json['country_code'],
        'coverImage': json['cover_image'] == null ? undefined : json['cover_image'],
        'coverImageThumbnail': json['cover_image_thumbnail'] == null ? undefined : json['cover_image_thumbnail'],
        'createdAtSeconds': json['created_at_seconds'] == null ? undefined : json['created_at_seconds'],
        'dangerousUser': json['dangerous_user'] == null ? undefined : json['dangerous_user'],
        'emailConfirmed': json['email_confirmed'] == null ? undefined : json['email_confirmed'],
        'facebookConnected': json['facebook_connected'] == null ? undefined : json['facebook_connected'],
        'followPending': json['follow_pending'] == null ? undefined : json['follow_pending'],
        'followedBy': json['followed_by'] == null ? undefined : json['followed_by'],
        'followedByPending': json['followed_by_pending'] == null ? undefined : json['followed_by_pending'],
        'followersCount': json['followers_count'] == null ? undefined : json['followers_count'],
        'following': json['following'] == null ? undefined : json['following'],
        'followingsCount': json['followings_count'] == null ? undefined : json['followings_count'],
        'frame': json['frame'] == null ? undefined : json['frame'],
        'frameThumbnail': json['frame_thumbnail'] == null ? undefined : json['frame_thumbnail'],
        'fromDifferentGenerationAndTrusted': json['from_different_generation_and_trusted'] == null ? undefined : json['from_different_generation_and_trusted'],
        'gender': json['gender'] == null ? undefined : json['gender'],
        'generation': json['generation'] == null ? undefined : json['generation'],
        'giftingAbility': json['gifting_ability'] == null ? undefined : RealmGiftingAbilityFromJSON(json['gifting_ability']),
        'googleConnected': json['google_connected'] == null ? undefined : json['google_connected'],
        'groupPhoneOn': json['group_phone_on'] == null ? undefined : json['group_phone_on'],
        'groupUser': json['group_user'] == null ? undefined : GroupUserFromJSON(json['group_user']),
        'groupVideoOn': json['group_video_on'] == null ? undefined : json['group_video_on'],
        'groupsUsersCount': json['groups_users_count'] == null ? undefined : json['groups_users_count'],
        'hidden': json['hidden'] == null ? undefined : json['hidden'],
        'hideVip': json['hide_vip'] == null ? undefined : json['hide_vip'],
        'id': json['id'] == null ? undefined : json['id'],
        'interestsCount': json['interests_count'] == null ? undefined : json['interests_count'],
        'interestsSelected': json['interests_selected'] == null ? undefined : json['interests_selected'],
        'isPrivate': json['is_private'] == null ? undefined : json['is_private'],
        'lastLoggedInAtSeconds': json['last_logged_in_at_seconds'] == null ? undefined : json['last_logged_in_at_seconds'],
        'lineConnected': json['line_connected'] == null ? undefined : json['line_connected'],
        'loginStreakCount': json['login_streak_count'] == null ? undefined : json['login_streak_count'],
        'maskedEmail': json['masked_email'] == null ? undefined : json['masked_email'],
        'mediaCount': json['media_count'] == null ? undefined : json['media_count'],
        'mobileNumber': json['mobile_number'] == null ? undefined : json['mobile_number'],
        'newUser': json['new_user'] == null ? undefined : json['new_user'],
        'nickname': json['nickname'] == null ? undefined : json['nickname'],
        'onlineStatus': json['online_status'] == null ? undefined : OnlineStatusEnumFromJSON(json['online_status']),
        'phoneOn': json['phone_on'] == null ? undefined : json['phone_on'],
        'postsCount': json['posts_count'] == null ? undefined : json['posts_count'],
        'prefecture': json['prefecture'] == null ? undefined : json['prefecture'],
        'profileIcon': json['profile_icon'] == null ? undefined : json['profile_icon'],
        'profileIconThumbnail': json['profile_icon_thumbnail'] == null ? undefined : json['profile_icon_thumbnail'],
        'pushNotification': json['push_notification'] == null ? undefined : json['push_notification'],
        'recentlyKenta': json['recently_kenta'] == null ? undefined : json['recently_kenta'],
        'restrictedReviewBy': json['restricted_review_by'] == null ? undefined : json['restricted_review_by'],
        'reviewsCount': json['reviews_count'] == null ? undefined : json['reviews_count'],
        'title': json['title'] == null ? undefined : json['title'],
        'totalGiftsReceivedCount': json['total_gifts_received_count'] == null ? undefined : json['total_gifts_received_count'],
        'twitterId': json['twitter_id'] == null ? undefined : json['twitter_id'],
        'twoFaEnabled': json['two_fa_enabled'] == null ? undefined : json['two_fa_enabled'],
        'updatedTimeMillis': json['updated_time_millis'] == null ? undefined : json['updated_time_millis'],
        'username': json['username'] == null ? undefined : json['username'],
        'uuid': json['uuid'] == null ? undefined : json['uuid'],
        'videoOn': json['video_on'] == null ? undefined : json['video_on'],
        'vip': json['vip'] == null ? undefined : json['vip'],
        'vipUntilSeconds': json['vip_until_seconds'] == null ? undefined : json['vip_until_seconds'],
        'worldIdConnected': json['world_id_connected'] == null ? undefined : json['world_id_connected'],
    };
}

export function RealmUserToJSON(json: any): RealmUser {
    return RealmUserToJSONTyped(json, false);
}

export function RealmUserToJSONTyped(value?: RealmUser | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'age_verified': value['ageVerified'],
        'apple_connected': value['appleConnected'],
        'biography': value['biography'],
        'birth_date': value['birthDate'],
        'blocking_limit': value['blockingLimit'],
        'chat_request': value['chatRequest'],
        'connected_by': value['connectedBy'],
        'contact_phones': value['contactPhones'],
        'country_code': value['countryCode'],
        'cover_image': value['coverImage'],
        'cover_image_thumbnail': value['coverImageThumbnail'],
        'created_at_seconds': value['createdAtSeconds'],
        'dangerous_user': value['dangerousUser'],
        'email_confirmed': value['emailConfirmed'],
        'facebook_connected': value['facebookConnected'],
        'follow_pending': value['followPending'],
        'followed_by': value['followedBy'],
        'followed_by_pending': value['followedByPending'],
        'followers_count': value['followersCount'],
        'following': value['following'],
        'followings_count': value['followingsCount'],
        'frame': value['frame'],
        'frame_thumbnail': value['frameThumbnail'],
        'from_different_generation_and_trusted': value['fromDifferentGenerationAndTrusted'],
        'gender': value['gender'],
        'generation': value['generation'],
        'gifting_ability': RealmGiftingAbilityToJSON(value['giftingAbility']),
        'google_connected': value['googleConnected'],
        'group_phone_on': value['groupPhoneOn'],
        'group_user': GroupUserToJSON(value['groupUser']),
        'group_video_on': value['groupVideoOn'],
        'groups_users_count': value['groupsUsersCount'],
        'hidden': value['hidden'],
        'hide_vip': value['hideVip'],
        'id': value['id'],
        'interests_count': value['interestsCount'],
        'interests_selected': value['interestsSelected'],
        'is_private': value['isPrivate'],
        'last_logged_in_at_seconds': value['lastLoggedInAtSeconds'],
        'line_connected': value['lineConnected'],
        'login_streak_count': value['loginStreakCount'],
        'masked_email': value['maskedEmail'],
        'media_count': value['mediaCount'],
        'mobile_number': value['mobileNumber'],
        'new_user': value['newUser'],
        'nickname': value['nickname'],
        'online_status': OnlineStatusEnumToJSON(value['onlineStatus']),
        'phone_on': value['phoneOn'],
        'posts_count': value['postsCount'],
        'prefecture': value['prefecture'],
        'profile_icon': value['profileIcon'],
        'profile_icon_thumbnail': value['profileIconThumbnail'],
        'push_notification': value['pushNotification'],
        'recently_kenta': value['recentlyKenta'],
        'restricted_review_by': value['restrictedReviewBy'],
        'reviews_count': value['reviewsCount'],
        'title': value['title'],
        'total_gifts_received_count': value['totalGiftsReceivedCount'],
        'twitter_id': value['twitterId'],
        'two_fa_enabled': value['twoFaEnabled'],
        'updated_time_millis': value['updatedTimeMillis'],
        'username': value['username'],
        'uuid': value['uuid'],
        'video_on': value['videoOn'],
        'vip': value['vip'],
        'vip_until_seconds': value['vipUntilSeconds'],
        'world_id_connected': value['worldIdConnected'],
    };
}

