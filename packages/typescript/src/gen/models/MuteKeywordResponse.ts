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
import type { MuteKeyword } from './MuteKeyword.js';
import {
    MuteKeywordFromJSON,
    MuteKeywordFromJSONTyped,
    MuteKeywordToJSON,
    MuteKeywordToJSONTyped,
} from './MuteKeyword.js';

/**
 * 
 * @export
 * @interface MuteKeywordResponse
 */
export interface MuteKeywordResponse {
    /**
     * 
     * @type {Array<MuteKeyword>}
     * @memberof MuteKeywordResponse
     */
    hiddenWords?: Array<MuteKeyword> | null;
}

/**
 * Check if a given object implements the MuteKeywordResponse interface.
 */
export function instanceOfMuteKeywordResponse(value: object): value is MuteKeywordResponse {
    return true;
}

export function MuteKeywordResponseFromJSON(json: any): MuteKeywordResponse {
    return MuteKeywordResponseFromJSONTyped(json, false);
}

export function MuteKeywordResponseFromJSONTyped(json: any, ignoreDiscriminator: boolean): MuteKeywordResponse {
    if (json == null) {
        return json;
    }
    return {
        
        'hiddenWords': json['hidden_words'] == null ? undefined : ((json['hidden_words'] as Array<any>).map(MuteKeywordFromJSON)),
    };
}

export function MuteKeywordResponseToJSON(json: any): MuteKeywordResponse {
    return MuteKeywordResponseToJSONTyped(json, false);
}

export function MuteKeywordResponseToJSONTyped(value?: MuteKeywordResponse | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'hidden_words': value['hiddenWords'] == null ? undefined : ((value['hiddenWords'] as Array<any>).map(MuteKeywordToJSON)),
    };
}

