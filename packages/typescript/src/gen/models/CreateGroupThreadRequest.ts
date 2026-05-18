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
/**
 * 
 * @export
 * @interface CreateGroupThreadRequest
 */
export interface CreateGroupThreadRequest {
    /**
     * 
     * @type {number}
     * @memberof CreateGroupThreadRequest
     */
    groupId?: number | null;
    /**
     * 
     * @type {string}
     * @memberof CreateGroupThreadRequest
     */
    threadIconFilename?: string | null;
    /**
     * 
     * @type {string}
     * @memberof CreateGroupThreadRequest
     */
    title?: string | null;
}

/**
 * Check if a given object implements the CreateGroupThreadRequest interface.
 */
export function instanceOfCreateGroupThreadRequest(value: object): value is CreateGroupThreadRequest {
    return true;
}

export function CreateGroupThreadRequestFromJSON(json: any): CreateGroupThreadRequest {
    return CreateGroupThreadRequestFromJSONTyped(json, false);
}

export function CreateGroupThreadRequestFromJSONTyped(json: any, ignoreDiscriminator: boolean): CreateGroupThreadRequest {
    if (json == null) {
        return json;
    }
    return {
        
        'groupId': json['group_id'] == null ? undefined : json['group_id'],
        'threadIconFilename': json['thread_icon_filename'] == null ? undefined : json['thread_icon_filename'],
        'title': json['title'] == null ? undefined : json['title'],
    };
}

export function CreateGroupThreadRequestToJSON(json: any): CreateGroupThreadRequest {
    return CreateGroupThreadRequestToJSONTyped(json, false);
}

export function CreateGroupThreadRequestToJSONTyped(value?: CreateGroupThreadRequest | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'group_id': value['groupId'],
        'thread_icon_filename': value['threadIconFilename'],
        'title': value['title'],
    };
}

