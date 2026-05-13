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
import type { PalDTO } from './PalDTO';
import {
    PalDTOFromJSON,
    PalDTOFromJSONTyped,
    PalDTOToJSON,
    PalDTOToJSONTyped,
} from './PalDTO';

/**
 * 
 * @export
 * @interface PalResponse
 */
export interface PalResponse {
    /**
     * 
     * @type {PalDTO}
     * @memberof PalResponse
     */
    result?: PalDTO | null;
}

/**
 * Check if a given object implements the PalResponse interface.
 */
export function instanceOfPalResponse(value: object): value is PalResponse {
    return true;
}

export function PalResponseFromJSON(json: any): PalResponse {
    return PalResponseFromJSONTyped(json, false);
}

export function PalResponseFromJSONTyped(json: any, ignoreDiscriminator: boolean): PalResponse {
    if (json == null) {
        return json;
    }
    return {
        
        'result': json['result'] == null ? undefined : PalDTOFromJSON(json['result']),
    };
}

export function PalResponseToJSON(json: any): PalResponse {
    return PalResponseToJSONTyped(json, false);
}

export function PalResponseToJSONTyped(value?: PalResponse | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'result': PalDTOToJSON(value['result']),
    };
}

