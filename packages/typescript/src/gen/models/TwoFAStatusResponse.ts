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
import type { TwoFaAuthRequiredDTO } from './TwoFaAuthRequiredDTO';
import {
    TwoFaAuthRequiredDTOFromJSON,
    TwoFaAuthRequiredDTOFromJSONTyped,
    TwoFaAuthRequiredDTOToJSON,
    TwoFaAuthRequiredDTOToJSONTyped,
} from './TwoFaAuthRequiredDTO';

/**
 * 
 * @export
 * @interface TwoFAStatusResponse
 */
export interface TwoFAStatusResponse {
    /**
     * 
     * @type {TwoFaAuthRequiredDTO}
     * @memberof TwoFAStatusResponse
     */
    twoFaAuthRequired?: TwoFaAuthRequiredDTO | null;
    /**
     * 
     * @type {boolean}
     * @memberof TwoFAStatusResponse
     */
    twoFaEnabled?: boolean | null;
}

/**
 * Check if a given object implements the TwoFAStatusResponse interface.
 */
export function instanceOfTwoFAStatusResponse(value: object): value is TwoFAStatusResponse {
    return true;
}

export function TwoFAStatusResponseFromJSON(json: any): TwoFAStatusResponse {
    return TwoFAStatusResponseFromJSONTyped(json, false);
}

export function TwoFAStatusResponseFromJSONTyped(json: any, ignoreDiscriminator: boolean): TwoFAStatusResponse {
    if (json == null) {
        return json;
    }
    return {
        
        'twoFaAuthRequired': json['two_fa_auth_required'] == null ? undefined : TwoFaAuthRequiredDTOFromJSON(json['two_fa_auth_required']),
        'twoFaEnabled': json['two_fa_enabled'] == null ? undefined : json['two_fa_enabled'],
    };
}

export function TwoFAStatusResponseToJSON(json: any): TwoFAStatusResponse {
    return TwoFAStatusResponseToJSONTyped(json, false);
}

export function TwoFAStatusResponseToJSONTyped(value?: TwoFAStatusResponse | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'two_fa_auth_required': TwoFaAuthRequiredDTOToJSON(value['twoFaAuthRequired']),
        'two_fa_enabled': value['twoFaEnabled'],
    };
}

