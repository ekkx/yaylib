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
import type { PopularWord } from './PopularWord.js';
import {
    PopularWordFromJSON,
    PopularWordFromJSONTyped,
    PopularWordToJSON,
    PopularWordToJSONTyped,
} from './PopularWord.js';

/**
 * 
 * @export
 * @interface PopularWordsResponse
 */
export interface PopularWordsResponse {
    /**
     * 
     * @type {Array<PopularWord>}
     * @memberof PopularWordsResponse
     */
    popularWords?: Array<PopularWord> | null;
}

/**
 * Check if a given object implements the PopularWordsResponse interface.
 */
export function instanceOfPopularWordsResponse(value: object): value is PopularWordsResponse {
    return true;
}

export function PopularWordsResponseFromJSON(json: any): PopularWordsResponse {
    return PopularWordsResponseFromJSONTyped(json, false);
}

export function PopularWordsResponseFromJSONTyped(json: any, ignoreDiscriminator: boolean): PopularWordsResponse {
    if (json == null) {
        return json;
    }
    return {
        
        'popularWords': json['popular_words'] == null ? undefined : ((json['popular_words'] as Array<any>).map(PopularWordFromJSON)),
    };
}

export function PopularWordsResponseToJSON(json: any): PopularWordsResponse {
    return PopularWordsResponseToJSONTyped(json, false);
}

export function PopularWordsResponseToJSONTyped(value?: PopularWordsResponse | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'popular_words': value['popularWords'] == null ? undefined : ((value['popularWords'] as Array<any>).map(PopularWordToJSON)),
    };
}

