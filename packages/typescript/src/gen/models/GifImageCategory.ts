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
import type { GifImage } from './GifImage';
import {
    GifImageFromJSON,
    GifImageFromJSONTyped,
    GifImageToJSON,
    GifImageToJSONTyped,
} from './GifImage';

/**
 * 
 * @export
 * @interface GifImageCategory
 */
export interface GifImageCategory {
    /**
     * 
     * @type {Array<GifImage>}
     * @memberof GifImageCategory
     */
    gifs?: Array<GifImage> | null;
    /**
     * 
     * @type {number}
     * @memberof GifImageCategory
     */
    id?: number | null;
    /**
     * 
     * @type {string}
     * @memberof GifImageCategory
     */
    language?: string | null;
    /**
     * 
     * @type {string}
     * @memberof GifImageCategory
     */
    name?: string | null;
}

/**
 * Check if a given object implements the GifImageCategory interface.
 */
export function instanceOfGifImageCategory(value: object): value is GifImageCategory {
    return true;
}

export function GifImageCategoryFromJSON(json: any): GifImageCategory {
    return GifImageCategoryFromJSONTyped(json, false);
}

export function GifImageCategoryFromJSONTyped(json: any, ignoreDiscriminator: boolean): GifImageCategory {
    if (json == null) {
        return json;
    }
    return {
        
        'gifs': json['gifs'] == null ? undefined : ((json['gifs'] as Array<any>).map(GifImageFromJSON)),
        'id': json['id'] == null ? undefined : json['id'],
        'language': json['language'] == null ? undefined : json['language'],
        'name': json['name'] == null ? undefined : json['name'],
    };
}

export function GifImageCategoryToJSON(json: any): GifImageCategory {
    return GifImageCategoryToJSONTyped(json, false);
}

export function GifImageCategoryToJSONTyped(value?: GifImageCategory | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'gifs': value['gifs'] == null ? undefined : ((value['gifs'] as Array<any>).map(GifImageToJSON)),
        'id': value['id'],
        'language': value['language'],
        'name': value['name'],
    };
}

