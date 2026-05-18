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
 * @interface InAppPurchaseProduct
 */
export interface InAppPurchaseProduct {
    /**
     * 
     * @type {number}
     * @memberof InAppPurchaseProduct
     */
    amount?: number | null;
    /**
     * 
     * @type {number}
     * @memberof InAppPurchaseProduct
     */
    bonus?: number | null;
    /**
     * 
     * @type {string}
     * @memberof InAppPurchaseProduct
     */
    id?: string | null;
    /**
     * 
     * @type {boolean}
     * @memberof InAppPurchaseProduct
     */
    purchasable?: boolean | null;
}

/**
 * Check if a given object implements the InAppPurchaseProduct interface.
 */
export function instanceOfInAppPurchaseProduct(value: object): value is InAppPurchaseProduct {
    return true;
}

export function InAppPurchaseProductFromJSON(json: any): InAppPurchaseProduct {
    return InAppPurchaseProductFromJSONTyped(json, false);
}

export function InAppPurchaseProductFromJSONTyped(json: any, ignoreDiscriminator: boolean): InAppPurchaseProduct {
    if (json == null) {
        return json;
    }
    return {
        
        'amount': json['amount'] == null ? undefined : json['amount'],
        'bonus': json['bonus'] == null ? undefined : json['bonus'],
        'id': json['id'] == null ? undefined : json['id'],
        'purchasable': json['purchasable'] == null ? undefined : json['purchasable'],
    };
}

export function InAppPurchaseProductToJSON(json: any): InAppPurchaseProduct {
    return InAppPurchaseProductToJSONTyped(json, false);
}

export function InAppPurchaseProductToJSONTyped(value?: InAppPurchaseProduct | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'amount': value['amount'],
        'bonus': value['bonus'],
        'id': value['id'],
        'purchasable': value['purchasable'],
    };
}

