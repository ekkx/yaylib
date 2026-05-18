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
 * @interface UserCustomDefinitionsResponse
 */
export interface UserCustomDefinitionsResponse {
    /**
     * 
     * @type {number}
     * @memberof UserCustomDefinitionsResponse
     */
    age?: number | null;
    /**
     * 
     * @type {number}
     * @memberof UserCustomDefinitionsResponse
     */
    createdAt?: number | null;
    /**
     * 
     * @type {number}
     * @memberof UserCustomDefinitionsResponse
     */
    followersCount?: number | null;
    /**
     * 
     * @type {number}
     * @memberof UserCustomDefinitionsResponse
     */
    followingsCount?: number | null;
    /**
     * 
     * @type {number}
     * @memberof UserCustomDefinitionsResponse
     */
    lastLoggedinAt?: number | null;
    /**
     * 
     * @type {number}
     * @memberof UserCustomDefinitionsResponse
     */
    reportedCount?: number | null;
    /**
     * 
     * @type {string}
     * @memberof UserCustomDefinitionsResponse
     */
    status?: string | null;
}

/**
 * Check if a given object implements the UserCustomDefinitionsResponse interface.
 */
export function instanceOfUserCustomDefinitionsResponse(value: object): value is UserCustomDefinitionsResponse {
    return true;
}

export function UserCustomDefinitionsResponseFromJSON(json: any): UserCustomDefinitionsResponse {
    return UserCustomDefinitionsResponseFromJSONTyped(json, false);
}

export function UserCustomDefinitionsResponseFromJSONTyped(json: any, ignoreDiscriminator: boolean): UserCustomDefinitionsResponse {
    if (json == null) {
        return json;
    }
    return {
        
        'age': json['age'] == null ? undefined : json['age'],
        'createdAt': json['created_at'] == null ? undefined : json['created_at'],
        'followersCount': json['followers_count'] == null ? undefined : json['followers_count'],
        'followingsCount': json['followings_count'] == null ? undefined : json['followings_count'],
        'lastLoggedinAt': json['last_loggedin_at'] == null ? undefined : json['last_loggedin_at'],
        'reportedCount': json['reported_count'] == null ? undefined : json['reported_count'],
        'status': json['status'] == null ? undefined : json['status'],
    };
}

export function UserCustomDefinitionsResponseToJSON(json: any): UserCustomDefinitionsResponse {
    return UserCustomDefinitionsResponseToJSONTyped(json, false);
}

export function UserCustomDefinitionsResponseToJSONTyped(value?: UserCustomDefinitionsResponse | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'age': value['age'],
        'created_at': value['createdAt'],
        'followers_count': value['followersCount'],
        'followings_count': value['followingsCount'],
        'last_loggedin_at': value['lastLoggedinAt'],
        'reported_count': value['reportedCount'],
        'status': value['status'],
    };
}

