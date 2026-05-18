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
/**
 * 
 * @export
 * @interface PriceDTO
 */
export interface PriceDTO {
    /**
     * 
     * @type {string}
     * @memberof PriceDTO
     */
    jpy?: string | null;
    /**
     * 
     * @type {string}
     * @memberof PriceDTO
     */
    usd?: string | null;
}

/**
 * Check if a given object implements the PriceDTO interface.
 */
export function instanceOfPriceDTO(value: object): value is PriceDTO {
    return true;
}

export function PriceDTOFromJSON(json: any): PriceDTO {
    return PriceDTOFromJSONTyped(json, false);
}

export function PriceDTOFromJSONTyped(json: any, ignoreDiscriminator: boolean): PriceDTO {
    if (json == null) {
        return json;
    }
    return {
        
        'jpy': json['jpy'] == null ? undefined : json['jpy'],
        'usd': json['usd'] == null ? undefined : json['usd'],
    };
}

export function PriceDTOToJSON(json: any): PriceDTO {
    return PriceDTOToJSONTyped(json, false);
}

export function PriceDTOToJSONTyped(value?: PriceDTO | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'jpy': value['jpy'],
        'usd': value['usd'],
    };
}

