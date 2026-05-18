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
 * @interface TwoStepAuthRequestInfo
 */
export interface TwoStepAuthRequestInfo {
    /**
     * 
     * @type {string}
     * @memberof TwoStepAuthRequestInfo
     */
    qrCodeUrl?: string | null;
    /**
     * 
     * @type {string}
     * @memberof TwoStepAuthRequestInfo
     */
    setupKey?: string | null;
}

/**
 * Check if a given object implements the TwoStepAuthRequestInfo interface.
 */
export function instanceOfTwoStepAuthRequestInfo(value: object): value is TwoStepAuthRequestInfo {
    return true;
}

export function TwoStepAuthRequestInfoFromJSON(json: any): TwoStepAuthRequestInfo {
    return TwoStepAuthRequestInfoFromJSONTyped(json, false);
}

export function TwoStepAuthRequestInfoFromJSONTyped(json: any, ignoreDiscriminator: boolean): TwoStepAuthRequestInfo {
    if (json == null) {
        return json;
    }
    return {
        
        'qrCodeUrl': json['qr_code_url'] == null ? undefined : json['qr_code_url'],
        'setupKey': json['setup_key'] == null ? undefined : json['setup_key'],
    };
}

export function TwoStepAuthRequestInfoToJSON(json: any): TwoStepAuthRequestInfo {
    return TwoStepAuthRequestInfoToJSONTyped(json, false);
}

export function TwoStepAuthRequestInfoToJSONTyped(value?: TwoStepAuthRequestInfo | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'qr_code_url': value['qrCodeUrl'],
        'setup_key': value['setupKey'],
    };
}

