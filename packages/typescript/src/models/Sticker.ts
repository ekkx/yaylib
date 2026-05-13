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
 * @interface Sticker
 */
export interface Sticker {
    /**
     * 
     * @type {string}
     * @memberof Sticker
     */
    extension?: string | null;
    /**
     * 
     * @type {number}
     * @memberof Sticker
     */
    height?: number | null;
    /**
     * 
     * @type {number}
     * @memberof Sticker
     */
    id?: number | null;
    /**
     * 
     * @type {number}
     * @memberof Sticker
     */
    stickerPackId?: number | null;
    /**
     * 
     * @type {string}
     * @memberof Sticker
     */
    url?: string | null;
    /**
     * 
     * @type {number}
     * @memberof Sticker
     */
    width?: number | null;
}

/**
 * Check if a given object implements the Sticker interface.
 */
export function instanceOfSticker(value: object): value is Sticker {
    return true;
}

export function StickerFromJSON(json: any): Sticker {
    return StickerFromJSONTyped(json, false);
}

export function StickerFromJSONTyped(json: any, ignoreDiscriminator: boolean): Sticker {
    if (json == null) {
        return json;
    }
    return {
        
        'extension': json['extension'] == null ? undefined : json['extension'],
        'height': json['height'] == null ? undefined : json['height'],
        'id': json['id'] == null ? undefined : json['id'],
        'stickerPackId': json['sticker_pack_id'] == null ? undefined : json['sticker_pack_id'],
        'url': json['url'] == null ? undefined : json['url'],
        'width': json['width'] == null ? undefined : json['width'],
    };
}

export function StickerToJSON(json: any): Sticker {
    return StickerToJSONTyped(json, false);
}

export function StickerToJSONTyped(value?: Sticker | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'extension': value['extension'],
        'height': value['height'],
        'id': value['id'],
        'sticker_pack_id': value['stickerPackId'],
        'url': value['url'],
        'width': value['width'],
    };
}

