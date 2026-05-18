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
import type { PalRaceLeagueDTO } from './PalRaceLeagueDTO.js';
import {
    PalRaceLeagueDTOFromJSON,
    PalRaceLeagueDTOFromJSONTyped,
    PalRaceLeagueDTOToJSON,
    PalRaceLeagueDTOToJSONTyped,
} from './PalRaceLeagueDTO.js';

/**
 * 
 * @export
 * @interface PalRaceLeaguesResponse
 */
export interface PalRaceLeaguesResponse {
    /**
     * 
     * @type {Array<PalRaceLeagueDTO>}
     * @memberof PalRaceLeaguesResponse
     */
    leagues?: Array<PalRaceLeagueDTO> | null;
}

/**
 * Check if a given object implements the PalRaceLeaguesResponse interface.
 */
export function instanceOfPalRaceLeaguesResponse(value: object): value is PalRaceLeaguesResponse {
    return true;
}

export function PalRaceLeaguesResponseFromJSON(json: any): PalRaceLeaguesResponse {
    return PalRaceLeaguesResponseFromJSONTyped(json, false);
}

export function PalRaceLeaguesResponseFromJSONTyped(json: any, ignoreDiscriminator: boolean): PalRaceLeaguesResponse {
    if (json == null) {
        return json;
    }
    return {
        
        'leagues': json['leagues'] == null ? undefined : ((json['leagues'] as Array<any>).map(PalRaceLeagueDTOFromJSON)),
    };
}

export function PalRaceLeaguesResponseToJSON(json: any): PalRaceLeaguesResponse {
    return PalRaceLeaguesResponseToJSONTyped(json, false);
}

export function PalRaceLeaguesResponseToJSONTyped(value?: PalRaceLeaguesResponse | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'leagues': value['leagues'] == null ? undefined : ((value['leagues'] as Array<any>).map(PalRaceLeagueDTOToJSON)),
    };
}

