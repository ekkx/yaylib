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
import type { PalDTO } from './PalDTO.js';
import {
    PalDTOFromJSON,
    PalDTOFromJSONTyped,
    PalDTOToJSON,
    PalDTOToJSONTyped,
} from './PalDTO.js';

/**
 * 
 * @export
 * @interface PalListResponse
 */
export interface PalListResponse {
    /**
     * 
     * @type {string}
     * @memberof PalListResponse
     */
    cursor?: string | null;
    /**
     * 
     * @type {Array<PalDTO>}
     * @memberof PalListResponse
     */
    result?: Array<PalDTO> | null;
    /**
     * 
     * @type {number}
     * @memberof PalListResponse
     */
    totalCount?: number | null;
}

/**
 * Check if a given object implements the PalListResponse interface.
 */
export function instanceOfPalListResponse(value: object): value is PalListResponse {
    return true;
}

export function PalListResponseFromJSON(json: any): PalListResponse {
    return PalListResponseFromJSONTyped(json, false);
}

export function PalListResponseFromJSONTyped(json: any, ignoreDiscriminator: boolean): PalListResponse {
    if (json == null) {
        return json;
    }
    return {
        
        'cursor': json['cursor'] == null ? undefined : json['cursor'],
        'result': json['result'] == null ? undefined : ((json['result'] as Array<any>).map(PalDTOFromJSON)),
        'totalCount': json['total_count'] == null ? undefined : json['total_count'],
    };
}

export function PalListResponseToJSON(json: any): PalListResponse {
    return PalListResponseToJSONTyped(json, false);
}

export function PalListResponseToJSONTyped(value?: PalListResponse | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'cursor': value['cursor'],
        'result': value['result'] == null ? undefined : ((value['result'] as Array<any>).map(PalDTOToJSON)),
        'total_count': value['totalCount'],
    };
}

