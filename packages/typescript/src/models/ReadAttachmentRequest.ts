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
 * @interface ReadAttachmentRequest
 */
export interface ReadAttachmentRequest {
    /**
     * 
     * @type {Array<number>}
     * @memberof ReadAttachmentRequest
     */
    attachmentMsgIds?: Array<number> | null;
}

/**
 * Check if a given object implements the ReadAttachmentRequest interface.
 */
export function instanceOfReadAttachmentRequest(value: object): value is ReadAttachmentRequest {
    return true;
}

export function ReadAttachmentRequestFromJSON(json: any): ReadAttachmentRequest {
    return ReadAttachmentRequestFromJSONTyped(json, false);
}

export function ReadAttachmentRequestFromJSONTyped(json: any, ignoreDiscriminator: boolean): ReadAttachmentRequest {
    if (json == null) {
        return json;
    }
    return {
        
        'attachmentMsgIds': json['attachment_msg_ids'] == null ? undefined : json['attachment_msg_ids'],
    };
}

export function ReadAttachmentRequestToJSON(json: any): ReadAttachmentRequest {
    return ReadAttachmentRequestToJSONTyped(json, false);
}

export function ReadAttachmentRequestToJSONTyped(value?: ReadAttachmentRequest | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'attachment_msg_ids': value['attachmentMsgIds'],
    };
}

