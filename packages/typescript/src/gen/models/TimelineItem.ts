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
import type { ModelPost } from './ModelPost';
import {
    ModelPostFromJSON,
    ModelPostFromJSONTyped,
    ModelPostToJSON,
    ModelPostToJSONTyped,
} from './ModelPost';
import type { GroupCategory } from './GroupCategory';
import {
    GroupCategoryFromJSON,
    GroupCategoryFromJSONTyped,
    GroupCategoryToJSON,
    GroupCategoryToJSONTyped,
} from './GroupCategory';

/**
 * 
 * @export
 * @interface TimelineItem
 */
export interface TimelineItem {
    /**
     * 
     * @type {Array<GroupCategory>}
     * @memberof TimelineItem
     */
    categories?: Array<GroupCategory> | null;
    /**
     * 
     * @type {boolean}
     * @memberof TimelineItem
     */
    isRepliedByNextPost?: boolean | null;
    /**
     * 
     * @type {boolean}
     * @memberof TimelineItem
     */
    isReplyingToPrevPost?: boolean | null;
    /**
     * 
     * @type {ModelPost}
     * @memberof TimelineItem
     */
    post?: ModelPost | null;
    /**
     * 
     * @type {boolean}
     * @memberof TimelineItem
     */
    shouldShowConversationLine?: boolean | null;
    /**
     * 
     * @type {boolean}
     * @memberof TimelineItem
     */
    shouldShowMentionBox?: boolean | null;
    /**
     * 
     * @type {boolean}
     * @memberof TimelineItem
     */
    shouldShowRepostButton?: boolean | null;
    /**
     * 
     * @type {boolean}
     * @memberof TimelineItem
     */
    shouldShowTranslateButton?: boolean | null;
    /**
     * 
     * @type {boolean}
     * @memberof TimelineItem
     */
    shouldShowViewReposts?: boolean | null;
    /**
     * 
     * @type {string}
     * @memberof TimelineItem
     */
    translatedText?: string | null;
}

/**
 * Check if a given object implements the TimelineItem interface.
 */
export function instanceOfTimelineItem(value: object): value is TimelineItem {
    return true;
}

export function TimelineItemFromJSON(json: any): TimelineItem {
    return TimelineItemFromJSONTyped(json, false);
}

export function TimelineItemFromJSONTyped(json: any, ignoreDiscriminator: boolean): TimelineItem {
    if (json == null) {
        return json;
    }
    return {
        
        'categories': json['categories'] == null ? undefined : ((json['categories'] as Array<any>).map(GroupCategoryFromJSON)),
        'isRepliedByNextPost': json['is_replied_by_next_post'] == null ? undefined : json['is_replied_by_next_post'],
        'isReplyingToPrevPost': json['is_replying_to_prev_post'] == null ? undefined : json['is_replying_to_prev_post'],
        'post': json['post'] == null ? undefined : ModelPostFromJSON(json['post']),
        'shouldShowConversationLine': json['should_show_conversation_line'] == null ? undefined : json['should_show_conversation_line'],
        'shouldShowMentionBox': json['should_show_mention_box'] == null ? undefined : json['should_show_mention_box'],
        'shouldShowRepostButton': json['should_show_repost_button'] == null ? undefined : json['should_show_repost_button'],
        'shouldShowTranslateButton': json['should_show_translate_button'] == null ? undefined : json['should_show_translate_button'],
        'shouldShowViewReposts': json['should_show_view_reposts'] == null ? undefined : json['should_show_view_reposts'],
        'translatedText': json['translated_text'] == null ? undefined : json['translated_text'],
    };
}

export function TimelineItemToJSON(json: any): TimelineItem {
    return TimelineItemToJSONTyped(json, false);
}

export function TimelineItemToJSONTyped(value?: TimelineItem | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'categories': value['categories'] == null ? undefined : ((value['categories'] as Array<any>).map(GroupCategoryToJSON)),
        'is_replied_by_next_post': value['isRepliedByNextPost'],
        'is_replying_to_prev_post': value['isReplyingToPrevPost'],
        'post': ModelPostToJSON(value['post']),
        'should_show_conversation_line': value['shouldShowConversationLine'],
        'should_show_mention_box': value['shouldShowMentionBox'],
        'should_show_repost_button': value['shouldShowRepostButton'],
        'should_show_translate_button': value['shouldShowTranslateButton'],
        'should_show_view_reposts': value['shouldShowViewReposts'],
        'translated_text': value['translatedText'],
    };
}

