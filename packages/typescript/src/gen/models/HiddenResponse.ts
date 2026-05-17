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
import type { RealmUser } from './RealmUser';
import {
    RealmUserFromJSON,
    RealmUserFromJSONTyped,
    RealmUserToJSON,
    RealmUserToJSONTyped,
} from './RealmUser';

/**
 * 
 * @export
 * @interface HiddenResponse
 */
export interface HiddenResponse {
    /**
     * 
     * @type {Array<RealmUser>}
     * @memberof HiddenResponse
     */
    hiddenUsers?: Array<RealmUser> | null;
    /**
     * 
     * @type {number}
     * @memberof HiddenResponse
     */
    limit?: number | null;
    /**
     * 
     * @type {string}
     * @memberof HiddenResponse
     */
    nextPageValue?: string | null;
    /**
     * 
     * @type {number}
     * @memberof HiddenResponse
     */
    totalCount?: number | null;
}

/**
 * Check if a given object implements the HiddenResponse interface.
 */
export function instanceOfHiddenResponse(value: object): value is HiddenResponse {
    return true;
}

export function HiddenResponseFromJSON(json: any): HiddenResponse {
    return HiddenResponseFromJSONTyped(json, false);
}

export function HiddenResponseFromJSONTyped(json: any, ignoreDiscriminator: boolean): HiddenResponse {
    if (json == null) {
        return json;
    }
    return {
        
        'hiddenUsers': json['hidden_users'] == null ? undefined : ((json['hidden_users'] as Array<any>).map(RealmUserFromJSON)),
        'limit': json['limit'] == null ? undefined : json['limit'],
        'nextPageValue': json['next_page_value'] == null ? undefined : String(json['next_page_value']),
        'totalCount': json['total_count'] == null ? undefined : json['total_count'],
    };
}

export function HiddenResponseToJSON(json: any): HiddenResponse {
    return HiddenResponseToJSONTyped(json, false);
}

export function HiddenResponseToJSONTyped(value?: HiddenResponse | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'hidden_users': value['hiddenUsers'] == null ? undefined : ((value['hiddenUsers'] as Array<any>).map(RealmUserToJSON)),
        'limit': value['limit'],
        'next_page_value': value['nextPageValue'],
        'total_count': value['totalCount'],
    };
}

