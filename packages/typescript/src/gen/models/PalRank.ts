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
import type { Pal } from './Pal.js';
import {
    PalFromJSON,
    PalFromJSONTyped,
    PalToJSON,
    PalToJSONTyped,
} from './Pal.js';

/**
 * 
 * @export
 * @interface PalRank
 */
export interface PalRank {
    /**
     * 
     * @type {boolean}
     * @memberof PalRank
     */
    own?: boolean | null;
    /**
     * 
     * @type {Pal}
     * @memberof PalRank
     */
    pal?: Pal | null;
    /**
     * 
     * @type {number}
     * @memberof PalRank
     */
    place?: number | null;
    /**
     * 
     * @type {string}
     * @memberof PalRank
     */
    yayUserName?: string | null;
}

/**
 * Check if a given object implements the PalRank interface.
 */
export function instanceOfPalRank(value: object): value is PalRank {
    return true;
}

export function PalRankFromJSON(json: any): PalRank {
    return PalRankFromJSONTyped(json, false);
}

export function PalRankFromJSONTyped(json: any, ignoreDiscriminator: boolean): PalRank {
    if (json == null) {
        return json;
    }
    return {
        
        'own': json['own'] == null ? undefined : json['own'],
        'pal': json['pal'] == null ? undefined : PalFromJSON(json['pal']),
        'place': json['place'] == null ? undefined : json['place'],
        'yayUserName': json['yay_user_name'] == null ? undefined : json['yay_user_name'],
    };
}

export function PalRankToJSON(json: any): PalRank {
    return PalRankToJSONTyped(json, false);
}

export function PalRankToJSONTyped(value?: PalRank | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'own': value['own'],
        'pal': PalToJSON(value['pal']),
        'place': value['place'],
        'yay_user_name': value['yayUserName'],
    };
}

