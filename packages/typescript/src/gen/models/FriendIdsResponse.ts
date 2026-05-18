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
 * @interface FriendIdsResponse
 */
export interface FriendIdsResponse {
    /**
     * 
     * @type {Array<string>}
     * @memberof FriendIdsResponse
     */
    ids?: Array<string> | null;
}

/**
 * Check if a given object implements the FriendIdsResponse interface.
 */
export function instanceOfFriendIdsResponse(value: object): value is FriendIdsResponse {
    return true;
}

export function FriendIdsResponseFromJSON(json: any): FriendIdsResponse {
    return FriendIdsResponseFromJSONTyped(json, false);
}

export function FriendIdsResponseFromJSONTyped(json: any, ignoreDiscriminator: boolean): FriendIdsResponse {
    if (json == null) {
        return json;
    }
    return {
        
        'ids': json['ids'] == null ? undefined : json['ids'],
    };
}

export function FriendIdsResponseToJSON(json: any): FriendIdsResponse {
    return FriendIdsResponseToJSONTyped(json, false);
}

export function FriendIdsResponseToJSONTyped(value?: FriendIdsResponse | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'ids': value['ids'],
    };
}

