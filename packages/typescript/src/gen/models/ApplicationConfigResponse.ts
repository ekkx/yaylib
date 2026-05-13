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
import type { Application } from './Application';
import {
    ApplicationFromJSON,
    ApplicationFromJSONTyped,
    ApplicationToJSON,
    ApplicationToJSONTyped,
} from './Application';

/**
 * 
 * @export
 * @interface ApplicationConfigResponse
 */
export interface ApplicationConfigResponse {
    /**
     * 
     * @type {Application}
     * @memberof ApplicationConfigResponse
     */
    app?: Application | null;
}

/**
 * Check if a given object implements the ApplicationConfigResponse interface.
 */
export function instanceOfApplicationConfigResponse(value: object): value is ApplicationConfigResponse {
    return true;
}

export function ApplicationConfigResponseFromJSON(json: any): ApplicationConfigResponse {
    return ApplicationConfigResponseFromJSONTyped(json, false);
}

export function ApplicationConfigResponseFromJSONTyped(json: any, ignoreDiscriminator: boolean): ApplicationConfigResponse {
    if (json == null) {
        return json;
    }
    return {
        
        'app': json['app'] == null ? undefined : ApplicationFromJSON(json['app']),
    };
}

export function ApplicationConfigResponseToJSON(json: any): ApplicationConfigResponse {
    return ApplicationConfigResponseToJSONTyped(json, false);
}

export function ApplicationConfigResponseToJSONTyped(value?: ApplicationConfigResponse | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'app': ApplicationToJSON(value['app']),
    };
}

