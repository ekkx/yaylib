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
import type { FeaturesPalDTO } from './FeaturesPalDTO.js';
import {
    FeaturesPalDTOFromJSON,
    FeaturesPalDTOFromJSONTyped,
    FeaturesPalDTOToJSON,
    FeaturesPalDTOToJSONTyped,
} from './FeaturesPalDTO.js';

/**
 * 
 * @export
 * @interface FeaturesDTO
 */
export interface FeaturesDTO {
    /**
     * 
     * @type {FeaturesPalDTO}
     * @memberof FeaturesDTO
     */
    pal?: FeaturesPalDTO | null;
}

/**
 * Check if a given object implements the FeaturesDTO interface.
 */
export function instanceOfFeaturesDTO(value: object): value is FeaturesDTO {
    return true;
}

export function FeaturesDTOFromJSON(json: any): FeaturesDTO {
    return FeaturesDTOFromJSONTyped(json, false);
}

export function FeaturesDTOFromJSONTyped(json: any, ignoreDiscriminator: boolean): FeaturesDTO {
    if (json == null) {
        return json;
    }
    return {
        
        'pal': json['pal'] == null ? undefined : FeaturesPalDTOFromJSON(json['pal']),
    };
}

export function FeaturesDTOToJSON(json: any): FeaturesDTO {
    return FeaturesDTOToJSONTyped(json, false);
}

export function FeaturesDTOToJSONTyped(value?: FeaturesDTO | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'pal': FeaturesPalDTOToJSON(value['pal']),
    };
}

