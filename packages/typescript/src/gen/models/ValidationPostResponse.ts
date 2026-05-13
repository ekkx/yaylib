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
 * @interface ValidationPostResponse
 */
export interface ValidationPostResponse {
    /**
     * 
     * @type {boolean}
     * @memberof ValidationPostResponse
     */
    isAllowToPost?: boolean | null;
}

/**
 * Check if a given object implements the ValidationPostResponse interface.
 */
export function instanceOfValidationPostResponse(value: object): value is ValidationPostResponse {
    return true;
}

export function ValidationPostResponseFromJSON(json: any): ValidationPostResponse {
    return ValidationPostResponseFromJSONTyped(json, false);
}

export function ValidationPostResponseFromJSONTyped(json: any, ignoreDiscriminator: boolean): ValidationPostResponse {
    if (json == null) {
        return json;
    }
    return {
        
        'isAllowToPost': json['is_allow_to_post'] == null ? undefined : json['is_allow_to_post'],
    };
}

export function ValidationPostResponseToJSON(json: any): ValidationPostResponse {
    return ValidationPostResponseToJSONTyped(json, false);
}

export function ValidationPostResponseToJSONTyped(value?: ValidationPostResponse | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'is_allow_to_post': value['isAllowToPost'],
    };
}

