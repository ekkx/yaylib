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
 * @interface PalRaceDetailsResponse
 */
export interface PalRaceDetailsResponse {
    /**
     * 
     * @type {RaceHistory}
     * @memberof PalRaceDetailsResponse
     */
    race?: RaceHistory | null;
}

/**
 * Check if a given object implements the PalRaceDetailsResponse interface.
 */
export function instanceOfPalRaceDetailsResponse(value: object): value is PalRaceDetailsResponse {
    return true;
}

export function PalRaceDetailsResponseFromJSON(json: any): PalRaceDetailsResponse {
    return PalRaceDetailsResponseFromJSONTyped(json, false);
}

export function PalRaceDetailsResponseFromJSONTyped(json: any, ignoreDiscriminator: boolean): PalRaceDetailsResponse {
    if (json == null) {
        return json;
    }
    return {
        
        'race': json['race'] == null ? undefined : RaceHistoryFromJSON(json['race']),
    };
}

export function PalRaceDetailsResponseToJSON(json: any): PalRaceDetailsResponse {
    return PalRaceDetailsResponseToJSONTyped(json, false);
}

export function PalRaceDetailsResponseToJSONTyped(value?: PalRaceDetailsResponse | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'race': RaceHistoryToJSON(value['race']),
    };
}

