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
import type { GifImageCategory } from './GifImageCategory';
import {
    GifImageCategoryFromJSON,
    GifImageCategoryFromJSONTyped,
    GifImageCategoryToJSON,
    GifImageCategoryToJSONTyped,
} from './GifImageCategory';

/**
 * 
 * @export
 * @interface GifsDataResponse
 */
export interface GifsDataResponse {
    /**
     * 
     * @type {Array<GifImageCategory>}
     * @memberof GifsDataResponse
     */
    gifCategories?: Array<GifImageCategory> | null;
}

/**
 * Check if a given object implements the GifsDataResponse interface.
 */
export function instanceOfGifsDataResponse(value: object): value is GifsDataResponse {
    return true;
}

export function GifsDataResponseFromJSON(json: any): GifsDataResponse {
    return GifsDataResponseFromJSONTyped(json, false);
}

export function GifsDataResponseFromJSONTyped(json: any, ignoreDiscriminator: boolean): GifsDataResponse {
    if (json == null) {
        return json;
    }
    return {
        
        'gifCategories': json['gif_categories'] == null ? undefined : ((json['gif_categories'] as Array<any>).map(GifImageCategoryFromJSON)),
    };
}

export function GifsDataResponseToJSON(json: any): GifsDataResponse {
    return GifsDataResponseToJSONTyped(json, false);
}

export function GifsDataResponseToJSONTyped(value?: GifsDataResponse | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'gif_categories': value['gifCategories'] == null ? undefined : ((value['gifCategories'] as Array<any>).map(GifImageCategoryToJSON)),
    };
}

