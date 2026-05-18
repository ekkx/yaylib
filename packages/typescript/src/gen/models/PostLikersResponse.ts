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
import type { RealmUser } from './RealmUser.js';
import {
    RealmUserFromJSON,
    RealmUserFromJSONTyped,
    RealmUserToJSON,
    RealmUserToJSONTyped,
} from './RealmUser.js';

/**
 * 
 * @export
 * @interface PostLikersResponse
 */
export interface PostLikersResponse {
    /**
     * 
     * @type {number}
     * @memberof PostLikersResponse
     */
    lastId?: number | null;
    /**
     * 
     * @type {Array<RealmUser>}
     * @memberof PostLikersResponse
     */
    users?: Array<RealmUser> | null;
}

/**
 * Check if a given object implements the PostLikersResponse interface.
 */
export function instanceOfPostLikersResponse(value: object): value is PostLikersResponse {
    return true;
}

export function PostLikersResponseFromJSON(json: any): PostLikersResponse {
    return PostLikersResponseFromJSONTyped(json, false);
}

export function PostLikersResponseFromJSONTyped(json: any, ignoreDiscriminator: boolean): PostLikersResponse {
    if (json == null) {
        return json;
    }
    return {
        
        'lastId': json['last_id'] == null ? undefined : json['last_id'],
        'users': json['users'] == null ? undefined : ((json['users'] as Array<any>).map(RealmUserFromJSON)),
    };
}

export function PostLikersResponseToJSON(json: any): PostLikersResponse {
    return PostLikersResponseToJSONTyped(json, false);
}

export function PostLikersResponseToJSONTyped(value?: PostLikersResponse | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'last_id': value['lastId'],
        'users': value['users'] == null ? undefined : ((value['users'] as Array<any>).map(RealmUserToJSON)),
    };
}

