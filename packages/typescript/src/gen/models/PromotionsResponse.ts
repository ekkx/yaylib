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
import type { Promotion } from './Promotion';
import {
    PromotionFromJSON,
    PromotionFromJSONTyped,
    PromotionToJSON,
    PromotionToJSONTyped,
} from './Promotion';

/**
 * 
 * @export
 * @interface PromotionsResponse
 */
export interface PromotionsResponse {
    /**
     * 
     * @type {Array<Promotion>}
     * @memberof PromotionsResponse
     */
    promotions?: Array<Promotion> | null;
}

/**
 * Check if a given object implements the PromotionsResponse interface.
 */
export function instanceOfPromotionsResponse(value: object): value is PromotionsResponse {
    return true;
}

export function PromotionsResponseFromJSON(json: any): PromotionsResponse {
    return PromotionsResponseFromJSONTyped(json, false);
}

export function PromotionsResponseFromJSONTyped(json: any, ignoreDiscriminator: boolean): PromotionsResponse {
    if (json == null) {
        return json;
    }
    return {
        
        'promotions': json['promotions'] == null ? undefined : ((json['promotions'] as Array<any>).map(PromotionFromJSON)),
    };
}

export function PromotionsResponseToJSON(json: any): PromotionsResponse {
    return PromotionsResponseToJSONTyped(json, false);
}

export function PromotionsResponseToJSONTyped(value?: PromotionsResponse | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'promotions': value['promotions'] == null ? undefined : ((value['promotions'] as Array<any>).map(PromotionToJSON)),
    };
}

