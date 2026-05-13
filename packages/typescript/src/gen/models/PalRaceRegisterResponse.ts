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
import type { RaceHistory } from './RaceHistory';
import {
    RaceHistoryFromJSON,
    RaceHistoryFromJSONTyped,
    RaceHistoryToJSON,
    RaceHistoryToJSONTyped,
} from './RaceHistory';

/**
 * 
 * @export
 * @interface PalRaceRegisterResponse
 */
export interface PalRaceRegisterResponse {
    /**
     * 
     * @type {RaceHistory}
     * @memberof PalRaceRegisterResponse
     */
    raceHistory?: RaceHistory | null;
}

/**
 * Check if a given object implements the PalRaceRegisterResponse interface.
 */
export function instanceOfPalRaceRegisterResponse(value: object): value is PalRaceRegisterResponse {
    return true;
}

export function PalRaceRegisterResponseFromJSON(json: any): PalRaceRegisterResponse {
    return PalRaceRegisterResponseFromJSONTyped(json, false);
}

export function PalRaceRegisterResponseFromJSONTyped(json: any, ignoreDiscriminator: boolean): PalRaceRegisterResponse {
    if (json == null) {
        return json;
    }
    return {
        
        'raceHistory': json['race_history'] == null ? undefined : RaceHistoryFromJSON(json['race_history']),
    };
}

export function PalRaceRegisterResponseToJSON(json: any): PalRaceRegisterResponse {
    return PalRaceRegisterResponseToJSONTyped(json, false);
}

export function PalRaceRegisterResponseToJSONTyped(value?: PalRaceRegisterResponse | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'race_history': RaceHistoryToJSON(value['raceHistory']),
    };
}

