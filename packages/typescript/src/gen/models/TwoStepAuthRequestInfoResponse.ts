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
 * @interface TwoStepAuthRequestInfoResponse
 */
export interface TwoStepAuthRequestInfoResponse {
    /**
     * 
     * @type {string}
     * @memberof TwoStepAuthRequestInfoResponse
     */
    qrCode?: string | null;
    /**
     * 
     * @type {string}
     * @memberof TwoStepAuthRequestInfoResponse
     */
    secret?: string | null;
}

/**
 * Check if a given object implements the TwoStepAuthRequestInfoResponse interface.
 */
export function instanceOfTwoStepAuthRequestInfoResponse(value: object): value is TwoStepAuthRequestInfoResponse {
    return true;
}

export function TwoStepAuthRequestInfoResponseFromJSON(json: any): TwoStepAuthRequestInfoResponse {
    return TwoStepAuthRequestInfoResponseFromJSONTyped(json, false);
}

export function TwoStepAuthRequestInfoResponseFromJSONTyped(json: any, ignoreDiscriminator: boolean): TwoStepAuthRequestInfoResponse {
    if (json == null) {
        return json;
    }
    return {
        
        'qrCode': json['qr_code'] == null ? undefined : json['qr_code'],
        'secret': json['secret'] == null ? undefined : json['secret'],
    };
}

export function TwoStepAuthRequestInfoResponseToJSON(json: any): TwoStepAuthRequestInfoResponse {
    return TwoStepAuthRequestInfoResponseToJSONTyped(json, false);
}

export function TwoStepAuthRequestInfoResponseToJSONTyped(value?: TwoStepAuthRequestInfoResponse | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'qr_code': value['qrCode'],
        'secret': value['secret'],
    };
}

