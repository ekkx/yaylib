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
 * @interface PresignedUrlResponse
 */
export interface PresignedUrlResponse {
    /**
     * 
     * @type {string}
     * @memberof PresignedUrlResponse
     */
    presignedUrl?: string | null;
}

/**
 * Check if a given object implements the PresignedUrlResponse interface.
 */
export function instanceOfPresignedUrlResponse(value: object): value is PresignedUrlResponse {
    return true;
}

export function PresignedUrlResponseFromJSON(json: any): PresignedUrlResponse {
    return PresignedUrlResponseFromJSONTyped(json, false);
}

export function PresignedUrlResponseFromJSONTyped(json: any, ignoreDiscriminator: boolean): PresignedUrlResponse {
    if (json == null) {
        return json;
    }
    return {
        
        'presignedUrl': json['presigned_url'] == null ? undefined : json['presigned_url'],
    };
}

export function PresignedUrlResponseToJSON(json: any): PresignedUrlResponse {
    return PresignedUrlResponseToJSONTyped(json, false);
}

export function PresignedUrlResponseToJSONTyped(value?: PresignedUrlResponse | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'presigned_url': value['presignedUrl'],
    };
}

