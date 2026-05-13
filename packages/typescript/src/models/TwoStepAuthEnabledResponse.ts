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
 * @interface TwoStepAuthEnabledResponse
 */
export interface TwoStepAuthEnabledResponse {
    /**
     * 
     * @type {Array<string>}
     * @memberof TwoStepAuthEnabledResponse
     */
    recoveryCodes?: Array<string> | null;
}

/**
 * Check if a given object implements the TwoStepAuthEnabledResponse interface.
 */
export function instanceOfTwoStepAuthEnabledResponse(value: object): value is TwoStepAuthEnabledResponse {
    return true;
}

export function TwoStepAuthEnabledResponseFromJSON(json: any): TwoStepAuthEnabledResponse {
    return TwoStepAuthEnabledResponseFromJSONTyped(json, false);
}

export function TwoStepAuthEnabledResponseFromJSONTyped(json: any, ignoreDiscriminator: boolean): TwoStepAuthEnabledResponse {
    if (json == null) {
        return json;
    }
    return {
        
        'recoveryCodes': json['recovery_codes'] == null ? undefined : json['recovery_codes'],
    };
}

export function TwoStepAuthEnabledResponseToJSON(json: any): TwoStepAuthEnabledResponse {
    return TwoStepAuthEnabledResponseToJSONTyped(json, false);
}

export function TwoStepAuthEnabledResponseToJSONTyped(value?: TwoStepAuthEnabledResponse | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'recovery_codes': value['recoveryCodes'],
    };
}

