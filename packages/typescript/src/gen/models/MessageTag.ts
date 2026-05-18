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
 * @interface MessageTag
 */
export interface MessageTag {
    /**
     * 
     * @type {number}
     * @memberof MessageTag
     */
    length?: number | null;
    /**
     * 
     * @type {number}
     * @memberof MessageTag
     */
    offset?: number | null;
    /**
     * 
     * @type {string}
     * @memberof MessageTag
     */
    type?: string | null;
    /**
     * 
     * @type {number}
     * @memberof MessageTag
     */
    userId?: number | null;
}

/**
 * Check if a given object implements the MessageTag interface.
 */
export function instanceOfMessageTag(value: object): value is MessageTag {
    return true;
}

export function MessageTagFromJSON(json: any): MessageTag {
    return MessageTagFromJSONTyped(json, false);
}

export function MessageTagFromJSONTyped(json: any, ignoreDiscriminator: boolean): MessageTag {
    if (json == null) {
        return json;
    }
    return {
        
        'length': json['length'] == null ? undefined : json['length'],
        'offset': json['offset'] == null ? undefined : json['offset'],
        'type': json['type'] == null ? undefined : json['type'],
        'userId': json['user_id'] == null ? undefined : json['user_id'],
    };
}

export function MessageTagToJSON(json: any): MessageTag {
    return MessageTagToJSONTyped(json, false);
}

export function MessageTagToJSONTyped(value?: MessageTag | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'length': value['length'],
        'offset': value['offset'],
        'type': value['type'],
        'user_id': value['userId'],
    };
}

