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
 * @interface ChannelTypedMessage
 */
export interface ChannelTypedMessage {
    /**
     * 
     * @type {string}
     * @memberof ChannelTypedMessage
     */
    type?: string | null;
}

/**
 * Check if a given object implements the ChannelTypedMessage interface.
 */
export function instanceOfChannelTypedMessage(value: object): value is ChannelTypedMessage {
    return true;
}

export function ChannelTypedMessageFromJSON(json: any): ChannelTypedMessage {
    return ChannelTypedMessageFromJSONTyped(json, false);
}

export function ChannelTypedMessageFromJSONTyped(json: any, ignoreDiscriminator: boolean): ChannelTypedMessage {
    if (json == null) {
        return json;
    }
    return {
        
        'type': json['type'] == null ? undefined : json['type'],
    };
}

export function ChannelTypedMessageToJSON(json: any): ChannelTypedMessage {
    return ChannelTypedMessageToJSONTyped(json, false);
}

export function ChannelTypedMessageToJSONTyped(value?: ChannelTypedMessage | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'type': value['type'],
    };
}

