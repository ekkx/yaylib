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
 * @interface Callback
 */
export interface Callback {
    /**
     * 
     * @type {object}
     * @memberof Callback
     */
    callbackMethod?: object | null;
    /**
     * 
     * @type {number}
     * @memberof Callback
     */
    id?: number | null;
    /**
     * 
     * @type {string}
     * @memberof Callback
     */
    responseString?: string | null;
}

/**
 * Check if a given object implements the Callback interface.
 */
export function instanceOfCallback(value: object): value is Callback {
    return true;
}

export function CallbackFromJSON(json: any): Callback {
    return CallbackFromJSONTyped(json, false);
}

export function CallbackFromJSONTyped(json: any, ignoreDiscriminator: boolean): Callback {
    if (json == null) {
        return json;
    }
    return {
        
        'callbackMethod': json['callback_method'] == null ? undefined : json['callback_method'],
        'id': json['id'] == null ? undefined : json['id'],
        'responseString': json['response_string'] == null ? undefined : json['response_string'],
    };
}

export function CallbackToJSON(json: any): Callback {
    return CallbackToJSONTyped(json, false);
}

export function CallbackToJSONTyped(value?: Callback | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'callback_method': value['callbackMethod'],
        'id': value['id'],
        'response_string': value['responseString'],
    };
}

