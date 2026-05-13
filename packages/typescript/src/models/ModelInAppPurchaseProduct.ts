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
 * @interface ModelInAppPurchaseProduct
 */
export interface ModelInAppPurchaseProduct {
    /**
     * 
     * @type {number}
     * @memberof ModelInAppPurchaseProduct
     */
    amount?: number | null;
    /**
     * 
     * @type {number}
     * @memberof ModelInAppPurchaseProduct
     */
    bonus?: number | null;
    /**
     * 
     * @type {string}
     * @memberof ModelInAppPurchaseProduct
     */
    currencyCode?: string | null;
    /**
     * 
     * @type {string}
     * @memberof ModelInAppPurchaseProduct
     */
    id?: string | null;
    /**
     * 
     * @type {number}
     * @memberof ModelInAppPurchaseProduct
     */
    price?: number | null;
    /**
     * 
     * @type {boolean}
     * @memberof ModelInAppPurchaseProduct
     */
    purchasable?: boolean | null;
}

/**
 * Check if a given object implements the ModelInAppPurchaseProduct interface.
 */
export function instanceOfModelInAppPurchaseProduct(value: object): value is ModelInAppPurchaseProduct {
    return true;
}

export function ModelInAppPurchaseProductFromJSON(json: any): ModelInAppPurchaseProduct {
    return ModelInAppPurchaseProductFromJSONTyped(json, false);
}

export function ModelInAppPurchaseProductFromJSONTyped(json: any, ignoreDiscriminator: boolean): ModelInAppPurchaseProduct {
    if (json == null) {
        return json;
    }
    return {
        
        'amount': json['amount'] == null ? undefined : json['amount'],
        'bonus': json['bonus'] == null ? undefined : json['bonus'],
        'currencyCode': json['currency_code'] == null ? undefined : json['currency_code'],
        'id': json['id'] == null ? undefined : json['id'],
        'price': json['price'] == null ? undefined : json['price'],
        'purchasable': json['purchasable'] == null ? undefined : json['purchasable'],
    };
}

export function ModelInAppPurchaseProductToJSON(json: any): ModelInAppPurchaseProduct {
    return ModelInAppPurchaseProductToJSONTyped(json, false);
}

export function ModelInAppPurchaseProductToJSONTyped(value?: ModelInAppPurchaseProduct | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'amount': value['amount'],
        'bonus': value['bonus'],
        'currency_code': value['currencyCode'],
        'id': value['id'],
        'price': value['price'],
        'purchasable': value['purchasable'],
    };
}

