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
import type { FootprintDTO } from './FootprintDTO';
import {
    FootprintDTOFromJSON,
    FootprintDTOFromJSONTyped,
    FootprintDTOToJSON,
    FootprintDTOToJSONTyped,
} from './FootprintDTO';

/**
 * 
 * @export
 * @interface FootprintsResponse
 */
export interface FootprintsResponse {
    /**
     * 
     * @type {Array<FootprintDTO>}
     * @memberof FootprintsResponse
     */
    footprints?: Array<FootprintDTO> | null;
    /**
     * 
     * @type {string}
     * @memberof FootprintsResponse
     */
    nextPageValue?: string | null;
}

/**
 * Check if a given object implements the FootprintsResponse interface.
 */
export function instanceOfFootprintsResponse(value: object): value is FootprintsResponse {
    return true;
}

export function FootprintsResponseFromJSON(json: any): FootprintsResponse {
    return FootprintsResponseFromJSONTyped(json, false);
}

export function FootprintsResponseFromJSONTyped(json: any, ignoreDiscriminator: boolean): FootprintsResponse {
    if (json == null) {
        return json;
    }
    return {
        
        'footprints': json['footprints'] == null ? undefined : ((json['footprints'] as Array<any>).map(FootprintDTOFromJSON)),
        'nextPageValue': json['next_page_value'] == null ? undefined : json['next_page_value'],
    };
}

export function FootprintsResponseToJSON(json: any): FootprintsResponse {
    return FootprintsResponseToJSONTyped(json, false);
}

export function FootprintsResponseToJSONTyped(value?: FootprintsResponse | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'footprints': value['footprints'] == null ? undefined : ((value['footprints'] as Array<any>).map(FootprintDTOToJSON)),
        'next_page_value': value['nextPageValue'],
    };
}

