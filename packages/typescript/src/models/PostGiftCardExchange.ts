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
 * @interface PostGiftCardExchange
 */
export interface PostGiftCardExchange {
    /**
     * 
     * @type {string}
     * @memberof PostGiftCardExchange
     */
    email?: string | null;
    /**
     * 
     * @type {string}
     * @memberof PostGiftCardExchange
     */
    giftCard?: string | null;
}

/**
 * Check if a given object implements the PostGiftCardExchange interface.
 */
export function instanceOfPostGiftCardExchange(value: object): value is PostGiftCardExchange {
    return true;
}

export function PostGiftCardExchangeFromJSON(json: any): PostGiftCardExchange {
    return PostGiftCardExchangeFromJSONTyped(json, false);
}

export function PostGiftCardExchangeFromJSONTyped(json: any, ignoreDiscriminator: boolean): PostGiftCardExchange {
    if (json == null) {
        return json;
    }
    return {
        
        'email': json['email'] == null ? undefined : json['email'],
        'giftCard': json['gift_card'] == null ? undefined : json['gift_card'],
    };
}

export function PostGiftCardExchangeToJSON(json: any): PostGiftCardExchange {
    return PostGiftCardExchangeToJSONTyped(json, false);
}

export function PostGiftCardExchangeToJSONTyped(value?: PostGiftCardExchange | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'email': value['email'],
        'gift_card': value['giftCard'],
    };
}

