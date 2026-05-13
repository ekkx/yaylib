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
/**
 * 
 * @export
 * @interface PalGachaAvailabilityResponse
 */
export interface PalGachaAvailabilityResponse {
    /**
     * 
     * @type {string}
     * @memberof PalGachaAvailabilityResponse
     */
    state?: string | null;
}

/**
 * Check if a given object implements the PalGachaAvailabilityResponse interface.
 */
export function instanceOfPalGachaAvailabilityResponse(value: object): value is PalGachaAvailabilityResponse {
    return true;
}

export function PalGachaAvailabilityResponseFromJSON(json: any): PalGachaAvailabilityResponse {
    return PalGachaAvailabilityResponseFromJSONTyped(json, false);
}

export function PalGachaAvailabilityResponseFromJSONTyped(json: any, ignoreDiscriminator: boolean): PalGachaAvailabilityResponse {
    if (json == null) {
        return json;
    }
    return {
        
        'state': json['state'] == null ? undefined : json['state'],
    };
}

export function PalGachaAvailabilityResponseToJSON(json: any): PalGachaAvailabilityResponse {
    return PalGachaAvailabilityResponseToJSONTyped(json, false);
}

export function PalGachaAvailabilityResponseToJSONTyped(value?: PalGachaAvailabilityResponse | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'state': value['state'],
    };
}

