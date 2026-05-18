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
import type { GroupGiftHistory } from './GroupGiftHistory.js';
import {
    GroupGiftHistoryFromJSON,
    GroupGiftHistoryFromJSONTyped,
    GroupGiftHistoryToJSON,
    GroupGiftHistoryToJSONTyped,
} from './GroupGiftHistory.js';

/**
 * 
 * @export
 * @interface GroupGiftHistoryResponse
 */
export interface GroupGiftHistoryResponse {
    /**
     * 
     * @type {Array<GroupGiftHistory>}
     * @memberof GroupGiftHistoryResponse
     */
    giftHistory?: Array<GroupGiftHistory> | null;
    /**
     * 
     * @type {string}
     * @memberof GroupGiftHistoryResponse
     */
    nextPageValue?: string | null;
}

/**
 * Check if a given object implements the GroupGiftHistoryResponse interface.
 */
export function instanceOfGroupGiftHistoryResponse(value: object): value is GroupGiftHistoryResponse {
    return true;
}

export function GroupGiftHistoryResponseFromJSON(json: any): GroupGiftHistoryResponse {
    return GroupGiftHistoryResponseFromJSONTyped(json, false);
}

export function GroupGiftHistoryResponseFromJSONTyped(json: any, ignoreDiscriminator: boolean): GroupGiftHistoryResponse {
    if (json == null) {
        return json;
    }
    return {
        
        'giftHistory': json['gift_history'] == null ? undefined : ((json['gift_history'] as Array<any>).map(GroupGiftHistoryFromJSON)),
        'nextPageValue': json['next_page_value'] == null ? undefined : String(json['next_page_value']),
    };
}

export function GroupGiftHistoryResponseToJSON(json: any): GroupGiftHistoryResponse {
    return GroupGiftHistoryResponseToJSONTyped(json, false);
}

export function GroupGiftHistoryResponseToJSONTyped(value?: GroupGiftHistoryResponse | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'gift_history': value['giftHistory'] == null ? undefined : ((value['giftHistory'] as Array<any>).map(GroupGiftHistoryToJSON)),
        'next_page_value': value['nextPageValue'],
    };
}

