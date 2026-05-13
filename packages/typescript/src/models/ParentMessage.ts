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
 * @interface ParentMessage
 */
export interface ParentMessage {
    /**
     * 
     * @type {string}
     * @memberof ParentMessage
     */
    attachment?: string | null;
    /**
     * 
     * @type {string}
     * @memberof ParentMessage
     */
    attachmentThumbnail?: string | null;
    /**
     * 
     * @type {number}
     * @memberof ParentMessage
     */
    createdAt?: number | null;
    /**
     * 
     * @type {number}
     * @memberof ParentMessage
     */
    fontSize?: number | null;
    /**
     * 
     * @type {GifImage}
     * @memberof ParentMessage
     */
    gif?: GifImage | null;
    /**
     * 
     * @type {number}
     * @memberof ParentMessage
     */
    id?: number | null;
    /**
     * 
     * @type {MessageType}
     * @memberof ParentMessage
     */
    messageType?: MessageType | null;
    /**
     * 
     * @type {number}
     * @memberof ParentMessage
     */
    roomId?: number | null;
    /**
     * 
     * @type {Sticker}
     * @memberof ParentMessage
     */
    sticker?: Sticker | null;
    /**
     * 
     * @type {string}
     * @memberof ParentMessage
     */
    text?: string | null;
    /**
     * 
     * @type {number}
     * @memberof ParentMessage
     */
    userId?: number | null;
    /**
     * 
     * @type {string}
     * @memberof ParentMessage
     */
    videoThumbnailUrl?: string | null;
    /**
     * 
     * @type {string}
     * @memberof ParentMessage
     */
    videoUrl?: string | null;
}



/**
 * Check if a given object implements the ParentMessage interface.
 */
export function instanceOfParentMessage(value: object): value is ParentMessage {
    return true;
}

export function ParentMessageFromJSON(json: any): ParentMessage {
    return ParentMessageFromJSONTyped(json, false);
}

export function ParentMessageFromJSONTyped(json: any, ignoreDiscriminator: boolean): ParentMessage {
    if (json == null) {
        return json;
    }
    return {
        
        'attachment': json['attachment'] == null ? undefined : json['attachment'],
        'attachmentThumbnail': json['attachment_thumbnail'] == null ? undefined : json['attachment_thumbnail'],
        'createdAt': json['created_at'] == null ? undefined : json['created_at'],
        'fontSize': json['font_size'] == null ? undefined : json['font_size'],
        'gif': json['gif'] == null ? undefined : GifImageFromJSON(json['gif']),
        'id': json['id'] == null ? undefined : json['id'],
        'messageType': json['message_type'] == null ? undefined : MessageTypeFromJSON(json['message_type']),
        'roomId': json['room_id'] == null ? undefined : json['room_id'],
        'sticker': json['sticker'] == null ? undefined : StickerFromJSON(json['sticker']),
        'text': json['text'] == null ? undefined : json['text'],
        'userId': json['user_id'] == null ? undefined : json['user_id'],
        'videoThumbnailUrl': json['video_thumbnail_url'] == null ? undefined : json['video_thumbnail_url'],
        'videoUrl': json['video_url'] == null ? undefined : json['video_url'],
    };
}

export function ParentMessageToJSON(json: any): ParentMessage {
    return ParentMessageToJSONTyped(json, false);
}

export function ParentMessageToJSONTyped(value?: ParentMessage | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'attachment': value['attachment'],
        'attachment_thumbnail': value['attachmentThumbnail'],
        'created_at': value['createdAt'],
        'font_size': value['fontSize'],
        'gif': GifImageToJSON(value['gif']),
        'id': value['id'],
        'message_type': MessageTypeToJSON(value['messageType']),
        'room_id': value['roomId'],
        'sticker': StickerToJSON(value['sticker']),
        'text': value['text'],
        'user_id': value['userId'],
        'video_thumbnail_url': value['videoThumbnailUrl'],
        'video_url': value['videoUrl'],
    };
}

