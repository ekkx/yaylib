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
import type { FeaturesResultResponse } from './FeaturesResultResponse.js';
import {
    FeaturesResultResponseFromJSON,
    FeaturesResultResponseFromJSONTyped,
    FeaturesResultResponseToJSON,
    FeaturesResultResponseToJSONTyped,
} from './FeaturesResultResponse.js';

/**
 * 
 * @export
 * @interface Web3FeaturesResponse
 */
export interface Web3FeaturesResponse {
    /**
     * 
     * @type {FeaturesResultResponse}
     * @memberof Web3FeaturesResponse
     */
    result?: FeaturesResultResponse | null;
}

/**
 * Check if a given object implements the Web3FeaturesResponse interface.
 */
export function instanceOfWeb3FeaturesResponse(value: object): value is Web3FeaturesResponse {
    return true;
}

export function Web3FeaturesResponseFromJSON(json: any): Web3FeaturesResponse {
    return Web3FeaturesResponseFromJSONTyped(json, false);
}

export function Web3FeaturesResponseFromJSONTyped(json: any, ignoreDiscriminator: boolean): Web3FeaturesResponse {
    if (json == null) {
        return json;
    }
    return {
        
        'result': json['result'] == null ? undefined : FeaturesResultResponseFromJSON(json['result']),
    };
}

export function Web3FeaturesResponseToJSON(json: any): Web3FeaturesResponse {
    return Web3FeaturesResponseToJSONTyped(json, false);
}

export function Web3FeaturesResponseToJSONTyped(value?: Web3FeaturesResponse | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'result': FeaturesResultResponseToJSON(value['result']),
    };
}

