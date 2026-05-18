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
import type { EmplDTO } from './EmplDTO.js';
import {
    EmplDTOFromJSON,
    EmplDTOFromJSONTyped,
    EmplDTOToJSON,
    EmplDTOToJSONTyped,
} from './EmplDTO.js';

/**
 * 
 * @export
 * @interface BagDTO
 */
export interface BagDTO {
    /**
     * 
     * @type {number}
     * @memberof BagDTO
     */
    eggsCount?: number | null;
    /**
     * 
     * @type {EmplDTO}
     * @memberof BagDTO
     */
    empl?: EmplDTO | null;
    /**
     * 
     * @type {string}
     * @memberof BagDTO
     */
    expiringEmpl?: string | null;
    /**
     * 
     * @type {number}
     * @memberof BagDTO
     */
    palsCount?: number | null;
    /**
     * 
     * @type {number}
     * @memberof BagDTO
     */
    racePalAvailable?: number | null;
}

/**
 * Check if a given object implements the BagDTO interface.
 */
export function instanceOfBagDTO(value: object): value is BagDTO {
    return true;
}

export function BagDTOFromJSON(json: any): BagDTO {
    return BagDTOFromJSONTyped(json, false);
}

export function BagDTOFromJSONTyped(json: any, ignoreDiscriminator: boolean): BagDTO {
    if (json == null) {
        return json;
    }
    return {
        
        'eggsCount': json['eggs_count'] == null ? undefined : json['eggs_count'],
        'empl': json['empl'] == null ? undefined : EmplDTOFromJSON(json['empl']),
        'expiringEmpl': json['expiring_empl'] == null ? undefined : json['expiring_empl'],
        'palsCount': json['pals_count'] == null ? undefined : json['pals_count'],
        'racePalAvailable': json['race_pal_available'] == null ? undefined : json['race_pal_available'],
    };
}

export function BagDTOToJSON(json: any): BagDTO {
    return BagDTOToJSONTyped(json, false);
}

export function BagDTOToJSONTyped(value?: BagDTO | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'eggs_count': value['eggsCount'],
        'empl': EmplDTOToJSON(value['empl']),
        'expiring_empl': value['expiringEmpl'],
        'pals_count': value['palsCount'],
        'race_pal_available': value['racePalAvailable'],
    };
}

