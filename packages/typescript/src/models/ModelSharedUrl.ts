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
 * @interface ModelSharedUrl
 */
export interface ModelSharedUrl {
    /**
     * 
     * @type {string}
     * @memberof ModelSharedUrl
     */
    description?: string | null;
    /**
     * 
     * @type {string}
     * @memberof ModelSharedUrl
     */
    imageUrl?: string | null;
    /**
     * 
     * @type {string}
     * @memberof ModelSharedUrl
     */
    title?: string | null;
    /**
     * 
     * @type {string}
     * @memberof ModelSharedUrl
     */
    url?: string | null;
}

/**
 * Check if a given object implements the ModelSharedUrl interface.
 */
export function instanceOfModelSharedUrl(value: object): value is ModelSharedUrl {
    return true;
}

export function ModelSharedUrlFromJSON(json: any): ModelSharedUrl {
    return ModelSharedUrlFromJSONTyped(json, false);
}

export function ModelSharedUrlFromJSONTyped(json: any, ignoreDiscriminator: boolean): ModelSharedUrl {
    if (json == null) {
        return json;
    }
    return {
        
        'description': json['description'] == null ? undefined : json['description'],
        'imageUrl': json['image_url'] == null ? undefined : json['image_url'],
        'title': json['title'] == null ? undefined : json['title'],
        'url': json['url'] == null ? undefined : json['url'],
    };
}

export function ModelSharedUrlToJSON(json: any): ModelSharedUrl {
    return ModelSharedUrlToJSONTyped(json, false);
}

export function ModelSharedUrlToJSONTyped(value?: ModelSharedUrl | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'description': value['description'],
        'image_url': value['imageUrl'],
        'title': value['title'],
        'url': value['url'],
    };
}

