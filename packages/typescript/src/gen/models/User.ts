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
import type { ReviewRestriction } from './ReviewRestriction.js';
import {
    ReviewRestrictionFromJSON,
    ReviewRestrictionFromJSONTyped,
    ReviewRestrictionToJSON,
    ReviewRestrictionToJSONTyped,
} from './ReviewRestriction.js';
import type { OnlineStatus } from './OnlineStatus.js';
import {
    OnlineStatusFromJSON,
    OnlineStatusFromJSONTyped,
    OnlineStatusToJSON,
    OnlineStatusToJSONTyped,
} from './OnlineStatus.js';
import type { Title } from './Title.js';
import {
    TitleFromJSON,
    TitleFromJSONTyped,
    TitleToJSON,
    TitleToJSONTyped,
} from './Title.js';
import type { Gender } from './Gender.js';
import {
    GenderFromJSON,
    GenderFromJSONTyped,
    GenderToJSON,
    GenderToJSONTyped,
} from './Gender.js';
import type { GiftingAbility } from './GiftingAbility.js';
import {
    GiftingAbilityFromJSON,
    GiftingAbilityFromJSONTyped,
    GiftingAbilityToJSON,
    GiftingAbilityToJSONTyped,
} from './GiftingAbility.js';

/**
 * 
 * @export
 * @interface User
 */
export interface User {
    /**
     * 
     * @type {string}
     * @memberof User
     */
    biography?: string | null;
    /**
     * 
     * @type {string}
     * @memberof User
     */
    birthDate?: string | null;
    /**
     * 
     * @type {number}
     * @memberof User
     */
    blockingLimit?: number | null;
    /**
     * 
     * @type {object}
     * @memberof User
     */
    connectedBy?: object | null;
    /**
     * 
     * @type {Array<string>}
     * @memberof User
     */
    contactPhones?: Array<string> | null;
    /**
     * 
     * @type {object}
     * @memberof User
     */
    country?: object | null;
    /**
     * 
     * @type {string}
     * @memberof User
     */
    coverImage?: string | null;
    /**
     * 
     * @type {string}
     * @memberof User
     */
    coverImageThumbnail?: string | null;
    /**
     * 
     * @type {number}
     * @memberof User
     */
    createdAtMillis?: number | null;
    /**
     * 
     * @type {number}
     * @memberof User
     */
    followersCount?: number | null;
    /**
     * 
     * @type {number}
     * @memberof User
     */
    followingsCount?: number | null;
    /**
     * 
     * @type {Gender}
     * @memberof User
     */
    gender?: Gender | null;
    /**
     * 
     * @type {number}
     * @memberof User
     */
    giftReceivedCounter?: number | null;
    /**
     * 
     * @type {GiftingAbility}
     * @memberof User
     */
    giftingAbility?: GiftingAbility | null;
    /**
     * 
     * @type {number}
     * @memberof User
     */
    groupsUsersCount?: number | null;
    /**
     * 
     * @type {number}
     * @memberof User
     */
    id?: number | null;
    /**
     * 
     * @type {number}
     * @memberof User
     */
    interestsCount?: number | null;
    /**
     * 
     * @type {boolean}
     * @memberof User
     */
    isAgeVerified?: boolean | null;
    /**
     * 
     * @type {boolean}
     * @memberof User
     */
    isAppleConnected?: boolean | null;
    /**
     * 
     * @type {boolean}
     * @memberof User
     */
    isChatRequestOn?: boolean | null;
    /**
     * 
     * @type {boolean}
     * @memberof User
     */
    isDangerous?: boolean | null;
    /**
     * 
     * @type {boolean}
     * @memberof User
     */
    isEmailConfirmed?: boolean | null;
    /**
     * 
     * @type {boolean}
     * @memberof User
     */
    isFacebookConnected?: boolean | null;
    /**
     * 
     * @type {boolean}
     * @memberof User
     */
    isFollowPending?: boolean | null;
    /**
     * 
     * @type {boolean}
     * @memberof User
     */
    isFollowedBy?: boolean | null;
    /**
     * 
     * @type {boolean}
     * @memberof User
     */
    isFollowedByPending?: boolean | null;
    /**
     * 
     * @type {boolean}
     * @memberof User
     */
    isFollowing?: boolean | null;
    /**
     * 
     * @type {boolean}
     * @memberof User
     */
    isFromDifferentGenerationAndTrusted?: boolean | null;
    /**
     * 
     * @type {boolean}
     * @memberof User
     */
    isGoogleConnected?: boolean | null;
    /**
     * 
     * @type {boolean}
     * @memberof User
     */
    isGroupPhoneOn?: boolean | null;
    /**
     * 
     * @type {boolean}
     * @memberof User
     */
    isGroupVideoOn?: boolean | null;
    /**
     * 
     * @type {boolean}
     * @memberof User
     */
    isHidden?: boolean | null;
    /**
     * 
     * @type {boolean}
     * @memberof User
     */
    isHideVip?: boolean | null;
    /**
     * 
     * @type {boolean}
     * @memberof User
     */
    isInterestsSelected?: boolean | null;
    /**
     * 
     * @type {boolean}
     * @memberof User
     */
    isLineConnected?: boolean | null;
    /**
     * 
     * @type {boolean}
     * @memberof User
     */
    isMuted?: boolean | null;
    /**
     * 
     * @type {boolean}
     * @memberof User
     */
    isNew?: boolean | null;
    /**
     * 
     * @type {boolean}
     * @memberof User
     */
    isPhoneOn?: boolean | null;
    /**
     * 
     * @type {boolean}
     * @memberof User
     */
    isPrivate?: boolean | null;
    /**
     * 
     * @type {boolean}
     * @memberof User
     */
    isRecentlyPenalized?: boolean | null;
    /**
     * 
     * @type {boolean}
     * @memberof User
     */
    isRegistered?: boolean | null;
    /**
     * 
     * @type {boolean}
     * @memberof User
     */
    isTwoFaEnabled?: boolean | null;
    /**
     * 
     * @type {boolean}
     * @memberof User
     */
    isVideoOn?: boolean | null;
    /**
     * 
     * @type {boolean}
     * @memberof User
     */
    isVip?: boolean | null;
    /**
     * 
     * @type {boolean}
     * @memberof User
     */
    isWorldIdConnected?: boolean | null;
    /**
     * 
     * @type {number}
     * @memberof User
     */
    lastLoggedInAtMillis?: number | null;
    /**
     * 
     * @type {number}
     * @memberof User
     */
    loginStreakCount?: number | null;
    /**
     * 
     * @type {string}
     * @memberof User
     */
    maskedEmail?: string | null;
    /**
     * 
     * @type {number}
     * @memberof User
     */
    matchingId?: number | null;
    /**
     * 
     * @type {number}
     * @memberof User
     */
    mediaCount?: number | null;
    /**
     * 
     * @type {string}
     * @memberof User
     */
    mobileNumber?: string | null;
    /**
     * 
     * @type {string}
     * @memberof User
     */
    nickname?: string | null;
    /**
     * 
     * @type {OnlineStatus}
     * @memberof User
     */
    onlineStatus?: OnlineStatus | null;
    /**
     * 
     * @type {number}
     * @memberof User
     */
    postsCount?: number | null;
    /**
     * 
     * @type {string}
     * @memberof User
     */
    prefecture?: string | null;
    /**
     * 
     * @type {string}
     * @memberof User
     */
    profileIcon?: string | null;
    /**
     * 
     * @type {string}
     * @memberof User
     */
    profileIconFrame?: string | null;
    /**
     * 
     * @type {string}
     * @memberof User
     */
    profileIconFrameThumbnail?: string | null;
    /**
     * 
     * @type {string}
     * @memberof User
     */
    profileIconThumbnail?: string | null;
    /**
     * 
     * @type {ReviewRestriction}
     * @memberof User
     */
    reviewRestriction?: ReviewRestriction | null;
    /**
     * 
     * @type {number}
     * @memberof User
     */
    reviewsCount?: number | null;
    /**
     * 
     * @type {Title}
     * @memberof User
     */
    title?: Title | null;
    /**
     * 
     * @type {string}
     * @memberof User
     */
    twitterId?: string | null;
    /**
     * 
     * @type {string}
     * @memberof User
     */
    username?: string | null;
    /**
     * 
     * @type {string}
     * @memberof User
     */
    uuid?: string | null;
    /**
     * 
     * @type {number}
     * @memberof User
     */
    vipUntilMillis?: number | null;
}

/**
 * Check if a given object implements the User interface.
 */
export function instanceOfUser(value: object): value is User {
    return true;
}

export function UserFromJSON(json: any): User {
    return UserFromJSONTyped(json, false);
}

export function UserFromJSONTyped(json: any, ignoreDiscriminator: boolean): User {
    if (json == null) {
        return json;
    }
    return {
        
        'biography': json['biography'] == null ? undefined : json['biography'],
        'birthDate': json['birth_date'] == null ? undefined : json['birth_date'],
        'blockingLimit': json['blocking_limit'] == null ? undefined : json['blocking_limit'],
        'connectedBy': json['connected_by'] == null ? undefined : json['connected_by'],
        'contactPhones': json['contact_phones'] == null ? undefined : json['contact_phones'],
        'country': json['country'] == null ? undefined : json['country'],
        'coverImage': json['cover_image'] == null ? undefined : json['cover_image'],
        'coverImageThumbnail': json['cover_image_thumbnail'] == null ? undefined : json['cover_image_thumbnail'],
        'createdAtMillis': json['created_at_millis'] == null ? undefined : json['created_at_millis'],
        'followersCount': json['followers_count'] == null ? undefined : json['followers_count'],
        'followingsCount': json['followings_count'] == null ? undefined : json['followings_count'],
        'gender': json['gender'] == null ? undefined : GenderFromJSON(json['gender']),
        'giftReceivedCounter': json['gift_received_counter'] == null ? undefined : json['gift_received_counter'],
        'giftingAbility': json['gifting_ability'] == null ? undefined : GiftingAbilityFromJSON(json['gifting_ability']),
        'groupsUsersCount': json['groups_users_count'] == null ? undefined : json['groups_users_count'],
        'id': json['id'] == null ? undefined : json['id'],
        'interestsCount': json['interests_count'] == null ? undefined : json['interests_count'],
        'isAgeVerified': json['is_age_verified'] == null ? undefined : json['is_age_verified'],
        'isAppleConnected': json['is_apple_connected'] == null ? undefined : json['is_apple_connected'],
        'isChatRequestOn': json['is_chat_request_on'] == null ? undefined : json['is_chat_request_on'],
        'isDangerous': json['is_dangerous'] == null ? undefined : json['is_dangerous'],
        'isEmailConfirmed': json['is_email_confirmed'] == null ? undefined : json['is_email_confirmed'],
        'isFacebookConnected': json['is_facebook_connected'] == null ? undefined : json['is_facebook_connected'],
        'isFollowPending': json['is_follow_pending'] == null ? undefined : json['is_follow_pending'],
        'isFollowedBy': json['is_followed_by'] == null ? undefined : json['is_followed_by'],
        'isFollowedByPending': json['is_followed_by_pending'] == null ? undefined : json['is_followed_by_pending'],
        'isFollowing': json['is_following'] == null ? undefined : json['is_following'],
        'isFromDifferentGenerationAndTrusted': json['is_from_different_generation_and_trusted'] == null ? undefined : json['is_from_different_generation_and_trusted'],
        'isGoogleConnected': json['is_google_connected'] == null ? undefined : json['is_google_connected'],
        'isGroupPhoneOn': json['is_group_phone_on'] == null ? undefined : json['is_group_phone_on'],
        'isGroupVideoOn': json['is_group_video_on'] == null ? undefined : json['is_group_video_on'],
        'isHidden': json['is_hidden'] == null ? undefined : json['is_hidden'],
        'isHideVip': json['is_hide_vip'] == null ? undefined : json['is_hide_vip'],
        'isInterestsSelected': json['is_interests_selected'] == null ? undefined : json['is_interests_selected'],
        'isLineConnected': json['is_line_connected'] == null ? undefined : json['is_line_connected'],
        'isMuted': json['is_muted'] == null ? undefined : json['is_muted'],
        'isNew': json['is_new'] == null ? undefined : json['is_new'],
        'isPhoneOn': json['is_phone_on'] == null ? undefined : json['is_phone_on'],
        'isPrivate': json['is_private'] == null ? undefined : json['is_private'],
        'isRecentlyPenalized': json['is_recently_penalized'] == null ? undefined : json['is_recently_penalized'],
        'isRegistered': json['is_registered'] == null ? undefined : json['is_registered'],
        'isTwoFaEnabled': json['is_two_fa_enabled'] == null ? undefined : json['is_two_fa_enabled'],
        'isVideoOn': json['is_video_on'] == null ? undefined : json['is_video_on'],
        'isVip': json['is_vip'] == null ? undefined : json['is_vip'],
        'isWorldIdConnected': json['is_world_id_connected'] == null ? undefined : json['is_world_id_connected'],
        'lastLoggedInAtMillis': json['last_logged_in_at_millis'] == null ? undefined : json['last_logged_in_at_millis'],
        'loginStreakCount': json['login_streak_count'] == null ? undefined : json['login_streak_count'],
        'maskedEmail': json['masked_email'] == null ? undefined : json['masked_email'],
        'matchingId': json['matching_id'] == null ? undefined : json['matching_id'],
        'mediaCount': json['media_count'] == null ? undefined : json['media_count'],
        'mobileNumber': json['mobile_number'] == null ? undefined : json['mobile_number'],
        'nickname': json['nickname'] == null ? undefined : json['nickname'],
        'onlineStatus': json['online_status'] == null ? undefined : OnlineStatusFromJSON(json['online_status']),
        'postsCount': json['posts_count'] == null ? undefined : json['posts_count'],
        'prefecture': json['prefecture'] == null ? undefined : json['prefecture'],
        'profileIcon': json['profile_icon'] == null ? undefined : json['profile_icon'],
        'profileIconFrame': json['profile_icon_frame'] == null ? undefined : json['profile_icon_frame'],
        'profileIconFrameThumbnail': json['profile_icon_frame_thumbnail'] == null ? undefined : json['profile_icon_frame_thumbnail'],
        'profileIconThumbnail': json['profile_icon_thumbnail'] == null ? undefined : json['profile_icon_thumbnail'],
        'reviewRestriction': json['review_restriction'] == null ? undefined : ReviewRestrictionFromJSON(json['review_restriction']),
        'reviewsCount': json['reviews_count'] == null ? undefined : json['reviews_count'],
        'title': json['title'] == null ? undefined : TitleFromJSON(json['title']),
        'twitterId': json['twitter_id'] == null ? undefined : json['twitter_id'],
        'username': json['username'] == null ? undefined : json['username'],
        'uuid': json['uuid'] == null ? undefined : json['uuid'],
        'vipUntilMillis': json['vip_until_millis'] == null ? undefined : json['vip_until_millis'],
    };
}

export function UserToJSON(json: any): User {
    return UserToJSONTyped(json, false);
}

export function UserToJSONTyped(value?: User | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'biography': value['biography'],
        'birth_date': value['birthDate'],
        'blocking_limit': value['blockingLimit'],
        'connected_by': value['connectedBy'],
        'contact_phones': value['contactPhones'],
        'country': value['country'],
        'cover_image': value['coverImage'],
        'cover_image_thumbnail': value['coverImageThumbnail'],
        'created_at_millis': value['createdAtMillis'],
        'followers_count': value['followersCount'],
        'followings_count': value['followingsCount'],
        'gender': GenderToJSON(value['gender']),
        'gift_received_counter': value['giftReceivedCounter'],
        'gifting_ability': GiftingAbilityToJSON(value['giftingAbility']),
        'groups_users_count': value['groupsUsersCount'],
        'id': value['id'],
        'interests_count': value['interestsCount'],
        'is_age_verified': value['isAgeVerified'],
        'is_apple_connected': value['isAppleConnected'],
        'is_chat_request_on': value['isChatRequestOn'],
        'is_dangerous': value['isDangerous'],
        'is_email_confirmed': value['isEmailConfirmed'],
        'is_facebook_connected': value['isFacebookConnected'],
        'is_follow_pending': value['isFollowPending'],
        'is_followed_by': value['isFollowedBy'],
        'is_followed_by_pending': value['isFollowedByPending'],
        'is_following': value['isFollowing'],
        'is_from_different_generation_and_trusted': value['isFromDifferentGenerationAndTrusted'],
        'is_google_connected': value['isGoogleConnected'],
        'is_group_phone_on': value['isGroupPhoneOn'],
        'is_group_video_on': value['isGroupVideoOn'],
        'is_hidden': value['isHidden'],
        'is_hide_vip': value['isHideVip'],
        'is_interests_selected': value['isInterestsSelected'],
        'is_line_connected': value['isLineConnected'],
        'is_muted': value['isMuted'],
        'is_new': value['isNew'],
        'is_phone_on': value['isPhoneOn'],
        'is_private': value['isPrivate'],
        'is_recently_penalized': value['isRecentlyPenalized'],
        'is_registered': value['isRegistered'],
        'is_two_fa_enabled': value['isTwoFaEnabled'],
        'is_video_on': value['isVideoOn'],
        'is_vip': value['isVip'],
        'is_world_id_connected': value['isWorldIdConnected'],
        'last_logged_in_at_millis': value['lastLoggedInAtMillis'],
        'login_streak_count': value['loginStreakCount'],
        'masked_email': value['maskedEmail'],
        'matching_id': value['matchingId'],
        'media_count': value['mediaCount'],
        'mobile_number': value['mobileNumber'],
        'nickname': value['nickname'],
        'online_status': OnlineStatusToJSON(value['onlineStatus']),
        'posts_count': value['postsCount'],
        'prefecture': value['prefecture'],
        'profile_icon': value['profileIcon'],
        'profile_icon_frame': value['profileIconFrame'],
        'profile_icon_frame_thumbnail': value['profileIconFrameThumbnail'],
        'profile_icon_thumbnail': value['profileIconThumbnail'],
        'review_restriction': ReviewRestrictionToJSON(value['reviewRestriction']),
        'reviews_count': value['reviewsCount'],
        'title': TitleToJSON(value['title']),
        'twitter_id': value['twitterId'],
        'username': value['username'],
        'uuid': value['uuid'],
        'vip_until_millis': value['vipUntilMillis'],
    };
}

