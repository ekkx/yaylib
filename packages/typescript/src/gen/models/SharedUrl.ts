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
 * @interface SharedUrl
 */
export interface SharedUrl {
    /**
     * 
     * @type {string}
     * @memberof SharedUrl
     */
    description?: string | null;
    /**
     * 
     * @type {string}
     * @memberof SharedUrl
     */
    imageUrl?: string | null;
    /**
     * 
     * @type {string}
     * @memberof SharedUrl
     */
    title?: string | null;
    /**
     * 
     * @type {string}
     * @memberof SharedUrl
     */
    url?: string | null;
}

/**
 * Check if a given object implements the SharedUrl interface.
 */
export function instanceOfSharedUrl(value: object): value is SharedUrl {
    return true;
}

export function SharedUrlFromJSON(json: any): SharedUrl {
    return SharedUrlFromJSONTyped(json, false);
}

export function SharedUrlFromJSONTyped(json: any, ignoreDiscriminator: boolean): SharedUrl {
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

export function SharedUrlToJSON(json: any): SharedUrl {
    return SharedUrlToJSONTyped(json, false);
}

export function SharedUrlToJSONTyped(value?: SharedUrl | null, ignoreDiscriminator: boolean = false): any {
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

