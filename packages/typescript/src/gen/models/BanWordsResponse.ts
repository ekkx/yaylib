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
import type { BanWord } from './BanWord';
import {
    BanWordFromJSON,
    BanWordFromJSONTyped,
    BanWordToJSON,
    BanWordToJSONTyped,
} from './BanWord';

/**
 * 
 * @export
 * @interface BanWordsResponse
 */
export interface BanWordsResponse {
    /**
     * 
     * @type {Array<BanWord>}
     * @memberof BanWordsResponse
     */
    banWords?: Array<BanWord> | null;
}

/**
 * Check if a given object implements the BanWordsResponse interface.
 */
export function instanceOfBanWordsResponse(value: object): value is BanWordsResponse {
    return true;
}

export function BanWordsResponseFromJSON(json: any): BanWordsResponse {
    return BanWordsResponseFromJSONTyped(json, false);
}

export function BanWordsResponseFromJSONTyped(json: any, ignoreDiscriminator: boolean): BanWordsResponse {
    if (json == null) {
        return json;
    }
    return {
        
        'banWords': json['ban_words'] == null ? undefined : ((json['ban_words'] as Array<any>).map(BanWordFromJSON)),
    };
}

export function BanWordsResponseToJSON(json: any): BanWordsResponse {
    return BanWordsResponseToJSONTyped(json, false);
}

export function BanWordsResponseToJSONTyped(value?: BanWordsResponse | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'ban_words': value['banWords'] == null ? undefined : ((value['banWords'] as Array<any>).map(BanWordToJSON)),
    };
}

