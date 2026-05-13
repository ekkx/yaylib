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
 * @interface Promotion
 */
export interface Promotion {
    /**
     * 
     * @type {number}
     * @memberof Promotion
     */
    id?: number | null;
    /**
     * 
     * @type {string}
     * @memberof Promotion
     */
    imageUrl?: string | null;
    /**
     * 
     * @type {number}
     * @memberof Promotion
     */
    order?: number | null;
    /**
     * 
     * @type {string}
     * @memberof Promotion
     */
    promotionUrl?: string | null;
    /**
     * 
     * @type {string}
     * @memberof Promotion
     */
    title?: string | null;
}

/**
 * Check if a given object implements the Promotion interface.
 */
export function instanceOfPromotion(value: object): value is Promotion {
    return true;
}

export function PromotionFromJSON(json: any): Promotion {
    return PromotionFromJSONTyped(json, false);
}

export function PromotionFromJSONTyped(json: any, ignoreDiscriminator: boolean): Promotion {
    if (json == null) {
        return json;
    }
    return {
        
        'id': json['id'] == null ? undefined : json['id'],
        'imageUrl': json['image_url'] == null ? undefined : json['image_url'],
        'order': json['order'] == null ? undefined : json['order'],
        'promotionUrl': json['promotion_url'] == null ? undefined : json['promotion_url'],
        'title': json['title'] == null ? undefined : json['title'],
    };
}

export function PromotionToJSON(json: any): Promotion {
    return PromotionToJSONTyped(json, false);
}

export function PromotionToJSONTyped(value?: Promotion | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'id': value['id'],
        'image_url': value['imageUrl'],
        'order': value['order'],
        'promotion_url': value['promotionUrl'],
        'title': value['title'],
    };
}

