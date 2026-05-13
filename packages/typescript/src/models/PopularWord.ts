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
 * @interface PopularWord
 */
export interface PopularWord {
    /**
     * 
     * @type {number}
     * @memberof PopularWord
     */
    id?: number | null;
    /**
     * 
     * @type {string}
     * @memberof PopularWord
     */
    type?: string | null;
    /**
     * 
     * @type {string}
     * @memberof PopularWord
     */
    word?: string | null;
}

/**
 * Check if a given object implements the PopularWord interface.
 */
export function instanceOfPopularWord(value: object): value is PopularWord {
    return true;
}

export function PopularWordFromJSON(json: any): PopularWord {
    return PopularWordFromJSONTyped(json, false);
}

export function PopularWordFromJSONTyped(json: any, ignoreDiscriminator: boolean): PopularWord {
    if (json == null) {
        return json;
    }
    return {
        
        'id': json['id'] == null ? undefined : json['id'],
        'type': json['type'] == null ? undefined : json['type'],
        'word': json['word'] == null ? undefined : json['word'],
    };
}

export function PopularWordToJSON(json: any): PopularWord {
    return PopularWordToJSONTyped(json, false);
}

export function PopularWordToJSONTyped(value?: PopularWord | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'id': value['id'],
        'type': value['type'],
        'word': value['word'],
    };
}

