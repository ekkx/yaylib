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
import type { PalGachaHistoryResponseItem } from './PalGachaHistoryResponseItem';
import {
    PalGachaHistoryResponseItemFromJSON,
    PalGachaHistoryResponseItemFromJSONTyped,
    PalGachaHistoryResponseItemToJSON,
    PalGachaHistoryResponseItemToJSONTyped,
} from './PalGachaHistoryResponseItem';

/**
 * 
 * @export
 * @interface PalGachaHistoryResponse
 */
export interface PalGachaHistoryResponse {
    /**
     * 
     * @type {string}
     * @memberof PalGachaHistoryResponse
     */
    cursor?: string | null;
    /**
     * 
     * @type {Array<PalGachaHistoryResponseItem>}
     * @memberof PalGachaHistoryResponse
     */
    result?: Array<PalGachaHistoryResponseItem> | null;
}

/**
 * Check if a given object implements the PalGachaHistoryResponse interface.
 */
export function instanceOfPalGachaHistoryResponse(value: object): value is PalGachaHistoryResponse {
    return true;
}

export function PalGachaHistoryResponseFromJSON(json: any): PalGachaHistoryResponse {
    return PalGachaHistoryResponseFromJSONTyped(json, false);
}

export function PalGachaHistoryResponseFromJSONTyped(json: any, ignoreDiscriminator: boolean): PalGachaHistoryResponse {
    if (json == null) {
        return json;
    }
    return {
        
        'cursor': json['cursor'] == null ? undefined : json['cursor'],
        'result': json['result'] == null ? undefined : ((json['result'] as Array<any>).map(PalGachaHistoryResponseItemFromJSON)),
    };
}

export function PalGachaHistoryResponseToJSON(json: any): PalGachaHistoryResponse {
    return PalGachaHistoryResponseToJSONTyped(json, false);
}

export function PalGachaHistoryResponseToJSONTyped(value?: PalGachaHistoryResponse | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'cursor': value['cursor'],
        'result': value['result'] == null ? undefined : ((value['result'] as Array<any>).map(PalGachaHistoryResponseItemToJSON)),
    };
}

