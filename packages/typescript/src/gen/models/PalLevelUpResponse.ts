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
import type { LevelUpDetails } from './LevelUpDetails';
import {
    LevelUpDetailsFromJSON,
    LevelUpDetailsFromJSONTyped,
    LevelUpDetailsToJSON,
    LevelUpDetailsToJSONTyped,
} from './LevelUpDetails';

/**
 * 
 * @export
 * @interface PalLevelUpResponse
 */
export interface PalLevelUpResponse {
    /**
     * 
     * @type {LevelUpDetails}
     * @memberof PalLevelUpResponse
     */
    result?: LevelUpDetails | null;
}

/**
 * Check if a given object implements the PalLevelUpResponse interface.
 */
export function instanceOfPalLevelUpResponse(value: object): value is PalLevelUpResponse {
    return true;
}

export function PalLevelUpResponseFromJSON(json: any): PalLevelUpResponse {
    return PalLevelUpResponseFromJSONTyped(json, false);
}

export function PalLevelUpResponseFromJSONTyped(json: any, ignoreDiscriminator: boolean): PalLevelUpResponse {
    if (json == null) {
        return json;
    }
    return {
        
        'result': json['result'] == null ? undefined : LevelUpDetailsFromJSON(json['result']),
    };
}

export function PalLevelUpResponseToJSON(json: any): PalLevelUpResponse {
    return PalLevelUpResponseToJSONTyped(json, false);
}

export function PalLevelUpResponseToJSONTyped(value?: PalLevelUpResponse | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'result': LevelUpDetailsToJSON(value['result']),
    };
}

