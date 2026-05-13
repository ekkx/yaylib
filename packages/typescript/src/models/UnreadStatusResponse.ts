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
 * @interface UnreadStatusResponse
 */
export interface UnreadStatusResponse {
    /**
     * 
     * @type {boolean}
     * @memberof UnreadStatusResponse
     */
    isUnread?: boolean | null;
}

/**
 * Check if a given object implements the UnreadStatusResponse interface.
 */
export function instanceOfUnreadStatusResponse(value: object): value is UnreadStatusResponse {
    return true;
}

export function UnreadStatusResponseFromJSON(json: any): UnreadStatusResponse {
    return UnreadStatusResponseFromJSONTyped(json, false);
}

export function UnreadStatusResponseFromJSONTyped(json: any, ignoreDiscriminator: boolean): UnreadStatusResponse {
    if (json == null) {
        return json;
    }
    return {
        
        'isUnread': json['is_unread'] == null ? undefined : json['is_unread'],
    };
}

export function UnreadStatusResponseToJSON(json: any): UnreadStatusResponse {
    return UnreadStatusResponseToJSONTyped(json, false);
}

export function UnreadStatusResponseToJSONTyped(value?: UnreadStatusResponse | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'is_unread': value['isUnread'],
    };
}

