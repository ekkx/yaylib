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
import type { RealmGenre } from './RealmGenre';
import {
    RealmGenreFromJSON,
    RealmGenreFromJSONTyped,
    RealmGenreToJSON,
    RealmGenreToJSONTyped,
} from './RealmGenre';

/**
 * 
 * @export
 * @interface GenresResponse
 */
export interface GenresResponse {
    /**
     * 
     * @type {Array<RealmGenre>}
     * @memberof GenresResponse
     */
    genres?: Array<RealmGenre> | null;
    /**
     * 
     * @type {number}
     * @memberof GenresResponse
     */
    nextPageValue?: number | null;
}

/**
 * Check if a given object implements the GenresResponse interface.
 */
export function instanceOfGenresResponse(value: object): value is GenresResponse {
    return true;
}

export function GenresResponseFromJSON(json: any): GenresResponse {
    return GenresResponseFromJSONTyped(json, false);
}

export function GenresResponseFromJSONTyped(json: any, ignoreDiscriminator: boolean): GenresResponse {
    if (json == null) {
        return json;
    }
    return {
        
        'genres': json['genres'] == null ? undefined : ((json['genres'] as Array<any>).map(RealmGenreFromJSON)),
        'nextPageValue': json['next_page_value'] == null ? undefined : json['next_page_value'],
    };
}

export function GenresResponseToJSON(json: any): GenresResponse {
    return GenresResponseToJSONTyped(json, false);
}

export function GenresResponseToJSONTyped(value?: GenresResponse | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'genres': value['genres'] == null ? undefined : ((value['genres'] as Array<any>).map(RealmGenreToJSON)),
        'next_page_value': value['nextPageValue'],
    };
}

