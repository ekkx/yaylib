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
import type { GifImage } from './GifImage';
import {
    GifImageFromJSON,
    GifImageFromJSONTyped,
    GifImageToJSON,
    GifImageToJSONTyped,
} from './GifImage';
import type { Sticker } from './Sticker';
import {
    StickerFromJSON,
    StickerFromJSONTyped,
    StickerToJSON,
    StickerToJSONTyped,
} from './Sticker';
import type { RealmConferenceCall } from './RealmConferenceCall';
import {
    RealmConferenceCallFromJSON,
    RealmConferenceCallFromJSONTyped,
    RealmConferenceCallToJSON,
    RealmConferenceCallToJSONTyped,
} from './RealmConferenceCall';
import type { ChatInvitation } from './ChatInvitation';
import {
    ChatInvitationFromJSON,
    ChatInvitationFromJSONTyped,
    ChatInvitationToJSON,
    ChatInvitationToJSONTyped,
} from './ChatInvitation';
import type { ParentMessage } from './ParentMessage';
import {
    ParentMessageFromJSON,
    ParentMessageFromJSONTyped,
    ParentMessageToJSON,
    ParentMessageToJSONTyped,
} from './ParentMessage';
import type { MessageType } from './MessageType';
import {
    MessageTypeFromJSON,
    MessageTypeFromJSONTyped,
    MessageTypeToJSON,
    MessageTypeToJSONTyped,
} from './MessageType';

/**
 * 
 * @export
 * @interface RealmMessage
 */
export interface RealmMessage {
    /**
     * 
     * @type {string}
     * @memberof RealmMessage
     */
    attachment?: string | null;
    /**
     * 
     * @type {string}
     * @memberof RealmMessage
     */
    attachmentAndroid?: string | null;
    /**
     * 
     * @type {number}
     * @memberof RealmMessage
     */
    attachmentReadCount?: number | null;
    /**
     * 
     * @type {string}
     * @memberof RealmMessage
     */
    attachmentThumbnail?: string | null;
    /**
     * 
     * @type {RealmConferenceCall}
     * @memberof RealmMessage
     */
    conferenceCall?: RealmConferenceCall | null;
    /**
     * 
     * @type {number}
     * @memberof RealmMessage
     */
    createdAt?: number | null;
    /**
     * 
     * @type {number}
     * @memberof RealmMessage
     */
    fontSize?: number | null;
    /**
     * 
     * @type {GifImage}
     * @memberof RealmMessage
     */
    gif?: GifImage | null;
    /**
     * 
     * @type {number}
     * @memberof RealmMessage
     */
    id?: number | null;
    /**
     * 
     * @type {ChatInvitation}
     * @memberof RealmMessage
     */
    invitation?: ChatInvitation | null;
    /**
     * 
     * @type {boolean}
     * @memberof RealmMessage
     */
    isError?: boolean | null;
    /**
     * 
     * @type {boolean}
     * @memberof RealmMessage
     */
    isSent?: boolean | null;
    /**
     * 
     * @type {MessageType}
     * @memberof RealmMessage
     */
    messageType?: MessageType | null;
    /**
     * 
     * @type {ParentMessage}
     * @memberof RealmMessage
     */
    parent?: ParentMessage | null;
    /**
     * 
     * @type {number}
     * @memberof RealmMessage
     */
    refreshRetryCount?: number | null;
    /**
     * 
     * @type {number}
     * @memberof RealmMessage
     */
    roomId?: number | null;
    /**
     * 
     * @type {Sticker}
     * @memberof RealmMessage
     */
    sticker?: Sticker | null;
    /**
     * 
     * @type {string}
     * @memberof RealmMessage
     */
    text?: string | null;
    /**
     * 
     * @type {number}
     * @memberof RealmMessage
     */
    userId?: number | null;
    /**
     * 
     * @type {boolean}
     * @memberof RealmMessage
     */
    videoProcessed?: boolean | null;
    /**
     * 
     * @type {string}
     * @memberof RealmMessage
     */
    videoThumbnailUrl?: string | null;
    /**
     * 
     * @type {string}
     * @memberof RealmMessage
     */
    videoUrl?: string | null;
}



/**
 * Check if a given object implements the RealmMessage interface.
 */
export function instanceOfRealmMessage(value: object): value is RealmMessage {
    return true;
}

export function RealmMessageFromJSON(json: any): RealmMessage {
    return RealmMessageFromJSONTyped(json, false);
}

export function RealmMessageFromJSONTyped(json: any, ignoreDiscriminator: boolean): RealmMessage {
    if (json == null) {
        return json;
    }
    return {
        
        'attachment': json['attachment'] == null ? undefined : json['attachment'],
        'attachmentAndroid': json['attachment_android'] == null ? undefined : json['attachment_android'],
        'attachmentReadCount': json['attachment_read_count'] == null ? undefined : json['attachment_read_count'],
        'attachmentThumbnail': json['attachment_thumbnail'] == null ? undefined : json['attachment_thumbnail'],
        'conferenceCall': json['conference_call'] == null ? undefined : RealmConferenceCallFromJSON(json['conference_call']),
        'createdAt': json['created_at'] == null ? undefined : json['created_at'],
        'fontSize': json['font_size'] == null ? undefined : json['font_size'],
        'gif': json['gif'] == null ? undefined : GifImageFromJSON(json['gif']),
        'id': json['id'] == null ? undefined : json['id'],
        'invitation': json['invitation'] == null ? undefined : ChatInvitationFromJSON(json['invitation']),
        'isError': json['is_error'] == null ? undefined : json['is_error'],
        'isSent': json['is_sent'] == null ? undefined : json['is_sent'],
        'messageType': json['message_type'] == null ? undefined : MessageTypeFromJSON(json['message_type']),
        'parent': json['parent'] == null ? undefined : ParentMessageFromJSON(json['parent']),
        'refreshRetryCount': json['refresh_retry_count'] == null ? undefined : json['refresh_retry_count'],
        'roomId': json['room_id'] == null ? undefined : json['room_id'],
        'sticker': json['sticker'] == null ? undefined : StickerFromJSON(json['sticker']),
        'text': json['text'] == null ? undefined : json['text'],
        'userId': json['user_id'] == null ? undefined : json['user_id'],
        'videoProcessed': json['video_processed'] == null ? undefined : json['video_processed'],
        'videoThumbnailUrl': json['video_thumbnail_url'] == null ? undefined : json['video_thumbnail_url'],
        'videoUrl': json['video_url'] == null ? undefined : json['video_url'],
    };
}

export function RealmMessageToJSON(json: any): RealmMessage {
    return RealmMessageToJSONTyped(json, false);
}

export function RealmMessageToJSONTyped(value?: RealmMessage | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'attachment': value['attachment'],
        'attachment_android': value['attachmentAndroid'],
        'attachment_read_count': value['attachmentReadCount'],
        'attachment_thumbnail': value['attachmentThumbnail'],
        'conference_call': RealmConferenceCallToJSON(value['conferenceCall']),
        'created_at': value['createdAt'],
        'font_size': value['fontSize'],
        'gif': GifImageToJSON(value['gif']),
        'id': value['id'],
        'invitation': ChatInvitationToJSON(value['invitation']),
        'is_error': value['isError'],
        'is_sent': value['isSent'],
        'message_type': MessageTypeToJSON(value['messageType']),
        'parent': ParentMessageToJSON(value['parent']),
        'refresh_retry_count': value['refreshRetryCount'],
        'room_id': value['roomId'],
        'sticker': StickerToJSON(value['sticker']),
        'text': value['text'],
        'user_id': value['userId'],
        'video_processed': value['videoProcessed'],
        'video_thumbnail_url': value['videoThumbnailUrl'],
        'video_url': value['videoUrl'],
    };
}

