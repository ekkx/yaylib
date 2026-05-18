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
import type { BagDTO } from './BagDTO.js';
import {
    BagDTOFromJSON,
    BagDTOFromJSONTyped,
    BagDTOToJSON,
    BagDTOToJSONTyped,
} from './BagDTO.js';

/**
 * 
 * @export
 * @interface BagResponse
 */
export interface BagResponse {
    /**
     * 
     * @type {BagDTO}
     * @memberof BagResponse
     */
    result?: BagDTO | null;
}

/**
 * Check if a given object implements the BagResponse interface.
 */
export function instanceOfBagResponse(value: object): value is BagResponse {
    return true;
}

export function BagResponseFromJSON(json: any): BagResponse {
    return BagResponseFromJSONTyped(json, false);
}

export function BagResponseFromJSONTyped(json: any, ignoreDiscriminator: boolean): BagResponse {
    if (json == null) {
        return json;
    }
    return {
        
        'result': json['result'] == null ? undefined : BagDTOFromJSON(json['result']),
    };
}

export function BagResponseToJSON(json: any): BagResponse {
    return BagResponseToJSONTyped(json, false);
}

export function BagResponseToJSONTyped(value?: BagResponse | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'result': BagDTOToJSON(value['result']),
    };
}

