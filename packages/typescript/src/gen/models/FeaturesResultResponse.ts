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
import type { FeaturesDTO } from './FeaturesDTO';
import {
    FeaturesDTOFromJSON,
    FeaturesDTOFromJSONTyped,
    FeaturesDTOToJSON,
    FeaturesDTOToJSONTyped,
} from './FeaturesDTO';

/**
 * 
 * @export
 * @interface FeaturesResultResponse
 */
export interface FeaturesResultResponse {
    /**
     * 
     * @type {FeaturesDTO}
     * @memberof FeaturesResultResponse
     */
    features?: FeaturesDTO | null;
}

/**
 * Check if a given object implements the FeaturesResultResponse interface.
 */
export function instanceOfFeaturesResultResponse(value: object): value is FeaturesResultResponse {
    return true;
}

export function FeaturesResultResponseFromJSON(json: any): FeaturesResultResponse {
    return FeaturesResultResponseFromJSONTyped(json, false);
}

export function FeaturesResultResponseFromJSONTyped(json: any, ignoreDiscriminator: boolean): FeaturesResultResponse {
    if (json == null) {
        return json;
    }
    return {
        
        'features': json['features'] == null ? undefined : FeaturesDTOFromJSON(json['features']),
    };
}

export function FeaturesResultResponseToJSON(json: any): FeaturesResultResponse {
    return FeaturesResultResponseToJSONTyped(json, false);
}

export function FeaturesResultResponseToJSONTyped(value?: FeaturesResultResponse | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'features': FeaturesDTOToJSON(value['features']),
    };
}

