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
 * @interface TwoFaAuthRequiredDTO
 */
export interface TwoFaAuthRequiredDTO {
    /**
     * 
     * @type {boolean}
     * @memberof TwoFaAuthRequiredDTO
     */
    login?: boolean | null;
    /**
     * 
     * @type {boolean}
     * @memberof TwoFaAuthRequiredDTO
     */
    wallet?: boolean | null;
}

/**
 * Check if a given object implements the TwoFaAuthRequiredDTO interface.
 */
export function instanceOfTwoFaAuthRequiredDTO(value: object): value is TwoFaAuthRequiredDTO {
    return true;
}

export function TwoFaAuthRequiredDTOFromJSON(json: any): TwoFaAuthRequiredDTO {
    return TwoFaAuthRequiredDTOFromJSONTyped(json, false);
}

export function TwoFaAuthRequiredDTOFromJSONTyped(json: any, ignoreDiscriminator: boolean): TwoFaAuthRequiredDTO {
    if (json == null) {
        return json;
    }
    return {
        
        'login': json['login'] == null ? undefined : json['login'],
        'wallet': json['wallet'] == null ? undefined : json['wallet'],
    };
}

export function TwoFaAuthRequiredDTOToJSON(json: any): TwoFaAuthRequiredDTO {
    return TwoFaAuthRequiredDTOToJSONTyped(json, false);
}

export function TwoFaAuthRequiredDTOToJSONTyped(value?: TwoFaAuthRequiredDTO | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'login': value['login'],
        'wallet': value['wallet'],
    };
}

