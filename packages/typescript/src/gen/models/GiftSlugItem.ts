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
 * @interface GiftSlugItem
 */
export interface GiftSlugItem {
    /**
     * 
     * @type {string}
     * @memberof GiftSlugItem
     */
    icon?: string | null;
    /**
     * 
     * @type {number}
     * @memberof GiftSlugItem
     */
    id?: number | null;
    /**
     * 
     * @type {number}
     * @memberof GiftSlugItem
     */
    quantity?: number | null;
    /**
     * 
     * @type {string}
     * @memberof GiftSlugItem
     */
    slug?: string | null;
}

/**
 * Check if a given object implements the GiftSlugItem interface.
 */
export function instanceOfGiftSlugItem(value: object): value is GiftSlugItem {
    return true;
}

export function GiftSlugItemFromJSON(json: any): GiftSlugItem {
    return GiftSlugItemFromJSONTyped(json, false);
}

export function GiftSlugItemFromJSONTyped(json: any, ignoreDiscriminator: boolean): GiftSlugItem {
    if (json == null) {
        return json;
    }
    return {
        
        'icon': json['icon'] == null ? undefined : json['icon'],
        'id': json['id'] == null ? undefined : json['id'],
        'quantity': json['quantity'] == null ? undefined : json['quantity'],
        'slug': json['slug'] == null ? undefined : json['slug'],
    };
}

export function GiftSlugItemToJSON(json: any): GiftSlugItem {
    return GiftSlugItemToJSONTyped(json, false);
}

export function GiftSlugItemToJSONTyped(value?: GiftSlugItem | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'icon': value['icon'],
        'id': value['id'],
        'quantity': value['quantity'],
        'slug': value['slug'],
    };
}

