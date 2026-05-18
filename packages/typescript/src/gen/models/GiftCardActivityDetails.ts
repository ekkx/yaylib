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
 * @interface GiftCardActivityDetails
 */
export interface GiftCardActivityDetails {
    /**
     * 
     * @type {number}
     * @memberof GiftCardActivityDetails
     */
    createdAt?: number | null;
    /**
     * 
     * @type {string}
     * @memberof GiftCardActivityDetails
     */
    email?: string | null;
    /**
     * 
     * @type {number}
     * @memberof GiftCardActivityDetails
     */
    emplAmount?: number | null;
    /**
     * 
     * @type {number}
     * @memberof GiftCardActivityDetails
     */
    giftCardAmount?: number | null;
    /**
     * 
     * @type {object}
     * @memberof GiftCardActivityDetails
     */
    status?: object | null;
}

/**
 * Check if a given object implements the GiftCardActivityDetails interface.
 */
export function instanceOfGiftCardActivityDetails(value: object): value is GiftCardActivityDetails {
    return true;
}

export function GiftCardActivityDetailsFromJSON(json: any): GiftCardActivityDetails {
    return GiftCardActivityDetailsFromJSONTyped(json, false);
}

export function GiftCardActivityDetailsFromJSONTyped(json: any, ignoreDiscriminator: boolean): GiftCardActivityDetails {
    if (json == null) {
        return json;
    }
    return {
        
        'createdAt': json['created_at'] == null ? undefined : json['created_at'],
        'email': json['email'] == null ? undefined : json['email'],
        'emplAmount': json['empl_amount'] == null ? undefined : json['empl_amount'],
        'giftCardAmount': json['gift_card_amount'] == null ? undefined : json['gift_card_amount'],
        'status': json['status'] == null ? undefined : json['status'],
    };
}

export function GiftCardActivityDetailsToJSON(json: any): GiftCardActivityDetails {
    return GiftCardActivityDetailsToJSONTyped(json, false);
}

export function GiftCardActivityDetailsToJSONTyped(value?: GiftCardActivityDetails | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'created_at': value['createdAt'],
        'email': value['email'],
        'empl_amount': value['emplAmount'],
        'gift_card_amount': value['giftCardAmount'],
        'status': value['status'],
    };
}

