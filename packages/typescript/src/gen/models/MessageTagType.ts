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
 * @interface MessageTagType
 */
export interface MessageTagType {
    /**
     * 
     * @type {string}
     * @memberof MessageTagType
     */
    apiValue?: string | null;
}

/**
 * Check if a given object implements the MessageTagType interface.
 */
export function instanceOfMessageTagType(value: object): value is MessageTagType {
    return true;
}

export function MessageTagTypeFromJSON(json: any): MessageTagType {
    return MessageTagTypeFromJSONTyped(json, false);
}

export function MessageTagTypeFromJSONTyped(json: any, ignoreDiscriminator: boolean): MessageTagType {
    if (json == null) {
        return json;
    }
    return {
        
        'apiValue': json['api_value'] == null ? undefined : json['api_value'],
    };
}

export function MessageTagTypeToJSON(json: any): MessageTagType {
    return MessageTagTypeToJSONTyped(json, false);
}

export function MessageTagTypeToJSONTyped(value?: MessageTagType | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'api_value': value['apiValue'],
    };
}

