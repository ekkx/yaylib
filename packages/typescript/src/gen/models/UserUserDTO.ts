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
 * @interface UserUserDTO
 */
export interface UserUserDTO {
    /**
     * 
     * @type {boolean}
     * @memberof UserUserDTO
     */
    ageVerified?: boolean | null;
    /**
     * 
     * @type {string}
     * @memberof UserUserDTO
     */
    biography?: string | null;
    /**
     * 
     * @type {string}
     * @memberof UserUserDTO
     */
    coverImage?: string | null;
    /**
     * 
     * @type {string}
     * @memberof UserUserDTO
     */
    coverImageThumbnail?: string | null;
    /**
     * 
     * @type {number}
     * @memberof UserUserDTO
     */
    createdAt?: number | null;
    /**
     * 
     * @type {boolean}
     * @memberof UserUserDTO
     */
    dangerousUser?: boolean | null;
    /**
     * 
     * @type {boolean}
     * @memberof UserUserDTO
     */
    followPending?: boolean | null;
    /**
     * 
     * @type {boolean}
     * @memberof UserUserDTO
     */
    followedBy?: boolean | null;
    /**
     * 
     * @type {number}
     * @memberof UserUserDTO
     */
    followersCount?: number | null;
    /**
     * 
     * @type {boolean}
     * @memberof UserUserDTO
     */
    following?: boolean | null;
    /**
     * 
     * @type {number}
     * @memberof UserUserDTO
     */
    followingsCount?: number | null;
    /**
     * 
     * @type {number}
     * @memberof UserUserDTO
     */
    gender?: number | null;
    /**
     * 
     * @type {number}
     * @memberof UserUserDTO
     */
    groupsUsersCount?: number | null;
    /**
     * 
     * @type {boolean}
     * @memberof UserUserDTO
     */
    hidden?: boolean | null;
    /**
     * 
     * @type {boolean}
     * @memberof UserUserDTO
     */
    hideVip?: boolean | null;
    /**
     * 
     * @type {number}
     * @memberof UserUserDTO
     */
    id?: number | null;
    /**
     * 
     * @type {boolean}
     * @memberof UserUserDTO
     */
    isPrivate?: boolean | null;
    /**
     * 
     * @type {boolean}
     * @memberof UserUserDTO
     */
    newUser?: boolean | null;
    /**
     * 
     * @type {string}
     * @memberof UserUserDTO
     */
    nickname?: string | null;
    /**
     * 
     * @type {number}
     * @memberof UserUserDTO
     */
    postsCount?: number | null;
    /**
     * 
     * @type {string}
     * @memberof UserUserDTO
     */
    prefecture?: string | null;
    /**
     * 
     * @type {string}
     * @memberof UserUserDTO
     */
    profileIcon?: string | null;
    /**
     * 
     * @type {string}
     * @memberof UserUserDTO
     */
    profileIconThumbnail?: string | null;
    /**
     * 
     * @type {boolean}
     * @memberof UserUserDTO
     */
    recentlyKenta?: boolean | null;
    /**
     * 
     * @type {number}
     * @memberof UserUserDTO
     */
    reviewsCount?: number | null;
    /**
     * 
     * @type {string}
     * @memberof UserUserDTO
     */
    title?: string | null;
    /**
     * 
     * @type {boolean}
     * @memberof UserUserDTO
     */
    vip?: boolean | null;
}

/**
 * Check if a given object implements the UserUserDTO interface.
 */
export function instanceOfUserUserDTO(value: object): value is UserUserDTO {
    return true;
}

export function UserUserDTOFromJSON(json: any): UserUserDTO {
    return UserUserDTOFromJSONTyped(json, false);
}

export function UserUserDTOFromJSONTyped(json: any, ignoreDiscriminator: boolean): UserUserDTO {
    if (json == null) {
        return json;
    }
    return {
        
        'ageVerified': json['age_verified'] == null ? undefined : json['age_verified'],
        'biography': json['biography'] == null ? undefined : json['biography'],
        'coverImage': json['cover_image'] == null ? undefined : json['cover_image'],
        'coverImageThumbnail': json['cover_image_thumbnail'] == null ? undefined : json['cover_image_thumbnail'],
        'createdAt': json['created_at'] == null ? undefined : json['created_at'],
        'dangerousUser': json['dangerous_user'] == null ? undefined : json['dangerous_user'],
        'followPending': json['follow_pending'] == null ? undefined : json['follow_pending'],
        'followedBy': json['followed_by'] == null ? undefined : json['followed_by'],
        'followersCount': json['followers_count'] == null ? undefined : json['followers_count'],
        'following': json['following'] == null ? undefined : json['following'],
        'followingsCount': json['followings_count'] == null ? undefined : json['followings_count'],
        'gender': json['gender'] == null ? undefined : json['gender'],
        'groupsUsersCount': json['groups_users_count'] == null ? undefined : json['groups_users_count'],
        'hidden': json['hidden'] == null ? undefined : json['hidden'],
        'hideVip': json['hide_vip'] == null ? undefined : json['hide_vip'],
        'id': json['id'] == null ? undefined : json['id'],
        'isPrivate': json['is_private'] == null ? undefined : json['is_private'],
        'newUser': json['new_user'] == null ? undefined : json['new_user'],
        'nickname': json['nickname'] == null ? undefined : json['nickname'],
        'postsCount': json['posts_count'] == null ? undefined : json['posts_count'],
        'prefecture': json['prefecture'] == null ? undefined : json['prefecture'],
        'profileIcon': json['profile_icon'] == null ? undefined : json['profile_icon'],
        'profileIconThumbnail': json['profile_icon_thumbnail'] == null ? undefined : json['profile_icon_thumbnail'],
        'recentlyKenta': json['recently_kenta'] == null ? undefined : json['recently_kenta'],
        'reviewsCount': json['reviews_count'] == null ? undefined : json['reviews_count'],
        'title': json['title'] == null ? undefined : json['title'],
        'vip': json['vip'] == null ? undefined : json['vip'],
    };
}

export function UserUserDTOToJSON(json: any): UserUserDTO {
    return UserUserDTOToJSONTyped(json, false);
}

export function UserUserDTOToJSONTyped(value?: UserUserDTO | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'age_verified': value['ageVerified'],
        'biography': value['biography'],
        'cover_image': value['coverImage'],
        'cover_image_thumbnail': value['coverImageThumbnail'],
        'created_at': value['createdAt'],
        'dangerous_user': value['dangerousUser'],
        'follow_pending': value['followPending'],
        'followed_by': value['followedBy'],
        'followers_count': value['followersCount'],
        'following': value['following'],
        'followings_count': value['followingsCount'],
        'gender': value['gender'],
        'groups_users_count': value['groupsUsersCount'],
        'hidden': value['hidden'],
        'hide_vip': value['hideVip'],
        'id': value['id'],
        'is_private': value['isPrivate'],
        'new_user': value['newUser'],
        'nickname': value['nickname'],
        'posts_count': value['postsCount'],
        'prefecture': value['prefecture'],
        'profile_icon': value['profileIcon'],
        'profile_icon_thumbnail': value['profileIconThumbnail'],
        'recently_kenta': value['recentlyKenta'],
        'reviews_count': value['reviewsCount'],
        'title': value['title'],
        'vip': value['vip'],
    };
}

