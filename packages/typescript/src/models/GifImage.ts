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
 * @interface GifImage
 */
export interface GifImage {
    /**
     * 
     * @type {number}
     * @memberof GifImage
     */
    height?: number | null;
    /**
     * 
     * @type {number}
     * @memberof GifImage
     */
    id?: number | null;
    /**
     * 
     * @type {string}
     * @memberof GifImage
     */
    url?: string | null;
    /**
     * 
     * @type {number}
     * @memberof GifImage
     */
    width?: number | null;
}

/**
 * Check if a given object implements the GifImage interface.
 */
export function instanceOfGifImage(value: object): value is GifImage {
    return true;
}

export function GifImageFromJSON(json: any): GifImage {
    return GifImageFromJSONTyped(json, false);
}

export function GifImageFromJSONTyped(json: any, ignoreDiscriminator: boolean): GifImage {
    if (json == null) {
        return json;
    }
    return {
        
        'height': json['height'] == null ? undefined : json['height'],
        'id': json['id'] == null ? undefined : json['id'],
        'url': json['url'] == null ? undefined : json['url'],
        'width': json['width'] == null ? undefined : json['width'],
    };
}

export function GifImageToJSON(json: any): GifImage {
    return GifImageToJSONTyped(json, false);
}

export function GifImageToJSONTyped(value?: GifImage | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'height': value['height'],
        'id': value['id'],
        'url': value['url'],
        'width': value['width'],
    };
}

