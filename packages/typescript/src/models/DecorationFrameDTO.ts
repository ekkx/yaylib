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
 * @interface DecorationFrameDTO
 */
export interface DecorationFrameDTO {
    /**
     * 
     * @type {boolean}
     * @memberof DecorationFrameDTO
     */
    current?: boolean | null;
    /**
     * 
     * @type {string}
     * @memberof DecorationFrameDTO
     */
    frame?: string | null;
    /**
     * 
     * @type {string}
     * @memberof DecorationFrameDTO
     */
    frameThumbnail?: string | null;
    /**
     * 
     * @type {string}
     * @memberof DecorationFrameDTO
     */
    frameType?: string | null;
    /**
     * 
     * @type {number}
     * @memberof DecorationFrameDTO
     */
    id?: number | null;
    /**
     * 
     * @type {boolean}
     * @memberof DecorationFrameDTO
     */
    owned?: boolean | null;
    /**
     * 
     * @type {number}
     * @memberof DecorationFrameDTO
     */
    price?: number | null;
    /**
     * 
     * @type {boolean}
     * @memberof DecorationFrameDTO
     */
    vipOnly?: boolean | null;
}

/**
 * Check if a given object implements the DecorationFrameDTO interface.
 */
export function instanceOfDecorationFrameDTO(value: object): value is DecorationFrameDTO {
    return true;
}

export function DecorationFrameDTOFromJSON(json: any): DecorationFrameDTO {
    return DecorationFrameDTOFromJSONTyped(json, false);
}

export function DecorationFrameDTOFromJSONTyped(json: any, ignoreDiscriminator: boolean): DecorationFrameDTO {
    if (json == null) {
        return json;
    }
    return {
        
        'current': json['current'] == null ? undefined : json['current'],
        'frame': json['frame'] == null ? undefined : json['frame'],
        'frameThumbnail': json['frame_thumbnail'] == null ? undefined : json['frame_thumbnail'],
        'frameType': json['frame_type'] == null ? undefined : json['frame_type'],
        'id': json['id'] == null ? undefined : json['id'],
        'owned': json['owned'] == null ? undefined : json['owned'],
        'price': json['price'] == null ? undefined : json['price'],
        'vipOnly': json['vip_only'] == null ? undefined : json['vip_only'],
    };
}

export function DecorationFrameDTOToJSON(json: any): DecorationFrameDTO {
    return DecorationFrameDTOToJSONTyped(json, false);
}

export function DecorationFrameDTOToJSONTyped(value?: DecorationFrameDTO | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'current': value['current'],
        'frame': value['frame'],
        'frame_thumbnail': value['frameThumbnail'],
        'frame_type': value['frameType'],
        'id': value['id'],
        'owned': value['owned'],
        'price': value['price'],
        'vip_only': value['vipOnly'],
    };
}

