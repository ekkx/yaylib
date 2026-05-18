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
/**
 * 
 * @export
 * @interface JSEvent
 */
export interface JSEvent {
    /**
     * 
     * @type {string}
     * @memberof JSEvent
     */
    args?: string | null;
    /**
     * 
     * @type {string}
     * @memberof JSEvent
     */
    name?: string | null;
}

/**
 * Check if a given object implements the JSEvent interface.
 */
export function instanceOfJSEvent(value: object): value is JSEvent {
    return true;
}

export function JSEventFromJSON(json: any): JSEvent {
    return JSEventFromJSONTyped(json, false);
}

export function JSEventFromJSONTyped(json: any, ignoreDiscriminator: boolean): JSEvent {
    if (json == null) {
        return json;
    }
    return {
        
        'args': json['args'] == null ? undefined : json['args'],
        'name': json['name'] == null ? undefined : json['name'],
    };
}

export function JSEventToJSON(json: any): JSEvent {
    return JSEventToJSONTyped(json, false);
}

export function JSEventToJSONTyped(value?: JSEvent | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'args': value['args'],
        'name': value['name'],
    };
}

