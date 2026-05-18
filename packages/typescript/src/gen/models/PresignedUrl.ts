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
 * @interface PresignedUrl
 */
export interface PresignedUrl {
    /**
     * 
     * @type {string}
     * @memberof PresignedUrl
     */
    filename?: string | null;
    /**
     * 
     * @type {string}
     * @memberof PresignedUrl
     */
    url?: string | null;
}

/**
 * Check if a given object implements the PresignedUrl interface.
 */
export function instanceOfPresignedUrl(value: object): value is PresignedUrl {
    return true;
}

export function PresignedUrlFromJSON(json: any): PresignedUrl {
    return PresignedUrlFromJSONTyped(json, false);
}

export function PresignedUrlFromJSONTyped(json: any, ignoreDiscriminator: boolean): PresignedUrl {
    if (json == null) {
        return json;
    }
    return {
        
        'filename': json['filename'] == null ? undefined : json['filename'],
        'url': json['url'] == null ? undefined : json['url'],
    };
}

export function PresignedUrlToJSON(json: any): PresignedUrl {
    return PresignedUrlToJSONTyped(json, false);
}

export function PresignedUrlToJSONTyped(value?: PresignedUrl | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'filename': value['filename'],
        'url': value['url'],
    };
}

