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
import type { PalFreePoolStatus } from './PalFreePoolStatus';
import {
    PalFreePoolStatusFromJSON,
    PalFreePoolStatusFromJSONTyped,
    PalFreePoolStatusToJSON,
    PalFreePoolStatusToJSONTyped,
} from './PalFreePoolStatus';

/**
 * 
 * @export
 * @interface PalFreePoolStatusResponse
 */
export interface PalFreePoolStatusResponse {
    /**
     * 
     * @type {PalFreePoolStatus}
     * @memberof PalFreePoolStatusResponse
     */
    result?: PalFreePoolStatus | null;
}

/**
 * Check if a given object implements the PalFreePoolStatusResponse interface.
 */
export function instanceOfPalFreePoolStatusResponse(value: object): value is PalFreePoolStatusResponse {
    return true;
}

export function PalFreePoolStatusResponseFromJSON(json: any): PalFreePoolStatusResponse {
    return PalFreePoolStatusResponseFromJSONTyped(json, false);
}

export function PalFreePoolStatusResponseFromJSONTyped(json: any, ignoreDiscriminator: boolean): PalFreePoolStatusResponse {
    if (json == null) {
        return json;
    }
    return {
        
        'result': json['result'] == null ? undefined : PalFreePoolStatusFromJSON(json['result']),
    };
}

export function PalFreePoolStatusResponseToJSON(json: any): PalFreePoolStatusResponse {
    return PalFreePoolStatusResponseToJSONTyped(json, false);
}

export function PalFreePoolStatusResponseToJSONTyped(value?: PalFreePoolStatusResponse | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'result': PalFreePoolStatusToJSON(value['result']),
    };
}

