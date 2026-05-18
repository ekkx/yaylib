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
import type { Gift } from './Gift.js';
import {
    GiftFromJSON,
    GiftFromJSONTyped,
    GiftToJSON,
    GiftToJSONTyped,
} from './Gift.js';

/**
 * 
 * @export
 * @interface PostGift
 */
export interface PostGift {
    /**
     * 
     * @type {number}
     * @memberof PostGift
     */
    count?: number | null;
    /**
     * 
     * @type {Gift}
     * @memberof PostGift
     */
    gift?: Gift | null;
}

/**
 * Check if a given object implements the PostGift interface.
 */
export function instanceOfPostGift(value: object): value is PostGift {
    return true;
}

export function PostGiftFromJSON(json: any): PostGift {
    return PostGiftFromJSONTyped(json, false);
}

export function PostGiftFromJSONTyped(json: any, ignoreDiscriminator: boolean): PostGift {
    if (json == null) {
        return json;
    }
    return {
        
        'count': json['count'] == null ? undefined : json['count'],
        'gift': json['gift'] == null ? undefined : GiftFromJSON(json['gift']),
    };
}

export function PostGiftToJSON(json: any): PostGift {
    return PostGiftToJSONTyped(json, false);
}

export function PostGiftToJSONTyped(value?: PostGift | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'count': value['count'],
        'gift': GiftToJSON(value['gift']),
    };
}

