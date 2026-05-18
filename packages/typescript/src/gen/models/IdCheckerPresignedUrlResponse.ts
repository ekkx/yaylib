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
 * @interface IdCheckerPresignedUrlResponse
 */
export interface IdCheckerPresignedUrlResponse {
    /**
     * 
     * @type {string}
     * @memberof IdCheckerPresignedUrlResponse
     */
    presignedUrl?: string | null;
}

/**
 * Check if a given object implements the IdCheckerPresignedUrlResponse interface.
 */
export function instanceOfIdCheckerPresignedUrlResponse(value: object): value is IdCheckerPresignedUrlResponse {
    return true;
}

export function IdCheckerPresignedUrlResponseFromJSON(json: any): IdCheckerPresignedUrlResponse {
    return IdCheckerPresignedUrlResponseFromJSONTyped(json, false);
}

export function IdCheckerPresignedUrlResponseFromJSONTyped(json: any, ignoreDiscriminator: boolean): IdCheckerPresignedUrlResponse {
    if (json == null) {
        return json;
    }
    return {
        
        'presignedUrl': json['presigned_url'] == null ? undefined : json['presigned_url'],
    };
}

export function IdCheckerPresignedUrlResponseToJSON(json: any): IdCheckerPresignedUrlResponse {
    return IdCheckerPresignedUrlResponseToJSONTyped(json, false);
}

export function IdCheckerPresignedUrlResponseToJSONTyped(value?: IdCheckerPresignedUrlResponse | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'presigned_url': value['presignedUrl'],
    };
}

