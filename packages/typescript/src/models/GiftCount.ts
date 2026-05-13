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
 * @interface GiftCount
 */
export interface GiftCount {
    /**
     * 
     * @type {number}
     * @memberof GiftCount
     */
    id?: number | null;
    /**
     * 
     * @type {number}
     * @memberof GiftCount
     */
    quantity?: number | null;
}

/**
 * Check if a given object implements the GiftCount interface.
 */
export function instanceOfGiftCount(value: object): value is GiftCount {
    return true;
}

export function GiftCountFromJSON(json: any): GiftCount {
    return GiftCountFromJSONTyped(json, false);
}

export function GiftCountFromJSONTyped(json: any, ignoreDiscriminator: boolean): GiftCount {
    if (json == null) {
        return json;
    }
    return {
        
        'id': json['id'] == null ? undefined : json['id'],
        'quantity': json['quantity'] == null ? undefined : json['quantity'],
    };
}

export function GiftCountToJSON(json: any): GiftCount {
    return GiftCountToJSONTyped(json, false);
}

export function GiftCountToJSONTyped(value?: GiftCount | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'id': value['id'],
        'quantity': value['quantity'],
    };
}

