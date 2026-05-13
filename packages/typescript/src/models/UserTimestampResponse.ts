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
 * @interface UserTimestampResponse
 */
export interface UserTimestampResponse {
    /**
     * 
     * @type {string}
     * @memberof UserTimestampResponse
     */
    country?: string | null;
    /**
     * 
     * @type {string}
     * @memberof UserTimestampResponse
     */
    ipAddress?: string | null;
    /**
     * 
     * @type {number}
     * @memberof UserTimestampResponse
     */
    time?: number | null;
}

/**
 * Check if a given object implements the UserTimestampResponse interface.
 */
export function instanceOfUserTimestampResponse(value: object): value is UserTimestampResponse {
    return true;
}

export function UserTimestampResponseFromJSON(json: any): UserTimestampResponse {
    return UserTimestampResponseFromJSONTyped(json, false);
}

export function UserTimestampResponseFromJSONTyped(json: any, ignoreDiscriminator: boolean): UserTimestampResponse {
    if (json == null) {
        return json;
    }
    return {
        
        'country': json['country'] == null ? undefined : json['country'],
        'ipAddress': json['ip_address'] == null ? undefined : json['ip_address'],
        'time': json['time'] == null ? undefined : json['time'],
    };
}

export function UserTimestampResponseToJSON(json: any): UserTimestampResponse {
    return UserTimestampResponseToJSONTyped(json, false);
}

export function UserTimestampResponseToJSONTyped(value?: UserTimestampResponse | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'country': value['country'],
        'ip_address': value['ipAddress'],
        'time': value['time'],
    };
}

