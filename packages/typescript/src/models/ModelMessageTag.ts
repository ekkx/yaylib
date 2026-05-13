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
import type { MessageTagType } from './MessageTagType';
import {
    MessageTagTypeFromJSON,
    MessageTagTypeFromJSONTyped,
    MessageTagTypeToJSON,
    MessageTagTypeToJSONTyped,
} from './MessageTagType';

/**
 * 
 * @export
 * @interface ModelMessageTag
 */
export interface ModelMessageTag {
    /**
     * 
     * @type {number}
     * @memberof ModelMessageTag
     */
    endAt?: number | null;
    /**
     * 
     * @type {number}
     * @memberof ModelMessageTag
     */
    length?: number | null;
    /**
     * 
     * @type {number}
     * @memberof ModelMessageTag
     */
    offset?: number | null;
    /**
     * 
     * @type {MessageTagType}
     * @memberof ModelMessageTag
     */
    type?: MessageTagType | null;
    /**
     * 
     * @type {number}
     * @memberof ModelMessageTag
     */
    userId?: number | null;
}

/**
 * Check if a given object implements the ModelMessageTag interface.
 */
export function instanceOfModelMessageTag(value: object): value is ModelMessageTag {
    return true;
}

export function ModelMessageTagFromJSON(json: any): ModelMessageTag {
    return ModelMessageTagFromJSONTyped(json, false);
}

export function ModelMessageTagFromJSONTyped(json: any, ignoreDiscriminator: boolean): ModelMessageTag {
    if (json == null) {
        return json;
    }
    return {
        
        'endAt': json['end_at'] == null ? undefined : json['end_at'],
        'length': json['length'] == null ? undefined : json['length'],
        'offset': json['offset'] == null ? undefined : json['offset'],
        'type': json['type'] == null ? undefined : MessageTagTypeFromJSON(json['type']),
        'userId': json['user_id'] == null ? undefined : json['user_id'],
    };
}

export function ModelMessageTagToJSON(json: any): ModelMessageTag {
    return ModelMessageTagToJSONTyped(json, false);
}

export function ModelMessageTagToJSONTyped(value?: ModelMessageTag | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'end_at': value['endAt'],
        'length': value['length'],
        'offset': value['offset'],
        'type': MessageTagTypeToJSON(value['type']),
        'user_id': value['userId'],
    };
}

