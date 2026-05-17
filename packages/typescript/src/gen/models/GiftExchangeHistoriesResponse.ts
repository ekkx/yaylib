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
import type { GiftExchangeHistory } from './GiftExchangeHistory';
import {
    GiftExchangeHistoryFromJSON,
    GiftExchangeHistoryFromJSONTyped,
    GiftExchangeHistoryToJSON,
    GiftExchangeHistoryToJSONTyped,
} from './GiftExchangeHistory';

/**
 * 
 * @export
 * @interface GiftExchangeHistoriesResponse
 */
export interface GiftExchangeHistoriesResponse {
    /**
     * 
     * @type {Array<GiftExchangeHistory>}
     * @memberof GiftExchangeHistoriesResponse
     */
    giftExchangeHistories?: Array<GiftExchangeHistory> | null;
    /**
     * 
     * @type {string}
     * @memberof GiftExchangeHistoriesResponse
     */
    nextPageValue?: string | null;
}

/**
 * Check if a given object implements the GiftExchangeHistoriesResponse interface.
 */
export function instanceOfGiftExchangeHistoriesResponse(value: object): value is GiftExchangeHistoriesResponse {
    return true;
}

export function GiftExchangeHistoriesResponseFromJSON(json: any): GiftExchangeHistoriesResponse {
    return GiftExchangeHistoriesResponseFromJSONTyped(json, false);
}

export function GiftExchangeHistoriesResponseFromJSONTyped(json: any, ignoreDiscriminator: boolean): GiftExchangeHistoriesResponse {
    if (json == null) {
        return json;
    }
    return {
        
        'giftExchangeHistories': json['gift_exchange_histories'] == null ? undefined : ((json['gift_exchange_histories'] as Array<any>).map(GiftExchangeHistoryFromJSON)),
        'nextPageValue': json['next_page_value'] == null ? undefined : String(json['next_page_value']),
    };
}

export function GiftExchangeHistoriesResponseToJSON(json: any): GiftExchangeHistoriesResponse {
    return GiftExchangeHistoriesResponseToJSONTyped(json, false);
}

export function GiftExchangeHistoriesResponseToJSONTyped(value?: GiftExchangeHistoriesResponse | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'gift_exchange_histories': value['giftExchangeHistories'] == null ? undefined : ((value['giftExchangeHistories'] as Array<any>).map(GiftExchangeHistoryToJSON)),
        'next_page_value': value['nextPageValue'],
    };
}

