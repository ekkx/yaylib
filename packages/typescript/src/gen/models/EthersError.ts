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
 * @interface EthersError
 */
export interface EthersError {
    /**
     * 
     * @type {string}
     * @memberof EthersError
     */
    error?: string | null;
}

/**
 * Check if a given object implements the EthersError interface.
 */
export function instanceOfEthersError(value: object): value is EthersError {
    return true;
}

export function EthersErrorFromJSON(json: any): EthersError {
    return EthersErrorFromJSONTyped(json, false);
}

export function EthersErrorFromJSONTyped(json: any, ignoreDiscriminator: boolean): EthersError {
    if (json == null) {
        return json;
    }
    return {
        
        'error': json['error'] == null ? undefined : json['error'],
    };
}

export function EthersErrorToJSON(json: any): EthersError {
    return EthersErrorToJSONTyped(json, false);
}

export function EthersErrorToJSONTyped(value?: EthersError | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'error': value['error'],
    };
}

