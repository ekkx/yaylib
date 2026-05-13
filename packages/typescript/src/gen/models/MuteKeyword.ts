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
 * @interface MuteKeyword
 */
export interface MuteKeyword {
    /**
     * 
     * @type {object}
     * @memberof MuteKeyword
     */
    context?: object | null;
    /**
     * 
     * @type {number}
     * @memberof MuteKeyword
     */
    id?: number | null;
    /**
     * 
     * @type {string}
     * @memberof MuteKeyword
     */
    word?: string | null;
}

/**
 * Check if a given object implements the MuteKeyword interface.
 */
export function instanceOfMuteKeyword(value: object): value is MuteKeyword {
    return true;
}

export function MuteKeywordFromJSON(json: any): MuteKeyword {
    return MuteKeywordFromJSONTyped(json, false);
}

export function MuteKeywordFromJSONTyped(json: any, ignoreDiscriminator: boolean): MuteKeyword {
    if (json == null) {
        return json;
    }
    return {
        
        'context': json['context'] == null ? undefined : json['context'],
        'id': json['id'] == null ? undefined : json['id'],
        'word': json['word'] == null ? undefined : json['word'],
    };
}

export function MuteKeywordToJSON(json: any): MuteKeyword {
    return MuteKeywordToJSONTyped(json, false);
}

export function MuteKeywordToJSONTyped(value?: MuteKeyword | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'context': value['context'],
        'id': value['id'],
        'word': value['word'],
    };
}

