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
import type { RealmGift } from './RealmGift.js';
import {
    RealmGiftFromJSON,
    RealmGiftFromJSONTyped,
    RealmGiftToJSON,
    RealmGiftToJSONTyped,
} from './RealmGift.js';

/**
 * 
 * @export
 * @interface GiftsResponse
 */
export interface GiftsResponse {
    /**
     * 
     * @type {Array<RealmGift>}
     * @memberof GiftsResponse
     */
    gifts?: Array<RealmGift> | null;
    /**
     * 
     * @type {string}
     * @memberof GiftsResponse
     */
    nextPageValue?: string | null;
    /**
     * 
     * @type {number}
     * @memberof GiftsResponse
     */
    totalCount?: number | null;
}

/**
 * Check if a given object implements the GiftsResponse interface.
 */
export function instanceOfGiftsResponse(value: object): value is GiftsResponse {
    return true;
}

export function GiftsResponseFromJSON(json: any): GiftsResponse {
    return GiftsResponseFromJSONTyped(json, false);
}

export function GiftsResponseFromJSONTyped(json: any, ignoreDiscriminator: boolean): GiftsResponse {
    if (json == null) {
        return json;
    }
    return {
        
        'gifts': json['gifts'] == null ? undefined : ((json['gifts'] as Array<any>).map(RealmGiftFromJSON)),
        'nextPageValue': json['next_page_value'] == null ? undefined : String(json['next_page_value']),
        'totalCount': json['total_count'] == null ? undefined : json['total_count'],
    };
}

export function GiftsResponseToJSON(json: any): GiftsResponse {
    return GiftsResponseToJSONTyped(json, false);
}

export function GiftsResponseToJSONTyped(value?: GiftsResponse | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'gifts': value['gifts'] == null ? undefined : ((value['gifts'] as Array<any>).map(RealmGiftToJSON)),
        'next_page_value': value['nextPageValue'],
        'total_count': value['totalCount'],
    };
}

