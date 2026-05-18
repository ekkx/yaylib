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
import type { GiftHistory } from './GiftHistory.js';
import {
    GiftHistoryFromJSON,
    GiftHistoryFromJSONTyped,
    GiftHistoryToJSON,
    GiftHistoryToJSONTyped,
} from './GiftHistory.js';

/**
 * 
 * @export
 * @interface CallGiftHistoryResponse
 */
export interface CallGiftHistoryResponse {
    /**
     * 
     * @type {string}
     * @memberof CallGiftHistoryResponse
     */
    nextPageValue?: string | null;
    /**
     * 
     * @type {Array<GiftHistory>}
     * @memberof CallGiftHistoryResponse
     */
    sentGifts?: Array<GiftHistory> | null;
}

/**
 * Check if a given object implements the CallGiftHistoryResponse interface.
 */
export function instanceOfCallGiftHistoryResponse(value: object): value is CallGiftHistoryResponse {
    return true;
}

export function CallGiftHistoryResponseFromJSON(json: any): CallGiftHistoryResponse {
    return CallGiftHistoryResponseFromJSONTyped(json, false);
}

export function CallGiftHistoryResponseFromJSONTyped(json: any, ignoreDiscriminator: boolean): CallGiftHistoryResponse {
    if (json == null) {
        return json;
    }
    return {
        
        'nextPageValue': json['next_page_value'] == null ? undefined : String(json['next_page_value']),
        'sentGifts': json['sent_gifts'] == null ? undefined : ((json['sent_gifts'] as Array<any>).map(GiftHistoryFromJSON)),
    };
}

export function CallGiftHistoryResponseToJSON(json: any): CallGiftHistoryResponse {
    return CallGiftHistoryResponseToJSONTyped(json, false);
}

export function CallGiftHistoryResponseToJSONTyped(value?: CallGiftHistoryResponse | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'next_page_value': value['nextPageValue'],
        'sent_gifts': value['sentGifts'] == null ? undefined : ((value['sentGifts'] as Array<any>).map(GiftHistoryToJSON)),
    };
}

