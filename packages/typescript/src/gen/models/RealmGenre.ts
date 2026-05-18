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
 * @interface RealmGenre
 */
export interface RealmGenre {
    /**
     * 
     * @type {string}
     * @memberof RealmGenre
     */
    iconUrl?: string | null;
    /**
     * 
     * @type {number}
     * @memberof RealmGenre
     */
    id?: number | null;
    /**
     * 
     * @type {string}
     * @memberof RealmGenre
     */
    title?: string | null;
}

/**
 * Check if a given object implements the RealmGenre interface.
 */
export function instanceOfRealmGenre(value: object): value is RealmGenre {
    return true;
}

export function RealmGenreFromJSON(json: any): RealmGenre {
    return RealmGenreFromJSONTyped(json, false);
}

export function RealmGenreFromJSONTyped(json: any, ignoreDiscriminator: boolean): RealmGenre {
    if (json == null) {
        return json;
    }
    return {
        
        'iconUrl': json['icon_url'] == null ? undefined : json['icon_url'],
        'id': json['id'] == null ? undefined : json['id'],
        'title': json['title'] == null ? undefined : json['title'],
    };
}

export function RealmGenreToJSON(json: any): RealmGenre {
    return RealmGenreToJSONTyped(json, false);
}

export function RealmGenreToJSONTyped(value?: RealmGenre | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'icon_url': value['iconUrl'],
        'id': value['id'],
        'title': value['title'],
    };
}

