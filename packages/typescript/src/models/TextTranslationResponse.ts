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
 * @interface TextTranslationResponse
 */
export interface TextTranslationResponse {
    /**
     * 
     * @type {string}
     * @memberof TextTranslationResponse
     */
    translatedText?: string | null;
}

/**
 * Check if a given object implements the TextTranslationResponse interface.
 */
export function instanceOfTextTranslationResponse(value: object): value is TextTranslationResponse {
    return true;
}

export function TextTranslationResponseFromJSON(json: any): TextTranslationResponse {
    return TextTranslationResponseFromJSONTyped(json, false);
}

export function TextTranslationResponseFromJSONTyped(json: any, ignoreDiscriminator: boolean): TextTranslationResponse {
    if (json == null) {
        return json;
    }
    return {
        
        'translatedText': json['translated_text'] == null ? undefined : json['translated_text'],
    };
}

export function TextTranslationResponseToJSON(json: any): TextTranslationResponse {
    return TextTranslationResponseToJSONTyped(json, false);
}

export function TextTranslationResponseToJSONTyped(value?: TextTranslationResponse | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'translated_text': value['translatedText'],
    };
}

