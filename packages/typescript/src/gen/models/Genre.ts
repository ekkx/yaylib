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
 * @interface Genre
 */
export interface Genre {
    /**
     * 
     * @type {string}
     * @memberof Genre
     */
    iconUrl?: string | null;
    /**
     * 
     * @type {number}
     * @memberof Genre
     */
    id?: number | null;
    /**
     * 
     * @type {string}
     * @memberof Genre
     */
    title?: string | null;
    /**
     * 
     * @type {string}
     * @memberof Genre
     */
    type?: string | null;
}

/**
 * Check if a given object implements the Genre interface.
 */
export function instanceOfGenre(value: object): value is Genre {
    return true;
}

export function GenreFromJSON(json: any): Genre {
    return GenreFromJSONTyped(json, false);
}

export function GenreFromJSONTyped(json: any, ignoreDiscriminator: boolean): Genre {
    if (json == null) {
        return json;
    }
    return {
        
        'iconUrl': json['icon_url'] == null ? undefined : json['icon_url'],
        'id': json['id'] == null ? undefined : json['id'],
        'title': json['title'] == null ? undefined : json['title'],
        'type': json['type'] == null ? undefined : json['type'],
    };
}

export function GenreToJSON(json: any): Genre {
    return GenreToJSONTyped(json, false);
}

export function GenreToJSONTyped(value?: Genre | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'icon_url': value['iconUrl'],
        'id': value['id'],
        'title': value['title'],
        'type': value['type'],
    };
}

