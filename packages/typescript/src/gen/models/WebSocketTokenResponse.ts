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
 * @interface WebSocketTokenResponse
 */
export interface WebSocketTokenResponse {
    /**
     * 
     * @type {string}
     * @memberof WebSocketTokenResponse
     */
    token?: string | null;
}

/**
 * Check if a given object implements the WebSocketTokenResponse interface.
 */
export function instanceOfWebSocketTokenResponse(value: object): value is WebSocketTokenResponse {
    return true;
}

export function WebSocketTokenResponseFromJSON(json: any): WebSocketTokenResponse {
    return WebSocketTokenResponseFromJSONTyped(json, false);
}

export function WebSocketTokenResponseFromJSONTyped(json: any, ignoreDiscriminator: boolean): WebSocketTokenResponse {
    if (json == null) {
        return json;
    }
    return {
        
        'token': json['token'] == null ? undefined : json['token'],
    };
}

export function WebSocketTokenResponseToJSON(json: any): WebSocketTokenResponse {
    return WebSocketTokenResponseToJSONTyped(json, false);
}

export function WebSocketTokenResponseToJSONTyped(value?: WebSocketTokenResponse | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'token': value['token'],
    };
}

