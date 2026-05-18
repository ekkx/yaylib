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

/**
 * 
 * @export
 * @interface Group
 */
export interface Group {
    /**
     * 
     * @type {boolean}
     * @memberof Group
     */
    allowMembersToPostImageAndVideo?: boolean | null;
    /**
     * 
     * @type {boolean}
     * @memberof Group
     */
    allowMembersToPostUrl?: boolean | null;
    /**
     * 
     * @type {boolean}
     * @memberof Group
     */
    allowOwnershipTransfer?: boolean | null;
    /**
     * 
     * @type {string}
     * @memberof Group
     */
    allowThreadCreationBy?: string | null;
    /**
     * 
     * @type {boolean}
     * @memberof Group
     */
    callTimelineDisplay?: boolean | null;
    /**
     * 
     * @type {string}
     * @memberof Group
     */
    coverImage?: string | null;
    /**
     * 
     * @type {string}
     * @memberof Group
     */
    coverImageThumbnail?: string | null;
    /**
     * 
     * @type {string}
     * @memberof Group
     */
    description?: string | null;
    /**
     * 
     * @type {number}
     * @memberof Group
     */
    gender?: number | null;
    /**
     * 
     * @type {number}
     * @memberof Group
     */
    generationGroupsLimit?: number | null;
    /**
     * 
     * @type {number}
     * @memberof Group
     */
    groupCategoryId?: number | null;
    /**
     * 
     * @type {string}
     * @memberof Group
     */
    groupIcon?: string | null;
    /**
     * 
     * @type {string}
     * @memberof Group
     */
    groupIconThumbnail?: string | null;
    /**
     * 
     * @type {number}
     * @memberof Group
     */
    groupsUsersCount?: number | null;
    /**
     * 
     * @type {string}
     * @memberof Group
     */
    guidelines?: string | null;
    /**
     * 
     * @type {boolean}
     * @memberof Group
     */
    hideConferenceCall?: boolean | null;
    /**
     * 
     * @type {boolean}
     * @memberof Group
     */
    hideFromGameEight?: boolean | null;
    /**
     * 
     * @type {boolean}
     * @memberof Group
     */
    hideReportedPosts?: boolean | null;
    /**
     * 
     * @type {number}
     * @memberof Group
     */
    highlightedCount?: number | null;
    /**
     * 
     * @type {number}
     * @memberof Group
     */
    id?: number | null;
    /**
     * 
     * @type {boolean}
     * @memberof Group
     */
    invitedToJoin?: boolean | null;
    /**
     * 
     * @type {boolean}
     * @memberof Group
     */
    isJoined?: boolean | null;
    /**
     * 
     * @type {boolean}
     * @memberof Group
     */
    isPending?: boolean | null;
    /**
     * 
     * @type {boolean}
     * @memberof Group
     */
    isPrivate?: boolean | null;
    /**
     * 
     * @type {boolean}
     * @memberof Group
     */
    isRelated?: boolean | null;
    /**
     * 
     * @type {boolean}
     * @memberof Group
     */
    joinedCommunityCampaign?: boolean | null;
    /**
     * 
     * @type {Array<number>}
     * @memberof Group
     */
    moderatorIds?: Array<number> | null;
    /**
     * 
     * @type {boolean}
     * @memberof Group
     */
    onlyMobileVerified?: boolean | null;
    /**
     * 
     * @type {boolean}
     * @memberof Group
     */
    onlyVerifiedAge?: boolean | null;
    /**
     * 
     * @type {RealmUser}
     * @memberof Group
     */
    owner?: RealmUser | null;
    /**
     * 
     * @type {number}
     * @memberof Group
     */
    pendingCount?: number | null;
    /**
     * 
     * @type {Array<number>}
     * @memberof Group
     */
    pendingDeputizeIds?: Array<number> | null;
    /**
     * 
     * @type {number}
     * @memberof Group
     */
    pendingTransferId?: number | null;
    /**
     * 
     * @type {number}
     * @memberof Group
     */
    postsCount?: number | null;
    /**
     * 
     * @type {number}
     * @memberof Group
     */
    relatedCount?: number | null;
    /**
     * 
     * @type {boolean}
     * @memberof Group
     */
    safeMode?: boolean | null;
    /**
     * 
     * @type {boolean}
     * @memberof Group
     */
    secret?: boolean | null;
    /**
     * 
     * @type {boolean}
     * @memberof Group
     */
    seizable?: boolean | null;
    /**
     * 
     * @type {number}
     * @memberof Group
     */
    seizableBefore?: number | null;
    /**
     * 
     * @type {number}
     * @memberof Group
     */
    subCategoryId?: number | null;
    /**
     * 
     * @type {number}
     * @memberof Group
     */
    threadsCount?: number | null;
    /**
     * 
     * @type {string}
     * @memberof Group
     */
    title?: string | null;
    /**
     * 
     * @type {string}
     * @memberof Group
     */
    topic?: string | null;
    /**
     * 
     * @type {number}
     * @memberof Group
     */
    unreadCounts?: number | null;
    /**
     * 
     * @type {number}
     * @memberof Group
     */
    unreadThreadsCount?: number | null;
    /**
     * 
     * @type {number}
     * @memberof Group
     */
    updatedAt?: number | null;
    /**
     * 
     * @type {number}
     * @memberof Group
     */
    userId?: number | null;
    /**
     * 
     * @type {number}
     * @memberof Group
     */
    viewsCount?: number | null;
    /**
     * 
     * @type {boolean}
     * @memberof Group
     */
    walkthroughRequested?: boolean | null;
}

/**
 * Check if a given object implements the Group interface.
 */
export function instanceOfGroup(value: object): value is Group {
    return true;
}

export function GroupFromJSON(json: any): Group {
    return GroupFromJSONTyped(json, false);
}

export function GroupFromJSONTyped(json: any, ignoreDiscriminator: boolean): Group {
    if (json == null) {
        return json;
    }
    return {
        
        'allowMembersToPostImageAndVideo': json['allow_members_to_post_image_and_video'] == null ? undefined : json['allow_members_to_post_image_and_video'],
        'allowMembersToPostUrl': json['allow_members_to_post_url'] == null ? undefined : json['allow_members_to_post_url'],
        'allowOwnershipTransfer': json['allow_ownership_transfer'] == null ? undefined : json['allow_ownership_transfer'],
        'allowThreadCreationBy': json['allow_thread_creation_by'] == null ? undefined : json['allow_thread_creation_by'],
        'callTimelineDisplay': json['call_timeline_display'] == null ? undefined : json['call_timeline_display'],
        'coverImage': json['cover_image'] == null ? undefined : json['cover_image'],
        'coverImageThumbnail': json['cover_image_thumbnail'] == null ? undefined : json['cover_image_thumbnail'],
        'description': json['description'] == null ? undefined : json['description'],
        'gender': json['gender'] == null ? undefined : json['gender'],
        'generationGroupsLimit': json['generation_groups_limit'] == null ? undefined : json['generation_groups_limit'],
        'groupCategoryId': json['group_category_id'] == null ? undefined : json['group_category_id'],
        'groupIcon': json['group_icon'] == null ? undefined : json['group_icon'],
        'groupIconThumbnail': json['group_icon_thumbnail'] == null ? undefined : json['group_icon_thumbnail'],
        'groupsUsersCount': json['groups_users_count'] == null ? undefined : json['groups_users_count'],
        'guidelines': json['guidelines'] == null ? undefined : json['guidelines'],
        'hideConferenceCall': json['hide_conference_call'] == null ? undefined : json['hide_conference_call'],
        'hideFromGameEight': json['hide_from_game_eight'] == null ? undefined : json['hide_from_game_eight'],
        'hideReportedPosts': json['hide_reported_posts'] == null ? undefined : json['hide_reported_posts'],
        'highlightedCount': json['highlighted_count'] == null ? undefined : json['highlighted_count'],
        'id': json['id'] == null ? undefined : json['id'],
        'invitedToJoin': json['invited_to_join'] == null ? undefined : json['invited_to_join'],
        'isJoined': json['is_joined'] == null ? undefined : json['is_joined'],
        'isPending': json['is_pending'] == null ? undefined : json['is_pending'],
        'isPrivate': json['is_private'] == null ? undefined : json['is_private'],
        'isRelated': json['is_related'] == null ? undefined : json['is_related'],
        'joinedCommunityCampaign': json['joined_community_campaign'] == null ? undefined : json['joined_community_campaign'],
        'moderatorIds': json['moderator_ids'] == null ? undefined : json['moderator_ids'],
        'onlyMobileVerified': json['only_mobile_verified'] == null ? undefined : json['only_mobile_verified'],
        'onlyVerifiedAge': json['only_verified_age'] == null ? undefined : json['only_verified_age'],
        'owner': json['owner'] == null ? undefined : RealmUserFromJSON(json['owner']),
        'pendingCount': json['pending_count'] == null ? undefined : json['pending_count'],
        'pendingDeputizeIds': json['pending_deputize_ids'] == null ? undefined : json['pending_deputize_ids'],
        'pendingTransferId': json['pending_transfer_id'] == null ? undefined : json['pending_transfer_id'],
        'postsCount': json['posts_count'] == null ? undefined : json['posts_count'],
        'relatedCount': json['related_count'] == null ? undefined : json['related_count'],
        'safeMode': json['safe_mode'] == null ? undefined : json['safe_mode'],
        'secret': json['secret'] == null ? undefined : json['secret'],
        'seizable': json['seizable'] == null ? undefined : json['seizable'],
        'seizableBefore': json['seizable_before'] == null ? undefined : json['seizable_before'],
        'subCategoryId': json['sub_category_id'] == null ? undefined : json['sub_category_id'],
        'threadsCount': json['threads_count'] == null ? undefined : json['threads_count'],
        'title': json['title'] == null ? undefined : json['title'],
        'topic': json['topic'] == null ? undefined : json['topic'],
        'unreadCounts': json['unread_counts'] == null ? undefined : json['unread_counts'],
        'unreadThreadsCount': json['unread_threads_count'] == null ? undefined : json['unread_threads_count'],
        'updatedAt': json['updated_at'] == null ? undefined : json['updated_at'],
        'userId': json['user_id'] == null ? undefined : json['user_id'],
        'viewsCount': json['views_count'] == null ? undefined : json['views_count'],
        'walkthroughRequested': json['walkthrough_requested'] == null ? undefined : json['walkthrough_requested'],
    };
}

export function GroupToJSON(json: any): Group {
    return GroupToJSONTyped(json, false);
}

export function GroupToJSONTyped(value?: Group | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'allow_members_to_post_image_and_video': value['allowMembersToPostImageAndVideo'],
        'allow_members_to_post_url': value['allowMembersToPostUrl'],
        'allow_ownership_transfer': value['allowOwnershipTransfer'],
        'allow_thread_creation_by': value['allowThreadCreationBy'],
        'call_timeline_display': value['callTimelineDisplay'],
        'cover_image': value['coverImage'],
        'cover_image_thumbnail': value['coverImageThumbnail'],
        'description': value['description'],
        'gender': value['gender'],
        'generation_groups_limit': value['generationGroupsLimit'],
        'group_category_id': value['groupCategoryId'],
        'group_icon': value['groupIcon'],
        'group_icon_thumbnail': value['groupIconThumbnail'],
        'groups_users_count': value['groupsUsersCount'],
        'guidelines': value['guidelines'],
        'hide_conference_call': value['hideConferenceCall'],
        'hide_from_game_eight': value['hideFromGameEight'],
        'hide_reported_posts': value['hideReportedPosts'],
        'highlighted_count': value['highlightedCount'],
        'id': value['id'],
        'invited_to_join': value['invitedToJoin'],
        'is_joined': value['isJoined'],
        'is_pending': value['isPending'],
        'is_private': value['isPrivate'],
        'is_related': value['isRelated'],
        'joined_community_campaign': value['joinedCommunityCampaign'],
        'moderator_ids': value['moderatorIds'],
        'only_mobile_verified': value['onlyMobileVerified'],
        'only_verified_age': value['onlyVerifiedAge'],
        'owner': RealmUserToJSON(value['owner']),
        'pending_count': value['pendingCount'],
        'pending_deputize_ids': value['pendingDeputizeIds'],
        'pending_transfer_id': value['pendingTransferId'],
        'posts_count': value['postsCount'],
        'related_count': value['relatedCount'],
        'safe_mode': value['safeMode'],
        'secret': value['secret'],
        'seizable': value['seizable'],
        'seizable_before': value['seizableBefore'],
        'sub_category_id': value['subCategoryId'],
        'threads_count': value['threadsCount'],
        'title': value['title'],
        'topic': value['topic'],
        'unread_counts': value['unreadCounts'],
        'unread_threads_count': value['unreadThreadsCount'],
        'updated_at': value['updatedAt'],
        'user_id': value['userId'],
        'views_count': value['viewsCount'],
        'walkthrough_requested': value['walkthroughRequested'],
    };
}

