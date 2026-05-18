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
import type { Settings } from './Settings.js';
import {
    SettingsFromJSON,
    SettingsFromJSONTyped,
    SettingsToJSON,
    SettingsToJSONTyped,
} from './Settings.js';

/**
 * 
 * @export
 * @interface AdditionalSettingsResponse
 */
export interface AdditionalSettingsResponse {
    /**
     * 
     * @type {Settings}
     * @memberof AdditionalSettingsResponse
     */
    settings?: Settings | null;
}

/**
 * Check if a given object implements the AdditionalSettingsResponse interface.
 */
export function instanceOfAdditionalSettingsResponse(value: object): value is AdditionalSettingsResponse {
    return true;
}

export function AdditionalSettingsResponseFromJSON(json: any): AdditionalSettingsResponse {
    return AdditionalSettingsResponseFromJSONTyped(json, false);
}

export function AdditionalSettingsResponseFromJSONTyped(json: any, ignoreDiscriminator: boolean): AdditionalSettingsResponse {
    if (json == null) {
        return json;
    }
    return {
        
        'settings': json['settings'] == null ? undefined : SettingsFromJSON(json['settings']),
    };
}

export function AdditionalSettingsResponseToJSON(json: any): AdditionalSettingsResponse {
    return AdditionalSettingsResponseToJSONTyped(json, false);
}

export function AdditionalSettingsResponseToJSONTyped(value?: AdditionalSettingsResponse | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'settings': SettingsToJSON(value['settings']),
    };
}

