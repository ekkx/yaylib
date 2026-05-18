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
import type { PendingEmplActivityDetails } from './PendingEmplActivityDetails.js';
import {
    PendingEmplActivityDetailsFromJSON,
    PendingEmplActivityDetailsFromJSONTyped,
    PendingEmplActivityDetailsToJSON,
    PendingEmplActivityDetailsToJSONTyped,
} from './PendingEmplActivityDetails.js';

/**
 * 
 * @export
 * @interface PendingEmplActivity
 */
export interface PendingEmplActivity {
    /**
     * 
     * @type {object}
     * @memberof PendingEmplActivity
     */
    activity?: object | null;
    /**
     * 
     * @type {PendingEmplActivityDetails}
     * @memberof PendingEmplActivity
     */
    details?: PendingEmplActivityDetails | null;
}

/**
 * Check if a given object implements the PendingEmplActivity interface.
 */
export function instanceOfPendingEmplActivity(value: object): value is PendingEmplActivity {
    return true;
}

export function PendingEmplActivityFromJSON(json: any): PendingEmplActivity {
    return PendingEmplActivityFromJSONTyped(json, false);
}

export function PendingEmplActivityFromJSONTyped(json: any, ignoreDiscriminator: boolean): PendingEmplActivity {
    if (json == null) {
        return json;
    }
    return {
        
        'activity': json['activity'] == null ? undefined : json['activity'],
        'details': json['details'] == null ? undefined : PendingEmplActivityDetailsFromJSON(json['details']),
    };
}

export function PendingEmplActivityToJSON(json: any): PendingEmplActivity {
    return PendingEmplActivityToJSONTyped(json, false);
}

export function PendingEmplActivityToJSONTyped(value?: PendingEmplActivity | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'activity': value['activity'],
        'details': PendingEmplActivityDetailsToJSON(value['details']),
    };
}

