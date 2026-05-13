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
 * @interface TwoStepAuthEnabled
 */
export interface TwoStepAuthEnabled {
    /**
     * 
     * @type {Array<string>}
     * @memberof TwoStepAuthEnabled
     */
    recoveryCodeList?: Array<string> | null;
}

/**
 * Check if a given object implements the TwoStepAuthEnabled interface.
 */
export function instanceOfTwoStepAuthEnabled(value: object): value is TwoStepAuthEnabled {
    return true;
}

export function TwoStepAuthEnabledFromJSON(json: any): TwoStepAuthEnabled {
    return TwoStepAuthEnabledFromJSONTyped(json, false);
}

export function TwoStepAuthEnabledFromJSONTyped(json: any, ignoreDiscriminator: boolean): TwoStepAuthEnabled {
    if (json == null) {
        return json;
    }
    return {
        
        'recoveryCodeList': json['recovery_code_list'] == null ? undefined : json['recovery_code_list'],
    };
}

export function TwoStepAuthEnabledToJSON(json: any): TwoStepAuthEnabled {
    return TwoStepAuthEnabledToJSONTyped(json, false);
}

export function TwoStepAuthEnabledToJSONTyped(value?: TwoStepAuthEnabled | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'recovery_code_list': value['recoveryCodeList'],
    };
}

