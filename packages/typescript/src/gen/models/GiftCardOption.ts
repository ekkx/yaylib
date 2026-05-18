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
 * @interface GiftCardOption
 */
export interface GiftCardOption {
    /**
     * 
     * @type {number}
     * @memberof GiftCardOption
     */
    empl?: number | null;
    /**
     * 
     * @type {string}
     * @memberof GiftCardOption
     */
    id?: string | null;
    /**
     * 
     * @type {number}
     * @memberof GiftCardOption
     */
    value?: number | null;
}

/**
 * Check if a given object implements the GiftCardOption interface.
 */
export function instanceOfGiftCardOption(value: object): value is GiftCardOption {
    return true;
}

export function GiftCardOptionFromJSON(json: any): GiftCardOption {
    return GiftCardOptionFromJSONTyped(json, false);
}

export function GiftCardOptionFromJSONTyped(json: any, ignoreDiscriminator: boolean): GiftCardOption {
    if (json == null) {
        return json;
    }
    return {
        
        'empl': json['empl'] == null ? undefined : json['empl'],
        'id': json['id'] == null ? undefined : json['id'],
        'value': json['value'] == null ? undefined : json['value'],
    };
}

export function GiftCardOptionToJSON(json: any): GiftCardOption {
    return GiftCardOptionToJSONTyped(json, false);
}

export function GiftCardOptionToJSONTyped(value?: GiftCardOption | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'empl': value['empl'],
        'id': value['id'],
        'value': value['value'],
    };
}

