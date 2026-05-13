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
 * @interface PalDetailsResponse
 */
export interface PalDetailsResponse {
    /**
     * 
     * @type {PalDTO}
     * @memberof PalDetailsResponse
     */
    result?: PalDTO | null;
}

/**
 * Check if a given object implements the PalDetailsResponse interface.
 */
export function instanceOfPalDetailsResponse(value: object): value is PalDetailsResponse {
    return true;
}

export function PalDetailsResponseFromJSON(json: any): PalDetailsResponse {
    return PalDetailsResponseFromJSONTyped(json, false);
}

export function PalDetailsResponseFromJSONTyped(json: any, ignoreDiscriminator: boolean): PalDetailsResponse {
    if (json == null) {
        return json;
    }
    return {
        
        'result': json['result'] == null ? undefined : PalDTOFromJSON(json['result']),
    };
}

export function PalDetailsResponseToJSON(json: any): PalDetailsResponse {
    return PalDetailsResponseToJSONTyped(json, false);
}

export function PalDetailsResponseToJSONTyped(value?: PalDetailsResponse | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'result': PalDTOToJSON(value['result']),
    };
}

