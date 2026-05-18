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
import type { EmplFee } from './EmplFee.js';
import {
    EmplFeeFromJSON,
    EmplFeeFromJSONTyped,
    EmplFeeToJSON,
    EmplFeeToJSONTyped,
} from './EmplFee.js';

/**
 * 
 * @export
 * @interface EmplFeeResponse
 */
export interface EmplFeeResponse {
    /**
     * 
     * @type {EmplFee}
     * @memberof EmplFeeResponse
     */
    result?: EmplFee | null;
}

/**
 * Check if a given object implements the EmplFeeResponse interface.
 */
export function instanceOfEmplFeeResponse(value: object): value is EmplFeeResponse {
    return true;
}

export function EmplFeeResponseFromJSON(json: any): EmplFeeResponse {
    return EmplFeeResponseFromJSONTyped(json, false);
}

export function EmplFeeResponseFromJSONTyped(json: any, ignoreDiscriminator: boolean): EmplFeeResponse {
    if (json == null) {
        return json;
    }
    return {
        
        'result': json['result'] == null ? undefined : EmplFeeFromJSON(json['result']),
    };
}

export function EmplFeeResponseToJSON(json: any): EmplFeeResponse {
    return EmplFeeResponseToJSONTyped(json, false);
}

export function EmplFeeResponseToJSONTyped(value?: EmplFeeResponse | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'result': EmplFeeToJSON(value['result']),
    };
}

