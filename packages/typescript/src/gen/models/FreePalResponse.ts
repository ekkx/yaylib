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
 * @interface FreePalResponse
 */
export interface FreePalResponse {
    /**
     * 
     * @type {PalDTO}
     * @memberof FreePalResponse
     */
    result?: PalDTO | null;
}

/**
 * Check if a given object implements the FreePalResponse interface.
 */
export function instanceOfFreePalResponse(value: object): value is FreePalResponse {
    return true;
}

export function FreePalResponseFromJSON(json: any): FreePalResponse {
    return FreePalResponseFromJSONTyped(json, false);
}

export function FreePalResponseFromJSONTyped(json: any, ignoreDiscriminator: boolean): FreePalResponse {
    if (json == null) {
        return json;
    }
    return {
        
        'result': json['result'] == null ? undefined : PalDTOFromJSON(json['result']),
    };
}

export function FreePalResponseToJSON(json: any): FreePalResponse {
    return FreePalResponseToJSONTyped(json, false);
}

export function FreePalResponseToJSONTyped(value?: FreePalResponse | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'result': PalDTOToJSON(value['result']),
    };
}

