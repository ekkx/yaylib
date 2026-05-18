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
 * @interface GiftTransactionsResponse
 */
export interface GiftTransactionsResponse {
    /**
     * 
     * @type {boolean}
     * @memberof GiftTransactionsResponse
     */
    hideGiftsReceived?: boolean | null;
    /**
     * 
     * @type {string}
     * @memberof GiftTransactionsResponse
     */
    nextPageValue?: string | null;
    /**
     * 
     * @type {Array<GiftHistory>}
     * @memberof GiftTransactionsResponse
     */
    sentGifts?: Array<GiftHistory> | null;
}

/**
 * Check if a given object implements the GiftTransactionsResponse interface.
 */
export function instanceOfGiftTransactionsResponse(value: object): value is GiftTransactionsResponse {
    return true;
}

export function GiftTransactionsResponseFromJSON(json: any): GiftTransactionsResponse {
    return GiftTransactionsResponseFromJSONTyped(json, false);
}

export function GiftTransactionsResponseFromJSONTyped(json: any, ignoreDiscriminator: boolean): GiftTransactionsResponse {
    if (json == null) {
        return json;
    }
    return {
        
        'hideGiftsReceived': json['hide_gifts_received'] == null ? undefined : json['hide_gifts_received'],
        'nextPageValue': json['next_page_value'] == null ? undefined : String(json['next_page_value']),
        'sentGifts': json['sent_gifts'] == null ? undefined : ((json['sent_gifts'] as Array<any>).map(GiftHistoryFromJSON)),
    };
}

export function GiftTransactionsResponseToJSON(json: any): GiftTransactionsResponse {
    return GiftTransactionsResponseToJSONTyped(json, false);
}

export function GiftTransactionsResponseToJSONTyped(value?: GiftTransactionsResponse | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'hide_gifts_received': value['hideGiftsReceived'],
        'next_page_value': value['nextPageValue'],
        'sent_gifts': value['sentGifts'] == null ? undefined : ((value['sentGifts'] as Array<any>).map(GiftHistoryToJSON)),
    };
}

