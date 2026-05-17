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
import type { ThreadInfo } from './ThreadInfo';
import {
    ThreadInfoFromJSON,
    ThreadInfoFromJSONTyped,
    ThreadInfoToJSON,
    ThreadInfoToJSONTyped,
} from './ThreadInfo';

/**
 * 
 * @export
 * @interface GroupThreadListResponse
 */
export interface GroupThreadListResponse {
    /**
     * 
     * @type {string}
     * @memberof GroupThreadListResponse
     */
    nextPageValue?: string | null;
    /**
     * 
     * @type {Array<ThreadInfo>}
     * @memberof GroupThreadListResponse
     */
    threads?: Array<ThreadInfo> | null;
}

/**
 * Check if a given object implements the GroupThreadListResponse interface.
 */
export function instanceOfGroupThreadListResponse(value: object): value is GroupThreadListResponse {
    return true;
}

export function GroupThreadListResponseFromJSON(json: any): GroupThreadListResponse {
    return GroupThreadListResponseFromJSONTyped(json, false);
}

export function GroupThreadListResponseFromJSONTyped(json: any, ignoreDiscriminator: boolean): GroupThreadListResponse {
    if (json == null) {
        return json;
    }
    return {
        
        'nextPageValue': json['next_page_value'] == null ? undefined : String(json['next_page_value']),
        'threads': json['threads'] == null ? undefined : ((json['threads'] as Array<any>).map(ThreadInfoFromJSON)),
    };
}

export function GroupThreadListResponseToJSON(json: any): GroupThreadListResponse {
    return GroupThreadListResponseToJSONTyped(json, false);
}

export function GroupThreadListResponseToJSONTyped(value?: GroupThreadListResponse | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'next_page_value': value['nextPageValue'],
        'threads': value['threads'] == null ? undefined : ((value['threads'] as Array<any>).map(ThreadInfoToJSON)),
    };
}

