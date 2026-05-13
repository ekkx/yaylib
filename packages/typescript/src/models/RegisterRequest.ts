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
 * @interface RegisterRequest
 */
export interface RegisterRequest {
    /**
     * 
     * @type {string}
     * @memberof RegisterRequest
     */
    signedDataType?: string | null;
    /**
     * 
     * @type {string}
     * @memberof RegisterRequest
     */
    walletAddress?: string | null;
}

/**
 * Check if a given object implements the RegisterRequest interface.
 */
export function instanceOfRegisterRequest(value: object): value is RegisterRequest {
    return true;
}

export function RegisterRequestFromJSON(json: any): RegisterRequest {
    return RegisterRequestFromJSONTyped(json, false);
}

export function RegisterRequestFromJSONTyped(json: any, ignoreDiscriminator: boolean): RegisterRequest {
    if (json == null) {
        return json;
    }
    return {
        
        'signedDataType': json['signed_data_type'] == null ? undefined : json['signed_data_type'],
        'walletAddress': json['wallet_address'] == null ? undefined : json['wallet_address'],
    };
}

export function RegisterRequestToJSON(json: any): RegisterRequest {
    return RegisterRequestToJSONTyped(json, false);
}

export function RegisterRequestToJSONTyped(value?: RegisterRequest | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'signed_data_type': value['signedDataType'],
        'wallet_address': value['walletAddress'],
    };
}

