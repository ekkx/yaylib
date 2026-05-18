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
 * @interface FollowRequestCountResponse
 */
export interface FollowRequestCountResponse {
    /**
     * 
     * @type {number}
     * @memberof FollowRequestCountResponse
     */
    usersCount?: number | null;
}

/**
 * Check if a given object implements the FollowRequestCountResponse interface.
 */
export function instanceOfFollowRequestCountResponse(value: object): value is FollowRequestCountResponse {
    return true;
}

export function FollowRequestCountResponseFromJSON(json: any): FollowRequestCountResponse {
    return FollowRequestCountResponseFromJSONTyped(json, false);
}

export function FollowRequestCountResponseFromJSONTyped(json: any, ignoreDiscriminator: boolean): FollowRequestCountResponse {
    if (json == null) {
        return json;
    }
    return {
        
        'usersCount': json['users_count'] == null ? undefined : json['users_count'],
    };
}

export function FollowRequestCountResponseToJSON(json: any): FollowRequestCountResponse {
    return FollowRequestCountResponseToJSONTyped(json, false);
}

export function FollowRequestCountResponseToJSONTyped(value?: FollowRequestCountResponse | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'users_count': value['usersCount'],
    };
}

