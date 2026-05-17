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
import type { Sticker } from './Sticker';
import {
    StickerFromJSON,
    StickerFromJSONTyped,
    StickerToJSON,
    StickerToJSONTyped,
} from './Sticker';

/**
 * 
 * @export
 * @interface StickerPack
 */
export interface StickerPack {
    /**
     * 
     * @type {string}
     * @memberof StickerPack
     */
    cover?: string | null;
    /**
     * 
     * @type {string}
     * @memberof StickerPack
     */
    description?: string | null;
    /**
     * 
     * @type {number}
     * @memberof StickerPack
     */
    id?: number | null;
    /**
     * 
     * @type {string}
     * @memberof StickerPack
     */
    name?: string | null;
    /**
     * 
     * @type {number}
     * @memberof StickerPack
     */
    order?: number | null;
    /**
     * 
     * @type {Array<Sticker>}
     * @memberof StickerPack
     */
    stickers?: Array<Sticker> | null;
}

/**
 * Check if a given object implements the StickerPack interface.
 */
export function instanceOfStickerPack(value: object): value is StickerPack {
    return true;
}

export function StickerPackFromJSON(json: any): StickerPack {
    return StickerPackFromJSONTyped(json, false);
}

export function StickerPackFromJSONTyped(json: any, ignoreDiscriminator: boolean): StickerPack {
    if (json == null) {
        return json;
    }
    return {
        
        'cover': json['cover'] == null ? undefined : json['cover'],
        'description': json['description'] == null ? undefined : json['description'],
        'id': json['id'] == null ? undefined : json['id'],
        'name': json['name'] == null ? undefined : json['name'],
        'order': json['order'] == null ? undefined : json['order'],
        'stickers': json['stickers'] == null ? undefined : ((json['stickers'] as Array<any>).map(StickerFromJSON)),
    };
}

export function StickerPackToJSON(json: any): StickerPack {
    return StickerPackToJSONTyped(json, false);
}

export function StickerPackToJSONTyped(value?: StickerPack | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'cover': value['cover'],
        'description': value['description'],
        'id': value['id'],
        'name': value['name'],
        'order': value['order'],
        'stickers': value['stickers'] == null ? undefined : ((value['stickers'] as Array<any>).map(StickerToJSON)),
    };
}

