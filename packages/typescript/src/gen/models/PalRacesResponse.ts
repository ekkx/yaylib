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
import type { RaceHistory } from './RaceHistory.js';
import {
    RaceHistoryFromJSON,
    RaceHistoryFromJSONTyped,
    RaceHistoryToJSON,
    RaceHistoryToJSONTyped,
} from './RaceHistory.js';

/**
 * 
 * @export
 * @interface PalRacesResponse
 */
export interface PalRacesResponse {
    /**
     * 
     * @type {Array<RaceHistory>}
     * @memberof PalRacesResponse
     */
    races?: Array<RaceHistory> | null;
}

/**
 * Check if a given object implements the PalRacesResponse interface.
 */
export function instanceOfPalRacesResponse(value: object): value is PalRacesResponse {
    return true;
}

export function PalRacesResponseFromJSON(json: any): PalRacesResponse {
    return PalRacesResponseFromJSONTyped(json, false);
}

export function PalRacesResponseFromJSONTyped(json: any, ignoreDiscriminator: boolean): PalRacesResponse {
    if (json == null) {
        return json;
    }
    return {
        
        'races': json['races'] == null ? undefined : ((json['races'] as Array<any>).map(RaceHistoryFromJSON)),
    };
}

export function PalRacesResponseToJSON(json: any): PalRacesResponse {
    return PalRacesResponseToJSONTyped(json, false);
}

export function PalRacesResponseToJSONTyped(value?: PalRacesResponse | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'races': value['races'] == null ? undefined : ((value['races'] as Array<any>).map(RaceHistoryToJSON)),
    };
}

