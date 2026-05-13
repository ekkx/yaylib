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
 * @interface EmplActivityResponse
 */
export interface EmplActivityResponse {
    /**
     * 
     * @type {string}
     * @memberof EmplActivityResponse
     */
    cursor?: string | null;
    /**
     * 
     * @type {Array<EmplActivityDTO>}
     * @memberof EmplActivityResponse
     */
    result?: Array<EmplActivityDTO> | null;
}

/**
 * Check if a given object implements the EmplActivityResponse interface.
 */
export function instanceOfEmplActivityResponse(value: object): value is EmplActivityResponse {
    return true;
}

export function EmplActivityResponseFromJSON(json: any): EmplActivityResponse {
    return EmplActivityResponseFromJSONTyped(json, false);
}

export function EmplActivityResponseFromJSONTyped(json: any, ignoreDiscriminator: boolean): EmplActivityResponse {
    if (json == null) {
        return json;
    }
    return {
        
        'cursor': json['cursor'] == null ? undefined : json['cursor'],
        'result': json['result'] == null ? undefined : ((json['result'] as Array<any>).map(EmplActivityDTOFromJSON)),
    };
}

export function EmplActivityResponseToJSON(json: any): EmplActivityResponse {
    return EmplActivityResponseToJSONTyped(json, false);
}

export function EmplActivityResponseToJSONTyped(value?: EmplActivityResponse | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'cursor': value['cursor'],
        'result': value['result'] == null ? undefined : ((value['result'] as Array<any>).map(EmplActivityDTOToJSON)),
    };
}

