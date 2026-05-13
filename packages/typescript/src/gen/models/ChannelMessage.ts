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
import type { EventMessage } from './EventMessage';
import {
    EventMessageFromJSON,
    EventMessageFromJSONTyped,
    EventMessageToJSON,
    EventMessageToJSONTyped,
} from './EventMessage';

/**
 * 
 * @export
 * @interface ChannelMessage
 */
export interface ChannelMessage {
    /**
     * 
     * @type {string}
     * @memberof ChannelMessage
     */
    identifier?: string | null;
    /**
     * 
     * @type {EventMessage}
     * @memberof ChannelMessage
     */
    message?: EventMessage | null;
    /**
     * 
     * @type {string}
     * @memberof ChannelMessage
     */
    type?: string | null;
}

/**
 * Check if a given object implements the ChannelMessage interface.
 */
export function instanceOfChannelMessage(value: object): value is ChannelMessage {
    return true;
}

export function ChannelMessageFromJSON(json: any): ChannelMessage {
    return ChannelMessageFromJSONTyped(json, false);
}

export function ChannelMessageFromJSONTyped(json: any, ignoreDiscriminator: boolean): ChannelMessage {
    if (json == null) {
        return json;
    }
    return {
        
        'identifier': json['identifier'] == null ? undefined : json['identifier'],
        'message': json['message'] == null ? undefined : EventMessageFromJSON(json['message']),
        'type': json['type'] == null ? undefined : json['type'],
    };
}

export function ChannelMessageToJSON(json: any): ChannelMessage {
    return ChannelMessageToJSONTyped(json, false);
}

export function ChannelMessageToJSONTyped(value?: ChannelMessage | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'identifier': value['identifier'],
        'message': EventMessageToJSON(value['message']),
        'type': value['type'],
    };
}

