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
 * @interface Bgm
 */
export interface Bgm {
    /**
     * 
     * @type {number}
     * @memberof Bgm
     */
    id?: number | null;
    /**
     * 
     * @type {string}
     * @memberof Bgm
     */
    musicUrl?: string | null;
    /**
     * 
     * @type {number}
     * @memberof Bgm
     */
    order?: number | null;
    /**
     * 
     * @type {string}
     * @memberof Bgm
     */
    title?: string | null;
}

/**
 * Check if a given object implements the Bgm interface.
 */
export function instanceOfBgm(value: object): value is Bgm {
    return true;
}

export function BgmFromJSON(json: any): Bgm {
    return BgmFromJSONTyped(json, false);
}

export function BgmFromJSONTyped(json: any, ignoreDiscriminator: boolean): Bgm {
    if (json == null) {
        return json;
    }
    return {
        
        'id': json['id'] == null ? undefined : json['id'],
        'musicUrl': json['music_url'] == null ? undefined : json['music_url'],
        'order': json['order'] == null ? undefined : json['order'],
        'title': json['title'] == null ? undefined : json['title'],
    };
}

export function BgmToJSON(json: any): Bgm {
    return BgmToJSONTyped(json, false);
}

export function BgmToJSONTyped(value?: Bgm | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'id': value['id'],
        'music_url': value['musicUrl'],
        'order': value['order'],
        'title': value['title'],
    };
}

