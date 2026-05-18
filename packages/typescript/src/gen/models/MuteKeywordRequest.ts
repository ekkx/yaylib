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
 * @interface MuteKeywordRequest
 */
export interface MuteKeywordRequest {
    /**
     * 
     * @type {Array<string>}
     * @memberof MuteKeywordRequest
     */
    context?: Array<string> | null;
    /**
     * 
     * @type {string}
     * @memberof MuteKeywordRequest
     */
    word?: string | null;
}

/**
 * Check if a given object implements the MuteKeywordRequest interface.
 */
export function instanceOfMuteKeywordRequest(value: object): value is MuteKeywordRequest {
    return true;
}

export function MuteKeywordRequestFromJSON(json: any): MuteKeywordRequest {
    return MuteKeywordRequestFromJSONTyped(json, false);
}

export function MuteKeywordRequestFromJSONTyped(json: any, ignoreDiscriminator: boolean): MuteKeywordRequest {
    if (json == null) {
        return json;
    }
    return {
        
        'context': json['context'] == null ? undefined : json['context'],
        'word': json['word'] == null ? undefined : json['word'],
    };
}

export function MuteKeywordRequestToJSON(json: any): MuteKeywordRequest {
    return MuteKeywordRequestToJSONTyped(json, false);
}

export function MuteKeywordRequestToJSONTyped(value?: MuteKeywordRequest | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'context': value['context'],
        'word': value['word'],
    };
}

