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
 * @interface TotalChatRequestResponse
 */
export interface TotalChatRequestResponse {
    /**
     * 
     * @type {number}
     * @memberof TotalChatRequestResponse
     */
    total?: number | null;
}

/**
 * Check if a given object implements the TotalChatRequestResponse interface.
 */
export function instanceOfTotalChatRequestResponse(value: object): value is TotalChatRequestResponse {
    return true;
}

export function TotalChatRequestResponseFromJSON(json: any): TotalChatRequestResponse {
    return TotalChatRequestResponseFromJSONTyped(json, false);
}

export function TotalChatRequestResponseFromJSONTyped(json: any, ignoreDiscriminator: boolean): TotalChatRequestResponse {
    if (json == null) {
        return json;
    }
    return {
        
        'total': json['total'] == null ? undefined : json['total'],
    };
}

export function TotalChatRequestResponseToJSON(json: any): TotalChatRequestResponse {
    return TotalChatRequestResponseToJSONTyped(json, false);
}

export function TotalChatRequestResponseToJSONTyped(value?: TotalChatRequestResponse | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'total': value['total'],
    };
}

