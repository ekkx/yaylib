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
 * @interface WebSocketInteractor
 */
export interface WebSocketInteractor {
    /**
     * 
     * @type {object}
     * @memberof WebSocketInteractor
     */
    d?: object | null;
    /**
     * 
     * @type {object}
     * @memberof WebSocketInteractor
     */
    f?: object | null;
    /**
     * 
     * @type {object}
     * @memberof WebSocketInteractor
     */
    h?: object | null;
    /**
     * 
     * @type {object}
     * @memberof WebSocketInteractor
     */
    i?: object | null;
}

/**
 * Check if a given object implements the WebSocketInteractor interface.
 */
export function instanceOfWebSocketInteractor(value: object): value is WebSocketInteractor {
    return true;
}

export function WebSocketInteractorFromJSON(json: any): WebSocketInteractor {
    return WebSocketInteractorFromJSONTyped(json, false);
}

export function WebSocketInteractorFromJSONTyped(json: any, ignoreDiscriminator: boolean): WebSocketInteractor {
    if (json == null) {
        return json;
    }
    return {
        
        'd': json['d'] == null ? undefined : json['d'],
        'f': json['f'] == null ? undefined : json['f'],
        'h': json['h'] == null ? undefined : json['h'],
        'i': json['i'] == null ? undefined : json['i'],
    };
}

export function WebSocketInteractorToJSON(json: any): WebSocketInteractor {
    return WebSocketInteractorToJSONTyped(json, false);
}

export function WebSocketInteractorToJSONTyped(value?: WebSocketInteractor | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'd': value['d'],
        'f': value['f'],
        'h': value['h'],
        'i': value['i'],
    };
}

