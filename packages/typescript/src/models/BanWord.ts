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
 * @interface BanWord
 */
export interface BanWord {
    /**
     * 
     * @type {number}
     * @memberof BanWord
     */
    id?: number | null;
    /**
     * 
     * @type {string}
     * @memberof BanWord
     */
    type?: string | null;
    /**
     * 
     * @type {string}
     * @memberof BanWord
     */
    word?: string | null;
}

/**
 * Check if a given object implements the BanWord interface.
 */
export function instanceOfBanWord(value: object): value is BanWord {
    return true;
}

export function BanWordFromJSON(json: any): BanWord {
    return BanWordFromJSONTyped(json, false);
}

export function BanWordFromJSONTyped(json: any, ignoreDiscriminator: boolean): BanWord {
    if (json == null) {
        return json;
    }
    return {
        
        'id': json['id'] == null ? undefined : json['id'],
        'type': json['type'] == null ? undefined : json['type'],
        'word': json['word'] == null ? undefined : json['word'],
    };
}

export function BanWordToJSON(json: any): BanWord {
    return BanWordToJSONTyped(json, false);
}

export function BanWordToJSONTyped(value?: BanWord | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'id': value['id'],
        'type': value['type'],
        'word': value['word'],
    };
}

