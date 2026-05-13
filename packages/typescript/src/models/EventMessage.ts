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
/**
 * 
 * @export
 * @interface EventMessage
 */
export interface EventMessage {
    /**
     * 
     * @type {object}
     * @memberof EventMessage
     */
    data?: object | null;
}

/**
 * Check if a given object implements the EventMessage interface.
 */
export function instanceOfEventMessage(value: object): value is EventMessage {
    return true;
}

export function EventMessageFromJSON(json: any): EventMessage {
    return EventMessageFromJSONTyped(json, false);
}

export function EventMessageFromJSONTyped(json: any, ignoreDiscriminator: boolean): EventMessage {
    if (json == null) {
        return json;
    }
    return {
        
        'data': json['data'] == null ? undefined : json['data'],
    };
}

export function EventMessageToJSON(json: any): EventMessage {
    return EventMessageToJSONTyped(json, false);
}

export function EventMessageToJSONTyped(value?: EventMessage | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'data': value['data'],
    };
}

