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
import type { PalDTO } from './PalDTO';
import {
    PalDTOFromJSON,
    PalDTOFromJSONTyped,
    PalDTOToJSON,
    PalDTOToJSONTyped,
} from './PalDTO';

/**
 * 
 * @export
 * @interface NftInternalList
 */
export interface NftInternalList {
    /**
     * 
     * @type {string}
     * @memberof NftInternalList
     */
    cursor?: string | null;
    /**
     * 
     * @type {Array<PalDTO>}
     * @memberof NftInternalList
     */
    result?: Array<PalDTO> | null;
    /**
     * 
     * @type {number}
     * @memberof NftInternalList
     */
    totalCount?: number | null;
}

/**
 * Check if a given object implements the NftInternalList interface.
 */
export function instanceOfNftInternalList(value: object): value is NftInternalList {
    return true;
}

export function NftInternalListFromJSON(json: any): NftInternalList {
    return NftInternalListFromJSONTyped(json, false);
}

export function NftInternalListFromJSONTyped(json: any, ignoreDiscriminator: boolean): NftInternalList {
    if (json == null) {
        return json;
    }
    return {
        
        'cursor': json['cursor'] == null ? undefined : json['cursor'],
        'result': json['result'] == null ? undefined : ((json['result'] as Array<any>).map(PalDTOFromJSON)),
        'totalCount': json['total_count'] == null ? undefined : json['total_count'],
    };
}

export function NftInternalListToJSON(json: any): NftInternalList {
    return NftInternalListToJSONTyped(json, false);
}

export function NftInternalListToJSONTyped(value?: NftInternalList | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'cursor': value['cursor'],
        'result': value['result'] == null ? undefined : ((value['result'] as Array<any>).map(PalDTOToJSON)),
        'total_count': value['totalCount'],
    };
}

