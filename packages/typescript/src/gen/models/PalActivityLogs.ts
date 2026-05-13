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
import type { PalActivityLog } from './PalActivityLog';
import {
    PalActivityLogFromJSON,
    PalActivityLogFromJSONTyped,
    PalActivityLogToJSON,
    PalActivityLogToJSONTyped,
} from './PalActivityLog';

/**
 * 
 * @export
 * @interface PalActivityLogs
 */
export interface PalActivityLogs {
    /**
     * 
     * @type {Array<PalActivityLog>}
     * @memberof PalActivityLogs
     */
    result?: Array<PalActivityLog> | null;
}

/**
 * Check if a given object implements the PalActivityLogs interface.
 */
export function instanceOfPalActivityLogs(value: object): value is PalActivityLogs {
    return true;
}

export function PalActivityLogsFromJSON(json: any): PalActivityLogs {
    return PalActivityLogsFromJSONTyped(json, false);
}

export function PalActivityLogsFromJSONTyped(json: any, ignoreDiscriminator: boolean): PalActivityLogs {
    if (json == null) {
        return json;
    }
    return {
        
        'result': json['result'] == null ? undefined : ((json['result'] as Array<any>).map(PalActivityLogFromJSON)),
    };
}

export function PalActivityLogsToJSON(json: any): PalActivityLogs {
    return PalActivityLogsToJSONTyped(json, false);
}

export function PalActivityLogsToJSONTyped(value?: PalActivityLogs | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'result': value['result'] == null ? undefined : ((value['result'] as Array<any>).map(PalActivityLogToJSON)),
    };
}

