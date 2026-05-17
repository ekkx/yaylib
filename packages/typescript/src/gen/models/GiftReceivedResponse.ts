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
import type { ReceivedGift } from './ReceivedGift';
import {
    ReceivedGiftFromJSON,
    ReceivedGiftFromJSONTyped,
    ReceivedGiftToJSON,
    ReceivedGiftToJSONTyped,
} from './ReceivedGift';

/**
 * 
 * @export
 * @interface GiftReceivedResponse
 */
export interface GiftReceivedResponse {
    /**
     * 
     * @type {string}
     * @memberof GiftReceivedResponse
     */
    nextPageValue?: string | null;
    /**
     * 
     * @type {Array<ReceivedGift>}
     * @memberof GiftReceivedResponse
     */
    receivedGifts?: Array<ReceivedGift> | null;
    /**
     * 
     * @type {number}
     * @memberof GiftReceivedResponse
     */
    totalCount?: number | null;
}

/**
 * Check if a given object implements the GiftReceivedResponse interface.
 */
export function instanceOfGiftReceivedResponse(value: object): value is GiftReceivedResponse {
    return true;
}

export function GiftReceivedResponseFromJSON(json: any): GiftReceivedResponse {
    return GiftReceivedResponseFromJSONTyped(json, false);
}

export function GiftReceivedResponseFromJSONTyped(json: any, ignoreDiscriminator: boolean): GiftReceivedResponse {
    if (json == null) {
        return json;
    }
    return {
        
        'nextPageValue': json['next_page_value'] == null ? undefined : String(json['next_page_value']),
        'receivedGifts': json['received_gifts'] == null ? undefined : ((json['received_gifts'] as Array<any>).map(ReceivedGiftFromJSON)),
        'totalCount': json['total_count'] == null ? undefined : json['total_count'],
    };
}

export function GiftReceivedResponseToJSON(json: any): GiftReceivedResponse {
    return GiftReceivedResponseToJSONTyped(json, false);
}

export function GiftReceivedResponseToJSONTyped(value?: GiftReceivedResponse | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'next_page_value': value['nextPageValue'],
        'received_gifts': value['receivedGifts'] == null ? undefined : ((value['receivedGifts'] as Array<any>).map(ReceivedGiftToJSON)),
        'total_count': value['totalCount'],
    };
}

