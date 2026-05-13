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
 * @interface BlockedUserIdsResponse
 */
export interface BlockedUserIdsResponse {
    /**
     * 
     * @type {Array<number>}
     * @memberof BlockedUserIdsResponse
     */
    blockIds?: Array<number> | null;
}

/**
 * Check if a given object implements the BlockedUserIdsResponse interface.
 */
export function instanceOfBlockedUserIdsResponse(value: object): value is BlockedUserIdsResponse {
    return true;
}

export function BlockedUserIdsResponseFromJSON(json: any): BlockedUserIdsResponse {
    return BlockedUserIdsResponseFromJSONTyped(json, false);
}

export function BlockedUserIdsResponseFromJSONTyped(json: any, ignoreDiscriminator: boolean): BlockedUserIdsResponse {
    if (json == null) {
        return json;
    }
    return {
        
        'blockIds': json['block_ids'] == null ? undefined : json['block_ids'],
    };
}

export function BlockedUserIdsResponseToJSON(json: any): BlockedUserIdsResponse {
    return BlockedUserIdsResponseToJSONTyped(json, false);
}

export function BlockedUserIdsResponseToJSONTyped(value?: BlockedUserIdsResponse | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'block_ids': value['blockIds'],
    };
}

