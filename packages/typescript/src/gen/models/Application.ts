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
import type { Config } from './Config.js';
import {
    ConfigFromJSON,
    ConfigFromJSONTyped,
    ConfigToJSON,
    ConfigToJSONTyped,
} from './Config.js';

/**
 * 
 * @export
 * @interface Application
 */
export interface Application {
    /**
     * 
     * @type {Config}
     * @memberof Application
     */
    settings?: Config | null;
}

/**
 * Check if a given object implements the Application interface.
 */
export function instanceOfApplication(value: object): value is Application {
    return true;
}

export function ApplicationFromJSON(json: any): Application {
    return ApplicationFromJSONTyped(json, false);
}

export function ApplicationFromJSONTyped(json: any, ignoreDiscriminator: boolean): Application {
    if (json == null) {
        return json;
    }
    return {
        
        'settings': json['settings'] == null ? undefined : ConfigFromJSON(json['settings']),
    };
}

export function ApplicationToJSON(json: any): Application {
    return ApplicationToJSONTyped(json, false);
}

export function ApplicationToJSONTyped(value?: Application | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'settings': ConfigToJSON(value['settings']),
    };
}

