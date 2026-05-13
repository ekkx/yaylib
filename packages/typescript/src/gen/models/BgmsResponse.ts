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
import type { Bgm } from './Bgm';
import {
    BgmFromJSON,
    BgmFromJSONTyped,
    BgmToJSON,
    BgmToJSONTyped,
} from './Bgm';

/**
 * 
 * @export
 * @interface BgmsResponse
 */
export interface BgmsResponse {
    /**
     * 
     * @type {Array<Bgm>}
     * @memberof BgmsResponse
     */
    bgm?: Array<Bgm> | null;
}

/**
 * Check if a given object implements the BgmsResponse interface.
 */
export function instanceOfBgmsResponse(value: object): value is BgmsResponse {
    return true;
}

export function BgmsResponseFromJSON(json: any): BgmsResponse {
    return BgmsResponseFromJSONTyped(json, false);
}

export function BgmsResponseFromJSONTyped(json: any, ignoreDiscriminator: boolean): BgmsResponse {
    if (json == null) {
        return json;
    }
    return {
        
        'bgm': json['bgm'] == null ? undefined : ((json['bgm'] as Array<any>).map(BgmFromJSON)),
    };
}

export function BgmsResponseToJSON(json: any): BgmsResponse {
    return BgmsResponseToJSONTyped(json, false);
}

export function BgmsResponseToJSONTyped(value?: BgmsResponse | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'bgm': value['bgm'] == null ? undefined : ((value['bgm'] as Array<any>).map(BgmToJSON)),
    };
}

