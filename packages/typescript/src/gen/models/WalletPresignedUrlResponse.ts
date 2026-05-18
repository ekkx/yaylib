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
 * @interface WalletPresignedUrlResponse
 */
export interface WalletPresignedUrlResponse {
    /**
     * 
     * @type {string}
     * @memberof WalletPresignedUrlResponse
     */
    presignedUrl?: string | null;
}

/**
 * Check if a given object implements the WalletPresignedUrlResponse interface.
 */
export function instanceOfWalletPresignedUrlResponse(value: object): value is WalletPresignedUrlResponse {
    return true;
}

export function WalletPresignedUrlResponseFromJSON(json: any): WalletPresignedUrlResponse {
    return WalletPresignedUrlResponseFromJSONTyped(json, false);
}

export function WalletPresignedUrlResponseFromJSONTyped(json: any, ignoreDiscriminator: boolean): WalletPresignedUrlResponse {
    if (json == null) {
        return json;
    }
    return {
        
        'presignedUrl': json['presigned_url'] == null ? undefined : json['presigned_url'],
    };
}

export function WalletPresignedUrlResponseToJSON(json: any): WalletPresignedUrlResponse {
    return WalletPresignedUrlResponseToJSONTyped(json, false);
}

export function WalletPresignedUrlResponseToJSONTyped(value?: WalletPresignedUrlResponse | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'presigned_url': value['presignedUrl'],
    };
}

