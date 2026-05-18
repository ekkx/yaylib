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
import type { PalActivityLogDetails } from './PalActivityLogDetails.js';
import {
    PalActivityLogDetailsFromJSON,
    PalActivityLogDetailsFromJSONTyped,
    PalActivityLogDetailsToJSON,
    PalActivityLogDetailsToJSONTyped,
} from './PalActivityLogDetails.js';

/**
 * 
 * @export
 * @interface PalActivityLog
 */
export interface PalActivityLog {
    /**
     * 
     * @type {number}
     * @memberof PalActivityLog
     */
    createdAtSeconds?: number | null;
    /**
     * 
     * @type {PalActivityLogDetails}
     * @memberof PalActivityLog
     */
    details?: PalActivityLogDetails | null;
    /**
     * 
     * @type {number}
     * @memberof PalActivityLog
     */
    id?: number | null;
    /**
     * 
     * @type {object}
     * @memberof PalActivityLog
     */
    type?: object | null;
}

/**
 * Check if a given object implements the PalActivityLog interface.
 */
export function instanceOfPalActivityLog(value: object): value is PalActivityLog {
    return true;
}

export function PalActivityLogFromJSON(json: any): PalActivityLog {
    return PalActivityLogFromJSONTyped(json, false);
}

export function PalActivityLogFromJSONTyped(json: any, ignoreDiscriminator: boolean): PalActivityLog {
    if (json == null) {
        return json;
    }
    return {
        
        'createdAtSeconds': json['created_at_seconds'] == null ? undefined : json['created_at_seconds'],
        'details': json['details'] == null ? undefined : PalActivityLogDetailsFromJSON(json['details']),
        'id': json['id'] == null ? undefined : json['id'],
        'type': json['type'] == null ? undefined : json['type'],
    };
}

export function PalActivityLogToJSON(json: any): PalActivityLog {
    return PalActivityLogToJSONTyped(json, false);
}

export function PalActivityLogToJSONTyped(value?: PalActivityLog | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'created_at_seconds': value['createdAtSeconds'],
        'details': PalActivityLogDetailsToJSON(value['details']),
        'id': value['id'],
        'type': value['type'],
    };
}

