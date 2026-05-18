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
import type { Activity } from './Activity.js';
import {
    ActivityFromJSON,
    ActivityFromJSONTyped,
    ActivityToJSON,
    ActivityToJSONTyped,
} from './Activity.js';

/**
 * 
 * @export
 * @interface ActivitiesResponse
 */
export interface ActivitiesResponse {
    /**
     * 
     * @type {Array<Activity>}
     * @memberof ActivitiesResponse
     */
    activities?: Array<Activity> | null;
    /**
     * 
     * @type {number}
     * @memberof ActivitiesResponse
     */
    lastTimestamp?: number | null;
}

/**
 * Check if a given object implements the ActivitiesResponse interface.
 */
export function instanceOfActivitiesResponse(value: object): value is ActivitiesResponse {
    return true;
}

export function ActivitiesResponseFromJSON(json: any): ActivitiesResponse {
    return ActivitiesResponseFromJSONTyped(json, false);
}

export function ActivitiesResponseFromJSONTyped(json: any, ignoreDiscriminator: boolean): ActivitiesResponse {
    if (json == null) {
        return json;
    }
    return {
        
        'activities': json['activities'] == null ? undefined : ((json['activities'] as Array<any>).map(ActivityFromJSON)),
        'lastTimestamp': json['last_timestamp'] == null ? undefined : json['last_timestamp'],
    };
}

export function ActivitiesResponseToJSON(json: any): ActivitiesResponse {
    return ActivitiesResponseToJSONTyped(json, false);
}

export function ActivitiesResponseToJSONTyped(value?: ActivitiesResponse | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'activities': value['activities'] == null ? undefined : ((value['activities'] as Array<any>).map(ActivityToJSON)),
        'last_timestamp': value['lastTimestamp'],
    };
}

