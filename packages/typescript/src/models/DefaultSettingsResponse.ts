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
import type { TimelineSettings } from './TimelineSettings';
import {
    TimelineSettingsFromJSON,
    TimelineSettingsFromJSONTyped,
    TimelineSettingsToJSON,
    TimelineSettingsToJSONTyped,
} from './TimelineSettings';

/**
 * 
 * @export
 * @interface DefaultSettingsResponse
 */
export interface DefaultSettingsResponse {
    /**
     * 
     * @type {TimelineSettings}
     * @memberof DefaultSettingsResponse
     */
    timelineSettings?: TimelineSettings | null;
}

/**
 * Check if a given object implements the DefaultSettingsResponse interface.
 */
export function instanceOfDefaultSettingsResponse(value: object): value is DefaultSettingsResponse {
    return true;
}

export function DefaultSettingsResponseFromJSON(json: any): DefaultSettingsResponse {
    return DefaultSettingsResponseFromJSONTyped(json, false);
}

export function DefaultSettingsResponseFromJSONTyped(json: any, ignoreDiscriminator: boolean): DefaultSettingsResponse {
    if (json == null) {
        return json;
    }
    return {
        
        'timelineSettings': json['timeline_settings'] == null ? undefined : TimelineSettingsFromJSON(json['timeline_settings']),
    };
}

export function DefaultSettingsResponseToJSON(json: any): DefaultSettingsResponse {
    return DefaultSettingsResponseToJSONTyped(json, false);
}

export function DefaultSettingsResponseToJSONTyped(value?: DefaultSettingsResponse | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'timeline_settings': TimelineSettingsToJSON(value['timelineSettings']),
    };
}

