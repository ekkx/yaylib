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
 * @interface InfuraResponse
 */
export interface InfuraResponse {
    /**
     * 
     * @type {number}
     * @memberof InfuraResponse
     */
    id?: number | null;
    /**
     * 
     * @type {any}
     * @memberof InfuraResponse
     */
    result?: any | null;
}

/**
 * Check if a given object implements the InfuraResponse interface.
 */
export function instanceOfInfuraResponse(value: object): value is InfuraResponse {
    return true;
}

export function InfuraResponseFromJSON(json: any): InfuraResponse {
    return InfuraResponseFromJSONTyped(json, false);
}

export function InfuraResponseFromJSONTyped(json: any, ignoreDiscriminator: boolean): InfuraResponse {
    if (json == null) {
        return json;
    }
    return {
        
        'id': json['id'] == null ? undefined : json['id'],
        'result': json['result'] == null ? undefined : json['result'],
    };
}

export function InfuraResponseToJSON(json: any): InfuraResponse {
    return InfuraResponseToJSONTyped(json, false);
}

export function InfuraResponseToJSONTyped(value?: InfuraResponse | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'id': value['id'],
        'result': value['result'],
    };
}

