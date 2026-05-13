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
import type { ModelGroup } from './ModelGroup';
import {
    ModelGroupFromJSON,
    ModelGroupFromJSONTyped,
    ModelGroupToJSON,
    ModelGroupToJSONTyped,
} from './ModelGroup';
import type { PostGift } from './PostGift';
import {
    PostGiftFromJSON,
    PostGiftFromJSONTyped,
    PostGiftToJSON,
    PostGiftToJSONTyped,
} from './PostGift';
import type { User } from './User';
import {
    UserFromJSON,
    UserFromJSONTyped,
    UserToJSON,
    UserToJSONTyped,
} from './User';
import type { ModelThreadInfo } from './ModelThreadInfo';
import {
    ModelThreadInfoFromJSON,
    ModelThreadInfoFromJSONTyped,
    ModelThreadInfoToJSON,
    ModelThreadInfoToJSONTyped,
} from './ModelThreadInfo';
import type { ModelVideo } from './ModelVideo';
import {
    ModelVideoFromJSON,
    ModelVideoFromJSONTyped,
    ModelVideoToJSON,
    ModelVideoToJSONTyped,
} from './ModelVideo';
import type { ModelMessageTag } from './ModelMessageTag';
import {
    ModelMessageTagFromJSON,
    ModelMessageTagFromJSONTyped,
    ModelMessageTagToJSON,
    ModelMessageTagToJSONTyped,
} from './ModelMessageTag';
import type { ConferenceCall } from './ConferenceCall';
import {
    ConferenceCallFromJSON,
    ConferenceCallFromJSONTyped,
    ConferenceCallToJSON,
    ConferenceCallToJSONTyped,
} from './ConferenceCall';
import type { ModelShareable } from './ModelShareable';
import {
    ModelShareableFromJSON,
    ModelShareableFromJSONTyped,
    ModelShareableToJSON,
    ModelShareableToJSONTyped,
} from './ModelShareable';
import type { ModelSharedUrl } from './ModelSharedUrl';
import {
    ModelSharedUrlFromJSON,
    ModelSharedUrlFromJSONTyped,
    ModelSharedUrlToJSON,
    ModelSharedUrlToJSONTyped,
} from './ModelSharedUrl';
import type { ModelSurvey } from './ModelSurvey';
import {
    ModelSurveyFromJSON,
    ModelSurveyFromJSONTyped,
    ModelSurveyToJSON,
    ModelSurveyToJSONTyped,
} from './ModelSurvey';

/**
 * 
 * @export
 * @interface ModelPost
 */
export interface ModelPost {
    /**
     * 
     * @type {string}
     * @memberof ModelPost
     */
    attachment?: string | null;
    /**
     * 
     * @type {string}
     * @memberof ModelPost
     */
    attachment2?: string | null;
    /**
     * 
     * @type {string}
     * @memberof ModelPost
     */
    attachment2Thumbnail?: string | null;
    /**
     * 
     * @type {string}
     * @memberof ModelPost
     */
    attachment3?: string | null;
    /**
     * 
     * @type {string}
     * @memberof ModelPost
     */
    attachment3Thumbnail?: string | null;
    /**
     * 
     * @type {string}
     * @memberof ModelPost
     */
    attachment4?: string | null;
    /**
     * 
     * @type {string}
     * @memberof ModelPost
     */
    attachment4Thumbnail?: string | null;
    /**
     * 
     * @type {string}
     * @memberof ModelPost
     */
    attachment5?: string | null;
    /**
     * 
     * @type {string}
     * @memberof ModelPost
     */
    attachment5Thumbnail?: string | null;
    /**
     * 
     * @type {string}
     * @memberof ModelPost
     */
    attachment6?: string | null;
    /**
     * 
     * @type {string}
     * @memberof ModelPost
     */
    attachment6Thumbnail?: string | null;
    /**
     * 
     * @type {string}
     * @memberof ModelPost
     */
    attachment7?: string | null;
    /**
     * 
     * @type {string}
     * @memberof ModelPost
     */
    attachment7Thumbnail?: string | null;
    /**
     * 
     * @type {string}
     * @memberof ModelPost
     */
    attachment8?: string | null;
    /**
     * 
     * @type {string}
     * @memberof ModelPost
     */
    attachment8Thumbnail?: string | null;
    /**
     * 
     * @type {string}
     * @memberof ModelPost
     */
    attachment9?: string | null;
    /**
     * 
     * @type {string}
     * @memberof ModelPost
     */
    attachment9Thumbnail?: string | null;
    /**
     * 
     * @type {string}
     * @memberof ModelPost
     */
    attachmentThumbnail?: string | null;
    /**
     * 
     * @type {number}
     * @memberof ModelPost
     */
    categoryId?: number | null;
    /**
     * 
     * @type {number}
     * @memberof ModelPost
     */
    color?: number | null;
    /**
     * 
     * @type {ConferenceCall}
     * @memberof ModelPost
     */
    conferenceCall?: ConferenceCall | null;
    /**
     * 
     * @type {number}
     * @memberof ModelPost
     */
    conversationId?: number | null;
    /**
     * 
     * @type {number}
     * @memberof ModelPost
     */
    createdAtSeconds?: number | null;
    /**
     * 
     * @type {number}
     * @memberof ModelPost
     */
    editedAtSeconds?: number | null;
    /**
     * 
     * @type {number}
     * @memberof ModelPost
     */
    fontSize?: number | null;
    /**
     * 
     * @type {string}
     * @memberof ModelPost
     */
    gameTitle?: string | null;
    /**
     * 
     * @type {Array<PostGift>}
     * @memberof ModelPost
     */
    gifts?: Array<PostGift> | null;
    /**
     * 
     * @type {ModelGroup}
     * @memberof ModelPost
     */
    group?: ModelGroup | null;
    /**
     * 
     * @type {number}
     * @memberof ModelPost
     */
    groupId?: number | null;
    /**
     * 
     * @type {string}
     * @memberof ModelPost
     */
    groupUserTitle?: string | null;
    /**
     * 
     * @type {number}
     * @memberof ModelPost
     */
    id?: number | null;
    /**
     * 
     * @type {number}
     * @memberof ModelPost
     */
    inReplyTo?: number | null;
    /**
     * 
     * @type {ModelPost}
     * @memberof ModelPost
     */
    inReplyToPost?: ModelPost | null;
    /**
     * 
     * @type {number}
     * @memberof ModelPost
     */
    inReplyToPostCount?: number | null;
    /**
     * 
     * @type {boolean}
     * @memberof ModelPost
     */
    isEdited?: boolean | null;
    /**
     * 
     * @type {boolean}
     * @memberof ModelPost
     */
    isFailToSend?: boolean | null;
    /**
     * 
     * @type {boolean}
     * @memberof ModelPost
     */
    isGroup?: boolean | null;
    /**
     * 
     * @type {boolean}
     * @memberof ModelPost
     */
    isHighlighted?: boolean | null;
    /**
     * 
     * @type {boolean}
     * @memberof ModelPost
     */
    isHot?: boolean | null;
    /**
     * 
     * @type {boolean}
     * @memberof ModelPost
     */
    isIdFake?: boolean | null;
    /**
     * 
     * @type {boolean}
     * @memberof ModelPost
     */
    isLiked?: boolean | null;
    /**
     * 
     * @type {boolean}
     * @memberof ModelPost
     */
    isMuted?: boolean | null;
    /**
     * 
     * @type {boolean}
     * @memberof ModelPost
     */
    isPinned?: boolean | null;
    /**
     * 
     * @type {boolean}
     * @memberof ModelPost
     */
    isRecommended?: boolean | null;
    /**
     * 
     * @type {boolean}
     * @memberof ModelPost
     */
    isRepostable?: boolean | null;
    /**
     * 
     * @type {boolean}
     * @memberof ModelPost
     */
    isReposted?: boolean | null;
    /**
     * 
     * @type {boolean}
     * @memberof ModelPost
     */
    isRootOfConversation?: boolean | null;
    /**
     * 
     * @type {boolean}
     * @memberof ModelPost
     */
    isRootOfConversationComment?: boolean | null;
    /**
     * 
     * @type {Array<User>}
     * @memberof ModelPost
     */
    likers?: Array<User> | null;
    /**
     * 
     * @type {number}
     * @memberof ModelPost
     */
    likersCount?: number | null;
    /**
     * 
     * @type {number}
     * @memberof ModelPost
     */
    likesCount?: number | null;
    /**
     * 
     * @type {Array<User>}
     * @memberof ModelPost
     */
    mentions?: Array<User> | null;
    /**
     * 
     * @type {Array<ModelMessageTag>}
     * @memberof ModelPost
     */
    messageTags?: Array<ModelMessageTag> | null;
    /**
     * 
     * @type {object}
     * @memberof ModelPost
     */
    postType?: object | null;
    /**
     * 
     * @type {number}
     * @memberof ModelPost
     */
    reportedCount?: number | null;
    /**
     * 
     * @type {number}
     * @memberof ModelPost
     */
    repostsCount?: number | null;
    /**
     * 
     * @type {ModelShareable}
     * @memberof ModelPost
     */
    shareable?: ModelShareable | null;
    /**
     * 
     * @type {ModelThreadInfo}
     * @memberof ModelPost
     */
    sharedThread?: ModelThreadInfo | null;
    /**
     * 
     * @type {ModelSharedUrl}
     * @memberof ModelPost
     */
    sharedUrl?: ModelSharedUrl | null;
    /**
     * 
     * @type {boolean}
     * @memberof ModelPost
     */
    shouldShowOriginalSource?: boolean | null;
    /**
     * 
     * @type {boolean}
     * @memberof ModelPost
     */
    shouldShowSeeMoreConversation?: boolean | null;
    /**
     * 
     * @type {ModelSurvey}
     * @memberof ModelPost
     */
    survey?: ModelSurvey | null;
    /**
     * 
     * @type {string}
     * @memberof ModelPost
     */
    tag?: string | null;
    /**
     * 
     * @type {string}
     * @memberof ModelPost
     */
    text?: string | null;
    /**
     * 
     * @type {ModelThreadInfo}
     * @memberof ModelPost
     */
    thread?: ModelThreadInfo | null;
    /**
     * 
     * @type {number}
     * @memberof ModelPost
     */
    threadId?: number | null;
    /**
     * 
     * @type {number}
     * @memberof ModelPost
     */
    totalGiftsCount?: number | null;
    /**
     * 
     * @type {number}
     * @memberof ModelPost
     */
    updatedAtSeconds?: number | null;
    /**
     * 
     * @type {User}
     * @memberof ModelPost
     */
    user?: User | null;
    /**
     * 
     * @type {Array<ModelVideo>}
     * @memberof ModelPost
     */
    videos?: Array<ModelVideo> | null;
}

/**
 * Check if a given object implements the ModelPost interface.
 */
export function instanceOfModelPost(value: object): value is ModelPost {
    return true;
}

export function ModelPostFromJSON(json: any): ModelPost {
    return ModelPostFromJSONTyped(json, false);
}

export function ModelPostFromJSONTyped(json: any, ignoreDiscriminator: boolean): ModelPost {
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
        'categoryId': json['category_id'] == null ? undefined : json['category_id'],
        'color': json['color'] == null ? undefined : json['color'],
        'conferenceCall': json['conference_call'] == null ? undefined : ConferenceCallFromJSON(json['conference_call']),
        'conversationId': json['conversation_id'] == null ? undefined : json['conversation_id'],
        'createdAtSeconds': json['created_at_seconds'] == null ? undefined : json['created_at_seconds'],
        'editedAtSeconds': json['edited_at_seconds'] == null ? undefined : json['edited_at_seconds'],
        'fontSize': json['font_size'] == null ? undefined : json['font_size'],
        'gameTitle': json['game_title'] == null ? undefined : json['game_title'],
        'gifts': json['gifts'] == null ? undefined : ((json['gifts'] as Array<any>).map(PostGiftFromJSON)),
        'group': json['group'] == null ? undefined : ModelGroupFromJSON(json['group']),
        'groupId': json['group_id'] == null ? undefined : json['group_id'],
        'groupUserTitle': json['group_user_title'] == null ? undefined : json['group_user_title'],
        'id': json['id'] == null ? undefined : json['id'],
        'inReplyTo': json['in_reply_to'] == null ? undefined : json['in_reply_to'],
        'inReplyToPost': json['in_reply_to_post'] == null ? undefined : ModelPostFromJSON(json['in_reply_to_post']),
        'inReplyToPostCount': json['in_reply_to_post_count'] == null ? undefined : json['in_reply_to_post_count'],
        'isEdited': json['is_edited'] == null ? undefined : json['is_edited'],
        'isFailToSend': json['is_fail_to_send'] == null ? undefined : json['is_fail_to_send'],
        'isGroup': json['is_group'] == null ? undefined : json['is_group'],
        'isHighlighted': json['is_highlighted'] == null ? undefined : json['is_highlighted'],
        'isHot': json['is_hot'] == null ? undefined : json['is_hot'],
        'isIdFake': json['is_id_fake'] == null ? undefined : json['is_id_fake'],
        'isLiked': json['is_liked'] == null ? undefined : json['is_liked'],
        'isMuted': json['is_muted'] == null ? undefined : json['is_muted'],
        'isPinned': json['is_pinned'] == null ? undefined : json['is_pinned'],
        'isRecommended': json['is_recommended'] == null ? undefined : json['is_recommended'],
        'isRepostable': json['is_repostable'] == null ? undefined : json['is_repostable'],
        'isReposted': json['is_reposted'] == null ? undefined : json['is_reposted'],
        'isRootOfConversation': json['is_root_of_conversation'] == null ? undefined : json['is_root_of_conversation'],
        'isRootOfConversationComment': json['is_root_of_conversation_comment'] == null ? undefined : json['is_root_of_conversation_comment'],
        'likers': json['likers'] == null ? undefined : ((json['likers'] as Array<any>).map(UserFromJSON)),
        'likersCount': json['likers_count'] == null ? undefined : json['likers_count'],
        'likesCount': json['likes_count'] == null ? undefined : json['likes_count'],
        'mentions': json['mentions'] == null ? undefined : ((json['mentions'] as Array<any>).map(UserFromJSON)),
        'messageTags': json['message_tags'] == null ? undefined : ((json['message_tags'] as Array<any>).map(ModelMessageTagFromJSON)),
        'postType': json['post_type'] == null ? undefined : json['post_type'],
        'reportedCount': json['reported_count'] == null ? undefined : json['reported_count'],
        'repostsCount': json['reposts_count'] == null ? undefined : json['reposts_count'],
        'shareable': json['shareable'] == null ? undefined : ModelShareableFromJSON(json['shareable']),
        'sharedThread': json['shared_thread'] == null ? undefined : ModelThreadInfoFromJSON(json['shared_thread']),
        'sharedUrl': json['shared_url'] == null ? undefined : ModelSharedUrlFromJSON(json['shared_url']),
        'shouldShowOriginalSource': json['should_show_original_source'] == null ? undefined : json['should_show_original_source'],
        'shouldShowSeeMoreConversation': json['should_show_see_more_conversation'] == null ? undefined : json['should_show_see_more_conversation'],
        'survey': json['survey'] == null ? undefined : ModelSurveyFromJSON(json['survey']),
        'tag': json['tag'] == null ? undefined : json['tag'],
        'text': json['text'] == null ? undefined : json['text'],
        'thread': json['thread'] == null ? undefined : ModelThreadInfoFromJSON(json['thread']),
        'threadId': json['thread_id'] == null ? undefined : json['thread_id'],
        'totalGiftsCount': json['total_gifts_count'] == null ? undefined : json['total_gifts_count'],
        'updatedAtSeconds': json['updated_at_seconds'] == null ? undefined : json['updated_at_seconds'],
        'user': json['user'] == null ? undefined : UserFromJSON(json['user']),
        'videos': json['videos'] == null ? undefined : ((json['videos'] as Array<any>).map(ModelVideoFromJSON)),
    };
}

export function ModelPostToJSON(json: any): ModelPost {
    return ModelPostToJSONTyped(json, false);
}

export function ModelPostToJSONTyped(value?: ModelPost | null, ignoreDiscriminator: boolean = false): any {
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
        'category_id': value['categoryId'],
        'color': value['color'],
        'conference_call': ConferenceCallToJSON(value['conferenceCall']),
        'conversation_id': value['conversationId'],
        'created_at_seconds': value['createdAtSeconds'],
        'edited_at_seconds': value['editedAtSeconds'],
        'font_size': value['fontSize'],
        'game_title': value['gameTitle'],
        'gifts': value['gifts'] == null ? undefined : ((value['gifts'] as Array<any>).map(PostGiftToJSON)),
        'group': ModelGroupToJSON(value['group']),
        'group_id': value['groupId'],
        'group_user_title': value['groupUserTitle'],
        'id': value['id'],
        'in_reply_to': value['inReplyTo'],
        'in_reply_to_post': ModelPostToJSON(value['inReplyToPost']),
        'in_reply_to_post_count': value['inReplyToPostCount'],
        'is_edited': value['isEdited'],
        'is_fail_to_send': value['isFailToSend'],
        'is_group': value['isGroup'],
        'is_highlighted': value['isHighlighted'],
        'is_hot': value['isHot'],
        'is_id_fake': value['isIdFake'],
        'is_liked': value['isLiked'],
        'is_muted': value['isMuted'],
        'is_pinned': value['isPinned'],
        'is_recommended': value['isRecommended'],
        'is_repostable': value['isRepostable'],
        'is_reposted': value['isReposted'],
        'is_root_of_conversation': value['isRootOfConversation'],
        'is_root_of_conversation_comment': value['isRootOfConversationComment'],
        'likers': value['likers'] == null ? undefined : ((value['likers'] as Array<any>).map(UserToJSON)),
        'likers_count': value['likersCount'],
        'likes_count': value['likesCount'],
        'mentions': value['mentions'] == null ? undefined : ((value['mentions'] as Array<any>).map(UserToJSON)),
        'message_tags': value['messageTags'] == null ? undefined : ((value['messageTags'] as Array<any>).map(ModelMessageTagToJSON)),
        'post_type': value['postType'],
        'reported_count': value['reportedCount'],
        'reposts_count': value['repostsCount'],
        'shareable': ModelShareableToJSON(value['shareable']),
        'shared_thread': ModelThreadInfoToJSON(value['sharedThread']),
        'shared_url': ModelSharedUrlToJSON(value['sharedUrl']),
        'should_show_original_source': value['shouldShowOriginalSource'],
        'should_show_see_more_conversation': value['shouldShowSeeMoreConversation'],
        'survey': ModelSurveyToJSON(value['survey']),
        'tag': value['tag'],
        'text': value['text'],
        'thread': ModelThreadInfoToJSON(value['thread']),
        'thread_id': value['threadId'],
        'total_gifts_count': value['totalGiftsCount'],
        'updated_at_seconds': value['updatedAtSeconds'],
        'user': UserToJSON(value['user']),
        'videos': value['videos'] == null ? undefined : ((value['videos'] as Array<any>).map(ModelVideoToJSON)),
    };
}

