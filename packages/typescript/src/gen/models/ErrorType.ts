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
 * @interface ErrorType
 */
export interface ErrorType {
    /**
     * 
     * @type {number}
     * @memberof ErrorType
     */
    code?: number | null;
}

/**
 * Check if a given object implements the ErrorType interface.
 */
export function instanceOfErrorType(value: object): value is ErrorType {
    return true;
}

export function ErrorTypeFromJSON(json: any): ErrorType {
    return ErrorTypeFromJSONTyped(json, false);
}

export function ErrorTypeFromJSONTyped(json: any, ignoreDiscriminator: boolean): ErrorType {
    if (json == null) {
        return json;
    }
    return {
        
        'code': json['code'] == null ? undefined : json['code'],
    };
}

export function ErrorTypeToJSON(json: any): ErrorType {
    return ErrorTypeToJSONTyped(json, false);
}

export function ErrorTypeToJSONTyped(value?: ErrorType | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'code': value['code'],
    };
}

