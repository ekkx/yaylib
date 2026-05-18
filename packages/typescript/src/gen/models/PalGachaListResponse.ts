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
import type { PalGacha } from './PalGacha.js';
import {
    PalGachaFromJSON,
    PalGachaFromJSONTyped,
    PalGachaToJSON,
    PalGachaToJSONTyped,
} from './PalGacha.js';

/**
 * 
 * @export
 * @interface PalGachaListResponse
 */
export interface PalGachaListResponse {
    /**
     * 
     * @type {Array<PalGacha>}
     * @memberof PalGachaListResponse
     */
    result?: Array<PalGacha> | null;
}

/**
 * Check if a given object implements the PalGachaListResponse interface.
 */
export function instanceOfPalGachaListResponse(value: object): value is PalGachaListResponse {
    return true;
}

export function PalGachaListResponseFromJSON(json: any): PalGachaListResponse {
    return PalGachaListResponseFromJSONTyped(json, false);
}

export function PalGachaListResponseFromJSONTyped(json: any, ignoreDiscriminator: boolean): PalGachaListResponse {
    if (json == null) {
        return json;
    }
    return {
        
        'result': json['result'] == null ? undefined : ((json['result'] as Array<any>).map(PalGachaFromJSON)),
    };
}

export function PalGachaListResponseToJSON(json: any): PalGachaListResponse {
    return PalGachaListResponseToJSONTyped(json, false);
}

export function PalGachaListResponseToJSONTyped(value?: PalGachaListResponse | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'result': value['result'] == null ? undefined : ((value['result'] as Array<any>).map(PalGachaToJSON)),
    };
}

