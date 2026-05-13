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
import type { Generation } from './Generation';
import {
    GenerationFromJSON,
    GenerationFromJSONTyped,
    GenerationToJSON,
    GenerationToJSONTyped,
} from './Generation';
import type { User } from './User';
import {
    UserFromJSON,
    UserFromJSONTyped,
    UserToJSON,
    UserToJSONTyped,
} from './User';
import type { GroupRole } from './GroupRole';
import {
    GroupRoleFromJSON,
    GroupRoleFromJSONTyped,
    GroupRoleToJSON,
    GroupRoleToJSONTyped,
} from './GroupRole';
import type { Gender } from './Gender';
import {
    GenderFromJSON,
    GenderFromJSONTyped,
    GenderToJSON,
    GenderToJSONTyped,
} from './Gender';

/**
 * 
 * @export
 * @interface ModelGroup
 */
export interface ModelGroup {
    /**
     * 
     * @type {boolean}
     * @memberof ModelGroup
     */
    allowMembersToPostMedia?: boolean | null;
    /**
     * 
     * @type {boolean}
     * @memberof ModelGroup
     */
    allowMembersToPostUrl?: boolean | null;
    /**
     * 
     * @type {boolean}
     * @memberof ModelGroup
     */
    allowOwnershipTransfer?: boolean | null;
    /**
     * 
     * @type {GroupRole}
     * @memberof ModelGroup
     */
    allowThreadCreationBy?: GroupRole | null;
    /**
     * 
     * @type {string}
     * @memberof ModelGroup
     */
    coverImage?: string | null;
    /**
     * 
     * @type {string}
     * @memberof ModelGroup
     */
    coverImageThumbnail?: string | null;
    /**
     * 
     * @type {string}
     * @memberof ModelGroup
     */
    description?: string | null;
    /**
     * 
     * @type {string}
     * @memberof ModelGroup
     */
    displayIconThumbnail?: string | null;
    /**
     * 
     * @type {Gender}
     * @memberof ModelGroup
     */
    gender?: Gender | null;
    /**
     * 
     * @type {Generation}
     * @memberof ModelGroup
     */
    generation?: Generation | null;
    /**
     * 
     * @type {number}
     * @memberof ModelGroup
     */
    groupCategoryId?: number | null;
    /**
     * 
     * @type {number}
     * @memberof ModelGroup
     */
    groupsUsersCount?: number | null;
    /**
     * 
     * @type {string}
     * @memberof ModelGroup
     */
    guidelines?: string | null;
    /**
     * 
     * @type {number}
     * @memberof ModelGroup
     */
    highlightedCount?: number | null;
    /**
     * 
     * @type {string}
     * @memberof ModelGroup
     */
    iconImage?: string | null;
    /**
     * 
     * @type {string}
     * @memberof ModelGroup
     */
    iconImageThumbnail?: string | null;
    /**
     * 
     * @type {number}
     * @memberof ModelGroup
     */
    id?: number | null;
    /**
     * 
     * @type {boolean}
     * @memberof ModelGroup
     */
    invitedToJoin?: boolean | null;
    /**
     * 
     * @type {boolean}
     * @memberof ModelGroup
     */
    isCallTimelineDisplay?: boolean | null;
    /**
     * 
     * @type {boolean}
     * @memberof ModelGroup
     */
    isHideConferenceCall?: boolean | null;
    /**
     * 
     * @type {boolean}
     * @memberof ModelGroup
     */
    isHideFromGameEight?: boolean | null;
    /**
     * 
     * @type {boolean}
     * @memberof ModelGroup
     */
    isHideReportedPosts?: boolean | null;
    /**
     * 
     * @type {boolean}
     * @memberof ModelGroup
     */
    isJoined?: boolean | null;
    /**
     * 
     * @type {boolean}
     * @memberof ModelGroup
     */
    isOnlyVerifiedAge?: boolean | null;
    /**
     * 
     * @type {boolean}
     * @memberof ModelGroup
     */
    isOnlyVerifiedPhone?: boolean | null;
    /**
     * 
     * @type {boolean}
     * @memberof ModelGroup
     */
    isPending?: boolean | null;
    /**
     * 
     * @type {boolean}
     * @memberof ModelGroup
     */
    isPinned?: boolean | null;
    /**
     * 
     * @type {boolean}
     * @memberof ModelGroup
     */
    isPrivate?: boolean | null;
    /**
     * 
     * @type {boolean}
     * @memberof ModelGroup
     */
    isSafeMode?: boolean | null;
    /**
     * 
     * @type {boolean}
     * @memberof ModelGroup
     */
    isSecret?: boolean | null;
    /**
     * 
     * @type {boolean}
     * @memberof ModelGroup
     */
    isSeizable?: boolean | null;
    /**
     * 
     * @type {boolean}
     * @memberof ModelGroup
     */
    isWalkthroughRequested?: boolean | null;
    /**
     * 
     * @type {object}
     * @memberof ModelGroup
     */
    joinStatus?: object | null;
    /**
     * 
     * @type {boolean}
     * @memberof ModelGroup
     */
    joinedCommunityCampaign?: boolean | null;
    /**
     * 
     * @type {Array<number>}
     * @memberof ModelGroup
     */
    moderatorIds?: Array<number> | null;
    /**
     * 
     * @type {User}
     * @memberof ModelGroup
     */
    owner?: User | null;
    /**
     * 
     * @type {number}
     * @memberof ModelGroup
     */
    pendingCount?: number | null;
    /**
     * 
     * @type {Array<number>}
     * @memberof ModelGroup
     */
    pendingDeputizeIds?: Array<number> | null;
    /**
     * 
     * @type {number}
     * @memberof ModelGroup
     */
    pendingTransferId?: number | null;
    /**
     * 
     * @type {number}
     * @memberof ModelGroup
     */
    postsCount?: number | null;
    /**
     * 
     * @type {number}
     * @memberof ModelGroup
     */
    relatedCount?: number | null;
    /**
     * 
     * @type {number}
     * @memberof ModelGroup
     */
    seizableBeforeMillis?: number | null;
    /**
     * 
     * @type {number}
     * @memberof ModelGroup
     */
    subCategoryId?: number | null;
    /**
     * 
     * @type {number}
     * @memberof ModelGroup
     */
    threadsCount?: number | null;
    /**
     * 
     * @type {object}
     * @memberof ModelGroup
     */
    title?: object | null;
    /**
     * 
     * @type {string}
     * @memberof ModelGroup
     */
    topic?: string | null;
    /**
     * 
     * @type {number}
     * @memberof ModelGroup
     */
    unreadCounts?: number | null;
    /**
     * 
     * @type {number}
     * @memberof ModelGroup
     */
    unreadThreadsCount?: number | null;
    /**
     * 
     * @type {number}
     * @memberof ModelGroup
     */
    updatedAtSeconds?: number | null;
    /**
     * 
     * @type {number}
     * @memberof ModelGroup
     */
    viewsCount?: number | null;
}

/**
 * Check if a given object implements the ModelGroup interface.
 */
export function instanceOfModelGroup(value: object): value is ModelGroup {
    return true;
}

export function ModelGroupFromJSON(json: any): ModelGroup {
    return ModelGroupFromJSONTyped(json, false);
}

export function ModelGroupFromJSONTyped(json: any, ignoreDiscriminator: boolean): ModelGroup {
    if (json == null) {
        return json;
    }
    return {
        
        'allowMembersToPostMedia': json['allow_members_to_post_media'] == null ? undefined : json['allow_members_to_post_media'],
        'allowMembersToPostUrl': json['allow_members_to_post_url'] == null ? undefined : json['allow_members_to_post_url'],
        'allowOwnershipTransfer': json['allow_ownership_transfer'] == null ? undefined : json['allow_ownership_transfer'],
        'allowThreadCreationBy': json['allow_thread_creation_by'] == null ? undefined : GroupRoleFromJSON(json['allow_thread_creation_by']),
        'coverImage': json['cover_image'] == null ? undefined : json['cover_image'],
        'coverImageThumbnail': json['cover_image_thumbnail'] == null ? undefined : json['cover_image_thumbnail'],
        'description': json['description'] == null ? undefined : json['description'],
        'displayIconThumbnail': json['display_icon_thumbnail'] == null ? undefined : json['display_icon_thumbnail'],
        'gender': json['gender'] == null ? undefined : GenderFromJSON(json['gender']),
        'generation': json['generation'] == null ? undefined : GenerationFromJSON(json['generation']),
        'groupCategoryId': json['group_category_id'] == null ? undefined : json['group_category_id'],
        'groupsUsersCount': json['groups_users_count'] == null ? undefined : json['groups_users_count'],
        'guidelines': json['guidelines'] == null ? undefined : json['guidelines'],
        'highlightedCount': json['highlighted_count'] == null ? undefined : json['highlighted_count'],
        'iconImage': json['icon_image'] == null ? undefined : json['icon_image'],
        'iconImageThumbnail': json['icon_image_thumbnail'] == null ? undefined : json['icon_image_thumbnail'],
        'id': json['id'] == null ? undefined : json['id'],
        'invitedToJoin': json['invited_to_join'] == null ? undefined : json['invited_to_join'],
        'isCallTimelineDisplay': json['is_call_timeline_display'] == null ? undefined : json['is_call_timeline_display'],
        'isHideConferenceCall': json['is_hide_conference_call'] == null ? undefined : json['is_hide_conference_call'],
        'isHideFromGameEight': json['is_hide_from_game_eight'] == null ? undefined : json['is_hide_from_game_eight'],
        'isHideReportedPosts': json['is_hide_reported_posts'] == null ? undefined : json['is_hide_reported_posts'],
        'isJoined': json['is_joined'] == null ? undefined : json['is_joined'],
        'isOnlyVerifiedAge': json['is_only_verified_age'] == null ? undefined : json['is_only_verified_age'],
        'isOnlyVerifiedPhone': json['is_only_verified_phone'] == null ? undefined : json['is_only_verified_phone'],
        'isPending': json['is_pending'] == null ? undefined : json['is_pending'],
        'isPinned': json['is_pinned'] == null ? undefined : json['is_pinned'],
        'isPrivate': json['is_private'] == null ? undefined : json['is_private'],
        'isSafeMode': json['is_safe_mode'] == null ? undefined : json['is_safe_mode'],
        'isSecret': json['is_secret'] == null ? undefined : json['is_secret'],
        'isSeizable': json['is_seizable'] == null ? undefined : json['is_seizable'],
        'isWalkthroughRequested': json['is_walkthrough_requested'] == null ? undefined : json['is_walkthrough_requested'],
        'joinStatus': json['join_status'] == null ? undefined : json['join_status'],
        'joinedCommunityCampaign': json['joined_community_campaign'] == null ? undefined : json['joined_community_campaign'],
        'moderatorIds': json['moderator_ids'] == null ? undefined : json['moderator_ids'],
        'owner': json['owner'] == null ? undefined : UserFromJSON(json['owner']),
        'pendingCount': json['pending_count'] == null ? undefined : json['pending_count'],
        'pendingDeputizeIds': json['pending_deputize_ids'] == null ? undefined : json['pending_deputize_ids'],
        'pendingTransferId': json['pending_transfer_id'] == null ? undefined : json['pending_transfer_id'],
        'postsCount': json['posts_count'] == null ? undefined : json['posts_count'],
        'relatedCount': json['related_count'] == null ? undefined : json['related_count'],
        'seizableBeforeMillis': json['seizable_before_millis'] == null ? undefined : json['seizable_before_millis'],
        'subCategoryId': json['sub_category_id'] == null ? undefined : json['sub_category_id'],
        'threadsCount': json['threads_count'] == null ? undefined : json['threads_count'],
        'title': json['title'] == null ? undefined : json['title'],
        'topic': json['topic'] == null ? undefined : json['topic'],
        'unreadCounts': json['unread_counts'] == null ? undefined : json['unread_counts'],
        'unreadThreadsCount': json['unread_threads_count'] == null ? undefined : json['unread_threads_count'],
        'updatedAtSeconds': json['updated_at_seconds'] == null ? undefined : json['updated_at_seconds'],
        'viewsCount': json['views_count'] == null ? undefined : json['views_count'],
    };
}

export function ModelGroupToJSON(json: any): ModelGroup {
    return ModelGroupToJSONTyped(json, false);
}

export function ModelGroupToJSONTyped(value?: ModelGroup | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'allow_members_to_post_media': value['allowMembersToPostMedia'],
        'allow_members_to_post_url': value['allowMembersToPostUrl'],
        'allow_ownership_transfer': value['allowOwnershipTransfer'],
        'allow_thread_creation_by': GroupRoleToJSON(value['allowThreadCreationBy']),
        'cover_image': value['coverImage'],
        'cover_image_thumbnail': value['coverImageThumbnail'],
        'description': value['description'],
        'display_icon_thumbnail': value['displayIconThumbnail'],
        'gender': GenderToJSON(value['gender']),
        'generation': GenerationToJSON(value['generation']),
        'group_category_id': value['groupCategoryId'],
        'groups_users_count': value['groupsUsersCount'],
        'guidelines': value['guidelines'],
        'highlighted_count': value['highlightedCount'],
        'icon_image': value['iconImage'],
        'icon_image_thumbnail': value['iconImageThumbnail'],
        'id': value['id'],
        'invited_to_join': value['invitedToJoin'],
        'is_call_timeline_display': value['isCallTimelineDisplay'],
        'is_hide_conference_call': value['isHideConferenceCall'],
        'is_hide_from_game_eight': value['isHideFromGameEight'],
        'is_hide_reported_posts': value['isHideReportedPosts'],
        'is_joined': value['isJoined'],
        'is_only_verified_age': value['isOnlyVerifiedAge'],
        'is_only_verified_phone': value['isOnlyVerifiedPhone'],
        'is_pending': value['isPending'],
        'is_pinned': value['isPinned'],
        'is_private': value['isPrivate'],
        'is_safe_mode': value['isSafeMode'],
        'is_secret': value['isSecret'],
        'is_seizable': value['isSeizable'],
        'is_walkthrough_requested': value['isWalkthroughRequested'],
        'join_status': value['joinStatus'],
        'joined_community_campaign': value['joinedCommunityCampaign'],
        'moderator_ids': value['moderatorIds'],
        'owner': UserToJSON(value['owner']),
        'pending_count': value['pendingCount'],
        'pending_deputize_ids': value['pendingDeputizeIds'],
        'pending_transfer_id': value['pendingTransferId'],
        'posts_count': value['postsCount'],
        'related_count': value['relatedCount'],
        'seizable_before_millis': value['seizableBeforeMillis'],
        'sub_category_id': value['subCategoryId'],
        'threads_count': value['threadsCount'],
        'title': value['title'],
        'topic': value['topic'],
        'unread_counts': value['unreadCounts'],
        'unread_threads_count': value['unreadThreadsCount'],
        'updated_at_seconds': value['updatedAtSeconds'],
        'views_count': value['viewsCount'],
    };
}

