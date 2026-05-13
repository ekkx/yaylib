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
 * @interface ChannelCommand
 */
export interface ChannelCommand {
    /**
     * 
     * @type {string}
     * @memberof ChannelCommand
     */
    command?: string | null;
    /**
     * 
     * @type {string}
     * @memberof ChannelCommand
     */
    identifier?: string | null;
}

/**
 * Check if a given object implements the ChannelCommand interface.
 */
export function instanceOfChannelCommand(value: object): value is ChannelCommand {
    return true;
}

export function ChannelCommandFromJSON(json: any): ChannelCommand {
    return ChannelCommandFromJSONTyped(json, false);
}

export function ChannelCommandFromJSONTyped(json: any, ignoreDiscriminator: boolean): ChannelCommand {
    if (json == null) {
        return json;
    }
    return {
        
        'command': json['command'] == null ? undefined : json['command'],
        'identifier': json['identifier'] == null ? undefined : json['identifier'],
    };
}

export function ChannelCommandToJSON(json: any): ChannelCommand {
    return ChannelCommandToJSONTyped(json, false);
}

export function ChannelCommandToJSONTyped(value?: ChannelCommand | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'command': value['command'],
        'identifier': value['identifier'],
    };
}

