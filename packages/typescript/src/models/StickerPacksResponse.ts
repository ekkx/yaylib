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
import type { StickerPack } from './StickerPack';
import {
    StickerPackFromJSON,
    StickerPackFromJSONTyped,
    StickerPackToJSON,
    StickerPackToJSONTyped,
} from './StickerPack';

/**
 * 
 * @export
 * @interface StickerPacksResponse
 */
export interface StickerPacksResponse {
    /**
     * 
     * @type {Array<StickerPack>}
     * @memberof StickerPacksResponse
     */
    stickerPacks?: Array<StickerPack> | null;
}

/**
 * Check if a given object implements the StickerPacksResponse interface.
 */
export function instanceOfStickerPacksResponse(value: object): value is StickerPacksResponse {
    return true;
}

export function StickerPacksResponseFromJSON(json: any): StickerPacksResponse {
    return StickerPacksResponseFromJSONTyped(json, false);
}

export function StickerPacksResponseFromJSONTyped(json: any, ignoreDiscriminator: boolean): StickerPacksResponse {
    if (json == null) {
        return json;
    }
    return {
        
        'stickerPacks': json['sticker_packs'] == null ? undefined : ((json['sticker_packs'] as Array<any>).map(StickerPackFromJSON)),
    };
}

export function StickerPacksResponseToJSON(json: any): StickerPacksResponse {
    return StickerPacksResponseToJSONTyped(json, false);
}

export function StickerPacksResponseToJSONTyped(value?: StickerPacksResponse | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'sticker_packs': value['stickerPacks'] == null ? undefined : ((value['stickerPacks'] as Array<any>).map(StickerPackToJSON)),
    };
}

