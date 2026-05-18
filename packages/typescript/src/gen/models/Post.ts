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
import type { Shareable } from './Shareable.js';
import {
    ShareableFromJSON,
    ShareableFromJSONTyped,
    ShareableToJSON,
    ShareableToJSONTyped,
} from './Shareable.js';
import type { Survey } from './Survey.js';
import {
    SurveyFromJSON,
    SurveyFromJSONTyped,
    SurveyToJSON,
    SurveyToJSONTyped,
} from './Survey.js';
import type { PostType } from './PostType.js';
import {
    PostTypeFromJSON,
    PostTypeFromJSONTyped,
    PostTypeToJSON,
    PostTypeToJSONTyped,
} from './PostType.js';
import type { RealmConferenceCall } from './RealmConferenceCall.js';
import {
    RealmConferenceCallFromJSON,
    RealmConferenceCallFromJSONTyped,
    RealmConferenceCallToJSON,
    RealmConferenceCallToJSONTyped,
} from './RealmConferenceCall.js';
import type { ThreadInfo } from './ThreadInfo.js';
import {
    ThreadInfoFromJSON,
    ThreadInfoFromJSONTyped,
    ThreadInfoToJSON,
    ThreadInfoToJSONTyped,
} from './ThreadInfo.js';
import type { Video } from './Video.js';
import {
    VideoFromJSON,
    VideoFromJSONTyped,
    VideoToJSON,
    VideoToJSONTyped,
} from './Video.js';
import type { GiftCount } from './GiftCount.js';
import {
    GiftCountFromJSON,
    GiftCountFromJSONTyped,
    GiftCountToJSON,
    GiftCountToJSONTyped,
} from './GiftCount.js';
import type { SharedUrl } from './SharedUrl.js';
import {
    SharedUrlFromJSON,
    SharedUrlFromJSONTyped,
    SharedUrlToJSON,
    SharedUrlToJSONTyped,
} from './SharedUrl.js';
import type { MessageTag } from './MessageTag.js';
import {
    MessageTagFromJSON,
    MessageTagFromJSONTyped,
    MessageTagToJSON,
    MessageTagToJSONTyped,
} from './MessageTag.js';
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
 * @interface Post
 */
export interface Post {
    /**
     * 
     * @type {string}
     * @memberof Post
     */
    attachment?: string | null;
    /**
     * 
     * @type {string}
     * @memberof Post
     */
    attachment2?: string | null;
    /**
     * 
     * @type {string}
     * @memberof Post
     */
    attachment2Thumbnail?: string | null;
    /**
     * 
     * @type {string}
     * @memberof Post
     */
    attachment3?: string | null;
    /**
     * 
     * @type {string}
     * @memberof Post
     */
    attachment3Thumbnail?: string | null;
    /**
     * 
     * @type {string}
     * @memberof Post
     */
    attachment4?: string | null;
    /**
     * 
     * @type {string}
     * @memberof Post
     */
    attachment4Thumbnail?: string | null;
    /**
     * 
     * @type {string}
     * @memberof Post
     */
    attachment5?: string | null;
    /**
     * 
     * @type {string}
     * @memberof Post
     */
    attachment5Thumbnail?: string | null;
    /**
     * 
     * @type {string}
     * @memberof Post
     */
    attachment6?: string | null;
    /**
     * 
     * @type {string}
     * @memberof Post
     */
    attachment6Thumbnail?: string | null;
    /**
     * 
     * @type {string}
     * @memberof Post
     */
    attachment7?: string | null;
    /**
     * 
     * @type {string}
     * @memberof Post
     */
    attachment7Thumbnail?: string | null;
    /**
     * 
     * @type {string}
     * @memberof Post
     */
    attachment8?: string | null;
    /**
     * 
     * @type {string}
     * @memberof Post
     */
    attachment8Thumbnail?: string | null;
    /**
     * 
     * @type {string}
     * @memberof Post
     */
    attachment9?: string | null;
    /**
     * 
     * @type {string}
     * @memberof Post
     */
    attachment9Thumbnail?: string | null;
    /**
     * 
     * @type {string}
     * @memberof Post
     */
    attachmentThumbnail?: string | null;
    /**
     * 
     * @type {number}
     * @memberof Post
     */
    color?: number | null;
    /**
     * 
     * @type {RealmConferenceCall}
     * @memberof Post
     */
    conferenceCall?: RealmConferenceCall | null;
    /**
     * 
     * @type {number}
     * @memberof Post
     */
    conversationId?: number | null;
    /**
     * 
     * @type {number}
     * @memberof Post
     */
    createdAt?: number | null;
    /**
     * 
     * @type {number}
     * @memberof Post
     */
    editedAt?: number | null;
    /**
     * 
     * @type {number}
     * @memberof Post
     */
    fontSize?: number | null;
    /**
     * 
     * @type {Array<GiftCount>}
     * @memberof Post
     */
    giftsCount?: Array<GiftCount> | null;
    /**
     * 
     * @type {Group}
     * @memberof Post
     */
    group?: Group | null;
    /**
     * 
     * @type {number}
     * @memberof Post
     */
    groupId?: number | null;
    /**
     * 
     * @type {boolean}
     * @memberof Post
     */
    highlighted?: boolean | null;
    /**
     * 
     * @type {number}
     * @memberof Post
     */
    id?: number | null;
    /**
     * 
     * @type {number}
     * @memberof Post
     */
    inReplyTo?: number | null;
    /**
     * 
     * @type {Post}
     * @memberof Post
     */
    inReplyToPost?: Post | null;
    /**
     * 
     * @type {number}
     * @memberof Post
     */
    inReplyToPostCount?: number | null;
    /**
     * 
     * @type {boolean}
     * @memberof Post
     */
    isFailToSend?: boolean | null;
    /**
     * 
     * @type {boolean}
     * @memberof Post
     */
    isMuted?: boolean | null;
    /**
     * 
     * @type {boolean}
     * @memberof Post
     */
    liked?: boolean | null;
    /**
     * 
     * @type {Array<RealmUser>}
     * @memberof Post
     */
    likers?: Array<RealmUser> | null;
    /**
     * 
     * @type {number}
     * @memberof Post
     */
    likersCount?: number | null;
    /**
     * 
     * @type {number}
     * @memberof Post
     */
    likesCount?: number | null;
    /**
     * 
     * @type {Array<RealmUser>}
     * @memberof Post
     */
    mentions?: Array<RealmUser> | null;
    /**
     * 
     * @type {Array<MessageTag>}
     * @memberof Post
     */
    messageTags?: Array<MessageTag> | null;
    /**
     * 
     * @type {PostType}
     * @memberof Post
     */
    postType?: PostType | null;
    /**
     * 
     * @type {number}
     * @memberof Post
     */
    reportedCount?: number | null;
    /**
     * 
     * @type {boolean}
     * @memberof Post
     */
    repostable?: boolean | null;
    /**
     * 
     * @type {boolean}
     * @memberof Post
     */
    reposted?: boolean | null;
    /**
     * 
     * @type {number}
     * @memberof Post
     */
    repostsCount?: number | null;
    /**
     * 
     * @type {Shareable}
     * @memberof Post
     */
    shareable?: Shareable | null;
    /**
     * 
     * @type {ThreadInfo}
     * @memberof Post
     */
    sharedThread?: ThreadInfo | null;
    /**
     * 
     * @type {SharedUrl}
     * @memberof Post
     */
    sharedUrl?: SharedUrl | null;
    /**
     * 
     * @type {Survey}
     * @memberof Post
     */
    survey?: Survey | null;
    /**
     * 
     * @type {string}
     * @memberof Post
     */
    tag?: string | null;
    /**
     * 
     * @type {string}
     * @memberof Post
     */
    text?: string | null;
    /**
     * 
     * @type {ThreadInfo}
     * @memberof Post
     */
    thread?: ThreadInfo | null;
    /**
     * 
     * @type {number}
     * @memberof Post
     */
    threadId?: number | null;
    /**
     * 
     * @type {number}
     * @memberof Post
     */
    totalGiftsCount?: number | null;
    /**
     * 
     * @type {number}
     * @memberof Post
     */
    updatedAt?: number | null;
    /**
     * 
     * @type {RealmUser}
     * @memberof Post
     */
    user?: RealmUser | null;
    /**
     * 
     * @type {Array<Video>}
     * @memberof Post
     */
    videos?: Array<Video> | null;
}



/**
 * Check if a given object implements the Post interface.
 */
export function instanceOfPost(value: object): value is Post {
    return true;
}

export function PostFromJSON(json: any): Post {
    return PostFromJSONTyped(json, false);
}

export function PostFromJSONTyped(json: any, ignoreDiscriminator: boolean): Post {
    if (json == null) {
        return json;
    }
    return {
        
        'attachment': json['attachment'] == null ? undefined : json['attachment'],
        'attachment2': json['attachment_2'] == null ? undefined : json['attachment_2'],
        'attachment2Thumbnail': json['attachment_2_thumbnail'] == null ? undefined : json['attachment_2_thumbnail'],
        'attachment3': json['attachment_3'] == null ? undefined : json['attachment_3'],
        'attachment3Thumbnail': json['attachment_3_thumbnail'] == null ? undefined : json['attachment_3_thumbnail'],
        'attachment4': json['attachment_4'] == null ? undefined : json['attachment_4'],
        'attachment4Thumbnail': json['attachment_4_thumbnail'] == null ? undefined : json['attachment_4_thumbnail'],
        'attachment5': json['attachment_5'] == null ? undefined : json['attachment_5'],
        'attachment5Thumbnail': json['attachment_5_thumbnail'] == null ? undefined : json['attachment_5_thumbnail'],
        'attachment6': json['attachment_6'] == null ? undefined : json['attachment_6'],
        'attachment6Thumbnail': json['attachment_6_thumbnail'] == null ? undefined : json['attachment_6_thumbnail'],
        'attachment7': json['attachment_7'] == null ? undefined : json['attachment_7'],
        'attachment7Thumbnail': json['attachment_7_thumbnail'] == null ? undefined : json['attachment_7_thumbnail'],
        'attachment8': json['attachment_8'] == null ? undefined : json['attachment_8'],
        'attachment8Thumbnail': json['attachment_8_thumbnail'] == null ? undefined : json['attachment_8_thumbnail'],
        'attachment9': json['attachment_9'] == null ? undefined : json['attachment_9'],
        'attachment9Thumbnail': json['attachment_9_thumbnail'] == null ? undefined : json['attachment_9_thumbnail'],
        'attachmentThumbnail': json['attachment_thumbnail'] == null ? undefined : json['attachment_thumbnail'],
        'color': json['color'] == null ? undefined : json['color'],
        'conferenceCall': json['conference_call'] == null ? undefined : RealmConferenceCallFromJSON(json['conference_call']),
        'conversationId': json['conversation_id'] == null ? undefined : json['conversation_id'],
        'createdAt': json['created_at'] == null ? undefined : json['created_at'],
        'editedAt': json['edited_at'] == null ? undefined : json['edited_at'],
        'fontSize': json['font_size'] == null ? undefined : json['font_size'],
        'giftsCount': json['gifts_count'] == null ? undefined : ((json['gifts_count'] as Array<any>).map(GiftCountFromJSON)),
        'group': json['group'] == null ? undefined : GroupFromJSON(json['group']),
        'groupId': json['group_id'] == null ? undefined : json['group_id'],
        'highlighted': json['highlighted'] == null ? undefined : json['highlighted'],
        'id': json['id'] == null ? undefined : json['id'],
        'inReplyTo': json['in_reply_to'] == null ? undefined : json['in_reply_to'],
        'inReplyToPost': json['in_reply_to_post'] == null ? undefined : PostFromJSON(json['in_reply_to_post']),
        'inReplyToPostCount': json['in_reply_to_post_count'] == null ? undefined : json['in_reply_to_post_count'],
        'isFailToSend': json['is_fail_to_send'] == null ? undefined : json['is_fail_to_send'],
        'isMuted': json['is_muted'] == null ? undefined : json['is_muted'],
        'liked': json['liked'] == null ? undefined : json['liked'],
        'likers': json['likers'] == null ? undefined : ((json['likers'] as Array<any>).map(RealmUserFromJSON)),
        'likersCount': json['likers_count'] == null ? undefined : json['likers_count'],
        'likesCount': json['likes_count'] == null ? undefined : json['likes_count'],
        'mentions': json['mentions'] == null ? undefined : ((json['mentions'] as Array<any>).map(RealmUserFromJSON)),
        'messageTags': json['message_tags'] == null ? undefined : ((json['message_tags'] as Array<any>).map(MessageTagFromJSON)),
        'postType': json['post_type'] == null ? undefined : PostTypeFromJSON(json['post_type']),
        'reportedCount': json['reported_count'] == null ? undefined : json['reported_count'],
        'repostable': json['repostable'] == null ? undefined : json['repostable'],
        'reposted': json['reposted'] == null ? undefined : json['reposted'],
        'repostsCount': json['reposts_count'] == null ? undefined : json['reposts_count'],
        'shareable': json['shareable'] == null ? undefined : ShareableFromJSON(json['shareable']),
        'sharedThread': json['shared_thread'] == null ? undefined : ThreadInfoFromJSON(json['shared_thread']),
        'sharedUrl': json['shared_url'] == null ? undefined : SharedUrlFromJSON(json['shared_url']),
        'survey': json['survey'] == null ? undefined : SurveyFromJSON(json['survey']),
        'tag': json['tag'] == null ? undefined : json['tag'],
        'text': json['text'] == null ? undefined : json['text'],
        'thread': json['thread'] == null ? undefined : ThreadInfoFromJSON(json['thread']),
        'threadId': json['thread_id'] == null ? undefined : json['thread_id'],
        'totalGiftsCount': json['total_gifts_count'] == null ? undefined : json['total_gifts_count'],
        'updatedAt': json['updated_at'] == null ? undefined : json['updated_at'],
        'user': json['user'] == null ? undefined : RealmUserFromJSON(json['user']),
        'videos': json['videos'] == null ? undefined : ((json['videos'] as Array<any>).map(VideoFromJSON)),
    };
}

export function PostToJSON(json: any): Post {
    return PostToJSONTyped(json, false);
}

export function PostToJSONTyped(value?: Post | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'attachment': value['attachment'],
        'attachment_2': value['attachment2'],
        'attachment_2_thumbnail': value['attachment2Thumbnail'],
        'attachment_3': value['attachment3'],
        'attachment_3_thumbnail': value['attachment3Thumbnail'],
        'attachment_4': value['attachment4'],
        'attachment_4_thumbnail': value['attachment4Thumbnail'],
        'attachment_5': value['attachment5'],
        'attachment_5_thumbnail': value['attachment5Thumbnail'],
        'attachment_6': value['attachment6'],
        'attachment_6_thumbnail': value['attachment6Thumbnail'],
        'attachment_7': value['attachment7'],
        'attachment_7_thumbnail': value['attachment7Thumbnail'],
        'attachment_8': value['attachment8'],
        'attachment_8_thumbnail': value['attachment8Thumbnail'],
        'attachment_9': value['attachment9'],
        'attachment_9_thumbnail': value['attachment9Thumbnail'],
        'attachment_thumbnail': value['attachmentThumbnail'],
        'color': value['color'],
        'conference_call': RealmConferenceCallToJSON(value['conferenceCall']),
        'conversation_id': value['conversationId'],
        'created_at': value['createdAt'],
        'edited_at': value['editedAt'],
        'font_size': value['fontSize'],
        'gifts_count': value['giftsCount'] == null ? undefined : ((value['giftsCount'] as Array<any>).map(GiftCountToJSON)),
        'group': GroupToJSON(value['group']),
        'group_id': value['groupId'],
        'highlighted': value['highlighted'],
        'id': value['id'],
        'in_reply_to': value['inReplyTo'],
        'in_reply_to_post': PostToJSON(value['inReplyToPost']),
        'in_reply_to_post_count': value['inReplyToPostCount'],
        'is_fail_to_send': value['isFailToSend'],
        'is_muted': value['isMuted'],
        'liked': value['liked'],
        'likers': value['likers'] == null ? undefined : ((value['likers'] as Array<any>).map(RealmUserToJSON)),
        'likers_count': value['likersCount'],
        'likes_count': value['likesCount'],
        'mentions': value['mentions'] == null ? undefined : ((value['mentions'] as Array<any>).map(RealmUserToJSON)),
        'message_tags': value['messageTags'] == null ? undefined : ((value['messageTags'] as Array<any>).map(MessageTagToJSON)),
        'post_type': PostTypeToJSON(value['postType']),
        'reported_count': value['reportedCount'],
        'repostable': value['repostable'],
        'reposted': value['reposted'],
        'reposts_count': value['repostsCount'],
        'shareable': ShareableToJSON(value['shareable']),
        'shared_thread': ThreadInfoToJSON(value['sharedThread']),
        'shared_url': SharedUrlToJSON(value['sharedUrl']),
        'survey': SurveyToJSON(value['survey']),
        'tag': value['tag'],
        'text': value['text'],
        'thread': ThreadInfoToJSON(value['thread']),
        'thread_id': value['threadId'],
        'total_gifts_count': value['totalGiftsCount'],
        'updated_at': value['updatedAt'],
        'user': RealmUserToJSON(value['user']),
        'videos': value['videos'] == null ? undefined : ((value['videos'] as Array<any>).map(VideoToJSON)),
    };
}

