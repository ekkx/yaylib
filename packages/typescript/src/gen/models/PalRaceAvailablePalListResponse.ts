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
import type { PalRaceItemDTO } from './PalRaceItemDTO';
import {
    PalRaceItemDTOFromJSON,
    PalRaceItemDTOFromJSONTyped,
    PalRaceItemDTOToJSON,
    PalRaceItemDTOToJSONTyped,
} from './PalRaceItemDTO';

/**
 * 
 * @export
 * @interface PalRaceAvailablePalListResponse
 */
export interface PalRaceAvailablePalListResponse {
    /**
     * 
     * @type {Array<PalRaceItemDTO>}
     * @memberof PalRaceAvailablePalListResponse
     */
    pals?: Array<PalRaceItemDTO> | null;
}

/**
 * Check if a given object implements the PalRaceAvailablePalListResponse interface.
 */
export function instanceOfPalRaceAvailablePalListResponse(value: object): value is PalRaceAvailablePalListResponse {
    return true;
}

export function PalRaceAvailablePalListResponseFromJSON(json: any): PalRaceAvailablePalListResponse {
    return PalRaceAvailablePalListResponseFromJSONTyped(json, false);
}

export function PalRaceAvailablePalListResponseFromJSONTyped(json: any, ignoreDiscriminator: boolean): PalRaceAvailablePalListResponse {
    if (json == null) {
        return json;
    }
    return {
        
        'pals': json['pals'] == null ? undefined : ((json['pals'] as Array<any>).map(PalRaceItemDTOFromJSON)),
    };
}

export function PalRaceAvailablePalListResponseToJSON(json: any): PalRaceAvailablePalListResponse {
    return PalRaceAvailablePalListResponseToJSONTyped(json, false);
}

export function PalRaceAvailablePalListResponseToJSONTyped(value?: PalRaceAvailablePalListResponse | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'pals': value['pals'] == null ? undefined : ((value['pals'] as Array<any>).map(PalRaceItemDTOToJSON)),
    };
}

