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
 * @interface CommonErrorResponse
 */
export interface CommonErrorResponse {
    /**
     * 
     * @type {number}
     * @memberof CommonErrorResponse
     */
    banUntil?: number | null;
    /**
     * 
     * @type {number}
     * @memberof CommonErrorResponse
     */
    errorCode?: number | null;
    /**
     * 
     * @type {string}
     * @memberof CommonErrorResponse
     */
    message?: string | null;
    /**
     * 
     * @type {number}
     * @memberof CommonErrorResponse
     */
    remainingQuota?: number | null;
    /**
     * 
     * @type {number}
     * @memberof CommonErrorResponse
     */
    retryIn?: number | null;
}

/**
 * Check if a given object implements the CommonErrorResponse interface.
 */
export function instanceOfCommonErrorResponse(value: object): value is CommonErrorResponse {
    return true;
}

export function CommonErrorResponseFromJSON(json: any): CommonErrorResponse {
    return CommonErrorResponseFromJSONTyped(json, false);
}

export function CommonErrorResponseFromJSONTyped(json: any, ignoreDiscriminator: boolean): CommonErrorResponse {
    if (json == null) {
        return json;
    }
    return {
        
        'banUntil': json['ban_until'] == null ? undefined : json['ban_until'],
        'errorCode': json['error_code'] == null ? undefined : json['error_code'],
        'message': json['message'] == null ? undefined : json['message'],
        'remainingQuota': json['remaining_quota'] == null ? undefined : json['remaining_quota'],
        'retryIn': json['retry_in'] == null ? undefined : json['retry_in'],
    };
}

export function CommonErrorResponseToJSON(json: any): CommonErrorResponse {
    return CommonErrorResponseToJSONTyped(json, false);
}

export function CommonErrorResponseToJSONTyped(value?: CommonErrorResponse | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'ban_until': value['banUntil'],
        'error_code': value['errorCode'],
        'message': value['message'],
        'remaining_quota': value['remainingQuota'],
        'retry_in': value['retryIn'],
    };
}

