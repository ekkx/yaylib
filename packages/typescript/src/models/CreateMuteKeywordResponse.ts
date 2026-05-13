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
import type { MuteKeyword } from './MuteKeyword';
import {
    MuteKeywordFromJSON,
    MuteKeywordFromJSONTyped,
    MuteKeywordToJSON,
    MuteKeywordToJSONTyped,
} from './MuteKeyword';

/**
 * 
 * @export
 * @interface CreateMuteKeywordResponse
 */
export interface CreateMuteKeywordResponse {
    /**
     * 
     * @type {MuteKeyword}
     * @memberof CreateMuteKeywordResponse
     */
    hiddenWord?: MuteKeyword | null;
}

/**
 * Check if a given object implements the CreateMuteKeywordResponse interface.
 */
export function instanceOfCreateMuteKeywordResponse(value: object): value is CreateMuteKeywordResponse {
    return true;
}

export function CreateMuteKeywordResponseFromJSON(json: any): CreateMuteKeywordResponse {
    return CreateMuteKeywordResponseFromJSONTyped(json, false);
}

export function CreateMuteKeywordResponseFromJSONTyped(json: any, ignoreDiscriminator: boolean): CreateMuteKeywordResponse {
    if (json == null) {
        return json;
    }
    return {
        
        'hiddenWord': json['hidden_word'] == null ? undefined : MuteKeywordFromJSON(json['hidden_word']),
    };
}

export function CreateMuteKeywordResponseToJSON(json: any): CreateMuteKeywordResponse {
    return CreateMuteKeywordResponseToJSONTyped(json, false);
}

export function CreateMuteKeywordResponseToJSONTyped(value?: CreateMuteKeywordResponse | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'hidden_word': MuteKeywordToJSON(value['hiddenWord']),
    };
}

