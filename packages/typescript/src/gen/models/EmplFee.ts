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
 * @interface EmplFee
 */
export interface EmplFee {
    /**
     * 
     * @type {string}
     * @memberof EmplFee
     */
    gasFee?: string | null;
}

/**
 * Check if a given object implements the EmplFee interface.
 */
export function instanceOfEmplFee(value: object): value is EmplFee {
    return true;
}

export function EmplFeeFromJSON(json: any): EmplFee {
    return EmplFeeFromJSONTyped(json, false);
}

export function EmplFeeFromJSONTyped(json: any, ignoreDiscriminator: boolean): EmplFee {
    if (json == null) {
        return json;
    }
    return {
        
        'gasFee': json['gas_fee'] == null ? undefined : json['gas_fee'],
    };
}

export function EmplFeeToJSON(json: any): EmplFee {
    return EmplFeeToJSONTyped(json, false);
}

export function EmplFeeToJSONTyped(value?: EmplFee | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'gas_fee': value['gasFee'],
    };
}

