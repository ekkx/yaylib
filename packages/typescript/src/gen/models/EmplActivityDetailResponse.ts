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
import type { EmplActivityDTO } from './EmplActivityDTO';
import {
    EmplActivityDTOFromJSON,
    EmplActivityDTOFromJSONTyped,
    EmplActivityDTOToJSON,
    EmplActivityDTOToJSONTyped,
} from './EmplActivityDTO';

/**
 * 
 * @export
 * @interface EmplActivityDetailResponse
 */
export interface EmplActivityDetailResponse {
    /**
     * 
     * @type {EmplActivityDTO}
     * @memberof EmplActivityDetailResponse
     */
    result?: EmplActivityDTO | null;
}

/**
 * Check if a given object implements the EmplActivityDetailResponse interface.
 */
export function instanceOfEmplActivityDetailResponse(value: object): value is EmplActivityDetailResponse {
    return true;
}

export function EmplActivityDetailResponseFromJSON(json: any): EmplActivityDetailResponse {
    return EmplActivityDetailResponseFromJSONTyped(json, false);
}

export function EmplActivityDetailResponseFromJSONTyped(json: any, ignoreDiscriminator: boolean): EmplActivityDetailResponse {
    if (json == null) {
        return json;
    }
    return {
        
        'result': json['result'] == null ? undefined : EmplActivityDTOFromJSON(json['result']),
    };
}

export function EmplActivityDetailResponseToJSON(json: any): EmplActivityDetailResponse {
    return EmplActivityDetailResponseToJSONTyped(json, false);
}

export function EmplActivityDetailResponseToJSONTyped(value?: EmplActivityDetailResponse | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'result': EmplActivityDTOToJSON(value['result']),
    };
}

