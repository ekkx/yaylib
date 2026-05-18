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
 * @interface InfuraRequest
 */
export interface InfuraRequest {
    /**
     * 
     * @type {number}
     * @memberof InfuraRequest
     */
    id?: number | null;
    /**
     * 
     * @type {string}
     * @memberof InfuraRequest
     */
    jsonrpc?: string | null;
    /**
     * 
     * @type {string}
     * @memberof InfuraRequest
     */
    method?: string | null;
    /**
     * 
     * @type {Array<any>}
     * @memberof InfuraRequest
     */
    params?: Array<any> | null;
}

/**
 * Check if a given object implements the InfuraRequest interface.
 */
export function instanceOfInfuraRequest(value: object): value is InfuraRequest {
    return true;
}

export function InfuraRequestFromJSON(json: any): InfuraRequest {
    return InfuraRequestFromJSONTyped(json, false);
}

export function InfuraRequestFromJSONTyped(json: any, ignoreDiscriminator: boolean): InfuraRequest {
    if (json == null) {
        return json;
    }
    return {
        
        'id': json['id'] == null ? undefined : json['id'],
        'jsonrpc': json['jsonrpc'] == null ? undefined : json['jsonrpc'],
        'method': json['method'] == null ? undefined : json['method'],
        'params': json['params'] == null ? undefined : json['params'],
    };
}

export function InfuraRequestToJSON(json: any): InfuraRequest {
    return InfuraRequestToJSONTyped(json, false);
}

export function InfuraRequestToJSONTyped(value?: InfuraRequest | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'id': value['id'],
        'jsonrpc': value['jsonrpc'],
        'method': value['method'],
        'params': value['params'],
    };
}

