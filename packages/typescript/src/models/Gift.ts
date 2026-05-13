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
 * @interface Gift
 */
export interface Gift {
    /**
     * 
     * @type {string}
     * @memberof Gift
     */
    currency?: string | null;
    /**
     * 
     * @type {string}
     * @memberof Gift
     */
    icon?: string | null;
    /**
     * 
     * @type {string}
     * @memberof Gift
     */
    iconThumbnail?: string | null;
    /**
     * 
     * @type {number}
     * @memberof Gift
     */
    id?: number | null;
    /**
     * 
     * @type {number}
     * @memberof Gift
     */
    price?: number | null;
    /**
     * 
     * @type {string}
     * @memberof Gift
     */
    title?: string | null;
}

/**
 * Check if a given object implements the Gift interface.
 */
export function instanceOfGift(value: object): value is Gift {
    return true;
}

export function GiftFromJSON(json: any): Gift {
    return GiftFromJSONTyped(json, false);
}

export function GiftFromJSONTyped(json: any, ignoreDiscriminator: boolean): Gift {
    if (json == null) {
        return json;
    }
    return {
        
        'currency': json['currency'] == null ? undefined : json['currency'],
        'icon': json['icon'] == null ? undefined : json['icon'],
        'iconThumbnail': json['icon_thumbnail'] == null ? undefined : json['icon_thumbnail'],
        'id': json['id'] == null ? undefined : json['id'],
        'price': json['price'] == null ? undefined : json['price'],
        'title': json['title'] == null ? undefined : json['title'],
    };
}

export function GiftToJSON(json: any): Gift {
    return GiftToJSONTyped(json, false);
}

export function GiftToJSONTyped(value?: Gift | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'currency': value['currency'],
        'icon': value['icon'],
        'icon_thumbnail': value['iconThumbnail'],
        'id': value['id'],
        'price': value['price'],
        'title': value['title'],
    };
}

