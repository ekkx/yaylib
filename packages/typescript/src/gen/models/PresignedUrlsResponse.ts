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
import type { PresignedUrl } from './PresignedUrl.js';
import {
    PresignedUrlFromJSON,
    PresignedUrlFromJSONTyped,
    PresignedUrlToJSON,
    PresignedUrlToJSONTyped,
} from './PresignedUrl.js';

/**
 * 
 * @export
 * @interface PresignedUrlsResponse
 */
export interface PresignedUrlsResponse {
    /**
     * 
     * @type {Array<PresignedUrl>}
     * @memberof PresignedUrlsResponse
     */
    presignedUrls?: Array<PresignedUrl> | null;
}

/**
 * Check if a given object implements the PresignedUrlsResponse interface.
 */
export function instanceOfPresignedUrlsResponse(value: object): value is PresignedUrlsResponse {
    return true;
}

export function PresignedUrlsResponseFromJSON(json: any): PresignedUrlsResponse {
    return PresignedUrlsResponseFromJSONTyped(json, false);
}

export function PresignedUrlsResponseFromJSONTyped(json: any, ignoreDiscriminator: boolean): PresignedUrlsResponse {
    if (json == null) {
        return json;
    }
    return {
        
        'presignedUrls': json['presigned_urls'] == null ? undefined : ((json['presigned_urls'] as Array<any>).map(PresignedUrlFromJSON)),
    };
}

export function PresignedUrlsResponseToJSON(json: any): PresignedUrlsResponse {
    return PresignedUrlsResponseToJSONTyped(json, false);
}

export function PresignedUrlsResponseToJSONTyped(value?: PresignedUrlsResponse | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'presigned_urls': value['presignedUrls'] == null ? undefined : ((value['presignedUrls'] as Array<any>).map(PresignedUrlToJSON)),
    };
}

