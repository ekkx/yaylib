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
 * @interface RealmGift
 */
export interface RealmGift {
    /**
     * 
     * @type {string}
     * @memberof RealmGift
     */
    currency?: string | null;
    /**
     * 
     * @type {string}
     * @memberof RealmGift
     */
    icon?: string | null;
    /**
     * 
     * @type {string}
     * @memberof RealmGift
     */
    iconThumbnail?: string | null;
    /**
     * 
     * @type {number}
     * @memberof RealmGift
     */
    id?: number | null;
    /**
     * 
     * @type {number}
     * @memberof RealmGift
     */
    price?: number | null;
    /**
     * 
     * @type {string}
     * @memberof RealmGift
     */
    slug?: string | null;
    /**
     * 
     * @type {string}
     * @memberof RealmGift
     */
    title?: string | null;
}

/**
 * Check if a given object implements the RealmGift interface.
 */
export function instanceOfRealmGift(value: object): value is RealmGift {
    return true;
}

export function RealmGiftFromJSON(json: any): RealmGift {
    return RealmGiftFromJSONTyped(json, false);
}

export function RealmGiftFromJSONTyped(json: any, ignoreDiscriminator: boolean): RealmGift {
    if (json == null) {
        return json;
    }
    return {
        
        'currency': json['currency'] == null ? undefined : json['currency'],
        'icon': json['icon'] == null ? undefined : json['icon'],
        'iconThumbnail': json['icon_thumbnail'] == null ? undefined : json['icon_thumbnail'],
        'id': json['id'] == null ? undefined : json['id'],
        'price': json['price'] == null ? undefined : json['price'],
        'slug': json['slug'] == null ? undefined : json['slug'],
        'title': json['title'] == null ? undefined : json['title'],
    };
}

export function RealmGiftToJSON(json: any): RealmGift {
    return RealmGiftToJSONTyped(json, false);
}

export function RealmGiftToJSONTyped(value?: RealmGift | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'currency': value['currency'],
        'icon': value['icon'],
        'icon_thumbnail': value['iconThumbnail'],
        'id': value['id'],
        'price': value['price'],
        'slug': value['slug'],
        'title': value['title'],
    };
}

