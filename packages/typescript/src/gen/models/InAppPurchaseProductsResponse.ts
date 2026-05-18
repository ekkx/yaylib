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
import type { ProductQuota } from './ProductQuota.js';
import {
    ProductQuotaFromJSON,
    ProductQuotaFromJSONTyped,
    ProductQuotaToJSON,
    ProductQuotaToJSONTyped,
} from './ProductQuota.js';
import type { InAppPurchaseProduct } from './InAppPurchaseProduct.js';
import {
    InAppPurchaseProductFromJSON,
    InAppPurchaseProductFromJSONTyped,
    InAppPurchaseProductToJSON,
    InAppPurchaseProductToJSONTyped,
} from './InAppPurchaseProduct.js';

/**
 * 
 * @export
 * @interface InAppPurchaseProductsResponse
 */
export interface InAppPurchaseProductsResponse {
    /**
     * 
     * @type {Array<InAppPurchaseProduct>}
     * @memberof InAppPurchaseProductsResponse
     */
    iapProducts?: Array<InAppPurchaseProduct> | null;
    /**
     * 
     * @type {ProductQuota}
     * @memberof InAppPurchaseProductsResponse
     */
    quota?: ProductQuota | null;
}

/**
 * Check if a given object implements the InAppPurchaseProductsResponse interface.
 */
export function instanceOfInAppPurchaseProductsResponse(value: object): value is InAppPurchaseProductsResponse {
    return true;
}

export function InAppPurchaseProductsResponseFromJSON(json: any): InAppPurchaseProductsResponse {
    return InAppPurchaseProductsResponseFromJSONTyped(json, false);
}

export function InAppPurchaseProductsResponseFromJSONTyped(json: any, ignoreDiscriminator: boolean): InAppPurchaseProductsResponse {
    if (json == null) {
        return json;
    }
    return {
        
        'iapProducts': json['iap_products'] == null ? undefined : ((json['iap_products'] as Array<any>).map(InAppPurchaseProductFromJSON)),
        'quota': json['quota'] == null ? undefined : ProductQuotaFromJSON(json['quota']),
    };
}

export function InAppPurchaseProductsResponseToJSON(json: any): InAppPurchaseProductsResponse {
    return InAppPurchaseProductsResponseToJSONTyped(json, false);
}

export function InAppPurchaseProductsResponseToJSONTyped(value?: InAppPurchaseProductsResponse | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'iap_products': value['iapProducts'] == null ? undefined : ((value['iapProducts'] as Array<any>).map(InAppPurchaseProductToJSON)),
        'quota': ProductQuotaToJSON(value['quota']),
    };
}

